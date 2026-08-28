// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetServiceEndpointRequest struct {
}

func (s GetServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
