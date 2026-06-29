// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanOrgInviteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTokenPlanOrgInviteConfigRequest struct {
}

func (s GetTokenPlanOrgInviteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanOrgInviteConfigRequest) GoString() string {
	return s.String()
}

func (s *GetTokenPlanOrgInviteConfigRequest) Validate() error {
	return dara.Validate(s)
}
