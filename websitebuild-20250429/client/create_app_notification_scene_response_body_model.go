// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppNotificationSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppNotificationSceneResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppNotificationSceneResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppNotificationSceneResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppNotificationSceneResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppNotificationSceneResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppNotificationSceneResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppNotificationSceneResponseBodyModule) *CreateAppNotificationSceneResponseBody
	GetModule() *CreateAppNotificationSceneResponseBodyModule
	SetRequestId(v string) *CreateAppNotificationSceneResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppNotificationSceneResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppNotificationSceneResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppNotificationSceneResponseBody
	GetSynchro() *bool
}

type CreateAppNotificationSceneResponseBody struct {
	// The detailed reason why access is denied.
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
	// spring-cloud-b
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
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *CreateAppNotificationSceneResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Indicates whether the request is synchronously processed.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateAppNotificationSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppNotificationSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppNotificationSceneResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppNotificationSceneResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppNotificationSceneResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppNotificationSceneResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppNotificationSceneResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppNotificationSceneResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppNotificationSceneResponseBody) GetModule() *CreateAppNotificationSceneResponseBodyModule {
	return s.Module
}

func (s *CreateAppNotificationSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppNotificationSceneResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppNotificationSceneResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppNotificationSceneResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppNotificationSceneResponseBody) SetAccessDeniedDetail(v string) *CreateAppNotificationSceneResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetAllowRetry(v bool) *CreateAppNotificationSceneResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetAppName(v string) *CreateAppNotificationSceneResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetDynamicCode(v string) *CreateAppNotificationSceneResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetDynamicMessage(v string) *CreateAppNotificationSceneResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetErrorArgs(v []interface{}) *CreateAppNotificationSceneResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetModule(v *CreateAppNotificationSceneResponseBodyModule) *CreateAppNotificationSceneResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetRequestId(v string) *CreateAppNotificationSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetRootErrorCode(v string) *CreateAppNotificationSceneResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetRootErrorMsg(v string) *CreateAppNotificationSceneResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) SetSynchro(v bool) *CreateAppNotificationSceneResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppNotificationSceneResponseBodyModule struct {
	// The ID of the created scenario.
	//
	// example:
	//
	// 16257
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateAppNotificationSceneResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppNotificationSceneResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppNotificationSceneResponseBodyModule) GetId() *string {
	return s.Id
}

func (s *CreateAppNotificationSceneResponseBodyModule) SetId(v string) *CreateAppNotificationSceneResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CreateAppNotificationSceneResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
