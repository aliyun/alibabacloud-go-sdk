// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeSQLLogRecordsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeSQLLogRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeSQLLogRecordsRequest
	GetEndTime() *string
	SetForm(v string) *DescribeSQLLogRecordsRequest
	GetForm() *string
	SetOwnerAccount(v string) *DescribeSQLLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSQLLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLLogRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest
	GetQueryKeywords() *string
	SetResourceOwnerAccount(v string) *DescribeSQLLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLLogRecordsRequest
	GetResourceOwnerId() *int64
	SetSQLId(v int64) *DescribeSQLLogRecordsRequest
	GetSQLId() *int64
	SetStartTime(v string) *DescribeSQLLogRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeSQLLogRecordsRequest
	GetUser() *string
}

type DescribeSQLLogRecordsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database. You can enter only one database name. If you specify this parameter, this operation returns the logs that are generated only for the specified database. If you do not specify this parameter, this operation returns the logs that are generated for all databases on the instance.
	//
	// example:
	//
	// Database
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span between the start time and the end time must be less than 15 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-11T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to generate an SQL audit log file or return SQL audit logs. Valid values:
	//
	// 	- **File**: If you set this parameter to File, this operation generates an SQL audit log file and returns only common response parameters. After you call this operation, you must call the DescribeSQLLogFiles operation to obtain the download URL of the SQL audit log file.
	//
	// 	- **Stream*	- (default): If you set this parameter to Stream, this operation returns SQL audit logs.
	//
	// >  If you set this parameter to **File**, only ApsaraDB RDS for MySQL instances that use local disks and ApsaraDB RDS for SQL Server instances are supported, and a maximum of 1 million logs are returned.
	//
	// example:
	//
	// Stream
	Form         *string `json:"Form,omitempty" xml:"Form,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used for the query.
	//
	// 	- When you call this operation and set the **Form*	- parameter to **File*	- to generate an audit file, you cannot filter log entries by keyword.
	//
	// 	- You can specify up to 10 keywords. The keywords are evaluated by using the **AND*	- operator. Separate multiple keywords with spaces.
	//
	// 	- If a field name in the specified SQL statement is enclosed in grave accents (\\`) and you want to use the field name as a keyword, you must enter the grave accents (\\`) as part of the field name. For example, if the field name is \\`id\\`, enter \\`id\\` instead of id.
	//
	// >  After you enter a keyword, the system matches the keyword based on the **Database**, **User**, and **QueryKeywords*	- parameters. The parameters are evaluated by using the **AND*	- operator.
	//
	// example:
	//
	// table_name
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the SQL statement.
	//
	// example:
	//
	// 25623548
	SQLId *int64 `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// The beginning of the time range to query. You can query data in the last 15 days before the current date. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-01T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username of the account. You can enter only one username. If you specify this parameter, this operation returns the logs that are generated only for the specified account. If you do not specify this parameter, this operation returns the logs that are generated for all accounts on the instance.
	//
	// example:
	//
	// user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSQLLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSQLLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogRecordsRequest) GetForm() *string {
	return s.Form
}

func (s *DescribeSQLLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLLogRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSQLLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLLogRecordsRequest) GetSQLId() *int64 {
	return s.SQLId
}

func (s *DescribeSQLLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeSQLLogRecordsRequest) SetClientToken(v string) *DescribeSQLLogRecordsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDatabase(v string) *DescribeSQLLogRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetEndTime(v string) *DescribeSQLLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetForm(v string) *DescribeSQLLogRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetOwnerAccount(v string) *DescribeSQLLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetOwnerId(v int64) *DescribeSQLLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageNumber(v int32) *DescribeSQLLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageSize(v int32) *DescribeSQLLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSQLLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSQLLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetSQLId(v int64) *DescribeSQLLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetStartTime(v string) *DescribeSQLLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetUser(v string) *DescribeSQLLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
