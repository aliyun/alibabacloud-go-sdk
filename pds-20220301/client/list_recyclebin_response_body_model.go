// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecyclebinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*File) *ListRecyclebinResponseBody
	GetItems() []*File
	SetNextMarker(v string) *ListRecyclebinResponseBody
	GetNextMarker() *string
}

type ListRecyclebinResponseBody struct {
	// The information about the files and folders in the recycle bin.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhM1
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListRecyclebinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecyclebinResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecyclebinResponseBody) GetItems() []*File {
	return s.Items
}

func (s *ListRecyclebinResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListRecyclebinResponseBody) SetItems(v []*File) *ListRecyclebinResponseBody {
	s.Items = v
	return s
}

func (s *ListRecyclebinResponseBody) SetNextMarker(v string) *ListRecyclebinResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListRecyclebinResponseBody) Validate() error {
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
