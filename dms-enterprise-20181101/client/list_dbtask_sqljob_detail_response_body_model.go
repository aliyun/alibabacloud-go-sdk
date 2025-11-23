// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBTaskSQLJobDetailList(v []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) *ListDBTaskSQLJobDetailResponseBody
	GetDBTaskSQLJobDetailList() []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList
	SetErrorCode(v string) *ListDBTaskSQLJobDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDBTaskSQLJobDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDBTaskSQLJobDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDBTaskSQLJobDetailResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDBTaskSQLJobDetailResponseBody
	GetTotalCount() *int64
}

type ListDBTaskSQLJobDetailResponseBody struct {
	// The details of SQL tasks.
	DBTaskSQLJobDetailList []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList `json:"DBTaskSQLJobDetailList,omitempty" xml:"DBTaskSQLJobDetailList,omitempty" type:"Repeated"`
	// The error code that is returned.
	//
	// example:
	//
	// MissingJobId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// JobId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F044E33-FE09-58F1-8C61-A0F612EC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of SQL tasks.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDBTaskSQLJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetDBTaskSQLJobDetailList() []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	return s.DBTaskSQLJobDetailList
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDBTaskSQLJobDetailResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetDBTaskSQLJobDetailList(v []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) *ListDBTaskSQLJobDetailResponseBody {
	s.DBTaskSQLJobDetailList = v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetErrorCode(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetErrorMessage(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetRequestId(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetSuccess(v bool) *ListDBTaskSQLJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetTotalCount(v int64) *ListDBTaskSQLJobDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) Validate() error {
	if s.DBTaskSQLJobDetailList != nil {
		for _, item := range s.DBTaskSQLJobDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList struct {
	// The number of rows affected by the SQL task.
	//
	// example:
	//
	// 0
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The SQL statement that was executed in the SQL task.
	//
	// example:
	//
	// update a set id = 1 where id  = 1;
	CurrentSql *string `json:"CurrentSql,omitempty" xml:"CurrentSql,omitempty"`
	// The ID of the physical database.
	//
	// example:
	//
	// 1988****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The point in time when the SQL task ended.
	//
	// example:
	//
	// 2021-12-16 00:00:01
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of times that the SQL statement was executed.
	//
	// example:
	//
	// 1
	ExecuteCount *int64 `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	// The ID of the details of the SQL task.
	//
	// example:
	//
	// 24723****
	JobDetailId *int64 `json:"JobDetailId,omitempty" xml:"JobDetailId,omitempty"`
	// The ID of the SQL task.
	//
	// example:
	//
	// 1276****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The details of the operational log.
	//
	// example:
	//
	// log_info
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
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
	// Indicates whether the SQL statement was skipped. Valid values:
	//
	// 	- **true**: The SQL statement was skipped.
	//
	// 	- **false**: The SQL statement was not skipped.
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The type of the SQL statement, such as DELETE, UPDATE, or ALTER_TABLE.
	//
	// example:
	//
	// CREATE_TABLE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The point in time when the SQL task started.
	//
	// example:
	//
	// 2021-12-16 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the SQL task. Valid values:
	//
	// 	- **INIT**: The SQL task was initialized.
	//
	// 	- **PENDING**: The SQL task waited to be run.
	//
	// 	- **BE_SCHEDULED**: The SQL task waited to be scheduled.
	//
	// 	- **FAIL**: The SQL task failed.
	//
	// 	- **SUCCESS**: The SQL task was successful.
	//
	// 	- **PAUSE**: The SQL task was paused.
	//
	// 	- **DELETE**: The SQL task was deleted.
	//
	// 	- **RUNNING**: The SQL task was being run.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The duration of the SQL task. Unit: milliseconds.
	//
	// example:
	//
	// 38
	TimeDelay *int64 `json:"TimeDelay,omitempty" xml:"TimeDelay,omitempty"`
}

func (s ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetCurrentSql() *string {
	return s.CurrentSql
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetExecuteCount() *int64 {
	return s.ExecuteCount
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetJobDetailId() *int64 {
	return s.JobDetailId
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetLog() *string {
	return s.Log
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetLogic() *bool {
	return s.Logic
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetSkip() *bool {
	return s.Skip
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetSqlType() *string {
	return s.SqlType
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetStatus() *string {
	return s.Status
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GetTimeDelay() *int64 {
	return s.TimeDelay
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetAffectRows(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.AffectRows = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetCurrentSql(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.CurrentSql = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetDbId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.DbId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetEndTime(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.EndTime = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetExecuteCount(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.ExecuteCount = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetJobDetailId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.JobDetailId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetJobId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetLog(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Log = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetLogic(v bool) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Logic = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetSkip(v bool) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Skip = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetSqlType(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.SqlType = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetStartTime(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.StartTime = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetStatus(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Status = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetTimeDelay(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.TimeDelay = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) Validate() error {
	return dara.Validate(s)
}
