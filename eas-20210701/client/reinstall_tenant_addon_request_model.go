// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallTenantAddonRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ReinstallTenantAddonRequest struct {
}

func (s ReinstallTenantAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinstallTenantAddonRequest) GoString() string {
	return s.String()
}

func (s *ReinstallTenantAddonRequest) Validate() error {
	return dara.Validate(s)
}
