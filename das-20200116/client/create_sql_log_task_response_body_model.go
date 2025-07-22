// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlLogTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSqlLogTaskResponseBody
	GetCode() *string
	SetData(v *CreateSqlLogTaskResponseBodyData) *CreateSqlLogTaskResponseBody
	GetData() *CreateSqlLogTaskResponseBodyData
	SetMessage(v string) *CreateSqlLogTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSqlLogTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSqlLogTaskResponseBody
	GetSuccess() *string
}

type CreateSqlLogTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateSqlLogTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, error information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83D9D59B-057A-54A9-BFF9-CF2B42F05645
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSqlLogTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlLogTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlLogTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSqlLogTaskResponseBody) GetData() *CreateSqlLogTaskResponseBodyData {
	return s.Data
}

func (s *CreateSqlLogTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSqlLogTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlLogTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSqlLogTaskResponseBody) SetCode(v string) *CreateSqlLogTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSqlLogTaskResponseBody) SetData(v *CreateSqlLogTaskResponseBodyData) *CreateSqlLogTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateSqlLogTaskResponseBody) SetMessage(v string) *CreateSqlLogTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSqlLogTaskResponseBody) SetRequestId(v string) *CreateSqlLogTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlLogTaskResponseBody) SetSuccess(v string) *CreateSqlLogTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSqlLogTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSqlLogTaskResponseBodyData struct {
	// The time when the task was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1681363254423
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// pc-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Export_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1596177993000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- **INIT**: The task is to be scheduled.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FAILED**: The task failed.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **COMPLETED**: The task is complete.
	//
	// >  You can view the result of a task that is in the **COMPLETED*	- state.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 54f8041743ca3a9ac5cb9342d050527c
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSqlLogTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlLogTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSqlLogTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateSqlLogTaskResponseBodyData) GetEnd() *int64 {
	return s.End
}

func (s *CreateSqlLogTaskResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSqlLogTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateSqlLogTaskResponseBodyData) GetStart() *int64 {
	return s.Start
}

func (s *CreateSqlLogTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateSqlLogTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSqlLogTaskResponseBodyData) SetCreateTime(v int64) *CreateSqlLogTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetEnd(v int64) *CreateSqlLogTaskResponseBodyData {
	s.End = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetInstanceId(v string) *CreateSqlLogTaskResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetName(v string) *CreateSqlLogTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetStart(v int64) *CreateSqlLogTaskResponseBodyData {
	s.Start = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetStatus(v string) *CreateSqlLogTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) SetTaskId(v string) *CreateSqlLogTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateSqlLogTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
