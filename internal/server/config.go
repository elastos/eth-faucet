package server

type Config struct {
	provider        string
	network         string
	symbol          string
	httpPort        int
	interval        int
	payout          float64
	proxyCount      int
	hcaptchaSiteKey string
	hcaptchaSecret  string
}

func NewConfig(provider, network, symbol string, httpPort, interval, proxyCount int, payout float64, hcaptchaSiteKey, hcaptchaSecret string) *Config {
	return &Config{
		provider:        provider,
		network:         network,
		symbol:          symbol,
		httpPort:        httpPort,
		interval:        interval,
		payout:          payout,
		proxyCount:      proxyCount,
		hcaptchaSiteKey: hcaptchaSiteKey,
		hcaptchaSecret:  hcaptchaSecret,
	}
}
