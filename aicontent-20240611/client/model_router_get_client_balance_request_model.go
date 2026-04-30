// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterGetClientBalanceRequest struct {
}

func (s ModelRouterGetClientBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceRequest) Validate() error {
	return dara.Validate(s)
}
