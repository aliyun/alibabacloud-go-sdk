// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListRobotNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListRobotNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRobotNodeRequest
	GetResourceOwnerId() *int64
	SetRobotId(v int64) *ListRobotNodeRequest
	GetRobotId() *int64
}

type ListRobotNodeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The robot ID, which is the script ID.
	//
	// You can obtain it in the <b>outbound robot</b> > <b>Script Management</b> section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// 100002674****
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
}

func (s ListRobotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRobotNodeRequest) GoString() string {
	return s.String()
}

func (s *ListRobotNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRobotNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRobotNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRobotNodeRequest) GetRobotId() *int64 {
	return s.RobotId
}

func (s *ListRobotNodeRequest) SetOwnerId(v int64) *ListRobotNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRobotNodeRequest) SetResourceOwnerAccount(v string) *ListRobotNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRobotNodeRequest) SetResourceOwnerId(v int64) *ListRobotNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRobotNodeRequest) SetRobotId(v int64) *ListRobotNodeRequest {
	s.RobotId = &v
	return s
}

func (s *ListRobotNodeRequest) Validate() error {
	return dara.Validate(s)
}
