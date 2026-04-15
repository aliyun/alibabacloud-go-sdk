// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFlowEndpointRequest struct {
}

func (s DeleteFlowEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowEndpointRequest) Validate() error {
	return dara.Validate(s)
}
