// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiccsRobotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListAiccsRobotRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAiccsRobotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAiccsRobotRequest
	GetResourceOwnerId() *int64
	SetRobotName(v string) *ListAiccsRobotRequest
	GetRobotName() *string
}

type ListAiccsRobotRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The robot name. You can view the robot name in the [Script Management](https://aiccs.console.aliyun.com/patter/list) interface.
	//
	// example:
	//
	// 测试机器人
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
}

func (s ListAiccsRobotRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiccsRobotRequest) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAiccsRobotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAiccsRobotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAiccsRobotRequest) GetRobotName() *string {
	return s.RobotName
}

func (s *ListAiccsRobotRequest) SetOwnerId(v int64) *ListAiccsRobotRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAiccsRobotRequest) SetResourceOwnerAccount(v string) *ListAiccsRobotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAiccsRobotRequest) SetResourceOwnerId(v int64) *ListAiccsRobotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAiccsRobotRequest) SetRobotName(v string) *ListAiccsRobotRequest {
	s.RobotName = &v
	return s
}

func (s *ListAiccsRobotRequest) Validate() error {
	return dara.Validate(s)
}
