// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListImportTasksResponseBodyItems) *ListImportTasksResponseBody
	GetItems() []*ListImportTasksResponseBodyItems
	SetMaxResults(v int32) *ListImportTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListImportTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImportTasksResponseBody
	GetRequestId() *string
}

type ListImportTasksResponseBody struct {
	Items []*ListImportTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListImportTasksResponseBody) GetItems() []*ListImportTasksResponseBodyItems {
	return s.Items
}

func (s *ListImportTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImportTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImportTasksResponseBody) SetItems(v []*ListImportTasksResponseBodyItems) *ListImportTasksResponseBody {
	s.Items = v
	return s
}

func (s *ListImportTasksResponseBody) SetMaxResults(v int32) *ListImportTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImportTasksResponseBody) SetNextToken(v string) *ListImportTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImportTasksResponseBody) SetRequestId(v string) *ListImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImportTasksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImportTasksResponseBodyItems struct {
	// example:
	//
	// 2018-05-30T14:30:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 5.7
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// Importing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rm-bp*****
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// example:
	//
	// 342900000
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 362c6c7a-4d20-4eac-898c-1495ceab374c
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// import
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListImportTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListImportTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListImportTasksResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListImportTasksResponseBodyItems) GetDbVersion() *string {
	return s.DbVersion
}

func (s *ListImportTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListImportTasksResponseBodyItems) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *ListImportTasksResponseBodyItems) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListImportTasksResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *ListImportTasksResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *ListImportTasksResponseBodyItems) SetCreatedTime(v string) *ListImportTasksResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetDbVersion(v string) *ListImportTasksResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetStatus(v string) *ListImportTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetTargetInstanceName(v string) *ListImportTasksResponseBodyItems {
	s.TargetInstanceName = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetTaskId(v int64) *ListImportTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetTaskName(v string) *ListImportTasksResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) SetTaskType(v string) *ListImportTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *ListImportTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
