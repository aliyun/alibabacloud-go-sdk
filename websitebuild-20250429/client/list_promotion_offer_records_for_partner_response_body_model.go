// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionOfferRecordsForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListPromotionOfferRecordsForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListPromotionOfferRecordsForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListPromotionOfferRecordsForPartnerResponseBody
	GetMaxResults() *int32
	SetModule(v *ListPromotionOfferRecordsForPartnerResponseBodyModule) *ListPromotionOfferRecordsForPartnerResponseBody
	GetModule() *ListPromotionOfferRecordsForPartnerResponseBodyModule
	SetNextToken(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListPromotionOfferRecordsForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListPromotionOfferRecordsForPartnerResponseBody
	GetSynchro() *bool
}

type ListPromotionOfferRecordsForPartnerResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether a retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name. The application with this name is queried.
	//
	// example:
	//
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the %s variable in the ErrMessage response parameter.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response data.
	Module *ListPromotionOfferRecordsForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// test
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListPromotionOfferRecordsForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionOfferRecordsForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetModule() *ListPromotionOfferRecordsForPartnerResponseBodyModule {
	return s.Module
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetAccessDeniedDetail(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetAllowRetry(v bool) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetAppName(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetDynamicCode(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetDynamicMessage(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetErrorArgs(v []interface{}) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetMaxResults(v int32) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetModule(v *ListPromotionOfferRecordsForPartnerResponseBodyModule) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetNextToken(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetRequestId(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetRootErrorCode(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetRootErrorMsg(v string) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) SetSynchro(v bool) *ListPromotionOfferRecordsForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPromotionOfferRecordsForPartnerResponseBodyModule struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of issuance records.
	Records []*ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPromotionOfferRecordsForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionOfferRecordsForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) GetRecords() []*ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	return s.Records
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) SetPageNum(v int32) *ListPromotionOfferRecordsForPartnerResponseBodyModule {
	s.PageNum = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) SetPageSize(v int32) *ListPromotionOfferRecordsForPartnerResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) SetRecords(v []*ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) *ListPromotionOfferRecordsForPartnerResponseBodyModule {
	s.Records = v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) SetTotalCount(v int32) *ListPromotionOfferRecordsForPartnerResponseBodyModule {
	s.TotalCount = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModule) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords struct {
	// The activity code.
	//
	// example:
	//
	// acwfradoj5u
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// The activity name.
	//
	// example:
	//
	// IP网段过滤统计
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The employee code.
	//
	// example:
	//
	// 1234
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// SYSTEM_ERROR
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The issuance time.
	//
	// example:
	//
	// 2025-06-01
	IssuanceTime *string `json:"IssuanceTime,omitempty" xml:"IssuanceTime,omitempty"`
	// The benefit snapshot in JSON format.
	//
	// example:
	//
	// {}
	OfferSnapshot *string `json:"OfferSnapshot,omitempty" xml:"OfferSnapshot,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 5094
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The operation remark.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The issuance status.
	//
	// example:
	//
	// FE_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123241414
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GoString() string {
	return s.String()
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetActivityName() *string {
	return s.ActivityName
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetFailReason() *string {
	return s.FailReason
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetIssuanceTime() *string {
	return s.IssuanceTime
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetOfferSnapshot() *string {
	return s.OfferSnapshot
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetRecordId() *string {
	return s.RecordId
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetRemark() *string {
	return s.Remark
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetStatus() *string {
	return s.Status
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) GetUserId() *string {
	return s.UserId
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetActivityCode(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.ActivityCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetActivityName(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.ActivityName = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetEmployeeCode(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.EmployeeCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetFailReason(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.FailReason = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetIssuanceTime(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.IssuanceTime = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetOfferSnapshot(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.OfferSnapshot = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetRecordId(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.RecordId = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetRemark(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.Remark = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetStatus(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.Status = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) SetUserId(v string) *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords {
	s.UserId = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponseBodyModuleRecords) Validate() error {
	return dara.Validate(s)
}
