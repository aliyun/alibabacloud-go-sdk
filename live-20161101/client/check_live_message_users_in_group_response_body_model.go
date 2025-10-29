// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckLiveMessageUsersInGroupResponseBody
	GetRequestId() *string
	SetUsers(v []*CheckLiveMessageUsersInGroupResponseBodyUsers) *CheckLiveMessageUsersInGroupResponseBody
	GetUsers() []*CheckLiveMessageUsersInGroupResponseBodyUsers
}

type CheckLiveMessageUsersInGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F1F68D81-1543-1FE4-B56E-82200DD2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of users queried.
	Users []*CheckLiveMessageUsersInGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CheckLiveMessageUsersInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckLiveMessageUsersInGroupResponseBody) GetUsers() []*CheckLiveMessageUsersInGroupResponseBodyUsers {
	return s.Users
}

func (s *CheckLiveMessageUsersInGroupResponseBody) SetRequestId(v string) *CheckLiveMessageUsersInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponseBody) SetUsers(v []*CheckLiveMessageUsersInGroupResponseBodyUsers) *CheckLiveMessageUsersInGroupResponseBody {
	s.Users = v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckLiveMessageUsersInGroupResponseBodyUsers struct {
	// Indicates whether the user is in the group.
	//
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// uid1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckLiveMessageUsersInGroupResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersInGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersInGroupResponseBodyUsers) GetOnline() *bool {
	return s.Online
}

func (s *CheckLiveMessageUsersInGroupResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *CheckLiveMessageUsersInGroupResponseBodyUsers) SetOnline(v bool) *CheckLiveMessageUsersInGroupResponseBodyUsers {
	s.Online = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponseBodyUsers) SetUserId(v string) *CheckLiveMessageUsersInGroupResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
