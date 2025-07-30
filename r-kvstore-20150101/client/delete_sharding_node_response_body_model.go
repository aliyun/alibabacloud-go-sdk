// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShardingNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DeleteShardingNodeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DeleteShardingNodeResponseBody
	GetRequestId() *string
}

type DeleteShardingNodeResponseBody struct {
	// The ID of the order. On the Orders page in the Billing Management console, you can obtain the details of the order based on the order ID.
	//
	// example:
	//
	// 22179******0904
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteShardingNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteShardingNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteShardingNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteShardingNodeResponseBody) SetOrderId(v string) *DeleteShardingNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *DeleteShardingNodeResponseBody) SetRequestId(v string) *DeleteShardingNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteShardingNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
