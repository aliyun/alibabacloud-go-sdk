// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseInstanceInfoForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QuerySupabaseInstanceInfoForAdminResponseBodyModule) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetModule() *QuerySupabaseInstanceInfoForAdminResponseBodyModule
	SetRequestId(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QuerySupabaseInstanceInfoForAdminResponseBody
	GetSynchro() *bool
}

type QuerySupabaseInstanceInfoForAdminResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                        `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QuerySupabaseInstanceInfoForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s QuerySupabaseInstanceInfoForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseInstanceInfoForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetModule() *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	return s.Module
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetAccessDeniedDetail(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetAllowRetry(v bool) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetAppName(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetDynamicCode(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetDynamicMessage(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetErrorArgs(v []interface{}) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetModule(v *QuerySupabaseInstanceInfoForAdminResponseBodyModule) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.Module = v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetRequestId(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetRootErrorCode(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetRootErrorMsg(v string) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) SetSynchro(v bool) *QuerySupabaseInstanceInfoForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySupabaseInstanceInfoForAdminResponseBodyModule struct {
	// example:
	//
	// 1111
	AnonKey *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	// example:
	//
	// WD20250703155602000001
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
	// RUNNING
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1111
	SupabaseDashboardPassword *string `json:"SupabaseDashboardPassword,omitempty" xml:"SupabaseDashboardPassword,omitempty"`
	// example:
	//
	// 1111
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
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySupabaseInstanceInfoForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseInstanceInfoForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetAnonKey() *string {
	return s.AnonKey
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetDbInstanceCreateTime() *string {
	return s.DbInstanceCreateTime
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetDbPublicUrl() *string {
	return s.DbPublicUrl
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetDbType() *string {
	return s.DbType
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetExtra() *string {
	return s.Extra
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetInstanceCreateFinishedTime() *string {
	return s.InstanceCreateFinishedTime
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetInstanceCreateStatus() *string {
	return s.InstanceCreateStatus
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetRdsDatabasePassword() *string {
	return s.RdsDatabasePassword
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetServiceKey() *string {
	return s.ServiceKey
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabaseDashboardPassword() *string {
	return s.SupabaseDashboardPassword
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabaseDashboardUserName() *string {
	return s.SupabaseDashboardUserName
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabaseInstanceCreateTime() *string {
	return s.SupabaseInstanceCreateTime
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabaseInstanceId() *string {
	return s.SupabaseInstanceId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabaseKongUrl() *string {
	return s.SupabaseKongUrl
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabasePrivateIp() *string {
	return s.SupabasePrivateIp
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabasePublicIp() *string {
	return s.SupabasePublicIp
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetSupabasePublicUrl() *string {
	return s.SupabasePublicUrl
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetAnonKey(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.AnonKey = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetBizId(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetDbInstanceCreateTime(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.DbInstanceCreateTime = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetDbInstanceId(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.DbInstanceId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetDbPublicUrl(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.DbPublicUrl = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetDbType(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.DbType = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetExtra(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.Extra = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetInstanceCreateFinishedTime(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.InstanceCreateFinishedTime = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetInstanceCreateStatus(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.InstanceCreateStatus = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetIsDeleted(v int32) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.IsDeleted = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetRdsDatabasePassword(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.RdsDatabasePassword = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetServiceKey(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.ServiceKey = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetStatus(v int32) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.Status = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabaseDashboardPassword(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabaseDashboardPassword = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabaseDashboardUserName(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabaseDashboardUserName = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabaseInstanceCreateTime(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabaseInstanceCreateTime = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabaseInstanceId(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabaseInstanceId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabaseKongUrl(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabaseKongUrl = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabasePrivateIp(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabasePrivateIp = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabasePublicIp(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabasePublicIp = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetSupabasePublicUrl(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.SupabasePublicUrl = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetTenantId(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.TenantId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) SetUserId(v string) *QuerySupabaseInstanceInfoForAdminResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
