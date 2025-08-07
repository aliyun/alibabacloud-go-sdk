// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBEndpointAddressResponseBody
	GetRequestId() *string
}

type CreateDBEndpointAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBEndpointAddressResponseBody) SetRequestId(v string) *CreateDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
