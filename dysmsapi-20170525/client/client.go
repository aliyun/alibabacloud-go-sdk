// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddExtCodeSignRequest struct {
	// 扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExtCode              *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AddExtCodeSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AddExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignRequest) SetExtCode(v string) *AddExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *AddExtCodeSignRequest) SetOwnerId(v int64) *AddExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *AddExtCodeSignRequest) SetResourceOwnerAccount(v string) *AddExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddExtCodeSignRequest) SetResourceOwnerId(v int64) *AddExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddExtCodeSignRequest) SetSignName(v string) *AddExtCodeSignRequest {
	s.SignName = &v
	return s
}

type AddExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddExtCodeSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *AddExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetCode(v string) *AddExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetData(v bool) *AddExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetMessage(v string) *AddExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetRequestId(v string) *AddExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetSuccess(v bool) *AddExtCodeSignResponseBody {
	s.Success = &v
	return s
}

type AddExtCodeSignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddExtCodeSignResponse) String() string {
	return tea.Prettify(s)
}

func (s AddExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignResponse) SetHeaders(v map[string]*string) *AddExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *AddExtCodeSignResponse) SetStatusCode(v int32) *AddExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *AddExtCodeSignResponse) SetBody(v *AddExtCodeSignResponseBody) *AddExtCodeSignResponse {
	s.Body = v
	return s
}

type AddShortUrlRequest struct {
	// The validity period of the short URL. Unit: days. The maximum validity period is 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	EffectiveDays        *string `json:"EffectiveDays,omitempty" xml:"EffectiveDays,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name of the short URL. The name cannot exceed 13 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// The Alibaba Cloud Short Link service.
	ShortUrlName *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	// The source URL. The URL cannot exceed 1,000 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the short URL.
	Data *AddShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The time when the short URL expires.
	//
	// > The value of **ExpireDate*	- is on the hour.
	//
	// example:
	//
	// 2021-09-19 00:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The short URL.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	// The source URL.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the signature application. The description cannot exceed 200 characters in length. The description is one of the reference information for signature review. We recommend that you describe the use scenarios of your services in detail, and provide information that can verify the services, such as a website URL, a domain name with an ICP filing, an app download URL, an official account name, or a mini program name. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is the abbreviation of our company.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature files.
	//
	// This parameter is required.
	SignFileList []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The name of the signature.
	//
	// >
	//
	// 	- The signature name is not case-sensitive. For example, [Alibaba Cloud Communication] and [alibaba cloud communication] are considered as the same name.
	//
	// 	- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The source of the signature. Valid values:
	//
	// 	- **0**: the full name or abbreviation of an enterprise or institution
	//
	// 	- **1**: the full name or abbreviation of a website that has obtained an ICP filing from the Ministry of Industry and Information Technology (MIIT) of China
	//
	// 	- **2**: the full name or abbreviation of an app
	//
	// 	- **3**: the full name or abbreviation of an official account or mini-program
	//
	// 	- **4**: the full name or abbreviation of an e-commerce store
	//
	// 	- **5**: the full name or abbreviation of a trademark
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The type of the signature. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: general-purpose
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
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
	// The Base64-encoded string of the qualification document. An image cannot exceed 2 MB in size. In some scenarios, you must upload supporting documents to apply for signatures. For more information, see [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uL****
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// The format of the qualification document. You can upload multiple images. Images in JPG, PNG, GIF, or JPEG format are supported.
	//
	// In some scenarios, you must upload supporting documents to apply for signatures. For more information, see [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// > If you apply for a signature for other users or if the signature source is the name of an enterprise or public institution, you must upload a certificate and a letter of authorization. For more information, see [Certificate](https://help.aliyun.com/document_detail/108076.html) and [Letter of authorization](https://help.aliyun.com/document_detail/56741.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the message template. It is one of the reference information for template review. The description cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Apply for a template to send verification codes.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The content of the template. The content can be up to 500 characters in length. For more information, see [Message template specifications](https://help.aliyun.com/document_detail/108253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// > Only enterprise users can send promotional messages, or send messages to countries or regions outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_15255****
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The list of mobile phone numbers that receive messages.
	//
	// This parameter is required.
	Mobiles []map[string]interface{} `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CheckMobilesCardSupportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CheckMobilesCardSupportResponseBody) SetData(v *CheckMobilesCardSupportResponseBodyData) *CheckMobilesCardSupportResponseBody {
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

type CheckMobilesCardSupportResponseBodyData struct {
	// The list of returned results.
	QueryResult []*CheckMobilesCardSupportResponseBodyDataQueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty" type:"Repeated"`
}

func (s CheckMobilesCardSupportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBodyData) SetQueryResult(v []*CheckMobilesCardSupportResponseBodyDataQueryResult) *CheckMobilesCardSupportResponseBodyData {
	s.QueryResult = v
	return s
}

type CheckMobilesCardSupportResponseBodyDataQueryResult struct {
	// The mobile phone number.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// Indicates whether the mobile phone number supports card messages.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Support *bool `json:"support,omitempty" xml:"support,omitempty"`
}

func (s CheckMobilesCardSupportResponseBodyDataQueryResult) String() string {
	return tea.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBodyDataQueryResult) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) SetMobile(v string) *CheckMobilesCardSupportResponseBodyDataQueryResult {
	s.Mobile = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) SetSupport(v bool) *CheckMobilesCardSupportResponseBodyDataQueryResult {
	s.Support = &v
	return s
}

type CheckMobilesCardSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMobilesCardSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ConversionDataIntlRequest struct {
	// The conversion rate.
	//
	// > The value of this parameter is a double, and ranges from 0 to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.53
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time point at which the conversion rate is monitored. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > If you do not specify this parameter, the current timestamp is used by default.
	//
	// example:
	//
	// 1349055900000
	ReportTime           *int64  `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ConversionDataIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataIntlRequest) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlRequest) SetConversionRate(v string) *ConversionDataIntlRequest {
	s.ConversionRate = &v
	return s
}

func (s *ConversionDataIntlRequest) SetOwnerId(v int64) *ConversionDataIntlRequest {
	s.OwnerId = &v
	return s
}

func (s *ConversionDataIntlRequest) SetReportTime(v int64) *ConversionDataIntlRequest {
	s.ReportTime = &v
	return s
}

func (s *ConversionDataIntlRequest) SetResourceOwnerAccount(v string) *ConversionDataIntlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConversionDataIntlRequest) SetResourceOwnerId(v int64) *ConversionDataIntlRequest {
	s.ResourceOwnerId = &v
	return s
}

type ConversionDataIntlResponseBody struct {
	// The status code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.101345.0.0.74326ff2J5EZyt).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConversionDataIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataIntlResponseBody) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlResponseBody) SetCode(v string) *ConversionDataIntlResponseBody {
	s.Code = &v
	return s
}

func (s *ConversionDataIntlResponseBody) SetMessage(v string) *ConversionDataIntlResponseBody {
	s.Message = &v
	return s
}

func (s *ConversionDataIntlResponseBody) SetRequestId(v string) *ConversionDataIntlResponseBody {
	s.RequestId = &v
	return s
}

type ConversionDataIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversionDataIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversionDataIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataIntlResponse) GoString() string {
	return s.String()
}

func (s *ConversionDataIntlResponse) SetHeaders(v map[string]*string) *ConversionDataIntlResponse {
	s.Headers = v
	return s
}

func (s *ConversionDataIntlResponse) SetStatusCode(v int32) *ConversionDataIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversionDataIntlResponse) SetBody(v *ConversionDataIntlResponseBody) *ConversionDataIntlResponse {
	s.Body = v
	return s
}

type CreateCardSmsTemplateRequest struct {
	// The mobile phone manufacturer. Valid values:
	//
	// 	- **HuaWei**: HUAWEI
	//
	// 	- **XiaoMi**: Xiaomi
	//
	// 	- **OPPO**: OPPO
	//
	// 	- **VIVO**: vivo
	//
	// 	- **MEIZU**: MEIZU
	//
	// > If this parameter is not specified, the system automatically specifies a supported mobile phone manufacturer.
	//
	// example:
	//
	// XiaoMi
	Factorys *string `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
	// The description of the message template.
	//
	// example:
	//
	// Image and Text Template
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The content of the card message template.
	//
	// >
	//
	// 	- For information about fields such as Template, ExtendInfo, TemplateContent, TmpCard, and Action, see [Parameters of card message templates](https://help.aliyun.com/document_detail/434929.html).
	//
	// 	- Message template content varies based on the template type. For more information, see [Sample message templates](https://help.aliyun.com/document_detail/435361.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//        "extendInfo":{
	//
	//               "scene":"HMOVM",
	//
	//               "purpose":"2",
	//
	//               "userExt":{
	//
	//                      "outId":"1234554321"
	//
	//               }
	//
	//        },
	//
	//        "templateContent":{
	//
	//               "pages":[
	//
	//                      {
	//
	// "tmpCards":[
	//
	//                                    {
	//
	//                                           "type":"IMAGE",
	//
	//                                           "srcType":1,
	//
	//                                           "src":"28755",
	//
	//                                           "actionType":"OPEN_APP",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://s.tb.cn/c.KxzZ",
	//
	//                                                  "merchantName":"test-template",
	//
	//                                                  "packageName":[
	//
	//                                                         "com.taobao.taobao"],
	//
	//                                                  "floorUrl":"https://s.tb.cn/c.KxzZ"
	//
	//                                           },
	//
	//                                           "positionNumber":1
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"this is a test msg.",
	//
	//                                           "isTextTitle":true,
	//
	//                                           "positionNumber":2
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"Promotional information",
	//
	//                                           "isTextTitle":false,
	//
	//                                           "positionNumber":3
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"BUTTON",
	//
	//                                           "content":"Promotional information,",
	//
	//                                           "actionType":"OPEN_BROWSER",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://www.aliyun.com",
	//
	//                                                  "merchantName":"Currently on the Alibaba Cloud official website."
	//
	// },
	//
	//                                           "positionNumber":4
	//
	//                                    }]
	//
	//                      }]
	//
	//        },
	//
	//        "cardSignName":"aliyun",
	//
	//        "cardType":5
	//
	// }
	Template map[string]interface{} `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the card message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Image and Text Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	// The mobile phone manufacturer. Valid values:
	//
	// 	- **HuaWei**: HUAWEI
	//
	// 	- **XiaoMi**: Xiaomi
	//
	// 	- **OPPO**: OPPO
	//
	// 	- **VIVO**: vivo
	//
	// 	- **MEIZU**: MEIZU
	//
	// > If this parameter is not specified, the system automatically specifies a supported mobile phone manufacturer.
	//
	// example:
	//
	// XiaoMi
	Factorys *string `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
	// The description of the message template.
	//
	// example:
	//
	// Image and Text Template
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The content of the card message template.
	//
	// >
	//
	// 	- For information about fields such as Template, ExtendInfo, TemplateContent, TmpCard, and Action, see [Parameters of card message templates](https://help.aliyun.com/document_detail/434929.html).
	//
	// 	- Message template content varies based on the template type. For more information, see [Sample message templates](https://help.aliyun.com/document_detail/435361.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//        "extendInfo":{
	//
	//               "scene":"HMOVM",
	//
	//               "purpose":"2",
	//
	//               "userExt":{
	//
	//                      "outId":"1234554321"
	//
	//               }
	//
	//        },
	//
	//        "templateContent":{
	//
	//               "pages":[
	//
	//                      {
	//
	// "tmpCards":[
	//
	//                                    {
	//
	//                                           "type":"IMAGE",
	//
	//                                           "srcType":1,
	//
	//                                           "src":"28755",
	//
	//                                           "actionType":"OPEN_APP",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://s.tb.cn/c.KxzZ",
	//
	//                                                  "merchantName":"test-template",
	//
	//                                                  "packageName":[
	//
	//                                                         "com.taobao.taobao"],
	//
	//                                                  "floorUrl":"https://s.tb.cn/c.KxzZ"
	//
	//                                           },
	//
	//                                           "positionNumber":1
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"this is a test msg.",
	//
	//                                           "isTextTitle":true,
	//
	//                                           "positionNumber":2
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"Promotional information",
	//
	//                                           "isTextTitle":false,
	//
	//                                           "positionNumber":3
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"BUTTON",
	//
	//                                           "content":"Promotional information,",
	//
	//                                           "actionType":"OPEN_BROWSER",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://www.aliyun.com",
	//
	//                                                  "merchantName":"Currently on the Alibaba Cloud official website."
	//
	// },
	//
	//                                           "positionNumber":4
	//
	//                                    }]
	//
	//                      }]
	//
	//        },
	//
	//        "cardSignName":"aliyun",
	//
	//        "cardType":5
	//
	// }
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the card message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Image and Text Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CreateCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The code of the message template.
	//
	// You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview).
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// CARD_SMS_60000****
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateSmartShortUrlRequest struct {
	// example:
	//
	// 示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15900195***
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s CreateSmartShortUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartShortUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlRequest) SetOutId(v string) *CreateSmartShortUrlRequest {
	s.OutId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetOwnerId(v int64) *CreateSmartShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetPhoneNumbers(v string) *CreateSmartShortUrlRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetResourceOwnerAccount(v string) *CreateSmartShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetResourceOwnerId(v int64) *CreateSmartShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetSourceUrl(v string) *CreateSmartShortUrlRequest {
	s.SourceUrl = &v
	return s
}

type CreateSmartShortUrlResponseBody struct {
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*CreateSmartShortUrlResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmartShortUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartShortUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponseBody) SetCode(v string) *CreateSmartShortUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetMessage(v string) *CreateSmartShortUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetModel(v []*CreateSmartShortUrlResponseBodyModel) *CreateSmartShortUrlResponseBody {
	s.Model = v
	return s
}

func (s *CreateSmartShortUrlResponseBody) SetRequestId(v string) *CreateSmartShortUrlResponseBody {
	s.RequestId = &v
	return s
}

type CreateSmartShortUrlResponseBodyModel struct {
	// example:
	//
	// 示例值
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 11
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 示例值
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// example:
	//
	// 示例值示例值
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s CreateSmartShortUrlResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartShortUrlResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponseBodyModel) SetDomain(v string) *CreateSmartShortUrlResponseBodyModel {
	s.Domain = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetExpiration(v int64) *CreateSmartShortUrlResponseBodyModel {
	s.Expiration = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetPhoneNumber(v string) *CreateSmartShortUrlResponseBodyModel {
	s.PhoneNumber = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetShortName(v string) *CreateSmartShortUrlResponseBodyModel {
	s.ShortName = &v
	return s
}

func (s *CreateSmartShortUrlResponseBodyModel) SetShortUrl(v string) *CreateSmartShortUrlResponseBodyModel {
	s.ShortUrl = &v
	return s
}

type CreateSmartShortUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartShortUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmartShortUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlResponse) SetHeaders(v map[string]*string) *CreateSmartShortUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartShortUrlResponse) SetStatusCode(v int32) *CreateSmartShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartShortUrlResponse) SetBody(v *CreateSmartShortUrlResponseBody) *CreateSmartShortUrlResponse {
	s.Body = v
	return s
}

type CreateSmsSignRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, enter the domain name with HTTP or HTTPS that has been registered with the MIIT.
	//
	// - For launched apps, provide a display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, input the full name, ensuring they are online.
	//
	// - For e-commerce platform store names, applicable only to enterprise users, provide a display link with HTTP or HTTPS for the store.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// Additional information to supplement uploaded business proof documents or screenshots, which helps reviewers understand your business details.
	//
	// This parameter is optional; please fill it out based on your actual needs.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [Apply for Qualification](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// > The scenario explanation is one of the reference materials for signature review. Please provide a detailed description of the usage scenarios for your live services, along with links to verify these services such as website URLs with MIIT备案, app store display links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A comprehensive application explanation enhances the efficiency of signature and template reviews. Refer to the **Application Scenario*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// SMS signature for the login scenario using verification code.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature name. Please adhere to the [Signature Specifications](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-0p8-qn8-mmy).
	//
	// > - Signature names are case-insensitive; e.g., 【Aliyun Communication】 and 【aliyun communication】 are considered identical.
	//
	// > - If your verification code signature and general signature names are the same, the system defaults to using the general signature for sending SMS messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Signature source. Values:
	//
	// - **0**: Full name or abbreviation of an enterprise or institution.
	//
	// - **1**: Full name or abbreviation of a MIIT-registered website.
	//
	// - **2**: Full name or abbreviation of an App.
	//
	// - **3**: Full name or abbreviation of an official account or mini-program.
	//
	// - **4**: Full name or abbreviation of an e-commerce platform store.
	//
	// - **5**: Full name or abbreviation of a trademark.
	//
	// For detailed information on signature sources, refer to [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-xup-k46-yi4).
	//
	// > This interface does not support applying for signatures with sources as **Test or Learning*	- and **Trial Use**. If you need to apply for signatures with these sources, please go to the [SMS Service Console](https://dysms.console.aliyun.com/domestic/text/sign/add/qualification).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. Values:
	//
	// - **0**: Verification Code
	//
	// - **1**: General (Default)
	//
	// > It is recommended to use the default value: **General**.
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Choose whether the applied signature is for self-use or third-party use.
	//
	// - false: Self-use (default)
	//
	// - true: Third-party use
	//
	// 	Notice: Please select self-use qualification ID when the signature is for self-use; choose third-party use qualification ID when it\\"s for third-party use.
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s CreateSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) SetApplySceneContent(v string) *CreateSmsSignRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsSignRequest) SetMoreData(v []*string) *CreateSmsSignRequest {
	s.MoreData = v
	return s
}

func (s *CreateSmsSignRequest) SetOwnerId(v int64) *CreateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetQualificationId(v int64) *CreateSmsSignRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateSmsSignRequest) SetRemark(v string) *CreateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerAccount(v string) *CreateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerId(v int64) *CreateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignName(v string) *CreateSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignSource(v int32) *CreateSmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignType(v int32) *CreateSmsSignRequest {
	s.SignType = &v
	return s
}

func (s *CreateSmsSignRequest) SetThirdParty(v bool) *CreateSmsSignRequest {
	s.ThirdParty = &v
	return s
}

type CreateSmsSignShrinkRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, enter the domain name with HTTP or HTTPS that has been registered with the MIIT.
	//
	// - For launched apps, provide a display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, input the full name, ensuring they are online.
	//
	// - For e-commerce platform store names, applicable only to enterprise users, provide a display link with HTTP or HTTPS for the store.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// Additional information to supplement uploaded business proof documents or screenshots, which helps reviewers understand your business details.
	//
	// This parameter is optional; please fill it out based on your actual needs.
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [Apply for Qualification](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// > The scenario explanation is one of the reference materials for signature review. Please provide a detailed description of the usage scenarios for your live services, along with links to verify these services such as website URLs with MIIT备案, app store display links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A comprehensive application explanation enhances the efficiency of signature and template reviews. Refer to the **Application Scenario*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// SMS signature for the login scenario using verification code.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature name. Please adhere to the [Signature Specifications](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-0p8-qn8-mmy).
	//
	// > - Signature names are case-insensitive; e.g., 【Aliyun Communication】 and 【aliyun communication】 are considered identical.
	//
	// > - If your verification code signature and general signature names are the same, the system defaults to using the general signature for sending SMS messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Signature source. Values:
	//
	// - **0**: Full name or abbreviation of an enterprise or institution.
	//
	// - **1**: Full name or abbreviation of a MIIT-registered website.
	//
	// - **2**: Full name or abbreviation of an App.
	//
	// - **3**: Full name or abbreviation of an official account or mini-program.
	//
	// - **4**: Full name or abbreviation of an e-commerce platform store.
	//
	// - **5**: Full name or abbreviation of a trademark.
	//
	// For detailed information on signature sources, refer to [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-xup-k46-yi4).
	//
	// > This interface does not support applying for signatures with sources as **Test or Learning*	- and **Trial Use**. If you need to apply for signatures with these sources, please go to the [SMS Service Console](https://dysms.console.aliyun.com/domestic/text/sign/add/qualification).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. Values:
	//
	// - **0**: Verification Code
	//
	// - **1**: General (Default)
	//
	// > It is recommended to use the default value: **General**.
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Choose whether the applied signature is for self-use or third-party use.
	//
	// - false: Self-use (default)
	//
	// - true: Third-party use
	//
	// 	Notice: Please select self-use qualification ID when the signature is for self-use; choose third-party use qualification ID when it\\"s for third-party use.
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s CreateSmsSignShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignShrinkRequest) SetApplySceneContent(v string) *CreateSmsSignShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetMoreDataShrink(v string) *CreateSmsSignShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetOwnerId(v int64) *CreateSmsSignShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetQualificationId(v int64) *CreateSmsSignShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetRemark(v string) *CreateSmsSignShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetResourceOwnerAccount(v string) *CreateSmsSignShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetResourceOwnerId(v int64) *CreateSmsSignShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetSignName(v string) *CreateSmsSignShrinkRequest {
	s.SignName = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetSignSource(v int32) *CreateSmsSignShrinkRequest {
	s.SignSource = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetSignType(v int32) *CreateSmsSignShrinkRequest {
	s.SignType = &v
	return s
}

func (s *CreateSmsSignShrinkRequest) SetThirdParty(v bool) *CreateSmsSignShrinkRequest {
	s.ThirdParty = &v
	return s
}

type CreateSmsSignResponseBody struct {
	// Request status code.
	//
	// - OK indicates a successful request.
	//
	// - For other error codes, refer to the [Error Code List](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Work order number.
	//
	// This parameter is used by auditors when querying the audit. You will need to provide this work order number if you require expedited review.
	//
	// example:
	//
	// 2004415****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request and can be used for troubleshooting and issue localization.
	//
	// example:
	//
	// CCA2BCFF-2BA7-427C-90EE-AC6994748607
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Signature name.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s CreateSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBody) SetCode(v string) *CreateSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetMessage(v string) *CreateSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetOrderId(v string) *CreateSmsSignResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetRequestId(v string) *CreateSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetSignName(v string) *CreateSmsSignResponseBody {
	s.SignName = &v
	return s
}

type CreateSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponse) SetHeaders(v map[string]*string) *CreateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsSignResponse) SetStatusCode(v int32) *CreateSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsSignResponse) SetBody(v *CreateSmsSignResponseBody) *CreateSmsSignResponse {
	s.Body = v
	return s
}

type CreateSmsTemplateRequest struct {
	// If there is an applicable scenario, you can fill it in.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional message.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional materials you can upload, such as business proof documents or screenshots, to help reviewers understand your business details.
	//
	// This parameter is optional; please fill it in according to actual needs.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The signature name that the template needs to be associated with. The associated SMS signature must have passed the review.
	//
	// This parameter is mandatory when the TemplateType parameter is **0**, **1**, or **2**.
	//
	// <notice>Associating a signature can expedite the review process. Note that this associated signature is unrelated to the signature selected when sending SMS messages.</notice>
	//
	// example:
	//
	// Aliyun
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Please describe the business scenario where you use SMS or provide an online link to the scenario, along with a complete example of the SMS (with variable contents filled), as complete information helps increase the template approval rate. Failure to follow guidelines or leaving this field blank may affect the approval of your template.
	//
	// example:
	//
	// Request verification code SMS.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS specifications; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. For variable specifications, see [TemplateContent Variable Parameter Filling Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}. It is valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyunCode
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For filling in variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional message.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For details on the differences between personal and enterprise user rights, please refer to [Usage Instructions](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateRequest) SetApplySceneContent(v string) *CreateSmsTemplateRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetIntlType(v int32) *CreateSmsTemplateRequest {
	s.IntlType = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetMoreData(v []*string) *CreateSmsTemplateRequest {
	s.MoreData = v
	return s
}

func (s *CreateSmsTemplateRequest) SetOwnerId(v int64) *CreateSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetRelatedSignName(v string) *CreateSmsTemplateRequest {
	s.RelatedSignName = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetRemark(v string) *CreateSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetResourceOwnerAccount(v string) *CreateSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetResourceOwnerId(v int64) *CreateSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateContent(v string) *CreateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateName(v string) *CreateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateRule(v string) *CreateSmsTemplateRequest {
	s.TemplateRule = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateType(v int32) *CreateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateSmsTemplateShrinkRequest struct {
	// If there is an applicable scenario, you can fill it in.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional message.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional materials you can upload, such as business proof documents or screenshots, to help reviewers understand your business details.
	//
	// This parameter is optional; please fill it in according to actual needs.
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The signature name that the template needs to be associated with. The associated SMS signature must have passed the review.
	//
	// This parameter is mandatory when the TemplateType parameter is **0**, **1**, or **2**.
	//
	// <notice>Associating a signature can expedite the review process. Note that this associated signature is unrelated to the signature selected when sending SMS messages.</notice>
	//
	// example:
	//
	// Aliyun
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Please describe the business scenario where you use SMS or provide an online link to the scenario, along with a complete example of the SMS (with variable contents filled), as complete information helps increase the template approval rate. Failure to follow guidelines or leaving this field blank may affect the approval of your template.
	//
	// example:
	//
	// Request verification code SMS.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS specifications; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. For variable specifications, see [TemplateContent Variable Parameter Filling Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}. It is valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyunCode
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For filling in variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional message.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For details on the differences between personal and enterprise user rights, please refer to [Usage Instructions](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateSmsTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateShrinkRequest) SetApplySceneContent(v string) *CreateSmsTemplateShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetIntlType(v int32) *CreateSmsTemplateShrinkRequest {
	s.IntlType = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetMoreDataShrink(v string) *CreateSmsTemplateShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetOwnerId(v int64) *CreateSmsTemplateShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetRelatedSignName(v string) *CreateSmsTemplateShrinkRequest {
	s.RelatedSignName = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetRemark(v string) *CreateSmsTemplateShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetResourceOwnerAccount(v string) *CreateSmsTemplateShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetResourceOwnerId(v int64) *CreateSmsTemplateShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateContent(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateName(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateRule(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateRule = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateType(v int32) *CreateSmsTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type CreateSmsTemplateResponseBody struct {
	// Request status code.
	//
	// 	- OK indicates a successful request.
	//
	// 	- For other error codes, refer to the **Error Codes*	- section of this chapter or the product\\"s [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Work order ID.
	//
	// This parameter is used by auditors when querying audits. If you need expedited review, you must provide this work order number.
	//
	// example:
	//
	// 2005020****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID generated by Alibaba Cloud for this request, which is a unique identifier that can be used for troubleshooting and issue定位.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SMS template code.
	//
	// After submitting the template application, you can use the SMS template code to query the template audit details via the [GetSmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-getsmstemplate?) API. You can also [configure delivery receipts](https://help.aliyun.com/zh/sms/developer-reference/configure-delivery-receipts-1?spm), and obtain the template audit status messages through TemplateSmsReport.
	//
	// example:
	//
	// SMS_10000****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// SMS template name.
	//
	// example:
	//
	// aliyunCode
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponseBody) SetCode(v string) *CreateSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetMessage(v string) *CreateSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetOrderId(v string) *CreateSmsTemplateResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetRequestId(v string) *CreateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetTemplateCode(v string) *CreateSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *CreateSmsTemplateResponseBody) SetTemplateName(v string) *CreateSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

type CreateSmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsTemplateResponse) SetStatusCode(v int32) *CreateSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsTemplateResponse) SetBody(v *CreateSmsTemplateResponseBody) *CreateSmsTemplateResponse {
	s.Body = v
	return s
}

type DeleteExtCodeSignRequest struct {
	// 扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExtCode              *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DeleteExtCodeSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignRequest) SetExtCode(v string) *DeleteExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetOwnerId(v int64) *DeleteExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetResourceOwnerAccount(v string) *DeleteExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetResourceOwnerId(v int64) *DeleteExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteExtCodeSignRequest) SetSignName(v string) *DeleteExtCodeSignRequest {
	s.SignName = &v
	return s
}

type DeleteExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExtCodeSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *DeleteExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetCode(v string) *DeleteExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetData(v bool) *DeleteExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetMessage(v string) *DeleteExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetRequestId(v string) *DeleteExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetSuccess(v bool) *DeleteExtCodeSignResponseBody {
	s.Success = &v
	return s
}

type DeleteExtCodeSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExtCodeSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignResponse) SetHeaders(v map[string]*string) *DeleteExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteExtCodeSignResponse) SetStatusCode(v int32) *DeleteExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExtCodeSignResponse) SetBody(v *DeleteExtCodeSignResponseBody) *DeleteExtCodeSignResponse {
	s.Body = v
	return s
}

type DeleteShortUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source address. The address can be up to 1,000 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The signature.
	//
	// > The signature must be submitted by the current Alibaba Cloud account, and has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The code of the message template.
	//
	// You can log on to the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm) and obtain the message template code on the **Message Templates*	- tab. You can also obtain the message template code by calling the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CCA2BCFF-2BA7-427C-90EE-AC6994748607
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_20375****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetCardSmsDetailsRequest struct {
	// example:
	//
	// 123456^0
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// example:
	//
	// 12346^0
	BizDigitId *string `json:"BizDigitId,omitempty" xml:"BizDigitId,omitempty"`
	// example:
	//
	// 1234576^0
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20240112
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
}

func (s GetCardSmsDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsRequest) SetBizCardId(v string) *GetCardSmsDetailsRequest {
	s.BizCardId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetBizDigitId(v string) *GetCardSmsDetailsRequest {
	s.BizDigitId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetBizSmsId(v string) *GetCardSmsDetailsRequest {
	s.BizSmsId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetCurrentPage(v int64) *GetCardSmsDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetOwnerId(v int64) *GetCardSmsDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetPageSize(v int64) *GetCardSmsDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetPhoneNumber(v string) *GetCardSmsDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetResourceOwnerAccount(v string) *GetCardSmsDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetResourceOwnerId(v int64) *GetCardSmsDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetSendDate(v string) *GetCardSmsDetailsRequest {
	s.SendDate = &v
	return s
}

type GetCardSmsDetailsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 卡片短信发送结果
	CardSendDetailDTO *GetCardSmsDetailsResponseBodyCardSendDetailDTO `json:"CardSendDetailDTO,omitempty" xml:"CardSendDetailDTO,omitempty" type:"Struct"`
	// 状态码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态描述
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardSmsDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBody) SetAccessDeniedDetail(v string) *GetCardSmsDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetCardSendDetailDTO(v *GetCardSmsDetailsResponseBodyCardSendDetailDTO) *GetCardSmsDetailsResponseBody {
	s.CardSendDetailDTO = v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetCode(v string) *GetCardSmsDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetMessage(v string) *GetCardSmsDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCardSmsDetailsResponseBody) SetSuccess(v bool) *GetCardSmsDetailsResponseBody {
	s.Success = &v
	return s
}

type GetCardSmsDetailsResponseBodyCardSendDetailDTO struct {
	// 页码
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 页数
	//
	// example:
	//
	// 10
	PageSize *int64                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*GetCardSmsDetailsResponseBodyCardSendDetailDTORecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总量
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTO) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetCurrentPage(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.CurrentPage = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetPageSize(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.PageSize = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetRecords(v []*GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.Records = v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTO) SetTotalCount(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTO {
	s.TotalCount = &v
	return s
}

type GetCardSmsDetailsResponseBodyCardSendDetailDTORecords struct {
	// 发送错误码
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// 客户传输outId
	//
	// example:
	//
	// 12345678
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// 接收短信手机号
	//
	// example:
	//
	// 156****9080
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// 接收时间
	//
	// example:
	//
	// 2024-09-27 11:26:35
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// 接收短信类型
	//
	// example:
	//
	// CARD_SMS
	ReceiveType *string `json:"ReceiveType,omitempty" xml:"ReceiveType,omitempty"`
	// 渲染时间
	//
	// example:
	//
	// 2024-09-27 12:13:39
	RenderDate *string `json:"RenderDate,omitempty" xml:"RenderDate,omitempty"`
	// 解析状态.。0：未解析；1：解析成功；3：未解析
	//
	// example:
	//
	// 1
	RenderStatus *int64 `json:"RenderStatus,omitempty" xml:"RenderStatus,omitempty"`
	// 短信发送时间
	//
	// example:
	//
	// 2024-09-27 11:26:32
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// 发送状态 1：发送中；2：发送失败；3：发送成功；4：寻址失败
	//
	// example:
	//
	// 3
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// 短信内容。只有文本短信有值
	//
	// example:
	//
	// 您收到一条短信消息
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	// 模板code
	//
	// example:
	//
	// CARD_SMS_6***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetErrCode(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ErrCode = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetOutId(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.OutId = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetPhoneNumber(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.PhoneNumber = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetReceiveDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ReceiveDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetReceiveType(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.ReceiveType = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetRenderDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.RenderDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetRenderStatus(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.RenderStatus = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSendDate(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SendDate = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSendStatus(v int64) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SendStatus = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetSmsContent(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.SmsContent = &v
	return s
}

func (s *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords) SetTemplateCode(v string) *GetCardSmsDetailsResponseBodyCardSendDetailDTORecords {
	s.TemplateCode = &v
	return s
}

type GetCardSmsDetailsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardSmsDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardSmsDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardSmsDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsResponse) SetHeaders(v map[string]*string) *GetCardSmsDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetCardSmsDetailsResponse) SetStatusCode(v int32) *GetCardSmsDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardSmsDetailsResponse) SetBody(v *GetCardSmsDetailsResponseBody) *GetCardSmsDetailsResponse {
	s.Body = v
	return s
}

type GetCardSmsLinkRequest struct {
	// The code type of the URLs.
	//
	// 	- **1**: group texting
	//
	// 	- **2**: personalization
	//
	// example:
	//
	// 2
	CardCodeType *int32 `json:"CardCodeType,omitempty" xml:"CardCodeType,omitempty"`
	// The type of the short URLs.
	//
	// 	- 1: standard short code.
	//
	// 	- 2: custom short code.
	//
	// > If the **CardLinkType*	- is not specified, standard short codes are generated. If you need to generate custom short codes, contact Alibaba Cloud SMS technical support.
	//
	// example:
	//
	// 1
	CardLinkType *int32 `json:"CardLinkType,omitempty" xml:"CardLinkType,omitempty"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_****
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The variables of the message template.
	//
	// example:
	//
	// [{},{}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The custom short code. It can contain 4 to 8 digits or letters.
	//
	// > If the CardLinkType parameter is set to 2, the CustomShortCodeJson parameter is required.
	//
	// example:
	//
	// abCde
	CustomShortCodeJson *string `json:"CustomShortCodeJson,omitempty" xml:"CustomShortCodeJson,omitempty"`
	// The original domain name. You must submit domain names for approval in advance.
	//
	// >
	//
	// 	- If the **CardLinkType*	- parameter is set to **2**, the **Domain*	- parameter is required.
	//
	// 	- The **Domain*	- parameter cannot exceed 100 characters in length. If the parameter is not specified, a default domain name is used.
	//
	// example:
	//
	// xxx.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The extension field.
	//
	// example:
	//
	// BC20220608102511660860762****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile phone numbers of recipients, custom identifiers, or system identifiers.
	//
	// >
	//
	// 	- A maximum of 10,000 mobile phone numbers are supported.
	//
	// 	- You can enter custom identifier. Each identifier can be a maximum of 60 characters in length.
	//
	// 	- You can apply for a maximum of 10 OPPO templates at a time.
	//
	// example:
	//
	// [\\"1390000****
	//
	// \\",\\"1370000****
	//
	// \\"]
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"aliyun\\", \\"aliyun2\\"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetCardSmsLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The mobile phone numbers that support card messages.
	//
	// example:
	//
	// [\\"1390000****\\",\\"1370000****\\"]
	CardPhoneNumbers *string `json:"CardPhoneNumbers,omitempty" xml:"CardPhoneNumbers,omitempty"`
	// The signatures must correspond to the mobile numbers and short URLs in sequence.
	//
	// example:
	//
	// ["aliyun","aliyun2"]
	CardSignNames *string `json:"CardSignNames,omitempty" xml:"CardSignNames,omitempty"`
	// The short URLs.
	//
	// example:
	//
	// [\\"mw2m.cn/LAaGGa\\",\\"mw2m.cn/LAAaes\\"]
	CardSmsLinks *string `json:"CardSmsLinks,omitempty" xml:"CardSmsLinks,omitempty"`
	// The review status of the card message template.
	//
	// 	- **0**: pending approval
	//
	// 	- **1**: approved
	//
	// 	- **2**: rejected
	//
	// > Unapproved card messages are rolled back.
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// The mobile phone numbers that do not support card messages.
	//
	// example:
	//
	// 1390000****
	NotMediaMobiles *string `json:"NotMediaMobiles,omitempty" xml:"NotMediaMobiles,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardSmsLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The extended fields.
	//
	// > If you set the ResourceType parameter to **2**, this parameter is required.
	//
	// example:
	//
	// {\\"img_rate\\":\\"oneToOne\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The size of the resource. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The description of the resource.
	//
	// example:
	//
	// remark
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The address of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://alicom-fc-media/1947741454322274/alicom-fc-media/pic/202205191526575398603697152.png
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The type of the resource.
	//
	// 	- **1**: text.
	//
	// 	- **2**: image. A small image cannot exceed 100 KB in size, and a large image cannot exceed 2 MB in size. The image must be clear. Supported format: JPG, JPEG, and PNG.
	//
	// 	- **3**: audio.
	//
	// 	- **4**: video. Supported format: MP4.
	//
	// >
	//
	// 	- If you set the ResourceType parameter to 2, the **img_rate*	- required is required. Valid values:
	//
	// 	- 1:1
	//
	// 	- 16:9
	//
	// 	- 3:1
	//
	// 	- 48:65
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetMediaResourceIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F07CF237-F6E3-5F77-B91B-F9B7C5DE84AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The download URL of the resource.
	//
	// example:
	//
	// http://test-example.com/download.jpg
	ResUrlDownload *string `json:"ResUrlDownload,omitempty" xml:"ResUrlDownload,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// SMS_14571****
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaResourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetOSSInfoForCardTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The AccessKey ID.
	//
	// example:
	//
	// LTAIxetqt1Dg****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 599333677478****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// alicom-cardsms-resources
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 1634209418
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The hostname.
	//
	// example:
	//
	// https://alicom-cardsms-resources.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The signature policy.
	//
	// example:
	//
	// eyJxxx0=
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The path of the policy.
	//
	// example:
	//
	// 1631792777
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSInfoForCardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetOSSInfoForUploadFileRequest struct {
	// Business type, default value is **fcMediaSms**.
	//
	// When creating signatures and templates, and uploading **additional materials**, this value is **fcMediaSms**.
	//
	// example:
	//
	// fcMediaSms
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOSSInfoForUploadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForUploadFileRequest) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileRequest) SetBizType(v string) *GetOSSInfoForUploadFileRequest {
	s.BizType = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetOwnerId(v int64) *GetOSSInfoForUploadFileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetResourceOwnerAccount(v string) *GetOSSInfoForUploadFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOSSInfoForUploadFileRequest) SetResourceOwnerId(v int64) *GetOSSInfoForUploadFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetOSSInfoForUploadFileResponseBody struct {
	// Request status code.
	//
	// - OK return represents a successful request.
	//
	// - For other error codes, please refer to the [Error Code List](https://help.aliyun.com/document_detail/101346.htm).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Return result.
	Model *GetOSSInfoForUploadFileResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request, can be used for troubleshooting and issue定位.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700EDCFD0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates success. Values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOSSInfoForUploadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForUploadFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileResponseBody) SetCode(v string) *GetOSSInfoForUploadFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBody) SetMessage(v string) *GetOSSInfoForUploadFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBody) SetModel(v *GetOSSInfoForUploadFileResponseBodyModel) *GetOSSInfoForUploadFileResponseBody {
	s.Model = v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBody) SetRequestId(v string) *GetOSSInfoForUploadFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBody) SetSuccess(v bool) *GetOSSInfoForUploadFileResponseBody {
	s.Success = &v
	return s
}

type GetOSSInfoForUploadFileResponseBodyModel struct {
	// AccessKey ID used for signing.
	//
	// example:
	//
	// LTAIxetqt1Dg****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Expiration time.
	//
	// example:
	//
	// 1719297445
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Host address.
	//
	// example:
	//
	// https://alicom-fc-media.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Signature policy.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyN***Ni0yNVQwNjozNzoyNS45NzBaI**iY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJzdGFydHMtd2l0***sIiRrZXkiLCIiXV19
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Signature information calculated based on **AccessKey Secret*	- and **Policy**. When calling the OSS API, OSS verifies this signature information to confirm the legitimacy of the Post request.
	//
	// example:
	//
	// BXnwCWPrhVb*****aoZHZfli5KE=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// Policy path.
	//
	// example:
	//
	// 123456
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetOSSInfoForUploadFileResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForUploadFileResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetAccessKeyId(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.AccessKeyId = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetExpireTime(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.ExpireTime = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetHost(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.Host = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetPolicy(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.Policy = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetSignature(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.Signature = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponseBodyModel) SetStartPath(v string) *GetOSSInfoForUploadFileResponseBodyModel {
	s.StartPath = &v
	return s
}

type GetOSSInfoForUploadFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSInfoForUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSInfoForUploadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOSSInfoForUploadFileResponse) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileResponse) SetHeaders(v map[string]*string) *GetOSSInfoForUploadFileResponse {
	s.Headers = v
	return s
}

func (s *GetOSSInfoForUploadFileResponse) SetStatusCode(v int32) *GetOSSInfoForUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponse) SetBody(v *GetOSSInfoForUploadFileResponseBody) *GetOSSInfoForUploadFileResponse {
	s.Body = v
	return s
}

type GetSmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature name. Must be an SMS signature already applied for by this account.
	//
	// - Obtain from the return parameters after calling the [CreateSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-createsmssign?spm) API.
	//
	// - View the signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s GetSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSmsSignRequest) GoString() string {
	return s.String()
}

func (s *GetSmsSignRequest) SetOwnerId(v int64) *GetSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsSignRequest) SetResourceOwnerAccount(v string) *GetSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsSignRequest) SetResourceOwnerId(v int64) *GetSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsSignRequest) SetSignName(v string) *GetSmsSignRequest {
	s.SignName = &v
	return s
}

type GetSmsSignResponseBody struct {
	// Content of application scenarios.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplyScene *string `json:"ApplyScene,omitempty" xml:"ApplyScene,omitempty"`
	// Audit information.
	AuditInfo *GetSmsSignResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// Request status code.
	//
	// - OK indicates a successful request.
	//
	// - For other error codes, see [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Creation date and time of the SMS signature.
	//
	// example:
	//
	// 2024-06-03 10:02:34
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// 更多资料信息，补充上传业务证明文件或业务截图文件列表。
	FileUrlList []*string `json:"FileUrlList,omitempty" xml:"FileUrlList,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Work order number.
	//
	// Used by reviewers when querying the review. You need to provide this work order number if you require expedited review.
	//
	// example:
	//
	// 20044156924
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Credential ID, the credential ID associated when applying for the signature.
	//
	// example:
	//
	// 2004393****
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// example:
	//
	// Send verification code text message during login.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request and can be used for troubleshooting and issue localization.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SMS signature code.
	//
	// example:
	//
	// SIGN_100000077042023_17174665*****_ZM2kG
	SignCode *string `json:"SignCode,omitempty" xml:"SignCode,omitempty"`
	// SMS signature name.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Signature review status. Values:
	//
	// - **0**: Under review.
	//
	// - **1**: Approved.
	//
	// - **2**: Review failed, please check the Reason parameter for the failure cause.
	//
	// - **10**: Review canceled.
	//
	// example:
	//
	// 2
	SignStatus *int64 `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
	// Signature tag indicating whether the signature is user-defined, system-provided, test, or trial. Values:
	//
	// - 2: User-defined signature
	//
	// - 3: System-provided signature
	//
	// - 4: Test signature
	//
	// - 5: Trial signature
	//
	// example:
	//
	// 2
	SignTag *string `json:"SignTag,omitempty" xml:"SignTag,omitempty"`
	// scenarios for using signatures.
	//
	// example:
	//
	// App.
	SignUsage *string `json:"SignUsage,omitempty" xml:"SignUsage,omitempty"`
	// Signature usage indication—self-use or third-party use.
	//
	// - false: Self-use (default)
	//
	// - true: Third-party use
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s GetSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBody) SetApplyScene(v string) *GetSmsSignResponseBody {
	s.ApplyScene = &v
	return s
}

func (s *GetSmsSignResponseBody) SetAuditInfo(v *GetSmsSignResponseBodyAuditInfo) *GetSmsSignResponseBody {
	s.AuditInfo = v
	return s
}

func (s *GetSmsSignResponseBody) SetCode(v string) *GetSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsSignResponseBody) SetCreateDate(v string) *GetSmsSignResponseBody {
	s.CreateDate = &v
	return s
}

func (s *GetSmsSignResponseBody) SetFileUrlList(v []*string) *GetSmsSignResponseBody {
	s.FileUrlList = v
	return s
}

func (s *GetSmsSignResponseBody) SetMessage(v string) *GetSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsSignResponseBody) SetOrderId(v string) *GetSmsSignResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetQualificationId(v int64) *GetSmsSignResponseBody {
	s.QualificationId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRemark(v string) *GetSmsSignResponseBody {
	s.Remark = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRequestId(v string) *GetSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignCode(v string) *GetSmsSignResponseBody {
	s.SignCode = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignName(v string) *GetSmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignStatus(v int64) *GetSmsSignResponseBody {
	s.SignStatus = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignTag(v string) *GetSmsSignResponseBody {
	s.SignTag = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignUsage(v string) *GetSmsSignResponseBody {
	s.SignUsage = &v
	return s
}

func (s *GetSmsSignResponseBody) SetThirdParty(v bool) *GetSmsSignResponseBody {
	s.ThirdParty = &v
	return s
}

type GetSmsSignResponseBodyAuditInfo struct {
	// Audit date and time.
	//
	// example:
	//
	// 2024-06-03 12:02:34
	AuditDate *string `json:"AuditDate,omitempty" xml:"AuditDate,omitempty"`
	// Reasons for not passing the review.
	//
	// example:
	//
	// reason for rejection.
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
}

func (s GetSmsSignResponseBodyAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSmsSignResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBodyAuditInfo) SetAuditDate(v string) *GetSmsSignResponseBodyAuditInfo {
	s.AuditDate = &v
	return s
}

func (s *GetSmsSignResponseBodyAuditInfo) SetRejectInfo(v string) *GetSmsSignResponseBodyAuditInfo {
	s.RejectInfo = &v
	return s
}

type GetSmsSignResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSmsSignResponse) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponse) SetHeaders(v map[string]*string) *GetSmsSignResponse {
	s.Headers = v
	return s
}

func (s *GetSmsSignResponse) SetStatusCode(v int32) *GetSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsSignResponse) SetBody(v *GetSmsSignResponseBody) *GetSmsSignResponse {
	s.Body = v
	return s
}

type GetSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// SMS template code.
	//
	// - Obtain the SMS template code from the return parameters of the [CreateSmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-createsmstemplate?spm) API.
	//
	// - View the SMS template code on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateRequest) SetOwnerId(v int64) *GetSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsTemplateRequest) SetResourceOwnerAccount(v string) *GetSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsTemplateRequest) SetResourceOwnerId(v int64) *GetSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsTemplateRequest) SetTemplateCode(v string) *GetSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

type GetSmsTemplateResponseBody struct {
	// Application scenario content.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplyScene *string `json:"ApplyScene,omitempty" xml:"ApplyScene,omitempty"`
	// Audit information.
	AuditInfo *GetSmsTemplateResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// Request status code.
	//
	// 	- OK indicates a successful request.
	//
	// 	- For other error codes, please refer to [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the SMS template was created.
	//
	// example:
	//
	// 2024-06-03 10:02:34
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// File information, compatible with signatures created by the [AddSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-addsmstemplate?spm) API.
	FileUrlList *GetSmsTemplateResponseBodyFileUrlList `json:"FileUrlList,omitempty" xml:"FileUrlList,omitempty" type:"Struct"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Additional documentation information, supplementing uploaded business proof files or operational screenshots file list.
	MoreDataFileUrlList *GetSmsTemplateResponseBodyMoreDataFileUrlList `json:"MoreDataFileUrlList,omitempty" xml:"MoreDataFileUrlList,omitempty" type:"Struct"`
	// Work order number.
	//
	// This parameter is used by auditors when querying the audit. You need to provide this work order number when requesting expedited review.
	//
	// example:
	//
	// 2003019****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The SMS signature associated with the template when applied.
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Explanation for the SMS template application, which is one of the reference information for template review.
	//
	// example:
	//
	// 申请验证码模板
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request and can be used for troubleshooting and issue定位.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SMS template code.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Content of the SMS template.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Name of the SMS template.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template review status. Return values:
	//
	// - **0**: Under review.
	//
	// - **1**: Approved.
	//
	// - **2**: Not approved, with reasons for failure returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.41fd339f3bPSCQ), invoke the [UpdateSmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmstemplate?spm) API or modify the SMS template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
	//
	// - **10**: Review canceled.
	//
	// example:
	//
	// 2
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// Template identifier, indicating whether the template is user-defined or system-provided. Values:
	//
	// - **2**: User-defined template.
	//
	// - **3**: System-provided template.
	//
	// example:
	//
	// 2
	TemplateTag *int32 `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-certified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. For details on the differences between personal and enterprise user rights, please refer to [Usage Notes](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// example:
	//
	// 0
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// Template variable rules.
	//
	// For detailed rules of template variables, refer to the [Example Document](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	VariableAttribute *string `json:"VariableAttribute,omitempty" xml:"VariableAttribute,omitempty"`
}

func (s GetSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBody) SetApplyScene(v string) *GetSmsTemplateResponseBody {
	s.ApplyScene = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetAuditInfo(v *GetSmsTemplateResponseBodyAuditInfo) *GetSmsTemplateResponseBody {
	s.AuditInfo = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetCode(v string) *GetSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetCreateDate(v string) *GetSmsTemplateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetFileUrlList(v *GetSmsTemplateResponseBodyFileUrlList) *GetSmsTemplateResponseBody {
	s.FileUrlList = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetIntlType(v int32) *GetSmsTemplateResponseBody {
	s.IntlType = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetMessage(v string) *GetSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetMoreDataFileUrlList(v *GetSmsTemplateResponseBodyMoreDataFileUrlList) *GetSmsTemplateResponseBody {
	s.MoreDataFileUrlList = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetOrderId(v string) *GetSmsTemplateResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRelatedSignName(v string) *GetSmsTemplateResponseBody {
	s.RelatedSignName = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRemark(v string) *GetSmsTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRequestId(v string) *GetSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateCode(v string) *GetSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateContent(v string) *GetSmsTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateName(v string) *GetSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateStatus(v string) *GetSmsTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateTag(v int32) *GetSmsTemplateResponseBody {
	s.TemplateTag = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateType(v string) *GetSmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetVariableAttribute(v string) *GetSmsTemplateResponseBody {
	s.VariableAttribute = &v
	return s
}

type GetSmsTemplateResponseBodyAuditInfo struct {
	// Audit date and time.
	//
	// example:
	//
	// 2024-06-03 11:20:34
	AuditDate *string `json:"AuditDate,omitempty" xml:"AuditDate,omitempty"`
	// Reasons for failed audit.
	//
	// example:
	//
	// 模板内容中包含错别字。
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
}

func (s GetSmsTemplateResponseBodyAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyAuditInfo) SetAuditDate(v string) *GetSmsTemplateResponseBodyAuditInfo {
	s.AuditDate = &v
	return s
}

func (s *GetSmsTemplateResponseBodyAuditInfo) SetRejectInfo(v string) *GetSmsTemplateResponseBodyAuditInfo {
	s.RejectInfo = &v
	return s
}

type GetSmsTemplateResponseBodyFileUrlList struct {
	FileUrl []*string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateResponseBodyFileUrlList) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateResponseBodyFileUrlList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyFileUrlList) SetFileUrl(v []*string) *GetSmsTemplateResponseBodyFileUrlList {
	s.FileUrl = v
	return s
}

type GetSmsTemplateResponseBodyMoreDataFileUrlList struct {
	MoreDataFileUrl []*string `json:"MoreDataFileUrl,omitempty" xml:"MoreDataFileUrl,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateResponseBodyMoreDataFileUrlList) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateResponseBodyMoreDataFileUrlList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyMoreDataFileUrlList) SetMoreDataFileUrl(v []*string) *GetSmsTemplateResponseBodyMoreDataFileUrlList {
	s.MoreDataFileUrl = v
	return s
}

type GetSmsTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponse) SetHeaders(v map[string]*string) *GetSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetSmsTemplateResponse) SetStatusCode(v int32) *GetSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsTemplateResponse) SetBody(v *GetSmsTemplateResponseBody) *GetSmsTemplateResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token used to query the next page.
	//
	// example:
	//
	// 23432453245
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the cloud service. Set the value to **dysms**.
	//
	// example:
	//
	// dysms
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region ID. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The code of the message template. Specify either the Tag or the ResourceId parameter.
	//
	// example:
	//
	// SMS_23423423
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to TEMPLATE.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list. Specify either the Tag or the ResourceId parameter. You can specify a maximum of 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The token used to query the next page.
	//
	// example:
	//
	// "23432453245"
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags.
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
	// The code of the message template.
	//
	// example:
	//
	// SMS_23423****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource.
	//
	// example:
	//
	// ALIYUN::DYSMS::TEMPLATE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The scenario description of your released services. Provide the information of your services, such as a website URL, a domain name with an ICP filing, an app download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
	//
	// > The description can be up to 200 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is the abbreviation of our company.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of signature files.
	//
	// This parameter is required.
	SignFileList []*ModifySmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The source of the signature. Valid values:
	//
	// 	- **0**: full name or abbreviation of an enterprise or institution.
	//
	// 	- **1**: full name or abbreviation of a website with Ministry of Industry and Information Technology (MIIT) filing.
	//
	// 	- **2**: full name or abbreviation of an app.
	//
	// 	- **3**: full name or abbreviation of a WeChat official account or applet.
	//
	// 	- **4**: full name or abbreviation of an e-commerce store.
	//
	// 	- **5**: full name or abbreviation of a trademark.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The type of the signature. Valid values:
	//
	// 	- **0**: verification-code signature
	//
	// 	- **1**: general-purpose signature
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
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
	// The base64-encoded string of the signed files. The size of the image cannot exceed 2 MB.
	//
	// In some scenarios, documents are required to prove your identity. For more information, see [Signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// R0lGODlhHAAmAKIHAKqqqsvLy0hISObm5vf394uLiwAA
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	// The format of the documents. You can upload multiple images. JPG, PNG, GIF, and JPEG are supported.
	//
	// In some scenarios, documents are required to prove your identity. For more information, see [Signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// > If the signature is used for other purposes or the signature source is an enterprise or public institution, you must upload some documents and an authorization letter. For more information, see [Documents](https://help.aliyun.com/document_detail/108076.html) and [Letter of authorization](https://help.aliyun.com/document_detail/56741.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// jpg
	FileSuffix *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the message template. It is one of the reference information for template review. The description cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Modify the parameters of the template.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the template. The content must be 1 to 500 characters in length.
	//
	// > When you modify a template, design the template content based on the review comments.
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name must be 1 to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: text message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_15255****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The code of the message template.
	//
	// You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_4139
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCardSmsTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The array of objects.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCardSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-10-11 00:00:01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-10-10 00:00:01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The array of message templates.
	//
	// This parameter is required.
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCardSmsTemplateReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *QueryCardSmsTemplateReportResponseBody) SetData(v *QueryCardSmsTemplateReportResponseBodyData) *QueryCardSmsTemplateReportResponseBody {
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

type QueryCardSmsTemplateReportResponseBodyData struct {
	// The details of the data returned.
	Model []map[string]interface{} `json:"model,omitempty" xml:"model,omitempty" type:"Repeated"`
}

func (s QueryCardSmsTemplateReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCardSmsTemplateReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCardSmsTemplateReportResponseBodyData) SetModel(v []map[string]interface{}) *QueryCardSmsTemplateReportResponseBodyData {
	s.Model = v
	return s
}

type QueryCardSmsTemplateReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCardSmsTemplateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryExtCodeSignRequest struct {
	// 扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExtCode *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QueryExtCodeSignRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignRequest) SetExtCode(v string) *QueryExtCodeSignRequest {
	s.ExtCode = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetOwnerId(v int64) *QueryExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetPageNo(v int64) *QueryExtCodeSignRequest {
	s.PageNo = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetPageSize(v int64) *QueryExtCodeSignRequest {
	s.PageSize = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetResourceOwnerAccount(v string) *QueryExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetResourceOwnerId(v int64) *QueryExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryExtCodeSignRequest) SetSignName(v string) *QueryExtCodeSignRequest {
	s.SignName = &v
	return s
}

type QueryExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryExtCodeSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryExtCodeSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *QueryExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetCode(v string) *QueryExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetData(v *QueryExtCodeSignResponseBodyData) *QueryExtCodeSignResponseBody {
	s.Data = v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetMessage(v string) *QueryExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetRequestId(v string) *QueryExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExtCodeSignResponseBody) SetSuccess(v bool) *QueryExtCodeSignResponseBody {
	s.Success = &v
	return s
}

type QueryExtCodeSignResponseBodyData struct {
	List []*QueryExtCodeSignResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryExtCodeSignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryExtCodeSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBodyData) SetList(v []*QueryExtCodeSignResponseBodyDataList) *QueryExtCodeSignResponseBodyData {
	s.List = v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetPageNo(v int64) *QueryExtCodeSignResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetPageSize(v int64) *QueryExtCodeSignResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyData) SetTotal(v int64) *QueryExtCodeSignResponseBodyData {
	s.Total = &v
	return s
}

type QueryExtCodeSignResponseBodyDataList struct {
	// 是否可回收
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 扩展码A3
	//
	// example:
	//
	// 01
	ExtCode *string `json:"ExtCode,omitempty" xml:"ExtCode,omitempty"`
	// 近1个月发送成功条数（只读）
	//
	// example:
	//
	// 69
	SendCount *int64 `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
	// 签名
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 来源
	//
	// example:
	//
	// 示例值示例值示例值
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s QueryExtCodeSignResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryExtCodeSignResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponseBodyDataList) SetActive(v int64) *QueryExtCodeSignResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetExtCode(v string) *QueryExtCodeSignResponseBodyDataList {
	s.ExtCode = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSendCount(v int64) *QueryExtCodeSignResponseBodyDataList {
	s.SendCount = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSignName(v string) *QueryExtCodeSignResponseBodyDataList {
	s.SignName = &v
	return s
}

func (s *QueryExtCodeSignResponseBodyDataList) SetSource(v string) *QueryExtCodeSignResponseBodyDataList {
	s.Source = &v
	return s
}

type QueryExtCodeSignResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryExtCodeSignResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *QueryExtCodeSignResponse) SetHeaders(v map[string]*string) *QueryExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *QueryExtCodeSignResponse) SetStatusCode(v int32) *QueryExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExtCodeSignResponse) SetBody(v *QueryExtCodeSignResponseBody) *QueryExtCodeSignResponse {
	s.Body = v
	return s
}

type QueryMobilesCardSupportRequest struct {
	// The list of mobile phone numbers.
	//
	// This parameter is required.
	Mobiles []map[string]interface{} `json:"Mobiles,omitempty" xml:"Mobiles,omitempty" type:"Repeated"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_0000
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryMobilesCardSupportRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportRequest) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportRequest) SetMobiles(v []map[string]interface{}) *QueryMobilesCardSupportRequest {
	s.Mobiles = v
	return s
}

func (s *QueryMobilesCardSupportRequest) SetTemplateCode(v string) *QueryMobilesCardSupportRequest {
	s.TemplateCode = &v
	return s
}

type QueryMobilesCardSupportShrinkRequest struct {
	// The list of mobile phone numbers.
	//
	// This parameter is required.
	MobilesShrink *string `json:"Mobiles,omitempty" xml:"Mobiles,omitempty"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_0000
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryMobilesCardSupportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportShrinkRequest) SetMobilesShrink(v string) *QueryMobilesCardSupportShrinkRequest {
	s.MobilesShrink = &v
	return s
}

func (s *QueryMobilesCardSupportShrinkRequest) SetTemplateCode(v string) *QueryMobilesCardSupportShrinkRequest {
	s.TemplateCode = &v
	return s
}

type QueryMobilesCardSupportResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryMobilesCardSupportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 08C17DFE-2E10-54F4-BAFB-7180039CC217
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMobilesCardSupportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBody) SetCode(v string) *QueryMobilesCardSupportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetData(v *QueryMobilesCardSupportResponseBodyData) *QueryMobilesCardSupportResponseBody {
	s.Data = v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetRequestId(v string) *QueryMobilesCardSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetSuccess(v bool) *QueryMobilesCardSupportResponseBody {
	s.Success = &v
	return s
}

type QueryMobilesCardSupportResponseBodyData struct {
	// The list of returned results.
	QueryResult []*QueryMobilesCardSupportResponseBodyDataQueryResult `json:"QueryResult,omitempty" xml:"QueryResult,omitempty" type:"Repeated"`
}

func (s QueryMobilesCardSupportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyData) SetQueryResult(v []*QueryMobilesCardSupportResponseBodyDataQueryResult) *QueryMobilesCardSupportResponseBodyData {
	s.QueryResult = v
	return s
}

type QueryMobilesCardSupportResponseBodyDataQueryResult struct {
	// The mobile phone number.
	//
	// example:
	//
	// 1380000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the mobile phone number supports card messages. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetMobile(v string) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Mobile = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetSupport(v bool) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Support = &v
	return s
}

type QueryMobilesCardSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMobilesCardSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMobilesCardSupportResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMobilesCardSupportResponse) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponse) SetHeaders(v map[string]*string) *QueryMobilesCardSupportResponse {
	s.Headers = v
	return s
}

func (s *QueryMobilesCardSupportResponse) SetStatusCode(v int32) *QueryMobilesCardSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMobilesCardSupportResponse) SetBody(v *QueryMobilesCardSupportResponseBody) *QueryMobilesCardSupportResponse {
	s.Body = v
	return s
}

type QueryPageSmartShortUrlLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20181225
	CreateDateEnd *int64 `json:"CreateDateEnd,omitempty" xml:"CreateDateEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20181225
	CreateDateStart *int64 `json:"CreateDateStart,omitempty" xml:"CreateDateStart,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// http://ays.cn/****
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryPageSmartShortUrlLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPageSmartShortUrlLogRequest) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogRequest) SetCreateDateEnd(v int64) *QueryPageSmartShortUrlLogRequest {
	s.CreateDateEnd = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetCreateDateStart(v int64) *QueryPageSmartShortUrlLogRequest {
	s.CreateDateStart = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetOwnerId(v int64) *QueryPageSmartShortUrlLogRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPageNo(v int64) *QueryPageSmartShortUrlLogRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPageSize(v int64) *QueryPageSmartShortUrlLogRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPhoneNumber(v string) *QueryPageSmartShortUrlLogRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetResourceOwnerAccount(v string) *QueryPageSmartShortUrlLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetResourceOwnerId(v int64) *QueryPageSmartShortUrlLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetShortUrl(v string) *QueryPageSmartShortUrlLogRequest {
	s.ShortUrl = &v
	return s
}

type QueryPageSmartShortUrlLogResponseBody struct {
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QueryPageSmartShortUrlLogResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetCode(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetMessage(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetModel(v *QueryPageSmartShortUrlLogResponseBodyModel) *QueryPageSmartShortUrlLogResponseBody {
	s.Model = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetRequestId(v string) *QueryPageSmartShortUrlLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBody) SetSuccess(v bool) *QueryPageSmartShortUrlLogResponseBody {
	s.Success = &v
	return s
}

type QueryPageSmartShortUrlLogResponseBodyModel struct {
	List []*QueryPageSmartShortUrlLogResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 74
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 15
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 66
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 86
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetList(v []*QueryPageSmartShortUrlLogResponseBodyModelList) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.List = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetPageNo(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetPageSize(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetTotalCount(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModel) SetTotalPage(v int64) *QueryPageSmartShortUrlLogResponseBodyModel {
	s.TotalPage = &v
	return s
}

type QueryPageSmartShortUrlLogResponseBodyModelList struct {
	// example:
	//
	// 87
	ClickState *int64 `json:"ClickState,omitempty" xml:"ClickState,omitempty"`
	// example:
	//
	// 51
	ClickTime *int64 `json:"ClickTime,omitempty" xml:"ClickTime,omitempty"`
	// example:
	//
	// 64
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 示例值示例值
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetClickState(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ClickState = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetClickTime(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ClickTime = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetCreateTime(v int64) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetPhoneNumber(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.PhoneNumber = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetShortName(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ShortName = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponseBodyModelList) SetShortUrl(v string) *QueryPageSmartShortUrlLogResponseBodyModelList {
	s.ShortUrl = &v
	return s
}

type QueryPageSmartShortUrlLogResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPageSmartShortUrlLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPageSmartShortUrlLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPageSmartShortUrlLogResponse) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogResponse) SetHeaders(v map[string]*string) *QueryPageSmartShortUrlLogResponse {
	s.Headers = v
	return s
}

func (s *QueryPageSmartShortUrlLogResponse) SetStatusCode(v int32) *QueryPageSmartShortUrlLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPageSmartShortUrlLogResponse) SetBody(v *QueryPageSmartShortUrlLogResponseBody) *QueryPageSmartShortUrlLogResponse {
	s.Body = v
	return s
}

type QuerySendDetailsRequest struct {
	// The ID of the delivery receipt. The delivery receipt ID is the value of the BizId parameter that is returned when you call the SendSms or SendBatchSms operation.
	//
	// example:
	//
	// 134523^435****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The page number of the first page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of items displayed per page.
	//
	// Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mobile numbers of the recipients. Format:
	//
	// 	- If you send messages in the Chinese mainland, specify an 11-digit mobile number, for example, 1390000\\*\\*\\*\\*.
	//
	// 	- If you send messages to countries or regions outside the Chinese mainland, specify this parameter in the \\<Area code>\\<Mobile number> format. Example: 8520000\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The date when the message was sent. You can query messages that were sent within the last 30 days.
	//
	// Format: yyyyMMdd. Example: 20181225.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20181228
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
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
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the message.
	SmsSendDetailDTOs *QuerySendDetailsResponseBodySmsSendDetailDTOs `json:"SmsSendDetailDTOs,omitempty" xml:"SmsSendDetailDTOs,omitempty" type:"Struct"`
	// The number of sent messages.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The content of the message.
	//
	// example:
	//
	// 【Aliyun】This is a test message.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status code returned by the carrier.
	//
	// 	- If the message is delivered, "DELIVERED" is returned.
	//
	// 	- For information about the error codes that may be returned if the message is not delivered, see [error codes](https://help.aliyun.com/document_detail/101347.html).
	//
	// example:
	//
	// DELIVERED
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The extended field.
	//
	// example:
	//
	// 123
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile numbers of the recipients.
	//
	// example:
	//
	// 1390000****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The date and time when the message was received.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// The date and time when the message was sent.
	//
	// example:
	//
	// 2019-01-08 16:44:10
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The delivery status of the message. Valid values:
	//
	// 	- **1**: The message has not received a delivery receipt yet.
	//
	// 	- **2**: The message failed to be delivered.
	//
	// 	- **3**: The message was delivered.
	//
	// example:
	//
	// 3
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The ID of the message template.
	//
	// example:
	//
	// SMS_12231****
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The end of the time range to query. Format: yyyyMMdd. Example: 20181225.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201003
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The site from where the message is sent. Valid values:
	//
	// 	- **1**: China site
	//
	// 	- **2**: international site
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsGlobe *int32 `json:"IsGlobe,omitempty" xml:"IsGlobe,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries to return on each page. Valid values: **1 to 50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The beginning of the time range to query. Format: yyyyMMdd. Example: 20181225.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201002
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The type of the message template. Valid values: Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification
	//
	// 	- **2**: promotional message (Enterprise users only)
	//
	// 	- **3**: international purpose (Enterprise users only)
	//
	// 	- **7**: digital message
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QuerySendStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The details of the data returned.
	TargetList []*QuerySendStatisticsResponseBodyDataTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
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
	// The number of messages without a delivery receipt.
	//
	// example:
	//
	// 1
	NoRespondedCount *int64 `json:"NoRespondedCount,omitempty" xml:"NoRespondedCount,omitempty"`
	// The number of messages with a delivery receipt that indicates a failure.
	//
	// example:
	//
	// 2
	RespondedFailCount *int64 `json:"RespondedFailCount,omitempty" xml:"RespondedFailCount,omitempty"`
	// The number of messages with a delivery receipt that indicates a success.
	//
	// example:
	//
	// 17
	RespondedSuccessCount *int64 `json:"RespondedSuccessCount,omitempty" xml:"RespondedSuccessCount,omitempty"`
	// The date when the message is sent. Format: yyyyMMdd. Example: 20181225.
	//
	// example:
	//
	// 20201010
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The number of delivered messages.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The short URL. You can query the short URL by calling the [AddShortUrl](https://help.aliyun.com/document_detail/186774.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the short URL.
	Data *QueryShortUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The time when the short URL was created.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the short URL expires.
	//
	// example:
	//
	// 2019-01-22 11:21:11
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The PV.
	//
	// example:
	//
	// 300
	PageViewCount *string `json:"PageViewCount,omitempty" xml:"PageViewCount,omitempty"`
	// The short URL.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	// The service name of the short URL.
	//
	// example:
	//
	// The Alibaba Cloud Short Link service.
	ShortUrlName *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	// The status of the short URL. Valid values:
	//
	// 	- **expired**
	//
	// 	- **effective**
	//
	// 	- **audit**
	//
	// 	- **reject**
	//
	// example:
	//
	// expired
	ShortUrlStatus *string `json:"ShortUrlStatus,omitempty" xml:"ShortUrlStatus,omitempty"`
	// The source address.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	// The UV.
	//
	// example:
	//
	// 23
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The date and time when the signature was created.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The remarks of the review. Valid values:
	//
	// 	- If the signature is in the **Approved*	- or **Pending Approval*	- state, No Remarks is returned.
	//
	// 	- If the signature is in the **Not Approved*	- state, the reason why the signature is rejected is returned.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC89A90C-978F-46AC-B80D-54738371E7CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The status of the signature. Valid values:
	//
	// 	- **0**: The signature is pending approval.
	//
	// 	- **1**: The signature is approved.
	//
	// 	- **2**: The signature is rejected. The Reason parameter indicates the reason why the signature is rejected.
	//
	// 	- **10**: The signature is cancelled.
	//
	// example:
	//
	// 1
	SignStatus *int32 `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of signatures per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of signatures per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried message signatures.
	SmsSignList []*QuerySmsSignListResponseBodySmsSignList `json:"SmsSignList,omitempty" xml:"SmsSignList,omitempty" type:"Repeated"`
	// The total number of signatures.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The approval status of the signature. Valid values:
	//
	// 	- **AUDIT_STATE_INIT**: The signature is pending approval.
	//
	// 	- **AUDIT_STATE_PASS**: The signature is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The signature is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **AUDIT_STATE_CANCEL**: The approval is canceled.
	//
	// example:
	//
	// AUDIT_STATE_NOT_PASS
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The type of the signature scenario. The return value ends with "type". Valid values:
	//
	// 	- Verification code type
	//
	// 	- General-purpose type
	//
	// example:
	//
	// Verification code type
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The time when the signature was created. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-08 16:44:13
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 236****5
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the signature is rejected is returned.
	Reason *QuerySmsSignListResponseBodySmsSignListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The name of the signature.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	// The time when the signature was rejected. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-08 19:02:13
	RejectDate *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	// The reason why the signature was rejected.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The remarks about the rejection.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsSignListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_1525***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the message template was created.
	//
	// example:
	//
	// 2019-06-04 11:42:17
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the message template is rejected is returned.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A974B78-02BF-4C79-ADF3-90CFBA1B55B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_16703****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the message template.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The approval status of the message template. Valid values:
	//
	// 	- **0**: The message template is pending approval.
	//
	// 	- **1**: The message template is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The message template is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **10**: The approval is canceled.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of templates per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of templates per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E47708****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried message templates.
	SmsTemplateList []*QuerySmsTemplateListResponseBodySmsTemplateList `json:"SmsTemplateList,omitempty" xml:"SmsTemplateList,omitempty" type:"Repeated"`
	// The total number of templates.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The approval status of the message template. Valid values:
	//
	// 	- **AUDIT_STATE_INIT**: The message template is pending approval.
	//
	// 	- **AUDIT_STATE_PASS**: The message template is approved.
	//
	// 	- **AUDIT_STATE_NOT_PASS**: The message template is rejected. You can view the reason in the Reason response parameter.
	//
	// 	- **AUDIT_STATE_CANCEL*	- or **AUDIT_SATE_CANCEL**: The approval is canceled.
	//
	// example:
	//
	// AUDIT_STATE_PASS
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The time when the message template was created. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2020-06-04 11:42:17
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 2361****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The type of the message template. We recommend that you specify this parameter. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// 	- **7**: digital message
	//
	// > The template type is the same as the value of the TemplateType parameter in the AddSmsTemplate and ModifySmsTemplate operations.
	//
	// example:
	//
	// 0
	OuterTemplateType *int32 `json:"OuterTemplateType,omitempty" xml:"OuterTemplateType,omitempty"`
	// The approval remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_PASS*	- or **AUDIT_STATE_INIT**, the value of Reason is No Approval Remarks.
	//
	// 	- If the value of AuditStatus is **AUDIT_STATE_NOT_PASS**, the reason why the message template is rejected is returned.
	Reason *QuerySmsTemplateListResponseBodySmsTemplateListReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// example:
	//
	// SMS_1525***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the message template.
	//
	// example:
	//
	// 123456789
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the message template.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message template. Valid values:
	//
	// 	- **0**: notification message
	//
	// 	- **1**: promotional message
	//
	// 	- **2**: verification code
	//
	// 	- **6**: message sent to countries or regions outside the Chinese mainland
	//
	// 	- **7**: digital message
	//
	// example:
	//
	// 7
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
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
	// The time when the message template was rejected. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-06-04 16:01:17
	RejectDate *string `json:"RejectDate,omitempty" xml:"RejectDate,omitempty"`
	// The reason why the message template was rejected.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The remarks about the rejection.
	//
	// example:
	//
	// The document cannot verify the authenticity of the information. Please upload it again.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_3245
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The variables of the card message template.
	//
	// example:
	//
	// [{\\"customurl\\":\\"http://www.alibaba.com\\",\\"dyncParams\\":\\"{\\\\\\"a\\\\\\":\\\\\\"hello\\\\\\",\\\\\\"b\\\\\\":\\\\\\"world\\\\\\"}\\"}]
	CardTemplateParamJson *string `json:"CardTemplateParamJson,omitempty" xml:"CardTemplateParamJson,omitempty"`
	// The code of the digital message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// DIGITAL_SMS_234080176
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The variables of the digital message template.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	DigitalTemplateParamJson *string `json:"DigitalTemplateParamJson,omitempty" xml:"DigitalTemplateParamJson,omitempty"`
	// The rollback type. Valid values:
	//
	// 	- **SMS**: text message
	//
	// 	- **DIGITALSMS**: digital message
	//
	// 	- **NONE**: none
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	FallbackType *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	// The ID that is reserved for the caller of the operation.
	//
	// example:
	//
	// 16545681783595370
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The mobile numbers of the recipients.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"1390000****\\",\\"1370000****\\"]"
	PhoneNumberJson *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"aliyun\\",\\"aliyuncode\\"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// The code of the text message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// SMS_234251075
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The variables of the text message template.
	//
	// example:
	//
	// [{"a":1,"b":2},{"a":9,"b":8}]
	SmsTemplateParamJson *string `json:"SmsTemplateParamJson,omitempty" xml:"SmsTemplateParamJson,omitempty"`
	// The extension code of the upstream message.
	//
	// example:
	//
	// [\\"6\\",\\"6\\"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Alibaba Cloud console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the **template code*	- on the **Templates*	- tab.
	//
	// > You must specify a message template that is created in the SMS console and approved by Alibaba Cloud. If you send messages to countries or regions outside the Chinese mainland, use the corresponding message templates.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid. In addition, the sequence of variable values must be the same as that of the mobile numbers and signatures.
	//
	// example:
	//
	// [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
	TemplateParamJson *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
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

func (s *SendBatchCardSmsRequest) SetTemplateCode(v string) *SendBatchCardSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendBatchCardSmsRequest) SetTemplateParamJson(v string) *SendBatchCardSmsRequest {
	s.TemplateParamJson = &v
	return s
}

type SendBatchCardSmsResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SendBatchCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The ID of the card message.
	//
	// example:
	//
	// 123
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// The ID of the digital message.
	//
	// example:
	//
	// 3214
	BizDigitalId *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	// The ID of the text message.
	//
	// example:
	//
	// 3256
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// The review status of the card message template.
	//
	// 	- **0**: pending approval
	//
	// 	- **1**: approved
	//
	// 	- **2**: rejected
	//
	// > Unapproved card messages are rolled back.
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// The mobile phone number from which the card message is sent.
	//
	// example:
	//
	// 1390000****
	MediaMobiles *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	// The mobile phone number whose card message is rolled back.
	//
	// example:
	//
	// 1390000****
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBatchCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The extension field of the external record. The value is a string that contains no more than 256 characters.
	//
	// > The parameter is optional.
	//
	// example:
	//
	// abcdefg
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mobile number of the recipient. Format:
	//
	// 	- Message delivery to the Chinese mainland: +/+86/0086/86 or an 11-digit mobile number without a prefix. Example: 1590000\\*\\*\\*\\*.
	//
	// 	- Message delivery to countries or regions outside the Chinese mainland: Dialing code + Mobile number. Example: 852000012\\*\\*\\*\\*.
	//
	// > We recommend that you call the SendSms operation to send verification codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1590000****","1350000****"]
	PhoneNumberJson      *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the signature in the **Signature*	- column on the **Signatures*	- tab.
	//
	// > The signatures must be approved and correspond to the mobile numbers in sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Aliyun","Alibaba"]
	SignNameJson *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	// The extension code of the MO message. Format: JSON array.
	//
	// > The parameter is optional.
	//
	// example:
	//
	// ["90999","90998"]
	SmsUpExtendCodeJson *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
	// The code of the message template.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the message template in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// > The message templates must be created on the Go Globe page and approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid. In addition, the sequence of variable values must be the same as that of the mobile numbers and signatures.
	//
	// example:
	//
	// [{"name":"TemplateParamJson"},{"name":"TemplateParamJson"}]
	TemplateParamJson *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
}

func (s SendBatchSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBatchSmsRequest) GoString() string {
	return s.String()
}

func (s *SendBatchSmsRequest) SetOutId(v string) *SendBatchSmsRequest {
	s.OutId = &v
	return s
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
	// The ID of the delivery receipt. You can use one of the following methods to query the delivery status of a message based on the ID.
	//
	// 	- Call the [QuerySendDetails](https://help.aliyun.com/document_detail/102352.html) operation.
	//
	// 	- Log on to the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview). In the left-side navigation pane, choose **Analytics*	- > **Delivery Report**.
	//
	// example:
	//
	// 9006197469364984400
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D230E
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBatchSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The objects of the message template.
	//
	// This parameter is required.
	CardObjects []*SendCardSmsRequestCardObjects `json:"CardObjects,omitempty" xml:"CardObjects,omitempty" type:"Repeated"`
	// The code of the message template. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// CARD_SMS_70
	CardTemplateCode *string `json:"CardTemplateCode,omitempty" xml:"CardTemplateCode,omitempty"`
	// The code of the digital message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved.
	//
	// example:
	//
	// SMS_003
	DigitalTemplateCode *string `json:"DigitalTemplateCode,omitempty" xml:"DigitalTemplateCode,omitempty"`
	// The variables of the digital message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {\\"msg\\",\\"xxxd\\"}
	DigitalTemplateParam *string `json:"DigitalTemplateParam,omitempty" xml:"DigitalTemplateParam,omitempty"`
	// The rollback type. Valid values:
	//
	// 	- **SMS**: text message
	//
	// 	- **DIGITALSMS**: digital message
	//
	// 	- **NONE**: none
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	FallbackType *string `json:"FallbackType,omitempty" xml:"FallbackType,omitempty"`
	// The ID that is reserved for the caller of the operation.
	//
	// example:
	//
	// 38d76c9b-4a9a-4c89-afae-61fd8e0e****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The signature. You can view the template code in the **Signature*	- column on the **Signaturess*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > The signature must be approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The code of the text message template that applies when the card message is rolled back. You can view the template code in the **Template Code*	- column on the **Templates*	- tab of the **Go China*	- page in the Alibaba Cloud SMS console.
	//
	// > Make sure that the message template has been approved. If you set the **FallbackType*	- parameter to **SMS**, this parameter is required.
	//
	// example:
	//
	// SIER_TEST_01
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	// The variables of the text message template.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {\\"uri\\":\\"Zg11tZ\\"}
	SmsTemplateParam *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	// The extension code of the upstream message. Upstream messages are messages sent to the communication service provider. Upstream messages are used to customize a service, complete an inquiry, or send a request. You are charged for sending upstream messages based on the billing standards of the service provider.
	//
	// > If you do not need upstream messages, ignore this parameter.
	//
	// example:
	//
	// 1
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the text message template.
	//
	// Log on to the Alibaba Cloud SMS console. In the left-side navigation pane, click **Go Globe*	- or **Go China**. You can view the message template in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// > The message templates must be created on the Go Globe page and approved.
	//
	// example:
	//
	// SMS_2322****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The variables of the message template. Format: JSON.
	//
	// > If you need to add line breaks to the JSON template, make sure that the format is valid.
	//
	// example:
	//
	// {
	//
	//       \\"code\\": \\"1111\\"
	//
	// }
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
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

func (s *SendCardSmsRequest) SetTemplateCode(v string) *SendCardSmsRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendCardSmsRequest) SetTemplateParam(v string) *SendCardSmsRequest {
	s.TemplateParam = &v
	return s
}

type SendCardSmsRequestCardObjects struct {
	// The URL to which the message is redirected if the message fails to be rendered.
	//
	// example:
	//
	// https://alibaba.com
	CustomUrl *string `json:"customUrl,omitempty" xml:"customUrl,omitempty"`
	// The variables. Special characters, such as $ and {}, do not need to be entered.
	//
	// example:
	//
	// {\\"param3\\":\\"three\\",\\"param1\\":\\"one\\",\\"param2\\":\\"two\\"}
	DyncParams *string `json:"dyncParams,omitempty" xml:"dyncParams,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SendCardSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D28D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The ID of the card message.
	//
	// example:
	//
	// 123
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// The ID of the digital message.
	//
	// example:
	//
	// 232
	BizDigitalId *string `json:"BizDigitalId,omitempty" xml:"BizDigitalId,omitempty"`
	// The ID of the text message.
	//
	// example:
	//
	// 524
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// The review status of the card message template.
	//
	// 	- **0**: pending approval
	//
	// 	- **1**: approved
	//
	// 	- **2**: rejected
	//
	// > Unapproved card messages are rolled back.
	//
	// example:
	//
	// 0
	CardTmpState *int32 `json:"CardTmpState,omitempty" xml:"CardTmpState,omitempty"`
	// The mobile phone number from which the card message is sent.
	//
	// example:
	//
	// 1390000****
	MediaMobiles *string `json:"MediaMobiles,omitempty" xml:"MediaMobiles,omitempty"`
	// The mobile phone number whose card message is rolled back.
	//
	// example:
	//
	// 1390000****
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The extension field.
	//
	// > You can ignore this parameter if you do not have special requirements.
	//
	// example:
	//
	// abcdefgh
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mobile numbers of the recipients. Format:
	//
	// 	- If you send messages to the Chinese mainland, specify mobile numbers that are prefixed with +, +86, 0086, or 86, or 11-digit mobile numbers without prefixes. Example: 1390000\\*\\*\\*\\*.
	//
	// 	- If you send messages to countries or regions outside the Chinese mainland, specify this parameter in the \\<Area code>\\<Mobile number> format. Example: 852000012\\*\\*\\*\\*.
	//
	// You can send messages to multiple mobile numbers, separate the mobile numbers with commas (,). You can specify up to 1,000 mobile numbers in each request. Compared with sending messages to a single mobile number, sending messages to multiple mobile numbers requires longer response time.
	//
	// > We recommend that you send one verification code message to a mobile number in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the signature on the **Signatures*	- tab.
	//
	// > You must specify a signature that is created in the SMS console and approved by Alibaba Cloud. For more information about SMS signature specifications, see [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The extension code of the upstream message. Upstream messages are messages sent to the communication service provider. Upstream messages are used to customize a service, complete an inquiry, or send a request. You are charged for sending upstream messages based on the billing standards of the service provider.
	//
	// > The extension code is automatically generated by the system when the signature is generated. You do not need to specify the extension code. You can ignore this parameter if you do not have special requirements.
	//
	// example:
	//
	// 90999
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the **template code*	- on the **Templates*	- tab.
	//
	// > You must specify a message template that is created in the SMS console and approved by Alibaba Cloud. If you send messages to countries or regions outside the Chinese mainland, use the corresponding message templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15305****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template. You can specify multiple parameter values. Example:{"name":"Sam","number":"1390000\\*\\*\\*\\*"}.
	//
	// >
	//
	// 	- If line breaks are required in JSON-formatted data, they must meet the relevant requirements that are specified in the standard JSON protocol.
	//
	// 	- For more information about template variables, see [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
	//
	// example:
	//
	// {"code":"1111"}
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
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
	// The ID of the delivery receipt.
	//
	// You can call the [QuerySendDetails](~~QuerySendDetails~~) operation to query the delivery status based on the receipt ID.
	//
	// example:
	//
	// 9006197469364984****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SmsConversionIntlRequest struct {
	// The time when the OTP message was delivered. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	- If you leave the parameter empty, the current timestamp is specified by default.
	//
	// 	- If you specify the parameter, the timestamp must be greater than the message sending time and less than the current timestamp.
	//
	// example:
	//
	// 1349055900000
	ConversionTime *int64 `json:"ConversionTime,omitempty" xml:"ConversionTime,omitempty"`
	// Specifies whether customers replied to the OTP message. Valid values: true and false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// The ID of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1008030300****
	MessageId            *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SmsConversionIntlRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionIntlRequest) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlRequest) SetConversionTime(v int64) *SmsConversionIntlRequest {
	s.ConversionTime = &v
	return s
}

func (s *SmsConversionIntlRequest) SetDelivered(v bool) *SmsConversionIntlRequest {
	s.Delivered = &v
	return s
}

func (s *SmsConversionIntlRequest) SetMessageId(v string) *SmsConversionIntlRequest {
	s.MessageId = &v
	return s
}

func (s *SmsConversionIntlRequest) SetOwnerId(v int64) *SmsConversionIntlRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsConversionIntlRequest) SetResourceOwnerAccount(v string) *SmsConversionIntlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsConversionIntlRequest) SetResourceOwnerId(v int64) *SmsConversionIntlRequest {
	s.ResourceOwnerId = &v
	return s
}

type SmsConversionIntlResponseBody struct {
	// The response code. If OK is returned, the request is successful. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.101345.0.0.74326ff2J5EZyt).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SmsConversionIntlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionIntlResponseBody) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlResponseBody) SetCode(v string) *SmsConversionIntlResponseBody {
	s.Code = &v
	return s
}

func (s *SmsConversionIntlResponseBody) SetMessage(v string) *SmsConversionIntlResponseBody {
	s.Message = &v
	return s
}

func (s *SmsConversionIntlResponseBody) SetRequestId(v string) *SmsConversionIntlResponseBody {
	s.RequestId = &v
	return s
}

type SmsConversionIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsConversionIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsConversionIntlResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionIntlResponse) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlResponse) SetHeaders(v map[string]*string) *SmsConversionIntlResponse {
	s.Headers = v
	return s
}

func (s *SmsConversionIntlResponse) SetStatusCode(v int32) *SmsConversionIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsConversionIntlResponse) SetBody(v *SmsConversionIntlResponseBody) *SmsConversionIntlResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the cloud service. Set the value to **dysms**.
	//
	// example:
	//
	// dysms
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region ID. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// SMS_23423423
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to **TEMPLATE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The array of tag keys. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The array of tag values. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- Other values indicate that the request fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether tags were attached. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C****
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to delete all tags from the message template. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	All     *bool  `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the cloud service. Set the value to **dysms**.
	//
	// example:
	//
	// dysms
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The region. Set the value to cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The array of message template codes. You can specify 1 to 20 message templates.
	//
	// example:
	//
	// SMS_23423423
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to TEMPLATE.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The array of tag keys. You can specify 1 to 20 tag keys.
	//
	// example:
	//
	// TestKey
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C****
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateExtCodeSignRequest struct {
	// 要修改的扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ExistExtCode *string `json:"ExistExtCode,omitempty" xml:"ExistExtCode,omitempty"`
	// 修改后的扩展码A3
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	NewExtCode           *string `json:"NewExtCode,omitempty" xml:"NewExtCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpdateExtCodeSignRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtCodeSignRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignRequest) SetExistExtCode(v string) *UpdateExtCodeSignRequest {
	s.ExistExtCode = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetNewExtCode(v string) *UpdateExtCodeSignRequest {
	s.NewExtCode = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetOwnerId(v int64) *UpdateExtCodeSignRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetResourceOwnerAccount(v string) *UpdateExtCodeSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetResourceOwnerId(v int64) *UpdateExtCodeSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateExtCodeSignRequest) SetSignName(v string) *UpdateExtCodeSignRequest {
	s.SignName = &v
	return s
}

type UpdateExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExtCodeSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *UpdateExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetCode(v string) *UpdateExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetData(v bool) *UpdateExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetMessage(v string) *UpdateExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetRequestId(v string) *UpdateExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetSuccess(v bool) *UpdateExtCodeSignResponseBody {
	s.Success = &v
	return s
}

type UpdateExtCodeSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtCodeSignResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignResponse) SetHeaders(v map[string]*string) *UpdateExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtCodeSignResponse) SetStatusCode(v int32) *UpdateExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExtCodeSignResponse) SetBody(v *UpdateExtCodeSignResponseBody) *UpdateExtCodeSignResponse {
	s.Body = v
	return s
}

type UpdateSmsSignRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the domain name registered with MIIT, including HTTP or HTTPS.
	//
	// - For launched apps, provide the display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, fill in the full name, ensuring they are online.
	//
	// - For e-commerce platform store names (for enterprise users only), provide the display link with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// Additional materials, such as uploading business proof documents or screenshots of business operations, to help reviewers understand your business details.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [apply for qualifications](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// > The scenario explanation is one of the reference information for signature review. Please provide a detailed description of the usage scenarios of the launched business, along with verifiable information such as website links, registered domain addresses, app store download links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A well-informed application explanation will enhance the efficiency of signature and template reviews. Refer to the **Application Scenarios*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// 登录场景申请验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature not yet approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云验证码
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Source of the signature. Values:
	//
	// - **0**: Full name or abbreviation of enterprises and institutions.
	//
	// - **1**: Full name or abbreviation of MIIT-registered websites.
	//
	// - **2**: Full name or abbreviation of app applications.
	//
	// - **3**: Full name or abbreviation of public accounts or mini-programs.
	//
	// - **4**: Full name or abbreviation of e-commerce platform store names.
	//
	// - **5**: Full name or abbreviation of trademarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. It is recommended to use the default value.
	//
	// - **0**: Verification code
	//
	// - **1**: General (default)
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Whether the signature is for self-use or others.
	//
	// - false: Self-use
	//
	// - true: Others
	//
	// 	Notice: When the signature is for self-use, select the self-use qualification ID; when it\\"s for others, choose the others\\" qualification ID.
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s UpdateSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignRequest) SetApplySceneContent(v string) *UpdateSmsSignRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsSignRequest) SetMoreData(v []*string) *UpdateSmsSignRequest {
	s.MoreData = v
	return s
}

func (s *UpdateSmsSignRequest) SetOwnerId(v int64) *UpdateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetQualificationId(v int64) *UpdateSmsSignRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetRemark(v string) *UpdateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsSignRequest) SetResourceOwnerAccount(v string) *UpdateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsSignRequest) SetResourceOwnerId(v int64) *UpdateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignName(v string) *UpdateSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignSource(v int32) *UpdateSmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignType(v int32) *UpdateSmsSignRequest {
	s.SignType = &v
	return s
}

func (s *UpdateSmsSignRequest) SetThirdParty(v bool) *UpdateSmsSignRequest {
	s.ThirdParty = &v
	return s
}

type UpdateSmsSignShrinkRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the domain name registered with MIIT, including HTTP or HTTPS.
	//
	// - For launched apps, provide the display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, fill in the full name, ensuring they are online.
	//
	// - For e-commerce platform store names (for enterprise users only), provide the display link with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// Additional materials, such as uploading business proof documents or screenshots of business operations, to help reviewers understand your business details.
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [apply for qualifications](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// > The scenario explanation is one of the reference information for signature review. Please provide a detailed description of the usage scenarios of the launched business, along with verifiable information such as website links, registered domain addresses, app store download links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A well-informed application explanation will enhance the efficiency of signature and template reviews. Refer to the **Application Scenarios*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// 登录场景申请验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature not yet approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云验证码
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Source of the signature. Values:
	//
	// - **0**: Full name or abbreviation of enterprises and institutions.
	//
	// - **1**: Full name or abbreviation of MIIT-registered websites.
	//
	// - **2**: Full name or abbreviation of app applications.
	//
	// - **3**: Full name or abbreviation of public accounts or mini-programs.
	//
	// - **4**: Full name or abbreviation of e-commerce platform store names.
	//
	// - **5**: Full name or abbreviation of trademarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. It is recommended to use the default value.
	//
	// - **0**: Verification code
	//
	// - **1**: General (default)
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Whether the signature is for self-use or others.
	//
	// - false: Self-use
	//
	// - true: Others
	//
	// 	Notice: When the signature is for self-use, select the self-use qualification ID; when it\\"s for others, choose the others\\" qualification ID.
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s UpdateSmsSignShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignShrinkRequest) SetApplySceneContent(v string) *UpdateSmsSignShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetMoreDataShrink(v string) *UpdateSmsSignShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetOwnerId(v int64) *UpdateSmsSignShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetQualificationId(v int64) *UpdateSmsSignShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetRemark(v string) *UpdateSmsSignShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetResourceOwnerAccount(v string) *UpdateSmsSignShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetResourceOwnerId(v int64) *UpdateSmsSignShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignName(v string) *UpdateSmsSignShrinkRequest {
	s.SignName = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignSource(v int32) *UpdateSmsSignShrinkRequest {
	s.SignSource = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignType(v int32) *UpdateSmsSignShrinkRequest {
	s.SignType = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetThirdParty(v bool) *UpdateSmsSignShrinkRequest {
	s.ThirdParty = &v
	return s
}

type UpdateSmsSignResponseBody struct {
	// Request status code.
	//
	// 	- OK indicates a successful request.
	//
	// 	- For other error codes, refer to [Error Code List](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Work order number.
	//
	// This parameter is used by auditors when querying audits. You need to provide this work order number for expedited review.
	//
	// example:
	//
	// 2004417****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of this call request, uniquely generated by Alibaba Cloud, which can be used for troubleshooting and issue localization.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The modified signature name.
	//
	// example:
	//
	// 登录验证
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpdateSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignResponseBody) SetCode(v string) *UpdateSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmsSignResponseBody) SetMessage(v string) *UpdateSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmsSignResponseBody) SetOrderId(v string) *UpdateSmsSignResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateSmsSignResponseBody) SetRequestId(v string) *UpdateSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmsSignResponseBody) SetSignName(v string) *UpdateSmsSignResponseBody {
	s.SignName = &v
	return s
}

type UpdateSmsSignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignResponse) SetHeaders(v map[string]*string) *UpdateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsSignResponse) SetStatusCode(v int32) *UpdateSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsSignResponse) SetBody(v *UpdateSmsSignResponseBody) *UpdateSmsSignResponse {
	s.Body = v
	return s
}

type UpdateSmsTemplateRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the MIIT-registered domain with HTTP or HTTPS.
	//
	// - For launched apps, provide the app store display link with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, enter the full name of the public account or mini-program, ensuring they are online.
	//
	// - For e-commerce platform stores, applicable only to enterprise users, enter the display link of the e-commerce store with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional information, such as uploading business proof documents or screenshots, to help reviewers understand your business details. Optional and can be left unset.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// SMS signature associated with the template during the application.
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Explanation for the SMS template application, which serves as a reference for template review.
	//
	// example:
	//
	// 登录场景使用验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template Code of an unapproved template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS regulations; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. Variable specifications can be found in [TemplateContent Parameter Variable Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For guidance on filling variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-certified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. Details on differences between personal and enterprise user rights are available in [Usage Guidelines](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s UpdateSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateRequest) SetApplySceneContent(v string) *UpdateSmsTemplateRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetIntlType(v int32) *UpdateSmsTemplateRequest {
	s.IntlType = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetMoreData(v []*string) *UpdateSmsTemplateRequest {
	s.MoreData = v
	return s
}

func (s *UpdateSmsTemplateRequest) SetOwnerId(v int64) *UpdateSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetRelatedSignName(v string) *UpdateSmsTemplateRequest {
	s.RelatedSignName = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetRemark(v string) *UpdateSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetResourceOwnerAccount(v string) *UpdateSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetResourceOwnerId(v int64) *UpdateSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateCode(v string) *UpdateSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateContent(v string) *UpdateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateName(v string) *UpdateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateRule(v string) *UpdateSmsTemplateRequest {
	s.TemplateRule = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateType(v int32) *UpdateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

type UpdateSmsTemplateShrinkRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the MIIT-registered domain with HTTP or HTTPS.
	//
	// - For launched apps, provide the app store display link with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, enter the full name of the public account or mini-program, ensuring they are online.
	//
	// - For e-commerce platform stores, applicable only to enterprise users, enter the display link of the e-commerce store with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional information, such as uploading business proof documents or screenshots, to help reviewers understand your business details. Optional and can be left unset.
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// SMS signature associated with the template during the application.
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Explanation for the SMS template application, which serves as a reference for template review.
	//
	// example:
	//
	// 登录场景使用验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template Code of an unapproved template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS regulations; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. Variable specifications can be found in [TemplateContent Parameter Variable Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For guidance on filling variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-certified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. Details on differences between personal and enterprise user rights are available in [Usage Guidelines](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s UpdateSmsTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateShrinkRequest) SetApplySceneContent(v string) *UpdateSmsTemplateShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetIntlType(v int32) *UpdateSmsTemplateShrinkRequest {
	s.IntlType = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetMoreDataShrink(v string) *UpdateSmsTemplateShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetOwnerId(v int64) *UpdateSmsTemplateShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetRelatedSignName(v string) *UpdateSmsTemplateShrinkRequest {
	s.RelatedSignName = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetRemark(v string) *UpdateSmsTemplateShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetResourceOwnerAccount(v string) *UpdateSmsTemplateShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetResourceOwnerId(v int64) *UpdateSmsTemplateShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateCode(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateContent(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateName(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateRule(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateRule = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateType(v int32) *UpdateSmsTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

type UpdateSmsTemplateResponseBody struct {
	// Request status code.
	//
	// 	- OK indicates a successful request.
	//
	// 	- For other error codes, refer to the **Error Codes*	- section of this chapter or the product\\"s [API Error Codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Work order number.
	//
	// This parameter is used by auditors when querying audits. You need to provide this work order number when requesting expedited review.
	//
	// example:
	//
	// 20041271****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Code of the SMS template.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Name of the SMS template.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateResponseBody) SetCode(v string) *UpdateSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmsTemplateResponseBody) SetMessage(v string) *UpdateSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmsTemplateResponseBody) SetOrderId(v string) *UpdateSmsTemplateResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateSmsTemplateResponseBody) SetRequestId(v string) *UpdateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmsTemplateResponseBody) SetTemplateCode(v string) *UpdateSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateResponseBody) SetTemplateName(v string) *UpdateSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

type UpdateSmsTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateResponse) SetHeaders(v map[string]*string) *UpdateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsTemplateResponse) SetStatusCode(v int32) *UpdateSmsTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsTemplateResponse) SetBody(v *UpdateSmsTemplateResponseBody) *UpdateSmsTemplateResponse {
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
		"ap-southeast-5": tea.String("dysmsapi.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     tea.String("dysmsapi-proxy.cn-beijing.aliyuncs.com"),
		"cn-hongkong":    tea.String("dysmsapi-xman.cn-hongkong.aliyuncs.com"),
		"eu-central-1":   tea.String("dysmsapi.eu-central-1.aliyuncs.com"),
		"us-east-1":      tea.String("dysmsapi.us-east-1.aliyuncs.com"),
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

// Summary:
//
// 添加验证码签名信息
//
// @param request - AddExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddExtCodeSignResponse
func (client *Client) AddExtCodeSignWithOptions(request *AddExtCodeSignRequest, runtime *util.RuntimeOptions) (_result *AddExtCodeSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtCode)) {
		query["ExtCode"] = request.ExtCode
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

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddExtCodeSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加验证码签名信息
//
// @param request - AddExtCodeSignRequest
//
// @return AddExtCodeSignResponse
func (client *Client) AddExtCodeSign(request *AddExtCodeSignRequest) (_result *AddExtCodeSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddExtCodeSignResponse{}
	_body, _err := client.AddExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a short URL.
//
// Description:
//
//   Before you call this operation, you must register the primary domain name of the source URL in the Short Message Service (SMS) console. After the domain name is registered, you can call this operation to create a short URL. For more information, see [Domain name registration](https://help.aliyun.com/document_detail/302325.html#title-mau-zdh-hd0).
//
// 	- You can create up to 3,000 short URLs within a natural day.
//
// 	- After a short URL is generated, a security review is required. Generally, the review takes 10 minutes to 2 hours to complete. Before the security review is passed, the short URL cannot be directly accessed.
//
// @param request - AddShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShortUrlResponse
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

// Summary:
//
// Creates a short URL.
//
// Description:
//
//   Before you call this operation, you must register the primary domain name of the source URL in the Short Message Service (SMS) console. After the domain name is registered, you can call this operation to create a short URL. For more information, see [Domain name registration](https://help.aliyun.com/document_detail/302325.html#title-mau-zdh-hd0).
//
// 	- You can create up to 3,000 short URLs within a natural day.
//
// 	- After a short URL is generated, a security review is required. Generally, the review takes 10 minutes to 2 hours to complete. Before the security review is passed, the short URL cannot be directly accessed.
//
// @param request - AddShortUrlRequest
//
// @return AddShortUrlResponse
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

// Summary:
//
// Creates a signature.
//
// Description:
//
// You can call the AddSmsSign operation or use the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm#/overview) to create an SMS signature. The signature must comply with the [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html). You can call the QuerySmsSign operation or use the SMS console to query the review status of the signature.
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation only once per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
// 	- You cannot cancel the review of a signature.
//
// 	- Individual users can create only one verification code signature, and can create only one general-purpose signature within a natural day. If you need to apply for multiple signatures, we recommend that you upgrade your account to an enterprise user.
//
// 	- If you need to use the same signature for messages sent to recipients both in and outside the Chinese mainland, the signature must be a general-purpose signature.
//
// 	- If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
// 	- An SMS signature must undergo a thorough review process before it can be approved for use.
//
// @param request - AddSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsSignResponse
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

// Summary:
//
// Creates a signature.
//
// Description:
//
// You can call the AddSmsSign operation or use the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm#/overview) to create an SMS signature. The signature must comply with the [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html). You can call the QuerySmsSign operation or use the SMS console to query the review status of the signature.
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation only once per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
// 	- You cannot cancel the review of a signature.
//
// 	- Individual users can create only one verification code signature, and can create only one general-purpose signature within a natural day. If you need to apply for multiple signatures, we recommend that you upgrade your account to an enterprise user.
//
// 	- If you need to use the same signature for messages sent to recipients both in and outside the Chinese mainland, the signature must be a general-purpose signature.
//
// 	- If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
// 	- An SMS signature must undergo a thorough review process before it can be approved for use.
//
// @param request - AddSmsSignRequest
//
// @return AddSmsSignResponse
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

// Summary:
//
// Creates a message template.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to apply for a message template. The template must comply with the [message template specifications](https://help.aliyun.com/document_detail/108253.html). You can call the [QuerySmsTemplate](https://help.aliyun.com/document_detail/419289.html) operation or use the Alibaba Cloud SMS console to check whether the message template is approved.
//
// >
//
// 	- Message templates pending approval can be withdrawn. You can withdraw a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- If you call the AddSmsTemplate operation, you can apply for a maximum of 100 message templates in a calendar day. After you apply for a message template, we recommend that you wait for at least 30 seconds before you apply for another one. If you use the Alibaba Cloud SMS console, you can apply for an unlimited number of message templates.
//
// 	- Messages sent to the Chinese mainland and messages sent to countries or regions outside the Chinese mainland use separate message templates. Create message templates based on your needs.
//
// 	- If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
// 	- A signature must undergo a thorough review process before it can be approved for use. For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsTemplateResponse
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

// Summary:
//
// Creates a message template.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to apply for a message template. The template must comply with the [message template specifications](https://help.aliyun.com/document_detail/108253.html). You can call the [QuerySmsTemplate](https://help.aliyun.com/document_detail/419289.html) operation or use the Alibaba Cloud SMS console to check whether the message template is approved.
//
// >
//
// 	- Message templates pending approval can be withdrawn. You can withdraw a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- If you call the AddSmsTemplate operation, you can apply for a maximum of 100 message templates in a calendar day. After you apply for a message template, we recommend that you wait for at least 30 seconds before you apply for another one. If you use the Alibaba Cloud SMS console, you can apply for an unlimited number of message templates.
//
// 	- Messages sent to the Chinese mainland and messages sent to countries or regions outside the Chinese mainland use separate message templates. Create message templates based on your needs.
//
// 	- If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
// 	- A signature must undergo a thorough review process before it can be approved for use. For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSmsTemplateRequest
//
// @return AddSmsTemplateResponse
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

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 2,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMobilesCardSupportResponse
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

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 2,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckMobilesCardSupportRequest
//
// @return CheckMobilesCardSupportResponse
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

// Summary:
//
// Sends conversion rate information to Alibaba Cloud SMS.
//
// @param request - ConversionDataIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversionDataIntlResponse
func (client *Client) ConversionDataIntlWithOptions(request *ConversionDataIntlRequest, runtime *util.RuntimeOptions) (_result *ConversionDataIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversionRate)) {
		query["ConversionRate"] = request.ConversionRate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReportTime)) {
		query["ReportTime"] = request.ReportTime
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
		Action:      tea.String("ConversionDataIntl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConversionDataIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends conversion rate information to Alibaba Cloud SMS.
//
// @param request - ConversionDataIntlRequest
//
// @return ConversionDataIntlResponse
func (client *Client) ConversionDataIntl(request *ConversionDataIntlRequest) (_result *ConversionDataIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConversionDataIntlResponse{}
	_body, _err := client.ConversionDataIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a card message template.
//
// Description:
//
//   The CreateCardSmsTemplate operation saves the card message template information, submits it to the mobile phone manufacturer for approval, and returns the message template ID.
//
// 	- If the type of the message template is not supported or events that are not supported by the mobile phone manufacturer are specified, the template is not submitted. For more information, see [Supported message templates](https://help.aliyun.com/document_detail/434611.html).
//
// 	- For information about sample card message templates, see [Sample card message templates](https://help.aliyun.com/document_detail/435361.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreateCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCardSmsTemplateResponse
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

// Summary:
//
// Creates a card message template.
//
// Description:
//
//   The CreateCardSmsTemplate operation saves the card message template information, submits it to the mobile phone manufacturer for approval, and returns the message template ID.
//
// 	- If the type of the message template is not supported or events that are not supported by the mobile phone manufacturer are specified, the template is not submitted. For more information, see [Supported message templates](https://help.aliyun.com/document_detail/434611.html).
//
// 	- For information about sample card message templates, see [Sample card message templates](https://help.aliyun.com/document_detail/435361.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateCardSmsTemplateRequest
//
// @return CreateCardSmsTemplateResponse
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

// Summary:
//
// 创建短链
//
// @param request - CreateSmartShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartShortUrlResponse
func (client *Client) CreateSmartShortUrlWithOptions(request *CreateSmartShortUrlRequest, runtime *util.RuntimeOptions) (_result *CreateSmartShortUrlResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SourceUrl)) {
		query["SourceUrl"] = request.SourceUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSmartShortUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSmartShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短链
//
// @param request - CreateSmartShortUrlRequest
//
// @return CreateSmartShortUrlResponse
func (client *Client) CreateSmartShortUrl(request *CreateSmartShortUrlRequest) (_result *CreateSmartShortUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmartShortUrlResponse{}
	_body, _err := client.CreateSmartShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create SMS Signature
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Individual authenticated users can apply for one formal signature per natural day under the same Alibaba Cloud account, while enterprise authenticated users have no current restrictions. For details on the differences in rights between individual and enterprise users, please refer to [User Guide](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm).
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// - After submitting the signature application, you can query the signature review status and details via the [GetSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-getsmssign?spm) interface. You can also [Configure Receipt Messages](https://help.aliyun.com/zh/sms/developer-reference/configure-delivery-receipts-1?spm) and obtain signature review status messages through SignSmsReport.
//
// @param tmpReq - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithOptions(tmpReq *CreateSmsSignRequest, runtime *util.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSmsSignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MoreData)) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, tea.String("MoreData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplySceneContent)) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !tea.BoolValue(util.IsUnset(request.MoreDataShrink)) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
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

	if !tea.BoolValue(util.IsUnset(request.ThirdParty)) {
		query["ThirdParty"] = request.ThirdParty
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create SMS Signature
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Individual authenticated users can apply for one formal signature per natural day under the same Alibaba Cloud account, while enterprise authenticated users have no current restrictions. For details on the differences in rights between individual and enterprise users, please refer to [User Guide](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm).
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// - After submitting the signature application, you can query the signature review status and details via the [GetSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-getsmssign?spm) interface. You can also [Configure Receipt Messages](https://help.aliyun.com/zh/sms/developer-reference/configure-delivery-receipts-1?spm) and obtain signature review status messages through SignSmsReport.
//
// @param request - CreateSmsSignRequest
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on the Update of SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - It is recommended to apply for SMS templates via the API with at least a 30-second interval between each request.
//
// - The template information applied through the API will be synchronized in the SMS service console. For operations related to templates in the console, please refer to SMS Templates.
//
// - After submitting the template application, you can query the audit status and details using the GetSmsTemplate interface. You can also configure delivery receipts to obtain the audit status messages via TemplateSmsReport.
//
// - Domestic SMS templates are not interchangeable with international/Hong Kong, Macao, and Taiwan SMS templates. Please apply for templates based on your business scenario.
//
// - Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For differences in rights between personal and enterprise users, please refer to Usage Instructions.
//
// @param tmpReq - CreateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplateWithOptions(tmpReq *CreateSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MoreData)) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, tea.String("MoreData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplySceneContent)) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !tea.BoolValue(util.IsUnset(request.IntlType)) {
		query["IntlType"] = request.IntlType
	}

	if !tea.BoolValue(util.IsUnset(request.MoreDataShrink)) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedSignName)) {
		query["RelatedSignName"] = request.RelatedSignName
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

	if !tea.BoolValue(util.IsUnset(request.TemplateRule)) {
		query["TemplateRule"] = request.TemplateRule
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on the Update of SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - It is recommended to apply for SMS templates via the API with at least a 30-second interval between each request.
//
// - The template information applied through the API will be synchronized in the SMS service console. For operations related to templates in the console, please refer to SMS Templates.
//
// - After submitting the template application, you can query the audit status and details using the GetSmsTemplate interface. You can also configure delivery receipts to obtain the audit status messages via TemplateSmsReport.
//
// - Domestic SMS templates are not interchangeable with international/Hong Kong, Macao, and Taiwan SMS templates. Please apply for templates based on your business scenario.
//
// - Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For differences in rights between personal and enterprise users, please refer to Usage Instructions.
//
// @param request - CreateSmsTemplateRequest
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplate(request *CreateSmsTemplateRequest) (_result *CreateSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CreateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除验证码签名
//
// @param request - DeleteExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExtCodeSignResponse
func (client *Client) DeleteExtCodeSignWithOptions(request *DeleteExtCodeSignRequest, runtime *util.RuntimeOptions) (_result *DeleteExtCodeSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtCode)) {
		query["ExtCode"] = request.ExtCode
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

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExtCodeSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除验证码签名
//
// @param request - DeleteExtCodeSignRequest
//
// @return DeleteExtCodeSignResponse
func (client *Client) DeleteExtCodeSign(request *DeleteExtCodeSignRequest) (_result *DeleteExtCodeSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExtCodeSignResponse{}
	_body, _err := client.DeleteExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a short URL. After you delete a short URL, it cannot be changed to its original state.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteShortUrlResponse
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

// Summary:
//
// Deletes a short URL. After you delete a short URL, it cannot be changed to its original state.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteShortUrlRequest
//
// @return DeleteShortUrlResponse
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

// Summary:
//
// Deletes a signature.
//
// Description:
//
//   You cannot delete a signature that has not been approved.
//
// 	- After you delete a signature, you cannot recover it. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsSignResponse
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

// Summary:
//
// Deletes a signature.
//
// Description:
//
//   You cannot delete a signature that has not been approved.
//
// 	- After you delete a signature, you cannot recover it. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsSignRequest
//
// @return DeleteSmsSignResponse
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

// Summary:
//
// Deletes a message template.
//
// Description:
//
//   Message templates pending approval can be withdrawn. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- You cannot recover deleted message templates. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsTemplateResponse
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

// Summary:
//
// Deletes a message template.
//
// Description:
//
//   Message templates pending approval can be withdrawn. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
// 	- You cannot recover deleted message templates. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsTemplateRequest
//
// @return DeleteSmsTemplateResponse
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

// Summary:
//
// 查询卡片发送详情
//
// @param request - GetCardSmsDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsDetailsResponse
func (client *Client) GetCardSmsDetailsWithOptions(request *GetCardSmsDetailsRequest, runtime *util.RuntimeOptions) (_result *GetCardSmsDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCardId)) {
		query["BizCardId"] = request.BizCardId
	}

	if !tea.BoolValue(util.IsUnset(request.BizDigitId)) {
		query["BizDigitId"] = request.BizDigitId
	}

	if !tea.BoolValue(util.IsUnset(request.BizSmsId)) {
		query["BizSmsId"] = request.BizSmsId
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
		Action:      tea.String("GetCardSmsDetails"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardSmsDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询卡片发送详情
//
// @param request - GetCardSmsDetailsRequest
//
// @return GetCardSmsDetailsResponse
func (client *Client) GetCardSmsDetails(request *GetCardSmsDetailsRequest) (_result *GetCardSmsDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardSmsDetailsResponse{}
	_body, _err := client.GetCardSmsDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the short URLs of a card messages template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCardSmsLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsLinkResponse
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

// Summary:
//
// Queries the short URLs of a card messages template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCardSmsLinkRequest
//
// @return GetCardSmsLinkResponse
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

// Summary:
//
// Converts a resource uploaded to the specified Object Storage Service (OSS) bucket for unified management. Then, a resource ID is returned. You can manage the resource based on the ID.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMediaResourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaResourceIdResponse
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

// Summary:
//
// Converts a resource uploaded to the specified Object Storage Service (OSS) bucket for unified management. Then, a resource ID is returned. You can manage the resource based on the ID.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMediaResourceIdRequest
//
// @return GetMediaResourceIdResponse
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

// Summary:
//
// Queries the OSS configuration information about card messages.
//
// Description:
//
// Resources such as images and videos used for card message templates can be uploaded to Object Storage Service (OSS) buckets for storage. For more information, see [Upload files to OSS](https://help.aliyun.com/document_detail/437303.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetOSSInfoForCardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSInfoForCardTemplateResponse
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

// Summary:
//
// Queries the OSS configuration information about card messages.
//
// Description:
//
// Resources such as images and videos used for card message templates can be uploaded to Object Storage Service (OSS) buckets for storage. For more information, see [Upload files to OSS](https://help.aliyun.com/document_detail/437303.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return GetOSSInfoForCardTemplateResponse
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

// Summary:
//
// SMS File Upload, Get Authorization Info
//
// Description:
//
// - When creating signatures or templates, you can upload materials such as login pages with links, backend page screenshots, software copyrights, supplementary agreements, etc. This helps the review personnel understand your business details. If there are multiple materials, they can be combined into one file, supporting png, jpg, jpeg, doc, docx, pdf formats.
//
// - For additional materials needed when creating signatures or templates, you can upload them to the OSS file system for storage. For file upload operations, refer to [OSS File Upload](https://help.aliyun.com/zh/sms/upload-files-through-oss).
//
// @param request - GetOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSInfoForUploadFileResponse
func (client *Client) GetOSSInfoForUploadFileWithOptions(request *GetOSSInfoForUploadFileRequest, runtime *util.RuntimeOptions) (_result *GetOSSInfoForUploadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOSSInfoForUploadFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOSSInfoForUploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// SMS File Upload, Get Authorization Info
//
// Description:
//
// - When creating signatures or templates, you can upload materials such as login pages with links, backend page screenshots, software copyrights, supplementary agreements, etc. This helps the review personnel understand your business details. If there are multiple materials, they can be combined into one file, supporting png, jpg, jpeg, doc, docx, pdf formats.
//
// - For additional materials needed when creating signatures or templates, you can upload them to the OSS file system for storage. For file upload operations, refer to [OSS File Upload](https://help.aliyun.com/zh/sms/upload-files-through-oss).
//
// @param request - GetOSSInfoForUploadFileRequest
//
// @return GetOSSInfoForUploadFileResponse
func (client *Client) GetOSSInfoForUploadFile(request *GetOSSInfoForUploadFileRequest) (_result *GetOSSInfoForUploadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOSSInfoForUploadFileResponse{}
	_body, _err := client.GetOSSInfoForUploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query SMS Signature Details
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Generally, after submitting the signature, Alibaba Cloud expects to complete the review within 2 hours (Review Business Hours: Monday to Sunday 9:00~21:00, with legal holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the signature fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm), invoke the [UpdateSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmssign?spm) API, or modify the unapproved SMS signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page.
//
// @param request - GetSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSignWithOptions(request *GetSmsSignRequest, runtime *util.RuntimeOptions) (_result *GetSmsSignResponse, _err error) {
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
		Action:      tea.String("GetSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query SMS Signature Details
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Generally, after submitting the signature, Alibaba Cloud expects to complete the review within 2 hours (Review Business Hours: Monday to Sunday 9:00~21:00, with legal holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the signature fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm), invoke the [UpdateSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmssign?spm) API, or modify the unapproved SMS signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page.
//
// @param request - GetSmsSignRequest
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSign(request *GetSmsSignRequest) (_result *GetSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSmsSignResponse{}
	_body, _err := client.GetSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Text SMS Template Details
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Under normal circumstances, Alibaba Cloud expects to complete the review within 2 hours after template submission (review working hours: Monday to Sunday 9:00~21:00, with statutory holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the template fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.41fd339f3bPSCQ), invoke the [ModifySmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-modifysmstemplate?spm=a2c4g.11186623.0.0.5b1f6e8bQloFit) API or modify the SMS template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// - The current QuerySmsTemplate interface queries the audit details of a single template by template code. The [QuerySmsTemplateList](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-querysmstemplatelist?spm=a2c4g.11186623.0.0.24086e8bO8cFn4) interface can query the template details of all templates under your current account.
//
// @param request - GetSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsTemplateResponse
func (client *Client) GetSmsTemplateWithOptions(request *GetSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *GetSmsTemplateResponse, _err error) {
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
		Action:      tea.String("GetSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Text SMS Template Details
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Under normal circumstances, Alibaba Cloud expects to complete the review within 2 hours after template submission (review working hours: Monday to Sunday 9:00~21:00, with statutory holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the template fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.41fd339f3bPSCQ), invoke the [ModifySmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-modifysmstemplate?spm=a2c4g.11186623.0.0.5b1f6e8bQloFit) API or modify the SMS template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// - The current QuerySmsTemplate interface queries the audit details of a single template by template code. The [QuerySmsTemplateList](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-querysmstemplatelist?spm=a2c4g.11186623.0.0.24086e8bO8cFn4) interface can query the template details of all templates under your current account.
//
// @param request - GetSmsTemplateRequest
//
// @return GetSmsTemplateResponse
func (client *Client) GetSmsTemplate(request *GetSmsTemplateRequest) (_result *GetSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSmsTemplateResponse{}
	_body, _err := client.GetSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries the tags of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Modifies a rejected signature and submit it for approval. Signatures that are pending approval or have been approved cannot be modified.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to modify an existing signature and submit the signature for approval. The signature must comply with the [signature specifications](https://help.aliyun.com/document_detail/108076.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
// 	- Signatures pending approval cannot be modified.
//
// 	- You cannot modify a signature after it is approved. If you no longer need the signature, you can delete it.
//
// 	- If you are an individual user, you cannot apply for a new signature on the same day that your signature is rejected or deleted. We recommend that you modify the rejected signature and submit it again.
//
// @param request - ModifySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsSignResponse
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

// Summary:
//
// Modifies a rejected signature and submit it for approval. Signatures that are pending approval or have been approved cannot be modified.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to modify an existing signature and submit the signature for approval. The signature must comply with the [signature specifications](https://help.aliyun.com/document_detail/108076.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
// 	- Signatures pending approval cannot be modified.
//
// 	- You cannot modify a signature after it is approved. If you no longer need the signature, you can delete it.
//
// 	- If you are an individual user, you cannot apply for a new signature on the same day that your signature is rejected or deleted. We recommend that you modify the rejected signature and submit it again.
//
// @param request - ModifySmsSignRequest
//
// @return ModifySmsSignResponse
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

// Summary:
//
// Modifies the information of an unapproved message template and submits it for review again.
//
// Description:
//
// After you apply for a message template, if the template fails to pass the review, you can call this operation to modify the template and submit the template again. You can call this operation to modify only a template for a specific message type.
//
// The template content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsTemplateResponse
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

// Summary:
//
// Modifies the information of an unapproved message template and submits it for review again.
//
// Description:
//
// After you apply for a message template, if the template fails to pass the review, you can call this operation to modify the template and submit the template again. You can call this operation to modify only a template for a specific message type.
//
// The template content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifySmsTemplateRequest
//
// @return ModifySmsTemplateResponse
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

// Summary:
//
// Queries the review status of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateResponse
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

// Summary:
//
// Queries the review status of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateRequest
//
// @return QueryCardSmsTemplateResponse
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

// Summary:
//
// Queries sent card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateReportResponse
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

// Summary:
//
// Queries sent card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateReportRequest
//
// @return QueryCardSmsTemplateReportResponse
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

// Summary:
//
// 查询验证码签名
//
// @param request - QueryExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExtCodeSignResponse
func (client *Client) QueryExtCodeSignWithOptions(request *QueryExtCodeSignRequest, runtime *util.RuntimeOptions) (_result *QueryExtCodeSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtCode)) {
		query["ExtCode"] = request.ExtCode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryExtCodeSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询验证码签名
//
// @param request - QueryExtCodeSignRequest
//
// @return QueryExtCodeSignResponse
func (client *Client) QueryExtCodeSign(request *QueryExtCodeSignRequest) (_result *QueryExtCodeSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryExtCodeSignResponse{}
	_body, _err := client.QueryExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// @param tmpReq - QueryMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMobilesCardSupportResponse
func (client *Client) QueryMobilesCardSupportWithOptions(tmpReq *QueryMobilesCardSupportRequest, runtime *util.RuntimeOptions) (_result *QueryMobilesCardSupportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryMobilesCardSupportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Mobiles)) {
		request.MobilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Mobiles, tea.String("Mobiles"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MobilesShrink)) {
		query["Mobiles"] = request.MobilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMobilesCardSupport"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMobilesCardSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// @param request - QueryMobilesCardSupportRequest
//
// @return QueryMobilesCardSupportResponse
func (client *Client) QueryMobilesCardSupport(request *QueryMobilesCardSupportRequest) (_result *QueryMobilesCardSupportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMobilesCardSupportResponse{}
	_body, _err := client.QueryMobilesCardSupportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 点击明细查询
//
// @param request - QueryPageSmartShortUrlLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPageSmartShortUrlLogResponse
func (client *Client) QueryPageSmartShortUrlLogWithOptions(request *QueryPageSmartShortUrlLogRequest, runtime *util.RuntimeOptions) (_result *QueryPageSmartShortUrlLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateDateEnd)) {
		query["CreateDateEnd"] = request.CreateDateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateDateStart)) {
		query["CreateDateStart"] = request.CreateDateStart
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
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

	if !tea.BoolValue(util.IsUnset(request.ShortUrl)) {
		query["ShortUrl"] = request.ShortUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPageSmartShortUrlLog"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPageSmartShortUrlLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 点击明细查询
//
// @param request - QueryPageSmartShortUrlLogRequest
//
// @return QueryPageSmartShortUrlLogResponse
func (client *Client) QueryPageSmartShortUrlLog(request *QueryPageSmartShortUrlLogRequest) (_result *QueryPageSmartShortUrlLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPageSmartShortUrlLogResponse{}
	_body, _err := client.QueryPageSmartShortUrlLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a message.
//
// @param request - QuerySendDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendDetailsResponse
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

// Summary:
//
// Queries the information about a message.
//
// @param request - QuerySendDetailsRequest
//
// @return QuerySendDetailsResponse
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

// Summary:
//
// Queries message delivery details.
//
// Description:
//
// You can call the operation to query message delivery details, including the number of delivered messages, the number of messages with delivery receipts, and the time that a message is sent. If a large number of messages are sent on the specified date, you can specify the number of items displayed on each page and the number of pages to view the details by page.
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySendStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendStatisticsResponse
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

// Summary:
//
// Queries message delivery details.
//
// Description:
//
// You can call the operation to query message delivery details, including the number of delivered messages, the number of messages with delivery receipts, and the time that a message is sent. If a large number of messages are sent on the specified date, you can specify the number of items displayed on each page and the number of pages to view the details by page.
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySendStatisticsRequest
//
// @return QuerySendStatisticsResponse
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

// Summary:
//
// Queries the status of a short URL.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShortUrlResponse
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

// Summary:
//
// Queries the status of a short URL.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryShortUrlRequest
//
// @return QueryShortUrlResponse
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

// Summary:
//
// Queries the status of a signature.
//
// Description:
//
// After you apply for an SMS signature, you can query its status by using the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm) or calling the operation. If the signature is rejected, you can modify the signature based on the reason why it is rejected.
//
// ### QPS limits
//
// You can call this API operation up to 500 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignResponse
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

// Summary:
//
// Queries the status of a signature.
//
// Description:
//
// After you apply for an SMS signature, you can query its status by using the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm) or calling the operation. If the signature is rejected, you can modify the signature based on the reason why it is rejected.
//
// ### QPS limits
//
// You can call this API operation up to 500 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignRequest
//
// @return QuerySmsSignResponse
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

// Summary:
//
// Queries message signatures by page.
//
// Description:
//
// You can call this operation to query the details of message signatures, including the name, creation time, and approval status of each signature. If a message template is rejected, the reason is returned. Modify the message signature based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignListResponse
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

// Summary:
//
// Queries message signatures by page.
//
// Description:
//
// You can call this operation to query the details of message signatures, including the name, creation time, and approval status of each signature. If a message template is rejected, the reason is returned. Modify the message signature based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignListRequest
//
// @return QuerySmsSignListResponse
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

// Summary:
//
// Queries the approval status of a message template.
//
// Description:
//
// After you create a message template, you can call this operation to query the approval status of the template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateResponse
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

// Summary:
//
// Queries the approval status of a message template.
//
// Description:
//
// After you create a message template, you can call this operation to query the approval status of the template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateRequest
//
// @return QuerySmsTemplateResponse
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

// Summary:
//
// Queries message templates.
//
// Description:
//
// You can call this operation to query the details of message templates, including the name, creation time, and approval status of each template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateListResponse
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

// Summary:
//
// Queries message templates.
//
// Description:
//
// You can call this operation to query the details of message templates, including the name, creation time, and approval status of each template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateListRequest
//
// @return QuerySmsTemplateListResponse
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

// Summary:
//
// Sends multiple card messages at a time.
//
// Description:
//
// You can call the operation to send multiple card messages to a maximum of mobile phone numbers at a time. Different signatures and rollback settings can be specified for the mobile phone numbers.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendBatchCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchCardSmsResponse
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

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParamJson)) {
		query["TemplateParamJson"] = request.TemplateParamJson
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

// Summary:
//
// Sends multiple card messages at a time.
//
// Description:
//
// You can call the operation to send multiple card messages to a maximum of mobile phone numbers at a time. Different signatures and rollback settings can be specified for the mobile phone numbers.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendBatchCardSmsRequest
//
// @return SendBatchCardSmsResponse
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

// Summary:
//
// Uses a single message template and multiple signatures to send messages to multiple recipients.
//
// Description:
//
// You can call the operation to send messages to a maximum of 100 recipients at a time.
//
// @param request - SendBatchSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchSmsResponse
func (client *Client) SendBatchSmsWithOptions(request *SendBatchSmsRequest, runtime *util.RuntimeOptions) (_result *SendBatchSmsResponse, _err error) {
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

// Summary:
//
// Uses a single message template and multiple signatures to send messages to multiple recipients.
//
// Description:
//
// You can call the operation to send messages to a maximum of 100 recipients at a time.
//
// @param request - SendBatchSmsRequest
//
// @return SendBatchSmsResponse
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

// Summary:
//
// Sends a card message.
//
// Description:
//
//   Make sure that the message template that you want to use has been approved. If the mobile phone number of a recipient does not support card messages, the SendCardSms operation allows the rollback feature to ensure successful delivery.
//
// 	- When you call the SendCardSms operation to send card messages, the operation checks whether the mobile phone numbers of the recipients support card messages. If the mobile phone numbers do not support card messages, you can specify whether to enable rollback. Otherwise, the card message cannot be delivered.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCardSmsResponse
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

// Summary:
//
// Sends a card message.
//
// Description:
//
//   Make sure that the message template that you want to use has been approved. If the mobile phone number of a recipient does not support card messages, the SendCardSms operation allows the rollback feature to ensure successful delivery.
//
// 	- When you call the SendCardSms operation to send card messages, the operation checks whether the mobile phone numbers of the recipients support card messages. If the mobile phone numbers do not support card messages, you can specify whether to enable rollback. Otherwise, the card message cannot be delivered.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendCardSmsRequest
//
// @return SendCardSmsResponse
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

// Summary:
//
// Sends a message. Before you call this operation, submit a message signature and message template, and make sure that the signature and template are approved.
//
// Description:
//
//   This operation is mainly used to send a single message. In special scenarios, you can send multiple messages with the same content to a maximum of 1,000 mobile numbers. Note that group sending may be delayed.
//
// 	- To send messages with different signatures and template content to multiple mobile numbers in a single request, call the [SendBatchSms](https://help.aliyun.com/document_detail/102364.html) operation.
//
// 	- You are charged for using Alibaba Cloud Short Message Service (SMS) based on the amount of messages sent. For more information, see [Pricing](https://www.aliyun.com/price/product#/sms/detail).
//
// 	- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
//
// @param request - SendSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSmsResponse
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

// Summary:
//
// Sends a message. Before you call this operation, submit a message signature and message template, and make sure that the signature and template are approved.
//
// Description:
//
//   This operation is mainly used to send a single message. In special scenarios, you can send multiple messages with the same content to a maximum of 1,000 mobile numbers. Note that group sending may be delayed.
//
// 	- To send messages with different signatures and template content to multiple mobile numbers in a single request, call the [SendBatchSms](https://help.aliyun.com/document_detail/102364.html) operation.
//
// 	- You are charged for using Alibaba Cloud Short Message Service (SMS) based on the amount of messages sent. For more information, see [Pricing](https://www.aliyun.com/price/product#/sms/detail).
//
// 	- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
//
// @param request - SendSmsRequest
//
// @return SendSmsResponse
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

// Summary:
//
// Reports the status of an OTP message to Alibaba Cloud SMS.
//
// Description:
//
// Metrics:
//
// 	- Requested OTP messages
//
// 	- Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations: 1. Call the SmsConversion operation in an asynchronous manner by configuring queues or events. 2. Manually degrade your services or use a circuit breaker to automatically degrade services.
//
// @param request - SmsConversionIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsConversionIntlResponse
func (client *Client) SmsConversionIntlWithOptions(request *SmsConversionIntlRequest, runtime *util.RuntimeOptions) (_result *SmsConversionIntlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversionTime)) {
		query["ConversionTime"] = request.ConversionTime
	}

	if !tea.BoolValue(util.IsUnset(request.Delivered)) {
		query["Delivered"] = request.Delivered
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsConversionIntl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsConversionIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports the status of an OTP message to Alibaba Cloud SMS.
//
// Description:
//
// Metrics:
//
// 	- Requested OTP messages
//
// 	- Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations: 1. Call the SmsConversion operation in an asynchronous manner by configuring queues or events. 2. Manually degrade your services or use a circuit breaker to automatically degrade services.
//
// @param request - SmsConversionIntlRequest
//
// @return SmsConversionIntlResponse
func (client *Client) SmsConversionIntl(request *SmsConversionIntlRequest) (_result *SmsConversionIntlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmsConversionIntlResponse{}
	_body, _err := client.SmsConversionIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches tags to a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
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

// Summary:
//
// Attaches tags to a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Deletes tags from a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
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

// Summary:
//
// Deletes tags from a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// 修改验证码签名
//
// @param request - UpdateExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtCodeSignResponse
func (client *Client) UpdateExtCodeSignWithOptions(request *UpdateExtCodeSignRequest, runtime *util.RuntimeOptions) (_result *UpdateExtCodeSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExistExtCode)) {
		query["ExistExtCode"] = request.ExistExtCode
	}

	if !tea.BoolValue(util.IsUnset(request.NewExtCode)) {
		query["NewExtCode"] = request.NewExtCode
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

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExtCodeSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改验证码签名
//
// @param request - UpdateExtCodeSignRequest
//
// @return UpdateExtCodeSignResponse
func (client *Client) UpdateExtCodeSign(request *UpdateExtCodeSignRequest) (_result *UpdateExtCodeSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExtCodeSignResponse{}
	_body, _err := client.UpdateExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update Text SMS Signature
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only signatures that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm) and call this interface to modify and resubmit for review after modification.
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, please see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// @param tmpReq - UpdateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsSignResponse
func (client *Client) UpdateSmsSignWithOptions(tmpReq *UpdateSmsSignRequest, runtime *util.RuntimeOptions) (_result *UpdateSmsSignResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSmsSignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MoreData)) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, tea.String("MoreData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplySceneContent)) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !tea.BoolValue(util.IsUnset(request.MoreDataShrink)) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
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

	if !tea.BoolValue(util.IsUnset(request.ThirdParty)) {
		query["ThirdParty"] = request.ThirdParty
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSmsSign"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Text SMS Signature
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only signatures that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm) and call this interface to modify and resubmit for review after modification.
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, please see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// @param request - UpdateSmsSignRequest
//
// @return UpdateSmsSignResponse
func (client *Client) UpdateSmsSign(request *UpdateSmsSignRequest) (_result *UpdateSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.UpdateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update Text SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on SMS Service Update: Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only templates that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Template Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.4bf5561ajcFtMQ) and call this interface to modify and resubmit for review.
//
// - Modifications made through the interface will be synchronized in the SMS service console. For related operations on templates in the console, see [SMS Templates](https://help.aliyun.com/zh/sms/user-guide/message-templates/?spm=a2c4g.11186623.0.0.35a947464Itaxp).
//
// ### QPS Limit
//
// The single-user QPS limit for this interface is 1000 times/second. Exceeding this limit will result in API throttling, which may impact your business. Please make calls reasonably.
//
// @param tmpReq - UpdateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsTemplateResponse
func (client *Client) UpdateSmsTemplateWithOptions(tmpReq *UpdateSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MoreData)) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, tea.String("MoreData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplySceneContent)) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !tea.BoolValue(util.IsUnset(request.IntlType)) {
		query["IntlType"] = request.IntlType
	}

	if !tea.BoolValue(util.IsUnset(request.MoreDataShrink)) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedSignName)) {
		query["RelatedSignName"] = request.RelatedSignName
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

	if !tea.BoolValue(util.IsUnset(request.TemplateRule)) {
		query["TemplateRule"] = request.TemplateRule
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSmsTemplate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Text SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on SMS Service Update: Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only templates that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Template Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.4bf5561ajcFtMQ) and call this interface to modify and resubmit for review.
//
// - Modifications made through the interface will be synchronized in the SMS service console. For related operations on templates in the console, see [SMS Templates](https://help.aliyun.com/zh/sms/user-guide/message-templates/?spm=a2c4g.11186623.0.0.35a947464Itaxp).
//
// ### QPS Limit
//
// The single-user QPS limit for this interface is 1000 times/second. Exceeding this limit will result in API throttling, which may impact your business. Please make calls reasonably.
//
// @param request - UpdateSmsTemplateRequest
//
// @return UpdateSmsTemplateResponse
func (client *Client) UpdateSmsTemplate(request *UpdateSmsTemplateRequest) (_result *UpdateSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.UpdateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
