// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderSCIMCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteIdentityProviderSCIMCredentialRequest struct {
}

func (s DeleteIdentityProviderSCIMCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderSCIMCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderSCIMCredentialRequest) Validate() error {
	return dara.Validate(s)
}
