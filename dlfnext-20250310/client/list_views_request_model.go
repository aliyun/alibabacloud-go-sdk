// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListViewsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListViewsRequest
	GetPageToken() *string
	SetViewNamePattern(v string) *ListViewsRequest
	GetViewNamePattern() *string
}

type ListViewsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// view%
	ViewNamePattern *string `json:"viewNamePattern,omitempty" xml:"viewNamePattern,omitempty"`
}

func (s ListViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListViewsRequest) GoString() string {
	return s.String()
}

func (s *ListViewsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListViewsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListViewsRequest) GetViewNamePattern() *string {
	return s.ViewNamePattern
}

func (s *ListViewsRequest) SetMaxResults(v int32) *ListViewsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListViewsRequest) SetPageToken(v string) *ListViewsRequest {
	s.PageToken = &v
	return s
}

func (s *ListViewsRequest) SetViewNamePattern(v string) *ListViewsRequest {
	s.ViewNamePattern = &v
	return s
}

func (s *ListViewsRequest) Validate() error {
	return dara.Validate(s)
}
