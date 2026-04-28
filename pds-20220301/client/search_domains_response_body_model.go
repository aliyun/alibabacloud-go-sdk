// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Domain) *SearchDomainsResponseBody
	GetItems() []*Domain
	SetNextMarker(v string) *SearchDomainsResponseBody
	GetNextMarker() *string
}

type SearchDomainsResponseBody struct {
	// The queried domains.
	Items []*Domain `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDomainsResponseBody) GetItems() []*Domain {
	return s.Items
}

func (s *SearchDomainsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchDomainsResponseBody) SetItems(v []*Domain) *SearchDomainsResponseBody {
	s.Items = v
	return s
}

func (s *SearchDomainsResponseBody) SetNextMarker(v string) *SearchDomainsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchDomainsResponseBody) Validate() error {
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
