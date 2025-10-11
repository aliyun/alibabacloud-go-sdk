// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationEndpointAddressResponseBody
	GetApplicationId() *string
	SetEndpointId(v string) *DeleteApplicationEndpointAddressResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *DeleteApplicationEndpointAddressResponseBody
	GetRequestId() *string
}

type DeleteApplicationEndpointAddressResponseBody struct {
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

func (s DeleteApplicationEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationEndpointAddressResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationEndpointAddressResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteApplicationEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationEndpointAddressResponseBody) SetApplicationId(v string) *DeleteApplicationEndpointAddressResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationEndpointAddressResponseBody) SetEndpointId(v string) *DeleteApplicationEndpointAddressResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DeleteApplicationEndpointAddressResponseBody) SetRequestId(v string) *DeleteApplicationEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
