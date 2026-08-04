// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterDeleteUserRequest struct {
}

func (s ModelRouterDeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteUserRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
