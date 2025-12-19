// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOAuth2CredentialProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOAuth2CredentialProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOAuth2CredentialProvidersRequest
	GetNextToken() *string
}

type ListOAuth2CredentialProvidersRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mmxm9MUzOLyiXaWmj3GOT8
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListOAuth2CredentialProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOAuth2CredentialProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListOAuth2CredentialProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOAuth2CredentialProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOAuth2CredentialProvidersRequest) SetMaxResults(v int32) *ListOAuth2CredentialProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOAuth2CredentialProvidersRequest) SetNextToken(v string) *ListOAuth2CredentialProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOAuth2CredentialProvidersRequest) Validate() error {
	return dara.Validate(s)
}
