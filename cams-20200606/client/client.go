// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckContactsRequest struct {
	ChannelType          *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Contacts             *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	From                 *string `json:"From,omitempty" xml:"From,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckContactsRequest) GoString() string {
	return s.String()
}

func (s *CheckContactsRequest) SetChannelType(v string) *CheckContactsRequest {
	s.ChannelType = &v
	return s
}

func (s *CheckContactsRequest) SetContacts(v string) *CheckContactsRequest {
	s.Contacts = &v
	return s
}

func (s *CheckContactsRequest) SetFrom(v string) *CheckContactsRequest {
	s.From = &v
	return s
}

func (s *CheckContactsRequest) SetOwnerId(v int64) *CheckContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckContactsRequest) SetResourceOwnerAccount(v string) *CheckContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckContactsRequest) SetResourceOwnerId(v int64) *CheckContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckContactsResponseBody struct {
	Contacts      []*CheckContactsResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CheckContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckContactsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckContactsResponseBody) SetContacts(v []*CheckContactsResponseBodyContacts) *CheckContactsResponseBody {
	s.Contacts = v
	return s
}

func (s *CheckContactsResponseBody) SetRequestId(v string) *CheckContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckContactsResponseBody) SetResultCode(v string) *CheckContactsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CheckContactsResponseBody) SetResultMessage(v string) *CheckContactsResponseBody {
	s.ResultMessage = &v
	return s
}

type CheckContactsResponseBodyContacts struct {
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckContactsResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s CheckContactsResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *CheckContactsResponseBodyContacts) SetPhoneNumber(v string) *CheckContactsResponseBodyContacts {
	s.PhoneNumber = &v
	return s
}

func (s *CheckContactsResponseBodyContacts) SetStatus(v string) *CheckContactsResponseBodyContacts {
	s.Status = &v
	return s
}

type CheckContactsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckContactsResponse) GoString() string {
	return s.String()
}

func (s *CheckContactsResponse) SetHeaders(v map[string]*string) *CheckContactsResponse {
	s.Headers = v
	return s
}

func (s *CheckContactsResponse) SetBody(v *CheckContactsResponseBody) *CheckContactsResponse {
	s.Body = v
	return s
}

type CreateChatappTemplateRequest struct {
	// 模板分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 消息模板组件
	// 值需要通过把json结构转成String的方式传入
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// 变量例子
	// 值需要通过把json结构转成String的方式传入
	Example *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequest) SetCategory(v string) *CreateChatappTemplateRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetComponents(v string) *CreateChatappTemplateRequest {
	s.Components = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetExample(v string) *CreateChatappTemplateRequest {
	s.Example = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetLanguage(v string) *CreateChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetName(v string) *CreateChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetOwnerId(v int64) *CreateChatappTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetResourceOwnerAccount(v string) *CreateChatappTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetResourceOwnerId(v int64) *CreateChatappTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetTemplateType(v string) *CreateChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateChatappTemplateResponseBody struct {
	// 返回结果 OK 为正常
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据结点
	// {"templateCode": "744c4b5c79c9432497a075bdfca36bf5"，"templateName": "hello_whatsapp"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 提示信息，当返回异常时有值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBody) SetCode(v string) *CreateChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetData(v string) *CreateChatappTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetMessage(v string) *CreateChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatappTemplateResponseBody) SetRequestId(v string) *CreateChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateChatappTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponse) SetHeaders(v map[string]*string) *CreateChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateChatappTemplateResponse) SetBody(v *CreateChatappTemplateResponseBody) *CreateChatappTemplateResponse {
	s.Body = v
	return s
}

type DeleteChatappTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 模板编码
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateRequest) SetOwnerId(v int64) *DeleteChatappTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetResourceOwnerAccount(v string) *DeleteChatappTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetResourceOwnerId(v int64) *DeleteChatappTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateCode(v string) *DeleteChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

type DeleteChatappTemplateResponseBody struct {
	// 返回结果 OK 为正常
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 提示信息，当返回异常时有值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponseBody) SetCode(v string) *DeleteChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetMessage(v string) *DeleteChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatappTemplateResponseBody) SetRequestId(v string) *DeleteChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChatappTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateResponse) SetHeaders(v map[string]*string) *DeleteChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatappTemplateResponse) SetBody(v *DeleteChatappTemplateResponseBody) *DeleteChatappTemplateResponse {
	s.Body = v
	return s
}

type GetChatappTemplateDetailRequest struct {
	// 语言
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 模板分类
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailRequest) SetLanguage(v string) *GetChatappTemplateDetailRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetOwnerId(v int64) *GetChatappTemplateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetResourceOwnerAccount(v string) *GetChatappTemplateDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetResourceOwnerId(v int64) *GetChatappTemplateDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateCode(v string) *GetChatappTemplateDetailRequest {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateDetailResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据对像
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappTemplateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBody) SetCode(v string) *GetChatappTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetData(v string) *GetChatappTemplateDetailResponseBody {
	s.Data = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetMessage(v string) *GetChatappTemplateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBody) SetRequestId(v string) *GetChatappTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetChatappTemplateDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChatappTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChatappTemplateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponse) SetHeaders(v map[string]*string) *GetChatappTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetBody(v *GetChatappTemplateDetailResponseBody) *GetChatappTemplateDetailResponse {
	s.Body = v
	return s
}

type ListChatappTemplateRequest struct {
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 分页字段
	Page                 *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequest) SetAuditStatus(v string) *ListChatappTemplateRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateRequest) SetLanguage(v string) *ListChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateRequest) SetName(v string) *ListChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateRequest) SetOwnerId(v int64) *ListChatappTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetPage(v string) *ListChatappTemplateRequest {
	s.Page = &v
	return s
}

func (s *ListChatappTemplateRequest) SetResourceOwnerAccount(v string) *ListChatappTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatappTemplateRequest) SetResourceOwnerId(v int64) *ListChatappTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListChatappTemplateResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 模板列表
	ListTemplate []*ListChatappTemplateResponseBodyListTemplate `json:"ListTemplate,omitempty" xml:"ListTemplate,omitempty" type:"Repeated"`
	Message      *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListChatappTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBody) SetCode(v string) *ListChatappTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetListTemplate(v []*ListChatappTemplateResponseBodyListTemplate) *ListChatappTemplateResponseBody {
	s.ListTemplate = v
	return s
}

func (s *ListChatappTemplateResponseBody) SetMessage(v string) *ListChatappTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatappTemplateResponseBody) SetRequestId(v string) *ListChatappTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ListChatappTemplateResponseBodyListTemplate struct {
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 模板分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板编码
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListChatappTemplateResponseBodyListTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponseBodyListTemplate) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetAuditStatus(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetCategory(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Category = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetLanguage(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateCode(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappTemplateResponseBodyListTemplate) SetTemplateName(v string) *ListChatappTemplateResponseBodyListTemplate {
	s.TemplateName = &v
	return s
}

type ListChatappTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChatappTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateResponse) SetHeaders(v map[string]*string) *ListChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListChatappTemplateResponse) SetBody(v *ListChatappTemplateResponseBody) *ListChatappTemplateResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	Caption              *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	ChannelType          *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	From                 *string `json:"From,omitempty" xml:"From,omitempty"`
	Link                 *string `json:"Link,omitempty" xml:"Link,omitempty"`
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateBodyParams   *string `json:"TemplateBodyParams,omitempty" xml:"TemplateBodyParams,omitempty"`
	TemplateButtonParams *string `json:"TemplateButtonParams,omitempty" xml:"TemplateButtonParams,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateHeaderParams *string `json:"TemplateHeaderParams,omitempty" xml:"TemplateHeaderParams,omitempty"`
	Text                 *string `json:"Text,omitempty" xml:"Text,omitempty"`
	To                   *string `json:"To,omitempty" xml:"To,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetCaption(v string) *SendMessageRequest {
	s.Caption = &v
	return s
}

func (s *SendMessageRequest) SetChannelType(v string) *SendMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendMessageRequest) SetFileName(v string) *SendMessageRequest {
	s.FileName = &v
	return s
}

func (s *SendMessageRequest) SetFrom(v string) *SendMessageRequest {
	s.From = &v
	return s
}

func (s *SendMessageRequest) SetLink(v string) *SendMessageRequest {
	s.Link = &v
	return s
}

func (s *SendMessageRequest) SetMessageType(v string) *SendMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendMessageRequest) SetOwnerId(v int64) *SendMessageRequest {
	s.OwnerId = &v
	return s
}

func (s *SendMessageRequest) SetResourceOwnerAccount(v string) *SendMessageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendMessageRequest) SetResourceOwnerId(v int64) *SendMessageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendMessageRequest) SetTemplateBodyParams(v string) *SendMessageRequest {
	s.TemplateBodyParams = &v
	return s
}

func (s *SendMessageRequest) SetTemplateButtonParams(v string) *SendMessageRequest {
	s.TemplateButtonParams = &v
	return s
}

func (s *SendMessageRequest) SetTemplateCode(v string) *SendMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageRequest) SetTemplateHeaderParams(v string) *SendMessageRequest {
	s.TemplateHeaderParams = &v
	return s
}

func (s *SendMessageRequest) SetText(v string) *SendMessageRequest {
	s.Text = &v
	return s
}

func (s *SendMessageRequest) SetTo(v string) *SendMessageRequest {
	s.To = &v
	return s
}

func (s *SendMessageRequest) SetType(v string) *SendMessageRequest {
	s.Type = &v
	return s
}

type SendMessageResponseBody struct {
	Module        *SendMessageResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId     *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                        `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetModule(v *SendMessageResponseBodyModule) *SendMessageResponseBody {
	s.Module = v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) SetResultCode(v string) *SendMessageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *SendMessageResponseBody) SetResultMessage(v string) *SendMessageResponseBody {
	s.ResultMessage = &v
	return s
}

type SendMessageResponseBodyModule struct {
	FromId    *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	ToId      *string `json:"ToId,omitempty" xml:"ToId,omitempty"`
}

func (s SendMessageResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBodyModule) SetFromId(v string) *SendMessageResponseBodyModule {
	s.FromId = &v
	return s
}

func (s *SendMessageResponseBodyModule) SetMessageId(v string) *SendMessageResponseBodyModule {
	s.MessageId = &v
	return s
}

func (s *SendMessageResponseBodyModule) SetToId(v string) *SendMessageResponseBodyModule {
	s.ToId = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cams"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckContactsWithOptions(request *CheckContactsRequest, runtime *util.RuntimeOptions) (_result *CheckContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.Contacts)) {
		body["Contacts"] = request.Contacts
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckContacts"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckContacts(request *CheckContactsRequest) (_result *CheckContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckContactsResponse{}
	_body, _err := client.CheckContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateChatappTemplateWithOptions(request *CreateChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.Example)) {
		query["Example"] = request.Example
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateChatappTemplate(request *CreateChatappTemplateRequest) (_result *CreateChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CreateChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteChatappTemplateWithOptions(request *DeleteChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteChatappTemplate(request *DeleteChatappTemplateRequest) (_result *DeleteChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.DeleteChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChatappTemplateDetailWithOptions(request *GetChatappTemplateDetailRequest, runtime *util.RuntimeOptions) (_result *GetChatappTemplateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatappTemplateDetail"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChatappTemplateDetail(request *GetChatappTemplateDetailRequest) (_result *GetChatappTemplateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.GetChatappTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChatappTemplateWithOptions(request *ListChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *ListChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChatappTemplate"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChatappTemplate(request *ListChatappTemplateRequest) (_result *ListChatappTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.ListChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caption)) {
		body["Caption"] = request.Caption
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Link)) {
		body["Link"] = request.Link
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBodyParams)) {
		body["TemplateBodyParams"] = request.TemplateBodyParams
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateButtonParams)) {
		body["TemplateButtonParams"] = request.TemplateButtonParams
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateHeaderParams)) {
		body["TemplateHeaderParams"] = request.TemplateHeaderParams
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
