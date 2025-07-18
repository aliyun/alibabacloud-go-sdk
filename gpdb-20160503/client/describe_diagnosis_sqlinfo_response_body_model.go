// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSQLInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetDatabase() *string
	SetDuration(v int32) *DescribeDiagnosisSQLInfoResponseBody
	GetDuration() *int32
	SetMaxOutputRows(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetMaxOutputRows() *string
	SetQueryID(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetQueryID() *string
	SetQueryPlan(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetQueryPlan() *string
	SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetRequestId() *string
	SetSQLStmt(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetSQLStmt() *string
	SetSessionID(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetSessionID() *string
	SetSortedMetrics(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetSortedMetrics() *string
	SetStartTime(v int64) *DescribeDiagnosisSQLInfoResponseBody
	GetStartTime() *int64
	SetStatus(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetStatus() *string
	SetTextPlan(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetTextPlan() *string
	SetUser(v string) *DescribeDiagnosisSQLInfoResponseBody
	GetUser() *string
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	// The name of the database.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The execution duration of the query. Unit: seconds.
	//
	// example:
	//
	// 16
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum number of output rows.
	//
	// example:
	//
	// 10
	MaxOutputRows *string `json:"MaxOutputRows,omitempty" xml:"MaxOutputRows,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 71403480878****
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The information about the operator.
	//
	// example:
	//
	// {\\"children\\":********\\"startTime\\":1660719602199}
	QueryPlan *string `json:"QueryPlan,omitempty" xml:"QueryPlan,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425AAA6A-63E0-1929-A1CE-3D9036CBC463
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select n_live_tup, n_live_tup + n_dead_tup, pg_relation_size(table_name), last_vacuum from pg_stat_user_tables where relid = table_name::regclass
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// The ID of the session that contains the query.
	//
	// example:
	//
	// 658****
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The sequence of metrics.
	//
	// example:
	//
	// {\\"costSort\\":******:\\"Seq Scan-9\\",\\"value\\":0.0}]}
	SortedMetrics *string `json:"SortedMetrics,omitempty" xml:"SortedMetrics,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660719602199
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the execution plan.
	//
	// example:
	//
	// ******
	TextPlan *string `json:"TextPlan,omitempty" xml:"TextPlan,omitempty"`
	// The username.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetMaxOutputRows() *string {
	return s.MaxOutputRows
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetQueryID() *string {
	return s.QueryID
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetQueryPlan() *string {
	return s.QueryPlan
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetSQLStmt() *string {
	return s.SQLStmt
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetSessionID() *string {
	return s.SessionID
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetSortedMetrics() *string {
	return s.SortedMetrics
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetTextPlan() *string {
	return s.TextPlan
}

func (s *DescribeDiagnosisSQLInfoResponseBody) GetUser() *string {
	return s.User
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDatabase(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDuration(v int32) *DescribeDiagnosisSQLInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetMaxOutputRows(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.MaxOutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetQueryPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.QueryPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSQLStmt(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSessionID(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetSortedMetrics(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.SortedMetrics = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStartTime(v int64) *DescribeDiagnosisSQLInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStatus(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetTextPlan(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.TextPlan = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetUser(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.User = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
