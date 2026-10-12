// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDomainId(v int64) *ListDataAssetsShrinkRequest
	GetAssetDomainId() *int64
	SetCategoryUuid(v string) *ListDataAssetsShrinkRequest
	GetCategoryUuid() *string
	SetDataAssetIdsShrink(v string) *ListDataAssetsShrinkRequest
	GetDataAssetIdsShrink() *string
	SetDataAssetType(v string) *ListDataAssetsShrinkRequest
	GetDataAssetType() *string
	SetEnvType(v string) *ListDataAssetsShrinkRequest
	GetEnvType() *string
	SetName(v string) *ListDataAssetsShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataAssetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAssetsShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataAssetsShrinkRequest
	GetProjectId() *int64
	SetTagsShrink(v string) *ListDataAssetsShrinkRequest
	GetTagsShrink() *string
}

type ListDataAssetsShrinkRequest struct {
	// The ID of the asset domain.
	//
	// example:
	//
	// 1001
	AssetDomainId *int64 `json:"AssetDomainId,omitempty" xml:"AssetDomainId,omitempty"`
	// The ID of the asset category.
	//
	// example:
	//
	// cate-xxxxxxxx
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// The list of unique IDs of the data assets.
	DataAssetIdsShrink *string `json:"DataAssetIds,omitempty" xml:"DataAssetIds,omitempty"`
	// The asset type of the data asset.
	//
	// - DataWorks table (ACS::DataWorks::Table)
	//
	// - DataWorks scheduled task (ACS::DataWorks::Task)
	//
	// example:
	//
	// ACS::DataWorks::Task
	DataAssetType *string `json:"DataAssetType,omitempty" xml:"DataAssetType,omitempty"`
	// The workspace environment to which the data asset belongs. Valid values:
	//
	// - Dev: development environment
	//
	// - Prod: production environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the asset. Fuzzy search by name is supported.
	//
	// example:
	//
	// Asset domain name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of tags associated with the data assets. This parameter is used to filter query results based on tags:
	//
	// - Multiple values are evaluated with an OR logical operator. For example, if you specify `["key1:v1", "key2:v1", "key3:v1"]`, the system queries data assets that contain at least one of the specified tags.
	//
	// - If this parameter is not specified or is left empty, no tag-based filtering is performed.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListDataAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataAssetsShrinkRequest) GetAssetDomainId() *int64 {
	return s.AssetDomainId
}

func (s *ListDataAssetsShrinkRequest) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *ListDataAssetsShrinkRequest) GetDataAssetIdsShrink() *string {
	return s.DataAssetIdsShrink
}

func (s *ListDataAssetsShrinkRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *ListDataAssetsShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataAssetsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListDataAssetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAssetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataAssetsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListDataAssetsShrinkRequest) SetAssetDomainId(v int64) *ListDataAssetsShrinkRequest {
	s.AssetDomainId = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetCategoryUuid(v string) *ListDataAssetsShrinkRequest {
	s.CategoryUuid = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetDataAssetIdsShrink(v string) *ListDataAssetsShrinkRequest {
	s.DataAssetIdsShrink = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetDataAssetType(v string) *ListDataAssetsShrinkRequest {
	s.DataAssetType = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetEnvType(v string) *ListDataAssetsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetName(v string) *ListDataAssetsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetPageNumber(v int32) *ListDataAssetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetPageSize(v int32) *ListDataAssetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetProjectId(v int64) *ListDataAssetsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) SetTagsShrink(v string) *ListDataAssetsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListDataAssetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
