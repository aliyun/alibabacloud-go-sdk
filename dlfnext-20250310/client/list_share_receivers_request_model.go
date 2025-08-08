// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareReceiversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListShareReceiversRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListShareReceiversRequest
	GetPageToken() *string
}

type ListShareReceiversRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListShareReceiversRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShareReceiversRequest) GoString() string {
	return s.String()
}

func (s *ListShareReceiversRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListShareReceiversRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListShareReceiversRequest) SetMaxResults(v int32) *ListShareReceiversRequest {
	s.MaxResults = &v
	return s
}

func (s *ListShareReceiversRequest) SetPageToken(v string) *ListShareReceiversRequest {
	s.PageToken = &v
	return s
}

func (s *ListShareReceiversRequest) Validate() error {
	return dara.Validate(s)
}
