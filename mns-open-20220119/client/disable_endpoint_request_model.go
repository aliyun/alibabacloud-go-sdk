// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *DisableEndpointRequest
	GetEndpointType() *string
}

type DisableEndpointRequest struct {
	// The type of the endpoint. Value:
	//
	// 	- **public**: indicates an public endpoint. (Only the public endpoint is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DisableEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableEndpointRequest) GoString() string {
	return s.String()
}

func (s *DisableEndpointRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DisableEndpointRequest) SetEndpointType(v string) *DisableEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *DisableEndpointRequest) Validate() error {
	return dara.Validate(s)
}
