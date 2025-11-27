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
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20793850608****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C5CFDEE-F774-4DED-89A2-1D76EC63C575
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
