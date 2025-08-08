// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListShareResourcesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListShareResourcesRequest
	GetPageToken() *string
}

type ListShareResourcesRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListShareResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShareResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListShareResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListShareResourcesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListShareResourcesRequest) SetMaxResults(v int32) *ListShareResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListShareResourcesRequest) SetPageToken(v string) *ListShareResourcesRequest {
	s.PageToken = &v
	return s
}

func (s *ListShareResourcesRequest) Validate() error {
	return dara.Validate(s)
}
