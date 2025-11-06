// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskInfoHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageCursor(v *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) *QueryTaskInfoHistoryResponseBody
	GetCurrentPageCursor() *QueryTaskInfoHistoryResponseBodyCurrentPageCursor
	SetNextPageCursor(v *QueryTaskInfoHistoryResponseBodyNextPageCursor) *QueryTaskInfoHistoryResponseBody
	GetNextPageCursor() *QueryTaskInfoHistoryResponseBodyNextPageCursor
	SetObjects(v []*QueryTaskInfoHistoryResponseBodyObjects) *QueryTaskInfoHistoryResponseBody
	GetObjects() []*QueryTaskInfoHistoryResponseBodyObjects
	SetPageSize(v int32) *QueryTaskInfoHistoryResponseBody
	GetPageSize() *int32
	SetPrePageCursor(v *QueryTaskInfoHistoryResponseBodyPrePageCursor) *QueryTaskInfoHistoryResponseBody
	GetPrePageCursor() *QueryTaskInfoHistoryResponseBodyPrePageCursor
	SetRequestId(v string) *QueryTaskInfoHistoryResponseBody
	GetRequestId() *string
}

type QueryTaskInfoHistoryResponseBody struct {
	CurrentPageCursor *QueryTaskInfoHistoryResponseBodyCurrentPageCursor `json:"CurrentPageCursor,omitempty" xml:"CurrentPageCursor,omitempty" type:"Struct"`
	NextPageCursor    *QueryTaskInfoHistoryResponseBodyNextPageCursor    `json:"NextPageCursor,omitempty" xml:"NextPageCursor,omitempty" type:"Struct"`
	Objects           []*QueryTaskInfoHistoryResponseBodyObjects         `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageSize      *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePageCursor *QueryTaskInfoHistoryResponseBodyPrePageCursor `json:"PrePageCursor,omitempty" xml:"PrePageCursor,omitempty" type:"Struct"`
	// example:
	//
	// EB3FCCBA-CA1F-4D31-9F34-test
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBody) GetCurrentPageCursor() *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	return s.CurrentPageCursor
}

func (s *QueryTaskInfoHistoryResponseBody) GetNextPageCursor() *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	return s.NextPageCursor
}

func (s *QueryTaskInfoHistoryResponseBody) GetObjects() []*QueryTaskInfoHistoryResponseBodyObjects {
	return s.Objects
}

func (s *QueryTaskInfoHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskInfoHistoryResponseBody) GetPrePageCursor() *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	return s.PrePageCursor
}

func (s *QueryTaskInfoHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskInfoHistoryResponseBody) SetCurrentPageCursor(v *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) *QueryTaskInfoHistoryResponseBody {
	s.CurrentPageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetNextPageCursor(v *QueryTaskInfoHistoryResponseBodyNextPageCursor) *QueryTaskInfoHistoryResponseBody {
	s.NextPageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetObjects(v []*QueryTaskInfoHistoryResponseBodyObjects) *QueryTaskInfoHistoryResponseBody {
	s.Objects = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetPageSize(v int32) *QueryTaskInfoHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetPrePageCursor(v *QueryTaskInfoHistoryResponseBodyPrePageCursor) *QueryTaskInfoHistoryResponseBody {
	s.PrePageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetRequestId(v string) *QueryTaskInfoHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) Validate() error {
	if s.CurrentPageCursor != nil {
		if err := s.CurrentPageCursor.Validate(); err != nil {
			return err
		}
	}
	if s.NextPageCursor != nil {
		if err := s.NextPageCursor.Validate(); err != nil {
			return err
		}
	}
	if s.Objects != nil {
		for _, item := range s.Objects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrePageCursor != nil {
		if err := s.PrePageCursor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskInfoHistoryResponseBodyCurrentPageCursor struct {
	// example:
	//
	// 127.0.0.1
	Clientip *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	// example:
	//
	// 2017-11-01 17:22:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1509528171000
	CreateTimeLong *int64 `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	// example:
	//
	// aa634d3f-927e-4d17-9d2c-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 1
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CHG_DNS
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyCurrentPageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetClientip() *string {
	return s.Clientip
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetCreateTimeLong() *int64 {
	return s.CreateTimeLong
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) Validate() error {
	return dara.Validate(s)
}

type QueryTaskInfoHistoryResponseBodyNextPageCursor struct {
	// example:
	//
	// 127.0.0.1
	Clientip *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	// example:
	//
	// 2017-10-27 13:07:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1509080827000
	CreateTimeLong *int64 `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	// example:
	//
	// 8f112aa1-98be-48c3-82f8-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 15
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CHG_DNS
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyNextPageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyNextPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetClientip() *string {
	return s.Clientip
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetCreateTimeLong() *int64 {
	return s.CreateTimeLong
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) Validate() error {
	return dara.Validate(s)
}

type QueryTaskInfoHistoryResponseBodyObjects struct {
	// example:
	//
	// 127.0.0.1
	Clientip *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	// example:
	//
	// 2017-11-01 17:22:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1509528171000
	CreateTimeLong *int64 `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	// example:
	//
	// aa634d3f-927e-4d17-9d2c-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 1
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CHG_DNS
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyObjects) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetClientip() *string {
	return s.Clientip
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetCreateTimeLong() *int64 {
	return s.CreateTimeLong
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyObjects {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) Validate() error {
	return dara.Validate(s)
}

type QueryTaskInfoHistoryResponseBodyPrePageCursor struct {
	// example:
	//
	// 127.0.0.1
	Clientip *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	// example:
	//
	// 2017-11-01 17:19:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1509527987000
	CreateTimeLong *int64 `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	// example:
	//
	// f9baa3d5-33b9-4c81-8847-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 15
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 3
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CHG_DNS
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyPrePageCursor) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyPrePageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetClientip() *string {
	return s.Clientip
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetCreateTimeLong() *int64 {
	return s.CreateTimeLong
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) Validate() error {
	return dara.Validate(s)
}
