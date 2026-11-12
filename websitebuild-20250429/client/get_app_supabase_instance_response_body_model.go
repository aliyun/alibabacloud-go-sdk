// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppSupabaseInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppSupabaseInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppSupabaseInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppSupabaseInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppSupabaseInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppSupabaseInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppSupabaseInstanceResponseBodyModule) *GetAppSupabaseInstanceResponseBody
	GetModule() *GetAppSupabaseInstanceResponseBodyModule
	SetRequestId(v string) *GetAppSupabaseInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppSupabaseInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppSupabaseInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppSupabaseInstanceResponseBody
	GetSynchro() *bool
}

type GetAppSupabaseInstanceResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
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
	// The dynamic error message, which is used to replace `%s` in the error message of the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The task object.
	Module *GetAppSupabaseInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The backup parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppSupabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppSupabaseInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppSupabaseInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppSupabaseInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppSupabaseInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppSupabaseInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppSupabaseInstanceResponseBody) GetModule() *GetAppSupabaseInstanceResponseBodyModule {
	return s.Module
}

func (s *GetAppSupabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSupabaseInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppSupabaseInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppSupabaseInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppSupabaseInstanceResponseBody) SetAccessDeniedDetail(v string) *GetAppSupabaseInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetAllowRetry(v bool) *GetAppSupabaseInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetAppName(v string) *GetAppSupabaseInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetDynamicCode(v string) *GetAppSupabaseInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetDynamicMessage(v string) *GetAppSupabaseInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetErrorArgs(v []interface{}) *GetAppSupabaseInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetModule(v *GetAppSupabaseInstanceResponseBodyModule) *GetAppSupabaseInstanceResponseBody {
	s.Module = v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetRequestId(v string) *GetAppSupabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetRootErrorCode(v string) *GetAppSupabaseInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetRootErrorMsg(v string) *GetAppSupabaseInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) SetSynchro(v bool) *GetAppSupabaseInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppSupabaseInstanceResponseBodyModule struct {
	// The anonymous key.
	//
	// example:
	//
	// 1111
	AnonKey *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	// The business ID of the application instance.
	//
	// example:
	//
	// WS20250915163734000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// selectdb-cn-2bl4djolb02
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The public URL of the database.
	//
	// example:
	//
	// 111
	DbPublicUrl *string `json:"DbPublicUrl,omitempty" xml:"DbPublicUrl,omitempty"`
	// The instance creation status.
	//
	// example:
	//
	// 1111
	InstanceCreateStatus *string `json:"InstanceCreateStatus,omitempty" xml:"InstanceCreateStatus,omitempty"`
	// The read-only business schema name resolved by Builder trusted binding, which cannot be used as input for tenant switching.
	//
	// example:
	//
	// tenant_00000000-0000-4000-8000-000000000001_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The service key.
	//
	// example:
	//
	// e80f5a7a08514709a2fb
	ServiceKey *string `json:"ServiceKey,omitempty" xml:"ServiceKey,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Supabase instance ID.
	//
	// example:
	//
	// 1111
	SupabaseInstanceId *string `json:"SupabaseInstanceId,omitempty" xml:"SupabaseInstanceId,omitempty"`
	// The public URL of Supabase.
	//
	// example:
	//
	// 111
	SupabasePublicUrl *string `json:"SupabasePublicUrl,omitempty" xml:"SupabasePublicUrl,omitempty"`
	// The read-only tenant mode. The schema is the legacy Builder pool, and db/agenticdb is the new platform. History records can be empty.
	//
	// example:
	//
	// schema
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
}

func (s GetAppSupabaseInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetAnonKey() *string {
	return s.AnonKey
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetDbPublicUrl() *string {
	return s.DbPublicUrl
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetInstanceCreateStatus() *string {
	return s.InstanceCreateStatus
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetServiceKey() *string {
	return s.ServiceKey
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetSupabaseInstanceId() *string {
	return s.SupabaseInstanceId
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetSupabasePublicUrl() *string {
	return s.SupabasePublicUrl
}

func (s *GetAppSupabaseInstanceResponseBodyModule) GetTenantMode() *string {
	return s.TenantMode
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetAnonKey(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.AnonKey = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetBizId(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetDbInstanceId(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.DbInstanceId = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetDbPublicUrl(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.DbPublicUrl = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetInstanceCreateStatus(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.InstanceCreateStatus = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetSchemaName(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.SchemaName = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetServiceKey(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.ServiceKey = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetStatus(v int32) *GetAppSupabaseInstanceResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetSupabaseInstanceId(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.SupabaseInstanceId = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetSupabasePublicUrl(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.SupabasePublicUrl = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) SetTenantMode(v string) *GetAppSupabaseInstanceResponseBodyModule {
	s.TenantMode = &v
	return s
}

func (s *GetAppSupabaseInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
