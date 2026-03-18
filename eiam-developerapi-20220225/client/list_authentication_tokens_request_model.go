// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthenticationTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ListAuthenticationTokensRequest
	GetConsumerId() *string
	SetCredentialProviderIdentifier(v string) *ListAuthenticationTokensRequest
	GetCredentialProviderIdentifier() *string
	SetExpired(v bool) *ListAuthenticationTokensRequest
	GetExpired() *bool
	SetMaxResults(v int64) *ListAuthenticationTokensRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListAuthenticationTokensRequest
	GetNextToken() *string
	SetRevoked(v bool) *ListAuthenticationTokensRequest
	GetRevoked() *bool
}

type ListAuthenticationTokensRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
	// example:
	//
	// false
	Expired *bool `json:"expired,omitempty" xml:"expired,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
}

func (s ListAuthenticationTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthenticationTokensRequest) GoString() string {
	return s.String()
}

func (s *ListAuthenticationTokensRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListAuthenticationTokensRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *ListAuthenticationTokensRequest) GetExpired() *bool {
	return s.Expired
}

func (s *ListAuthenticationTokensRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAuthenticationTokensRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthenticationTokensRequest) GetRevoked() *bool {
	return s.Revoked
}

func (s *ListAuthenticationTokensRequest) SetConsumerId(v string) *ListAuthenticationTokensRequest {
	s.ConsumerId = &v
	return s
}

func (s *ListAuthenticationTokensRequest) SetCredentialProviderIdentifier(v string) *ListAuthenticationTokensRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *ListAuthenticationTokensRequest) SetExpired(v bool) *ListAuthenticationTokensRequest {
	s.Expired = &v
	return s
}

func (s *ListAuthenticationTokensRequest) SetMaxResults(v int64) *ListAuthenticationTokensRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthenticationTokensRequest) SetNextToken(v string) *ListAuthenticationTokensRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthenticationTokensRequest) SetRevoked(v bool) *ListAuthenticationTokensRequest {
	s.Revoked = &v
	return s
}

func (s *ListAuthenticationTokensRequest) Validate() error {
	return dara.Validate(s)
}
