// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProject interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Project
	GetCreateTime() *string
	SetDatasetCount(v int64) *Project
	GetDatasetCount() *int64
	SetDatasetMaxBindCount(v int64) *Project
	GetDatasetMaxBindCount() *int64
	SetDatasetMaxEntityCount(v int64) *Project
	GetDatasetMaxEntityCount() *int64
	SetDatasetMaxFileCount(v int64) *Project
	GetDatasetMaxFileCount() *int64
	SetDatasetMaxRelationCount(v int64) *Project
	GetDatasetMaxRelationCount() *int64
	SetDatasetMaxTotalFileSize(v int64) *Project
	GetDatasetMaxTotalFileSize() *int64
	SetDescription(v string) *Project
	GetDescription() *string
	SetEngineConcurrency(v int64) *Project
	GetEngineConcurrency() *int64
	SetFileCount(v int64) *Project
	GetFileCount() *int64
	SetProjectMaxDatasetCount(v int64) *Project
	GetProjectMaxDatasetCount() *int64
	SetProjectName(v string) *Project
	GetProjectName() *string
	SetProjectQueriesPerSecond(v int64) *Project
	GetProjectQueriesPerSecond() *int64
	SetServiceRole(v string) *Project
	GetServiceRole() *string
	SetTags(v []*ProjectTags) *Project
	GetTags() []*ProjectTags
	SetTemplateId(v string) *Project
	GetTemplateId() *string
	SetTotalFileSize(v int64) *Project
	GetTotalFileSize() *int64
	SetUpdateTime(v string) *Project
	GetUpdateTime() *string
}

type Project struct {
	// The timestamp when the project was created. The timestamp is in the RFC3339Nano format.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current number of datasets in the project.
	//
	// example:
	//
	// 5
	DatasetCount *int64 `json:"DatasetCount,omitempty" xml:"DatasetCount,omitempty"`
	// The maximum number of bindings that a dataset can have. Valid values: 1 to 10. Default value: 10.
	//
	// example:
	//
	// 10
	DatasetMaxBindCount *int64 `json:"DatasetMaxBindCount,omitempty" xml:"DatasetMaxBindCount,omitempty"`
	// The maximum number of metadata entities in a dataset. Default value: 10000000000.
	//
	// >  This parameter is reserved and does not actually apply a limit.
	//
	// example:
	//
	// 10000000000
	DatasetMaxEntityCount *int64 `json:"DatasetMaxEntityCount,omitempty" xml:"DatasetMaxEntityCount,omitempty"`
	// The maximum number of files in a dataset. Valid values: 1 to 100000000. Default value: 100000000.
	//
	// example:
	//
	// 100000000
	DatasetMaxFileCount *int64 `json:"DatasetMaxFileCount,omitempty" xml:"DatasetMaxFileCount,omitempty"`
	// The maximum number of metadata relationships in a dataset. Default value: 100000000000.
	//
	// >  This parameter is reserved and does not actually apply a limit.
	//
	// example:
	//
	// 100000000000
	DatasetMaxRelationCount *int64 `json:"DatasetMaxRelationCount,omitempty" xml:"DatasetMaxRelationCount,omitempty"`
	// The maximum total file size for a dataset. If the total file size exceeds this limit, indexes can no longer be added. Default value: 90000000000000000. Unit: bytes.
	//
	// example:
	//
	// 90000000000000000
	DatasetMaxTotalFileSize *int64 `json:"DatasetMaxTotalFileSize,omitempty" xml:"DatasetMaxTotalFileSize,omitempty"`
	// The project description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of tasks that the project can process per second. This corresponds to the maximum number of operators that can run in parallel in the project. Default value: 100.
	//
	// 	- If the number of synchronous tasks that run in parallel exceeds this limit, the task execution time will be extended until a timeout occurs.
	//
	// 	- If the number of asynchronous tasks that run in parallel exceeds this limit, the tasks will be queued. This causes delayed task completion. If a task remains in the queue for longer than the specified time limit (usually dozens of seconds), the task will fail.
	//
	// example:
	//
	// 100
	EngineConcurrency *int64 `json:"EngineConcurrency,omitempty" xml:"EngineConcurrency,omitempty"`
	// The current number of files in the project.
	//
	// example:
	//
	// 10
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The maximum number of datasets that a project can contain. Valid values: 1 to 1000000000. Default value: 1000000000.
	//
	// example:
	//
	// 1000000000
	ProjectMaxDatasetCount *int64 `json:"ProjectMaxDatasetCount,omitempty" xml:"ProjectMaxDatasetCount,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The maximum number of requests that can be processed by the project per second. This corresponds to the maximum number of API operations that can be called in the project per second. Default value: 100.
	//
	// example:
	//
	// 100
	ProjectQueriesPerSecond *int64 `json:"ProjectQueriesPerSecond,omitempty" xml:"ProjectQueriesPerSecond,omitempty"`
	// The service role.
	//
	// example:
	//
	// AliyunIMMDefaultRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	// The tag list.
	Tags []*ProjectTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the workflow template.
	//
	// example:
	//
	// DefaultId
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The current total size of files in the project. Unit: bytes.
	//
	// example:
	//
	// 100000
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// The timestamp when the project was last modified. The timestamp is in the RFC3339Nano format.
	//
	// >  If a project is not modified after it is created, the timestamp when the project was created is the same as the timestamp when the project was last modified.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Project) String() string {
	return dara.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Project) GetDatasetCount() *int64 {
	return s.DatasetCount
}

func (s *Project) GetDatasetMaxBindCount() *int64 {
	return s.DatasetMaxBindCount
}

func (s *Project) GetDatasetMaxEntityCount() *int64 {
	return s.DatasetMaxEntityCount
}

func (s *Project) GetDatasetMaxFileCount() *int64 {
	return s.DatasetMaxFileCount
}

func (s *Project) GetDatasetMaxRelationCount() *int64 {
	return s.DatasetMaxRelationCount
}

func (s *Project) GetDatasetMaxTotalFileSize() *int64 {
	return s.DatasetMaxTotalFileSize
}

func (s *Project) GetDescription() *string {
	return s.Description
}

func (s *Project) GetEngineConcurrency() *int64 {
	return s.EngineConcurrency
}

func (s *Project) GetFileCount() *int64 {
	return s.FileCount
}

func (s *Project) GetProjectMaxDatasetCount() *int64 {
	return s.ProjectMaxDatasetCount
}

func (s *Project) GetProjectName() *string {
	return s.ProjectName
}

func (s *Project) GetProjectQueriesPerSecond() *int64 {
	return s.ProjectQueriesPerSecond
}

func (s *Project) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *Project) GetTags() []*ProjectTags {
	return s.Tags
}

func (s *Project) GetTemplateId() *string {
	return s.TemplateId
}

func (s *Project) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *Project) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Project) SetCreateTime(v string) *Project {
	s.CreateTime = &v
	return s
}

func (s *Project) SetDatasetCount(v int64) *Project {
	s.DatasetCount = &v
	return s
}

func (s *Project) SetDatasetMaxBindCount(v int64) *Project {
	s.DatasetMaxBindCount = &v
	return s
}

func (s *Project) SetDatasetMaxEntityCount(v int64) *Project {
	s.DatasetMaxEntityCount = &v
	return s
}

func (s *Project) SetDatasetMaxFileCount(v int64) *Project {
	s.DatasetMaxFileCount = &v
	return s
}

func (s *Project) SetDatasetMaxRelationCount(v int64) *Project {
	s.DatasetMaxRelationCount = &v
	return s
}

func (s *Project) SetDatasetMaxTotalFileSize(v int64) *Project {
	s.DatasetMaxTotalFileSize = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetEngineConcurrency(v int64) *Project {
	s.EngineConcurrency = &v
	return s
}

func (s *Project) SetFileCount(v int64) *Project {
	s.FileCount = &v
	return s
}

func (s *Project) SetProjectMaxDatasetCount(v int64) *Project {
	s.ProjectMaxDatasetCount = &v
	return s
}

func (s *Project) SetProjectName(v string) *Project {
	s.ProjectName = &v
	return s
}

func (s *Project) SetProjectQueriesPerSecond(v int64) *Project {
	s.ProjectQueriesPerSecond = &v
	return s
}

func (s *Project) SetServiceRole(v string) *Project {
	s.ServiceRole = &v
	return s
}

func (s *Project) SetTags(v []*ProjectTags) *Project {
	s.Tags = v
	return s
}

func (s *Project) SetTemplateId(v string) *Project {
	s.TemplateId = &v
	return s
}

func (s *Project) SetTotalFileSize(v int64) *Project {
	s.TotalFileSize = &v
	return s
}

func (s *Project) SetUpdateTime(v string) *Project {
	s.UpdateTime = &v
	return s
}

func (s *Project) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ProjectTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ProjectTags) String() string {
	return dara.Prettify(s)
}

func (s ProjectTags) GoString() string {
	return s.String()
}

func (s *ProjectTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ProjectTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ProjectTags) SetTagKey(v string) *ProjectTags {
	s.TagKey = &v
	return s
}

func (s *ProjectTags) SetTagValue(v string) *ProjectTags {
	s.TagValue = &v
	return s
}

func (s *ProjectTags) Validate() error {
	return dara.Validate(s)
}
