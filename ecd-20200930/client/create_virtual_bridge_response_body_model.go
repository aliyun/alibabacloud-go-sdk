// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBridgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v string) *CreateVirtualBridgeResponseBody
	GetBridgeId() *string
	SetOrderId(v int64) *CreateVirtualBridgeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateVirtualBridgeResponseBody
	GetRequestId() *string
}

type CreateVirtualBridgeResponseBody struct {
	// The virtual bridge ID.
	//
	// example:
	//
	// vb-sjfaijfish***
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 123456789
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 93AD30C1-16B8-5C54-AD23-A51FF53F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirtualBridgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBridgeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualBridgeResponseBody) GetBridgeId() *string {
	return s.BridgeId
}

func (s *CreateVirtualBridgeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateVirtualBridgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualBridgeResponseBody) SetBridgeId(v string) *CreateVirtualBridgeResponseBody {
	s.BridgeId = &v
	return s
}

func (s *CreateVirtualBridgeResponseBody) SetOrderId(v int64) *CreateVirtualBridgeResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateVirtualBridgeResponseBody) SetRequestId(v string) *CreateVirtualBridgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualBridgeResponseBody) Validate() error {
	return dara.Validate(s)
}
