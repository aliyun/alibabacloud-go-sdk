// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6GatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpv6GatewayAttributeResponseBody
	GetRequestId() *string
}

type ModifyIpv6GatewayAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DFEDBEE-E5AB-49E8-A2DC-CC114C67AF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpv6GatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6GatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpv6GatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpv6GatewayAttributeResponseBody) SetRequestId(v string) *ModifyIpv6GatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
