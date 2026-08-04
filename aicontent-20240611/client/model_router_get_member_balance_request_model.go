// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterGetMemberBalanceRequest struct {
}

func (s ModelRouterGetMemberBalanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceRequest) Validate() error {
	return dara.Validate(s)
}
