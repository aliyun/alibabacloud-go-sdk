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

type AddShortUrlRequest struct {
	EffectiveDays        *string `json:"EffectiveDays,omitempty" xml:"EffectiveDays,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ShortUrlName         *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	SourceUrl            *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s AddShortUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShortUrlRequest) GoString() string {
	return s.String()
}

func (s *AddShortUrlRequest) SetEffectiveDays(v string) *AddShortUrlRequest {
	s.EffectiveDays = &v
	return s
}

func (s *AddShortUrlRequest) SetOwnerId(v int64) *AddShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *AddShortUrlRequest) SetResourceOwnerAccount(v string) *AddShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddShortUrlRequest) SetResourceOwnerId(v int64) *AddShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddShortUrlRequest) SetShortUrlName(v string) *AddShortUrlRequest {
	s.ShortUrlName = &v
	return s
}

func (s *AddShortUrlRequest) SetSourceUrl(v string) *AddShortUrlRequest {
	s.SourceUrl = &v
	return s
}

type AddShortUrlResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AddShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddShortUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponseBody) SetCode(v string) *AddShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *AddShortUrlResponseBody) SetData(v *AddShortUrlResponseBodyData) *AddShortUrlResponseBody {
	s.Data = v
	return s
}

func (s *AddShortUrlResponseBody) SetMessage(v string) *AddShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *AddShortUrlResponseBody) SetRequestId(v string) *AddShortUrlResponseBody {
	s.RequestId = &v
	return s
}

type AddShortUrlResponseBodyData struct {
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	ShortUrl   *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	SourceUrl  *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s AddShortUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddShortUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponseBodyData) SetExpireDate(v string) *AddShortUrlResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *AddShortUrlResponseBodyData) SetShortUrl(v string) *AddShortUrlResponseBodyData {
	s.ShortUrl = &v
	return s
}

func (s *AddShortUrlResponseBodyData) SetSourceUrl(v string) *AddShortUrlResponseBodyData {
	s.SourceUrl = &v
	return s
}

type AddShortUrlResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddShortUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShortUrlResponse) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponse) SetHeaders(v map[string]*string) *AddShortUrlResponse {
	s.Headers = v
	return s
}

func (s *AddShortUrlResponse) SetBody(v *AddShortUrlResponseBody) *AddShortUrlResponse {
	s.Body = v
	return s
}

type AddSmsSignRequest struct {
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignFileList         []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	SignName             *string                          `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignSource           *int32                           `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
}

func (s AddSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSmsSignRequest) GoString() string {
	return s.String()
}

func (s *AddSmsSignRequest) SetOwnerId(v int64) *AddSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSmsSignRequest) SetRemark(v string) *AddSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *AddSmsSignRequest) SetResourceOwnerAccount(v string) *AddSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsSignRequest) SetResourceOwnerId(v int64) *AddSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmsSignRequest) SetSignFileList(v []*AddSmsSignRequestSignFileList) *AddSmsSignRequest {
	s.SignFileList = v
	return s
}

func (s *AddSmsSignRequest) SetSignName(v string) *AddSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *AddSmsSignRequest) SetSignSource(v int32) *AddSmsSignRequest {
	s.SignSource = &v
	return s
}

type AddSmsSignRequestSignFileList struct {
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	FileSuffix   *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func (s AddSmsSignRequestSignFileList) String() string {
	return tea.Prettify(s)
}

func (s AddSmsSignRequestSignFileList) GoString() string {
	return s.String()
}

func (s *AddSmsSignRequestSignFileList) SetFileContents(v string) *AddSmsSignRequestSignFileList {
	s.FileContents = &v
	return s
}

func (s *AddSmsSignRequestSignFileList) SetFileSuffix(v string) *AddSmsSignRequestSignFileList {
	s.FileSuffix = &v
	return s
}

type AddSmsSignResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignName  *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AddSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmsSignResponseBody) SetCode(v string) *AddSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *AddSmsSignResponseBody) SetMessage(v string) *AddSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *AddSmsSignResponseBody) SetRequestId(v string) *AddSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmsSignResponseBody) SetSignName(v string) *AddSmsSignResponseBody {
	s.SignName = &v
	return s
}

type AddSmsSignResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSmsSignResponse) GoString() string {
	return s.String()
}

func (s *AddSmsSignResponse) SetHeaders(v map[string]*string) *AddSmsSignResponse {
	s.Headers = v
	return s
}

func (s *AddSmsSignResponse) SetBody(v *AddSmsSignResponseBody) *AddSmsSignResponse {
	s.Body = v
	return s
}

type AddSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateContent      *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateRequest) SetOwnerId(v int64) *AddSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSmsTemplateRequest) SetRemark(v string) *AddSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *AddSmsTemplateRequest) SetResourceOwnerAccount(v string) *AddSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsTemplateRequest) SetResourceOwnerId(v int64) *AddSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateContent(v string) *AddSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateName(v string) *AddSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateType(v int32) *AddSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

type AddSmsTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s AddSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateResponseBody) SetCode(v string) *AddSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetMessage(v string) *AddSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetRequestId(v string) *AddSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSmsTemplateResponseBody) SetTemplateCode(v string) *AddSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

type AddSmsTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateResponse) SetHeaders(v map[string]*string) *AddSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddSmsTemplateResponse) SetBody(v *AddSmsTemplateResponseBody) *AddSmsTemplateResponse {
	s.Body = v
	return s
}

type CreateCardSmsTemplateRequest struct {
	Memo         *string                `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Template     map[string]interface{} `json:"Template,omitempty" xml:"Template,omitempty"`
	TemplateName *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateCardSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCardSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateRequest) SetMemo(v string) *CreateCardSmsTemplateRequest {
	s.Memo = &v
	return s
}

func (s *CreateCardSmsTemplateRequest) SetTemplate(v map[string]interface{}) *CreateCardSmsTemplateRequest {
	s.Template = v
	return s
}

func (s *CreateCardSmsTemplateRequest) SetTemplateName(v string) *CreateCardSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

type CreateCardSmsTemplateShrinkRequest struct {
	Memo           *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateCardSmsTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCardSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateShrinkRequest) SetMemo(v string) *CreateCardSmsTemplateShrinkRequest {
	s.Memo = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) SetTemplateShrink(v string) *CreateCardSmsTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) SetTemplateName(v string) *CreateCardSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

type CreateCardSmsTemplateResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCardSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCardSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponseBody) SetCode(v string) *CreateCardSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetData(v *CreateCardSmsTemplateResponseBodyData) *CreateCardSmsTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetRequestId(v string) *CreateCardSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCardSmsTemplateResponseBody) SetSuccess(v bool) *CreateCardSmsTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateCardSmsTemplateResponseBodyData struct {
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateCardSmsTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCardSmsTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponseBodyData) SetTemplateCode(v string) *CreateCardSmsTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

type CreateCardSmsTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCardSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCardSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateCardSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCardSmsTemplateResponse) SetBody(v *CreateCardSmsTemplateResponseBody) *CreateCardSmsTemplateResponse {
	s.Body = v
	return s
}

type DeleteShortUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceUrl            *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s DeleteShortUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteShortUrlRequest) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlRequest) SetOwnerId(v int64) *DeleteShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteShortUrlRequest) SetResourceOwnerAccount(v string) *DeleteShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteShortUrlRequest) SetResourceOwnerId(v int64) *DeleteShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteShortUrlRequest) SetSourceUrl(v string) *DeleteShortUrlRequest {
	s.SourceUrl = &v
	return s
}

type DeleteShortUrlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteShortUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlResponseBody) SetCode(v string) *DeleteShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteShortUrlResponseBody) SetMessage(v string) *DeleteShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteShortUrlResponseBody) SetRequestId(v string) *DeleteShortUrlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteShortUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteShortUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShortUrlResponse) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlResponse) SetHeaders(v map[string]*string) *DeleteShortUrlResponse {
	s.Headers = v
	return s
}

func (s *DeleteShortUrlResponse) SetBody(v *DeleteShortUrlResponseBody) *DeleteShortUrlResponse {
	s.Body = v
	return s
}

type DeleteSmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DeleteSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignRequest) SetOwnerId(v int64) *DeleteSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmsSignRequest) SetResourceOwnerAccount(v string) *DeleteSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmsSignRequest) SetResourceOwnerId(v int64) *DeleteSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmsSignRequest) SetSignName(v string) *DeleteSmsSignRequest {
	s.SignName = &v
	return s
}

type DeleteSmsSignResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignName  *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DeleteSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponseBody) SetCode(v string) *DeleteSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetMessage(v string) *DeleteSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetRequestId(v string) *DeleteSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmsSignResponseBody) SetSignName(v string) *DeleteSmsSignResponseBody {
	s.SignName = &v
	return s
}

type DeleteSmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponse) SetHeaders(v map[string]*string) *DeleteSmsSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsSignResponse) SetBody(v *DeleteSmsSignResponseBody) *DeleteSmsSignResponse {
	s.Body = v
	return s
}

type DeleteSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateRequest) SetOwnerId(v int64) *DeleteSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetResourceOwnerAccount(v string) *DeleteSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetResourceOwnerId(v int64) *DeleteSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetTemplateCode(v string) *DeleteSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

type DeleteSmsTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponseBody) SetCode(v string) *DeleteSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetMessage(v string) *DeleteSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetRequestId(v string) *DeleteSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmsTemplateResponseBody) SetTemplateCode(v string) *DeleteSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

type DeleteSmsTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponse) SetHeaders(v map[string]*string) *DeleteSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsTemplateResponse) SetBody(v *DeleteSmsTemplateResponseBody) *DeleteSmsTemplateResponse {
	s.Body = v
	return s
}

type GetMediaResourceIdRequest struct {
	ExtendInfo   *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FileSize     *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	ResourceType *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMediaResourceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaResourceIdRequest) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdRequest) SetExtendInfo(v string) *GetMediaResourceIdRequest {
	s.ExtendInfo = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetFileSize(v int64) *GetMediaResourceIdRequest {
	s.FileSize = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetMemo(v string) *GetMediaResourceIdRequest {
	s.Memo = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetOssKey(v string) *GetMediaResourceIdRequest {
	s.OssKey = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetResourceType(v int32) *GetMediaResourceIdRequest {
	s.ResourceType = &v
	return s
}

type GetMediaResourceIdResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMediaResourceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMediaResourceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaResourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponseBody) SetCode(v string) *GetMediaResourceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetData(v *GetMediaResourceIdResponseBodyData) *GetMediaResourceIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetRequestId(v string) *GetMediaResourceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaResourceIdResponseBody) SetSuccess(v bool) *GetMediaResourceIdResponseBody {
	s.Success = &v
	return s
}

type GetMediaResourceIdResponseBodyData struct {
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetMediaResourceIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMediaResourceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponseBodyData) SetResourceId(v int64) *GetMediaResourceIdResponseBodyData {
	s.ResourceId = &v
	return s
}

type GetMediaResourceIdResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaResourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaResourceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaResourceIdResponse) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponse) SetHeaders(v map[string]*string) *GetMediaResourceIdResponse {
	s.Headers = v
	return s
}

func (s *GetMediaResourceIdResponse) SetBody(v *GetMediaResourceIdResponseBody) *GetMediaResourceIdResponse {
	s.Body = v
	return s
}

type GetOSSInfoForCardTemplateResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOSSInfoForCardTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOSSInfoForCardTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetCode(v string) *GetOSSInfoForCardTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetData(v *GetOSSInfoForCardTemplateResponseBodyData) *GetOSSInfoForCardTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetRequestId(v string) *GetOSSInfoForCardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetSuccess(v bool) *GetOSSInfoForCardTemplateResponseBody {
	s.Success = &v
	return s
}

type GetOSSInfoForCardTemplateResponseBodyData struct {
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AliUid      *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOSSInfoForCardTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetAccessKeyId(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetAliUid(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetExpireTime(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetHost(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetPolicy(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetSignature(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Signature = &v
	return s
}

type GetOSSInfoForCardTemplateResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOSSInfoForCardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOSSInfoForCardTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponse) SetHeaders(v map[string]*string) *GetOSSInfoForCardTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetOSSInfoForCardTemplateResponse) SetBody(v *GetOSSInfoForCardTemplateResponseBody) *GetOSSInfoForCardTemplateResponse {
	s.Body = v
	return s
}

type ModifySmsSignRequest struct {
	OwnerId              *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignFileList         []*ModifySmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	SignName             *string                             `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignSource           *int32                              `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
}

func (s ModifySmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsSignRequest) GoString() string {
	return s.String()
}

func (s *ModifySmsSignRequest) SetOwnerId(v int64) *ModifySmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmsSignRequest) SetRemark(v string) *ModifySmsSignRequest {
	s.Remark = &v
	return s
}

func (s *ModifySmsSignRequest) SetResourceOwnerAccount(v string) *ModifySmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsSignRequest) SetResourceOwnerId(v int64) *ModifySmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignFileList(v []*ModifySmsSignRequestSignFileList) *ModifySmsSignRequest {
	s.SignFileList = v
	return s
}

func (s *ModifySmsSignRequest) SetSignName(v string) *ModifySmsSignRequest {
	s.SignName = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignSource(v int32) *ModifySmsSignRequest {
	s.SignSource = &v
	return s
}

type ModifySmsSignRequestSignFileList struct {
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	FileSuffix   *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func (s ModifySmsSignRequestSignFileList) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsSignRequestSignFileList) GoString() string {
	return s.String()
}

func (s *ModifySmsSignRequestSignFileList) SetFileContents(v string) *ModifySmsSignRequestSignFileList {
	s.FileContents = &v
	return s
}

func (s *ModifySmsSignRequestSignFileList) SetFileSuffix(v string) *ModifySmsSignRequestSignFileList {
	s.FileSuffix = &v
	return s
}

type ModifySmsSignResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignName  *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s ModifySmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmsSignResponseBody) SetCode(v string) *ModifySmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetMessage(v string) *ModifySmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetRequestId(v string) *ModifySmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmsSignResponseBody) SetSignName(v string) *ModifySmsSignResponseBody {
	s.SignName = &v
	return s
}

type ModifySmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsSignResponse) GoString() string {
	return s.String()
}

func (s *ModifySmsSignResponse) SetHeaders(v map[string]*string) *ModifySmsSignResponse {
	s.Headers = v
	return s
}

func (s *ModifySmsSignResponse) SetBody(v *ModifySmsSignResponseBody) *ModifySmsSignResponse {
	s.Body = v
	return s
}

type ModifySmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateContent      *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifySmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateRequest) SetOwnerId(v int64) *ModifySmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetRemark(v string) *ModifySmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetResourceOwnerAccount(v string) *ModifySmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetResourceOwnerId(v int64) *ModifySmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateCode(v string) *ModifySmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateContent(v string) *ModifySmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateName(v string) *ModifySmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateType(v int32) *ModifySmsTemplateRequest {
	s.TemplateType = &v
	return s
}

type ModifySmsTemplateResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ModifySmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateResponseBody) SetCode(v string) *ModifySmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetMessage(v string) *ModifySmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetRequestId(v string) *ModifySmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmsTemplateResponseBody) SetTemplateCode(v string) *ModifySmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

type ModifySmsTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateResponse) SetHeaders(v map[string]*string) *ModifySmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifySmsTemplateResponse) SetBody(v *ModifySmsTemplateResponseBody) *ModifySmsTemplateResponse {
	s.Body = v
	return s
}

type QueryCardSmsTemplateRequest struct {
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryCardSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateRequest) SetTemplateCode(v string) *QueryCardSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

type QueryCardSmsTemplateResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCardSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponseBody) SetCode(v string) *QueryCardSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetData(v *QueryCardSmsTemplateResponseBodyData) *QueryCardSmsTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetRequestId(v string) *QueryCardSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCardSmsTemplateResponseBody) SetSuccess(v bool) *QueryCardSmsTemplateResponseBody {
	s.Success = &v
	return s
}

type QueryCardSmsTemplateResponseBodyData struct {
	Templates []map[string]interface{} `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponseBodyData) SetTemplates(v []map[string]interface{}) *QueryCardSmsTemplateResponseBodyData {
	s.Templates = v
	return s
}

type QueryCardSmsTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateResponse) SetHeaders(v map[string]*string) *QueryCardSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryCardSmsTemplateResponse) SetBody(v *QueryCardSmsTemplateResponseBody) *QueryCardSmsTemplateResponse {
	s.Body = v
	return s
}

type QuerySendDetailsRequest struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CurrentPage          *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SendDate             *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
}

func (s QuerySendDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsRequest) SetBizId(v string) *QuerySendDetailsRequest {
	s.BizId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetCurrentPage(v int64) *QuerySendDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuerySendDetailsRequest) SetOwnerId(v int64) *QuerySendDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPageSize(v int64) *QuerySendDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPhoneNumber(v string) *QuerySendDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerAccount(v string) *QuerySendDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerId(v int64) *QuerySendDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetSendDate(v string) *QuerySendDetailsRequest {
	s.SendDate = &v
	return s
}

type QuerySendDetailsResponseBody struct {
	Code              *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message           *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsSendDetailDTOs *QuerySendDetailsResponseBodySmsSendDetailDTOs `json:"SmsSendDetailDTOs,omitempty" xml:"SmsSendDetailDTOs,omitempty" type:"Struct"`
	TotalCount        *string                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBody) SetCode(v string) *QuerySendDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetMessage(v string) *QuerySendDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetRequestId(v string) *QuerySendDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetSmsSendDetailDTOs(v *QuerySendDetailsResponseBodySmsSendDetailDTOs) *QuerySendDetailsResponseBody {
	s.SmsSendDetailDTOs = v
	return s
}

func (s *QuerySendDetailsResponseBody) SetTotalCount(v string) *QuerySendDetailsResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySendDetailsResponseBodySmsSendDetailDTOs struct {
	SmsSendDetailDTO []*QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO `json:"SmsSendDetailDTO,omitempty" xml:"SmsSendDetailDTO,omitempty" type:"Repeated"`
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOs) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOs) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOs) SetSmsSendDetailDTO(v []*QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) *QuerySendDetailsResponseBodySmsSendDetailDTOs {
	s.SmsSendDetailDTO = v
	return s
}

type QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrCode      *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	OutId        *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PhoneNum     *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	ReceiveDate  *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	SendDate     *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	SendStatus   *int64  `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetContent(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.Content = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetErrCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.ErrCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetOutId(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.OutId = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetPhoneNum(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.PhoneNum = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetReceiveDate(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.ReceiveDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetSendDate(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetSendStatus(v int64) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.SendStatus = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetTemplateCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.TemplateCode = &v
	return s
}

type QuerySendDetailsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySendDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySendDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponse) SetHeaders(v map[string]*string) *QuerySendDetailsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendDetailsResponse) SetBody(v *QuerySendDetailsResponseBody) *QuerySendDetailsResponse {
	s.Body = v
	return s
}

type QueryShortUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ShortUrl             *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryShortUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryShortUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryShortUrlRequest) SetOwnerId(v int64) *QueryShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryShortUrlRequest) SetResourceOwnerAccount(v string) *QueryShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryShortUrlRequest) SetResourceOwnerId(v int64) *QueryShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryShortUrlRequest) SetShortUrl(v string) *QueryShortUrlRequest {
	s.ShortUrl = &v
	return s
}

type QueryShortUrlResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryShortUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponseBody) SetCode(v string) *QueryShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryShortUrlResponseBody) SetData(v *QueryShortUrlResponseBodyData) *QueryShortUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryShortUrlResponseBody) SetMessage(v string) *QueryShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *QueryShortUrlResponseBody) SetRequestId(v string) *QueryShortUrlResponseBody {
	s.RequestId = &v
	return s
}

type QueryShortUrlResponseBodyData struct {
	CreateDate         *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	ExpireDate         *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	PageViewCount      *string `json:"PageViewCount,omitempty" xml:"PageViewCount,omitempty"`
	ShortUrl           *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	ShortUrlName       *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	ShortUrlStatus     *string `json:"ShortUrlStatus,omitempty" xml:"ShortUrlStatus,omitempty"`
	SourceUrl          *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	UniqueVisitorCount *string `json:"UniqueVisitorCount,omitempty" xml:"UniqueVisitorCount,omitempty"`
}

func (s QueryShortUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryShortUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponseBodyData) SetCreateDate(v string) *QueryShortUrlResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetExpireDate(v string) *QueryShortUrlResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetPageViewCount(v string) *QueryShortUrlResponseBodyData {
	s.PageViewCount = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrl(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrl = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrlName(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrlName = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetShortUrlStatus(v string) *QueryShortUrlResponseBodyData {
	s.ShortUrlStatus = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetSourceUrl(v string) *QueryShortUrlResponseBodyData {
	s.SourceUrl = &v
	return s
}

func (s *QueryShortUrlResponseBodyData) SetUniqueVisitorCount(v string) *QueryShortUrlResponseBodyData {
	s.UniqueVisitorCount = &v
	return s
}

type QueryShortUrlResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryShortUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShortUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponse) SetHeaders(v map[string]*string) *QueryShortUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryShortUrlResponse) SetBody(v *QueryShortUrlResponseBody) *QueryShortUrlResponse {
	s.Body = v
	return s
}

type QuerySmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QuerySmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsSignRequest) SetOwnerId(v int64) *QuerySmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsSignRequest) SetResourceOwnerAccount(v string) *QuerySmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsSignRequest) SetResourceOwnerId(v int64) *QuerySmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsSignRequest) SetSignName(v string) *QuerySmsSignRequest {
	s.SignName = &v
	return s
}

type QuerySmsSignResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignStatus *int32  `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
}

func (s QuerySmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsSignResponseBody) SetCode(v string) *QuerySmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetCreateDate(v string) *QuerySmsSignResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetMessage(v string) *QuerySmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetReason(v string) *QuerySmsSignResponseBody {
	s.Reason = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetRequestId(v string) *QuerySmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignName(v string) *QuerySmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignStatus(v int32) *QuerySmsSignResponseBody {
	s.SignStatus = &v
	return s
}

type QuerySmsSignResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsSignResponse) SetHeaders(v map[string]*string) *QuerySmsSignResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsSignResponse) SetBody(v *QuerySmsSignResponseBody) *QuerySmsSignResponse {
	s.Body = v
	return s
}

type QuerySmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateRequest) SetOwnerId(v int64) *QuerySmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetResourceOwnerAccount(v string) *QuerySmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetResourceOwnerId(v int64) *QuerySmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetTemplateCode(v string) *QuerySmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

type QuerySmsTemplateResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateStatus  *int32  `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateResponseBody) SetCode(v string) *QuerySmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetCreateDate(v string) *QuerySmsTemplateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetMessage(v string) *QuerySmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetReason(v string) *QuerySmsTemplateResponseBody {
	s.Reason = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetRequestId(v string) *QuerySmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateCode(v string) *QuerySmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateContent(v string) *QuerySmsTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateName(v string) *QuerySmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateStatus(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateType(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

type QuerySmsTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateResponse) SetHeaders(v map[string]*string) *QuerySmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsTemplateResponse) SetBody(v *QuerySmsTemplateResponseBody) *QuerySmsTemplateResponse {
	s.Body = v
	return s
}

type SendBatchSmsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumberJson      *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignNameJson         *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	SmsUpExtendCodeJson  *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParamJson    *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
}

func (s SendBatchSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBatchSmsRequest) GoString() string {
	return s.String()
}

func (s *SendBatchSmsRequest) SetOwnerId(v int64) *SendBatchSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SendBatchSmsRequest) SetPhoneNumberJson(v string) *SendBatchSmsRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetResourceOwnerAccount(v string) *SendBatchSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendBatchSmsRequest) SetResourceOwnerId(v int64) *SendBatchSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendBatchSmsRequest) SetSignNameJson(v string) *SendBatchSmsRequest {
	s.SignNameJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetSmsUpExtendCodeJson(v string) *SendBatchSmsRequest {
	s.SmsUpExtendCodeJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetTemplateCode(v string) *SendBatchSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendBatchSmsRequest) SetTemplateParamJson(v string) *SendBatchSmsRequest {
	s.TemplateParamJson = &v
	return s
}

type SendBatchSmsResponseBody struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendBatchSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBatchSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchSmsResponseBody) SetBizId(v string) *SendBatchSmsResponseBody {
	s.BizId = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetCode(v string) *SendBatchSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetMessage(v string) *SendBatchSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetRequestId(v string) *SendBatchSmsResponseBody {
	s.RequestId = &v
	return s
}

type SendBatchSmsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendBatchSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendBatchSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendBatchSmsResponse) GoString() string {
	return s.String()
}

func (s *SendBatchSmsResponse) SetHeaders(v map[string]*string) *SendBatchSmsResponse {
	s.Headers = v
	return s
}

func (s *SendBatchSmsResponse) SetBody(v *SendBatchSmsResponseBody) *SendBatchSmsResponse {
	s.Body = v
	return s
}

type SendCardSmsRequest struct {
	CardObjects          map[string]interface{} `json:"CardObjects,omitempty" xml:"CardObjects,omitempty"`
	CardTemplateCode     *string                `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	DigitalTemplateCode  *string                `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	DigitalTemplateParam *string                `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
	FallbackType         *string                `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	OutId                *string                `json:"OutId,omitempty" xml:"OutId,omitempty"`
	SignName             *string                `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsTemplateCode      *string                `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam     *string                `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	SmsUpExtendCode      *string                `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
}

func (s SendCardSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsRequest) GoString() string {
	return s.String()
}

func (s *SendCardSmsRequest) SetCardObjects(v map[string]interface{}) *SendCardSmsRequest {
	s.CardObjects = v
	return s
}

func (s *SendCardSmsRequest) SetCardTemplateCode(v string) *SendCardSmsRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetDigitalTemplateCode(v string) *SendCardSmsRequest {
	s.DigitalTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetDigitalTemplateParam(v string) *SendCardSmsRequest {
	s.DigitalTemplateParam = &v
	return s
}

func (s *SendCardSmsRequest) SetFallbackType(v string) *SendCardSmsRequest {
	s.FallbackType = &v
	return s
}

func (s *SendCardSmsRequest) SetOutId(v string) *SendCardSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendCardSmsRequest) SetSignName(v string) *SendCardSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsTemplateCode(v string) *SendCardSmsRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsTemplateParam(v string) *SendCardSmsRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *SendCardSmsRequest) SetSmsUpExtendCode(v string) *SendCardSmsRequest {
	s.SmsUpExtendCode = &v
	return s
}

type SendCardSmsShrinkRequest struct {
	CardObjectsShrink    *string `json:"CardObjects,omitempty" xml:"CardObjects,omitempty"`
	CardTemplateCode     *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	DigitalTemplateCode  *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	DigitalTemplateParam *string `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
	FallbackType         *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsTemplateCode      *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam     *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	SmsUpExtendCode      *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
}

func (s SendCardSmsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendCardSmsShrinkRequest) SetCardObjectsShrink(v string) *SendCardSmsShrinkRequest {
	s.CardObjectsShrink = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetCardTemplateCode(v string) *SendCardSmsShrinkRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetDigitalTemplateCode(v string) *SendCardSmsShrinkRequest {
	s.DigitalTemplateCode = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetDigitalTemplateParam(v string) *SendCardSmsShrinkRequest {
	s.DigitalTemplateParam = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetFallbackType(v string) *SendCardSmsShrinkRequest {
	s.FallbackType = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetOutId(v string) *SendCardSmsShrinkRequest {
	s.OutId = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetSignName(v string) *SendCardSmsShrinkRequest {
	s.SignName = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetSmsTemplateCode(v string) *SendCardSmsShrinkRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetSmsTemplateParam(v string) *SendCardSmsShrinkRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *SendCardSmsShrinkRequest) SetSmsUpExtendCode(v string) *SendCardSmsShrinkRequest {
	s.SmsUpExtendCode = &v
	return s
}

type SendCardSmsResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SendCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCardSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponseBody) SetCode(v string) *SendCardSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendCardSmsResponseBody) SetData(v *SendCardSmsResponseBodyData) *SendCardSmsResponseBody {
	s.Data = v
	return s
}

func (s *SendCardSmsResponseBody) SetRequestId(v string) *SendCardSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCardSmsResponseBody) SetSuccess(v bool) *SendCardSmsResponseBody {
	s.Success = &v
	return s
}

type SendCardSmsResponseBodyData struct {
	BizCardId       *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	BizDigitalId    *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	BizSmsId        *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	CardTmpState    *int32  `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	MediaMobiles    *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s SendCardSmsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponseBodyData) SetBizCardId(v string) *SendCardSmsResponseBodyData {
	s.BizCardId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetBizDigitalId(v string) *SendCardSmsResponseBodyData {
	s.BizDigitalId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetBizSmsId(v string) *SendCardSmsResponseBodyData {
	s.BizSmsId = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetCardTmpState(v int32) *SendCardSmsResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetMediaMobiles(v string) *SendCardSmsResponseBodyData {
	s.MediaMobiles = &v
	return s
}

func (s *SendCardSmsResponseBodyData) SetNotMediaMobiles(v string) *SendCardSmsResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

type SendCardSmsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCardSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsResponse) GoString() string {
	return s.String()
}

func (s *SendCardSmsResponse) SetHeaders(v map[string]*string) *SendCardSmsResponse {
	s.Headers = v
	return s
}

func (s *SendCardSmsResponse) SetBody(v *SendCardSmsResponseBody) *SendCardSmsResponse {
	s.Body = v
	return s
}

type SendMessageToGlobeRequest struct {
	From                 *string `json:"From,omitempty" xml:"From,omitempty"`
	Message              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	To                   *string `json:"To,omitempty" xml:"To,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeRequest) SetFrom(v string) *SendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetMessage(v string) *SendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetOwnerId(v int64) *SendMessageToGlobeRequest {
	s.OwnerId = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetResourceOwnerAccount(v string) *SendMessageToGlobeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetResourceOwnerId(v int64) *SendMessageToGlobeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetTo(v string) *SendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetType(v string) *SendMessageToGlobeRequest {
	s.Type = &v
	return s
}

type SendMessageToGlobeResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	From         *string                                     `json:"From,omitempty" xml:"From,omitempty"`
	MessageId    *string                                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	NumberDetail *SendMessageToGlobeResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Segments     *string                                     `json:"Segments,omitempty" xml:"Segments,omitempty"`
	To           *string                                     `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBody) SetCode(v string) *SendMessageToGlobeResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetFrom(v string) *SendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetMessageId(v string) *SendMessageToGlobeResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetNumberDetail(v *SendMessageToGlobeResponseBodyNumberDetail) *SendMessageToGlobeResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetRequestId(v string) *SendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetSegments(v string) *SendMessageToGlobeResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetTo(v string) *SendMessageToGlobeResponseBody {
	s.To = &v
	return s
}

type SendMessageToGlobeResponseBodyNumberDetail struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SendMessageToGlobeResponseBodyNumberDetail) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCarrier(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCountry(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetRegion(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Region = &v
	return s
}

type SendMessageToGlobeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageToGlobeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponse) SetHeaders(v map[string]*string) *SendMessageToGlobeResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToGlobeResponse) SetBody(v *SendMessageToGlobeResponseBody) *SendMessageToGlobeResponse {
	s.Body = v
	return s
}

type SendSmsRequest struct {
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsUpExtendCode      *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam        *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsRequest) GoString() string {
	return s.String()
}

func (s *SendSmsRequest) SetOutId(v string) *SendSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendSmsRequest) SetOwnerId(v int64) *SendSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *SendSmsRequest) SetPhoneNumbers(v string) *SendSmsRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendSmsRequest) SetResourceOwnerAccount(v string) *SendSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendSmsRequest) SetResourceOwnerId(v int64) *SendSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendSmsRequest) SetSignName(v string) *SendSmsRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsRequest) SetSmsUpExtendCode(v string) *SendSmsRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsRequest) SetTemplateCode(v string) *SendSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsRequest) SetTemplateParam(v string) *SendSmsRequest {
	s.TemplateParam = &v
	return s
}

type SendSmsResponseBody struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBody) SetBizId(v string) *SendSmsResponseBody {
	s.BizId = &v
	return s
}

func (s *SendSmsResponseBody) SetCode(v string) *SendSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsResponseBody) SetMessage(v string) *SendSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsResponseBody) SetRequestId(v string) *SendSmsResponseBody {
	s.RequestId = &v
	return s
}

type SendSmsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponse) GoString() string {
	return s.String()
}

func (s *SendSmsResponse) SetHeaders(v map[string]*string) *SendSmsResponse {
	s.Headers = v
	return s
}

func (s *SendSmsResponse) SetBody(v *SendSmsResponseBody) *SendSmsResponse {
	s.Body = v
	return s
}

type SendSmsSmartRequest struct {
	ExtendCode      *string `json:"ExtendCode,omitempty" xml:"ExtendCode,omitempty"`
	ModelCode       *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	NumberType      *string `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	PhoneNumbers    *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	SignName        *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam   *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendSmsSmartRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsSmartRequest) GoString() string {
	return s.String()
}

func (s *SendSmsSmartRequest) SetExtendCode(v string) *SendSmsSmartRequest {
	s.ExtendCode = &v
	return s
}

func (s *SendSmsSmartRequest) SetModelCode(v string) *SendSmsSmartRequest {
	s.ModelCode = &v
	return s
}

func (s *SendSmsSmartRequest) SetNumberType(v string) *SendSmsSmartRequest {
	s.NumberType = &v
	return s
}

func (s *SendSmsSmartRequest) SetPhoneNumbers(v string) *SendSmsSmartRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendSmsSmartRequest) SetSignName(v string) *SendSmsSmartRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsSmartRequest) SetSmsUpExtendCode(v string) *SendSmsSmartRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsSmartRequest) SetTemplateCode(v string) *SendSmsSmartRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsSmartRequest) SetTemplateParam(v string) *SendSmsSmartRequest {
	s.TemplateParam = &v
	return s
}

type SendSmsSmartResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SendSmsSmartResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendSmsSmartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsSmartResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsSmartResponseBody) SetCode(v string) *SendSmsSmartResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsSmartResponseBody) SetData(v *SendSmsSmartResponseBodyData) *SendSmsSmartResponseBody {
	s.Data = v
	return s
}

func (s *SendSmsSmartResponseBody) SetRequestId(v string) *SendSmsSmartResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSmsSmartResponseBody) SetSuccess(v bool) *SendSmsSmartResponseBody {
	s.Success = &v
	return s
}

type SendSmsSmartResponseBodyData struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s SendSmsSmartResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendSmsSmartResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendSmsSmartResponseBodyData) SetBizId(v string) *SendSmsSmartResponseBodyData {
	s.BizId = &v
	return s
}

type SendSmsSmartResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendSmsSmartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendSmsSmartResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSmsSmartResponse) GoString() string {
	return s.String()
}

func (s *SendSmsSmartResponse) SetHeaders(v map[string]*string) *SendSmsSmartResponse {
	s.Headers = v
	return s
}

func (s *SendSmsSmartResponse) SetBody(v *SendSmsSmartResponseBody) *SendSmsSmartResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"ap-southeast-1": tea.String("dysmsapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": tea.String("dysmsapi-xman.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     tea.String("dysmsapi-proxy.cn-beijing.aliyuncs.com"),
		"cn-hongkong":    tea.String("dysmsapi-xman.cn-hongkong.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dysmsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddShortUrlWithOptions(request *AddShortUrlRequest, runtime *util.RuntimeOptions) (_result *AddShortUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddShortUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddShortUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddShortUrl(request *AddShortUrlRequest) (_result *AddShortUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddShortUrlResponse{}
	_body, _err := client.AddShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSmsSignWithOptions(request *AddSmsSignRequest, runtime *util.RuntimeOptions) (_result *AddSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSmsSign"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSmsSign(request *AddSmsSignRequest) (_result *AddSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSmsSignResponse{}
	_body, _err := client.AddSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSmsTemplateWithOptions(request *AddSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *AddSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (_result *AddSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSmsTemplateResponse{}
	_body, _err := client.AddSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCardSmsTemplateWithOptions(tmpReq *CreateCardSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateCardSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCardSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Template)) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, tea.String("Template"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCardSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCardSmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCardSmsTemplate(request *CreateCardSmsTemplateRequest) (_result *CreateCardSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCardSmsTemplateResponse{}
	_body, _err := client.CreateCardSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteShortUrlWithOptions(request *DeleteShortUrlRequest, runtime *util.RuntimeOptions) (_result *DeleteShortUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteShortUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteShortUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteShortUrl(request *DeleteShortUrlRequest) (_result *DeleteShortUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteShortUrlResponse{}
	_body, _err := client.DeleteShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSmsSignWithOptions(request *DeleteSmsSignRequest, runtime *util.RuntimeOptions) (_result *DeleteSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSmsSign"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (_result *DeleteSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.DeleteSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSmsTemplateWithOptions(request *DeleteSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (_result *DeleteSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.DeleteSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaResourceIdWithOptions(request *GetMediaResourceIdRequest, runtime *util.RuntimeOptions) (_result *GetMediaResourceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaResourceIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaResourceId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaResourceId(request *GetMediaResourceIdRequest) (_result *GetMediaResourceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaResourceIdResponse{}
	_body, _err := client.GetMediaResourceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOSSInfoForCardTemplateWithOptions(runtime *util.RuntimeOptions) (_result *GetOSSInfoForCardTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetOSSInfoForCardTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOSSInfoForCardTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOSSInfoForCardTemplate() (_result *GetOSSInfoForCardTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOSSInfoForCardTemplateResponse{}
	_body, _err := client.GetOSSInfoForCardTemplateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySmsSignWithOptions(request *ModifySmsSignRequest, runtime *util.RuntimeOptions) (_result *ModifySmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySmsSign"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySmsSign(request *ModifySmsSignRequest) (_result *ModifySmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySmsSignResponse{}
	_body, _err := client.ModifySmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySmsTemplateWithOptions(request *ModifySmsTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifySmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (_result *ModifySmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySmsTemplateResponse{}
	_body, _err := client.ModifySmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCardSmsTemplateWithOptions(request *QueryCardSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *QueryCardSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCardSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCardSmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardSmsTemplate(request *QueryCardSmsTemplateRequest) (_result *QueryCardSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardSmsTemplateResponse{}
	_body, _err := client.QueryCardSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySendDetailsWithOptions(request *QuerySendDetailsRequest, runtime *util.RuntimeOptions) (_result *QuerySendDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySendDetails"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySendDetails(request *QuerySendDetailsRequest) (_result *QuerySendDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.QuerySendDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryShortUrlWithOptions(request *QueryShortUrlRequest, runtime *util.RuntimeOptions) (_result *QueryShortUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryShortUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryShortUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryShortUrl(request *QueryShortUrlRequest) (_result *QueryShortUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryShortUrlResponse{}
	_body, _err := client.QueryShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySmsSignWithOptions(request *QuerySmsSignRequest, runtime *util.RuntimeOptions) (_result *QuerySmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySmsSign"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySmsSign(request *QuerySmsSignRequest) (_result *QuerySmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsSignResponse{}
	_body, _err := client.QuerySmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySmsTemplateWithOptions(request *QuerySmsTemplateRequest, runtime *util.RuntimeOptions) (_result *QuerySmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySmsTemplate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySmsTemplate(request *QuerySmsTemplateRequest) (_result *QuerySmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsTemplateResponse{}
	_body, _err := client.QuerySmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendBatchSmsWithOptions(request *SendBatchSmsRequest, runtime *util.RuntimeOptions) (_result *SendBatchSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendBatchSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendBatchSms"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendBatchSms(request *SendBatchSmsRequest) (_result *SendBatchSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendBatchSmsResponse{}
	_body, _err := client.SendBatchSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCardSmsWithOptions(tmpReq *SendCardSmsRequest, runtime *util.RuntimeOptions) (_result *SendCardSmsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendCardSmsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CardObjects)) {
		request.CardObjectsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CardObjects, tea.String("CardObjects"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCardSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCardSms"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCardSms(request *SendCardSmsRequest) (_result *SendCardSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCardSmsResponse{}
	_body, _err := client.SendCardSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageToGlobeWithOptions(request *SendMessageToGlobeRequest, runtime *util.RuntimeOptions) (_result *SendMessageToGlobeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMessageToGlobe"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessageToGlobe(request *SendMessageToGlobeRequest) (_result *SendMessageToGlobeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.SendMessageToGlobeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSmsWithOptions(request *SendSmsRequest, runtime *util.RuntimeOptions) (_result *SendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendSmsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendSms"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSms(request *SendSmsRequest) (_result *SendSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendSmsResponse{}
	_body, _err := client.SendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendSmsSmartWithOptions(request *SendSmsSmartRequest, runtime *util.RuntimeOptions) (_result *SendSmsSmartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendSmsSmartResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendSmsSmart"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendSmsSmart(request *SendSmsSmartRequest) (_result *SendSmsSmartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendSmsSmartResponse{}
	_body, _err := client.SendSmsSmartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
