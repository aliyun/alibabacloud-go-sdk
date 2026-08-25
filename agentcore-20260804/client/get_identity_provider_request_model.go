// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIdentityProviderRequest struct {
}

func (s GetIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
