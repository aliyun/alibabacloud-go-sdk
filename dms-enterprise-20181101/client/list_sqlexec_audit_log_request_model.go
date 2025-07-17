// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLExecAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListSQLExecAuditLogRequest
	GetEndTime() *string
	SetExecState(v string) *ListSQLExecAuditLogRequest
	GetExecState() *string
	SetOpUserName(v string) *ListSQLExecAuditLogRequest
	GetOpUserName() *string
	SetPageNumber(v int32) *ListSQLExecAuditLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSQLExecAuditLogRequest
	GetPageSize() *int32
	SetSearchName(v string) *ListSQLExecAuditLogRequest
	GetSearchName() *string
	SetSqlType(v string) *ListSQLExecAuditLogRequest
	GetSqlType() *string
	SetStartTime(v string) *ListSQLExecAuditLogRequest
	GetStartTime() *string
	SetTid(v int64) *ListSQLExecAuditLogRequest
	GetTid() *int64
}

type ListSQLExecAuditLogRequest struct {
	// The end of the time range to query.
	//
	// >  The end time supports fuzzy match. Specify the time in the YYYY-MM-DD hh:mm:ss format. We recommend that you use the StartTime and EndTime parameters to specify a time range that does not exceed one day. The returned entries can be displayed by page to improve query efficiency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-08 11:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- **FAIL**: The SQL statement fails to be executed.
	//
	// 	- **NOEXE**: The SQL statement has not been executed.
	//
	// 	- **RUNNING**: The SQL statement is being executed.
	//
	// 	- **CANCEL**: The execution of the SQL statement is canceled.
	//
	// 	- **SUCCESS**: The SQL statement is executed.
	//
	// example:
	//
	// SUCCESS
	ExecState *string `json:"ExecState,omitempty" xml:"ExecState,omitempty"`
	// The nickname of the user who wrote the SQL statement.
	//
	// example:
	//
	// test_OpUserName
	OpUserName *string `json:"OpUserName,omitempty" xml:"OpUserName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The value cannot exceed 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the database or instance based on which you want to query SQL statements.
	//
	// >  If the SQL statements to be queried are at the instance level, you can set this parameter to an instance name. If the SQL statements to be queried are at the database level, you can set this parameter to a database name.
	//
	// example:
	//
	// test_SearchName
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// 	- **SELECT**: the SQL statement that is used to query data.
	//
	// 	- **INSERT**: the SQL statement that is used to insert data.
	//
	// 	- **DELETE**: the SQL statement that is used to delete data.
	//
	// 	- **CREATE_TABLE**: the SQL statement that is used to create tables.
	//
	// >  To view more types of SQL statements, log on to the DMS console and click Security and Specifications. In the left-side navigation pane, click **Operation Audit**. Then, you can view all supported types of SQL statements from the **SQL type*	- drop-down list.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query.
	//
	// >  The start time supports fuzzy match. Specify the time in the YYYY-MM-DD hh:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-08 11:04:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSQLExecAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSQLExecAuditLogRequest) GetExecState() *string {
	return s.ExecState
}

func (s *ListSQLExecAuditLogRequest) GetOpUserName() *string {
	return s.OpUserName
}

func (s *ListSQLExecAuditLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSQLExecAuditLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSQLExecAuditLogRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListSQLExecAuditLogRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *ListSQLExecAuditLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSQLExecAuditLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSQLExecAuditLogRequest) SetEndTime(v string) *ListSQLExecAuditLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetExecState(v string) *ListSQLExecAuditLogRequest {
	s.ExecState = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetOpUserName(v string) *ListSQLExecAuditLogRequest {
	s.OpUserName = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetPageNumber(v int32) *ListSQLExecAuditLogRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetPageSize(v int32) *ListSQLExecAuditLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetSearchName(v string) *ListSQLExecAuditLogRequest {
	s.SearchName = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetSqlType(v string) *ListSQLExecAuditLogRequest {
	s.SqlType = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetStartTime(v string) *ListSQLExecAuditLogRequest {
	s.StartTime = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) SetTid(v int64) *ListSQLExecAuditLogRequest {
	s.Tid = &v
	return s
}

func (s *ListSQLExecAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
