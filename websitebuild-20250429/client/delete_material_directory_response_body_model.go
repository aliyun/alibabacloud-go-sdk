// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteMaterialDirectoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteMaterialDirectoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteMaterialDirectoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteMaterialDirectoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteMaterialDirectoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteMaterialDirectoryResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *DeleteMaterialDirectoryResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteMaterialDirectoryResponseBody
	GetErrorMsg() *string
	SetModule(v bool) *DeleteMaterialDirectoryResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteMaterialDirectoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteMaterialDirectoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteMaterialDirectoryResponseBody
	GetRootErrorMsg() *string
	SetSuccess(v bool) *DeleteMaterialDirectoryResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *DeleteMaterialDirectoryResponseBody
	GetSynchro() *bool
}

type DeleteMaterialDirectoryResponseBody struct {
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
	// or
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

func (s DeleteMaterialDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialDirectoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteMaterialDirectoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteMaterialDirectoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteMaterialDirectoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteMaterialDirectoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteMaterialDirectoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteMaterialDirectoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMaterialDirectoryResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteMaterialDirectoryResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteMaterialDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaterialDirectoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteMaterialDirectoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteMaterialDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMaterialDirectoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteMaterialDirectoryResponseBody) SetAccessDeniedDetail(v string) *DeleteMaterialDirectoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetAllowRetry(v bool) *DeleteMaterialDirectoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetAppName(v string) *DeleteMaterialDirectoryResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetDynamicCode(v string) *DeleteMaterialDirectoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetDynamicMessage(v string) *DeleteMaterialDirectoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetErrorArgs(v []interface{}) *DeleteMaterialDirectoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetErrorCode(v string) *DeleteMaterialDirectoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetErrorMsg(v string) *DeleteMaterialDirectoryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetModule(v bool) *DeleteMaterialDirectoryResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetRequestId(v string) *DeleteMaterialDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetRootErrorCode(v string) *DeleteMaterialDirectoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetRootErrorMsg(v string) *DeleteMaterialDirectoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetSuccess(v bool) *DeleteMaterialDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) SetSynchro(v bool) *DeleteMaterialDirectoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteMaterialDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
