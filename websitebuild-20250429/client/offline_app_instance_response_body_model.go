// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OfflineAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *OfflineAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *OfflineAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *OfflineAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *OfflineAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *OfflineAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v map[string]interface{}) *OfflineAppInstanceResponseBody
	GetModule() map[string]interface{}
	SetRequestId(v string) *OfflineAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *OfflineAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *OfflineAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *OfflineAppInstanceResponseBody
	GetSynchro() *bool
}

type OfflineAppInstanceResponseBody struct {
	// The deprecated parameter. You can ignore this parameter.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// DocdbSortingCode
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the shift was successful.
	//
	// example:
	//
	// {\\"TotalPageNum\\": 1, \\"ResultLimit\\": False, \\"CurrentPageNum\\": 0, \\"PageSize\\": 0, \\"TotalItemNum\\": 0}
	Module map[string]interface{} `json:"Module,omitempty" xml:"Module,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s OfflineAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OfflineAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *OfflineAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *OfflineAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *OfflineAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *OfflineAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *OfflineAppInstanceResponseBody) GetModule() map[string]interface{} {
	return s.Module
}

func (s *OfflineAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *OfflineAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *OfflineAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *OfflineAppInstanceResponseBody) SetAccessDeniedDetail(v string) *OfflineAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetAllowRetry(v bool) *OfflineAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetAppName(v string) *OfflineAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetDynamicCode(v string) *OfflineAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetDynamicMessage(v string) *OfflineAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetErrorArgs(v []interface{}) *OfflineAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetModule(v map[string]interface{}) *OfflineAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetRequestId(v string) *OfflineAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetRootErrorCode(v string) *OfflineAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetRootErrorMsg(v string) *OfflineAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) SetSynchro(v bool) *OfflineAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *OfflineAppInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
