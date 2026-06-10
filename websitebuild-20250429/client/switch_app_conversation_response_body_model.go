// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAppConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SwitchAppConversationResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *SwitchAppConversationResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SwitchAppConversationResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SwitchAppConversationResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SwitchAppConversationResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SwitchAppConversationResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *SwitchAppConversationResponseBodyModule) *SwitchAppConversationResponseBody
	GetModule() *SwitchAppConversationResponseBodyModule
	SetRequestId(v string) *SwitchAppConversationResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *SwitchAppConversationResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *SwitchAppConversationResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *SwitchAppConversationResponseBody
	GetSynchro() *bool
}

type SwitchAppConversationResponseBody struct {
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
	// watermark
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                            `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *SwitchAppConversationResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s SwitchAppConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchAppConversationResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchAppConversationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SwitchAppConversationResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SwitchAppConversationResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SwitchAppConversationResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SwitchAppConversationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SwitchAppConversationResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SwitchAppConversationResponseBody) GetModule() *SwitchAppConversationResponseBodyModule {
	return s.Module
}

func (s *SwitchAppConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchAppConversationResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *SwitchAppConversationResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *SwitchAppConversationResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SwitchAppConversationResponseBody) SetAccessDeniedDetail(v string) *SwitchAppConversationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetAllowRetry(v bool) *SwitchAppConversationResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetAppName(v string) *SwitchAppConversationResponseBody {
	s.AppName = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetDynamicCode(v string) *SwitchAppConversationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetDynamicMessage(v string) *SwitchAppConversationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetErrorArgs(v []interface{}) *SwitchAppConversationResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SwitchAppConversationResponseBody) SetModule(v *SwitchAppConversationResponseBodyModule) *SwitchAppConversationResponseBody {
	s.Module = v
	return s
}

func (s *SwitchAppConversationResponseBody) SetRequestId(v string) *SwitchAppConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetRootErrorCode(v string) *SwitchAppConversationResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetRootErrorMsg(v string) *SwitchAppConversationResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *SwitchAppConversationResponseBody) SetSynchro(v bool) *SwitchAppConversationResponseBody {
	s.Synchro = &v
	return s
}

func (s *SwitchAppConversationResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SwitchAppConversationResponseBodyModule struct {
	// example:
	//
	// 12343131221311
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 123
	ChatNum *int32 `json:"ChatNum,omitempty" xml:"ChatNum,omitempty"`
	// example:
	//
	// 8642d886-0322-43a9-b12f-6629b067978c
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {\\"appId\\":\\"WS20260504134737000001\\",\\"inputTokens\\":1395,\\"outputTokens\\":38}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
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
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SwitchAppConversationResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s SwitchAppConversationResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SwitchAppConversationResponseBodyModule) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *SwitchAppConversationResponseBodyModule) GetBotId() *string {
	return s.BotId
}

func (s *SwitchAppConversationResponseBodyModule) GetChatNum() *int32 {
	return s.ChatNum
}

func (s *SwitchAppConversationResponseBodyModule) GetConversationId() *string {
	return s.ConversationId
}

func (s *SwitchAppConversationResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *SwitchAppConversationResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *SwitchAppConversationResponseBodyModule) GetMetaData() *string {
	return s.MetaData
}

func (s *SwitchAppConversationResponseBodyModule) GetSectionId() *string {
	return s.SectionId
}

func (s *SwitchAppConversationResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *SwitchAppConversationResponseBodyModule) GetTitle() *string {
	return s.Title
}

func (s *SwitchAppConversationResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *SwitchAppConversationResponseBodyModule) SetAliyunPk(v string) *SwitchAppConversationResponseBodyModule {
	s.AliyunPk = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetBotId(v string) *SwitchAppConversationResponseBodyModule {
	s.BotId = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetChatNum(v int32) *SwitchAppConversationResponseBodyModule {
	s.ChatNum = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetConversationId(v string) *SwitchAppConversationResponseBodyModule {
	s.ConversationId = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetGmtCreate(v string) *SwitchAppConversationResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetGmtModified(v string) *SwitchAppConversationResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetMetaData(v string) *SwitchAppConversationResponseBodyModule {
	s.MetaData = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetSectionId(v string) *SwitchAppConversationResponseBodyModule {
	s.SectionId = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetSiteId(v string) *SwitchAppConversationResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetTitle(v string) *SwitchAppConversationResponseBodyModule {
	s.Title = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) SetUserId(v string) *SwitchAppConversationResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *SwitchAppConversationResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
