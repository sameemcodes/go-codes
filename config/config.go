package config

var Config *Configuration

type Configuration struct {
	ServerAddr string

	KafkaConfig     KafkaConfig
	NewRelicConfig  NewRelicConfig
	RedisConfig     RedisConfig
	LoggerConfig    LoggerConfig
	TimeZoneConfig  TimeZoneConfig
	ConstantsConfig ConstantsConfig
}

type LoggerConfig struct {
	OutputPaths      []string
	ErrorOutputPaths []string
}

type NewRelicConfig struct {
	AppName    string
	LicenseKey string
	Enabled    bool
}

type KafkaConfig struct {
	BrokerList            string
	ConsumerGrp           string
	OffsetReset           string
	FeedTopic             KafkaTopicConfig
	IndexTopic            KafkaTopicConfig
	BookTopic             KafkaTopicConfig
	BhavCopyTopic         KafkaTopicConfig
	IndexEodTopic         KafkaTopicConfig
	ResetPriceTopic       KafkaTopicConfig
	ResetEquityPriceTopic KafkaTopicConfig
	IsEnabled             bool
}

type KafkaTopicConfig struct {
	Name        string
	Concurrency int
}

type RedisConfig struct {
	Addr         string
	Password     string
	PoolSize     int
	MinIdleConns int
	DB           int
}

type ConstantsConfig struct {
	PricePrefix string
	BookPrefix  string
	IndexPrefix string
}

type TimeZoneConfig struct {
	TimeZone string
}
