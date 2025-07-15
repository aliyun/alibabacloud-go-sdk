// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetConfigShrink(v string) *CreateDatasetShrinkRequest
	GetDatasetConfigShrink() *string
	SetDatasetDescription(v string) *CreateDatasetShrinkRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateDatasetShrinkRequest
	GetDatasetName() *string
	SetDatasetType(v string) *CreateDatasetShrinkRequest
	GetDatasetType() *string
	SetDocumentHandleConfigShrink(v string) *CreateDatasetShrinkRequest
	GetDocumentHandleConfigShrink() *string
	SetInvokeType(v string) *CreateDatasetShrinkRequest
	GetInvokeType() *string
	SetSearchDatasetEnable(v int32) *CreateDatasetShrinkRequest
	GetSearchDatasetEnable() *int32
	SetWorkspaceId(v string) *CreateDatasetShrinkRequest
	GetWorkspaceId() *string
}

type CreateDatasetShrinkRequest struct {
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// example:
	//
	// 企业自定义数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// businessDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType                *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	DocumentHandleConfigShrink *string `json:"DocumentHandleConfig,omitempty" xml:"DocumentHandleConfig,omitempty"`
	// example:
	//
	// portal
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetShrinkRequest) GetDatasetConfigShrink() *string {
	return s.DatasetConfigShrink
}

func (s *CreateDatasetShrinkRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateDatasetShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetShrinkRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateDatasetShrinkRequest) GetDocumentHandleConfigShrink() *string {
	return s.DocumentHandleConfigShrink
}

func (s *CreateDatasetShrinkRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *CreateDatasetShrinkRequest) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *CreateDatasetShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetShrinkRequest) SetDatasetConfigShrink(v string) *CreateDatasetShrinkRequest {
	s.DatasetConfigShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetDescription(v string) *CreateDatasetShrinkRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetName(v string) *CreateDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetType(v string) *CreateDatasetShrinkRequest {
	s.DatasetType = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDocumentHandleConfigShrink(v string) *CreateDatasetShrinkRequest {
	s.DocumentHandleConfigShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetInvokeType(v string) *CreateDatasetShrinkRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetSearchDatasetEnable(v int32) *CreateDatasetShrinkRequest {
	s.SearchDatasetEnable = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetWorkspaceId(v string) *CreateDatasetShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
