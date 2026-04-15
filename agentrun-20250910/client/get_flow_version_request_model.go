// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFlowVersionRequest struct {
}

func (s GetFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *GetFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
