// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationEndpointAddressRequest
	GetApplicationId() *string
	SetEndpointId(v string) *CreateApplicationEndpointAddressRequest
	GetEndpointId() *string
	SetNetType(v string) *CreateApplicationEndpointAddressRequest
	GetNetType() *string
}

type CreateApplicationEndpointAddressRequest struct {
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

func (s CreateApplicationEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationEndpointAddressRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationEndpointAddressRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateApplicationEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *CreateApplicationEndpointAddressRequest) SetApplicationId(v string) *CreateApplicationEndpointAddressRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationEndpointAddressRequest) SetEndpointId(v string) *CreateApplicationEndpointAddressRequest {
	s.EndpointId = &v
	return s
}

func (s *CreateApplicationEndpointAddressRequest) SetNetType(v string) *CreateApplicationEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *CreateApplicationEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
