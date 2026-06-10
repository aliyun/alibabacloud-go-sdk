// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryMaterialTaskDetailResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryMaterialTaskDetailResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryMaterialTaskDetailResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryMaterialTaskDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryMaterialTaskDetailResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryMaterialTaskDetailResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *QueryMaterialTaskDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryMaterialTaskDetailResponseBody
	GetErrorMsg() *string
	SetModule(v *AppMaterialTask) *QueryMaterialTaskDetailResponseBody
	GetModule() *AppMaterialTask
	SetRequestId(v string) *QueryMaterialTaskDetailResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryMaterialTaskDetailResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryMaterialTaskDetailResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *QueryMaterialTaskDetailResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *QueryMaterialTaskDetailResponseBody
	GetSynchro() *bool
}

type QueryMaterialTaskDetailResponseBody struct {
	// access denied details
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
	// AppName. It can contain digits, letters, and hyphens (-). It must start with a letter, cannot end with a hyphen (-), and cannot exceed 36 characters in length.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic message
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// error code. The ErrorCode field is returned only when the Request fails. If the Request succeeds, the ErrorCode field is not returned. For more information, see the error code List in this topic.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// error message
	//
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Response data
	//
	// example:
	//
	// {\\"Success\\": True}
	Module *AppMaterialTask `json:"Module,omitempty" xml:"Module,omitempty"`
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
	// Indicates whether the Request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryMaterialTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialTaskDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryMaterialTaskDetailResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryMaterialTaskDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryMaterialTaskDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryMaterialTaskDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryMaterialTaskDetailResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryMaterialTaskDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMaterialTaskDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryMaterialTaskDetailResponseBody) GetModule() *AppMaterialTask {
	return s.Module
}

func (s *QueryMaterialTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMaterialTaskDetailResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryMaterialTaskDetailResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryMaterialTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMaterialTaskDetailResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryMaterialTaskDetailResponseBody) SetAccessDeniedDetail(v string) *QueryMaterialTaskDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetAllowRetry(v bool) *QueryMaterialTaskDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetAppName(v string) *QueryMaterialTaskDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetDynamicCode(v string) *QueryMaterialTaskDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetDynamicMessage(v string) *QueryMaterialTaskDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetErrorArgs(v []interface{}) *QueryMaterialTaskDetailResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetErrorCode(v string) *QueryMaterialTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetErrorMsg(v string) *QueryMaterialTaskDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetModule(v *AppMaterialTask) *QueryMaterialTaskDetailResponseBody {
	s.Module = v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetRequestId(v string) *QueryMaterialTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetRootErrorCode(v string) *QueryMaterialTaskDetailResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetRootErrorMsg(v string) *QueryMaterialTaskDetailResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetSuccess(v bool) *QueryMaterialTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) SetSynchro(v bool) *QueryMaterialTaskDetailResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryMaterialTaskDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}
