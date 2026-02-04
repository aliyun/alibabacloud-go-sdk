// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainRedirectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppDomainRedirectResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppDomainRedirectResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppDomainRedirectResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppDomainRedirectResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppDomainRedirectResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppDomainRedirectResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *DeleteAppDomainRedirectResponseBodyModule) *DeleteAppDomainRedirectResponseBody
	GetModule() *DeleteAppDomainRedirectResponseBodyModule
	SetRequestId(v string) *DeleteAppDomainRedirectResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppDomainRedirectResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppDomainRedirectResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppDomainRedirectResponseBody
	GetSynchro() *bool
}

type DeleteAppDomainRedirectResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name, to query this named application
	//
	// example:
	//
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *DeleteAppDomainRedirectResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Exception message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Whether it is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DeleteAppDomainRedirectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainRedirectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainRedirectResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppDomainRedirectResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppDomainRedirectResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppDomainRedirectResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppDomainRedirectResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppDomainRedirectResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppDomainRedirectResponseBody) GetModule() *DeleteAppDomainRedirectResponseBodyModule {
	return s.Module
}

func (s *DeleteAppDomainRedirectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppDomainRedirectResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppDomainRedirectResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppDomainRedirectResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppDomainRedirectResponseBody) SetAccessDeniedDetail(v string) *DeleteAppDomainRedirectResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetAllowRetry(v bool) *DeleteAppDomainRedirectResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetAppName(v string) *DeleteAppDomainRedirectResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetDynamicCode(v string) *DeleteAppDomainRedirectResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetDynamicMessage(v string) *DeleteAppDomainRedirectResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetErrorArgs(v []interface{}) *DeleteAppDomainRedirectResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetModule(v *DeleteAppDomainRedirectResponseBodyModule) *DeleteAppDomainRedirectResponseBody {
	s.Module = v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetRequestId(v string) *DeleteAppDomainRedirectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetRootErrorCode(v string) *DeleteAppDomainRedirectResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetRootErrorMsg(v string) *DeleteAppDomainRedirectResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) SetSynchro(v bool) *DeleteAppDomainRedirectResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppDomainRedirectResponseBodyModule struct {
	// Whether the operation was successful
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppDomainRedirectResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainRedirectResponseBodyModule) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainRedirectResponseBodyModule) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppDomainRedirectResponseBodyModule) SetSuccess(v bool) *DeleteAppDomainRedirectResponseBodyModule {
	s.Success = &v
	return s
}

func (s *DeleteAppDomainRedirectResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
