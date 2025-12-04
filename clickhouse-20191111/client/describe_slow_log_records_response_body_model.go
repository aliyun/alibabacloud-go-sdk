// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlowLogRecordsResponseBody
	GetRequestId() *string
	SetSlowLogRecords(v *DescribeSlowLogRecordsResponseBodySlowLogRecords) *DescribeSlowLogRecordsResponseBody
	GetSlowLogRecords() *DescribeSlowLogRecordsResponseBodySlowLogRecords
}

type DescribeSlowLogRecordsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DF203CC8-5F68-5E3F-8050-3C77DD65731A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the slow query logs.
	SlowLogRecords *DescribeSlowLogRecordsResponseBodySlowLogRecords `json:"SlowLogRecords,omitempty" xml:"SlowLogRecords,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogRecordsResponseBody) GetSlowLogRecords() *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	return s.SlowLogRecords
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetSlowLogRecords(v *DescribeSlowLogRecordsResponseBodySlowLogRecords) *DescribeSlowLogRecordsResponseBody {
	s.SlowLogRecords = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) Validate() error {
	if s.SlowLogRecords != nil {
		if err := s.SlowLogRecords.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodySlowLogRecords struct {
	// Details about the slow query logs.
	Data *DescribeSlowLogRecordsResponseBodySlowLogRecordsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of rows in the result set.
	//
	// example:
	//
	// 1
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	RowsBeforeLimitAtLeast *string `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	// The statistics of the results.
	Statistics *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The schema of the table in the database.
	TableSchema *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) GetData() *DescribeSlowLogRecordsResponseBodySlowLogRecordsData {
	return s.Data
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) GetRows() *string {
	return s.Rows
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) GetRowsBeforeLimitAtLeast() *string {
	return s.RowsBeforeLimitAtLeast
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) GetStatistics() *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	return s.Statistics
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) GetTableSchema() *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema {
	return s.TableSchema
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetData(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Data = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetRows(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetRowsBeforeLimitAtLeast(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetStatistics(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Statistics = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetTableSchema(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.TableSchema = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsData struct {
	ResultSet []*DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) GetResultSet() []*DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) SetResultSet(v []*DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) *DescribeSlowLogRecordsResponseBodySlowLogRecordsData {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) Validate() error {
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

type DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet struct {
	// The IP address of the client that initiated the query.
	//
	// example:
	//
	// ::ffff:100.104.XX.XX
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"b51496f2-6b0b-4546-aff9-e17951cb9410\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The username that is used to initiate the query.
	//
	// example:
	//
	// test_users
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The peak memory usage for the query. Unit: bytes.
	//
	// example:
	//
	// 1048576
	MemoryUsage *string `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// The statement that was executed in the query.
	//
	// example:
	//
	// Select 	- from table
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-22 20:00:01
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The size of the data read by executing the statement. Unit: bytes.
	//
	// example:
	//
	// 1048576
	ReadBytes *string `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// The number of rows read by executing the statement.
	//
	// example:
	//
	// 10027008
	ReadRows *string `json:"ReadRows,omitempty" xml:"ReadRows,omitempty"`
	// The size of the result data. Unit: bytes.
	//
	// example:
	//
	// 1024
	ResultBytes *string `json:"ResultBytes,omitempty" xml:"ResultBytes,omitempty"`
	// The query status. Valid values:
	//
	// 	- **QueryFinish**: The query is complete.
	//
	// 	- **Processing**: The query is running.
	//
	// example:
	//
	// QueryFinish
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetInitialAddress() *string {
	return s.InitialAddress
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetMemoryUsage() *string {
	return s.MemoryUsage
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetQuery() *string {
	return s.Query
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetQueryDurationMs() *string {
	return s.QueryDurationMs
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetReadBytes() *string {
	return s.ReadBytes
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetReadRows() *string {
	return s.ReadRows
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetResultBytes() *string {
	return s.ResultBytes
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GetType() *string {
	return s.Type
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialAddress(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialQueryId(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialUser(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQuery(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQueryDurationMs(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetReadBytes(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ReadBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetReadRows(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ReadRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetResultBytes(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ResultBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.Type = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics struct {
	// The total size of data that were read. Unit: bytes.
	//
	// example:
	//
	// 123456
	BytesRead *int32 `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	// The time consumed by the slow query. Unit: milliseconds.
	//
	// example:
	//
	// 21.35
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The total number of rows that were read.
	//
	// example:
	//
	// 2016722
	RowsRead *int32 `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) GetBytesRead() *int32 {
	return s.BytesRead
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) GetElapsedTime() *float32 {
	return s.ElapsedTime
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) GetRowsRead() *int32 {
	return s.RowsRead
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetBytesRead(v int32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetElapsedTime(v float32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetRowsRead(v int32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.RowsRead = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema struct {
	ResultSet []*DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) GetResultSet() []*DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) SetResultSet(v []*DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) Validate() error {
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

type DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet struct {
	// The name of the column.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) GetName() *string {
	return s.Name
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) GetType() *string {
	return s.Type
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) SetName(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet {
	s.Type = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) Validate() error {
	return dara.Validate(s)
}
