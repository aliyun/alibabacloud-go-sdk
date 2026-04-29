// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawChannelRequest
	GetApplicationId() *string
	SetChannelConfig(v map[string]interface{}) *UpgradePolarClawChannelRequest
	GetChannelConfig() map[string]interface{}
	SetChannelId(v string) *UpgradePolarClawChannelRequest
	GetChannelId() *string
	SetNpmPackage(v string) *UpgradePolarClawChannelRequest
	GetNpmPackage() *string
	SetPluginId(v string) *UpgradePolarClawChannelRequest
	GetPluginId() *string
	SetRestart(v bool) *UpgradePolarClawChannelRequest
	GetRestart() *bool
}

type UpgradePolarClawChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// {
	//
	//     "enabled": true,
	//
	//     "dmPolicy": "open",
	//
	//     "allowFrom": [
	//
	//         "*",
	//
	//         "ou_abc"
	//
	//     ]
	//
	// }
	ChannelConfig map[string]interface{} `json:"ChannelConfig,omitempty" xml:"ChannelConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// @larksuite/openclaw-feishu@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// example:
	//
	// openclaw-feishu
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s UpgradePolarClawChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawChannelRequest) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawChannelRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawChannelRequest) GetChannelConfig() map[string]interface{} {
	return s.ChannelConfig
}

func (s *UpgradePolarClawChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpgradePolarClawChannelRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *UpgradePolarClawChannelRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UpgradePolarClawChannelRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpgradePolarClawChannelRequest) SetApplicationId(v string) *UpgradePolarClawChannelRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawChannelRequest) SetChannelConfig(v map[string]interface{}) *UpgradePolarClawChannelRequest {
	s.ChannelConfig = v
	return s
}

func (s *UpgradePolarClawChannelRequest) SetChannelId(v string) *UpgradePolarClawChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *UpgradePolarClawChannelRequest) SetNpmPackage(v string) *UpgradePolarClawChannelRequest {
	s.NpmPackage = &v
	return s
}

func (s *UpgradePolarClawChannelRequest) SetPluginId(v string) *UpgradePolarClawChannelRequest {
	s.PluginId = &v
	return s
}

func (s *UpgradePolarClawChannelRequest) SetRestart(v bool) *UpgradePolarClawChannelRequest {
	s.Restart = &v
	return s
}

func (s *UpgradePolarClawChannelRequest) Validate() error {
	return dara.Validate(s)
}
