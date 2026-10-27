// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSCIMRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIdentityProviderSCIMRequest struct {
}

func (s GetIdentityProviderSCIMRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSCIMRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSCIMRequest) Validate() error {
	return dara.Validate(s)
}
