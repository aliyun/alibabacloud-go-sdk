// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPromotionOfferForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GrantPromotionOfferForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GrantPromotionOfferForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GrantPromotionOfferForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GrantPromotionOfferForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GrantPromotionOfferForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GrantPromotionOfferForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GrantPromotionOfferForPartnerResponseBodyModule) *GrantPromotionOfferForPartnerResponseBody
	GetModule() *GrantPromotionOfferForPartnerResponseBodyModule
	SetRequestId(v string) *GrantPromotionOfferForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GrantPromotionOfferForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GrantPromotionOfferForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GrantPromotionOfferForPartnerResponseBody
	GetSynchro() *bool
}

type GrantPromotionOfferForPartnerResponseBody struct {
	// The details of the permission verification failure.
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
	// The application name. The name can contain digits, letters, and hyphens (-). It must start with a letter and cannot end with a hyphen (-). The name cannot exceed 36 characters in length.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This parameter is not in use. Ignore it.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *GrantPromotionOfferForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The exception message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GrantPromotionOfferForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantPromotionOfferForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetModule() *GrantPromotionOfferForPartnerResponseBodyModule {
	return s.Module
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GrantPromotionOfferForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetAccessDeniedDetail(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetAllowRetry(v bool) *GrantPromotionOfferForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetAppName(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetDynamicCode(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetDynamicMessage(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetErrorArgs(v []interface{}) *GrantPromotionOfferForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetModule(v *GrantPromotionOfferForPartnerResponseBodyModule) *GrantPromotionOfferForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetRequestId(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetRootErrorCode(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetRootErrorMsg(v string) *GrantPromotionOfferForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) SetSynchro(v bool) *GrantPromotionOfferForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GrantPromotionOfferForPartnerResponseBodyModule struct {
	// The message.
	//
	// example:
	//
	// 系统异常
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The benefit summary.
	//
	// example:
	//
	// {}
	OfferSummary *string `json:"OfferSummary,omitempty" xml:"OfferSummary,omitempty"`
	// The distribution record ID.
	//
	// example:
	//
	// 5094
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantPromotionOfferForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GrantPromotionOfferForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) GetMessage() *string {
	return s.Message
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) GetOfferSummary() *string {
	return s.OfferSummary
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) GetRecordId() *string {
	return s.RecordId
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) SetMessage(v string) *GrantPromotionOfferForPartnerResponseBodyModule {
	s.Message = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) SetOfferSummary(v string) *GrantPromotionOfferForPartnerResponseBodyModule {
	s.OfferSummary = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) SetRecordId(v string) *GrantPromotionOfferForPartnerResponseBodyModule {
	s.RecordId = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) SetSuccess(v bool) *GrantPromotionOfferForPartnerResponseBodyModule {
	s.Success = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
