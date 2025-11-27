// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyRCInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyRCInstanceResponseBody
	GetRequestId() *string
}

type ModifyRCInstanceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 100789370230206
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF82B07-28D2-48D1-B5D6-7E78FED277C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRCInstanceResponseBody) SetOrderId(v int64) *ModifyRCInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyRCInstanceResponseBody) SetRequestId(v string) *ModifyRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
