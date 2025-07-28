// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTopNMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityEventTopNMetricResponseBody
	GetRequestId() *string
	SetSecurityEventTopNValues(v []*DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) *DescribeSecurityEventTopNMetricResponseBody
	GetSecurityEventTopNValues() []*DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues
	SetTopNMetaData(v *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) *DescribeSecurityEventTopNMetricResponseBody
	GetTopNMetaData() *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData
}

type DescribeSecurityEventTopNMetricResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-*****B4C7726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top N data entries returned.
	SecurityEventTopNValues []*DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues `json:"SecurityEventTopNValues,omitempty" xml:"SecurityEventTopNValues,omitempty" type:"Repeated"`
	// The metadata of the data entries returned.
	TopNMetaData *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData `json:"TopNMetaData,omitempty" xml:"TopNMetaData,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventTopNMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventTopNMetricResponseBody) GetSecurityEventTopNValues() []*DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues {
	return s.SecurityEventTopNValues
}

func (s *DescribeSecurityEventTopNMetricResponseBody) GetTopNMetaData() *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData {
	return s.TopNMetaData
}

func (s *DescribeSecurityEventTopNMetricResponseBody) SetRequestId(v string) *DescribeSecurityEventTopNMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBody) SetSecurityEventTopNValues(v []*DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) *DescribeSecurityEventTopNMetricResponseBody {
	s.SecurityEventTopNValues = v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBody) SetTopNMetaData(v *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) *DescribeSecurityEventTopNMetricResponseBody {
	s.TopNMetaData = v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues struct {
	// The additional information, such as the protection module for a protection rule whose ID is returned.
	//
	// example:
	//
	// waf_base
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The field value, which varies based on the metric.
	//
	// example:
	//
	// 10000
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The count for the data entry.
	//
	// example:
	//
	// 1111
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) GetName() *string {
	return s.Name
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) GetValue() *int64 {
	return s.Value
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) SetAttribute(v string) *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues {
	s.Attribute = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) SetName(v string) *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues {
	s.Name = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) SetValue(v int64) *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues {
	s.Value = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodySecurityEventTopNValues) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTopNMetricResponseBodyTopNMetaData struct {
	// The time range that is used for the query.
	DateRange *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
	// The unit of the statistics returned. It is fixed as requests.
	//
	// example:
	//
	// requests
	Units *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) GetDateRange() *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) GetUnits() *string {
	return s.Units
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) SetDateRange(v *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) SetUnits(v string) *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData {
	s.Units = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaData) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange struct {
	// The end of the time range to query. The value is a Unix timestamp. Unit: seconds. This value is the same as the value of EndDate in the request parameters.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. The value is a Unix timestamp. Unit: seconds. This value is the same as the value of StartDate in the request parameters.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) SetEndDate(v int64) *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) SetStartDate(v int64) *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricResponseBodyTopNMetaDataDateRange) Validate() error {
	return dara.Validate(s)
}
