// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSCIMCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateIdentityProviderSCIMCredentialRequest struct {
}

func (s CreateIdentityProviderSCIMCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialRequest) Validate() error {
	return dara.Validate(s)
}
