// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedSharesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListReceivedSharesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListReceivedSharesRequest
	GetPageToken() *string
}

type ListReceivedSharesRequest struct {
	// The maximum number of records to return for this request.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token to retrieve the next page of results. If the response does not include this token, pass an empty string ("").
	//
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListReceivedSharesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedSharesRequest) GoString() string {
	return s.String()
}

func (s *ListReceivedSharesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReceivedSharesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListReceivedSharesRequest) SetMaxResults(v int32) *ListReceivedSharesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReceivedSharesRequest) SetPageToken(v string) *ListReceivedSharesRequest {
	s.PageToken = &v
	return s
}

func (s *ListReceivedSharesRequest) Validate() error {
	return dara.Validate(s)
}
