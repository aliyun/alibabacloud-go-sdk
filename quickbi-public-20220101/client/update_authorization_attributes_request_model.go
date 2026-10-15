// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditPermissionType(v int32) *UpdateAuthorizationAttributesRequest
	GetEditPermissionType() *int32
	SetResourceId(v string) *UpdateAuthorizationAttributesRequest
	GetResourceId() *string
	SetResourceType(v string) *UpdateAuthorizationAttributesRequest
	GetResourceType() *string
	SetUsePermissionType(v int32) *UpdateAuthorizationAttributesRequest
	GetUsePermissionType() *int32
}

type UpdateAuthorizationAttributesRequest struct {
	// The edit permission type. Valid values:
	//
	// - 0: Private.
	//
	// - 11: All workspace members except guests. For data forms, datasets, and data sources, this value indicates all workspace developers.
	//
	// - 12: Specified members. For data forms, datasets, and data sources, this value indicates specified developers.
	//
	// example:
	//
	// 0
	EditPermissionType *int32 `json:"EditPermissionType,omitempty" xml:"EditPermissionType,omitempty"`
	// The ID of the resource to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - dashboard: dashboard.
	//
	// - screen: data visualization.
	//
	// - report: workbook.
	//
	// - ANALYSIS: ad hoc analysis.
	//
	// - dashboardOfflineQuery: self-service data retrieval.
	//
	// - dataForm: data form.
	//
	// - cube: dataset.
	//
	// - datasource: data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// cube
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The usage permission type for datasets and data sources. Valid values:
	//
	// - 1: All workspace members except guests.
	//
	// - 2: Specified members.
	//
	// example:
	//
	// 1
	UsePermissionType *int32 `json:"UsePermissionType,omitempty" xml:"UsePermissionType,omitempty"`
}

func (s UpdateAuthorizationAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributesRequest) GetEditPermissionType() *int32 {
	return s.EditPermissionType
}

func (s *UpdateAuthorizationAttributesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateAuthorizationAttributesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateAuthorizationAttributesRequest) GetUsePermissionType() *int32 {
	return s.UsePermissionType
}

func (s *UpdateAuthorizationAttributesRequest) SetEditPermissionType(v int32) *UpdateAuthorizationAttributesRequest {
	s.EditPermissionType = &v
	return s
}

func (s *UpdateAuthorizationAttributesRequest) SetResourceId(v string) *UpdateAuthorizationAttributesRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateAuthorizationAttributesRequest) SetResourceType(v string) *UpdateAuthorizationAttributesRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateAuthorizationAttributesRequest) SetUsePermissionType(v int32) *UpdateAuthorizationAttributesRequest {
	s.UsePermissionType = &v
	return s
}

func (s *UpdateAuthorizationAttributesRequest) Validate() error {
	return dara.Validate(s)
}
