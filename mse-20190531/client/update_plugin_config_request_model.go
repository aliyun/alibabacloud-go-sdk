// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdatePluginConfigRequest
	GetAcceptLanguage() *string
	SetConfig(v string) *UpdatePluginConfigRequest
	GetConfig() *string
	SetConfigLevel(v int32) *UpdatePluginConfigRequest
	GetConfigLevel() *int32
	SetEnable(v bool) *UpdatePluginConfigRequest
	GetEnable() *bool
	SetGatewayId(v int64) *UpdatePluginConfigRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdatePluginConfigRequest
	GetGatewayUniqueId() *string
	SetGmtCreate(v string) *UpdatePluginConfigRequest
	GetGmtCreate() *string
	SetGmtModified(v string) *UpdatePluginConfigRequest
	GetGmtModified() *string
	SetId(v int64) *UpdatePluginConfigRequest
	GetId() *int64
	SetPluginId(v int64) *UpdatePluginConfigRequest
	GetPluginId() *int64
	SetResourceIdList(v []*int64) *UpdatePluginConfigRequest
	GetResourceIdList() []*int64
}

type UpdatePluginConfigRequest struct {
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
	PluginId       *int64   `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	ResourceIdList []*int64 `json:"ResourceIdList,omitempty" xml:"ResourceIdList,omitempty" type:"Repeated"`
}

func (s UpdatePluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdatePluginConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdatePluginConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdatePluginConfigRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *UpdatePluginConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdatePluginConfigRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdatePluginConfigRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdatePluginConfigRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdatePluginConfigRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdatePluginConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePluginConfigRequest) GetPluginId() *int64 {
	return s.PluginId
}

func (s *UpdatePluginConfigRequest) GetResourceIdList() []*int64 {
	return s.ResourceIdList
}

func (s *UpdatePluginConfigRequest) SetAcceptLanguage(v string) *UpdatePluginConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetConfig(v string) *UpdatePluginConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetConfigLevel(v int32) *UpdatePluginConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetEnable(v bool) *UpdatePluginConfigRequest {
	s.Enable = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetGatewayId(v int64) *UpdatePluginConfigRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetGatewayUniqueId(v string) *UpdatePluginConfigRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetGmtCreate(v string) *UpdatePluginConfigRequest {
	s.GmtCreate = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetGmtModified(v string) *UpdatePluginConfigRequest {
	s.GmtModified = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetId(v int64) *UpdatePluginConfigRequest {
	s.Id = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetPluginId(v int64) *UpdatePluginConfigRequest {
	s.PluginId = &v
	return s
}

func (s *UpdatePluginConfigRequest) SetResourceIdList(v []*int64) *UpdatePluginConfigRequest {
	s.ResourceIdList = v
	return s
}

func (s *UpdatePluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
