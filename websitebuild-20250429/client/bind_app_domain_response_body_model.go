// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAppDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *BindAppDomainResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *BindAppDomainResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *BindAppDomainResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *BindAppDomainResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *BindAppDomainResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *BindAppDomainResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *BindAppDomainResponseBodyModule) *BindAppDomainResponseBody
	GetModule() *BindAppDomainResponseBodyModule
	SetRequestId(v string) *BindAppDomainResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *BindAppDomainResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *BindAppDomainResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *BindAppDomainResponseBody
	GetSynchro() *bool
}

type BindAppDomainResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/vs1w2yd3p8264pz/vs1w2yd3p8264pz.diff.zip?Expires=1739592470&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=3DRuMtCeTinVYxo7XAOEIOEmZGE%3D
	DynamicMessage *string                          `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                    `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *BindAppDomainResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s BindAppDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *BindAppDomainResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *BindAppDomainResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *BindAppDomainResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *BindAppDomainResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *BindAppDomainResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *BindAppDomainResponseBody) GetModule() *BindAppDomainResponseBodyModule {
	return s.Module
}

func (s *BindAppDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAppDomainResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *BindAppDomainResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *BindAppDomainResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *BindAppDomainResponseBody) SetAccessDeniedDetail(v string) *BindAppDomainResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *BindAppDomainResponseBody) SetAllowRetry(v bool) *BindAppDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *BindAppDomainResponseBody) SetAppName(v string) *BindAppDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *BindAppDomainResponseBody) SetDynamicCode(v string) *BindAppDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BindAppDomainResponseBody) SetDynamicMessage(v string) *BindAppDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BindAppDomainResponseBody) SetErrorArgs(v []interface{}) *BindAppDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *BindAppDomainResponseBody) SetModule(v *BindAppDomainResponseBodyModule) *BindAppDomainResponseBody {
	s.Module = v
	return s
}

func (s *BindAppDomainResponseBody) SetRequestId(v string) *BindAppDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAppDomainResponseBody) SetRootErrorCode(v string) *BindAppDomainResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *BindAppDomainResponseBody) SetRootErrorMsg(v string) *BindAppDomainResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *BindAppDomainResponseBody) SetSynchro(v bool) *BindAppDomainResponseBody {
	s.Synchro = &v
	return s
}

func (s *BindAppDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type BindAppDomainResponseBodyModule struct {
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAppDomainResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *BindAppDomainResponseBodyModule) SetSuccess(v bool) *BindAppDomainResponseBodyModule {
	s.Success = &v
	return s
}

func (s *BindAppDomainResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
