// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListTasksRequest
	GetGroupId() *string
	SetKeyword(v string) *ListTasksRequest
	GetKeyword() *string
	SetKmsKeyId(v string) *ListTasksRequest
	GetKmsKeyId() *string
	SetModuleId(v string) *ListTasksRequest
	GetModuleId() *string
	SetPageNumber(v int32) *ListTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListTasksRequest
	GetProjectId() *string
	SetStatus(v string) *ListTasksRequest
	GetStatus() *string
	SetTag(v []*ListTasksRequestTag) *ListTasksRequest
	GetTag() []*ListTasksRequestTag
	SetTaskId(v string) *ListTasksRequest
	GetTaskId() *string
}

type ListTasksRequest struct {
	// The group ID.
	//
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The keyword for fuzzy search by task ID or task name.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 21a90f5d-a469-4ac4-a8ea-f6e1e7470e6f
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The module ID.
	//
	// example:
	//
	// mod-1525e992f1b62139d1c437d64ae
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The job status. Valid values:
	//
	// - Planning: The job is in the Plan execution phase.
	//
	// - Planned: The job has completed the Plan execution.
	//
	// - PlannedAndFinished: After the Plan execution is completed, no diff is found, and the job enters the final state.
	//
	// - Applying: The job is in the Apply execution phase.
	//
	// - Applied: The job has completed the Apply execution.
	//
	// - Errored: The job execution encountered errors and entered the final state.
	//
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of task tags.
	Tag []*ListTasksRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTasksRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *ListTasksRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksRequest) GetTag() []*ListTasksRequestTag {
	return s.Tag
}

func (s *ListTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksRequest) SetGroupId(v string) *ListTasksRequest {
	s.GroupId = &v
	return s
}

func (s *ListTasksRequest) SetKeyword(v string) *ListTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListTasksRequest) SetKmsKeyId(v string) *ListTasksRequest {
	s.KmsKeyId = &v
	return s
}

func (s *ListTasksRequest) SetModuleId(v string) *ListTasksRequest {
	s.ModuleId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v string) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTag(v []*ListTasksRequestTag) *ListTasksRequest {
	s.Tag = v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksRequestTag struct {
	// The tag key of the task.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value of the task.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTasksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequestTag) GoString() string {
	return s.String()
}

func (s *ListTasksRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTasksRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTasksRequestTag) SetTagKey(v string) *ListTasksRequestTag {
	s.TagKey = &v
	return s
}

func (s *ListTasksRequestTag) SetTagValue(v string) *ListTasksRequestTag {
	s.TagValue = &v
	return s
}

func (s *ListTasksRequestTag) Validate() error {
	return dara.Validate(s)
}
