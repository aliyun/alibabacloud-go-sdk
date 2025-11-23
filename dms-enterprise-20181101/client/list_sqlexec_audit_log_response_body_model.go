// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLExecAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSQLExecAuditLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSQLExecAuditLogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSQLExecAuditLogResponseBody
	GetRequestId() *string
	SetSQLExecAuditLogList(v *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) *ListSQLExecAuditLogResponseBody
	GetSQLExecAuditLogList() *ListSQLExecAuditLogResponseBodySQLExecAuditLogList
	SetSuccess(v bool) *ListSQLExecAuditLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSQLExecAuditLogResponseBody
	GetTotalCount() *int64
}

type ListSQLExecAuditLogResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingStartTime
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// StartTime is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 39BC9C86-95AE-58F2-9862-A7C3D896****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entries returned.
	SQLExecAuditLogList *ListSQLExecAuditLogResponseBodySQLExecAuditLogList `json:"SQLExecAuditLogList,omitempty" xml:"SQLExecAuditLogList,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSQLExecAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSQLExecAuditLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSQLExecAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSQLExecAuditLogResponseBody) GetSQLExecAuditLogList() *ListSQLExecAuditLogResponseBodySQLExecAuditLogList {
	return s.SQLExecAuditLogList
}

func (s *ListSQLExecAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSQLExecAuditLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSQLExecAuditLogResponseBody) SetErrorCode(v string) *ListSQLExecAuditLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetErrorMessage(v string) *ListSQLExecAuditLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetRequestId(v string) *ListSQLExecAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetSQLExecAuditLogList(v *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) *ListSQLExecAuditLogResponseBody {
	s.SQLExecAuditLogList = v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetSuccess(v bool) *ListSQLExecAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) SetTotalCount(v int64) *ListSQLExecAuditLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBody) Validate() error {
	if s.SQLExecAuditLogList != nil {
		if err := s.SQLExecAuditLogList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSQLExecAuditLogResponseBodySQLExecAuditLogList struct {
	SQLExecAuditLog []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog `json:"SQLExecAuditLog,omitempty" xml:"SQLExecAuditLog,omitempty" type:"Repeated"`
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogList) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogList) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) GetSQLExecAuditLog() []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	return s.SQLExecAuditLog
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) SetSQLExecAuditLog(v []*ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) *ListSQLExecAuditLogResponseBodySQLExecAuditLogList {
	s.SQLExecAuditLog = v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogList) Validate() error {
	if s.SQLExecAuditLog != nil {
		for _, item := range s.SQLExecAuditLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog struct {
	// The number of rows affected by the SQL statement. For example, if you execute an SQL statement to query data, the number of retrieved rows is returned.
	//
	// example:
	//
	// 2
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 2157****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The amount of time consumed by the execution of the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 18
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
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
	// The ID of the instance.
	//
	// example:
	//
	// 185***
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the database.
	//
	// >  If the SQL statement takes effect on an instance, the name of the instance is returned.
	//
	// example:
	//
	// polar123@pc-bp1h9tgq4st9g****.mysql.polardb.rds.aliyuncs.com:3306[polar_qw_test]
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The time when the operation specified by the SQL statement was performed on the instance or database.
	//
	// example:
	//
	// 2021-11-08 11:04:27
	OpTime *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	// The comment on the SQL statement.
	//
	// example:
	//
	// success
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The SQL statement that was written.
	//
	// example:
	//
	// SELECT 	- FROM `polar123`.`p_qw` ORDER BY `id` DESC
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
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
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// polar123
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The ID of the user who wrote the SQL statement.
	//
	// example:
	//
	// 12****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user who wrote the SQL statement.
	//
	// example:
	//
	// test_UserName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) String() string {
	return dara.Prettify(s)
}

func (s ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GoString() string {
	return s.String()
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetDbId() *int64 {
	return s.DbId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetExecState() *string {
	return s.ExecState
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetLogic() *bool {
	return s.Logic
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetOpTime() *string {
	return s.OpTime
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetRemark() *string {
	return s.Remark
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSQL() *string {
	return s.SQL
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSQLType() *string {
	return s.SQLType
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetUserId() *int64 {
	return s.UserId
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) GetUserName() *string {
	return s.UserName
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetAffectRows(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.AffectRows = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetDbId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.DbId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetElapsedTime(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.ElapsedTime = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetExecState(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.ExecState = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetInstanceId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.InstanceId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetInstanceName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.InstanceName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetLogic(v bool) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.Logic = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetOpTime(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.OpTime = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetRemark(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.Remark = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSQL(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SQL = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSQLType(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SQLType = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetSchemaName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.SchemaName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetUserId(v int64) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.UserId = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) SetUserName(v string) *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog {
	s.UserName = &v
	return s
}

func (s *ListSQLExecAuditLogResponseBodySQLExecAuditLogListSQLExecAuditLog) Validate() error {
	return dara.Validate(s)
}
