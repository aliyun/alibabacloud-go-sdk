// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MoveMaterialFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *MoveMaterialFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *MoveMaterialFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *MoveMaterialFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *MoveMaterialFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *MoveMaterialFileResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *MoveMaterialFileResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *MoveMaterialFileResponseBody
	GetErrorMsg() *string
	SetModule(v bool) *MoveMaterialFileResponseBody
	GetModule() *bool
	SetRequestId(v string) *MoveMaterialFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *MoveMaterialFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *MoveMaterialFileResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *MoveMaterialFileResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *MoveMaterialFileResponseBody
	GetSynchro() *bool
}

type MoveMaterialFileResponseBody struct {
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
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s MoveMaterialFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveMaterialFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MoveMaterialFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *MoveMaterialFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *MoveMaterialFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *MoveMaterialFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *MoveMaterialFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *MoveMaterialFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MoveMaterialFileResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MoveMaterialFileResponseBody) GetModule() *bool {
	return s.Module
}

func (s *MoveMaterialFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveMaterialFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *MoveMaterialFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *MoveMaterialFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveMaterialFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *MoveMaterialFileResponseBody) SetAccessDeniedDetail(v string) *MoveMaterialFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetAllowRetry(v bool) *MoveMaterialFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetAppName(v string) *MoveMaterialFileResponseBody {
	s.AppName = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetDynamicCode(v string) *MoveMaterialFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetDynamicMessage(v string) *MoveMaterialFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetErrorArgs(v []interface{}) *MoveMaterialFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *MoveMaterialFileResponseBody) SetErrorCode(v string) *MoveMaterialFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetErrorMsg(v string) *MoveMaterialFileResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetModule(v bool) *MoveMaterialFileResponseBody {
	s.Module = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetRequestId(v string) *MoveMaterialFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetRootErrorCode(v string) *MoveMaterialFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetRootErrorMsg(v string) *MoveMaterialFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetSuccess(v bool) *MoveMaterialFileResponseBody {
	s.Success = &v
	return s
}

func (s *MoveMaterialFileResponseBody) SetSynchro(v bool) *MoveMaterialFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *MoveMaterialFileResponseBody) Validate() error {
	return dara.Validate(s)
}
