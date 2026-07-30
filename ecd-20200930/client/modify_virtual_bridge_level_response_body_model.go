// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyVirtualBridgeLevelResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyVirtualBridgeLevelResponseBody
	GetRequestId() *string
}

type ModifyVirtualBridgeLevelResponseBody struct {
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
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVirtualBridgeLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeLevelResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyVirtualBridgeLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVirtualBridgeLevelResponseBody) SetOrderId(v int64) *ModifyVirtualBridgeLevelResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyVirtualBridgeLevelResponseBody) SetRequestId(v string) *ModifyVirtualBridgeLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVirtualBridgeLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
