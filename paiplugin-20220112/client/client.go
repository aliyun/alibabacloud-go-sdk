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

type CreateSignatureRequest struct {
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) SetDescription(v string) *CreateSignatureRequest {
	s.Description = &v
	return s
}

func (s *CreateSignatureRequest) SetName(v string) *CreateSignatureRequest {
	s.Name = &v
	return s
}

type CreateSignatureResponseBody struct {
	// 返回数据
	Data *CreateSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s CreateSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBody) SetData(v *CreateSignatureResponseBodyData) *CreateSignatureResponseBody {
	s.Data = v
	return s
}

func (s *CreateSignatureResponseBody) SetErrorCode(v int32) *CreateSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSignatureResponseBody) SetErrorMessage(v string) *CreateSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

type CreateSignatureResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 签名Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名审核状态。取值：
	// - 0：审核中
	// - 1：审核通过
	// - 2：审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBodyData) SetCreatedTime(v string) *CreateSignatureResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetId(v string) *CreateSignatureResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetName(v string) *CreateSignatureResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetStatus(v int32) *CreateSignatureResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateSignatureResponseBodyData) SetUpdatedTime(v string) *CreateSignatureResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateSignatureResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponse) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponse) SetHeaders(v map[string]*string) *CreateSignatureResponse {
	s.Headers = v
	return s
}

func (s *CreateSignatureResponse) SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// 模板内容，请注意控制总字数在70个字以内，超出部分按长短信收费，按67个字为单位记一条短信，必须在结尾添加”回T退订“
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名Id
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 模板类型：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetName(v string) *CreateTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplateRequest) SetSignatureId(v string) *CreateTemplateRequest {
	s.SignatureId = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v int32) *CreateTemplateRequest {
	s.Type = &v
	return s
}

type CreateTemplateResponseBody struct {
	// 返回数据
	Data *CreateTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetData(v *CreateTemplateResponseBodyData) *CreateTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorCode(v int32) *CreateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorMessage(v string) *CreateTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

type CreateTemplateResponseBodyData struct {
	// 模板内容，长度:2-30
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Id UUId
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名Id
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBodyData) SetContent(v string) *CreateTemplateResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetCreatedTime(v string) *CreateTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetDescription(v string) *CreateTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetId(v string) *CreateTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetName(v string) *CreateTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetReason(v string) *CreateTemplateResponseBodyData {
	s.Reason = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetSignatureId(v string) *CreateTemplateResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetStatus(v int32) *CreateTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetTemplateCode(v string) *CreateTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetType(v int32) *CreateTemplateResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetUpdatedTime(v string) *CreateTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteSignatureResponseBody struct {
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DeleteSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponseBody) SetData(v string) *DeleteSignatureResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSignatureResponseBody) SetErrorCode(v int32) *DeleteSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSignatureResponseBody) SetErrorMessage(v string) *DeleteSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

type DeleteSignatureResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponse) SetHeaders(v map[string]*string) *DeleteSignatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteSignatureResponse) SetBody(v *DeleteSignatureResponseBody) *DeleteSignatureResponse {
	s.Body = v
	return s
}

type DeleteTemplateResponseBody struct {
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetData(v string) *DeleteTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorCode(v int32) *DeleteTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorMessage(v string) *DeleteTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type GetSignatureResponseBody struct {
	// 返回数据
	Data *GetSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s GetSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignatureResponseBody) SetData(v *GetSignatureResponseBodyData) *GetSignatureResponseBody {
	s.Data = v
	return s
}

func (s *GetSignatureResponseBody) SetErrorCode(v int32) *GetSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSignatureResponseBody) SetErrorMessage(v string) *GetSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

type GetSignatureResponseBodyData struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 签名Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核建议。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名审核状态。取值：
	// - 0：审核中
	// - 1：审核通过
	// - 2：审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSignatureResponseBodyData) SetCreatedTime(v string) *GetSignatureResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetDescription(v string) *GetSignatureResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetId(v string) *GetSignatureResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetName(v string) *GetSignatureResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetReason(v string) *GetSignatureResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetStatus(v int32) *GetSignatureResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetUpdatedTime(v string) *GetSignatureResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetSignatureResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponse) GoString() string {
	return s.String()
}

func (s *GetSignatureResponse) SetHeaders(v map[string]*string) *GetSignatureResponse {
	s.Headers = v
	return s
}

func (s *GetSignatureResponse) SetBody(v *GetSignatureResponseBody) *GetSignatureResponse {
	s.Body = v
	return s
}

type GetTemplateResponseBody struct {
	// 返回数据
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v int32) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorMessage(v string) *GetTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

type GetTemplateResponseBodyData struct {
	// 模板内容，长度:2-30
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Id UUId
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名Id
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyData) SetContent(v string) *GetTemplateResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetCreatedTime(v string) *GetTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDescription(v string) *GetTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetId(v string) *GetTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetName(v string) *GetTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetReason(v string) *GetTemplateResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetSignatureId(v string) *GetTemplateResponseBodyData {
	s.SignatureId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetStatus(v int32) *GetTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTemplateCode(v string) *GetTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetType(v int32) *GetTemplateResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetUpdatedTime(v string) *GetTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListSignaturesRequest struct {
	// 签名名称，使用%name%模糊匹配。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 审核状态。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSignaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesRequest) GoString() string {
	return s.String()
}

func (s *ListSignaturesRequest) SetName(v string) *ListSignaturesRequest {
	s.Name = &v
	return s
}

func (s *ListSignaturesRequest) SetPageNumber(v int32) *ListSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSignaturesRequest) SetPageSize(v int32) *ListSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSignaturesRequest) SetStatus(v int32) *ListSignaturesRequest {
	s.Status = &v
	return s
}

type ListSignaturesResponseBody struct {
	// 返回数据
	Data *ListSignaturesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListSignaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBody) SetData(v *ListSignaturesResponseBodyData) *ListSignaturesResponseBody {
	s.Data = v
	return s
}

func (s *ListSignaturesResponseBody) SetErrorCode(v int32) *ListSignaturesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSignaturesResponseBody) SetErrorMessage(v string) *ListSignaturesResponseBody {
	s.ErrorMessage = &v
	return s
}

type ListSignaturesResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 签名列表。
	Signatures []*ListSignaturesResponseBodyDataSignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// 签名数量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSignaturesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBodyData) SetPageNumber(v int32) *ListSignaturesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSignaturesResponseBodyData) SetPageSize(v int32) *ListSignaturesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSignaturesResponseBodyData) SetSignatures(v []*ListSignaturesResponseBodyDataSignatures) *ListSignaturesResponseBodyData {
	s.Signatures = v
	return s
}

func (s *ListSignaturesResponseBodyData) SetTotalCount(v int32) *ListSignaturesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSignaturesResponseBodyDataSignatures struct {
	// 创建时间 (UTC+8)。
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 签名Id。
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 签名审核状态。取值：
	// - 0：审核中
	// - 1：审核通过
	// - 2：审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)。
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListSignaturesResponseBodyDataSignatures) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponseBodyDataSignatures) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponseBodyDataSignatures) SetCreatedTime(v string) *ListSignaturesResponseBodyDataSignatures {
	s.CreatedTime = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetId(v string) *ListSignaturesResponseBodyDataSignatures {
	s.Id = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetName(v string) *ListSignaturesResponseBodyDataSignatures {
	s.Name = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetStatus(v int32) *ListSignaturesResponseBodyDataSignatures {
	s.Status = &v
	return s
}

func (s *ListSignaturesResponseBodyDataSignatures) SetUpdatedTime(v string) *ListSignaturesResponseBodyDataSignatures {
	s.UpdatedTime = &v
	return s
}

type ListSignaturesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSignaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSignaturesResponse) GoString() string {
	return s.String()
}

func (s *ListSignaturesResponse) SetHeaders(v map[string]*string) *ListSignaturesResponse {
	s.Headers = v
	return s
}

func (s *ListSignaturesResponse) SetBody(v *ListSignaturesResponseBody) *ListSignaturesResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// 内容类型过滤，使用%content%模糊匹配
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 模板名称过滤，使用%name%模糊匹配
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 审核状态过滤
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板类型过滤
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetContent(v string) *ListTemplatesRequest {
	s.Content = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetStatus(v int32) *ListTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplatesRequest) SetType(v int32) *ListTemplatesRequest {
	s.Type = &v
	return s
}

type ListTemplatesResponseBody struct {
	// 返回数据
	Data *ListTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetData(v *ListTemplatesResponseBodyData) *ListTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorCode(v int32) *ListTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorMessage(v string) *ListTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

type ListTemplatesResponseBodyData struct {
	// 分页数，从1开始，默认为1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 模板列表
	Templates []*ListTemplatesResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// 总模板数量
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyData) SetPageNumber(v int32) *ListTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBodyData) SetPageSize(v int32) *ListTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBodyData) SetTemplates(v []*ListTemplatesResponseBodyDataTemplates) *ListTemplatesResponseBodyData {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBodyData) SetTotalCount(v int32) *ListTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyDataTemplates struct {
	// 模板内容，长度:2-30
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Id UUId
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核意见。
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 签名Id
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板类型：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListTemplatesResponseBodyDataTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyDataTemplates) SetContent(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Content = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetCreatedTime(v string) *ListTemplatesResponseBodyDataTemplates {
	s.CreatedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetDescription(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetId(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Id = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetName(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetReason(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Reason = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetSignatureId(v string) *ListTemplatesResponseBodyDataTemplates {
	s.SignatureId = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetStatus(v int32) *ListTemplatesResponseBodyDataTemplates {
	s.Status = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetTemplateCode(v string) *ListTemplatesResponseBodyDataTemplates {
	s.TemplateCode = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetType(v int32) *ListTemplatesResponseBodyDataTemplates {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetUpdatedTime(v string) *ListTemplatesResponseBodyDataTemplates {
	s.UpdatedTime = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	// 人群ID，用于关联人群。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 外部拓展字段。
	OutIds []*string `json:"OutIds,omitempty" xml:"OutIds,omitempty" type:"Repeated"`
	// 手机号，每个手机号对应一个模板变量、上行拓展码和外部拓展字段。
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// 发送计划ID，用于关联发送计划。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，同时只能指定签名名称或签名ID其中之一。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 短信上行拓展码。
	SmsUpExtendCodes []*string `json:"SmsUpExtendCodes,omitempty" xml:"SmsUpExtendCodes,omitempty" type:"Repeated"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，同时只能指定模板Code或模板ID其中之一。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 短信模板变量对应的实际值，JSON格式。支持传入多个参数，示例：{"name":"张三","number":"15038****76"}。
	TemplateParams []*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty" type:"Repeated"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetGroupId(v string) *SendMessageRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageRequest) SetOutIds(v []*string) *SendMessageRequest {
	s.OutIds = v
	return s
}

func (s *SendMessageRequest) SetPhoneNumbers(v []*string) *SendMessageRequest {
	s.PhoneNumbers = v
	return s
}

func (s *SendMessageRequest) SetScheduleId(v string) *SendMessageRequest {
	s.ScheduleId = &v
	return s
}

func (s *SendMessageRequest) SetSignName(v string) *SendMessageRequest {
	s.SignName = &v
	return s
}

func (s *SendMessageRequest) SetSignatureId(v string) *SendMessageRequest {
	s.SignatureId = &v
	return s
}

func (s *SendMessageRequest) SetSmsUpExtendCodes(v []*string) *SendMessageRequest {
	s.SmsUpExtendCodes = v
	return s
}

func (s *SendMessageRequest) SetTemplateCode(v string) *SendMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageRequest) SetTemplateId(v string) *SendMessageRequest {
	s.TemplateId = &v
	return s
}

func (s *SendMessageRequest) SetTemplateParams(v []*string) *SendMessageRequest {
	s.TemplateParams = v
	return s
}

type SendMessageResponseBody struct {
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetData(v string) *SendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorCode(v int32) *SendMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorMessage(v string) *SendMessageResponseBody {
	s.ErrorMessage = &v
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiplugin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateSignature(request *CreateSignatureRequest) (_result *CreateSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSignatureResponse{}
	_body, _err := client.CreateSignatureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSignatureWithOptions(request *CreateSignatureRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSignature(Id *string) (_result *DeleteSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DeleteSignatureWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSignatureWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	Id = openapiutil.GetEncodeParam(Id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures/" + tea.StringValue(Id)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(Id *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	Id = openapiutil.GetEncodeParam(Id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates/" + tea.StringValue(Id)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignature(Id *string) (_result *GetSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSignatureResponse{}
	_body, _err := client.GetSignatureWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignatureWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSignatureResponse, _err error) {
	Id = openapiutil.GetEncodeParam(Id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSignature"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures/" + tea.StringValue(Id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(Id *string) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(Id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(Id *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	Id = openapiutil.GetEncodeParam(Id)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates/" + tea.StringValue(Id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSignatures(request *ListSignaturesRequest) (_result *ListSignaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSignaturesResponse{}
	_body, _err := client.ListSignaturesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSignaturesWithOptions(request *ListSignaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSignaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSignatures"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/signatures"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OutIds)) {
		body["OutIds"] = request.OutIds
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		body["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleId)) {
		body["ScheduleId"] = request.ScheduleId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		body["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCodes)) {
		body["SmsUpExtendCodes"] = request.SmsUpExtendCodes
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParams)) {
		body["TemplateParams"] = request.TemplateParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
