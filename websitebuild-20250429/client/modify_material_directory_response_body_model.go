// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyMaterialDirectoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ModifyMaterialDirectoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ModifyMaterialDirectoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ModifyMaterialDirectoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyMaterialDirectoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ModifyMaterialDirectoryResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *ModifyMaterialDirectoryResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ModifyMaterialDirectoryResponseBody
	GetErrorMsg() *string
	SetModule(v bool) *ModifyMaterialDirectoryResponseBody
	GetModule() *bool
	SetRequestId(v string) *ModifyMaterialDirectoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ModifyMaterialDirectoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ModifyMaterialDirectoryResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *ModifyMaterialDirectoryResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *ModifyMaterialDirectoryResponseBody
	GetSynchro() *bool
}

type ModifyMaterialDirectoryResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
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
	// The dynamic error message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code. If the request is successful, this field is not returned. If the request fails, this field is returned. For more information, see the error codes in this topic.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the folder was modified successfully.
	//
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
	// The root error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The root error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether the request was processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ModifyMaterialDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaterialDirectoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyMaterialDirectoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ModifyMaterialDirectoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ModifyMaterialDirectoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyMaterialDirectoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyMaterialDirectoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ModifyMaterialDirectoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyMaterialDirectoryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ModifyMaterialDirectoryResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ModifyMaterialDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaterialDirectoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ModifyMaterialDirectoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ModifyMaterialDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMaterialDirectoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ModifyMaterialDirectoryResponseBody) SetAccessDeniedDetail(v string) *ModifyMaterialDirectoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetAllowRetry(v bool) *ModifyMaterialDirectoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetAppName(v string) *ModifyMaterialDirectoryResponseBody {
	s.AppName = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetDynamicCode(v string) *ModifyMaterialDirectoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetDynamicMessage(v string) *ModifyMaterialDirectoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetErrorArgs(v []interface{}) *ModifyMaterialDirectoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetErrorCode(v string) *ModifyMaterialDirectoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetErrorMsg(v string) *ModifyMaterialDirectoryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetModule(v bool) *ModifyMaterialDirectoryResponseBody {
	s.Module = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetRequestId(v string) *ModifyMaterialDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetRootErrorCode(v string) *ModifyMaterialDirectoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetRootErrorMsg(v string) *ModifyMaterialDirectoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetSuccess(v bool) *ModifyMaterialDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) SetSynchro(v bool) *ModifyMaterialDirectoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *ModifyMaterialDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
