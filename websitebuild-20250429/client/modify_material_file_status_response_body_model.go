// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyMaterialFileStatusResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ModifyMaterialFileStatusResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ModifyMaterialFileStatusResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ModifyMaterialFileStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyMaterialFileStatusResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ModifyMaterialFileStatusResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *ModifyMaterialFileStatusResponseBody
	GetModule() *bool
	SetRequestId(v string) *ModifyMaterialFileStatusResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ModifyMaterialFileStatusResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ModifyMaterialFileStatusResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ModifyMaterialFileStatusResponseBody
	GetSynchro() *bool
}

type ModifyMaterialFileStatusResponseBody struct {
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
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ModifyMaterialFileStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyMaterialFileStatusResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ModifyMaterialFileStatusResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ModifyMaterialFileStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyMaterialFileStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyMaterialFileStatusResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ModifyMaterialFileStatusResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ModifyMaterialFileStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaterialFileStatusResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ModifyMaterialFileStatusResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ModifyMaterialFileStatusResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ModifyMaterialFileStatusResponseBody) SetAccessDeniedDetail(v string) *ModifyMaterialFileStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetAllowRetry(v bool) *ModifyMaterialFileStatusResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetAppName(v string) *ModifyMaterialFileStatusResponseBody {
	s.AppName = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetDynamicCode(v string) *ModifyMaterialFileStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetDynamicMessage(v string) *ModifyMaterialFileStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetErrorArgs(v []interface{}) *ModifyMaterialFileStatusResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetModule(v bool) *ModifyMaterialFileStatusResponseBody {
	s.Module = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetRequestId(v string) *ModifyMaterialFileStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetRootErrorCode(v string) *ModifyMaterialFileStatusResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetRootErrorMsg(v string) *ModifyMaterialFileStatusResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) SetSynchro(v bool) *ModifyMaterialFileStatusResponseBody {
	s.Synchro = &v
	return s
}

func (s *ModifyMaterialFileStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
