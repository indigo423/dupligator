package config

type LocalConfig struct {
	Interface  string `yaml:"interface"`
	Address    string `yaml:"address"`
	GatewayMAC string `yaml:"gateway_mac"`
	Port       int    `yaml:"port"`
}

type SourceConfig struct {
	Name                string `yaml:"name"`
	SourceIP            string `yaml:"source_ip"`
	StickyBytesStart    int    `yaml:"sticky_bytes_start"`
	StickyBytesLength   int    `yaml:"sticky_bytes_length"`
	StickyBytesEnd      int
	DefaultCatchallOnly bool     `yaml:"default_catchall_only"`
	Receivers           []string `yaml:"receivers"`
}

type ReceiverConfig struct {
	Name            string `yaml:"name"`
	Spoof           bool   `yaml:"spoof"`
	PreserveUdpPort bool   `yaml:"preserve_udp_port"`
	Proto           string `yaml:"proto"`
	Ip              string `yaml:"ip"`
	IPvPref         int    `yaml:"ipv_pref"`
	Port            int    `yaml:"port"`
}

type Config struct {
	LocalV4Config LocalConfig      `yaml:"local_v4"`
	LocalV6Config LocalConfig      `yaml:"local_v6"`
	Sources       []SourceConfig   `yaml:"sources"`
	Receivers     []ReceiverConfig `yaml:"receivers"`
}
