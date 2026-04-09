// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRuntimeEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentRuntimeEndpointRequest struct {
}

func (s DeleteAgentRuntimeEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRuntimeEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRuntimeEndpointRequest) Validate() error {
	return dara.Validate(s)
}
