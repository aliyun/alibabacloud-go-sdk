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

type CreateVerifySchemeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeName           *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PackName             *string `json:"PackName,omitempty" xml:"PackName,omitempty"`
	PackSign             *string `json:"PackSign,omitempty" xml:"PackSign,omitempty"`
	BundleId             *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
}

func (s CreateVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeRequest) SetOwnerId(v int64) *CreateVerifySchemeRequest {
	s.OwnerId = &v
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

func (s *CreateVerifySchemeRequest) SetAppName(v string) *CreateVerifySchemeRequest {
	s.AppName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOsType(v string) *CreateVerifySchemeRequest {
	s.OsType = &v
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

func (s *CreateVerifySchemeRequest) SetBundleId(v string) *CreateVerifySchemeRequest {
	s.BundleId = &v
	return s
}

type CreateVerifySchemeResponseBody struct {
	Message             *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GateVerifySchemeDTO *CreateVerifySchemeResponseBodyGateVerifySchemeDTO `json:"GateVerifySchemeDTO,omitempty" xml:"GateVerifySchemeDTO,omitempty" type:"Struct"`
	Code                *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBody) SetMessage(v string) *CreateVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetRequestId(v string) *CreateVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetGateVerifySchemeDTO(v *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) *CreateVerifySchemeResponseBody {
	s.GateVerifySchemeDTO = v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetCode(v string) *CreateVerifySchemeResponseBody {
	s.Code = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeCode           *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
}

func (s DeleteVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeRequest) GoString() string {
	return s.String()
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

func (s *DeleteVerifySchemeRequest) SetCustomerId(v int64) *DeleteVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

type DeleteVerifySchemeResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponseBody) SetMessage(v string) *DeleteVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetRequestId(v string) *DeleteVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetCode(v string) *DeleteVerifySchemeResponseBody {
	s.Code = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemeCode           *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
}

func (s DescribeVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeRequest) GoString() string {
	return s.String()
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

func (s *DescribeVerifySchemeRequest) SetCustomerId(v int64) *DescribeVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

type DescribeVerifySchemeResponseBody struct {
	Message              *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                 *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	SchemeQueryResultDTO *DescribeVerifySchemeResponseBodySchemeQueryResultDTO `json:"SchemeQueryResultDTO,omitempty" xml:"SchemeQueryResultDTO,omitempty" type:"Struct"`
}

func (s DescribeVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBody) SetMessage(v string) *DescribeVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetRequestId(v string) *DescribeVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetCode(v string) *DescribeVerifySchemeResponseBody {
	s.Code = &v
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

type GetAuthorizationUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	SchemeId             *int64  `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
}

func (s GetAuthorizationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlRequest) SetOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.OwnerId = &v
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

func (s *GetAuthorizationUrlRequest) SetPhoneNo(v string) *GetAuthorizationUrlRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetSchemeId(v int64) *GetAuthorizationUrlRequest {
	s.SchemeId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetEndDate(v string) *GetAuthorizationUrlRequest {
	s.EndDate = &v
	return s
}

type GetAuthorizationUrlResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAuthorizationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAuthorizationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBody) SetMessage(v string) *GetAuthorizationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetRequestId(v string) *GetAuthorizationUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetData(v *GetAuthorizationUrlResponseBodyData) *GetAuthorizationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetCode(v string) *GetAuthorizationUrlResponseBody {
	s.Code = &v
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

type GetAuthTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Origin               *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
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

func (s *GetAuthTokenRequest) SetOrigin(v string) *GetAuthTokenRequest {
	s.Origin = &v
	return s
}

type GetAuthTokenResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenInfo *GetAuthTokenResponseBodyTokenInfo `json:"TokenInfo,omitempty" xml:"TokenInfo,omitempty" type:"Struct"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
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

func (s *GetAuthTokenResponseBody) SetCode(v string) *GetAuthTokenResponseBody {
	s.Code = &v
	return s
}

type GetAuthTokenResponseBodyTokenInfo struct {
	JwtToken    *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s GetAuthTokenResponseBodyTokenInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyTokenInfo) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetJwtToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.JwtToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetAccessToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.AccessToken = &v
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

type GetCertifyResultRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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

func (s *GetCertifyResultRequest) SetProductCode(v string) *GetCertifyResultRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCertifyResultRequest) SetToken(v string) *GetCertifyResultRequest {
	s.Token = &v
	return s
}

type GetCertifyResultResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetCertifyResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetCertifyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertifyResultResponseBody) SetMessage(v string) *GetCertifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetCertifyResultResponseBody) SetRequestId(v string) *GetCertifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertifyResultResponseBody) SetData(v []*GetCertifyResultResponseBodyData) *GetCertifyResultResponseBody {
	s.Data = v
	return s
}

func (s *GetCertifyResultResponseBody) SetCode(v string) *GetCertifyResultResponseBody {
	s.Code = &v
	return s
}

type GetCertifyResultResponseBodyData struct {
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	VerifyDesc   *string `json:"VerifyDesc,omitempty" xml:"VerifyDesc,omitempty"`
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	DeviceToken  *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
}

func (s GetCertifyResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCertifyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCertifyResultResponseBodyData) SetMaterialInfo(v string) *GetCertifyResultResponseBodyData {
	s.MaterialInfo = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetVerifyDesc(v string) *GetCertifyResultResponseBodyData {
	s.VerifyDesc = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetIdentityInfo(v string) *GetCertifyResultResponseBodyData {
	s.IdentityInfo = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetVerifyResult(v string) *GetCertifyResultResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *GetCertifyResultResponseBodyData) SetDeviceToken(v string) *GetCertifyResultResponseBodyData {
	s.DeviceToken = &v
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
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccessToken          *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s GetMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMobileRequest) GoString() string {
	return s.String()
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

func (s *GetMobileRequest) SetAccessToken(v string) *GetMobileRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMobileRequest) SetOutId(v string) *GetMobileRequest {
	s.OutId = &v
	return s
}

type GetMobileResponseBody struct {
	Message            *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code               *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	GetMobileResultDTO *GetMobileResponseBodyGetMobileResultDTO `json:"GetMobileResultDTO,omitempty" xml:"GetMobileResultDTO,omitempty" type:"Struct"`
}

func (s GetMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponseBody) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBody) SetMessage(v string) *GetMobileResponseBody {
	s.Message = &v
	return s
}

func (s *GetMobileResponseBody) SetRequestId(v string) *GetMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMobileResponseBody) SetCode(v string) *GetMobileResponseBody {
	s.Code = &v
	return s
}

func (s *GetMobileResponseBody) SetGetMobileResultDTO(v *GetMobileResponseBodyGetMobileResultDTO) *GetMobileResponseBody {
	s.GetMobileResultDTO = v
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

type TwiceTelVerifyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Since                *string `json:"Since,omitempty" xml:"Since,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s TwiceTelVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s TwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *TwiceTelVerifyRequest) SetOwnerId(v int64) *TwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *TwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *TwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TwiceTelVerifyRequest) SetResourceOwnerId(v int64) *TwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TwiceTelVerifyRequest) SetSince(v string) *TwiceTelVerifyRequest {
	s.Since = &v
	return s
}

func (s *TwiceTelVerifyRequest) SetPhoneNumber(v string) *TwiceTelVerifyRequest {
	s.PhoneNumber = &v
	return s
}

type TwiceTelVerifyResponseBody struct {
	Message              *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TwiceTelVerifyResult *TwiceTelVerifyResponseBodyTwiceTelVerifyResult `json:"TwiceTelVerifyResult,omitempty" xml:"TwiceTelVerifyResult,omitempty" type:"Struct"`
	Code                 *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TwiceTelVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *TwiceTelVerifyResponseBody) SetMessage(v string) *TwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *TwiceTelVerifyResponseBody) SetRequestId(v string) *TwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *TwiceTelVerifyResponseBody) SetTwiceTelVerifyResult(v *TwiceTelVerifyResponseBodyTwiceTelVerifyResult) *TwiceTelVerifyResponseBody {
	s.TwiceTelVerifyResult = v
	return s
}

func (s *TwiceTelVerifyResponseBody) SetCode(v string) *TwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

type TwiceTelVerifyResponseBodyTwiceTelVerifyResult struct {
	Carrier      *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	VerifyResult *int32  `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s TwiceTelVerifyResponseBodyTwiceTelVerifyResult) String() string {
	return tea.Prettify(s)
}

func (s TwiceTelVerifyResponseBodyTwiceTelVerifyResult) GoString() string {
	return s.String()
}

func (s *TwiceTelVerifyResponseBodyTwiceTelVerifyResult) SetCarrier(v string) *TwiceTelVerifyResponseBodyTwiceTelVerifyResult {
	s.Carrier = &v
	return s
}

func (s *TwiceTelVerifyResponseBodyTwiceTelVerifyResult) SetVerifyResult(v int32) *TwiceTelVerifyResponseBodyTwiceTelVerifyResult {
	s.VerifyResult = &v
	return s
}

type TwiceTelVerifyResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TwiceTelVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TwiceTelVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s TwiceTelVerifyResponse) GoString() string {
	return s.String()
}

func (s *TwiceTelVerifyResponse) SetHeaders(v map[string]*string) *TwiceTelVerifyResponse {
	s.Headers = v
	return s
}

func (s *TwiceTelVerifyResponse) SetBody(v *TwiceTelVerifyResponseBody) *TwiceTelVerifyResponse {
	s.Body = v
	return s
}

type VerifyMobileRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AccessCode           *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s VerifyMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileRequest) GoString() string {
	return s.String()
}

func (s *VerifyMobileRequest) SetOwnerId(v int64) *VerifyMobileRequest {
	s.OwnerId = &v
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

func (s *VerifyMobileRequest) SetAccessCode(v string) *VerifyMobileRequest {
	s.AccessCode = &v
	return s
}

func (s *VerifyMobileRequest) SetPhoneNumber(v string) *VerifyMobileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyMobileRequest) SetOutId(v string) *VerifyMobileRequest {
	s.OutId = &v
	return s
}

type VerifyMobileResponseBody struct {
	GateVerifyResultDTO *VerifyMobileResponseBodyGateVerifyResultDTO `json:"GateVerifyResultDTO,omitempty" xml:"GateVerifyResultDTO,omitempty" type:"Struct"`
	Message             *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s VerifyMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBody) GoString() string {
	return s.String()
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

func (s *VerifyMobileResponseBody) SetCode(v string) *VerifyMobileResponseBody {
	s.Code = &v
	return s
}

type VerifyMobileResponseBodyGateVerifyResultDTO struct {
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	VerifyId     *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyResult(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyResult = &v
	return s
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyId(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyId = &v
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
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
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

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetPhoneNumber(v string) *VerifyPhoneWithTokenRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetSpToken(v string) *VerifyPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

type VerifyPhoneWithTokenResponseBody struct {
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GateVerify *VerifyPhoneWithTokenResponseBodyGateVerify `json:"GateVerify,omitempty" xml:"GateVerify,omitempty" type:"Struct"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBody) SetMessage(v string) *VerifyPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetRequestId(v string) *VerifyPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetGateVerify(v *VerifyPhoneWithTokenResponseBodyGateVerify) *VerifyPhoneWithTokenResponseBody {
	s.GateVerify = v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetCode(v string) *VerifyPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

type VerifyPhoneWithTokenResponseBodyGateVerify struct {
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	VerifyId     *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyResult(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyResult = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyId(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyId = &v
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVerifyScheme"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVerifyScheme"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyScheme"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAuthorizationUrlWithOptions(request *GetAuthorizationUrlRequest, runtime *util.RuntimeOptions) (_result *GetAuthorizationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuthorizationUrl"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuthToken"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetCertifyResultWithOptions(request *GetCertifyResultRequest, runtime *util.RuntimeOptions) (_result *GetCertifyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCertifyResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCertifyResult"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMobileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMobile"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TwiceTelVerifyWithOptions(request *TwiceTelVerifyRequest, runtime *util.RuntimeOptions) (_result *TwiceTelVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TwiceTelVerifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TwiceTelVerify"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TwiceTelVerify(request *TwiceTelVerifyRequest) (_result *TwiceTelVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TwiceTelVerifyResponse{}
	_body, _err := client.TwiceTelVerifyWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyMobileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyMobile"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyPhoneWithToken"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
