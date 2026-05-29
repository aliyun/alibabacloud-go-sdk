// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRoleRequest struct {
}

func (s DeleteRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) Validate() error {
	return dara.Validate(s)
}
