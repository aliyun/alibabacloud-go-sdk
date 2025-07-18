// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *UpgradeDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeDBInstanceResponseBody
	GetRequestId() *string
}

type UpgradeDBInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-rj***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// *********
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25C11EE5-B7E8-481A-A07C-BD**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceResponseBody) SetDBInstanceId(v string) *UpgradeDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetOrderId(v string) *UpgradeDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) SetRequestId(v string) *UpgradeDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
