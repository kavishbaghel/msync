package config

import (
	"fmt"
	"strings"
)

func (c *SyncConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("config cannot be nil")
	}
	if strings.TrimSpace(c.Name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if strings.TrimSpace(c.Source) == "" {
		return fmt.Errorf("source cannot be empty")
	}
	if strings.TrimSpace(c.Destination) == "" {
		return fmt.Errorf("destination cannot be empty")
	}
	if c.Trigger != "manual" && c.Trigger != "automatic" {
		return fmt.Errorf("trigger must be either 'manual' or 'automatic'")
	}
	if c.Trigger == "automatic" && (c.Schedule == nil || strings.TrimSpace(*c.Schedule) == "") {
		return fmt.Errorf("schedule cannot be empty when trigger is automatic")
	}
	return nil
}
