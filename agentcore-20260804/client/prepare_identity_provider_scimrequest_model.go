// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareIdentityProviderSCIMRequest interface {
	dara.Model
	String() string
	GoString() string
}

type PrepareIdentityProviderSCIMRequest struct {
}

func (s PrepareIdentityProviderSCIMRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMRequest) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMRequest) Validate() error {
	return dara.Validate(s)
}
