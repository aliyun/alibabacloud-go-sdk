// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *GetEndpointAttributeRequest
	GetEndpointType() *string
}

type GetEndpointAttributeRequest struct {
	// The type of the endpoint. Value:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s GetEndpointAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEndpointAttributeRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetEndpointAttributeRequest) SetEndpointType(v string) *GetEndpointAttributeRequest {
	s.EndpointType = &v
	return s
}

func (s *GetEndpointAttributeRequest) Validate() error {
	return dara.Validate(s)
}
