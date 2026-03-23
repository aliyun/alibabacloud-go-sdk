// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionResponseBody
	GetDBInstanceId() *string
	SetOrderId(v string) *UpgradeDBInstanceMajorVersionResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeDBInstanceMajorVersionResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *UpgradeDBInstanceMajorVersionResponseBody
	GetTaskId() *int64
}

type UpgradeDBInstanceMajorVersionResponseBody struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) SetOrderId(v string) *UpgradeDBInstanceMajorVersionResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceMajorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) SetTaskId(v int64) *UpgradeDBInstanceMajorVersionResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
