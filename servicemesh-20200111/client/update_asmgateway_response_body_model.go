// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateASMGatewayResponseBody
	GetRequestId() *string
}

type UpdateASMGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateASMGatewayResponseBody) SetRequestId(v string) *UpdateASMGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateASMGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
