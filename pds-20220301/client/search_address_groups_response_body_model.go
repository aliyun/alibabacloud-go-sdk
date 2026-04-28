// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAddressGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AddressGroup) *SearchAddressGroupsResponseBody
	GetItems() []*AddressGroup
}

type SearchAddressGroupsResponseBody struct {
	// The location-based groups.
	Items []*AddressGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s SearchAddressGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAddressGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsResponseBody) GetItems() []*AddressGroup {
	return s.Items
}

func (s *SearchAddressGroupsResponseBody) SetItems(v []*AddressGroup) *SearchAddressGroupsResponseBody {
	s.Items = v
	return s
}

func (s *SearchAddressGroupsResponseBody) Validate() error {
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
