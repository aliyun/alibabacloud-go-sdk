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

type AssumeRoleRequest struct {
	DurationSeconds *int64  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RoleArn         *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
}

func (s AssumeRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleRequest) SetDurationSeconds(v int64) *AssumeRoleRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleRequest) SetPolicy(v string) *AssumeRoleRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleArn(v string) *AssumeRoleRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleSessionName(v string) *AssumeRoleRequest {
	s.RoleSessionName = &v
	return s
}

type AssumeRoleResponseBody struct {
	AssumedRoleUser *AssumeRoleResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials     *AssumeRoleResponseBodyCredentials     `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	RequestId       *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssumeRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBody) SetAssumedRoleUser(v *AssumeRoleResponseBodyAssumedRoleUser) *AssumeRoleResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleResponseBody) SetCredentials(v *AssumeRoleResponseBodyCredentials) *AssumeRoleResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleResponseBody) SetRequestId(v string) *AssumeRoleResponseBody {
	s.RequestId = &v
	return s
}

type AssumeRoleResponseBodyAssumedRoleUser struct {
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

type AssumeRoleResponseBodyCredentials struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetExpiration(v string) *AssumeRoleResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

type AssumeRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssumeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssumeRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponse) SetHeaders(v map[string]*string) *AssumeRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleResponse) SetStatusCode(v int32) *AssumeRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleResponse) SetBody(v *AssumeRoleResponseBody) *AssumeRoleResponse {
	s.Body = v
	return s
}

type AssumeRoleWithOIDCRequest struct {
	DurationSeconds *int64  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	OIDCProviderArn *string `json:"OIDCProviderArn,omitempty" xml:"OIDCProviderArn,omitempty"`
	OIDCToken       *string `json:"OIDCToken,omitempty" xml:"OIDCToken,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RoleArn         *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
}

func (s AssumeRoleWithOIDCRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCRequest) SetDurationSeconds(v int64) *AssumeRoleWithOIDCRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetOIDCProviderArn(v string) *AssumeRoleWithOIDCRequest {
	s.OIDCProviderArn = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetOIDCToken(v string) *AssumeRoleWithOIDCRequest {
	s.OIDCToken = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetPolicy(v string) *AssumeRoleWithOIDCRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleArn(v string) *AssumeRoleWithOIDCRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleSessionName(v string) *AssumeRoleWithOIDCRequest {
	s.RoleSessionName = &v
	return s
}

type AssumeRoleWithOIDCResponseBody struct {
	AssumedRoleUser *AssumeRoleWithOIDCResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials     *AssumeRoleWithOIDCResponseBodyCredentials     `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	OIDCTokenInfo   *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo   `json:"OIDCTokenInfo,omitempty" xml:"OIDCTokenInfo,omitempty" type:"Struct"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBody) SetAssumedRoleUser(v *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) *AssumeRoleWithOIDCResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetCredentials(v *AssumeRoleWithOIDCResponseBodyCredentials) *AssumeRoleWithOIDCResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetOIDCTokenInfo(v *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) *AssumeRoleWithOIDCResponseBody {
	s.OIDCTokenInfo = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetRequestId(v string) *AssumeRoleWithOIDCResponseBody {
	s.RequestId = &v
	return s
}

type AssumeRoleWithOIDCResponseBodyAssumedRoleUser struct {
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

type AssumeRoleWithOIDCResponseBodyCredentials struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

type AssumeRoleWithOIDCResponseBodyOIDCTokenInfo struct {
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	Issuer    *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Subject   *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetClientIds(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.ClientIds = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetIssuer(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetSubject(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Subject = &v
	return s
}

type AssumeRoleWithOIDCResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssumeRoleWithOIDCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssumeRoleWithOIDCResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponse) SetHeaders(v map[string]*string) *AssumeRoleWithOIDCResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetStatusCode(v int32) *AssumeRoleWithOIDCResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetBody(v *AssumeRoleWithOIDCResponseBody) *AssumeRoleWithOIDCResponse {
	s.Body = v
	return s
}

type AssumeRoleWithSAMLRequest struct {
	DurationSeconds *int64  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RoleArn         *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SAMLAssertion   *string `json:"SAMLAssertion,omitempty" xml:"SAMLAssertion,omitempty"`
	SAMLProviderArn *string `json:"SAMLProviderArn,omitempty" xml:"SAMLProviderArn,omitempty"`
}

func (s AssumeRoleWithSAMLRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLRequest) SetDurationSeconds(v int64) *AssumeRoleWithSAMLRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetPolicy(v string) *AssumeRoleWithSAMLRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetRoleArn(v string) *AssumeRoleWithSAMLRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetSAMLAssertion(v string) *AssumeRoleWithSAMLRequest {
	s.SAMLAssertion = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetSAMLProviderArn(v string) *AssumeRoleWithSAMLRequest {
	s.SAMLProviderArn = &v
	return s
}

type AssumeRoleWithSAMLResponseBody struct {
	AssumedRoleUser   *AssumeRoleWithSAMLResponseBodyAssumedRoleUser   `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials       *AssumeRoleWithSAMLResponseBodyCredentials       `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SAMLAssertionInfo *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo `json:"SAMLAssertionInfo,omitempty" xml:"SAMLAssertionInfo,omitempty" type:"Struct"`
}

func (s AssumeRoleWithSAMLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBody) SetAssumedRoleUser(v *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) *AssumeRoleWithSAMLResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetCredentials(v *AssumeRoleWithSAMLResponseBodyCredentials) *AssumeRoleWithSAMLResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetRequestId(v string) *AssumeRoleWithSAMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetSAMLAssertionInfo(v *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) *AssumeRoleWithSAMLResponseBody {
	s.SAMLAssertionInfo = v
	return s
}

type AssumeRoleWithSAMLResponseBodyAssumedRoleUser struct {
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

type AssumeRoleWithSAMLResponseBodyCredentials struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

type AssumeRoleWithSAMLResponseBodySAMLAssertionInfo struct {
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Recipient   *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
	Subject     *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetIssuer(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetRecipient(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Recipient = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubject(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Subject = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubjectType(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.SubjectType = &v
	return s
}

type AssumeRoleWithSAMLResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssumeRoleWithSAMLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssumeRoleWithSAMLResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponse) SetHeaders(v map[string]*string) *AssumeRoleWithSAMLResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetStatusCode(v int32) *AssumeRoleWithSAMLResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetBody(v *AssumeRoleWithSAMLResponseBody) *AssumeRoleWithSAMLResponse {
	s.Body = v
	return s
}

type GetCallerIdentityResponseBody struct {
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Arn          *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	PrincipalId  *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleId       *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCallerIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallerIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponseBody) SetAccountId(v string) *GetCallerIdentityResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetArn(v string) *GetCallerIdentityResponseBody {
	s.Arn = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetIdentityType(v string) *GetCallerIdentityResponseBody {
	s.IdentityType = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetPrincipalId(v string) *GetCallerIdentityResponseBody {
	s.PrincipalId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRequestId(v string) *GetCallerIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRoleId(v string) *GetCallerIdentityResponseBody {
	s.RoleId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetUserId(v string) *GetCallerIdentityResponseBody {
	s.UserId = &v
	return s
}

type GetCallerIdentityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCallerIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallerIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallerIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponse) SetHeaders(v map[string]*string) *GetCallerIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetCallerIdentityResponse) SetStatusCode(v int32) *GetCallerIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallerIdentityResponse) SetBody(v *GetCallerIdentityResponseBody) *GetCallerIdentityResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("sts.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("sts.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("sts.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("sts.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("sts.aliyuncs.com"),
		"cn-edge-1":                   tea.String("sts.aliyuncs.com"),
		"cn-fujian":                   tea.String("sts.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("sts.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("sts.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("sts.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("sts-vpc.cn-north-2-gov-1.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("sts.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("sts.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("sts.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("sts.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("sts-vpc.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("sts.aliyuncs.com"),
		"cn-wuhan":                    tea.String("sts.aliyuncs.com"),
		"cn-yushanfang":               tea.String("sts.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("sts.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("sts.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("sts.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("sts.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("sts.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("sts.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AssumeRoleWithOptions(request *AssumeRoleRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.RoleSessionName)) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRole"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssumeRole(request *AssumeRoleRequest) (_result *AssumeRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleResponse{}
	_body, _err := client.AssumeRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssumeRoleWithOIDCWithOptions(request *AssumeRoleWithOIDCRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleWithOIDCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderArn)) {
		query["OIDCProviderArn"] = request.OIDCProviderArn
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCToken)) {
		query["OIDCToken"] = request.OIDCToken
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.RoleSessionName)) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRoleWithOIDC"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssumeRoleWithOIDC(request *AssumeRoleWithOIDCRequest) (_result *AssumeRoleWithOIDCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.AssumeRoleWithOIDCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssumeRoleWithSAMLWithOptions(request *AssumeRoleWithSAMLRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleWithSAMLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLAssertion)) {
		query["SAMLAssertion"] = request.SAMLAssertion
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLProviderArn)) {
		query["SAMLProviderArn"] = request.SAMLProviderArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRoleWithSAML"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssumeRoleWithSAML(request *AssumeRoleWithSAMLRequest) (_result *AssumeRoleWithSAMLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.AssumeRoleWithSAMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallerIdentityWithOptions(runtime *util.RuntimeOptions) (_result *GetCallerIdentityResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetCallerIdentity"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallerIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallerIdentity() (_result *GetCallerIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallerIdentityResponse{}
	_body, _err := client.GetCallerIdentityWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
