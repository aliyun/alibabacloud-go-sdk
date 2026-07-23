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
	SetParameters(v interface{}) *UpdateConnectionRequest
	GetParameters() interface{}
	SetType(v string) *UpdateConnectionRequest
	GetType() *string
}

type UpdateConnectionRequest struct {
	// The data structure of the authentication parameters.
	AuthParameters *UpdateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The name of the connection to be updated. The maximum length is 127 characters. The minimum length is 2 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description. The maximum length is 255 characters.
	//
	// example:
	//
	// Description of the connection configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data structure of the network configuration.
	//
	// This parameter is required.
	NetworkParameters *UpdateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
	// The data source connection parameters (JSON object). For specific field definitions, call the GetConnectionType API and refer to the ParamsSchema in the response.
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	Parameters interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The connection type. Valid values: MySQL, PostgreSQL, Elasticsearch, and Http.
	//
	// example:
	//
	// Http
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *UpdateConnectionRequest) GetParameters() interface{} {
	return s.Parameters
}

func (s *UpdateConnectionRequest) GetType() *string {
	return s.Type
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

func (s *UpdateConnectionRequest) SetParameters(v interface{}) *UpdateConnectionRequest {
	s.Parameters = v
	return s
}

func (s *UpdateConnectionRequest) SetType(v string) *UpdateConnectionRequest {
	s.Type = &v
	return s
}

func (s *UpdateConnectionRequest) Validate() error {
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

type UpdateConnectionRequestAuthParameters struct {
	// The data structure of API Key authentication.
	ApiKeyAuthParameters *UpdateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type. Valid values:
	//
	// - BASIC: BASIC_AUTH. This authorization method is the basic authentication method implemented by browsers in compliance with the HTTP protocol. During communication using the HTTP protocol, the HTTP protocol defines basic authentication that allows the HTTP server to authenticate the user identity of the client. Add Authorization: Basic followed by one space and the Base64-encoded value of `username:password` to the request header in a fixed format. Username and Password are required.
	//
	// - API KEY: API_KEY_AUTH. Add Token: TokenValue to the request header in a fixed format. ApiKeyName and ApiKeyValue are required.
	//
	// - OAUTH: OAUTH_AUTH. OAuth 2.0 is an authorization mechanism. In normal cases, in a system that does not use an authorization mechanism such as OAuth 2.0, the client can directly access resources on the resource server. To ensure secure data access for users, the Access Token mechanism is added between the client and the resource server. The client must carry an Access Token to access protected resources. Therefore, OAuth 2.0 prevents resources from being accessed by malicious clients, thereby improving the security of the system. AuthorizationEndpoint, OAuthHttpParameters, and HttpMethod are required.
	//
	// example:
	//
	// BASIC_AUTH
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The data structure of Basic authentication.
	BasicAuthParameters *UpdateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The data structure of the OAuth authentication parameters.
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

type UpdateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key of the API Key.
	//
	// example:
	//
	// name
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API Key.
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
	// The password for Basic authentication.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for Basic authentication.
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
	// The request URL for obtaining the OAuth token. The maximum length is 127 characters.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The data structure of the client parameters.
	ClientParameters *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
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
	// The request parameters for OAuth authentication.
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

type UpdateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
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
	// The list of data structures for body request parameters.
	BodyParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The list of request header parameters.
	HeaderParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The data structure of request query parameters.
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

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Specifies whether the parameter is used for authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the body request parameter.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the body request parameter.
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
	// Specifies whether the parameter is used for authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request header parameter.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header parameter.
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
	// Specifies whether the parameter is used for authentication.
	//
	// example:
	//
	// false
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request query parameter.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request query parameter.
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
	// - Public network: PublicNetwork
	//
	// - Virtual Private Cloud (VPC): PrivateNetwork
	//
	//
	// > When you select PrivateNetwork, VpcId, VswitcheId, and SecurityGroupId are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// PublicNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// eb-167adad548759-security_grop/sg-bp1addad26peuh9qh9rtyb
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the VPC.
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
