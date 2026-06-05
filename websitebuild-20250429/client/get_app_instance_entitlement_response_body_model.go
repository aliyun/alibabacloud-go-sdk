// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceEntitlementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppInstanceEntitlementResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppInstanceEntitlementResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppInstanceEntitlementResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppInstanceEntitlementResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppInstanceEntitlementResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppInstanceEntitlementResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppInstanceEntitlementResponseBodyModule) *GetAppInstanceEntitlementResponseBody
	GetModule() *GetAppInstanceEntitlementResponseBodyModule
	SetRequestId(v string) *GetAppInstanceEntitlementResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppInstanceEntitlementResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppInstanceEntitlementResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppInstanceEntitlementResponseBody
	GetSynchro() *bool
}

type GetAppInstanceEntitlementResponseBody struct {
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
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                      `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppInstanceEntitlementResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAppInstanceEntitlementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceEntitlementResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppInstanceEntitlementResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppInstanceEntitlementResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppInstanceEntitlementResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppInstanceEntitlementResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppInstanceEntitlementResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppInstanceEntitlementResponseBody) GetModule() *GetAppInstanceEntitlementResponseBodyModule {
	return s.Module
}

func (s *GetAppInstanceEntitlementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppInstanceEntitlementResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppInstanceEntitlementResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppInstanceEntitlementResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppInstanceEntitlementResponseBody) SetAccessDeniedDetail(v string) *GetAppInstanceEntitlementResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetAllowRetry(v bool) *GetAppInstanceEntitlementResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetAppName(v string) *GetAppInstanceEntitlementResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetDynamicCode(v string) *GetAppInstanceEntitlementResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetDynamicMessage(v string) *GetAppInstanceEntitlementResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetErrorArgs(v []interface{}) *GetAppInstanceEntitlementResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetModule(v *GetAppInstanceEntitlementResponseBodyModule) *GetAppInstanceEntitlementResponseBody {
	s.Module = v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetRequestId(v string) *GetAppInstanceEntitlementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetRootErrorCode(v string) *GetAppInstanceEntitlementResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetRootErrorMsg(v string) *GetAppInstanceEntitlementResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) SetSynchro(v bool) *GetAppInstanceEntitlementResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppInstanceEntitlementResponseBodyModule struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string                                             `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Items []*GetAppInstanceEntitlementResponseBodyModuleItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetAppInstanceEntitlementResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceEntitlementResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppInstanceEntitlementResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceEntitlementResponseBodyModule) GetItems() []*GetAppInstanceEntitlementResponseBodyModuleItems {
	return s.Items
}

func (s *GetAppInstanceEntitlementResponseBodyModule) SetBizId(v string) *GetAppInstanceEntitlementResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModule) SetItems(v []*GetAppInstanceEntitlementResponseBodyModuleItems) *GetAppInstanceEntitlementResponseBodyModule {
	s.Items = v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModule) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppInstanceEntitlementResponseBodyModuleItems struct {
	Allocated *bool `json:"Allocated,omitempty" xml:"Allocated,omitempty"`
	// example:
	//
	// on
	Available *bool `json:"Available,omitempty" xml:"Available,omitempty"`
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
	Entitled   *bool `json:"Entitled,omitempty" xml:"Entitled,omitempty"`
	// example:
	//
	// 5
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// waf_v2intl_public_intl-sg-i5c43rcpw04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 4e46d24b56bfa944b5e6f2305715bc4e.jpg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// alipay-isv
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// 10
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 1234
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// example:
	//
	// SmsCount
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// example:
	//
	// filesystem
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 0
	Running *bool `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// question
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 10
	UsagePercent *int32 `json:"UsagePercent,omitempty" xml:"UsagePercent,omitempty"`
	// example:
	//
	// 3295422523872
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s GetAppInstanceEntitlementResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceEntitlementResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetAllocated() *bool {
	return s.Allocated
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetAvailable() *bool {
	return s.Available
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetCode() *string {
	return s.Code
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetConfigured() *bool {
	return s.Configured
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetEntitled() *bool {
	return s.Entitled
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetFeatureType() *string {
	return s.FeatureType
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetName() *string {
	return s.Name
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetQuota() *int64 {
	return s.Quota
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetRemaining() *int64 {
	return s.Remaining
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetRunning() *bool {
	return s.Running
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetType() *string {
	return s.Type
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetUsagePercent() *int32 {
	return s.UsagePercent
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) GetUsed() *int64 {
	return s.Used
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetAllocated(v bool) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Allocated = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetAvailable(v bool) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Available = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetCode(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Code = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetConfigured(v bool) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Configured = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetEntitled(v bool) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Entitled = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetFeatureType(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.FeatureType = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetInstanceId(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.InstanceId = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetName(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Name = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetPluginId(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.PluginId = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetQuota(v int64) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Quota = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetRemaining(v int64) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Remaining = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetResourceCode(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.ResourceCode = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetResourceType(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.ResourceType = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetRunning(v bool) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Running = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetType(v string) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Type = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetUsagePercent(v int32) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.UsagePercent = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) SetUsed(v int64) *GetAppInstanceEntitlementResponseBodyModuleItems {
	s.Used = &v
	return s
}

func (s *GetAppInstanceEntitlementResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
