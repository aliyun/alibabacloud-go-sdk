// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenByConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *RevokeAuthenticationTokenByConsumerRequest
	GetConsumerId() *string
	SetCredentialProviderIdentifier(v string) *RevokeAuthenticationTokenByConsumerRequest
	GetCredentialProviderIdentifier() *string
}

type RevokeAuthenticationTokenByConsumerRequest struct {
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

func (s RevokeAuthenticationTokenByConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenByConsumerRequest) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenByConsumerRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *RevokeAuthenticationTokenByConsumerRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *RevokeAuthenticationTokenByConsumerRequest) SetConsumerId(v string) *RevokeAuthenticationTokenByConsumerRequest {
	s.ConsumerId = &v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerRequest) SetCredentialProviderIdentifier(v string) *RevokeAuthenticationTokenByConsumerRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerRequest) Validate() error {
	return dara.Validate(s)
}
