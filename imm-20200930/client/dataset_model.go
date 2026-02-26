// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataset interface {
	dara.Model
	String() string
	GoString() string
	SetBindCount(v int64) *Dataset
	GetBindCount() *int64
	SetCreateTime(v string) *Dataset
	GetCreateTime() *string
	SetDatasetConfig(v *DatasetConfig) *Dataset
	GetDatasetConfig() *DatasetConfig
	SetDatasetMaxBindCount(v int64) *Dataset
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *Dataset
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *Dataset
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *Dataset
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *Dataset
	GetDatasetMaxTotalFileSize() *int64
	SetDatasetName(v string) *Dataset
	GetDatasetName() *string
	SetDescription(v string) *Dataset
	GetDescription() *string
	SetFileCount(v int64) *Dataset
	GetFileCount() *int64
	SetProjectName(v string) *Dataset
	GetProjectName() *string
	SetTemplateId(v string) *Dataset
	GetTemplateId() *string
	SetTotalFileSize(v int64) *Dataset
	GetTotalFileSize() *int64
	SetUpdateTime(v string) *Dataset
	GetUpdateTime() *string
	SetWorkflowParameters(v []*WorkflowParameter) *Dataset
	GetWorkflowParameters() []*WorkflowParameter
}

type Dataset struct {
	// The current number of OSS buckets that are bound to the dataset.
	//
	// example:
	//
	// 2
	BindCount *int64 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The timestamp when the dataset was created. The timestamp must be in the RFC3339Nano format.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	CreateTime    *string        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetConfig *DatasetConfig `json:"DatasetConfig,omitempty" xml:"DatasetConfig,omitempty"`
	// The maximum number of bindings for the dataset.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities for the dataset.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files for the dataset.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships for the dataset.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total size of files in the dataset. Unit: bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// dataset001
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The dataset description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The current number of files in the dataset.
	//
	// example:
	//
	// 10
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the workflow template.
	//
	// example:
	//
	// DefaultId
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The total size of files in the dataset. Unit: bytes.
	//
	// example:
	//
	// 100000
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// The timestamp when the dataset was last modified. The timestamp must be in the RFC3339Nano format.
	//
	// >  If a dataset has never been modified after it was created, the timestamp when the dataset was last modified is the same as the timestamp when the dataset was created.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 自定义参数
	//
	// if can be null:
	// true
	WorkflowParameters []*WorkflowParameter `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty" type:"Repeated"`
}

func (s Dataset) String() string {
	return dara.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) GetBindCount() *int64 {
	return s.BindCount
}

func (s *Dataset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Dataset) GetDatasetConfig() *DatasetConfig {
	return s.DatasetConfig
}

func (s *Dataset) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *Dataset) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *Dataset) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *Dataset) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *Dataset) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *Dataset) GetDatasetName() *string {
	return s.DatasetName
}

func (s *Dataset) GetDescription() *string {
	return s.Description
}

func (s *Dataset) GetFileCount() *int64 {
	return s.FileCount
}

func (s *Dataset) GetProjectName() *string {
	return s.ProjectName
}

func (s *Dataset) GetTemplateId() *string {
	return s.TemplateId
}

func (s *Dataset) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *Dataset) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Dataset) GetWorkflowParameters() []*WorkflowParameter {
	return s.WorkflowParameters
}

func (s *Dataset) SetBindCount(v int64) *Dataset {
	s.BindCount = &v
	return s
}

func (s *Dataset) SetCreateTime(v string) *Dataset {
	s.CreateTime = &v
	return s
}

func (s *Dataset) SetDatasetConfig(v *DatasetConfig) *Dataset {
	s.DatasetConfig = v
	return s
}

func (s *Dataset) SetDatasetMaxBindCount(v int64) *Dataset {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxEntityCount(v int64) *Dataset {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxFileCount(v int64) *Dataset {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxRelationCount(v int64) *Dataset {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *Dataset) SetDatasetMaxTotalFileSize(v int64) *Dataset {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *Dataset) SetDatasetName(v string) *Dataset {
	s.DatasetName = &v
	return s
}

func (s *Dataset) SetDescription(v string) *Dataset {
	s.Description = &v
	return s
}

func (s *Dataset) SetFileCount(v int64) *Dataset {
	s.FileCount = &v
	return s
}

func (s *Dataset) SetProjectName(v string) *Dataset {
	s.ProjectName = &v
	return s
}

func (s *Dataset) SetTemplateId(v string) *Dataset {
	s.TemplateId = &v
	return s
}

func (s *Dataset) SetTotalFileSize(v int64) *Dataset {
	s.TotalFileSize = &v
	return s
}

func (s *Dataset) SetUpdateTime(v string) *Dataset {
	s.UpdateTime = &v
	return s
}

func (s *Dataset) SetWorkflowParameters(v []*WorkflowParameter) *Dataset {
	s.WorkflowParameters = v
	return s
}

func (s *Dataset) Validate() error {
	if s.DatasetConfig != nil {
		if err := s.DatasetConfig.Validate(); err != nil {
			return err
		}
	}
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
