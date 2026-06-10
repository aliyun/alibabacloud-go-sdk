// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsvPaymentPluginConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListIsvPaymentPluginConfigsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListIsvPaymentPluginConfigsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListIsvPaymentPluginConfigsResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListIsvPaymentPluginConfigsResponseBodyModule) *ListIsvPaymentPluginConfigsResponseBody
	GetModule() []*ListIsvPaymentPluginConfigsResponseBodyModule
	SetNextToken(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListIsvPaymentPluginConfigsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListIsvPaymentPluginConfigsResponseBody
	GetSynchro() *bool
}

type ListIsvPaymentPluginConfigsResponseBody struct {
	// Detailed reason for access denial.
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
	// App name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic code; currently unused. Ignore this field.
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
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Value range: 10–100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Response data
	Module []*ListIsvPaymentPluginConfigsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
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
	// abnormal message
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

func (s ListIsvPaymentPluginConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIsvPaymentPluginConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetModule() []*ListIsvPaymentPluginConfigsResponseBodyModule {
	return s.Module
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListIsvPaymentPluginConfigsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetAccessDeniedDetail(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetAllowRetry(v bool) *ListIsvPaymentPluginConfigsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetAppName(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetDynamicCode(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetDynamicMessage(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetErrorArgs(v []interface{}) *ListIsvPaymentPluginConfigsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetMaxResults(v int32) *ListIsvPaymentPluginConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetModule(v []*ListIsvPaymentPluginConfigsResponseBodyModule) *ListIsvPaymentPluginConfigsResponseBody {
	s.Module = v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetNextToken(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetRequestId(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetRootErrorCode(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetRootErrorMsg(v string) *ListIsvPaymentPluginConfigsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) SetSynchro(v bool) *ListIsvPaymentPluginConfigsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBody) Validate() error {
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

type ListIsvPaymentPluginConfigsResponseBodyModule struct {
	// Business ID
	//
	// example:
	//
	// WS20250915163734000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Extension information
	Extend map[string]*string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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
	// Primary key
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
	// jugaocai.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// User ID
	//
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListIsvPaymentPluginConfigsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListIsvPaymentPluginConfigsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetExtend() map[string]*string {
	return s.Extend
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetPluginDesc() *string {
	return s.PluginDesc
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetPluginName() *string {
	return s.PluginName
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetSiteName() *string {
	return s.SiteName
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetBizId(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetExtend(v map[string]*string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.Extend = v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetGmtCreateTime(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetGmtModifiedTime(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetId(v int64) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.Id = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetPluginConfig(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.PluginConfig = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetPluginDesc(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.PluginDesc = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetPluginId(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetPluginName(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.PluginName = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetSiteName(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.SiteName = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) SetUserId(v string) *ListIsvPaymentPluginConfigsResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ListIsvPaymentPluginConfigsResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
