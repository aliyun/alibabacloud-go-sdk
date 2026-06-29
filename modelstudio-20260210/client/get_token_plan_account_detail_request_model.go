// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanAccountDetailRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTokenPlanAccountDetailRequest struct {
}

func (s GetTokenPlanAccountDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailRequest) Validate() error {
	return dara.Validate(s)
}
