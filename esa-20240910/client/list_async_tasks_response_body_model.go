// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTasks(v []*ListAsyncTasksResponseBodyAsyncTasks) *ListAsyncTasksResponseBody
	GetAsyncTasks() []*ListAsyncTasksResponseBodyAsyncTasks
	SetPageNumber(v int32) *ListAsyncTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAsyncTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAsyncTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAsyncTasksResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListAsyncTasksResponseBody
	GetTotalPage() *int32
}

type ListAsyncTasksResponseBody struct {
	AsyncTasks []*ListAsyncTasksResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListAsyncTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBody) GetAsyncTasks() []*ListAsyncTasksResponseBodyAsyncTasks {
	return s.AsyncTasks
}

func (s *ListAsyncTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAsyncTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAsyncTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsyncTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAsyncTasksResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListAsyncTasksResponseBody) SetAsyncTasks(v []*ListAsyncTasksResponseBodyAsyncTasks) *ListAsyncTasksResponseBody {
	s.AsyncTasks = v
	return s
}

func (s *ListAsyncTasksResponseBody) SetPageNumber(v int32) *ListAsyncTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetPageSize(v int32) *ListAsyncTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetRequestId(v string) *ListAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetTotalCount(v int32) *ListAsyncTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetTotalPage(v int32) *ListAsyncTasksResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListAsyncTasksResponseBody) Validate() error {
	if s.AsyncTasks != nil {
		for _, item := range s.AsyncTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAsyncTasksResponseBodyAsyncTasks struct {
	// example:
	//
	// 2024-09-19 09:36:46
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-09-19 09:37:04
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 4081****752512
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 4081****752512
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// www.example.com
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// site
	ResourceType    *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TaskDescription map[string]*string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// example:
	//
	// linke-quality-sign
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	// example:
	//
	// success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// free_cert
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 1077***12880
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAsyncTasksResponseBodyAsyncTasks) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBodyAsyncTasks) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetId() *int64 {
	return s.Id
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetTaskDescription() map[string]*string {
	return s.TaskDescription
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetTaskKey() *string {
	return s.TaskKey
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) GetUserId() *int64 {
	return s.UserId
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetGmtCreate(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.GmtCreate = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetGmtModified(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.GmtModified = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetId(v int64) *ListAsyncTasksResponseBodyAsyncTasks {
	s.Id = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetResourceId(v int64) *ListAsyncTasksResponseBodyAsyncTasks {
	s.ResourceId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetResourceName(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.ResourceName = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetResourceType(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.ResourceType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetTaskDescription(v map[string]*string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.TaskDescription = v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetTaskKey(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.TaskKey = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetTaskStatus(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetTaskType(v string) *ListAsyncTasksResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) SetUserId(v int64) *ListAsyncTasksResponseBodyAsyncTasks {
	s.UserId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyAsyncTasks) Validate() error {
	return dara.Validate(s)
}
