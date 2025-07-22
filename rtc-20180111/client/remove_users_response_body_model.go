// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUsersResponseBody
	GetRequestId() *string
	SetUsers(v []*RemoveUsersResponseBodyUsers) *RemoveUsersResponseBody
	GetUsers() []*RemoveUsersResponseBodyUsers
}

type RemoveUsersResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*RemoveUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUsersResponseBody) GetUsers() []*RemoveUsersResponseBodyUsers {
	return s.Users
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersResponseBody) SetUsers(v []*RemoveUsersResponseBodyUsers) *RemoveUsersResponseBody {
	s.Users = v
	return s
}

func (s *RemoveUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveUsersResponseBodyUsers struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyUsers) GetCode() *int32 {
	return s.Code
}

func (s *RemoveUsersResponseBodyUsers) GetMessage() *string {
	return s.Message
}

func (s *RemoveUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUsersResponseBodyUsers) SetCode(v int32) *RemoveUsersResponseBodyUsers {
	s.Code = &v
	return s
}

func (s *RemoveUsersResponseBodyUsers) SetMessage(v string) *RemoveUsersResponseBodyUsers {
	s.Message = &v
	return s
}

func (s *RemoveUsersResponseBodyUsers) SetUserId(v string) *RemoveUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *RemoveUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
