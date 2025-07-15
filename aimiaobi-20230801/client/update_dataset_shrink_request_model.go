// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetConfigShrink(v string) *UpdateDatasetShrinkRequest
	GetDatasetConfigShrink() *string
	SetDatasetDescription(v string) *UpdateDatasetShrinkRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *UpdateDatasetShrinkRequest
	GetDatasetId() *int64
	SetSearchDatasetEnable(v int32) *UpdateDatasetShrinkRequest
	GetSearchDatasetEnable() *int32
	SetWorkspaceId(v string) *UpdateDatasetShrinkRequest
	GetWorkspaceId() *string
}

type UpdateDatasetShrinkRequest struct {
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// example:
	//
	// 企业自定义数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetShrinkRequest) GetDatasetConfigShrink() *string {
	return s.DatasetConfigShrink
}

func (s *UpdateDatasetShrinkRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *UpdateDatasetShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *UpdateDatasetShrinkRequest) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *UpdateDatasetShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetShrinkRequest) SetDatasetConfigShrink(v string) *UpdateDatasetShrinkRequest {
	s.DatasetConfigShrink = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetDescription(v string) *UpdateDatasetShrinkRequest {
	s.DatasetDescription = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetId(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetSearchDatasetEnable(v int32) *UpdateDatasetShrinkRequest {
	s.SearchDatasetEnable = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetWorkspaceId(v string) *UpdateDatasetShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
