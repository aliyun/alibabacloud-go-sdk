// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppChatMessageResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppChatMessageResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppChatMessageResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppChatMessageResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppChatMessageResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppChatMessageResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UpdateAppChatMessageResponseBodyModule) *UpdateAppChatMessageResponseBody
	GetModule() *UpdateAppChatMessageResponseBodyModule
	SetRequestId(v string) *UpdateAppChatMessageResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppChatMessageResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppChatMessageResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppChatMessageResponseBody
	GetSynchro() *bool
}

type UpdateAppChatMessageResponseBody struct {
	// Detailed reason for access denial.
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
	// App name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic code; not currently used. Please ignore.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message, used to replace the `%s` placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the provided request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// faulty parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The returned object.
	Module *UpdateAppChatMessageResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s UpdateAppChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppChatMessageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppChatMessageResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppChatMessageResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppChatMessageResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppChatMessageResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppChatMessageResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppChatMessageResponseBody) GetModule() *UpdateAppChatMessageResponseBodyModule {
	return s.Module
}

func (s *UpdateAppChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppChatMessageResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppChatMessageResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppChatMessageResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppChatMessageResponseBody) SetAccessDeniedDetail(v string) *UpdateAppChatMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetAllowRetry(v bool) *UpdateAppChatMessageResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetAppName(v string) *UpdateAppChatMessageResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetDynamicCode(v string) *UpdateAppChatMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetDynamicMessage(v string) *UpdateAppChatMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetErrorArgs(v []interface{}) *UpdateAppChatMessageResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetModule(v *UpdateAppChatMessageResponseBodyModule) *UpdateAppChatMessageResponseBody {
	s.Module = v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetRequestId(v string) *UpdateAppChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetRootErrorCode(v string) *UpdateAppChatMessageResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetRootErrorMsg(v string) *UpdateAppChatMessageResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) SetSynchro(v bool) *UpdateAppChatMessageResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppChatMessageResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppChatMessageResponseBodyModule struct {
	// Bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// The ID of the chat.
	//
	// example:
	//
	// 53467af9-8c4e-4498-9032-1f26978007f8
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// Current conversation status
	//
	// example:
	//
	// 123
	ChatStatus *string `json:"ChatStatus,omitempty" xml:"ChatStatus,omitempty"`
	// The ID of the data class API being invoked.
	//
	// example:
	//
	// {\\"code\\": \\"200\\", \\"success\\": True, \\"content\\": True, \\"requestId\\": \\"4AF53F7B-FEA9-5966-B0F8-BAF9A1EEFE34\\", \\"businessError\\": False, \\"message\\": \\"success\\", \\"httpStatusCode\\": 200}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Content type.
	//
	// example:
	//
	// image/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Session ID.
	//
	// example:
	//
	// 872be9bc-3097-433d-b462-596202455102
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Creation Time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 1740479834
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Message ID
	//
	// example:
	//
	// 471791769135220858
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Business extension metadata (in Map format, must be a JSON string)
	//
	// example:
	//
	// {\\"appId\\":\\"WS20260418211121000001\\",\\"inputTokens\\":273,\\"outputTokens\\":1}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The probability of not wearing a mask, not wearing a uniform, or not wearing a hat.
	//
	// example:
	//
	// 1
	No *int32 `json:"No,omitempty" xml:"No,omitempty"`
	// Indicates the role of a conversation participant. Valid values include:
	//
	// - user: User
	//
	// - assistant: Assistant
	//
	// - system: System
	//
	// - function: Function
	//
	// - plugin: Plugin
	//
	// - tool: Tool
	//
	// example:
	//
	// LoC
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Section ID of the checklist item.
	//
	// example:
	//
	// 11
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// Site ID, which can be obtained by invoking the [ListSites](~~ListSites~~) API.
	//
	// example:
	//
	// 865181640657408
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// File type
	//
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAppChatMessageResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppChatMessageResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UpdateAppChatMessageResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetChatId() *string {
	return s.ChatId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetChatStatus() *string {
	return s.ChatStatus
}

func (s *UpdateAppChatMessageResponseBodyModule) GetContent() *string {
	return s.Content
}

func (s *UpdateAppChatMessageResponseBodyModule) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateAppChatMessageResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *UpdateAppChatMessageResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *UpdateAppChatMessageResponseBodyModule) GetMessageId() *string {
	return s.MessageId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetMetaData() *string {
	return s.MetaData
}

func (s *UpdateAppChatMessageResponseBodyModule) GetNo() *int32 {
	return s.No
}

func (s *UpdateAppChatMessageResponseBodyModule) GetRole() *string {
	return s.Role
}

func (s *UpdateAppChatMessageResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *UpdateAppChatMessageResponseBodyModule) GetType() *string {
	return s.Type
}

func (s *UpdateAppChatMessageResponseBodyModule) SetBotId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetChatId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.ChatId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetChatStatus(v string) *UpdateAppChatMessageResponseBodyModule {
	s.ChatStatus = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetContent(v string) *UpdateAppChatMessageResponseBodyModule {
	s.Content = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetContentType(v string) *UpdateAppChatMessageResponseBodyModule {
	s.ContentType = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetConversationId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetGmtCreateTime(v string) *UpdateAppChatMessageResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetGmtModifiedTime(v string) *UpdateAppChatMessageResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetMessageId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.MessageId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetMetaData(v string) *UpdateAppChatMessageResponseBodyModule {
	s.MetaData = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetNo(v int32) *UpdateAppChatMessageResponseBodyModule {
	s.No = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetRole(v string) *UpdateAppChatMessageResponseBodyModule {
	s.Role = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetSectionId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetSiteId(v string) *UpdateAppChatMessageResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) SetType(v string) *UpdateAppChatMessageResponseBodyModule {
	s.Type = &v
	return s
}

func (s *UpdateAppChatMessageResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
