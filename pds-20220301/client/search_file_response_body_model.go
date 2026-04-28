// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*File) *SearchFileResponseBody
	GetItems() []*File
	SetNextMarker(v string) *SearchFileResponseBody
	GetNextMarker() *string
	SetTotalCount(v int64) *SearchFileResponseBody
	GetTotalCount() *int64
}

type SearchFileResponseBody struct {
	// The information about the files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The total number of retrieved files.
	//
	// example:
	//
	// 1022
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s SearchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFileResponseBody) GetItems() []*File {
	return s.Items
}

func (s *SearchFileResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchFileResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchFileResponseBody) SetItems(v []*File) *SearchFileResponseBody {
	s.Items = v
	return s
}

func (s *SearchFileResponseBody) SetNextMarker(v string) *SearchFileResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchFileResponseBody) SetTotalCount(v int64) *SearchFileResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFileResponseBody) Validate() error {
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
