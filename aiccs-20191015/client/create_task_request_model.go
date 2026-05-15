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
	CallString *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JSON
	CallStringType *string `json:"CallStringType,omitempty" xml:"CallStringType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0571****5678,0571****5679
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// example:
	//
	// 1
	RetryFlag *int32 `json:"RetryFlag,omitempty" xml:"RetryFlag,omitempty"`
	// example:
	//
	// 2
	RetryInterval *int32 `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	// example:
	//
	// 200010,200011
	RetryStatusCode *string `json:"RetryStatusCode,omitempty" xml:"RetryStatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RobotId *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	SeatCount *string `json:"SeatCount,omitempty" xml:"SeatCount,omitempty"`
	StartNow  *bool   `json:"StartNow,omitempty" xml:"StartNow,omitempty"`
	// This parameter is required.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkDay *string `json:"WorkDay,omitempty" xml:"WorkDay,omitempty"`
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
