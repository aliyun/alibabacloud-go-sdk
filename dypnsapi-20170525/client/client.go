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

type CreateVerifySchemeRequest struct {
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BundleId             *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PackName             *string `json:"PackName,omitempty" xml:"PackName,omitempty"`
	PackSign             *string `json:"PackSign,omitempty" xml:"PackSign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeName           *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
}

func (s CreateVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeRequest) SetAppName(v string) *CreateVerifySchemeRequest {
	s.AppName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetBundleId(v string) *CreateVerifySchemeRequest {
	s.BundleId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOsType(v string) *CreateVerifySchemeRequest {
	s.OsType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOwnerId(v int64) *CreateVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackName(v string) *CreateVerifySchemeRequest {
	s.PackName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackSign(v string) *CreateVerifySchemeRequest {
	s.PackSign = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerAccount(v string) *CreateVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerId(v int64) *CreateVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSchemeName(v string) *CreateVerifySchemeRequest {
	s.SchemeName = &v
	return s
}

type CreateVerifySchemeResponseBody struct {
	Code                *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	GateVerifySchemeDTO *CreateVerifySchemeResponseBodyGateVerifySchemeDTO `json:"GateVerifySchemeDTO,omitempty" xml:"GateVerifySchemeDTO,omitempty" type:"Struct"`
	Message             *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBody) SetCode(v string) *CreateVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetGateVerifySchemeDTO(v *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) *CreateVerifySchemeResponseBody {
	s.GateVerifySchemeDTO = v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetMessage(v string) *CreateVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetRequestId(v string) *CreateVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

type CreateVerifySchemeResponseBodyGateVerifySchemeDTO struct {
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) SetSchemeCode(v string) *CreateVerifySchemeResponseBodyGateVerifySchemeDTO {
	s.SchemeCode = &v
	return s
}

type CreateVerifySchemeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponse) SetHeaders(v map[string]*string) *CreateVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySchemeResponse) SetBody(v *CreateVerifySchemeResponseBody) *CreateVerifySchemeResponse {
	s.Body = v
	return s
}

type DeleteVerifySchemeRequest struct {
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeCode           *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DeleteVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeRequest) SetCustomerId(v int64) *DeleteVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerAccount(v string) *DeleteVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetSchemeCode(v string) *DeleteVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

type DeleteVerifySchemeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponseBody) SetCode(v string) *DeleteVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetMessage(v string) *DeleteVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetRequestId(v string) *DeleteVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetResult(v bool) *DeleteVerifySchemeResponseBody {
	s.Result = &v
	return s
}

type DeleteVerifySchemeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponse) SetHeaders(v map[string]*string) *DeleteVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifySchemeResponse) SetBody(v *DeleteVerifySchemeResponseBody) *DeleteVerifySchemeResponse {
	s.Body = v
	return s
}

type DescribeVerifySchemeRequest struct {
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeCode           *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DescribeVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeRequest) SetCustomerId(v int64) *DescribeVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerAccount(v string) *DescribeVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetSchemeCode(v string) *DescribeVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

type DescribeVerifySchemeResponseBody struct {
	Code                 *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchemeQueryResultDTO *DescribeVerifySchemeResponseBodySchemeQueryResultDTO `json:"SchemeQueryResultDTO,omitempty" xml:"SchemeQueryResultDTO,omitempty" type:"Struct"`
}

func (s DescribeVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBody) SetCode(v string) *DescribeVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetMessage(v string) *DescribeVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetRequestId(v string) *DescribeVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetSchemeQueryResultDTO(v *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) *DescribeVerifySchemeResponseBody {
	s.SchemeQueryResultDTO = v
	return s
}

type DescribeVerifySchemeResponseBodySchemeQueryResultDTO struct {
	AppEncryptInfo *string `json:"AppEncryptInfo,omitempty" xml:"AppEncryptInfo,omitempty"`
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) SetAppEncryptInfo(v string) *DescribeVerifySchemeResponseBodySchemeQueryResultDTO {
	s.AppEncryptInfo = &v
	return s
}

type DescribeVerifySchemeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponse) SetHeaders(v map[string]*string) *DescribeVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySchemeResponse) SetBody(v *DescribeVerifySchemeResponseBody) *DescribeVerifySchemeResponse {
	s.Body = v
	return s
}

type GetAuthTokenRequest struct {
	Origin               *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetOrigin(v string) *GetAuthTokenRequest {
	s.Origin = &v
	return s
}

func (s *GetAuthTokenRequest) SetOwnerId(v int64) *GetAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerAccount(v string) *GetAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerId(v int64) *GetAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetUrl(v string) *GetAuthTokenRequest {
	s.Url = &v
	return s
}

type GetAuthTokenResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenInfo *GetAuthTokenResponseBodyTokenInfo `json:"TokenInfo,omitempty" xml:"TokenInfo,omitempty" type:"Struct"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetCode(v string) *GetAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetMessage(v string) *GetAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetTokenInfo(v *GetAuthTokenResponseBodyTokenInfo) *GetAuthTokenResponseBody {
	s.TokenInfo = v
	return s
}

type GetAuthTokenResponseBodyTokenInfo struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	JwtToken    *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetAuthTokenResponseBodyTokenInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyTokenInfo) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetAccessToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.AccessToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetJwtToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.JwtToken = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type GetAuthorizationUrlRequest struct {
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeId             *int64  `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s GetAuthorizationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlRequest) SetEndDate(v string) *GetAuthorizationUrlRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetPhoneNo(v string) *GetAuthorizationUrlRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerAccount(v string) *GetAuthorizationUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetSchemeId(v int64) *GetAuthorizationUrlRequest {
	s.SchemeId = &v
	return s
}

type GetAuthorizationUrlResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAuthorizationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBody) SetCode(v string) *GetAuthorizationUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetData(v *GetAuthorizationUrlResponseBodyData) *GetAuthorizationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetMessage(v string) *GetAuthorizationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetRequestId(v string) *GetAuthorizationUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthorizationUrlResponseBodyData struct {
	AuthorizationUrl *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
}

func (s GetAuthorizationUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBodyData) SetAuthorizationUrl(v string) *GetAuthorizationUrlResponseBodyData {
	s.AuthorizationUrl = &v
	return s
}

type GetAuthorizationUrlResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthorizationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthorizationUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponse) SetHeaders(v map[string]*string) *GetAuthorizationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationUrlResponse) SetBody(v *GetAuthorizationUrlResponseBody) *GetAuthorizationUrlResponse {
	s.Body = v
	return s
}

type GetCertifyResultRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Token                *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetCertifyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultRequest) GoString() string {
	return s.String()
}

func (s *GetCertifyResultRequest) SetOwnerId(v int64) *GetCertifyResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCertifyResultRequest) SetResourceOwnerAccount(v string) *GetCertifyResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCertifyResultRequest) SetResourceOwnerId(v int64) *GetCertifyResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCertifyResultRequest) SetToken(v string) *GetCertifyResultRequest {
	s.Token = &v
	return s
}

type GetCertifyResultResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetCertifyResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCertifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertifyResultResponseBody) SetCode(v string) *GetCertifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetCertifyResultResponseBody) SetData(v []*GetCertifyResultResponseBodyData) *GetCertifyResultResponseBody {
	s.Data = v
	return s
}

func (s *GetCertifyResultResponseBody) SetMessage(v string) *GetCertifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetCertifyResultResponseBody) SetRequestId(v string) *GetCertifyResultResponseBody {
	s.RequestId = &v
	return s
}

type GetCertifyResultResponseBodyData struct {
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	VerifyDesc   *string `json:"VerifyDesc,omitempty" xml:"VerifyDesc,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s GetCertifyResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCertifyResultResponseBodyData) SetDeviceToken(v string) *GetCertifyResultResponseBodyData {
	s.DeviceToken = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetIdentityInfo(v string) *GetCertifyResultResponseBodyData {
	s.IdentityInfo = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetMaterialInfo(v string) *GetCertifyResultResponseBodyData {
	s.MaterialInfo = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetVerifyDesc(v string) *GetCertifyResultResponseBodyData {
	s.VerifyDesc = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetVerifyResult(v string) *GetCertifyResultResponseBodyData {
	s.VerifyResult = &v
	return s
}

type GetCertifyResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCertifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCertifyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultResponse) GoString() string {
	return s.String()
}

func (s *GetCertifyResultResponse) SetHeaders(v map[string]*string) *GetCertifyResultResponse {
	s.Headers = v
	return s
}

func (s *GetCertifyResultResponse) SetBody(v *GetCertifyResultResponseBody) *GetCertifyResultResponse {
	s.Body = v
	return s
}

type GetMobileRequest struct {
	AccessToken          *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMobileRequest) GoString() string {
	return s.String()
}

func (s *GetMobileRequest) SetAccessToken(v string) *GetMobileRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMobileRequest) SetOutId(v string) *GetMobileRequest {
	s.OutId = &v
	return s
}

func (s *GetMobileRequest) SetOwnerId(v int64) *GetMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerAccount(v string) *GetMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerId(v int64) *GetMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMobileResponseBody struct {
	Code               *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	GetMobileResultDTO *GetMobileResponseBodyGetMobileResultDTO `json:"GetMobileResultDTO,omitempty" xml:"GetMobileResultDTO,omitempty" type:"Struct"`
	Message            *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponseBody) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBody) SetCode(v string) *GetMobileResponseBody {
	s.Code = &v
	return s
}

func (s *GetMobileResponseBody) SetGetMobileResultDTO(v *GetMobileResponseBodyGetMobileResultDTO) *GetMobileResponseBody {
	s.GetMobileResultDTO = v
	return s
}

func (s *GetMobileResponseBody) SetMessage(v string) *GetMobileResponseBody {
	s.Message = &v
	return s
}

func (s *GetMobileResponseBody) SetRequestId(v string) *GetMobileResponseBody {
	s.RequestId = &v
	return s
}

type GetMobileResponseBodyGetMobileResultDTO struct {
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetMobileResponseBodyGetMobileResultDTO) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponseBodyGetMobileResultDTO) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBodyGetMobileResultDTO) SetMobile(v string) *GetMobileResponseBodyGetMobileResultDTO {
	s.Mobile = &v
	return s
}

type GetMobileResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMobileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponse) GoString() string {
	return s.String()
}

func (s *GetMobileResponse) SetHeaders(v map[string]*string) *GetMobileResponse {
	s.Headers = v
	return s
}

func (s *GetMobileResponse) SetBody(v *GetMobileResponseBody) *GetMobileResponse {
	s.Body = v
	return s
}

type GetSmsAuthTokensRequest struct {
	BundleId             *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	Expire               *int64  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PackageName          *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SceneCode            *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsCodeExpire        *int32  `json:"SmsCodeExpire,omitempty" xml:"SmsCodeExpire,omitempty"`
	SmsTemplateCode      *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
}

func (s GetSmsAuthTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensRequest) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensRequest) SetBundleId(v string) *GetSmsAuthTokensRequest {
	s.BundleId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetExpire(v int64) *GetSmsAuthTokensRequest {
	s.Expire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOsType(v string) *GetSmsAuthTokensRequest {
	s.OsType = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetPackageName(v string) *GetSmsAuthTokensRequest {
	s.PackageName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerAccount(v string) *GetSmsAuthTokensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSceneCode(v string) *GetSmsAuthTokensRequest {
	s.SceneCode = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSignName(v string) *GetSmsAuthTokensRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsCodeExpire(v int32) *GetSmsAuthTokensRequest {
	s.SmsCodeExpire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsTemplateCode(v string) *GetSmsAuthTokensRequest {
	s.SmsTemplateCode = &v
	return s
}

type GetSmsAuthTokensResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSmsAuthTokensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSmsAuthTokensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBody) SetCode(v string) *GetSmsAuthTokensResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetData(v *GetSmsAuthTokensResponseBodyData) *GetSmsAuthTokensResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetMessage(v string) *GetSmsAuthTokensResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetRequestId(v string) *GetSmsAuthTokensResponseBody {
	s.RequestId = &v
	return s
}

type GetSmsAuthTokensResponseBodyData struct {
	BizToken           *string `json:"BizToken,omitempty" xml:"BizToken,omitempty"`
	ExpireTime         *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	StsAccessKeyId     *string `json:"StsAccessKeyId,omitempty" xml:"StsAccessKeyId,omitempty"`
	StsAccessKeySecret *string `json:"StsAccessKeySecret,omitempty" xml:"StsAccessKeySecret,omitempty"`
	StsToken           *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s GetSmsAuthTokensResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBodyData) SetBizToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.BizToken = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetExpireTime(v int64) *GetSmsAuthTokensResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeyId(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeyId = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeySecret(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeySecret = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsToken = &v
	return s
}

type GetSmsAuthTokensResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSmsAuthTokensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSmsAuthTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponse) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponse) SetHeaders(v map[string]*string) *GetSmsAuthTokensResponse {
	s.Headers = v
	return s
}

func (s *GetSmsAuthTokensResponse) SetBody(v *GetSmsAuthTokensResponseBody) *GetSmsAuthTokensResponse {
	s.Body = v
	return s
}

type QueryGateVerifyBillingPublicRequest struct {
	AuthenticationType   *int32  `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	Month                *string `json:"Month,omitempty" xml:"Month,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s QueryGateVerifyBillingPublicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyBillingPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetMonth(v string) *QueryGateVerifyBillingPublicRequest {
	s.Month = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetOwnerId(v int64) *QueryGateVerifyBillingPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyBillingPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type QueryGateVerifyBillingPublicResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryGateVerifyBillingPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetCode(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetData(v *QueryGateVerifyBillingPublicResponseBodyData) *QueryGateVerifyBillingPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetMessage(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetRequestId(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.RequestId = &v
	return s
}

type QueryGateVerifyBillingPublicResponseBodyData struct {
	AmountSum        *string                                                         `json:"AmountSum,omitempty" xml:"AmountSum,omitempty"`
	SceneBillingList []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList `json:"SceneBillingList,omitempty" xml:"SceneBillingList,omitempty" type:"Repeated"`
}

func (s QueryGateVerifyBillingPublicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetAmountSum(v string) *QueryGateVerifyBillingPublicResponseBodyData {
	s.AmountSum = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetSceneBillingList(v []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) *QueryGateVerifyBillingPublicResponseBodyData {
	s.SceneBillingList = v
	return s
}

type QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList struct {
	Add         *string `json:"Add,omitempty" xml:"Add,omitempty"`
	Amount      *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ItemName    *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	SceneCode   *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	SceneName   *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	SinglePrice *string `json:"SinglePrice,omitempty" xml:"SinglePrice,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAdd(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Add = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAmount(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Amount = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAppName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.AppName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetItemName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.ItemName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneCode(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSinglePrice(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SinglePrice = &v
	return s
}

type QueryGateVerifyBillingPublicResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGateVerifyBillingPublicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGateVerifyBillingPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyBillingPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) SetBody(v *QueryGateVerifyBillingPublicResponseBody) *QueryGateVerifyBillingPublicResponse {
	s.Body = v
	return s
}

type QueryGateVerifyStatisticPublicRequest struct {
	AuthenticationType   *int32  `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SceneCode            *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryGateVerifyStatisticPublicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyStatisticPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetEndDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOsType(v string) *QueryGateVerifyStatisticPublicRequest {
	s.OsType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOwnerId(v int64) *QueryGateVerifyStatisticPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetSceneCode(v string) *QueryGateVerifyStatisticPublicRequest {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetStartDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.StartDate = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryGateVerifyStatisticPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetCode(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetData(v *QueryGateVerifyStatisticPublicResponseBodyData) *QueryGateVerifyStatisticPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetMessage(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetRequestId(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.RequestId = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBodyData struct {
	DayStatistic []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic `json:"DayStatistic,omitempty" xml:"DayStatistic,omitempty" type:"Repeated"`
	Total        *int64                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalFail    *int64                                                        `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	TotalSuccess *int64                                                        `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	TotalUnknown *int64                                                        `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetDayStatistic(v []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.DayStatistic = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotal(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalUnknown = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic struct {
	StatisticDateStr *string `json:"StatisticDateStr,omitempty" xml:"StatisticDateStr,omitempty"`
	TotalFail        *int64  `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	TotalSuccess     *int64  `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	TotalUnknown     *int64  `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetStatisticDateStr(v string) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.StatisticDateStr = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalUnknown = &v
	return s
}

type QueryGateVerifyStatisticPublicResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGateVerifyStatisticPublicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGateVerifyStatisticPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyStatisticPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) SetBody(v *QueryGateVerifyStatisticPublicResponseBody) *QueryGateVerifyStatisticPublicResponse {
	s.Body = v
	return s
}

type VerifyMobileRequest struct {
	AccessCode           *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileRequest) GoString() string {
	return s.String()
}

func (s *VerifyMobileRequest) SetAccessCode(v string) *VerifyMobileRequest {
	s.AccessCode = &v
	return s
}

func (s *VerifyMobileRequest) SetOutId(v string) *VerifyMobileRequest {
	s.OutId = &v
	return s
}

func (s *VerifyMobileRequest) SetOwnerId(v int64) *VerifyMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyMobileRequest) SetPhoneNumber(v string) *VerifyMobileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerAccount(v string) *VerifyMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerId(v int64) *VerifyMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

type VerifyMobileResponseBody struct {
	Code                *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	GateVerifyResultDTO *VerifyMobileResponseBodyGateVerifyResultDTO `json:"GateVerifyResultDTO,omitempty" xml:"GateVerifyResultDTO,omitempty" type:"Struct"`
	Message             *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBody) SetCode(v string) *VerifyMobileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyMobileResponseBody) SetGateVerifyResultDTO(v *VerifyMobileResponseBodyGateVerifyResultDTO) *VerifyMobileResponseBody {
	s.GateVerifyResultDTO = v
	return s
}

func (s *VerifyMobileResponseBody) SetMessage(v string) *VerifyMobileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyMobileResponseBody) SetRequestId(v string) *VerifyMobileResponseBody {
	s.RequestId = &v
	return s
}

type VerifyMobileResponseBodyGateVerifyResultDTO struct {
	VerifyId     *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyId(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyId = &v
	return s
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyResult(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyResult = &v
	return s
}

type VerifyMobileResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyMobileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponse) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponse) SetHeaders(v map[string]*string) *VerifyMobileResponse {
	s.Headers = v
	return s
}

func (s *VerifyMobileResponse) SetBody(v *VerifyMobileResponseBody) *VerifyMobileResponse {
	s.Body = v
	return s
}

type VerifyPhoneWithTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpToken              *string `json:"SpToken,omitempty" xml:"SpToken,omitempty"`
}

func (s VerifyPhoneWithTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenRequest) SetOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetPhoneNumber(v string) *VerifyPhoneWithTokenRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetSpToken(v string) *VerifyPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

type VerifyPhoneWithTokenResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	GateVerify *VerifyPhoneWithTokenResponseBodyGateVerify `json:"GateVerify,omitempty" xml:"GateVerify,omitempty" type:"Struct"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBody) SetCode(v string) *VerifyPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetGateVerify(v *VerifyPhoneWithTokenResponseBodyGateVerify) *VerifyPhoneWithTokenResponseBody {
	s.GateVerify = v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetMessage(v string) *VerifyPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetRequestId(v string) *VerifyPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

type VerifyPhoneWithTokenResponseBodyGateVerify struct {
	VerifyId     *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyId(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyId = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyResult(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyResult = &v
	return s
}

type VerifyPhoneWithTokenResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyPhoneWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyPhoneWithTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponse) SetHeaders(v map[string]*string) *VerifyPhoneWithTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyPhoneWithTokenResponse) SetBody(v *VerifyPhoneWithTokenResponseBody) *VerifyPhoneWithTokenResponse {
	s.Body = v
	return s
}

type VerifySmsCodeRequest struct {
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	SmsCode     *string `json:"SmsCode,omitempty" xml:"SmsCode,omitempty"`
	SmsToken    *string `json:"SmsToken,omitempty" xml:"SmsToken,omitempty"`
}

func (s VerifySmsCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeRequest) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeRequest) SetPhoneNumber(v string) *VerifySmsCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsCode(v string) *VerifySmsCodeRequest {
	s.SmsCode = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsToken(v string) *VerifySmsCodeRequest {
	s.SmsToken = &v
	return s
}

type VerifySmsCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifySmsCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponseBody) SetCode(v string) *VerifySmsCodeResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetData(v bool) *VerifySmsCodeResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetMessage(v string) *VerifySmsCodeResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetRequestId(v string) *VerifySmsCodeResponseBody {
	s.RequestId = &v
	return s
}

type VerifySmsCodeResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifySmsCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifySmsCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponse) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponse) SetHeaders(v map[string]*string) *VerifySmsCodeResponse {
	s.Headers = v
	return s
}

func (s *VerifySmsCodeResponse) SetBody(v *VerifySmsCodeResponseBody) *VerifySmsCodeResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dypnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateVerifySchemeWithOptions(request *CreateVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackName)) {
		query["PackName"] = request.PackName
	}

	if !tea.BoolValue(util.IsUnset(request.PackSign)) {
		query["PackSign"] = request.PackSign
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVerifyScheme(request *CreateVerifySchemeRequest) (_result *CreateVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CreateVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVerifySchemeWithOptions(request *DeleteVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *DeleteVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["CustomerId"] = request.CustomerId
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

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVerifyScheme(request *DeleteVerifySchemeRequest) (_result *DeleteVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.DeleteVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifySchemeWithOptions(request *DescribeVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["CustomerId"] = request.CustomerId
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

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyScheme(request *DescribeVerifySchemeRequest) (_result *DescribeVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.DescribeVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["Origin"] = request.Origin
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

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthorizationUrlWithOptions(request *GetAuthorizationUrlRequest, runtime *util.RuntimeOptions) (_result *GetAuthorizationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNo)) {
		query["PhoneNo"] = request.PhoneNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeId)) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthorizationUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthorizationUrl(request *GetAuthorizationUrlRequest) (_result *GetAuthorizationUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.GetAuthorizationUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCertifyResultWithOptions(request *GetCertifyResultRequest, runtime *util.RuntimeOptions) (_result *GetCertifyResultResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCertifyResult"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCertifyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCertifyResult(request *GetCertifyResultRequest) (_result *GetCertifyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCertifyResultResponse{}
	_body, _err := client.GetCertifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMobileWithOptions(request *GetMobileRequest, runtime *util.RuntimeOptions) (_result *GetMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMobile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMobile(request *GetMobileRequest) (_result *GetMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMobileResponse{}
	_body, _err := client.GetMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSmsAuthTokensWithOptions(request *GetSmsAuthTokensRequest, runtime *util.RuntimeOptions) (_result *GetSmsAuthTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsCodeExpire)) {
		query["SmsCodeExpire"] = request.SmsCodeExpire
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateCode)) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSmsAuthTokens"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSmsAuthTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSmsAuthTokens(request *GetSmsAuthTokensRequest) (_result *GetSmsAuthTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSmsAuthTokensResponse{}
	_body, _err := client.GetSmsAuthTokensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGateVerifyBillingPublicWithOptions(request *QueryGateVerifyBillingPublicRequest, runtime *util.RuntimeOptions) (_result *QueryGateVerifyBillingPublicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationType)) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !tea.BoolValue(util.IsUnset(request.Month)) {
		query["Month"] = request.Month
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGateVerifyBillingPublic"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGateVerifyBillingPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGateVerifyBillingPublic(request *QueryGateVerifyBillingPublicRequest) (_result *QueryGateVerifyBillingPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGateVerifyBillingPublicResponse{}
	_body, _err := client.QueryGateVerifyBillingPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGateVerifyStatisticPublicWithOptions(request *QueryGateVerifyStatisticPublicRequest, runtime *util.RuntimeOptions) (_result *QueryGateVerifyStatisticPublicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationType)) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGateVerifyStatisticPublic"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGateVerifyStatisticPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGateVerifyStatisticPublic(request *QueryGateVerifyStatisticPublicRequest) (_result *QueryGateVerifyStatisticPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGateVerifyStatisticPublicResponse{}
	_body, _err := client.QueryGateVerifyStatisticPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyMobileWithOptions(request *VerifyMobileRequest, runtime *util.RuntimeOptions) (_result *VerifyMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessCode)) {
		query["AccessCode"] = request.AccessCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyMobile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyMobile(request *VerifyMobileRequest) (_result *VerifyMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMobileResponse{}
	_body, _err := client.VerifyMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyPhoneWithTokenWithOptions(request *VerifyPhoneWithTokenRequest, runtime *util.RuntimeOptions) (_result *VerifyPhoneWithTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.SpToken)) {
		query["SpToken"] = request.SpToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPhoneWithToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyPhoneWithToken(request *VerifyPhoneWithTokenRequest) (_result *VerifyPhoneWithTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.VerifyPhoneWithTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifySmsCodeWithOptions(request *VerifySmsCodeRequest, runtime *util.RuntimeOptions) (_result *VerifySmsCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SmsCode)) {
		query["SmsCode"] = request.SmsCode
	}

	if !tea.BoolValue(util.IsUnset(request.SmsToken)) {
		query["SmsToken"] = request.SmsToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySmsCode"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifySmsCode(request *VerifySmsCodeRequest) (_result *VerifySmsCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.VerifySmsCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
