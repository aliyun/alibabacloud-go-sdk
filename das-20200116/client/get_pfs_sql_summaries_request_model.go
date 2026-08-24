// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsSqlSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GetPfsSqlSummariesRequest
	GetAsc() *bool
	SetEndTime(v int64) *GetPfsSqlSummariesRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetPfsSqlSummariesRequest
	GetInstanceId() *string
	SetKeywords(v string) *GetPfsSqlSummariesRequest
	GetKeywords() *string
	SetNodeId(v string) *GetPfsSqlSummariesRequest
	GetNodeId() *string
	SetOrderBy(v string) *GetPfsSqlSummariesRequest
	GetOrderBy() *string
	SetPageNo(v int32) *GetPfsSqlSummariesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetPfsSqlSummariesRequest
	GetPageSize() *int32
	SetSqlId(v string) *GetPfsSqlSummariesRequest
	GetSqlId() *string
	SetStartTime(v int64) *GetPfsSqlSummariesRequest
	GetStartTime() *int64
}

type GetPfsSqlSummariesRequest struct {
	// Sort in ascending order. Default is **false**.
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// End time of the query, in Unix timestamp format, in milliseconds.
	//
	// > The end time must be later than the start time. You can query data for any seven-day period within the last 30 days.
	//
	// example:
	//
	// 1679297005999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Database instance ID.
	//
	// example:
	//
	// rm-uf61swc4cru0b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Keywords in the SQL text. Separate multiple keywords with spaces.
	//
	// example:
	//
	// select update
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Node ID.
	//
	// > Provide a node ID for RDS MySQL Cluster Edition or PolarDB for MySQL database instances.
	//
	// example:
	//
	// r-****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Field to sort by. Default is **count**.
	//
	// - **count**: Number of executions.
	//
	// - **avgLatency**: Average execution time.
	//
	// - **rowsExamined**: Total number of scanned rows.
	//
	// - **rowsExaminedAvg**: Average number of scanned rows.
	//
	// - **rowsSentAvg**: Average number of returned rows.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// Page number for paged queries. Start from 1. Default is 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Maximum number of records per page for paged queries. Default is 10. Maximum is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// SQL ID.
	//
	// > If you specify an SQL ID, the system returns statistics for that SQL ID only. If you leave this parameter empty, the system returns statistics for the entire database instance.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// Start time of the query, in Unix timestamp format, in milliseconds.
	//
	// example:
	//
	// 1675833788056
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPfsSqlSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPfsSqlSummariesRequest) GoString() string {
	return s.String()
}

func (s *GetPfsSqlSummariesRequest) GetAsc() *bool {
	return s.Asc
}

func (s *GetPfsSqlSummariesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPfsSqlSummariesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPfsSqlSummariesRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *GetPfsSqlSummariesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPfsSqlSummariesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetPfsSqlSummariesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetPfsSqlSummariesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPfsSqlSummariesRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetPfsSqlSummariesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPfsSqlSummariesRequest) SetAsc(v bool) *GetPfsSqlSummariesRequest {
	s.Asc = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetEndTime(v int64) *GetPfsSqlSummariesRequest {
	s.EndTime = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetInstanceId(v string) *GetPfsSqlSummariesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetKeywords(v string) *GetPfsSqlSummariesRequest {
	s.Keywords = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetNodeId(v string) *GetPfsSqlSummariesRequest {
	s.NodeId = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetOrderBy(v string) *GetPfsSqlSummariesRequest {
	s.OrderBy = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetPageNo(v int32) *GetPfsSqlSummariesRequest {
	s.PageNo = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetPageSize(v int32) *GetPfsSqlSummariesRequest {
	s.PageSize = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetSqlId(v string) *GetPfsSqlSummariesRequest {
	s.SqlId = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) SetStartTime(v int64) *GetPfsSqlSummariesRequest {
	s.StartTime = &v
	return s
}

func (s *GetPfsSqlSummariesRequest) Validate() error {
	return dara.Validate(s)
}
