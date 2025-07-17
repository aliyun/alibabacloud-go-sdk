// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxySQLExecAuditLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListProxySQLExecAuditLogRequest
	GetEndTime() *int64
	SetExecState(v string) *ListProxySQLExecAuditLogRequest
	GetExecState() *string
	SetOpUserName(v string) *ListProxySQLExecAuditLogRequest
	GetOpUserName() *string
	SetPageNumber(v int32) *ListProxySQLExecAuditLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProxySQLExecAuditLogRequest
	GetPageSize() *int32
	SetSQLType(v string) *ListProxySQLExecAuditLogRequest
	GetSQLType() *string
	SetSearchName(v string) *ListProxySQLExecAuditLogRequest
	GetSearchName() *string
	SetStartTime(v int64) *ListProxySQLExecAuditLogRequest
	GetStartTime() *int64
	SetTid(v int64) *ListProxySQLExecAuditLogRequest
	GetTid() *int64
}

type ListProxySQLExecAuditLogRequest struct {
	// The end of the time range to query. The value of this parameter must be a timestamp that follows the UNIX time format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636962846000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- **FAIL**: The execution of the SQL statement fails.
	//
	// 	- **CANCEL**: The execution of the SQL statement is canceled.
	//
	// 	- **SUCCESS**: The SQL statement is executed.
	//
	// example:
	//
	// SUCCESS
	ExecState *string `json:"ExecState,omitempty" xml:"ExecState,omitempty"`
	// The alias of the user.
	//
	// example:
	//
	// testNickName
	OpUserName *string `json:"OpUserName,omitempty" xml:"OpUserName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum values: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of SQL statement. Valid values:
	//
	// 	- **SELECT**
	//
	// 	- **INSERT**
	//
	// 	- **DELETE**
	//
	// 	- **CREATE_TABLE**
	//
	// >  You can choose Operation Audit > Secure Access Proxy in the top navigation bar of the DMS console to view more types of SQL statements.
	//
	// example:
	//
	// SELECT
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// The name of the database instance.
	//
	// example:
	//
	// test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The beginning of the time range to query. The value of this parameter must be a timestamp that follows the UNIX time format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636876446000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 14****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListProxySQLExecAuditLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProxySQLExecAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ListProxySQLExecAuditLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListProxySQLExecAuditLogRequest) GetExecState() *string {
	return s.ExecState
}

func (s *ListProxySQLExecAuditLogRequest) GetOpUserName() *string {
	return s.OpUserName
}

func (s *ListProxySQLExecAuditLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProxySQLExecAuditLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProxySQLExecAuditLogRequest) GetSQLType() *string {
	return s.SQLType
}

func (s *ListProxySQLExecAuditLogRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListProxySQLExecAuditLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListProxySQLExecAuditLogRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListProxySQLExecAuditLogRequest) SetEndTime(v int64) *ListProxySQLExecAuditLogRequest {
	s.EndTime = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetExecState(v string) *ListProxySQLExecAuditLogRequest {
	s.ExecState = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetOpUserName(v string) *ListProxySQLExecAuditLogRequest {
	s.OpUserName = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetPageNumber(v int32) *ListProxySQLExecAuditLogRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetPageSize(v int32) *ListProxySQLExecAuditLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetSQLType(v string) *ListProxySQLExecAuditLogRequest {
	s.SQLType = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetSearchName(v string) *ListProxySQLExecAuditLogRequest {
	s.SearchName = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetStartTime(v int64) *ListProxySQLExecAuditLogRequest {
	s.StartTime = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) SetTid(v int64) *ListProxySQLExecAuditLogRequest {
	s.Tid = &v
	return s
}

func (s *ListProxySQLExecAuditLogRequest) Validate() error {
	return dara.Validate(s)
}
