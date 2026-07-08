// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevel(v string) *CreateDatasetShrinkRequest
	GetAccessLevel() *string
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
	// example:
	//
	// private
	AccessLevel *string `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	// The dataset search configuration.
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// The description of the dataset. This is the display name in the console. Use a human-readable name.
	//
	// example:
	//
	// 企业知识库
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// The name of the dataset. The name must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// businessDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The type of the dataset. Valid values:
	//
	// - CustomSemanticSearch: A custom semantic index. This is the default value. Upload documents to build the dataset.
	//
	// - ThirdSearch: A third-party search source (API). Configure your own search API.
	//
	// example:
	//
	// CustomSemanticSearch
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// Dataset index configuration.
	DocumentHandleConfigShrink *string `json:"DocumentHandleConfig,omitempty" xml:"DocumentHandleConfig,omitempty"`
	// The invocation method. Currently, only portal is supported, which indicates an invocation from the console.
	//
	// - If left empty: When DatasetType is ThirdSearch, datasetConfig.SearchSourceConfigs (third-party API definition) is required.
	//
	// - If set to portal: When DatasetType is ThirdSearch, the system initializes a SearchSourceConfigs (third-party API demo) example by default for your reference.
	//
	// example:
	//
	// portal
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The dataset search switch. Valid values:
	//
	// - 0: Disabled for all.
	//
	// - 1: Visible only to Miao Search.
	//
	// - 2: Visible only to Miao Bi.
	//
	// - 3: Visible to both Miao Search and Miao Bi. This is the default value.
	//
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Obtain a workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
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

func (s *CreateDatasetShrinkRequest) GetAccessLevel() *string {
	return s.AccessLevel
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

func (s *CreateDatasetShrinkRequest) SetAccessLevel(v string) *CreateDatasetShrinkRequest {
	s.AccessLevel = &v
	return s
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
