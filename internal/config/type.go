package config

type SyncConfig struct {
	Name              string  `yaml:"name"`               // unique name of the sync task
	Source            string  `yaml:"source"`             // path to the local storage location
	SourcePrefix      string  `yaml:"source_prefix"`      // optional prefix in the source location
	Destination       string  `yaml:"destination"`        // path to the remote storage location
	DestinationPrefix string  `yaml:"destination_prefix"` // optional prefix in the destination location
	Trigger           string  `yaml:"trigger"`            // options: manual, automatic; determines when the sync should occur
	Schedule          *string `yaml:"schedule"`           // cron expression for automatic sync; ignored if trigger is manual
}
