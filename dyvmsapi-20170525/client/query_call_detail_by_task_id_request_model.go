// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *QueryCallDetailByTaskIdRequest
	GetCallee() *string
	SetOwnerId(v int64) *QueryCallDetailByTaskIdRequest
	GetOwnerId() *int64
	SetQueryDate(v int64) *QueryCallDetailByTaskIdRequest
	GetQueryDate() *int64
	SetResourceOwnerAccount(v string) *QueryCallDetailByTaskIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCallDetailByTaskIdRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *QueryCallDetailByTaskIdRequest
	GetTaskId() *string
}

type QueryCallDetailByTaskIdRequest struct {
	// The called number. You can view the outbound call records of only one called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Callee  *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the outbound robocall task. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-09 09:02:03
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the outbound robocall task. The task ID is returned after the outbound robocall task is successfully delivered. You can view the task ID on the [Task Management](https://dyvms.console.aliyun.com/job/list) page of the Voice Messaging Service console, or call the **BatchRobotSmartCall*	- operation to obtain the **task ID**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4001112222
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryCallDetailByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdRequest) GetCallee() *string {
	return s.Callee
}

func (s *QueryCallDetailByTaskIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCallDetailByTaskIdRequest) GetQueryDate() *int64 {
	return s.QueryDate
}

func (s *QueryCallDetailByTaskIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCallDetailByTaskIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCallDetailByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryCallDetailByTaskIdRequest) SetCallee(v string) *QueryCallDetailByTaskIdRequest {
	s.Callee = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetOwnerId(v int64) *QueryCallDetailByTaskIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetQueryDate(v int64) *QueryCallDetailByTaskIdRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetResourceOwnerAccount(v string) *QueryCallDetailByTaskIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetResourceOwnerId(v int64) *QueryCallDetailByTaskIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetTaskId(v string) *QueryCallDetailByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
