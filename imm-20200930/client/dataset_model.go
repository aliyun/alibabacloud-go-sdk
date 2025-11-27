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
	BindCount               *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetMaxBindCount     *int64  `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	DatasetMaxEntityCount   *int64  `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	DatasetMaxFileCount     *int64  `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	DatasetMaxRelationCount *int64  `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	DatasetMaxTotalFileSize *int64  `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	DatasetName             *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileCount               *int64  `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	ProjectName             *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TotalFileSize           *int64  `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	UpdateTime              *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
