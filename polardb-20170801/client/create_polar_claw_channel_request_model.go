// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePolarClawChannelRequest
	GetApplicationId() *string
	SetChannelConfig(v map[string]interface{}) *CreatePolarClawChannelRequest
	GetChannelConfig() map[string]interface{}
	SetChannelId(v string) *CreatePolarClawChannelRequest
	GetChannelId() *string
	SetNpmPackage(v string) *CreatePolarClawChannelRequest
	GetNpmPackage() *string
	SetPluginId(v string) *CreatePolarClawChannelRequest
	GetPluginId() *string
	SetRestart(v bool) *CreatePolarClawChannelRequest
	GetRestart() *bool
}

type CreatePolarClawChannelRequest struct {
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
	ChannelConfig map[string]interface{} `json:"ChannelConfig,omitempty" xml:"ChannelConfig,omitempty"`
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

func (s CreatePolarClawChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawChannelRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawChannelRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawChannelRequest) GetChannelConfig() map[string]interface{} {
	return s.ChannelConfig
}

func (s *CreatePolarClawChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreatePolarClawChannelRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *CreatePolarClawChannelRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *CreatePolarClawChannelRequest) GetRestart() *bool {
	return s.Restart
}

func (s *CreatePolarClawChannelRequest) SetApplicationId(v string) *CreatePolarClawChannelRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawChannelRequest) SetChannelConfig(v map[string]interface{}) *CreatePolarClawChannelRequest {
	s.ChannelConfig = v
	return s
}

func (s *CreatePolarClawChannelRequest) SetChannelId(v string) *CreatePolarClawChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *CreatePolarClawChannelRequest) SetNpmPackage(v string) *CreatePolarClawChannelRequest {
	s.NpmPackage = &v
	return s
}

func (s *CreatePolarClawChannelRequest) SetPluginId(v string) *CreatePolarClawChannelRequest {
	s.PluginId = &v
	return s
}

func (s *CreatePolarClawChannelRequest) SetRestart(v bool) *CreatePolarClawChannelRequest {
	s.Restart = &v
	return s
}

func (s *CreatePolarClawChannelRequest) Validate() error {
	return dara.Validate(s)
}
