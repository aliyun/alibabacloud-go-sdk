// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *PayOrderResponseBody
	GetMetadata() interface{}
	SetOrderId(v int64) *PayOrderResponseBody
	GetOrderId() *int64
	SetPayStatus(v int64) *PayOrderResponseBody
	GetPayStatus() *int64
	SetPayerId(v int64) *PayOrderResponseBody
	GetPayerId() *int64
	SetRequestId(v string) *PayOrderResponseBody
	GetRequestId() *string
}

type PayOrderResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 260591304500425
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	PayStatus *int64 `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// example:
	//
	// 201437655683478597
	PayerId *int64 `json:"PayerId,omitempty" xml:"PayerId,omitempty"`
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *PayOrderResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *PayOrderResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *PayOrderResponseBody) GetPayStatus() *int64 {
	return s.PayStatus
}

func (s *PayOrderResponseBody) GetPayerId() *int64 {
	return s.PayerId
}

func (s *PayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PayOrderResponseBody) SetMetadata(v interface{}) *PayOrderResponseBody {
	s.Metadata = v
	return s
}

func (s *PayOrderResponseBody) SetOrderId(v int64) *PayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *PayOrderResponseBody) SetPayStatus(v int64) *PayOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *PayOrderResponseBody) SetPayerId(v int64) *PayOrderResponseBody {
	s.PayerId = &v
	return s
}

func (s *PayOrderResponseBody) SetRequestId(v string) *PayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *PayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
