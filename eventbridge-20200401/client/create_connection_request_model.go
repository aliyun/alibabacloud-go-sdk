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
	SetParameters(v interface{}) *CreateConnectionRequest
	GetParameters() interface{}
	SetType(v string) *CreateConnectionRequest
	GetType() *string
}

type CreateConnectionRequest struct {
	// The authentication configuration.
	AuthParameters *CreateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The connection configuration name. Maximum length: 127 characters. Minimum length: 2 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection configuration. Maximum length: 255 characters.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The network configuration.
	//
	// This parameter is required.
	NetworkParameters *CreateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
	// The data source connection parameters (JSON object). This parameter is required when Type is set to a data source type. This parameter is not required for the Http type. For specific field definitions, call the GetConnectionType operation and refer to ParamsSchema in the response.
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	Parameters interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The connection type. Valid values: MySQL, PostgreSQL, Elasticsearch, and Http. This parameter is required for data source connections. If this parameter is not specified, the default value Http is used. The Http type is used for HTTP protocol targets such as API Destination. Data source types are used for data connections in the integration marketplace.
	//
	// example:
	//
	// Http
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *CreateConnectionRequest) GetParameters() interface{} {
	return s.Parameters
}

func (s *CreateConnectionRequest) GetType() *string {
	return s.Type
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

func (s *CreateConnectionRequest) SetParameters(v interface{}) *CreateConnectionRequest {
	s.Parameters = v
	return s
}

func (s *CreateConnectionRequest) SetType(v string) *CreateConnectionRequest {
	s.Type = &v
	return s
}

func (s *CreateConnectionRequest) Validate() error {
	if s.AuthParameters != nil {
		if err := s.AuthParameters.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkParameters != nil {
		if err := s.NetworkParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectionRequestAuthParameters struct {
	// The API key authentication configuration.
	ApiKeyAuthParameters *CreateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type:
	//
	// - BASIC: BASIC_AUTH. This authorization method is a basic authorization method implemented by browsers in compliance with the HTTP protocol. During HTTP communication, the HTTP protocol defines a basic authentication method that allows an HTTP server to authenticate clients. Add `Authorization: Basic Base64Encoded(username:password)` in the fixed format to the request header. Username and Password are required.
	//
	// - API KEY: API_KEY_AUTH. Add `Token: TokenValue` in the fixed format to the request header. ApiKeyName and ApiKeyValue are required.
	//
	// - OAUTH: OAUTH_AUTH. OAuth 2.0 is an authorization mechanism. In a system that does not use an authorization mechanism such as OAuth 2.0, the client can directly access resources on the resource server. To ensure secure data access, an Access Token mechanism is added. The client must carry an Access Token to access protected resources. OAuth 2.0 prevents resources from being accessed by malicious clients, which improves system security. AuthorizationEndpoint, OAuthHttpParameters, and HttpMethod are required.
	//
	// example:
	//
	// BASIC_AUTH
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The basic authentication configuration.
	BasicAuthParameters *CreateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The OAuth authentication configuration.
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
	if s.ApiKeyAuthParameters != nil {
		if err := s.ApiKeyAuthParameters.Validate(); err != nil {
			return err
		}
	}
	if s.BasicAuthParameters != nil {
		if err := s.BasicAuthParameters.Validate(); err != nil {
			return err
		}
	}
	if s.OAuthParameters != nil {
		if err := s.OAuthParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key name of the API key.
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
	// The authorization endpoint URL. Maximum length: 127 characters.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The client parameter configuration.
	ClientParameters *CreateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP method. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// - HEAD
	//
	// - DELETE
	//
	// - PUT
	//
	// - PATCH
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The OAuth authentication request parameters.
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
	if s.ClientParameters != nil {
		if err := s.ClientParameters.Validate(); err != nil {
			return err
		}
	}
	if s.OAuthHttpParameters != nil {
		if err := s.OAuthHttpParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	//
	// example:
	//
	// ClientID
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The client secret of the application.
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
	// The list of body request parameter configurations.
	BodyParameters []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The list of header parameter configurations.
	HeaderParameters []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The structure of the URI of the request path parameters.
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
	if s.BodyParameters != nil {
		for _, item := range s.BodyParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HeaderParameters != nil {
		for _, item := range s.HeaderParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryStringParameters != nil {
		for _, item := range s.QueryStringParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Specifies whether the value is a secret.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the body request parameter.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the body request parameter.
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
	// Specifies whether the value is a secret.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the header parameter.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header parameter.
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
	// Specifies whether the value is a secret.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the URI of the request path parameter.
	//
	// example:
	//
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the URI of the request path parameter.
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
	// - Public network: PublicNetwork
	//
	// - Virtual private cloud (VPC): PrivateNetwork
	//
	// 	Notice: If you select PrivateNetwork, VpcId, VswitcheId, and SecurityGroupId are required.
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
	// The ID of the VPC.
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
