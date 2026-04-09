// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteOrganizationalUnitRequest struct {
}

func (s DeleteOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
