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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database. By default, all databases are queried. You can also enter a database name to query. Only one database name can be entered at a time.
	//
	// example:
	//
	// Database
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end time of the query. The end time must be later than the start time, and the interval between the start time and end time must be 7 days or less. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > If DAS Enterprise Edition V3 is activated and you use the SQL Explorer and Audit feature it provides, you can query data within the hot data storage duration. You can call [DescribeSqlLogConfig](https://help.aliyun.com/document_detail/2778837.html) to query the activated Enterprise Edition information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-06T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to generate an audit file or return a list of SQL records. Valid values:
	//
	// 	- **File**: If you set this parameter to File, an audit file is generated. Only common parameters are returned. You must call the DescribeSQLLogFiles operation to obtain the download URL of the file.
	//
	// 	- **Stream**: This is the default value. A list of SQL records is returned.
	//
	// > If this parameter is set to **File**, only MySQL (with Premium Local SSDs) and SQL Server instances are supported, and a maximum of 1,000,000 log entries are recorded.
	//
	// example:
	//
	// Stream
	Form         *string `json:"Form,omitempty" xml:"Form,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type.
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
	// The keywords that are used for the query.
	//
	// - When you generate an audit file by calling this operation (the **Form*	- request parameter is set to **File**), keyword-based filtering is not supported.
	//
	// - Separate multiple keywords with spaces. You can specify up to 10 keywords. The logical relationship among keywords is **and**.
	//
	// - If a field name in the SQL statement uses backticks (\\`), you must also include the backticks when using the field name as a keyword. For example, if the field name is \\`id\\`, enter \\`id\\` instead of id.
	//
	// > After you enter keywords, the system matches the keywords against the **Database**, **User**, and **QueryKeywords*	- parameters simultaneously. The logical relationship among the three request parameters is **and**.
	//
	// example:
	//
	// table_name
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	SQLId *int64 `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// The start time of the query. You can query data within the last 7 days from the current date. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > If DAS Enterprise Edition V3 is activated and you use the SQL Explorer and Audit feature it provides, you can query data within the hot data storage duration. You can call [DescribeSqlLogConfig](https://help.aliyun.com/document_detail/2778837.html) to query the activated Enterprise Edition information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-06-01T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username. By default, all users are queried. You can also enter a username to query. Only one username can be entered at a time.
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
