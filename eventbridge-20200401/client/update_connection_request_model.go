// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthParameters(v *UpdateConnectionRequestAuthParameters) *UpdateConnectionRequest
	GetAuthParameters() *UpdateConnectionRequestAuthParameters
	SetConnectionName(v string) *UpdateConnectionRequest
	GetConnectionName() *string
	SetDescription(v string) *UpdateConnectionRequest
	GetDescription() *string
	SetNetworkParameters(v *UpdateConnectionRequestNetworkParameters) *UpdateConnectionRequest
	GetNetworkParameters() *UpdateConnectionRequestNetworkParameters
}

type UpdateConnectionRequest struct {
	// The parameters that are configured for authentication.
	AuthParameters *UpdateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The name of the connection that you want to update. The name must be 2 to 127 characters in length.
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
	// The description of the connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	//
	// This parameter is required.
	NetworkParameters *UpdateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequest) GetAuthParameters() *UpdateConnectionRequestAuthParameters {
	return s.AuthParameters
}

func (s *UpdateConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectionRequest) GetNetworkParameters() *UpdateConnectionRequestNetworkParameters {
	return s.NetworkParameters
}

func (s *UpdateConnectionRequest) SetAuthParameters(v *UpdateConnectionRequestAuthParameters) *UpdateConnectionRequest {
	s.AuthParameters = v
	return s
}

func (s *UpdateConnectionRequest) SetConnectionName(v string) *UpdateConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionRequest) SetDescription(v string) *UpdateConnectionRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionRequest) SetNetworkParameters(v *UpdateConnectionRequestNetworkParameters) *UpdateConnectionRequest {
	s.NetworkParameters = v
	return s
}

func (s *UpdateConnectionRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParameters struct {
	// The parameters configured for API key authentication.
	ApiKeyAuthParameters *UpdateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
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
	BasicAuthParameters *UpdateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are configured for OAuth authentication.
	OAuthParameters *UpdateConnectionRequestAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequestAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParameters) GetApiKeyAuthParameters() *UpdateConnectionRequestAuthParametersApiKeyAuthParameters {
	return s.ApiKeyAuthParameters
}

func (s *UpdateConnectionRequestAuthParameters) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *UpdateConnectionRequestAuthParameters) GetBasicAuthParameters() *UpdateConnectionRequestAuthParametersBasicAuthParameters {
	return s.BasicAuthParameters
}

func (s *UpdateConnectionRequestAuthParameters) GetOAuthParameters() *UpdateConnectionRequestAuthParametersOAuthParameters {
	return s.OAuthParameters
}

func (s *UpdateConnectionRequestAuthParameters) SetApiKeyAuthParameters(v *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) *UpdateConnectionRequestAuthParameters {
	s.ApiKeyAuthParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParameters) SetAuthorizationType(v string) *UpdateConnectionRequestAuthParameters {
	s.AuthorizationType = &v
	return s
}

func (s *UpdateConnectionRequestAuthParameters) SetBasicAuthParameters(v *UpdateConnectionRequestAuthParametersBasicAuthParameters) *UpdateConnectionRequestAuthParameters {
	s.BasicAuthParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParameters) SetOAuthParameters(v *UpdateConnectionRequestAuthParametersOAuthParameters) *UpdateConnectionRequestAuthParameters {
	s.OAuthParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	//
	// example:
	//
	// name
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	//
	// example:
	//
	// demo
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersApiKeyAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *UpdateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *UpdateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersBasicAuthParameters struct {
	// The password for basic authentication.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for basic authentication.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersBasicAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) GetPassword() *string {
	return s.Password
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) GetUsername() *string {
	return s.Username
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) SetPassword(v string) *UpdateConnectionRequestAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) SetUsername(v string) *UpdateConnectionRequestAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token. The endpoint can be up to 127 characters in length.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are configured for the client.
	ClientParameters *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
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
	OAuthHttpParameters *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) GetClientParameters() *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters {
	return s.ClientParameters
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) GetOAuthHttpParameters() *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	return s.OAuthHttpParameters
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) SetAuthorizationEndpoint(v string) *UpdateConnectionRequestAuthParametersOAuthParameters {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) SetClientParameters(v *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) *UpdateConnectionRequestAuthParametersOAuthParameters {
	s.ClientParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) SetHttpMethod(v string) *UpdateConnectionRequestAuthParametersOAuthParameters {
	s.HttpMethod = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) SetOAuthHttpParameters(v *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) *UpdateConnectionRequestAuthParametersOAuthParameters {
	s.OAuthHttpParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
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
	// ClientSecret
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) GetClientID() *string {
	return s.ClientID
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientID(v string) *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The parameters that are configured for the request body.
	BodyParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request header.
	HeaderParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request path.
	QueryStringParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetBodyParameters() []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	return s.BodyParameters
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetHeaderParameters() []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GetQueryStringParameters() []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	return s.QueryStringParameters
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetBodyParameters(v []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetHeaderParameters(v []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) SetQueryStringParameters(v []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters {
	s.QueryStringParameters = v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
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
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	//
	// example:
	//
	// demo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetKey() *string {
	return s.Key
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetIsValueSecret(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.IsValueSecret = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetKey(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetValue(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Value = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
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
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// demo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetKey() *string {
	return s.Key
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetIsValueSecret(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.IsValueSecret = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetKey(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetValue(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
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
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	//
	// example:
	//
	// demo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetKey() *string {
	return s.Key
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetIsValueSecret(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.IsValueSecret = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetKey(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetValue(v string) *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateConnectionRequestNetworkParameters struct {
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
	// eb-167adad548759-security_grop/sg-bp1addad26peuh9qh9rtyb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// eb-test/vpc-bp1symadadwnwgmqud
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1iu4x7aeradadown1og8,vsw-bp193sqmadadlaszpeqbt2c
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s UpdateConnectionRequestNetworkParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestNetworkParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestNetworkParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateConnectionRequestNetworkParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateConnectionRequestNetworkParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateConnectionRequestNetworkParameters) GetVswitcheId() *string {
	return s.VswitcheId
}

func (s *UpdateConnectionRequestNetworkParameters) SetNetworkType(v string) *UpdateConnectionRequestNetworkParameters {
	s.NetworkType = &v
	return s
}

func (s *UpdateConnectionRequestNetworkParameters) SetSecurityGroupId(v string) *UpdateConnectionRequestNetworkParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateConnectionRequestNetworkParameters) SetVpcId(v string) *UpdateConnectionRequestNetworkParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateConnectionRequestNetworkParameters) SetVswitcheId(v string) *UpdateConnectionRequestNetworkParameters {
	s.VswitcheId = &v
	return s
}

func (s *UpdateConnectionRequestNetworkParameters) Validate() error {
	return dara.Validate(s)
}
