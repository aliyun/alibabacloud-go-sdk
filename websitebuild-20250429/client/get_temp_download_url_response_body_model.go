// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetTempDownloadUrlResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetTempDownloadUrlResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetTempDownloadUrlResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetTempDownloadUrlResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetTempDownloadUrlResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetTempDownloadUrlResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *GetTempDownloadUrlResponseBody
	GetModule() *string
	SetRequestId(v string) *GetTempDownloadUrlResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetTempDownloadUrlResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetTempDownloadUrlResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetTempDownloadUrlResponseBody
	GetSynchro() *bool
}

type GetTempDownloadUrlResponseBody struct {
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

func (s GetTempDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTempDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTempDownloadUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetTempDownloadUrlResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetTempDownloadUrlResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetTempDownloadUrlResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetTempDownloadUrlResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetTempDownloadUrlResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetTempDownloadUrlResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetTempDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTempDownloadUrlResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetTempDownloadUrlResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetTempDownloadUrlResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetTempDownloadUrlResponseBody) SetAccessDeniedDetail(v string) *GetTempDownloadUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetAllowRetry(v bool) *GetTempDownloadUrlResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetAppName(v string) *GetTempDownloadUrlResponseBody {
	s.AppName = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetDynamicCode(v string) *GetTempDownloadUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetDynamicMessage(v string) *GetTempDownloadUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetErrorArgs(v []interface{}) *GetTempDownloadUrlResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetModule(v string) *GetTempDownloadUrlResponseBody {
	s.Module = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetRequestId(v string) *GetTempDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetRootErrorCode(v string) *GetTempDownloadUrlResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetRootErrorMsg(v string) *GetTempDownloadUrlResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) SetSynchro(v bool) *GetTempDownloadUrlResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetTempDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
