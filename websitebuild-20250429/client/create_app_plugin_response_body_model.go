// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppPluginResponseBodyModule) *CreateAppPluginResponseBody
	GetModule() *CreateAppPluginResponseBodyModule
	SetRequestId(v string) *CreateAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppPluginResponseBody
	GetSynchro() *bool
}

type CreateAppPluginResponseBody struct {
	// The detailed reasons for the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether a retry is allowed.
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
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
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
	Module *CreateAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppPluginResponseBody) GetModule() *CreateAppPluginResponseBodyModule {
	return s.Module
}

func (s *CreateAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppPluginResponseBody) SetAccessDeniedDetail(v string) *CreateAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetAllowRetry(v bool) *CreateAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetAppName(v string) *CreateAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetDynamicCode(v string) *CreateAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetDynamicMessage(v string) *CreateAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetErrorArgs(v []interface{}) *CreateAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppPluginResponseBody) SetModule(v *CreateAppPluginResponseBodyModule) *CreateAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppPluginResponseBody) SetRequestId(v string) *CreateAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetRootErrorCode(v string) *CreateAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetRootErrorMsg(v string) *CreateAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppPluginResponseBody) SetSynchro(v bool) *CreateAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppPluginResponseBodyModule struct {
	// The ID of the bound API Gateway plug-in.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppPluginResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *CreateAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *CreateAppPluginResponseBodyModule) SetPluginId(v string) *CreateAppPluginResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *CreateAppPluginResponseBodyModule) SetStatus(v string) *CreateAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *CreateAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
