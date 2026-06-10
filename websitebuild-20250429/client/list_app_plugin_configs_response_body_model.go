// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppPluginConfigsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppPluginConfigsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppPluginConfigsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppPluginConfigsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppPluginConfigsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppPluginConfigsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppPluginConfigsResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListAppPluginConfigsResponseBodyModule) *ListAppPluginConfigsResponseBody
	GetModule() []*ListAppPluginConfigsResponseBodyModule
	SetNextToken(v string) *ListAppPluginConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppPluginConfigsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppPluginConfigsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppPluginConfigsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppPluginConfigsResponseBody
	GetSynchro() *bool
}

type ListAppPluginConfigsResponseBody struct {
	// The detailed reason why access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// Application Name. Query the application with this name.
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
	// Dynamic error message used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The number of results returned per query.
	//
	// Valid range: 10 to 100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// task object
	Module []*ListAppPluginConfigsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// The token to start the next query. It is empty if there is no next query.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// is synchronous processing enabled
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppPluginConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPluginConfigsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppPluginConfigsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppPluginConfigsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppPluginConfigsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppPluginConfigsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppPluginConfigsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppPluginConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPluginConfigsResponseBody) GetModule() []*ListAppPluginConfigsResponseBodyModule {
	return s.Module
}

func (s *ListAppPluginConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPluginConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPluginConfigsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppPluginConfigsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppPluginConfigsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppPluginConfigsResponseBody) SetAccessDeniedDetail(v string) *ListAppPluginConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetAllowRetry(v bool) *ListAppPluginConfigsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetAppName(v string) *ListAppPluginConfigsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetDynamicCode(v string) *ListAppPluginConfigsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetDynamicMessage(v string) *ListAppPluginConfigsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetErrorArgs(v []interface{}) *ListAppPluginConfigsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetMaxResults(v int32) *ListAppPluginConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetModule(v []*ListAppPluginConfigsResponseBodyModule) *ListAppPluginConfigsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetNextToken(v string) *ListAppPluginConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetRequestId(v string) *ListAppPluginConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetRootErrorCode(v string) *ListAppPluginConfigsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetRootErrorMsg(v string) *ListAppPluginConfigsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) SetSynchro(v bool) *ListAppPluginConfigsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppPluginConfigsResponseBody) Validate() error {
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

type ListAppPluginConfigsResponseBodyModule struct {
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Indicates whether to display.
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
	// Extension information
	Extend map[string]*string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Creation Time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-10-11T06:25:13Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-05-20T01:59:17.000Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// primary key
	//
	// example:
	//
	// 16257
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Plugin configuration
	//
	// example:
	//
	// {}
	PluginConfig *string `json:"PluginConfig,omitempty" xml:"PluginConfig,omitempty"`
	// Plugin description
	//
	// example:
	//
	// a simple test plugin
	PluginDesc *string `json:"PluginDesc,omitempty" xml:"PluginDesc,omitempty"`
	// Plugin ID
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// Plugin name
	//
	// example:
	//
	// tf_testaccapigatewayplugin29311
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// Site name
	//
	// example:
	//
	// wkweb.cn
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// User ID
	//
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAppPluginConfigsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginConfigsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppPluginConfigsResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *ListAppPluginConfigsResponseBodyModule) GetDisplay() *int32 {
	return s.Display
}

func (s *ListAppPluginConfigsResponseBodyModule) GetEnabled() *int32 {
	return s.Enabled
}

func (s *ListAppPluginConfigsResponseBodyModule) GetExtend() map[string]*string {
	return s.Extend
}

func (s *ListAppPluginConfigsResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAppPluginConfigsResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAppPluginConfigsResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *ListAppPluginConfigsResponseBodyModule) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *ListAppPluginConfigsResponseBodyModule) GetPluginDesc() *string {
	return s.PluginDesc
}

func (s *ListAppPluginConfigsResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAppPluginConfigsResponseBodyModule) GetPluginName() *string {
	return s.PluginName
}

func (s *ListAppPluginConfigsResponseBodyModule) GetSiteName() *string {
	return s.SiteName
}

func (s *ListAppPluginConfigsResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *ListAppPluginConfigsResponseBodyModule) SetBizId(v string) *ListAppPluginConfigsResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetDisplay(v int32) *ListAppPluginConfigsResponseBodyModule {
	s.Display = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetEnabled(v int32) *ListAppPluginConfigsResponseBodyModule {
	s.Enabled = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetExtend(v map[string]*string) *ListAppPluginConfigsResponseBodyModule {
	s.Extend = v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetGmtCreateTime(v string) *ListAppPluginConfigsResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetGmtModifiedTime(v string) *ListAppPluginConfigsResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetId(v int64) *ListAppPluginConfigsResponseBodyModule {
	s.Id = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetPluginConfig(v string) *ListAppPluginConfigsResponseBodyModule {
	s.PluginConfig = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetPluginDesc(v string) *ListAppPluginConfigsResponseBodyModule {
	s.PluginDesc = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetPluginId(v string) *ListAppPluginConfigsResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetPluginName(v string) *ListAppPluginConfigsResponseBodyModule {
	s.PluginName = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetSiteName(v string) *ListAppPluginConfigsResponseBodyModule {
	s.SiteName = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) SetUserId(v string) *ListAppPluginConfigsResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ListAppPluginConfigsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
