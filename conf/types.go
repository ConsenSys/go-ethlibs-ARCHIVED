package conf

type ContextKeyMethod int

// HealthCheckEnvSpec represents the health check env variables struct.
type HealthCheckEnvSpec struct {
	ClientBlockLeeway     string `required:"true" envconfig:"CLIENT_BLOCK_LEEWAY"`
	ClientChainID         string `required:"true" envconfig:"CLIENT_CHAIN_ID"`
	ClientEndpoint        string `required:"true" envconfig:"CLIENT_ENDPOINT"`
	ExplorerEndpoint      string `required:"true" envconfig:"EXPLORER_ENDPOINT"`
	MinPeerCount		  string `required:"true" envconfig:"CLIENT_MIN_PEERCOUNT" default:"30"`
	ListenerPort          string `required:"false" envconfig:"LISTENER_PORT" default:"8080"`
	LogLevel              string `required:"false" envconfig:"LOG_LEVEL" default:"info"`
	MetricsEnabled        string `required:"false" envconfig:"METRICS_ENABLED" default:"false"`
	MetricsRecordInterval string `required:"false" envconfig:"METRICS_RECORD_INTERVAL" default:"10s"`
}

// MonitoringEnvSpec represents the monitoring env variables struct.
type MonitoringEnvSpec struct {
	Env                  	string `required:"true"`
	ClientEndpoint       	 string `required:"true" envconfig:"CLIENT_ENDPOINT"`
	ExplorerEndpoint      	string `required:"true" envconfig:"EXPLORER_ENDPOINT"`
	Network              	string `required:"true"`
	InfuraProjectID      	string `required:"true" split_words:"true"`
	InfuraProjectSecret  	string `required:"true" split_words:"true"`
	ClientChainID        	string `required:"true" envconfig:"CLIENT_CHAIN_ID"`
	DatadogAgentHost     	string `required:"true" envconfig:"DD_AGENT_HOST"`

	ClientBlockLeeway    	string `required:"false" envconfig:"CLIENT_BLOCK_LEEWAY" default:"10"`
	CheckInterval        	string `required:"false" envconfig:"CHECK_INTERVAL" default:"10s"`
	Internal             	string `required:"false" default:"false"`
	LogLevel             	string `required:"false" envconfig:"LOG_LEVEL" default:"info"`
	MinPeers              	string `required:"false" envconfig:"MIN_PEERS" default:"3"`
	LogBlockOffset		  	string `required:"false" envconfig:"LOG_BLOCK_OFFSET" default:"16"`
	EthCallContract       	string `required:"false" envconfig:"ETH_CALL_CONTRACT"`
	EthCallContractData   	string `required:"false" envconfig:"ETH_CALL_CONTRACT_DATA"`
	EthCallContractBlock  	string `required:"false" envconfig:"ETH_CALL_CONTRACT_BLOCK"`
	EthCallContractResult 	string `required:"false" envconfig:"ETH_CALL_CONTRACT_RESULT"`
}
