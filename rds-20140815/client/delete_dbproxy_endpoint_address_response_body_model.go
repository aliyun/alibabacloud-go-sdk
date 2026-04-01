// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBProxyEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBProxyEndpointAddressResponseBody
	GetRequestId() *string
}

type DeleteDBProxyEndpointAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 343356C6-64B2-4245-ADEB-C9BD165EDD11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBProxyEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBProxyEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBProxyEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBProxyEndpointAddressResponseBody) SetRequestId(v string) *DeleteDBProxyEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
