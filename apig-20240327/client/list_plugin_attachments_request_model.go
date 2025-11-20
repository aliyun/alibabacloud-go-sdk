// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *ListPluginAttachmentsRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *ListPluginAttachmentsRequest
	GetAttachResourceType() *string
	SetAttachResourceTypes(v string) *ListPluginAttachmentsRequest
	GetAttachResourceTypes() *string
	SetEnvironmentId(v string) *ListPluginAttachmentsRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *ListPluginAttachmentsRequest
	GetGatewayId() *string
	SetPageNumber(v int32) *ListPluginAttachmentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPluginAttachmentsRequest
	GetPageSize() *int32
	SetPluginId(v string) *ListPluginAttachmentsRequest
	GetPluginId() *string
	SetWithParentResource(v bool) *ListPluginAttachmentsRequest
	GetWithParentResource() *bool
}

type ListPluginAttachmentsRequest struct {
	// The resource attachment ID.
	//
	// example:
	//
	// hr-cv2h58em1hkg7c6vt43g
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// The resource attachment type (not yet in use).
	//
	// example:
	//
	// GatewayRoute
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The resource attachment types, separated by commas.
	//
	// example:
	//
	// GatewayRoute
	AttachResourceTypes *string `json:"attachResourceTypes,omitempty" xml:"attachResourceTypes,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-crlnqhtlhtgqflkqislg
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-cr79f75lhtgme744084g
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
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
	// The plug-in ID.
	//
	// example:
	//
	// pl-ct8181um1hkiqns9f6e0
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// Specifies whether to return parent resource attachments.
	//
	// example:
	//
	// false
	WithParentResource *bool `json:"withParentResource,omitempty" xml:"withParentResource,omitempty"`
}

func (s ListPluginAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListPluginAttachmentsRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *ListPluginAttachmentsRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPluginAttachmentsRequest) GetAttachResourceTypes() *string {
	return s.AttachResourceTypes
}

func (s *ListPluginAttachmentsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListPluginAttachmentsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPluginAttachmentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPluginAttachmentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPluginAttachmentsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListPluginAttachmentsRequest) GetWithParentResource() *bool {
	return s.WithParentResource
}

func (s *ListPluginAttachmentsRequest) SetAttachResourceId(v string) *ListPluginAttachmentsRequest {
	s.AttachResourceId = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetAttachResourceType(v string) *ListPluginAttachmentsRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetAttachResourceTypes(v string) *ListPluginAttachmentsRequest {
	s.AttachResourceTypes = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetEnvironmentId(v string) *ListPluginAttachmentsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetGatewayId(v string) *ListPluginAttachmentsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetPageNumber(v int32) *ListPluginAttachmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetPageSize(v int32) *ListPluginAttachmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetPluginId(v string) *ListPluginAttachmentsRequest {
	s.PluginId = &v
	return s
}

func (s *ListPluginAttachmentsRequest) SetWithParentResource(v bool) *ListPluginAttachmentsRequest {
	s.WithParentResource = &v
	return s
}

func (s *ListPluginAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
