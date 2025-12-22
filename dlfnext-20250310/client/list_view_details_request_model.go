// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListViewDetailsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListViewDetailsRequest
	GetPageToken() *string
	SetViewNamePattern(v string) *ListViewDetailsRequest
	GetViewNamePattern() *string
}

type ListViewDetailsRequest struct {
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

func (s ListViewDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListViewDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListViewDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListViewDetailsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListViewDetailsRequest) GetViewNamePattern() *string {
	return s.ViewNamePattern
}

func (s *ListViewDetailsRequest) SetMaxResults(v int32) *ListViewDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListViewDetailsRequest) SetPageToken(v string) *ListViewDetailsRequest {
	s.PageToken = &v
	return s
}

func (s *ListViewDetailsRequest) SetViewNamePattern(v string) *ListViewDetailsRequest {
	s.ViewNamePattern = &v
	return s
}

func (s *ListViewDetailsRequest) Validate() error {
	return dara.Validate(s)
}
