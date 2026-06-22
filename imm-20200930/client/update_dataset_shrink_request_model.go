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
	// The dataset configuration.
	DatasetConfigShrink *string `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// The maximum number of bindings for the dataset. Valid values: 1 to 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities, such as data files, file relationships, and cluster groups, in the dataset. The maximum value is 2^63 - 1.
	//
	// > This parameter is reserved and not enforced in practice.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in the dataset. Valid values: 1 to 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in the dataset. The maximum value is 2^63 - 1.
	//
	// > This parameter is reserved and not enforced in practice.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total size of all files in the dataset, in bytes. If this limit is exceeded, you can no longer add new index entries. The maximum value is 2^63 - 1.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The dataset name. For information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The dataset description.
	//
	// example:
	//
	// immtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The workflow template ID. For more information, see [Workflow templates and operators](https://help.aliyun.com/document_detail/466304.html).
	//
	// example:
	//
	// Official:ImageManagement
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is invalid.
	WorkflowParametersShrink *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
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

func (s *UpdateDatasetShrinkRequest) SetDatasetConfigShrink(v string) *UpdateDatasetShrinkRequest {
	s.DatasetConfigShrink = &v
	return s
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
