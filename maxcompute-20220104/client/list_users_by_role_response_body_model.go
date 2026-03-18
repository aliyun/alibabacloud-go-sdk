// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersByRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListUsersByRoleResponseBodyData) *ListUsersByRoleResponseBody
	GetData() *ListUsersByRoleResponseBodyData
	SetRequestId(v string) *ListUsersByRoleResponseBody
	GetRequestId() *string
}

type ListUsersByRoleResponseBody struct {
	Data      *ListUsersByRoleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersByRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersByRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBody) GetData() *ListUsersByRoleResponseBodyData {
	return s.Data
}

func (s *ListUsersByRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersByRoleResponseBody) SetData(v *ListUsersByRoleResponseBodyData) *ListUsersByRoleResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersByRoleResponseBody) SetRequestId(v string) *ListUsersByRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersByRoleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersByRoleResponseBodyData struct {
	Users []*ListUsersByRoleResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersByRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersByRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyData) GetUsers() []*ListUsersByRoleResponseBodyDataUsers {
	return s.Users
}

func (s *ListUsersByRoleResponseBodyData) SetUsers(v []*ListUsersByRoleResponseBodyDataUsers) *ListUsersByRoleResponseBodyData {
	s.Users = v
	return s
}

func (s *ListUsersByRoleResponseBodyData) Validate() error {
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

type ListUsersByRoleResponseBodyDataUsers struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListUsersByRoleResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersByRoleResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyDataUsers) GetName() *string {
	return s.Name
}

func (s *ListUsersByRoleResponseBodyDataUsers) SetName(v string) *ListUsersByRoleResponseBodyDataUsers {
	s.Name = &v
	return s
}

func (s *ListUsersByRoleResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
