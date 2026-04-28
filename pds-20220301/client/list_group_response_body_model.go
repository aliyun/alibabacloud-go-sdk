// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Group) *ListGroupResponseBody
	GetItems() []*Group
	SetNextMarker(v string) *ListGroupResponseBody
	GetNextMarker() *string
}

type ListGroupResponseBody struct {
	// The information about the groups.
	Items []*Group `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBody) GetItems() []*Group {
	return s.Items
}

func (s *ListGroupResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListGroupResponseBody) SetItems(v []*Group) *ListGroupResponseBody {
	s.Items = v
	return s
}

func (s *ListGroupResponseBody) SetNextMarker(v string) *ListGroupResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListGroupResponseBody) Validate() error {
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
