package config

import "testing"

func strPtr(s string) *string {
	return &s
}

func TestSyncConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     SyncConfig
		wantErr bool
	}{
		{
			name: "valid manual config with nil schedule",
			cfg: SyncConfig{
				Name:        "example-sync",
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "manual",
				Schedule:    nil,
			},
			wantErr: false,
		},
		{
			name: "valid automatic config with schedule",
			cfg: SyncConfig{
				Name:        "example-sync",
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "automatic",
				Schedule:    strPtr("0 * * * *"),
			},
			wantErr: false,
		},
		{
			name: "missing name",
			cfg: SyncConfig{
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "manual",
			},
			wantErr: true,
		},
		{
			name: "missing source",
			cfg: SyncConfig{
				Name:        "example-sync",
				Destination: "s3://bucket-name",
				Trigger:     "manual",
			},
			wantErr: true,
		},
		{
			name: "missing destination",
			cfg: SyncConfig{
				Name:    "example-sync",
				Source:  "/Users/user1/data",
				Trigger: "manual",
			},
			wantErr: true,
		},
		{
			name: "invalid trigger",
			cfg: SyncConfig{
				Name:        "example-sync",
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "on-start",
			},
			wantErr: true,
		},
		{
			name: "automatic trigger with nil schedule",
			cfg: SyncConfig{
				Name:        "example-sync",
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "automatic",
				Schedule:    nil,
			},
			wantErr: true,
		},
		{
			name: "automatic trigger with empty schedule",
			cfg: SyncConfig{
				Name:        "example-sync",
				Source:      "/Users/user1/data",
				Destination: "s3://bucket-name",
				Trigger:     "automatic",
				Schedule:    strPtr(""),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr && err == nil {
				t.Fatalf("expected an error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error but got %v", err)
			}
		})
	}
}
