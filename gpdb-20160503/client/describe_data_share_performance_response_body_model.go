// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSharePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDataSharePerformanceResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDataSharePerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeys) *DescribeDataSharePerformanceResponseBody
	GetPerformanceKeys() []*DescribeDataSharePerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDataSharePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDataSharePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDataSharePerformanceResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2022-08-03T15:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details of data sharing performance metrics.
	PerformanceKeys []*DescribeDataSharePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BBE00C04-A3E8-4114-881D-0480A72CB92E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2022-08-03T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDataSharePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDataSharePerformanceResponseBody) GetPerformanceKeys() []*DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDataSharePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSharePerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataSharePerformanceResponseBody) SetDBClusterId(v string) *DescribeDataSharePerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetEndTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetPerformanceKeys(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeys) *DescribeDataSharePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetRequestId(v string) *DescribeDataSharePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) SetStartTime(v string) *DescribeDataSharePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeys struct {
	// The name of the performance metric.
	//
	// example:
	//
	// adbpg_datashare_topic_count
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the performance metric.
	Series []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	//
	// example:
	//
	// int
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) GetName() *string {
	return s.Name
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) GetSeries() []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	return s.Series
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetSeries(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Series = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) SetUnit(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeys {
	s.Unit = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries struct {
	// The name of the performance metric.
	//
	// example:
	//
	// adbpg_datashare_topic_count
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// One or more values of the performance metric.
	Values []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) GetName() *string {
	return s.Name
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) GetValues() []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues {
	return s.Values
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetName(v string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) SetValues(v []*DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries {
	s.Values = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues struct {
	// The value of the performance metric at a point in time.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) GetPoint() []*string {
	return s.Point
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues {
	s.Point = v
	return s
}

func (s *DescribeDataSharePerformanceResponseBodyPerformanceKeysSeriesValues) Validate() error {
	return dara.Validate(s)
}
