// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIStaffConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAIStaffConversationResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAIStaffConversationResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAIStaffConversationResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAIStaffConversationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAIStaffConversationResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAIStaffConversationResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAIStaffConversationResponseBodyModule) *CreateAIStaffConversationResponseBody
	GetModule() *CreateAIStaffConversationResponseBodyModule
	SetRequestId(v string) *CreateAIStaffConversationResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAIStaffConversationResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAIStaffConversationResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAIStaffConversationResponseBody
	GetSynchro() *bool
}

type CreateAIStaffConversationResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                      `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CreateAIStaffConversationResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateAIStaffConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIStaffConversationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAIStaffConversationResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAIStaffConversationResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAIStaffConversationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAIStaffConversationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAIStaffConversationResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAIStaffConversationResponseBody) GetModule() *CreateAIStaffConversationResponseBodyModule {
	return s.Module
}

func (s *CreateAIStaffConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIStaffConversationResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAIStaffConversationResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAIStaffConversationResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAIStaffConversationResponseBody) SetAccessDeniedDetail(v string) *CreateAIStaffConversationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetAllowRetry(v bool) *CreateAIStaffConversationResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetAppName(v string) *CreateAIStaffConversationResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetDynamicCode(v string) *CreateAIStaffConversationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetDynamicMessage(v string) *CreateAIStaffConversationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetErrorArgs(v []interface{}) *CreateAIStaffConversationResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetModule(v *CreateAIStaffConversationResponseBodyModule) *CreateAIStaffConversationResponseBody {
	s.Module = v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetRequestId(v string) *CreateAIStaffConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetRootErrorCode(v string) *CreateAIStaffConversationResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetRootErrorMsg(v string) *CreateAIStaffConversationResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) SetSynchro(v bool) *CreateAIStaffConversationResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAIStaffConversationResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAIStaffConversationResponseBodyModule struct {
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 53467af9-8c4e-4498-9032-1f26978007f8
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// example:
	//
	// 8642d886-0322-43a9-b12f-6629b067978c
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 11
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// example:
	//
	// 865181640657408
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// Feel Like Makin\\" Love
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateAIStaffConversationResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAIStaffConversationResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAIStaffConversationResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *CreateAIStaffConversationResponseBodyModule) GetChatId() *string {
	return s.ChatId
}

func (s *CreateAIStaffConversationResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateAIStaffConversationResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *CreateAIStaffConversationResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *CreateAIStaffConversationResponseBodyModule) GetTitle() *string {
	return s.Title
}

func (s *CreateAIStaffConversationResponseBodyModule) SetBotId(v string) *CreateAIStaffConversationResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) SetChatId(v string) *CreateAIStaffConversationResponseBodyModule {
	s.ChatId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) SetConversationId(v string) *CreateAIStaffConversationResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) SetSectionId(v string) *CreateAIStaffConversationResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) SetSiteId(v string) *CreateAIStaffConversationResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) SetTitle(v string) *CreateAIStaffConversationResponseBodyModule {
	s.Title = &v
	return s
}

func (s *CreateAIStaffConversationResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
