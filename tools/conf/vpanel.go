package conf

import (
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vpanel"
)

type VPanelConfig struct {
	Api          string              `json:"api"`
	Id           string              `json:"id"`
	Key          string              `json:"key"`
	Features     *FeaturesConfig     `json:"features"`
	Defaults     *VMessDefaultConfig `json:"default"`
	DetourConfig *VMessDetourConfig  `json:"detour"`
}

// Build implements Buildable
func (c *VPanelConfig) Build() (*serial.TypedMessage, error) {
	config := new(vpanel.Config)

	if c.Defaults != nil {
		config.Default = c.Defaults.Build()
	}

	if c.DetourConfig != nil {
		config.Detour = c.DetourConfig.Build()
	} else if c.Features != nil && c.Features.Detour != nil {
		config.Detour = c.Features.Detour.Build()
	}

	config.Api = c.Api
	config.Id = c.Id
	config.Key = c.Key

	return serial.ToTypedMessage(config), nil
}
