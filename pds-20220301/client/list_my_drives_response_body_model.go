// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyDrivesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Drive) *ListMyDrivesResponseBody
	GetItems() []*Drive
	SetNextMarker(v string) *ListMyDrivesResponseBody
	GetNextMarker() *string
}

type ListMyDrivesResponseBody struct {
	// The queried drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListMyDrivesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyDrivesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyDrivesResponseBody) GetItems() []*Drive {
	return s.Items
}

func (s *ListMyDrivesResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListMyDrivesResponseBody) SetItems(v []*Drive) *ListMyDrivesResponseBody {
	s.Items = v
	return s
}

func (s *ListMyDrivesResponseBody) SetNextMarker(v string) *ListMyDrivesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListMyDrivesResponseBody) Validate() error {
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
