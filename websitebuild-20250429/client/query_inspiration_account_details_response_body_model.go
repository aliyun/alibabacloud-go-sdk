// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationAccountDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInspirationAccountDetailsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryInspirationAccountDetailsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryInspirationAccountDetailsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryInspirationAccountDetailsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryInspirationAccountDetailsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryInspirationAccountDetailsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryInspirationAccountDetailsResponseBodyModule) *QueryInspirationAccountDetailsResponseBody
	GetModule() *QueryInspirationAccountDetailsResponseBodyModule
	SetRequestId(v string) *QueryInspirationAccountDetailsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryInspirationAccountDetailsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryInspirationAccountDetailsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryInspirationAccountDetailsResponseBody
	GetSynchro() *bool
}

type QueryInspirationAccountDetailsResponseBody struct {
	// Access denied details
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application Name. Query the application with this name.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *QueryInspirationAccountDetailsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether processing is synchronous
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryInspirationAccountDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInspirationAccountDetailsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryInspirationAccountDetailsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryInspirationAccountDetailsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryInspirationAccountDetailsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryInspirationAccountDetailsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryInspirationAccountDetailsResponseBody) GetModule() *QueryInspirationAccountDetailsResponseBodyModule {
	return s.Module
}

func (s *QueryInspirationAccountDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInspirationAccountDetailsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryInspirationAccountDetailsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryInspirationAccountDetailsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryInspirationAccountDetailsResponseBody) SetAccessDeniedDetail(v string) *QueryInspirationAccountDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetAllowRetry(v bool) *QueryInspirationAccountDetailsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetAppName(v string) *QueryInspirationAccountDetailsResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetDynamicCode(v string) *QueryInspirationAccountDetailsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetDynamicMessage(v string) *QueryInspirationAccountDetailsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetErrorArgs(v []interface{}) *QueryInspirationAccountDetailsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetModule(v *QueryInspirationAccountDetailsResponseBodyModule) *QueryInspirationAccountDetailsResponseBody {
	s.Module = v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetRequestId(v string) *QueryInspirationAccountDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetRootErrorCode(v string) *QueryInspirationAccountDetailsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetRootErrorMsg(v string) *QueryInspirationAccountDetailsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) SetSynchro(v bool) *QueryInspirationAccountDetailsResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationAccountDetailsResponseBodyModule struct {
	// Current page number.
	//
	// example:
	//
	// 12
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// Request result.
	Data []*QueryInspirationAccountDetailsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Decision weight
	Next *QueryInspirationAccountDetailsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether a previous page exists
	//
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// In addition to paging limits, the server-side processes at most the latest 1,000 records for the current query. If the result exceeds 1,000 records, **ResultLimit*	- is **true**. You can narrow the Time Range and search again. Otherwise, **ResultLimit*	- is **false**.
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryInspirationAccountDetailsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetData() []*QueryInspirationAccountDetailsResponseBodyModuleData {
	return s.Data
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetNext() *QueryInspirationAccountDetailsResponseBodyModuleNext {
	return s.Next
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetCurrentPageNum(v int32) *QueryInspirationAccountDetailsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetData(v []*QueryInspirationAccountDetailsResponseBodyModuleData) *QueryInspirationAccountDetailsResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetNext(v *QueryInspirationAccountDetailsResponseBodyModuleNext) *QueryInspirationAccountDetailsResponseBodyModule {
	s.Next = v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetNextPage(v bool) *QueryInspirationAccountDetailsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetPageSize(v int32) *QueryInspirationAccountDetailsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetPrePage(v bool) *QueryInspirationAccountDetailsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetResultLimit(v bool) *QueryInspirationAccountDetailsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetTotalItemNum(v int32) *QueryInspirationAccountDetailsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) SetTotalPageNum(v int32) *QueryInspirationAccountDetailsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationAccountDetailsResponseBodyModuleData struct {
	// Acquisition Time
	//
	// example:
	//
	// 2026-03-03 12:00:00
	AcquisitionTime *string `json:"AcquisitionTime,omitempty" xml:"AcquisitionTime,omitempty"`
	// Balance (initQuota - used)
	//
	// example:
	//
	// 12
	Balance    *int64  `json:"Balance,omitempty" xml:"Balance,omitempty"`
	BalanceStr *string `json:"BalanceStr,omitempty" xml:"BalanceStr,omitempty"`
	// Validity Period expiration time
	//
	// example:
	//
	// 2025-04-11 10:26:27 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether it has expired (endDate < current time). The frontend uses this to gray out the display.
	//
	// example:
	//
	// False
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Quantity obtained (initial quota)
	//
	// example:
	//
	// 123123
	InitQuota    *int64  `json:"InitQuota,omitempty" xml:"InitQuota,omitempty"`
	InitQuotaStr *string `json:"InitQuotaStr,omitempty" xml:"InitQuotaStr,omitempty"`
	// Source type code (such as FREE_TRIAL_GIFT, INSTANCE_GIFT, UPGRADE_GRANT, PURCHASED)
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Source Type Display Name
	//
	// example:
	//
	// FREE_TRIAL_GIFT
	SourceTypeName *string `json:"SourceTypeName,omitempty" xml:"SourceTypeName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryInspirationAccountDetailsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetAcquisitionTime() *string {
	return s.AcquisitionTime
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetBalance() *int64 {
	return s.Balance
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetBalanceStr() *string {
	return s.BalanceStr
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetExpired() *bool {
	return s.Expired
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetInitQuota() *int64 {
	return s.InitQuota
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetInitQuotaStr() *string {
	return s.InitQuotaStr
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetSourceTypeName() *string {
	return s.SourceTypeName
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetAcquisitionTime(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.AcquisitionTime = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetBalance(v int64) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.Balance = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetBalanceStr(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.BalanceStr = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetEndDate(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.EndDate = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetExpired(v bool) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.Expired = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetInitQuota(v int64) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.InitQuota = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetInitQuotaStr(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.InitQuotaStr = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetSourceType(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.SourceType = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetSourceTypeName(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.SourceTypeName = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) SetStatus(v string) *QueryInspirationAccountDetailsResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type QueryInspirationAccountDetailsResponseBodyModuleNext struct {
	// Acquisition time
	//
	// example:
	//
	// 2026-03-03 12:00:00
	AcquisitionTime *string `json:"AcquisitionTime,omitempty" xml:"AcquisitionTime,omitempty"`
	// Balance (initQuota - used)
	//
	// example:
	//
	// 12
	Balance    *int64  `json:"Balance,omitempty" xml:"Balance,omitempty"`
	BalanceStr *string `json:"BalanceStr,omitempty" xml:"BalanceStr,omitempty"`
	// Expiration time of the validity period
	//
	// example:
	//
	// 2026-02-25 10:11:25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether it has expired (endDate < current time). The frontend uses this to gray out the display.
	//
	// example:
	//
	// False
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Quantity obtained (initial quota)
	//
	// example:
	//
	// 123123
	InitQuota    *int64  `json:"InitQuota,omitempty" xml:"InitQuota,omitempty"`
	InitQuotaStr *string `json:"InitQuotaStr,omitempty" xml:"InitQuotaStr,omitempty"`
	// Source type code (such as FREE_TRIAL_GIFT, INSTANCE_GIFT, UPGRADE_GRANT, PURCHASED)
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Display Name of the source type
	//
	// example:
	//
	// FREE_TRIAL_GIFT
	SourceTypeName *string `json:"SourceTypeName,omitempty" xml:"SourceTypeName,omitempty"`
}

func (s QueryInspirationAccountDetailsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationAccountDetailsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetAcquisitionTime() *string {
	return s.AcquisitionTime
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetBalance() *int64 {
	return s.Balance
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetBalanceStr() *string {
	return s.BalanceStr
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetExpired() *bool {
	return s.Expired
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetInitQuota() *int64 {
	return s.InitQuota
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetInitQuotaStr() *string {
	return s.InitQuotaStr
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) GetSourceTypeName() *string {
	return s.SourceTypeName
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetAcquisitionTime(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.AcquisitionTime = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetBalance(v int64) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.Balance = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetBalanceStr(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.BalanceStr = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetEndDate(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.EndDate = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetExpired(v bool) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.Expired = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetInitQuota(v int64) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.InitQuota = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetInitQuotaStr(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.InitQuotaStr = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetSourceType(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.SourceType = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) SetSourceTypeName(v string) *QueryInspirationAccountDetailsResponseBodyModuleNext {
	s.SourceTypeName = &v
	return s
}

func (s *QueryInspirationAccountDetailsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
