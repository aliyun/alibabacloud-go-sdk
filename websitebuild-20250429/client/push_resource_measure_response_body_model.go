// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceMeasureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PushResourceMeasureResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *PushResourceMeasureResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PushResourceMeasureResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PushResourceMeasureResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PushResourceMeasureResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *PushResourceMeasureResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *PushResourceMeasureResponseBody
	GetModule() *bool
	SetRequestId(v string) *PushResourceMeasureResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *PushResourceMeasureResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *PushResourceMeasureResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *PushResourceMeasureResponseBody
	GetSynchro() *bool
}

type PushResourceMeasureResponseBody struct {
	// Details of access denied
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name. Query the application with this name.
	//
	// example:
	//
	// spring-cloud-b
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
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether processing is synchronous
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s PushResourceMeasureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushResourceMeasureResponseBody) GoString() string {
	return s.String()
}

func (s *PushResourceMeasureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PushResourceMeasureResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PushResourceMeasureResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PushResourceMeasureResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PushResourceMeasureResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PushResourceMeasureResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *PushResourceMeasureResponseBody) GetModule() *bool {
	return s.Module
}

func (s *PushResourceMeasureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushResourceMeasureResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *PushResourceMeasureResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *PushResourceMeasureResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *PushResourceMeasureResponseBody) SetAccessDeniedDetail(v string) *PushResourceMeasureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetAllowRetry(v bool) *PushResourceMeasureResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetAppName(v string) *PushResourceMeasureResponseBody {
	s.AppName = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetDynamicCode(v string) *PushResourceMeasureResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetDynamicMessage(v string) *PushResourceMeasureResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetErrorArgs(v []interface{}) *PushResourceMeasureResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *PushResourceMeasureResponseBody) SetModule(v bool) *PushResourceMeasureResponseBody {
	s.Module = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetRequestId(v string) *PushResourceMeasureResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetRootErrorCode(v string) *PushResourceMeasureResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetRootErrorMsg(v string) *PushResourceMeasureResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *PushResourceMeasureResponseBody) SetSynchro(v bool) *PushResourceMeasureResponseBody {
	s.Synchro = &v
	return s
}

func (s *PushResourceMeasureResponseBody) Validate() error {
	return dara.Validate(s)
}
