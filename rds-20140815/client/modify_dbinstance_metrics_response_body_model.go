// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceMetricsResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *ModifyDBInstanceMetricsResponseBody
	GetRequestId() *string
	SetScope(v string) *ModifyDBInstanceMetricsResponseBody
	GetScope() *string
}

type ModifyDBInstanceMetricsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// pgm-bp1s1j103lo6****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B55934BB-FFAA-5276-80A8-E0FDB12810B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The application scope of this modification. Valid values:
	//
	// 	- **instance**: This modification is applied only to the current instance.
	//
	// 	- **region**: This modification is applied to all ApsaraDB RDS for PostgreSQL instances that are equipped with the same type of storage media as the current instance in the region to which the current instance belongs.
	//
	// example:
	//
	// instance
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ModifyDBInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMetricsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceMetricsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *ModifyDBInstanceMetricsResponseBody) SetDBInstanceId(v string) *ModifyDBInstanceMetricsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) SetRequestId(v string) *ModifyDBInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) SetScope(v string) *ModifyDBInstanceMetricsResponseBody {
	s.Scope = &v
	return s
}

func (s *ModifyDBInstanceMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
