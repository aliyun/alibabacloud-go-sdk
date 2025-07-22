// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasSQLLogHotDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDasSQLLogHotDataResponseBody
	GetCode() *string
	SetData(v *GetDasSQLLogHotDataResponseBodyData) *GetDasSQLLogHotDataResponseBody
	GetData() *GetDasSQLLogHotDataResponseBodyData
	SetMessage(v string) *GetDasSQLLogHotDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDasSQLLogHotDataResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDasSQLLogHotDataResponseBody
	GetSuccess() *string
}

type GetDasSQLLogHotDataResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetDasSQLLogHotDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDasSQLLogHotDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDasSQLLogHotDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDasSQLLogHotDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDasSQLLogHotDataResponseBody) GetData() *GetDasSQLLogHotDataResponseBodyData {
	return s.Data
}

func (s *GetDasSQLLogHotDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDasSQLLogHotDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDasSQLLogHotDataResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDasSQLLogHotDataResponseBody) SetCode(v string) *GetDasSQLLogHotDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBody) SetData(v *GetDasSQLLogHotDataResponseBodyData) *GetDasSQLLogHotDataResponseBody {
	s.Data = v
	return s
}

func (s *GetDasSQLLogHotDataResponseBody) SetMessage(v string) *GetDasSQLLogHotDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBody) SetRequestId(v string) *GetDasSQLLogHotDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBody) SetSuccess(v string) *GetDasSQLLogHotDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDasSQLLogHotDataResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The details of the data returned.
	List []*GetDasSQLLogHotDataResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetDasSQLLogHotDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDasSQLLogHotDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDasSQLLogHotDataResponseBodyData) GetExtra() interface{} {
	return s.Extra
}

func (s *GetDasSQLLogHotDataResponseBodyData) GetList() []*GetDasSQLLogHotDataResponseBodyDataList {
	return s.List
}

func (s *GetDasSQLLogHotDataResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetDasSQLLogHotDataResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDasSQLLogHotDataResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetDasSQLLogHotDataResponseBodyData) SetExtra(v interface{}) *GetDasSQLLogHotDataResponseBodyData {
	s.Extra = v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyData) SetList(v []*GetDasSQLLogHotDataResponseBodyDataList) *GetDasSQLLogHotDataResponseBodyData {
	s.List = v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyData) SetPageNo(v int64) *GetDasSQLLogHotDataResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyData) SetPageSize(v int64) *GetDasSQLLogHotDataResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyData) SetTotal(v int64) *GetDasSQLLogHotDataResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDasSQLLogHotDataResponseBodyDataList struct {
	// The account of the database.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The execution time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-23 T12:11:20Z
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The extended information. This parameter is a reserved parameter.
	//
	// example:
	//
	// None
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 47.100.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The execution duration. Unit: microseconds.
	//
	// example:
	//
	// 10000
	Latancy *int64 `json:"Latancy,omitempty" xml:"Latancy,omitempty"`
	// The lock wait duration. Unit: microseconds.
	//
	// example:
	//
	// 1
	LockTime *int64 `json:"LockTime,omitempty" xml:"LockTime,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 12
	LogicRead *int64  `json:"LogicRead,omitempty" xml:"LogicRead,omitempty"`
	NodeId    *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The execution time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-23 T12:11:20.999Z
	OriginTime *string `json:"OriginTime,omitempty" xml:"OriginTime,omitempty"`
	// The number of physical asynchronous reads.
	//
	// example:
	//
	// 0
	PhysicAsyncRead *int64 `json:"PhysicAsyncRead,omitempty" xml:"PhysicAsyncRead,omitempty"`
	// The number of physical synchronous reads.
	//
	// example:
	//
	// 0
	PhysicSyncRead *int64 `json:"PhysicSyncRead,omitempty" xml:"PhysicSyncRead,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 1
	ReturnRows *int64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	// The content of the SQL statement.
	//
	// example:
	//
	// select 1
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The number of rows scanned by the SQL statement.
	//
	// example:
	//
	// 29
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// 	- **SELECT**
	//
	// 	- **UPDATE**
	//
	// 	- **DELETE**
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The execution result. If a **0*	- is returned, the SQL statement was successfully executed. If an error code is returned, the SQL statement failed to be executed.
	//
	// example:
	//
	// 0
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// 657
	ThreadID *int64 `json:"ThreadID,omitempty" xml:"ThreadID,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// 0
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
	// The number of updated rows.
	//
	// example:
	//
	// 30
	UpdateRows *int64 `json:"UpdateRows,omitempty" xml:"UpdateRows,omitempty"`
}

func (s GetDasSQLLogHotDataResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetDasSQLLogHotDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetAccountName() *string {
	return s.AccountName
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetDBName() *string {
	return s.DBName
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetExecuteTime() *string {
	return s.ExecuteTime
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetExt() *string {
	return s.Ext
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetHostAddress() *string {
	return s.HostAddress
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetLatancy() *int64 {
	return s.Latancy
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetLockTime() *int64 {
	return s.LockTime
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetLogicRead() *int64 {
	return s.LogicRead
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetOriginTime() *string {
	return s.OriginTime
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetPhysicAsyncRead() *int64 {
	return s.PhysicAsyncRead
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetPhysicSyncRead() *int64 {
	return s.PhysicSyncRead
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetReturnRows() *int64 {
	return s.ReturnRows
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetSQLText() *string {
	return s.SQLText
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetSqlType() *string {
	return s.SqlType
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetThreadID() *int64 {
	return s.ThreadID
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetTransactionId() *string {
	return s.TransactionId
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) GetUpdateRows() *int64 {
	return s.UpdateRows
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetAccountName(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.AccountName = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetDBName(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.DBName = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetExecuteTime(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.ExecuteTime = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetExt(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.Ext = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetHostAddress(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.HostAddress = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetLatancy(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.Latancy = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetLockTime(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.LockTime = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetLogicRead(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.LogicRead = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetNodeId(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.NodeId = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetOriginTime(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.OriginTime = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetPhysicAsyncRead(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.PhysicAsyncRead = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetPhysicSyncRead(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.PhysicSyncRead = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetReturnRows(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.ReturnRows = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetSQLText(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.SQLText = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetScanRows(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.ScanRows = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetSqlType(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.SqlType = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetState(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.State = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetThreadID(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.ThreadID = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetTransactionId(v string) *GetDasSQLLogHotDataResponseBodyDataList {
	s.TransactionId = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) SetUpdateRows(v int64) *GetDasSQLLogHotDataResponseBodyDataList {
	s.UpdateRows = &v
	return s
}

func (s *GetDasSQLLogHotDataResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
