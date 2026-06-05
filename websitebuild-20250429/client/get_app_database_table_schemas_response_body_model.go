// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDatabaseTableSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppDatabaseTableSchemasResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppDatabaseTableSchemasResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppDatabaseTableSchemasResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppDatabaseTableSchemasResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppDatabaseTableSchemasResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppDatabaseTableSchemasResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *GetAppDatabaseTableSchemasResponseBody
	GetModule() *string
	SetRequestId(v string) *GetAppDatabaseTableSchemasResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppDatabaseTableSchemasResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppDatabaseTableSchemasResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppDatabaseTableSchemasResponseBody
	GetSynchro() *bool
}

type GetAppDatabaseTableSchemasResponseBody struct {
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
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
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

func (s GetAppDatabaseTableSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppDatabaseTableSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppDatabaseTableSchemasResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetAccessDeniedDetail(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetAllowRetry(v bool) *GetAppDatabaseTableSchemasResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetAppName(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetDynamicCode(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetDynamicMessage(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetErrorArgs(v []interface{}) *GetAppDatabaseTableSchemasResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetModule(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.Module = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetRequestId(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetRootErrorCode(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetRootErrorMsg(v string) *GetAppDatabaseTableSchemasResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) SetSynchro(v bool) *GetAppDatabaseTableSchemasResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}
