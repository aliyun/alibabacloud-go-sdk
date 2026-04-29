// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawChannelShrinkRequest
	GetApplicationId() *string
	SetChannelConfigShrink(v string) *UpgradePolarClawChannelShrinkRequest
	GetChannelConfigShrink() *string
	SetChannelId(v string) *UpgradePolarClawChannelShrinkRequest
	GetChannelId() *string
	SetNpmPackage(v string) *UpgradePolarClawChannelShrinkRequest
	GetNpmPackage() *string
	SetPluginId(v string) *UpgradePolarClawChannelShrinkRequest
	GetPluginId() *string
	SetRestart(v bool) *UpgradePolarClawChannelShrinkRequest
	GetRestart() *bool
}

type UpgradePolarClawChannelShrinkRequest struct {
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
	ChannelConfigShrink *string `json:"ChannelConfig,omitempty" xml:"ChannelConfig,omitempty"`
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

func (s UpgradePolarClawChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawChannelShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawChannelShrinkRequest) GetChannelConfigShrink() *string {
	return s.ChannelConfigShrink
}

func (s *UpgradePolarClawChannelShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpgradePolarClawChannelShrinkRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *UpgradePolarClawChannelShrinkRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UpgradePolarClawChannelShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpgradePolarClawChannelShrinkRequest) SetApplicationId(v string) *UpgradePolarClawChannelShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) SetChannelConfigShrink(v string) *UpgradePolarClawChannelShrinkRequest {
	s.ChannelConfigShrink = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) SetChannelId(v string) *UpgradePolarClawChannelShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) SetNpmPackage(v string) *UpgradePolarClawChannelShrinkRequest {
	s.NpmPackage = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) SetPluginId(v string) *UpgradePolarClawChannelShrinkRequest {
	s.PluginId = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) SetRestart(v bool) *UpgradePolarClawChannelShrinkRequest {
	s.Restart = &v
	return s
}

func (s *UpgradePolarClawChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
