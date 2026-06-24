// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConsoleOperationLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSqlConsoleOperationLogResponseBodyData) *GetSqlConsoleOperationLogResponseBody
	GetData() []*GetSqlConsoleOperationLogResponseBodyData
	SetErrorCode(v string) *GetSqlConsoleOperationLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSqlConsoleOperationLogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetSqlConsoleOperationLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlConsoleOperationLogResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetSqlConsoleOperationLogResponseBody
	GetTotalCount() *int64
}

type GetSqlConsoleOperationLogResponseBody struct {
	// The response struct.
	Data []*GetSqlConsoleOperationLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// C0A813EB-4623-523A-8598-86390CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of logs.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSqlConsoleOperationLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConsoleOperationLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConsoleOperationLogResponseBody) GetData() []*GetSqlConsoleOperationLogResponseBodyData {
	return s.Data
}

func (s *GetSqlConsoleOperationLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSqlConsoleOperationLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSqlConsoleOperationLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConsoleOperationLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlConsoleOperationLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSqlConsoleOperationLogResponseBody) SetData(v []*GetSqlConsoleOperationLogResponseBodyData) *GetSqlConsoleOperationLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) SetErrorCode(v string) *GetSqlConsoleOperationLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) SetErrorMessage(v string) *GetSqlConsoleOperationLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) SetRequestId(v string) *GetSqlConsoleOperationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) SetSuccess(v bool) *GetSqlConsoleOperationLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) SetTotalCount(v int64) *GetSqlConsoleOperationLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSqlConsoleOperationLogResponseBodyData struct {
	// The number of affected rows.
	//
	// example:
	//
	// 100
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The execution duration. Unit: milliseconds.
	//
	// example:
	//
	// 10
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The database search name.
	//
	// example:
	//
	// mysql@xxx.com
	DatabaseSearchName *string `json:"DatabaseSearchName,omitempty" xml:"DatabaseSearchName,omitempty"`
	// The error message.
	//
	// example:
	//
	// Access Denied
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The database schema.
	//
	// example:
	//
	// mysql
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM user;
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The SQL type.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the logs.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the statement is executed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The username.
	//
	// example:
	//
	// user
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetSqlConsoleOperationLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConsoleOperationLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetCost() *int64 {
	return s.Cost
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetSqlType() *string {
	return s.SqlType
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlConsoleOperationLogResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetAffectRows(v int64) *GetSqlConsoleOperationLogResponseBodyData {
	s.AffectRows = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetCost(v int64) *GetSqlConsoleOperationLogResponseBodyData {
	s.Cost = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetDatabaseSearchName(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.DatabaseSearchName = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetError(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetSchema(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetSql(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetSqlType(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.SqlType = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetStartTime(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetSuccess(v bool) *GetSqlConsoleOperationLogResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) SetUsername(v string) *GetSqlConsoleOperationLogResponseBodyData {
	s.Username = &v
	return s
}

func (s *GetSqlConsoleOperationLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
