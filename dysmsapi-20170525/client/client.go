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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddShortUrlResponse) SetStatusCode(v int32) *AddShortUrlResponse {
	s.StatusCode = &v
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
	SignType             *int32                           `json:"SignType,omitempty" xml:"SignType,omitempty"`
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

func (s *AddSmsSignRequest) SetSignType(v int32) *AddSmsSignRequest {
	s.SignType = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddSmsSignResponse) SetStatusCode(v int32) *AddSmsSignResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddSmsTemplateResponse) SetStatusCode(v int32) *AddSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSmsTemplateResponse) SetBody(v *AddSmsTemplateResponseBody) *AddSmsTemplateResponse {
	s.Body = v
	return s
}

type CheckMobilesCardSupportRequest struct {
	Mobiles      []map[string]interface{} `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
	TemplateCode *string                  `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CheckMobilesCardSupportRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMobilesCardSupportRequest) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportRequest) SetMobiles(v []map[string]interface{}) *CheckMobilesCardSupportRequest {
	s.Mobiles = v
	return s
}

func (s *CheckMobilesCardSupportRequest) SetTemplateCode(v string) *CheckMobilesCardSupportRequest {
	s.TemplateCode = &v
	return s
}

type CheckMobilesCardSupportResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckMobilesCardSupportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBody) SetCode(v string) *CheckMobilesCardSupportResponseBody {
	s.Code = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetData(v []map[string]interface{}) *CheckMobilesCardSupportResponseBody {
	s.Data = v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetRequestId(v string) *CheckMobilesCardSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetSuccess(v bool) *CheckMobilesCardSupportResponseBody {
	s.Success = &v
	return s
}

type CheckMobilesCardSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckMobilesCardSupportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMobilesCardSupportResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMobilesCardSupportResponse) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponse) SetHeaders(v map[string]*string) *CheckMobilesCardSupportResponse {
	s.Headers = v
	return s
}

func (s *CheckMobilesCardSupportResponse) SetStatusCode(v int32) *CheckMobilesCardSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMobilesCardSupportResponse) SetBody(v *CheckMobilesCardSupportResponseBody) *CheckMobilesCardSupportResponse {
	s.Body = v
	return s
}

type CreateCardSmsTemplateRequest struct {
	Factorys     *string                `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
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

func (s *CreateCardSmsTemplateRequest) SetFactorys(v string) *CreateCardSmsTemplateRequest {
	s.Factorys = &v
	return s
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
	Factorys       *string `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
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

func (s *CreateCardSmsTemplateShrinkRequest) SetFactorys(v string) *CreateCardSmsTemplateShrinkRequest {
	s.Factorys = &v
	return s
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCardSmsTemplateResponse) SetStatusCode(v int32) *CreateCardSmsTemplateResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteShortUrlResponse) SetStatusCode(v int32) *DeleteShortUrlResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSmsSignResponse) SetStatusCode(v int32) *DeleteSmsSignResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSmsTemplateResponse) SetStatusCode(v int32) *DeleteSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmsTemplateResponse) SetBody(v *DeleteSmsTemplateResponseBody) *DeleteSmsTemplateResponse {
	s.Body = v
	return s
}

type GetCardSmsLinkRequest struct {
	CardCodeType          *int32  `json:"CardCodeType,omitempty" xml:"CardCodeType,omitempty"`
	CardLinkType          *int32  `json:"CardLinkType,omitempty" xml:"CardLinkType,omitempty"`
	CardTemplateCode      *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	CustomShortCodeJson   *string `json:"CustomShortCodeJson,omitempty" xml:"CustomShortCodeJson,omitempty"`
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OutId                 *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PhoneNumberJson       *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	SignNameJson          *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
}

func (s GetCardSmsLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsLinkRequest) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkRequest) SetCardCodeType(v int32) *GetCardSmsLinkRequest {
	s.CardCodeType = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardLinkType(v int32) *GetCardSmsLinkRequest {
	s.CardLinkType = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardTemplateCode(v string) *GetCardSmsLinkRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCardTemplateParamJson(v string) *GetCardSmsLinkRequest {
	s.CardTemplateParamJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetCustomShortCodeJson(v string) *GetCardSmsLinkRequest {
	s.CustomShortCodeJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetDomain(v string) *GetCardSmsLinkRequest {
	s.Domain = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetOutId(v string) *GetCardSmsLinkRequest {
	s.OutId = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetPhoneNumberJson(v string) *GetCardSmsLinkRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *GetCardSmsLinkRequest) SetSignNameJson(v string) *GetCardSmsLinkRequest {
	s.SignNameJson = &v
	return s
}

type GetCardSmsLinkResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCardSmsLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardSmsLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBody) SetCode(v string) *GetCardSmsLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetData(v *GetCardSmsLinkResponseBodyData) *GetCardSmsLinkResponseBody {
	s.Data = v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetRequestId(v string) *GetCardSmsLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardSmsLinkResponseBody) SetSuccess(v bool) *GetCardSmsLinkResponseBody {
	s.Success = &v
	return s
}

type GetCardSmsLinkResponseBodyData struct {
	CardPhoneNumbers *string `json:"CardPhoneNumbers,omitempty" xml:"CardPhoneNumbers,omitempty"`
	CardSignNames    *string `json:"CardSignNames,omitempty" xml:"CardSignNames,omitempty"`
	CardSmsLinks     *string `json:"CardSmsLinks,omitempty" xml:"CardSmsLinks,omitempty"`
	CardTmpState     *int32  `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	NotMediaMobiles  *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s GetCardSmsLinkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponseBodyData) SetCardPhoneNumbers(v string) *GetCardSmsLinkResponseBodyData {
	s.CardPhoneNumbers = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSignNames(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSignNames = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardSmsLinks(v string) *GetCardSmsLinkResponseBodyData {
	s.CardSmsLinks = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetCardTmpState(v int32) *GetCardSmsLinkResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *GetCardSmsLinkResponseBodyData) SetNotMediaMobiles(v string) *GetCardSmsLinkResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

type GetCardSmsLinkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardSmsLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardSmsLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsLinkResponse) GoString() string {
	return s.String()
}

func (s *GetCardSmsLinkResponse) SetHeaders(v map[string]*string) *GetCardSmsLinkResponse {
	s.Headers = v
	return s
}

func (s *GetCardSmsLinkResponse) SetStatusCode(v int32) *GetCardSmsLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardSmsLinkResponse) SetBody(v *GetCardSmsLinkResponseBody) *GetCardSmsLinkResponse {
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
	ResUrlDownload *string `json:"ResUrlDownload,omitempty" xml:"ResUrlDownload,omitempty"`
	ResourceId     *int64  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetMediaResourceIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMediaResourceIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdResponseBodyData) SetResUrlDownload(v string) *GetMediaResourceIdResponseBodyData {
	s.ResUrlDownload = &v
	return s
}

func (s *GetMediaResourceIdResponseBodyData) SetResourceId(v int64) *GetMediaResourceIdResponseBodyData {
	s.ResourceId = &v
	return s
}

type GetMediaResourceIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMediaResourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMediaResourceIdResponse) SetStatusCode(v int32) *GetMediaResourceIdResponse {
	s.StatusCode = &v
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
	Bucket      *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	StartPath   *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
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

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetBucket(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Bucket = &v
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

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetStartPath(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.StartPath = &v
	return s
}

type GetOSSInfoForCardTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOSSInfoForCardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOSSInfoForCardTemplateResponse) SetStatusCode(v int32) *GetOSSInfoForCardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponse) SetBody(v *GetOSSInfoForCardTemplateResponseBody) *GetOSSInfoForCardTemplateResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string                       `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetProdCode(v string) *ListTagResourcesRequest {
	s.ProdCode = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
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
	SignType             *int32                              `json:"SignType,omitempty" xml:"SignType,omitempty"`
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

func (s *ModifySmsSignRequest) SetSignType(v int32) *ModifySmsSignRequest {
	s.SignType = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySmsSignResponse) SetStatusCode(v int32) *ModifySmsSignResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySmsTemplateResponse) SetStatusCode(v int32) *ModifySmsTemplateResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryCardSmsTemplateResponse) SetStatusCode(v int32) *QueryCardSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardSmsTemplateResponse) SetBody(v *QueryCardSmsTemplateResponseBody) *QueryCardSmsTemplateResponse {
	s.Body = v
	return s
}

type QueryCardSmsTemplateReportRequest struct {
	EndDate       *string   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate     *string   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TemplateCodes []*string `json:"TemplateCodes,omitempty" xml:"TemplateCodes,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateReportRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateReportRequest) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportRequest) SetEndDate(v string) *QueryCardSmsTemplateReportRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCardSmsTemplateReportRequest) SetStartDate(v string) *QueryCardSmsTemplateReportRequest {
	s.StartDate = &v
	return s
}

func (s *QueryCardSmsTemplateReportRequest) SetTemplateCodes(v []*string) *QueryCardSmsTemplateReportRequest {
	s.TemplateCodes = v
	return s
}

type QueryCardSmsTemplateReportResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCardSmsTemplateReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponseBody) SetCode(v string) *QueryCardSmsTemplateReportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetData(v []map[string]interface{}) *QueryCardSmsTemplateReportResponseBody {
	s.Data = v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetRequestId(v string) *QueryCardSmsTemplateReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponseBody) SetSuccess(v bool) *QueryCardSmsTemplateReportResponseBody {
	s.Success = &v
	return s
}

type QueryCardSmsTemplateReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCardSmsTemplateReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCardSmsTemplateReportResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponse) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponse) SetHeaders(v map[string]*string) *QueryCardSmsTemplateReportResponse {
	s.Headers = v
	return s
}

func (s *QueryCardSmsTemplateReportResponse) SetStatusCode(v int32) *QueryCardSmsTemplateReportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCardSmsTemplateReportResponse) SetBody(v *QueryCardSmsTemplateReportResponseBody) *QueryCardSmsTemplateReportResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySendDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySendDetailsResponse) SetStatusCode(v int32) *QuerySendDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySendDetailsResponse) SetBody(v *QuerySendDetailsResponseBody) *QuerySendDetailsResponse {
	s.Body = v
	return s
}

type QuerySendStatisticsRequest struct {
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	IsGlobe              *int32  `json:"IsGlobe,omitempty" xml:"IsGlobe,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageIndex            *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySendStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsRequest) SetEndDate(v string) *QuerySendStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetIsGlobe(v int32) *QuerySendStatisticsRequest {
	s.IsGlobe = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetOwnerId(v int64) *QuerySendStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetPageIndex(v int32) *QuerySendStatisticsRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetPageSize(v int32) *QuerySendStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetResourceOwnerAccount(v string) *QuerySendStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetResourceOwnerId(v int64) *QuerySendStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetSignName(v string) *QuerySendStatisticsRequest {
	s.SignName = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetStartDate(v string) *QuerySendStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetTemplateType(v int32) *QuerySendStatisticsRequest {
	s.TemplateType = &v
	return s
}

type QuerySendStatisticsResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QuerySendStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySendStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySendStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBody) SetCode(v string) *QuerySendStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetData(v *QuerySendStatisticsResponseBodyData) *QuerySendStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetMessage(v string) *QuerySendStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendStatisticsResponseBody) SetRequestId(v string) *QuerySendStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type QuerySendStatisticsResponseBodyData struct {
	TargetList []*QuerySendStatisticsResponseBodyDataTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	TotalSize  *int64                                           `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QuerySendStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySendStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBodyData) SetTargetList(v []*QuerySendStatisticsResponseBodyDataTargetList) *QuerySendStatisticsResponseBodyData {
	s.TargetList = v
	return s
}

func (s *QuerySendStatisticsResponseBodyData) SetTotalSize(v int64) *QuerySendStatisticsResponseBodyData {
	s.TotalSize = &v
	return s
}

type QuerySendStatisticsResponseBodyDataTargetList struct {
	NoRespondedCount      *int64  `json:"NoRespondedCount,omitempty" xml:"NoRespondedCount,omitempty"`
	RespondedFailCount    *int64  `json:"RespondedFailCount,omitempty" xml:"RespondedFailCount,omitempty"`
	RespondedSuccessCount *int64  `json:"RespondedSuccessCount,omitempty" xml:"RespondedSuccessCount,omitempty"`
	SendDate              *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	TotalCount            *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendStatisticsResponseBodyDataTargetList) String() string {
	return tea.Prettify(s)
}

func (s QuerySendStatisticsResponseBodyDataTargetList) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetNoRespondedCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.NoRespondedCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetRespondedFailCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.RespondedFailCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetRespondedSuccessCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.RespondedSuccessCount = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetSendDate(v string) *QuerySendStatisticsResponseBodyDataTargetList {
	s.SendDate = &v
	return s
}

func (s *QuerySendStatisticsResponseBodyDataTargetList) SetTotalCount(v int64) *QuerySendStatisticsResponseBodyDataTargetList {
	s.TotalCount = &v
	return s
}

type QuerySendStatisticsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySendStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySendStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySendStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponse) SetHeaders(v map[string]*string) *QuerySendStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendStatisticsResponse) SetStatusCode(v int32) *QuerySendStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySendStatisticsResponse) SetBody(v *QuerySendStatisticsResponseBody) *QuerySendStatisticsResponse {
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryShortUrlResponse) SetStatusCode(v int32) *QueryShortUrlResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySmsSignResponse) SetStatusCode(v int32) *QuerySmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsSignResponse) SetBody(v *QuerySmsSignResponseBody) *QuerySmsSignResponse {
	s.Body = v
	return s
}

type QuerySmsSignListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageIndex            *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsSignListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignListRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListRequest) SetOwnerId(v int64) *QuerySmsSignListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsSignListRequest) SetPageIndex(v int32) *QuerySmsSignListRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySmsSignListRequest) SetPageSize(v int32) *QuerySmsSignListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmsSignListRequest) SetResourceOwnerAccount(v string) *QuerySmsSignListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsSignListRequest) SetResourceOwnerId(v int64) *QuerySmsSignListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySmsSignListResponseBody struct {
	Code        *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage *int32                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Message     *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsSignList []*QuerySmsSignListResponseBodySmsSignList `json:"SmsSignList,omitempty" xml:"SmsSignList,omitempty" type:"Repeated"`
	TotalCount  *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySmsSignListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBody) SetCode(v string) *QuerySmsSignListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetCurrentPage(v int32) *QuerySmsSignListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetMessage(v string) *QuerySmsSignListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetPageSize(v int32) *QuerySmsSignListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetRequestId(v string) *QuerySmsSignListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsSignListResponseBody) SetSmsSignList(v []*QuerySmsSignListResponseBodySmsSignList) *QuerySmsSignListResponseBody {
	s.SmsSignList = v
	return s
}

func (s *QuerySmsSignListResponseBody) SetTotalCount(v int64) *QuerySmsSignListResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySmsSignListResponseBodySmsSignList struct {
	AuditStatus  *string                                        `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	BusinessType *string                                        `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CreateDate   *string                                        `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	OrderId      *string                                        `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Reason       *QuerySmsSignListResponseBodySmsSignListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	SignName     *string                                        `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QuerySmsSignListResponseBodySmsSignList) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignListResponseBodySmsSignList) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetAuditStatus(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.AuditStatus = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetBusinessType(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.BusinessType = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetCreateDate(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetOrderId(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.OrderId = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetReason(v *QuerySmsSignListResponseBodySmsSignListReason) *QuerySmsSignListResponseBodySmsSignList {
	s.Reason = v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignList) SetSignName(v string) *QuerySmsSignListResponseBodySmsSignList {
	s.SignName = &v
	return s
}

type QuerySmsSignListResponseBodySmsSignListReason struct {
	RejectDate    *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	RejectInfo    *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	RejectSubInfo *string `json:"RejectSubInfo,omitempty" xml:"RejectSubInfo,omitempty"`
}

func (s QuerySmsSignListResponseBodySmsSignListReason) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignListResponseBodySmsSignListReason) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectDate(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectDate = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectInfo(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectInfo = &v
	return s
}

func (s *QuerySmsSignListResponseBodySmsSignListReason) SetRejectSubInfo(v string) *QuerySmsSignListResponseBodySmsSignListReason {
	s.RejectSubInfo = &v
	return s
}

type QuerySmsSignListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySmsSignListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySmsSignListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignListResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListResponse) SetHeaders(v map[string]*string) *QuerySmsSignListResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsSignListResponse) SetStatusCode(v int32) *QuerySmsSignListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsSignListResponse) SetBody(v *QuerySmsSignListResponseBody) *QuerySmsSignListResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySmsTemplateResponse) SetStatusCode(v int32) *QuerySmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsTemplateResponse) SetBody(v *QuerySmsTemplateResponseBody) *QuerySmsTemplateResponse {
	s.Body = v
	return s
}

type QuerySmsTemplateListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageIndex            *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateListRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListRequest) SetOwnerId(v int64) *QuerySmsTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetPageIndex(v int32) *QuerySmsTemplateListRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetPageSize(v int32) *QuerySmsTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetResourceOwnerAccount(v string) *QuerySmsTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetResourceOwnerId(v int64) *QuerySmsTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySmsTemplateListResponseBody struct {
	Code            *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage     *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Message         *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize        *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsTemplateList []*QuerySmsTemplateListResponseBodySmsTemplateList `json:"SmsTemplateList,omitempty" xml:"SmsTemplateList,omitempty" type:"Repeated"`
	TotalCount      *int64                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySmsTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBody) SetCode(v string) *QuerySmsTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetCurrentPage(v int32) *QuerySmsTemplateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetMessage(v string) *QuerySmsTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetPageSize(v int32) *QuerySmsTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetRequestId(v string) *QuerySmsTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetSmsTemplateList(v []*QuerySmsTemplateListResponseBodySmsTemplateList) *QuerySmsTemplateListResponseBody {
	s.SmsTemplateList = v
	return s
}

func (s *QuerySmsTemplateListResponseBody) SetTotalCount(v int64) *QuerySmsTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySmsTemplateListResponseBodySmsTemplateList struct {
	AuditStatus       *string                                                `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CreateDate        *string                                                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	OrderId           *string                                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OuterTemplateType *int32                                                 `json:"OuterTemplateType,omitempty" xml:"OuterTemplateType,omitempty"`
	Reason            *QuerySmsTemplateListResponseBodySmsTemplateListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	TemplateCode      *string                                                `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateContent   *string                                                `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	TemplateName      *string                                                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType      *int32                                                 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySmsTemplateListResponseBodySmsTemplateList) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateListResponseBodySmsTemplateList) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetAuditStatus(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.AuditStatus = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetCreateDate(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetOrderId(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.OrderId = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetOuterTemplateType(v int32) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.OuterTemplateType = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetReason(v *QuerySmsTemplateListResponseBodySmsTemplateListReason) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.Reason = v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateCode(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateCode = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateContent(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateContent = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateName(v string) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateName = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateList) SetTemplateType(v int32) *QuerySmsTemplateListResponseBodySmsTemplateList {
	s.TemplateType = &v
	return s
}

type QuerySmsTemplateListResponseBodySmsTemplateListReason struct {
	RejectDate    *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	RejectInfo    *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	RejectSubInfo *string `json:"RejectSubInfo,omitempty" xml:"RejectSubInfo,omitempty"`
}

func (s QuerySmsTemplateListResponseBodySmsTemplateListReason) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateListResponseBodySmsTemplateListReason) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectDate(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectDate = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectInfo(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectInfo = &v
	return s
}

func (s *QuerySmsTemplateListResponseBodySmsTemplateListReason) SetRejectSubInfo(v string) *QuerySmsTemplateListResponseBodySmsTemplateListReason {
	s.RejectSubInfo = &v
	return s
}

type QuerySmsTemplateListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySmsTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySmsTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListResponse) SetHeaders(v map[string]*string) *QuerySmsTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsTemplateListResponse) SetStatusCode(v int32) *QuerySmsTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsTemplateListResponse) SetBody(v *QuerySmsTemplateListResponseBody) *QuerySmsTemplateListResponse {
	s.Body = v
	return s
}

type SendBatchCardSmsRequest struct {
	CardTemplateCode         *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	CardTemplateParamJson    *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	DigitalTemplateCode      *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	DigitalTemplateParamJson *string `json:"DigitalTemplateParamJson,omitempty" xml:"DigitalTemplateParamJson,omitempty"`
	FallbackType             *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	OutId                    *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PhoneNumberJson          *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	SignNameJson             *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	SmsTemplateCode          *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParamJson     *string `json:"SmsTemplateParamJson,omitempty" xml:"SmsTemplateParamJson,omitempty"`
	SmsUpExtendCodeJson      *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
}

func (s SendBatchCardSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBatchCardSmsRequest) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsRequest) SetCardTemplateCode(v string) *SendBatchCardSmsRequest {
	s.CardTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetCardTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.CardTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetDigitalTemplateCode(v string) *SendBatchCardSmsRequest {
	s.DigitalTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetDigitalTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.DigitalTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetFallbackType(v string) *SendBatchCardSmsRequest {
	s.FallbackType = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetOutId(v string) *SendBatchCardSmsRequest {
	s.OutId = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetPhoneNumberJson(v string) *SendBatchCardSmsRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSignNameJson(v string) *SendBatchCardSmsRequest {
	s.SignNameJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsTemplateCode(v string) *SendBatchCardSmsRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.SmsTemplateParamJson = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetSmsUpExtendCodeJson(v string) *SendBatchCardSmsRequest {
	s.SmsUpExtendCodeJson = &v
	return s
}

type SendBatchCardSmsResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SendBatchCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendBatchCardSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBatchCardSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponseBody) SetCode(v string) *SendBatchCardSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetData(v *SendBatchCardSmsResponseBodyData) *SendBatchCardSmsResponseBody {
	s.Data = v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetRequestId(v string) *SendBatchCardSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendBatchCardSmsResponseBody) SetSuccess(v bool) *SendBatchCardSmsResponseBody {
	s.Success = &v
	return s
}

type SendBatchCardSmsResponseBodyData struct {
	BizCardId       *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	BizDigitalId    *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	BizSmsId        *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	CardTmpState    *int32  `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	MediaMobiles    *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
}

func (s SendBatchCardSmsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendBatchCardSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponseBodyData) SetBizCardId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizCardId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetBizDigitalId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizDigitalId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetBizSmsId(v string) *SendBatchCardSmsResponseBodyData {
	s.BizSmsId = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetCardTmpState(v int32) *SendBatchCardSmsResponseBodyData {
	s.CardTmpState = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetMediaMobiles(v string) *SendBatchCardSmsResponseBodyData {
	s.MediaMobiles = &v
	return s
}

func (s *SendBatchCardSmsResponseBodyData) SetNotMediaMobiles(v string) *SendBatchCardSmsResponseBodyData {
	s.NotMediaMobiles = &v
	return s
}

type SendBatchCardSmsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendBatchCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendBatchCardSmsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendBatchCardSmsResponse) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponse) SetHeaders(v map[string]*string) *SendBatchCardSmsResponse {
	s.Headers = v
	return s
}

func (s *SendBatchCardSmsResponse) SetStatusCode(v int32) *SendBatchCardSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBatchCardSmsResponse) SetBody(v *SendBatchCardSmsResponseBody) *SendBatchCardSmsResponse {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendBatchSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendBatchSmsResponse) SetStatusCode(v int32) *SendBatchSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBatchSmsResponse) SetBody(v *SendBatchSmsResponseBody) *SendBatchSmsResponse {
	s.Body = v
	return s
}

type SendCardSmsRequest struct {
	CardObjects          []*SendCardSmsRequestCardObjects `json:"CardObjects,omitempty" xml:"CardObjects,omitempty" type:"Repeated"`
	CardTemplateCode     *string                          `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	DigitalTemplateCode  *string                          `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	DigitalTemplateParam *string                          `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
	FallbackType         *string                          `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	OutId                *string                          `json:"OutId,omitempty" xml:"OutId,omitempty"`
	SignName             *string                          `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsTemplateCode      *string                          `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam     *string                          `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	SmsUpExtendCode      *string                          `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
}

func (s SendCardSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsRequest) GoString() string {
	return s.String()
}

func (s *SendCardSmsRequest) SetCardObjects(v []*SendCardSmsRequestCardObjects) *SendCardSmsRequest {
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

type SendCardSmsRequestCardObjects struct {
	CustomUrl  *string `json:"customUrl,omitempty" xml:"customUrl,omitempty"`
	DyncParams *string `json:"dyncParams,omitempty" xml:"dyncParams,omitempty"`
	Mobile     *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s SendCardSmsRequestCardObjects) String() string {
	return tea.Prettify(s)
}

func (s SendCardSmsRequestCardObjects) GoString() string {
	return s.String()
}

func (s *SendCardSmsRequestCardObjects) SetCustomUrl(v string) *SendCardSmsRequestCardObjects {
	s.CustomUrl = &v
	return s
}

func (s *SendCardSmsRequestCardObjects) SetDyncParams(v string) *SendCardSmsRequestCardObjects {
	s.DyncParams = &v
	return s
}

func (s *SendCardSmsRequestCardObjects) SetMobile(v string) *SendCardSmsRequestCardObjects {
	s.Mobile = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendCardSmsResponse) SetStatusCode(v int32) *SendCardSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCardSmsResponse) SetBody(v *SendCardSmsResponseBody) *SendCardSmsResponse {
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendSmsResponse) SetStatusCode(v int32) *SendSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSmsResponse) SetBody(v *SendSmsResponseBody) *SendSmsResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string                   `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetProdCode(v string) *TagResourcesRequest {
	s.ProdCode = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v string) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string   `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetProdCode(v string) *UntagResourcesRequest {
	s.ProdCode = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v string) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	if !tea.BoolValue(util.IsUnset(request.EffectiveDays)) {
		body["EffectiveDays"] = request.EffectiveDays
	}

	if !tea.BoolValue(util.IsUnset(request.ShortUrlName)) {
		body["ShortUrlName"] = request.ShortUrlName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		body["SourceUrl"] = request.SourceUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddShortUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignSource)) {
		query["SignSource"] = request.SignSource
	}

	if !tea.BoolValue(util.IsUnset(request.SignType)) {
		query["SignType"] = request.SignType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SignFileList)) {
		body["SignFileList"] = request.SignFileList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateContent)) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CheckMobilesCardSupportWithOptions(request *CheckMobilesCardSupportRequest, runtime *util.RuntimeOptions) (_result *CheckMobilesCardSupportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobiles)) {
		query["Mobiles"] = request.Mobiles
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckMobilesCardSupport"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMobilesCardSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMobilesCardSupport(request *CheckMobilesCardSupportRequest) (_result *CheckMobilesCardSupportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMobilesCardSupportResponse{}
	_body, _err := client.CheckMobilesCardSupportWithOptions(request, runtime)
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

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Factorys)) {
		query["Factorys"] = request.Factorys
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		query["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateShrink)) {
		query["Template"] = request.TemplateShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCardSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCardSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		body["SourceUrl"] = request.SourceUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteShortUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Action:      tea.String("DeleteSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetCardSmsLinkWithOptions(request *GetCardSmsLinkRequest, runtime *util.RuntimeOptions) (_result *GetCardSmsLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardCodeType)) {
		query["CardCodeType"] = request.CardCodeType
	}

	if !tea.BoolValue(util.IsUnset(request.CardLinkType)) {
		query["CardLinkType"] = request.CardLinkType
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateCode)) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateParamJson)) {
		query["CardTemplateParamJson"] = request.CardTemplateParamJson
	}

	if !tea.BoolValue(util.IsUnset(request.CustomShortCodeJson)) {
		query["CustomShortCodeJson"] = request.CustomShortCodeJson
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberJson)) {
		query["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !tea.BoolValue(util.IsUnset(request.SignNameJson)) {
		query["SignNameJson"] = request.SignNameJson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardSmsLink"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardSmsLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardSmsLink(request *GetCardSmsLinkRequest) (_result *GetCardSmsLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardSmsLinkResponse{}
	_body, _err := client.GetCardSmsLinkWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtendInfo)) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		query["FileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		query["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.OssKey)) {
		query["OssKey"] = request.OssKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMediaResourceId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaResourceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetOSSInfoForCardTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSInfoForCardTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignSource)) {
		query["SignSource"] = request.SignSource
	}

	if !tea.BoolValue(util.IsUnset(request.SignType)) {
		query["SignType"] = request.SignType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SignFileList)) {
		body["SignFileList"] = request.SignFileList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
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

	if !tea.BoolValue(util.IsUnset(request.TemplateContent)) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryCardSmsTemplateReportWithOptions(request *QueryCardSmsTemplateReportRequest, runtime *util.RuntimeOptions) (_result *QueryCardSmsTemplateReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCodes)) {
		query["TemplateCodes"] = request.TemplateCodes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCardSmsTemplateReport"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCardSmsTemplateReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCardSmsTemplateReport(request *QueryCardSmsTemplateReportRequest) (_result *QueryCardSmsTemplateReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCardSmsTemplateReportResponse{}
	_body, _err := client.QueryCardSmsTemplateReportWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDate)) {
		query["SendDate"] = request.SendDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySendDetails"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QuerySendStatisticsWithOptions(request *QuerySendStatisticsRequest, runtime *util.RuntimeOptions) (_result *QuerySendStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.IsGlobe)) {
		query["IsGlobe"] = request.IsGlobe
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySendStatistics"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySendStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySendStatistics(request *QuerySendStatisticsRequest) (_result *QuerySendStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySendStatisticsResponse{}
	_body, _err := client.QuerySendStatisticsWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.ShortUrl)) {
		body["ShortUrl"] = request.ShortUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryShortUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QuerySmsSignListWithOptions(request *QuerySmsSignListRequest, runtime *util.RuntimeOptions) (_result *QuerySmsSignListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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
		Action:      tea.String("QuerySmsSignList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsSignListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySmsSignList(request *QuerySmsSignListRequest) (_result *QuerySmsSignListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsSignListResponse{}
	_body, _err := client.QuerySmsSignListWithOptions(request, runtime)
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
		Action:      tea.String("QuerySmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QuerySmsTemplateListWithOptions(request *QuerySmsTemplateListRequest, runtime *util.RuntimeOptions) (_result *QuerySmsTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
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
		Action:      tea.String("QuerySmsTemplateList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySmsTemplateList(request *QuerySmsTemplateListRequest) (_result *QuerySmsTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsTemplateListResponse{}
	_body, _err := client.QuerySmsTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendBatchCardSmsWithOptions(request *SendBatchCardSmsRequest, runtime *util.RuntimeOptions) (_result *SendBatchCardSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardTemplateCode)) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateParamJson)) {
		query["CardTemplateParamJson"] = request.CardTemplateParamJson
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalTemplateCode)) {
		query["DigitalTemplateCode"] = request.DigitalTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalTemplateParamJson)) {
		query["DigitalTemplateParamJson"] = request.DigitalTemplateParamJson
	}

	if !tea.BoolValue(util.IsUnset(request.FallbackType)) {
		query["FallbackType"] = request.FallbackType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberJson)) {
		query["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !tea.BoolValue(util.IsUnset(request.SignNameJson)) {
		query["SignNameJson"] = request.SignNameJson
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateCode)) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateParamJson)) {
		query["SmsTemplateParamJson"] = request.SmsTemplateParamJson
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCodeJson)) {
		query["SmsUpExtendCodeJson"] = request.SmsUpExtendCodeJson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendBatchCardSms"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendBatchCardSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendBatchCardSms(request *SendBatchCardSmsRequest) (_result *SendBatchCardSmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendBatchCardSmsResponse{}
	_body, _err := client.SendBatchCardSmsWithOptions(request, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumberJson)) {
		body["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !tea.BoolValue(util.IsUnset(request.SignNameJson)) {
		body["SignNameJson"] = request.SignNameJson
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCodeJson)) {
		body["SmsUpExtendCodeJson"] = request.SmsUpExtendCodeJson
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParamJson)) {
		body["TemplateParamJson"] = request.TemplateParamJson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendBatchSms"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendBatchSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SendCardSmsWithOptions(request *SendCardSmsRequest, runtime *util.RuntimeOptions) (_result *SendCardSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardObjects)) {
		query["CardObjects"] = request.CardObjects
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateCode)) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalTemplateCode)) {
		query["DigitalTemplateCode"] = request.DigitalTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalTemplateParam)) {
		query["DigitalTemplateParam"] = request.DigitalTemplateParam
	}

	if !tea.BoolValue(util.IsUnset(request.FallbackType)) {
		query["FallbackType"] = request.FallbackType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateCode)) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateParam)) {
		query["SmsTemplateParam"] = request.SmsTemplateParam
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCode)) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCardSms"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCardSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SendSmsWithOptions(request *SendSmsRequest, runtime *util.RuntimeOptions) (_result *SendSmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCode)) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParam)) {
		query["TemplateParam"] = request.TemplateParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendSms"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
