// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagDataAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataAssetIdsShrink(v string) *UnTagDataAssetsShrinkRequest
	GetDataAssetIdsShrink() *string
	SetDataAssetType(v string) *UnTagDataAssetsShrinkRequest
	GetDataAssetType() *string
	SetEnvType(v string) *UnTagDataAssetsShrinkRequest
	GetEnvType() *string
	SetProjectId(v int64) *UnTagDataAssetsShrinkRequest
	GetProjectId() *int64
	SetTagsShrink(v string) *UnTagDataAssetsShrinkRequest
	GetTagsShrink() *string
}

type UnTagDataAssetsShrinkRequest struct {
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
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tags that you want to remove.
	//
	// This parameter is required.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UnTagDataAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagDataAssetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnTagDataAssetsShrinkRequest) GetDataAssetIdsShrink() *string {
	return s.DataAssetIdsShrink
}

func (s *UnTagDataAssetsShrinkRequest) GetDataAssetType() *string {
	return s.DataAssetType
}

func (s *UnTagDataAssetsShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *UnTagDataAssetsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UnTagDataAssetsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UnTagDataAssetsShrinkRequest) SetDataAssetIdsShrink(v string) *UnTagDataAssetsShrinkRequest {
	s.DataAssetIdsShrink = &v
	return s
}

func (s *UnTagDataAssetsShrinkRequest) SetDataAssetType(v string) *UnTagDataAssetsShrinkRequest {
	s.DataAssetType = &v
	return s
}

func (s *UnTagDataAssetsShrinkRequest) SetEnvType(v string) *UnTagDataAssetsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *UnTagDataAssetsShrinkRequest) SetProjectId(v int64) *UnTagDataAssetsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UnTagDataAssetsShrinkRequest) SetTagsShrink(v string) *UnTagDataAssetsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UnTagDataAssetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
