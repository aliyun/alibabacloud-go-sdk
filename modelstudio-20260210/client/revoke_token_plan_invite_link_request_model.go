// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenPlanInviteLinkRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RevokeTokenPlanInviteLinkRequest struct {
}

func (s RevokeTokenPlanInviteLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenPlanInviteLinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenPlanInviteLinkRequest) Validate() error {
	return dara.Validate(s)
}
