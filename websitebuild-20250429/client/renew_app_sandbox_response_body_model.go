// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppSandboxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RenewAppSandboxResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RenewAppSandboxResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RenewAppSandboxResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RenewAppSandboxResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RenewAppSandboxResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RenewAppSandboxResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RenewAppSandboxResponseBodyModule) *RenewAppSandboxResponseBody
	GetModule() *RenewAppSandboxResponseBodyModule
	SetRequestId(v string) *RenewAppSandboxResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RenewAppSandboxResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RenewAppSandboxResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RenewAppSandboxResponseBody
	GetSynchro() *bool
}

type RenewAppSandboxResponseBody struct {
	// The detailed reason why access was denied.
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
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the move was successful.
	Module *RenewAppSandboxResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s RenewAppSandboxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAppSandboxResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppSandboxResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RenewAppSandboxResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RenewAppSandboxResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RenewAppSandboxResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RenewAppSandboxResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RenewAppSandboxResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RenewAppSandboxResponseBody) GetModule() *RenewAppSandboxResponseBodyModule {
	return s.Module
}

func (s *RenewAppSandboxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAppSandboxResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RenewAppSandboxResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RenewAppSandboxResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RenewAppSandboxResponseBody) SetAccessDeniedDetail(v string) *RenewAppSandboxResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetAllowRetry(v bool) *RenewAppSandboxResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetAppName(v string) *RenewAppSandboxResponseBody {
	s.AppName = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetDynamicCode(v string) *RenewAppSandboxResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetDynamicMessage(v string) *RenewAppSandboxResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetErrorArgs(v []interface{}) *RenewAppSandboxResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RenewAppSandboxResponseBody) SetModule(v *RenewAppSandboxResponseBodyModule) *RenewAppSandboxResponseBody {
	s.Module = v
	return s
}

func (s *RenewAppSandboxResponseBody) SetRequestId(v string) *RenewAppSandboxResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetRootErrorCode(v string) *RenewAppSandboxResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetRootErrorMsg(v string) *RenewAppSandboxResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RenewAppSandboxResponseBody) SetSynchro(v bool) *RenewAppSandboxResponseBody {
	s.Synchro = &v
	return s
}

func (s *RenewAppSandboxResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewAppSandboxResponseBodyModule struct {
	// The preview URLs.
	PreviewUrls map[string]*string `json:"PreviewUrls,omitempty" xml:"PreviewUrls,omitempty"`
}

func (s RenewAppSandboxResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RenewAppSandboxResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RenewAppSandboxResponseBodyModule) GetPreviewUrls() map[string]*string {
	return s.PreviewUrls
}

func (s *RenewAppSandboxResponseBodyModule) SetPreviewUrls(v map[string]*string) *RenewAppSandboxResponseBodyModule {
	s.PreviewUrls = v
	return s
}

func (s *RenewAppSandboxResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
