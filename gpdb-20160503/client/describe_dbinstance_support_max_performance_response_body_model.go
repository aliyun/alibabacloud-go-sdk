// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSupportMaxPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBody
	GetDBInstanceId() *string
	SetPerformances(v *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) *DescribeDBInstanceSupportMaxPerformanceResponseBody
	GetPerformances() *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances
	SetRequestId(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceSupportMaxPerformanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The queried performance metric.
	Performances *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) GetPerformances() *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances {
	return s.Performances
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) SetPerformances(v *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) *DescribeDBInstanceSupportMaxPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) SetRequestId(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances struct {
	Performance []*DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) GetPerformance() []*DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance {
	return s.Performance
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) SetPerformance(v []*DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances {
	s.Performance = v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformances) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance struct {
	// The performance bottleneck type.
	//
	// example:
	//
	// ecs or disk
	Bottleneck *string `json:"Bottleneck,omitempty" xml:"Bottleneck,omitempty"`
	// The name of the performance metric.
	//
	// example:
	//
	// adbpg_status,adbpg_disk_status,adbpg_connection_status,adbgp_segment_disk_usage_percent_max,adbpg_master_disk_usage_percent_max,adbpg_disk_usage_percent
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The unit of the performance metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the performance metric.
	//
	// example:
	//
	// 90
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) GetBottleneck() *string {
	return s.Bottleneck
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) SetBottleneck(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance {
	s.Bottleneck = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) SetKey(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) SetUnit(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance {
	s.Unit = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) SetValue(v string) *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceSupportMaxPerformanceResponseBodyPerformancesPerformance) Validate() error {
	return dara.Validate(s)
}
