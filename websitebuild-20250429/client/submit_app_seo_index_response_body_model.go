// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAppSeoIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitAppSeoIndexResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SubmitAppSeoIndexResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SubmitAppSeoIndexResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SubmitAppSeoIndexResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SubmitAppSeoIndexResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SubmitAppSeoIndexResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *SubmitAppSeoIndexResponseBody
	GetModule() *bool
	SetRequestId(v string) *SubmitAppSeoIndexResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SubmitAppSeoIndexResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SubmitAppSeoIndexResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SubmitAppSeoIndexResponseBody
	GetSynchro() *bool
}

type SubmitAppSeoIndexResponseBody struct {
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
	// spring-cloud-b
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

func (s SubmitAppSeoIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAppSeoIndexResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAppSeoIndexResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitAppSeoIndexResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SubmitAppSeoIndexResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SubmitAppSeoIndexResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SubmitAppSeoIndexResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SubmitAppSeoIndexResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SubmitAppSeoIndexResponseBody) GetModule() *bool {
	return s.Module
}

func (s *SubmitAppSeoIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAppSeoIndexResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SubmitAppSeoIndexResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SubmitAppSeoIndexResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SubmitAppSeoIndexResponseBody) SetAccessDeniedDetail(v string) *SubmitAppSeoIndexResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetAllowRetry(v bool) *SubmitAppSeoIndexResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetAppName(v string) *SubmitAppSeoIndexResponseBody {
	s.AppName = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetDynamicCode(v string) *SubmitAppSeoIndexResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetDynamicMessage(v string) *SubmitAppSeoIndexResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetErrorArgs(v []interface{}) *SubmitAppSeoIndexResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetModule(v bool) *SubmitAppSeoIndexResponseBody {
	s.Module = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetRequestId(v string) *SubmitAppSeoIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetRootErrorCode(v string) *SubmitAppSeoIndexResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetRootErrorMsg(v string) *SubmitAppSeoIndexResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) SetSynchro(v bool) *SubmitAppSeoIndexResponseBody {
	s.Synchro = &v
	return s
}

func (s *SubmitAppSeoIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
