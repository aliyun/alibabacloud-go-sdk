// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDataAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoTraceEnabled(v bool) *TagDataAssetsRequest
	GetAutoTraceEnabled() *bool
	SetDataAssetIds(v []*string) *TagDataAssetsRequest
	GetDataAssetIds() []*string
	SetDataAssetType(v string) *TagDataAssetsRequest
	GetDataAssetType() *string
	SetEnvType(v string) *TagDataAssetsRequest
	GetEnvType() *string
	SetProjectId(v int64) *TagDataAssetsRequest
	GetProjectId() *int64
	SetTags(v []*TagDataAssetsRequestTags) *TagDataAssetsRequest
	GetTags() []*TagDataAssetsRequestTags
}

type TagDataAssetsRequest struct {
	// Specifies whether to enable lineage-based automatic backtracking.
	//
	// example:
	//
	// false
	AutoTraceEnabled *bool `json:"AutoTraceEnabled,omitempty" xml:"AutoTraceEnabled,omitempty"`
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
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tags that you want to add to data assets.
	//
	// This parameter is required.
	Tags []*TagDataAssetsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagDataAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s TagDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *TagDataAssetsRequest) GetAutoTraceEnabled() *bool {
	return s.AutoTraceEnabled
}

func (s *TagDataAssetsRequest) GetDataAssetIds() []*string {
	return s.DataAssetIds
}

func (s *TagDataAssetsRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *TagDataAssetsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *TagDataAssetsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TagDataAssetsRequest) GetTags() []*TagDataAssetsRequestTags {
	return s.Tags
}

func (s *TagDataAssetsRequest) SetAutoTraceEnabled(v bool) *TagDataAssetsRequest {
	s.AutoTraceEnabled = &v
	return s
}

func (s *TagDataAssetsRequest) SetDataAssetIds(v []*string) *TagDataAssetsRequest {
	s.DataAssetIds = v
	return s
}

func (s *TagDataAssetsRequest) SetDataAssetType(v string) *TagDataAssetsRequest {
	s.DataAssetType = &v
	return s
}

func (s *TagDataAssetsRequest) SetEnvType(v string) *TagDataAssetsRequest {
	s.EnvType = &v
	return s
}

func (s *TagDataAssetsRequest) SetProjectId(v int64) *TagDataAssetsRequest {
	s.ProjectId = &v
	return s
}

func (s *TagDataAssetsRequest) SetTags(v []*TagDataAssetsRequestTags) *TagDataAssetsRequest {
	s.Tags = v
	return s
}

func (s *TagDataAssetsRequest) Validate() error {
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

type TagDataAssetsRequestTags struct {
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

func (s TagDataAssetsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s TagDataAssetsRequestTags) GoString() string {
	return s.String()
}

func (s *TagDataAssetsRequestTags) GetKey() *string {
	return s.Key
}

func (s *TagDataAssetsRequestTags) GetValue() *string {
	return s.Value
}

func (s *TagDataAssetsRequestTags) SetKey(v string) *TagDataAssetsRequestTags {
	s.Key = &v
	return s
}

func (s *TagDataAssetsRequestTags) SetValue(v string) *TagDataAssetsRequestTags {
	s.Value = &v
	return s
}

func (s *TagDataAssetsRequestTags) Validate() error {
	return dara.Validate(s)
}
