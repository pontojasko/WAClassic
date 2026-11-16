package main

import "testing"

func TestChatDisplayNamePrefersCurrentContactForIndividualChat(t *testing.T) {
	got := chatDisplayName("+55•••••81", "Saved Contact", "+5581999999981", false)
	if got != "Saved Contact" {
		t.Fatalf("chatDisplayName() = %q, want current contact name", got)
	}
}

func TestChatDisplayNameKeepsGroupName(t *testing.T) {
	got := chatDisplayName("Family Group", "", "", true)
	if got != "Family Group" {
		t.Fatalf("chatDisplayName() = %q, want stored group name", got)
	}
}

func TestChatDisplayNameFallsBackWhenContactIsUnavailable(t *testing.T) {
	if got := chatDisplayName("Old Name", "", "Phone", false); got != "Old Name" {
		t.Fatalf("chatDisplayName() = %q, want stored name", got)
	}
	if got := chatDisplayName("", "", "Phone", false); got != "Phone" {
		t.Fatalf("chatDisplayName() = %q, want fallback", got)
	}
}
