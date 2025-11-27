// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetMaxBindCount(v int64) *CreateDatasetRequest
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *CreateDatasetRequest
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *CreateDatasetRequest
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *CreateDatasetRequest
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *CreateDatasetRequest
	GetDatasetMaxTotalFileSize() *int64
	SetDatasetName(v string) *CreateDatasetRequest
	GetDatasetName() *string
	SetDescription(v string) *CreateDatasetRequest
	GetDescription() *string
	SetProjectName(v string) *CreateDatasetRequest
	GetProjectName() *string
	SetTemplateId(v string) *CreateDatasetRequest
	GetTemplateId() *string
	SetWorkflowParameters(v []*WorkflowParameter) *CreateDatasetRequest
	GetWorkflowParameters() []*WorkflowParameter
}

type CreateDatasetRequest struct {
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
	WorkflowParameters []*WorkflowParameter `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty" type:"Repeated"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *CreateDatasetRequest) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *CreateDatasetRequest) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *CreateDatasetRequest) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *CreateDatasetRequest) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *CreateDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateDatasetRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateDatasetRequest) GetWorkflowParameters() []*WorkflowParameter {
	return s.WorkflowParameters
}

func (s *CreateDatasetRequest) SetDatasetMaxBindCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxEntityCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxFileCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxRelationCount(v int64) *CreateDatasetRequest {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetMaxTotalFileSize(v int64) *CreateDatasetRequest {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetProjectName(v string) *CreateDatasetRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDatasetRequest) SetTemplateId(v string) *CreateDatasetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDatasetRequest) SetWorkflowParameters(v []*WorkflowParameter) *CreateDatasetRequest {
	s.WorkflowParameters = v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
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
