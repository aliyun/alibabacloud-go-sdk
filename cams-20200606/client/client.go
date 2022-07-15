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

type CheckChatappContactsRequest struct {
	// 通道类型
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// 需要查询的用户列表,单次调用最多查询10个。注意：用户号码必须加国家码
	Contacts []*string `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// ISV客户wabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 发送号码,所选择ChannelType下的Business账号，即在控制台上审核通过的Number
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s CheckChatappContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckChatappContactsRequest) GoString() string {
	return s.String()
}

func (s *CheckChatappContactsRequest) SetChannelType(v string) *CheckChatappContactsRequest {
	s.ChannelType = &v
	return s
}

func (s *CheckChatappContactsRequest) SetContacts(v []*string) *CheckChatappContactsRequest {
	s.Contacts = v
	return s
}

func (s *CheckChatappContactsRequest) SetCustWabaId(v string) *CheckChatappContactsRequest {
	s.CustWabaId = &v
	return s
}

func (s *CheckChatappContactsRequest) SetFrom(v string) *CheckChatappContactsRequest {
	s.From = &v
	return s
}

type CheckChatappContactsShrinkRequest struct {
	// 通道类型
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// 需要查询的用户列表,单次调用最多查询10个。注意：用户号码必须加国家码
	ContactsShrink *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	// ISV客户wabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 发送号码,所选择ChannelType下的Business账号，即在控制台上审核通过的Number
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s CheckChatappContactsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckChatappContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckChatappContactsShrinkRequest) SetChannelType(v string) *CheckChatappContactsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *CheckChatappContactsShrinkRequest) SetContactsShrink(v string) *CheckChatappContactsShrinkRequest {
	s.ContactsShrink = &v
	return s
}

func (s *CheckChatappContactsShrinkRequest) SetCustWabaId(v string) *CheckChatappContactsShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *CheckChatappContactsShrinkRequest) SetFrom(v string) *CheckChatappContactsShrinkRequest {
	s.From = &v
	return s
}

type CheckChatappContactsResponseBody struct {
	// 返回结果 OK 为正常
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*CheckChatappContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 提示信息，当返回异常时有值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckChatappContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckChatappContactsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckChatappContactsResponseBody) SetCode(v string) *CheckChatappContactsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckChatappContactsResponseBody) SetData(v []*CheckChatappContactsResponseBodyData) *CheckChatappContactsResponseBody {
	s.Data = v
	return s
}

func (s *CheckChatappContactsResponseBody) SetMessage(v string) *CheckChatappContactsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckChatappContactsResponseBody) SetRequestId(v string) *CheckChatappContactsResponseBody {
	s.RequestId = &v
	return s
}

type CheckChatappContactsResponseBodyData struct {
	// 号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 状态
	// 有效账号为"valid"，无法账号为"invalid"，查询失败返回"failed"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckChatappContactsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckChatappContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckChatappContactsResponseBodyData) SetPhoneNumber(v string) *CheckChatappContactsResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CheckChatappContactsResponseBodyData) SetStatus(v string) *CheckChatappContactsResponseBodyData {
	s.Status = &v
	return s
}

type CheckChatappContactsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckChatappContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckChatappContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckChatappContactsResponse) GoString() string {
	return s.String()
}

func (s *CheckChatappContactsResponse) SetHeaders(v map[string]*string) *CheckChatappContactsResponse {
	s.Headers = v
	return s
}

func (s *CheckChatappContactsResponse) SetStatusCode(v int32) *CheckChatappContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckChatappContactsResponse) SetBody(v *CheckChatappContactsResponseBody) *CheckChatappContactsResponse {
	s.Body = v
	return s
}

type CreateChatappTemplateRequest struct {
	// 模板分类
	Category   *string                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	Components []*CreateChatappTemplateRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// ISV客户WabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 变量，KEY-VALUE结构
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *CreateChatappTemplateRequest) SetComponents(v []*CreateChatappTemplateRequestComponents) *CreateChatappTemplateRequest {
	s.Components = v
	return s
}

func (s *CreateChatappTemplateRequest) SetCustWabaId(v string) *CreateChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateRequest) SetExample(v map[string]*string) *CreateChatappTemplateRequest {
	s.Example = v
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

func (s *CreateChatappTemplateRequest) SetTemplateType(v string) *CreateChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateChatappTemplateRequestComponents struct {
	// 按钮
	Buttons []*CreateChatappTemplateRequestComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// 描述，当Type为Header，且Format为IMGAGE/DOCUMENT/VIDEO 可以增加描述
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// 文件名称，当Type为Header，且Format为DOCUMENT时可以给文件指定名称
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 格式
	// TEXT-文本 IMGAGE-图片 DOCUMENT-文档 VIDEO-视频
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// 所发送消息的文本
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 组件类型
	// 值：BODY、HEADER、FOOTER 和 BUTTONS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 素材路径
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateChatappTemplateRequestComponents) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponents) SetButtons(v []*CreateChatappTemplateRequestComponentsButtons) *CreateChatappTemplateRequestComponents {
	s.Buttons = v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetCaption(v string) *CreateChatappTemplateRequestComponents {
	s.Caption = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFileName(v string) *CreateChatappTemplateRequestComponents {
	s.FileName = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetFormat(v string) *CreateChatappTemplateRequestComponents {
	s.Format = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetText(v string) *CreateChatappTemplateRequestComponents {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetType(v string) *CreateChatappTemplateRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponents) SetUrl(v string) *CreateChatappTemplateRequestComponents {
	s.Url = &v
	return s
}

type CreateChatappTemplateRequestComponentsButtons struct {
	// 号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 所发送消息的文本
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 按钮类型
	// PHONE_NUMBER（电话）,URL（网页按钮）和QUICK_REPLY（快速回复）
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 点击按钮后将访问的网址
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 网址类型 static-静态dynamic-动态
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s CreateChatappTemplateRequestComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateRequestComponentsButtons) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetPhoneNumber(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetText(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Text = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Type = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrl(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.Url = &v
	return s
}

func (s *CreateChatappTemplateRequestComponentsButtons) SetUrlType(v string) *CreateChatappTemplateRequestComponentsButtons {
	s.UrlType = &v
	return s
}

type CreateChatappTemplateShrinkRequest struct {
	// 模板分类
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// ISV客户WabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 变量，KEY-VALUE结构
	ExampleShrink *string `json:"Example,omitempty" xml:"Example,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateShrinkRequest) SetCategory(v string) *CreateChatappTemplateShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetComponentsShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetCustWabaId(v string) *CreateChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetExampleShrink(v string) *CreateChatappTemplateShrinkRequest {
	s.ExampleShrink = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetLanguage(v string) *CreateChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetName(v string) *CreateChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateChatappTemplateShrinkRequest) SetTemplateType(v string) *CreateChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type CreateChatappTemplateResponseBody struct {
	// 返回结果 OK 为正常
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateChatappTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s *CreateChatappTemplateResponseBody) SetData(v *CreateChatappTemplateResponseBodyData) *CreateChatappTemplateResponseBody {
	s.Data = v
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

type CreateChatappTemplateResponseBodyData struct {
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateChatappTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateChatappTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateCode(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateChatappTemplateResponseBodyData) SetTemplateName(v string) *CreateChatappTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

type CreateChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateChatappTemplateResponse) SetStatusCode(v int32) *CreateChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatappTemplateResponse) SetBody(v *CreateChatappTemplateResponseBody) *CreateChatappTemplateResponse {
	s.Body = v
	return s
}

type DeleteChatappTemplateRequest struct {
	// ISV客户wabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 模板编码
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteChatappTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateRequest) SetCustWabaId(v string) *DeleteChatappTemplateRequest {
	s.CustWabaId = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteChatappTemplateResponse) SetStatusCode(v int32) *DeleteChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatappTemplateResponse) SetBody(v *DeleteChatappTemplateResponseBody) *DeleteChatappTemplateResponse {
	s.Body = v
	return s
}

type GetChatappTemplateDetailRequest struct {
	// ISV客户WabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板分类
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailRequest) SetCustWabaId(v string) *GetChatappTemplateDetailRequest {
	s.CustWabaId = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetLanguage(v string) *GetChatappTemplateDetailRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailRequest) SetTemplateCode(v string) *GetChatappTemplateDetailRequest {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateDetailResponseBody struct {
	// 返回结果 OK 为正常
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// TemplateDetail
	Data    *GetChatappTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
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

func (s *GetChatappTemplateDetailResponseBody) SetData(v *GetChatappTemplateDetailResponseBodyData) *GetChatappTemplateDetailResponseBody {
	s.Data = v
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

type GetChatappTemplateDetailResponseBodyData struct {
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 模板分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 消息模板组件
	Components []*GetChatappTemplateDetailResponseBodyDataComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// 变量例子
	Example map[string]*string `json:"Example,omitempty" xml:"Example,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板编码
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyData) SetAuditStatus(v string) *GetChatappTemplateDetailResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetCategory(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetComponents(v []*GetChatappTemplateDetailResponseBodyDataComponents) *GetChatappTemplateDetailResponseBodyData {
	s.Components = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetExample(v map[string]*string) *GetChatappTemplateDetailResponseBodyData {
	s.Example = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetLanguage(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetName(v string) *GetChatappTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateDetailResponseBodyData {
	s.TemplateCode = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponents struct {
	// 仅适用于 BUTTONS 类型。
	// 与按钮相关的参数。
	Buttons []*GetChatappTemplateDetailResponseBodyDataComponentsButtons `json:"Buttons,omitempty" xml:"Buttons,omitempty" type:"Repeated"`
	// 描述，当Type为Header，且Format为IMGAGE/DOCUMENT/VIDEO 可以增加描述
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// 文件名称，当Type为Header，且Format为DOCUMENT时可以给文件指定名称
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 格式
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// 所发送消息的文本
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 组件类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 素材路径
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponents) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetButtons(v []*GetChatappTemplateDetailResponseBodyDataComponentsButtons) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Buttons = v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetCaption(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Caption = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFileName(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.FileName = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetFormat(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Format = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponents) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponents {
	s.Url = &v
	return s
}

type GetChatappTemplateDetailResponseBodyDataComponentsButtons struct {
	// 电话号码
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 所发送消息的文本
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 按钮类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 当按钮类型是
	// URL 时有效
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// WEB地址类型
	// static-静态
	// dynamic-动态
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) String() string {
	return tea.Prettify(s)
}

func (s GetChatappTemplateDetailResponseBodyDataComponentsButtons) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetPhoneNumber(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetText(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Text = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrl(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.Url = &v
	return s
}

func (s *GetChatappTemplateDetailResponseBodyDataComponentsButtons) SetUrlType(v string) *GetChatappTemplateDetailResponseBodyDataComponentsButtons {
	s.UrlType = &v
	return s
}

type GetChatappTemplateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetChatappTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetChatappTemplateDetailResponse) SetStatusCode(v int32) *GetChatappTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappTemplateDetailResponse) SetBody(v *GetChatappTemplateDetailResponseBody) *GetChatappTemplateDetailResponse {
	s.Body = v
	return s
}

type ListChatappTemplateRequest struct {
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// ISV客户WabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Page *ListChatappTemplateRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
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

func (s *ListChatappTemplateRequest) SetCustWabaId(v string) *ListChatappTemplateRequest {
	s.CustWabaId = &v
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

func (s *ListChatappTemplateRequest) SetPage(v *ListChatappTemplateRequestPage) *ListChatappTemplateRequest {
	s.Page = v
	return s
}

type ListChatappTemplateRequestPage struct {
	// 查询开始数
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// 每次查询返回的条数
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatappTemplateRequestPage) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequestPage) SetIndex(v int32) *ListChatappTemplateRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatappTemplateRequestPage) SetSize(v int32) *ListChatappTemplateRequestPage {
	s.Size = &v
	return s
}

type ListChatappTemplateShrinkRequest struct {
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// ISV客户WabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 模板名称
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListChatappTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateShrinkRequest) SetAuditStatus(v string) *ListChatappTemplateShrinkRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCustWabaId(v string) *ListChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetLanguage(v string) *ListChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetName(v string) *ListChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetPageShrink(v string) *ListChatappTemplateShrinkRequest {
	s.PageShrink = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListChatappTemplateResponse) SetStatusCode(v int32) *ListChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatappTemplateResponse) SetBody(v *ListChatappTemplateResponseBody) *ListChatappTemplateResponse {
	s.Body = v
	return s
}

type SendChatappMessageRequest struct {
	// 通道类型 whatsapp/viber/line
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// 消息内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// ISV客户wabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 回落消息内容
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// 回落策略ID，可在控制台创建策略并查看
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// 发送方
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 消息类型
	MessageType *string   `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Payload     []*string `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Repeated"`
	// 模板编码
	TemplateCode   *string            `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParams map[string]*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// 接收号码
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// 消息大类
	// template--模板
	// message--非模板
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageRequest) SetChannelType(v string) *SendChatappMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageRequest) SetContent(v string) *SendChatappMessageRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageRequest) SetCustWabaId(v string) *SendChatappMessageRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackContent(v string) *SendChatappMessageRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageRequest) SetFallBackId(v string) *SendChatappMessageRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageRequest) SetFrom(v string) *SendChatappMessageRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageRequest) SetLanguage(v string) *SendChatappMessageRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageRequest) SetMessageType(v string) *SendChatappMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageRequest) SetPayload(v []*string) *SendChatappMessageRequest {
	s.Payload = v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateCode(v string) *SendChatappMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageRequest) SetTemplateParams(v map[string]*string) *SendChatappMessageRequest {
	s.TemplateParams = v
	return s
}

func (s *SendChatappMessageRequest) SetTo(v string) *SendChatappMessageRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageRequest) SetType(v string) *SendChatappMessageRequest {
	s.Type = &v
	return s
}

type SendChatappMessageShrinkRequest struct {
	// 通道类型 whatsapp/viber/line
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// 消息内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// ISV客户wabaId
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// 回落消息内容
	FallBackContent *string `json:"FallBackContent,omitempty" xml:"FallBackContent,omitempty"`
	// 回落策略ID，可在控制台创建策略并查看
	FallBackId *string `json:"FallBackId,omitempty" xml:"FallBackId,omitempty"`
	// 发送方
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// 语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 消息类型
	MessageType   *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// 模板编码
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParamsShrink *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// 接收号码
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// 消息大类
	// template--模板
	// message--非模板
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendChatappMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatappMessageShrinkRequest) SetChannelType(v string) *SendChatappMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetContent(v string) *SendChatappMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetCustWabaId(v string) *SendChatappMessageShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackContent(v string) *SendChatappMessageShrinkRequest {
	s.FallBackContent = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFallBackId(v string) *SendChatappMessageShrinkRequest {
	s.FallBackId = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetFrom(v string) *SendChatappMessageShrinkRequest {
	s.From = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetLanguage(v string) *SendChatappMessageShrinkRequest {
	s.Language = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetMessageType(v string) *SendChatappMessageShrinkRequest {
	s.MessageType = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetPayloadShrink(v string) *SendChatappMessageShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateCode(v string) *SendChatappMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTemplateParamsShrink(v string) *SendChatappMessageShrinkRequest {
	s.TemplateParamsShrink = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetTo(v string) *SendChatappMessageShrinkRequest {
	s.To = &v
	return s
}

func (s *SendChatappMessageShrinkRequest) SetType(v string) *SendChatappMessageShrinkRequest {
	s.Type = &v
	return s
}

type SendChatappMessageResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 消息ID
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponseBody) SetCode(v string) *SendChatappMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessage(v string) *SendChatappMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessageId(v string) *SendChatappMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetRequestId(v string) *SendChatappMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendChatappMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendChatappMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendChatappMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendChatappMessageResponse) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponse) SetHeaders(v map[string]*string) *SendChatappMessageResponse {
	s.Headers = v
	return s
}

func (s *SendChatappMessageResponse) SetStatusCode(v int32) *SendChatappMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendChatappMessageResponse) SetBody(v *SendChatappMessageResponseBody) *SendChatappMessageResponse {
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

func (client *Client) CheckChatappContactsWithOptions(tmpReq *CheckChatappContactsRequest, runtime *util.RuntimeOptions) (_result *CheckChatappContactsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CheckChatappContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Contacts)) {
		request.ContactsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contacts, tea.String("Contacts"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.ContactsShrink)) {
		body["Contacts"] = request.ContactsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckChatappContacts"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckChatappContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckChatappContacts(request *CheckChatappContactsRequest) (_result *CheckChatappContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckChatappContactsResponse{}
	_body, _err := client.CheckChatappContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateChatappTemplateWithOptions(tmpReq *CreateChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Example)) {
		request.ExampleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Example, tea.String("Example"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		body["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.ExampleShrink)) {
		body["Example"] = request.ExampleShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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
	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
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
	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
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

func (client *Client) ListChatappTemplateWithOptions(tmpReq *ListChatappTemplateRequest, runtime *util.RuntimeOptions) (_result *ListChatappTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Page))) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Page), tea.String("Page"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageShrink)) {
		query["Page"] = request.PageShrink
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

func (client *Client) SendChatappMessageWithOptions(tmpReq *SendChatappMessageRequest, runtime *util.RuntimeOptions) (_result *SendChatappMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendChatappMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TemplateParams)) {
		request.TemplateParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateParams, tea.String("TemplateParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.CustWabaId)) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackContent)) {
		body["FallBackContent"] = request.FallBackContent
	}

	if !tea.BoolValue(util.IsUnset(request.FallBackId)) {
		body["FallBackId"] = request.FallBackId
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParamsShrink)) {
		body["TemplateParams"] = request.TemplateParamsShrink
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
		Action:      tea.String("SendChatappMessage"),
		Version:     tea.String("2020-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendChatappMessage(request *SendChatappMessageRequest) (_result *SendChatappMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.SendChatappMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
