// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSupabaseAuthConfigResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSupabaseAuthConfigResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSupabaseAuthConfigResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSupabaseAuthConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSupabaseAuthConfigResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSupabaseAuthConfigResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppSupabaseAuthConfigResponseBodyModule) *GetAppSupabaseAuthConfigResponseBody
	GetModule() *GetAppSupabaseAuthConfigResponseBodyModule
	SetRequestId(v string) *GetAppSupabaseAuthConfigResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSupabaseAuthConfigResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSupabaseAuthConfigResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSupabaseAuthConfigResponseBody
	GetSynchro() *bool
}

type GetAppSupabaseAuthConfigResponseBody struct {
	// Detailed reason for access denial.
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
	// App Name.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Error arguments returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Task object
	Module *GetAppSupabaseAuthConfigResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// SYSTEM.EROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Error message
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

func (s GetAppSupabaseAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetModule() *GetAppSupabaseAuthConfigResponseBodyModule {
	return s.Module
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSupabaseAuthConfigResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetAccessDeniedDetail(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetAllowRetry(v bool) *GetAppSupabaseAuthConfigResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetAppName(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetDynamicCode(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetDynamicMessage(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetErrorArgs(v []interface{}) *GetAppSupabaseAuthConfigResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetModule(v *GetAppSupabaseAuthConfigResponseBodyModule) *GetAppSupabaseAuthConfigResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetRequestId(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetRootErrorCode(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetRootErrorMsg(v string) *GetAppSupabaseAuthConfigResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) SetSynchro(v bool) *GetAppSupabaseAuthConfigResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppSupabaseAuthConfigResponseBodyModule struct {
	// Configuration value. Valid values:
	//
	// - cc_rule: CC rule.
	//
	// - ddos_dispatch: DDoS filter interaction scheduling.
	//
	// - edge_safe: Edge application security.
	//
	// - blocked_regions: Geo-blocking.
	//
	// - http_acl_policy: Precise ACL rule.
	//
	// - bot_manager: Bot traffic management.
	//
	// - ip_reputation: IP reputation investigation.
	Configs map[string]interface{} `json:"Configs,omitempty" xml:"Configs,omitempty"`
}

func (s GetAppSupabaseAuthConfigResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseAuthConfigResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseAuthConfigResponseBodyModule) GetConfigs() map[string]interface{} {
	return s.Configs
}

func (s *GetAppSupabaseAuthConfigResponseBodyModule) SetConfigs(v map[string]interface{}) *GetAppSupabaseAuthConfigResponseBodyModule {
	s.Configs = v
	return s
}

func (s *GetAppSupabaseAuthConfigResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
