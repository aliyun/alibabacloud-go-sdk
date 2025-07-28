// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTopNMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkFlowTopNValues(v []*DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) *DescribeNetworkFlowTopNMetricResponseBody
	GetNetworkFlowTopNValues() []*DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues
	SetRequestId(v string) *DescribeNetworkFlowTopNMetricResponseBody
	GetRequestId() *string
	SetTopNMetaData(v *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) *DescribeNetworkFlowTopNMetricResponseBody
	GetTopNMetaData() *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData
}

type DescribeNetworkFlowTopNMetricResponseBody struct {
	// The top statistical data array returned.
	NetworkFlowTopNValues []*DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues `json:"NetworkFlowTopNValues,omitempty" xml:"NetworkFlowTopNValues,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-D33C8B4C7726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the returned data.
	TopNMetaData *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData `json:"TopNMetaData,omitempty" xml:"TopNMetaData,omitempty" type:"Struct"`
}

func (s DescribeNetworkFlowTopNMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) GetNetworkFlowTopNValues() []*DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues {
	return s.NetworkFlowTopNValues
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) GetTopNMetaData() *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData {
	return s.TopNMetaData
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) SetNetworkFlowTopNValues(v []*DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) *DescribeNetworkFlowTopNMetricResponseBody {
	s.NetworkFlowTopNValues = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) SetRequestId(v string) *DescribeNetworkFlowTopNMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) SetTopNMetaData(v *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) *DescribeNetworkFlowTopNMetricResponseBody {
	s.TopNMetaData = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues struct {
	// Returns additional information, such as the country, province, or city to which an IP address belongs.
	//
	// example:
	//
	// ""
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The value of this field varies depending on the queried Metric.
	//
	// example:
	//
	// 127.0.0.1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Counts for top ranking.
	//
	// example:
	//
	// 1123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) GetName() *string {
	return s.Name
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) GetValue() *int64 {
	return s.Value
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) SetAttribute(v string) *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues {
	s.Attribute = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) SetName(v string) *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues {
	s.Name = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) SetValue(v int64) *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues {
	s.Value = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyNetworkFlowTopNValues) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData struct {
	// The query time range.
	DateRange *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
	// The unit of the returned data. It is fixed as requests.
	//
	// example:
	//
	// requests
	Units *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) GetDateRange() *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange {
	return s.DateRange
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) GetUnits() *string {
	return s.Units
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) SetDateRange(v *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData {
	s.DateRange = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) SetUnits(v string) *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData {
	s.Units = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaData) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange struct {
	// The end time of the query range (Unix timestamp, seconds). Same as the EndDate request parameter.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query range (Unix timestamp, seconds). Same as the StartDate request parameter.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) SetEndDate(v int64) *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) SetStartDate(v int64) *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponseBodyTopNMetaDataDateRange) Validate() error {
	return dara.Validate(s)
}
