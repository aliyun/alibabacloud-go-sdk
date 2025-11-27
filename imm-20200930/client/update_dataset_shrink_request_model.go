// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *UpdateDatasetShrinkRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *UpdateDatasetShrinkRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *UpdateDatasetShrinkRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *UpdateDatasetShrinkRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetShrinkRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDatasetName(v string) *UpdateDatasetShrinkRequest
	GetDatasetName() *string
	SetDescription(v string) *UpdateDatasetShrinkRequest
	GetDescription() *string
	SetProjectName(v string) *UpdateDatasetShrinkRequest
	GetProjectName() *string
	SetTemplateId(v string) *UpdateDatasetShrinkRequest
	GetTemplateId() *string
	SetWorkflowParametersShrink(v string) *UpdateDatasetShrinkRequest
	GetWorkflowParametersShrink() *string
}

type UpdateDatasetShrinkRequest struct {
	// The maximum number of bindings per dataset. The value range is from 1 to 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities (including data files, file relationships, clustering groups, etc.) per dataset, with a maximum value of 2^63-1.
	//
	// >Reserved parameter, no actual restriction in use.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files per dataset. The value range is from 1 to 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships per dataset, with a maximum value of 2^63-1.
	//
	// >Reserved parameter, no actual restriction in use.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total size of files in each dataset. Once the limit is exceeded, no more indexes can be added. The maximum value is 2^63-1, measured in bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// Dataset name, for how to obtain it, please refer to [Create Dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// Description of the dataset.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Project name, for how to obtain it, please refer to [Create Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Workflow template ID. For more information, please refer to [Workflow Templates and Operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// Official:ImageManagement
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Invalid parameter.
	WorkflowParametersShrink *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
}

func (s UpdateDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetShrinkRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *UpdateDatasetShrinkRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *UpdateDatasetShrinkRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *UpdateDatasetShrinkRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *UpdateDatasetShrinkRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *UpdateDatasetShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateDatasetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateDatasetShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateDatasetShrinkRequest) GetWorkflowParametersShrink() *string {
	return s.WorkflowParametersShrink
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxBindCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxEntityCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxFileCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxRelationCount(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetShrinkRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDatasetName(v string) *UpdateDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetDescription(v string) *UpdateDatasetShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetProjectName(v string) *UpdateDatasetShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetTemplateId(v string) *UpdateDatasetShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) SetWorkflowParametersShrink(v string) *UpdateDatasetShrinkRequest {
	s.WorkflowParametersShrink = &v
	return s
}

func (s *UpdateDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
