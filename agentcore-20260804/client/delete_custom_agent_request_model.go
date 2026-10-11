// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCustomAgentRequest struct {
}

func (s DeleteCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
