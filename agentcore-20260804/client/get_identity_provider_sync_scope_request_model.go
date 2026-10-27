// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncScopeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIdentityProviderSyncScopeRequest struct {
}

func (s GetIdentityProviderSyncScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeRequest) Validate() error {
	return dara.Validate(s)
}
