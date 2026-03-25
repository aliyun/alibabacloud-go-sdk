// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasLike(v string) *ListPluginClassesRequest
	GetAliasLike() *string
	SetDirection(v string) *ListPluginClassesRequest
	GetDirection() *string
	SetExcludeBuiltinAiProxy(v bool) *ListPluginClassesRequest
	GetExcludeBuiltinAiProxy() *bool
	SetGatewayId(v string) *ListPluginClassesRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListPluginClassesRequest
	GetGatewayType() *string
	SetInstalled(v bool) *ListPluginClassesRequest
	GetInstalled() *bool
	SetNameLike(v string) *ListPluginClassesRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListPluginClassesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPluginClassesRequest
	GetPageSize() *int32
	SetSource(v string) *ListPluginClassesRequest
	GetSource() *string
	SetType(v string) *ListPluginClassesRequest
	GetType() *string
}

type ListPluginClassesRequest struct {
	// The alias keyword for a fuzzy search.
	AliasLike *string `json:"aliasLike,omitempty" xml:"aliasLike,omitempty"`
	// The traffic direction. Valid values:
	//
	// 	- OutBound
	//
	// 	- InBound
	//
	// 	- Both
	//
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Specifies whether to exclude built-in plugins.
	//
	// example:
	//
	// true
	ExcludeBuiltinAiProxy *bool `json:"excludeBuiltinAiProxy,omitempty" xml:"excludeBuiltinAiProxy,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-d1j8tjum1hkhxxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance type. Valid values: **AI*	- and **API**.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// Indicates whether the plugin is installed.
	//
	// example:
	//
	// false
	Installed *bool `json:"installed,omitempty" xml:"installed,omitempty"`
	// The plugin name for a fuzzy search.
	//
	// example:
	//
	// oauth
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The plugin source. Valid values:
	//
	// 	- HigressOfficial
	//
	// 	- HigressCommunity
	//
	// 	- Custom
	//
	// example:
	//
	// HigressOfficial
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The type of the plugin.
	//
	// example:
	//
	// Auth
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPluginClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginClassesRequest) GoString() string {
	return s.String()
}

func (s *ListPluginClassesRequest) GetAliasLike() *string {
	return s.AliasLike
}

func (s *ListPluginClassesRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListPluginClassesRequest) GetExcludeBuiltinAiProxy() *bool {
	return s.ExcludeBuiltinAiProxy
}

func (s *ListPluginClassesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPluginClassesRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListPluginClassesRequest) GetInstalled() *bool {
	return s.Installed
}

func (s *ListPluginClassesRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListPluginClassesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginClassesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginClassesRequest) GetSource() *string {
	return s.Source
}

func (s *ListPluginClassesRequest) GetType() *string {
	return s.Type
}

func (s *ListPluginClassesRequest) SetAliasLike(v string) *ListPluginClassesRequest {
	s.AliasLike = &v
	return s
}

func (s *ListPluginClassesRequest) SetDirection(v string) *ListPluginClassesRequest {
	s.Direction = &v
	return s
}

func (s *ListPluginClassesRequest) SetExcludeBuiltinAiProxy(v bool) *ListPluginClassesRequest {
	s.ExcludeBuiltinAiProxy = &v
	return s
}

func (s *ListPluginClassesRequest) SetGatewayId(v string) *ListPluginClassesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPluginClassesRequest) SetGatewayType(v string) *ListPluginClassesRequest {
	s.GatewayType = &v
	return s
}

func (s *ListPluginClassesRequest) SetInstalled(v bool) *ListPluginClassesRequest {
	s.Installed = &v
	return s
}

func (s *ListPluginClassesRequest) SetNameLike(v string) *ListPluginClassesRequest {
	s.NameLike = &v
	return s
}

func (s *ListPluginClassesRequest) SetPageNumber(v int32) *ListPluginClassesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPluginClassesRequest) SetPageSize(v int32) *ListPluginClassesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPluginClassesRequest) SetSource(v string) *ListPluginClassesRequest {
	s.Source = &v
	return s
}

func (s *ListPluginClassesRequest) SetType(v string) *ListPluginClassesRequest {
	s.Type = &v
	return s
}

func (s *ListPluginClassesRequest) Validate() error {
	return dara.Validate(s)
}
