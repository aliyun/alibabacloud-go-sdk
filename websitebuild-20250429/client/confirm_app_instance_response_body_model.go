// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ConfirmAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ConfirmAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ConfirmAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ConfirmAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ConfirmAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ConfirmAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ConfirmAppInstanceResponseBodyModule) *ConfirmAppInstanceResponseBody
	GetModule() *ConfirmAppInstanceResponseBodyModule
	SetRequestId(v string) *ConfirmAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ConfirmAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ConfirmAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ConfirmAppInstanceResponseBody
	GetSynchro() *bool
}

type ConfirmAppInstanceResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                         `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *ConfirmAppInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ConfirmAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ConfirmAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ConfirmAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ConfirmAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ConfirmAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ConfirmAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ConfirmAppInstanceResponseBody) GetModule() *ConfirmAppInstanceResponseBodyModule {
	return s.Module
}

func (s *ConfirmAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ConfirmAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ConfirmAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ConfirmAppInstanceResponseBody) SetAccessDeniedDetail(v string) *ConfirmAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetAllowRetry(v bool) *ConfirmAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetAppName(v string) *ConfirmAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetDynamicCode(v string) *ConfirmAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetDynamicMessage(v string) *ConfirmAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetErrorArgs(v []interface{}) *ConfirmAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetModule(v *ConfirmAppInstanceResponseBodyModule) *ConfirmAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetRequestId(v string) *ConfirmAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetRootErrorCode(v string) *ConfirmAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetRootErrorMsg(v string) *ConfirmAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) SetSynchro(v bool) *ConfirmAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *ConfirmAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfirmAppInstanceResponseBodyModule struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// idaas-cn-7mz2uc8v902
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 247748990880615
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// abcd.scd.wanwang.xin
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
}

func (s ConfirmAppInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAppInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ConfirmAppInstanceResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *ConfirmAppInstanceResponseBodyModule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfirmAppInstanceResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *ConfirmAppInstanceResponseBodyModule) GetSiteHost() *string {
	return s.SiteHost
}

func (s *ConfirmAppInstanceResponseBodyModule) SetBizId(v string) *ConfirmAppInstanceResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *ConfirmAppInstanceResponseBodyModule) SetInstanceId(v string) *ConfirmAppInstanceResponseBodyModule {
	s.InstanceId = &v
	return s
}

func (s *ConfirmAppInstanceResponseBodyModule) SetOrderId(v string) *ConfirmAppInstanceResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *ConfirmAppInstanceResponseBodyModule) SetSiteHost(v string) *ConfirmAppInstanceResponseBodyModule {
	s.SiteHost = &v
	return s
}

func (s *ConfirmAppInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
