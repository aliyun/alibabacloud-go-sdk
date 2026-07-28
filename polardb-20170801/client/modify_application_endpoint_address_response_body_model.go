// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationEndpointAddressResponseBody
	GetApplicationId() *string
	SetEndpointId(v string) *ModifyApplicationEndpointAddressResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *ModifyApplicationEndpointAddressResponseBody
	GetRequestId() *string
}

type ModifyApplicationEndpointAddressResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 155462B9-205F-4FFC-BB43-4855FE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationEndpointAddressResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationEndpointAddressResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ModifyApplicationEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationEndpointAddressResponseBody) SetApplicationId(v string) *ModifyApplicationEndpointAddressResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressResponseBody) SetEndpointId(v string) *ModifyApplicationEndpointAddressResponseBody {
	s.EndpointId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressResponseBody) SetRequestId(v string) *ModifyApplicationEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
