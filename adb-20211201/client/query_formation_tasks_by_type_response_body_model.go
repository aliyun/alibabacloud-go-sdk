// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTasksByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFormationTasksByTypeResponseBody
	GetCode() *string
	SetData(v []*QueryFormationTasksByTypeResponseBodyData) *QueryFormationTasksByTypeResponseBody
	GetData() []*QueryFormationTasksByTypeResponseBodyData
	SetHttpStatusCode(v int32) *QueryFormationTasksByTypeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryFormationTasksByTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryFormationTasksByTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFormationTasksByTypeResponseBody
	GetSuccess() *bool
}

type QueryFormationTasksByTypeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task list.
	Data []*QueryFormationTasksByTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message. OK is returned if the call was successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFormationTasksByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTasksByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormationTasksByTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFormationTasksByTypeResponseBody) GetData() []*QueryFormationTasksByTypeResponseBodyData {
	return s.Data
}

func (s *QueryFormationTasksByTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryFormationTasksByTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFormationTasksByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFormationTasksByTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFormationTasksByTypeResponseBody) SetCode(v string) *QueryFormationTasksByTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) SetData(v []*QueryFormationTasksByTypeResponseBodyData) *QueryFormationTasksByTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) SetHttpStatusCode(v int32) *QueryFormationTasksByTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) SetMessage(v string) *QueryFormationTasksByTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) SetRequestId(v string) *QueryFormationTasksByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) SetSuccess(v bool) *QueryFormationTasksByTypeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBody) Validate() error {
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

type QueryFormationTasksByTypeResponseBodyData struct {
	// The creation time.
	//
	// example:
	//
	// 2026-07-08 17:05:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The execution duration of the most recent task.
	//
	// example:
	//
	// 2
	LastTaskInstCostTime *string `json:"LastTaskInstCostTime,omitempty" xml:"LastTaskInstCostTime,omitempty"`
	// The instance ID of the most recent task.
	//
	// example:
	//
	// 46872
	LastTaskInstID *string `json:"LastTaskInstID,omitempty" xml:"LastTaskInstID,omitempty"`
	// The error message of the most recent task.
	//
	// example:
	//
	// xxx
	LastTaskInstMessage *string `json:"LastTaskInstMessage,omitempty" xml:"LastTaskInstMessage,omitempty"`
	// The instance status of the most recent node.
	//
	// example:
	//
	// SUCCESS
	LastTaskInstState *string `json:"LastTaskInstState,omitempty" xml:"LastTaskInstState,omitempty"`
	// The scheduling status.
	//
	// example:
	//
	// STOP
	ScheduleState *string `json:"ScheduleState,omitempty" xml:"ScheduleState,omitempty"`
	// The database name.
	//
	// example:
	//
	// sales_db
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The source type.
	//
	// example:
	//
	// OSSWAREHOUSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The scheduling frequency.
	//
	// example:
	//
	// {\\"cron\\":\\"\\",\\"type\\":\\"run_on_demand\\"}
	SyncTime *string `json:"SyncTime,omitempty" xml:"SyncTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type.
	//
	// example:
	//
	// CRAWLER
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryFormationTasksByTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTasksByTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetLastTaskInstCostTime() *string {
	return s.LastTaskInstCostTime
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetLastTaskInstID() *string {
	return s.LastTaskInstID
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetLastTaskInstMessage() *string {
	return s.LastTaskInstMessage
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetLastTaskInstState() *string {
	return s.LastTaskInstState
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetScheduleState() *string {
	return s.ScheduleState
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetSyncTime() *string {
	return s.SyncTime
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryFormationTasksByTypeResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetCreateTime(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetLastTaskInstCostTime(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.LastTaskInstCostTime = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetLastTaskInstID(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.LastTaskInstID = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetLastTaskInstMessage(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.LastTaskInstMessage = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetLastTaskInstState(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.LastTaskInstState = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetScheduleState(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.ScheduleState = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetSchema(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.Schema = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetSourceType(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetSyncTime(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.SyncTime = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetTaskId(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetTaskName(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) SetTaskType(v string) *QueryFormationTasksByTypeResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *QueryFormationTasksByTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
