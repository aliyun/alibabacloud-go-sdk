// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppChatMessagesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppChatMessagesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppChatMessagesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppChatMessagesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppChatMessagesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppChatMessagesResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppChatMessagesResponseBody
	GetMaxResults() *int32
	SetModule(v []*ListAppChatMessagesResponseBodyModule) *ListAppChatMessagesResponseBody
	GetModule() []*ListAppChatMessagesResponseBodyModule
	SetNextToken(v string) *ListAppChatMessagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppChatMessagesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppChatMessagesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppChatMessagesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppChatMessagesResponseBody
	GetSynchro() *bool
}

type ListAppChatMessagesResponseBody struct {
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
	// App name.
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
	// dynamic error message used to replace `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, it indicates that the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Value range: 10 to 100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Response data
	Module []*ListAppChatMessagesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// Token for starting the next query. It is empty when there is no next query.
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
	// Error code
	//
	// example:
	//
	// SYSTEM.EROR
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

func (s ListAppChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppChatMessagesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppChatMessagesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppChatMessagesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppChatMessagesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppChatMessagesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppChatMessagesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppChatMessagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppChatMessagesResponseBody) GetModule() []*ListAppChatMessagesResponseBodyModule {
	return s.Module
}

func (s *ListAppChatMessagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppChatMessagesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppChatMessagesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppChatMessagesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppChatMessagesResponseBody) SetAccessDeniedDetail(v string) *ListAppChatMessagesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetAllowRetry(v bool) *ListAppChatMessagesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetAppName(v string) *ListAppChatMessagesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetDynamicCode(v string) *ListAppChatMessagesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetDynamicMessage(v string) *ListAppChatMessagesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetErrorArgs(v []interface{}) *ListAppChatMessagesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetMaxResults(v int32) *ListAppChatMessagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetModule(v []*ListAppChatMessagesResponseBodyModule) *ListAppChatMessagesResponseBody {
	s.Module = v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetNextToken(v string) *ListAppChatMessagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetRequestId(v string) *ListAppChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetRootErrorCode(v string) *ListAppChatMessagesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetRootErrorMsg(v string) *ListAppChatMessagesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) SetSynchro(v bool) *ListAppChatMessagesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppChatMessagesResponseBody) Validate() error {
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

type ListAppChatMessagesResponseBodyModule struct {
	// Bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Unique ID of a single message
	//
	// example:
	//
	// chat-xrz3etcl2bsygwlx8g
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// Current chat status
	//
	// example:
	//
	// created
	ChatStatus *string `json:"ChatStatus,omitempty" xml:"ChatStatus,omitempty"`
	// ID of the data class API being invoked.
	//
	// example:
	//
	// {\\"code\\": \\"200\\", \\"success\\": True, \\"content\\": True, \\"requestId\\": \\"028993DE-097E-5F4E-AC48-64A2D5ED5F30\\", \\"businessError\\": False, \\"message\\": \\"success\\", \\"httpStatusCode\\": 200}
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
	// Message ID
	//
	// example:
	//
	// ee60f5a6-88ca-4074-ad37-515f065bbbd2
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Business extension metadata (in Map format, must be a JSON string)
	//
	// example:
	//
	// {\\"appId\\":\\"WS20260506101154000001\\",\\"inputTokens\\":1148,\\"outputTokens\\":60}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// Region ordinal number. This parameter is deprecated.
	//
	// example:
	//
	// 1
	No *int32 `json:"No,omitempty" xml:"No,omitempty"`
	// Indicates the role of the participant in the conversation. Valid values include:
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

func (s ListAppChatMessagesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppChatMessagesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppChatMessagesResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *ListAppChatMessagesResponseBodyModule) GetChatId() *string {
	return s.ChatId
}

func (s *ListAppChatMessagesResponseBodyModule) GetChatStatus() *string {
	return s.ChatStatus
}

func (s *ListAppChatMessagesResponseBodyModule) GetContent() *string {
	return s.Content
}

func (s *ListAppChatMessagesResponseBodyModule) GetContentType() *string {
	return s.ContentType
}

func (s *ListAppChatMessagesResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAppChatMessagesResponseBodyModule) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAppChatMessagesResponseBodyModule) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAppChatMessagesResponseBodyModule) GetMessageId() *string {
	return s.MessageId
}

func (s *ListAppChatMessagesResponseBodyModule) GetMetaData() *string {
	return s.MetaData
}

func (s *ListAppChatMessagesResponseBodyModule) GetNo() *int32 {
	return s.No
}

func (s *ListAppChatMessagesResponseBodyModule) GetRole() *string {
	return s.Role
}

func (s *ListAppChatMessagesResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *ListAppChatMessagesResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAppChatMessagesResponseBodyModule) GetType() *string {
	return s.Type
}

func (s *ListAppChatMessagesResponseBodyModule) SetBotId(v string) *ListAppChatMessagesResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetChatId(v string) *ListAppChatMessagesResponseBodyModule {
	s.ChatId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetChatStatus(v string) *ListAppChatMessagesResponseBodyModule {
	s.ChatStatus = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetContent(v string) *ListAppChatMessagesResponseBodyModule {
	s.Content = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetContentType(v string) *ListAppChatMessagesResponseBodyModule {
	s.ContentType = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetConversationId(v string) *ListAppChatMessagesResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetGmtCreateTime(v string) *ListAppChatMessagesResponseBodyModule {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetGmtModifiedTime(v string) *ListAppChatMessagesResponseBodyModule {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetMessageId(v string) *ListAppChatMessagesResponseBodyModule {
	s.MessageId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetMetaData(v string) *ListAppChatMessagesResponseBodyModule {
	s.MetaData = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetNo(v int32) *ListAppChatMessagesResponseBodyModule {
	s.No = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetRole(v string) *ListAppChatMessagesResponseBodyModule {
	s.Role = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetSectionId(v string) *ListAppChatMessagesResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetSiteId(v string) *ListAppChatMessagesResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) SetType(v string) *ListAppChatMessagesResponseBodyModule {
	s.Type = &v
	return s
}

func (s *ListAppChatMessagesResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
