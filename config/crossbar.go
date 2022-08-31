package config

type CrossbarConfigAllow struct {
	Call      bool `json:"call"`
	Register  bool `json:"register"`
	Publish   bool `json:"publish"`
	Subscribe bool `json:"subscribe"`
}

type CrossbarConfigDisclose struct {
	Caller    bool `json:"caller"`
	Publisher bool `json:"publisher"`
}

type CrossbarConfigPermission struct {
	Uri      string                 `json:"uri"`
	Match    string                 `json:"match"`
	Allow    CrossbarConfigAllow    `json:"allow"`
	Disclose CrossbarConfigDisclose `json:"disclose"`
	Cache    bool                   `json:"cache"`
}

type CrossbarConfigRole struct {
	Name        string                     `json:"name"`
	Permissions []CrossbarConfigPermission `json:"permissions"`
}

type CrossbarConfigRealm struct {
	Name  string               `json:"name"`
	Roles []CrossbarConfigRole `json:"roles"`
}

type CrossbarConfigEndpoint struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type CrossbarConfigTransport struct {
	Type     string
	Endpoint string
}

type CrossbarConfigWorker struct {
	Type       string                    `json:"type"`
	Realms     []CrossbarConfigRealm     `json:"realms"`
	Transports []CrossbarConfigTransport `json:"transports"`
}

type CrossbarConfig struct {
	Version    int                    `json:"version"`
	Controller *interface{}           `json:"-"`
	Workers    []CrossbarConfigWorker `json:"workers"`
}
