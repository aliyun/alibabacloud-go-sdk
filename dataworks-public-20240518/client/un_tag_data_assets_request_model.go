// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagDataAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataAssetIds(v []*string) *UnTagDataAssetsRequest
	GetDataAssetIds() []*string
	SetDataAssetType(v string) *UnTagDataAssetsRequest
	GetDataAssetType() *string
	SetEnvType(v string) *UnTagDataAssetsRequest
	GetEnvType() *string
	SetProjectId(v int64) *UnTagDataAssetsRequest
	GetProjectId() *int64
	SetTags(v []*UnTagDataAssetsRequestTags) *UnTagDataAssetsRequest
	GetTags() []*UnTagDataAssetsRequestTags
}

type UnTagDataAssetsRequest struct {
	// The data asset IDs.
	//
	// This parameter is required.
	DataAssetIds []*string `json:"DataAssetIds,omitempty" xml:"DataAssetIds,omitempty" type:"Repeated"`
	// The type of the data asset. Valid values:
	//
	// 	- ACS::DataWorks::Table
	//
	// 	- ACS::DataWorks::Task
	//
	// This parameter is required.
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tags that you want to remove.
	//
	// This parameter is required.
	Tags []*UnTagDataAssetsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UnTagDataAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *UnTagDataAssetsRequest) GetDataAssetIds() []*string {
	return s.DataAssetIds
}

func (s *UnTagDataAssetsRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *UnTagDataAssetsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UnTagDataAssetsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UnTagDataAssetsRequest) GetTags() []*UnTagDataAssetsRequestTags {
	return s.Tags
}

func (s *UnTagDataAssetsRequest) SetDataAssetIds(v []*string) *UnTagDataAssetsRequest {
	s.DataAssetIds = v
	return s
}

func (s *UnTagDataAssetsRequest) SetDataAssetType(v string) *UnTagDataAssetsRequest {
	s.DataAssetType = &v
	return s
}

func (s *UnTagDataAssetsRequest) SetEnvType(v string) *UnTagDataAssetsRequest {
	s.EnvType = &v
	return s
}

func (s *UnTagDataAssetsRequest) SetProjectId(v int64) *UnTagDataAssetsRequest {
	s.ProjectId = &v
	return s
}

func (s *UnTagDataAssetsRequest) SetTags(v []*UnTagDataAssetsRequestTags) *UnTagDataAssetsRequest {
	s.Tags = v
	return s
}

func (s *UnTagDataAssetsRequest) Validate() error {
	return dara.Validate(s)
}

type UnTagDataAssetsRequestTags struct {
	// The tag key.
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

func (s UnTagDataAssetsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UnTagDataAssetsRequestTags) GoString() string {
	return s.String()
}

func (s *UnTagDataAssetsRequestTags) GetKey() *string {
	return s.Key
}

func (s *UnTagDataAssetsRequestTags) GetValue() *string {
	return s.Value
}

func (s *UnTagDataAssetsRequestTags) SetKey(v string) *UnTagDataAssetsRequestTags {
	s.Key = &v
	return s
}

func (s *UnTagDataAssetsRequestTags) SetValue(v string) *UnTagDataAssetsRequestTags {
	s.Value = &v
	return s
}

func (s *UnTagDataAssetsRequestTags) Validate() error {
	return dara.Validate(s)
}
