// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveSQLRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeActiveSQLRecordsResponseBody
	GetDBInstanceId() *string
	SetQueries(v []*DescribeActiveSQLRecordsResponseBodyQueries) *DescribeActiveSQLRecordsResponseBody
	GetQueries() []*DescribeActiveSQLRecordsResponseBodyQueries
	SetRequestId(v string) *DescribeActiveSQLRecordsResponseBody
	GetRequestId() *string
}

type DescribeActiveSQLRecordsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The queried SQL records.
	Queries []*DescribeActiveSQLRecordsResponseBodyQueries `json:"Queries,omitempty" xml:"Queries,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActiveSQLRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveSQLRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveSQLRecordsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeActiveSQLRecordsResponseBody) GetQueries() []*DescribeActiveSQLRecordsResponseBodyQueries {
	return s.Queries
}

func (s *DescribeActiveSQLRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveSQLRecordsResponseBody) SetDBInstanceId(v string) *DescribeActiveSQLRecordsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBody) SetQueries(v []*DescribeActiveSQLRecordsResponseBodyQueries) *DescribeActiveSQLRecordsResponseBody {
	s.Queries = v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBody) SetRequestId(v string) *DescribeActiveSQLRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveSQLRecordsResponseBodyQueries struct {
	// The IP address of the client.
	//
	// example:
	//
	// 0.0.0.0
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// xg_analyse
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The progress ID.
	//
	// example:
	//
	// 6164
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
	// The SQL statement of the query.
	//
	// example:
	//
	// Select 	- from t1,t2 where t1.id=t2.id;
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The execution duration of the query. Unit: seconds.
	//
	// example:
	//
	// 60s
	QueryDuration *string `json:"QueryDuration,omitempty" xml:"QueryDuration,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2022-05-07T06:59Z
	QueryStart *string `json:"QueryStart,omitempty" xml:"QueryStart,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 6612
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// Indicates whether the SQL statement is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	SqlTruncated *string `json:"SqlTruncated,omitempty" xml:"SqlTruncated,omitempty"`
	// The threshold that is used to truncate the SQL statement.
	//
	// example:
	//
	// 1024
	SqlTruncatedThreshold *string `json:"SqlTruncatedThreshold,omitempty" xml:"SqlTruncatedThreshold,omitempty"`
	// The status of the asynchronous request. Valid values:
	//
	// 	- **running**
	//
	// 	- **block**
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testuser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeActiveSQLRecordsResponseBodyQueries) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveSQLRecordsResponseBodyQueries) GoString() string {
	return s.String()
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetDatabase() *string {
	return s.Database
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetPID() *string {
	return s.PID
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetQuery() *string {
	return s.Query
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetQueryDuration() *string {
	return s.QueryDuration
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetQueryStart() *string {
	return s.QueryStart
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetSessionID() *string {
	return s.SessionID
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetSqlTruncated() *string {
	return s.SqlTruncated
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetSqlTruncatedThreshold() *string {
	return s.SqlTruncatedThreshold
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetState() *string {
	return s.State
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) GetUser() *string {
	return s.User
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetClientAddr(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.ClientAddr = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetDatabase(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.Database = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetPID(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.PID = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetQuery(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.Query = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetQueryDuration(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.QueryDuration = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetQueryStart(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.QueryStart = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetSessionID(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.SessionID = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetSqlTruncated(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.SqlTruncated = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetSqlTruncatedThreshold(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.SqlTruncatedThreshold = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetState(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.State = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) SetUser(v string) *DescribeActiveSQLRecordsResponseBodyQueries {
	s.User = &v
	return s
}

func (s *DescribeActiveSQLRecordsResponseBodyQueries) Validate() error {
	return dara.Validate(s)
}
