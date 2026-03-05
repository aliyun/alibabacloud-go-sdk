// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInspirationBalanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryInspirationBalanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryInspirationBalanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryInspirationBalanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryInspirationBalanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryInspirationBalanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryInspirationBalanceResponseBodyModule) *QueryInspirationBalanceResponseBody
	GetModule() *QueryInspirationBalanceResponseBodyModule
	SetRequestId(v string) *QueryInspirationBalanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryInspirationBalanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryInspirationBalanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryInspirationBalanceResponseBody
	GetSynchro() *bool
}

type QueryInspirationBalanceResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                    `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                              `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QueryInspirationBalanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryInspirationBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInspirationBalanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryInspirationBalanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryInspirationBalanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryInspirationBalanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryInspirationBalanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryInspirationBalanceResponseBody) GetModule() *QueryInspirationBalanceResponseBodyModule {
	return s.Module
}

func (s *QueryInspirationBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInspirationBalanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryInspirationBalanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryInspirationBalanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryInspirationBalanceResponseBody) SetAccessDeniedDetail(v string) *QueryInspirationBalanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetAllowRetry(v bool) *QueryInspirationBalanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetAppName(v string) *QueryInspirationBalanceResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetDynamicCode(v string) *QueryInspirationBalanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetDynamicMessage(v string) *QueryInspirationBalanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetErrorArgs(v []interface{}) *QueryInspirationBalanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetModule(v *QueryInspirationBalanceResponseBodyModule) *QueryInspirationBalanceResponseBody {
	s.Module = v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetRequestId(v string) *QueryInspirationBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetRootErrorCode(v string) *QueryInspirationBalanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetRootErrorMsg(v string) *QueryInspirationBalanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) SetSynchro(v bool) *QueryInspirationBalanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryInspirationBalanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationBalanceResponseBodyModule struct {
	// example:
	//
	// 8
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// example:
	//
	// 10
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 1
	TotalUsed *int64 `json:"TotalUsed,omitempty" xml:"TotalUsed,omitempty"`
}

func (s QueryInspirationBalanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceResponseBodyModule) GetRemaining() *int64 {
	return s.Remaining
}

func (s *QueryInspirationBalanceResponseBodyModule) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *QueryInspirationBalanceResponseBodyModule) GetTotalUsed() *int64 {
	return s.TotalUsed
}

func (s *QueryInspirationBalanceResponseBodyModule) SetRemaining(v int64) *QueryInspirationBalanceResponseBodyModule {
	s.Remaining = &v
	return s
}

func (s *QueryInspirationBalanceResponseBodyModule) SetTotalQuota(v int64) *QueryInspirationBalanceResponseBodyModule {
	s.TotalQuota = &v
	return s
}

func (s *QueryInspirationBalanceResponseBodyModule) SetTotalUsed(v int64) *QueryInspirationBalanceResponseBodyModule {
	s.TotalUsed = &v
	return s
}

func (s *QueryInspirationBalanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
