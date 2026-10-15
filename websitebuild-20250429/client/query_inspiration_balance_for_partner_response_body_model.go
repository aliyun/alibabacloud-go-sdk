// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationBalanceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryInspirationBalanceForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryInspirationBalanceForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryInspirationBalanceForPartnerResponseBodyModule) *QueryInspirationBalanceForPartnerResponseBody
	GetModule() *QueryInspirationBalanceForPartnerResponseBodyModule
	SetRequestId(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryInspirationBalanceForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryInspirationBalanceForPartnerResponseBody
	GetSynchro() *bool
}

type QueryInspirationBalanceForPartnerResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *QueryInspirationBalanceForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryInspirationBalanceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetModule() *QueryInspirationBalanceForPartnerResponseBodyModule {
	return s.Module
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryInspirationBalanceForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetAccessDeniedDetail(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetAllowRetry(v bool) *QueryInspirationBalanceForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetAppName(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetDynamicCode(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetDynamicMessage(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetErrorArgs(v []interface{}) *QueryInspirationBalanceForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetModule(v *QueryInspirationBalanceForPartnerResponseBodyModule) *QueryInspirationBalanceForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetRequestId(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetRootErrorCode(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetRootErrorMsg(v string) *QueryInspirationBalanceForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) SetSynchro(v bool) *QueryInspirationBalanceForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationBalanceForPartnerResponseBodyModule struct {
	// The remaining amount, calculated as totalQuota minus totalUsed.
	//
	// example:
	//
	// 8
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// The remaining amount as a precise value. This field is of the String type and supports decimal display.
	//
	// example:
	//
	// 10
	RemainingStr *string `json:"RemainingStr,omitempty" xml:"RemainingStr,omitempty"`
	// The total quota, which is the sum of initQuota for all valid accounts.
	//
	// example:
	//
	// 10
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The total quota as a precise value. This field is of the String type and supports decimal display.
	//
	// example:
	//
	// 10
	TotalQuotaStr *string `json:"TotalQuotaStr,omitempty" xml:"TotalQuotaStr,omitempty"`
	// The total consumed amount, which is the sum of used for all valid accounts.
	//
	// example:
	//
	// 1
	TotalUsed *int64 `json:"TotalUsed,omitempty" xml:"TotalUsed,omitempty"`
	// The total consumed amount as a precise value. This field is of the String type and supports decimal display.
	//
	// example:
	//
	// 10
	TotalUsedStr *string `json:"TotalUsedStr,omitempty" xml:"TotalUsedStr,omitempty"`
}

func (s QueryInspirationBalanceForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetRemaining() *int64 {
	return s.Remaining
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetRemainingStr() *string {
	return s.RemainingStr
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetTotalQuotaStr() *string {
	return s.TotalQuotaStr
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetTotalUsed() *int64 {
	return s.TotalUsed
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) GetTotalUsedStr() *string {
	return s.TotalUsedStr
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetRemaining(v int64) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.Remaining = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetRemainingStr(v string) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.RemainingStr = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetTotalQuota(v int64) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.TotalQuota = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetTotalQuotaStr(v string) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.TotalQuotaStr = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetTotalUsed(v int64) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.TotalUsed = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) SetTotalUsedStr(v string) *QueryInspirationBalanceForPartnerResponseBodyModule {
	s.TotalUsedStr = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
