// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterQueryModelGroupRequest struct {
}

func (s ModelRouterQueryModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
