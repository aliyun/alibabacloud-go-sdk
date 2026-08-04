// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetUserRolesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterGetUserRolesRequest struct {
}

func (s ModelRouterGetUserRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetUserRolesRequest) Validate() error {
	return dara.Validate(s)
}
