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
	RequestId       *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssumedRoleUser *AssumeRoleResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials     *AssumeRoleResponseBodyCredentials     `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
}

func (s AssumeRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBody) SetRequestId(v string) *AssumeRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleResponseBody) SetAssumedRoleUser(v *AssumeRoleResponseBodyAssumedRoleUser) *AssumeRoleResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleResponseBody) SetCredentials(v *AssumeRoleResponseBodyCredentials) *AssumeRoleResponseBody {
	s.Credentials = v
	return s
}

type AssumeRoleResponseBodyAssumedRoleUser struct {
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s AssumeRoleResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

type AssumeRoleResponseBodyCredentials struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s AssumeRoleResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetExpiration(v string) *AssumeRoleResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

type AssumeRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssumeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssumeRoleResponse) SetBody(v *AssumeRoleResponseBody) *AssumeRoleResponse {
	s.Body = v
	return s
}

type AssumeRoleWithOIDCRequest struct {
	// OIDC Provider的ARN
	OIDCProviderArn *string `json:"OIDCProviderArn,omitempty" xml:"OIDCProviderArn,omitempty"`
	// 需要扮演的角色的ARN
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// OIDC的ID Token，需输入原始Token，无需Base64解码
	OIDCToken *string `json:"OIDCToken,omitempty" xml:"OIDCToken,omitempty"`
	// 权限策略。 生成STS Token时可以指定一个额外的权限策略，以进一步限制STS Token的权限。若不指定则返回的Token拥有指定角色的所有权限。
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Session过期时间，单位为秒。
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// 用户自定义参数。此参数用来区分不同的令牌，可用于用户级别的访问审计。
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
}

func (s AssumeRoleWithOIDCRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCRequest) SetOIDCProviderArn(v string) *AssumeRoleWithOIDCRequest {
	s.OIDCProviderArn = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleArn(v string) *AssumeRoleWithOIDCRequest {
	s.RoleArn = &v
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

func (s *AssumeRoleWithOIDCRequest) SetDurationSeconds(v int64) *AssumeRoleWithOIDCRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleSessionName(v string) *AssumeRoleWithOIDCRequest {
	s.RoleSessionName = &v
	return s
}

type AssumeRoleWithOIDCResponseBody struct {
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OIDCTokenInfo   *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo   `json:"OIDCTokenInfo,omitempty" xml:"OIDCTokenInfo,omitempty" type:"Struct"`
	AssumedRoleUser *AssumeRoleWithOIDCResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials     *AssumeRoleWithOIDCResponseBodyCredentials     `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
}

func (s AssumeRoleWithOIDCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBody) SetRequestId(v string) *AssumeRoleWithOIDCResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetOIDCTokenInfo(v *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) *AssumeRoleWithOIDCResponseBody {
	s.OIDCTokenInfo = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetAssumedRoleUser(v *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) *AssumeRoleWithOIDCResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetCredentials(v *AssumeRoleWithOIDCResponseBodyCredentials) *AssumeRoleWithOIDCResponseBody {
	s.Credentials = v
	return s
}

type AssumeRoleWithOIDCResponseBodyOIDCTokenInfo struct {
	Subject   *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Issuer    *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetSubject(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Subject = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetIssuer(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetClientIds(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.ClientIds = &v
	return s
}

type AssumeRoleWithOIDCResponseBodyAssumedRoleUser struct {
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

type AssumeRoleWithOIDCResponseBodyCredentials struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

type AssumeRoleWithOIDCResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssumeRoleWithOIDCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssumeRoleWithOIDCResponse) SetBody(v *AssumeRoleWithOIDCResponseBody) *AssumeRoleWithOIDCResponse {
	s.Body = v
	return s
}

type AssumeRoleWithSAMLRequest struct {
	SAMLProviderArn *string `json:"SAMLProviderArn,omitempty" xml:"SAMLProviderArn,omitempty"`
	RoleArn         *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SAMLAssertion   *string `json:"SAMLAssertion,omitempty" xml:"SAMLAssertion,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	DurationSeconds *int64  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
}

func (s AssumeRoleWithSAMLRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLRequest) SetSAMLProviderArn(v string) *AssumeRoleWithSAMLRequest {
	s.SAMLProviderArn = &v
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

func (s *AssumeRoleWithSAMLRequest) SetPolicy(v string) *AssumeRoleWithSAMLRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetDurationSeconds(v int64) *AssumeRoleWithSAMLRequest {
	s.DurationSeconds = &v
	return s
}

type AssumeRoleWithSAMLResponseBody struct {
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SAMLAssertionInfo *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo `json:"SAMLAssertionInfo,omitempty" xml:"SAMLAssertionInfo,omitempty" type:"Struct"`
	AssumedRoleUser   *AssumeRoleWithSAMLResponseBodyAssumedRoleUser   `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials       *AssumeRoleWithSAMLResponseBodyCredentials       `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
}

func (s AssumeRoleWithSAMLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBody) SetRequestId(v string) *AssumeRoleWithSAMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetSAMLAssertionInfo(v *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) *AssumeRoleWithSAMLResponseBody {
	s.SAMLAssertionInfo = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetAssumedRoleUser(v *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) *AssumeRoleWithSAMLResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetCredentials(v *AssumeRoleWithSAMLResponseBodyCredentials) *AssumeRoleWithSAMLResponseBody {
	s.Credentials = v
	return s
}

type AssumeRoleWithSAMLResponseBodySAMLAssertionInfo struct {
	SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
	Subject     *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Recipient   *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubjectType(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.SubjectType = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubject(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Subject = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetIssuer(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetRecipient(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Recipient = &v
	return s
}

type AssumeRoleWithSAMLResponseBodyAssumedRoleUser struct {
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
	Arn           *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

type AssumeRoleWithSAMLResponseBodyCredentials struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

type AssumeRoleWithSAMLResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssumeRoleWithSAMLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssumeRoleWithSAMLResponse) SetBody(v *AssumeRoleWithSAMLResponseBody) *AssumeRoleWithSAMLResponse {
	s.Body = v
	return s
}

type GetCallerIdentityResponseBody struct {
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	AccountId    *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrincipalId  *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Arn          *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	RoleId       *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s GetCallerIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallerIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponseBody) SetIdentityType(v string) *GetCallerIdentityResponseBody {
	s.IdentityType = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetAccountId(v string) *GetCallerIdentityResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRequestId(v string) *GetCallerIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetPrincipalId(v string) *GetCallerIdentityResponseBody {
	s.PrincipalId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetUserId(v string) *GetCallerIdentityResponseBody {
	s.UserId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetArn(v string) *GetCallerIdentityResponseBody {
	s.Arn = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRoleId(v string) *GetCallerIdentityResponseBody {
	s.RoleId = &v
	return s
}

type GetCallerIdentityResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCallerIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssumeRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssumeRole"), tea.String("2015-04-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssumeRoleWithOIDC"), tea.String("2015-04-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssumeRoleWithSAML"), tea.String("2015-04-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_result = &GetCallerIdentityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCallerIdentity"), tea.String("2015-04-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
