// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenByConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *ReinstateAuthenticationTokenByConsumerRequest
	GetConsumerId() *string
	SetCredentialProviderIdentifier(v string) *ReinstateAuthenticationTokenByConsumerRequest
	GetCredentialProviderIdentifier() *string
}

type ReinstateAuthenticationTokenByConsumerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_jwt_subject
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
}

func (s ReinstateAuthenticationTokenByConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenByConsumerRequest) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenByConsumerRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ReinstateAuthenticationTokenByConsumerRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *ReinstateAuthenticationTokenByConsumerRequest) SetConsumerId(v string) *ReinstateAuthenticationTokenByConsumerRequest {
	s.ConsumerId = &v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerRequest) SetCredentialProviderIdentifier(v string) *ReinstateAuthenticationTokenByConsumerRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *ReinstateAuthenticationTokenByConsumerRequest) Validate() error {
	return dara.Validate(s)
}
