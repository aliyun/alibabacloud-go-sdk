// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Drive) *ListDriveResponseBody
	GetItems() []*Drive
	SetNextMarker(v string) *ListDriveResponseBody
	GetNextMarker() *string
}

type ListDriveResponseBody struct {
	// The queried drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListDriveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDriveResponseBody) GoString() string {
	return s.String()
}

func (s *ListDriveResponseBody) GetItems() []*Drive {
	return s.Items
}

func (s *ListDriveResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListDriveResponseBody) SetItems(v []*Drive) *ListDriveResponseBody {
	s.Items = v
	return s
}

func (s *ListDriveResponseBody) SetNextMarker(v string) *ListDriveResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListDriveResponseBody) Validate() error {
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
