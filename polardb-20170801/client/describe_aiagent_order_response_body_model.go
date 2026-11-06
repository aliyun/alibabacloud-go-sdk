// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentVersion(v string) *DescribeAIAgentOrderResponseBody
	GetAgentVersion() *string
	SetCapacityPack(v string) *DescribeAIAgentOrderResponseBody
	GetCapacityPack() *string
	SetExpireTime(v string) *DescribeAIAgentOrderResponseBody
	GetExpireTime() *string
	SetHasOrder(v bool) *DescribeAIAgentOrderResponseBody
	GetHasOrder() *bool
	SetOrderId(v string) *DescribeAIAgentOrderResponseBody
	GetOrderId() *string
	SetOrderName(v string) *DescribeAIAgentOrderResponseBody
	GetOrderName() *string
	SetOrderNum(v string) *DescribeAIAgentOrderResponseBody
	GetOrderNum() *string
	SetOrderStatus(v string) *DescribeAIAgentOrderResponseBody
	GetOrderStatus() *string
	SetRequestId(v string) *DescribeAIAgentOrderResponseBody
	GetRequestId() *string
}

type DescribeAIAgentOrderResponseBody struct {
	// The AI assistant version.
	//
	// example:
	//
	// 1.0
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The capacity package specifications.
	//
	// example:
	//
	// small
	CapacityPack *string `json:"CapacityPack,omitempty" xml:"CapacityPack,omitempty"`
	// The expiration time of the order in RFC 3339 format.
	//
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the current account has an AI assistant order.
	HasOrder *bool `json:"HasOrder,omitempty" xml:"HasOrder,omitempty"`
	// The order ID.
	//
	// example:
	//
	// order-001
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The order name.
	//
	// example:
	//
	// pc-copilot-001
	OrderName *string `json:"OrderName,omitempty" xml:"OrderName,omitempty"`
	// The order number.
	//
	// example:
	//
	// ord-001
	OrderNum *string `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// The order status.
	//
	// example:
	//
	// active
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIAgentOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentOrderResponseBody) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeAIAgentOrderResponseBody) GetCapacityPack() *string {
	return s.CapacityPack
}

func (s *DescribeAIAgentOrderResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAIAgentOrderResponseBody) GetHasOrder() *bool {
	return s.HasOrder
}

func (s *DescribeAIAgentOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeAIAgentOrderResponseBody) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribeAIAgentOrderResponseBody) GetOrderNum() *string {
	return s.OrderNum
}

func (s *DescribeAIAgentOrderResponseBody) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *DescribeAIAgentOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIAgentOrderResponseBody) SetAgentVersion(v string) *DescribeAIAgentOrderResponseBody {
	s.AgentVersion = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetCapacityPack(v string) *DescribeAIAgentOrderResponseBody {
	s.CapacityPack = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetExpireTime(v string) *DescribeAIAgentOrderResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetHasOrder(v bool) *DescribeAIAgentOrderResponseBody {
	s.HasOrder = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetOrderId(v string) *DescribeAIAgentOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetOrderName(v string) *DescribeAIAgentOrderResponseBody {
	s.OrderName = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetOrderNum(v string) *DescribeAIAgentOrderResponseBody {
	s.OrderNum = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetOrderStatus(v string) *DescribeAIAgentOrderResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) SetRequestId(v string) *DescribeAIAgentOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIAgentOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
