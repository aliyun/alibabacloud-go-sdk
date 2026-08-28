// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIMChannelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentIMChannelRequest struct {
}

func (s GetAgentIMChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIMChannelRequest) GoString() string {
	return s.String()
}

func (s *GetAgentIMChannelRequest) Validate() error {
	return dara.Validate(s)
}
