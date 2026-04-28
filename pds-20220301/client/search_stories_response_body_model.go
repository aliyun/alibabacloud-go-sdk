// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchStoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Story) *SearchStoriesResponseBody
	GetItems() []*Story
	SetNextMarker(v string) *SearchStoriesResponseBody
	GetNextMarker() *string
}

type SearchStoriesResponseBody struct {
	Items []*Story `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJ***
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchStoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchStoriesResponseBody) GetItems() []*Story {
	return s.Items
}

func (s *SearchStoriesResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchStoriesResponseBody) SetItems(v []*Story) *SearchStoriesResponseBody {
	s.Items = v
	return s
}

func (s *SearchStoriesResponseBody) SetNextMarker(v string) *SearchStoriesResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchStoriesResponseBody) Validate() error {
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
