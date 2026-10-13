package main

import "testing"

func TestIsAIAssistant(t *testing.T) {
	for name, want := range map[string]bool{
		"ИИ-помощник":  true,
		"HR бот":       true,
		"Робот Анна":   true,
		"Анна Иванова": false,
		"Работа":       false,
	} {
		if got := isAIAssistant(name); got != want {
			t.Errorf("isAIAssistant(%q) = %v, want %v", name, got, want)
		}
	}
}
