// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFlowEndpointRequest struct {
}

func (s GetFlowEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetFlowEndpointRequest) Validate() error {
	return dara.Validate(s)
}
