// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentRequest struct {
}

func (s DeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
