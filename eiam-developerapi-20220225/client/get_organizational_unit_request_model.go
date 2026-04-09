// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetOrganizationalUnitRequest struct {
}

func (s GetOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
