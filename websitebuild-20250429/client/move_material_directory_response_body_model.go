// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveMaterialDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MoveMaterialDirectoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *MoveMaterialDirectoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *MoveMaterialDirectoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *MoveMaterialDirectoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *MoveMaterialDirectoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *MoveMaterialDirectoryResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *MoveMaterialDirectoryResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *MoveMaterialDirectoryResponseBody
	GetErrorMsg() *string
	SetModule(v bool) *MoveMaterialDirectoryResponseBody
	GetModule() *bool
	SetRequestId(v string) *MoveMaterialDirectoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *MoveMaterialDirectoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *MoveMaterialDirectoryResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *MoveMaterialDirectoryResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *MoveMaterialDirectoryResponseBody
	GetSynchro() *bool
}

type MoveMaterialDirectoryResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// application Name. Query the application with this Name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// error code. The value is described as follows: If the request succeeded, the ErrorCode field is not returned. If the request failed, the ErrorCode field is returned. For more information, see the error code List in this topic.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// error message.
	//
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// is shift succeeded
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
	// error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// identity indicating whether the request succeeded or failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Backup parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s MoveMaterialDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveMaterialDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *MoveMaterialDirectoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MoveMaterialDirectoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *MoveMaterialDirectoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *MoveMaterialDirectoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *MoveMaterialDirectoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *MoveMaterialDirectoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *MoveMaterialDirectoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MoveMaterialDirectoryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *MoveMaterialDirectoryResponseBody) GetModule() *bool {
	return s.Module
}

func (s *MoveMaterialDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveMaterialDirectoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *MoveMaterialDirectoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *MoveMaterialDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveMaterialDirectoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *MoveMaterialDirectoryResponseBody) SetAccessDeniedDetail(v string) *MoveMaterialDirectoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetAllowRetry(v bool) *MoveMaterialDirectoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetAppName(v string) *MoveMaterialDirectoryResponseBody {
	s.AppName = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetDynamicCode(v string) *MoveMaterialDirectoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetDynamicMessage(v string) *MoveMaterialDirectoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetErrorArgs(v []interface{}) *MoveMaterialDirectoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetErrorCode(v string) *MoveMaterialDirectoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetErrorMsg(v string) *MoveMaterialDirectoryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetModule(v bool) *MoveMaterialDirectoryResponseBody {
	s.Module = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetRequestId(v string) *MoveMaterialDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetRootErrorCode(v string) *MoveMaterialDirectoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetRootErrorMsg(v string) *MoveMaterialDirectoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetSuccess(v bool) *MoveMaterialDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) SetSynchro(v bool) *MoveMaterialDirectoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *MoveMaterialDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
