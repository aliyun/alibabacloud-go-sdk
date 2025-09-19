// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnbindAppDomainResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UnbindAppDomainResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UnbindAppDomainResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UnbindAppDomainResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UnbindAppDomainResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UnbindAppDomainResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UnbindAppDomainResponseBodyModule) *UnbindAppDomainResponseBody
	GetModule() *UnbindAppDomainResponseBodyModule
	SetRequestId(v string) *UnbindAppDomainResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UnbindAppDomainResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UnbindAppDomainResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UnbindAppDomainResponseBody
	GetSynchro() *bool
}

type UnbindAppDomainResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/wlr10v1g25549kq/wlr10v1g25549kq.diff.zip?Expires=1730174953&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=F93JAFEuVN63YzNQyUy2xOaOtKs%3D
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                      `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *UnbindAppDomainResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s UnbindAppDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAppDomainResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnbindAppDomainResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UnbindAppDomainResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UnbindAppDomainResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UnbindAppDomainResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UnbindAppDomainResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UnbindAppDomainResponseBody) GetModule() *UnbindAppDomainResponseBodyModule {
	return s.Module
}

func (s *UnbindAppDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAppDomainResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UnbindAppDomainResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UnbindAppDomainResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UnbindAppDomainResponseBody) SetAccessDeniedDetail(v string) *UnbindAppDomainResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetAllowRetry(v bool) *UnbindAppDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetAppName(v string) *UnbindAppDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetDynamicCode(v string) *UnbindAppDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetDynamicMessage(v string) *UnbindAppDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetErrorArgs(v []interface{}) *UnbindAppDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UnbindAppDomainResponseBody) SetModule(v *UnbindAppDomainResponseBodyModule) *UnbindAppDomainResponseBody {
	s.Module = v
	return s
}

func (s *UnbindAppDomainResponseBody) SetRequestId(v string) *UnbindAppDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetRootErrorCode(v string) *UnbindAppDomainResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetRootErrorMsg(v string) *UnbindAppDomainResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UnbindAppDomainResponseBody) SetSynchro(v bool) *UnbindAppDomainResponseBody {
	s.Synchro = &v
	return s
}

func (s *UnbindAppDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnbindAppDomainResponseBodyModule struct {
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAppDomainResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppDomainResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UnbindAppDomainResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindAppDomainResponseBodyModule) SetSuccess(v bool) *UnbindAppDomainResponseBodyModule {
	s.Success = &v
	return s
}

func (s *UnbindAppDomainResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
