// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoLockTime(v string) *UnlockUsersRequest
	GetAutoLockTime() *string
	SetUsers(v []*string) *UnlockUsersRequest
	GetUsers() []*string
}

type UnlockUsersRequest struct {
	// The date on which the convenience users are automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The usernames of the convenience users that you want to unlock.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s UnlockUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersRequest) GoString() string {
	return s.String()
}

func (s *UnlockUsersRequest) GetAutoLockTime() *string {
	return s.AutoLockTime
}

func (s *UnlockUsersRequest) GetUsers() []*string {
	return s.Users
}

func (s *UnlockUsersRequest) SetAutoLockTime(v string) *UnlockUsersRequest {
	s.AutoLockTime = &v
	return s
}

func (s *UnlockUsersRequest) SetUsers(v []*string) *UnlockUsersRequest {
	s.Users = v
	return s
}

func (s *UnlockUsersRequest) Validate() error {
	return dara.Validate(s)
}
