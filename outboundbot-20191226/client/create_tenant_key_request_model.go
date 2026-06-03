// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CreateTenantKeyRequest struct {
}

func (s CreateTenantKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantKeyRequest) Validate() error {
	return dara.Validate(s)
}
