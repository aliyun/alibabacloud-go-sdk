// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDataAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoTraceEnabled(v bool) *TagDataAssetsShrinkRequest
	GetAutoTraceEnabled() *bool
	SetDataAssetIdsShrink(v string) *TagDataAssetsShrinkRequest
	GetDataAssetIdsShrink() *string
	SetDataAssetType(v string) *TagDataAssetsShrinkRequest
	GetDataAssetType() *string
	SetEnvType(v string) *TagDataAssetsShrinkRequest
	GetEnvType() *string
	SetProjectId(v int64) *TagDataAssetsShrinkRequest
	GetProjectId() *int64
	SetTagsShrink(v string) *TagDataAssetsShrinkRequest
	GetTagsShrink() *string
}

type TagDataAssetsShrinkRequest struct {
	// Specifies whether to enable lineage-based automatic backtracking.
	//
	// example:
	//
	// false
	AutoTraceEnabled *bool `json:"AutoTraceEnabled,omitempty" xml:"AutoTraceEnabled,omitempty"`
	// The data asset IDs.
	//
	// This parameter is required.
	DataAssetIdsShrink *string `json:"DataAssetIds,omitempty" xml:"DataAssetIds,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagDataAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TagDataAssetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagDataAssetsShrinkRequest) GetAutoTraceEnabled() *bool {
	return s.AutoTraceEnabled
}

func (s *TagDataAssetsShrinkRequest) GetDataAssetIdsShrink() *string {
	return s.DataAssetIdsShrink
}

func (s *TagDataAssetsShrinkRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *TagDataAssetsShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *TagDataAssetsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TagDataAssetsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *TagDataAssetsShrinkRequest) SetAutoTraceEnabled(v bool) *TagDataAssetsShrinkRequest {
	s.AutoTraceEnabled = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) SetDataAssetIdsShrink(v string) *TagDataAssetsShrinkRequest {
	s.DataAssetIdsShrink = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) SetDataAssetType(v string) *TagDataAssetsShrinkRequest {
	s.DataAssetType = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) SetEnvType(v string) *TagDataAssetsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) SetProjectId(v int64) *TagDataAssetsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) SetTagsShrink(v string) *TagDataAssetsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TagDataAssetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
