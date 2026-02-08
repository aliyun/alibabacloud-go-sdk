// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyMaterialFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ModifyMaterialFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ModifyMaterialFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ModifyMaterialFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyMaterialFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ModifyMaterialFileResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *ModifyMaterialFileResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ModifyMaterialFileResponseBody
	GetErrorMsg() *string
	SetModule(v bool) *ModifyMaterialFileResponseBody
	GetModule() *bool
	SetRequestId(v string) *ModifyMaterialFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ModifyMaterialFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ModifyMaterialFileResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *ModifyMaterialFileResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *ModifyMaterialFileResponseBody
	GetSynchro() *bool
}

type ModifyMaterialFileResponseBody struct {
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
	// dewuApp
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

func (s ModifyMaterialFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyMaterialFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ModifyMaterialFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ModifyMaterialFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyMaterialFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyMaterialFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ModifyMaterialFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyMaterialFileResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ModifyMaterialFileResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ModifyMaterialFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaterialFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ModifyMaterialFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ModifyMaterialFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMaterialFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ModifyMaterialFileResponseBody) SetAccessDeniedDetail(v string) *ModifyMaterialFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetAllowRetry(v bool) *ModifyMaterialFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetAppName(v string) *ModifyMaterialFileResponseBody {
	s.AppName = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetDynamicCode(v string) *ModifyMaterialFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetDynamicMessage(v string) *ModifyMaterialFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetErrorArgs(v []interface{}) *ModifyMaterialFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetErrorCode(v string) *ModifyMaterialFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetErrorMsg(v string) *ModifyMaterialFileResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetModule(v bool) *ModifyMaterialFileResponseBody {
	s.Module = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetRequestId(v string) *ModifyMaterialFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetRootErrorCode(v string) *ModifyMaterialFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetRootErrorMsg(v string) *ModifyMaterialFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetSuccess(v bool) *ModifyMaterialFileResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) SetSynchro(v bool) *ModifyMaterialFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *ModifyMaterialFileResponseBody) Validate() error {
	return dara.Validate(s)
}
