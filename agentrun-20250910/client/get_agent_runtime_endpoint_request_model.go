// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentRuntimeEndpointRequest struct {
}

func (s GetAgentRuntimeEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeEndpointRequest) Validate() error {
	return dara.Validate(s)
}
