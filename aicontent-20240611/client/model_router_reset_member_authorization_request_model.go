// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterResetMemberAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterResetMemberAuthorizationRequest struct {
}

func (s ModelRouterResetMemberAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterResetMemberAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterResetMemberAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
