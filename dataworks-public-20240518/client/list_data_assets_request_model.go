// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataAssetIds(v []*string) *ListDataAssetsRequest
	GetDataAssetIds() []*string
	SetDataAssetType(v string) *ListDataAssetsRequest
	GetDataAssetType() *string
	SetEnvType(v string) *ListDataAssetsRequest
	GetEnvType() *string
	SetPageNumber(v int32) *ListDataAssetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAssetsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataAssetsRequest
	GetProjectId() *int64
	SetTags(v []*ListDataAssetsRequestTags) *ListDataAssetsRequest
	GetTags() []*ListDataAssetsRequestTags
}

type ListDataAssetsRequest struct {
	// The data asset IDs.
	DataAssetIds []*string `json:"DataAssetIds,omitempty" xml:"DataAssetIds,omitempty" type:"Repeated"`
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
	Tags []*ListDataAssetsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListDataAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListDataAssetsRequest) GetDataAssetIds() []*string {
	return s.DataAssetIds
}

func (s *ListDataAssetsRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *ListDataAssetsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataAssetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataAssetsRequest) GetTags() []*ListDataAssetsRequestTags {
	return s.Tags
}

func (s *ListDataAssetsRequest) SetDataAssetIds(v []*string) *ListDataAssetsRequest {
	s.DataAssetIds = v
	return s
}

func (s *ListDataAssetsRequest) SetDataAssetType(v string) *ListDataAssetsRequest {
	s.DataAssetType = &v
	return s
}

func (s *ListDataAssetsRequest) SetEnvType(v string) *ListDataAssetsRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataAssetsRequest) SetPageNumber(v int32) *ListDataAssetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAssetsRequest) SetPageSize(v int32) *ListDataAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetsRequest) SetProjectId(v int64) *ListDataAssetsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataAssetsRequest) SetTags(v []*ListDataAssetsRequestTags) *ListDataAssetsRequest {
	s.Tags = v
	return s
}

func (s *ListDataAssetsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAssetsRequestTags struct {
	// The tag key.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, and the following characters: `-@#*<>|[]()+=&%$!~`. It cannot start with `dw:`.
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataAssetsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsRequestTags) GoString() string {
	return s.String()
}

func (s *ListDataAssetsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListDataAssetsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListDataAssetsRequestTags) SetKey(v string) *ListDataAssetsRequestTags {
	s.Key = &v
	return s
}

func (s *ListDataAssetsRequestTags) SetValue(v string) *ListDataAssetsRequestTags {
	s.Value = &v
	return s
}

func (s *ListDataAssetsRequestTags) Validate() error {
	return dara.Validate(s)
}
