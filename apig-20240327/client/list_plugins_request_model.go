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
	// The resource attachment ID.
	//
	// example:
	//
	// api-cuip2pum1hksng6oni3g
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// The resource attachment type.
	//
	// - HttpApi: HttpApi.
	//
	// - Operation: Operation of HttpApi.
	//
	// - GatewayRoute: Gateway route.
	//
	// - GatewayService: Gateway service.
	//
	// - GatewayServicePort: Gateway service port.
	//
	// - Domain: Gateway domain.
	//
	// - Gateway: Gateway.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The gateway instance ID for filtering.
	//
	// example:
	//
	// gw-csrhgn6m1hkt65qbxxgg
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance type. Valid values: **AI*	- and **API**.
	//
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// Specifies whether to include built-in AI plug-ins in the returned results. Default: false.
	//
	// example:
	//
	// false
	IncludeBuiltinAiGateway *bool `json:"includeBuiltinAiGateway,omitempty" xml:"includeBuiltinAiGateway,omitempty"`
	// The page number to return. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The plug-in type ID for filtering.
	//
	// example:
	//
	// pls-dn82a9djd8z****
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	// The plug-in type name for filtering.
	//
	// example:
	//
	// key-auth
	PluginClassName *string `json:"pluginClassName,omitempty" xml:"pluginClassName,omitempty"`
	// Specifies whether the returned results should include plug-in attachment information corresponding to the attachResourceId.
	//
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
