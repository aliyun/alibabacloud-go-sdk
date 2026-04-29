// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePolarClawChannelShrinkRequest
	GetApplicationId() *string
	SetChannelConfigShrink(v string) *CreatePolarClawChannelShrinkRequest
	GetChannelConfigShrink() *string
	SetChannelId(v string) *CreatePolarClawChannelShrinkRequest
	GetChannelId() *string
	SetNpmPackage(v string) *CreatePolarClawChannelShrinkRequest
	GetNpmPackage() *string
	SetPluginId(v string) *CreatePolarClawChannelShrinkRequest
	GetPluginId() *string
	SetRestart(v bool) *CreatePolarClawChannelShrinkRequest
	GetRestart() *bool
}

type CreatePolarClawChannelShrinkRequest struct {
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
	//         "*"
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
	// This parameter is required.
	//
	// example:
	//
	// @larksuite/openclaw-lark@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s CreatePolarClawChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawChannelShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawChannelShrinkRequest) GetChannelConfigShrink() *string {
	return s.ChannelConfigShrink
}

func (s *CreatePolarClawChannelShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreatePolarClawChannelShrinkRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *CreatePolarClawChannelShrinkRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *CreatePolarClawChannelShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *CreatePolarClawChannelShrinkRequest) SetApplicationId(v string) *CreatePolarClawChannelShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) SetChannelConfigShrink(v string) *CreatePolarClawChannelShrinkRequest {
	s.ChannelConfigShrink = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) SetChannelId(v string) *CreatePolarClawChannelShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) SetNpmPackage(v string) *CreatePolarClawChannelShrinkRequest {
	s.NpmPackage = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) SetPluginId(v string) *CreatePolarClawChannelShrinkRequest {
	s.PluginId = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) SetRestart(v bool) *CreatePolarClawChannelShrinkRequest {
	s.Restart = &v
	return s
}

func (s *CreatePolarClawChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
