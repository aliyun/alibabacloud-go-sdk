// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *UpdateDatasetRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *UpdateDatasetRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *UpdateDatasetRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *UpdateDatasetRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDatasetName(v string) *UpdateDatasetRequest
	GetDatasetName() *string
	SetDescription(v string) *UpdateDatasetRequest
	GetDescription() *string
	SetProjectName(v string) *UpdateDatasetRequest
	GetProjectName() *string
	SetTemplateId(v string) *UpdateDatasetRequest
	GetTemplateId() *string
	SetWorkflowParameters(v []*WorkflowParameter) *UpdateDatasetRequest
	GetWorkflowParameters() []*WorkflowParameter
}

type UpdateDatasetRequest struct {
	// The maximum number of bindings for each dataset. Valid values: 1 to 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities (including data files, file relationships, and clustering groups) in each dataset. The maximum value is 2^63-1.
	//
	// >This is a reserved parameter and is not enforced during use.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in each dataset. Valid values: 1 to 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in each dataset. The maximum value is 2^63-1.
	//
	// >This is a reserved parameter and is not enforced during use.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total file size in each dataset. After this limit is exceeded, no more indexes can be added. The maximum value is 2^63-1. Unit: bytes.
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
	// The description of the dataset.
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
	// Deprecated
	//
	// Invalid parameter.
	WorkflowParameters []*WorkflowParameter `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *UpdateDatasetRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *UpdateDatasetRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *UpdateDatasetRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *UpdateDatasetRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *UpdateDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateDatasetRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateDatasetRequest) GetWorkflowParameters() []*WorkflowParameter {
	return s.WorkflowParameters
}

func (s *UpdateDatasetRequest) SetDatasetMaxBindCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxEntityCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxFileCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxRelationCount(v int64) *UpdateDatasetRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetMaxTotalFileSize(v int64) *UpdateDatasetRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *UpdateDatasetRequest) SetDatasetName(v string) *UpdateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetProjectName(v string) *UpdateDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDatasetRequest) SetTemplateId(v string) *UpdateDatasetRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDatasetRequest) SetWorkflowParameters(v []*WorkflowParameter) *UpdateDatasetRequest {
	s.WorkflowParameters = v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	if s.WorkflowParameters != nil {
		for _, item := range s.WorkflowParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
