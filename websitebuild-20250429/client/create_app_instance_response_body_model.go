// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppInstanceResponseBodyModule) *CreateAppInstanceResponseBody
	GetModule() *CreateAppInstanceResponseBodyModule
	SetRequestId(v string) *CreateAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppInstanceResponseBody
	GetSynchro() *bool
}

type CreateAppInstanceResponseBody struct {
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
	// Application name, query this application by name
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it means that the input parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *CreateAppInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppInstanceResponseBody) GetModule() *CreateAppInstanceResponseBodyModule {
	return s.Module
}

func (s *CreateAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppInstanceResponseBody) SetAccessDeniedDetail(v string) *CreateAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetAllowRetry(v bool) *CreateAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetAppName(v string) *CreateAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetDynamicCode(v string) *CreateAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetDynamicMessage(v string) *CreateAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetErrorArgs(v []interface{}) *CreateAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppInstanceResponseBody) SetModule(v *CreateAppInstanceResponseBodyModule) *CreateAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppInstanceResponseBody) SetRequestId(v string) *CreateAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetRootErrorCode(v string) *CreateAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetRootErrorMsg(v string) *CreateAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetSynchro(v bool) *CreateAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppInstanceResponseBodyModule struct {
	// Business ID
	//
	// example:
	//
	// WS20250915163734000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// idaas-cn-7mz2uc8v902
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Order ID
	//
	// example:
	//
	// 250822465990301
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// siteId
	//
	// example:
	//
	// abcd.scd.wanwang.xin
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
}

func (s CreateAppInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppInstanceResponseBodyModule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAppInstanceResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAppInstanceResponseBodyModule) GetSiteHost() *string {
	return s.SiteHost
}

func (s *CreateAppInstanceResponseBodyModule) SetBizId(v string) *CreateAppInstanceResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *CreateAppInstanceResponseBodyModule) SetInstanceId(v string) *CreateAppInstanceResponseBodyModule {
	s.InstanceId = &v
	return s
}

func (s *CreateAppInstanceResponseBodyModule) SetOrderId(v string) *CreateAppInstanceResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *CreateAppInstanceResponseBodyModule) SetSiteHost(v string) *CreateAppInstanceResponseBodyModule {
	s.SiteHost = &v
	return s
}

func (s *CreateAppInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
