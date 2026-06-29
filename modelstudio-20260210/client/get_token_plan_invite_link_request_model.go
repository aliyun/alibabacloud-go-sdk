// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanInviteLinkRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTokenPlanInviteLinkRequest struct {
}

func (s GetTokenPlanInviteLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanInviteLinkRequest) GoString() string {
	return s.String()
}

func (s *GetTokenPlanInviteLinkRequest) Validate() error {
	return dara.Validate(s)
}
