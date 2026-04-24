// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOperatorOrderRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelOperatorOrderRequest struct {
}

func (s GetModelOperatorOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelOperatorOrderRequest) GoString() string {
	return s.String()
}

func (s *GetModelOperatorOrderRequest) Validate() error {
	return dara.Validate(s)
}
