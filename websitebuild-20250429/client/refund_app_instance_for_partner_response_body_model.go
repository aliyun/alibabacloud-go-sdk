// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundAppInstanceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RefundAppInstanceForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RefundAppInstanceForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RefundAppInstanceForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RefundAppInstanceForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RefundAppInstanceForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RefundAppInstanceForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RefundAppInstanceForPartnerResponseBodyModule) *RefundAppInstanceForPartnerResponseBody
	GetModule() *RefundAppInstanceForPartnerResponseBodyModule
	SetRequestId(v string) *RefundAppInstanceForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RefundAppInstanceForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RefundAppInstanceForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RefundAppInstanceForPartnerResponseBody
	GetSynchro() *bool
}

type RefundAppInstanceForPartnerResponseBody struct {
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
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string                                        `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                  `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *RefundAppInstanceForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s RefundAppInstanceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundAppInstanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *RefundAppInstanceForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RefundAppInstanceForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RefundAppInstanceForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RefundAppInstanceForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RefundAppInstanceForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RefundAppInstanceForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RefundAppInstanceForPartnerResponseBody) GetModule() *RefundAppInstanceForPartnerResponseBodyModule {
	return s.Module
}

func (s *RefundAppInstanceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundAppInstanceForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RefundAppInstanceForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RefundAppInstanceForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RefundAppInstanceForPartnerResponseBody) SetAccessDeniedDetail(v string) *RefundAppInstanceForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetAllowRetry(v bool) *RefundAppInstanceForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetAppName(v string) *RefundAppInstanceForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetDynamicCode(v string) *RefundAppInstanceForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetDynamicMessage(v string) *RefundAppInstanceForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetErrorArgs(v []interface{}) *RefundAppInstanceForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetModule(v *RefundAppInstanceForPartnerResponseBodyModule) *RefundAppInstanceForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetRequestId(v string) *RefundAppInstanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetRootErrorCode(v string) *RefundAppInstanceForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetRootErrorMsg(v string) *RefundAppInstanceForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) SetSynchro(v bool) *RefundAppInstanceForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundAppInstanceForPartnerResponseBodyModule struct {
	// example:
	//
	// 250822465990301
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RefundAppInstanceForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RefundAppInstanceForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RefundAppInstanceForPartnerResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *RefundAppInstanceForPartnerResponseBodyModule) SetOrderId(v string) *RefundAppInstanceForPartnerResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
