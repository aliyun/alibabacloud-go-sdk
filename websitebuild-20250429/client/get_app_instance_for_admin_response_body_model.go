// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppInstanceForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppInstanceForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppInstanceForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppInstanceForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppInstanceForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppInstanceForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *AppInstanceAggregate) *GetAppInstanceForAdminResponseBody
	GetModule() *AppInstanceAggregate
	SetRequestId(v string) *GetAppInstanceForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppInstanceForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppInstanceForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppInstanceForAdminResponseBody
	GetSynchro() *bool
}

type GetAppInstanceForAdminResponseBody struct {
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
	// The application name that was queried.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	//
	// example:
	//
	// true
	Module *AppInstanceAggregate `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppInstanceForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppInstanceForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppInstanceForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppInstanceForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppInstanceForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppInstanceForAdminResponseBody) GetModule() *AppInstanceAggregate {
	return s.Module
}

func (s *GetAppInstanceForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppInstanceForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppInstanceForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppInstanceForAdminResponseBody) SetAccessDeniedDetail(v string) *GetAppInstanceForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetAllowRetry(v bool) *GetAppInstanceForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetAppName(v string) *GetAppInstanceForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetDynamicCode(v string) *GetAppInstanceForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetDynamicMessage(v string) *GetAppInstanceForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetErrorArgs(v []interface{}) *GetAppInstanceForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetModule(v *AppInstanceAggregate) *GetAppInstanceForAdminResponseBody {
	s.Module = v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetRequestId(v string) *GetAppInstanceForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetRootErrorCode(v string) *GetAppInstanceForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetRootErrorMsg(v string) *GetAppInstanceForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) SetSynchro(v bool) *GetAppInstanceForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppInstanceForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}
