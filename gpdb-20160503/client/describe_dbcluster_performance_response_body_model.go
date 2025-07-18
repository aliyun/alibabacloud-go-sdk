// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody
	GetPerformanceKeys() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBClusterPerformanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-03T15:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the performance metric. For more information, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	PerformanceKeys []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8E8990F0-C81E-4C94-8F51-5F**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the `YYYY-MM-DDTHH:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-03T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterPerformanceResponseBody) GetPerformanceKeys() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBClusterPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	// The name of the performance metric. For more information, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	//
	// example:
	//
	// adbpg_group_cpu_used_percent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the performance metric of a node.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GetSeries() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	return s.Series
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries struct {
	// The name of the compute node or compute group.
	//
	// example:
	//
	// standby-********-cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- **master**: primary coordinator node
	//
	// 	- **standby**: standby coordinator node
	//
	// 	- **segment**: compute node
	//
	// example:
	//
	// standby
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The value of the performance metric collected at a point in time.
	Values []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GetRole() *string {
	return s.Role
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) GetValues() []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues {
	return s.Values
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetRole(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues struct {
	// The value of the performance metric and the time when the metric value was collected.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) GetPoint() []*string {
	return s.Point
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysSeriesValues) Validate() error {
	return dara.Validate(s)
}
