// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIdentityProviderSyncJobRequest struct {
}

func (s GetIdentityProviderSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncJobRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
