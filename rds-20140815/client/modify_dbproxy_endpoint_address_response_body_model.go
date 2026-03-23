// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBProxyEndpointAddressResponseBody
	GetRequestId() *string
}

type ModifyDBProxyEndpointAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBProxyEndpointAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBProxyEndpointAddressResponseBody) SetRequestId(v string) *ModifyDBProxyEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
