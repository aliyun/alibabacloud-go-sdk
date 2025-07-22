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
	// Specifies whether to sort the returned entries in ascending order. Default value: **false**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. You can view the data of up to seven days within the last month.
	//
	// example:
	//
	// 1679297005999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf61swc4cru0b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keywords of the SQL template. Separate multiple keywords with spaces.
	//
	// example:
	//
	// select update
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The node ID.
	//
	// >  This parameter is required if the database instance is an ApsaraDB RDS for MySQL Cluster Edition instance or a PolarDB for MySQL cluster.
	//
	// example:
	//
	// r-****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The field by which to sort the returned entries. Default value: **count**.
	//
	// 	- **count**: the number of executions.
	//
	// 	- **avgRt**: the average execution duration.
	//
	// 	- **rtRate**: the execution duration percentage.
	//
	// 	- **rowsExamined**: the total number of scanned rows.
	//
	// 	- **avgRowsExamined**: the average number of scanned rows.
	//
	// 	- **avgRowsReturned**: the average number of returned rows.
	//
	// example:
	//
	// count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The SQL ID.
	//
	// >  If this parameter is specified, the full request statistics of the specified SQL query are collected. If this parameter is left empty, the full request statistics of the entire database instance are collected.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
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
