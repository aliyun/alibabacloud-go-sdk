// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpv4GatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpv4GatewayAttributeResponseBody
	GetRequestId() *string
}

type UpdateIpv4GatewayAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 671CEB03-C98D-5916-950C-C55B0BD06621
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpv4GatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpv4GatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpv4GatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpv4GatewayAttributeResponseBody) SetRequestId(v string) *UpdateIpv4GatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
