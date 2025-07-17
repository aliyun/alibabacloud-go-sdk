// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataAssetIdsShrink(v string) *ListDataAssetsShrinkRequest
	GetDataAssetIdsShrink() *string
	SetDataAssetType(v string) *ListDataAssetsShrinkRequest
	GetDataAssetType() *string
	SetEnvType(v string) *ListDataAssetsShrinkRequest
	GetEnvType() *string
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
	// The data asset IDs.
	DataAssetIdsShrink *string `json:"DataAssetIds,omitempty" xml:"DataAssetIds,omitempty"`
	// The type of the data asset. Valid values:
	//
	// 	- ACS::DataWorks::Table
	//
	// 	- ACS::DataWorks::Task
	//
	// example:
	//
	// ACS::DataWorks::Task
	DataAssetType *string `json:"DataAssetType,omitempty" xml:"DataAssetType,omitempty"`
	// The environment of the workspace to which the data asset belongs. Valid values:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tags that are added to data assets. This parameter specifies a filter condition.
	//
	// 	- You can specify multiple tags, which are in the logical OR relation. For example, you can query the data assets that contain one of the following tags: `["key1:v1", "key2:v1", "key3:v1"]`.
	//
	// 	- If you do not configure this parameter, tag-based filtering is not performed.
	//
	// This parameter is required.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListDataAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsShrinkRequest) GoString() string {
	return s.String()
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
