// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAuditLogRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetAccessDeniedDetail() *string
	SetDBClusterId(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetDBClusterId() *string
	SetItems(v []*DescribeSparkAuditLogRecordsResponseBodyItems) *DescribeSparkAuditLogRecordsResponseBody
	GetItems() []*DescribeSparkAuditLogRecordsResponseBodyItems
	SetPageNumber(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSparkAuditLogRecordsResponseBody
	GetTotalCount() *string
}

type DescribeSparkAuditLogRecordsResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp1j7******78j8i
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The list of SQL audit logs.
	Items []*DescribeSparkAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D5867CF-423F-559F-BBB1-199A289E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSparkAuditLogRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetItems() []*DescribeSparkAuditLogRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkAuditLogRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetAccessDeniedDetail(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetItems(v []*DescribeSparkAuditLogRecordsResponseBodyItems) *DescribeSparkAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetPageSize(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetRequestId(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeSparkAuditLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSparkAuditLogRecordsResponseBodyItems struct {
	// The Spark application ID.
	//
	// example:
	//
	// s202411061011hzc5d6476000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The SQL execution error message.
	//
	// example:
	//
	// notFoundIp
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The SQL execution error stack trace.
	//
	// example:
	//
	// notFoundIpException
	ErrorTrace *string `json:"ErrorTrace,omitempty" xml:"ErrorTrace,omitempty"`
	// The start time of the SQL statement. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2022-01-23T16:05:08Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The ID of the query executed within the Spark application.
	//
	// example:
	//
	// 1
	InnerQueryId *string `json:"InnerQueryId,omitempty" xml:"InnerQueryId,omitempty"`
	// Whether it can be diagnosed.
	//
	// example:
	//
	// true
	IsDiagnosable *bool `json:"IsDiagnosable,omitempty" xml:"IsDiagnosable,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 999f2439-6b10-xxxx-a5d3-daf3b35c****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// test_job
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM adb_hdfs_import_source
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The ID of the statement.
	//
	// example:
	//
	// fbd22066-1c03-xxxx-aa16-6ae28288****
	StatementId *string `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
	// The source from which the query was initiated.
	//
	// Valid values:
	//
	// 	- SQL_EDITOR: SQL_EDITOR.
	//
	// 	- JDBC: JDBC.
	//
	// example:
	//
	// SQL_EDITOR
	StatementSource *string `json:"StatementSource,omitempty" xml:"StatementSource,omitempty"`
	// The execution status of the SQL statement.
	//
	// Valid values:
	//
	// 	- cancel: The task is canceled .
	//
	// 	- finished: The execution succeeds .
	//
	// 	- error: The execution fails .
	//
	// 	- timeout: The execution of the command timed out.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The duration of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 40000
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The username that is used to execute SQL statements.
	//
	// example:
	//
	// test_user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSparkAuditLogRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetErrorTrace() *string {
	return s.ErrorTrace
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetInnerQueryId() *string {
	return s.InnerQueryId
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetIsDiagnosable() *bool {
	return s.IsDiagnosable
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetStatementId() *string {
	return s.StatementId
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetStatementSource() *string {
	return s.StatementSource
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetAppId(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.AppId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetClientIp(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ClientIp = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetErrorMsg(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetErrorTrace(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ErrorTrace = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetInnerQueryId(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.InnerQueryId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetIsDiagnosable(v bool) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.IsDiagnosable = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetProcessId(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ProcessId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetResourceGroupName(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetStatementId(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.StatementId = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetStatementSource(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.StatementSource = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetStatus(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetTotalTime(v int64) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) SetUser(v string) *DescribeSparkAuditLogRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
