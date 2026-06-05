// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppConversationsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppConversationsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppConversationsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppConversationsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppConversationsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppConversationsResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppConversationsResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppConversationsResponseBodyModule) *ListAppConversationsResponseBody
	GetModule() *ListAppConversationsResponseBodyModule
	SetNextToken(v string) *ListAppConversationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppConversationsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppConversationsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppConversationsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppConversationsResponseBody
	GetSynchro() *bool
}

type ListAppConversationsResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Module     *ListAppConversationsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s ListAppConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppConversationsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppConversationsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppConversationsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppConversationsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppConversationsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppConversationsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppConversationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppConversationsResponseBody) GetModule() *ListAppConversationsResponseBodyModule {
	return s.Module
}

func (s *ListAppConversationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppConversationsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppConversationsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppConversationsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppConversationsResponseBody) SetAccessDeniedDetail(v string) *ListAppConversationsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetAllowRetry(v bool) *ListAppConversationsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetAppName(v string) *ListAppConversationsResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetDynamicCode(v string) *ListAppConversationsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetDynamicMessage(v string) *ListAppConversationsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetErrorArgs(v []interface{}) *ListAppConversationsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppConversationsResponseBody) SetMaxResults(v int32) *ListAppConversationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetModule(v *ListAppConversationsResponseBodyModule) *ListAppConversationsResponseBody {
	s.Module = v
	return s
}

func (s *ListAppConversationsResponseBody) SetNextToken(v string) *ListAppConversationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetRequestId(v string) *ListAppConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetRootErrorCode(v string) *ListAppConversationsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetRootErrorMsg(v string) *ListAppConversationsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppConversationsResponseBody) SetSynchro(v bool) *ListAppConversationsResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppConversationsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppConversationsResponseBodyModule struct {
	Data []*ListAppConversationsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppConversationsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppConversationsResponseBodyModule) GetData() []*ListAppConversationsResponseBodyModuleData {
	return s.Data
}

func (s *ListAppConversationsResponseBodyModule) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppConversationsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppConversationsResponseBodyModule) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppConversationsResponseBodyModule) SetData(v []*ListAppConversationsResponseBodyModuleData) *ListAppConversationsResponseBodyModule {
	s.Data = v
	return s
}

func (s *ListAppConversationsResponseBodyModule) SetPageNum(v int32) *ListAppConversationsResponseBodyModule {
	s.PageNum = &v
	return s
}

func (s *ListAppConversationsResponseBodyModule) SetPageSize(v int32) *ListAppConversationsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppConversationsResponseBodyModule) SetTotalCount(v int32) *ListAppConversationsResponseBodyModule {
	s.TotalCount = &v
	return s
}

func (s *ListAppConversationsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppConversationsResponseBodyModuleData struct {
	// aliyun_pk
	//
	// example:
	//
	// ***
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 1
	ChatNum *int32 `json:"ChatNum,omitempty" xml:"ChatNum,omitempty"`
	// example:
	//
	// 799EAC1246C855CAC75B77955E43D841
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-27T17:51:25.415+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-10-02T02:21:07Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// {\\"MD5\\":\\"296f6c01e7fea2697ffe1cf41082b774\\",\\"driver\\":\\"vhd\\",\\"flag\\":\\"12845825\\",\\"imds_support\\":\\"v1\\",\\"io_optimized\\":true,\\"nvme_supported\\":true,\\"uefi_preferred\\":false}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// 255
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	// example:
	//
	// 1068725896006128
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// Sometimes When We Touch
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAppConversationsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *ListAppConversationsResponseBodyModuleData) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListAppConversationsResponseBodyModuleData) GetBotId() *string {
	return s.BotId
}

func (s *ListAppConversationsResponseBodyModuleData) GetChatNum() *int32 {
	return s.ChatNum
}

func (s *ListAppConversationsResponseBodyModuleData) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAppConversationsResponseBodyModuleData) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAppConversationsResponseBodyModuleData) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAppConversationsResponseBodyModuleData) GetMetaData() *string {
	return s.MetaData
}

func (s *ListAppConversationsResponseBodyModuleData) GetSectionId() *string {
	return s.SectionId
}

func (s *ListAppConversationsResponseBodyModuleData) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAppConversationsResponseBodyModuleData) GetTitle() *string {
	return s.Title
}

func (s *ListAppConversationsResponseBodyModuleData) GetUserId() *string {
	return s.UserId
}

func (s *ListAppConversationsResponseBodyModuleData) SetAliyunPk(v string) *ListAppConversationsResponseBodyModuleData {
	s.AliyunPk = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetBotId(v string) *ListAppConversationsResponseBodyModuleData {
	s.BotId = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetChatNum(v int32) *ListAppConversationsResponseBodyModuleData {
	s.ChatNum = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetConversationId(v string) *ListAppConversationsResponseBodyModuleData {
	s.ConversationId = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetGmtCreateTime(v string) *ListAppConversationsResponseBodyModuleData {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetGmtModifiedTime(v string) *ListAppConversationsResponseBodyModuleData {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetMetaData(v string) *ListAppConversationsResponseBodyModuleData {
	s.MetaData = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetSectionId(v string) *ListAppConversationsResponseBodyModuleData {
	s.SectionId = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetSiteId(v string) *ListAppConversationsResponseBodyModuleData {
	s.SiteId = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetTitle(v string) *ListAppConversationsResponseBodyModuleData {
	s.Title = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) SetUserId(v string) *ListAppConversationsResponseBodyModuleData {
	s.UserId = &v
	return s
}

func (s *ListAppConversationsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}
