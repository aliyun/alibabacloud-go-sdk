// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *DeleteAgenticVpcEndpointRequest
	GetEndpointId() *string
}

type DeleteAgenticVpcEndpointRequest struct {
	// The ID of the endpoint to be deleted. You can delete the endpoint even if it has associated data sources, but this may interrupt access.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-7xvra4ad6619266ba43d
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s DeleteAgenticVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticVpcEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteAgenticVpcEndpointRequest) SetEndpointId(v string) *DeleteAgenticVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteAgenticVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
