// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersOnlineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckLiveMessageUsersOnlineResponseBody
	GetRequestId() *string
	SetUserList(v []*CheckLiveMessageUsersOnlineResponseBodyUserList) *CheckLiveMessageUsersOnlineResponseBody
	GetUserList() []*CheckLiveMessageUsersOnlineResponseBodyUserList
}

type CheckLiveMessageUsersOnlineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 178F572F-AECF-100B-937A-B8047B4D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of users queried.
	UserList []*CheckLiveMessageUsersOnlineResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s CheckLiveMessageUsersOnlineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersOnlineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckLiveMessageUsersOnlineResponseBody) GetUserList() []*CheckLiveMessageUsersOnlineResponseBodyUserList {
	return s.UserList
}

func (s *CheckLiveMessageUsersOnlineResponseBody) SetRequestId(v string) *CheckLiveMessageUsersOnlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponseBody) SetUserList(v []*CheckLiveMessageUsersOnlineResponseBodyUserList) *CheckLiveMessageUsersOnlineResponseBody {
	s.UserList = v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckLiveMessageUsersOnlineResponseBodyUserList struct {
	// Indicates whether the user is online.
	//
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The user ID.
	//
	// example:
	//
	// uid1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckLiveMessageUsersOnlineResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersOnlineResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersOnlineResponseBodyUserList) GetOnline() *bool {
	return s.Online
}

func (s *CheckLiveMessageUsersOnlineResponseBodyUserList) GetUserId() *string {
	return s.UserId
}

func (s *CheckLiveMessageUsersOnlineResponseBodyUserList) SetOnline(v bool) *CheckLiveMessageUsersOnlineResponseBodyUserList {
	s.Online = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponseBodyUserList) SetUserId(v string) *CheckLiveMessageUsersOnlineResponseBodyUserList {
	s.UserId = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
