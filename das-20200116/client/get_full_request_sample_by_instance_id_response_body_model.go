// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestSampleByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetFullRequestSampleByInstanceIdResponseBody
	GetCode() *int64
	SetData(v []*GetFullRequestSampleByInstanceIdResponseBodyData) *GetFullRequestSampleByInstanceIdResponseBody
	GetData() []*GetFullRequestSampleByInstanceIdResponseBodyData
	SetMessage(v string) *GetFullRequestSampleByInstanceIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFullRequestSampleByInstanceIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFullRequestSampleByInstanceIdResponseBody
	GetSuccess() *bool
}

type GetFullRequestSampleByInstanceIdResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetFullRequestSampleByInstanceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A74B755-98B7-59DB-8724-1321B394****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFullRequestSampleByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestSampleByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) GetData() []*GetFullRequestSampleByInstanceIdResponseBodyData {
	return s.Data
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) SetCode(v int64) *GetFullRequestSampleByInstanceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) SetData(v []*GetFullRequestSampleByInstanceIdResponseBodyData) *GetFullRequestSampleByInstanceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) SetMessage(v string) *GetFullRequestSampleByInstanceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) SetRequestId(v string) *GetFullRequestSampleByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) SetSuccess(v bool) *GetFullRequestSampleByInstanceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBody) Validate() error {
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

type GetFullRequestSampleByInstanceIdResponseBodyData struct {
	// The name of the database.
	//
	// example:
	//
	// dbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of rows fetched by PolarDB-X 2.0 compute nodes.
	//
	// example:
	//
	// 0
	Frows *int64 `json:"Frows,omitempty" xml:"Frows,omitempty"`
	// The lock wait duration. Unit: seconds.
	//
	// example:
	//
	// 0.0137
	LockWaitTime *float64 `json:"LockWaitTime,omitempty" xml:"LockWaitTime,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 165848
	LogicalRead *float64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 172.17.XX.XX
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The number of physical asynchronous reads.
	//
	// example:
	//
	// 0
	PhysicalAsyncRead *float64 `json:"PhysicalAsyncRead,omitempty" xml:"PhysicalAsyncRead,omitempty"`
	// The number of physical synchronous reads.
	//
	// example:
	//
	// 0
	PhysicalSyncRead *float64 `json:"PhysicalSyncRead,omitempty" xml:"PhysicalSyncRead,omitempty"`
	// The number of rows updated or returned on PolarDB-X 2.0 compute nodes.
	//
	// example:
	//
	// 0
	Rows *int64 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The total number of scanned rows.
	//
	// > This parameter is returned only for ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and PolarDB for MySQL databases.
	//
	// example:
	//
	// 2048576
	RowsExamined *int64 `json:"RowsExamined,omitempty" xml:"RowsExamined,omitempty"`
	// The number of rows returned by the SQL statement.
	//
	// example:
	//
	// 14
	RowsReturned *int64 `json:"RowsReturned,omitempty" xml:"RowsReturned,omitempty"`
	// The amount of time consumed to execute the SQL statement. Unit: seconds.
	//
	// example:
	//
	// 0.409789
	Rt *float64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// The number of scanned rows.
	//
	// example:
	//
	// 0
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The number of requests sent from PolarDB-X 2.0 compute nodes to data nodes.
	//
	// example:
	//
	// 0
	Scnt *int64 `json:"Scnt,omitempty" xml:"Scnt,omitempty"`
	// The sample SQL statement.
	//
	// example:
	//
	// select 	- from testdb01 where ****
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The SQL statement ID.
	//
	// example:
	//
	// 651b56fe9418d48edb8fdf0980ec****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The type of the SQL statement. Valid values: **SELECT**, **INSERT**, **UPDATE**, **DELETE**, **LOGIN**, **LOGOUT**, **MERGE**, **ALTER**, **CREATEINDEX**, **DROPINDEX**, **CREATE**, **DROP**, **SET**, **DESC**, **REPLACE**, **CALL**, **BEGIN**, **DESCRIBE**, **ROLLBACK**, **FLUSH**, **USE**, **SHOW**, **START**, **COMMIT**, and **RENAME**.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The time when the SQL statement was executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660100753556
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The number of updated rows.
	//
	// example:
	//
	// 0
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
	// The name of the user who executes the SQL statement.
	//
	// example:
	//
	// testuser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetFullRequestSampleByInstanceIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestSampleByInstanceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetDatabase() *string {
	return s.Database
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetFrows() *int64 {
	return s.Frows
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetLockWaitTime() *float64 {
	return s.LockWaitTime
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetLogicalRead() *float64 {
	return s.LogicalRead
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetOriginHost() *string {
	return s.OriginHost
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetPhysicalAsyncRead() *float64 {
	return s.PhysicalAsyncRead
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetPhysicalSyncRead() *float64 {
	return s.PhysicalSyncRead
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetRows() *int64 {
	return s.Rows
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetRowsExamined() *int64 {
	return s.RowsExamined
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetRowsReturned() *int64 {
	return s.RowsReturned
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetRt() *float64 {
	return s.Rt
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetScnt() *int64 {
	return s.Scnt
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetSqlId() *string {
	return s.SqlId
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetSqlType() *string {
	return s.SqlType
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetUpdateRows() *int64 {
	return s.UpdateRows
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetDatabase(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Database = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetFrows(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Frows = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetLockWaitTime(v float64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.LockWaitTime = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetLogicalRead(v float64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.LogicalRead = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetOriginHost(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.OriginHost = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetPhysicalAsyncRead(v float64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.PhysicalAsyncRead = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetPhysicalSyncRead(v float64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.PhysicalSyncRead = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetRows(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetRowsExamined(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.RowsExamined = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetRowsReturned(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.RowsReturned = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetRt(v float64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Rt = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetScanRows(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.ScanRows = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetScnt(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Scnt = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetSql(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetSqlId(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetSqlType(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.SqlType = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetTimestamp(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetUpdateRows(v int64) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.UpdateRows = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) SetUser(v string) *GetFullRequestSampleByInstanceIdResponseBodyData {
	s.User = &v
	return s
}

func (s *GetFullRequestSampleByInstanceIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
