package types

type DockerConfig struct {
	Endpoint string `yaml:"endpoint"`
}

type ETCDConfig struct {
	Prefix   string   `yaml:"prefix"`
	Machines []string `yaml:"etcd"`
}

type MetricsConfig struct {
	Step      int64    `yaml:"step"`
	Transfers []string `yaml:"transfers"`
}

type APIConfig struct {
	Addr string `yaml:"addr"`
}

type LogConfig struct {
	Forwards []string `yaml:"forwards"`
	Stdout   bool     `yaml:"stdout"`
}

type Config struct {
	PidFile             string `yaml:"pid"`
	HostName            string
	Zone                string `yaml:"zone"`
	HealthCheckInterval int    `yaml:"health_check_interval"`
	HealthCheckTimeout  int    `yaml:"health_check_timeout"`

	Docker  DockerConfig
	Etcd    ETCDConfig
	Metrics MetricsConfig
	API     APIConfig
	Log     LogConfig
}
