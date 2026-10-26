// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *InstallAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *InstallAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *InstallAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *InstallAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *InstallAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *InstallAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *InstallAppPluginResponseBodyModule) *InstallAppPluginResponseBody
	GetModule() *InstallAppPluginResponseBodyModule
	SetRequestId(v string) *InstallAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *InstallAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *InstallAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *InstallAppPluginResponseBody
	GetSynchro() *bool
}

type InstallAppPluginResponseBody struct {
	// The detailed reason for access denial.
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
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the error message returned in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *InstallAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s InstallAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *InstallAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *InstallAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *InstallAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *InstallAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *InstallAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *InstallAppPluginResponseBody) GetModule() *InstallAppPluginResponseBodyModule {
	return s.Module
}

func (s *InstallAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *InstallAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *InstallAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *InstallAppPluginResponseBody) SetAccessDeniedDetail(v string) *InstallAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetAllowRetry(v bool) *InstallAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetAppName(v string) *InstallAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetDynamicCode(v string) *InstallAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetDynamicMessage(v string) *InstallAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetErrorArgs(v []interface{}) *InstallAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *InstallAppPluginResponseBody) SetModule(v *InstallAppPluginResponseBodyModule) *InstallAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *InstallAppPluginResponseBody) SetRequestId(v string) *InstallAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetRootErrorCode(v string) *InstallAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetRootErrorMsg(v string) *InstallAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *InstallAppPluginResponseBody) SetSynchro(v bool) *InstallAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *InstallAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallAppPluginResponseBodyModule struct {
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InstallAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InstallAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *InstallAppPluginResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *InstallAppPluginResponseBodyModule) SetStatus(v string) *InstallAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *InstallAppPluginResponseBodyModule) SetSuccess(v bool) *InstallAppPluginResponseBodyModule {
	s.Success = &v
	return s
}

func (s *InstallAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
