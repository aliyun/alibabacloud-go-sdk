// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserExclusiveCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreateUserExclusiveCredentialResponseBody
	GetCredentialId() *string
	SetCredentialIdentifier(v string) *CreateUserExclusiveCredentialResponseBody
	GetCredentialIdentifier() *string
}

type CreateUserExclusiveCredentialResponseBody struct {
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// 凭据标识。
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"credentialIdentifier,omitempty" xml:"credentialIdentifier,omitempty"`
}

func (s CreateUserExclusiveCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialResponseBody) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateUserExclusiveCredentialResponseBody) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *CreateUserExclusiveCredentialResponseBody) SetCredentialId(v string) *CreateUserExclusiveCredentialResponseBody {
	s.CredentialId = &v
	return s
}

func (s *CreateUserExclusiveCredentialResponseBody) SetCredentialIdentifier(v string) *CreateUserExclusiveCredentialResponseBody {
	s.CredentialIdentifier = &v
	return s
}

func (s *CreateUserExclusiveCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
