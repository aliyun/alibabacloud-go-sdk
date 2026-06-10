// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSupabaseSecretsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSupabaseSecretsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSupabaseSecretsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSupabaseSecretsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSupabaseSecretsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSupabaseSecretsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppSupabaseSecretsResponseBodyModule) *GetAppSupabaseSecretsResponseBody
	GetModule() *GetAppSupabaseSecretsResponseBodyModule
	SetRequestId(v string) *GetAppSupabaseSecretsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSupabaseSecretsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSupabaseSecretsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSupabaseSecretsResponseBody
	GetSynchro() *bool
}

type GetAppSupabaseSecretsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed. Valid values:
	//
	// - false: Retry is not allowed.
	//
	// - true: Retry is allowed.
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
	// Dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > For example, if **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Returned object.
	Module *GetAppSupabaseSecretsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// ID of the request
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
	// Abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the operation is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppSupabaseSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseSecretsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSupabaseSecretsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSupabaseSecretsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSupabaseSecretsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSupabaseSecretsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSupabaseSecretsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSupabaseSecretsResponseBody) GetModule() *GetAppSupabaseSecretsResponseBodyModule {
	return s.Module
}

func (s *GetAppSupabaseSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSupabaseSecretsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSupabaseSecretsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSupabaseSecretsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSupabaseSecretsResponseBody) SetAccessDeniedDetail(v string) *GetAppSupabaseSecretsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetAllowRetry(v bool) *GetAppSupabaseSecretsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetAppName(v string) *GetAppSupabaseSecretsResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetDynamicCode(v string) *GetAppSupabaseSecretsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetDynamicMessage(v string) *GetAppSupabaseSecretsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetErrorArgs(v []interface{}) *GetAppSupabaseSecretsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetModule(v *GetAppSupabaseSecretsResponseBodyModule) *GetAppSupabaseSecretsResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetRequestId(v string) *GetAppSupabaseSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetRootErrorCode(v string) *GetAppSupabaseSecretsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetRootErrorMsg(v string) *GetAppSupabaseSecretsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) SetSynchro(v bool) *GetAppSupabaseSecretsResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppSupabaseSecretsResponseBodyModule struct {
	// Instance ID.
	Secrets []*GetAppSupabaseSecretsResponseBodyModuleSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s GetAppSupabaseSecretsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseSecretsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseSecretsResponseBodyModule) GetSecrets() []*GetAppSupabaseSecretsResponseBodyModuleSecrets {
	return s.Secrets
}

func (s *GetAppSupabaseSecretsResponseBodyModule) SetSecrets(v []*GetAppSupabaseSecretsResponseBodyModuleSecrets) *GetAppSupabaseSecretsResponseBodyModule {
	s.Secrets = v
	return s
}

func (s *GetAppSupabaseSecretsResponseBodyModule) Validate() error {
	if s.Secrets != nil {
		for _, item := range s.Secrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppSupabaseSecretsResponseBodyModuleSecrets struct {
	// Tag key
	//
	// example:
	//
	// curl GCEA6fET.popscan.xaliyun.com
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Application name
	//
	// example:
	//
	// LOGO3.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The credential type. Valid values:
	//
	// - Generic: generic secret.
	//
	// - Rds: RDS credential.
	//
	// - Redis: Redis/Tair credential.
	//
	// - RAMCredentials: RAM credential.
	//
	// - ECS: ECS credential.
	//
	// - PolarDB: PolarDB credential.
	//
	// example:
	//
	// Opaque
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// Record value
	//
	// example:
	//
	// 58.16.60.28
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAppSupabaseSecretsResponseBodyModuleSecrets) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseSecretsResponseBodyModuleSecrets) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) GetKey() *string {
	return s.Key
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) GetName() *string {
	return s.Name
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) GetSecretType() *string {
	return s.SecretType
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) GetValue() *string {
	return s.Value
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) SetKey(v string) *GetAppSupabaseSecretsResponseBodyModuleSecrets {
	s.Key = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) SetName(v string) *GetAppSupabaseSecretsResponseBodyModuleSecrets {
	s.Name = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) SetSecretType(v string) *GetAppSupabaseSecretsResponseBodyModuleSecrets {
	s.SecretType = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) SetValue(v string) *GetAppSupabaseSecretsResponseBodyModuleSecrets {
	s.Value = &v
	return s
}

func (s *GetAppSupabaseSecretsResponseBodyModuleSecrets) Validate() error {
	return dara.Validate(s)
}
