// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *ModifyDBInstanceSpecResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyDBInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceSpecResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDBInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceSpecResponseBody) SetDBInstanceId(v string) *ModifyDBInstanceSpecResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecResponseBody) SetOrderId(v int64) *ModifyDBInstanceSpecResponseBody {
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
