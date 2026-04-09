// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMemberRoleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RemoveMemberRoleRequest struct {
}

func (s RemoveMemberRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveMemberRoleRequest) GoString() string {
	return s.String()
}

func (s *RemoveMemberRoleRequest) Validate() error {
	return dara.Validate(s)
}
