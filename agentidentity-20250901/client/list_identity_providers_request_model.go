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
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6m6UfcO+UDmLjng/InRW9IU
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
