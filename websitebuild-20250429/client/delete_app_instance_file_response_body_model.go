// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteAppInstanceFileResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *DeleteAppInstanceFileResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DeleteAppInstanceFileResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DeleteAppInstanceFileResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteAppInstanceFileResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DeleteAppInstanceFileResponseBody
	GetErrorArgs() []interface{}
	SetModule(v bool) *DeleteAppInstanceFileResponseBody
	GetModule() *bool
	SetRequestId(v string) *DeleteAppInstanceFileResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *DeleteAppInstanceFileResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *DeleteAppInstanceFileResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *DeleteAppInstanceFileResponseBody
	GetSynchro() *bool
}

type DeleteAppInstanceFileResponseBody struct {
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

func (s DeleteAppInstanceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceFileResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteAppInstanceFileResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DeleteAppInstanceFileResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppInstanceFileResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteAppInstanceFileResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteAppInstanceFileResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DeleteAppInstanceFileResponseBody) GetModule() *bool {
	return s.Module
}

func (s *DeleteAppInstanceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppInstanceFileResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *DeleteAppInstanceFileResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *DeleteAppInstanceFileResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DeleteAppInstanceFileResponseBody) SetAccessDeniedDetail(v string) *DeleteAppInstanceFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetAllowRetry(v bool) *DeleteAppInstanceFileResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetAppName(v string) *DeleteAppInstanceFileResponseBody {
	s.AppName = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetDynamicCode(v string) *DeleteAppInstanceFileResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetDynamicMessage(v string) *DeleteAppInstanceFileResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetErrorArgs(v []interface{}) *DeleteAppInstanceFileResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetModule(v bool) *DeleteAppInstanceFileResponseBody {
	s.Module = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetRequestId(v string) *DeleteAppInstanceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetRootErrorCode(v string) *DeleteAppInstanceFileResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetRootErrorMsg(v string) *DeleteAppInstanceFileResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) SetSynchro(v bool) *DeleteAppInstanceFileResponseBody {
	s.Synchro = &v
	return s
}

func (s *DeleteAppInstanceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
