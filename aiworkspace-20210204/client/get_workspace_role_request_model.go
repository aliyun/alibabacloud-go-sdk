// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRoleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetWorkspaceRoleRequest struct {
}

func (s GetWorkspaceRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleRequest) Validate() error {
	return dara.Validate(s)
}
