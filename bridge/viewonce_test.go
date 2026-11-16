package main

import (
	"testing"

	"github.com/polymorfa/hypermeow/proto/waE2E"
	"github.com/polymorfa/hypermeow/types/events"
)

func TestDescribeEventDoesNotPersistViewOnceMedia(t *testing.T) {
	for name, test := range map[string]struct {
		msg      *waE2E.Message
		wantKind string
	}{
		"photo": {msg: &waE2E.Message{ImageMessage: &waE2E.ImageMessage{Caption: stringPointer("private caption")}}, wantKind: "view-once-photo"},
		"video": {msg: &waE2E.Message{VideoMessage: &waE2E.VideoMessage{Caption: stringPointer("private caption")}}, wantKind: "view-once-video"},
	} {
		kind, text, raw := describeEvent(&events.Message{Message: test.msg, IsViewOnce: true})
		if kind != test.wantKind {
			t.Errorf("%s kind = %q, want %q", name, kind, test.wantKind)
		}
		if text != "" {
			t.Errorf("%s caption leaked into stored text: %q", name, text)
		}
		if len(raw) != 0 {
			t.Errorf("%s payload should not be persisted, got %d bytes", name, len(raw))
		}
	}
}

func TestDescribeEventKeepsOrdinaryImageMedia(t *testing.T) {
	kind, text, raw := describeEvent(&events.Message{Message: &waE2E.Message{
		ImageMessage: &waE2E.ImageMessage{Caption: stringPointer("ordinary photo")},
	}})
	if kind != "image" || text != "ordinary photo" || len(raw) == 0 {
		t.Fatalf("ordinary image = kind %q, text %q, %d raw bytes", kind, text, len(raw))
	}
}

func stringPointer(value string) *string { return &value }
