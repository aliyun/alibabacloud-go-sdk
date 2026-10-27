// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateIdentityProviderSyncJobRequest struct {
}

func (s CreateIdentityProviderSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSyncJobRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
