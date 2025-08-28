// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskCallDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *QueryRobotTaskCallDetailRequest
	GetCallee() *string
	SetOwnerId(v int64) *QueryRobotTaskCallDetailRequest
	GetOwnerId() *int64
	SetQueryDate(v int64) *QueryRobotTaskCallDetailRequest
	GetQueryDate() *int64
	SetResourceOwnerAccount(v string) *QueryRobotTaskCallDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotTaskCallDetailRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *QueryRobotTaskCallDetailRequest
	GetTaskId() *int64
}

type QueryRobotTaskCallDetailRequest struct {
	// The called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130****0000
	Callee  *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timestamp of the time at which the call details you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-21 08:23:21
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](https://help.aliyun.com/document_detail/393531.html) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRobotTaskCallDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskCallDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailRequest) GetCallee() *string {
	return s.Callee
}

func (s *QueryRobotTaskCallDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotTaskCallDetailRequest) GetQueryDate() *int64 {
	return s.QueryDate
}

func (s *QueryRobotTaskCallDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotTaskCallDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotTaskCallDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryRobotTaskCallDetailRequest) SetCallee(v string) *QueryRobotTaskCallDetailRequest {
	s.Callee = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetOwnerId(v int64) *QueryRobotTaskCallDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetQueryDate(v int64) *QueryRobotTaskCallDetailRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskCallDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetResourceOwnerId(v int64) *QueryRobotTaskCallDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetTaskId(v int64) *QueryRobotTaskCallDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) Validate() error {
	return dara.Validate(s)
}
