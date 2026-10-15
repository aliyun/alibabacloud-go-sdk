// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveChannelDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SaveChannelDraftResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SaveChannelDraftResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SaveChannelDraftResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SaveChannelDraftResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SaveChannelDraftResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SaveChannelDraftResponseBody
	GetErrorArgs() []interface{}
	SetModule(v string) *SaveChannelDraftResponseBody
	GetModule() *string
	SetRequestId(v string) *SaveChannelDraftResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SaveChannelDraftResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SaveChannelDraftResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SaveChannelDraftResponseBody
	GetSynchro() *bool
}

type SaveChannelDraftResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether a retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- return parameter.
	//
	// > For example, if **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- that you specified is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
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
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s SaveChannelDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveChannelDraftResponseBody) GoString() string {
	return s.String()
}

func (s *SaveChannelDraftResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SaveChannelDraftResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SaveChannelDraftResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SaveChannelDraftResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SaveChannelDraftResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SaveChannelDraftResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SaveChannelDraftResponseBody) GetModule() *string {
	return s.Module
}

func (s *SaveChannelDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveChannelDraftResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SaveChannelDraftResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SaveChannelDraftResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SaveChannelDraftResponseBody) SetAccessDeniedDetail(v string) *SaveChannelDraftResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetAllowRetry(v bool) *SaveChannelDraftResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetAppName(v string) *SaveChannelDraftResponseBody {
	s.AppName = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetDynamicCode(v string) *SaveChannelDraftResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetDynamicMessage(v string) *SaveChannelDraftResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetErrorArgs(v []interface{}) *SaveChannelDraftResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SaveChannelDraftResponseBody) SetModule(v string) *SaveChannelDraftResponseBody {
	s.Module = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetRequestId(v string) *SaveChannelDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetRootErrorCode(v string) *SaveChannelDraftResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetRootErrorMsg(v string) *SaveChannelDraftResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SaveChannelDraftResponseBody) SetSynchro(v bool) *SaveChannelDraftResponseBody {
	s.Synchro = &v
	return s
}

func (s *SaveChannelDraftResponseBody) Validate() error {
	return dara.Validate(s)
}
