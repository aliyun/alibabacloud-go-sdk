// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationEndpointAddressRequest
	GetApplicationId() *string
	SetEndpointId(v string) *DeleteApplicationEndpointAddressRequest
	GetEndpointId() *string
	SetNetType(v string) *DeleteApplicationEndpointAddressRequest
	GetNetType() *string
}

type DeleteApplicationEndpointAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DeleteApplicationEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationEndpointAddressRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationEndpointAddressRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteApplicationEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *DeleteApplicationEndpointAddressRequest) SetApplicationId(v string) *DeleteApplicationEndpointAddressRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationEndpointAddressRequest) SetEndpointId(v string) *DeleteApplicationEndpointAddressRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteApplicationEndpointAddressRequest) SetNetType(v string) *DeleteApplicationEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *DeleteApplicationEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
