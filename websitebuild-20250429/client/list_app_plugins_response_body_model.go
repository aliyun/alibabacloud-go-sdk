// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppPluginsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppPluginsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppPluginsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppPluginsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppPluginsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppPluginsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppPluginsResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListAppPluginsResponseBodyModule) *ListAppPluginsResponseBody
	GetModule() []*ListAppPluginsResponseBodyModule
	SetNextToken(v string) *ListAppPluginsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppPluginsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppPluginsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppPluginsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppPluginsResponseBody
	GetSynchro() *bool
}

type ListAppPluginsResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application name. Query the application with this name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Value range: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Returned object.
	Module []*ListAppPluginsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// Token for starting the next query. It is empty if there is no next query.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Abnormal message
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

func (s ListAppPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPluginsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppPluginsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppPluginsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppPluginsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppPluginsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppPluginsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppPluginsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPluginsResponseBody) GetModule() []*ListAppPluginsResponseBodyModule {
	return s.Module
}

func (s *ListAppPluginsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPluginsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppPluginsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppPluginsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppPluginsResponseBody) SetAccessDeniedDetail(v string) *ListAppPluginsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetAllowRetry(v bool) *ListAppPluginsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetAppName(v string) *ListAppPluginsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetDynamicCode(v string) *ListAppPluginsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetDynamicMessage(v string) *ListAppPluginsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetErrorArgs(v []interface{}) *ListAppPluginsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppPluginsResponseBody) SetMaxResults(v int32) *ListAppPluginsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetModule(v []*ListAppPluginsResponseBodyModule) *ListAppPluginsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppPluginsResponseBody) SetNextToken(v string) *ListAppPluginsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetRequestId(v string) *ListAppPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetRootErrorCode(v string) *ListAppPluginsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetRootErrorMsg(v string) *ListAppPluginsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppPluginsResponseBody) SetSynchro(v bool) *ListAppPluginsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppPluginsResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppPluginsResponseBodyModule struct {
	// plugin code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Configuration form in React JSON Schema format
	//
	// example:
	//
	// ***
	ConfigItems *string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty"`
	// plugin Description
	//
	// example:
	//
	// fail to decode json
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Image display. Valid values:
	//
	// - **0*	- (None): Do not display.
	//
	// - **1*	- (Always): Always display.
	//
	// example:
	//
	// always
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// Indicates whether scheduled delivery of resource snapshots is enabled.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Shutdown.
	//
	// example:
	//
	// True
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// environment
	//
	// example:
	//
	// pre
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Creation Time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 1740479834
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// hook definitions
	//
	// example:
	//
	// []
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// plugin Description
	//
	// example:
	//
	// @lALPM2AwTOg9IUHNAUDNAUA
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// primary key
	//
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the item has been deleted: 0—No, 1—Yes
	//
	// example:
	//
	// false
	IsDeleted *int32 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// plugin Name
	//
	// example:
	//
	// 文件名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// skill header, used by the model for selection
	//
	// example:
	//
	// header
	SkillHeader *string `json:"SkillHeader,omitempty" xml:"SkillHeader,omitempty"`
	// Categorization label
	//
	// example:
	//
	// [{\\"Key\\": \\"kubernetes.io/cluster-id\\", \\"Value\\": \\"cc67198b13db948c9848599654da5586e\\"}, {\\"Key\\": \\"created-by\\", \\"Value\\": \\"alibabacloud-imagecache-controller\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListAppPluginsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppPluginsResponseBodyModule) GetCode() *string {
	return s.Code
}

func (s *ListAppPluginsResponseBodyModule) GetConfigItems() *string {
	return s.ConfigItems
}

func (s *ListAppPluginsResponseBodyModule) GetDesc() *string {
	return s.Desc
}

func (s *ListAppPluginsResponseBodyModule) GetDisplay() *int32 {
	return s.Display
}

func (s *ListAppPluginsResponseBodyModule) GetEnabled() *int32 {
	return s.Enabled
}

func (s *ListAppPluginsResponseBodyModule) GetEnv() *string {
	return s.Env
}

func (s *ListAppPluginsResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAppPluginsResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAppPluginsResponseBodyModule) GetHooks() *string {
	return s.Hooks
}

func (s *ListAppPluginsResponseBodyModule) GetIcon() *string {
	return s.Icon
}

func (s *ListAppPluginsResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *ListAppPluginsResponseBodyModule) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *ListAppPluginsResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *ListAppPluginsResponseBodyModule) GetSkillHeader() *string {
	return s.SkillHeader
}

func (s *ListAppPluginsResponseBodyModule) GetTags() *string {
	return s.Tags
}

func (s *ListAppPluginsResponseBodyModule) SetCode(v string) *ListAppPluginsResponseBodyModule {
	s.Code = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetConfigItems(v string) *ListAppPluginsResponseBodyModule {
	s.ConfigItems = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetDesc(v string) *ListAppPluginsResponseBodyModule {
	s.Desc = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetDisplay(v int32) *ListAppPluginsResponseBodyModule {
	s.Display = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetEnabled(v int32) *ListAppPluginsResponseBodyModule {
	s.Enabled = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetEnv(v string) *ListAppPluginsResponseBodyModule {
	s.Env = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetGmtCreateTime(v string) *ListAppPluginsResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetGmtModifiedTime(v string) *ListAppPluginsResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetHooks(v string) *ListAppPluginsResponseBodyModule {
	s.Hooks = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetIcon(v string) *ListAppPluginsResponseBodyModule {
	s.Icon = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetId(v int64) *ListAppPluginsResponseBodyModule {
	s.Id = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetIsDeleted(v int32) *ListAppPluginsResponseBodyModule {
	s.IsDeleted = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetName(v string) *ListAppPluginsResponseBodyModule {
	s.Name = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetSkillHeader(v string) *ListAppPluginsResponseBodyModule {
	s.SkillHeader = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) SetTags(v string) *ListAppPluginsResponseBodyModule {
	s.Tags = &v
	return s
}

func (s *ListAppPluginsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
