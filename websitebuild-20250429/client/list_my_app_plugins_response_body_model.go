// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAppPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMyAppPluginsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListMyAppPluginsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListMyAppPluginsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListMyAppPluginsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMyAppPluginsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListMyAppPluginsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListMyAppPluginsResponseBodyModule) *ListMyAppPluginsResponseBody
	GetModule() *ListMyAppPluginsResponseBodyModule
	SetRequestId(v string) *ListMyAppPluginsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListMyAppPluginsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListMyAppPluginsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListMyAppPluginsResponseBody
	GetSynchro() *bool
}

type ListMyAppPluginsResponseBody struct {
	// The detailed reason why access is denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name. The name must start with a letter and can contain digits, letters, underscores (_), and hyphens (-). It can be up to 36 characters in length.
	//
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The task object.
	Module *ListMyAppPluginsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
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

func (s ListMyAppPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMyAppPluginsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListMyAppPluginsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListMyAppPluginsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMyAppPluginsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMyAppPluginsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListMyAppPluginsResponseBody) GetModule() *ListMyAppPluginsResponseBodyModule {
	return s.Module
}

func (s *ListMyAppPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMyAppPluginsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListMyAppPluginsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListMyAppPluginsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListMyAppPluginsResponseBody) SetAccessDeniedDetail(v string) *ListMyAppPluginsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetAllowRetry(v bool) *ListMyAppPluginsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetAppName(v string) *ListMyAppPluginsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetDynamicCode(v string) *ListMyAppPluginsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetDynamicMessage(v string) *ListMyAppPluginsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetErrorArgs(v []interface{}) *ListMyAppPluginsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetModule(v *ListMyAppPluginsResponseBodyModule) *ListMyAppPluginsResponseBody {
	s.Module = v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetRequestId(v string) *ListMyAppPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetRootErrorCode(v string) *ListMyAppPluginsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetRootErrorMsg(v string) *ListMyAppPluginsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) SetSynchro(v bool) *ListMyAppPluginsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListMyAppPluginsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyAppPluginsResponseBodyModule struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The request results.
	Data []*ListMyAppPluginsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The decision weight.
	Next *ListMyAppPluginsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
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
	// The result limit. In addition to the pagination limit, the server processes a maximum of the latest 1,000 records for the current query. If the number of results exceeds 1,000, **ResultLimit*	- is **true**. In this case, narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
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

func (s ListMyAppPluginsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListMyAppPluginsResponseBodyModule) GetData() []*ListMyAppPluginsResponseBodyModuleData {
	return s.Data
}

func (s *ListMyAppPluginsResponseBodyModule) GetNext() *ListMyAppPluginsResponseBodyModuleNext {
	return s.Next
}

func (s *ListMyAppPluginsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListMyAppPluginsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyAppPluginsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListMyAppPluginsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListMyAppPluginsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListMyAppPluginsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListMyAppPluginsResponseBodyModule) SetCurrentPageNum(v int32) *ListMyAppPluginsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetData(v []*ListMyAppPluginsResponseBodyModuleData) *ListMyAppPluginsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetNext(v *ListMyAppPluginsResponseBodyModuleNext) *ListMyAppPluginsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetNextPage(v bool) *ListMyAppPluginsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetPageSize(v int32) *ListMyAppPluginsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetPrePage(v bool) *ListMyAppPluginsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetResultLimit(v bool) *ListMyAppPluginsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetTotalItemNum(v int32) *ListMyAppPluginsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) SetTotalPageNum(v int32) *ListMyAppPluginsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModule) Validate() error {
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

type ListMyAppPluginsResponseBodyModuleData struct {
	// The disk specifications.
	//
	// example:
	//
	// NormalMultimaster
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The scenario description.
	//
	// example:
	//
	// lshm-mysql-coypt
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-26T10:12:18.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
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
	// The plugin description.
	//
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// alipay-isv
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plugin name.
	//
	// example:
	//
	// transformer
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
	// - **PUBLIC**: Public.
	//
	// - **PRIVATE**: Private.
	//
	// example:
	//
	// visible
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMyAppPluginsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetCategory() *string {
	return s.Category
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetDescription() *string {
	return s.Description
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetHotCount() *int32 {
	return s.HotCount
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetIcon() *string {
	return s.Icon
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetVersion() *string {
	return s.Version
}

func (s *ListMyAppPluginsResponseBodyModuleData) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetCategory(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Category = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetDescription(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Description = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetGmtCreate(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.GmtCreate = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetGmtModified(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.GmtModified = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetHotCount(v int32) *ListMyAppPluginsResponseBodyModuleData {
	s.HotCount = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetIcon(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Icon = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetPluginId(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.PluginId = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetPluginName(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.PluginName = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetStatus(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetVersion(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Version = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) SetVisibility(v string) *ListMyAppPluginsResponseBodyModuleData {
	s.Visibility = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListMyAppPluginsResponseBodyModuleNext struct {
	// The controlType converted based on the associated public category.
	//
	// example:
	//
	// CRAC
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
	// The modification time.
	//
	// example:
	//
	// 2026-03-18T10:03:56+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The plugin description.
	//
	// example:
	//
	// @lADPDetfgMsFFUvNAkjNAkg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// alipay-isv
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plugin name.
	//
	// example:
	//
	// trafficControl
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
	// The visibility. Valid values:
	//
	// - **PUBLIC**: Public.
	//
	// - **PRIVATE**: Private.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMyAppPluginsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetCategory() *string {
	return s.Category
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetDescription() *string {
	return s.Description
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetIcon() *string {
	return s.Icon
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetStatus() *string {
	return s.Status
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetVersion() *string {
	return s.Version
}

func (s *ListMyAppPluginsResponseBodyModuleNext) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetCategory(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Category = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetDescription(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Description = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetGmtCreate(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.GmtCreate = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetGmtModified(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.GmtModified = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetIcon(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Icon = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetPluginId(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.PluginId = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetPluginName(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.PluginName = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetStatus(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Status = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetVersion(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Version = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) SetVisibility(v string) *ListMyAppPluginsResponseBodyModuleNext {
	s.Visibility = &v
	return s
}

func (s *ListMyAppPluginsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
