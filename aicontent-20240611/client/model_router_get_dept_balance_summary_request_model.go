// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetDeptBalanceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ModelRouterGetDeptBalanceSummaryRequest struct {
}

func (s ModelRouterGetDeptBalanceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetDeptBalanceSummaryRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetDeptBalanceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
