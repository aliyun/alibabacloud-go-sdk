// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMaterialTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitMaterialTaskResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SubmitMaterialTaskResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SubmitMaterialTaskResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SubmitMaterialTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SubmitMaterialTaskResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SubmitMaterialTaskResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *SubmitMaterialTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SubmitMaterialTaskResponseBody
	GetErrorMsg() *string
	SetModule(v *SubmitMaterialTaskResponseBodyModule) *SubmitMaterialTaskResponseBody
	GetModule() *SubmitMaterialTaskResponseBodyModule
	SetRequestId(v string) *SubmitMaterialTaskResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SubmitMaterialTaskResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SubmitMaterialTaskResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *SubmitMaterialTaskResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *SubmitMaterialTaskResponseBody
	GetSynchro() *bool
}

type SubmitMaterialTaskResponseBody struct {
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
	// SYSTEM_ERRROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Module   *SubmitMaterialTaskResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s SubmitMaterialTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMaterialTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMaterialTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitMaterialTaskResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SubmitMaterialTaskResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SubmitMaterialTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SubmitMaterialTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SubmitMaterialTaskResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SubmitMaterialTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitMaterialTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitMaterialTaskResponseBody) GetModule() *SubmitMaterialTaskResponseBodyModule {
	return s.Module
}

func (s *SubmitMaterialTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMaterialTaskResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SubmitMaterialTaskResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SubmitMaterialTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitMaterialTaskResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SubmitMaterialTaskResponseBody) SetAccessDeniedDetail(v string) *SubmitMaterialTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetAllowRetry(v bool) *SubmitMaterialTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetAppName(v string) *SubmitMaterialTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetDynamicCode(v string) *SubmitMaterialTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetDynamicMessage(v string) *SubmitMaterialTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetErrorArgs(v []interface{}) *SubmitMaterialTaskResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetErrorCode(v string) *SubmitMaterialTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetErrorMsg(v string) *SubmitMaterialTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetModule(v *SubmitMaterialTaskResponseBodyModule) *SubmitMaterialTaskResponseBody {
	s.Module = v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetRequestId(v string) *SubmitMaterialTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetRootErrorCode(v string) *SubmitMaterialTaskResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetRootErrorMsg(v string) *SubmitMaterialTaskResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetSuccess(v bool) *SubmitMaterialTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) SetSynchro(v bool) *SubmitMaterialTaskResponseBody {
	s.Synchro = &v
	return s
}

func (s *SubmitMaterialTaskResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitMaterialTaskResponseBodyModule struct {
	// example:
	//
	// 01baf7bcdd5c3a4c8d481cdd57c15837
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitMaterialTaskResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s SubmitMaterialTaskResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SubmitMaterialTaskResponseBodyModule) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitMaterialTaskResponseBodyModule) SetTaskId(v string) *SubmitMaterialTaskResponseBodyModule {
	s.TaskId = &v
	return s
}

func (s *SubmitMaterialTaskResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
