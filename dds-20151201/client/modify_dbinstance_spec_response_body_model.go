// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDBInstanceSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDBInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceSpecResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 21012408824****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E9F1104-19E7-59F0-AB7F-F4EBFDEA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDBInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceSpecResponseBody) SetOrderId(v string) *ModifyDBInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstanceSpecResponseBody) SetRequestId(v string) *ModifyDBInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
