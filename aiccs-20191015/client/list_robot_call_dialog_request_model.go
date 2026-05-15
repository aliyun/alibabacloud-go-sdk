// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotCallDialogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *ListRobotCallDialogRequest
	GetCallId() *string
	SetCreateTime(v string) *ListRobotCallDialogRequest
	GetCreateTime() *string
	SetOwnerId(v int64) *ListRobotCallDialogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListRobotCallDialogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRobotCallDialogRequest
	GetResourceOwnerId() *int64
}

type ListRobotCallDialogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 125165515022^11195613****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-20 00:00:00
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListRobotCallDialogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRobotCallDialogRequest) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogRequest) GetCallId() *string {
	return s.CallId
}

func (s *ListRobotCallDialogRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRobotCallDialogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRobotCallDialogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRobotCallDialogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRobotCallDialogRequest) SetCallId(v string) *ListRobotCallDialogRequest {
	s.CallId = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetCreateTime(v string) *ListRobotCallDialogRequest {
	s.CreateTime = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetOwnerId(v int64) *ListRobotCallDialogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetResourceOwnerAccount(v string) *ListRobotCallDialogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRobotCallDialogRequest) SetResourceOwnerId(v int64) *ListRobotCallDialogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRobotCallDialogRequest) Validate() error {
	return dara.Validate(s)
}
