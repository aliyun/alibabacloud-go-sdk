// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityEventLogsResponseBody
	GetRequestId() *string
	SetSecurityEventLogs(v []interface{}) *DescribeSecurityEventLogsResponseBody
	GetSecurityEventLogs() []interface{}
	SetSecurityEventLogsTotalCount(v int64) *DescribeSecurityEventLogsResponseBody
	GetSecurityEventLogsTotalCount() *int64
	SetSecurityEventMetaData(v *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) *DescribeSecurityEventLogsResponseBody
	GetSecurityEventMetaData() *DescribeSecurityEventLogsResponseBodySecurityEventMetaData
}

type DescribeSecurityEventLogsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-D33C8B4C7726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The attack logs returned.
	SecurityEventLogs []interface{} `json:"SecurityEventLogs,omitempty" xml:"SecurityEventLogs,omitempty" type:"Repeated"`
	// The total number of logs returned.
	//
	// example:
	//
	// 1000
	SecurityEventLogsTotalCount *int64 `json:"SecurityEventLogsTotalCount,omitempty" xml:"SecurityEventLogsTotalCount,omitempty"`
	// The metadata of the time series data returned.
	SecurityEventMetaData *DescribeSecurityEventLogsResponseBodySecurityEventMetaData `json:"SecurityEventMetaData,omitempty" xml:"SecurityEventMetaData,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventLogsResponseBody) GetSecurityEventLogs() []interface{} {
	return s.SecurityEventLogs
}

func (s *DescribeSecurityEventLogsResponseBody) GetSecurityEventLogsTotalCount() *int64 {
	return s.SecurityEventLogsTotalCount
}

func (s *DescribeSecurityEventLogsResponseBody) GetSecurityEventMetaData() *DescribeSecurityEventLogsResponseBodySecurityEventMetaData {
	return s.SecurityEventMetaData
}

func (s *DescribeSecurityEventLogsResponseBody) SetRequestId(v string) *DescribeSecurityEventLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventLogsResponseBody) SetSecurityEventLogs(v []interface{}) *DescribeSecurityEventLogsResponseBody {
	s.SecurityEventLogs = v
	return s
}

func (s *DescribeSecurityEventLogsResponseBody) SetSecurityEventLogsTotalCount(v int64) *DescribeSecurityEventLogsResponseBody {
	s.SecurityEventLogsTotalCount = &v
	return s
}

func (s *DescribeSecurityEventLogsResponseBody) SetSecurityEventMetaData(v *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) *DescribeSecurityEventLogsResponseBody {
	s.SecurityEventMetaData = v
	return s
}

func (s *DescribeSecurityEventLogsResponseBody) Validate() error {
	if s.SecurityEventMetaData != nil {
		if err := s.SecurityEventMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventLogsResponseBodySecurityEventMetaData struct {
	// The time range that is used for the query.
	DateRange *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
	// The unit of the statistics returned. The value is fixed as requests.
	//
	// example:
	//
	// requests
	Units *string `json:"Units,omitempty" xml:"Units,omitempty"`
}

func (s DescribeSecurityEventLogsResponseBodySecurityEventMetaData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsResponseBodySecurityEventMetaData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) GetDateRange() *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) GetUnits() *string {
	return s.Units
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) SetDateRange(v *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) *DescribeSecurityEventLogsResponseBodySecurityEventMetaData {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) SetUnits(v string) *DescribeSecurityEventLogsResponseBodySecurityEventMetaData {
	s.Units = &v
	return s
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaData) Validate() error {
	if s.DateRange != nil {
		if err := s.DateRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. This value is the same as the value of EndDate in the request parameters.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds. This value is the same as the value of StartDate in the request parameters.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) SetEndDate(v int64) *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) SetStartDate(v int64) *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventLogsResponseBodySecurityEventMetaDataDateRange) Validate() error {
	return dara.Validate(s)
}
