package event

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf, _ := loadConfig("test_resources/test_event.yml")
	if conf.CTFd.Name != "Test CTFd" {
		t.Fatalf("Unexpected CTFd name (expected 'Test CTFd'): %s", conf.CTFd.Name)
	}
	if conf.Guac.Host != "localhost" {
		t.Fatalf("Unexpected Guacamole host (expected 'localhost'): %s", conf.Guac.Host)
	}
	if conf.RevProxy.Host != "127.0.0.1" {
		t.Fatalf("Unexpected reverse proxy host (expected '127.0.0.1'): %s", conf.RevProxy.Host)
	}
}
