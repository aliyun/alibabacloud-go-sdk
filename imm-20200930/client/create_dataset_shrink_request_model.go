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
	SetDatasetMaxBindCount(v int64) *CreateDatasetShrinkRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *CreateDatasetShrinkRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *CreateDatasetShrinkRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *CreateDatasetShrinkRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *CreateDatasetShrinkRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDatasetName(v string) *CreateDatasetShrinkRequest
	GetDatasetName() *string
	SetDescription(v string) *CreateDatasetShrinkRequest
	GetDescription() *string
	SetProjectName(v string) *CreateDatasetShrinkRequest
	GetProjectName() *string
	SetTemplateId(v string) *CreateDatasetShrinkRequest
	GetTemplateId() *string
	SetWorkflowParametersShrink(v string) *CreateDatasetShrinkRequest
	GetWorkflowParametersShrink() *string
}

type CreateDatasetShrinkRequest struct {
	// The dataset configuration.
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// The maximum number of bindings for each dataset. Valid values: 1 to 10. The default value is 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in each dataset. The default value is 10,000,000,000.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in each dataset. Valid values: 1 to 100,000,000. The default value is 100,000,000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in each dataset. The default value is 100,000,000,000.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total size of files in each dataset. You cannot add more indexes after exceeding this limit. The default value is 90,000,000,000,000,000 bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The dataset name. It must be unique within the same project. Naming conventions are as follows:
	//
	// - Length: 1 to 128 characters.
	//
	// - Only English letters, digits, hyphens (-), and underscores (_) are allowed.
	//
	// - Must start with an English letter or an underscore (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description of the dataset. Length: 1 to 256 English or Chinese characters. The default value is empty.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The workflow template ID. For more information, see [Workflow Templates and Operators](https://help.aliyun.com/document_detail/466304.html). The default value is empty.
	//
	// example:
	//
	// Official:ImageManagement
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Invalid parameter.
	WorkflowParametersShrink *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
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

func (s *CreateDatasetShrinkRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *CreateDatasetShrinkRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *CreateDatasetShrinkRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *CreateDatasetShrinkRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *CreateDatasetShrinkRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *CreateDatasetShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateDatasetShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateDatasetShrinkRequest) GetWorkflowParametersShrink() *string {
	return s.WorkflowParametersShrink
}

func (s *CreateDatasetShrinkRequest) SetDatasetConfigShrink(v string) *CreateDatasetShrinkRequest {
	s.DatasetConfigShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetMaxBindCount(v int64) *CreateDatasetShrinkRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetMaxEntityCount(v int64) *CreateDatasetShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetMaxFileCount(v int64) *CreateDatasetShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetMaxRelationCount(v int64) *CreateDatasetShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetMaxTotalFileSize(v int64) *CreateDatasetShrinkRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDatasetName(v string) *CreateDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetDescription(v string) *CreateDatasetShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetProjectName(v string) *CreateDatasetShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetTemplateId(v string) *CreateDatasetShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDatasetShrinkRequest) SetWorkflowParametersShrink(v string) *CreateDatasetShrinkRequest {
	s.WorkflowParametersShrink = &v
	return s
}

func (s *CreateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
