// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyInstalledAppPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMyInstalledAppPluginsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListMyInstalledAppPluginsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListMyInstalledAppPluginsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListMyInstalledAppPluginsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMyInstalledAppPluginsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListMyInstalledAppPluginsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListMyInstalledAppPluginsResponseBodyModule) *ListMyInstalledAppPluginsResponseBody
	GetModule() *ListMyInstalledAppPluginsResponseBodyModule
	SetRequestId(v string) *ListMyInstalledAppPluginsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListMyInstalledAppPluginsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListMyInstalledAppPluginsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListMyInstalledAppPluginsResponseBody
	GetSynchro() *bool
}

type ListMyInstalledAppPluginsResponseBody struct {
	// The detailed reason why access is denied.
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
	// The application name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` in the **ErrMessage*	- return parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *ListMyInstalledAppPluginsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// System exception.
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListMyInstalledAppPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMyInstalledAppPluginsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListMyInstalledAppPluginsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListMyInstalledAppPluginsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMyInstalledAppPluginsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMyInstalledAppPluginsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListMyInstalledAppPluginsResponseBody) GetModule() *ListMyInstalledAppPluginsResponseBodyModule {
	return s.Module
}

func (s *ListMyInstalledAppPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMyInstalledAppPluginsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListMyInstalledAppPluginsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListMyInstalledAppPluginsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListMyInstalledAppPluginsResponseBody) SetAccessDeniedDetail(v string) *ListMyInstalledAppPluginsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetAllowRetry(v bool) *ListMyInstalledAppPluginsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetAppName(v string) *ListMyInstalledAppPluginsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetDynamicCode(v string) *ListMyInstalledAppPluginsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetDynamicMessage(v string) *ListMyInstalledAppPluginsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetErrorArgs(v []interface{}) *ListMyInstalledAppPluginsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetModule(v *ListMyInstalledAppPluginsResponseBodyModule) *ListMyInstalledAppPluginsResponseBody {
	s.Module = v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetRequestId(v string) *ListMyInstalledAppPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetRootErrorCode(v string) *ListMyInstalledAppPluginsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetRootErrorMsg(v string) *ListMyInstalledAppPluginsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) SetSynchro(v bool) *ListMyInstalledAppPluginsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyInstalledAppPluginsResponseBodyModule struct {
	// The current page number.
	//
	// example:
	//
	// 12
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The request results.
	Data []*ListMyInstalledAppPluginsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The decision weight.
	Next *ListMyInstalledAppPluginsResponseBodyModuleNext `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
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
	// Apart from pagination limits, the server processes a maximum of 1000 recent records for the current query. If the results exceed 1000 entries, **ResultLimit*	- is **true**. Narrow the time range and search again. Otherwise, **ResultLimit*	- is **false**.
	//
	// example:
	//
	// False
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListMyInstalledAppPluginsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetData() []*ListMyInstalledAppPluginsResponseBodyModuleData {
	return s.Data
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetNext() *ListMyInstalledAppPluginsResponseBodyModuleNext {
	return s.Next
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetCurrentPageNum(v int32) *ListMyInstalledAppPluginsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetData(v []*ListMyInstalledAppPluginsResponseBodyModuleData) *ListMyInstalledAppPluginsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetNext(v *ListMyInstalledAppPluginsResponseBodyModuleNext) *ListMyInstalledAppPluginsResponseBodyModule {
	s.Next = v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetNextPage(v bool) *ListMyInstalledAppPluginsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetPageSize(v int32) *ListMyInstalledAppPluginsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetPrePage(v bool) *ListMyInstalledAppPluginsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetResultLimit(v bool) *ListMyInstalledAppPluginsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetTotalItemNum(v int32) *ListMyInstalledAppPluginsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) SetTotalPageNum(v int32) *ListMyInstalledAppPluginsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModule) Validate() error {
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

type ListMyInstalledAppPluginsResponseBodyModuleData struct {
	// The Ray cluster type. Valid values:
	//
	// - BASIC: basic type, non-high-availability.
	//
	// - HIGH_AVAILABILITY: high-availability type.
	//
	// example:
	//
	// high_performance
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
	// 1763605196000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-10-11 16:05:43
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The popularity count.
	//
	// example:
	//
	// 1
	HotCount *int32 `json:"HotCount,omitempty" xml:"HotCount,omitempty"`
	// The plug-in icon.
	//
	// example:
	//
	// @lALPM2AwTOg9IUHNAUDNAUA
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
	// wx*
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The file status.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The API version number.
	//
	// example:
	//
	// 2025-04-29
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The visibility. Valid values:
	//
	// - **PUBLIC**: Public.
	//
	// - **PRIVATE**: Private.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMyInstalledAppPluginsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetCategory() *string {
	return s.Category
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetDescription() *string {
	return s.Description
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetHotCount() *int32 {
	return s.HotCount
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetIcon() *string {
	return s.Icon
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetStatus() *string {
	return s.Status
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetVersion() *string {
	return s.Version
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetCategory(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Category = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetDescription(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Description = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetGmtCreate(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.GmtCreate = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetGmtModified(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.GmtModified = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetHotCount(v int32) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.HotCount = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetIcon(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Icon = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetPluginId(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.PluginId = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetPluginName(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.PluginName = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetStatus(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetVersion(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Version = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) SetVisibility(v string) *ListMyInstalledAppPluginsResponseBodyModuleData {
	s.Visibility = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type ListMyInstalledAppPluginsResponseBodyModuleNext struct {
	// The category.
	//
	// example:
	//
	// {\\"Depth\\": 2, \\"ParentCategoryId\\": 5260, \\"Name\\": \\"2\\", \\"CategoryId\\": 5280, \\"CreateTime\\": \\"2025-08-26 10:20:37\\"}
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
	// The plug-in icon.
	//
	// example:
	//
	// @lADPD0-jULc8yJ3NC9DNC9A
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The bound API Gateway plug-in ID.
	//
	// example:
	//
	// alipay-isv
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// errorMapping
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// Valid values: unknown, init, testing, online.
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
	// Filters by visibility level. Multiple values can be specified.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMyInstalledAppPluginsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetCategory() *string {
	return s.Category
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetDescription() *string {
	return s.Description
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetIcon() *string {
	return s.Icon
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetPluginId() *string {
	return s.PluginId
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetPluginName() *string {
	return s.PluginName
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetStatus() *string {
	return s.Status
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetVersion() *string {
	return s.Version
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) GetVisibility() *string {
	return s.Visibility
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetCategory(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Category = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetDescription(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Description = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetGmtCreate(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.GmtCreate = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetGmtModified(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.GmtModified = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetIcon(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Icon = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetPluginId(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.PluginId = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetPluginName(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.PluginName = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetStatus(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Status = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetVersion(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Version = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) SetVisibility(v string) *ListMyInstalledAppPluginsResponseBodyModuleNext {
	s.Visibility = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}
