// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSQLLogsV2ResponseBody
	GetAccessDeniedDetail() *string
	SetItems(v []*DescribeSQLLogsV2ResponseBodyItems) *DescribeSQLLogsV2ResponseBody
	GetItems() []*DescribeSQLLogsV2ResponseBodyItems
	SetPageNumber(v int32) *DescribeSQLLogsV2ResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSQLLogsV2ResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSQLLogsV2ResponseBody
	GetRequestId() *string
}

type DescribeSQLLogsV2ResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// account name invalid
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The queried SQL execution logs.
	Items []*DescribeSQLLogsV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A7941C94-B92F-46A0-BD3E-2D**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSQLLogsV2ResponseBody) GetItems() []*DescribeSQLLogsV2ResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogsV2ResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogsV2ResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSQLLogsV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogsV2ResponseBody) SetAccessDeniedDetail(v string) *DescribeSQLLogsV2ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetItems(v []*DescribeSQLLogsV2ResponseBodyItems) *DescribeSQLLogsV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageNumber(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogsV2ResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) SetRequestId(v string) *DescribeSQLLogsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogsV2ResponseBodyItems struct {
	// The database account that executes the SQL statement.
	//
	// example:
	//
	// testadmin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adbpgadmin
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The role of the database.
	//
	// example:
	//
	// master
	DBRole *string `json:"DBRole,omitempty" xml:"DBRole,omitempty"`
	// The error code.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The execution duration of the SQL statement.
	//
	// example:
	//
	// 2
	ExecuteCost *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The type of the query language.
	//
	// example:
	//
	// DQL
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The time when the SQL statement was executed.
	//
	// example:
	//
	// 2021-03-15T17:02:32Z
	OperationExecuteTime *string `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 2548026401648157601713924318883
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// select 1
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The number of entries scanned.
	//
	// example:
	//
	// 1
	ScanRowCounts *int64 `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty"`
	// The ID of the session.
	//
	// example:
	//
	// efc33bd7-f1dc-4b24-b4fb-ab0d5329b7bb
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 100.**.**.90
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The number of the source port.
	//
	// example:
	//
	// 50514
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeSQLLogsV2ResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetDBRole() *string {
	return s.DBRole
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetExecuteCost() *float32 {
	return s.ExecuteCost
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetOperationExecuteTime() *string {
	return s.OperationExecuteTime
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetReturnRowCounts() *int64 {
	return s.ReturnRowCounts
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetSQLText() *string {
	return s.SQLText
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetScanRowCounts() *int64 {
	return s.ScanRowCounts
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DescribeSQLLogsV2ResponseBodyItems) GetSourcePort() *int32 {
	return s.SourcePort
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetAccountName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBName(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetDBRole(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetErrorCode(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetErrorMsg(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteCost(v float32) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetExecuteState(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationClass(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationExecuteTime(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetOperationType(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetQueryId(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.QueryId = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetReturnRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSQLText(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetScanRowCounts(v int64) *DescribeSQLLogsV2ResponseBodyItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSessionId(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SessionId = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourceIP(v string) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) SetSourcePort(v int32) *DescribeSQLLogsV2ResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSQLLogsV2ResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
