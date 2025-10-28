// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlockingDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBlockingDetailListResponseBody
	GetCode() *string
	SetData(v *GetBlockingDetailListResponseBodyData) *GetBlockingDetailListResponseBody
	GetData() *GetBlockingDetailListResponseBodyData
	SetMessage(v string) *GetBlockingDetailListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBlockingDetailListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetBlockingDetailListResponseBody
	GetSuccess() *string
}

type GetBlockingDetailListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetBlockingDetailListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
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

func (s GetBlockingDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBlockingDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBlockingDetailListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBlockingDetailListResponseBody) GetData() *GetBlockingDetailListResponseBodyData {
	return s.Data
}

func (s *GetBlockingDetailListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBlockingDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBlockingDetailListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetBlockingDetailListResponseBody) SetCode(v string) *GetBlockingDetailListResponseBody {
	s.Code = &v
	return s
}

func (s *GetBlockingDetailListResponseBody) SetData(v *GetBlockingDetailListResponseBodyData) *GetBlockingDetailListResponseBody {
	s.Data = v
	return s
}

func (s *GetBlockingDetailListResponseBody) SetMessage(v string) *GetBlockingDetailListResponseBody {
	s.Message = &v
	return s
}

func (s *GetBlockingDetailListResponseBody) SetRequestId(v string) *GetBlockingDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBlockingDetailListResponseBody) SetSuccess(v string) *GetBlockingDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBlockingDetailListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBlockingDetailListResponseBodyData struct {
	// The details of the data returned.
	List []*GetBlockingDetailListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 19
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetBlockingDetailListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBlockingDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBlockingDetailListResponseBodyData) GetList() []*GetBlockingDetailListResponseBodyDataList {
	return s.List
}

func (s *GetBlockingDetailListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetBlockingDetailListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetBlockingDetailListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetBlockingDetailListResponseBodyData) SetList(v []*GetBlockingDetailListResponseBodyDataList) *GetBlockingDetailListResponseBodyData {
	s.List = v
	return s
}

func (s *GetBlockingDetailListResponseBodyData) SetPageNo(v int64) *GetBlockingDetailListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyData) SetPageSize(v int64) *GetBlockingDetailListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyData) SetTotal(v int64) *GetBlockingDetailListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBlockingDetailListResponseBodyDataList struct {
	// The batch ID.
	//
	// example:
	//
	// 1683530096156
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The client name.
	//
	// example:
	//
	// .Net SqlClient Data Provider
	ClientAppName *string `json:"ClientAppName,omitempty" xml:"ClientAppName,omitempty"`
	// The time when the blocking data was collected. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1700065800000
	CurrentCollectionTime *int64 `json:"CurrentCollectionTime,omitempty" xml:"CurrentCollectionTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// school
	DataBase *string `json:"DataBase,omitempty" xml:"DataBase,omitempty"`
	// The client hostname.
	//
	// example:
	//
	// ALLBNMGTAPPRD01
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The username that is used for the logon.
	//
	// example:
	//
	// Cheney603
	LoginId *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
	// The hash value of the SQL statement.
	//
	// example:
	//
	// 6977DD06CD9CAFF2
	QueryHash *string `json:"QueryHash,omitempty" xml:"QueryHash,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 1717
	Spid *string `json:"Spid,omitempty" xml:"Spid,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select 	- from test1
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The time when the execution started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The duration of the blocking. Unit: milliseconds.
	//
	// example:
	//
	// 30000
	WaitTimeMs *int64 `json:"WaitTimeMs,omitempty" xml:"WaitTimeMs,omitempty"`
	// The wait type. For more information about wait types, see [sys.dm_os_wait_stats (Transact-SQL)](https://learn.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-os-wait-stats-transact-sql?view=sql-server-ver15).
	//
	// example:
	//
	// MISCELLANEOUS
	WaitType *string `json:"WaitType,omitempty" xml:"WaitType,omitempty"`
}

func (s GetBlockingDetailListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetBlockingDetailListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetBlockingDetailListResponseBodyDataList) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetBlockingDetailListResponseBodyDataList) GetClientAppName() *string {
	return s.ClientAppName
}

func (s *GetBlockingDetailListResponseBodyDataList) GetCurrentCollectionTime() *int64 {
	return s.CurrentCollectionTime
}

func (s *GetBlockingDetailListResponseBodyDataList) GetDataBase() *string {
	return s.DataBase
}

func (s *GetBlockingDetailListResponseBodyDataList) GetHostName() *string {
	return s.HostName
}

func (s *GetBlockingDetailListResponseBodyDataList) GetLoginId() *string {
	return s.LoginId
}

func (s *GetBlockingDetailListResponseBodyDataList) GetQueryHash() *string {
	return s.QueryHash
}

func (s *GetBlockingDetailListResponseBodyDataList) GetSpid() *string {
	return s.Spid
}

func (s *GetBlockingDetailListResponseBodyDataList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetBlockingDetailListResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetBlockingDetailListResponseBodyDataList) GetWaitTimeMs() *int64 {
	return s.WaitTimeMs
}

func (s *GetBlockingDetailListResponseBodyDataList) GetWaitType() *string {
	return s.WaitType
}

func (s *GetBlockingDetailListResponseBodyDataList) SetBatchId(v int64) *GetBlockingDetailListResponseBodyDataList {
	s.BatchId = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetClientAppName(v string) *GetBlockingDetailListResponseBodyDataList {
	s.ClientAppName = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetCurrentCollectionTime(v int64) *GetBlockingDetailListResponseBodyDataList {
	s.CurrentCollectionTime = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetDataBase(v string) *GetBlockingDetailListResponseBodyDataList {
	s.DataBase = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetHostName(v string) *GetBlockingDetailListResponseBodyDataList {
	s.HostName = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetLoginId(v string) *GetBlockingDetailListResponseBodyDataList {
	s.LoginId = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetQueryHash(v string) *GetBlockingDetailListResponseBodyDataList {
	s.QueryHash = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetSpid(v string) *GetBlockingDetailListResponseBodyDataList {
	s.Spid = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetSqlText(v string) *GetBlockingDetailListResponseBodyDataList {
	s.SqlText = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetStartTime(v string) *GetBlockingDetailListResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetWaitTimeMs(v int64) *GetBlockingDetailListResponseBodyDataList {
	s.WaitTimeMs = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) SetWaitType(v string) *GetBlockingDetailListResponseBodyDataList {
	s.WaitType = &v
	return s
}

func (s *GetBlockingDetailListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
