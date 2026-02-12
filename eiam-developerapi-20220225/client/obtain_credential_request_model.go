// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialIdentifier(v string) *ObtainCredentialRequest
	GetCredentialIdentifier() *string
}

type ObtainCredentialRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"credentialIdentifier,omitempty" xml:"credentialIdentifier,omitempty"`
}

func (s ObtainCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialRequest) GoString() string {
	return s.String()
}

func (s *ObtainCredentialRequest) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *ObtainCredentialRequest) SetCredentialIdentifier(v string) *ObtainCredentialRequest {
	s.CredentialIdentifier = &v
	return s
}

func (s *ObtainCredentialRequest) Validate() error {
	return dara.Validate(s)
}
