// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddSmsSignRequest struct {
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string                          `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignSource           *int32                           `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	Remark               *string                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SignFileList         []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
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

func (s *AddSmsSignRequest) SetResourceOwnerAccount(v string) *AddSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsSignRequest) SetResourceOwnerId(v int64) *AddSmsSignRequest {
	s.ResourceOwnerId = &v
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

func (s *AddSmsSignRequest) SetRemark(v string) *AddSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *AddSmsSignRequest) SetSignFileList(v []*AddSmsSignRequestSignFileList) *AddSmsSignRequest {
	s.SignFileList = v
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
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateContent      *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *AddSmsTemplateRequest) SetResourceOwnerAccount(v string) *AddSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsTemplateRequest) SetResourceOwnerId(v int64) *AddSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateType(v int32) *AddSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateName(v string) *AddSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateContent(v string) *AddSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *AddSmsTemplateRequest) SetRemark(v string) *AddSmsTemplateRequest {
	s.Remark = &v
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

type ModifySmsSignRequest struct {
	OwnerId              *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string                             `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignSource           *int32                              `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	Remark               *string                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SignFileList         []*ModifySmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
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

func (s *ModifySmsSignRequest) SetResourceOwnerAccount(v string) *ModifySmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsSignRequest) SetResourceOwnerId(v int64) *ModifySmsSignRequest {
	s.ResourceOwnerId = &v
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

func (s *ModifySmsSignRequest) SetRemark(v string) *ModifySmsSignRequest {
	s.Remark = &v
	return s
}

func (s *ModifySmsSignRequest) SetSignFileList(v []*ModifySmsSignRequestSignFileList) *ModifySmsSignRequest {
	s.SignFileList = v
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
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateType         *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateContent      *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *ModifySmsTemplateRequest) SetResourceOwnerAccount(v string) *ModifySmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetResourceOwnerId(v int64) *ModifySmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateType(v int32) *ModifySmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateName(v string) *ModifySmsTemplateRequest {
	s.TemplateName = &v
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

func (s *ModifySmsTemplateRequest) SetRemark(v string) *ModifySmsTemplateRequest {
	s.Remark = &v
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

type QuerySendDetailsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SendDate             *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage          *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QuerySendDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsRequest) SetOwnerId(v int64) *QuerySendDetailsRequest {
	s.OwnerId = &v
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

func (s *QuerySendDetailsRequest) SetPhoneNumber(v string) *QuerySendDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QuerySendDetailsRequest) SetBizId(v string) *QuerySendDetailsRequest {
	s.BizId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetSendDate(v string) *QuerySendDetailsRequest {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPageSize(v int64) *QuerySendDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendDetailsRequest) SetCurrentPage(v int64) *QuerySendDetailsRequest {
	s.CurrentPage = &v
	return s
}

type QuerySendDetailsResponseBody struct {
	Code              *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message           *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *string                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SmsSendDetailDTOs *QuerySendDetailsResponseBodySmsSendDetailDTOs `json:"SmsSendDetailDTOs,omitempty" xml:"SmsSendDetailDTOs,omitempty" type:"Struct"`
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

func (s *QuerySendDetailsResponseBody) SetTotalCount(v string) *QuerySendDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetSmsSendDetailDTOs(v *QuerySendDetailsResponseBodySmsSendDetailDTOs) *QuerySendDetailsResponseBody {
	s.SmsSendDetailDTOs = v
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
	ErrCode      *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	OutId        *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	ReceiveDate  *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	SendDate     *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	PhoneNum     *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SendStatus   *int64  `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetErrCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.ErrCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetTemplateCode(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.TemplateCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetOutId(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.OutId = &v
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

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetPhoneNum(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.PhoneNum = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetContent(v string) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.Content = &v
	return s
}

func (s *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO) SetSendStatus(v int64) *QuerySendDetailsResponseBodySmsSendDetailDTOsSmsSendDetailDTO {
	s.SendStatus = &v
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignStatus *int32  `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QuerySmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsSignResponseBody) SetRequestId(v string) *QuerySmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignStatus(v int32) *QuerySmsSignResponseBody {
	s.SignStatus = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetCode(v string) *QuerySmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetMessage(v string) *QuerySmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetCreateDate(v string) *QuerySmsSignResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetReason(v string) *QuerySmsSignResponseBody {
	s.Reason = &v
	return s
}

func (s *QuerySmsSignResponseBody) SetSignName(v string) *QuerySmsSignResponseBody {
	s.SignName = &v
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
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateStatus  *int32  `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s QuerySmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateResponseBody) SetTemplateContent(v string) *QuerySmsTemplateResponseBody {
	s.TemplateContent = &v
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

func (s *QuerySmsTemplateResponseBody) SetTemplateStatus(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetCode(v string) *QuerySmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateType(v int32) *QuerySmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetMessage(v string) *QuerySmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetTemplateName(v string) *QuerySmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetCreateDate(v string) *QuerySmsTemplateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsTemplateResponseBody) SetReason(v string) *QuerySmsTemplateResponseBody {
	s.Reason = &v
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
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumberJson      *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	SignNameJson         *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParamJson    *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
	SmsUpExtendCodeJson  *string `json:"SmsUpExtendCodeJson,omitempty" xml:"SmsUpExtendCodeJson,omitempty"`
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

func (s *SendBatchSmsRequest) SetResourceOwnerAccount(v string) *SendBatchSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendBatchSmsRequest) SetResourceOwnerId(v int64) *SendBatchSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendBatchSmsRequest) SetPhoneNumberJson(v string) *SendBatchSmsRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchSmsRequest) SetSignNameJson(v string) *SendBatchSmsRequest {
	s.SignNameJson = &v
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

func (s *SendBatchSmsRequest) SetSmsUpExtendCodeJson(v string) *SendBatchSmsRequest {
	s.SmsUpExtendCodeJson = &v
	return s
}

type SendBatchSmsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendBatchSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBatchSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchSmsResponseBody) SetCode(v string) *SendBatchSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetMessage(v string) *SendBatchSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendBatchSmsResponseBody) SetBizId(v string) *SendBatchSmsResponseBody {
	s.BizId = &v
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

type SendMessageToGlobeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	To                   *string `json:"To,omitempty" xml:"To,omitempty"`
	From                 *string `json:"From,omitempty" xml:"From,omitempty"`
	Message              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeRequest) GoString() string {
	return s.String()
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

func (s *SendMessageToGlobeRequest) SetFrom(v string) *SendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetMessage(v string) *SendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetType(v string) *SendMessageToGlobeRequest {
	s.Type = &v
	return s
}

type SendMessageToGlobeResponseBody struct {
	From         *string                                     `json:"From,omitempty" xml:"From,omitempty"`
	MessageId    *string                                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Segments     *string                                     `json:"Segments,omitempty" xml:"Segments,omitempty"`
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	To           *string                                     `json:"To,omitempty" xml:"To,omitempty"`
	NumberDetail *SendMessageToGlobeResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
}

func (s SendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBody) SetFrom(v string) *SendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetMessageId(v string) *SendMessageToGlobeResponseBody {
	s.MessageId = &v
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

func (s *SendMessageToGlobeResponseBody) SetCode(v string) *SendMessageToGlobeResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetTo(v string) *SendMessageToGlobeResponseBody {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetNumberDetail(v *SendMessageToGlobeResponseBodyNumberDetail) *SendMessageToGlobeResponseBody {
	s.NumberDetail = v
	return s
}

type SendMessageToGlobeResponseBodyNumberDetail struct {
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SendMessageToGlobeResponseBodyNumberDetail) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCountry(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCarrier(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Carrier = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam        *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	SmsUpExtendCode      *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s SendSmsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsRequest) GoString() string {
	return s.String()
}

func (s *SendSmsRequest) SetOwnerId(v int64) *SendSmsRequest {
	s.OwnerId = &v
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

func (s *SendSmsRequest) SetPhoneNumbers(v string) *SendSmsRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendSmsRequest) SetSignName(v string) *SendSmsRequest {
	s.SignName = &v
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

func (s *SendSmsRequest) SetSmsUpExtendCode(v string) *SendSmsRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsRequest) SetOutId(v string) *SendSmsRequest {
	s.OutId = &v
	return s
}

type SendSmsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendSmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsResponseBody) SetCode(v string) *SendSmsResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsResponseBody) SetMessage(v string) *SendSmsResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsResponseBody) SetBizId(v string) *SendSmsResponseBody {
	s.BizId = &v
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
