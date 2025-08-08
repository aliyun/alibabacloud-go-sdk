// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvidedSharesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProvidedSharesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListProvidedSharesRequest
	GetPageToken() *string
}

type ListProvidedSharesRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListProvidedSharesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProvidedSharesRequest) GoString() string {
	return s.String()
}

func (s *ListProvidedSharesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProvidedSharesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListProvidedSharesRequest) SetMaxResults(v int32) *ListProvidedSharesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProvidedSharesRequest) SetPageToken(v string) *ListProvidedSharesRequest {
	s.PageToken = &v
	return s
}

func (s *ListProvidedSharesRequest) Validate() error {
	return dara.Validate(s)
}
