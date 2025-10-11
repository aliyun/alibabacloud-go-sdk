// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationEndpointAddressResponseBody
	GetApplicationId() *string
	SetEndpointId(v string) *CreateApplicationEndpointAddressResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *CreateApplicationEndpointAddressResponseBody
	GetRequestId() *string
}

type CreateApplicationEndpointAddressResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationEndpointAddressResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationEndpointAddressResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateApplicationEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationEndpointAddressResponseBody) SetApplicationId(v string) *CreateApplicationEndpointAddressResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationEndpointAddressResponseBody) SetEndpointId(v string) *CreateApplicationEndpointAddressResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateApplicationEndpointAddressResponseBody) SetRequestId(v string) *CreateApplicationEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
