// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DisableUserRequest struct {
}

func (s DisableUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableUserRequest) GoString() string {
	return s.String()
}

func (s *DisableUserRequest) Validate() error {
	return dara.Validate(s)
}
