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
	// The language in which the response is displayed. Valid values:
	//
	// - zh: Chinese. This is the default value.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The plug-in configuration. The configuration content of a Wasm plug-in is in YAML format, and the configuration content of a Lua plug-in is Lua code.
	//
	// example:
	//
	// name: John
	//
	// age: 18
	//
	// friends:
	//
	// - David
	//
	// - Anne
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The application scope of the plug-in. Valid values:
	//
	// - 0: Global.
	//
	// - 1: Route.
	//
	// - 2: Domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// Specifies whether to enable the plug-in.
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
	// The ID of the gateway plug-in.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PluginId *int64 `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The list of domain name IDs or route IDs. The specific type depends on the value of ConfigLevel.
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
