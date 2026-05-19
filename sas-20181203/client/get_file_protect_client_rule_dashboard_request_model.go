// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFileProtectClientRuleDashboardRequest struct {
}

func (s GetFileProtectClientRuleDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleDashboardRequest) Validate() error {
	return dara.Validate(s)
}
