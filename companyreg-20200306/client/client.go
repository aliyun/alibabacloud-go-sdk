// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BindProduceAuthorizationRequest struct {
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType           *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s BindProduceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s BindProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationRequest) SetAuthorizedUserIds(v string) *BindProduceAuthorizationRequest {
	s.AuthorizedUserIds = &v
	return s
}

func (s *BindProduceAuthorizationRequest) SetBizId(v string) *BindProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *BindProduceAuthorizationRequest) SetBizType(v string) *BindProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

type BindProduceAuthorizationResponseBody struct {
	Data      *BindProduceAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindProduceAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBody) SetData(v *BindProduceAuthorizationResponseBodyData) *BindProduceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetErrorCode(v string) *BindProduceAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetErrorMsg(v string) *BindProduceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetRequestId(v string) *BindProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type BindProduceAuthorizationResponseBodyData struct {
	AuthorizedUserList []*BindProduceAuthorizationResponseBodyDataAuthorizedUserList `json:"AuthorizedUserList,omitempty" xml:"AuthorizedUserList,omitempty" type:"Repeated"`
	Message            *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindProduceAuthorizationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindProduceAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBodyData) SetAuthorizedUserList(v []*BindProduceAuthorizationResponseBodyDataAuthorizedUserList) *BindProduceAuthorizationResponseBodyData {
	s.AuthorizedUserList = v
	return s
}

func (s *BindProduceAuthorizationResponseBodyData) SetMessage(v string) *BindProduceAuthorizationResponseBodyData {
	s.Message = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyData) SetSuccess(v bool) *BindProduceAuthorizationResponseBodyData {
	s.Success = &v
	return s
}

type BindProduceAuthorizationResponseBodyDataAuthorizedUserList struct {
	AccountValidType *int32  `json:"AccountValidType,omitempty" xml:"AccountValidType,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BindProduceAuthorizationResponseBodyDataAuthorizedUserList) String() string {
	return tea.Prettify(s)
}

func (s BindProduceAuthorizationResponseBodyDataAuthorizedUserList) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetAccountValidType(v int32) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.AccountValidType = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetUserId(v string) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.UserId = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetUserName(v string) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.UserName = &v
	return s
}

type BindProduceAuthorizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindProduceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindProduceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s BindProduceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponse) SetHeaders(v map[string]*string) *BindProduceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *BindProduceAuthorizationResponse) SetStatusCode(v int32) *BindProduceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *BindProduceAuthorizationResponse) SetBody(v *BindProduceAuthorizationResponseBody) *BindProduceAuthorizationResponse {
	s.Body = v
	return s
}

type CloseIntentionForPartnerRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CloseIntentionForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseIntentionForPartnerRequest) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerRequest) SetBizType(v string) *CloseIntentionForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *CloseIntentionForPartnerRequest) SetIntentionBizId(v string) *CloseIntentionForPartnerRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CloseIntentionForPartnerRequest) SetNote(v string) *CloseIntentionForPartnerRequest {
	s.Note = &v
	return s
}

type CloseIntentionForPartnerResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseIntentionForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseIntentionForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerResponseBody) SetErrorCode(v string) *CloseIntentionForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetErrorMsg(v string) *CloseIntentionForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetRequestId(v string) *CloseIntentionForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetSuccess(v bool) *CloseIntentionForPartnerResponseBody {
	s.Success = &v
	return s
}

type CloseIntentionForPartnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseIntentionForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseIntentionForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseIntentionForPartnerResponse) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerResponse) SetHeaders(v map[string]*string) *CloseIntentionForPartnerResponse {
	s.Headers = v
	return s
}

func (s *CloseIntentionForPartnerResponse) SetStatusCode(v int32) *CloseIntentionForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseIntentionForPartnerResponse) SetBody(v *CloseIntentionForPartnerResponseBody) *CloseIntentionForPartnerResponse {
	s.Body = v
	return s
}

type CloseUserIntentionRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CloseUserIntentionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseUserIntentionRequest) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionRequest) SetBizType(v string) *CloseUserIntentionRequest {
	s.BizType = &v
	return s
}

func (s *CloseUserIntentionRequest) SetIntentionBizId(v string) *CloseUserIntentionRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CloseUserIntentionRequest) SetNote(v string) *CloseUserIntentionRequest {
	s.Note = &v
	return s
}

type CloseUserIntentionResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseUserIntentionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseUserIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionResponseBody) SetErrorCode(v string) *CloseUserIntentionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetErrorMsg(v string) *CloseUserIntentionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetRequestId(v string) *CloseUserIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetSuccess(v bool) *CloseUserIntentionResponseBody {
	s.Success = &v
	return s
}

type CloseUserIntentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseUserIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseUserIntentionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseUserIntentionResponse) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionResponse) SetHeaders(v map[string]*string) *CloseUserIntentionResponse {
	s.Headers = v
	return s
}

func (s *CloseUserIntentionResponse) SetStatusCode(v int32) *CloseUserIntentionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseUserIntentionResponse) SetBody(v *CloseUserIntentionResponseBody) *CloseUserIntentionResponse {
	s.Body = v
	return s
}

type CreateBusinessOpportunityRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Source      *int32  `json:"Source,omitempty" xml:"Source,omitempty"`
	VCode       *string `json:"VCode,omitempty" xml:"VCode,omitempty"`
}

func (s CreateBusinessOpportunityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessOpportunityRequest) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityRequest) SetBizType(v string) *CreateBusinessOpportunityRequest {
	s.BizType = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetContactName(v string) *CreateBusinessOpportunityRequest {
	s.ContactName = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetMobile(v string) *CreateBusinessOpportunityRequest {
	s.Mobile = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetSource(v int32) *CreateBusinessOpportunityRequest {
	s.Source = &v
	return s
}

func (s *CreateBusinessOpportunityRequest) SetVCode(v string) *CreateBusinessOpportunityRequest {
	s.VCode = &v
	return s
}

type CreateBusinessOpportunityResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBusinessOpportunityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessOpportunityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityResponseBody) SetErrorCode(v string) *CreateBusinessOpportunityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetErrorMessage(v string) *CreateBusinessOpportunityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetRequestId(v string) *CreateBusinessOpportunityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetSuccess(v bool) *CreateBusinessOpportunityResponseBody {
	s.Success = &v
	return s
}

type CreateBusinessOpportunityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBusinessOpportunityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBusinessOpportunityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessOpportunityResponse) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityResponse) SetHeaders(v map[string]*string) *CreateBusinessOpportunityResponse {
	s.Headers = v
	return s
}

func (s *CreateBusinessOpportunityResponse) SetStatusCode(v int32) *CreateBusinessOpportunityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBusinessOpportunityResponse) SetBody(v *CreateBusinessOpportunityResponseBody) *CreateBusinessOpportunityResponse {
	s.Body = v
	return s
}

type CreateProduceForPartnerRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s CreateProduceForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProduceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerRequest) SetBizId(v string) *CreateProduceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *CreateProduceForPartnerRequest) SetBizType(v string) *CreateProduceForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *CreateProduceForPartnerRequest) SetExtInfo(v string) *CreateProduceForPartnerRequest {
	s.ExtInfo = &v
	return s
}

type CreateProduceForPartnerResponseBody struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProduceForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProduceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerResponseBody) SetBizId(v string) *CreateProduceForPartnerResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetErrorCode(v string) *CreateProduceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetErrorMsg(v string) *CreateProduceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetRequestId(v string) *CreateProduceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetSuccess(v bool) *CreateProduceForPartnerResponseBody {
	s.Success = &v
	return s
}

type CreateProduceForPartnerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProduceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProduceForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProduceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerResponse) SetHeaders(v map[string]*string) *CreateProduceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *CreateProduceForPartnerResponse) SetStatusCode(v int32) *CreateProduceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProduceForPartnerResponse) SetBody(v *CreateProduceForPartnerResponseBody) *CreateProduceForPartnerResponse {
	s.Body = v
	return s
}

type DescribePartnerConfigRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
}

func (s DescribePartnerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigRequest) SetBizType(v string) *DescribePartnerConfigRequest {
	s.BizType = &v
	return s
}

func (s *DescribePartnerConfigRequest) SetPartnerCode(v string) *DescribePartnerConfigRequest {
	s.PartnerCode = &v
	return s
}

type DescribePartnerConfigResponseBody struct {
	Contact     *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	PartnerName *string `json:"PartnerName,omitempty" xml:"PartnerName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePartnerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigResponseBody) SetContact(v string) *DescribePartnerConfigResponseBody {
	s.Contact = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetPartnerCode(v string) *DescribePartnerConfigResponseBody {
	s.PartnerCode = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetPartnerName(v string) *DescribePartnerConfigResponseBody {
	s.PartnerName = &v
	return s
}

func (s *DescribePartnerConfigResponseBody) SetRequestId(v string) *DescribePartnerConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribePartnerConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePartnerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePartnerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigResponse) SetHeaders(v map[string]*string) *DescribePartnerConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePartnerConfigResponse) SetStatusCode(v int32) *DescribePartnerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePartnerConfigResponse) SetBody(v *DescribePartnerConfigResponseBody) *DescribePartnerConfigResponse {
	s.Body = v
	return s
}

type GenerateUploadFilePolicyRequest struct {
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GenerateUploadFilePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyRequest) SetBizType(v string) *GenerateUploadFilePolicyRequest {
	s.BizType = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) SetFileName(v string) *GenerateUploadFilePolicyRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) SetFileType(v string) *GenerateUploadFilePolicyRequest {
	s.FileType = &v
	return s
}

type GenerateUploadFilePolicyResponseBody struct {
	// OSSAccessKeyId
	AccessId      *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileDir       *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	FileUrl       *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// OSS Endpoint。
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadFilePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponseBody) SetAccessId(v string) *GenerateUploadFilePolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetEncodedPolicy(v string) *GenerateUploadFilePolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetExpireTime(v string) *GenerateUploadFilePolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileDir(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileUrl(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetHost(v string) *GenerateUploadFilePolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetSignature(v string) *GenerateUploadFilePolicyResponseBody {
	s.Signature = &v
	return s
}

type GenerateUploadFilePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateUploadFilePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateUploadFilePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetStatusCode(v int32) *GenerateUploadFilePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetBody(v *GenerateUploadFilePolicyResponseBody) *GenerateUploadFilePolicyResponse {
	s.Body = v
	return s
}

type GetAlipayUrlRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlipayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlRequest) SetBizType(v string) *GetAlipayUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetAlipayUrlRequest) SetOrderId(v int64) *GetAlipayUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetAlipayUrlRequest) SetReturnUrl(v string) *GetAlipayUrlRequest {
	s.ReturnUrl = &v
	return s
}

func (s *GetAlipayUrlRequest) SetType(v string) *GetAlipayUrlRequest {
	s.Type = &v
	return s
}

type GetAlipayUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlipayUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponseBody) SetData(v string) *GetAlipayUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetAlipayUrlResponseBody) SetRequestId(v string) *GetAlipayUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAlipayUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlipayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlipayUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponse) SetHeaders(v map[string]*string) *GetAlipayUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAlipayUrlResponse) SetStatusCode(v int32) *GetAlipayUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlipayUrlResponse) SetBody(v *GetAlipayUrlResponseBody) *GetAlipayUrlResponse {
	s.Body = v
	return s
}

type ListIntentionNoteRequest struct {
	BeginTime      *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIntentionNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntentionNoteRequest) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteRequest) SetBeginTime(v int64) *ListIntentionNoteRequest {
	s.BeginTime = &v
	return s
}

func (s *ListIntentionNoteRequest) SetEndTime(v int64) *ListIntentionNoteRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntentionNoteRequest) SetIntentionBizId(v string) *ListIntentionNoteRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListIntentionNoteRequest) SetPageNumber(v int32) *ListIntentionNoteRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIntentionNoteRequest) SetPageSize(v int32) *ListIntentionNoteRequest {
	s.PageSize = &v
	return s
}

type ListIntentionNoteResponseBody struct {
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListIntentionNoteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                               `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                               `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListIntentionNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntentionNoteResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponseBody) SetCurrentPageNum(v int32) *ListIntentionNoteResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetData(v []*ListIntentionNoteResponseBodyData) *ListIntentionNoteResponseBody {
	s.Data = v
	return s
}

func (s *ListIntentionNoteResponseBody) SetPageSize(v int32) *ListIntentionNoteResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetRequestId(v string) *ListIntentionNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetTotalItemNum(v int32) *ListIntentionNoteResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetTotalPageNum(v int32) *ListIntentionNoteResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListIntentionNoteResponseBodyData struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Source         *int32  `json:"Source,omitempty" xml:"Source,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIntentionNoteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntentionNoteResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponseBodyData) SetCreateTime(v string) *ListIntentionNoteResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetIntentionBizId(v string) *ListIntentionNoteResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetNote(v string) *ListIntentionNoteResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetSource(v int32) *ListIntentionNoteResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetType(v string) *ListIntentionNoteResponseBodyData {
	s.Type = &v
	return s
}

type ListIntentionNoteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIntentionNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntentionNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntentionNoteResponse) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponse) SetHeaders(v map[string]*string) *ListIntentionNoteResponse {
	s.Headers = v
	return s
}

func (s *ListIntentionNoteResponse) SetStatusCode(v int32) *ListIntentionNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentionNoteResponse) SetBody(v *ListIntentionNoteResponseBody) *ListIntentionNoteResponse {
	s.Body = v
	return s
}

type ListProduceAuthorizationRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListProduceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationRequest) SetBizId(v string) *ListProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetBizType(v string) *ListProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetPageNum(v int32) *ListProduceAuthorizationRequest {
	s.PageNum = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetPageSize(v int32) *ListProduceAuthorizationRequest {
	s.PageSize = &v
	return s
}

type ListProduceAuthorizationResponseBody struct {
	CurrentPageNum *int32                                      `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListProduceAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalItemNum   *int32                                      `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                      `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListProduceAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationResponseBody) SetCurrentPageNum(v int32) *ListProduceAuthorizationResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetData(v []*ListProduceAuthorizationResponseBodyData) *ListProduceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetPageSize(v int32) *ListProduceAuthorizationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetRequestId(v string) *ListProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetSuccess(v bool) *ListProduceAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetTotalItemNum(v int32) *ListProduceAuthorizationResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetTotalPageNum(v int32) *ListProduceAuthorizationResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListProduceAuthorizationResponseBodyData struct {
	AuthorizedUserId   *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	AuthorizedUserName *string `json:"AuthorizedUserName,omitempty" xml:"AuthorizedUserName,omitempty"`
}

func (s ListProduceAuthorizationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProduceAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationResponseBodyData) SetAuthorizedUserId(v string) *ListProduceAuthorizationResponseBodyData {
	s.AuthorizedUserId = &v
	return s
}

func (s *ListProduceAuthorizationResponseBodyData) SetAuthorizedUserName(v string) *ListProduceAuthorizationResponseBodyData {
	s.AuthorizedUserName = &v
	return s
}

type ListProduceAuthorizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProduceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProduceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProduceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationResponse) SetHeaders(v map[string]*string) *ListProduceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ListProduceAuthorizationResponse) SetStatusCode(v int32) *ListProduceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProduceAuthorizationResponse) SetBody(v *ListProduceAuthorizationResponseBody) *ListProduceAuthorizationResponse {
	s.Body = v
	return s
}

type ListUserDetailSolutionsRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDetailSolutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDetailSolutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsRequest) SetBizType(v string) *ListUserDetailSolutionsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetIntentionBizId(v string) *ListUserDetailSolutionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetPageNum(v int32) *ListUserDetailSolutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetPageSize(v int32) *ListUserDetailSolutionsRequest {
	s.PageSize = &v
	return s
}

type ListUserDetailSolutionsResponseBody struct {
	CurrentPageNum *int32                                     `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserDetailSolutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                     `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                     `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserDetailSolutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDetailSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponseBody) SetCurrentPageNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetData(v []*ListUserDetailSolutionsResponseBodyData) *ListUserDetailSolutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetPageSize(v int32) *ListUserDetailSolutionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetRequestId(v string) *ListUserDetailSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetTotalItemNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetTotalPageNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListUserDetailSolutionsResponseBodyData struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeliveryOrderBizId   *string `json:"DeliveryOrderBizId,omitempty" xml:"DeliveryOrderBizId,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	IntentionAssignBizId *string `json:"IntentionAssignBizId,omitempty" xml:"IntentionAssignBizId,omitempty"`
	IntentionBizId       *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PartnerCode          *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	Reason               *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime           *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserDetailSolutionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserDetailSolutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponseBodyData) SetBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetBizType(v string) *ListUserDetailSolutionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetCreateTime(v int64) *ListUserDetailSolutionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetDeliveryOrderBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.DeliveryOrderBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetExtInfo(v string) *ListUserDetailSolutionsResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetIntentionAssignBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.IntentionAssignBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetIntentionBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetPartnerCode(v string) *ListUserDetailSolutionsResponseBodyData {
	s.PartnerCode = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetReason(v string) *ListUserDetailSolutionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetStatus(v int32) *ListUserDetailSolutionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetUpdateTime(v int64) *ListUserDetailSolutionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetUserId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.UserId = &v
	return s
}

type ListUserDetailSolutionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserDetailSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserDetailSolutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDetailSolutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponse) SetHeaders(v map[string]*string) *ListUserDetailSolutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserDetailSolutionsResponse) SetStatusCode(v int32) *ListUserDetailSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDetailSolutionsResponse) SetBody(v *ListUserDetailSolutionsResponseBody) *ListUserDetailSolutionsResponse {
	s.Body = v
	return s
}

type ListUserIntentionNotesRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserIntentionNotesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionNotesRequest) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesRequest) SetBizType(v string) *ListUserIntentionNotesRequest {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetIntentionBizId(v string) *ListUserIntentionNotesRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetPageNum(v int32) *ListUserIntentionNotesRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionNotesRequest) SetPageSize(v int32) *ListUserIntentionNotesRequest {
	s.PageSize = &v
	return s
}

type ListUserIntentionNotesResponseBody struct {
	Data         []*ListUserIntentionNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum      *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalItemNum *int32                                    `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum *int32                                    `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserIntentionNotesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponseBody) SetData(v []*ListUserIntentionNotesResponseBodyData) *ListUserIntentionNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetPageNum(v int32) *ListUserIntentionNotesResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetPageSize(v int32) *ListUserIntentionNotesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetRequestId(v string) *ListUserIntentionNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetSuccess(v bool) *ListUserIntentionNotesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetTotalItemNum(v int32) *ListUserIntentionNotesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetTotalPageNum(v int32) *ListUserIntentionNotesResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListUserIntentionNotesResponseBodyData struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Note       *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s ListUserIntentionNotesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponseBodyData) SetCreateTime(v string) *ListUserIntentionNotesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserIntentionNotesResponseBodyData) SetNote(v string) *ListUserIntentionNotesResponseBodyData {
	s.Note = &v
	return s
}

type ListUserIntentionNotesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserIntentionNotesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserIntentionNotesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionNotesResponse) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponse) SetHeaders(v map[string]*string) *ListUserIntentionNotesResponse {
	s.Headers = v
	return s
}

func (s *ListUserIntentionNotesResponse) SetStatusCode(v int32) *ListUserIntentionNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserIntentionNotesResponse) SetBody(v *ListUserIntentionNotesResponseBody) *ListUserIntentionNotesResponse {
	s.Body = v
	return s
}

type ListUserIntentionsRequest struct {
	Area           *string `json:"Area,omitempty" xml:"Area,omitempty"`
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizTypes       *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortFiled      *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
	SortOrder      *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WithExtInfo    *bool   `json:"WithExtInfo,omitempty" xml:"WithExtInfo,omitempty"`
}

func (s ListUserIntentionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsRequest) SetArea(v string) *ListUserIntentionsRequest {
	s.Area = &v
	return s
}

func (s *ListUserIntentionsRequest) SetBizType(v string) *ListUserIntentionsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionsRequest) SetBizTypes(v string) *ListUserIntentionsRequest {
	s.BizTypes = &v
	return s
}

func (s *ListUserIntentionsRequest) SetIntentionBizId(v string) *ListUserIntentionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserIntentionsRequest) SetPageNum(v int32) *ListUserIntentionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionsRequest) SetPageSize(v int32) *ListUserIntentionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionsRequest) SetSortFiled(v string) *ListUserIntentionsRequest {
	s.SortFiled = &v
	return s
}

func (s *ListUserIntentionsRequest) SetSortOrder(v string) *ListUserIntentionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserIntentionsRequest) SetStatus(v int32) *ListUserIntentionsRequest {
	s.Status = &v
	return s
}

func (s *ListUserIntentionsRequest) SetWithExtInfo(v bool) *ListUserIntentionsRequest {
	s.WithExtInfo = &v
	return s
}

type ListUserIntentionsResponseBody struct {
	CurrentPageNum *int32                                `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserIntentionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserIntentionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponseBody) SetCurrentPageNum(v int32) *ListUserIntentionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetData(v []*ListUserIntentionsResponseBodyData) *ListUserIntentionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserIntentionsResponseBody) SetPageSize(v int32) *ListUserIntentionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetRequestId(v string) *ListUserIntentionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetTotalItemNum(v int32) *ListUserIntentionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetTotalPageNum(v int32) *ListUserIntentionsResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListUserIntentionsResponseBodyData struct {
	Area        *string `json:"Area,omitempty" xml:"Area,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Ext         *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserIntentionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponseBodyData) SetArea(v string) *ListUserIntentionsResponseBodyData {
	s.Area = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetBizId(v string) *ListUserIntentionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetBizType(v string) *ListUserIntentionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetContactName(v string) *ListUserIntentionsResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetCreateTime(v int64) *ListUserIntentionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetDescription(v string) *ListUserIntentionsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetExt(v string) *ListUserIntentionsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetMobile(v string) *ListUserIntentionsResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetReason(v string) *ListUserIntentionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetStatus(v int32) *ListUserIntentionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetUpdateTime(v int64) *ListUserIntentionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetUserId(v string) *ListUserIntentionsResponseBodyData {
	s.UserId = &v
	return s
}

type ListUserIntentionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserIntentionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserIntentionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserIntentionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponse) SetHeaders(v map[string]*string) *ListUserIntentionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserIntentionsResponse) SetStatusCode(v int32) *ListUserIntentionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserIntentionsResponse) SetBody(v *ListUserIntentionsResponseBody) *ListUserIntentionsResponse {
	s.Body = v
	return s
}

type ListUserProduceOperateLogsRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserProduceOperateLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserProduceOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsRequest) SetBizId(v string) *ListUserProduceOperateLogsRequest {
	s.BizId = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetBizType(v string) *ListUserProduceOperateLogsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetPageNum(v int32) *ListUserProduceOperateLogsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserProduceOperateLogsRequest) SetPageSize(v int32) *ListUserProduceOperateLogsRequest {
	s.PageSize = &v
	return s
}

type ListUserProduceOperateLogsResponseBody struct {
	Data         []*ListUserProduceOperateLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum      *int32                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalItemNum *int32                                        `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum *int32                                        `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserProduceOperateLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserProduceOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponseBody) SetData(v []*ListUserProduceOperateLogsResponseBodyData) *ListUserProduceOperateLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetPageNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetPageSize(v int32) *ListUserProduceOperateLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetRequestId(v string) *ListUserProduceOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetSuccess(v bool) *ListUserProduceOperateLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetTotalItemNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBody) SetTotalPageNum(v int32) *ListUserProduceOperateLogsResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListUserProduceOperateLogsResponseBodyData struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizStatus       *int32  `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Note            *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OperateName     *string `json:"OperateName,omitempty" xml:"OperateName,omitempty"`
	OperateTime     *int64  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperateUserType *string `json:"OperateUserType,omitempty" xml:"OperateUserType,omitempty"`
	ToBizStatus     *int32  `json:"ToBizStatus,omitempty" xml:"ToBizStatus,omitempty"`
}

func (s ListUserProduceOperateLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserProduceOperateLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizId(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizStatus(v int32) *ListUserProduceOperateLogsResponseBodyData {
	s.BizStatus = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetBizType(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetNote(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateName(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateName = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateTime(v int64) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateTime = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetOperateUserType(v string) *ListUserProduceOperateLogsResponseBodyData {
	s.OperateUserType = &v
	return s
}

func (s *ListUserProduceOperateLogsResponseBodyData) SetToBizStatus(v int32) *ListUserProduceOperateLogsResponseBodyData {
	s.ToBizStatus = &v
	return s
}

type ListUserProduceOperateLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserProduceOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserProduceOperateLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserProduceOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProduceOperateLogsResponse) SetHeaders(v map[string]*string) *ListUserProduceOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProduceOperateLogsResponse) SetStatusCode(v int32) *ListUserProduceOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProduceOperateLogsResponse) SetBody(v *ListUserProduceOperateLogsResponseBody) *ListUserProduceOperateLogsResponse {
	s.Body = v
	return s
}

type ListUserSolutionsRequest struct {
	ExistStatus    []*int64 `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty" type:"Repeated"`
	IntentionBizId *string  `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNum        *int32   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserSolutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSolutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsRequest) SetExistStatus(v []*int64) *ListUserSolutionsRequest {
	s.ExistStatus = v
	return s
}

func (s *ListUserSolutionsRequest) SetIntentionBizId(v string) *ListUserSolutionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsRequest) SetPageNum(v int32) *ListUserSolutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserSolutionsRequest) SetPageSize(v int32) *ListUserSolutionsRequest {
	s.PageSize = &v
	return s
}

type ListUserSolutionsShrinkRequest struct {
	ExistStatusShrink *string `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	IntentionBizId    *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserSolutionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSolutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsShrinkRequest) SetExistStatusShrink(v string) *ListUserSolutionsShrinkRequest {
	s.ExistStatusShrink = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetIntentionBizId(v string) *ListUserSolutionsShrinkRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetPageNum(v int32) *ListUserSolutionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetPageSize(v int32) *ListUserSolutionsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListUserSolutionsResponseBody struct {
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserSolutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                               `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                               `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserSolutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponseBody) SetCurrentPageNum(v int32) *ListUserSolutionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetData(v []*ListUserSolutionsResponseBodyData) *ListUserSolutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserSolutionsResponseBody) SetPageSize(v int32) *ListUserSolutionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetRequestId(v string) *ListUserSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetTotalItemNum(v int32) *ListUserSolutionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetTotalPageNum(v int32) *ListUserSolutionsResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListUserSolutionsResponseBodyData struct {
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeliveryOrderBizId   *string `json:"DeliveryOrderBizId,omitempty" xml:"DeliveryOrderBizId,omitempty"`
	IntentionAssignBizId *string `json:"IntentionAssignBizId,omitempty" xml:"IntentionAssignBizId,omitempty"`
	IntentionBizId       *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PartnerCode          *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	Reason               *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime           *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserSolutionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserSolutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponseBodyData) SetBizId(v string) *ListUserSolutionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetBizType(v string) *ListUserSolutionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetCreateTime(v int64) *ListUserSolutionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetDeliveryOrderBizId(v string) *ListUserSolutionsResponseBodyData {
	s.DeliveryOrderBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetIntentionAssignBizId(v string) *ListUserSolutionsResponseBodyData {
	s.IntentionAssignBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetIntentionBizId(v string) *ListUserSolutionsResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetPartnerCode(v string) *ListUserSolutionsResponseBodyData {
	s.PartnerCode = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetReason(v string) *ListUserSolutionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetStatus(v int32) *ListUserSolutionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetUpdateTime(v int64) *ListUserSolutionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetUserId(v string) *ListUserSolutionsResponseBodyData {
	s.UserId = &v
	return s
}

type ListUserSolutionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserSolutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserSolutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponse) SetHeaders(v map[string]*string) *ListUserSolutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserSolutionsResponse) SetStatusCode(v int32) *ListUserSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSolutionsResponse) SetBody(v *ListUserSolutionsResponseBody) *ListUserSolutionsResponse {
	s.Body = v
	return s
}

type OperateProduceForPartnerRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtInfo     *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s OperateProduceForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateProduceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerRequest) SetBizId(v string) *OperateProduceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetBizType(v string) *OperateProduceForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetExtInfo(v string) *OperateProduceForPartnerRequest {
	s.ExtInfo = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetOperateType(v string) *OperateProduceForPartnerRequest {
	s.OperateType = &v
	return s
}

type OperateProduceForPartnerResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateProduceForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateProduceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerResponseBody) SetErrorCode(v string) *OperateProduceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetErrorMsg(v string) *OperateProduceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetRequestId(v string) *OperateProduceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetSuccess(v bool) *OperateProduceForPartnerResponseBody {
	s.Success = &v
	return s
}

type OperateProduceForPartnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OperateProduceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateProduceForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateProduceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerResponse) SetHeaders(v map[string]*string) *OperateProduceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateProduceForPartnerResponse) SetStatusCode(v int32) *OperateProduceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateProduceForPartnerResponse) SetBody(v *OperateProduceForPartnerResponseBody) *OperateProduceForPartnerResponse {
	s.Body = v
	return s
}

type PutMeasureDataRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureDataRequest) SetBizType(v string) *PutMeasureDataRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureDataRequest) SetData(v string) *PutMeasureDataRequest {
	s.Data = &v
	return s
}

func (s *PutMeasureDataRequest) SetDataType(v string) *PutMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataRequest) SetEndTime(v string) *PutMeasureDataRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureDataRequest) SetStartTime(v string) *PutMeasureDataRequest {
	s.StartTime = &v
	return s
}

type PutMeasureDataResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponseBody) SetData(v bool) *PutMeasureDataResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetRequestId(v string) *PutMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

type PutMeasureDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMeasureDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponse) SetHeaders(v map[string]*string) *PutMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureDataResponse) SetStatusCode(v int32) *PutMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureDataResponse) SetBody(v *PutMeasureDataResponseBody) *PutMeasureDataResponse {
	s.Body = v
	return s
}

type PutMeasureReadyFlagRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReadyFlag *string `json:"ReadyFlag,omitempty" xml:"ReadyFlag,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagRequest) SetBizType(v string) *PutMeasureReadyFlagRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetDataType(v string) *PutMeasureReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetEndTime(v string) *PutMeasureReadyFlagRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetReadyFlag(v string) *PutMeasureReadyFlagRequest {
	s.ReadyFlag = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetStartTime(v string) *PutMeasureReadyFlagRequest {
	s.StartTime = &v
	return s
}

type PutMeasureReadyFlagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponseBody) SetData(v bool) *PutMeasureReadyFlagResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureReadyFlagResponseBody) SetRequestId(v string) *PutMeasureReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

type PutMeasureReadyFlagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutMeasureReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMeasureReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponse) SetHeaders(v map[string]*string) *PutMeasureReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetStatusCode(v int32) *PutMeasureReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetBody(v *PutMeasureReadyFlagResponseBody) *PutMeasureReadyFlagResponse {
	s.Body = v
	return s
}

type QueryAvailableNumbersRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s QueryAvailableNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableNumbersRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersRequest) SetBizType(v string) *QueryAvailableNumbersRequest {
	s.BizType = &v
	return s
}

type QueryAvailableNumbersResponseBody struct {
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersResponseBody) SetData(v []*string) *QueryAvailableNumbersResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetErrorCode(v string) *QueryAvailableNumbersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetErrorMsg(v string) *QueryAvailableNumbersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetRequestId(v string) *QueryAvailableNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetSuccess(v bool) *QueryAvailableNumbersResponseBody {
	s.Success = &v
	return s
}

type QueryAvailableNumbersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAvailableNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAvailableNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableNumbersResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersResponse) SetHeaders(v map[string]*string) *QueryAvailableNumbersResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableNumbersResponse) SetStatusCode(v int32) *QueryAvailableNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableNumbersResponse) SetBody(v *QueryAvailableNumbersResponseBody) *QueryAvailableNumbersResponse {
	s.Body = v
	return s
}

type QueryBagRemainingRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s QueryBagRemainingRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBagRemainingRequest) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingRequest) SetBizType(v string) *QueryBagRemainingRequest {
	s.BizType = &v
	return s
}

type QueryBagRemainingResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryBagRemainingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBagRemainingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingResponseBody) SetData(v int64) *QueryBagRemainingResponseBody {
	s.Data = &v
	return s
}

func (s *QueryBagRemainingResponseBody) SetRequestId(v string) *QueryBagRemainingResponseBody {
	s.RequestId = &v
	return s
}

type QueryBagRemainingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBagRemainingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBagRemainingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBagRemainingResponse) GoString() string {
	return s.String()
}

func (s *QueryBagRemainingResponse) SetHeaders(v map[string]*string) *QueryBagRemainingResponse {
	s.Headers = v
	return s
}

func (s *QueryBagRemainingResponse) SetStatusCode(v int32) *QueryBagRemainingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBagRemainingResponse) SetBody(v *QueryBagRemainingResponseBody) *QueryBagRemainingResponse {
	s.Body = v
	return s
}

type QueryCommodityConfigRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	QueryModule   *bool   `json:"QueryModule,omitempty" xml:"QueryModule,omitempty"`
}

func (s QueryCommodityConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCommodityConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCommodityConfigRequest) SetBizType(v string) *QueryCommodityConfigRequest {
	s.BizType = &v
	return s
}

func (s *QueryCommodityConfigRequest) SetCommodityCode(v string) *QueryCommodityConfigRequest {
	s.CommodityCode = &v
	return s
}

func (s *QueryCommodityConfigRequest) SetQueryModule(v bool) *QueryCommodityConfigRequest {
	s.QueryModule = &v
	return s
}

type QueryCommodityConfigResponseBody struct {
	Data      *QueryCommodityConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCommodityConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCommodityConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommodityConfigResponseBody) SetData(v *QueryCommodityConfigResponseBodyData) *QueryCommodityConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryCommodityConfigResponseBody) SetRequestId(v string) *QueryCommodityConfigResponseBody {
	s.RequestId = &v
	return s
}

type QueryCommodityConfigResponseBodyData struct {
	CommodityCode    *string                                                 `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityModules []*QueryCommodityConfigResponseBodyDataCommodityModules `json:"CommodityModules,omitempty" xml:"CommodityModules,omitempty" type:"Repeated"`
	Description      *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	IconUrl          *string                                                 `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	ProductLine      *string                                                 `json:"ProductLine,omitempty" xml:"ProductLine,omitempty"`
	ProtocolUrl      *string                                                 `json:"ProtocolUrl,omitempty" xml:"ProtocolUrl,omitempty"`
	StartingPrice    *string                                                 `json:"StartingPrice,omitempty" xml:"StartingPrice,omitempty"`
	Type             *int32                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryCommodityConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCommodityConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCommodityConfigResponseBodyData) SetCommodityCode(v string) *QueryCommodityConfigResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetCommodityModules(v []*QueryCommodityConfigResponseBodyDataCommodityModules) *QueryCommodityConfigResponseBodyData {
	s.CommodityModules = v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetDescription(v string) *QueryCommodityConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetIconUrl(v string) *QueryCommodityConfigResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetProductLine(v string) *QueryCommodityConfigResponseBodyData {
	s.ProductLine = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetProtocolUrl(v string) *QueryCommodityConfigResponseBodyData {
	s.ProtocolUrl = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetStartingPrice(v string) *QueryCommodityConfigResponseBodyData {
	s.StartingPrice = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyData) SetType(v int32) *QueryCommodityConfigResponseBodyData {
	s.Type = &v
	return s
}

type QueryCommodityConfigResponseBodyDataCommodityModules struct {
	LxModuleCode      *string `json:"LxModuleCode,omitempty" xml:"LxModuleCode,omitempty"`
	ModuleCode        *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleDescription *string `json:"ModuleDescription,omitempty" xml:"ModuleDescription,omitempty"`
	ModuleName        *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleTip         *string `json:"ModuleTip,omitempty" xml:"ModuleTip,omitempty"`
	ModuleType        *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	ModuleUrl         *string `json:"ModuleUrl,omitempty" xml:"ModuleUrl,omitempty"`
	ModuleValue       *string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty"`
	SortNumber        *int32  `json:"SortNumber,omitempty" xml:"SortNumber,omitempty"`
}

func (s QueryCommodityConfigResponseBodyDataCommodityModules) String() string {
	return tea.Prettify(s)
}

func (s QueryCommodityConfigResponseBodyDataCommodityModules) GoString() string {
	return s.String()
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetLxModuleCode(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.LxModuleCode = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleCode(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleCode = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleDescription(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleDescription = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleName(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleName = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleTip(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleTip = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleType(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleType = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleUrl(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleUrl = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetModuleValue(v string) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.ModuleValue = &v
	return s
}

func (s *QueryCommodityConfigResponseBodyDataCommodityModules) SetSortNumber(v int32) *QueryCommodityConfigResponseBodyDataCommodityModules {
	s.SortNumber = &v
	return s
}

type QueryCommodityConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCommodityConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCommodityConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCommodityConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCommodityConfigResponse) SetHeaders(v map[string]*string) *QueryCommodityConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCommodityConfigResponse) SetStatusCode(v int32) *QueryCommodityConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommodityConfigResponse) SetBody(v *QueryCommodityConfigResponseBody) *QueryCommodityConfigResponse {
	s.Body = v
	return s
}

type QueryInstanceRequest struct {
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceRequest) SetBizType(v string) *QueryInstanceRequest {
	s.BizType = &v
	return s
}

func (s *QueryInstanceRequest) SetInstanceId(v string) *QueryInstanceRequest {
	s.InstanceId = &v
	return s
}

type QueryInstanceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponseBody) SetData(v string) *QueryInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInstanceResponseBody) SetRequestId(v string) *QueryInstanceResponseBody {
	s.RequestId = &v
	return s
}

type QueryInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceResponse) SetHeaders(v map[string]*string) *QueryInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceResponse) SetStatusCode(v int32) *QueryInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceResponse) SetBody(v *QueryInstanceResponseBody) *QueryInstanceResponse {
	s.Body = v
	return s
}

type QueryPartnerIntentionListRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPartnerIntentionListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerIntentionListRequest) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListRequest) SetBizId(v string) *QueryPartnerIntentionListRequest {
	s.BizId = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetBizType(v string) *QueryPartnerIntentionListRequest {
	s.BizType = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetPageNum(v int64) *QueryPartnerIntentionListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetPageSize(v int64) *QueryPartnerIntentionListRequest {
	s.PageSize = &v
	return s
}

type QueryPartnerIntentionListResponseBody struct {
	CurrentPageNum *int64                                       `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryPartnerIntentionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int64                                       `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int64                                       `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryPartnerIntentionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerIntentionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponseBody) SetCurrentPageNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetData(v []*QueryPartnerIntentionListResponseBodyData) *QueryPartnerIntentionListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetPageSize(v int64) *QueryPartnerIntentionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetRequestId(v string) *QueryPartnerIntentionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetTotalItemNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetTotalPageNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryPartnerIntentionListResponseBodyData struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Mobile  *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s QueryPartnerIntentionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerIntentionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponseBodyData) SetBizId(v string) *QueryPartnerIntentionListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBodyData) SetBizType(v string) *QueryPartnerIntentionListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBodyData) SetMobile(v string) *QueryPartnerIntentionListResponseBodyData {
	s.Mobile = &v
	return s
}

type QueryPartnerIntentionListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPartnerIntentionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPartnerIntentionListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerIntentionListResponse) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponse) SetHeaders(v map[string]*string) *QueryPartnerIntentionListResponse {
	s.Headers = v
	return s
}

func (s *QueryPartnerIntentionListResponse) SetStatusCode(v int32) *QueryPartnerIntentionListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPartnerIntentionListResponse) SetBody(v *QueryPartnerIntentionListResponseBody) *QueryPartnerIntentionListResponse {
	s.Body = v
	return s
}

type QueryPartnerProduceListRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPartnerProduceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerProduceListRequest) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListRequest) SetBizId(v string) *QueryPartnerProduceListRequest {
	s.BizId = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetBizType(v string) *QueryPartnerProduceListRequest {
	s.BizType = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetPageNum(v int64) *QueryPartnerProduceListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetPageSize(v int64) *QueryPartnerProduceListRequest {
	s.PageSize = &v
	return s
}

type QueryPartnerProduceListResponseBody struct {
	CurrentPageNum *int64                                     `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryPartnerProduceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int64                                     `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int64                                     `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryPartnerProduceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerProduceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponseBody) SetCurrentPageNum(v int64) *QueryPartnerProduceListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetData(v []*QueryPartnerProduceListResponseBodyData) *QueryPartnerProduceListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetPageSize(v int64) *QueryPartnerProduceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetRequestId(v string) *QueryPartnerProduceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetTotalItemNum(v int64) *QueryPartnerProduceListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetTotalPageNum(v int64) *QueryPartnerProduceListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryPartnerProduceListResponseBodyData struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Mobile  *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s QueryPartnerProduceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerProduceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponseBodyData) SetBizId(v string) *QueryPartnerProduceListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryPartnerProduceListResponseBodyData) SetBizType(v string) *QueryPartnerProduceListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *QueryPartnerProduceListResponseBodyData) SetMobile(v string) *QueryPartnerProduceListResponseBodyData {
	s.Mobile = &v
	return s
}

type QueryPartnerProduceListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPartnerProduceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPartnerProduceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerProduceListResponse) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponse) SetHeaders(v map[string]*string) *QueryPartnerProduceListResponse {
	s.Headers = v
	return s
}

func (s *QueryPartnerProduceListResponse) SetStatusCode(v int32) *QueryPartnerProduceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPartnerProduceListResponse) SetBody(v *QueryPartnerProduceListResponseBody) *QueryPartnerProduceListResponse {
	s.Body = v
	return s
}

type QueryUserNeedAuthResponseBody struct {
	NeedAuth  *bool   `json:"NeedAuth,omitempty" xml:"NeedAuth,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserNeedAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserNeedAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserNeedAuthResponseBody) SetNeedAuth(v bool) *QueryUserNeedAuthResponseBody {
	s.NeedAuth = &v
	return s
}

func (s *QueryUserNeedAuthResponseBody) SetRequestId(v string) *QueryUserNeedAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserNeedAuthResponseBody) SetSuccess(v bool) *QueryUserNeedAuthResponseBody {
	s.Success = &v
	return s
}

type QueryUserNeedAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserNeedAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserNeedAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserNeedAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryUserNeedAuthResponse) SetHeaders(v map[string]*string) *QueryUserNeedAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryUserNeedAuthResponse) SetStatusCode(v int32) *QueryUserNeedAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserNeedAuthResponse) SetBody(v *QueryUserNeedAuthResponseBody) *QueryUserNeedAuthResponse {
	s.Body = v
	return s
}

type RecordPostBackRequest struct {
	BizId     *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType   *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Contactor *string `json:"contactor,omitempty" xml:"contactor,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	EntityKey *string `json:"entityKey,omitempty" xml:"entityKey,omitempty"`
}

func (s RecordPostBackRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordPostBackRequest) GoString() string {
	return s.String()
}

func (s *RecordPostBackRequest) SetBizId(v string) *RecordPostBackRequest {
	s.BizId = &v
	return s
}

func (s *RecordPostBackRequest) SetBizType(v string) *RecordPostBackRequest {
	s.BizType = &v
	return s
}

func (s *RecordPostBackRequest) SetContactor(v string) *RecordPostBackRequest {
	s.Contactor = &v
	return s
}

func (s *RecordPostBackRequest) SetContent(v string) *RecordPostBackRequest {
	s.Content = &v
	return s
}

func (s *RecordPostBackRequest) SetEntityKey(v string) *RecordPostBackRequest {
	s.EntityKey = &v
	return s
}

type RecordPostBackResponseBody struct {
	AllowRetry     *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName        *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	ErrorCode      *string       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *bool         `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId      *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecordPostBackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecordPostBackResponseBody) GoString() string {
	return s.String()
}

func (s *RecordPostBackResponseBody) SetAllowRetry(v bool) *RecordPostBackResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RecordPostBackResponseBody) SetAppName(v string) *RecordPostBackResponseBody {
	s.AppName = &v
	return s
}

func (s *RecordPostBackResponseBody) SetDynamicCode(v string) *RecordPostBackResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetDynamicMessage(v string) *RecordPostBackResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorArgs(v []interface{}) *RecordPostBackResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorCode(v string) *RecordPostBackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorMsg(v string) *RecordPostBackResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RecordPostBackResponseBody) SetHttpStatusCode(v int32) *RecordPostBackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetModule(v bool) *RecordPostBackResponseBody {
	s.Module = &v
	return s
}

func (s *RecordPostBackResponseBody) SetRequestId(v string) *RecordPostBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecordPostBackResponseBody) SetSuccess(v bool) *RecordPostBackResponseBody {
	s.Success = &v
	return s
}

type RecordPostBackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecordPostBackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecordPostBackResponse) String() string {
	return tea.Prettify(s)
}

func (s RecordPostBackResponse) GoString() string {
	return s.String()
}

func (s *RecordPostBackResponse) SetHeaders(v map[string]*string) *RecordPostBackResponse {
	s.Headers = v
	return s
}

func (s *RecordPostBackResponse) SetStatusCode(v int32) *RecordPostBackResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordPostBackResponse) SetBody(v *RecordPostBackResponseBody) *RecordPostBackResponse {
	s.Body = v
	return s
}

type RejectSolutionRequest struct {
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
}

func (s RejectSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectSolutionRequest) GoString() string {
	return s.String()
}

func (s *RejectSolutionRequest) SetNote(v string) *RejectSolutionRequest {
	s.Note = &v
	return s
}

func (s *RejectSolutionRequest) SetSolutionBizId(v string) *RejectSolutionRequest {
	s.SolutionBizId = &v
	return s
}

type RejectSolutionResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RejectSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectSolutionResponseBody) SetErrorCode(v string) *RejectSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RejectSolutionResponseBody) SetErrorMsg(v string) *RejectSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RejectSolutionResponseBody) SetRequestId(v string) *RejectSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectSolutionResponseBody) SetSuccess(v bool) *RejectSolutionResponseBody {
	s.Success = &v
	return s
}

type RejectSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RejectSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectSolutionResponse) GoString() string {
	return s.String()
}

func (s *RejectSolutionResponse) SetHeaders(v map[string]*string) *RejectSolutionResponse {
	s.Headers = v
	return s
}

func (s *RejectSolutionResponse) SetStatusCode(v int32) *RejectSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectSolutionResponse) SetBody(v *RejectSolutionResponseBody) *RejectSolutionResponse {
	s.Body = v
	return s
}

type RejectUserSolutionRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
}

func (s RejectUserSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectUserSolutionRequest) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionRequest) SetBizType(v string) *RejectUserSolutionRequest {
	s.BizType = &v
	return s
}

func (s *RejectUserSolutionRequest) SetNote(v string) *RejectUserSolutionRequest {
	s.Note = &v
	return s
}

func (s *RejectUserSolutionRequest) SetSolutionBizId(v string) *RejectUserSolutionRequest {
	s.SolutionBizId = &v
	return s
}

type RejectUserSolutionResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RejectUserSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectUserSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionResponseBody) SetErrorCode(v string) *RejectUserSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetErrorMsg(v string) *RejectUserSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetRequestId(v string) *RejectUserSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectUserSolutionResponseBody) SetSuccess(v bool) *RejectUserSolutionResponseBody {
	s.Success = &v
	return s
}

type RejectUserSolutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RejectUserSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectUserSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectUserSolutionResponse) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionResponse) SetHeaders(v map[string]*string) *RejectUserSolutionResponse {
	s.Headers = v
	return s
}

func (s *RejectUserSolutionResponse) SetStatusCode(v int32) *RejectUserSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectUserSolutionResponse) SetBody(v *RejectUserSolutionResponseBody) *RejectUserSolutionResponse {
	s.Body = v
	return s
}

type ReleaseProduceAuthorizationRequest struct {
	AuthorizedUserId *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s ReleaseProduceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationRequest) SetAuthorizedUserId(v string) *ReleaseProduceAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *ReleaseProduceAuthorizationRequest) SetBizId(v string) *ReleaseProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *ReleaseProduceAuthorizationRequest) SetBizType(v string) *ReleaseProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

type ReleaseProduceAuthorizationResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseProduceAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationResponseBody) SetErrorCode(v string) *ReleaseProduceAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetErrorMsg(v string) *ReleaseProduceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetRequestId(v string) *ReleaseProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponseBody) SetSuccess(v bool) *ReleaseProduceAuthorizationResponseBody {
	s.Success = &v
	return s
}

type ReleaseProduceAuthorizationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseProduceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseProduceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseProduceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ReleaseProduceAuthorizationResponse) SetHeaders(v map[string]*string) *ReleaseProduceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ReleaseProduceAuthorizationResponse) SetStatusCode(v int32) *ReleaseProduceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseProduceAuthorizationResponse) SetBody(v *ReleaseProduceAuthorizationResponseBody) *ReleaseProduceAuthorizationResponse {
	s.Body = v
	return s
}

type StartBackToBackCallRequest struct {
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallCenterNumber *string `json:"CallCenterNumber,omitempty" xml:"CallCenterNumber,omitempty"`
	Caller           *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	MobileKey        *string `json:"MobileKey,omitempty" xml:"MobileKey,omitempty"`
	SkillType        *int64  `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
}

func (s StartBackToBackCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBackToBackCallRequest) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallRequest) SetBizId(v string) *StartBackToBackCallRequest {
	s.BizId = &v
	return s
}

func (s *StartBackToBackCallRequest) SetBizType(v string) *StartBackToBackCallRequest {
	s.BizType = &v
	return s
}

func (s *StartBackToBackCallRequest) SetCallCenterNumber(v string) *StartBackToBackCallRequest {
	s.CallCenterNumber = &v
	return s
}

func (s *StartBackToBackCallRequest) SetCaller(v string) *StartBackToBackCallRequest {
	s.Caller = &v
	return s
}

func (s *StartBackToBackCallRequest) SetMobileKey(v string) *StartBackToBackCallRequest {
	s.MobileKey = &v
	return s
}

func (s *StartBackToBackCallRequest) SetSkillType(v int64) *StartBackToBackCallRequest {
	s.SkillType = &v
	return s
}

type StartBackToBackCallResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartBackToBackCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBackToBackCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallResponseBody) SetData(v bool) *StartBackToBackCallResponseBody {
	s.Data = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetErrorCode(v string) *StartBackToBackCallResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetErrorMsg(v string) *StartBackToBackCallResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetRequestId(v string) *StartBackToBackCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetSuccess(v bool) *StartBackToBackCallResponseBody {
	s.Success = &v
	return s
}

type StartBackToBackCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartBackToBackCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBackToBackCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBackToBackCallResponse) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallResponse) SetHeaders(v map[string]*string) *StartBackToBackCallResponse {
	s.Headers = v
	return s
}

func (s *StartBackToBackCallResponse) SetStatusCode(v int32) *StartBackToBackCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBackToBackCallResponse) SetBody(v *StartBackToBackCallResponseBody) *StartBackToBackCallResponse {
	s.Body = v
	return s
}

type SubmitIntentionNoteRequest struct {
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s SubmitIntentionNoteRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIntentionNoteRequest) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteRequest) SetIntentionBizId(v string) *SubmitIntentionNoteRequest {
	s.IntentionBizId = &v
	return s
}

func (s *SubmitIntentionNoteRequest) SetNote(v string) *SubmitIntentionNoteRequest {
	s.Note = &v
	return s
}

type SubmitIntentionNoteResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIntentionNoteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIntentionNoteResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteResponseBody) SetErrorCode(v string) *SubmitIntentionNoteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetErrorMsg(v string) *SubmitIntentionNoteResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetRequestId(v string) *SubmitIntentionNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetSuccess(v bool) *SubmitIntentionNoteResponseBody {
	s.Success = &v
	return s
}

type SubmitIntentionNoteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitIntentionNoteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitIntentionNoteResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIntentionNoteResponse) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteResponse) SetHeaders(v map[string]*string) *SubmitIntentionNoteResponse {
	s.Headers = v
	return s
}

func (s *SubmitIntentionNoteResponse) SetStatusCode(v int32) *SubmitIntentionNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIntentionNoteResponse) SetBody(v *SubmitIntentionNoteResponseBody) *SubmitIntentionNoteResponse {
	s.Body = v
	return s
}

type SubmitSolutionRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Solution       *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SubmitSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSolutionRequest) GoString() string {
	return s.String()
}

func (s *SubmitSolutionRequest) SetBizType(v string) *SubmitSolutionRequest {
	s.BizType = &v
	return s
}

func (s *SubmitSolutionRequest) SetIntentionBizId(v string) *SubmitSolutionRequest {
	s.IntentionBizId = &v
	return s
}

func (s *SubmitSolutionRequest) SetSolution(v string) *SubmitSolutionRequest {
	s.Solution = &v
	return s
}

func (s *SubmitSolutionRequest) SetUserId(v string) *SubmitSolutionRequest {
	s.UserId = &v
	return s
}

type SubmitSolutionResponseBody struct {
	ConfirmUrl    *string `json:"ConfirmUrl,omitempty" xml:"ConfirmUrl,omitempty"`
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg      *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SolutionBizId *string `json:"SolutionBizId,omitempty" xml:"SolutionBizId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSolutionResponseBody) SetConfirmUrl(v string) *SubmitSolutionResponseBody {
	s.ConfirmUrl = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetErrorCode(v string) *SubmitSolutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetErrorMsg(v string) *SubmitSolutionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetRequestId(v string) *SubmitSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetSolutionBizId(v string) *SubmitSolutionResponseBody {
	s.SolutionBizId = &v
	return s
}

func (s *SubmitSolutionResponseBody) SetSuccess(v bool) *SubmitSolutionResponseBody {
	s.Success = &v
	return s
}

type SubmitSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSolutionResponse) GoString() string {
	return s.String()
}

func (s *SubmitSolutionResponse) SetHeaders(v map[string]*string) *SubmitSolutionResponse {
	s.Headers = v
	return s
}

func (s *SubmitSolutionResponse) SetStatusCode(v int32) *SubmitSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSolutionResponse) SetBody(v *SubmitSolutionResponseBody) *SubmitSolutionResponse {
	s.Body = v
	return s
}

type TransferIntentionOwnerRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PersonId *int32  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s TransferIntentionOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferIntentionOwnerRequest) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerRequest) SetBizId(v string) *TransferIntentionOwnerRequest {
	s.BizId = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetBizType(v string) *TransferIntentionOwnerRequest {
	s.BizType = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetPersonId(v int32) *TransferIntentionOwnerRequest {
	s.PersonId = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetRemark(v string) *TransferIntentionOwnerRequest {
	s.Remark = &v
	return s
}

type TransferIntentionOwnerResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferIntentionOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferIntentionOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerResponseBody) SetData(v bool) *TransferIntentionOwnerResponseBody {
	s.Data = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetErrorCode(v string) *TransferIntentionOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetErrorMsg(v string) *TransferIntentionOwnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetRequestId(v string) *TransferIntentionOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetSuccess(v bool) *TransferIntentionOwnerResponseBody {
	s.Success = &v
	return s
}

type TransferIntentionOwnerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferIntentionOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferIntentionOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferIntentionOwnerResponse) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerResponse) SetHeaders(v map[string]*string) *TransferIntentionOwnerResponse {
	s.Headers = v
	return s
}

func (s *TransferIntentionOwnerResponse) SetStatusCode(v int32) *TransferIntentionOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferIntentionOwnerResponse) SetBody(v *TransferIntentionOwnerResponseBody) *TransferIntentionOwnerResponse {
	s.Body = v
	return s
}

type TransferProduceOwnerRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PersonId *int32  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s TransferProduceOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferProduceOwnerRequest) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerRequest) SetBizId(v string) *TransferProduceOwnerRequest {
	s.BizId = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetBizType(v string) *TransferProduceOwnerRequest {
	s.BizType = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetPersonId(v int32) *TransferProduceOwnerRequest {
	s.PersonId = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetRemark(v string) *TransferProduceOwnerRequest {
	s.Remark = &v
	return s
}

type TransferProduceOwnerResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferProduceOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferProduceOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerResponseBody) SetData(v bool) *TransferProduceOwnerResponseBody {
	s.Data = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetErrorCode(v string) *TransferProduceOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetErrorMsg(v string) *TransferProduceOwnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetRequestId(v string) *TransferProduceOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetSuccess(v bool) *TransferProduceOwnerResponseBody {
	s.Success = &v
	return s
}

type TransferProduceOwnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferProduceOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferProduceOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferProduceOwnerResponse) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerResponse) SetHeaders(v map[string]*string) *TransferProduceOwnerResponse {
	s.Headers = v
	return s
}

func (s *TransferProduceOwnerResponse) SetStatusCode(v int32) *TransferProduceOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferProduceOwnerResponse) SetBody(v *TransferProduceOwnerResponseBody) *TransferProduceOwnerResponse {
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("companyreg.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("companyreg.aliyuncs.com"),
		"ap-south-1":                  tea.String("companyreg.aliyuncs.com"),
		"ap-southeast-1":              tea.String("companyreg.aliyuncs.com"),
		"ap-southeast-2":              tea.String("companyreg.aliyuncs.com"),
		"ap-southeast-3":              tea.String("companyreg.aliyuncs.com"),
		"ap-southeast-5":              tea.String("companyreg.aliyuncs.com"),
		"cn-beijing":                  tea.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("companyreg.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("companyreg.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("companyreg.aliyuncs.com"),
		"cn-chengdu":                  tea.String("companyreg.aliyuncs.com"),
		"cn-edge-1":                   tea.String("companyreg.aliyuncs.com"),
		"cn-fujian":                   tea.String("companyreg.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("companyreg.aliyuncs.com"),
		"cn-hongkong":                 tea.String("companyreg.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("companyreg.aliyuncs.com"),
		"cn-huhehaote":                tea.String("companyreg.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("companyreg.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("companyreg.aliyuncs.com"),
		"cn-qingdao":                  tea.String("companyreg.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai":                 tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("companyreg.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("companyreg.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("companyreg.aliyuncs.com"),
		"cn-wuhan":                    tea.String("companyreg.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("companyreg.aliyuncs.com"),
		"cn-yushanfang":               tea.String("companyreg.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("companyreg.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("companyreg.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("companyreg.aliyuncs.com"),
		"eu-central-1":                tea.String("companyreg.aliyuncs.com"),
		"eu-west-1":                   tea.String("companyreg.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("companyreg.aliyuncs.com"),
		"me-east-1":                   tea.String("companyreg.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("companyreg.aliyuncs.com"),
		"us-east-1":                   tea.String("companyreg.aliyuncs.com"),
		"us-west-1":                   tea.String("companyreg.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("companyreg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindProduceAuthorizationWithOptions(request *BindProduceAuthorizationRequest, runtime *util.RuntimeOptions) (_result *BindProduceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserIds)) {
		body["AuthorizedUserIds"] = request.AuthorizedUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindProduceAuthorization"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindProduceAuthorization(request *BindProduceAuthorizationRequest) (_result *BindProduceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindProduceAuthorizationResponse{}
	_body, _err := client.BindProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseIntentionForPartnerWithOptions(request *CloseIntentionForPartnerRequest, runtime *util.RuntimeOptions) (_result *CloseIntentionForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseIntentionForPartner"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseIntentionForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseIntentionForPartner(request *CloseIntentionForPartnerRequest) (_result *CloseIntentionForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseIntentionForPartnerResponse{}
	_body, _err := client.CloseIntentionForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseUserIntentionWithOptions(request *CloseUserIntentionRequest, runtime *util.RuntimeOptions) (_result *CloseUserIntentionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseUserIntention"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseUserIntentionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseUserIntention(request *CloseUserIntentionRequest) (_result *CloseUserIntentionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseUserIntentionResponse{}
	_body, _err := client.CloseUserIntentionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBusinessOpportunityWithOptions(request *CreateBusinessOpportunityRequest, runtime *util.RuntimeOptions) (_result *CreateBusinessOpportunityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.VCode)) {
		query["VCode"] = request.VCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBusinessOpportunity"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBusinessOpportunityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBusinessOpportunity(request *CreateBusinessOpportunityRequest) (_result *CreateBusinessOpportunityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBusinessOpportunityResponse{}
	_body, _err := client.CreateBusinessOpportunityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProduceForPartnerWithOptions(request *CreateProduceForPartnerRequest, runtime *util.RuntimeOptions) (_result *CreateProduceForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		query["ExtInfo"] = request.ExtInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduceForPartner"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProduceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProduceForPartner(request *CreateProduceForPartnerRequest) (_result *CreateProduceForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProduceForPartnerResponse{}
	_body, _err := client.CreateProduceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePartnerConfigWithOptions(request *DescribePartnerConfigRequest, runtime *util.RuntimeOptions) (_result *DescribePartnerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerCode)) {
		query["PartnerCode"] = request.PartnerCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePartnerConfig"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePartnerConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePartnerConfig(request *DescribePartnerConfigRequest) (_result *DescribePartnerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePartnerConfigResponse{}
	_body, _err := client.DescribePartnerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicyWithOptions(request *GenerateUploadFilePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadFilePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadFilePolicy"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicy(request *GenerateUploadFilePolicyRequest) (_result *GenerateUploadFilePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.GenerateUploadFilePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlipayUrlWithOptions(request *GetAlipayUrlRequest, runtime *util.RuntimeOptions) (_result *GetAlipayUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlipayUrl"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlipayUrl(request *GetAlipayUrlRequest) (_result *GetAlipayUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.GetAlipayUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntentionNoteWithOptions(request *ListIntentionNoteRequest, runtime *util.RuntimeOptions) (_result *ListIntentionNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntentionNote"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntentionNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntentionNote(request *ListIntentionNoteRequest) (_result *ListIntentionNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntentionNoteResponse{}
	_body, _err := client.ListIntentionNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProduceAuthorizationWithOptions(request *ListProduceAuthorizationRequest, runtime *util.RuntimeOptions) (_result *ListProduceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProduceAuthorization"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProduceAuthorization(request *ListProduceAuthorizationRequest) (_result *ListProduceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProduceAuthorizationResponse{}
	_body, _err := client.ListProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserDetailSolutionsWithOptions(request *ListUserDetailSolutionsRequest, runtime *util.RuntimeOptions) (_result *ListUserDetailSolutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserDetailSolutions"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserDetailSolutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserDetailSolutions(request *ListUserDetailSolutionsRequest) (_result *ListUserDetailSolutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDetailSolutionsResponse{}
	_body, _err := client.ListUserDetailSolutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserIntentionNotesWithOptions(request *ListUserIntentionNotesRequest, runtime *util.RuntimeOptions) (_result *ListUserIntentionNotesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserIntentionNotes"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserIntentionNotesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserIntentionNotes(request *ListUserIntentionNotesRequest) (_result *ListUserIntentionNotesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserIntentionNotesResponse{}
	_body, _err := client.ListUserIntentionNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserIntentionsWithOptions(request *ListUserIntentionsRequest, runtime *util.RuntimeOptions) (_result *ListUserIntentionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortFiled)) {
		query["SortFiled"] = request.SortFiled
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.WithExtInfo)) {
		query["WithExtInfo"] = request.WithExtInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserIntentions"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserIntentionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserIntentions(request *ListUserIntentionsRequest) (_result *ListUserIntentionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserIntentionsResponse{}
	_body, _err := client.ListUserIntentionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserProduceOperateLogsWithOptions(request *ListUserProduceOperateLogsRequest, runtime *util.RuntimeOptions) (_result *ListUserProduceOperateLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserProduceOperateLogs"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserProduceOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserProduceOperateLogs(request *ListUserProduceOperateLogsRequest) (_result *ListUserProduceOperateLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserProduceOperateLogsResponse{}
	_body, _err := client.ListUserProduceOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserSolutionsWithOptions(tmpReq *ListUserSolutionsRequest, runtime *util.RuntimeOptions) (_result *ListUserSolutionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListUserSolutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExistStatus)) {
		request.ExistStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExistStatus, tea.String("ExistStatus"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExistStatusShrink)) {
		query["ExistStatus"] = request.ExistStatusShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserSolutions"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserSolutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserSolutions(request *ListUserSolutionsRequest) (_result *ListUserSolutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserSolutionsResponse{}
	_body, _err := client.ListUserSolutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateProduceForPartnerWithOptions(request *OperateProduceForPartnerRequest, runtime *util.RuntimeOptions) (_result *OperateProduceForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateProduceForPartner"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateProduceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateProduceForPartner(request *OperateProduceForPartnerRequest) (_result *OperateProduceForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateProduceForPartnerResponse{}
	_body, _err := client.OperateProduceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureDataWithOptions(request *PutMeasureDataRequest, runtime *util.RuntimeOptions) (_result *PutMeasureDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureData"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureData(request *PutMeasureDataRequest) (_result *PutMeasureDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.PutMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureReadyFlagWithOptions(request *PutMeasureReadyFlagRequest, runtime *util.RuntimeOptions) (_result *PutMeasureReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReadyFlag)) {
		query["ReadyFlag"] = request.ReadyFlag
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureReadyFlag"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureReadyFlag(request *PutMeasureReadyFlagRequest) (_result *PutMeasureReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.PutMeasureReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAvailableNumbersWithOptions(request *QueryAvailableNumbersRequest, runtime *util.RuntimeOptions) (_result *QueryAvailableNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvailableNumbers"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvailableNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAvailableNumbers(request *QueryAvailableNumbersRequest) (_result *QueryAvailableNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAvailableNumbersResponse{}
	_body, _err := client.QueryAvailableNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBagRemainingWithOptions(request *QueryBagRemainingRequest, runtime *util.RuntimeOptions) (_result *QueryBagRemainingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBagRemaining"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBagRemainingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBagRemaining(request *QueryBagRemainingRequest) (_result *QueryBagRemainingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBagRemainingResponse{}
	_body, _err := client.QueryBagRemainingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCommodityConfigWithOptions(request *QueryCommodityConfigRequest, runtime *util.RuntimeOptions) (_result *QueryCommodityConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.QueryModule)) {
		query["QueryModule"] = request.QueryModule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCommodityConfig"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCommodityConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCommodityConfig(request *QueryCommodityConfigRequest) (_result *QueryCommodityConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCommodityConfigResponse{}
	_body, _err := client.QueryCommodityConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceWithOptions(request *QueryInstanceRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstance"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstance(request *QueryInstanceRequest) (_result *QueryInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceResponse{}
	_body, _err := client.QueryInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPartnerIntentionListWithOptions(request *QueryPartnerIntentionListRequest, runtime *util.RuntimeOptions) (_result *QueryPartnerIntentionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPartnerIntentionList"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPartnerIntentionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPartnerIntentionList(request *QueryPartnerIntentionListRequest) (_result *QueryPartnerIntentionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPartnerIntentionListResponse{}
	_body, _err := client.QueryPartnerIntentionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPartnerProduceListWithOptions(request *QueryPartnerProduceListRequest, runtime *util.RuntimeOptions) (_result *QueryPartnerProduceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPartnerProduceList"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPartnerProduceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPartnerProduceList(request *QueryPartnerProduceListRequest) (_result *QueryPartnerProduceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPartnerProduceListResponse{}
	_body, _err := client.QueryPartnerProduceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserNeedAuthWithOptions(runtime *util.RuntimeOptions) (_result *QueryUserNeedAuthResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryUserNeedAuth"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserNeedAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserNeedAuth() (_result *QueryUserNeedAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserNeedAuthResponse{}
	_body, _err := client.QueryUserNeedAuthWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecordPostBackWithOptions(request *RecordPostBackRequest, runtime *util.RuntimeOptions) (_result *RecordPostBackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Contactor)) {
		query["contactor"] = request.Contactor
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.EntityKey)) {
		query["entityKey"] = request.EntityKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecordPostBack"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecordPostBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecordPostBack(request *RecordPostBackRequest) (_result *RecordPostBackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecordPostBackResponse{}
	_body, _err := client.RecordPostBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectSolutionWithOptions(request *RejectSolutionRequest, runtime *util.RuntimeOptions) (_result *RejectSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionBizId)) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectSolution"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectSolution(request *RejectSolutionRequest) (_result *RejectSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectSolutionResponse{}
	_body, _err := client.RejectSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectUserSolutionWithOptions(request *RejectUserSolutionRequest, runtime *util.RuntimeOptions) (_result *RejectUserSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionBizId)) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectUserSolution"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectUserSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectUserSolution(request *RejectUserSolutionRequest) (_result *RejectUserSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectUserSolutionResponse{}
	_body, _err := client.RejectUserSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseProduceAuthorizationWithOptions(request *ReleaseProduceAuthorizationRequest, runtime *util.RuntimeOptions) (_result *ReleaseProduceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		body["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseProduceAuthorization"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseProduceAuthorization(request *ReleaseProduceAuthorizationRequest) (_result *ReleaseProduceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseProduceAuthorizationResponse{}
	_body, _err := client.ReleaseProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBackToBackCallWithOptions(request *StartBackToBackCallRequest, runtime *util.RuntimeOptions) (_result *StartBackToBackCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallCenterNumber)) {
		query["CallCenterNumber"] = request.CallCenterNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.MobileKey)) {
		query["MobileKey"] = request.MobileKey
	}

	if !tea.BoolValue(util.IsUnset(request.SkillType)) {
		query["SkillType"] = request.SkillType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartBackToBackCall"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartBackToBackCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBackToBackCall(request *StartBackToBackCallRequest) (_result *StartBackToBackCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBackToBackCallResponse{}
	_body, _err := client.StartBackToBackCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitIntentionNoteWithOptions(request *SubmitIntentionNoteRequest, runtime *util.RuntimeOptions) (_result *SubmitIntentionNoteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitIntentionNote"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitIntentionNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitIntentionNote(request *SubmitIntentionNoteRequest) (_result *SubmitIntentionNoteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitIntentionNoteResponse{}
	_body, _err := client.SubmitIntentionNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSolutionWithOptions(request *SubmitSolutionRequest, runtime *util.RuntimeOptions) (_result *SubmitSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionBizId)) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Solution)) {
		query["Solution"] = request.Solution
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSolution"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSolution(request *SubmitSolutionRequest) (_result *SubmitSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSolutionResponse{}
	_body, _err := client.SubmitSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferIntentionOwnerWithOptions(request *TransferIntentionOwnerRequest, runtime *util.RuntimeOptions) (_result *TransferIntentionOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PersonId)) {
		query["PersonId"] = request.PersonId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferIntentionOwner"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferIntentionOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferIntentionOwner(request *TransferIntentionOwnerRequest) (_result *TransferIntentionOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferIntentionOwnerResponse{}
	_body, _err := client.TransferIntentionOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferProduceOwnerWithOptions(request *TransferProduceOwnerRequest, runtime *util.RuntimeOptions) (_result *TransferProduceOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.PersonId)) {
		query["PersonId"] = request.PersonId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferProduceOwner"),
		Version:     tea.String("2020-03-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferProduceOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferProduceOwner(request *TransferProduceOwnerRequest) (_result *TransferProduceOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferProduceOwnerResponse{}
	_body, _err := client.TransferProduceOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
