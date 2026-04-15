// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFlowRequest struct {
}

func (s GetFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) Validate() error {
	return dara.Validate(s)
}
