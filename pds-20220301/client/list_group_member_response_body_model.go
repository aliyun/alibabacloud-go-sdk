// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupItems(v []*Group) *ListGroupMemberResponseBody
	GetGroupItems() []*Group
	SetNextMarker(v string) *ListGroupMemberResponseBody
	GetNextMarker() *string
	SetUserItems(v []*User) *ListGroupMemberResponseBody
	GetUserItems() []*User
}

type ListGroupMemberResponseBody struct {
	// The information about the groups.
	GroupItems []*Group `json:"group_items,omitempty" xml:"group_items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhM1
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The information about the users.
	UserItems []*User `json:"user_items,omitempty" xml:"user_items,omitempty" type:"Repeated"`
}

func (s ListGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBody) GetGroupItems() []*Group {
	return s.GroupItems
}

func (s *ListGroupMemberResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListGroupMemberResponseBody) GetUserItems() []*User {
	return s.UserItems
}

func (s *ListGroupMemberResponseBody) SetGroupItems(v []*Group) *ListGroupMemberResponseBody {
	s.GroupItems = v
	return s
}

func (s *ListGroupMemberResponseBody) SetNextMarker(v string) *ListGroupMemberResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetUserItems(v []*User) *ListGroupMemberResponseBody {
	s.UserItems = v
	return s
}

func (s *ListGroupMemberResponseBody) Validate() error {
	if s.GroupItems != nil {
		for _, item := range s.GroupItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserItems != nil {
		for _, item := range s.UserItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
