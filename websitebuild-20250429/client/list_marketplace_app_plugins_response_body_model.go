// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketplaceAppPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMarketplaceAppPluginsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListMarketplaceAppPluginsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListMarketplaceAppPluginsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListMarketplaceAppPluginsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMarketplaceAppPluginsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListMarketplaceAppPluginsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListMarketplaceAppPluginsResponseBodyModule) *ListMarketplaceAppPluginsResponseBody
	GetModule() *ListMarketplaceAppPluginsResponseBodyModule
	SetRequestId(v string) *ListMarketplaceAppPluginsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListMarketplaceAppPluginsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListMarketplaceAppPluginsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListMarketplaceAppPluginsResponseBody
	GetSynchro() *bool
}

type ListMarketplaceAppPluginsResponseBody struct {
	// The detailed reason for the access denial.
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
	// The name of the application.
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
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *ListMarketplaceAppPluginsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListMarketplaceAppPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMarketplaceAppPluginsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListMarketplaceAppPluginsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListMarketplaceAppPluginsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMarketplaceAppPluginsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMarketplaceAppPluginsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListMarketplaceAppPluginsResponseBody) GetModule() *ListMarketplaceAppPluginsResponseBodyModule {
	return s.Module
}

func (s *ListMarketplaceAppPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMarketplaceAppPluginsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListMarketplaceAppPluginsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListMarketplaceAppPluginsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListMarketplaceAppPluginsResponseBody) SetAccessDeniedDetail(v string) *ListMarketplaceAppPluginsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetAllowRetry(v bool) *ListMarketplaceAppPluginsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetAppName(v string) *ListMarketplaceAppPluginsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetDynamicCode(v string) *ListMarketplaceAppPluginsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetDynamicMessage(v string) *ListMarketplaceAppPluginsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetErrorArgs(v []interface{}) *ListMarketplaceAppPluginsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetModule(v *ListMarketplaceAppPluginsResponseBodyModule) *ListMarketplaceAppPluginsResponseBody {
	s.Module = v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetRequestId(v string) *ListMarketplaceAppPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetRootErrorCode(v string) *ListMarketplaceAppPluginsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetRootErrorMsg(v string) *ListMarketplaceAppPluginsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) SetSynchro(v bool) *ListMarketplaceAppPluginsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMarketplaceAppPluginsResponseBodyModule struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The request results.
	Data []*ListMarketplaceAppPluginsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The next feature ID.
	Next *ListMarketplaceAppPluginsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// Indicates whether a next page exists.
	//
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether a previous page exists.
	//
	// example:
	//
	// False
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// In addition to the pagination limit, the server processes a maximum of the latest 1,000 records for the current query. If the number of results exceeds 1,000, **ResultLimit*	- is **true**. In this case, narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
	//
	// example:
	//
	// False
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListMarketplaceAppPluginsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetData() []*ListMarketplaceAppPluginsResponseBodyModuleData {
	return s.Data
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetNext() *ListMarketplaceAppPluginsResponseBodyModuleNext {
	return s.Next
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetCurrentPageNum(v int32) *ListMarketplaceAppPluginsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetData(v []*ListMarketplaceAppPluginsResponseBodyModuleData) *ListMarketplaceAppPluginsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetNext(v *ListMarketplaceAppPluginsResponseBodyModuleNext) *ListMarketplaceAppPluginsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetNextPage(v bool) *ListMarketplaceAppPluginsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetPageSize(v int32) *ListMarketplaceAppPluginsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetPrePage(v bool) *ListMarketplaceAppPluginsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetResultLimit(v bool) *ListMarketplaceAppPluginsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetTotalItemNum(v int32) *ListMarketplaceAppPluginsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) SetTotalPageNum(v int32) *ListMarketplaceAppPluginsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMarketplaceAppPluginsResponseBodyModuleData struct {
	// The category of the system cloud disk. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: basic disk.
	//
	// example:
	//
	// basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The scenario description.
	//
	// example:
	//
	// lshm-mysql-coypt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1756951455000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1556529068000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The popularity.
	//
	// example:
	//
	// 1
	HotCount *int32 `json:"HotCount,omitempty" xml:"HotCount,omitempty"`
	// The plug-in description.
	//
	// example:
	//
	// @lALPM2X6xumibpPNAgDNAgA
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// chat-ai
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// trafficControl
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The file status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The API version number.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The visibility. Valid values:
	//
	// - **PUBLIC**: public.
	//
	// - **PRIVATE**: private.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMarketplaceAppPluginsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetCategory() *string {
	return s.Category
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetDescription() *string {
	return s.Description
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetExtend() *string {
	return s.Extend
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetHotCount() *int32 {
	return s.HotCount
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetIcon() *string {
	return s.Icon
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetVersion() *string {
	return s.Version
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetCategory(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Category = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetDescription(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Description = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetExtend(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Extend = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetGmtCreate(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.GmtCreate = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetGmtModified(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.GmtModified = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetHotCount(v int32) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.HotCount = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetIcon(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Icon = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetPluginId(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.PluginId = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetPluginName(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.PluginName = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetStatus(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetVersion(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Version = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) SetVisibility(v string) *ListMarketplaceAppPluginsResponseBodyModuleData {
	s.Visibility = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListMarketplaceAppPluginsResponseBodyModuleNext struct {
	// The category of data cloud disk N. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// example:
	//
	// {\\"Status\\": 0, \\"Name\\": u\\"\\u81ea\\u52a8\\u521b\\u5efacom.alibaba.cloudrobot.apitest.resource.AbstractResource1773800213419\\", \\"BizCode\\": \\"\\", \\"CategoryId\\": 30000941833, \\"ParentCategoryId\\": 30000941832}
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The scenario description.
	//
	// example:
	//
	// lx_supabase_test_02
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-09-17 20:43:21
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-03-18T10:03:56+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The plug-in description.
	//
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// alipay-isv
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// tf_testaccapigatewayplugin68046
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The status. Valid values: unknown, init, testing, and online.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The API version number.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether the bucket to which the uploaded file belongs is publicly accessible. Valid values: **public*	- and **private**. A value of public indicates that the file is uploaded to a public bucket. A value of private indicates that the file is uploaded to a private bucket and requires authorization for access.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMarketplaceAppPluginsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListMarketplaceAppPluginsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetCategory() *string {
	return s.Category
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetDescription() *string {
	return s.Description
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetIcon() *string {
	return s.Icon
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetStatus() *string {
	return s.Status
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetVersion() *string {
	return s.Version
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetCategory(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Category = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetDescription(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Description = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetGmtCreate(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.GmtCreate = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetGmtModified(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.GmtModified = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetIcon(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Icon = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetPluginId(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.PluginId = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetPluginName(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.PluginName = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetStatus(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Status = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetVersion(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Version = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) SetVisibility(v string) *ListMarketplaceAppPluginsResponseBodyModuleNext {
	s.Visibility = &v
	return s
}

func (s *ListMarketplaceAppPluginsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
