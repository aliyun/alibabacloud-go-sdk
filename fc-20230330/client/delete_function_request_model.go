// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFunctionRequest struct {
}

func (s DeleteFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) Validate() error {
	return dara.Validate(s)
}
