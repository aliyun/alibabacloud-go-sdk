// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *DeleteAppPluginResponseBodyModule) *DeleteAppPluginResponseBody
	GetModule() *DeleteAppPluginResponseBodyModule
	SetRequestId(v string) *DeleteAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppPluginResponseBody
	GetSynchro() *bool
}

type DeleteAppPluginResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether a retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// dewuApp
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
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *DeleteAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Indicates whether the processing is synchronous.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DeleteAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppPluginResponseBody) GetModule() *DeleteAppPluginResponseBodyModule {
	return s.Module
}

func (s *DeleteAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppPluginResponseBody) SetAccessDeniedDetail(v string) *DeleteAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetAllowRetry(v bool) *DeleteAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetAppName(v string) *DeleteAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetDynamicCode(v string) *DeleteAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetDynamicMessage(v string) *DeleteAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetErrorArgs(v []interface{}) *DeleteAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppPluginResponseBody) SetModule(v *DeleteAppPluginResponseBodyModule) *DeleteAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *DeleteAppPluginResponseBody) SetRequestId(v string) *DeleteAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetRootErrorCode(v string) *DeleteAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetRootErrorMsg(v string) *DeleteAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppPluginResponseBody) SetSynchro(v bool) *DeleteAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppPluginResponseBodyModule struct {
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DeleteAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *DeleteAppPluginResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppPluginResponseBodyModule) SetStatus(v string) *DeleteAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *DeleteAppPluginResponseBodyModule) SetSuccess(v bool) *DeleteAppPluginResponseBodyModule {
	s.Success = &v
	return s
}

func (s *DeleteAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
