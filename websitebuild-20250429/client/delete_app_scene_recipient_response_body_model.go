// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSceneRecipientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppSceneRecipientResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppSceneRecipientResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppSceneRecipientResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppSceneRecipientResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppSceneRecipientResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppSceneRecipientResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteAppSceneRecipientResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteAppSceneRecipientResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppSceneRecipientResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppSceneRecipientResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppSceneRecipientResponseBody
	GetSynchro() *bool
}

type DeleteAppSceneRecipientResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. Currently not in use. Ignore this parameter.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Indicates whether the move was successful.
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
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
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

func (s DeleteAppSceneRecipientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSceneRecipientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppSceneRecipientResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppSceneRecipientResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppSceneRecipientResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppSceneRecipientResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppSceneRecipientResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppSceneRecipientResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppSceneRecipientResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteAppSceneRecipientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppSceneRecipientResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppSceneRecipientResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppSceneRecipientResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppSceneRecipientResponseBody) SetAccessDeniedDetail(v string) *DeleteAppSceneRecipientResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetAllowRetry(v bool) *DeleteAppSceneRecipientResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetAppName(v string) *DeleteAppSceneRecipientResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetDynamicCode(v string) *DeleteAppSceneRecipientResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetDynamicMessage(v string) *DeleteAppSceneRecipientResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetErrorArgs(v []interface{}) *DeleteAppSceneRecipientResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetModule(v bool) *DeleteAppSceneRecipientResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetRequestId(v string) *DeleteAppSceneRecipientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetRootErrorCode(v string) *DeleteAppSceneRecipientResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetRootErrorMsg(v string) *DeleteAppSceneRecipientResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) SetSynchro(v bool) *DeleteAppSceneRecipientResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppSceneRecipientResponseBody) Validate() error {
	return dara.Validate(s)
}
