// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateToOtherZoneResponseBody
	GetDBInstanceId() *string
	SetOrderId(v int64) *MigrateToOtherZoneResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *MigrateToOtherZoneResponseBody
	GetRequestId() *string
}

type MigrateToOtherZoneResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the order. This parameter is returned only when the instance runs MySQL.
	//
	// example:
	//
	// 213341575990728
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateToOtherZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateToOtherZoneResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *MigrateToOtherZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateToOtherZoneResponseBody) SetDBInstanceId(v string) *MigrateToOtherZoneResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) SetOrderId(v int64) *MigrateToOtherZoneResponseBody {
	s.OrderId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) SetRequestId(v string) *MigrateToOtherZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateToOtherZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
