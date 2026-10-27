// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderSCIMCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListIdentityProviderSCIMCredentialsRequest struct {
}

func (s ListIdentityProviderSCIMCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderSCIMCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderSCIMCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
