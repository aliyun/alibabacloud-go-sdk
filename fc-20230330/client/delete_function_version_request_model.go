// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFunctionVersionRequest struct {
}

func (s DeleteFunctionVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionVersionRequest) Validate() error {
	return dara.Validate(s)
}
