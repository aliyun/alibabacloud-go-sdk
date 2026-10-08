// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RecoveryDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *RecoveryDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RecoveryDBInstanceResponseBody
	GetRequestId() *string
}

type RecoveryDBInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-bp1v****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2270972****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E4CDD460-2618-51FE-BD0B-A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoveryDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoveryDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RecoveryDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RecoveryDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoveryDBInstanceResponseBody) SetDBInstanceId(v string) *RecoveryDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *RecoveryDBInstanceResponseBody) SetOrderId(v string) *RecoveryDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RecoveryDBInstanceResponseBody) SetRequestId(v string) *RecoveryDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
