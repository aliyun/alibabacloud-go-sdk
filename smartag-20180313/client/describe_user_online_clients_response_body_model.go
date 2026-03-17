// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserOnlineClientsResponseBody
	GetRequestId() *string
	SetUsers(v *DescribeUserOnlineClientsResponseBodyUsers) *DescribeUserOnlineClientsResponseBody
	GetUsers() *DescribeUserOnlineClientsResponseBodyUsers
}

type DescribeUserOnlineClientsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7108A98F-C47D-45F7-A4D8-C2E3022735DA
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     *DescribeUserOnlineClientsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s DescribeUserOnlineClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserOnlineClientsResponseBody) GetUsers() *DescribeUserOnlineClientsResponseBodyUsers {
	return s.Users
}

func (s *DescribeUserOnlineClientsResponseBody) SetRequestId(v string) *DescribeUserOnlineClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserOnlineClientsResponseBody) SetUsers(v *DescribeUserOnlineClientsResponseBodyUsers) *DescribeUserOnlineClientsResponseBody {
	s.Users = v
	return s
}

func (s *DescribeUserOnlineClientsResponseBody) Validate() error {
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserOnlineClientsResponseBodyUsers struct {
	User []*DescribeUserOnlineClientsResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeUserOnlineClientsResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientsResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientsResponseBodyUsers) GetUser() []*DescribeUserOnlineClientsResponseBodyUsersUser {
	return s.User
}

func (s *DescribeUserOnlineClientsResponseBodyUsers) SetUser(v []*DescribeUserOnlineClientsResponseBodyUsersUser) *DescribeUserOnlineClientsResponseBodyUsers {
	s.User = v
	return s
}

func (s *DescribeUserOnlineClientsResponseBodyUsers) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserOnlineClientsResponseBodyUsersUser struct {
	ClientIp   *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	OnlineTime *string `json:"OnlineTime,omitempty" xml:"OnlineTime,omitempty"`
}

func (s DescribeUserOnlineClientsResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientsResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientsResponseBodyUsersUser) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeUserOnlineClientsResponseBodyUsersUser) GetOnlineTime() *string {
	return s.OnlineTime
}

func (s *DescribeUserOnlineClientsResponseBodyUsersUser) SetClientIp(v string) *DescribeUserOnlineClientsResponseBodyUsersUser {
	s.ClientIp = &v
	return s
}

func (s *DescribeUserOnlineClientsResponseBodyUsersUser) SetOnlineTime(v string) *DescribeUserOnlineClientsResponseBodyUsersUser {
	s.OnlineTime = &v
	return s
}

func (s *DescribeUserOnlineClientsResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
