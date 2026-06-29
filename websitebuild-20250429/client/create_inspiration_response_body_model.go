// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspirationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInspirationResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateInspirationResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateInspirationResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateInspirationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateInspirationResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateInspirationResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateInspirationResponseBodyModule) *CreateInspirationResponseBody
	GetModule() *CreateInspirationResponseBodyModule
	SetRequestId(v string) *CreateInspirationResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateInspirationResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateInspirationResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateInspirationResponseBody
	GetSynchro() *bool
}

type CreateInspirationResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
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
	// The dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *CreateInspirationResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// SYSTEM.EROR
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

func (s CreateInspirationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInspirationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInspirationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInspirationResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateInspirationResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateInspirationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateInspirationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateInspirationResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateInspirationResponseBody) GetModule() *CreateInspirationResponseBodyModule {
	return s.Module
}

func (s *CreateInspirationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInspirationResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateInspirationResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateInspirationResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateInspirationResponseBody) SetAccessDeniedDetail(v string) *CreateInspirationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInspirationResponseBody) SetAllowRetry(v bool) *CreateInspirationResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateInspirationResponseBody) SetAppName(v string) *CreateInspirationResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateInspirationResponseBody) SetDynamicCode(v string) *CreateInspirationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInspirationResponseBody) SetDynamicMessage(v string) *CreateInspirationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInspirationResponseBody) SetErrorArgs(v []interface{}) *CreateInspirationResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateInspirationResponseBody) SetModule(v *CreateInspirationResponseBodyModule) *CreateInspirationResponseBody {
	s.Module = v
	return s
}

func (s *CreateInspirationResponseBody) SetRequestId(v string) *CreateInspirationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInspirationResponseBody) SetRootErrorCode(v string) *CreateInspirationResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateInspirationResponseBody) SetRootErrorMsg(v string) *CreateInspirationResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateInspirationResponseBody) SetSynchro(v bool) *CreateInspirationResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateInspirationResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInspirationResponseBodyModule struct {
	// The order ID.
	//
	// example:
	//
	// 250822465990301
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInspirationResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateInspirationResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateInspirationResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateInspirationResponseBodyModule) SetOrderId(v string) *CreateInspirationResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *CreateInspirationResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
