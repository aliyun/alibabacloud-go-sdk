// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*User) *ListUserResponseBody
	GetItems() []*User
	SetNextMarker(v string) *ListUserResponseBody
	GetNextMarker() *string
}

type ListUserResponseBody struct {
	// The information about the users.
	Items []*User `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserResponseBody) GetItems() []*User {
	return s.Items
}

func (s *ListUserResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListUserResponseBody) SetItems(v []*User) *ListUserResponseBody {
	s.Items = v
	return s
}

func (s *ListUserResponseBody) SetNextMarker(v string) *ListUserResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListUserResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
