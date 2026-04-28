// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Drive) *SearchDriveResponseBody
	GetItems() []*Drive
	SetNextMarker(v string) *SearchDriveResponseBody
	GetNextMarker() *string
}

type SearchDriveResponseBody struct {
	// The information about the drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDriveResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDriveResponseBody) GetItems() []*Drive {
	return s.Items
}

func (s *SearchDriveResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchDriveResponseBody) SetItems(v []*Drive) *SearchDriveResponseBody {
	s.Items = v
	return s
}

func (s *SearchDriveResponseBody) SetNextMarker(v string) *SearchDriveResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchDriveResponseBody) Validate() error {
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
