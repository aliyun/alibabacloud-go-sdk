// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The maximum number of bindings per dataset. The range is 1~10, with a default value of 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in each dataset. The default value is 10000000000.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in each dataset. The range is 1~100000000, with a default value of 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in each dataset. The default value is 100000000000.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total size of files in each dataset. Once the limit is exceeded, no more indexes can be added. The default value is 90000000000000000, in bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The name of the dataset, which must be unique under the same Project. Naming rules are as follows:
	//
	// - Length should be 1~128 characters.
	//
	// - Can only contain English letters, numbers, hyphens (-), and underscores (_).
	//
	// - Must start with an English letter or underscore (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Description of the dataset. The length should be 1~256 English or Chinese characters, with a default value of empty.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the project. For more information on how to obtain it, see [Create Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Workflow template ID. For more information, see [Workflow Templates and Operators](https://help.aliyun.com/document_detail/466304.html). The default value is empty.
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
