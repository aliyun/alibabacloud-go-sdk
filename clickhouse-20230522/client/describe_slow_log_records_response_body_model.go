// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSlowLogRecordsResponseBodyData) *DescribeSlowLogRecordsResponseBody
	GetData() *DescribeSlowLogRecordsResponseBodyData
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
}

type DescribeSlowLogRecordsResponseBody struct {
	// The data returned.
	Data *DescribeSlowLogRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DF203CC8-5F68-5E3F-8050-3C77DD65731A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetData() *DescribeSlowLogRecordsResponseBodyData {
	return s.Data
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) SetData(v *DescribeSlowLogRecordsResponseBodyData) *DescribeSlowLogRecordsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z32****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// TestCluster
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeSlowLogRecordsResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetResultSet() []*DescribeSlowLogRecordsResponseBodyDataResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogRecordsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDBInstanceID(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetDBInstanceName(v string) *DescribeSlowLogRecordsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetResultSet(v []*DescribeSlowLogRecordsResponseBodyDataResultSet) *DescribeSlowLogRecordsResponseBodyData {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) SetTotalCount(v int32) *DescribeSlowLogRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyData) Validate() error {
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

type DescribeSlowLogRecordsResponseBodyDataResultSet struct {
	// The address to which the query statement is sent.
	//
	// example:
	//
	// 0:0:0:0:0:ffff:1edd65ea
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"ae915a3ad30e77e67a7215d05b658cc6\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// bany
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The peak memory usage for the query. Unit: bytes.
	//
	// example:
	//
	// 4941696
	MemoryUsage *int64 `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// The query statement that is running.
	//
	// example:
	//
	// select 	- from test
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	QueryDurationMs *int64 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-09-11 16:00:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The size of the data that is scanned. Unit: bytes.
	//
	// example:
	//
	// 4507128020832
	ReadBytes *int64 `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// The number of read rows.
	//
	// example:
	//
	// 10
	ReadRows *int64 `json:"ReadRows,omitempty" xml:"ReadRows,omitempty"`
	// The size of the result data. Unit: bytes.
	//
	// example:
	//
	// 10
	ResultBytes *int64 `json:"ResultBytes,omitempty" xml:"ResultBytes,omitempty"`
	// The type of the slow query logs.
	//
	// example:
	//
	// ExceptionWhileProcessing
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetInitialAddress() *string {
	return s.InitialAddress
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetQuery() *string {
	return s.Query
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetQueryDurationMs() *int64 {
	return s.QueryDurationMs
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetReadBytes() *int64 {
	return s.ReadBytes
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetReadRows() *int64 {
	return s.ReadRows
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetResultBytes() *int64 {
	return s.ResultBytes
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) GetType() *string {
	return s.Type
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialAddress(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialQueryId(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetInitialUser(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetMemoryUsage(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQuery(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQueryDurationMs(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetReadBytes(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ReadBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetReadRows(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ReadRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetResultBytes(v int64) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.ResultBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodyDataResultSet {
	s.Type = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyDataResultSet) Validate() error {
	return dara.Validate(s)
}
