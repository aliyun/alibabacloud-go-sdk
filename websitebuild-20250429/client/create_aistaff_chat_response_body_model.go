// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAIStaffChatResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAIStaffChatResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAIStaffChatResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAIStaffChatResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAIStaffChatResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAIStaffChatResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *CreateAIStaffChatResponseBody
	GetModule() *bool
	SetRequestId(v string) *CreateAIStaffChatResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAIStaffChatResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAIStaffChatResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAIStaffChatResponseBody
	GetSynchro() *bool
}

type CreateAIStaffChatResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
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

func (s CreateAIStaffChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIStaffChatResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAIStaffChatResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAIStaffChatResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAIStaffChatResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAIStaffChatResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAIStaffChatResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAIStaffChatResponseBody) GetModule() *bool {
	return s.Module
}

func (s *CreateAIStaffChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIStaffChatResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAIStaffChatResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAIStaffChatResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAIStaffChatResponseBody) SetAccessDeniedDetail(v string) *CreateAIStaffChatResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetAllowRetry(v bool) *CreateAIStaffChatResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetAppName(v string) *CreateAIStaffChatResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetDynamicCode(v string) *CreateAIStaffChatResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetDynamicMessage(v string) *CreateAIStaffChatResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetErrorArgs(v []interface{}) *CreateAIStaffChatResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetModule(v bool) *CreateAIStaffChatResponseBody {
	s.Module = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetRequestId(v string) *CreateAIStaffChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetRootErrorCode(v string) *CreateAIStaffChatResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetRootErrorMsg(v string) *CreateAIStaffChatResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) SetSynchro(v bool) *CreateAIStaffChatResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAIStaffChatResponseBody) Validate() error {
	return dara.Validate(s)
}
