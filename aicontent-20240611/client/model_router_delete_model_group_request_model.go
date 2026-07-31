// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterDeleteModelGroupRequest struct {
}

func (s ModelRouterDeleteModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteModelGroupRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
