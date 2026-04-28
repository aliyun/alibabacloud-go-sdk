// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchShareLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ShareLink) *SearchShareLinkResponseBody
	GetItems() []*ShareLink
	SetNextMarker(v string) *SearchShareLinkResponseBody
	GetNextMarker() *string
	SetTotalCount(v int64) *SearchShareLinkResponseBody
	GetTotalCount() *int64
}

type SearchShareLinkResponseBody struct {
	// The share URLs.
	Items []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 101
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s SearchShareLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SearchShareLinkResponseBody) GetItems() []*ShareLink {
	return s.Items
}

func (s *SearchShareLinkResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchShareLinkResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchShareLinkResponseBody) SetItems(v []*ShareLink) *SearchShareLinkResponseBody {
	s.Items = v
	return s
}

func (s *SearchShareLinkResponseBody) SetNextMarker(v string) *SearchShareLinkResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchShareLinkResponseBody) SetTotalCount(v int64) *SearchShareLinkResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchShareLinkResponseBody) Validate() error {
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
