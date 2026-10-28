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
	SetBusinessChannel(v string) *UnlockUsersRequest
	GetBusinessChannel() *string
	SetUsers(v []*string) *UnlockUsersRequest
	GetUsers() []*string
}

type UnlockUsersRequest struct {
	// The date when the convenience user is automatically locked. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The list of usernames of the convenience users to unlock.
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

func (s *UnlockUsersRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *UnlockUsersRequest) GetUsers() []*string {
	return s.Users
}

func (s *UnlockUsersRequest) SetAutoLockTime(v string) *UnlockUsersRequest {
	s.AutoLockTime = &v
	return s
}

func (s *UnlockUsersRequest) SetBusinessChannel(v string) *UnlockUsersRequest {
	s.BusinessChannel = &v
	return s
}

func (s *UnlockUsersRequest) SetUsers(v []*string) *UnlockUsersRequest {
	s.Users = v
	return s
}

func (s *UnlockUsersRequest) Validate() error {
	return dara.Validate(s)
}
