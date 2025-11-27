// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstancePayTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ModifyDBInstancePayTypeResponseBody
	GetOrderId() *int64
}

type ModifyDBInstancePayTypeResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 100789370230206
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBInstancePayTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstancePayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstancePayTypeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDBInstancePayTypeResponseBody) SetOrderId(v int64) *ModifyDBInstancePayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstancePayTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
