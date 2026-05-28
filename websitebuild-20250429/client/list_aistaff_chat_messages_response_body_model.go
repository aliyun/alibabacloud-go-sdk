// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAIStaffChatMessagesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAIStaffChatMessagesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAIStaffChatMessagesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAIStaffChatMessagesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAIStaffChatMessagesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAIStaffChatMessagesResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ListAIStaffChatMessagesResponseBodyModule) *ListAIStaffChatMessagesResponseBody
	GetModule() *ListAIStaffChatMessagesResponseBodyModule
	SetRequestId(v string) *ListAIStaffChatMessagesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAIStaffChatMessagesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAIStaffChatMessagesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAIStaffChatMessagesResponseBody
	GetSynchro() *bool
}

type ListAIStaffChatMessagesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                    `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                              `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *ListAIStaffChatMessagesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAIStaffChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatMessagesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAIStaffChatMessagesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAIStaffChatMessagesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAIStaffChatMessagesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAIStaffChatMessagesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAIStaffChatMessagesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAIStaffChatMessagesResponseBody) GetModule() *ListAIStaffChatMessagesResponseBodyModule {
	return s.Module
}

func (s *ListAIStaffChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIStaffChatMessagesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAIStaffChatMessagesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAIStaffChatMessagesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAIStaffChatMessagesResponseBody) SetAccessDeniedDetail(v string) *ListAIStaffChatMessagesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetAllowRetry(v bool) *ListAIStaffChatMessagesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetAppName(v string) *ListAIStaffChatMessagesResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetDynamicCode(v string) *ListAIStaffChatMessagesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetDynamicMessage(v string) *ListAIStaffChatMessagesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetErrorArgs(v []interface{}) *ListAIStaffChatMessagesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetModule(v *ListAIStaffChatMessagesResponseBodyModule) *ListAIStaffChatMessagesResponseBody {
	s.Module = v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetRequestId(v string) *ListAIStaffChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetRootErrorCode(v string) *ListAIStaffChatMessagesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetRootErrorMsg(v string) *ListAIStaffChatMessagesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) SetSynchro(v bool) *ListAIStaffChatMessagesResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAIStaffChatMessagesResponseBodyModule struct {
	Messages []*ListAIStaffChatMessagesResponseBodyModuleMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s ListAIStaffChatMessagesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatMessagesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatMessagesResponseBodyModule) GetMessages() []*ListAIStaffChatMessagesResponseBodyModuleMessages {
	return s.Messages
}

func (s *ListAIStaffChatMessagesResponseBodyModule) SetMessages(v []*ListAIStaffChatMessagesResponseBodyModuleMessages) *ListAIStaffChatMessagesResponseBodyModule {
	s.Messages = v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModule) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIStaffChatMessagesResponseBodyModuleMessages struct {
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 54a0bfa0-41bd-4e96-acd9-fb13c0474452
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// example:
	//
	// success
	ChatStatus *string `json:"ChatStatus,omitempty" xml:"ChatStatus,omitempty"`
	// example:
	//
	// domain cnamenwww.buyhao8.com www.buyhao8.com.a1.initrr.comn
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// application/octet-stream
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// job-675163021891846144
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1723532098
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1591339051000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1dafa033-e72b-44c2-99b7-bc202c5b6198
	MessageId *string                `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	MetaData  map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 324
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// example:
	//
	// 928636774795776
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// SINGLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIStaffChatMessagesResponseBodyModuleMessages) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatMessagesResponseBodyModuleMessages) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetBotId() *string {
	return s.BotId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetChatId() *string {
	return s.ChatId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetChatStatus() *string {
	return s.ChatStatus
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetContent() *string {
	return s.Content
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetContentType() *string {
	return s.ContentType
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetRole() *string {
	return s.Role
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetSectionId() *string {
	return s.SectionId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) GetType() *string {
	return s.Type
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetBotId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.BotId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetChatId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.ChatId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetChatStatus(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.ChatStatus = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetContent(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.Content = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetContentType(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.ContentType = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetConversationId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.ConversationId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetGmtCreate(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.GmtCreate = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetGmtModified(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.GmtModified = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetMessageId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.MessageId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetMetaData(v map[string]interface{}) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.MetaData = v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetRole(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.Role = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetSectionId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.SectionId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetSiteId(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.SiteId = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) SetType(v string) *ListAIStaffChatMessagesResponseBodyModuleMessages {
	s.Type = &v
	return s
}

func (s *ListAIStaffChatMessagesResponseBodyModuleMessages) Validate() error {
	return dara.Validate(s)
}
