// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIdentityProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIdentityProvidersRequest
	GetNextToken() *string
}

type ListIdentityProvidersRequest struct {
	// The maximum number of records to return per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The cursor for querying the next page. Do not specify this parameter for the first call. For subsequent calls, specify the nextToken value returned from the previous call.
	//
	// example:
	//
	// aWRlbnRpdHktcHJvdmlkZXItb2Zmc2V0OjEw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListIdentityProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersRequest) SetMaxResults(v int32) *ListIdentityProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProvidersRequest) SetNextToken(v string) *ListIdentityProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersRequest) Validate() error {
	return dara.Validate(s)
}
