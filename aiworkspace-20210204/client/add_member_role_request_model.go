// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberRoleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type AddMemberRoleRequest struct {
}

func (s AddMemberRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMemberRoleRequest) GoString() string {
	return s.String()
}

func (s *AddMemberRoleRequest) Validate() error {
	return dara.Validate(s)
}
