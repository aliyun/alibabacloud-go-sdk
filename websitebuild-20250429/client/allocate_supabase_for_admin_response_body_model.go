// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateSupabaseForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AllocateSupabaseForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *AllocateSupabaseForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *AllocateSupabaseForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *AllocateSupabaseForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *AllocateSupabaseForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *AllocateSupabaseForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *AllocateSupabaseForAdminResponseBodyModule) *AllocateSupabaseForAdminResponseBody
	GetModule() *AllocateSupabaseForAdminResponseBodyModule
	SetRequestId(v string) *AllocateSupabaseForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *AllocateSupabaseForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *AllocateSupabaseForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *AllocateSupabaseForAdminResponseBody
	GetSynchro() *bool
}

type AllocateSupabaseForAdminResponseBody struct {
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
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                               `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *AllocateSupabaseForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s AllocateSupabaseForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateSupabaseForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateSupabaseForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AllocateSupabaseForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *AllocateSupabaseForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *AllocateSupabaseForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *AllocateSupabaseForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *AllocateSupabaseForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *AllocateSupabaseForAdminResponseBody) GetModule() *AllocateSupabaseForAdminResponseBodyModule {
	return s.Module
}

func (s *AllocateSupabaseForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateSupabaseForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *AllocateSupabaseForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *AllocateSupabaseForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *AllocateSupabaseForAdminResponseBody) SetAccessDeniedDetail(v string) *AllocateSupabaseForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetAllowRetry(v bool) *AllocateSupabaseForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetAppName(v string) *AllocateSupabaseForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetDynamicCode(v string) *AllocateSupabaseForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetDynamicMessage(v string) *AllocateSupabaseForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetErrorArgs(v []interface{}) *AllocateSupabaseForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetModule(v *AllocateSupabaseForAdminResponseBodyModule) *AllocateSupabaseForAdminResponseBody {
	s.Module = v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetRequestId(v string) *AllocateSupabaseForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetRootErrorCode(v string) *AllocateSupabaseForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetRootErrorMsg(v string) *AllocateSupabaseForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) SetSynchro(v bool) *AllocateSupabaseForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AllocateSupabaseForAdminResponseBodyModule struct {
	// example:
	//
	// 1111
	AnonKey *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	// example:
	//
	// WS20250915163734000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 111
	DbInstanceCreateTime *string `json:"DbInstanceCreateTime,omitempty" xml:"DbInstanceCreateTime,omitempty"`
	// example:
	//
	// selectdb-cn-2bl4djolb02
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 111
	DbPublicUrl *string `json:"DbPublicUrl,omitempty" xml:"DbPublicUrl,omitempty"`
	// example:
	//
	// ORACLE
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// {\\"appId\\":\\"APP_NTJAK8P11SNZDJ3M6BWC\\"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 111
	InstanceCreateFinishedTime *string `json:"InstanceCreateFinishedTime,omitempty" xml:"InstanceCreateFinishedTime,omitempty"`
	// example:
	//
	// 1111
	InstanceCreateStatus *string `json:"InstanceCreateStatus,omitempty" xml:"InstanceCreateStatus,omitempty"`
	// example:
	//
	// false
	IsDeleted *int32 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// 111
	RdsDatabasePassword *string `json:"RdsDatabasePassword,omitempty" xml:"RdsDatabasePassword,omitempty"`
	// example:
	//
	// e80f5a7a08514709a2fb
	ServiceKey *string `json:"ServiceKey,omitempty" xml:"ServiceKey,omitempty"`
	// example:
	//
	// NORMAL
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1111
	SupabaseDashboardPassword *string `json:"SupabaseDashboardPassword,omitempty" xml:"SupabaseDashboardPassword,omitempty"`
	// example:
	//
	// 111
	SupabaseDashboardUserName *string `json:"SupabaseDashboardUserName,omitempty" xml:"SupabaseDashboardUserName,omitempty"`
	// example:
	//
	// 111
	SupabaseInstanceCreateTime *string `json:"SupabaseInstanceCreateTime,omitempty" xml:"SupabaseInstanceCreateTime,omitempty"`
	// example:
	//
	// 1111
	SupabaseInstanceId *string `json:"SupabaseInstanceId,omitempty" xml:"SupabaseInstanceId,omitempty"`
	// Supabase Kong URL
	//
	// example:
	//
	// 111
	SupabaseKongUrl *string `json:"SupabaseKongUrl,omitempty" xml:"SupabaseKongUrl,omitempty"`
	// example:
	//
	// 1111
	SupabasePrivateIp *string `json:"SupabasePrivateIp,omitempty" xml:"SupabasePrivateIp,omitempty"`
	// example:
	//
	// 111
	SupabasePublicIp *string `json:"SupabasePublicIp,omitempty" xml:"SupabasePublicIp,omitempty"`
	// example:
	//
	// 111
	SupabasePublicUrl *string `json:"SupabasePublicUrl,omitempty" xml:"SupabasePublicUrl,omitempty"`
	// example:
	//
	// GFCBAMJH-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AllocateSupabaseForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AllocateSupabaseForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetAnonKey() *string {
	return s.AnonKey
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetDbInstanceCreateTime() *string {
	return s.DbInstanceCreateTime
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetDbPublicUrl() *string {
	return s.DbPublicUrl
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetDbType() *string {
	return s.DbType
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetExtra() *string {
	return s.Extra
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetInstanceCreateFinishedTime() *string {
	return s.InstanceCreateFinishedTime
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetInstanceCreateStatus() *string {
	return s.InstanceCreateStatus
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetRdsDatabasePassword() *string {
	return s.RdsDatabasePassword
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetServiceKey() *string {
	return s.ServiceKey
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabaseDashboardPassword() *string {
	return s.SupabaseDashboardPassword
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabaseDashboardUserName() *string {
	return s.SupabaseDashboardUserName
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabaseInstanceCreateTime() *string {
	return s.SupabaseInstanceCreateTime
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabaseInstanceId() *string {
	return s.SupabaseInstanceId
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabaseKongUrl() *string {
	return s.SupabaseKongUrl
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabasePrivateIp() *string {
	return s.SupabasePrivateIp
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabasePublicIp() *string {
	return s.SupabasePublicIp
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetSupabasePublicUrl() *string {
	return s.SupabasePublicUrl
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetTenantId() *string {
	return s.TenantId
}

func (s *AllocateSupabaseForAdminResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetAnonKey(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.AnonKey = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetBizId(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetDbInstanceCreateTime(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.DbInstanceCreateTime = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetDbInstanceId(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.DbInstanceId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetDbPublicUrl(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.DbPublicUrl = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetDbType(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.DbType = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetExtra(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.Extra = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetInstanceCreateFinishedTime(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.InstanceCreateFinishedTime = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetInstanceCreateStatus(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.InstanceCreateStatus = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetIsDeleted(v int32) *AllocateSupabaseForAdminResponseBodyModule {
	s.IsDeleted = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetRdsDatabasePassword(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.RdsDatabasePassword = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetServiceKey(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.ServiceKey = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetStatus(v int32) *AllocateSupabaseForAdminResponseBodyModule {
	s.Status = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabaseDashboardPassword(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabaseDashboardPassword = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabaseDashboardUserName(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabaseDashboardUserName = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabaseInstanceCreateTime(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabaseInstanceCreateTime = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabaseInstanceId(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabaseInstanceId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabaseKongUrl(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabaseKongUrl = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabasePrivateIp(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabasePrivateIp = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabasePublicIp(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabasePublicIp = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetSupabasePublicUrl(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.SupabasePublicUrl = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetTenantId(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.TenantId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) SetUserId(v string) *AllocateSupabaseForAdminResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *AllocateSupabaseForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
