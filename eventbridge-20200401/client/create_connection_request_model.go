// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthParameters(v *CreateConnectionRequestAuthParameters) *CreateConnectionRequest
	GetAuthParameters() *CreateConnectionRequestAuthParameters
	SetConnectionName(v string) *CreateConnectionRequest
	GetConnectionName() *string
	SetDescription(v string) *CreateConnectionRequest
	GetDescription() *string
	SetNetworkParameters(v *CreateConnectionRequestNetworkParameters) *CreateConnectionRequest
	GetNetworkParameters() *CreateConnectionRequestNetworkParameters
}

type CreateConnectionRequest struct {
	// The parameters that are configured for authentication.
	AuthParameters *CreateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	//
	// This parameter is required.
	NetworkParameters *CreateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequest) GetAuthParameters() *CreateConnectionRequestAuthParameters {
	return s.AuthParameters
}

func (s *CreateConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConnectionRequest) GetNetworkParameters() *CreateConnectionRequestNetworkParameters {
	return s.NetworkParameters
}

func (s *CreateConnectionRequest) SetAuthParameters(v *CreateConnectionRequestAuthParameters) *CreateConnectionRequest {
	s.AuthParameters = v
	return s
}

func (s *CreateConnectionRequest) SetConnectionName(v string) *CreateConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateConnectionRequest) SetDescription(v string) *CreateConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateConnectionRequest) SetNetworkParameters(v *CreateConnectionRequestNetworkParameters) *CreateConnectionRequest {
	s.NetworkParameters = v
	return s
}

func (s *CreateConnectionRequest) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParameters struct {
	// The parameters that are configured for API key authentication.
	ApiKeyAuthParameters *CreateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type. Valid values:
	//
	// 	- BASIC: basic authentication. Basic authentication is a simple authentication scheme built into the HTTP protocol. When you use the HTTP protocol for communications, the authentication method that the HTTP server uses to authenticate user identities on the client is defined in the protocol. The request header is in the Authorization: Basic Base64-encoded string (`Username:Password`) format. If you use this authentication method, you must configure Username and Password.
	//
	// 	- API_KEY_AUTH: API key authentication. The request header is in the Token: Token value format. If you use this authentication method, you must configure ApiKeyName and ApiKeyValue.
	//
	// 	- OAUTH_AUTH: OAuth authentication. OAuth2.0 is an authentication mechanism. In normal cases, a system that does not use OAuth2.0 can access the resources of the server from the client. To ensure access security, access tokens are used to authenticate users in OAuth 2.0. The client must use an access token to access protected resources. This way, OAuth 2.0 protects resources from being accessed from malicious clients and improves system security. If you use this authentication method, you must configure AuthorizationEndpoint, OAuthHttpParameters, and HttpMethod.
	//
	// example:
	//
	// BASIC_AUTH
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The parameters that are configured for basic authentication.
	BasicAuthParameters *CreateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are configured for OAuth authentication.
	OAuthParameters *CreateConnectionRequestAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequestAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParameters) GetApiKeyAuthParameters() *CreateConnectionRequestAuthParametersApiKeyAuthParameters {
	return s.ApiKeyAuthParameters
}

func (s *CreateConnectionRequestAuthParameters) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *CreateConnectionRequestAuthParameters) GetBasicAuthParameters() *CreateConnectionRequestAuthParametersBasicAuthParameters {
	return s.BasicAuthParameters
}

func (s *CreateConnectionRequestAuthParameters) GetOAuthParameters() *CreateConnectionRequestAuthParametersOAuthParameters {
	return s.OAuthParameters
}

func (s *CreateConnectionRequestAuthParameters) SetApiKeyAuthParameters(v *CreateConnectionRequestAuthParametersApiKeyAuthParameters) *CreateConnectionRequestAuthParameters {
	s.ApiKeyAuthParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParameters) SetAuthorizationType(v string) *CreateConnectionRequestAuthParameters {
	s.AuthorizationType = &v
	return s
}

func (s *CreateConnectionRequestAuthParameters) SetBasicAuthParameters(v *CreateConnectionRequestAuthParametersBasicAuthParameters) *CreateConnectionRequestAuthParameters {
	s.BasicAuthParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParameters) SetOAuthParameters(v *CreateConnectionRequestAuthParametersOAuthParameters) *CreateConnectionRequestAuthParameters {
	s.OAuthParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	//
	// example:
	//
	// Token
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	//
	// example:
	//
	// adkjnakddh****
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s CreateConnectionRequestAuthParametersApiKeyAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *CreateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *CreateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersBasicAuthParameters struct {
	// The password for basic authentication.
	//
	// example:
	//
	// *******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for basic authentication.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateConnectionRequestAuthParametersBasicAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) GetPassword() *string {
	return s.Password
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) GetUsername() *string {
	return s.Username
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) SetPassword(v string) *CreateConnectionRequestAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) SetUsername(v string) *CreateConnectionRequestAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParameters struct {
	// The endpoint of the authorized client. The endpoint can be up to 127 characters in length.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are configured for the client.
	ClientParameters *CreateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
	//
	// 	- DELETE
	//
	// 	- PUT
	//
	// 	- PATCH
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters of OAuth authentication.
	OAuthHttpParameters *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequestAuthParametersOAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) GetClientParameters() *CreateConnectionRequestAuthParametersOAuthParametersClientParameters {
	return s.ClientParameters
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) GetOAuthHttpParameters() *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	return s.OAuthHttpParameters
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) SetAuthorizationEndpoint(v string) *CreateConnectionRequestAuthParametersOAuthParameters {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) SetClientParameters(v *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) *CreateConnectionRequestAuthParametersOAuthParameters {
	s.ClientParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) SetHttpMethod(v string) *CreateConnectionRequestAuthParametersOAuthParameters {
	s.HttpMethod = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) SetOAuthHttpParameters(v *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) *CreateConnectionRequestAuthParametersOAuthParameters {
	s.OAuthHttpParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	//
	// example:
	//
	// ClientID
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The AccessKey secret of the client.
	//
	// example:
	//
	// Qo57Q~F249~S74GmNPA36pZJoJK4f4LY****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersClientParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) GetClientID() *string {
	return s.ClientID
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientID(v string) *CreateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *CreateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The parameters that are configured for the request body.
	BodyParameters []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request header.
	HeaderParameters []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request path.
	QueryStringParameters []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetBodyParameters() []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	return s.BodyParameters
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetHeaderParameters() []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetQueryStringParameters() []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	return s.QueryStringParameters
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetBodyParameters(v []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetHeaderParameters(v []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetQueryStringParameters(v []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.QueryStringParameters = v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request body.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	//
	// example:
	//
	// keyValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetKey() *string {
	return s.Key
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetValue() *string {
	return s.Value
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetIsValueSecret(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.IsValueSecret = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetKey(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetValue(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Value = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request header.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// keyValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetKey() *string {
	return s.Key
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetValue() *string {
	return s.Value
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetIsValueSecret(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.IsValueSecret = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetKey(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetValue(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Specifies whether to enable authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request path.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	//
	// example:
	//
	// valueDemo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetKey() *string {
	return s.Key
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetValue() *string {
	return s.Value
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetIsValueSecret(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.IsValueSecret = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetKey(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetValue(v string) *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestNetworkParameters struct {
	// 	- PublicNetwork: the Internet.
	//
	// 	- PrivateNetwork: virtual private cloud (VPC).
	//
	// >  If you set this parameter to PrivateNetwork, you must also configure VpcId, VswitchId, and SecurityGroupId.
	//
	// This parameter is required.
	//
	// example:
	//
	// PublicNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// eb-167adad548759-security_grop/sg-bp1addad26peuh9qh9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// eb-test/vpc-bp1symadadwnwg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1iu4x7aeradadown1og8,vsw-bp193sqmadadlaszpeq****
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s CreateConnectionRequestNetworkParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestNetworkParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestNetworkParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateConnectionRequestNetworkParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateConnectionRequestNetworkParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateConnectionRequestNetworkParameters) GetVswitcheId() *string {
	return s.VswitcheId
}

func (s *CreateConnectionRequestNetworkParameters) SetNetworkType(v string) *CreateConnectionRequestNetworkParameters {
	s.NetworkType = &v
	return s
}

func (s *CreateConnectionRequestNetworkParameters) SetSecurityGroupId(v string) *CreateConnectionRequestNetworkParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateConnectionRequestNetworkParameters) SetVpcId(v string) *CreateConnectionRequestNetworkParameters {
	s.VpcId = &v
	return s
}

func (s *CreateConnectionRequestNetworkParameters) SetVswitcheId(v string) *CreateConnectionRequestNetworkParameters {
	s.VswitcheId = &v
	return s
}

func (s *CreateConnectionRequestNetworkParameters) Validate() error {
	return dara.Validate(s)
}
