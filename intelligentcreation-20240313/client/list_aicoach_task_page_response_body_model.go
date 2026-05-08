// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAICoachTaskPageResponseBody
	GetRequestId() *string
	SetTaskList(v []*ListAICoachTaskPageResponseBodyTaskList) *ListAICoachTaskPageResponseBody
	GetTaskList() []*ListAICoachTaskPageResponseBodyTaskList
	SetTotal(v int64) *ListAICoachTaskPageResponseBody
	GetTotal() *int64
}

type ListAICoachTaskPageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D7F2B74F-63F2-5DD6-95E4-F408EAD6617E
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskList  []*ListAICoachTaskPageResponseBodyTaskList `json:"taskList,omitempty" xml:"taskList,omitempty" type:"Repeated"`
	Total     *int64                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAICoachTaskPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAICoachTaskPageResponseBody) GetTaskList() []*ListAICoachTaskPageResponseBodyTaskList {
	return s.TaskList
}

func (s *ListAICoachTaskPageResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAICoachTaskPageResponseBody) SetRequestId(v string) *ListAICoachTaskPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICoachTaskPageResponseBody) SetTaskList(v []*ListAICoachTaskPageResponseBodyTaskList) *ListAICoachTaskPageResponseBody {
	s.TaskList = v
	return s
}

func (s *ListAICoachTaskPageResponseBody) SetTotal(v int64) *ListAICoachTaskPageResponseBody {
	s.Total = &v
	return s
}

func (s *ListAICoachTaskPageResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICoachTaskPageResponseBodyTaskList struct {
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	GmtCreate  *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 222
	StudentId *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	// example:
	//
	// 11111111111
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAICoachTaskPageResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskPageResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageResponseBodyTaskList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListAICoachTaskPageResponseBodyTaskList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAICoachTaskPageResponseBodyTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListAICoachTaskPageResponseBodyTaskList) GetStudentId() *string {
	return s.StudentId
}

func (s *ListAICoachTaskPageResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetFinishTime(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.FinishTime = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetGmtCreate(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.GmtCreate = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetStatus(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetStudentId(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.StudentId = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) SetTaskId(v string) *ListAICoachTaskPageResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListAICoachTaskPageResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
