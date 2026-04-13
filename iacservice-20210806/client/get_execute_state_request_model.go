// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExecuteStateRequest struct {
}

func (s GetExecuteStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStateRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteStateRequest) Validate() error {
	return dara.Validate(s)
}
