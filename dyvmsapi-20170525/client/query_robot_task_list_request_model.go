// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryRobotTaskListRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryRobotTaskListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryRobotTaskListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryRobotTaskListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotTaskListRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *QueryRobotTaskListRequest
	GetStatus() *string
	SetTaskName(v string) *QueryRobotTaskListRequest
	GetTaskName() *string
	SetTime(v string) *QueryRobotTaskListRequest
	GetTime() *string
}

type QueryRobotTaskListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **READY**: The task is ready to start.
	//
	// 	- **DISPATCH**: The task is being parsed.
	//
	// 	- **EXCUTING**: The task is being executed.
	//
	// 	- **MANUAL_STOP**: The task is manually suspended.
	//
	// 	- **SYSTEM_STOP**: The task is suspended by the system.
	//
	// 	- **ARREARS_STOP**: The task is suspended due to overdue payments.
	//
	// 	- **CANCEL**: The task is manually canceled.
	//
	// 	- **SYSTEM_CANCEL**: The task is canceled by the system.
	//
	// 	- **FINISH**: The task is complete.
	//
	// example:
	//
	// EXCUTING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// Test Template
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The date when the task is created, in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2019-06-14
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryRobotTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotTaskListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryRobotTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRobotTaskListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotTaskListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotTaskListRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryRobotTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryRobotTaskListRequest) GetTime() *string {
	return s.Time
}

func (s *QueryRobotTaskListRequest) SetOwnerId(v int64) *QueryRobotTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetPageNo(v int32) *QueryRobotTaskListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetPageSize(v int32) *QueryRobotTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetResourceOwnerId(v int64) *QueryRobotTaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetStatus(v string) *QueryRobotTaskListRequest {
	s.Status = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTaskName(v string) *QueryRobotTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTime(v string) *QueryRobotTaskListRequest {
	s.Time = &v
	return s
}

func (s *QueryRobotTaskListRequest) Validate() error {
	return dara.Validate(s)
}
