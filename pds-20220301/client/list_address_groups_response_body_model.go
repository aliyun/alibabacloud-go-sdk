// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AddressGroup) *ListAddressGroupsResponseBody
	GetItems() []*AddressGroup
	SetNextMarker(v string) *ListAddressGroupsResponseBody
	GetNextMarker() *string
}

type ListAddressGroupsResponseBody struct {
	// The information about the location-based groups.
	Items []*AddressGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListAddressGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddressGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsResponseBody) GetItems() []*AddressGroup {
	return s.Items
}

func (s *ListAddressGroupsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListAddressGroupsResponseBody) SetItems(v []*AddressGroup) *ListAddressGroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListAddressGroupsResponseBody) SetNextMarker(v string) *ListAddressGroupsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListAddressGroupsResponseBody) Validate() error {
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
