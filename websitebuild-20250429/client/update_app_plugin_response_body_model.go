// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UpdateAppPluginResponseBodyModule) *UpdateAppPluginResponseBody
	GetModule() *UpdateAppPluginResponseBodyModule
	SetRequestId(v string) *UpdateAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppPluginResponseBody
	GetSynchro() *bool
}

type UpdateAppPluginResponseBody struct {
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
	// The application name must start with a letter and can contain letters, digits, underscores (_), and hyphens (-). The name can be up to 36 characters in length.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This parameter is not in use. Ignore it.
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
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the move is successful.
	Module *UpdateAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UpdateAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppPluginResponseBody) GetModule() *UpdateAppPluginResponseBodyModule {
	return s.Module
}

func (s *UpdateAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppPluginResponseBody) SetAccessDeniedDetail(v string) *UpdateAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetAllowRetry(v bool) *UpdateAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetAppName(v string) *UpdateAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetDynamicCode(v string) *UpdateAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetDynamicMessage(v string) *UpdateAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetErrorArgs(v []interface{}) *UpdateAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppPluginResponseBody) SetModule(v *UpdateAppPluginResponseBodyModule) *UpdateAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *UpdateAppPluginResponseBody) SetRequestId(v string) *UpdateAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetRootErrorCode(v string) *UpdateAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetRootErrorMsg(v string) *UpdateAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppPluginResponseBody) SetSynchro(v bool) *UpdateAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppPluginResponseBodyModule struct {
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UpdateAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *UpdateAppPluginResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAppPluginResponseBodyModule) SetStatus(v string) *UpdateAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *UpdateAppPluginResponseBodyModule) SetSuccess(v bool) *UpdateAppPluginResponseBodyModule {
	s.Success = &v
	return s
}

func (s *UpdateAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
