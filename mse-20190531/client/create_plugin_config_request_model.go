// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreatePluginConfigRequest
	GetAcceptLanguage() *string
	SetConfig(v string) *CreatePluginConfigRequest
	GetConfig() *string
	SetConfigLevel(v int32) *CreatePluginConfigRequest
	GetConfigLevel() *int32
	SetEnable(v bool) *CreatePluginConfigRequest
	GetEnable() *bool
	SetGatewayUniqueId(v string) *CreatePluginConfigRequest
	GetGatewayUniqueId() *string
	SetPluginId(v int64) *CreatePluginConfigRequest
	GetPluginId() *int64
	SetResourceIdList(v []*int64) *CreatePluginConfigRequest
	GetResourceIdList() []*int64
}

type CreatePluginConfigRequest struct {
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
	ResourceIdList []*int64 `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty" type:"Repeated"`
}

func (s CreatePluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginConfigRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreatePluginConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *CreatePluginConfigRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *CreatePluginConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreatePluginConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *CreatePluginConfigRequest) GetPluginId() *int64 {
	return s.PluginId
}

func (s *CreatePluginConfigRequest) GetResourceIdList() []*int64 {
	return s.ResourceIdList
}

func (s *CreatePluginConfigRequest) SetAcceptLanguage(v string) *CreatePluginConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreatePluginConfigRequest) SetConfig(v string) *CreatePluginConfigRequest {
	s.Config = &v
	return s
}

func (s *CreatePluginConfigRequest) SetConfigLevel(v int32) *CreatePluginConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *CreatePluginConfigRequest) SetEnable(v bool) *CreatePluginConfigRequest {
	s.Enable = &v
	return s
}

func (s *CreatePluginConfigRequest) SetGatewayUniqueId(v string) *CreatePluginConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *CreatePluginConfigRequest) SetPluginId(v int64) *CreatePluginConfigRequest {
	s.PluginId = &v
	return s
}

func (s *CreatePluginConfigRequest) SetResourceIdList(v []*int64) *CreatePluginConfigRequest {
	s.ResourceIdList = v
	return s
}

func (s *CreatePluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
