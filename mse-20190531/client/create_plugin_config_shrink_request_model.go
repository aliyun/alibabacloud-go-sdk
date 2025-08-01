// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreatePluginConfigShrinkRequest
	GetAcceptLanguage() *string
	SetConfig(v string) *CreatePluginConfigShrinkRequest
	GetConfig() *string
	SetConfigLevel(v int32) *CreatePluginConfigShrinkRequest
	GetConfigLevel() *int32
	SetEnable(v bool) *CreatePluginConfigShrinkRequest
	GetEnable() *bool
	SetGatewayUniqueId(v string) *CreatePluginConfigShrinkRequest
	GetGatewayUniqueId() *string
	SetPluginId(v int64) *CreatePluginConfigShrinkRequest
	GetPluginId() *int64
	SetResourceIdListShrink(v string) *CreatePluginConfigShrinkRequest
	GetResourceIdListShrink() *string
}

type CreatePluginConfigShrinkRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	Config         *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The application scope of the plug-in. Valid values:
	//
	// 	- 0: global
	//
	// 	- 1: route
	//
	// 	- 2: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// Indicates whether the plug-in is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-ubuwqygbq4783gqb2y3f87q****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The gateway plug-in ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PluginId *int64 `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The domain IDs or route IDs. They are distinguished based on ConfigLevel.
	ResourceIdListShrink *string `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty"`
}

func (s CreatePluginConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginConfigShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreatePluginConfigShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *CreatePluginConfigShrinkRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *CreatePluginConfigShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreatePluginConfigShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreatePluginConfigShrinkRequest) GetPluginId() *int64 {
	return s.PluginId
}

func (s *CreatePluginConfigShrinkRequest) GetResourceIdListShrink() *string {
	return s.ResourceIdListShrink
}

func (s *CreatePluginConfigShrinkRequest) SetAcceptLanguage(v string) *CreatePluginConfigShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetConfig(v string) *CreatePluginConfigShrinkRequest {
	s.Config = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetConfigLevel(v int32) *CreatePluginConfigShrinkRequest {
	s.ConfigLevel = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetEnable(v bool) *CreatePluginConfigShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetGatewayUniqueId(v string) *CreatePluginConfigShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetPluginId(v int64) *CreatePluginConfigShrinkRequest {
	s.PluginId = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) SetResourceIdListShrink(v string) *CreatePluginConfigShrinkRequest {
	s.ResourceIdListShrink = &v
	return s
}

func (s *CreatePluginConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
