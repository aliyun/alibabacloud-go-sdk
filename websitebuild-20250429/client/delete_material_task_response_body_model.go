// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteMaterialTaskResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteMaterialTaskResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteMaterialTaskResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteMaterialTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteMaterialTaskResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteMaterialTaskResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteMaterialTaskResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteMaterialTaskResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteMaterialTaskResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteMaterialTaskResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteMaterialTaskResponseBody
	GetSynchro() *bool
}

type DeleteMaterialTaskResponseBody struct {
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

func (s DeleteMaterialTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteMaterialTaskResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteMaterialTaskResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteMaterialTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteMaterialTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteMaterialTaskResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteMaterialTaskResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteMaterialTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaterialTaskResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteMaterialTaskResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteMaterialTaskResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteMaterialTaskResponseBody) SetAccessDeniedDetail(v string) *DeleteMaterialTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetAllowRetry(v bool) *DeleteMaterialTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetAppName(v string) *DeleteMaterialTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetDynamicCode(v string) *DeleteMaterialTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetDynamicMessage(v string) *DeleteMaterialTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetErrorArgs(v []interface{}) *DeleteMaterialTaskResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetModule(v bool) *DeleteMaterialTaskResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetRequestId(v string) *DeleteMaterialTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetRootErrorCode(v string) *DeleteMaterialTaskResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetRootErrorMsg(v string) *DeleteMaterialTaskResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) SetSynchro(v bool) *DeleteMaterialTaskResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteMaterialTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
