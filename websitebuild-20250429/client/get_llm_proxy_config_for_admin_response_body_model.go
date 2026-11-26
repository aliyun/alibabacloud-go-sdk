// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmProxyConfigForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLlmProxyConfigForAdminResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetLlmProxyConfigForAdminResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetLlmProxyConfigForAdminResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetLlmProxyConfigForAdminResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetLlmProxyConfigForAdminResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetLlmProxyConfigForAdminResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetLlmProxyConfigForAdminResponseBodyModule) *GetLlmProxyConfigForAdminResponseBody
	GetModule() *GetLlmProxyConfigForAdminResponseBodyModule
	SetRequestId(v string) *GetLlmProxyConfigForAdminResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetLlmProxyConfigForAdminResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetLlmProxyConfigForAdminResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetLlmProxyConfigForAdminResponseBody
	GetSynchro() *bool
}

type GetLlmProxyConfigForAdminResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether to allow retries. Valid values:
	//
	// - false: Retries are not allowed.
	//
	// - true: Retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application to query.
	//
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic code. This field is not used and can be ignored.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The returned object.
	Module *GetLlmProxyConfigForAdminResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetLlmProxyConfigForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLlmProxyConfigForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetModule() *GetLlmProxyConfigForAdminResponseBodyModule {
	return s.Module
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetLlmProxyConfigForAdminResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetAccessDeniedDetail(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetAllowRetry(v bool) *GetLlmProxyConfigForAdminResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetAppName(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.AppName = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetDynamicCode(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetDynamicMessage(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetErrorArgs(v []interface{}) *GetLlmProxyConfigForAdminResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetModule(v *GetLlmProxyConfigForAdminResponseBodyModule) *GetLlmProxyConfigForAdminResponseBody {
	s.Module = v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetRequestId(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetRootErrorCode(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetRootErrorMsg(v string) *GetLlmProxyConfigForAdminResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) SetSynchro(v bool) *GetLlmProxyConfigForAdminResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLlmProxyConfigForAdminResponseBodyModule struct {
	// The list of supported models.
	//
	// example:
	//
	// qwen3.5-plus
	AllowedModels *string `json:"AllowedModels,omitempty" xml:"AllowedModels,omitempty"`
	// The business ID of the application instance.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The list of blocked models.
	//
	// example:
	//
	// qwen3.5-plus
	BlockedModels *string `json:"BlockedModels,omitempty" xml:"BlockedModels,omitempty"`
	// Grants specific permissions to processes in the container. Currently, only NET_ADMIN and NET_RAW are supported.
	//
	// > NET_RAW is not supported by default. Submit a ticket to apply for this permission.
	//
	// example:
	//
	// adds2
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// The daily request limit.
	//
	// example:
	//
	// -1
	DailyLimit *int32 `json:"DailyLimit,omitempty" xml:"DailyLimit,omitempty"`
	// The daily token limit.
	//
	// example:
	//
	// -1
	DailyTokenLimit *int64 `json:"DailyTokenLimit,omitempty" xml:"DailyTokenLimit,omitempty"`
	// Specifies whether to enable scheduled delivery of resource snapshots.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The extended configuration in JSON format.
	//
	// example:
	//
	// {"deliveryNodeName":"Bind domain name and publish","deliveryNodeStatus":"Finish","deliveryOperatorRole":"Provider"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key.
	//
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP blacklist.
	//
	// example:
	//
	// 121.41.11.161,10.200.255.60
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP whitelist. Separate multiple IP addresses with commas.
	//
	// example:
	//
	// 121.41.11.161,10.200.255.60
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The number of requests per minute.
	//
	// example:
	//
	// -1
	RpmLimit *int32 `json:"RpmLimit,omitempty" xml:"RpmLimit,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason for suspension.
	//
	// example:
	//
	// xxx
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
}

func (s GetLlmProxyConfigForAdminResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetLlmProxyConfigForAdminResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetBlockedModels() *string {
	return s.BlockedModels
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetCapability() *string {
	return s.Capability
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetDailyLimit() *int32 {
	return s.DailyLimit
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetDailyTokenLimit() *int64 {
	return s.DailyTokenLimit
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetExtend() *string {
	return s.Extend
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetRpmLimit() *int32 {
	return s.RpmLimit
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) GetSuspendReason() *string {
	return s.SuspendReason
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetAllowedModels(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.AllowedModels = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetBizId(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetBlockedModels(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.BlockedModels = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetCapability(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.Capability = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetDailyLimit(v int32) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.DailyLimit = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetDailyTokenLimit(v int64) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.DailyTokenLimit = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetEnabled(v bool) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.Enabled = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetExtend(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.Extend = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetGmtCreate(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetGmtModified(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetId(v int64) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.Id = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetIpBlacklist(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.IpBlacklist = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetIpWhitelist(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.IpWhitelist = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetRpmLimit(v int32) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.RpmLimit = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetStatus(v int32) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) SetSuspendReason(v string) *GetLlmProxyConfigForAdminResponseBodyModule {
	s.SuspendReason = &v
	return s
}

func (s *GetLlmProxyConfigForAdminResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
