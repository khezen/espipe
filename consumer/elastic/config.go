package elastic

import "github.com/khezen/bulklog/auth"

// Config -
type Config struct {
	Enabled   bool              `yaml:"enabled"`
	Endpoint  string            `yaml:"endpoint"`
	Shards    int               `yaml:"shards"`
	AWSAuth   *auth.AWSConfig   `yaml:"aws_auth,omitempty"`
	BasicAuth *auth.BasicConfig `yaml:"basic_auth,omitempty"`
}
