// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBProxyEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBProxyEndpointAddressResponseBody
	GetRequestId() *string
}

type CreateDBProxyEndpointAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBProxyEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBProxyEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBProxyEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBProxyEndpointAddressResponseBody) SetRequestId(v string) *CreateDBProxyEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
