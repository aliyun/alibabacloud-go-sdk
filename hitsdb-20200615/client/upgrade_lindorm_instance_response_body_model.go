// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *UpgradeLindormInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *UpgradeLindormInstanceResponseBody
	GetRequestId() *string
}

type UpgradeLindormInstanceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2A7D4F9D-AA26-4E15-A2B1-3E4792C6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeLindormInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *UpgradeLindormInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeLindormInstanceResponseBody) SetOrderId(v int64) *UpgradeLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeLindormInstanceResponseBody) SetRequestId(v string) *UpgradeLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeLindormInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
