// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForDeleteDBNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *CreateOrderForDeleteDBNodesResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateOrderForDeleteDBNodesResponseBody
	GetRequestId() *string
}

type CreateOrderForDeleteDBNodesResponseBody struct {
	// The instance ID
	//
	// example:
	//
	// rm-7xv******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 221172852******
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 49BC2500-8078-5AC4-A545-20AA5945B0E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderForDeleteDBNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForDeleteDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderForDeleteDBNodesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateOrderForDeleteDBNodesResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateOrderForDeleteDBNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderForDeleteDBNodesResponseBody) SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponseBody) SetOrderId(v int64) *CreateOrderForDeleteDBNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponseBody) SetRequestId(v string) *CreateOrderForDeleteDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
