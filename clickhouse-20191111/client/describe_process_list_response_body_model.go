// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcessList(v *DescribeProcessListResponseBodyProcessList) *DescribeProcessListResponseBody
	GetProcessList() *DescribeProcessListResponseBodyProcessList
	SetRequestId(v string) *DescribeProcessListResponseBody
	GetRequestId() *string
}

type DescribeProcessListResponseBody struct {
	// The queries.
	ProcessList *DescribeProcessListResponseBodyProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FD61BB0D-788A-5185-A8E3-1B90BA8F6F04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) GetProcessList() *DescribeProcessListResponseBodyProcessList {
	return s.ProcessList
}

func (s *DescribeProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProcessListResponseBody) SetProcessList(v *DescribeProcessListResponseBodyProcessList) *DescribeProcessListResponseBody {
	s.ProcessList = v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessListResponseBody) Validate() error {
	if s.ProcessList != nil {
		if err := s.ProcessList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProcessListResponseBodyProcessList struct {
	// The details of the query.
	Data *DescribeProcessListResponseBodyProcessListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of rows returned for the query.
	//
	// example:
	//
	// 1145700
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	RowsBeforeLimitAtLeast *string `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	// The statistics of the results.
	Statistics *DescribeProcessListResponseBodyProcessListStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// Details of the columns.
	TableSchema *DescribeProcessListResponseBodyProcessListTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeProcessListResponseBodyProcessList) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessList) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessList) GetData() *DescribeProcessListResponseBodyProcessListData {
	return s.Data
}

func (s *DescribeProcessListResponseBodyProcessList) GetRows() *string {
	return s.Rows
}

func (s *DescribeProcessListResponseBodyProcessList) GetRowsBeforeLimitAtLeast() *string {
	return s.RowsBeforeLimitAtLeast
}

func (s *DescribeProcessListResponseBodyProcessList) GetStatistics() *DescribeProcessListResponseBodyProcessListStatistics {
	return s.Statistics
}

func (s *DescribeProcessListResponseBodyProcessList) GetTableSchema() *DescribeProcessListResponseBodyProcessListTableSchema {
	return s.TableSchema
}

func (s *DescribeProcessListResponseBodyProcessList) SetData(v *DescribeProcessListResponseBodyProcessListData) *DescribeProcessListResponseBodyProcessList {
	s.Data = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetRows(v string) *DescribeProcessListResponseBodyProcessList {
	s.Rows = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetRowsBeforeLimitAtLeast(v string) *DescribeProcessListResponseBodyProcessList {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetStatistics(v *DescribeProcessListResponseBodyProcessListStatistics) *DescribeProcessListResponseBodyProcessList {
	s.Statistics = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetTableSchema(v *DescribeProcessListResponseBodyProcessListTableSchema) *DescribeProcessListResponseBodyProcessList {
	s.TableSchema = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) Validate() error {
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

type DescribeProcessListResponseBodyProcessListData struct {
	ResultSet []*DescribeProcessListResponseBodyProcessListDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyProcessListData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListData) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListData) GetResultSet() []*DescribeProcessListResponseBodyProcessListDataResultSet {
	return s.ResultSet
}

func (s *DescribeProcessListResponseBodyProcessListData) SetResultSet(v []*DescribeProcessListResponseBodyProcessListDataResultSet) *DescribeProcessListResponseBodyProcessListData {
	s.ResultSet = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListData) Validate() error {
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

type DescribeProcessListResponseBodyProcessListDataResultSet struct {
	// The IP address of the client that initiates the query.
	//
	// example:
	//
	// ::ffff:10.1.XX.XX
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 2dd144fd-4230-4249-b15c-e63f964fbb5a
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The database account.
	//
	// example:
	//
	// test
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The SQL statement that is executed in the query.
	//
	// example:
	//
	// select 	- from test order by score limit 1;
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The value is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2021-02-02T09:14:48Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetInitialAddress() *string {
	return s.InitialAddress
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetQuery() *string {
	return s.Query
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetQueryDurationMs() *string {
	return s.QueryDurationMs
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialAddress(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialQueryId(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialUser(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQuery(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQueryDurationMs(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQueryStartTime(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) Validate() error {
	return dara.Validate(s)
}

type DescribeProcessListResponseBodyProcessListStatistics struct {
	// The size of the data that was scanned. Unit: bytes.
	//
	// example:
	//
	// 9141300000
	BytesRead *int32 `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	// The average response time.
	//
	// example:
	//
	// 4156
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The number of scanned rows.
	//
	// example:
	//
	// 1000000
	RowsRead *int32 `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListStatistics) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) GetBytesRead() *int32 {
	return s.BytesRead
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) GetElapsedTime() *float32 {
	return s.ElapsedTime
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) GetRowsRead() *int32 {
	return s.RowsRead
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetBytesRead(v int32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetElapsedTime(v float32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetRowsRead(v int32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.RowsRead = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeProcessListResponseBodyProcessListTableSchema struct {
	ResultSet []*DescribeProcessListResponseBodyProcessListTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyProcessListTableSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListTableSchema) GetResultSet() []*DescribeProcessListResponseBodyProcessListTableSchemaResultSet {
	return s.ResultSet
}

func (s *DescribeProcessListResponseBodyProcessListTableSchema) SetResultSet(v []*DescribeProcessListResponseBodyProcessListTableSchemaResultSet) *DescribeProcessListResponseBodyProcessListTableSchema {
	s.ResultSet = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListTableSchema) Validate() error {
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

type DescribeProcessListResponseBodyProcessListTableSchemaResultSet struct {
	// The column name.
	//
	// example:
	//
	// InitialUser
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The column type.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListTableSchemaResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) GetName() *string {
	return s.Name
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) GetType() *string {
	return s.Type
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) SetName(v string) *DescribeProcessListResponseBodyProcessListTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) SetType(v string) *DescribeProcessListResponseBodyProcessListTableSchemaResultSet {
	s.Type = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) Validate() error {
	return dara.Validate(s)
}
