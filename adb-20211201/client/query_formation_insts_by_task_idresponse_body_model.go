// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationInstsByTaskIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFormationInstsByTaskIDResponseBody
	GetCode() *string
	SetData(v []*QueryFormationInstsByTaskIDResponseBodyData) *QueryFormationInstsByTaskIDResponseBody
	GetData() []*QueryFormationInstsByTaskIDResponseBodyData
	SetHttpStatusCode(v int32) *QueryFormationInstsByTaskIDResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []map[string]interface{}) *QueryFormationInstsByTaskIDResponseBody
	GetItems() []map[string]interface{}
	SetMessage(v string) *QueryFormationInstsByTaskIDResponseBody
	GetMessage() *string
	SetPageNumber(v string) *QueryFormationInstsByTaskIDResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *QueryFormationInstsByTaskIDResponseBody
	GetPageSize() *string
	SetRequestId(v string) *QueryFormationInstsByTaskIDResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFormationInstsByTaskIDResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *QueryFormationInstsByTaskIDResponseBody
	GetTotalCount() *string
}

type QueryFormationInstsByTaskIDResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned task list.
	Data []*QueryFormationInstsByTaskIDResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The task list.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The response message. OK is returned if the request was successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryFormationInstsByTaskIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationInstsByTaskIDResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetData() []*QueryFormationInstsByTaskIDResponseBodyData {
	return s.Data
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFormationInstsByTaskIDResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetCode(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetData(v []*QueryFormationInstsByTaskIDResponseBodyData) *QueryFormationInstsByTaskIDResponseBody {
	s.Data = v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetHttpStatusCode(v int32) *QueryFormationInstsByTaskIDResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetItems(v []map[string]interface{}) *QueryFormationInstsByTaskIDResponseBody {
	s.Items = v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetMessage(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetPageNumber(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetPageSize(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetRequestId(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetSuccess(v bool) *QueryFormationInstsByTaskIDResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) SetTotalCount(v string) *QueryFormationInstsByTaskIDResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFormationInstsByTaskIDResponseBodyData struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2023-05-15T07:24:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The execution duration of the most recent task.
	//
	// example:
	//
	// 8
	LastTaskInstCostTime *string `json:"LastTaskInstCostTime,omitempty" xml:"LastTaskInstCostTime,omitempty"`
	// The instance ID of the most recent task.
	//
	// example:
	//
	// 1223
	LastTaskInstID *string `json:"LastTaskInstID,omitempty" xml:"LastTaskInstID,omitempty"`
	// The error message of the most recent task.
	//
	// example:
	//
	// error
	LastTaskInstMessage *string `json:"LastTaskInstMessage,omitempty" xml:"LastTaskInstMessage,omitempty"`
	// The instance status of the most recent node.
	//
	// example:
	//
	// FINISH
	LastTaskInstState *string `json:"LastTaskInstState,omitempty" xml:"LastTaskInstState,omitempty"`
	// The task status.
	//
	// example:
	//
	// NORMAL
	ScheduleState *string `json:"ScheduleState,omitempty" xml:"ScheduleState,omitempty"`
	// The schema ID assigned to the instance by the system.
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The task source. Valid values:
	//
	// - **system**: system.
	//
	// - **custom**: custom.
	//
	// example:
	//
	// shareScreen
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The synchronization time, accurate to milliseconds. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-01-20t10:01:01z
	SyncTime *string `json:"SyncTime,omitempty" xml:"SyncTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// doc_test_daily
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type.
	//
	// example:
	//
	// Update\\"\\"
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryFormationInstsByTaskIDResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationInstsByTaskIDResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetLastTaskInstCostTime() *string {
	return s.LastTaskInstCostTime
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetLastTaskInstID() *string {
	return s.LastTaskInstID
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetLastTaskInstMessage() *string {
	return s.LastTaskInstMessage
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetLastTaskInstState() *string {
	return s.LastTaskInstState
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetSyncTime() *string {
	return s.SyncTime
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetCreateTime(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetLastTaskInstCostTime(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.LastTaskInstCostTime = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetLastTaskInstID(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.LastTaskInstID = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetLastTaskInstMessage(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.LastTaskInstMessage = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetLastTaskInstState(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.LastTaskInstState = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetScheduleState(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.ScheduleState = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetSchema(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.Schema = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetSourceType(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetSyncTime(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.SyncTime = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetTaskId(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetTaskName(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) SetTaskType(v string) *QueryFormationInstsByTaskIDResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponseBodyData) Validate() error {
	return dara.Validate(s)
}
