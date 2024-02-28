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

type CreateApiDestinationRequest struct {
	// The name of the API destination. The name must be 2 to 127 characters in length.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// >
	// >  Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	HttpApiParameters *CreateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s CreateApiDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationRequest) SetApiDestinationName(v string) *CreateApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *CreateApiDestinationRequest) SetConnectionName(v string) *CreateApiDestinationRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateApiDestinationRequest) SetDescription(v string) *CreateApiDestinationRequest {
	s.Description = &v
	return s
}

func (s *CreateApiDestinationRequest) SetHttpApiParameters(v *CreateApiDestinationRequestHttpApiParameters) *CreateApiDestinationRequest {
	s.HttpApiParameters = v
	return s
}

type CreateApiDestinationRequestHttpApiParameters struct {
	// The endpoint of the API destination. The endpoint can be up to 127 characters in length.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// *   GET
	// *   POST
	// *   HEAD
	// *   DELETE
	// *   PUT
	// *   PATCH
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s CreateApiDestinationRequestHttpApiParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationRequestHttpApiParameters) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationRequestHttpApiParameters) SetEndpoint(v string) *CreateApiDestinationRequestHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *CreateApiDestinationRequestHttpApiParameters) SetMethod(v string) *CreateApiDestinationRequestHttpApiParameters {
	s.Method = &v
	return s
}

type CreateApiDestinationShrinkRequest struct {
	// The name of the API destination. The name must be 2 to 127 characters in length.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// >
	// >  Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	HttpApiParametersShrink *string `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty"`
}

func (s CreateApiDestinationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationShrinkRequest) SetApiDestinationName(v string) *CreateApiDestinationShrinkRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetConnectionName(v string) *CreateApiDestinationShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetDescription(v string) *CreateApiDestinationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetHttpApiParametersShrink(v string) *CreateApiDestinationShrinkRequest {
	s.HttpApiParametersShrink = &v
	return s
}

type CreateApiDestinationResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the API destination is created.
	Date *CreateApiDestinationResponseBodyDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponseBody) SetCode(v string) *CreateApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiDestinationResponseBody) SetDate(v *CreateApiDestinationResponseBodyDate) *CreateApiDestinationResponseBody {
	s.Date = v
	return s
}

func (s *CreateApiDestinationResponseBody) SetMessage(v string) *CreateApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiDestinationResponseBody) SetRequestId(v string) *CreateApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApiDestinationResponseBodyDate struct {
	// The name of the API destination.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s CreateApiDestinationResponseBodyDate) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationResponseBodyDate) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponseBodyDate) SetApiDestinationName(v string) *CreateApiDestinationResponseBodyDate {
	s.ApiDestinationName = &v
	return s
}

type CreateApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationResponse) SetHeaders(v map[string]*string) *CreateApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *CreateApiDestinationResponse) SetStatusCode(v int32) *CreateApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiDestinationResponse) SetBody(v *CreateApiDestinationResponseBody) *CreateApiDestinationResponse {
	s.Body = v
	return s
}

type CreateConnectionRequest struct {
	// The parameters that are configured for authentication.
	AuthParameters *CreateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	NetworkParameters *CreateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequest) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParameters struct {
	// The parameters that are configured for API key authentication.
	ApiKeyAuthParameters *CreateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type. Valid values:
	//
	// BASIC_AUTH: basic authentication.
	//
	// Introduction: Basic authentication is a simple authentication scheme built into the HTTP protocol. When you use the HTTP protocol for communications, the authentication method that the HTTP server uses to authenticate user identities on the client is defined in the protocol. The request header is in the Authorization: Basic Base64-encoded string (Username:Password) format.
	//
	// 1.  Username and Password are required.
	//
	// API_KEY_AUTH: API key authentication.
	//
	// Introduction: The request header is in the Token: Token value format.
	//
	// *   ApiKeyName and ApiKeyValue are required.
	//
	// OAUTH_AUTH: OAuth authentication.
	//
	// Introduction: OAuth2.0 is an authentication mechanism. In normal cases, a system that does not use OAuth2.0 can access the resources of the server from the client. To ensure access security, access tokens are used to authenticate users in OAuth 2.0. The client must use an access token to access protected resources. This way, OAuth 2.0 protects resources from being accessed from malicious clients and improves system security.
	//
	// *   AuthorizationEndpoint, OAuthHttpParameters, and HttpMethod are required.
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The parameters that are configured for basic authentication.
	BasicAuthParameters *CreateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are configured for OAuth authentication.
	OAuthParameters *CreateConnectionRequestAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequestAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s CreateConnectionRequestAuthParametersApiKeyAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *CreateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *CreateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

type CreateConnectionRequestAuthParametersBasicAuthParameters struct {
	// The password for basic authentication.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for basic authentication.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateConnectionRequestAuthParametersBasicAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) SetPassword(v string) *CreateConnectionRequestAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersBasicAuthParameters) SetUsername(v string) *CreateConnectionRequestAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

type CreateConnectionRequestAuthParametersOAuthParameters struct {
	// The IP address of the authorized endpoint. The default value of a column can be up to 127 characters in length.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are configured for the client.
	ClientParameters *CreateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// *   GET
	// *   POST
	// *   HEAD
	// *   DELETE
	// *   PUT
	// *   PATCH
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters that are configured for OAuth authentication.
	OAuthHttpParameters *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s CreateConnectionRequestAuthParametersOAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The AccessKey secret of the client.
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersClientParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientID(v string) *CreateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *CreateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *CreateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
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
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request body.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request header.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request path.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
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

type CreateConnectionRequestNetworkParameters struct {
	// The network type. Valid values:
	//
	// PublicNetwork and PrivateNetwork.
	//
	// *   Note: If you set this parameter to PrivateNetwork, you must configure VpcId, VswitcheId, and SecurityGroupId.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VPC. ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s CreateConnectionRequestNetworkParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionRequestNetworkParameters) GoString() string {
	return s.String()
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

type CreateConnectionShrinkRequest struct {
	// The parameters that are configured for authentication.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
}

func (s CreateConnectionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionShrinkRequest) SetAuthParametersShrink(v string) *CreateConnectionShrinkRequest {
	s.AuthParametersShrink = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetConnectionName(v string) *CreateConnectionShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetDescription(v string) *CreateConnectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetNetworkParametersShrink(v string) *CreateConnectionShrinkRequest {
	s.NetworkParametersShrink = &v
	return s
}

type CreateConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBody) SetCode(v string) *CreateConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConnectionResponseBody) SetData(v *CreateConnectionResponseBodyData) *CreateConnectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateConnectionResponseBody) SetMessage(v string) *CreateConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConnectionResponseBody) SetRequestId(v string) *CreateConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateConnectionResponseBodyData struct {
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s CreateConnectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBodyData) SetConnectionName(v string) *CreateConnectionResponseBodyData {
	s.ConnectionName = &v
	return s
}

type CreateConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponse) SetHeaders(v map[string]*string) *CreateConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectionResponse) SetStatusCode(v int32) *CreateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnectionResponse) SetBody(v *CreateConnectionResponseBody) *CreateConnectionResponse {
	s.Body = v
	return s
}

type CreateEventBusRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s CreateEventBusRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventBusRequest) GoString() string {
	return s.String()
}

func (s *CreateEventBusRequest) SetDescription(v string) *CreateEventBusRequest {
	s.Description = &v
	return s
}

func (s *CreateEventBusRequest) SetEventBusName(v string) *CreateEventBusRequest {
	s.EventBusName = &v
	return s
}

type CreateEventBusResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. The value true indicates that the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventBusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponseBody) SetCode(v string) *CreateEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventBusResponseBody) SetData(v *CreateEventBusResponseBodyData) *CreateEventBusResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventBusResponseBody) SetMessage(v string) *CreateEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventBusResponseBody) SetRequestId(v string) *CreateEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventBusResponseBody) SetSuccess(v bool) *CreateEventBusResponseBody {
	s.Success = &v
	return s
}

type CreateEventBusResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
}

func (s CreateEventBusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEventBusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponseBodyData) SetEventBusARN(v string) *CreateEventBusResponseBodyData {
	s.EventBusARN = &v
	return s
}

type CreateEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventBusResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventBusResponse) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponse) SetHeaders(v map[string]*string) *CreateEventBusResponse {
	s.Headers = v
	return s
}

func (s *CreateEventBusResponse) SetStatusCode(v int32) *CreateEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventBusResponse) SetBody(v *CreateEventBusResponseBody) *CreateEventBusResponse {
	s.Body = v
	return s
}

type CreateEventSourceRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus with which the event source is associated.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParameters *CreateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParameters *CreateEventSourceRequestSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Service (MNS). If you specify MNS as the event source, you must configure RegionId, IsBase64Decode, and QueueName.
	SourceMNSParameters *CreateEventSourceRequestSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParameters *CreateEventSourceRequestSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParameters *CreateEventSourceRequestSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Log Service.
	SourceSLSParameters *CreateEventSourceRequestSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParameters *CreateEventSourceRequestSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
}

func (s CreateEventSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequest) SetDescription(v string) *CreateEventSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateEventSourceRequest) SetEventBusName(v string) *CreateEventSourceRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventSourceRequest) SetEventSourceName(v string) *CreateEventSourceRequest {
	s.EventSourceName = &v
	return s
}

func (s *CreateEventSourceRequest) SetSourceHttpEventParameters(v *CreateEventSourceRequestSourceHttpEventParameters) *CreateEventSourceRequest {
	s.SourceHttpEventParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceKafkaParameters(v *CreateEventSourceRequestSourceKafkaParameters) *CreateEventSourceRequest {
	s.SourceKafkaParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceMNSParameters(v *CreateEventSourceRequestSourceMNSParameters) *CreateEventSourceRequest {
	s.SourceMNSParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceRabbitMQParameters(v *CreateEventSourceRequestSourceRabbitMQParameters) *CreateEventSourceRequest {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceRocketMQParameters(v *CreateEventSourceRequestSourceRocketMQParameters) *CreateEventSourceRequest {
	s.SourceRocketMQParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceSLSParameters(v *CreateEventSourceRequestSourceSLSParameters) *CreateEventSourceRequest {
	s.SourceSLSParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceScheduledEventParameters(v *CreateEventSourceRequestSourceScheduledEventParameters) *CreateEventSourceRequest {
	s.SourceScheduledEventParameters = v
	return s
}

type CreateEventSourceRequestSourceHttpEventParameters struct {
	// The CIDR block that is used for security settings. This parameter is required only if you set SecurityConfig to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request method supported by the generated webhook URL. You can select multiple values. Valid values:
	//
	// *   GET
	// *   POST
	// *   PUT
	// *   PATCH
	// *   DELETE
	// *   HEAD
	// *   OPTIONS
	// *   TRACE
	// *   CONNECT
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	// The security domain name. This parameter is required only if you set SecurityConfig to referer. You can enter a domain name.
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	// The type of security settings. Valid values:
	//
	// *   none: No configuration is required.
	// *   ip: CIDR block.
	// *   referer: security domain name.
	SecurityConfig *string `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	// The protocol type that is supported by the generated webhook URL. Valid values:
	//
	// *   HTTP
	// *   HTTPS
	// *   HTTP\&HTTPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEventSourceRequestSourceHttpEventParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetIp(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetMethod(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetReferer(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetSecurityConfig(v string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetType(v string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Type = &v
	return s
}

type CreateEventSourceRequestSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of consumers.
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a self-managed network.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the Message Queue for Apache Kafka instance belongs. This parameter is required only if you set Network to PublicNetwork.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the Message Queue for Apache Kafka instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache Kafka instance is associated. This parameter is required only if you set Network to PublicNetwork.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC in which the Message Queue for Apache Kafka instance resides. This parameter is required only if you set Network to PublicNetwork.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventSourceRequestSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetConsumerGroup(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetMaximumTasks(v int32) *CreateEventSourceRequestSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetNetwork(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetOffsetReset(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetRegionId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetSecurityGroupId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetTopic(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetVSwitchIds(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetVpcId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type CreateEventSourceRequestSourceMNSParameters struct {
	// Specifies whether to enable Base64 decoding. Valid values: true and false.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region where the MNS queue resides. Valid values: cn-qingdao, cn-beijing, cn-zhangjiakou, cn-huhehaote, cn-wulanchabu, cn-hangzhou, cn-shanghai, cn-shenzhen, cn-guangzhou, cn-chengdu, cn-hongkong, ap-southeast-1, ap-southeast-2, ap-southeast-3, ap-southeast-5, ap-northeast-1, eu-central-1, us-west-1, us-east-1, ap-south-1, me-east-1, and cn-north-2-gov-1.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEventSourceRequestSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v bool) *CreateEventSourceRequestSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetQueueName(v string) *CreateEventSourceRequestSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetRegionId(v string) *CreateEventSourceRequestSourceMNSParameters {
	s.RegionId = &v
	return s
}

type CreateEventSourceRequestSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance. For more information, see Limits.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s CreateEventSourceRequestSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetQueueName(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetRegionId(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type CreateEventSourceRequestSourceRocketMQParameters struct {
	// The authentication type. You can set this parameter to ACL or leave this parameter empty.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// None.
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the Message Queue for Apache RocketMQ instance.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the Message Queue for Apache RocketMQ instance belongs.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The type of the Message Queue for Apache RocketMQ instance. Valid values:
	//
	// *   Cloud\_4: Message Queue for Apache RocketMQ 4.0 instance.
	// *   Cloud\_5: Message Queue for Apache RocketMQ 5.0 instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache RocketMQ instance is associated.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the Message Queue for Apache RocketMQ instance resides.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which message consumption starts. Valid values: CONSUME_FROM_LAST_OFFSET: Start message consumption from the latest offset. CONSUME_FROM_FIRST_OFFSET: Start message consumption from the earliest offset. CONSUME_FROM_TIMESTAMP: Start message consumption from the offset at the specified point in time. Default value: CONSUME_FROM_LAST_OFFSET.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUME_FROM_TIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventSourceRequestSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetAuthType(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceEndpoint(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceNetwork(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstancePassword(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceUsername(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceVpcId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetOffset(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetRegionId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTag(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v int64) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type CreateEventSourceRequestSourceSLSParameters struct {
	// The starting consumer offset. The value begin specifies the earliest offset, and the value end specifies the latest offset. You can also specify a time in seconds to start consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console. For information about the permission policy of this role, see Create a custom event source of the Log Service type.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventSourceRequestSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetConsumePosition(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetLogStore(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetProject(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetRoleName(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.RoleName = &v
	return s
}

type CreateEventSourceRequestSourceScheduledEventParameters struct {
	// The cron expression.
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone in which the cron expression is executed.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateEventSourceRequestSourceScheduledEventParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetSchedule(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetTimeZone(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetUserData(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

type CreateEventSourceShrinkRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus with which the event source is associated.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParametersShrink *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	// The parameters that are configured if the event source is Message Service (MNS). If you specify MNS as the event source, you must configure RegionId, IsBase64Decode, and QueueName.
	SourceMNSParametersShrink *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParametersShrink *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParametersShrink *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// The parameters that are configured if the event source is Log Service.
	SourceSLSParametersShrink *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParametersShrink *string `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty"`
}

func (s CreateEventSourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSourceShrinkRequest) SetDescription(v string) *CreateEventSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetEventBusName(v string) *CreateEventSourceShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetEventSourceName(v string) *CreateEventSourceShrinkRequest {
	s.EventSourceName = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceHttpEventParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceHttpEventParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceKafkaParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceKafkaParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceMNSParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceMNSParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceRabbitMQParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceRabbitMQParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceRocketMQParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceRocketMQParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceSLSParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceSLSParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceScheduledEventParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceScheduledEventParametersShrink = &v
	return s
}

type CreateEventSourceResponseBody struct {
	// The returned response code. Valid values:
	//
	// *   Success: The request is successful.
	// *   Other codes: The request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateEventSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponseBody) SetCode(v string) *CreateEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetData(v *CreateEventSourceResponseBodyData) *CreateEventSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventSourceResponseBody) SetMessage(v string) *CreateEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetRequestId(v string) *CreateEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetSuccess(v bool) *CreateEventSourceResponseBody {
	s.Success = &v
	return s
}

type CreateEventSourceResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	EventSourceARN *string `json:"EventSourceARN,omitempty" xml:"EventSourceARN,omitempty"`
}

func (s CreateEventSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponseBodyData) SetEventSourceARN(v string) *CreateEventSourceResponseBodyData {
	s.EventSourceARN = &v
	return s
}

type CreateEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponse) SetHeaders(v map[string]*string) *CreateEventSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateEventSourceResponse) SetStatusCode(v int32) *CreateEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventSourceResponse) SetBody(v *CreateEventSourceResponseBody) *CreateEventSourceResponse {
	s.Body = v
	return s
}

type CreateEventStreamingRequest struct {
	// The description of the event stream.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptions *CreateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target. You must and can specify only one event target.
	Sink *CreateEventStreamingRequestSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source. You must and can specify only one event source.
	Source     *CreateEventStreamingRequestSource       `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Transforms []*CreateEventStreamingRequestTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s CreateEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequest) SetDescription(v string) *CreateEventStreamingRequest {
	s.Description = &v
	return s
}

func (s *CreateEventStreamingRequest) SetEventStreamingName(v string) *CreateEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *CreateEventStreamingRequest) SetFilterPattern(v string) *CreateEventStreamingRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateEventStreamingRequest) SetRunOptions(v *CreateEventStreamingRequestRunOptions) *CreateEventStreamingRequest {
	s.RunOptions = v
	return s
}

func (s *CreateEventStreamingRequest) SetSink(v *CreateEventStreamingRequestSink) *CreateEventStreamingRequest {
	s.Sink = v
	return s
}

func (s *CreateEventStreamingRequest) SetSource(v *CreateEventStreamingRequestSource) *CreateEventStreamingRequest {
	s.Source = v
	return s
}

func (s *CreateEventStreamingRequest) SetTransforms(v []*CreateEventStreamingRequestTransforms) *CreateEventStreamingRequest {
	s.Transforms = v
	return s
}

type CreateEventStreamingRequestRunOptions struct {
	// The batch window.
	BatchWindow *CreateEventStreamingRequestRunOptionsBatchWindow `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	// Specifies whether to enable dead-letter queues. By default, dead-letter queues are disabled. Messages that fail to be pushed are discarded after the maximum number of retries that is specified by the retry policy is reached.
	DeadLetterQueue *CreateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The exception tolerance policy. Valid values:
	//
	// *   NONE: does not tolerate exceptions.
	// *   ALL: tolerates all exceptions.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The maximum number of concurrent threads.
	MaximumTasks *int64 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The retry policy that you want to use if events fail to be pushed.
	RetryStrategy *CreateEventStreamingRequestRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestRunOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptions) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptions) SetBatchWindow(v *CreateEventStreamingRequestRunOptionsBatchWindow) *CreateEventStreamingRequestRunOptions {
	s.BatchWindow = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetDeadLetterQueue(v *CreateEventStreamingRequestRunOptionsDeadLetterQueue) *CreateEventStreamingRequestRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetErrorsTolerance(v string) *CreateEventStreamingRequestRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetMaximumTasks(v int64) *CreateEventStreamingRequestRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptions) SetRetryStrategy(v *CreateEventStreamingRequestRunOptionsRetryStrategy) *CreateEventStreamingRequestRunOptions {
	s.RetryStrategy = v
	return s
}

type CreateEventStreamingRequestRunOptionsBatchWindow struct {
	// The maximum number of events that is allowed in the batch window. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsBatchWindow) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) SetCountBasedWindow(v int32) *CreateEventStreamingRequestRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *CreateEventStreamingRequestRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

type CreateEventStreamingRequestRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsDeadLetterQueue) SetArn(v string) *CreateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

type CreateEventStreamingRequestRunOptionsRetryStrategy struct {
	// The maximum timeout period for a retry.
	MaximumEventAgeInSeconds *int64 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	MaximumRetryAttempts *int64 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry policy. Valid values:
	//
	// *   BACKOFF_RETRY
	// *   EXPONENTIAL_DECAY_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s CreateEventStreamingRequestRunOptionsRetryStrategy) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v int64) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumRetryAttempts(v int64) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *CreateEventStreamingRequestRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *CreateEventStreamingRequestRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

type CreateEventStreamingRequestSink struct {
	// The parameters that are configured if you specify the event target as DataHub.
	SinkDataHubParameters *CreateEventStreamingRequestSinkSinkDataHubParameters `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Function Compute.
	SinkFcParameters *CreateEventStreamingRequestSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Serverless Workflow.
	SinkFnfParameters *CreateEventStreamingRequestSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Message Queue for Apache Kafka.
	SinkKafkaParameters *CreateEventStreamingRequestSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as MNS.
	SinkMNSParameters        *CreateEventStreamingRequestSinkSinkMNSParameters        `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkPrometheusParameters *CreateEventStreamingRequestSinkSinkPrometheusParameters `json:"SinkPrometheusParameters,omitempty" xml:"SinkPrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Message Queue for RabbitMQ.
	SinkRabbitMQParameters *CreateEventStreamingRequestSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Message Queue for Apache RocketMQ.
	SinkRocketMQParameters *CreateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Simple Log Service.
	SinkSLSParameters *CreateEventStreamingRequestSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSink) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSink) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSink) SetSinkDataHubParameters(v *CreateEventStreamingRequestSinkSinkDataHubParameters) *CreateEventStreamingRequestSink {
	s.SinkDataHubParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkFcParameters(v *CreateEventStreamingRequestSinkSinkFcParameters) *CreateEventStreamingRequestSink {
	s.SinkFcParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkFnfParameters(v *CreateEventStreamingRequestSinkSinkFnfParameters) *CreateEventStreamingRequestSink {
	s.SinkFnfParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkKafkaParameters(v *CreateEventStreamingRequestSinkSinkKafkaParameters) *CreateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkMNSParameters(v *CreateEventStreamingRequestSinkSinkMNSParameters) *CreateEventStreamingRequestSink {
	s.SinkMNSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkPrometheusParameters(v *CreateEventStreamingRequestSinkSinkPrometheusParameters) *CreateEventStreamingRequestSink {
	s.SinkPrometheusParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRabbitMQParameters(v *CreateEventStreamingRequestSinkSinkRabbitMQParameters) *CreateEventStreamingRequestSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkRocketMQParameters(v *CreateEventStreamingRequestSinkSinkRocketMQParameters) *CreateEventStreamingRequestSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkSLSParameters(v *CreateEventStreamingRequestSinkSinkSLSParameters) *CreateEventStreamingRequestSink {
	s.SinkSLSParameters = v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParameters struct {
	// The BLOB topic.
	Body *CreateEventStreamingRequestSinkSinkDataHubParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The name of the DataHub project.
	Project *CreateEventStreamingRequestSinkSinkDataHubParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name.
	RoleName *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the DataHub topic.
	Topic *CreateEventStreamingRequestSinkSinkDataHubParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The TUBLE topic.
	TopicSchema *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	// The topic type. Valid values:
	//
	// *   TUPLE
	// *   BLOB
	TopicType *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetBody(v *CreateEventStreamingRequestSinkSinkDataHubParametersBody) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetProject(v *CreateEventStreamingRequestSinkSinkDataHubParametersProject) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Project = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.RoleName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopicSchema(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicSchema = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParameters) SetTopicType(v *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) *CreateEventStreamingRequestSinkSinkDataHubParameters {
	s.TopicType = v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The BLOB topic.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersProject struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the DataHub project.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersProject) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersProject) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersProject) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersProject {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersRoleName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopic struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the DataHub topic.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopic {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The TUBLE topic.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkDataHubParametersTopicType struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic type. Valid values:
	//
	// *   TUPLE
	// *   BLOB
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetForm(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType) SetValue(v string) *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParameters struct {
	// The message body that you want to deliver to Function Compute.
	Body *CreateEventStreamingRequestSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *CreateEventStreamingRequestSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The function name.
	FunctionName *CreateEventStreamingRequestSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation method. Valid values: Sync and Async.
	InvocationType *CreateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The service version.
	Qualifier *CreateEventStreamingRequestSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *CreateEventStreamingRequestSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkFcParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetBody(v *CreateEventStreamingRequestSinkSinkFcParametersBody) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetConcurrency(v *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetFunctionName(v *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetInvocationType(v *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetQualifier(v *CreateEventStreamingRequestSinkSinkFcParametersQualifier) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParameters) SetServiceName(v *CreateEventStreamingRequestSinkSinkFcParametersServiceName) *CreateEventStreamingRequestSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersConcurrency struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersConcurrency) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersConcurrency) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersFunctionName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The function name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersFunctionName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersFunctionName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersInvocationType struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation method. Valid values: Sync and Async.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersInvocationType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersInvocationType) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersQualifier struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service version.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersQualifier) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersQualifier) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFcParametersServiceName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFcParametersServiceName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFcParametersServiceName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *CreateEventStreamingRequestSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The input information of the execution.
	Input *CreateEventStreamingRequestSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *CreateEventStreamingRequestSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetExecutionName(v *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetFlowName(v *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetInput(v *CreateEventStreamingRequestSinkSinkFnfParametersInput) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) *CreateEventStreamingRequestSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

type CreateEventStreamingRequestSinkSinkFnfParametersExecutionName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFnfParametersFlowName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersFlowName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersFlowName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFnfParametersInput struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The input information of the execution.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersInput) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersInput) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkFnfParametersRoleName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkFnfParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParameters struct {
	// The acknowledgement (ACK) mode.
	//
	// *   If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	// *   If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	// *   If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks *CreateEventStreamingRequestSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *CreateEventStreamingRequestSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The topic name.
	Topic *CreateEventStreamingRequestSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body.
	Value *CreateEventStreamingRequestSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetAcks(v *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetKey(v *CreateEventStreamingRequestSinkSinkKafkaParametersKey) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParameters) SetValue(v *CreateEventStreamingRequestSinkSinkKafkaParametersValue) *CreateEventStreamingRequestSinkSinkKafkaParameters {
	s.Value = v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParametersAcks struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ACK mode.
	//
	// *   If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	// *   If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	// *   If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersAcks) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersAcks) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance ID.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParametersKey struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersKey) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParametersTopic struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkKafkaParametersValue struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersValue) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetForm(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkKafkaParametersValue) SetValue(v string) *CreateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkMNSParameters struct {
	// The message body.
	Body *CreateEventStreamingRequestSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Specifies whether to enable Base64 encoding.
	IsBase64Encode *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The name of the MNS queue.
	QueueName *CreateEventStreamingRequestSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetBody(v *CreateEventStreamingRequestSinkSinkMNSParametersBody) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetIsBase64Encode(v *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParameters) SetQueueName(v *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) *CreateEventStreamingRequestSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

type CreateEventStreamingRequestSinkSinkMNSParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies that Base64 encoding is enabled.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkMNSParametersQueueName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetForm(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkMNSParametersQueueName) SetValue(v string) *CreateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParameters struct {
	AuthorizationType *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty" type:"Struct"`
	Data              *CreateEventStreamingRequestSinkSinkPrometheusParametersData              `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NetworkType       *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password          *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword          `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	SecurityGroupId   *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	URL               *CreateEventStreamingRequestSinkSinkPrometheusParametersURL               `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	Username          *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername          `json:"Username,omitempty" xml:"Username,omitempty" type:"Struct"`
	VSwitchId         *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Struct"`
	VpcId             *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId             `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetAuthorizationType(v *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.AuthorizationType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetData(v *CreateEventStreamingRequestSinkSinkPrometheusParametersData) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Data = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetNetworkType(v *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.NetworkType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetPassword(v *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Password = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetURL(v *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.URL = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetUsername(v *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Username = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetVSwitchId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VSwitchId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) *CreateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VpcId = v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersData struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersData) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersData) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersData) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersURL struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersURL) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersURL) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersURL) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersUsername struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	// The message body.
	Body *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange to which you want to deliver events. This parameter is available only if you set TargetType to Exchange.
	Exchange *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The information about the Message Queue for RabbitMQ instance.
	InstanceId *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The properties that you want to use to filter messages.
	Properties *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue to which you want to deliver events. This parameter is available only if you set TargetType to Queue.
	QueueName *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The rule that you want to use to route messages. This parameter is available only if you set TargetType to Exchange.
	RoutingKey *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The type of the resource to which you want to deliver events.
	TargetType *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetBody(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetExchange(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetMessageId(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetProperties(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetQueueName(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetRoutingKey(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetTargetType(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParameters) SetVirtualHostName(v *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) *CreateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange on the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The rule that you want to use to route messages.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which you want to deliver events. Valid values:
	//
	// *   Exchange
	// *   Queue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *CreateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParameters struct {
	// The message body.
	Body *CreateEventStreamingRequestSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The instance endpoint.
	InstanceEndpoint *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	// The parameters that are configured if you specify the event target as Message Queue for Apache RocketMQ.
	InstanceId *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The instance password.
	InstancePassword *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	// The instance type.
	InstanceType *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	// The instance username.
	InstanceUsername *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	// The keys that you want to use to filter messages.
	Keys *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// *   PublicNetwork
	// *   PrivateNetwork
	Network *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The properties that you want to use to filter messages.
	Properties *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The security group ID.
	SecurityGroupId *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	// The tags that you want to use to filter messages.
	Tags *CreateEventStreamingRequestSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic on the Message Queue for Apache RocketMQ instance.
	Topic *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The vSwitch ID.
	VSwitchIds *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The VPC ID.
	VpcId *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetBody(v *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceEndpoint(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceEndpoint = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstancePassword(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstancePassword = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceType(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceType = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceUsername(v *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceUsername = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetKeys(v *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetNetwork(v *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Network = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetProperties(v *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetSecurityGroupId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.SecurityGroupId = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetTags(v *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetVSwitchIds(v *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VSwitchIds = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParameters) SetVpcId(v *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) *CreateEventStreamingRequestSinkSinkRocketMQParameters {
	s.VpcId = v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance endpoint.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance password.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance type.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance username.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersKeys struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The network type. Valid values:
	//
	// *   PublicNetwork
	// *   PrivateNetwork
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersProperties struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The security group ID.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersTags struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTags) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTags) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersTopic struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vSwitch ID.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The VPC ID.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetForm(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId) SetValue(v string) *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParameters struct {
	// The message body that you want to deliver to Simple Log Service.
	Body *CreateEventStreamingRequestSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *CreateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *CreateEventStreamingRequestSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *CreateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The topic that you want to use to store logs. This parameter corresponds to the **topic** reserved field in Simple Log Service.
	Topic *CreateEventStreamingRequestSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetBody(v *CreateEventStreamingRequestSinkSinkSLSParametersBody) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetLogStore(v *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetProject(v *CreateEventStreamingRequestSinkSinkSLSParametersProject) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetRoleName(v *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParameters) SetTopic(v *CreateEventStreamingRequestSinkSinkSLSParametersTopic) *CreateEventStreamingRequestSinkSinkSLSParameters {
	s.Topic = v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParametersBody struct {
	// The format into which you want to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which you want to transform events.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersBody) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParametersLogStore struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Simple Log Service Logstore.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersLogStore) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersLogStore) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParametersProject struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Simple Log Service project.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersProject) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersProject) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParametersRoleName struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersRoleName) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSinkSinkSLSParametersTopic struct {
	// The format into which you want to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic that you want to use to store logs. This parameter corresponds to the **topic** reserved field in Simple Log Service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetForm(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetTemplate(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *CreateEventStreamingRequestSinkSinkSLSParametersTopic) SetValue(v string) *CreateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

type CreateEventStreamingRequestSource struct {
	// The parameters that are configured if you specify Data Transmission Service (DTS) as the event source.
	SourceDTSParameters *CreateEventStreamingRequestSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for Apache Kafka as the event source.
	SourceKafkaParameters *CreateEventStreamingRequestSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Service (MNS) as the event source.
	SourceMNSParameters *CreateEventStreamingRequestSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for MQTT as the event source.
	SourceMQTTParameters       *CreateEventStreamingRequestSourceSourceMQTTParameters       `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourcePrometheusParameters *CreateEventStreamingRequestSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for RabbitMQ as the event source.
	SourceRabbitMQParameters *CreateEventStreamingRequestSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for Apache RocketMQ as the event source.
	SourceRocketMQParameters *CreateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Log Service as the event source.
	SourceSLSParameters *CreateEventStreamingRequestSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s CreateEventStreamingRequestSource) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSource) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSource) SetSourceDTSParameters(v *CreateEventStreamingRequestSourceSourceDTSParameters) *CreateEventStreamingRequestSource {
	s.SourceDTSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceKafkaParameters(v *CreateEventStreamingRequestSourceSourceKafkaParameters) *CreateEventStreamingRequestSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceMNSParameters(v *CreateEventStreamingRequestSourceSourceMNSParameters) *CreateEventStreamingRequestSource {
	s.SourceMNSParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceMQTTParameters(v *CreateEventStreamingRequestSourceSourceMQTTParameters) *CreateEventStreamingRequestSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourcePrometheusParameters(v *CreateEventStreamingRequestSourceSourcePrometheusParameters) *CreateEventStreamingRequestSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRabbitMQParameters(v *CreateEventStreamingRequestSourceSourceRabbitMQParameters) *CreateEventStreamingRequestSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceRocketMQParameters(v *CreateEventStreamingRequestSourceSourceRocketMQParameters) *CreateEventStreamingRequestSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *CreateEventStreamingRequestSource) SetSourceSLSParameters(v *CreateEventStreamingRequestSourceSourceSLSParameters) *CreateEventStreamingRequestSource {
	s.SourceSLSParameters = v
	return s
}

type CreateEventStreamingRequestSourceSourceDTSParameters struct {
	// The URL and port number of the change tracking instance.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The UNIX timestamp that is generated when the SDK client consumes the first data record.
	InitCheckPoint *int64 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic of the change tracking instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The consumer group username.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceDTSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetBrokerUrl(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetInitCheckPoint(v int64) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetPassword(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetSid(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetTaskId(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceDTSParameters) SetUsername(v string) *CreateEventStreamingRequestSourceSourceDTSParameters {
	s.Username = &v
	return s
}

type CreateEventStreamingRequestSourceSourceKafkaParameters struct {
	// The group ID of the consumer that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network setting. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	VSwitchIds    *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetConsumerGroup(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetNetwork(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetOffsetReset(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetValueDataType(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type CreateEventStreamingRequestSourceSourceMNSParameters struct {
	// Specifies whether to enable Base64 encoding. Default value: true.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetIsBase64Decode(v bool) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetQueueName(v string) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMNSParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

type CreateEventStreamingRequestSourceSourceMQTTParameters struct {
	BodyDataType *string `json:"BodyDataType,omitempty" xml:"BodyDataType,omitempty"`
	// The ID of the Message Queue for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Message Queue for MQTT instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic from which messages are sent.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetBodyDataType(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.BodyDataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceMQTTParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

type CreateEventStreamingRequestSourceSourcePrometheusParameters struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourcePrometheusParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetClusterId(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetDataType(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourcePrometheusParameters) SetLabels(v string) *CreateEventStreamingRequestSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

type CreateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The queue name of the Message Queue for RabbitMQ instance.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides. You can call the [DescribeRegions](~~62010~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetQueueName(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type CreateEventStreamingRequestSourceSourceRocketMQParameters struct {
	// The authentication method.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The SQL statement that you want to use to filter messages.
	FilterSql *string `json:"FilterSql,omitempty" xml:"FilterSql,omitempty"`
	// The method that you want to use to filter messages.
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The instance endpoint.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the Message Queue for Apache RocketMQ instance. Valid values:
	//
	// *   PublicNetwork
	// *   PrivateNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The instance password.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The security group ID of the Message Queue for Apache RocketMQ instance.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance username.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The vSwitch ID of the Message Queue for Apache RocketMQ instance.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The VPC ID of the Message Queue for Apache RocketMQ instance.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The network type. Valid values: PublicNetwork and PrivateNetwork.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// *   CONSUME_FROM_LAST_OFFSET: Messages are consumed from the latest offset.
	// *   CONSUME_FROM_FIRST_OFFSET: Messages are consumed from the earliest offset.
	// *   CONSUME_FROM_TIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID of the cross-border task.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The tag that you want to use to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUME_FROM_TIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID of the cross-border task.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID of the cross-border task.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetAuthType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterSql(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterSql = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetFilterType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetGroupID(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstancePassword(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceType(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceUsername(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetNetwork(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Network = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetOffset(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetRegionId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetSecurityGroupId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTag(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTimestamp(v int64) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetTopic(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetVSwitchIds(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceRocketMQParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceRocketMQParameters {
	s.VpcId = &v
	return s
}

type CreateEventStreamingRequestSourceSourceSLSParameters struct {
	// The consumer offset. The value begin specifies the earliest offset, and the value end specifies the latest offset. You can also specify a time in seconds to start consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetConsumePosition(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetLogStore(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetProject(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *CreateEventStreamingRequestSourceSourceSLSParameters) SetRoleName(v string) *CreateEventStreamingRequestSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

type CreateEventStreamingRequestTransforms struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s CreateEventStreamingRequestTransforms) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestTransforms) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingRequestTransforms) SetArn(v string) *CreateEventStreamingRequestTransforms {
	s.Arn = &v
	return s
}

type CreateEventStreamingShrinkRequest struct {
	// The description of the event stream.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptionsShrink *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	// The event target. You must and can specify only one event target.
	SinkShrink *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	// The event provider, which is also known as the event source. You must and can specify only one event source.
	SourceShrink     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TransformsShrink *string `json:"Transforms,omitempty" xml:"Transforms,omitempty"`
}

func (s CreateEventStreamingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingShrinkRequest) SetDescription(v string) *CreateEventStreamingShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetEventStreamingName(v string) *CreateEventStreamingShrinkRequest {
	s.EventStreamingName = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetFilterPattern(v string) *CreateEventStreamingShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetRunOptionsShrink(v string) *CreateEventStreamingShrinkRequest {
	s.RunOptionsShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetSinkShrink(v string) *CreateEventStreamingShrinkRequest {
	s.SinkShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetSourceShrink(v string) *CreateEventStreamingShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *CreateEventStreamingShrinkRequest) SetTransformsShrink(v string) *CreateEventStreamingShrinkRequest {
	s.TransformsShrink = &v
	return s
}

type CreateEventStreamingResponseBody struct {
	// The response code. Valid values:
	//
	// *   Success: The request is successful.
	// *   Other codes: The request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponseBody) SetCode(v string) *CreateEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetData(v *CreateEventStreamingResponseBodyData) *CreateEventStreamingResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventStreamingResponseBody) SetMessage(v string) *CreateEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetRequestId(v string) *CreateEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetSuccess(v bool) *CreateEventStreamingResponseBody {
	s.Success = &v
	return s
}

type CreateEventStreamingResponseBodyData struct {
	// The ARN of the event stream.
	EventStreamingARN *string `json:"EventStreamingARN,omitempty" xml:"EventStreamingARN,omitempty"`
}

func (s CreateEventStreamingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponseBodyData) SetEventStreamingARN(v string) *CreateEventStreamingResponseBodyData {
	s.EventStreamingARN = &v
	return s
}

type CreateEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponse) SetHeaders(v map[string]*string) *CreateEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *CreateEventStreamingResponse) SetStatusCode(v int32) *CreateEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventStreamingResponse) SetBody(v *CreateEventStreamingResponseBody) *CreateEventStreamingResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event targets.
	EventTargets []*CreateRuleRequestEventTargets `json:"EventTargets,omitempty" xml:"EventTargets,omitempty" type:"Repeated"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: enables the event rule. It is the default status of the event rule. DISABLE: disables the event rule.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleRequest) SetEventBusName(v string) *CreateRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateRuleRequest) SetEventTargets(v []*CreateRuleRequestEventTargets) *CreateRuleRequest {
	s.EventTargets = v
	return s
}

func (s *CreateRuleRequest) SetFilterPattern(v string) *CreateRuleRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleRequest) SetStatus(v string) *CreateRuleRequest {
	s.Status = &v
	return s
}

type CreateRuleRequestEventTargets struct {
	// The dead-letter queue. Events that are not processed or whose maximum retries are exceeded are written to the dead-letter queue. The dead-letter queue feature is supported by the following queue types: Message Queue for Apache RocketMQ, Message Service (MNS), Message Queue for Apache Kafka, and EventBridge.
	DeadLetterQueue *CreateRuleRequestEventTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint of the event target.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values: ALL: allows fault tolerance. If an error occurs, the event processing is not blocked. If the message fails to be sent after the maximum number of retries specified by the retry policy is reached, the message is delivered to the dead-letter queue or discarded based on your configurations. NONE: does not allow fault tolerance. If an error occurs and the message fails to be sent after the maximum number of retries specified by the retry policy is reached, the event processing is blocked.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the custom event target.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*CreateRuleRequestEventTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy that is used to push events. Valid values: BACKOFF_RETRY: backoff retry. If an event failed to be pushed, it can be retried up to three times. The interval between two consecutive retries is a random value between 10 and 20 seconds. EXPONENTIAL_DECAY_RETRY: exponential decay retry. If an event failed to be pushed, it can be retried up to 176 times. The interval between two consecutive retries exponentially increases to 512 seconds, and the total retry time is one day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 512, ..., and 512 seconds. The interval of 512 seconds is used for 167 retries.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The type of the event target. For more information, see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestEventTargets) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestEventTargets) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargets) SetDeadLetterQueue(v *CreateRuleRequestEventTargetsDeadLetterQueue) *CreateRuleRequestEventTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *CreateRuleRequestEventTargets) SetEndpoint(v string) *CreateRuleRequestEventTargets {
	s.Endpoint = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetErrorsTolerance(v string) *CreateRuleRequestEventTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetId(v string) *CreateRuleRequestEventTargets {
	s.Id = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetParamList(v []*CreateRuleRequestEventTargetsParamList) *CreateRuleRequestEventTargets {
	s.ParamList = v
	return s
}

func (s *CreateRuleRequestEventTargets) SetPushRetryStrategy(v string) *CreateRuleRequestEventTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *CreateRuleRequestEventTargets) SetType(v string) *CreateRuleRequestEventTargets {
	s.Type = &v
	return s
}

type CreateRuleRequestEventTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue. Events that are not processed or whose maximum retries are exceeded are written to the dead-letter queue. The ARN feature is supported by the following queue types: MNS and Message Queue for Apache RocketMQ.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s CreateRuleRequestEventTargetsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestEventTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargetsDeadLetterQueue) SetArn(v string) *CreateRuleRequestEventTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

type CreateRuleRequestEventTargetsParamList struct {
	// The format that is used by the event target parameter. For more information, see [Limits.](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target. For more information, see [Limits](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template that is used by the event target parameter.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestEventTargetsParamList) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestEventTargetsParamList) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestEventTargetsParamList) SetForm(v string) *CreateRuleRequestEventTargetsParamList {
	s.Form = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetResourceKey(v string) *CreateRuleRequestEventTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetTemplate(v string) *CreateRuleRequestEventTargetsParamList {
	s.Template = &v
	return s
}

func (s *CreateRuleRequestEventTargetsParamList) SetValue(v string) *CreateRuleRequestEventTargetsParamList {
	s.Value = &v
	return s
}

type CreateRuleShrinkRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event targets.
	EventTargetsShrink *string `json:"EventTargets,omitempty" xml:"EventTargets,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: enables the event rule. It is the default status of the event rule. DISABLE: disables the event rule.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleShrinkRequest) SetDescription(v string) *CreateRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetEventBusName(v string) *CreateRuleShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetEventTargetsShrink(v string) *CreateRuleShrinkRequest {
	s.EventTargetsShrink = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetFilterPattern(v string) *CreateRuleShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetRuleName(v string) *CreateRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRuleShrinkRequest) SetStatus(v string) *CreateRuleShrinkRequest {
	s.Status = &v
	return s
}

type CreateRuleResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetCode(v string) *CreateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRuleResponseBody) SetData(v *CreateRuleResponseBodyData) *CreateRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateRuleResponseBody) SetMessage(v string) *CreateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetSuccess(v bool) *CreateRuleResponseBody {
	s.Success = &v
	return s
}

type CreateRuleResponseBodyData struct {
	// The ARN of the event rule. The ARN is used for authorization.
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
}

func (s CreateRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBodyData) SetRuleARN(v string) *CreateRuleResponseBodyData {
	s.RuleARN = &v
	return s
}

type CreateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleForProductRequest struct {
	// The name of the cloud service or the name of the service-linked role with which the cloud service is associated.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s CreateServiceLinkedRoleForProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductRequest) SetProductName(v string) *CreateServiceLinkedRoleForProductRequest {
	s.ProductName = &v
	return s
}

type CreateServiceLinkedRoleForProductResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceLinkedRoleForProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetCode(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetMessage(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleForProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponseBody) SetSuccess(v bool) *CreateServiceLinkedRoleForProductResponseBody {
	s.Success = &v
	return s
}

type CreateServiceLinkedRoleForProductResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleForProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleForProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForProductResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleForProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponse) SetBody(v *CreateServiceLinkedRoleForProductResponseBody) *CreateServiceLinkedRoleForProductResponse {
	s.Body = v
	return s
}

type DeleteApiDestinationRequest struct {
	// The name of the API destination.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s DeleteApiDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationRequest) SetApiDestinationName(v string) *DeleteApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

type DeleteApiDestinationResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationResponseBody) SetCode(v string) *DeleteApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApiDestinationResponseBody) SetMessage(v string) *DeleteApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApiDestinationResponseBody) SetRequestId(v string) *DeleteApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiDestinationResponse) SetHeaders(v map[string]*string) *DeleteApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiDestinationResponse) SetStatusCode(v int32) *DeleteApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiDestinationResponse) SetBody(v *DeleteApiDestinationResponseBody) *DeleteApiDestinationResponse {
	s.Body = v
	return s
}

type DeleteConnectionRequest struct {
	// The name of the connection that you want to delete.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s DeleteConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) SetConnectionName(v string) *DeleteConnectionRequest {
	s.ConnectionName = &v
	return s
}

type DeleteConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponseBody) SetCode(v string) *DeleteConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetMessage(v string) *DeleteConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetRequestId(v string) *DeleteConnectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponse) SetHeaders(v map[string]*string) *DeleteConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectionResponse) SetStatusCode(v int32) *DeleteConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectionResponse) SetBody(v *DeleteConnectionResponseBody) *DeleteConnectionResponse {
	s.Body = v
	return s
}

type DeleteEventBusRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s DeleteEventBusRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBusRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventBusRequest) SetEventBusName(v string) *DeleteEventBusRequest {
	s.EventBusName = &v
	return s
}

type DeleteEventBusResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventBusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventBusResponseBody) SetCode(v string) *DeleteEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetMessage(v string) *DeleteEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetRequestId(v string) *DeleteEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetSuccess(v bool) *DeleteEventBusResponseBody {
	s.Success = &v
	return s
}

type DeleteEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventBusResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBusResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventBusResponse) SetHeaders(v map[string]*string) *DeleteEventBusResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventBusResponse) SetStatusCode(v int32) *DeleteEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventBusResponse) SetBody(v *DeleteEventBusResponseBody) *DeleteEventBusResponse {
	s.Body = v
	return s
}

type DeleteEventSourceRequest struct {
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
}

func (s DeleteEventSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceRequest) SetEventSourceName(v string) *DeleteEventSourceRequest {
	s.EventSourceName = &v
	return s
}

type DeleteEventSourceResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceResponseBody) SetCode(v string) *DeleteEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetMessage(v string) *DeleteEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetRequestId(v string) *DeleteEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetSuccess(v bool) *DeleteEventSourceResponseBody {
	s.Success = &v
	return s
}

type DeleteEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceResponse) SetHeaders(v map[string]*string) *DeleteEventSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventSourceResponse) SetStatusCode(v int32) *DeleteEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventSourceResponse) SetBody(v *DeleteEventSourceResponseBody) *DeleteEventSourceResponse {
	s.Body = v
	return s
}

type DeleteEventStreamingRequest struct {
	// The name of the event stream that you want to delete.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s DeleteEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingRequest) SetEventStreamingName(v string) *DeleteEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

type DeleteEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *bool `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingResponseBody) SetCode(v bool) *DeleteEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetMessage(v string) *DeleteEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetRequestId(v string) *DeleteEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetSuccess(v bool) *DeleteEventStreamingResponseBody {
	s.Success = &v
	return s
}

type DeleteEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingResponse) SetHeaders(v map[string]*string) *DeleteEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventStreamingResponse) SetStatusCode(v int32) *DeleteEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventStreamingResponse) SetBody(v *DeleteEventStreamingResponseBody) *DeleteEventStreamingResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule that you want to delete.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetEventBusName(v string) *DeleteRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleName(v string) *DeleteRuleRequest {
	s.RuleName = &v
	return s
}

type DeleteRuleResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetSuccess(v bool) *DeleteRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteTargetsRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The IDs of the event targets that you want to delete.
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
}

func (s DeleteTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTargetsRequest) SetEventBusName(v string) *DeleteTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteTargetsRequest) SetRuleName(v string) *DeleteTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteTargetsRequest) SetTargetIds(v []*string) *DeleteTargetsRequest {
	s.TargetIds = v
	return s
}

type DeleteTargetsShrinkRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The IDs of the event targets that you want to delete.
	TargetIdsShrink *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
}

func (s DeleteTargetsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTargetsShrinkRequest) SetEventBusName(v string) *DeleteTargetsShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *DeleteTargetsShrinkRequest) SetRuleName(v string) *DeleteTargetsShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteTargetsShrinkRequest) SetTargetIdsShrink(v string) *DeleteTargetsShrinkRequest {
	s.TargetIdsShrink = &v
	return s
}

type DeleteTargetsResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBody) SetCode(v string) *DeleteTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetData(v *DeleteTargetsResponseBodyData) *DeleteTargetsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTargetsResponseBody) SetMessage(v string) *DeleteTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetRequestId(v string) *DeleteTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetSuccess(v bool) *DeleteTargetsResponseBody {
	s.Success = &v
	return s
}

type DeleteTargetsResponseBodyData struct {
	// The information about the event body that failed to be processed.
	ErrorEntries []*DeleteTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	// The number of event bodies that failed to be processed. Valid values: 0: No event bodies failed to be processed. An integer other than 0: the number of event bodies that failed to be processed.
	ErrorEntriesCount *int32 `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
}

func (s DeleteTargetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBodyData) SetErrorEntries(v []*DeleteTargetsResponseBodyDataErrorEntries) *DeleteTargetsResponseBodyData {
	s.ErrorEntries = v
	return s
}

func (s *DeleteTargetsResponseBodyData) SetErrorEntriesCount(v int32) *DeleteTargetsResponseBodyData {
	s.ErrorEntriesCount = &v
	return s
}

type DeleteTargetsResponseBodyDataErrorEntries struct {
	// The ID of the event body that failed to be processed.
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// The returned error code.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned error message.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DeleteTargetsResponseBodyDataErrorEntries) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsResponseBodyDataErrorEntries) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetEntryId(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.EntryId = &v
	return s
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetErrorCode(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetErrorMessage(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.ErrorMessage = &v
	return s
}

type DeleteTargetsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponse) SetHeaders(v map[string]*string) *DeleteTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTargetsResponse) SetStatusCode(v int32) *DeleteTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTargetsResponse) SetBody(v *DeleteTargetsResponseBody) *DeleteTargetsResponse {
	s.Body = v
	return s
}

type DisableRuleRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DisableRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableRuleRequest) SetEventBusName(v string) *DisableRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *DisableRuleRequest) SetRuleName(v string) *DisableRuleRequest {
	s.RuleName = &v
	return s
}

type DisableRuleResponseBody struct {
	// The error code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRuleResponseBody) SetCode(v string) *DisableRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableRuleResponseBody) SetMessage(v string) *DisableRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRuleResponseBody) SetRequestId(v string) *DisableRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRuleResponseBody) SetSuccess(v bool) *DisableRuleResponseBody {
	s.Success = &v
	return s
}

type DisableRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableRuleResponse) SetHeaders(v map[string]*string) *DisableRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableRuleResponse) SetStatusCode(v int32) *DisableRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRuleResponse) SetBody(v *DisableRuleResponseBody) *DisableRuleResponse {
	s.Body = v
	return s
}

type EnableRuleRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s EnableRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableRuleRequest) SetEventBusName(v string) *EnableRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *EnableRuleRequest) SetRuleName(v string) *EnableRuleRequest {
	s.RuleName = &v
	return s
}

type EnableRuleResponseBody struct {
	// The error code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRuleResponseBody) SetCode(v string) *EnableRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableRuleResponseBody) SetMessage(v string) *EnableRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableRuleResponseBody) SetRequestId(v string) *EnableRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableRuleResponseBody) SetSuccess(v bool) *EnableRuleResponseBody {
	s.Success = &v
	return s
}

type EnableRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableRuleResponse) SetHeaders(v map[string]*string) *EnableRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableRuleResponse) SetStatusCode(v int32) *EnableRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRuleResponse) SetBody(v *EnableRuleResponseBody) *EnableRuleResponse {
	s.Body = v
	return s
}

type GetApiDestinationRequest struct {
	// The name of the API destination.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
}

func (s GetApiDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *GetApiDestinationRequest) SetApiDestinationName(v string) *GetApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

type GetApiDestinationResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetApiDestinationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApiDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBody) SetCode(v string) *GetApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiDestinationResponseBody) SetData(v *GetApiDestinationResponseBodyData) *GetApiDestinationResponseBody {
	s.Data = v
	return s
}

func (s *GetApiDestinationResponseBody) SetMessage(v string) *GetApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiDestinationResponseBody) SetRequestId(v string) *GetApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

type GetApiDestinationResponseBodyData struct {
	// The name of the API destination.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the API destination was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The request parameters that are configured for the API destination.
	HttpApiParameters *GetApiDestinationResponseBodyDataHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s GetApiDestinationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApiDestinationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBodyData) SetApiDestinationName(v string) *GetApiDestinationResponseBodyData {
	s.ApiDestinationName = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetConnectionName(v string) *GetApiDestinationResponseBodyData {
	s.ConnectionName = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetDescription(v string) *GetApiDestinationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetGmtCreate(v int64) *GetApiDestinationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetHttpApiParameters(v *GetApiDestinationResponseBodyDataHttpApiParameters) *GetApiDestinationResponseBodyData {
	s.HttpApiParameters = v
	return s
}

type GetApiDestinationResponseBodyDataHttpApiParameters struct {
	// The endpoint of the API destination.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// *   POST
	// *   GET
	// *   DELETE
	// *   PUT
	// *   HEAD
	// *   TRACE
	// *   PATCH
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s GetApiDestinationResponseBodyDataHttpApiParameters) String() string {
	return tea.Prettify(s)
}

func (s GetApiDestinationResponseBodyDataHttpApiParameters) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) SetEndpoint(v string) *GetApiDestinationResponseBodyDataHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) SetMethod(v string) *GetApiDestinationResponseBodyDataHttpApiParameters {
	s.Method = &v
	return s
}

type GetApiDestinationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponse) SetHeaders(v map[string]*string) *GetApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *GetApiDestinationResponse) SetStatusCode(v int32) *GetApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiDestinationResponse) SetBody(v *GetApiDestinationResponseBody) *GetApiDestinationResponse {
	s.Body = v
	return s
}

type GetConnectionRequest struct {
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
}

func (s GetConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionRequest) SetConnectionName(v string) *GetConnectionRequest {
	s.ConnectionName = &v
	return s
}

type GetConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBody) SetCode(v string) *GetConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnectionResponseBody) SetData(v *GetConnectionResponseBodyData) *GetConnectionResponseBody {
	s.Data = v
	return s
}

func (s *GetConnectionResponseBody) SetHttpCode(v int32) *GetConnectionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetConnectionResponseBody) SetMessage(v string) *GetConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectionResponseBody) SetRequestId(v string) *GetConnectionResponseBody {
	s.RequestId = &v
	return s
}

type GetConnectionResponseBodyData struct {
	// The queried connections.
	Connections []*GetConnectionResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
}

func (s GetConnectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyData) SetConnections(v []*GetConnectionResponseBodyDataConnections) *GetConnectionResponseBodyData {
	s.Connections = v
	return s
}

type GetConnectionResponseBodyDataConnections struct {
	// The authentication methods.
	AuthParameters *GetConnectionResponseBodyDataConnectionsAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The connection description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the connection was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The data source ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the network.
	NetworkParameters *GetConnectionResponseBodyDataConnectionsNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnections) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnections) SetAuthParameters(v *GetConnectionResponseBodyDataConnectionsAuthParameters) *GetConnectionResponseBodyDataConnections {
	s.AuthParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnections) SetConnectionName(v string) *GetConnectionResponseBodyDataConnections {
	s.ConnectionName = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnections) SetDescription(v string) *GetConnectionResponseBodyDataConnections {
	s.Description = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnections) SetGmtCreate(v int64) *GetConnectionResponseBodyDataConnections {
	s.GmtCreate = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnections) SetId(v int64) *GetConnectionResponseBodyDataConnections {
	s.Id = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnections) SetNetworkParameters(v *GetConnectionResponseBodyDataConnectionsNetworkParameters) *GetConnectionResponseBodyDataConnections {
	s.NetworkParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParameters struct {
	// The information about API key authentication.
	ApiKeyAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication method. Valid values:
	//
	// *   BASIC_AUTH: basic authentication.
	// *   API_KEY_AUTH: API key authentication.
	// *   OAUTH_AUTH: OAuth authentication.
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The information about basic authentication.
	BasicAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The information about OAuth authentication.
	OAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetApiKeyAuthParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.ApiKeyAuthParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetAuthorizationType(v string) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.AuthorizationType = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetBasicAuthParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.BasicAuthParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetOAuthParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.OAuthParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters struct {
	// The password of basic authentication.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username of basic authentication.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetPassword(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetUsername(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The information about the client.
	ClientParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// *   GET
	// *   POST
	// *   HEAD
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters of OAuth authentication.
	OAuthHttpParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) SetAuthorizationEndpoint(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) SetClientParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.ClientParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) SetHttpMethod(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.HttpMethod = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) SetOAuthHttpParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.OAuthHttpParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The AccessKey secret of the client.
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientID(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The information about the request body.
	BodyParameters []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The information about the request header.
	HeaderParameters []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The information about the request path.
	QueryStringParameters []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetBodyParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetHeaderParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetQueryStringParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.QueryStringParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request body.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetIsValueSecret(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.IsValueSecret = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Value = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request header.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetIsValueSecret(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.IsValueSecret = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request path.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetIsValueSecret(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.IsValueSecret = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsNetworkParameters struct {
	// *   PublicNetwork: the Internet.
	// *   PrivateNetwork: virtual private cloud (VPC).
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsNetworkParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsNetworkParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) SetNetworkType(v string) *GetConnectionResponseBodyDataConnectionsNetworkParameters {
	s.NetworkType = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) SetSecurityGroupId(v string) *GetConnectionResponseBodyDataConnectionsNetworkParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) SetVpcId(v string) *GetConnectionResponseBodyDataConnectionsNetworkParameters {
	s.VpcId = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) SetVswitcheId(v string) *GetConnectionResponseBodyDataConnectionsNetworkParameters {
	s.VswitcheId = &v
	return s
}

type GetConnectionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionResponse) SetHeaders(v map[string]*string) *GetConnectionResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionResponse) SetStatusCode(v int32) *GetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionResponse) SetBody(v *GetConnectionResponseBody) *GetConnectionResponse {
	s.Body = v
	return s
}

type GetEventBusRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s GetEventBusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventBusRequest) GoString() string {
	return s.String()
}

func (s *GetEventBusRequest) SetEventBusName(v string) *GetEventBusRequest {
	s.EventBusName = &v
	return s
}

type GetEventBusResponseBody struct {
	// The response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventBusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBody) SetCode(v string) *GetEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventBusResponseBody) SetData(v *GetEventBusResponseBodyData) *GetEventBusResponseBody {
	s.Data = v
	return s
}

func (s *GetEventBusResponseBody) SetMessage(v string) *GetEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventBusResponseBody) SetRequestId(v string) *GetEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventBusResponseBody) SetSuccess(v bool) *GetEventBusResponseBody {
	s.Success = &v
	return s
}

type GetEventBusResponseBodyData struct {
	// The timestamp that indicates when the event bus was created.
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s GetEventBusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEventBusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBodyData) SetCreateTimestamp(v int64) *GetEventBusResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetDescription(v string) *GetEventBusResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusARN(v string) *GetEventBusResponseBodyData {
	s.EventBusARN = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusName(v string) *GetEventBusResponseBodyData {
	s.EventBusName = &v
	return s
}

type GetEventBusResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventBusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventBusResponse) GoString() string {
	return s.String()
}

func (s *GetEventBusResponse) SetHeaders(v map[string]*string) *GetEventBusResponse {
	s.Headers = v
	return s
}

func (s *GetEventBusResponse) SetStatusCode(v int32) *GetEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventBusResponse) SetBody(v *GetEventBusResponseBody) *GetEventBusResponse {
	s.Body = v
	return s
}

type GetEventStreamingRequest struct {
	// The name of the event stream whose details you want to query.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s GetEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *GetEventStreamingRequest) SetEventStreamingName(v string) *GetEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

type GetEventStreamingResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For a list of error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBody) SetCode(v string) *GetEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetData(v *GetEventStreamingResponseBodyData) *GetEventStreamingResponseBody {
	s.Data = v
	return s
}

func (s *GetEventStreamingResponseBody) SetMessage(v string) *GetEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetRequestId(v string) *GetEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventStreamingResponseBody) SetSuccess(v bool) *GetEventStreamingResponseBody {
	s.Success = &v
	return s
}

type GetEventStreamingResponseBodyData struct {
	// The description of the event stream that is returned.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream that is returned.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptions *GetEventStreamingResponseBodyDataRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target.
	Sink *GetEventStreamingResponseBodyDataSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source.
	Source *GetEventStreamingResponseBodyDataSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The status of the event stream that is returned.
	Status     *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Transforms []*GetEventStreamingResponseBodyDataTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s GetEventStreamingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyData) SetDescription(v string) *GetEventStreamingResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetEventStreamingName(v string) *GetEventStreamingResponseBodyData {
	s.EventStreamingName = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetFilterPattern(v string) *GetEventStreamingResponseBodyData {
	s.FilterPattern = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetRunOptions(v *GetEventStreamingResponseBodyDataRunOptions) *GetEventStreamingResponseBodyData {
	s.RunOptions = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetSink(v *GetEventStreamingResponseBodyDataSink) *GetEventStreamingResponseBodyData {
	s.Sink = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetSource(v *GetEventStreamingResponseBodyDataSource) *GetEventStreamingResponseBodyData {
	s.Source = v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetStatus(v string) *GetEventStreamingResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEventStreamingResponseBodyData) SetTransforms(v []*GetEventStreamingResponseBodyDataTransforms) *GetEventStreamingResponseBodyData {
	s.Transforms = v
	return s
}

type GetEventStreamingResponseBodyDataRunOptions struct {
	// The batch window.
	BatchWindow *GetEventStreamingResponseBodyDataRunOptionsBatchWindow `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	// Indicates whether dead-letter queues are enabled. By default, dead-letter queues are disabled. Messages that fail to be pushed after allowed retries as specified by the retry policy are discarded.
	DeadLetterQueue *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The fault tolerance policy. The value NONE specifies that faults are not tolerated, and the value All specifies that all faults are tolerated.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The concurrency level.
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The information about the retry policy that is used if the event fails to be pushed.
	RetryStrategy *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataRunOptions) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptions) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetBatchWindow(v *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) *GetEventStreamingResponseBodyDataRunOptions {
	s.BatchWindow = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetDeadLetterQueue(v *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) *GetEventStreamingResponseBodyDataRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetErrorsTolerance(v string) *GetEventStreamingResponseBodyDataRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetMaximumTasks(v int32) *GetEventStreamingResponseBodyDataRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptions) SetRetryStrategy(v *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) *GetEventStreamingResponseBodyDataRunOptions {
	s.RetryStrategy = v
	return s
}

type GetEventStreamingResponseBodyDataRunOptionsBatchWindow struct {
	// The maximum number of events that are allowed in the batch window. If this threshold is reached, data in the window is pushed downstream. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. If this threshold is reached, data in the window is pushed downstream. When multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsBatchWindow) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) SetCountBasedWindow(v int32) *GetEventStreamingResponseBodyDataRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *GetEventStreamingResponseBodyDataRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

type GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue) SetArn(v string) *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

type GetEventStreamingResponseBodyDataRunOptionsRetryStrategy struct {
	// The maximum period of time during which retries are performed.
	MaximumEventAgeInSeconds *float32 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	MaximumRetryAttempts *float32 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry policy. Valid values: BACKOFFRETRY and EXPONENTIALDECAY_RETRY.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v float32) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetMaximumRetryAttempts(v float32) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

type GetEventStreamingResponseBodyDataSink struct {
	// The parameters that are returned if the event target is Function Compute.
	SinkFcParameters *GetEventStreamingResponseBodyDataSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The Sink Fnf parameters.
	SinkFnfParameters *GetEventStreamingResponseBodyDataSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Message Queue for Apache Kafka.
	SinkKafkaParameters *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Message Service (MNS).
	SinkMNSParameters *GetEventStreamingResponseBodyDataSinkSinkMNSParameters `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event target is Message Queue for RabbitMQ.
	SinkRabbitMQParameters *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Parameters
	SinkRocketMQParameters *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// Sink SLS Parameters
	SinkSLSParameters *GetEventStreamingResponseBodyDataSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSink) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSink) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkFcParameters(v *GetEventStreamingResponseBodyDataSinkSinkFcParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkFcParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkFnfParameters(v *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkFnfParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkKafkaParameters(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkMNSParameters(v *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkMNSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkRabbitMQParameters(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkRocketMQParameters(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSink) SetSinkSLSParameters(v *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) *GetEventStreamingResponseBodyDataSink {
	s.SinkSLSParameters = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParameters struct {
	// The message body that is sent to the function.
	Body *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The function name.
	FunctionName *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation type. Valid values: Sync: synchronous Async: asynchronous
	InvocationType *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetConcurrency(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetFunctionName(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetInvocationType(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetQualifier(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParameters) SetServiceName(v *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) *GetEventStreamingResponseBodyDataSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersBody struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The function name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation type.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The alias of the service to which the function belongs.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The execution input information.
	Input *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetExecutionName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetFlowName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetInput(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParameters) SetRoleName(v *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) *GetEventStreamingResponseBodyDataSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution input information.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role configuration.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParameters struct {
	// The acknowledgment information.
	Acks *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The target service type is Message Queue for Apache Kafka.
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The topic name.
	Topic *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message content.
	Value *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetAcks(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetKey(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters) SetValue(v *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters {
	s.Value = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The acknowledgment information.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The instance ID.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Indicates whether Base64 encoding is enabled.
	IsBase64Encode *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The target service type is MNS.
	QueueName *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetIsBase64Encode(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParameters) SetQueueName(v *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) *GetEventStreamingResponseBodyDataSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies that Base64 encoding is enabled.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange mode. This parameter is available only if TargetType is set to Exchange.
	Exchange *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The target service type is Message Queue for RabbitMQ instance.
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Properties *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue mode. This parameter is available only if TargetType is set to Queue.
	QueueName *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The routing rule for the message. This parameter is available only if TargetType is set to Exchange.
	RoutingKey *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The target type.
	TargetType *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetExchange(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetMessageId(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetProperties(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetQueueName(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetRoutingKey(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetTargetType(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters) SetVirtualHostName(v *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange in the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue in the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The routing rule for the message.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which the event is delivered. Valid values: Exchange: exchanges. Queue: queues.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The target service type is Message Queue for Apache RocketMQ.
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Keys *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Properties *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Tags *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic in the Message Queue for Apache RocketMQ instance.
	Topic *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetInstanceId(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetKeys(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetProperties(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetTags(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParameters struct {
	// The message content.
	Body *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The Simple Log Service Logstore.
	LogStore *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Simple Log Service project.
	Project *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Simple Log Service.
	Topic *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetBody(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetLogStore(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetProject(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetRoleName(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParameters) SetTopic(v *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) *GetEventStreamingResponseBodyDataSinkSinkSLSParameters {
	s.Topic = v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody struct {
	// The method that is used to transform the event.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which the event is transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before the transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service Logstore.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service project.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic struct {
	// The method that is used to transform the event. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template style.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetForm(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetTemplate(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic) SetValue(v string) *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

type GetEventStreamingResponseBodyDataSource struct {
	// The parameters that are returned if the event source is Data Transmission Service (DTS).
	SourceDTSParameters *GetEventStreamingResponseBodyDataSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// Source Kafka Parameters
	SourceKafkaParameters *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// Source MNS Parameters
	SourceMNSParameters *GetEventStreamingResponseBodyDataSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event source is Message Queue for MQTT.
	SourceMQTTParameters       *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters       `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourcePrometheusParameters *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// Source RabbitMQ Parameters
	SourceRabbitMQParameters *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// Source RocketMQ Parameters
	SourceRocketMQParameters *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if the event provider is Simple Log Service.
	SourceSLSParameters *GetEventStreamingResponseBodyDataSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s GetEventStreamingResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceDTSParameters(v *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceDTSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceKafkaParameters(v *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceMNSParameters(v *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceMNSParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceMQTTParameters(v *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourcePrometheusParameters(v *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceRabbitMQParameters(v *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceRocketMQParameters(v *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *GetEventStreamingResponseBodyDataSource) SetSourceSLSParameters(v *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) *GetEventStreamingResponseBodyDataSource {
	s.SourceSLSParameters = v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceDTSParameters struct {
	// The URL and port number of the data subscription channel.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The consumer offset. A consumer offset is a timestamp that indicates when the SDK client consumes the first data record. The value is a UNIX timestamp.
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The password of the consumer group.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the consumer group.
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The topic to which you want to subscribe by using the data subscription channel.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The account of the consumer group.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceDTSParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetBrokerUrl(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetInitCheckPoint(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetPassword(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetSid(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetTaskId(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceDTSParameters) SetUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceDTSParameters {
	s.Username = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The region ID of the Message Queue for Apache Kafka instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	VSwitchIds    *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetConsumerGroup(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetNetwork(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetOffsetReset(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetValueDataType(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceMNSParameters struct {
	// Indicates whether Base64 encoding is enabled.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID of the MNS queue.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetIsBase64Decode(v bool) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMNSParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceMQTTParameters struct {
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Message Queue for MQTT instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic in the Message Queue for MQTT instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetClusterId(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetDataType(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters) SetLabels(v string) *GetEventStreamingResponseBodyDataSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue in the Message Queue for RabbitMQ instance.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID of the Message Queue for RabbitMQ instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetQueueName(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters struct {
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group in the Message Queue for Apache RocketMQ instance.
	GroupID          *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceNetwork         *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	InstancePassword        *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceUsername        *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	InstanceVSwitchIds      *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	InstanceVpcId           *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The consumer offset of messages. Valid values: CONSUME_FROM_LAST_OFFSET: Start consumption from the latest offset. CONSUME_FROM_FIRST_OFFSET: Start consumption from the earliest offset. CONSUME_FROM_TIMESTAMP: Start consumption from the offset at the specified point in time.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region ID of the Message Queue for Apache RocketMQ instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags that are used to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp of the offset from which consumption starts. This parameter is valid only if you set the Offset parameter to CONSUME_FROM_TIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic to which the message belongs.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetAuthType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetGroupID(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstancePassword(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceType(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceUsername(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetOffset(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetRegionId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTag(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTimestamp(v int64) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetTopic(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceSLSParameters struct {
	// The starting consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The consumer group.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetConsumePosition(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetConsumerGroup(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetLogStore(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetProject(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceSLSParameters) SetRoleName(v string) *GetEventStreamingResponseBodyDataSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

type GetEventStreamingResponseBodyDataTransforms struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s GetEventStreamingResponseBodyDataTransforms) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataTransforms) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataTransforms) SetArn(v string) *GetEventStreamingResponseBodyDataTransforms {
	s.Arn = &v
	return s
}

type GetEventStreamingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponse) SetHeaders(v map[string]*string) *GetEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *GetEventStreamingResponse) SetStatusCode(v int32) *GetEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventStreamingResponse) SetBody(v *GetEventStreamingResponseBody) *GetEventStreamingResponse {
	s.Body = v
	return s
}

type GetRuleRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) SetEventBusName(v string) *GetRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *GetRuleRequest) SetRuleName(v string) *GetRuleRequest {
	s.RuleName = &v
	return s
}

type GetRuleResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetSuccess(v bool) *GetRuleResponseBody {
	s.Success = &v
	return s
}

type GetRuleResponseBodyData struct {
	// The timestamp that indicates when the event rule was created.
	CreatedTimestamp *int64 `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	// The description of the event rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual and stringExpression. You can specify up to five expressions in the map data structure in each field.
	//
	// You can specify up to five expressions in the map data structure in each field.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event rule.
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE (default): The event rule is enabled. DISABLE: The event rule is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event targets.
	Targets []*GetRuleResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) SetCreatedTimestamp(v int64) *GetRuleResponseBodyData {
	s.CreatedTimestamp = &v
	return s
}

func (s *GetRuleResponseBodyData) SetDescription(v string) *GetRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRuleResponseBodyData) SetEventBusName(v string) *GetRuleResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *GetRuleResponseBodyData) SetFilterPattern(v string) *GetRuleResponseBodyData {
	s.FilterPattern = &v
	return s
}

func (s *GetRuleResponseBodyData) SetRuleARN(v string) *GetRuleResponseBodyData {
	s.RuleARN = &v
	return s
}

func (s *GetRuleResponseBodyData) SetRuleName(v string) *GetRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetRuleResponseBodyData) SetStatus(v string) *GetRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRuleResponseBodyData) SetTargets(v []*GetRuleResponseBodyDataTargets) *GetRuleResponseBodyData {
	s.Targets = v
	return s
}

type GetRuleResponseBodyDataTargets struct {
	// The ID of the custom event target.
	DeadLetterQueue *GetRuleResponseBodyDataTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The information about the event target.
	DetailMap map[string]interface{} `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	// The endpoint of the event target.
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the custom event target.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*GetRuleResponseBodyDataTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy that is used to push events. Valid values: BACKOFF_RETRY: backoff retry. If an event failed to be pushed, it can be retried up to three times. The interval between two consecutive retries is a random value from 10 to 20. Unit: seconds. EXPONENTIAL_DECAY_RETRY: exponential decay retry. If an event failed to be pushed, it can be retried up to 176 times. The interval between two consecutive retries exponentially increases to 512 seconds, and the total retry time is one day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 512, ..., and 512 seconds. The interval of 512 seconds is used for 167 retries.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The transformer that is used to push events.
	PushSelector *string `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	// The type of the event target. For more information, see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRuleResponseBodyDataTargets) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargets) SetDeadLetterQueue(v *GetRuleResponseBodyDataTargetsDeadLetterQueue) *GetRuleResponseBodyDataTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetDetailMap(v map[string]interface{}) *GetRuleResponseBodyDataTargets {
	s.DetailMap = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetEndpoint(v string) *GetRuleResponseBodyDataTargets {
	s.Endpoint = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetErrorsTolerance(v string) *GetRuleResponseBodyDataTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetId(v string) *GetRuleResponseBodyDataTargets {
	s.Id = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetParamList(v []*GetRuleResponseBodyDataTargetsParamList) *GetRuleResponseBodyDataTargets {
	s.ParamList = v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetPushRetryStrategy(v string) *GetRuleResponseBodyDataTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetPushSelector(v string) *GetRuleResponseBodyDataTargets {
	s.PushSelector = &v
	return s
}

func (s *GetRuleResponseBodyDataTargets) SetType(v string) *GetRuleResponseBodyDataTargets {
	s.Type = &v
	return s
}

type GetRuleResponseBodyDataTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the event source.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s GetRuleResponseBodyDataTargetsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargetsDeadLetterQueue) SetArn(v string) *GetRuleResponseBodyDataTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

type GetRuleResponseBodyDataTargetsParamList struct {
	// The format that is used by the event target parameter. For more information, see [Limits.](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target. For more information, see [Limits.](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template that is used by the event target parameter.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRuleResponseBodyDataTargetsParamList) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyDataTargetsParamList) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetForm(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Form = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetResourceKey(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetTemplate(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Template = &v
	return s
}

func (s *GetRuleResponseBodyDataTargetsParamList) SetValue(v string) *GetRuleResponseBodyDataTargetsParamList {
	s.Value = &v
	return s
}

type GetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetStatusCode(v int32) *GetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

type ListAliyunOfficialEventSourcesResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For a list of error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListAliyunOfficialEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. If the operation is successful, the value true is returned.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetCode(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetData(v *ListAliyunOfficialEventSourcesResponseBodyData) *ListAliyunOfficialEventSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetMessage(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetRequestId(v string) *ListAliyunOfficialEventSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBody) SetSuccess(v bool) *ListAliyunOfficialEventSourcesResponseBody {
	s.Success = &v
	return s
}

type ListAliyunOfficialEventSourcesResponseBodyData struct {
	// The event sources.
	EventSourceList []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList `json:"EventSourceList,omitempty" xml:"EventSourceList,omitempty" type:"Repeated"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyData) SetEventSourceList(v []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) *ListAliyunOfficialEventSourcesResponseBodyData {
	s.EventSourceList = v
	return s
}

type ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList struct {
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The time when the event source was created. Unit: milliseconds.
	Ctime *float32 `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event source to which the event type belongs.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event types.
	EventTypes []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// The full name of the event source.
	FullName *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	// The name of the event source.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the event source. Valid value: Activated.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the event source.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetArn(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Arn = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetCtime(v float32) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Ctime = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetDescription(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Description = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetEventBusName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.EventBusName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetEventTypes(v []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.EventTypes = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetFullName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.FullName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Name = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetStatus(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Status = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList) SetType(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceList {
	s.Type = &v
	return s
}

type ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes struct {
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The name of the group to which the event type belongs.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The full name of the event type.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The short name of the event type.
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetEventSourceName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.EventSourceName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetGroupName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.GroupName = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.Name = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes) SetShortName(v string) *ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes {
	s.ShortName = &v
	return s
}

type ListAliyunOfficialEventSourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliyunOfficialEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliyunOfficialEventSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunOfficialEventSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAliyunOfficialEventSourcesResponse) SetHeaders(v map[string]*string) *ListAliyunOfficialEventSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponse) SetStatusCode(v int32) *ListAliyunOfficialEventSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliyunOfficialEventSourcesResponse) SetBody(v *ListAliyunOfficialEventSourcesResponseBody) *ListAliyunOfficialEventSourcesResponse {
	s.Body = v
	return s
}

type ListApiDestinationsRequest struct {
	// The prefix of the API destination name.
	ApiDestinationNamePrefix *string `json:"ApiDestinationNamePrefix,omitempty" xml:"ApiDestinationNamePrefix,omitempty"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging.
	//
	// *   Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If you set Limit and excess return values exist, this parameter is returned.
	//
	// *   Default value: 0.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApiDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsRequest) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsRequest) SetApiDestinationNamePrefix(v string) *ListApiDestinationsRequest {
	s.ApiDestinationNamePrefix = &v
	return s
}

func (s *ListApiDestinationsRequest) SetConnectionName(v string) *ListApiDestinationsRequest {
	s.ConnectionName = &v
	return s
}

func (s *ListApiDestinationsRequest) SetMaxResults(v int64) *ListApiDestinationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiDestinationsRequest) SetNextToken(v string) *ListApiDestinationsRequest {
	s.NextToken = &v
	return s
}

type ListApiDestinationsResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListApiDestinationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApiDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBody) SetCode(v string) *ListApiDestinationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiDestinationsResponseBody) SetData(v *ListApiDestinationsResponseBodyData) *ListApiDestinationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApiDestinationsResponseBody) SetMessage(v string) *ListApiDestinationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiDestinationsResponseBody) SetRequestId(v string) *ListApiDestinationsResponseBody {
	s.RequestId = &v
	return s
}

type ListApiDestinationsResponseBodyData struct {
	// The API destinations.
	ApiDestinations []*ListApiDestinationsResponseBodyDataApiDestinations `json:"ApiDestinations,omitempty" xml:"ApiDestinations,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	MaxResults *float32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	Total *float32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListApiDestinationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyData) SetApiDestinations(v []*ListApiDestinationsResponseBodyDataApiDestinations) *ListApiDestinationsResponseBodyData {
	s.ApiDestinations = v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetMaxResults(v float32) *ListApiDestinationsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetNextToken(v string) *ListApiDestinationsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListApiDestinationsResponseBodyData) SetTotal(v float32) *ListApiDestinationsResponseBodyData {
	s.Total = &v
	return s
}

type ListApiDestinationsResponseBodyDataApiDestinations struct {
	// The name of the API destination.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the API destination was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The request parameters that are configured for the API destination.
	HttpApiParameters *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s ListApiDestinationsResponseBodyDataApiDestinations) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsResponseBodyDataApiDestinations) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetApiDestinationName(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.ApiDestinationName = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetConnectionName(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.ConnectionName = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetDescription(v string) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.Description = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetGmtCreate(v int64) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.GmtCreate = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinations) SetHttpApiParameters(v *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) *ListApiDestinationsResponseBodyDataApiDestinations {
	s.HttpApiParameters = v
	return s
}

type ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters struct {
	// The endpoint of the API destination.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// - POST
	//
	// - GET
	//
	// - DELETE
	//
	// - PUT
	//
	// - HEAD
	//
	// - TRACE
	//
	// - PATCH
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) SetEndpoint(v string) *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters) SetMethod(v string) *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters {
	s.Method = &v
	return s
}

type ListApiDestinationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApiDestinationsResponse) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsResponse) SetHeaders(v map[string]*string) *ListApiDestinationsResponse {
	s.Headers = v
	return s
}

func (s *ListApiDestinationsResponse) SetStatusCode(v int32) *ListApiDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiDestinationsResponse) SetBody(v *ListApiDestinationsResponseBody) *ListApiDestinationsResponse {
	s.Body = v
	return s
}

type ListConnectionsRequest struct {
	// The key word that you specify to query connections. Connections can be queried by prefixes.
	ConnectionNamePrefix *string `json:"ConnectionNamePrefix,omitempty" xml:"ConnectionNamePrefix,omitempty"`
	// The maximum number of entries to be returned in a single call. You can use this parameter and the NextToken parameter to implement paging.
	//
	// *   Default value: 10.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If you set the Limit parameter and excess return values exist, this parameter is returned.
	//
	// *   Default value: 0.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) SetConnectionNamePrefix(v string) *ListConnectionsRequest {
	s.ConnectionNamePrefix = &v
	return s
}

func (s *ListConnectionsRequest) SetMaxResults(v int64) *ListConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsRequest) SetNextToken(v string) *ListConnectionsRequest {
	s.NextToken = &v
	return s
}

type ListConnectionsResponseBody struct {
	// The HTTP status code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the connections returned.
	Data *ListConnectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) SetCode(v string) *ListConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConnectionsResponseBody) SetData(v *ListConnectionsResponseBodyData) *ListConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConnectionsResponseBody) SetMessage(v string) *ListConnectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConnectionsResponseBody) SetRequestId(v string) *ListConnectionsResponseBody {
	s.RequestId = &v
	return s
}

type ListConnectionsResponseBodyData struct {
	// The value of the key in the request path.
	Connections []*ListConnectionsResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *float32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	Total *float32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListConnectionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyData) SetConnections(v []*ListConnectionsResponseBodyDataConnections) *ListConnectionsResponseBodyData {
	s.Connections = v
	return s
}

func (s *ListConnectionsResponseBodyData) SetMaxResults(v float32) *ListConnectionsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsResponseBodyData) SetNextToken(v string) *ListConnectionsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListConnectionsResponseBodyData) SetTotal(v float32) *ListConnectionsResponseBodyData {
	s.Total = &v
	return s
}

type ListConnectionsResponseBodyDataConnections struct {
	// The parameters that are configured for authentication.
	AuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The connection description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the connection was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The connection ID.
	Id                *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	NetworkParameters *ListConnectionsResponseBodyDataConnectionsNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnections) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnections) SetAuthParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParameters) *ListConnectionsResponseBodyDataConnections {
	s.AuthParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetConnectionName(v string) *ListConnectionsResponseBodyDataConnections {
	s.ConnectionName = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetDescription(v string) *ListConnectionsResponseBodyDataConnections {
	s.Description = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetGmtCreate(v int64) *ListConnectionsResponseBodyDataConnections {
	s.GmtCreate = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetId(v int64) *ListConnectionsResponseBodyDataConnections {
	s.Id = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnections) SetNetworkParameters(v *ListConnectionsResponseBodyDataConnectionsNetworkParameters) *ListConnectionsResponseBodyDataConnections {
	s.NetworkParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParameters struct {
	// The parameters that are configured for API key authentication.
	ApiKeyAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type. Valid values:
	//
	// - BASIC_AUTH: basic authentication.
	//
	// - API_KEY_AUTH: API key authentication.
	//
	// - OAUTH_AUTH: OAuth authentication.
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The parameters that are configured for basic authentication.
	BasicAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are configured for OAuth authentication.
	OAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetApiKeyAuthParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.ApiKeyAuthParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetAuthorizationType(v string) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.AuthorizationType = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetBasicAuthParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.BasicAuthParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetOAuthParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.OAuthParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
	// The API key.
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters struct {
	// The password for basic authentication.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for basic authentication.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetPassword(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetUsername(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are configured for the client.
	ClientParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// - HEAD
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters for OAuth authentication.
	OAuthHttpParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) SetAuthorizationEndpoint(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) SetClientParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.ClientParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) SetHttpMethod(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.HttpMethod = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) SetOAuthHttpParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters {
	s.OAuthHttpParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The client key secret of the application.
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientID(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The parameters that are configured for the request.
	BodyParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request header.
	HeaderParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request path.
	QueryStringParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetBodyParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetHeaderParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) SetQueryStringParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	s.QueryStringParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key in the request body.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the key in the request body.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetIsValueSecret(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.IsValueSecret = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key in the request header.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the key in the request header.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetIsValueSecret(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.IsValueSecret = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Indicates whether authentication is enabled.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key in the request path.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the key in the request path.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetIsValueSecret(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.IsValueSecret = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsNetworkParameters struct {
	// The network type. Valid values:PublicNetwork and PrivateNetwork.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The virtual private cloud (VPC) ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsNetworkParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsNetworkParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) SetNetworkType(v string) *ListConnectionsResponseBodyDataConnectionsNetworkParameters {
	s.NetworkType = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) SetSecurityGroupId(v string) *ListConnectionsResponseBodyDataConnectionsNetworkParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) SetVpcId(v string) *ListConnectionsResponseBodyDataConnectionsNetworkParameters {
	s.VpcId = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) SetVswitcheId(v string) *ListConnectionsResponseBodyDataConnectionsNetworkParameters {
	s.VswitcheId = &v
	return s
}

type ListConnectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponse) SetHeaders(v map[string]*string) *ListConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionsResponse) SetStatusCode(v int32) *ListConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionsResponse) SetBody(v *ListConnectionsResponseBody) *ListConnectionsResponse {
	s.Body = v
	return s
}

type ListEventBusesRequest struct {
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. Note: Up to 100 entries can be returned in a call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The prefix of the names of the event buses that you want to query.
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you set Limit and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListEventBusesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventBusesRequest) GoString() string {
	return s.String()
}

func (s *ListEventBusesRequest) SetLimit(v int32) *ListEventBusesRequest {
	s.Limit = &v
	return s
}

func (s *ListEventBusesRequest) SetNamePrefix(v string) *ListEventBusesRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListEventBusesRequest) SetNextToken(v string) *ListEventBusesRequest {
	s.NextToken = &v
	return s
}

type ListEventBusesResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListEventBusesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the event buses are successfully queried. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventBusesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventBusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBody) SetCode(v string) *ListEventBusesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventBusesResponseBody) SetData(v *ListEventBusesResponseBodyData) *ListEventBusesResponseBody {
	s.Data = v
	return s
}

func (s *ListEventBusesResponseBody) SetMessage(v string) *ListEventBusesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventBusesResponseBody) SetRequestId(v string) *ListEventBusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventBusesResponseBody) SetSuccess(v bool) *ListEventBusesResponseBody {
	s.Success = &v
	return s
}

type ListEventBusesResponseBodyData struct {
	// The timestamp that indicates when the event bus was created.
	EventBuses []*ListEventBusesResponseBodyDataEventBuses `json:"EventBuses,omitempty" xml:"EventBuses,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventBusesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEventBusesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBodyData) SetEventBuses(v []*ListEventBusesResponseBodyDataEventBuses) *ListEventBusesResponseBodyData {
	s.EventBuses = v
	return s
}

func (s *ListEventBusesResponseBodyData) SetNextToken(v string) *ListEventBusesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventBusesResponseBodyData) SetTotal(v int32) *ListEventBusesResponseBodyData {
	s.Total = &v
	return s
}

type ListEventBusesResponseBodyDataEventBuses struct {
	// The timestamp that indicates when the event bus was created.
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the queried event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the queried event bus.
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	// The name of the queried event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s ListEventBusesResponseBodyDataEventBuses) String() string {
	return tea.Prettify(s)
}

func (s ListEventBusesResponseBodyDataEventBuses) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetCreateTimestamp(v int64) *ListEventBusesResponseBodyDataEventBuses {
	s.CreateTimestamp = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetDescription(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.Description = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetEventBusARN(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.EventBusARN = &v
	return s
}

func (s *ListEventBusesResponseBodyDataEventBuses) SetEventBusName(v string) *ListEventBusesResponseBodyDataEventBuses {
	s.EventBusName = &v
	return s
}

type ListEventBusesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventBusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventBusesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventBusesResponse) GoString() string {
	return s.String()
}

func (s *ListEventBusesResponse) SetHeaders(v map[string]*string) *ListEventBusesResponse {
	s.Headers = v
	return s
}

func (s *ListEventBusesResponse) SetStatusCode(v int32) *ListEventBusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventBusesResponse) SetBody(v *ListEventBusesResponseBody) *ListEventBusesResponse {
	s.Body = v
	return s
}

type ListEventStreamingsRequest struct {
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. A maximum of 100 entries can be returned in a call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event stream that you want to query.
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ARN of the event target.
	SinkArn *string `json:"SinkArn,omitempty" xml:"SinkArn,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event source.
	SourceArn *string `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
}

func (s ListEventStreamingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsRequest) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsRequest) SetLimit(v int32) *ListEventStreamingsRequest {
	s.Limit = &v
	return s
}

func (s *ListEventStreamingsRequest) SetNamePrefix(v string) *ListEventStreamingsRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListEventStreamingsRequest) SetNextToken(v string) *ListEventStreamingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventStreamingsRequest) SetSinkArn(v string) *ListEventStreamingsRequest {
	s.SinkArn = &v
	return s
}

func (s *ListEventStreamingsRequest) SetSourceArn(v string) *ListEventStreamingsRequest {
	s.SourceArn = &v
	return s
}

type ListEventStreamingsResponseBody struct {
	// The response code. Valid values:
	//
	// Success: The request is successful.
	//
	// Other codes: The request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the event streams.
	Data *ListEventStreamingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. The value true indicates that the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventStreamingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBody) SetCode(v string) *ListEventStreamingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetData(v *ListEventStreamingsResponseBodyData) *ListEventStreamingsResponseBody {
	s.Data = v
	return s
}

func (s *ListEventStreamingsResponseBody) SetMessage(v string) *ListEventStreamingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetRequestId(v string) *ListEventStreamingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventStreamingsResponseBody) SetSuccess(v bool) *ListEventStreamingsResponseBody {
	s.Success = &v
	return s
}

type ListEventStreamingsResponseBodyData struct {
	// The event streams.
	EventStreamings []*ListEventStreamingsResponseBodyDataEventStreamings `json:"EventStreamings,omitempty" xml:"EventStreamings,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists. You must specify the pagination token in the next request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of records.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventStreamingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyData) SetEventStreamings(v []*ListEventStreamingsResponseBodyDataEventStreamings) *ListEventStreamingsResponseBodyData {
	s.EventStreamings = v
	return s
}

func (s *ListEventStreamingsResponseBodyData) SetNextToken(v string) *ListEventStreamingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventStreamingsResponseBodyData) SetTotal(v int32) *ListEventStreamingsResponseBodyData {
	s.Total = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamings struct {
	// The description of the event stream.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are returned for the runtime environment.
	RunOptions *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target.
	Sink *ListEventStreamingsResponseBodyDataEventStreamingsSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event provider, which is also known as the event source.
	Source *ListEventStreamingsResponseBodyDataEventStreamingsSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// The status of the event stream that is returned.
	Status     *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Transforms []*ListEventStreamingsResponseBodyDataEventStreamingsTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamings) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamings) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetDescription(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Description = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetEventStreamingName(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.EventStreamingName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetFilterPattern(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.FilterPattern = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetRunOptions(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.RunOptions = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetSink(v *ListEventStreamingsResponseBodyDataEventStreamingsSink) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Sink = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetSource(v *ListEventStreamingsResponseBodyDataEventStreamingsSource) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Source = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetStatus(v string) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Status = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamings) SetTransforms(v []*ListEventStreamingsResponseBodyDataEventStreamingsTransforms) *ListEventStreamingsResponseBodyDataEventStreamings {
	s.Transforms = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptions struct {
	// The batch window.
	BatchWindow *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	// Indicates whether dead-letter queues are enabled. By default, dead-letter queues are disabled. Messages that fail to be pushed are discarded after the maximum number of retries that is specified by the retry policy is reached.
	DeadLetterQueue *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The exception tolerance policy. Valid values: NONE and ALL.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The number of concurrent threads.
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The retry policy that is used if events fail to be pushed.
	RetryStrategy *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetBatchWindow(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.BatchWindow = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetDeadLetterQueue(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetErrorsTolerance(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetMaximumTasks(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions) SetRetryStrategy(v *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions {
	s.RetryStrategy = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow struct {
	// The maximum number of events that are allowed in the batch window. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. When this threshold is reached, data in the window is pushed to the downstream service. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) SetCountBasedWindow(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue) SetArn(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy struct {
	// The maximum timeout period for a retry.
	MaximumEventAgeInSeconds *float32 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	MaximumRetryAttempts *float32 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// Valid values: BACKOFFRETRY and EXPONENTIALDECAY_RETRY.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v float32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) SetMaximumRetryAttempts(v float32) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSink struct {
	// The parameters that are returned if Function Compute is specified as the event target.
	SinkFcParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Serverless Workflow is specified as the event target.
	SinkFnfParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache Kafka is specified as the event target.
	SinkKafkaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if MNS is specified as the event target.
	SinkMNSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for RabbitMQ is specified as the event target.
	SinkRabbitMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Parameters
	SinkRocketMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// Sink SLS Parameters
	SinkSLSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSink) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSink) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkFcParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkFcParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkFnfParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkFnfParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkMNSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkMNSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkRocketMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSink) SetSinkSLSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSink {
	s.SinkSLSParameters = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters struct {
	// The message body that is delivered to Function Compute.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The delivery concurrency. Minimum value: 1.
	Concurrency *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The function name.
	FunctionName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The invocation mode. Valid values:
	//
	// *   Sync: the synchronous mode
	// *   Async: the asynchronous mode
	InvocationType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The alias of the service to which the function belongs.
	Qualifier *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The service name.
	ServiceName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetConcurrency(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetFunctionName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetInvocationType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetQualifier(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters) SetServiceName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The function name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation mode. Valid values:
	//
	// *   Sync: the synchronous mode
	// *   Async: the asynchronous mode
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The alias of the service to which the function belongs.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The service name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters struct {
	// The execution name.
	ExecutionName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	// The flow name.
	FlowName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	// The input information of the execution.
	Input *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The role name.
	RoleName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetExecutionName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetFlowName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetInput(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters) SetRoleName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The execution name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The flow name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The input information of the execution.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters struct {
	// The acknowledgment (ACK) mode.
	//
	// *   If this parameter is set to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	// *   If this parameter is set to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	// *   If this parameter is set to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message key.
	Key *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The topic name.
	Topic *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The message body.
	Value *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetAcks(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters) SetValue(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters {
	s.Value = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ACK mode.
	//
	// *   If this parameter is set to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high.
	// *   If this parameter is set to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader.
	// *   If this parameter is set to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The topic name.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters struct {
	// The message body.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Indicates whether Base64 encoding is enabled.
	IsBase64Encode *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The name of the MNS queue.
	QueueName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetIsBase64Encode(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters) SetQueueName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Indicates that Base64 encoding is enabled.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the MNS queue.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters struct {
	// The message body.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The exchange to which events are delivered. This parameter is returned only if TargetType is set to Exchange.
	Exchange *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The ID of the Message Queue for RabbitMQ instance.
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The attributes that are used to filter messages.
	Properties *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The queue to which events are delivered. This parameter is returned only if TargetType is set to Queue.
	QueueName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The rule that is used to route messages. This parameter is returned only if TargetType is set to Exchange.
	RoutingKey *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The type of the resource to which events are delivered.
	TargetType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetExchange(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetMessageId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetProperties(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetQueueName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetRoutingKey(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetTargetType(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters) SetVirtualHostName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange on the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The rule that is used to route messages.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which events are delivered. Valid values: Exchange and Queue.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters struct {
	// The message body.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache RocketMQ is specified as the event target.
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The keys that are used to filter messages.
	Keys *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The attributes that are used to filter messages.
	Properties *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The tags that are used to filter messages.
	Tags *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance.
	Topic *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetInstanceId(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetKeys(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetProperties(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetTags(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters struct {
	// The message body that is delivered to Log Service.
	Body *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The Log Service Logstore.
	LogStore *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The Log Service project.
	Project *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	RoleName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	Topic *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetBody(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetLogStore(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetProject(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetRoleName(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters) SetTopic(v *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParameters {
	s.Topic = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service Logstore.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service project.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetForm(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetTemplate(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic) SetValue(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSource struct {
	// The parameters that are returned if Data Transmission Service (DTS) is specified as the event source.
	SourceDTSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache Kafka is specified as the event source.
	SourceKafkaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Service (MNS) is specified as the event source.
	SourceMNSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for MQTT is specified as the event source.
	SourceMQTTParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for RabbitMQ is specified as the event source.
	SourceRabbitMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// Source RocketMQ Parameters
	SourceRocketMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Log Service is specified as the event source.
	SourceSLSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSource) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSource) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceDTSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceDTSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceKafkaParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceMNSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceMNSParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceMQTTParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceRabbitMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceRocketMQParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSource) SetSourceSLSParameters(v *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) *ListEventStreamingsResponseBodyDataEventStreamingsSource {
	s.SourceSLSParameters = v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters struct {
	// The URL and port number of the change tracking instance.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The consumer offset. It is the timestamp that indicates when the SDK client consumes the first data record.
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic on the change tracking instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The consumer group username.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetBrokerUrl(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetInitCheckPoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetPassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetSid(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetTaskId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters) SetUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters {
	s.Username = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a VPC.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the Message Queue for Apache Kafka instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetConsumerGroup(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetOffsetReset(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters) SetVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters struct {
	// Indicates whether Base64 encoding is enabled.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetIsBase64Decode(v bool) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters struct {
	// The ID of the Message Queue for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Message Queue for MQTT instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic on the Message Queue for MQTT instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetQueueName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters struct {
	// The authentication method.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the group on the Message Queue for Apache RocketMQ instance.
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The instance endpoint.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// *   PublicNetwork
	// *   PrivateNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The instance password.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The security group ID of the instance.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance username.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The vSwitch ID of the instance.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The virtual private cloud (VPC) ID.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The consumer offset of the message. Valid values: CONSUMEFROMLASTOFFSET: Messages are consumed from the latest offset. CONSUMEFROMFIRSTOFFSET: Messages are consumed from the earliest offset. CONSUMEFROMTIMESTAMP: Messages are consumed from the offset at the specified point in time.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that indicates the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUMEFROMTIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic in which messages are stored.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetAuthType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetGroupID(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstancePassword(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceType(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceUsername(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetOffset(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetRegionId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTag(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTimestamp(v int64) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetTopic(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters struct {
	// The consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start message consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The ID of the consumer group that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetConsumePosition(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetConsumerGroup(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetLogStore(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetProject(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters) SetRoleName(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

type ListEventStreamingsResponseBodyDataEventStreamingsTransforms struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsTransforms) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsTransforms) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsTransforms) SetArn(v string) *ListEventStreamingsResponseBodyDataEventStreamingsTransforms {
	s.Arn = &v
	return s
}

type ListEventStreamingsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventStreamingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventStreamingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponse) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponse) SetHeaders(v map[string]*string) *ListEventStreamingsResponse {
	s.Headers = v
	return s
}

func (s *ListEventStreamingsResponse) SetStatusCode(v int32) *ListEventStreamingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventStreamingsResponse) SetBody(v *ListEventStreamingsResponseBody) *ListEventStreamingsResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of entries to be returned in a single call. You can use this parameter and the NextToken parameter to implement paging. A maximum of 100 entries can be returned in a single call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// If you set the Limit parameter and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix of the rule name.
	RuleNamePrefix *string `json:"RuleNamePrefix,omitempty" xml:"RuleNamePrefix,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetEventBusName(v string) *ListRulesRequest {
	s.EventBusName = &v
	return s
}

func (s *ListRulesRequest) SetLimit(v int32) *ListRulesRequest {
	s.Limit = &v
	return s
}

func (s *ListRulesRequest) SetNextToken(v string) *ListRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRulesRequest) SetRuleNamePrefix(v string) *ListRulesRequest {
	s.RuleNamePrefix = &v
	return s
}

type ListRulesResponseBody struct {
	// The error code. The value Success indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v *ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetSuccess(v bool) *ListRulesResponseBody {
	s.Success = &v
	return s
}

type ListRulesResponseBodyData struct {
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The rules.
	Rules []*ListRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) SetNextToken(v string) *ListRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRules(v []*ListRulesResponseBodyDataRules) *ListRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBodyData) SetTotal(v int32) *ListRulesResponseBodyData {
	s.Total = &v
	return s
}

type ListRulesResponseBodyDataRules struct {
	// The creation timestamp.
	CreatedTimestamp *int64 `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	// The rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the event rule.
	DetailMap map[string]interface{} `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual pattern stringExpression pattern Each field can have a maximum of five expressions in the map data structure.
	//
	// Each field can have a maximum of five expressions in the map data structure.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: The event rule is enabled. It is the default state of the event rule. DISABLE: The event rule is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event targets.
	Targets []*ListRulesResponseBodyDataRulesTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataRules) SetCreatedTimestamp(v int64) *ListRulesResponseBodyDataRules {
	s.CreatedTimestamp = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetDescription(v string) *ListRulesResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetDetailMap(v map[string]interface{}) *ListRulesResponseBodyDataRules {
	s.DetailMap = v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetEventBusName(v string) *ListRulesResponseBodyDataRules {
	s.EventBusName = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetFilterPattern(v string) *ListRulesResponseBodyDataRules {
	s.FilterPattern = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetRuleARN(v string) *ListRulesResponseBodyDataRules {
	s.RuleARN = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetRuleName(v string) *ListRulesResponseBodyDataRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetStatus(v string) *ListRulesResponseBodyDataRules {
	s.Status = &v
	return s
}

func (s *ListRulesResponseBodyDataRules) SetTargets(v []*ListRulesResponseBodyDataRulesTargets) *ListRulesResponseBodyDataRules {
	s.Targets = v
	return s
}

type ListRulesResponseBodyDataRulesTargets struct {
	// The endpoint of the event target.
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the custom event target.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The transformer that is used to push events.
	PushSelector *string `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	// The type of the event target. For more information, see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyDataRulesTargets) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyDataRulesTargets) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataRulesTargets) SetEndpoint(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Endpoint = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetErrorsTolerance(v string) *ListRulesResponseBodyDataRulesTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetId(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Id = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetPushSelector(v string) *ListRulesResponseBodyDataRulesTargets {
	s.PushSelector = &v
	return s
}

func (s *ListRulesResponseBodyDataRulesTargets) SetType(v string) *ListRulesResponseBodyDataRulesTargets {
	s.Type = &v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListTargetsRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the event rule.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of returned entries in a call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListTargetsRequest) SetArn(v string) *ListTargetsRequest {
	s.Arn = &v
	return s
}

func (s *ListTargetsRequest) SetEventBusName(v string) *ListTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *ListTargetsRequest) SetLimit(v int32) *ListTargetsRequest {
	s.Limit = &v
	return s
}

func (s *ListTargetsRequest) SetNextToken(v string) *ListTargetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTargetsRequest) SetRuleName(v string) *ListTargetsRequest {
	s.RuleName = &v
	return s
}

type ListTargetsResponseBody struct {
	// The returned response code. Valid values:
	//
	// *   Success: The request is successful.
	// *   Other codes: The request failed. For a list of error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBody) SetCode(v string) *ListTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTargetsResponseBody) SetData(v *ListTargetsResponseBodyData) *ListTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListTargetsResponseBody) SetMessage(v string) *ListTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTargetsResponseBody) SetRequestId(v string) *ListTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetsResponseBody) SetSuccess(v bool) *ListTargetsResponseBody {
	s.Success = &v
	return s
}

type ListTargetsResponseBodyData struct {
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The event targets.
	Targets []*ListTargetsResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The total number of entries.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTargetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyData) SetNextToken(v string) *ListTargetsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTargetsResponseBodyData) SetTargets(v []*ListTargetsResponseBodyDataTargets) *ListTargetsResponseBodyData {
	s.Targets = v
	return s
}

func (s *ListTargetsResponseBodyData) SetTotal(v int32) *ListTargetsResponseBodyData {
	s.Total = &v
	return s
}

type ListTargetsResponseBodyDataTargets struct {
	// The endpoint of the event target.
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The ID of the event target.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*ListTargetsResponseBodyDataTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the event target. For more information, see [Event target parameters](~~183698~~).
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTargetsResponseBodyDataTargets) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyDataTargets) SetEndpoint(v string) *ListTargetsResponseBodyDataTargets {
	s.Endpoint = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetErrorsTolerance(v string) *ListTargetsResponseBodyDataTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetEventBusName(v string) *ListTargetsResponseBodyDataTargets {
	s.EventBusName = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetId(v string) *ListTargetsResponseBodyDataTargets {
	s.Id = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetParamList(v []*ListTargetsResponseBodyDataTargetsParamList) *ListTargetsResponseBodyDataTargets {
	s.ParamList = v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetRuleName(v string) *ListTargetsResponseBodyDataTargets {
	s.RuleName = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargets) SetType(v string) *ListTargetsResponseBodyDataTargets {
	s.Type = &v
	return s
}

type ListTargetsResponseBodyDataTargetsParamList struct {
	// The format that is used by the event target parameter.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target.
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template that is used by the event target parameter.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTargetsResponseBodyDataTargetsParamList) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsResponseBodyDataTargetsParamList) GoString() string {
	return s.String()
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetForm(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Form = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetResourceKey(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetTemplate(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Template = &v
	return s
}

func (s *ListTargetsResponseBodyDataTargetsParamList) SetValue(v string) *ListTargetsResponseBodyDataTargetsParamList {
	s.Value = &v
	return s
}

type ListTargetsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListTargetsResponse) SetHeaders(v map[string]*string) *ListTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListTargetsResponse) SetStatusCode(v int32) *ListTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetsResponse) SetBody(v *ListTargetsResponseBody) *ListTargetsResponse {
	s.Body = v
	return s
}

type ListUserDefinedEventSourcesRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. Note: Up to 100 entries can be returned in a call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event source.
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserDefinedEventSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesRequest) SetEventBusName(v string) *ListUserDefinedEventSourcesRequest {
	s.EventBusName = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetLimit(v int32) *ListUserDefinedEventSourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetNamePrefix(v string) *ListUserDefinedEventSourcesRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListUserDefinedEventSourcesRequest) SetNextToken(v string) *ListUserDefinedEventSourcesRequest {
	s.NextToken = &v
	return s
}

type ListUserDefinedEventSourcesResponseBody struct {
	// The returned response code. Valid values:
	//
	// *   Success: The request is successful.
	// *   Other codes: The request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListUserDefinedEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBody) SetCode(v string) *ListUserDefinedEventSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetData(v *ListUserDefinedEventSourcesResponseBodyData) *ListUserDefinedEventSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetMessage(v string) *ListUserDefinedEventSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetRequestId(v string) *ListUserDefinedEventSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetSuccess(v bool) *ListUserDefinedEventSourcesResponseBody {
	s.Success = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyData struct {
	// The event sources.
	EventSourceList []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList `json:"EventSourceList,omitempty" xml:"EventSourceList,omitempty" type:"Repeated"`
	NextToken       *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total           *int32                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetEventSourceList(v []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList) *ListUserDefinedEventSourcesResponseBodyData {
	s.EventSourceList = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetNextToken(v string) *ListUserDefinedEventSourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetTotal(v int32) *ListUserDefinedEventSourcesResponseBodyData {
	s.Total = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceList struct {
	// The Alibaba Cloud Resource Name (ARN) of the queried event source.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The timestamp that indicates when the event source was created.
	Ctime *float32 `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The type of the event source.
	ExternalSourceType *string `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// The name of the queried event source.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that are returned if HTTP events are specified as the event source.
	SourceHttpEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache Kafka is specified as the event source.
	SourceKafkaParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Service (MNS) is specified as the event source.
	SourceMNSParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for RabbitMQ is specified as the event source.
	SourceRabbitMQParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache RocketMQ is specified as the event source.
	SourceRocketMQParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Log Service is specified as the event source.
	SourceSLSParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if scheduled events are specified as the event source.
	SourceScheduledEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
	// The status of the queried event source. The returned value Activated indicates that the event source is activated.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the queried event source. The returned value UserDefined indicates that the event source is a custom event source.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceList) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetArn(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Arn = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetCtime(v float32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Ctime = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetEventBusName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.EventBusName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetExternalSourceType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.ExternalSourceType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Name = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceHttpEventParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceHttpEventParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceKafkaParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceKafkaParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceMNSParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceMNSParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceRabbitMQParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceRocketMQParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceRocketMQParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceSLSParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceSLSParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceScheduledEventParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceScheduledEventParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetStatus(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Status = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Type = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters struct {
	// The CIDR block that is used for security settings. This parameter is required only if SecurityConfig is set to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request method that is supported by the generated webhook URL. You can select multiple values. Valid values:
	//
	// *   GET
	// *   POST
	// *   PUT
	// *   PATCH
	// *   DELETE
	// *   HEAD
	// *   OPTIONS
	// *   TRACE
	// *   CONNECT
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	// The Internet request URL.
	PublicWebHookUrl []*string `json:"PublicWebHookUrl,omitempty" xml:"PublicWebHookUrl,omitempty" type:"Repeated"`
	// The security domain name. This parameter is required only if SecurityConfig is set to referer. You can enter a domain name.
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	// The type of security settings. Valid values:
	//
	// *   none: No configuration is required.
	// *   ip: CIDR block.
	// *   referer: security domain name.
	SecurityConfig *string `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	// The protocol type that is supported by the generated webhook URL. Valid values:
	//
	// *   HTTP
	// *   HTTPS
	// *   HTTP\&HTTPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The internal request URL.
	VpcWebHookUrl []*string `json:"VpcWebHookUrl,omitempty" xml:"VpcWebHookUrl,omitempty" type:"Repeated"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetIp(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetMethod(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetPublicWebHookUrl(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.PublicWebHookUrl = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetReferer(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetSecurityConfig(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Type = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetVpcWebHookUrl(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.VpcWebHookUrl = v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic on the Message Queue for Apache Kafka instance.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of consumers.
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network type. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a self-managed network.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the Message Queue for Apache Kafka instance belongs.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache Kafka instance is associated.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC in which the Message Queue for Apache Kafka instance is deployed.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetConsumerGroup(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetMaximumTasks(v int32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetNetwork(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetOffsetReset(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetSecurityGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetTopic(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetVSwitchIds(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetVpcId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters struct {
	// Indicates whether Base64 decoding is enabled. By default, Base64 decoding is enabled.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetIsBase64Decode(v bool) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetQueueName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.RegionId = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetQueueName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetVirtualHostName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters struct {
	// The authentication type. This parameter can be set to ACL or left empty.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The endpoint that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of network over which the Message Queue for Apache RocketMQ instance is accessed.
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the Message Queue for Apache RocketMQ instance.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the Message Queue for Apache RocketMQ instance belongs.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type. Valid values: CLOUD\_4, CLOUD\_5, and SELF_BUILT. The value CLOUD\_4 indicates that the instance is a Message Queue for Apache RocketMQ 4.0 instance. The value CLOUD\_5 indicates that the instance is a Message Queue for Apache RocketMQ 5.0 instance. The value SELF_BUILT indicates that the instance is a self-managed RocketMQ instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache RocketMQ instance is associated.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the Message Queue for Apache RocketMQ instance is deployed.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// *   CONSUME_FROM_LAST_OFFSET: Messages are consumed from the latest offset.
	// *   CONSUME_FROM_FIRST_OFFSET: Messages are consumed from the earliest offset.
	// *   CONSUME_FROM_TIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that indicates the time from which messages are consumed. This parameter is valid only if Offset is set to CONSUME_FROM_TIMESTAMP.
	Timestamp *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetAuthType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.GroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceEndpoint(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceNetwork(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstancePassword(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceUsername(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceVpcId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetOffset(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTag(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTimestamp(v float32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTopic(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters struct {
	// The consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Simple Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console. For information about the permission policy of this role, see Create a custom event source of the Log Service type.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetConsumePosition(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetLogStore(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetProject(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetRoleName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.RoleName = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters struct {
	// The cron expression.
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone in which the cron expression is executed.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The JSON string.
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetSchedule(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetTimeZone(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetUserData(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

type ListUserDefinedEventSourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDefinedEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDefinedEventSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponse) SetHeaders(v map[string]*string) *ListUserDefinedEventSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserDefinedEventSourcesResponse) SetStatusCode(v int32) *ListUserDefinedEventSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponse) SetBody(v *ListUserDefinedEventSourcesResponseBody) *ListUserDefinedEventSourcesResponse {
	s.Body = v
	return s
}

type PauseEventStreamingRequest struct {
	// The name of the event stream that you want to stop.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s PauseEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingRequest) SetEventStreamingName(v string) *PauseEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

type PauseEventStreamingResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *bool `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingResponseBody) SetCode(v bool) *PauseEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetMessage(v string) *PauseEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetRequestId(v string) *PauseEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetSuccess(v bool) *PauseEventStreamingResponseBody {
	s.Success = &v
	return s
}

type PauseEventStreamingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingResponse) SetHeaders(v map[string]*string) *PauseEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *PauseEventStreamingResponse) SetStatusCode(v int32) *PauseEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseEventStreamingResponse) SetBody(v *PauseEventStreamingResponseBody) *PauseEventStreamingResponse {
	s.Body = v
	return s
}

type PutTargetsRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The event targets to be created or updated. For more information, see [Limits.](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	Targets []*PutTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s PutTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutTargetsRequest) SetEventBusName(v string) *PutTargetsRequest {
	s.EventBusName = &v
	return s
}

func (s *PutTargetsRequest) SetRuleName(v string) *PutTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutTargetsRequest) SetTargets(v []*PutTargetsRequestTargets) *PutTargetsRequest {
	s.Targets = v
	return s
}

type PutTargetsRequestTargets struct {
	// The dead-letter queue. Events that are not processed or whose maximum retries have been exceeded are written to the dead-letter queue. The dead-letter queue feature supports the following queue types: Message Queue for Apache RocketMQ, Message Service, Message Queue for Apache Kafka, and event bus.
	DeadLetterQueue *PutTargetsRequestTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The endpoint of the event target.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The fault tolerance policy. Valid values:
	//
	// * **ALL**: ignores the error. Fault tolerance is allowed. If an error occurs, event processing is not blocked. If the message exceeds the number of retries specified by the retry policy, the message is delivered to a dead-letter queue or discarded based on your configurations.
	//
	// * **NONE**: does not ignore the error. Fault tolerance is prohibited. If an error occurs and the message exceeds the number of retries specified by the retry policy, event processing is blocked.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The ID of the custom event target.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are configured for the event target.
	ParamList []*PutTargetsRequestTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// The retry policy for pushing the event. Valid values:
	//
	// * **BACKOFF_RETRY**: backoff retry. A failed event can be retried up to three times. The interval between two consecutive retries is a random value from 10 to 20. Unit: seconds.
	//
	// * **EXPONENTIAL_DECAY_RETRY**: exponential decay retry. The request can be retried up to 176 times. The interval between two consecutive retries exponentially increases to 512 seconds, and the total retry time is one day. The specific retry intervals are 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 512, ..., and 512 seconds. The interval of 512 seconds can be used up to one hundred and sixty-seven times in total.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	// The type of the event target. For more information, see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PutTargetsRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsRequestTargets) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargets) SetDeadLetterQueue(v *PutTargetsRequestTargetsDeadLetterQueue) *PutTargetsRequestTargets {
	s.DeadLetterQueue = v
	return s
}

func (s *PutTargetsRequestTargets) SetEndpoint(v string) *PutTargetsRequestTargets {
	s.Endpoint = &v
	return s
}

func (s *PutTargetsRequestTargets) SetErrorsTolerance(v string) *PutTargetsRequestTargets {
	s.ErrorsTolerance = &v
	return s
}

func (s *PutTargetsRequestTargets) SetId(v string) *PutTargetsRequestTargets {
	s.Id = &v
	return s
}

func (s *PutTargetsRequestTargets) SetParamList(v []*PutTargetsRequestTargetsParamList) *PutTargetsRequestTargets {
	s.ParamList = v
	return s
}

func (s *PutTargetsRequestTargets) SetPushRetryStrategy(v string) *PutTargetsRequestTargets {
	s.PushRetryStrategy = &v
	return s
}

func (s *PutTargetsRequestTargets) SetType(v string) *PutTargetsRequestTargets {
	s.Type = &v
	return s
}

type PutTargetsRequestTargetsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue. Events that are not processed or whose maximum retries have been exceeded are written to the dead-letter queue.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s PutTargetsRequestTargetsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsRequestTargetsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargetsDeadLetterQueue) SetArn(v string) *PutTargetsRequestTargetsDeadLetterQueue {
	s.Arn = &v
	return s
}

type PutTargetsRequestTargetsParamList struct {
	// The method that is used to deliver events to the event target. For more information,see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The resource parameter of the event target. For more information,see [Event target parameters.](https://www.alibabacloud.com/help/en/eventbridge/latest/event-target-parameters)
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// The template based on which events are delivered to the event target.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value of the event target parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutTargetsRequestTargetsParamList) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsRequestTargetsParamList) GoString() string {
	return s.String()
}

func (s *PutTargetsRequestTargetsParamList) SetForm(v string) *PutTargetsRequestTargetsParamList {
	s.Form = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetResourceKey(v string) *PutTargetsRequestTargetsParamList {
	s.ResourceKey = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetTemplate(v string) *PutTargetsRequestTargetsParamList {
	s.Template = &v
	return s
}

func (s *PutTargetsRequestTargetsParamList) SetValue(v string) *PutTargetsRequestTargetsParamList {
	s.Value = &v
	return s
}

type PutTargetsShrinkRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The event targets to be created or updated. For more information, see [Limits.](https://www.alibabacloud.com/help/en/eventbridge/latest/limits)
	TargetsShrink *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s PutTargetsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutTargetsShrinkRequest) SetEventBusName(v string) *PutTargetsShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *PutTargetsShrinkRequest) SetRuleName(v string) *PutTargetsShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *PutTargetsShrinkRequest) SetTargetsShrink(v string) *PutTargetsShrinkRequest {
	s.TargetsShrink = &v
	return s
}

type PutTargetsResponseBody struct {
	// The response code. Valid values:
	//
	// *   Success: The call succeeded.
	// *   Other codes: The call failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *PutTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request is successful.
	// *   **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBody) SetCode(v string) *PutTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutTargetsResponseBody) SetData(v *PutTargetsResponseBodyData) *PutTargetsResponseBody {
	s.Data = v
	return s
}

func (s *PutTargetsResponseBody) SetMessage(v string) *PutTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutTargetsResponseBody) SetRequestId(v string) *PutTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutTargetsResponseBody) SetSuccess(v bool) *PutTargetsResponseBody {
	s.Success = &v
	return s
}

type PutTargetsResponseBodyData struct {
	// The ID of the failed event target.
	ErrorEntries []*PutTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	// The number of failed event targets. Valid values:
	//
	// *   0: All event targets succeeded.
	// *   An integer other than 0: indicates the number of failed event targets.
	ErrorEntriesCount *int32 `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
}

func (s PutTargetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBodyData) SetErrorEntries(v []*PutTargetsResponseBodyDataErrorEntries) *PutTargetsResponseBodyData {
	s.ErrorEntries = v
	return s
}

func (s *PutTargetsResponseBodyData) SetErrorEntriesCount(v int32) *PutTargetsResponseBodyData {
	s.ErrorEntriesCount = &v
	return s
}

type PutTargetsResponseBodyDataErrorEntries struct {
	// The ID of the failed event target.
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// The error code returned.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s PutTargetsResponseBodyDataErrorEntries) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsResponseBodyDataErrorEntries) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetEntryId(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.EntryId = &v
	return s
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetErrorCode(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.ErrorCode = &v
	return s
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetErrorMessage(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.ErrorMessage = &v
	return s
}

type PutTargetsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutTargetsResponse) SetHeaders(v map[string]*string) *PutTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutTargetsResponse) SetStatusCode(v int32) *PutTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutTargetsResponse) SetBody(v *PutTargetsResponseBody) *PutTargetsResponse {
	s.Body = v
	return s
}

type QueryEventRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// EventSource is required for querying default bus events.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
}

func (s QueryEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRequest) SetEventBusName(v string) *QueryEventRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryEventRequest) SetEventId(v string) *QueryEventRequest {
	s.EventId = &v
	return s
}

func (s *QueryEventRequest) SetEventSource(v string) *QueryEventRequest {
	s.EventSource = &v
	return s
}

type QueryEventResponseBody struct {
	// The status code returned. The status code 200 indicates that the request was successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The content of the event.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventResponseBody) SetCode(v string) *QueryEventResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventResponseBody) SetData(v map[string]interface{}) *QueryEventResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventResponseBody) SetMessage(v string) *QueryEventResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventResponseBody) SetRequestId(v string) *QueryEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventResponseBody) SetSuccess(v bool) *QueryEventResponseBody {
	s.Success = &v
	return s
}

type QueryEventResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventResponse) GoString() string {
	return s.String()
}

func (s *QueryEventResponse) SetHeaders(v map[string]*string) *QueryEventResponse {
	s.Headers = v
	return s
}

func (s *QueryEventResponse) SetStatusCode(v int32) *QueryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventResponse) SetBody(v *QueryEventResponseBody) *QueryEventResponse {
	s.Body = v
	return s
}

type QueryEventTracesRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s QueryEventTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventTracesRequest) GoString() string {
	return s.String()
}

func (s *QueryEventTracesRequest) SetEventBusName(v string) *QueryEventTracesRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryEventTracesRequest) SetEventId(v string) *QueryEventTracesRequest {
	s.EventId = &v
	return s
}

type QueryEventTracesResponseBody struct {
	// The status code returned. The status code 200 indicates that the request was successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the event source.
	Data []*QueryEventTracesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventTracesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponseBody) SetCode(v string) *QueryEventTracesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetData(v []*QueryEventTracesResponseBodyData) *QueryEventTracesResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventTracesResponseBody) SetMessage(v string) *QueryEventTracesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetRequestId(v string) *QueryEventTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetSuccess(v bool) *QueryEventTracesResponseBody {
	s.Success = &v
	return s
}

type QueryEventTracesResponseBodyData struct {
	// The type of the event trace. Valid values: PutEvent: a delivery event. FilterEvent: a filtering event. PushEvent: a pushing event.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The execution time of the event trace.
	ActionTime *int64 `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	// The endpoint of the event target. This parameter is returned if the value of the Action parameter is PushEvent.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event source.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The delivery delay of the event target. This parameter is returned if the value of the Action parameter is PushEvent.
	NotifyLatency *string `json:"NotifyLatency,omitempty" xml:"NotifyLatency,omitempty"`
	// The event target delivery status.
	NotifyStatus *string `json:"NotifyStatus,omitempty" xml:"NotifyStatus,omitempty"`
	// The delivery time of the event target. This parameter is returned if the value of the Action parameter is PushEvent.
	NotifyTime *int64 `json:"NotifyTime,omitempty" xml:"NotifyTime,omitempty"`
	// The time when the event was delivered to the event bus. This parameter is returned if the value of the Action parameter is PutEvent.
	ReceivedTime *int64 `json:"ReceivedTime,omitempty" xml:"ReceivedTime,omitempty"`
	// The time when the event rule was matched. This parameter is returned if the value of the Action parameter is FilterEvent.
	RuleMatchingTime *string `json:"RuleMatchingTime,omitempty" xml:"RuleMatchingTime,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s QueryEventTracesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventTracesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponseBodyData) SetAction(v string) *QueryEventTracesResponseBodyData {
	s.Action = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetActionTime(v int64) *QueryEventTracesResponseBodyData {
	s.ActionTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEndpoint(v string) *QueryEventTracesResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventBusName(v string) *QueryEventTracesResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventId(v string) *QueryEventTracesResponseBodyData {
	s.EventId = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventSource(v string) *QueryEventTracesResponseBodyData {
	s.EventSource = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyLatency(v string) *QueryEventTracesResponseBodyData {
	s.NotifyLatency = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyStatus(v string) *QueryEventTracesResponseBodyData {
	s.NotifyStatus = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyTime(v int64) *QueryEventTracesResponseBodyData {
	s.NotifyTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetReceivedTime(v int64) *QueryEventTracesResponseBodyData {
	s.ReceivedTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetRuleMatchingTime(v string) *QueryEventTracesResponseBodyData {
	s.RuleMatchingTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetRuleName(v string) *QueryEventTracesResponseBodyData {
	s.RuleName = &v
	return s
}

type QueryEventTracesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventTracesResponse) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponse) SetHeaders(v map[string]*string) *QueryEventTracesResponse {
	s.Headers = v
	return s
}

func (s *QueryEventTracesResponse) SetStatusCode(v int32) *QueryEventTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventTracesResponse) SetBody(v *QueryEventTracesResponseBody) *QueryEventTracesResponse {
	s.Body = v
	return s
}

type QueryTracedEventByEventIdRequest struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event source.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
}

func (s QueryTracedEventByEventIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventByEventIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdRequest) SetEventBusName(v string) *QueryTracedEventByEventIdRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventByEventIdRequest) SetEventId(v string) *QueryTracedEventByEventIdRequest {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventByEventIdRequest) SetEventSource(v string) *QueryTracedEventByEventIdRequest {
	s.EventSource = &v
	return s
}

type QueryTracedEventByEventIdResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	Data []*QueryTracedEventByEventIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBody) SetCode(v string) *QueryTracedEventByEventIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetData(v []*QueryTracedEventByEventIdResponseBodyData) *QueryTracedEventByEventIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetMessage(v string) *QueryTracedEventByEventIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetRequestId(v string) *QueryTracedEventByEventIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBody) SetSuccess(v bool) *QueryTracedEventByEventIdResponseBody {
	s.Success = &v
	return s
}

type QueryTracedEventByEventIdResponseBodyData struct {
	// The events.
	Events []*QueryTracedEventByEventIdResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetEvents(v []*QueryTracedEventByEventIdResponseBodyDataEvents) *QueryTracedEventByEventIdResponseBodyData {
	s.Events = v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetNextToken(v string) *QueryTracedEventByEventIdResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyData) SetTotal(v int32) *QueryTracedEventByEventIdResponseBodyData {
	s.Total = &v
	return s
}

type QueryTracedEventByEventIdResponseBodyDataEvents struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the event was delivered to the event bus.
	EventReceivedTime *int64 `json:"EventReceivedTime,omitempty" xml:"EventReceivedTime,omitempty"`
	// The name of the event source.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s QueryTracedEventByEventIdResponseBodyDataEvents) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventByEventIdResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventBusName(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventId(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventReceivedTime(v int64) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventReceivedTime = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventSource(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventByEventIdResponseBodyDataEvents) SetEventType(v string) *QueryTracedEventByEventIdResponseBodyDataEvents {
	s.EventType = &v
	return s
}

type QueryTracedEventByEventIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTracedEventByEventIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTracedEventByEventIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventByEventIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTracedEventByEventIdResponse) SetHeaders(v map[string]*string) *QueryTracedEventByEventIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTracedEventByEventIdResponse) SetStatusCode(v int32) *QueryTracedEventByEventIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTracedEventByEventIdResponse) SetBody(v *QueryTracedEventByEventIdResponseBody) *QueryTracedEventByEventIdResponse {
	s.Body = v
	return s
}

type QueryTracedEventsRequest struct {
	// The end of the time range when event traces are queried. Unit: milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. Up to 100 entries can be returned in a call.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event rule that is matched.
	MatchedRule *string `json:"MatchedRule,omitempty" xml:"MatchedRule,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query event traces. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryTracedEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventsRequest) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsRequest) SetEndTime(v int64) *QueryTracedEventsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventBusName(v string) *QueryTracedEventsRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventSource(v string) *QueryTracedEventsRequest {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventType(v string) *QueryTracedEventsRequest {
	s.EventType = &v
	return s
}

func (s *QueryTracedEventsRequest) SetLimit(v int32) *QueryTracedEventsRequest {
	s.Limit = &v
	return s
}

func (s *QueryTracedEventsRequest) SetMatchedRule(v string) *QueryTracedEventsRequest {
	s.MatchedRule = &v
	return s
}

func (s *QueryTracedEventsRequest) SetNextToken(v string) *QueryTracedEventsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventsRequest) SetStartTime(v int64) *QueryTracedEventsRequest {
	s.StartTime = &v
	return s
}

type QueryTracedEventsResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryTracedEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTracedEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBody) SetCode(v string) *QueryTracedEventsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetData(v *QueryTracedEventsResponseBodyData) *QueryTracedEventsResponseBody {
	s.Data = v
	return s
}

func (s *QueryTracedEventsResponseBody) SetMessage(v string) *QueryTracedEventsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetRequestId(v string) *QueryTracedEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTracedEventsResponseBody) SetSuccess(v bool) *QueryTracedEventsResponseBody {
	s.Success = &v
	return s
}

type QueryTracedEventsResponseBodyData struct {
	// The event type.
	Events []*QueryTracedEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// If excess return values exist, this parameter is returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryTracedEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBodyData) SetEvents(v []*QueryTracedEventsResponseBodyDataEvents) *QueryTracedEventsResponseBodyData {
	s.Events = v
	return s
}

func (s *QueryTracedEventsResponseBodyData) SetNextToken(v string) *QueryTracedEventsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventsResponseBodyData) SetTotal(v int32) *QueryTracedEventsResponseBodyData {
	s.Total = &v
	return s
}

type QueryTracedEventsResponseBodyDataEvents struct {
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time when the event was delivered to the event bus.
	EventReceivedTime *int64 `json:"EventReceivedTime,omitempty" xml:"EventReceivedTime,omitempty"`
	// The name of the event source.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s QueryTracedEventsResponseBodyDataEvents) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventsResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventBusName(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventId(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventId = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventReceivedTime(v int64) *QueryTracedEventsResponseBodyDataEvents {
	s.EventReceivedTime = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventSource(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventsResponseBodyDataEvents) SetEventType(v string) *QueryTracedEventsResponseBodyDataEvents {
	s.EventType = &v
	return s
}

type QueryTracedEventsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTracedEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTracedEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTracedEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponse) SetHeaders(v map[string]*string) *QueryTracedEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryTracedEventsResponse) SetStatusCode(v int32) *QueryTracedEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTracedEventsResponse) SetBody(v *QueryTracedEventsResponseBody) *QueryTracedEventsResponse {
	s.Body = v
	return s
}

type StartEventStreamingRequest struct {
	// The name of the event stream that you want to enable.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
}

func (s StartEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s StartEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *StartEventStreamingRequest) SetEventStreamingName(v string) *StartEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

type StartEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *StartEventStreamingResponseBody) SetCode(v string) *StartEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetMessage(v string) *StartEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetRequestId(v string) *StartEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetSuccess(v bool) *StartEventStreamingResponseBody {
	s.Success = &v
	return s
}

type StartEventStreamingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s StartEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *StartEventStreamingResponse) SetHeaders(v map[string]*string) *StartEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *StartEventStreamingResponse) SetStatusCode(v int32) *StartEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEventStreamingResponse) SetBody(v *StartEventStreamingResponseBody) *StartEventStreamingResponse {
	s.Body = v
	return s
}

type TestEventPatternRequest struct {
	// The event.
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The event pattern.
	EventPattern *string `json:"EventPattern,omitempty" xml:"EventPattern,omitempty"`
}

func (s TestEventPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s TestEventPatternRequest) GoString() string {
	return s.String()
}

func (s *TestEventPatternRequest) SetEvent(v string) *TestEventPatternRequest {
	s.Event = &v
	return s
}

func (s *TestEventPatternRequest) SetEventPattern(v string) *TestEventPatternRequest {
	s.EventPattern = &v
	return s
}

type TestEventPatternResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *TestEventPatternResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestEventPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestEventPatternResponseBody) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponseBody) SetCode(v string) *TestEventPatternResponseBody {
	s.Code = &v
	return s
}

func (s *TestEventPatternResponseBody) SetData(v *TestEventPatternResponseBodyData) *TestEventPatternResponseBody {
	s.Data = v
	return s
}

func (s *TestEventPatternResponseBody) SetMessage(v string) *TestEventPatternResponseBody {
	s.Message = &v
	return s
}

func (s *TestEventPatternResponseBody) SetRequestId(v string) *TestEventPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestEventPatternResponseBody) SetSuccess(v bool) *TestEventPatternResponseBody {
	s.Success = &v
	return s
}

type TestEventPatternResponseBodyData struct {
	// The value true indicates that the event pattern matches the provided JSON format. The value false indicates that the event pattern does not match the provided JSON format.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TestEventPatternResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TestEventPatternResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponseBodyData) SetResult(v bool) *TestEventPatternResponseBodyData {
	s.Result = &v
	return s
}

type TestEventPatternResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestEventPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestEventPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s TestEventPatternResponse) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponse) SetHeaders(v map[string]*string) *TestEventPatternResponse {
	s.Headers = v
	return s
}

func (s *TestEventPatternResponse) SetStatusCode(v int32) *TestEventPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *TestEventPatternResponse) SetBody(v *TestEventPatternResponseBody) *TestEventPatternResponse {
	s.Body = v
	return s
}

type UpdateApiDestinationRequest struct {
	// The name of the API destination. The name must be 2 to 127 characters in length.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// Note: Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	HttpApiParameters *UpdateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s UpdateApiDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationRequest) SetApiDestinationName(v string) *UpdateApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetConnectionName(v string) *UpdateApiDestinationRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetDescription(v string) *UpdateApiDestinationRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetHttpApiParameters(v *UpdateApiDestinationRequestHttpApiParameters) *UpdateApiDestinationRequest {
	s.HttpApiParameters = v
	return s
}

type UpdateApiDestinationRequestHttpApiParameters struct {
	// The endpoint of the API destination. The endpoint can be up to 127 characters in length.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// - GET
	// - POST
	// - HEAD
	// - DELETE
	// - PUT
	// - PATCH
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s UpdateApiDestinationRequestHttpApiParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiDestinationRequestHttpApiParameters) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationRequestHttpApiParameters) SetEndpoint(v string) *UpdateApiDestinationRequestHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *UpdateApiDestinationRequestHttpApiParameters) SetMethod(v string) *UpdateApiDestinationRequestHttpApiParameters {
	s.Method = &v
	return s
}

type UpdateApiDestinationShrinkRequest struct {
	// The name of the API destination. The name must be 2 to 127 characters in length.
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// Note: Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	HttpApiParametersShrink *string `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty"`
}

func (s UpdateApiDestinationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiDestinationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationShrinkRequest) SetApiDestinationName(v string) *UpdateApiDestinationShrinkRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetConnectionName(v string) *UpdateApiDestinationShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetDescription(v string) *UpdateApiDestinationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetHttpApiParametersShrink(v string) *UpdateApiDestinationShrinkRequest {
	s.HttpApiParametersShrink = &v
	return s
}

type UpdateApiDestinationResponseBody struct {
	// api-destination-name
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The response code. If the request is successful, Success is returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApiDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationResponseBody) SetApiDestinationName(v string) *UpdateApiDestinationResponseBody {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetCode(v string) *UpdateApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetMessage(v string) *UpdateApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetRequestId(v string) *UpdateApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApiDestinationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiDestinationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationResponse) SetHeaders(v map[string]*string) *UpdateApiDestinationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiDestinationResponse) SetStatusCode(v int32) *UpdateApiDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiDestinationResponse) SetBody(v *UpdateApiDestinationResponseBody) *UpdateApiDestinationResponse {
	s.Body = v
	return s
}

type UpdateConnectionRequest struct {
	// The parameters that are configured for authentication.
	AuthParameters *UpdateConnectionRequestAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The name of the connection that you want to update. The name must be 2 to 127 characters in length.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	NetworkParameters *UpdateConnectionRequestNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequest) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParameters struct {
	// The parameters for API key authentication.
	ApiKeyAuthParameters *UpdateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication type. Valid values:
	//
	// BASIC_AUTH: basic authentication.
	//
	// Introduction: Basic authentication is a simple authentication scheme built into the HTTP protocol. When you use the HTTP protocol for communications, the authentication method that the HTTP server uses to authenticate user identities on the client is defined in the protocol. The request header is in the Authorization: Basic Base64-encoded string (Username:Password) format.
	//
	// 1.  Username and Password are required.
	//
	// API_KEY_AUTH: API key authentication.
	//
	// Introduction: The request header is in the Token : Token value format.
	//
	// *   ApiKeyName and ApiKeyValue are required.
	//
	// OAUTH_AUTH: OAuth authentication.
	//
	// Introduction: OAuth2.0 is an authentication mechanism. In normal cases, a system that does not use OAuth2.0 can access the resources of the server from the client. To ensure access security, access tokens are used to identify users in OAuth 2.0. The client must use an access token to access protected resources. This way, OAuth 2.0 protects resources from being accessed from malicious clients and improves system security.
	//
	// *   AuthorizationEndpoint, OAuthHttpParameters, and HttpMethod are required.
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The parameters that are configured for basic authentication.
	BasicAuthParameters *UpdateConnectionRequestAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are configured for OAuth authentication.
	OAuthParameters *UpdateConnectionRequestAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequestAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersApiKeyAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *UpdateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *UpdateConnectionRequestAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

type UpdateConnectionRequestAuthParametersBasicAuthParameters struct {
	// The password for basic authentication.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username for basic authentication.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersBasicAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) SetPassword(v string) *UpdateConnectionRequestAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersBasicAuthParameters) SetUsername(v string) *UpdateConnectionRequestAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

type UpdateConnectionRequestAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token. The endpoint can be up to 127 characters in length.
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are configured for the client.
	ClientParameters *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// *   GET
	// *   POST
	// *   HEAD
	// *   DELETE
	// *   PUT
	// *   PATCH
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters for OAuth authentication.
	OAuthHttpParameters *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParametersOAuthParametersClientParameters struct {
	// The client ID.
	ClientID *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
	// The AccessKey secret of the client.
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientID(v string) *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The parameters that are configured for the request body.
	BodyParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The value of the request header.
	HeaderParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The parameters that are configured for the request path.
	QueryStringParameters []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request body.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request header.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Specifies whether to enable authentication.
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	// The key of the request path.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
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

type UpdateConnectionRequestNetworkParameters struct {
	// PublicNetwork: the Internet.
	//
	// PrivateNetwork: virtual private cloud (VPC).
	//
	// Note: If you set this parameter to PrivateNetwork, you must configure VpcId, VswitcheId, and SecurityGroupId.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s UpdateConnectionRequestNetworkParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequestNetworkParameters) GoString() string {
	return s.String()
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

type UpdateConnectionShrinkRequest struct {
	// The parameters that are configured for authentication.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection that you want to update. The name must be 2 to 127 characters in length.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
}

func (s UpdateConnectionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionShrinkRequest) SetAuthParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.AuthParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetConnectionName(v string) *UpdateConnectionShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetDescription(v string) *UpdateConnectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetNetworkParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.NetworkParametersShrink = &v
	return s
}

type UpdateConnectionResponseBody struct {
	// The returned response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponseBody) SetCode(v string) *UpdateConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetMessage(v string) *UpdateConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetRequestId(v string) *UpdateConnectionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConnectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponse) SetHeaders(v map[string]*string) *UpdateConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectionResponse) SetStatusCode(v int32) *UpdateConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectionResponse) SetBody(v *UpdateConnectionResponseBody) *UpdateConnectionResponse {
	s.Body = v
	return s
}

type UpdateEventBusRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s UpdateEventBusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventBusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventBusRequest) SetDescription(v string) *UpdateEventBusRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventBusRequest) SetEventBusName(v string) *UpdateEventBusRequest {
	s.EventBusName = &v
	return s
}

type UpdateEventBusResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventBusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventBusResponseBody) SetCode(v string) *UpdateEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetMessage(v string) *UpdateEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetRequestId(v string) *UpdateEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetSuccess(v bool) *UpdateEventBusResponseBody {
	s.Success = &v
	return s
}

type UpdateEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventBusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventBusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventBusResponse) SetHeaders(v map[string]*string) *UpdateEventBusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventBusResponse) SetStatusCode(v int32) *UpdateEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventBusResponse) SetBody(v *UpdateEventBusResponseBody) *UpdateEventBusResponse {
	s.Body = v
	return s
}

type UpdateEventSourceRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event bus with which the event source is associated.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParameters *UpdateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParameters *UpdateEventSourceRequestSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Service (MNS).
	SourceMNSParameters *UpdateEventSourceRequestSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParameters *UpdateEventSourceRequestSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParameters *UpdateEventSourceRequestSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// SourceSLSParameters
	SourceSLSParameters *UpdateEventSourceRequestSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParameters *UpdateEventSourceRequestSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
}

func (s UpdateEventSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequest) SetDescription(v string) *UpdateEventSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventSourceRequest) SetEventBusName(v string) *UpdateEventSourceRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventSourceRequest) SetEventSourceName(v string) *UpdateEventSourceRequest {
	s.EventSourceName = &v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceHttpEventParameters(v *UpdateEventSourceRequestSourceHttpEventParameters) *UpdateEventSourceRequest {
	s.SourceHttpEventParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceKafkaParameters(v *UpdateEventSourceRequestSourceKafkaParameters) *UpdateEventSourceRequest {
	s.SourceKafkaParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceMNSParameters(v *UpdateEventSourceRequestSourceMNSParameters) *UpdateEventSourceRequest {
	s.SourceMNSParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceRabbitMQParameters(v *UpdateEventSourceRequestSourceRabbitMQParameters) *UpdateEventSourceRequest {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceRocketMQParameters(v *UpdateEventSourceRequestSourceRocketMQParameters) *UpdateEventSourceRequest {
	s.SourceRocketMQParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceSLSParameters(v *UpdateEventSourceRequestSourceSLSParameters) *UpdateEventSourceRequest {
	s.SourceSLSParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceScheduledEventParameters(v *UpdateEventSourceRequestSourceScheduledEventParameters) *UpdateEventSourceRequest {
	s.SourceScheduledEventParameters = v
	return s
}

type UpdateEventSourceRequestSourceHttpEventParameters struct {
	// The CIDR block that is used for security settings. This parameter is required only if SecurityConfig is set to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request method supported by the generated webhook URL. You can select multiple values. Valid values:
	//
	// *   GET
	// *   POST
	// *   PUT
	// *   PATCH
	// *   DELETE
	// *   HEAD
	// *   OPTIONS
	// *   TRACE
	// *   CONNECT
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	// The security domain name. This parameter is required only if SecurityConfig is set to referer. You can enter a domain name.
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	// The type of security settings. Valid values:
	//
	// *   none: No configuration is required.
	// *   ip: CIDR block.
	// *   referer: security domain name.
	SecurityConfig *string `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	// The protocol type that is supported by the generated webhook URL. Valid values:
	//
	// *   HTTP
	// *   HTTPS
	// *   HTTP\&HTTPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateEventSourceRequestSourceHttpEventParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetIp(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetMethod(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetReferer(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetSecurityConfig(v string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetType(v string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Type = &v
	return s
}

type UpdateEventSourceRequestSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of consumers.
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a self-managed network.
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the Message Queue for Apache Kafka instance belongs. This parameter is required only if you set Network to PublicNetwork.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the Message Queue for Apache Kafka instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache Kafka instance is associated. This parameter is required only if you set Network to PublicNetwork.
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC in which the Message Queue for Apache Kafka instance resides. This parameter is required only if you set Network to PublicNetwork.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventSourceRequestSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetConsumerGroup(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetMaximumTasks(v int32) *UpdateEventSourceRequestSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetNetwork(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetOffsetReset(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetSecurityGroupId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetTopic(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetVSwitchIds(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetVpcId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type UpdateEventSourceRequestSourceMNSParameters struct {
	// Indicates whether Base64 decoding is enabled. By default, Base64 decoding is enabled.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region where the MNS queue resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEventSourceRequestSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v bool) *UpdateEventSourceRequestSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetQueueName(v string) *UpdateEventSourceRequestSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceMNSParameters {
	s.RegionId = &v
	return s
}

type UpdateEventSourceRequestSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance. For more information, see [Limits](~~163289~~).
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s UpdateEventSourceRequestSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type UpdateEventSourceRequestSourceRocketMQParameters struct {
	// The authentication type. You can set this parameter to ACL or leave this parameter empty.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// None.
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the Message Queue for Apache RocketMQ instance.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the Message Queue for Apache RocketMQ instance belongs.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The type of the Message Queue for Apache RocketMQ instance. Valid values:
	//
	// *   Cloud\_4: Message Queue for Apache RocketMQ 4.0 instance.
	// *   Cloud\_5: Message Queue for Apache RocketMQ 5.0 instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username that is used to access the Message Queue for Apache RocketMQ instance.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache RocketMQ instance is associated.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the Message Queue for Apache RocketMQ instance resides.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which message consumption starts. Valid values:
	//
	// *   CONSUME_FROM_LAST_OFFSET: Start message consumption from the latest offset.
	// *   CONSUME_FROM_FIRST_OFFSET: Start message consumption from the earliest offset.
	// *   CONSUME_FROM_TIMESTAMP: Start message consumption from the offset at the specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUME_FROM_TIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance. For more information, see [Limits](~~163289~~).
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetAuthType(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceEndpoint(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceNetwork(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstancePassword(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceUsername(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceVpcId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetOffset(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTag(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v int64) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type UpdateEventSourceRequestSourceSLSParameters struct {
	// The starting consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start consumption.
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Log Service Logstore.
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console. For information about the permission policy of this role, see Create a custom event source of the Log Service type.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventSourceRequestSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetConsumePosition(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetLogStore(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetProject(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetRoleName(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.RoleName = &v
	return s
}

type UpdateEventSourceRequestSourceScheduledEventParameters struct {
	// The cron expression.
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone in which the cron expression is executed.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateEventSourceRequestSourceScheduledEventParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetSchedule(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetTimeZone(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetUserData(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

type UpdateEventSourceShrinkRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event bus with which the event source is associated.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParametersShrink *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	// The parameters that are configured if the event source is Message Service (MNS).
	SourceMNSParametersShrink *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParametersShrink *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParametersShrink *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// SourceSLSParameters
	SourceSLSParametersShrink *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParametersShrink *string `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty"`
}

func (s UpdateEventSourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceShrinkRequest) SetDescription(v string) *UpdateEventSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetEventBusName(v string) *UpdateEventSourceShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetEventSourceName(v string) *UpdateEventSourceShrinkRequest {
	s.EventSourceName = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceHttpEventParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceHttpEventParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceKafkaParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceKafkaParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceMNSParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceMNSParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceRabbitMQParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceRabbitMQParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceRocketMQParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceRocketMQParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceSLSParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceSLSParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceScheduledEventParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceScheduledEventParametersShrink = &v
	return s
}

type UpdateEventSourceResponseBody struct {
	// The returned response code. Valid values:
	//
	// *   Success: The request is successful.
	// *   Other codes: The request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceResponseBody) SetCode(v string) *UpdateEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetData(v bool) *UpdateEventSourceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetMessage(v string) *UpdateEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetRequestId(v string) *UpdateEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetSuccess(v bool) *UpdateEventSourceResponseBody {
	s.Success = &v
	return s
}

type UpdateEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceResponse) SetHeaders(v map[string]*string) *UpdateEventSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventSourceResponse) SetStatusCode(v int32) *UpdateEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventSourceResponse) SetBody(v *UpdateEventSourceResponseBody) *UpdateEventSourceResponse {
	s.Body = v
	return s
}

type UpdateEventStreamingRequest struct {
	// The description of the event stream.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptions *UpdateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	// The event target. You must and can specify only one event target.
	Sink *UpdateEventStreamingRequestSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	// The event source, which is also known as the event source. You must and can specify only one event source.
	Source     *UpdateEventStreamingRequestSource       `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Transforms []*UpdateEventStreamingRequestTransforms `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
}

func (s UpdateEventStreamingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequest) SetDescription(v string) *UpdateEventStreamingRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetEventStreamingName(v string) *UpdateEventStreamingRequest {
	s.EventStreamingName = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetFilterPattern(v string) *UpdateEventStreamingRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateEventStreamingRequest) SetRunOptions(v *UpdateEventStreamingRequestRunOptions) *UpdateEventStreamingRequest {
	s.RunOptions = v
	return s
}

func (s *UpdateEventStreamingRequest) SetSink(v *UpdateEventStreamingRequestSink) *UpdateEventStreamingRequest {
	s.Sink = v
	return s
}

func (s *UpdateEventStreamingRequest) SetSource(v *UpdateEventStreamingRequestSource) *UpdateEventStreamingRequest {
	s.Source = v
	return s
}

func (s *UpdateEventStreamingRequest) SetTransforms(v []*UpdateEventStreamingRequestTransforms) *UpdateEventStreamingRequest {
	s.Transforms = v
	return s
}

type UpdateEventStreamingRequestRunOptions struct {
	// The information about the batch window.
	BatchWindow *UpdateEventStreamingRequestRunOptionsBatchWindow `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	// Specifies whether to enable dead-letter queues. By default, dead-letter queues are disabled. Messages that fail to be pushed are discarded after the maximum number of retries specified by the retry policy is reached.
	DeadLetterQueue *UpdateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	// The fault tolerance policy. The value NONE specifies that faults are not tolerated, and the value All specifies that all faults are tolerated.
	ErrorsTolerance *string `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	// The concurrency level.
	MaximumTasks *int64 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The information about the retry policy that is used if the event fails to be pushed.
	RetryStrategy *UpdateEventStreamingRequestRunOptionsRetryStrategy `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestRunOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptions) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptions) SetBatchWindow(v *UpdateEventStreamingRequestRunOptionsBatchWindow) *UpdateEventStreamingRequestRunOptions {
	s.BatchWindow = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetDeadLetterQueue(v *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) *UpdateEventStreamingRequestRunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetErrorsTolerance(v string) *UpdateEventStreamingRequestRunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetMaximumTasks(v int64) *UpdateEventStreamingRequestRunOptions {
	s.MaximumTasks = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptions) SetRetryStrategy(v *UpdateEventStreamingRequestRunOptionsRetryStrategy) *UpdateEventStreamingRequestRunOptions {
	s.RetryStrategy = v
	return s
}

type UpdateEventStreamingRequestRunOptionsBatchWindow struct {
	// The maximum number of events that is allowed in the batch window. If the value specified by this parameter is reached, the data in the batch window is pushed to the downstream application. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// The maximum period of time during which events are allowed in the batch window. Unit: seconds. If the value specified by this parameter is reached, the data in the batch window is pushed to the downstream application. If multiple batch windows exist, data is pushed if triggering conditions are met in one of the windows.
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsBatchWindow) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsBatchWindow) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) SetCountBasedWindow(v int32) *UpdateEventStreamingRequestRunOptionsBatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsBatchWindow) SetTimeBasedWindow(v int32) *UpdateEventStreamingRequestRunOptionsBatchWindow {
	s.TimeBasedWindow = &v
	return s
}

type UpdateEventStreamingRequestRunOptionsDeadLetterQueue struct {
	// The Alibaba Cloud Resource Name (ARN) of the dead-letter queue.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsDeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsDeadLetterQueue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsDeadLetterQueue) SetArn(v string) *UpdateEventStreamingRequestRunOptionsDeadLetterQueue {
	s.Arn = &v
	return s
}

type UpdateEventStreamingRequestRunOptionsRetryStrategy struct {
	// The maximum period of time during which retries are performed.
	MaximumEventAgeInSeconds *int64 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	// The maximum number of retries.
	MaximumRetryAttempts *int64 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	// The retry policy that is used if an event failed to be pushed. Valid values: BACKOFF_RETRY and EXPONENTIAL_DECAY_RETRY.
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s UpdateEventStreamingRequestRunOptionsRetryStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestRunOptionsRetryStrategy) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumEventAgeInSeconds(v int64) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumEventAgeInSeconds = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetMaximumRetryAttempts(v int64) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.MaximumRetryAttempts = &v
	return s
}

func (s *UpdateEventStreamingRequestRunOptionsRetryStrategy) SetPushRetryStrategy(v string) *UpdateEventStreamingRequestRunOptionsRetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

type UpdateEventStreamingRequestSink struct {
	// The parameters that are configured if the event target is Function Compute.
	SinkFcParameters  *UpdateEventStreamingRequestSinkSinkFcParameters  `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	SinkFnfParameters *UpdateEventStreamingRequestSinkSinkFnfParameters `json:"SinkFnfParameters,omitempty" xml:"SinkFnfParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event target is Message Queue for Apache Kafka.
	SinkKafkaParameters *UpdateEventStreamingRequestSinkSinkKafkaParameters `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event target is MNS.
	SinkMNSParameters        *UpdateEventStreamingRequestSinkSinkMNSParameters        `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkPrometheusParameters *UpdateEventStreamingRequestSinkSinkPrometheusParameters `json:"SinkPrometheusParameters,omitempty" xml:"SinkPrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event target is Message Queue for RabbitMQ.
	SinkRabbitMQParameters *UpdateEventStreamingRequestSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Parameters
	SinkRocketMQParameters *UpdateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// Sink SLS Parameters
	SinkSLSParameters *UpdateEventStreamingRequestSinkSinkSLSParameters `json:"SinkSLSParameters,omitempty" xml:"SinkSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSink) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSink) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSink) SetSinkFcParameters(v *UpdateEventStreamingRequestSinkSinkFcParameters) *UpdateEventStreamingRequestSink {
	s.SinkFcParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkFnfParameters(v *UpdateEventStreamingRequestSinkSinkFnfParameters) *UpdateEventStreamingRequestSink {
	s.SinkFnfParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkKafkaParameters(v *UpdateEventStreamingRequestSinkSinkKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkMNSParameters(v *UpdateEventStreamingRequestSinkSinkMNSParameters) *UpdateEventStreamingRequestSink {
	s.SinkMNSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkPrometheusParameters(v *UpdateEventStreamingRequestSinkSinkPrometheusParameters) *UpdateEventStreamingRequestSink {
	s.SinkPrometheusParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRabbitMQParameters(v *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) *UpdateEventStreamingRequestSink {
	s.SinkRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkRocketMQParameters(v *UpdateEventStreamingRequestSinkSinkRocketMQParameters) *UpdateEventStreamingRequestSink {
	s.SinkRocketMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkSLSParameters(v *UpdateEventStreamingRequestSinkSinkSLSParameters) *UpdateEventStreamingRequestSink {
	s.SinkSLSParameters = v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParameters struct {
	// The message body that is sent to the function.
	Body *UpdateEventStreamingRequestSinkSinkFcParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The information about the delivery concurrency.
	Concurrency *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	// The information about the Function Compute function.
	FunctionName *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	// The information about the invocation type.
	InvocationType *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	// The information about the service to which the function belongs.
	Qualifier *UpdateEventStreamingRequestSinkSinkFcParametersQualifier `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	// The information about the Function Compute service.
	ServiceName *UpdateEventStreamingRequestSinkSinkFcParametersServiceName `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkFcParametersBody) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetConcurrency(v *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Concurrency = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetFunctionName(v *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.FunctionName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetInvocationType(v *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.InvocationType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetQualifier(v *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.Qualifier = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParameters) SetServiceName(v *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) *UpdateEventStreamingRequestSinkSinkFcParameters {
	s.ServiceName = v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersBody {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersConcurrency struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The delivery concurrency. Minimum value: 1.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersFunctionName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the Function Compute function.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersInvocationType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The invocation type. Valid values: Sync and Async.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersQualifier struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The alias of the service to which the function belongs.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersQualifier) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersQualifier) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersQualifier) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersQualifier {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFcParametersServiceName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the Function Compute service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersServiceName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFcParametersServiceName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFcParametersServiceName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFcParametersServiceName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFnfParameters struct {
	ExecutionName *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty" type:"Struct"`
	FlowName      *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName      `json:"FlowName,omitempty" xml:"FlowName,omitempty" type:"Struct"`
	Input         *UpdateEventStreamingRequestSinkSinkFnfParametersInput         `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	RoleName      *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName      `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetExecutionName(v *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.ExecutionName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetFlowName(v *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.FlowName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetInput(v *UpdateEventStreamingRequestSinkSinkFnfParametersInput) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.Input = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParameters) SetRoleName(v *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) *UpdateEventStreamingRequestSinkSinkFnfParameters {
	s.RoleName = v
	return s
}

type UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersExecutionName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFnfParametersFlowName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersFlowName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFnfParametersInput struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersInput) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersInput) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersInput {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkFnfParametersRoleName struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkFnfParametersRoleName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParameters struct {
	// The information about the acknowledgment (ACK) mode. If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high. If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader. If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Acks *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	// The information about the Message Queue for Apache Kafka instance.
	InstanceId *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The information about the message key.
	Key *UpdateEventStreamingRequestSinkSinkKafkaParametersKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// The information about the topic in Message Queue for Apache Kafka instance.
	Topic *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	// The information about the message value.
	Value *UpdateEventStreamingRequestSinkSinkKafkaParametersValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetAcks(v *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Acks = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetKey(v *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Key = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Topic = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParameters) SetValue(v *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) *UpdateEventStreamingRequestSinkSinkKafkaParameters {
	s.Value = v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersAcks struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ACK mode. If you set this parameter to 0, no response is returned from the broker. In this mode, the performance is high, but the risk of data loss is also high. If you set this parameter to 1, a response is returned when data is written to the leader. In this mode, the performance and the risk of data loss are moderate. Data loss may occur if a failure occurs on the leader. If you set this parameter to all, a response is returned when data is written to the leader and synchronized to the followers. In this mode, the performance is low, but the risk of data loss is also low. Data loss occurs if the leader and the followers fail at the same time.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The message key.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersKey) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersKey {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in Message Queue for Apache Kafka instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkKafkaParametersValue struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersValue) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkKafkaParametersValue) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetForm(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkKafkaParametersValue) SetValue(v string) *UpdateEventStreamingRequestSinkSinkKafkaParametersValue {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkMNSParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkMNSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// Specifies whether to enable Base64 encoding.
	IsBase64Encode *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	// The information about the MNS queue.
	QueueName *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkMNSParametersBody) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetIsBase64Encode(v *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.IsBase64Encode = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParameters) SetQueueName(v *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) *UpdateEventStreamingRequestSinkSinkMNSParameters {
	s.QueueName = v
	return s
}

type UpdateEventStreamingRequestSinkSinkMNSParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersBody {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// Specifies that Base64 encoding is enabled.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkMNSParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue in MNS.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParameters struct {
	AuthorizationType *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty" type:"Struct"`
	Data              *UpdateEventStreamingRequestSinkSinkPrometheusParametersData              `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NetworkType       *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType       `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" type:"Struct"`
	Password          *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword          `json:"Password,omitempty" xml:"Password,omitempty" type:"Struct"`
	SecurityGroupId   *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	URL               *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL               `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	Username          *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername          `json:"Username,omitempty" xml:"Username,omitempty" type:"Struct"`
	VSwitchId         *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Struct"`
	VpcId             *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId             `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetAuthorizationType(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.AuthorizationType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetData(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Data = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetNetworkType(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.NetworkType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetPassword(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Password = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetSecurityGroupId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.SecurityGroupId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetURL(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.URL = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetUsername(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.Username = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetVSwitchId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VSwitchId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParameters) SetVpcId(v *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) *UpdateEventStreamingRequestSinkSinkPrometheusParameters {
	s.VpcId = v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersAuthorizationType {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersData struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersData) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersData) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersData) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersData {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersNetworkType {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersPassword {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersSecurityGroupId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersURL struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersURL {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersUsername {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVSwitchId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkPrometheusParametersVpcId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The information about the exchange to which events are delivered. This parameter is available only if you set TargetType to Exchange.
	Exchange *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	// The information about the Message Queue for RabbitMQ instance.
	InstanceId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The message ID.
	MessageId *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Properties *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The information about the queue to which events are delivered. This parameter is available only if you set TargetType to Queue.
	QueueName *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	// The information about the routing rule of the message. This parameter is available only if you set TargetType to Exchange.
	RoutingKey *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	// The information about the type of the resource to which events are delivered.
	TargetType *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
	// The information about the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetExchange(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Exchange = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetMessageId(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.MessageId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetProperties(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.Properties = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetQueueName(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.QueueName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetRoutingKey(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.RoutingKey = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetTargetType(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.TargetType = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParameters) SetVirtualHostName(v *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) *UpdateEventStreamingRequestSinkSinkRabbitMQParameters {
	s.VirtualHostName = v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the exchange in the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the queue in the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The routing rule of the message.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The type of the resource to which events are delivered. Valid values: Exchange: exchanges. Queue: queues.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The vhost name of the Message Queue for RabbitMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRabbitMQParametersVirtualHostName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParameters struct {
	// The message content.
	Body *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The parameters that are configured if the event target is Message Queue for Apache RocketMQ.
	InstanceId *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Keys *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Properties *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The properties that are used to filter messages.
	Tags *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The information about the topic in the Message Queue for Apache RocketMQ instance.
	Topic *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetInstanceId(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.InstanceId = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetKeys(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Keys = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetProperties(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Properties = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetTags(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Tags = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) *UpdateEventStreamingRequestSinkSinkRocketMQParameters {
	s.Topic = v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersTags struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in the Message Queue for Apache RocketMQ instance.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParameters struct {
	// The message body that is sent to Log Service.
	Body *UpdateEventStreamingRequestSinkSinkSLSParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The information about the Log Service Logstore.
	LogStore *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	// The information about the Log Service project.
	Project *UpdateEventStreamingRequestSinkSinkSLSParametersProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The information about the role. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	RoleName *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	// The information about the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	Topic *UpdateEventStreamingRequestSinkSinkSLSParametersTopic `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetBody(v *UpdateEventStreamingRequestSinkSinkSLSParametersBody) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetLogStore(v *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.LogStore = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetProject(v *UpdateEventStreamingRequestSinkSinkSLSParametersProject) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Project = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetRoleName(v *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.RoleName = v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParameters) SetTopic(v *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) *UpdateEventStreamingRequestSinkSinkSLSParameters {
	s.Topic = v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParametersBody struct {
	// The method that is used to transform events.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template based on which events are transformed.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The value before event transformation.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersBody) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersBody {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParametersLogStore struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service Logstore.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParametersProject struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The Log Service project.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersProject) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersProject) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersProject) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersProject {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParametersRoleName struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the RAM console.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSinkSinkSLSParametersTopic struct {
	// The method that is used to transform events. Default value: CONSTANT.
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// None.
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the topic in which logs are stored. The topic corresponds to the topic reserved field in Log Service.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersTopic) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSinkSinkSLSParametersTopic) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetForm(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Form = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetTemplate(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Template = &v
	return s
}

func (s *UpdateEventStreamingRequestSinkSinkSLSParametersTopic) SetValue(v string) *UpdateEventStreamingRequestSinkSinkSLSParametersTopic {
	s.Value = &v
	return s
}

type UpdateEventStreamingRequestSource struct {
	// The parameters that are configured if you specify Data Transmission Service (DTS) as the event source.
	SourceDTSParameters *UpdateEventStreamingRequestSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for Apache Kafka as the event source.
	SourceKafkaParameters *UpdateEventStreamingRequestSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Service (MNS) as the event source.
	SourceMNSParameters *UpdateEventStreamingRequestSourceSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for MQTT as the event source.
	SourceMQTTParameters       *UpdateEventStreamingRequestSourceSourceMQTTParameters       `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourcePrometheusParameters *UpdateEventStreamingRequestSourceSourcePrometheusParameters `json:"SourcePrometheusParameters,omitempty" xml:"SourcePrometheusParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for RabbitMQ as the event source.
	SourceRabbitMQParameters *UpdateEventStreamingRequestSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Message Queue for Apache RocketMQ as the event source.
	SourceRocketMQParameters *UpdateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Log Service as the event source.
	SourceSLSParameters *UpdateEventStreamingRequestSourceSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
}

func (s UpdateEventStreamingRequestSource) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSource) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSource) SetSourceDTSParameters(v *UpdateEventStreamingRequestSourceSourceDTSParameters) *UpdateEventStreamingRequestSource {
	s.SourceDTSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceKafkaParameters(v *UpdateEventStreamingRequestSourceSourceKafkaParameters) *UpdateEventStreamingRequestSource {
	s.SourceKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceMNSParameters(v *UpdateEventStreamingRequestSourceSourceMNSParameters) *UpdateEventStreamingRequestSource {
	s.SourceMNSParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceMQTTParameters(v *UpdateEventStreamingRequestSourceSourceMQTTParameters) *UpdateEventStreamingRequestSource {
	s.SourceMQTTParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourcePrometheusParameters(v *UpdateEventStreamingRequestSourceSourcePrometheusParameters) *UpdateEventStreamingRequestSource {
	s.SourcePrometheusParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRabbitMQParameters(v *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) *UpdateEventStreamingRequestSource {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceRocketMQParameters(v *UpdateEventStreamingRequestSourceSourceRocketMQParameters) *UpdateEventStreamingRequestSource {
	s.SourceRocketMQParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSource) SetSourceSLSParameters(v *UpdateEventStreamingRequestSourceSourceSLSParameters) *UpdateEventStreamingRequestSource {
	s.SourceSLSParameters = v
	return s
}

type UpdateEventStreamingRequestSourceSourceDTSParameters struct {
	// The URL and port number of the change tracking instance.
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// The UNIX timestamp that is generated when the SDK client consumes the first data record.
	InitCheckPoint *int64 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// The consumer group password.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The consumer group ID.
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the tracked topic of the change tracking instance.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The consumer group username.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceDTSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceDTSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetBrokerUrl(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetInitCheckPoint(v int64) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetPassword(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Password = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetSid(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetTaskId(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceDTSParameters) SetUsername(v string) *UpdateEventStreamingRequestSourceSourceDTSParameters {
	s.Username = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceKafkaParameters struct {
	// The group ID of the consumer that subscribes to the topic.
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network setting. Default value: Default. The value PublicNetwork specifies a virtual private cloud (VPC).
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The offset.
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The vSwitch ID.
	VSwitchIds    *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ValueDataType *string `json:"ValueDataType,omitempty" xml:"ValueDataType,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetConsumerGroup(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetNetwork(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetOffsetReset(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetValueDataType(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.ValueDataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceMNSParameters struct {
	// Specifies whether to enable Base64 encoding. Default value: true.
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the MNS queue.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the MNS queue resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetIsBase64Decode(v bool) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetQueueName(v string) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMNSParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceMNSParameters {
	s.RegionId = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceMQTTParameters struct {
	// The ID of the Message Queue for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Message Queue for MQTT instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceMQTTParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceMQTTParameters {
	s.Topic = &v
	return s
}

type UpdateEventStreamingRequestSourceSourcePrometheusParameters struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourcePrometheusParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourcePrometheusParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetClusterId(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.ClusterId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetDataType(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.DataType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourcePrometheusParameters) SetLabels(v string) *UpdateEventStreamingRequestSourceSourcePrometheusParameters {
	s.Labels = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance.
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventStreamingRequestSourceSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceRocketMQParameters struct {
	// The authentication method.
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The instance endpoint.
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// PublicNetwork and PrivateNetwork.
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The instance password.
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The security group ID of the instance.
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance username.
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The vSwitch ID of the instance.
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The VPC ID of the instance.
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// *   CONSUMEFROMLASTOFFSET: Messages are consumed from the latest offset.
	// *   CONSUMEFROMFIRSTOFFSET: Messages are consumed from the earliest offset.
	// *   CONSUMEFROMTIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUMEFROMLASTOFFSET.
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the Message Queue for Apache RocketMQ instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that you want to use to filter messages.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUMEFROMTIMESTAMP.
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetAuthType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetGroupID(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceEndpoint(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceNetwork(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstancePassword(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceType(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceUsername(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceVpcId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetOffset(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetRegionId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTag(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTimestamp(v int64) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetTopic(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceSLSParameters struct {
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceSLSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceSLSParameters) SetRoleName(v string) *UpdateEventStreamingRequestSourceSourceSLSParameters {
	s.RoleName = &v
	return s
}

type UpdateEventStreamingRequestTransforms struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s UpdateEventStreamingRequestTransforms) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestTransforms) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestTransforms) SetArn(v string) *UpdateEventStreamingRequestTransforms {
	s.Arn = &v
	return s
}

type UpdateEventStreamingShrinkRequest struct {
	// The description of the event stream.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event stream.
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// The rule that is used to filter events. If you leave this parameter empty, all events are matched.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The parameters that are configured for the runtime environment.
	RunOptionsShrink *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	// The event target. You must and can specify only one event target.
	SinkShrink *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	// The event source, which is also known as the event source. You must and can specify only one event source.
	SourceShrink     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TransformsShrink *string `json:"Transforms,omitempty" xml:"Transforms,omitempty"`
}

func (s UpdateEventStreamingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingShrinkRequest) SetDescription(v string) *UpdateEventStreamingShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetEventStreamingName(v string) *UpdateEventStreamingShrinkRequest {
	s.EventStreamingName = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetFilterPattern(v string) *UpdateEventStreamingShrinkRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetRunOptionsShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.RunOptionsShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetSinkShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.SinkShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetSourceShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *UpdateEventStreamingShrinkRequest) SetTransformsShrink(v string) *UpdateEventStreamingShrinkRequest {
	s.TransformsShrink = &v
	return s
}

type UpdateEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingResponseBody) SetCode(v string) *UpdateEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetMessage(v string) *UpdateEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetRequestId(v string) *UpdateEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetSuccess(v bool) *UpdateEventStreamingResponseBody {
	s.Success = &v
	return s
}

type UpdateEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventStreamingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingResponse) SetHeaders(v map[string]*string) *UpdateEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventStreamingResponse) SetStatusCode(v int32) *UpdateEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventStreamingResponse) SetBody(v *UpdateEventStreamingResponseBody) *UpdateEventStreamingResponse {
	s.Body = v
	return s
}

type UpdateRuleRequest struct {
	// The description of the event bus.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus.
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event pattern, in JSON format. Valid values: stringEqual stringExpression Each field can have a maximum of five expressions in the map data structure.
	//
	// Each field can have a maximum of five expressions in the map data structure.
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	// The name of the event rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event rule. Valid values: ENABLE: The event rule is enabled. It is the default state of the event rule. DISABLE: The event rule is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) SetDescription(v string) *UpdateRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRuleRequest) SetEventBusName(v string) *UpdateRuleRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateRuleRequest) SetFilterPattern(v string) *UpdateRuleRequest {
	s.FilterPattern = &v
	return s
}

func (s *UpdateRuleRequest) SetRuleName(v string) *UpdateRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleRequest) SetStatus(v string) *UpdateRuleRequest {
	s.Status = &v
	return s
}

type UpdateRuleResponseBody struct {
	// The status code returned. The status code 200 indicates that the request was successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v bool) *UpdateRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetSuccess(v bool) *UpdateRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetStatusCode(v int32) *UpdateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eventbridge"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * You can call this API operation to create an API destination.
 *
 * @param tmpReq CreateApiDestinationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateApiDestinationResponse
 */
func (client *Client) CreateApiDestinationWithOptions(tmpReq *CreateApiDestinationRequest, runtime *util.RuntimeOptions) (_result *CreateApiDestinationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateApiDestinationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HttpApiParameters)) {
		request.HttpApiParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpApiParameters, tea.String("HttpApiParameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDestinationName)) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HttpApiParametersShrink)) {
		query["HttpApiParameters"] = request.HttpApiParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiDestination"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create an API destination.
 *
 * @param request CreateApiDestinationRequest
 * @return CreateApiDestinationResponse
 */
func (client *Client) CreateApiDestination(request *CreateApiDestinationRequest) (_result *CreateApiDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiDestinationResponse{}
	_body, _err := client.CreateApiDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create a connection.
 *
 * @param tmpReq CreateConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateConnectionResponse
 */
func (client *Client) CreateConnectionWithOptions(tmpReq *CreateConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateConnectionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthParameters)) {
		request.AuthParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthParameters, tea.String("AuthParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NetworkParameters)) {
		request.NetworkParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkParameters, tea.String("NetworkParameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthParametersShrink)) {
		query["AuthParameters"] = request.AuthParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkParametersShrink)) {
		query["NetworkParameters"] = request.NetworkParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConnection"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create a connection.
 *
 * @param request CreateConnectionRequest
 * @return CreateConnectionResponse
 */
func (client *Client) CreateConnection(request *CreateConnectionRequest) (_result *CreateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CreateConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create an event bus.
 *
 * @param request CreateEventBusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateEventBusResponse
 */
func (client *Client) CreateEventBusWithOptions(request *CreateEventBusRequest, runtime *util.RuntimeOptions) (_result *CreateEventBusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventBus"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventBusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create an event bus.
 *
 * @param request CreateEventBusRequest
 * @return CreateEventBusResponse
 */
func (client *Client) CreateEventBus(request *CreateEventBusRequest) (_result *CreateEventBusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventBusResponse{}
	_body, _err := client.CreateEventBusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to create an event source.
 *
 * @param tmpReq CreateEventSourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateEventSourceResponse
 */
func (client *Client) CreateEventSourceWithOptions(tmpReq *CreateEventSourceRequest, runtime *util.RuntimeOptions) (_result *CreateEventSourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEventSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SourceHttpEventParameters)) {
		request.SourceHttpEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceHttpEventParameters, tea.String("SourceHttpEventParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceKafkaParameters)) {
		request.SourceKafkaParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceKafkaParameters, tea.String("SourceKafkaParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceMNSParameters)) {
		request.SourceMNSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMNSParameters, tea.String("SourceMNSParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceRabbitMQParameters)) {
		request.SourceRabbitMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRabbitMQParameters, tea.String("SourceRabbitMQParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceRocketMQParameters)) {
		request.SourceRocketMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRocketMQParameters, tea.String("SourceRocketMQParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceSLSParameters)) {
		request.SourceSLSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSLSParameters, tea.String("SourceSLSParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceScheduledEventParameters)) {
		request.SourceScheduledEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceScheduledEventParameters, tea.String("SourceScheduledEventParameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		body["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventSourceName)) {
		body["EventSourceName"] = request.EventSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceHttpEventParametersShrink)) {
		body["SourceHttpEventParameters"] = request.SourceHttpEventParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceKafkaParametersShrink)) {
		body["SourceKafkaParameters"] = request.SourceKafkaParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceMNSParametersShrink)) {
		body["SourceMNSParameters"] = request.SourceMNSParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRabbitMQParametersShrink)) {
		body["SourceRabbitMQParameters"] = request.SourceRabbitMQParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRocketMQParametersShrink)) {
		body["SourceRocketMQParameters"] = request.SourceRocketMQParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSLSParametersShrink)) {
		body["SourceSLSParameters"] = request.SourceSLSParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceScheduledEventParametersShrink)) {
		body["SourceScheduledEventParameters"] = request.SourceScheduledEventParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventSource"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to create an event source.
 *
 * @param request CreateEventSourceRequest
 * @return CreateEventSourceResponse
 */
func (client *Client) CreateEventSource(request *CreateEventSourceRequest) (_result *CreateEventSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventSourceResponse{}
	_body, _err := client.CreateEventSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create an event stream.
 *
 * @param tmpReq CreateEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateEventStreamingResponse
 */
func (client *Client) CreateEventStreamingWithOptions(tmpReq *CreateEventStreamingRequest, runtime *util.RuntimeOptions) (_result *CreateEventStreamingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEventStreamingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RunOptions)) {
		request.RunOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RunOptions, tea.String("RunOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sink)) {
		request.SinkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sink, tea.String("Sink"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Source)) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, tea.String("Source"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Transforms)) {
		request.TransformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transforms, tea.String("Transforms"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterPattern)) {
		body["FilterPattern"] = request.FilterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.RunOptionsShrink)) {
		body["RunOptions"] = request.RunOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SinkShrink)) {
		body["Sink"] = request.SinkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceShrink)) {
		body["Source"] = request.SourceShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TransformsShrink)) {
		body["Transforms"] = request.TransformsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create an event stream.
 *
 * @param request CreateEventStreamingRequest
 * @return CreateEventStreamingResponse
 */
func (client *Client) CreateEventStreaming(request *CreateEventStreamingRequest) (_result *CreateEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventStreamingResponse{}
	_body, _err := client.CreateEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create an event rule.
 *
 * @param tmpReq CreateRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRuleResponse
 */
func (client *Client) CreateRuleWithOptions(tmpReq *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EventTargets)) {
		request.EventTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventTargets, tea.String("EventTargets"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventTargetsShrink)) {
		query["EventTargets"] = request.EventTargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FilterPattern)) {
		query["FilterPattern"] = request.FilterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create an event rule.
 *
 * @param request CreateRuleRequest
 * @return CreateRuleResponse
 */
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create a service-linked role for your cloud service.
 *
 * @param request CreateServiceLinkedRoleForProductRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateServiceLinkedRoleForProductResponse
 */
func (client *Client) CreateServiceLinkedRoleForProductWithOptions(request *CreateServiceLinkedRoleForProductRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleForProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRoleForProduct"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleForProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create a service-linked role for your cloud service.
 *
 * @param request CreateServiceLinkedRoleForProductRequest
 * @return CreateServiceLinkedRoleForProductResponse
 */
func (client *Client) CreateServiceLinkedRoleForProduct(request *CreateServiceLinkedRoleForProductRequest) (_result *CreateServiceLinkedRoleForProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleForProductResponse{}
	_body, _err := client.CreateServiceLinkedRoleForProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete an API destination.
 *
 * @param request DeleteApiDestinationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteApiDestinationResponse
 */
func (client *Client) DeleteApiDestinationWithOptions(request *DeleteApiDestinationRequest, runtime *util.RuntimeOptions) (_result *DeleteApiDestinationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDestinationName)) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiDestination"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApiDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete an API destination.
 *
 * @param request DeleteApiDestinationRequest
 * @return DeleteApiDestinationResponse
 */
func (client *Client) DeleteApiDestination(request *DeleteApiDestinationRequest) (_result *DeleteApiDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApiDestinationResponse{}
	_body, _err := client.DeleteApiDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete a connection.
 *
 * @param request DeleteConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteConnectionResponse
 */
func (client *Client) DeleteConnectionWithOptions(request *DeleteConnectionRequest, runtime *util.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConnection"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete a connection.
 *
 * @param request DeleteConnectionRequest
 * @return DeleteConnectionResponse
 */
func (client *Client) DeleteConnection(request *DeleteConnectionRequest) (_result *DeleteConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.DeleteConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete an event bus.
 *
 * @param request DeleteEventBusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteEventBusResponse
 */
func (client *Client) DeleteEventBusWithOptions(request *DeleteEventBusRequest, runtime *util.RuntimeOptions) (_result *DeleteEventBusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventBus"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventBusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete an event bus.
 *
 * @param request DeleteEventBusRequest
 * @return DeleteEventBusResponse
 */
func (client *Client) DeleteEventBus(request *DeleteEventBusRequest) (_result *DeleteEventBusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventBusResponse{}
	_body, _err := client.DeleteEventBusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete an event source.
 *
 * @param request DeleteEventSourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteEventSourceResponse
 */
func (client *Client) DeleteEventSourceWithOptions(request *DeleteEventSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteEventSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventSourceName)) {
		body["EventSourceName"] = request.EventSourceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventSource"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete an event source.
 *
 * @param request DeleteEventSourceRequest
 * @return DeleteEventSourceResponse
 */
func (client *Client) DeleteEventSource(request *DeleteEventSourceRequest) (_result *DeleteEventSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventSourceResponse{}
	_body, _err := client.DeleteEventSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete an event stream.
 *
 * @param request DeleteEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteEventStreamingResponse
 */
func (client *Client) DeleteEventStreamingWithOptions(request *DeleteEventStreamingRequest, runtime *util.RuntimeOptions) (_result *DeleteEventStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete an event stream.
 *
 * @param request DeleteEventStreamingRequest
 * @return DeleteEventStreamingResponse
 */
func (client *Client) DeleteEventStreaming(request *DeleteEventStreamingRequest) (_result *DeleteEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventStreamingResponse{}
	_body, _err := client.DeleteEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete an event rule.
 *
 * @param request DeleteRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteRuleResponse
 */
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete an event rule.
 *
 * @param request DeleteRuleRequest
 * @return DeleteRuleResponse
 */
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to delete one or more event targets of an event rule.
 *
 * @param tmpReq DeleteTargetsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTargetsResponse
 */
func (client *Client) DeleteTargetsWithOptions(tmpReq *DeleteTargetsRequest, runtime *util.RuntimeOptions) (_result *DeleteTargetsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteTargetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TargetIds)) {
		request.TargetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetIds, tea.String("TargetIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIdsShrink)) {
		query["TargetIds"] = request.TargetIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTargets"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to delete one or more event targets of an event rule.
 *
 * @param request DeleteTargetsRequest
 * @return DeleteTargetsResponse
 */
func (client *Client) DeleteTargets(request *DeleteTargetsRequest) (_result *DeleteTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTargetsResponse{}
	_body, _err := client.DeleteTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to disable an event rule.
 *
 * @param request DisableRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableRuleResponse
 */
func (client *Client) DisableRuleWithOptions(request *DisableRuleRequest, runtime *util.RuntimeOptions) (_result *DisableRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to disable an event rule.
 *
 * @param request DisableRuleRequest
 * @return DisableRuleResponse
 */
func (client *Client) DisableRule(request *DisableRuleRequest) (_result *DisableRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableRuleResponse{}
	_body, _err := client.DisableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to enable an event rule.
 *
 * @param request EnableRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return EnableRuleResponse
 */
func (client *Client) EnableRuleWithOptions(request *EnableRuleRequest, runtime *util.RuntimeOptions) (_result *EnableRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to enable an event rule.
 *
 * @param request EnableRuleRequest
 * @return EnableRuleResponse
 */
func (client *Client) EnableRule(request *EnableRuleRequest) (_result *EnableRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRuleResponse{}
	_body, _err := client.EnableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the information about an API destination.
 *
 * @param request GetApiDestinationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetApiDestinationResponse
 */
func (client *Client) GetApiDestinationWithOptions(request *GetApiDestinationRequest, runtime *util.RuntimeOptions) (_result *GetApiDestinationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDestinationName)) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApiDestination"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApiDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the information about an API destination.
 *
 * @param request GetApiDestinationRequest
 * @return GetApiDestinationResponse
 */
func (client *Client) GetApiDestination(request *GetApiDestinationRequest) (_result *GetApiDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApiDestinationResponse{}
	_body, _err := client.GetApiDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the configurations of a connection.
 *
 * @param request GetConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetConnectionResponse
 */
func (client *Client) GetConnectionWithOptions(request *GetConnectionRequest, runtime *util.RuntimeOptions) (_result *GetConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnection"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the configurations of a connection.
 *
 * @param request GetConnectionRequest
 * @return GetConnectionResponse
 */
func (client *Client) GetConnection(request *GetConnectionRequest) (_result *GetConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionResponse{}
	_body, _err := client.GetConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the detailed information about an event bus.
 *
 * @param request GetEventBusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetEventBusResponse
 */
func (client *Client) GetEventBusWithOptions(request *GetEventBusRequest, runtime *util.RuntimeOptions) (_result *GetEventBusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEventBus"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEventBusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the detailed information about an event bus.
 *
 * @param request GetEventBusRequest
 * @return GetEventBusResponse
 */
func (client *Client) GetEventBus(request *GetEventBusRequest) (_result *GetEventBusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEventBusResponse{}
	_body, _err := client.GetEventBusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the details of an event stream.
 *
 * @param request GetEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetEventStreamingResponse
 */
func (client *Client) GetEventStreamingWithOptions(request *GetEventStreamingRequest, runtime *util.RuntimeOptions) (_result *GetEventStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the details of an event stream.
 *
 * @param request GetEventStreamingRequest
 * @return GetEventStreamingResponse
 */
func (client *Client) GetEventStreaming(request *GetEventStreamingRequest) (_result *GetEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEventStreamingResponse{}
	_body, _err := client.GetEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the details of an event rule.
 *
 * @param request GetRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetRuleResponse
 */
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *util.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the details of an event rule.
 *
 * @param request GetRuleRequest
 * @return GetRuleResponse
 */
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query all Alibaba Cloud service event sources.
 *
 * @param request ListAliyunOfficialEventSourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAliyunOfficialEventSourcesResponse
 */
func (client *Client) ListAliyunOfficialEventSourcesWithOptions(runtime *util.RuntimeOptions) (_result *ListAliyunOfficialEventSourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListAliyunOfficialEventSources"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAliyunOfficialEventSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query all Alibaba Cloud service event sources.
 *
 * @return ListAliyunOfficialEventSourcesResponse
 */
func (client *Client) ListAliyunOfficialEventSources() (_result *ListAliyunOfficialEventSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliyunOfficialEventSourcesResponse{}
	_body, _err := client.ListAliyunOfficialEventSourcesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this API operation to query a list of API destinations.
 *
 * @param request ListApiDestinationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListApiDestinationsResponse
 */
func (client *Client) ListApiDestinationsWithOptions(request *ListApiDestinationsRequest, runtime *util.RuntimeOptions) (_result *ListApiDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDestinationNamePrefix)) {
		query["ApiDestinationNamePrefix"] = request.ApiDestinationNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApiDestinations"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApiDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this API operation to query a list of API destinations.
 *
 * @param request ListApiDestinationsRequest
 * @return ListApiDestinationsResponse
 */
func (client *Client) ListApiDestinations(request *ListApiDestinationsRequest) (_result *ListApiDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApiDestinationsResponse{}
	_body, _err := client.ListApiDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query connections.
 *
 * @param request ListConnectionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListConnectionsResponse
 */
func (client *Client) ListConnectionsWithOptions(request *ListConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionNamePrefix)) {
		body["ConnectionNamePrefix"] = request.ConnectionNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnections"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query connections.
 *
 * @param request ListConnectionsRequest
 * @return ListConnectionsResponse
 */
func (client *Client) ListConnections(request *ListConnectionsRequest) (_result *ListConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectionsResponse{}
	_body, _err := client.ListConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query all event buses.
 *
 * @param request ListEventBusesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListEventBusesResponse
 */
func (client *Client) ListEventBusesWithOptions(request *ListEventBusesRequest, runtime *util.RuntimeOptions) (_result *ListEventBusesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NamePrefix)) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventBuses"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventBusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query all event buses.
 *
 * @param request ListEventBusesRequest
 * @return ListEventBusesResponse
 */
func (client *Client) ListEventBuses(request *ListEventBusesRequest) (_result *ListEventBusesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventBusesResponse{}
	_body, _err := client.ListEventBusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query event streams.
 *
 * @param request ListEventStreamingsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListEventStreamingsResponse
 */
func (client *Client) ListEventStreamingsWithOptions(request *ListEventStreamingsRequest, runtime *util.RuntimeOptions) (_result *ListEventStreamingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NamePrefix)) {
		body["NamePrefix"] = request.NamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SinkArn)) {
		body["SinkArn"] = request.SinkArn
	}

	if !tea.BoolValue(util.IsUnset(request.SourceArn)) {
		body["SourceArn"] = request.SourceArn
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventStreamings"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventStreamingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query event streams.
 *
 * @param request ListEventStreamingsRequest
 * @return ListEventStreamingsResponse
 */
func (client *Client) ListEventStreamings(request *ListEventStreamingsRequest) (_result *ListEventStreamingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventStreamingsResponse{}
	_body, _err := client.ListEventStreamingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query all rules of an event bus.
 *
 * @param request ListRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRulesResponse
 */
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RuleNamePrefix)) {
		query["RuleNamePrefix"] = request.RuleNamePrefix
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRules"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query all rules of an event bus.
 *
 * @param request ListRulesRequest
 * @return ListRulesResponse
 */
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTargetsWithOptions(request *ListTargetsRequest, runtime *util.RuntimeOptions) (_result *ListTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Arn)) {
		query["Arn"] = request.Arn
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTargets"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTargets(request *ListTargetsRequest) (_result *ListTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTargetsResponse{}
	_body, _err := client.ListTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query custom event sources.
 *
 * @param request ListUserDefinedEventSourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListUserDefinedEventSourcesResponse
 */
func (client *Client) ListUserDefinedEventSourcesWithOptions(request *ListUserDefinedEventSourcesRequest, runtime *util.RuntimeOptions) (_result *ListUserDefinedEventSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NamePrefix)) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserDefinedEventSources"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserDefinedEventSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query custom event sources.
 *
 * @param request ListUserDefinedEventSourcesRequest
 * @return ListUserDefinedEventSourcesResponse
 */
func (client *Client) ListUserDefinedEventSources(request *ListUserDefinedEventSourcesRequest) (_result *ListUserDefinedEventSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDefinedEventSourcesResponse{}
	_body, _err := client.ListUserDefinedEventSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to stop an event stream that is running.
 *
 * @param request PauseEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PauseEventStreamingResponse
 */
func (client *Client) PauseEventStreamingWithOptions(request *PauseEventStreamingRequest, runtime *util.RuntimeOptions) (_result *PauseEventStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to stop an event stream that is running.
 *
 * @param request PauseEventStreamingRequest
 * @return PauseEventStreamingResponse
 */
func (client *Client) PauseEventStreaming(request *PauseEventStreamingRequest) (_result *PauseEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseEventStreamingResponse{}
	_body, _err := client.PauseEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to create or update event targets under a rule.
 *
 * @param tmpReq PutTargetsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutTargetsResponse
 */
func (client *Client) PutTargetsWithOptions(tmpReq *PutTargetsRequest, runtime *util.RuntimeOptions) (_result *PutTargetsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PutTargetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Targets)) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, tea.String("Targets"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetsShrink)) {
		query["Targets"] = request.TargetsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutTargets"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to create or update event targets under a rule.
 *
 * @param request PutTargetsRequest
 * @return PutTargetsResponse
 */
func (client *Client) PutTargets(request *PutTargetsRequest) (_result *PutTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutTargetsResponse{}
	_body, _err := client.PutTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query the content of an event.
 *
 * @param request QueryEventRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryEventResponse
 */
func (client *Client) QueryEventWithOptions(request *QueryEventRequest, runtime *util.RuntimeOptions) (_result *QueryEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.EventSource)) {
		query["EventSource"] = request.EventSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEvent"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query the content of an event.
 *
 * @param request QueryEventRequest
 * @return QueryEventResponse
 */
func (client *Client) QueryEvent(request *QueryEventRequest) (_result *QueryEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventResponse{}
	_body, _err := client.QueryEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query event traces.
 *
 * @param request QueryEventTracesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryEventTracesResponse
 */
func (client *Client) QueryEventTracesWithOptions(request *QueryEventTracesRequest, runtime *util.RuntimeOptions) (_result *QueryEventTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEventTraces"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query event traces.
 *
 * @param request QueryEventTracesRequest
 * @return QueryEventTracesResponse
 */
func (client *Client) QueryEventTraces(request *QueryEventTracesRequest) (_result *QueryEventTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEventTracesResponse{}
	_body, _err := client.QueryEventTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query event traces by event ID.
 *
 * @param request QueryTracedEventByEventIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryTracedEventByEventIdResponse
 */
func (client *Client) QueryTracedEventByEventIdWithOptions(request *QueryTracedEventByEventIdRequest, runtime *util.RuntimeOptions) (_result *QueryTracedEventByEventIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.EventSource)) {
		query["EventSource"] = request.EventSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTracedEventByEventId"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTracedEventByEventIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query event traces by event ID.
 *
 * @param request QueryTracedEventByEventIdRequest
 * @return QueryTracedEventByEventIdResponse
 */
func (client *Client) QueryTracedEventByEventId(request *QueryTracedEventByEventIdRequest) (_result *QueryTracedEventByEventIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTracedEventByEventIdResponse{}
	_body, _err := client.QueryTracedEventByEventIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to query event traces by time range.
 *
 * @param request QueryTracedEventsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryTracedEventsResponse
 */
func (client *Client) QueryTracedEventsWithOptions(request *QueryTracedEventsRequest, runtime *util.RuntimeOptions) (_result *QueryTracedEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventSource)) {
		query["EventSource"] = request.EventSource
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MatchedRule)) {
		query["MatchedRule"] = request.MatchedRule
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTracedEvents"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTracedEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to query event traces by time range.
 *
 * @param request QueryTracedEventsRequest
 * @return QueryTracedEventsResponse
 */
func (client *Client) QueryTracedEvents(request *QueryTracedEventsRequest) (_result *QueryTracedEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTracedEventsResponse{}
	_body, _err := client.QueryTracedEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to enable a created or deactivated event stream.
 *
 * @param request StartEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartEventStreamingResponse
 */
func (client *Client) StartEventStreamingWithOptions(request *StartEventStreamingRequest, runtime *util.RuntimeOptions) (_result *StartEventStreamingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to enable a created or deactivated event stream.
 *
 * @param request StartEventStreamingRequest
 * @return StartEventStreamingResponse
 */
func (client *Client) StartEventStreaming(request *StartEventStreamingRequest) (_result *StartEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartEventStreamingResponse{}
	_body, _err := client.StartEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to check whether the event pattern matches the provided JSON format.
 *
 * @param request TestEventPatternRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TestEventPatternResponse
 */
func (client *Client) TestEventPatternWithOptions(request *TestEventPatternRequest, runtime *util.RuntimeOptions) (_result *TestEventPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Event)) {
		body["Event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.EventPattern)) {
		body["EventPattern"] = request.EventPattern
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TestEventPattern"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestEventPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to check whether the event pattern matches the provided JSON format.
 *
 * @param request TestEventPatternRequest
 * @return TestEventPatternResponse
 */
func (client *Client) TestEventPattern(request *TestEventPatternRequest) (_result *TestEventPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestEventPatternResponse{}
	_body, _err := client.TestEventPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to update an API destination.
 *
 * @param tmpReq UpdateApiDestinationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateApiDestinationResponse
 */
func (client *Client) UpdateApiDestinationWithOptions(tmpReq *UpdateApiDestinationRequest, runtime *util.RuntimeOptions) (_result *UpdateApiDestinationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateApiDestinationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HttpApiParameters)) {
		request.HttpApiParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpApiParameters, tea.String("HttpApiParameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDestinationName)) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HttpApiParametersShrink)) {
		query["HttpApiParameters"] = request.HttpApiParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApiDestination"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApiDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to update an API destination.
 *
 * @param request UpdateApiDestinationRequest
 * @return UpdateApiDestinationResponse
 */
func (client *Client) UpdateApiDestination(request *UpdateApiDestinationRequest) (_result *UpdateApiDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApiDestinationResponse{}
	_body, _err := client.UpdateApiDestinationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to update a connection.
 *
 * @param tmpReq UpdateConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateConnectionResponse
 */
func (client *Client) UpdateConnectionWithOptions(tmpReq *UpdateConnectionRequest, runtime *util.RuntimeOptions) (_result *UpdateConnectionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AuthParameters)) {
		request.AuthParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthParameters, tea.String("AuthParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NetworkParameters)) {
		request.NetworkParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkParameters, tea.String("NetworkParameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthParametersShrink)) {
		query["AuthParameters"] = request.AuthParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionName)) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkParametersShrink)) {
		query["NetworkParameters"] = request.NetworkParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConnection"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to update a connection.
 *
 * @param request UpdateConnectionRequest
 * @return UpdateConnectionResponse
 */
func (client *Client) UpdateConnection(request *UpdateConnectionRequest) (_result *UpdateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.UpdateConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to modify an event bus.
 *
 * @param request UpdateEventBusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateEventBusResponse
 */
func (client *Client) UpdateEventBusWithOptions(request *UpdateEventBusRequest, runtime *util.RuntimeOptions) (_result *UpdateEventBusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventBus"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEventBusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to modify an event bus.
 *
 * @param request UpdateEventBusRequest
 * @return UpdateEventBusResponse
 */
func (client *Client) UpdateEventBus(request *UpdateEventBusRequest) (_result *UpdateEventBusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEventBusResponse{}
	_body, _err := client.UpdateEventBusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to update an event source.
 *
 * @param tmpReq UpdateEventSourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateEventSourceResponse
 */
func (client *Client) UpdateEventSourceWithOptions(tmpReq *UpdateEventSourceRequest, runtime *util.RuntimeOptions) (_result *UpdateEventSourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEventSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SourceHttpEventParameters)) {
		request.SourceHttpEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceHttpEventParameters, tea.String("SourceHttpEventParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceKafkaParameters)) {
		request.SourceKafkaParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceKafkaParameters, tea.String("SourceKafkaParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceMNSParameters)) {
		request.SourceMNSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMNSParameters, tea.String("SourceMNSParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceRabbitMQParameters)) {
		request.SourceRabbitMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRabbitMQParameters, tea.String("SourceRabbitMQParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceRocketMQParameters)) {
		request.SourceRocketMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRocketMQParameters, tea.String("SourceRocketMQParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceSLSParameters)) {
		request.SourceSLSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSLSParameters, tea.String("SourceSLSParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceScheduledEventParameters)) {
		request.SourceScheduledEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceScheduledEventParameters, tea.String("SourceScheduledEventParameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		body["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventSourceName)) {
		body["EventSourceName"] = request.EventSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceHttpEventParametersShrink)) {
		body["SourceHttpEventParameters"] = request.SourceHttpEventParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceKafkaParametersShrink)) {
		body["SourceKafkaParameters"] = request.SourceKafkaParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceMNSParametersShrink)) {
		body["SourceMNSParameters"] = request.SourceMNSParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRabbitMQParametersShrink)) {
		body["SourceRabbitMQParameters"] = request.SourceRabbitMQParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRocketMQParametersShrink)) {
		body["SourceRocketMQParameters"] = request.SourceRocketMQParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceSLSParametersShrink)) {
		body["SourceSLSParameters"] = request.SourceSLSParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceScheduledEventParametersShrink)) {
		body["SourceScheduledEventParameters"] = request.SourceScheduledEventParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventSource"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEventSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to update an event source.
 *
 * @param request UpdateEventSourceRequest
 * @return UpdateEventSourceResponse
 */
func (client *Client) UpdateEventSource(request *UpdateEventSourceRequest) (_result *UpdateEventSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEventSourceResponse{}
	_body, _err := client.UpdateEventSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to modify the information about an event stream, such as the basic information and the information about the event source, event filtering rule, and event target.
 *
 * @param tmpReq UpdateEventStreamingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateEventStreamingResponse
 */
func (client *Client) UpdateEventStreamingWithOptions(tmpReq *UpdateEventStreamingRequest, runtime *util.RuntimeOptions) (_result *UpdateEventStreamingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEventStreamingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RunOptions)) {
		request.RunOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RunOptions, tea.String("RunOptions"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sink)) {
		request.SinkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sink, tea.String("Sink"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Source)) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, tea.String("Source"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Transforms)) {
		request.TransformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transforms, tea.String("Transforms"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventStreamingName)) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterPattern)) {
		body["FilterPattern"] = request.FilterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.RunOptionsShrink)) {
		body["RunOptions"] = request.RunOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SinkShrink)) {
		body["Sink"] = request.SinkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceShrink)) {
		body["Source"] = request.SourceShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TransformsShrink)) {
		body["Transforms"] = request.TransformsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventStreaming"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEventStreamingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to modify the information about an event stream, such as the basic information and the information about the event source, event filtering rule, and event target.
 *
 * @param request UpdateEventStreamingRequest
 * @return UpdateEventStreamingResponse
 */
func (client *Client) UpdateEventStreaming(request *UpdateEventStreamingRequest) (_result *UpdateEventStreamingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEventStreamingResponse{}
	_body, _err := client.UpdateEventStreamingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this API operation to update the configurations of an event rule.
 *
 * @param request UpdateRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateRuleResponse
 */
func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		query["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterPattern)) {
		query["FilterPattern"] = request.FilterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRule"),
		Version:     tea.String("2020-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this API operation to update the configurations of an event rule.
 *
 * @param request UpdateRuleRequest
 * @return UpdateRuleResponse
 */
func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
