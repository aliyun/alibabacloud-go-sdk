// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Revision) *ListRevisionResponseBody
	GetItems() []*Revision
	SetNextMarker(v string) *ListRevisionResponseBody
	GetNextMarker() *string
}

type ListRevisionResponseBody struct {
	// The information about the versions.
	Items []*Revision `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListRevisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *ListRevisionResponseBody) GetItems() []*Revision {
	return s.Items
}

func (s *ListRevisionResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListRevisionResponseBody) SetItems(v []*Revision) *ListRevisionResponseBody {
	s.Items = v
	return s
}

func (s *ListRevisionResponseBody) SetNextMarker(v string) *ListRevisionResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListRevisionResponseBody) Validate() error {
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
