// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDBInstanceClassResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBInstanceClassResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceClassResponseBody struct {
	// example:
	//
	// 20211103105558
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBInstanceClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceClassResponseBody) SetOrderId(v string) *ModifyDBInstanceClassResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) SetRequestId(v string) *ModifyDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceClassResponseBody) Validate() error {
	return dara.Validate(s)
}
