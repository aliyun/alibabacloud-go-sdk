// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSlowLogTrendResponseBodyData) *DescribeSlowLogTrendResponseBody
	GetData() *DescribeSlowLogTrendResponseBodyData
	SetRequestId(v string) *DescribeSlowLogTrendResponseBody
	GetRequestId() *string
}

type DescribeSlowLogTrendResponseBody struct {
	// The returned result.
	Data *DescribeSlowLogTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7D3ECB0E-98CA-5E08-A9CA-F70C5A1E9BDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) GetData() *DescribeSlowLogTrendResponseBodyData {
	return s.Data
}

func (s *DescribeSlowLogTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogTrendResponseBody) SetData(v *DescribeSlowLogTrendResponseBodyData) *DescribeSlowLogTrendResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogTrendResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// clusterTest
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeSlowLogTrendResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeSlowLogTrendResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSlowLogTrendResponseBodyData) GetResultSet() []*DescribeSlowLogTrendResponseBodyDataResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogTrendResponseBodyData) SetDBInstanceID(v int32) *DescribeSlowLogTrendResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyData) SetDBInstanceName(v string) *DescribeSlowLogTrendResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyData) SetResultSet(v []*DescribeSlowLogTrendResponseBodyDataResultSet) *DescribeSlowLogTrendResponseBodyData {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyData) Validate() error {
	if s.ResultSet != nil {
		for _, item := range s.ResultSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogTrendResponseBodyDataResultSet struct {
	// The average execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	AvgQueryDurationMs *int64 `json:"AvgQueryDurationMs,omitempty" xml:"AvgQueryDurationMs,omitempty"`
	// The total number of SQL queries within the specified time range.
	//
	// example:
	//
	// 1
	Cnt *int64 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// The maximum execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	MaxQueryDurationMs *int64 `json:"MaxQueryDurationMs,omitempty" xml:"MaxQueryDurationMs,omitempty"`
	// The minimum execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	MinQueryDurationMs *int64 `json:"MinQueryDurationMs,omitempty" xml:"MinQueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-13 17:48:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) GetAvgQueryDurationMs() *int64 {
	return s.AvgQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) GetCnt() *int64 {
	return s.Cnt
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) GetMaxQueryDurationMs() *int64 {
	return s.MaxQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) GetMinQueryDurationMs() *int64 {
	return s.MinQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetAvgQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.AvgQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetCnt(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.Cnt = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetMaxQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.MaxQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetMinQueryDurationMs(v int64) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.MinQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogTrendResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyDataResultSet) Validate() error {
	return dara.Validate(s)
}
