// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallString(v string) *CreateTaskRequest
	GetCallString() *string
	SetCallStringType(v string) *CreateTaskRequest
	GetCallStringType() *string
	SetCaller(v string) *CreateTaskRequest
	GetCaller() *string
	SetOwnerId(v int64) *CreateTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTaskRequest
	GetResourceOwnerId() *int64
	SetRetryCount(v int32) *CreateTaskRequest
	GetRetryCount() *int32
	SetRetryFlag(v int32) *CreateTaskRequest
	GetRetryFlag() *int32
	SetRetryInterval(v int32) *CreateTaskRequest
	GetRetryInterval() *int32
	SetRetryStatusCode(v string) *CreateTaskRequest
	GetRetryStatusCode() *string
	SetRobotId(v string) *CreateTaskRequest
	GetRobotId() *string
	SetSeatCount(v string) *CreateTaskRequest
	GetSeatCount() *string
	SetStartNow(v bool) *CreateTaskRequest
	GetStartNow() *bool
	SetTaskName(v string) *CreateTaskRequest
	GetTaskName() *string
	SetWorkDay(v string) *CreateTaskRequest
	GetWorkDay() *string
	SetWorkTimeList(v string) *CreateTaskRequest
	GetWorkTimeList() *string
}

type CreateTaskRequest struct {
	// Call string (callee information and parameter list). Valid values:
	//
	// - **LIST**: `05715678****,05715679****`
	//
	// - **JSON**: `{"ParamNames":["name","age"],"CalleeList":[{"Callee":"1810000****","Params":["Zhang San","20"]},{"Callee":"1810001****","Params":["Li Si","21"]}]}`. In this example, ParamNames represents the List of Parameter Names; Params represents the List of parameter values.
	//
	// > - The order of the Parameter Name List and the parameter value List must correspond.
	//
	// - A maximum of 1 000 callee numbers is allowed.
	//
	// example:
	//
	// {"ParamNames":["name","age"],"CalleeList":[{"Callee":"1810000****","Params":["张三","20"]},{"Callee":"1810001****","Params":["李四","21"]}]}
	CallString *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
	// Call string type. Valid values:
	//
	// - **LIST*	-
	//
	// - **JSON**
	//
	// This parameter is required.
	//
	// example:
	//
	// JSON
	CallStringType *string `json:"CallStringType,omitempty" xml:"CallStringType,omitempty"`
	// Outbound caller number.
	//
	// > The number must be a purchased number. Separate multiple numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571****5678,0571****5679
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Retry Count.
	//
	// example:
	//
	// 2
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// Whether to enable automatic retry. Valid values:
	//
	// - **1**: Retry.
	//
	// - **0**: No retry.
	//
	// example:
	//
	// 1
	RetryFlag *int32 `json:"RetryFlag,omitempty" xml:"RetryFlag,omitempty"`
	// Retry interval. Unit: minute. Must be greater than 1.
	//
	// example:
	//
	// 2
	RetryInterval *int32 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// Call statuses that require redialing. Separate multiple statuses with commas (,). Valid values:
	//
	// - **200010**: Power off
	//
	// - **200011**: Service suspended
	//
	// - **200002**: Busy
	//
	// - **200012**: Call failed
	//
	// - **200005**: Unable to connect
	//
	// - **200003**: No acknowledgement
	//
	// example:
	//
	// 200010,200011
	RetryStatusCode *string `json:"RetryStatusCode,omitempty" xml:"RetryStatusCode,omitempty"`
	// ID of the specified robot (script ID), indicating which robot script to use for initiating calls.
	//
	// You can obtain the script ID on the [Script Management](https://aiccs.console.aliyun.com/patter/list) page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	RobotId *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// Concurrency (number of agents).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	SeatCount *string `json:"SeatCount,omitempty" xml:"SeatCount,omitempty"`
	// Indicates whether to start immediately.
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// 是
	StartNow *bool `json:"StartNow,omitempty" xml:"StartNow,omitempty"`
	// Task Name. Supports Chinese and English characters. Length: 0 to 30 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试任务
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Work day. Valid values:
	//
	// - **1**: Monday.
	//
	// - **2**: Tuesday.
	//
	// - **3**: Wednesday.
	//
	// - **4**: Thursday.
	//
	// - **5**: Friday.
	//
	// - **6**: Saturday.
	//
	// - **7**: Sunday.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkDay *string `json:"WorkDay,omitempty" xml:"WorkDay,omitempty"`
	// List of working hours (accurate to the minute).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00-12:00,13:00-14:00
	WorkTimeList *string `json:"WorkTimeList,omitempty" xml:"WorkTimeList,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetCallString() *string {
	return s.CallString
}

func (s *CreateTaskRequest) GetCallStringType() *string {
	return s.CallStringType
}

func (s *CreateTaskRequest) GetCaller() *string {
	return s.Caller
}

func (s *CreateTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTaskRequest) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *CreateTaskRequest) GetRetryFlag() *int32 {
	return s.RetryFlag
}

func (s *CreateTaskRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *CreateTaskRequest) GetRetryStatusCode() *string {
	return s.RetryStatusCode
}

func (s *CreateTaskRequest) GetRobotId() *string {
	return s.RobotId
}

func (s *CreateTaskRequest) GetSeatCount() *string {
	return s.SeatCount
}

func (s *CreateTaskRequest) GetStartNow() *bool {
	return s.StartNow
}

func (s *CreateTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskRequest) GetWorkDay() *string {
	return s.WorkDay
}

func (s *CreateTaskRequest) GetWorkTimeList() *string {
	return s.WorkTimeList
}

func (s *CreateTaskRequest) SetCallString(v string) *CreateTaskRequest {
	s.CallString = &v
	return s
}

func (s *CreateTaskRequest) SetCallStringType(v string) *CreateTaskRequest {
	s.CallStringType = &v
	return s
}

func (s *CreateTaskRequest) SetCaller(v string) *CreateTaskRequest {
	s.Caller = &v
	return s
}

func (s *CreateTaskRequest) SetOwnerId(v int64) *CreateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTaskRequest) SetResourceOwnerAccount(v string) *CreateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTaskRequest) SetResourceOwnerId(v int64) *CreateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTaskRequest) SetRetryCount(v int32) *CreateTaskRequest {
	s.RetryCount = &v
	return s
}

func (s *CreateTaskRequest) SetRetryFlag(v int32) *CreateTaskRequest {
	s.RetryFlag = &v
	return s
}

func (s *CreateTaskRequest) SetRetryInterval(v int32) *CreateTaskRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateTaskRequest) SetRetryStatusCode(v string) *CreateTaskRequest {
	s.RetryStatusCode = &v
	return s
}

func (s *CreateTaskRequest) SetRobotId(v string) *CreateTaskRequest {
	s.RobotId = &v
	return s
}

func (s *CreateTaskRequest) SetSeatCount(v string) *CreateTaskRequest {
	s.SeatCount = &v
	return s
}

func (s *CreateTaskRequest) SetStartNow(v bool) *CreateTaskRequest {
	s.StartNow = &v
	return s
}

func (s *CreateTaskRequest) SetTaskName(v string) *CreateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskRequest) SetWorkDay(v string) *CreateTaskRequest {
	s.WorkDay = &v
	return s
}

func (s *CreateTaskRequest) SetWorkTimeList(v string) *CreateTaskRequest {
	s.WorkTimeList = &v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}
