// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacegroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*FaceGroup) *ListFacegroupsResponseBody
	GetItems() []*FaceGroup
	SetNextMarker(v string) *ListFacegroupsResponseBody
	GetNextMarker() *string
	SetTotalCount(v int64) *ListFacegroupsResponseBody
	GetTotalCount() *int64
}

type ListFacegroupsResponseBody struct {
	// The information about the face-based groups.
	Items []*FaceGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	TotalCount *int64  `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListFacegroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFacegroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacegroupsResponseBody) GetItems() []*FaceGroup {
	return s.Items
}

func (s *ListFacegroupsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListFacegroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFacegroupsResponseBody) SetItems(v []*FaceGroup) *ListFacegroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListFacegroupsResponseBody) SetNextMarker(v string) *ListFacegroupsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListFacegroupsResponseBody) SetTotalCount(v int64) *ListFacegroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFacegroupsResponseBody) Validate() error {
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
