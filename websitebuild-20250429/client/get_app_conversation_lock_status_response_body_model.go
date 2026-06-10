// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationLockStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppConversationLockStatusResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppConversationLockStatusResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppConversationLockStatusResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppConversationLockStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppConversationLockStatusResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppConversationLockStatusResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *GetAppConversationLockStatusResponseBody
	GetModule() *bool
	SetRequestId(v string) *GetAppConversationLockStatusResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppConversationLockStatusResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppConversationLockStatusResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppConversationLockStatusResponseBody
	GetSynchro() *bool
}

type GetAppConversationLockStatusResponseBody struct {
	// permission denied information
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
	// AppName.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic code, currently unused; please ignore
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// response data
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
	// is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppConversationLockStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationLockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConversationLockStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppConversationLockStatusResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppConversationLockStatusResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppConversationLockStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppConversationLockStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppConversationLockStatusResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppConversationLockStatusResponseBody) GetModule() *bool {
	return s.Module
}

func (s *GetAppConversationLockStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppConversationLockStatusResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppConversationLockStatusResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppConversationLockStatusResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppConversationLockStatusResponseBody) SetAccessDeniedDetail(v string) *GetAppConversationLockStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetAllowRetry(v bool) *GetAppConversationLockStatusResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetAppName(v string) *GetAppConversationLockStatusResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetDynamicCode(v string) *GetAppConversationLockStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetDynamicMessage(v string) *GetAppConversationLockStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetErrorArgs(v []interface{}) *GetAppConversationLockStatusResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetModule(v bool) *GetAppConversationLockStatusResponseBody {
	s.Module = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetRequestId(v string) *GetAppConversationLockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetRootErrorCode(v string) *GetAppConversationLockStatusResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetRootErrorMsg(v string) *GetAppConversationLockStatusResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) SetSynchro(v bool) *GetAppConversationLockStatusResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppConversationLockStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
