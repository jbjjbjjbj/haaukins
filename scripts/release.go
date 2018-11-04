package main

import (
	"fmt"
	"github.com/aau-network-security/go-ntp/scripts/git"
	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/semver-bump/bump"
	"github.com/giantswarm/semver-bump/storage"
	"github.com/spf13/cobra"
	"os"
)

var (
	versionFile = "VERSION"
)

func newSemverBumper() (*bump.SemverBumper, error) {
	vs, err := storage.NewVersionStorage("file", "")
	if err != nil {
		return nil, err
	}
	return bump.NewSemverBumper(vs, versionFile), nil
}

func patch() *cobra.Command {
	return &cobra.Command{
		Use:   "patch",
		Short: "Release a patch version",
		Run: func(cmd *cobra.Command, args []string) {
			sb, err := newSemverBumper()
			if err != nil {
				fmt.Printf("Failed to create semver bumper: %s", err)
				return
			}
			curVer, err := sb.GetCurrentVersion()
			if err != nil {
				fmt.Printf("Failed to get current version: %s", err)
				return
			}
			newVer, err := sb.BumpPatchVersion("", "")
			if err != nil {
				fmt.Printf("Failed to bump version: %s", err)
				return
			}
			fmt.Printf("Releasing version %s (from %s)\n", curVer.String(), newVer.String())

			repo, err := git.NewRepo(".")
			if err := repo.Commit(newVer, versionFile); err != nil {
				fmt.Printf("Failed to commit version: %s", err)
				return
			}

			if err := repo.Tag(newVer); err != nil {
				fmt.Printf("Failed to create tag: %s", err)
				return
			}

			newBranchVer, err := semver.NewVersion(newVer.String())
			if err != nil {
				fmt.Printf("Failed to copy version: %s", err)
				return
			}

			newBranchVer.BumpPatch()
			if err := repo.CreateBranch(newBranchVer); err != nil {
				fmt.Printf("Failed to create branch: %s", err)
				return
			}
		},
	}
}

func main() {
	var rootCmd = &cobra.Command{Use: "release"}
	rootCmd.AddCommand(
		patch(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
