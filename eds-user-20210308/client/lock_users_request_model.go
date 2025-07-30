// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogoutSession(v bool) *LockUsersRequest
	GetLogoutSession() *bool
	SetUsers(v []*string) *LockUsersRequest
	GetUsers() []*string
}

type LockUsersRequest struct {
	LogoutSession *bool `json:"LogoutSession,omitempty" xml:"LogoutSession,omitempty"`
	// The usernames of the convenience users that you want to lock.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s LockUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s LockUsersRequest) GoString() string {
	return s.String()
}

func (s *LockUsersRequest) GetLogoutSession() *bool {
	return s.LogoutSession
}

func (s *LockUsersRequest) GetUsers() []*string {
	return s.Users
}

func (s *LockUsersRequest) SetLogoutSession(v bool) *LockUsersRequest {
	s.LogoutSession = &v
	return s
}

func (s *LockUsersRequest) SetUsers(v []*string) *LockUsersRequest {
	s.Users = v
	return s
}

func (s *LockUsersRequest) Validate() error {
	return dara.Validate(s)
}
