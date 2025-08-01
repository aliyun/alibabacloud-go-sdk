// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdatePluginConfigShrinkRequest
	GetAcceptLanguage() *string
	SetConfig(v string) *UpdatePluginConfigShrinkRequest
	GetConfig() *string
	SetConfigLevel(v int32) *UpdatePluginConfigShrinkRequest
	GetConfigLevel() *int32
	SetEnable(v bool) *UpdatePluginConfigShrinkRequest
	GetEnable() *bool
	SetGatewayId(v int64) *UpdatePluginConfigShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdatePluginConfigShrinkRequest
	GetGatewayUniqueId() *string
	SetGmtCreate(v string) *UpdatePluginConfigShrinkRequest
	GetGmtCreate() *string
	SetGmtModified(v string) *UpdatePluginConfigShrinkRequest
	GetGmtModified() *string
	SetId(v int64) *UpdatePluginConfigShrinkRequest
	GetId() *int64
	SetPluginId(v int64) *UpdatePluginConfigShrinkRequest
	GetPluginId() *int64
	SetResourceIdListShrink(v string) *UpdatePluginConfigShrinkRequest
	GetResourceIdListShrink() *string
}

type UpdatePluginConfigShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// zh: Chinese en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	Config         *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The application scope of the plug-in.
	//
	// 	- 0: global
	//
	// 	- 1: route
	//
	// 	- 2: domain name
	//
	// example:
	//
	// 0
	ConfigLevel *int32 `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	// Specifies whether to enable the plug-in.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Deprecated
	//
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-ubuwqygbq4783gqb2y3f87q****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Deprecated
	//
	// The creation time.
	//
	// example:
	//
	// 1667309705000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Deprecated
	//
	// The update time.
	//
	// example:
	//
	// 1667309705000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the plug-in configuration.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the gateway plug-in.
	//
	// example:
	//
	// 2
	PluginId             *int64  `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	ResourceIdListShrink *string `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty"`
}

func (s UpdatePluginConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePluginConfigShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdatePluginConfigShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdatePluginConfigShrinkRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *UpdatePluginConfigShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdatePluginConfigShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdatePluginConfigShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdatePluginConfigShrinkRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdatePluginConfigShrinkRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdatePluginConfigShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePluginConfigShrinkRequest) GetPluginId() *int64 {
	return s.PluginId
}

func (s *UpdatePluginConfigShrinkRequest) GetResourceIdListShrink() *string {
	return s.ResourceIdListShrink
}

func (s *UpdatePluginConfigShrinkRequest) SetAcceptLanguage(v string) *UpdatePluginConfigShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetConfig(v string) *UpdatePluginConfigShrinkRequest {
	s.Config = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetConfigLevel(v int32) *UpdatePluginConfigShrinkRequest {
	s.ConfigLevel = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetEnable(v bool) *UpdatePluginConfigShrinkRequest {
	s.Enable = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetGatewayId(v int64) *UpdatePluginConfigShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetGatewayUniqueId(v string) *UpdatePluginConfigShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetGmtCreate(v string) *UpdatePluginConfigShrinkRequest {
	s.GmtCreate = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetGmtModified(v string) *UpdatePluginConfigShrinkRequest {
	s.GmtModified = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetId(v int64) *UpdatePluginConfigShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetPluginId(v int64) *UpdatePluginConfigShrinkRequest {
	s.PluginId = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) SetResourceIdListShrink(v string) *UpdatePluginConfigShrinkRequest {
	s.ResourceIdListShrink = &v
	return s
}

func (s *UpdatePluginConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
