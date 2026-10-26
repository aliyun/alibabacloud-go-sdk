// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UninstallAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UninstallAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UninstallAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UninstallAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UninstallAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UninstallAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UninstallAppPluginResponseBodyModule) *UninstallAppPluginResponseBody
	GetModule() *UninstallAppPluginResponseBodyModule
	SetRequestId(v string) *UninstallAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UninstallAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UninstallAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UninstallAppPluginResponseBody
	GetSynchro() *bool
}

type UninstallAppPluginResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application.
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
	// The dynamic message. This parameter is not used and can be ignored.
	//
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *UninstallAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UninstallAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UninstallAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UninstallAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UninstallAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UninstallAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UninstallAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UninstallAppPluginResponseBody) GetModule() *UninstallAppPluginResponseBodyModule {
	return s.Module
}

func (s *UninstallAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UninstallAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UninstallAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UninstallAppPluginResponseBody) SetAccessDeniedDetail(v string) *UninstallAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetAllowRetry(v bool) *UninstallAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetAppName(v string) *UninstallAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetDynamicCode(v string) *UninstallAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetDynamicMessage(v string) *UninstallAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetErrorArgs(v []interface{}) *UninstallAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UninstallAppPluginResponseBody) SetModule(v *UninstallAppPluginResponseBodyModule) *UninstallAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *UninstallAppPluginResponseBody) SetRequestId(v string) *UninstallAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetRootErrorCode(v string) *UninstallAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetRootErrorMsg(v string) *UninstallAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UninstallAppPluginResponseBody) SetSynchro(v bool) *UninstallAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *UninstallAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UninstallAppPluginResponseBodyModule struct {
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UninstallAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UninstallAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *UninstallAppPluginResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *UninstallAppPluginResponseBodyModule) SetStatus(v string) *UninstallAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *UninstallAppPluginResponseBodyModule) SetSuccess(v bool) *UninstallAppPluginResponseBodyModule {
	s.Success = &v
	return s
}

func (s *UninstallAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
