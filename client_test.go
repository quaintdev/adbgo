package adbgo

import "testing"

func TestADBBridge_SendText(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		addr string
		// Named input parameters for target function.
		text    string
		wantErr bool
	}{
		{
			name:    "keyboard input",
			addr:    "192.168.1.104",
			text:    "test message",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewADBBridge(tt.addr)
			gotErr := a.SendText(tt.text)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SendText() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SendText() succeeded unexpectedly")
			}
		})
	}
}
