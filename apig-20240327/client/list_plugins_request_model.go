// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *ListPluginsRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *ListPluginsRequest
	GetAttachResourceType() *string
	SetGatewayId(v string) *ListPluginsRequest
	GetGatewayId() *string
	SetGatewayType(v string) *ListPluginsRequest
	GetGatewayType() *string
	SetIncludeBuiltinAiGateway(v bool) *ListPluginsRequest
	GetIncludeBuiltinAiGateway() *bool
	SetPageNumber(v int32) *ListPluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPluginsRequest
	GetPageSize() *int32
	SetPluginClassId(v string) *ListPluginsRequest
	GetPluginClassId() *string
	SetPluginClassName(v string) *ListPluginsRequest
	GetPluginClassName() *string
	SetWithAttachmentInfo(v bool) *ListPluginsRequest
	GetWithAttachmentInfo() *bool
}

type ListPluginsRequest struct {
	// example:
	//
	// api-cuip2pum1hksng6oni3g
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// example:
	//
	// gw-csrhgn6m1hkt65qbxxgg
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// AI
	GatewayType             *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	IncludeBuiltinAiGateway *bool   `json:"includeBuiltinAiGateway,omitempty" xml:"includeBuiltinAiGateway,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// pls-dn82a9djd8z****
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// example:
	//
	// key-auth
	PluginClassName *string `json:"pluginClassName,omitempty" xml:"pluginClassName,omitempty"`
	// example:
	//
	// false
	WithAttachmentInfo *bool `json:"withAttachmentInfo,omitempty" xml:"withAttachmentInfo,omitempty"`
}

func (s ListPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListPluginsRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *ListPluginsRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPluginsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPluginsRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListPluginsRequest) GetIncludeBuiltinAiGateway() *bool {
	return s.IncludeBuiltinAiGateway
}

func (s *ListPluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginsRequest) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *ListPluginsRequest) GetPluginClassName() *string {
	return s.PluginClassName
}

func (s *ListPluginsRequest) GetWithAttachmentInfo() *bool {
	return s.WithAttachmentInfo
}

func (s *ListPluginsRequest) SetAttachResourceId(v string) *ListPluginsRequest {
	s.AttachResourceId = &v
	return s
}

func (s *ListPluginsRequest) SetAttachResourceType(v string) *ListPluginsRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPluginsRequest) SetGatewayId(v string) *ListPluginsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPluginsRequest) SetGatewayType(v string) *ListPluginsRequest {
	s.GatewayType = &v
	return s
}

func (s *ListPluginsRequest) SetIncludeBuiltinAiGateway(v bool) *ListPluginsRequest {
	s.IncludeBuiltinAiGateway = &v
	return s
}

func (s *ListPluginsRequest) SetPageNumber(v int32) *ListPluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPluginsRequest) SetPageSize(v int32) *ListPluginsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPluginsRequest) SetPluginClassId(v string) *ListPluginsRequest {
	s.PluginClassId = &v
	return s
}

func (s *ListPluginsRequest) SetPluginClassName(v string) *ListPluginsRequest {
	s.PluginClassName = &v
	return s
}

func (s *ListPluginsRequest) SetWithAttachmentInfo(v bool) *ListPluginsRequest {
	s.WithAttachmentInfo = &v
	return s
}

func (s *ListPluginsRequest) Validate() error {
	return dara.Validate(s)
}
