// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ShareLink) *ListShareLinkResponseBody
	GetItems() []*ShareLink
	SetNextMarker(v string) *ListShareLinkResponseBody
	GetNextMarker() *string
}

type ListShareLinkResponseBody struct {
	// The information about the shares.
	Items []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListShareLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *ListShareLinkResponseBody) GetItems() []*ShareLink {
	return s.Items
}

func (s *ListShareLinkResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListShareLinkResponseBody) SetItems(v []*ShareLink) *ListShareLinkResponseBody {
	s.Items = v
	return s
}

func (s *ListShareLinkResponseBody) SetNextMarker(v string) *ListShareLinkResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListShareLinkResponseBody) Validate() error {
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
