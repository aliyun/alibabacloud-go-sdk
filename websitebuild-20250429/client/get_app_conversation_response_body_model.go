// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppConversationResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppConversationResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppConversationResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppConversationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppConversationResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppConversationResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppConversationResponseBodyModule) *GetAppConversationResponseBody
	GetModule() *GetAppConversationResponseBodyModule
	SetRequestId(v string) *GetAppConversationResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppConversationResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppConversationResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppConversationResponseBody
	GetSynchro() *bool
}

type GetAppConversationResponseBody struct {
	// access denied details
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
	// App Name.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// response Data
	Module *GetAppConversationResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// is processed synchronously
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConversationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppConversationResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppConversationResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppConversationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppConversationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppConversationResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppConversationResponseBody) GetModule() *GetAppConversationResponseBodyModule {
	return s.Module
}

func (s *GetAppConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppConversationResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppConversationResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppConversationResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppConversationResponseBody) SetAccessDeniedDetail(v string) *GetAppConversationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppConversationResponseBody) SetAllowRetry(v bool) *GetAppConversationResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppConversationResponseBody) SetAppName(v string) *GetAppConversationResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppConversationResponseBody) SetDynamicCode(v string) *GetAppConversationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppConversationResponseBody) SetDynamicMessage(v string) *GetAppConversationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppConversationResponseBody) SetErrorArgs(v []interface{}) *GetAppConversationResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppConversationResponseBody) SetModule(v *GetAppConversationResponseBodyModule) *GetAppConversationResponseBody {
	s.Module = v
	return s
}

func (s *GetAppConversationResponseBody) SetRequestId(v string) *GetAppConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConversationResponseBody) SetRootErrorCode(v string) *GetAppConversationResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppConversationResponseBody) SetRootErrorMsg(v string) *GetAppConversationResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppConversationResponseBody) SetSynchro(v bool) *GetAppConversationResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppConversationResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppConversationResponseBodyModule struct {
	// User ID
	//
	// example:
	//
	// 12343131221311
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Chat record
	//
	// example:
	//
	// 123
	ChatNum *int32 `json:"ChatNum,omitempty" xml:"ChatNum,omitempty"`
	// session id
	//
	// example:
	//
	// 872be9bc-3097-433d-b462-596202455102
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Creation Time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// None
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Update Time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-10-29T10:34:13Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Extension information (in JSON string format)
	//
	// example:
	//
	// {\\"appId\\":\\"WS20260507200853000001\\",\\"inputTokens\\":1411,\\"outputTokens\\":51}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// Section ID of the checklist item.
	//
	// example:
	//
	// 11
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// site ID, which can be obtained by invoking the [ListSites](~~ListSites~~) API.
	//
	// example:
	//
	// 865181640657408
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Current point, equivalent to news title
	//
	// example:
	//
	// Feel Like Makin\\" Love
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// User ID
	//
	// example:
	//
	// 用户id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAppConversationResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppConversationResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppConversationResponseBodyModule) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *GetAppConversationResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *GetAppConversationResponseBodyModule) GetChatNum() *int32 {
	return s.ChatNum
}

func (s *GetAppConversationResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppConversationResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAppConversationResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAppConversationResponseBodyModule) GetMetaData() *string {
	return s.MetaData
}

func (s *GetAppConversationResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *GetAppConversationResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *GetAppConversationResponseBodyModule) GetTitle() *string {
	return s.Title
}

func (s *GetAppConversationResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *GetAppConversationResponseBodyModule) SetAliyunPk(v string) *GetAppConversationResponseBodyModule {
	s.AliyunPk = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetBotId(v string) *GetAppConversationResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetChatNum(v int32) *GetAppConversationResponseBodyModule {
	s.ChatNum = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetConversationId(v string) *GetAppConversationResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetGmtCreateTime(v string) *GetAppConversationResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetGmtModifiedTime(v string) *GetAppConversationResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetMetaData(v string) *GetAppConversationResponseBodyModule {
	s.MetaData = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetSectionId(v string) *GetAppConversationResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetSiteId(v string) *GetAppConversationResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetTitle(v string) *GetAppConversationResponseBodyModule {
	s.Title = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) SetUserId(v string) *GetAppConversationResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *GetAppConversationResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
