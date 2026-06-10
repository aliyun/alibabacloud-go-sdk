// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDesktopSpecResponseBody
	GetOrderId() *string
	SetOrderIds(v []*int64) *ModifyDesktopSpecResponseBody
	GetOrderIds() []*int64
	SetRequestId(v string) *ModifyDesktopSpecResponseBody
	GetRequestId() *string
}

type ModifyDesktopSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 123456789
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// A list of order IDs.
	OrderIds []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDesktopSpecResponseBody) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *ModifyDesktopSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopSpecResponseBody) SetOrderId(v string) *ModifyDesktopSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopSpecResponseBody) SetOrderIds(v []*int64) *ModifyDesktopSpecResponseBody {
	s.OrderIds = v
	return s
}

func (s *ModifyDesktopSpecResponseBody) SetRequestId(v string) *ModifyDesktopSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
