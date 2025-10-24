// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProjectUsersResponseBodyData) *ListProjectUsersResponseBody
	GetData() *ListProjectUsersResponseBodyData
	SetRequestId(v string) *ListProjectUsersResponseBody
	GetRequestId() *string
}

type ListProjectUsersResponseBody struct {
	// The returned data.
	Data *ListProjectUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7b316643495896551555e855b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBody) GetData() *ListProjectUsersResponseBodyData {
	return s.Data
}

func (s *ListProjectUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectUsersResponseBody) SetData(v *ListProjectUsersResponseBodyData) *ListProjectUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectUsersResponseBody) SetRequestId(v string) *ListProjectUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectUsersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectUsersResponseBodyData struct {
	// An array that contains users.
	Users []*ListProjectUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListProjectUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyData) GetUsers() []*ListProjectUsersResponseBodyDataUsers {
	return s.Users
}

func (s *ListProjectUsersResponseBodyData) SetUsers(v []*ListProjectUsersResponseBodyDataUsers) *ListProjectUsersResponseBodyData {
	s.Users = v
	return s
}

func (s *ListProjectUsersResponseBodyData) Validate() error {
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

type ListProjectUsersResponseBodyDataUsers struct {
	// The name of the user.
	//
	// example:
	//
	// userA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProjectUsersResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s ListProjectUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyDataUsers) GetName() *string {
	return s.Name
}

func (s *ListProjectUsersResponseBodyDataUsers) SetName(v string) *ListProjectUsersResponseBodyDataUsers {
	s.Name = &v
	return s
}

func (s *ListProjectUsersResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
