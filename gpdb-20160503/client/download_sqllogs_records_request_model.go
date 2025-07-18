// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSQLLogsRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DownloadSQLLogsRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DownloadSQLLogsRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DownloadSQLLogsRecordsRequest
	GetEndTime() *string
	SetExecuteCost(v string) *DownloadSQLLogsRecordsRequest
	GetExecuteCost() *string
	SetExecuteState(v string) *DownloadSQLLogsRecordsRequest
	GetExecuteState() *string
	SetLang(v string) *DownloadSQLLogsRecordsRequest
	GetLang() *string
	SetMaxExecuteCost(v string) *DownloadSQLLogsRecordsRequest
	GetMaxExecuteCost() *string
	SetMinExecuteCost(v string) *DownloadSQLLogsRecordsRequest
	GetMinExecuteCost() *string
	SetOperationClass(v string) *DownloadSQLLogsRecordsRequest
	GetOperationClass() *string
	SetOperationType(v string) *DownloadSQLLogsRecordsRequest
	GetOperationType() *string
	SetPageNumber(v int32) *DownloadSQLLogsRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DownloadSQLLogsRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DownloadSQLLogsRecordsRequest
	GetQueryKeywords() *string
	SetSourceIP(v string) *DownloadSQLLogsRecordsRequest
	GetSourceIP() *string
	SetStartTime(v string) *DownloadSQLLogsRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DownloadSQLLogsRecordsRequest
	GetUser() *string
}

type DownloadSQLLogsRecordsRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-08T06:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution duration of the SQL statement. Unit: seconds.
	//
	// example:
	//
	// 1
	ExecuteCost *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution status of the SQL statement.
	//
	// 	- **1**: successful.
	//
	// 	- **0**: failed.
	//
	// example:
	//
	// 1
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The language of the file that contains the query diagnostic information. Valid values:
	//
	// 	- **zh**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum amount of time consumed by a slow query. Unit: seconds. Minimum value: 0.
	//
	// example:
	//
	// 999
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	// The minimum amount of time consumed by a slow query. Unit: seconds. Minimum value: 0.
	//
	// example:
	//
	// 1
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	// The type of the query language. Example: DQL, DML, or DDL.
	//
	// example:
	//
	// DQL
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The type of the SQL statement. Example: SELECT.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords that are used for query.
	//
	// example:
	//
	// select 1
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 100.XX.XX.90
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-07T06:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testuser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DownloadSQLLogsRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSQLLogsRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadSQLLogsRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DownloadSQLLogsRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DownloadSQLLogsRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadSQLLogsRecordsRequest) GetExecuteCost() *string {
	return s.ExecuteCost
}

func (s *DownloadSQLLogsRecordsRequest) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DownloadSQLLogsRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadSQLLogsRecordsRequest) GetMaxExecuteCost() *string {
	return s.MaxExecuteCost
}

func (s *DownloadSQLLogsRecordsRequest) GetMinExecuteCost() *string {
	return s.MinExecuteCost
}

func (s *DownloadSQLLogsRecordsRequest) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DownloadSQLLogsRecordsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DownloadSQLLogsRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DownloadSQLLogsRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DownloadSQLLogsRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DownloadSQLLogsRecordsRequest) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DownloadSQLLogsRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadSQLLogsRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DownloadSQLLogsRecordsRequest) SetDBInstanceId(v string) *DownloadSQLLogsRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetDatabase(v string) *DownloadSQLLogsRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetEndTime(v string) *DownloadSQLLogsRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetExecuteState(v string) *DownloadSQLLogsRecordsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetLang(v string) *DownloadSQLLogsRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetMaxExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetMinExecuteCost(v string) *DownloadSQLLogsRecordsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetOperationClass(v string) *DownloadSQLLogsRecordsRequest {
	s.OperationClass = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetOperationType(v string) *DownloadSQLLogsRecordsRequest {
	s.OperationType = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetPageNumber(v int32) *DownloadSQLLogsRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetPageSize(v int32) *DownloadSQLLogsRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetQueryKeywords(v string) *DownloadSQLLogsRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetSourceIP(v string) *DownloadSQLLogsRecordsRequest {
	s.SourceIP = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetStartTime(v string) *DownloadSQLLogsRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) SetUser(v string) *DownloadSQLLogsRecordsRequest {
	s.User = &v
	return s
}

func (s *DownloadSQLLogsRecordsRequest) Validate() error {
	return dara.Validate(s)
}
