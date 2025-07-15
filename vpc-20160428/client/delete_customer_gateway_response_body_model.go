// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomerGatewayResponseBody
	GetRequestId() *string
}

type DeleteCustomerGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomerGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomerGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomerGatewayResponseBody) SetRequestId(v string) *DeleteCustomerGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomerGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
