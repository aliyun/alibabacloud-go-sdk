// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlowLogTrendResponseBody
	GetRequestId() *string
	SetSlowLogTrend(v *DescribeSlowLogTrendResponseBodySlowLogTrend) *DescribeSlowLogTrendResponseBody
	GetSlowLogTrend() *DescribeSlowLogTrendResponseBodySlowLogTrend
}

type DescribeSlowLogTrendResponseBody struct {
	// example:
	//
	// 7D3ECB0E-98CA-5E08-A9CA-F70C5A1E9BDF
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowLogTrend *DescribeSlowLogTrendResponseBodySlowLogTrend `json:"SlowLogTrend,omitempty" xml:"SlowLogTrend,omitempty" type:"Struct"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowLogTrendResponseBody) GetSlowLogTrend() *DescribeSlowLogTrendResponseBodySlowLogTrend {
	return s.SlowLogTrend
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetSlowLogTrend(v *DescribeSlowLogTrendResponseBodySlowLogTrend) *DescribeSlowLogTrendResponseBody {
	s.SlowLogTrend = v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) Validate() error {
	if s.SlowLogTrend != nil {
		if err := s.SlowLogTrend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlowLogTrendResponseBodySlowLogTrend struct {
	Data *DescribeSlowLogTrendResponseBodySlowLogTrendData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 1
	RowsBeforeLimitAtLeast *string                                                  `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	Statistics             *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics  `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	TableSchema            *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrend) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrend) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) GetData() *DescribeSlowLogTrendResponseBodySlowLogTrendData {
	return s.Data
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) GetRows() *string {
	return s.Rows
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) GetRowsBeforeLimitAtLeast() *string {
	return s.RowsBeforeLimitAtLeast
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) GetStatistics() *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	return s.Statistics
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) GetTableSchema() *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema {
	return s.TableSchema
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetData(v *DescribeSlowLogTrendResponseBodySlowLogTrendData) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Data = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetRows(v string) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetRowsBeforeLimitAtLeast(v string) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetStatistics(v *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Statistics = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetTableSchema(v *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.TableSchema = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) Validate() error {
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

type DescribeSlowLogTrendResponseBodySlowLogTrendData struct {
	ResultSet []*DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendData) GetResultSet() []*DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendData) SetResultSet(v []*DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) *DescribeSlowLogTrendResponseBodySlowLogTrendData {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendData) Validate() error {
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

type DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet struct {
	// example:
	//
	// 2000
	AvgQueryDurationMs *string `json:"AvgQueryDurationMs,omitempty" xml:"AvgQueryDurationMs,omitempty"`
	// example:
	//
	// 4000
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 3000
	MaxQueryDurationMs *string `json:"MaxQueryDurationMs,omitempty" xml:"MaxQueryDurationMs,omitempty"`
	// example:
	//
	// 1000
	MinQueryDurationMs *string `json:"MinQueryDurationMs,omitempty" xml:"MinQueryDurationMs,omitempty"`
	// example:
	//
	// 2022-05-22 20:00:01
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GetAvgQueryDurationMs() *string {
	return s.AvgQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GetCount() *string {
	return s.Count
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GetMaxQueryDurationMs() *string {
	return s.MaxQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GetMinQueryDurationMs() *string {
	return s.MinQueryDurationMs
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetAvgQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.AvgQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetCount(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.Count = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetMaxQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.MaxQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetMinQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.MinQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogTrendResponseBodySlowLogTrendStatistics struct {
	// example:
	//
	// 697899
	BytesRead *int32 `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	// example:
	//
	// 0.001703578
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// example:
	//
	// 14721
	RowsRead *int32 `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) GetBytesRead() *int32 {
	return s.BytesRead
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) GetElapsedTime() *float32 {
	return s.ElapsedTime
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) GetRowsRead() *int32 {
	return s.RowsRead
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetBytesRead(v int32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetElapsedTime(v float32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetRowsRead(v int32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.RowsRead = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema struct {
	ResultSet []*DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) GetResultSet() []*DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet {
	return s.ResultSet
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) SetResultSet(v []*DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema {
	s.ResultSet = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) Validate() error {
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

type DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet struct {
	// example:
	//
	// query_start_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// UInt64
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) GetName() *string {
	return s.Name
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) GetType() *string {
	return s.Type
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) SetName(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) SetType(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet {
	s.Type = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) Validate() error {
	return dara.Validate(s)
}
