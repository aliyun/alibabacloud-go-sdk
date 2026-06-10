// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSitemapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSitemapResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSitemapResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSitemapResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSitemapResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSitemapResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSitemapResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *GetAppSitemapResponseBody
	GetModule() *string
	SetRequestId(v string) *GetAppSitemapResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSitemapResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSitemapResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSitemapResponseBody
	GetSynchro() *bool
}

type GetAppSitemapResponseBody struct {
	// Detailed reason for access denial.
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
	// App name.
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
	// Dynamic error message used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
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
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
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

func (s GetAppSitemapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSitemapResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSitemapResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSitemapResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSitemapResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSitemapResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSitemapResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSitemapResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSitemapResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetAppSitemapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSitemapResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSitemapResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSitemapResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSitemapResponseBody) SetAccessDeniedDetail(v string) *GetAppSitemapResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetAllowRetry(v bool) *GetAppSitemapResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetAppName(v string) *GetAppSitemapResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetDynamicCode(v string) *GetAppSitemapResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetDynamicMessage(v string) *GetAppSitemapResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetErrorArgs(v []interface{}) *GetAppSitemapResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSitemapResponseBody) SetModule(v string) *GetAppSitemapResponseBody {
	s.Module = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetRequestId(v string) *GetAppSitemapResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetRootErrorCode(v string) *GetAppSitemapResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetRootErrorMsg(v string) *GetAppSitemapResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSitemapResponseBody) SetSynchro(v bool) *GetAppSitemapResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSitemapResponseBody) Validate() error {
	return dara.Validate(s)
}
