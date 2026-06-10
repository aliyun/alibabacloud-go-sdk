// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppSiteValidationFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UploadAppSiteValidationFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UploadAppSiteValidationFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UploadAppSiteValidationFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UploadAppSiteValidationFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UploadAppSiteValidationFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UploadAppSiteValidationFileResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *UploadAppSiteValidationFileResponseBody
	GetModule() *bool
	SetRequestId(v string) *UploadAppSiteValidationFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UploadAppSiteValidationFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UploadAppSiteValidationFileResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UploadAppSiteValidationFileResponseBody
	GetSynchro() *bool
}

type UploadAppSiteValidationFileResponseBody struct {
	// access denied details
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// application name. Query the application with this name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic code, currently unused. Ignore it.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message.
	//
	// example:
	//
	// SYSTEM.ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// indicates whether the deletion succeeded
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
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s UploadAppSiteValidationFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAppSiteValidationFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAppSiteValidationFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UploadAppSiteValidationFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UploadAppSiteValidationFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UploadAppSiteValidationFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UploadAppSiteValidationFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UploadAppSiteValidationFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UploadAppSiteValidationFileResponseBody) GetModule() *bool {
	return s.Module
}

func (s *UploadAppSiteValidationFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAppSiteValidationFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UploadAppSiteValidationFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UploadAppSiteValidationFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UploadAppSiteValidationFileResponseBody) SetAccessDeniedDetail(v string) *UploadAppSiteValidationFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetAllowRetry(v bool) *UploadAppSiteValidationFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetAppName(v string) *UploadAppSiteValidationFileResponseBody {
	s.AppName = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetDynamicCode(v string) *UploadAppSiteValidationFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetDynamicMessage(v string) *UploadAppSiteValidationFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetErrorArgs(v []interface{}) *UploadAppSiteValidationFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetModule(v bool) *UploadAppSiteValidationFileResponseBody {
	s.Module = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetRequestId(v string) *UploadAppSiteValidationFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetRootErrorCode(v string) *UploadAppSiteValidationFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetRootErrorMsg(v string) *UploadAppSiteValidationFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) SetSynchro(v bool) *UploadAppSiteValidationFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *UploadAppSiteValidationFileResponseBody) Validate() error {
	return dara.Validate(s)
}
