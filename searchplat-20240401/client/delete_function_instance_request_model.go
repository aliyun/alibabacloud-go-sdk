// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFunctionInstanceRequest struct {
}

func (s DeleteFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
