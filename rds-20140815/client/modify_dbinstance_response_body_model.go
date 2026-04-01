// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *ModifyDBInstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyDBInstanceResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceResponseBody struct {
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 221172852******
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 17F57FEE-EA4F-4337-8D2E-9C23CAA63D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceResponseBody) SetDBInstanceId(v string) *ModifyDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceResponseBody) SetOrderId(v int64) *ModifyDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBInstanceResponseBody) SetRequestId(v string) *ModifyDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
