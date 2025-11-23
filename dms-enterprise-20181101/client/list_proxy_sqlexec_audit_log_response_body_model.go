// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxySQLExecAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProxySQLExecAuditLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListProxySQLExecAuditLogResponseBody
	GetErrorMessage() *string
	SetProxySQLExecAuditLogList(v *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) *ListProxySQLExecAuditLogResponseBody
	GetProxySQLExecAuditLogList() *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList
	SetRequestId(v string) *ListProxySQLExecAuditLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProxySQLExecAuditLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListProxySQLExecAuditLogResponseBody
	GetTotalCount() *int64
}

type ListProxySQLExecAuditLogResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidStartTime
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Specified parameter StartTime is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The audit information about the database instance that is provided by the secure access proxy feature.
	ProxySQLExecAuditLogList *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList `json:"ProxySQLExecAuditLogList,omitempty" xml:"ProxySQLExecAuditLogList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 50ECB006-2C35-5FCA-91B9-01987A0B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s ListProxySQLExecAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProxySQLExecAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListProxySQLExecAuditLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProxySQLExecAuditLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProxySQLExecAuditLogResponseBody) GetProxySQLExecAuditLogList() *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList {
	return s.ProxySQLExecAuditLogList
}

func (s *ListProxySQLExecAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProxySQLExecAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProxySQLExecAuditLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProxySQLExecAuditLogResponseBody) SetErrorCode(v string) *ListProxySQLExecAuditLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) SetErrorMessage(v string) *ListProxySQLExecAuditLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) SetProxySQLExecAuditLogList(v *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) *ListProxySQLExecAuditLogResponseBody {
	s.ProxySQLExecAuditLogList = v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) SetRequestId(v string) *ListProxySQLExecAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) SetSuccess(v bool) *ListProxySQLExecAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) SetTotalCount(v int64) *ListProxySQLExecAuditLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBody) Validate() error {
	if s.ProxySQLExecAuditLogList != nil {
		if err := s.ProxySQLExecAuditLogList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList struct {
	ProxySQLExecAuditLog []*ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog `json:"ProxySQLExecAuditLog,omitempty" xml:"ProxySQLExecAuditLog,omitempty" type:"Repeated"`
}

func (s ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) String() string {
	return dara.Prettify(s)
}

func (s ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) GoString() string {
	return s.String()
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) GetProxySQLExecAuditLog() []*ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	return s.ProxySQLExecAuditLog
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) SetProxySQLExecAuditLog(v []*ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList {
	s.ProxySQLExecAuditLog = v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogList) Validate() error {
	if s.ProxySQLExecAuditLog != nil {
		for _, item := range s.ProxySQLExecAuditLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog struct {
	// Indicates the total number of rows returned after the SQL statement was executed. If an SELECT SQL statement is executed, the return value of this parameter indicates the total number of the queried data rows.
	//
	// example:
	//
	// 1
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The amount of time that is consumed to execute the SQL statement. Unit: milliseconds.
	//
	// example:
	//
	// 1324
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
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
	// The ID of the database instance.
	//
	// example:
	//
	// 4***
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the database instance.
	//
	// example:
	//
	// pc-uf662nrg017c6****.mysql.polardb.rds.aliyuncs.com:3306【test】
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The time at which the user executes the SQL statement on the database instance. The value of this parameter must be a timestamp that follows the UNIX time format.
	//
	// example:
	//
	// 1636876446000
	OpTime *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
	// The description.
	//
	// example:
	//
	// success
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The SQL statement that was executed.
	//
	// example:
	//
	// select 1;
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The type of the SQL statement. Valid values:
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
	// The name of the database.
	//
	// example:
	//
	// test_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 4****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// testNickName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) String() string {
	return dara.Prettify(s)
}

func (s ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GoString() string {
	return s.String()
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetExecState() *string {
	return s.ExecState
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetOpTime() *string {
	return s.OpTime
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetRemark() *string {
	return s.Remark
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetSQL() *string {
	return s.SQL
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetSQLType() *string {
	return s.SQLType
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetUserId() *int64 {
	return s.UserId
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) GetUserName() *string {
	return s.UserName
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetAffectRows(v int64) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.AffectRows = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetElapsedTime(v int64) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.ElapsedTime = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetExecState(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.ExecState = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetInstanceId(v int64) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.InstanceId = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetInstanceName(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.InstanceName = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetOpTime(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.OpTime = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetRemark(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.Remark = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetSQL(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.SQL = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetSQLType(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.SQLType = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetSchemaName(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.SchemaName = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetUserId(v int64) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.UserId = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) SetUserName(v string) *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog {
	s.UserName = &v
	return s
}

func (s *ListProxySQLExecAuditLogResponseBodyProxySQLExecAuditLogListProxySQLExecAuditLog) Validate() error {
	return dara.Validate(s)
}
