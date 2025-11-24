// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateASMGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateASMGatewayResponseBody
	GetRequestId() *string
}

type CreateASMGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateASMGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateASMGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateASMGatewayResponseBody) SetRequestId(v string) *CreateASMGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateASMGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
