// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConnectionsResponseBody
	GetCode() *string
	SetData(v *ListConnectionsResponseBodyData) *ListConnectionsResponseBody
	GetData() *ListConnectionsResponseBodyData
	SetMessage(v string) *ListConnectionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConnectionsResponseBody
	GetRequestId() *string
}

type ListConnectionsResponseBody struct {
	// The HTTP status code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListConnectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// E3619976-8714-5D88-BBA2-6983D798A8BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConnectionsResponseBody) GetData() *ListConnectionsResponseBodyData {
	return s.Data
}

func (s *ListConnectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyData struct {
	// The connections.
	Connections []*ListConnectionsResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *float32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *float32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListConnectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyData) GetConnections() []*ListConnectionsResponseBodyDataConnections {
	return s.Connections
}

func (s *ListConnectionsResponseBodyData) GetMaxResults() *float32 {
	return s.MaxResults
}

func (s *ListConnectionsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectionsResponseBodyData) GetTotal() *float32 {
	return s.Total
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

func (s *ListConnectionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnections struct {
	// The parameters that are returned for authentication.
	AuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The connection name.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The connection description.
	//
	// example:
	//
	// The description of the connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the connection was created.
	//
	// example:
	//
	// 1592838994234
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the connection.
	//
	// example:
	//
	// 1141093
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameters that are returned for the network.
	NetworkParameters *ListConnectionsResponseBodyDataConnectionsNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnections) GetAuthParameters() *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	return s.AuthParameters
}

func (s *ListConnectionsResponseBodyDataConnections) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListConnectionsResponseBodyDataConnections) GetDescription() *string {
	return s.Description
}

func (s *ListConnectionsResponseBodyDataConnections) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListConnectionsResponseBodyDataConnections) GetId() *int64 {
	return s.Id
}

func (s *ListConnectionsResponseBodyDataConnections) GetNetworkParameters() *ListConnectionsResponseBodyDataConnectionsNetworkParameters {
	return s.NetworkParameters
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

func (s *ListConnectionsResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParameters struct {
	// The parameters that are returned for API key authentication.
	ApiKeyAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	// The authentication method. Valid values:
	//
	// 	- BASIC_AUTH: basic authentication.
	//
	// 	- API_KEY_AUTH: API key authentication.
	//
	// 	- OAUTH_AUTH: OAuth authentication.
	//
	// example:
	//
	// BASIC_AUTH
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The parameters that are returned for basic authentication.
	BasicAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The parameters that are returned for OAuth authentication.
	OAuthParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) GetApiKeyAuthParameters() *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	return s.ApiKeyAuthParameters
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) GetBasicAuthParameters() *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	return s.BasicAuthParameters
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) GetOAuthParameters() *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters {
	return s.OAuthParameters
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
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
	// asdkjnqkwejooa
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters struct {
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

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GetPassword() *string {
	return s.Password
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GetUsername() *string {
	return s.Username
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetPassword(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetUsername(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The parameters that are returned for the client.
	ClientParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request parameters of OAuth authentication.
	OAuthHttpParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GetClientParameters() *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	return s.ClientParameters
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) GetOAuthHttpParameters() *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	return s.OAuthHttpParameters
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters struct {
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

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GetClientID() *string {
	return s.ClientID
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientID(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters struct {
	// The parameters that are configured for the request.
	BodyParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	// The parameters that are returned for the request header.
	HeaderParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	// The parameters that are returned for the request path.
	QueryStringParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetBodyParameters() []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	return s.BodyParameters
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetHeaderParameters() []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetQueryStringParameters() []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	return s.QueryStringParameters
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
	// Indicates whether authentication is enabled.
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

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetKey() *string {
	return s.Key
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetValue() *string {
	return s.Value
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
	// Indicates whether authentication is enabled.
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

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetKey() *string {
	return s.Key
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetValue() *string {
	return s.Value
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
	// Indicates whether authentication is enabled.
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

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetKey() *string {
	return s.Key
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetValue() *string {
	return s.Value
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) Validate() error {
	return dara.Validate(s)
}

type ListConnectionsResponseBodyDataConnectionsNetworkParameters struct {
	// 	- PublicNetwork: the Internet.
	//
	// 	- PrivateNetwork: virtual private cloud (VPC).
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

func (s ListConnectionsResponseBodyDataConnectionsNetworkParameters) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsNetworkParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) GetVswitcheId() *string {
	return s.VswitcheId
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

func (s *ListConnectionsResponseBodyDataConnectionsNetworkParameters) Validate() error {
	return dara.Validate(s)
}
