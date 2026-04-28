// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Domain) *ListDomainsResponseBody
	GetItems() []*Domain
	SetNextMarker(v string) *ListDomainsResponseBody
	GetNextMarker() *string
}

type ListDomainsResponseBody struct {
	// The information about the domains.
	Items []*Domain `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) GetItems() []*Domain {
	return s.Items
}

func (s *ListDomainsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListDomainsResponseBody) SetItems(v []*Domain) *ListDomainsResponseBody {
	s.Items = v
	return s
}

func (s *ListDomainsResponseBody) SetNextMarker(v string) *ListDomainsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListDomainsResponseBody) Validate() error {
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
