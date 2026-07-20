// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListViewsResponseBody
	GetNextPageToken() *string
	SetViews(v []*string) *ListViewsResponseBody
	GetViews() []*string
}

type ListViewsResponseBody struct {
	// The token to retrieve the next page of results. A null value indicates that this is the last page.
	//
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	// A list of view names.
	Views []*string `json:"views,omitempty" xml:"views,omitempty" type:"Repeated"`
}

func (s ListViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListViewsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListViewsResponseBody) GetViews() []*string {
	return s.Views
}

func (s *ListViewsResponseBody) SetNextPageToken(v string) *ListViewsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListViewsResponseBody) SetViews(v []*string) *ListViewsResponseBody {
	s.Views = v
	return s
}

func (s *ListViewsResponseBody) Validate() error {
	return dara.Validate(s)
}
