// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListRobotParamsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListRobotParamsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRobotParamsRequest
	GetResourceOwnerId() *int64
	SetRobotId(v int64) *ListRobotParamsRequest
	GetRobotId() *int64
}

type ListRobotParamsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
}

func (s ListRobotParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRobotParamsRequest) GoString() string {
	return s.String()
}

func (s *ListRobotParamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRobotParamsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRobotParamsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRobotParamsRequest) GetRobotId() *int64 {
	return s.RobotId
}

func (s *ListRobotParamsRequest) SetOwnerId(v int64) *ListRobotParamsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRobotParamsRequest) SetResourceOwnerAccount(v string) *ListRobotParamsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRobotParamsRequest) SetResourceOwnerId(v int64) *ListRobotParamsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRobotParamsRequest) SetRobotId(v int64) *ListRobotParamsRequest {
	s.RobotId = &v
	return s
}

func (s *ListRobotParamsRequest) Validate() error {
	return dara.Validate(s)
}
