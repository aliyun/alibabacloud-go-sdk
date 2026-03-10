// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseAuthConfigsForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QuerySupabaseAuthConfigsForAdminResponseBodyModule) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetModule() *QuerySupabaseAuthConfigsForAdminResponseBodyModule
	SetRequestId(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QuerySupabaseAuthConfigsForAdminResponseBody
	GetSynchro() *bool
}

type QuerySupabaseAuthConfigsForAdminResponseBody struct {
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
	DynamicMessage *string                                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                       `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QuerySupabaseAuthConfigsForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s QuerySupabaseAuthConfigsForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseAuthConfigsForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetModule() *QuerySupabaseAuthConfigsForAdminResponseBodyModule {
	return s.Module
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetAccessDeniedDetail(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetAllowRetry(v bool) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetAppName(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetDynamicCode(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetDynamicMessage(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetErrorArgs(v []interface{}) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetModule(v *QuerySupabaseAuthConfigsForAdminResponseBodyModule) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.Module = v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetRequestId(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetRootErrorCode(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetRootErrorMsg(v string) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) SetSynchro(v bool) *QuerySupabaseAuthConfigsForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySupabaseAuthConfigsForAdminResponseBodyModule struct {
	Configs map[string]interface{} `json:"Configs,omitempty" xml:"Configs,omitempty"`
}

func (s QuerySupabaseAuthConfigsForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseAuthConfigsForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBodyModule) GetConfigs() map[string]interface{} {
	return s.Configs
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBodyModule) SetConfigs(v map[string]interface{}) *QuerySupabaseAuthConfigsForAdminResponseBodyModule {
	s.Configs = v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
