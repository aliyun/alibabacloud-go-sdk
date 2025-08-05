// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConnectionResponseBody
	GetCode() *string
	SetData(v *GetConnectionResponseBodyData) *GetConnectionResponseBody
	GetData() *GetConnectionResponseBodyData
	SetHttpCode(v int32) *GetConnectionResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *GetConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConnectionResponseBody
	GetRequestId() *string
}

type GetConnectionResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConnectionResponseBody) GetData() *GetConnectionResponseBodyData {
	return s.Data
}

func (s *GetConnectionResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *GetConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyData struct {
	// The queried connections.
	Connections []*GetConnectionResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
}

func (s GetConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyData) GetConnections() []*GetConnectionResponseBodyDataConnections {
	return s.Connections
}

func (s *GetConnectionResponseBodyData) SetConnections(v []*GetConnectionResponseBodyDataConnections) *GetConnectionResponseBodyData {
	s.Connections = v
	return s
}

func (s *GetConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnections struct {
	// The authentication methods.
	AuthParameters *GetConnectionResponseBodyDataConnectionsAuthParameters `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	// The connection name.
	//
	// example:
	//
	// demo
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The connection description.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the connection was created.
	//
	// example:
	//
	// 1669648905
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 5668
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the network.
	NetworkParameters *GetConnectionResponseBodyDataConnectionsNetworkParameters `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnections) GetAuthParameters() *GetConnectionResponseBodyDataConnectionsAuthParameters {
	return s.AuthParameters
}

func (s *GetConnectionResponseBodyDataConnections) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetConnectionResponseBodyDataConnections) GetDescription() *string {
	return s.Description
}

func (s *GetConnectionResponseBodyDataConnections) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetConnectionResponseBodyDataConnections) GetId() *int64 {
	return s.Id
}

func (s *GetConnectionResponseBodyDataConnections) GetNetworkParameters() *GetConnectionResponseBodyDataConnectionsNetworkParameters {
	return s.NetworkParameters
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

func (s *GetConnectionResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParameters struct {
	// The information about API key authentication.
	ApiKeyAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
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
	// The information about basic authentication.
	BasicAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	// The information about OAuth authentication.
	OAuthParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) GetApiKeyAuthParameters() *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	return s.ApiKeyAuthParameters
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) GetBasicAuthParameters() *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	return s.BasicAuthParameters
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) GetOAuthParameters() *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters {
	return s.OAuthParameters
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
	// The key of the API key.
	//
	// example:
	//
	// key
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The value of the API key.
	//
	// example:
	//
	// value
	ApiKeyValue *string `json:"ApiKeyValue,omitempty" xml:"ApiKeyValue,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyName(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyName = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) SetApiKeyValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters {
	s.ApiKeyValue = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters struct {
	// The password of basic authentication.
	//
	// example:
	//
	// ********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username of basic authentication.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GetPassword() *string {
	return s.Password
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) GetUsername() *string {
	return s.Username
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetPassword(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Password = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) SetUsername(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters {
	s.Username = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	// The endpoint that is used to obtain the OAuth token.
	//
	// example:
	//
	// http://localhost:8080/oauth/token
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The information about the client.
	ClientParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
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
	OAuthHttpParameters *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GetClientParameters() *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	return s.ClientParameters
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) GetOAuthHttpParameters() *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters {
	return s.OAuthHttpParameters
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters struct {
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
	// Qo57Q~F249~S74GmNPA36pZJoJK4f****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GetClientID() *string {
	return s.ClientID
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientID(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientID = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) SetClientSecret(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters {
	s.ClientSecret = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetBodyParameters() []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters {
	return s.BodyParameters
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetHeaderParameters() []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters {
	return s.HeaderParameters
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) GetQueryStringParameters() []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters {
	return s.QueryStringParameters
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters struct {
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
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request body.
	//
	// example:
	//
	// valueDemo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetKey() *string {
	return s.Key
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) GetValue() *string {
	return s.Value
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters struct {
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
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// keyDemo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetKey() *string {
	return s.Key
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) GetValue() *string {
	return s.Value
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters struct {
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
	// keyDemo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request path.
	//
	// example:
	//
	// valueDemo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetIsValueSecret() *string {
	return s.IsValueSecret
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetKey() *string {
	return s.Key
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) GetValue() *string {
	return s.Value
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersQueryStringParameters) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyDataConnectionsNetworkParameters struct {
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
	// vsw-bp1iu4x7aeradadown****,vsw-bp193sqmadadlaszpeq****
	VswitcheId *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsNetworkParameters) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsNetworkParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) GetVswitcheId() *string {
	return s.VswitcheId
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

func (s *GetConnectionResponseBodyDataConnectionsNetworkParameters) Validate() error {
	return dara.Validate(s)
}
