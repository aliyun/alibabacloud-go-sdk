// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*File) *ListReceivedFileResponseBody
	GetItems() []*File
	SetNextMarker(v string) *ListReceivedFileResponseBody
	GetNextMarker() *string
}

type ListReceivedFileResponseBody struct {
	// The queried files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// eym***
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListReceivedFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListReceivedFileResponseBody) GetItems() []*File {
	return s.Items
}

func (s *ListReceivedFileResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListReceivedFileResponseBody) SetItems(v []*File) *ListReceivedFileResponseBody {
	s.Items = v
	return s
}

func (s *ListReceivedFileResponseBody) SetNextMarker(v string) *ListReceivedFileResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListReceivedFileResponseBody) Validate() error {
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
