// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeletePolarClawChannelRequest
	GetApplicationId() *string
	SetChannelId(v string) *DeletePolarClawChannelRequest
	GetChannelId() *string
	SetPluginId(v string) *DeletePolarClawChannelRequest
	GetPluginId() *string
	SetRestart(v bool) *DeletePolarClawChannelRequest
	GetRestart() *bool
	SetUninstallPlugin(v bool) *DeletePolarClawChannelRequest
	GetUninstallPlugin() *bool
}

type DeletePolarClawChannelRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// feishu
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The plugin ID. This parameter is required if `UninstallPlugin` is set to `true`.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// Specifies whether to restart the gateway after the channel is deleted. Default value: `true`.
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// Specifies whether to uninstall the channel plugin. Default value: `false`.
	//
	// example:
	//
	// true
	UninstallPlugin *bool `json:"UninstallPlugin,omitempty" xml:"UninstallPlugin,omitempty"`
}

func (s DeletePolarClawChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawChannelRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarClawChannelRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeletePolarClawChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DeletePolarClawChannelRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DeletePolarClawChannelRequest) GetRestart() *bool {
	return s.Restart
}

func (s *DeletePolarClawChannelRequest) GetUninstallPlugin() *bool {
	return s.UninstallPlugin
}

func (s *DeletePolarClawChannelRequest) SetApplicationId(v string) *DeletePolarClawChannelRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeletePolarClawChannelRequest) SetChannelId(v string) *DeletePolarClawChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DeletePolarClawChannelRequest) SetPluginId(v string) *DeletePolarClawChannelRequest {
	s.PluginId = &v
	return s
}

func (s *DeletePolarClawChannelRequest) SetRestart(v bool) *DeletePolarClawChannelRequest {
	s.Restart = &v
	return s
}

func (s *DeletePolarClawChannelRequest) SetUninstallPlugin(v bool) *DeletePolarClawChannelRequest {
	s.UninstallPlugin = &v
	return s
}

func (s *DeletePolarClawChannelRequest) Validate() error {
	return dara.Validate(s)
}
