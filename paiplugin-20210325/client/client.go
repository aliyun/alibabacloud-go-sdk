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
	// 签名归属方的三证合一，OSS地址，必须以https开头，使用前需要授权
	Certificates *string `json:"Certificates,omitempty" xml:"Certificates,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 授权委托书(Power of attorney)， OSS地址，必须以https或oss开头，使用前需要授权，同上
	PowerOfAttorney *string `json:"PowerOfAttorney,omitempty" xml:"PowerOfAttorney,omitempty"`
	// 无需填写
	ProcessInstanceID *string `json:"ProcessInstanceID,omitempty" xml:"ProcessInstanceID,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) SetCertificates(v string) *CreateSignatureRequest {
	s.Certificates = &v
	return s
}

func (s *CreateSignatureRequest) SetDescription(v string) *CreateSignatureRequest {
	s.Description = &v
	return s
}

func (s *CreateSignatureRequest) SetName(v string) *CreateSignatureRequest {
	s.Name = &v
	return s
}

func (s *CreateSignatureRequest) SetPowerOfAttorney(v string) *CreateSignatureRequest {
	s.PowerOfAttorney = &v
	return s
}

func (s *CreateSignatureRequest) SetProcessInstanceID(v string) *CreateSignatureRequest {
	s.ProcessInstanceID = &v
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
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)
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

func (s *CreateSignatureResponseBodyData) SetID(v string) *CreateSignatureResponseBodyData {
	s.ID = &v
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

type CreateTemplateRequest struct {
	// 模板内容，请注意控制总字数在70个字以内，超出部分按长短信收费，按67个字为单位记一条短信，必须在结尾添加”回T退订“
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 无需填写
	ProcessInstanceID *string `json:"ProcessInstanceID,omitempty" xml:"ProcessInstanceID,omitempty"`
	// 签名ID
	SignatureID *string `json:"SignatureID,omitempty" xml:"SignatureID,omitempty"`
	// 模板类型：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	Type *int `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *CreateTemplateRequest) SetProcessInstanceID(v string) *CreateTemplateRequest {
	s.ProcessInstanceID = &v
	return s
}

func (s *CreateTemplateRequest) SetSignatureID(v string) *CreateTemplateRequest {
	s.SignatureID = &v
	return s
}

func (s *CreateTemplateRequest) SetType(v int) *CreateTemplateRequest {
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
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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

func (s *CreateTemplateResponseBodyData) SetID(v string) *CreateTemplateResponseBodyData {
	s.ID = &v
	return s
}

func (s *CreateTemplateResponseBodyData) SetName(v string) *CreateTemplateResponseBodyData {
	s.Name = &v
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

type ListTemplatesRequest struct {
	// 模板名称，用于名称过滤或搜索，使用%name%模糊匹配
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 审核状态过滤
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
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
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
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
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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

func (s *ListTemplatesResponseBodyDataTemplates) SetID(v string) *ListTemplatesResponseBodyDataTemplates {
	s.ID = &v
	return s
}

func (s *ListTemplatesResponseBodyDataTemplates) SetName(v string) *ListTemplatesResponseBodyDataTemplates {
	s.Name = &v
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

type DeleteScheduleResponseBody struct {
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DeleteScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponseBody) SetData(v string) *DeleteScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteScheduleResponseBody) SetErrorCode(v int32) *DeleteScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScheduleResponseBody) SetErrorMessage(v string) *DeleteScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

type DeleteScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponse) SetHeaders(v map[string]*string) *DeleteScheduleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleResponse) SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse {
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
	// 模板内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核结果说明
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板Code
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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

func (s *GetTemplateResponseBodyData) SetID(v string) *GetTemplateResponseBodyData {
	s.ID = &v
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

func (s *GetTemplateResponseBodyData) SetStatus(v int32) *GetTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTemplateCode(v string) *GetTemplateResponseBodyData {
	s.TemplateCode = &v
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
	// 签名名称，用于名称过滤或搜索，使用%name%模糊匹配
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 审核状态过滤
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
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 签名列表
	Signatures []*ListSignaturesResponseBodyDataSignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// 总签名数量
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
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)
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

func (s *ListSignaturesResponseBodyDataSignatures) SetID(v string) *ListSignaturesResponseBodyDataSignatures {
	s.ID = &v
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
	// 签名归属方的三证合一，OSS地址，必须以https开头，使用前需要授权
	Certificates *string `json:"Certificates,omitempty" xml:"Certificates,omitempty"`
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 申请说明
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID UUID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 签名名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 授权委托书(Power of attorney)， OSS地址，必须以https或oss开头，使用前需要授权，同上
	PowerOfAttorney *string `json:"PowerOfAttorney,omitempty" xml:"PowerOfAttorney,omitempty"`
	// 审核结果说明
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 审核状态
	// - 0 : 审核中
	// - 1 : 审核通过
	// - 2 : 审核不通过
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetSignatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSignatureResponseBodyData) SetCertificates(v string) *GetSignatureResponseBodyData {
	s.Certificates = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetCreatedTime(v string) *GetSignatureResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetDescription(v string) *GetSignatureResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetID(v string) *GetSignatureResponseBodyData {
	s.ID = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetName(v string) *GetSignatureResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSignatureResponseBodyData) SetPowerOfAttorney(v string) *GetSignatureResponseBodyData {
	s.PowerOfAttorney = &v
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

type CreateScheduleRequest struct {
	// 数据源地址
	// - 0: project/table
	// MaxCompute项目名和表名，使用前需要授权
	// - 1: oss地址 https://bucket.endpoint/path/to/file
	// OSS地址，必须以https开头，使用前需要授权，如 https://bucket.endpoint/path/to/file
	DataAddress *string `json:"DataAddress,omitempty" xml:"DataAddress,omitempty"`
	// 数据源类型
	// - 0: MaxCompute
	// - 1: CSV
	// 数据源类型为CSV时，可以提供不带header的CSV文件或带header的CSV文件
	// 不带header的CSV文件每行为一个手机号
	// 使用带header的CSV文件，如果不指定手机号列名，默认使用第一列为手机号
	// 其他列可用于替换模板中的形如${variable}的变量，实现个性化发送
	DataSource *int32 `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// 钉钉机器人关键词
	DingBotKeyword *string `json:"DingBotKeyword,omitempty" xml:"DingBotKeyword,omitempty"`
	// 钉钉机器人token
	DingBotToken *string `json:"DingBotToken,omitempty" xml:"DingBotToken,omitempty"`
	// 发送计划名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分区表达式
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// 手机号列名
	PhoneNumberColumn *string `json:"PhoneNumberColumn,omitempty" xml:"PhoneNumberColumn,omitempty"`
	// 发布时间 (UTC+8)，建议距现在10分钟后，不能距现在超过一年，否则会发生回绕，格式必须是2020-01-01 12:00:00
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// 签名ID
	SignatureID *string `json:"SignatureID,omitempty" xml:"SignatureID,omitempty"`
	// 模板号列名(可选)
	TemplateCodeColumn *string `json:"TemplateCodeColumn,omitempty" xml:"TemplateCodeColumn,omitempty"`
	// 模板ID
	TemplateID *string `json:"TemplateID,omitempty" xml:"TemplateID,omitempty"`
}

func (s CreateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleRequest) SetDataAddress(v string) *CreateScheduleRequest {
	s.DataAddress = &v
	return s
}

func (s *CreateScheduleRequest) SetDataSource(v int32) *CreateScheduleRequest {
	s.DataSource = &v
	return s
}

func (s *CreateScheduleRequest) SetDingBotKeyword(v string) *CreateScheduleRequest {
	s.DingBotKeyword = &v
	return s
}

func (s *CreateScheduleRequest) SetDingBotToken(v string) *CreateScheduleRequest {
	s.DingBotToken = &v
	return s
}

func (s *CreateScheduleRequest) SetName(v string) *CreateScheduleRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduleRequest) SetPartition(v string) *CreateScheduleRequest {
	s.Partition = &v
	return s
}

func (s *CreateScheduleRequest) SetPhoneNumberColumn(v string) *CreateScheduleRequest {
	s.PhoneNumberColumn = &v
	return s
}

func (s *CreateScheduleRequest) SetSendTime(v string) *CreateScheduleRequest {
	s.SendTime = &v
	return s
}

func (s *CreateScheduleRequest) SetSignatureID(v string) *CreateScheduleRequest {
	s.SignatureID = &v
	return s
}

func (s *CreateScheduleRequest) SetTemplateCodeColumn(v string) *CreateScheduleRequest {
	s.TemplateCodeColumn = &v
	return s
}

func (s *CreateScheduleRequest) SetTemplateID(v string) *CreateScheduleRequest {
	s.TemplateID = &v
	return s
}

type CreateScheduleResponseBody struct {
	// 返回数据
	Data *CreateScheduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s CreateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBody) SetData(v *CreateScheduleResponseBodyData) *CreateScheduleResponseBody {
	s.Data = v
	return s
}

func (s *CreateScheduleResponseBody) SetErrorCode(v int32) *CreateScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateScheduleResponseBody) SetErrorMessage(v string) *CreateScheduleResponseBody {
	s.ErrorMessage = &v
	return s
}

type CreateScheduleResponseBodyData struct {
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// ID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 发送计划名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 发布时间 (UTC+8)
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// 签名ID
	SignatureID *string `json:"SignatureID,omitempty" xml:"SignatureID,omitempty"`
	// 状态
	// - 0: 检查中
	// - 1: 检查成功
	// - 2: 检查失败
	// - 3: 发送中
	// - 4: 发送成功
	// - 5: 发送失败
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	TemplateID *string `json:"TemplateID,omitempty" xml:"TemplateID,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s CreateScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBodyData) SetCreatedTime(v string) *CreateScheduleResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetID(v string) *CreateScheduleResponseBodyData {
	s.ID = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetName(v string) *CreateScheduleResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetSendTime(v string) *CreateScheduleResponseBodyData {
	s.SendTime = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetSignatureID(v string) *CreateScheduleResponseBodyData {
	s.SignatureID = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetStatus(v int32) *CreateScheduleResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetTemplateID(v string) *CreateScheduleResponseBodyData {
	s.TemplateID = &v
	return s
}

func (s *CreateScheduleResponseBodyData) SetUpdatedTime(v string) *CreateScheduleResponseBodyData {
	s.UpdatedTime = &v
	return s
}

type CreateScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponse) SetHeaders(v map[string]*string) *CreateScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleResponse) SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse {
	s.Body = v
	return s
}

type ListSchedulesRequest struct {
	// 发送计划名称，用于名称过滤或搜索，使用%name%模糊匹配
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 发送状态过滤
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) SetName(v string) *ListSchedulesRequest {
	s.Name = &v
	return s
}

func (s *ListSchedulesRequest) SetPageNumber(v int32) *ListSchedulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSchedulesRequest) SetPageSize(v int32) *ListSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesRequest) SetStatus(v int32) *ListSchedulesRequest {
	s.Status = &v
	return s
}

type ListSchedulesResponseBody struct {
	// 返回数据
	Data *ListSchedulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) SetData(v *ListSchedulesResponseBodyData) *ListSchedulesResponseBody {
	s.Data = v
	return s
}

func (s *ListSchedulesResponseBody) SetErrorCode(v int32) *ListSchedulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSchedulesResponseBody) SetErrorMessage(v string) *ListSchedulesResponseBody {
	s.ErrorMessage = &v
	return s
}

type ListSchedulesResponseBodyData struct {
	// 分页数，从1开始，默认为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小，默认为10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 发送计划列表
	Schedules []*ListSchedulesResponseBodyDataSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	// 总发送计划数量
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchedulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodyData) SetPageNumber(v int32) *ListSchedulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSchedulesResponseBodyData) SetPageSize(v int32) *ListSchedulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSchedulesResponseBodyData) SetSchedules(v []*ListSchedulesResponseBodyDataSchedules) *ListSchedulesResponseBodyData {
	s.Schedules = v
	return s
}

func (s *ListSchedulesResponseBodyData) SetTotalCount(v int32) *ListSchedulesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSchedulesResponseBodyDataSchedules struct {
	// 创建时间 (UTC+8)
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// ID
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// 发送计划名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 发布时间 (UTC+8)
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// 签名ID
	SignatureID *string `json:"SignatureID,omitempty" xml:"SignatureID,omitempty"`
	// 状态
	// - 0: 检查中
	// - 1: 检查成功
	// - 2: 检查失败
	// - 3: 发送中
	// - 4: 发送成功
	// - 5: 发送失败
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	TemplateID *string `json:"TemplateID,omitempty" xml:"TemplateID,omitempty"`
	// 更新时间 (UTC+8)
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListSchedulesResponseBodyDataSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodyDataSchedules) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodyDataSchedules) SetCreatedTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.CreatedTime = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetID(v string) *ListSchedulesResponseBodyDataSchedules {
	s.ID = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetName(v string) *ListSchedulesResponseBodyDataSchedules {
	s.Name = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetSendTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.SendTime = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetSignatureID(v string) *ListSchedulesResponseBodyDataSchedules {
	s.SignatureID = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetStatus(v int32) *ListSchedulesResponseBodyDataSchedules {
	s.Status = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetTemplateID(v string) *ListSchedulesResponseBodyDataSchedules {
	s.TemplateID = &v
	return s
}

func (s *ListSchedulesResponseBodyDataSchedules) SetUpdatedTime(v string) *ListSchedulesResponseBodyDataSchedules {
	s.UpdatedTime = &v
	return s
}

type ListSchedulesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSchedulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponse) SetHeaders(v map[string]*string) *ListSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListSchedulesResponse) SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse {
	s.Body = v
	return s
}

type UploadMediaByURLRequest struct {
	UploadMetadatas []*UploadMediaByURLRequestUploadMetadatas `json:"UploadMetadatas,omitempty" xml:"UploadMetadatas,omitempty" type:"Repeated"`
	UploadURLs      *string                                   `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	UserData        *UploadMediaByURLRequestUserData          `json:"UserData,omitempty" xml:"UserData,omitempty" type:"Struct"`
}

func (s UploadMediaByURLRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequest) SetUploadMetadatas(v []*UploadMediaByURLRequestUploadMetadatas) *UploadMediaByURLRequest {
	s.UploadMetadatas = v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadURLs(v string) *UploadMediaByURLRequest {
	s.UploadURLs = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUserData(v *UploadMediaByURLRequestUserData) *UploadMediaByURLRequest {
	s.UserData = v
	return s
}

type UploadMediaByURLRequestUploadMetadatas struct {
	FileExtension *string                                             `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	S3UploadInfo  *UploadMediaByURLRequestUploadMetadatasS3UploadInfo `json:"S3UploadInfo,omitempty" xml:"S3UploadInfo,omitempty" type:"Struct"`
	SourceURL     *string                                             `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
	Title         *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UploadMediaByURLRequestUploadMetadatas) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLRequestUploadMetadatas) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequestUploadMetadatas) SetFileExtension(v string) *UploadMediaByURLRequestUploadMetadatas {
	s.FileExtension = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatas) SetS3UploadInfo(v *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) *UploadMediaByURLRequestUploadMetadatas {
	s.S3UploadInfo = v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatas) SetSourceURL(v string) *UploadMediaByURLRequestUploadMetadatas {
	s.SourceURL = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatas) SetTitle(v string) *UploadMediaByURLRequestUploadMetadatas {
	s.Title = &v
	return s
}

type UploadMediaByURLRequestUploadMetadatasS3UploadInfo struct {
	// 上传的临时AK
	S3AccessKey *string `json:"S3AccessKey,omitempty" xml:"S3AccessKey,omitempty"`
	// Bucket
	S3Bucket *string `json:"S3Bucket,omitempty" xml:"S3Bucket,omitempty"`
	// Endpoint
	S3Endpoint *string `json:"S3Endpoint,omitempty" xml:"S3Endpoint,omitempty"`
	// 上传的FileKey
	S3FileKey *string `json:"S3FileKey,omitempty" xml:"S3FileKey,omitempty"`
	// 供应商名称
	S3Provider *string `json:"S3Provider,omitempty" xml:"S3Provider,omitempty"`
	// 上传的临时SK
	S3SecretKey *string `json:"S3SecretKey,omitempty" xml:"S3SecretKey,omitempty"`
	// 上传的临时Token
	S3Token *string `json:"S3Token,omitempty" xml:"S3Token,omitempty"`
	Id      *int    `json:"id,omitempty" xml:"id,omitempty"`
	// Job Id
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s UploadMediaByURLRequestUploadMetadatasS3UploadInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLRequestUploadMetadatasS3UploadInfo) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3AccessKey(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3AccessKey = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3Bucket(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3Bucket = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3Endpoint(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3Endpoint = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3FileKey(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3FileKey = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3Provider(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3Provider = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3SecretKey(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3SecretKey = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetS3Token(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.S3Token = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetId(v int) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.Id = &v
	return s
}

func (s *UploadMediaByURLRequestUploadMetadatasS3UploadInfo) SetJobId(v string) *UploadMediaByURLRequestUploadMetadatasS3UploadInfo {
	s.JobId = &v
	return s
}

type UploadMediaByURLRequestUserData struct {
	Extend          map[string]interface{} `json:"Extend,omitempty" xml:"Extend,omitempty"`
	MessageCallback *string                `json:"MessageCallback,omitempty" xml:"MessageCallback,omitempty"`
}

func (s UploadMediaByURLRequestUserData) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLRequestUserData) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequestUserData) SetExtend(v map[string]interface{}) *UploadMediaByURLRequestUserData {
	s.Extend = v
	return s
}

func (s *UploadMediaByURLRequestUserData) SetMessageCallback(v string) *UploadMediaByURLRequestUserData {
	s.MessageCallback = &v
	return s
}

type UploadMediaByURLResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadJobs []*UploadMediaByURLResponseBodyUploadJobs `json:"UploadJobs,omitempty" xml:"UploadJobs,omitempty" type:"Repeated"`
}

func (s UploadMediaByURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBody) SetRequestId(v string) *UploadMediaByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMediaByURLResponseBody) SetUploadJobs(v []*UploadMediaByURLResponseBodyUploadJobs) *UploadMediaByURLResponseBody {
	s.UploadJobs = v
	return s
}

type UploadMediaByURLResponseBodyUploadJobs struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SourceURL *string `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
}

func (s UploadMediaByURLResponseBodyUploadJobs) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponseBodyUploadJobs) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetJobId(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.JobId = &v
	return s
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetSourceURL(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.SourceURL = &v
	return s
}

type UploadMediaByURLResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadMediaByURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMediaByURLResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponse) SetHeaders(v map[string]*string) *UploadMediaByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadMediaByURLResponse) SetBody(v *UploadMediaByURLResponseBody) *UploadMediaByURLResponse {
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
	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		body["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PowerOfAttorney)) {
		body["PowerOfAttorney"] = request.PowerOfAttorney
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceID)) {
		body["ProcessInstanceID"] = request.ProcessInstanceID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSignature"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/signatures"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(ID *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(ID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(ID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	ID = openapiutil.GetEncodeParam(ID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteTemplate"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/templates/"+tea.StringValue(ID)), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceID)) {
		body["ProcessInstanceID"] = request.ProcessInstanceID
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureID)) {
		body["SignatureID"] = request.SignatureID
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTemplate"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/templates"), tea.String("json"), req, runtime)
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
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTemplates"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchedule(ID *string) (_result *DeleteScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DeleteScheduleWithOptions(ID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScheduleWithOptions(ID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteScheduleResponse, _err error) {
	ID = openapiutil.GetEncodeParam(ID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSchedule"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/schedules/"+tea.StringValue(ID)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(ID *string) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(ID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(ID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	ID = openapiutil.GetEncodeParam(ID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTemplate"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/templates/"+tea.StringValue(ID)), tea.String("json"), req, runtime)
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
	_result = &ListSignaturesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSignatures"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/signatures"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignature(ID *string) (_result *GetSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSignatureResponse{}
	_body, _err := client.GetSignatureWithOptions(ID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignatureWithOptions(ID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSignatureResponse, _err error) {
	ID = openapiutil.GetEncodeParam(ID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetSignatureResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSignature"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/signatures/"+tea.StringValue(ID)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedule(request *CreateScheduleRequest) (_result *CreateScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduleResponse{}
	_body, _err := client.CreateScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScheduleWithOptions(request *CreateScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataAddress)) {
		body["DataAddress"] = request.DataAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DingBotKeyword)) {
		body["DingBotKeyword"] = request.DingBotKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.DingBotToken)) {
		body["DingBotToken"] = request.DingBotToken
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		body["Partition"] = request.Partition
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberColumn)) {
		body["PhoneNumberColumn"] = request.PhoneNumberColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SendTime)) {
		body["SendTime"] = request.SendTime
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureID)) {
		body["SignatureID"] = request.SignatureID
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCodeColumn)) {
		body["TemplateCodeColumn"] = request.TemplateCodeColumn
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateID)) {
		body["TemplateID"] = request.TemplateID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateScheduleResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSchedule"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/schedules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSchedules(request *ListSchedulesRequest) (_result *ListSchedulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSchedulesResponse{}
	_body, _err := client.ListSchedulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
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
	_result = &ListSchedulesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSchedules"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/schedules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMediaByURL(request *UploadMediaByURLRequest) (_result *UploadMediaByURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.UploadMediaByURLWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMediaByURLWithOptions(request *UploadMediaByURLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadMediaByURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UploadMetadatas)) {
		body["UploadMetadatas"] = request.UploadMetadatas
	}

	if !tea.BoolValue(util.IsUnset(request.UploadURLs)) {
		body["UploadURLs"] = request.UploadURLs
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserData))) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.DoROARequest(tea.String("UploadMediaByURL"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/media/api/v1/video/upload"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSignature(ID *string) (_result *DeleteSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DeleteSignatureWithOptions(ID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSignatureWithOptions(ID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	ID = openapiutil.GetEncodeParam(ID)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSignature"), tea.String("2021-03-25"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/signatures/"+tea.StringValue(ID)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
