// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *QueryRobotTaskDetailRequest
	GetId() *int64
	SetOwnerId(v int64) *QueryRobotTaskDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryRobotTaskDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRobotTaskDetailRequest
	GetResourceOwnerId() *int64
}

type QueryRobotTaskDetailRequest struct {
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1045001
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryRobotTaskDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRobotTaskDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRobotTaskDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRobotTaskDetailRequest) SetId(v int64) *QueryRobotTaskDetailRequest {
	s.Id = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetOwnerId(v int64) *QueryRobotTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetResourceOwnerId(v int64) *QueryRobotTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
