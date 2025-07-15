// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionSetupOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreatePhysicalConnectionSetupOrderResponseBody
	GetOrderId() *string
	SetPhysicalConnectionId(v string) *CreatePhysicalConnectionSetupOrderResponseBody
	GetPhysicalConnectionId() *string
	SetRequestId(v string) *CreatePhysicalConnectionSetupOrderResponseBody
	GetRequestId() *string
}

type CreatePhysicalConnectionSetupOrderResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 202844382740728
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// example:
	//
	// pc-2zegmc02v7ss4****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7A6301A-64BA-41EC-8284-8F4838C15D1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhysicalConnectionSetupOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionSetupOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) SetOrderId(v string) *CreatePhysicalConnectionSetupOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) SetPhysicalConnectionId(v string) *CreatePhysicalConnectionSetupOrderResponseBody {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) SetRequestId(v string) *CreatePhysicalConnectionSetupOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
