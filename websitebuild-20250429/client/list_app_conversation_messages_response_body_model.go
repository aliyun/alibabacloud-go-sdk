// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppConversationMessagesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppConversationMessagesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppConversationMessagesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppConversationMessagesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppConversationMessagesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppConversationMessagesResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppConversationMessagesResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListAppConversationMessagesResponseBodyModule) *ListAppConversationMessagesResponseBody
	GetModule() []*ListAppConversationMessagesResponseBodyModule
	SetNextToken(v string) *ListAppConversationMessagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppConversationMessagesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppConversationMessagesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppConversationMessagesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppConversationMessagesResponseBody
	GetSynchro() *bool
}

type ListAppConversationMessagesResponseBody struct {
	// permission denied information
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
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic error message used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// faulty parameter(s).
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Value range: 10–100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Response Data
	Module []*ListAppConversationMessagesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// Token for the start of the next query. It is empty if there is no next query.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
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
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppConversationMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppConversationMessagesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppConversationMessagesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppConversationMessagesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppConversationMessagesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppConversationMessagesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppConversationMessagesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppConversationMessagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppConversationMessagesResponseBody) GetModule() []*ListAppConversationMessagesResponseBodyModule {
	return s.Module
}

func (s *ListAppConversationMessagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppConversationMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppConversationMessagesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppConversationMessagesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppConversationMessagesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppConversationMessagesResponseBody) SetAccessDeniedDetail(v string) *ListAppConversationMessagesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetAllowRetry(v bool) *ListAppConversationMessagesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetAppName(v string) *ListAppConversationMessagesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetDynamicCode(v string) *ListAppConversationMessagesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetDynamicMessage(v string) *ListAppConversationMessagesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetErrorArgs(v []interface{}) *ListAppConversationMessagesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetMaxResults(v int32) *ListAppConversationMessagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetModule(v []*ListAppConversationMessagesResponseBodyModule) *ListAppConversationMessagesResponseBody {
	s.Module = v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetNextToken(v string) *ListAppConversationMessagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetRequestId(v string) *ListAppConversationMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetRootErrorCode(v string) *ListAppConversationMessagesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetRootErrorMsg(v string) *ListAppConversationMessagesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) SetSynchro(v bool) *ListAppConversationMessagesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppConversationMessagesResponseBody) Validate() error {
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

type ListAppConversationMessagesResponseBodyModule struct {
	// Bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Chat ID.
	//
	// example:
	//
	// chat-xrz3etcl2bsygwlx8g
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// Current chat status.
	//
	// example:
	//
	// created
	ChatStatus *string `json:"ChatStatus,omitempty" xml:"ChatStatus,omitempty"`
	// The ID of the data class API invoked.
	//
	// example:
	//
	// {\\"DeviceName\\": u\\"\\u667a\\u80fd\\u63d2\\u5ea716A\\", \\"BlidMac\\": \\"18bc5a53351c\\", \\"DeviceModel\\": \\"ABSP21-16M\\", \\"DeviceCategory\\": u\\"\\u63d2\\u5ea7\\", \\"DeviceManufacturer\\": \\"\\", \\"DeviceType\\": 140}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Content type
	//
	// example:
	//
	// image/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Session ID
	//
	// example:
	//
	// 872be9bc-3097-433d-b462-596202455102
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Message ID.
	//
	// example:
	//
	// 5f9e88ce-f223-4447-a23e-cb574df6c97a
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Business extension metadata (in Map format, must be a JSON string).
	//
	// example:
	//
	// {\\"appId\\":\\"WS20260507200853000001\\",\\"inputTokens\\":1411,\\"outputTokens\\":51}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// Region ordinal number. This parameter is deprecated.
	//
	// example:
	//
	// 1
	No *int32 `json:"No,omitempty" xml:"No,omitempty"`
	// Role of the conversation participant. Valid values include:
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
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Section ID of the inspection item.
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

func (s ListAppConversationMessagesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationMessagesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppConversationMessagesResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetChatId() *string {
	return s.ChatId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetChatStatus() *string {
	return s.ChatStatus
}

func (s *ListAppConversationMessagesResponseBodyModule) GetContent() *string {
	return s.Content
}

func (s *ListAppConversationMessagesResponseBodyModule) GetContentType() *string {
	return s.ContentType
}

func (s *ListAppConversationMessagesResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppConversationMessagesResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppConversationMessagesResponseBodyModule) GetMessageId() *string {
	return s.MessageId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetMetaData() *string {
	return s.MetaData
}

func (s *ListAppConversationMessagesResponseBodyModule) GetNo() *int32 {
	return s.No
}

func (s *ListAppConversationMessagesResponseBodyModule) GetRole() *string {
	return s.Role
}

func (s *ListAppConversationMessagesResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAppConversationMessagesResponseBodyModule) GetType() *string {
	return s.Type
}

func (s *ListAppConversationMessagesResponseBodyModule) SetBotId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetChatId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.ChatId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetChatStatus(v string) *ListAppConversationMessagesResponseBodyModule {
	s.ChatStatus = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetContent(v string) *ListAppConversationMessagesResponseBodyModule {
	s.Content = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetContentType(v string) *ListAppConversationMessagesResponseBodyModule {
	s.ContentType = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetConversationId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetGmtCreate(v string) *ListAppConversationMessagesResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetGmtModified(v string) *ListAppConversationMessagesResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetMessageId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.MessageId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetMetaData(v string) *ListAppConversationMessagesResponseBodyModule {
	s.MetaData = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetNo(v int32) *ListAppConversationMessagesResponseBodyModule {
	s.No = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetRole(v string) *ListAppConversationMessagesResponseBodyModule {
	s.Role = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetSectionId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetSiteId(v string) *ListAppConversationMessagesResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) SetType(v string) *ListAppConversationMessagesResponseBodyModule {
	s.Type = &v
	return s
}

func (s *ListAppConversationMessagesResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
