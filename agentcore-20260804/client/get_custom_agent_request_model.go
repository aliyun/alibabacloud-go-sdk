// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCustomAgentRequest struct {
}

func (s GetCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *GetCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
