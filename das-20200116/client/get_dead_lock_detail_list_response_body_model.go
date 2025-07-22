// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeadLockDetailListResponseBody
	GetCode() *string
	SetData(v *GetDeadLockDetailListResponseBodyData) *GetDeadLockDetailListResponseBody
	GetData() *GetDeadLockDetailListResponseBodyData
	SetMessage(v string) *GetDeadLockDetailListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeadLockDetailListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDeadLockDetailListResponseBody
	GetSuccess() *string
}

type GetDeadLockDetailListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetDeadLockDetailListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request is successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 840F51F7-9C01-538D-94F6-AE712905****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s GetDeadLockDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeadLockDetailListResponseBody) GetData() *GetDeadLockDetailListResponseBodyData {
	return s.Data
}

func (s *GetDeadLockDetailListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeadLockDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeadLockDetailListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDeadLockDetailListResponseBody) SetCode(v string) *GetDeadLockDetailListResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeadLockDetailListResponseBody) SetData(v *GetDeadLockDetailListResponseBodyData) *GetDeadLockDetailListResponseBody {
	s.Data = v
	return s
}

func (s *GetDeadLockDetailListResponseBody) SetMessage(v string) *GetDeadLockDetailListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeadLockDetailListResponseBody) SetRequestId(v string) *GetDeadLockDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeadLockDetailListResponseBody) SetSuccess(v string) *GetDeadLockDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeadLockDetailListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeadLockDetailListResponseBodyData struct {
	// The details of the data returned.
	List []*GetDeadLockDetailListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
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
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetDeadLockDetailListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListResponseBodyData) GetList() []*GetDeadLockDetailListResponseBodyDataList {
	return s.List
}

func (s *GetDeadLockDetailListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetDeadLockDetailListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDeadLockDetailListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetDeadLockDetailListResponseBodyData) SetList(v []*GetDeadLockDetailListResponseBodyDataList) *GetDeadLockDetailListResponseBodyData {
	s.List = v
	return s
}

func (s *GetDeadLockDetailListResponseBodyData) SetPageNo(v int64) *GetDeadLockDetailListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyData) SetPageSize(v int64) *GetDeadLockDetailListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyData) SetTotal(v int64) *GetDeadLockDetailListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeadLockDetailListResponseBodyDataList struct {
	// The time when the data was collected. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1702301170701
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The blocking details of the instance. The details are information about the session that caused the lock.
	BlockProcessList []*GetDeadLockDetailListResponseBodyDataListBlockProcessList `json:"BlockProcessList,omitempty" xml:"BlockProcessList,omitempty" type:"Repeated"`
	// The name of the client.
	//
	// example:
	//
	// Microsoft SQL Server Management Studio - Query
	ClientApp *string `json:"ClientApp,omitempty" xml:"ClientApp,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// school
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The hostname.
	//
	// example:
	//
	// sd74020124
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The time when the transaction was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1702301141000
	LastTranStarted *int64 `json:"LastTranStarted,omitempty" xml:"LastTranStarted,omitempty"`
	// The mode of the lock. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// U
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The size of the logs generated in the session. Unit: bytes.
	//
	// example:
	//
	// 352
	LogUsed *int64 `json:"LogUsed,omitempty" xml:"LogUsed,omitempty"`
	// The logon name of the user.
	//
	// example:
	//
	// sd74020124\\\\Administrator
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The locked object.
	//
	// example:
	//
	// school.dbo.test1
	ObjectOwned *string `json:"ObjectOwned,omitempty" xml:"ObjectOwned,omitempty"`
	// The object that the transaction requested to lock.
	//
	// example:
	//
	// school.dbo.test2
	ObjectRequested *string `json:"ObjectRequested,omitempty" xml:"ObjectRequested,omitempty"`
	// The lock mode held by the session. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// X
	OwnMode *string `json:"OwnMode,omitempty" xml:"OwnMode,omitempty"`
	// The ID of the session in which the transaction is started.
	//
	// example:
	//
	// 56
	Spid *int64 `json:"Spid,omitempty" xml:"Spid,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// update test2 set col1 =88
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The status of the transaction.
	//
	// example:
	//
	// suspended
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the session is the victim of the deadlock. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 1
	Victim *int64 `json:"Victim,omitempty" xml:"Victim,omitempty"`
	// The lock mode requested by the session. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// U
	WaitMode *string `json:"WaitMode,omitempty" xml:"WaitMode,omitempty"`
	// The resources requested by the transaction.
	//
	// example:
	//
	// RID: 5:1:376:0
	WaitResource *string `json:"WaitResource,omitempty" xml:"WaitResource,omitempty"`
	// The details of the resources requested by the transaction.
	//
	// example:
	//
	// RID:school:school.mdf:376:0
	WaitResourceDescription *string `json:"WaitResourceDescription,omitempty" xml:"WaitResourceDescription,omitempty"`
}

func (s GetDeadLockDetailListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetBlockProcessList() []*GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	return s.BlockProcessList
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetClientApp() *string {
	return s.ClientApp
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetHostName() *string {
	return s.HostName
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetLastTranStarted() *int64 {
	return s.LastTranStarted
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetLockMode() *string {
	return s.LockMode
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetLogUsed() *int64 {
	return s.LogUsed
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetLoginName() *string {
	return s.LoginName
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetObjectOwned() *string {
	return s.ObjectOwned
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetObjectRequested() *string {
	return s.ObjectRequested
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetOwnMode() *string {
	return s.OwnMode
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetSpid() *int64 {
	return s.Spid
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetVictim() *int64 {
	return s.Victim
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetWaitMode() *string {
	return s.WaitMode
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetWaitResource() *string {
	return s.WaitResource
}

func (s *GetDeadLockDetailListResponseBodyDataList) GetWaitResourceDescription() *string {
	return s.WaitResourceDescription
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetBatchId(v int64) *GetDeadLockDetailListResponseBodyDataList {
	s.BatchId = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetBlockProcessList(v []*GetDeadLockDetailListResponseBodyDataListBlockProcessList) *GetDeadLockDetailListResponseBodyDataList {
	s.BlockProcessList = v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetClientApp(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.ClientApp = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetDatabaseName(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.DatabaseName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetHostName(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.HostName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetLastTranStarted(v int64) *GetDeadLockDetailListResponseBodyDataList {
	s.LastTranStarted = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetLockMode(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.LockMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetLogUsed(v int64) *GetDeadLockDetailListResponseBodyDataList {
	s.LogUsed = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetLoginName(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetObjectOwned(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.ObjectOwned = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetObjectRequested(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.ObjectRequested = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetOwnMode(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.OwnMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetSpid(v int64) *GetDeadLockDetailListResponseBodyDataList {
	s.Spid = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetSqlText(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.SqlText = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetStatus(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetVictim(v int64) *GetDeadLockDetailListResponseBodyDataList {
	s.Victim = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetWaitMode(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.WaitMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetWaitResource(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.WaitResource = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) SetWaitResourceDescription(v string) *GetDeadLockDetailListResponseBodyDataList {
	s.WaitResourceDescription = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type GetDeadLockDetailListResponseBodyDataListBlockProcessList struct {
	// The name of the client that initiates the transaction in the session.
	//
	// example:
	//
	// Microsoft SQL Server Management Studio - Query
	ClientApp *string `json:"ClientApp,omitempty" xml:"ClientApp,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// school
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The hostname.
	//
	// example:
	//
	// sd74020124
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The time when the transaction was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1702301152000
	LastTranStarted *int64 `json:"LastTranStarted,omitempty" xml:"LastTranStarted,omitempty"`
	// The mode of the lock. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// U
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The size of the logs generated in the session. Unit: bytes.
	//
	// example:
	//
	// 352
	LogUsed *int64 `json:"LogUsed,omitempty" xml:"LogUsed,omitempty"`
	// The logon name of the user.
	//
	// example:
	//
	// sd74020124\\\\Administrator
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The locked object.
	//
	// example:
	//
	// school.dbo.test2
	ObjectOwned *string `json:"ObjectOwned,omitempty" xml:"ObjectOwned,omitempty"`
	// The object that the transaction requested to lock.
	//
	// example:
	//
	// school.dbo.test1
	ObjectRequested *string `json:"ObjectRequested,omitempty" xml:"ObjectRequested,omitempty"`
	// The lock mode held by the session. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// X
	OwnMode *string `json:"OwnMode,omitempty" xml:"OwnMode,omitempty"`
	// The ID of the session in which the transaction is started.
	//
	// example:
	//
	// 61
	Spid *int64 `json:"Spid,omitempty" xml:"Spid,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// update test1 set col1 =9
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The status of the transaction.
	//
	// example:
	//
	// suspended
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the session is the victim of the deadlock. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 0
	Victim *int64 `json:"Victim,omitempty" xml:"Victim,omitempty"`
	// The lock mode requested by the session. For more information, see [Lock modes](https://help.aliyun.com/document_detail/2362804.html).
	//
	// example:
	//
	// U
	WaitMode *string `json:"WaitMode,omitempty" xml:"WaitMode,omitempty"`
	// The resources requested by the transaction.
	//
	// example:
	//
	// RID: 5:1:312:0
	WaitResource *string `json:"WaitResource,omitempty" xml:"WaitResource,omitempty"`
	// The details of the resources requested by the transaction.
	//
	// example:
	//
	// RID:school:school.mdf:312:0
	WaitResourceDescription *string `json:"WaitResourceDescription,omitempty" xml:"WaitResourceDescription,omitempty"`
}

func (s GetDeadLockDetailListResponseBodyDataListBlockProcessList) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListResponseBodyDataListBlockProcessList) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetClientApp() *string {
	return s.ClientApp
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetHostName() *string {
	return s.HostName
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetLastTranStarted() *int64 {
	return s.LastTranStarted
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetLockMode() *string {
	return s.LockMode
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetLogUsed() *int64 {
	return s.LogUsed
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetLoginName() *string {
	return s.LoginName
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetObjectOwned() *string {
	return s.ObjectOwned
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetObjectRequested() *string {
	return s.ObjectRequested
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetOwnMode() *string {
	return s.OwnMode
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetSpid() *int64 {
	return s.Spid
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetStatus() *string {
	return s.Status
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetVictim() *int64 {
	return s.Victim
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetWaitMode() *string {
	return s.WaitMode
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetWaitResource() *string {
	return s.WaitResource
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) GetWaitResourceDescription() *string {
	return s.WaitResourceDescription
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetClientApp(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.ClientApp = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetDatabaseName(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.DatabaseName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetHostName(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.HostName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetLastTranStarted(v int64) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.LastTranStarted = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetLockMode(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.LockMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetLogUsed(v int64) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.LogUsed = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetLoginName(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.LoginName = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetObjectOwned(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.ObjectOwned = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetObjectRequested(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.ObjectRequested = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetOwnMode(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.OwnMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetSpid(v int64) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.Spid = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetSqlText(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.SqlText = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetStatus(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.Status = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetVictim(v int64) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.Victim = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetWaitMode(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.WaitMode = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetWaitResource(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.WaitResource = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) SetWaitResourceDescription(v string) *GetDeadLockDetailListResponseBodyDataListBlockProcessList {
	s.WaitResourceDescription = &v
	return s
}

func (s *GetDeadLockDetailListResponseBodyDataListBlockProcessList) Validate() error {
	return dara.Validate(s)
}
