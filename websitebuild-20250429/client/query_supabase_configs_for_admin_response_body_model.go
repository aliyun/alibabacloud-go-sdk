// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseConfigsForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QuerySupabaseConfigsForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QuerySupabaseConfigsForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QuerySupabaseConfigsForAdminResponseBodyModule) *QuerySupabaseConfigsForAdminResponseBody
	GetModule() *QuerySupabaseConfigsForAdminResponseBodyModule
	SetRequestId(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QuerySupabaseConfigsForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QuerySupabaseConfigsForAdminResponseBody
	GetSynchro() *bool
}

type QuerySupabaseConfigsForAdminResponseBody struct {
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
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                   `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QuerySupabaseConfigsForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s QuerySupabaseConfigsForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseConfigsForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetModule() *QuerySupabaseConfigsForAdminResponseBodyModule {
	return s.Module
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QuerySupabaseConfigsForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetAccessDeniedDetail(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetAllowRetry(v bool) *QuerySupabaseConfigsForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetAppName(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetDynamicCode(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetDynamicMessage(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetErrorArgs(v []interface{}) *QuerySupabaseConfigsForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetModule(v *QuerySupabaseConfigsForAdminResponseBodyModule) *QuerySupabaseConfigsForAdminResponseBody {
	s.Module = v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetRequestId(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetRootErrorCode(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetRootErrorMsg(v string) *QuerySupabaseConfigsForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) SetSynchro(v bool) *QuerySupabaseConfigsForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySupabaseConfigsForAdminResponseBodyModule struct {
	Configs map[string]interface{} `json:"Configs,omitempty" xml:"Configs,omitempty"`
}

func (s QuerySupabaseConfigsForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseConfigsForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QuerySupabaseConfigsForAdminResponseBodyModule) GetConfigs() map[string]interface{} {
	return s.Configs
}

func (s *QuerySupabaseConfigsForAdminResponseBodyModule) SetConfigs(v map[string]interface{}) *QuerySupabaseConfigsForAdminResponseBodyModule {
	s.Configs = v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
