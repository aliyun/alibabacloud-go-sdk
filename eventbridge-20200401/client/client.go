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
	ApiDestinationName *string                                       `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName     *string                                       `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description        *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpApiParameters  *CreateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
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
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
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
	ApiDestinationName      *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName          *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Date      *CreateApiDestinationResponseBodyDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AuthParameters    *CreateConnectionRequestAuthParameters    `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	ConnectionName    *string                                   `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description       *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
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
	ApiKeyAuthParameters *CreateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	AuthorizationType    *string                                                    `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	BasicAuthParameters  *CreateConnectionRequestAuthParametersBasicAuthParameters  `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	OAuthParameters      *CreateConnectionRequestAuthParametersOAuthParameters      `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
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
	ApiKeyName  *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	AuthorizationEndpoint *string                                                                  `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	ClientParameters      *CreateConnectionRequestAuthParametersOAuthParametersClientParameters    `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	HttpMethod            *string                                                                  `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	OAuthHttpParameters   *CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
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
	ClientID     *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
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
	BodyParameters        []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*CreateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitcheId      *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
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
	AuthParametersShrink    *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	ConnectionName          *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 事件源描述详情
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// 事件源英文Code
	EventSourceName           *string                                            `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	SourceHttpEventParameters *CreateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	SourceKafkaParameters     *CreateEventSourceRequestSourceKafkaParameters     `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	SourceMNSParameters       *CreateEventSourceRequestSourceMNSParameters       `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceRabbitMQParameters  *CreateEventSourceRequestSourceRabbitMQParameters  `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQParameters  *CreateEventSourceRequestSourceRocketMQParameters  `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// SourceSLSParameters
	SourceSLSParameters            *CreateEventSourceRequestSourceSLSParameters            `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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
	Ip             []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Method         []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	Referer        []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	SecurityConfig *string   `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaximumTasks    *int32  `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	IsBase64Decode *string `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEventSourceRequestSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v string) *CreateEventSourceRequestSourceMNSParameters {
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
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupID                 *string  `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceId              *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceNetwork         *string  `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	InstanceSecurityGroupId *string  `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	InstanceType            *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceVSwitchIds      *string  `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	InstanceVpcId           *string  `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	Offset                  *string  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId                *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag                     *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp               *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic                   *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventSourceRequestSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
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

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
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

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v float32) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type CreateEventSourceRequestSourceSLSParameters struct {
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
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

type CreateEventSourceShrinkRequest struct {
	// 事件源描述详情
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// 事件源英文Code
	EventSourceName                 *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	SourceKafkaParametersShrink     *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	SourceMNSParametersShrink       *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	SourceRabbitMQParametersShrink  *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	SourceRocketMQParametersShrink  *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// SourceSLSParameters
	SourceSLSParametersShrink            *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
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
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateEventSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description        *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string                                `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string                                `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptions         *CreateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	Sink               *CreateEventStreamingRequestSink       `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	Source             *CreateEventStreamingRequestSource     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
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

type CreateEventStreamingRequestRunOptions struct {
	BatchWindow     *CreateEventStreamingRequestRunOptionsBatchWindow     `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	DeadLetterQueue *CreateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	ErrorsTolerance *string                                               `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	MaximumTasks    *int64                                                `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	RetryStrategy   *CreateEventStreamingRequestRunOptionsRetryStrategy   `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
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
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	TimeBasedWindow  *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
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
	MaximumEventAgeInSeconds *int64  `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	MaximumRetryAttempts     *int64  `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	PushRetryStrategy        *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
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
	SinkDataHubParameters  *CreateEventStreamingRequestSinkSinkDataHubParameters  `json:"SinkDataHubParameters,omitempty" xml:"SinkDataHubParameters,omitempty" type:"Struct"`
	SinkFcParameters       *CreateEventStreamingRequestSinkSinkFcParameters       `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	SinkKafkaParameters    *CreateEventStreamingRequestSinkSinkKafkaParameters    `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	SinkMNSParameters      *CreateEventStreamingRequestSinkSinkMNSParameters      `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
	SinkRabbitMQParameters *CreateEventStreamingRequestSinkSinkRabbitMQParameters `json:"SinkRabbitMQParameters,omitempty" xml:"SinkRabbitMQParameters,omitempty" type:"Struct"`
	// Sink RocketMQ Parameters
	SinkRocketMQParameters *CreateEventStreamingRequestSinkSinkRocketMQParameters `json:"SinkRocketMQParameters,omitempty" xml:"SinkRocketMQParameters,omitempty" type:"Struct"`
	// Sink SLS Parameters
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

func (s *CreateEventStreamingRequestSink) SetSinkKafkaParameters(v *CreateEventStreamingRequestSinkSinkKafkaParameters) *CreateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *CreateEventStreamingRequestSink) SetSinkMNSParameters(v *CreateEventStreamingRequestSinkSinkMNSParameters) *CreateEventStreamingRequestSink {
	s.SinkMNSParameters = v
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
	Body        *CreateEventStreamingRequestSinkSinkDataHubParametersBody        `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Project     *CreateEventStreamingRequestSinkSinkDataHubParametersProject     `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName    *CreateEventStreamingRequestSinkSinkDataHubParametersRoleName    `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic       *CreateEventStreamingRequestSinkSinkDataHubParametersTopic       `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	TopicSchema *CreateEventStreamingRequestSinkSinkDataHubParametersTopicSchema `json:"TopicSchema,omitempty" xml:"TopicSchema,omitempty" type:"Struct"`
	TopicType   *CreateEventStreamingRequestSinkSinkDataHubParametersTopicType   `json:"TopicType,omitempty" xml:"TopicType,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body           *CreateEventStreamingRequestSinkSinkFcParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Concurrency    *CreateEventStreamingRequestSinkSinkFcParametersConcurrency    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	FunctionName   *CreateEventStreamingRequestSinkSinkFcParametersFunctionName   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	InvocationType *CreateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	Qualifier      *CreateEventStreamingRequestSinkSinkFcParametersQualifier      `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	ServiceName    *CreateEventStreamingRequestSinkSinkFcParametersServiceName    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type CreateEventStreamingRequestSinkSinkKafkaParameters struct {
	Acks       *CreateEventStreamingRequestSinkSinkKafkaParametersAcks       `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	InstanceId *CreateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Key        *CreateEventStreamingRequestSinkSinkKafkaParametersKey        `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	Topic      *CreateEventStreamingRequestSinkSinkKafkaParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	Value      *CreateEventStreamingRequestSinkSinkKafkaParametersValue      `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body           *CreateEventStreamingRequestSinkSinkMNSParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	IsBase64Encode *CreateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	QueueName      *CreateEventStreamingRequestSinkSinkMNSParametersQueueName      `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type CreateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	Body            *CreateEventStreamingRequestSinkSinkRabbitMQParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Exchange        *CreateEventStreamingRequestSinkSinkRabbitMQParametersExchange        `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	InstanceId      *CreateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	MessageId       *CreateEventStreamingRequestSinkSinkRabbitMQParametersMessageId       `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	Properties      *CreateEventStreamingRequestSinkSinkRabbitMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	QueueName       *CreateEventStreamingRequestSinkSinkRabbitMQParametersQueueName       `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	RoutingKey      *CreateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey      `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	TargetType      *CreateEventStreamingRequestSinkSinkRabbitMQParametersTargetType      `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body             *CreateEventStreamingRequestSinkSinkRocketMQParametersBody             `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	InstanceEndpoint *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceEndpoint `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty" type:"Struct"`
	InstanceId       *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceId       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	InstancePassword *CreateEventStreamingRequestSinkSinkRocketMQParametersInstancePassword `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty" type:"Struct"`
	InstanceType     *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceType     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Struct"`
	InstanceUsername *CreateEventStreamingRequestSinkSinkRocketMQParametersInstanceUsername `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty" type:"Struct"`
	Keys             *CreateEventStreamingRequestSinkSinkRocketMQParametersKeys             `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Network          *CreateEventStreamingRequestSinkSinkRocketMQParametersNetwork          `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Properties       *CreateEventStreamingRequestSinkSinkRocketMQParametersProperties       `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	SecurityGroupId  *CreateEventStreamingRequestSinkSinkRocketMQParametersSecurityGroupId  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Struct"`
	Tags             *CreateEventStreamingRequestSinkSinkRocketMQParametersTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic            *CreateEventStreamingRequestSinkSinkRocketMQParametersTopic            `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	VSwitchIds       *CreateEventStreamingRequestSinkSinkRocketMQParametersVSwitchIds       `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId            *CreateEventStreamingRequestSinkSinkRocketMQParametersVpcId            `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body     *CreateEventStreamingRequestSinkSinkSLSParametersBody     `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	LogStore *CreateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	Project  *CreateEventStreamingRequestSinkSinkSLSParametersProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName *CreateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic    *CreateEventStreamingRequestSinkSinkSLSParametersTopic    `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	SourceDTSParameters      *CreateEventStreamingRequestSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceKafkaParameters    *CreateEventStreamingRequestSourceSourceKafkaParameters    `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	SourceMNSParameters      *CreateEventStreamingRequestSourceSourceMNSParameters      `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceMQTTParameters     *CreateEventStreamingRequestSourceSourceMQTTParameters     `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceRabbitMQParameters *CreateEventStreamingRequestSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQParameters *CreateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	SourceSLSParameters      *CreateEventStreamingRequestSourceSourceSLSParameters      `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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
	BrokerUrl      *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	InitCheckPoint *int64  `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Sid            *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// VPC ID。
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

func (s *CreateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *CreateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type CreateEventStreamingRequestSourceSourceMNSParameters struct {
	IsBase64Decode *bool   `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateEventStreamingRequestSourceSourceMQTTParameters) GoString() string {
	return s.String()
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

type CreateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AuthType                *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	FilterSql               *string `json:"FilterSql,omitempty" xml:"FilterSql,omitempty"`
	FilterType              *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	GroupID                 *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceEndpoint        *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceNetwork         *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	InstancePassword        *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceUsername        *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	InstanceVSwitchIds      *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	InstanceVpcId           *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	Network                 *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Offset                  *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId         *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag                     *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp               *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic                   *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds              *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

type CreateEventStreamingShrinkRequest struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptionsShrink   *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	SinkShrink         *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	SourceShrink       *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

type CreateEventStreamingResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description   *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName  *string                          `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventTargets  []*CreateRuleRequestEventTargets `json:"EventTargets,omitempty" xml:"EventTargets,omitempty" type:"Repeated"`
	FilterPattern *string                          `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RuleName      *string                          `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status        *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
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
	DeadLetterQueue   *CreateRuleRequestEventTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	Endpoint          *string                                       `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ErrorsTolerance   *string                                       `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	Id                *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	ParamList         []*CreateRuleRequestEventTargetsParamList     `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	PushRetryStrategy *string                                       `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	Type              *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Form        *string `json:"Form,omitempty" xml:"Form,omitempty"`
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	Template    *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName       *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventTargetsShrink *string `json:"EventTargets,omitempty" xml:"EventTargets,omitempty"`
	FilterPattern      *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleForProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 事件源ID
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *bool   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string   `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	TargetIds    []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
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
	EventBusName    *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ErrorEntries      []*DeleteTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	ErrorEntriesCount *int32                                       `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
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
	EntryId      *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetApiDestinationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApiDestinationName *string                                             `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName     *string                                             `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description        *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *int64                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	HttpApiParameters  *GetApiDestinationResponseBodyDataHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
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
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetConnectionResponseBody) SetMessage(v string) *GetConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectionResponseBody) SetRequestId(v string) *GetConnectionResponseBody {
	s.RequestId = &v
	return s
}

type GetConnectionResponseBodyData struct {
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
	AuthParameters    *GetConnectionResponseBodyDataConnectionsAuthParameters    `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	ConnectionName    *string                                                    `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description       *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate         *int64                                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                *int64                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
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
	ApiKeyAuthParameters     *GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters     `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	AuthorizationType        *string                                                                         `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	BasicAuthParameters      *GetConnectionResponseBodyDataConnectionsAuthParametersBasicAuthParameters      `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	InvocationHttpParameters *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters `json:"InvocationHttpParameters,omitempty" xml:"InvocationHttpParameters,omitempty" type:"Struct"`
	OAuthParameters          *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters          `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
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

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetInvocationHttpParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.InvocationHttpParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParameters) SetOAuthParameters(v *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters) *GetConnectionResponseBodyDataConnectionsAuthParameters {
	s.OAuthParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
	ApiKeyName  *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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

type GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters struct {
	BodyParameters        []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	QueryStringParameters []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetBodyParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetHeaderParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetQueryStringParameters(v []*GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.QueryStringParameters = v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters struct {
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ValuValuee *string `json:"ValuValuee,omitempty" xml:"ValuValuee,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) SetValuValuee(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters {
	s.ValuValuee = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) SetValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) SetKey(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) SetValue(v string) *GetConnectionResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

type GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	AuthorizationEndpoint *string                                                                                   `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	ClientParameters      *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters    `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	HttpMethod            *string                                                                                   `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	OAuthHttpParameters   *GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
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
	ClientID     *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
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
	BodyParameters        []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*GetConnectionResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitcheId      *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusARN     *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	EventBusName    *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Description        *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string                                      `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string                                      `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptions         *GetEventStreamingResponseBodyDataRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	Sink               *GetEventStreamingResponseBodyDataSink       `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	Source             *GetEventStreamingResponseBodyDataSource     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status             *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
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

type GetEventStreamingResponseBodyDataRunOptions struct {
	BatchWindow     *GetEventStreamingResponseBodyDataRunOptionsBatchWindow     `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	DeadLetterQueue *GetEventStreamingResponseBodyDataRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	ErrorsTolerance *string                                                     `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	MaximumTasks    *int32                                                      `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	RetryStrategy   *GetEventStreamingResponseBodyDataRunOptionsRetryStrategy   `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
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
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	TimeBasedWindow  *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
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
	MaximumEventAgeInSeconds *float32 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	MaximumRetryAttempts     *float32 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	PushRetryStrategy        *string  `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
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
	SinkFcParameters       *GetEventStreamingResponseBodyDataSinkSinkFcParameters       `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	SinkKafkaParameters    *GetEventStreamingResponseBodyDataSinkSinkKafkaParameters    `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	SinkMNSParameters      *GetEventStreamingResponseBodyDataSinkSinkMNSParameters      `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
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
	Body           *GetEventStreamingResponseBodyDataSinkSinkFcParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Concurrency    *GetEventStreamingResponseBodyDataSinkSinkFcParametersConcurrency    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	FunctionName   *GetEventStreamingResponseBodyDataSinkSinkFcParametersFunctionName   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	InvocationType *GetEventStreamingResponseBodyDataSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	Qualifier      *GetEventStreamingResponseBodyDataSinkSinkFcParametersQualifier      `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	ServiceName    *GetEventStreamingResponseBodyDataSinkSinkFcParametersServiceName    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type GetEventStreamingResponseBodyDataSinkSinkKafkaParameters struct {
	Acks       *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersAcks       `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Key        *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersKey        `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	Topic      *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	Value      *GetEventStreamingResponseBodyDataSinkSinkKafkaParametersValue      `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body           *GetEventStreamingResponseBodyDataSinkSinkMNSParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	IsBase64Encode *GetEventStreamingResponseBodyDataSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	QueueName      *GetEventStreamingResponseBodyDataSinkSinkMNSParametersQueueName      `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body            *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Exchange        *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersExchange        `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	InstanceId      *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersInstanceId      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	MessageId       *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersMessageId       `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	Properties      *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	QueueName       *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersQueueName       `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	RoutingKey      *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersRoutingKey      `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	TargetType      *GetEventStreamingResponseBodyDataSinkSinkRabbitMQParametersTargetType      `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body       *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersBody       `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	InstanceId *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Keys       *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersKeys       `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Properties *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Tags       *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic      *GetEventStreamingResponseBodyDataSinkSinkRocketMQParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body     *GetEventStreamingResponseBodyDataSinkSinkSLSParametersBody     `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	LogStore *GetEventStreamingResponseBodyDataSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	Project  *GetEventStreamingResponseBodyDataSinkSinkSLSParametersProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName *GetEventStreamingResponseBodyDataSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic    *GetEventStreamingResponseBodyDataSinkSinkSLSParametersTopic    `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	SourceDTSParameters *GetEventStreamingResponseBodyDataSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// Source Kafka Parameters
	SourceKafkaParameters *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// Source MNS Parameters
	SourceMNSParameters  *GetEventStreamingResponseBodyDataSourceSourceMNSParameters  `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceMQTTParameters *GetEventStreamingResponseBodyDataSourceSourceMQTTParameters `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	// Source RabbitMQ Parameters
	SourceRabbitMQParameters *GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// Source RocketMQ Parameters
	SourceRocketMQParameters *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	SourceSLSParameters      *GetEventStreamingResponseBodyDataSourceSourceSLSParameters      `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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
	BrokerUrl      *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Sid            *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// VPC ID。
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

func (s *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters) SetVpcId(v string) *GetEventStreamingResponseBodyDataSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type GetEventStreamingResponseBodyDataSourceSourceMNSParameters struct {
	IsBase64Decode *bool   `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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

type GetEventStreamingResponseBodyDataSourceSourceRabbitMQParameters struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupID    *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Offset     *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp  *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetGroupID(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters) SetInstanceId(v string) *GetEventStreamingResponseBodyDataSourceSourceRocketMQParameters {
	s.InstanceId = &v
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
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

type GetEventStreamingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
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
	CreatedTimestamp *int64                            `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	Description      *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName     *string                           `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	FilterPattern    *string                           `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RuleARN          *string                           `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	RuleName         *string                           `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status           *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Targets          []*GetRuleResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	DeadLetterQueue   *GetRuleResponseBodyDataTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	DetailMap         map[string]interface{}                         `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	Endpoint          *string                                        `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Id                *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	ParamList         []*GetRuleResponseBodyDataTargetsParamList     `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	PushRetryStrategy *string                                        `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	PushSelector      *string                                        `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	Type              *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// TEMPLATE
	Form        *string `json:"Form,omitempty" xml:"Form,omitempty"`
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	Template    *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListAliyunOfficialEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Arn          *string                                                                    `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Ctime        *float32                                                                   `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Description  *string                                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName *string                                                                    `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventTypes   []*ListAliyunOfficialEventSourcesResponseBodyDataEventSourceListEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	FullName     *string                                                                    `json:"FullName,omitempty" xml:"FullName,omitempty"`
	Name         *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
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
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortName       *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAliyunOfficialEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiDestinationNamePrefix *string `json:"ApiDestinationNamePrefix,omitempty" xml:"ApiDestinationNamePrefix,omitempty"`
	ConnectionName           *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	MaxResults               *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListApiDestinationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApiDestinations []*ListApiDestinationsResponseBodyDataApiDestinations `json:"ApiDestinations,omitempty" xml:"ApiDestinations,omitempty" type:"Repeated"`
	MaxResults      *float32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total           *float32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
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
	ApiDestinationName *string                                                              `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName     *string                                                              `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description        *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *int64                                                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	HttpApiParameters  *ListApiDestinationsResponseBodyDataApiDestinationsHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
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
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApiDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConnectionNamePrefix *string `json:"ConnectionNamePrefix,omitempty" xml:"ConnectionNamePrefix,omitempty"`
	MaxResults           *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListConnectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Connections []*ListConnectionsResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	MaxResults  *float32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total       *float32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
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
	AuthParameters    *ListConnectionsResponseBodyDataConnectionsAuthParameters    `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	ConnectionName    *string                                                      `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description       *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate         *int64                                                       `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
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
	ApiKeyAuthParameters     *ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters     `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	AuthorizationType        *string                                                                           `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	BasicAuthParameters      *ListConnectionsResponseBodyDataConnectionsAuthParametersBasicAuthParameters      `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	InvocationHttpParameters *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters `json:"InvocationHttpParameters,omitempty" xml:"InvocationHttpParameters,omitempty" type:"Struct"`
	OAuthParameters          *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters          `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
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

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetInvocationHttpParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.InvocationHttpParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParameters) SetOAuthParameters(v *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters) *ListConnectionsResponseBodyDataConnectionsAuthParameters {
	s.OAuthParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersApiKeyAuthParameters struct {
	ApiKeyName  *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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

type ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters struct {
	BodyParameters        []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
	QueryStringParameters []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters `json:"QueryStringParameters,omitempty" xml:"QueryStringParameters,omitempty" type:"Repeated"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetBodyParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.BodyParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetHeaderParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.HeaderParameters = v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters) SetQueryStringParameters(v []*ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParameters {
	s.QueryStringParameters = v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersBodyParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersHeaderParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) SetKey(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters {
	s.Key = &v
	return s
}

func (s *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters) SetValue(v string) *ListConnectionsResponseBodyDataConnectionsAuthParametersInvocationHttpParametersQueryStringParameters {
	s.Value = &v
	return s
}

type ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParameters struct {
	AuthorizationEndpoint *string                                                                                     `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	ClientParameters      *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersClientParameters    `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	HttpMethod            *string                                                                                     `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	OAuthHttpParameters   *ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
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
	ClientID     *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
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
	BodyParameters        []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*ListConnectionsResponseBodyDataConnectionsAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitcheId      *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Limit      *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListEventBusesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
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
	EventBuses []*ListEventBusesResponseBodyDataEventBuses `json:"EventBuses,omitempty" xml:"EventBuses,omitempty" type:"Repeated"`
	NextToken  *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total      *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
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
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusARN     *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	EventBusName    *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEventBusesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Limit      *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SinkArn    *string `json:"SinkArn,omitempty" xml:"SinkArn,omitempty"`
	SourceArn  *string `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
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
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListEventStreamingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
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
	EventStreamings []*ListEventStreamingsResponseBodyDataEventStreamings `json:"EventStreamings,omitempty" xml:"EventStreamings,omitempty" type:"Repeated"`
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total           *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
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
	Description        *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string                                                       `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string                                                       `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptions         *ListEventStreamingsResponseBodyDataEventStreamingsRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	Sink               *ListEventStreamingsResponseBodyDataEventStreamingsSink       `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	Source             *ListEventStreamingsResponseBodyDataEventStreamingsSource     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Status             *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
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

type ListEventStreamingsResponseBodyDataEventStreamingsRunOptions struct {
	BatchWindow     *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsBatchWindow     `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	DeadLetterQueue *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	ErrorsTolerance *string                                                                      `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	MaximumTasks    *int32                                                                       `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	RetryStrategy   *ListEventStreamingsResponseBodyDataEventStreamingsRunOptionsRetryStrategy   `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
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
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	TimeBasedWindow  *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
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
	MaximumEventAgeInSeconds *float32 `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	MaximumRetryAttempts     *float32 `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	PushRetryStrategy        *string  `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
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
	SinkFcParameters       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParameters       `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	SinkKafkaParameters    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters    `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	SinkMNSParameters      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParameters      `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
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
	Body           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Concurrency    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersConcurrency    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	FunctionName   *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersFunctionName   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	InvocationType *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	Qualifier      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersQualifier      `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	ServiceName    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkFcParametersServiceName    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParameters struct {
	Acks       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersAcks       `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Key        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersKey        `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	Topic      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	Value      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkKafkaParametersValue      `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body           *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	IsBase64Encode *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	QueueName      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkMNSParametersQueueName      `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body            *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Exchange        *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersExchange        `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	InstanceId      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersInstanceId      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	MessageId       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersMessageId       `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	Properties      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	QueueName       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersQueueName       `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	RoutingKey      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersRoutingKey      `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	TargetType      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRabbitMQParametersTargetType      `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersBody       `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	InstanceId *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Keys       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersKeys       `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Properties *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Tags       *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic      *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkRocketMQParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body     *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersBody     `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	LogStore *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	Project  *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic    *ListEventStreamingsResponseBodyDataEventStreamingsSinkSinkSLSParametersTopic    `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	SourceDTSParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceDTSParameters `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	// Source Kafka Parameters
	SourceKafkaParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// Source MNS Parameters
	SourceMNSParameters  *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMNSParameters  `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceMQTTParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceMQTTParameters `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	// Source RabbitMQ Parameters
	SourceRabbitMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// Source RocketMQ Parameters
	SourceRocketMQParameters *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	SourceSLSParameters      *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceSLSParameters      `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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
	BrokerUrl      *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	InitCheckPoint *string `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Sid            *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	IsBase64Decode *bool   `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupID    *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Offset     *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp  *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetGroupID(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters) SetInstanceId(v string) *ListEventStreamingsResponseBodyDataEventStreamingsSourceSourceRocketMQParameters {
	s.InstanceId = &v
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
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

type ListEventStreamingsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEventStreamingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName   *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	Limit          *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
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
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Rules     []*ListRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Total     *int32                            `json:"Total,omitempty" xml:"Total,omitempty"`
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
	CreatedTimestamp *int64                                   `json:"CreatedTimestamp,omitempty" xml:"CreatedTimestamp,omitempty"`
	Description      *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DetailMap        map[string]interface{}                   `json:"DetailMap,omitempty" xml:"DetailMap,omitempty"`
	EventBusName     *string                                  `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	FilterPattern    *string                                  `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RuleARN          *string                                  `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
	RuleName         *string                                  `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status           *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Targets          []*ListRulesResponseBodyDataRulesTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	Endpoint     *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PushSelector *string `json:"PushSelector,omitempty" xml:"PushSelector,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Arn          *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	Limit        *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Targets   []*ListTargetsResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	Total     *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
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
	Endpoint     *string                                        `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventBusName *string                                        `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	Id           *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	ParamList    []*ListTargetsResponseBodyDataTargetsParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	RuleName     *string                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Type         *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Form        *string `json:"Form,omitempty" xml:"Form,omitempty"`
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	Template    *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListUserDefinedEventSourcesResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListUserDefinedEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	EventSourceList []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList `json:"EventSourceList,omitempty" xml:"EventSourceList,omitempty" type:"Repeated"`
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

type ListUserDefinedEventSourcesResponseBodyDataEventSourceList struct {
	Arn                       *string                                                                              `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Ctime                     *float32                                                                             `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	EventBusName              *string                                                                              `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	ExternalSourceType        *string                                                                              `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	Name                      *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceHttpEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	SourceMNSParameters       *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters       `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceRabbitMQParameters  *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters  `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQParameters  *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters  `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// SourceSLSParameters
	SourceSLSParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	Status              *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetStatus(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Status = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Type = &v
	return s
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters struct {
	Ip             []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Method         []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	Referer        []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	SecurityConfig *string   `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
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

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters struct {
	IsBase64Decode *bool   `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupId    *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Offset     *string  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId   *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp  *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic      *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.GroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceId = &v
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
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

type ListUserDefinedEventSourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserDefinedEventSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *bool   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PauseEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string                     `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName     *string                     `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Targets      []*PutTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	DeadLetterQueue   *PutTargetsRequestTargetsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	Endpoint          *string                                  `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ErrorsTolerance   *string                                  `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	Id                *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	ParamList         []*PutTargetsRequestTargetsParamList     `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	PushRetryStrategy *string                                  `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
	Type              *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Form        *string `json:"Form,omitempty" xml:"Form,omitempty"`
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	Template    *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	EventBusName  *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PutTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ErrorEntries      []*PutTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	ErrorEntriesCount *int32                                    `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
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
	EntryId      *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId      *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
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

type QueryEventResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId      *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
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
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryEventTracesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Action           *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionTime       *int64  `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	Endpoint         *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventBusName     *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId          *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventSource      *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	NotifyLatency    *string `json:"NotifyLatency,omitempty" xml:"NotifyLatency,omitempty"`
	NotifyStatus     *string `json:"NotifyStatus,omitempty" xml:"NotifyStatus,omitempty"`
	NotifyTime       *int64  `json:"NotifyTime,omitempty" xml:"NotifyTime,omitempty"`
	ReceivedTime     *int64  `json:"ReceivedTime,omitempty" xml:"ReceivedTime,omitempty"`
	RuleMatchingTime *string `json:"RuleMatchingTime,omitempty" xml:"RuleMatchingTime,omitempty"`
	RuleName         *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEventTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId      *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventSource  *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
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
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryTracedEventByEventIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Events    []*QueryTracedEventByEventIdResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	NextToken *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total     *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
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
	EventBusName      *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId           *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventReceivedTime *int64  `json:"EventReceivedTime,omitempty" xml:"EventReceivedTime,omitempty"`
	EventSource       *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	EventType         *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTracedEventByEventIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventSource  *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Limit        *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MatchedRule  *string `json:"MatchedRule,omitempty" xml:"MatchedRule,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryTracedEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Events    []*QueryTracedEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Total     *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
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
	EventBusName      *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventId           *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventReceivedTime *int64  `json:"EventReceivedTime,omitempty" xml:"EventReceivedTime,omitempty"`
	EventSource       *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	EventType         *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTracedEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code      *bool   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartEventStreamingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *StartEventStreamingResponseBody) SetCode(v bool) *StartEventStreamingResponseBody {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Event        *string `json:"Event,omitempty" xml:"Event,omitempty"`
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
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TestEventPatternResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TestEventPatternResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiDestinationName *string                                       `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName     *string                                       `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description        *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpApiParameters  *UpdateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
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
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
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
	ApiDestinationName      *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	ConnectionName          *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApiDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AuthParameters    *UpdateConnectionRequestAuthParameters    `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty" type:"Struct"`
	ConnectionName    *string                                   `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description       *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
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
	ApiKeyAuthParameters *UpdateConnectionRequestAuthParametersApiKeyAuthParameters `json:"ApiKeyAuthParameters,omitempty" xml:"ApiKeyAuthParameters,omitempty" type:"Struct"`
	AuthorizationType    *string                                                    `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	BasicAuthParameters  *UpdateConnectionRequestAuthParametersBasicAuthParameters  `json:"BasicAuthParameters,omitempty" xml:"BasicAuthParameters,omitempty" type:"Struct"`
	OAuthParameters      *UpdateConnectionRequestAuthParametersOAuthParameters      `json:"OAuthParameters,omitempty" xml:"OAuthParameters,omitempty" type:"Struct"`
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
	ApiKeyName  *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	AuthorizationEndpoint *string                                                                  `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	ClientParameters      *UpdateConnectionRequestAuthParametersOAuthParametersClientParameters    `json:"ClientParameters,omitempty" xml:"ClientParameters,omitempty" type:"Struct"`
	HttpMethod            *string                                                                  `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	OAuthHttpParameters   *UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParameters `json:"OAuthHttpParameters,omitempty" xml:"OAuthHttpParameters,omitempty" type:"Struct"`
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
	ClientID     *string `json:"ClientID,omitempty" xml:"ClientID,omitempty"`
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
	BodyParameters        []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersBodyParameters        `json:"BodyParameters,omitempty" xml:"BodyParameters,omitempty" type:"Repeated"`
	HeaderParameters      []*UpdateConnectionRequestAuthParametersOAuthParametersOAuthHttpParametersHeaderParameters      `json:"HeaderParameters,omitempty" xml:"HeaderParameters,omitempty" type:"Repeated"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	IsValueSecret *string `json:"IsValueSecret,omitempty" xml:"IsValueSecret,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitcheId      *string `json:"VswitcheId,omitempty" xml:"VswitcheId,omitempty"`
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
	AuthParametersShrink    *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	ConnectionName          *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 事件源描述详情
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName []byte  `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// 事件源英文Code
	EventSourceName           *string                                            `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	SourceHttpEventParameters *UpdateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	SourceKafkaParameters     *UpdateEventSourceRequestSourceKafkaParameters     `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	SourceMNSParameters       *UpdateEventSourceRequestSourceMNSParameters       `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceRabbitMQParameters  *UpdateEventSourceRequestSourceRabbitMQParameters  `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQParameters  *UpdateEventSourceRequestSourceRocketMQParameters  `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// SourceSLSParameters
	SourceSLSParameters            *UpdateEventSourceRequestSourceSLSParameters            `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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

func (s *UpdateEventSourceRequest) SetEventBusName(v []byte) *UpdateEventSourceRequest {
	s.EventBusName = v
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
	Ip             []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Method         []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	Referer        []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	SecurityConfig *string   `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaximumTasks    *int32  `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	IsBase64Decode *string `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEventSourceRequestSourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v string) *UpdateEventSourceRequestSourceMNSParameters {
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
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupID                 *string  `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceId              *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceNetwork         *string  `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	InstanceSecurityGroupId *string  `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	InstanceType            *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceVSwitchIds      *string  `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	InstanceVpcId           *string  `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	Offset                  *string  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId                *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag                     *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp               *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic                   *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
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

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
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

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v float32) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

type UpdateEventSourceRequestSourceSLSParameters struct {
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	LogStore        *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
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

type UpdateEventSourceShrinkRequest struct {
	// 事件源描述详情
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName []byte  `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// 事件源英文Code
	EventSourceName                 *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	SourceKafkaParametersShrink     *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	SourceMNSParametersShrink       *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	SourceRabbitMQParametersShrink  *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	SourceRocketMQParametersShrink  *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// SourceSLSParameters
	SourceSLSParametersShrink            *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
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

func (s *UpdateEventSourceShrinkRequest) SetEventBusName(v []byte) *UpdateEventSourceShrinkRequest {
	s.EventBusName = v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description        *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string                                `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string                                `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptions         *UpdateEventStreamingRequestRunOptions `json:"RunOptions,omitempty" xml:"RunOptions,omitempty" type:"Struct"`
	Sink               *UpdateEventStreamingRequestSink       `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Struct"`
	Source             *UpdateEventStreamingRequestSource     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
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

type UpdateEventStreamingRequestRunOptions struct {
	BatchWindow     *UpdateEventStreamingRequestRunOptionsBatchWindow     `json:"BatchWindow,omitempty" xml:"BatchWindow,omitempty" type:"Struct"`
	DeadLetterQueue *UpdateEventStreamingRequestRunOptionsDeadLetterQueue `json:"DeadLetterQueue,omitempty" xml:"DeadLetterQueue,omitempty" type:"Struct"`
	ErrorsTolerance *string                                               `json:"ErrorsTolerance,omitempty" xml:"ErrorsTolerance,omitempty"`
	MaximumTasks    *int64                                                `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	RetryStrategy   *UpdateEventStreamingRequestRunOptionsRetryStrategy   `json:"RetryStrategy,omitempty" xml:"RetryStrategy,omitempty" type:"Struct"`
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
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	TimeBasedWindow  *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
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
	MaximumEventAgeInSeconds *int64  `json:"MaximumEventAgeInSeconds,omitempty" xml:"MaximumEventAgeInSeconds,omitempty"`
	MaximumRetryAttempts     *int64  `json:"MaximumRetryAttempts,omitempty" xml:"MaximumRetryAttempts,omitempty"`
	PushRetryStrategy        *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
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
	SinkFcParameters       *UpdateEventStreamingRequestSinkSinkFcParameters       `json:"SinkFcParameters,omitempty" xml:"SinkFcParameters,omitempty" type:"Struct"`
	SinkKafkaParameters    *UpdateEventStreamingRequestSinkSinkKafkaParameters    `json:"SinkKafkaParameters,omitempty" xml:"SinkKafkaParameters,omitempty" type:"Struct"`
	SinkMNSParameters      *UpdateEventStreamingRequestSinkSinkMNSParameters      `json:"SinkMNSParameters,omitempty" xml:"SinkMNSParameters,omitempty" type:"Struct"`
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

func (s *UpdateEventStreamingRequestSink) SetSinkKafkaParameters(v *UpdateEventStreamingRequestSinkSinkKafkaParameters) *UpdateEventStreamingRequestSink {
	s.SinkKafkaParameters = v
	return s
}

func (s *UpdateEventStreamingRequestSink) SetSinkMNSParameters(v *UpdateEventStreamingRequestSinkSinkMNSParameters) *UpdateEventStreamingRequestSink {
	s.SinkMNSParameters = v
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
	Body           *UpdateEventStreamingRequestSinkSinkFcParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Concurrency    *UpdateEventStreamingRequestSinkSinkFcParametersConcurrency    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty" type:"Struct"`
	FunctionName   *UpdateEventStreamingRequestSinkSinkFcParametersFunctionName   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty" type:"Struct"`
	InvocationType *UpdateEventStreamingRequestSinkSinkFcParametersInvocationType `json:"InvocationType,omitempty" xml:"InvocationType,omitempty" type:"Struct"`
	Qualifier      *UpdateEventStreamingRequestSinkSinkFcParametersQualifier      `json:"Qualifier,omitempty" xml:"Qualifier,omitempty" type:"Struct"`
	ServiceName    *UpdateEventStreamingRequestSinkSinkFcParametersServiceName    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type UpdateEventStreamingRequestSinkSinkKafkaParameters struct {
	Acks       *UpdateEventStreamingRequestSinkSinkKafkaParametersAcks       `json:"Acks,omitempty" xml:"Acks,omitempty" type:"Struct"`
	InstanceId *UpdateEventStreamingRequestSinkSinkKafkaParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Key        *UpdateEventStreamingRequestSinkSinkKafkaParametersKey        `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	Topic      *UpdateEventStreamingRequestSinkSinkKafkaParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
	Value      *UpdateEventStreamingRequestSinkSinkKafkaParametersValue      `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body           *UpdateEventStreamingRequestSinkSinkMNSParametersBody           `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	IsBase64Encode *UpdateEventStreamingRequestSinkSinkMNSParametersIsBase64Encode `json:"IsBase64Encode,omitempty" xml:"IsBase64Encode,omitempty" type:"Struct"`
	QueueName      *UpdateEventStreamingRequestSinkSinkMNSParametersQueueName      `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type UpdateEventStreamingRequestSinkSinkRabbitMQParameters struct {
	Body            *UpdateEventStreamingRequestSinkSinkRabbitMQParametersBody            `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Exchange        *UpdateEventStreamingRequestSinkSinkRabbitMQParametersExchange        `json:"Exchange,omitempty" xml:"Exchange,omitempty" type:"Struct"`
	InstanceId      *UpdateEventStreamingRequestSinkSinkRabbitMQParametersInstanceId      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	MessageId       *UpdateEventStreamingRequestSinkSinkRabbitMQParametersMessageId       `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Struct"`
	Properties      *UpdateEventStreamingRequestSinkSinkRabbitMQParametersProperties      `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	QueueName       *UpdateEventStreamingRequestSinkSinkRabbitMQParametersQueueName       `json:"QueueName,omitempty" xml:"QueueName,omitempty" type:"Struct"`
	RoutingKey      *UpdateEventStreamingRequestSinkSinkRabbitMQParametersRoutingKey      `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty" type:"Struct"`
	TargetType      *UpdateEventStreamingRequestSinkSinkRabbitMQParametersTargetType      `json:"TargetType,omitempty" xml:"TargetType,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body       *UpdateEventStreamingRequestSinkSinkRocketMQParametersBody       `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	InstanceId *UpdateEventStreamingRequestSinkSinkRocketMQParametersInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	Keys       *UpdateEventStreamingRequestSinkSinkRocketMQParametersKeys       `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	Properties *UpdateEventStreamingRequestSinkSinkRocketMQParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Tags       *UpdateEventStreamingRequestSinkSinkRocketMQParametersTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic      *UpdateEventStreamingRequestSinkSinkRocketMQParametersTopic      `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Body     *UpdateEventStreamingRequestSinkSinkSLSParametersBody     `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	LogStore *UpdateEventStreamingRequestSinkSinkSLSParametersLogStore `json:"LogStore,omitempty" xml:"LogStore,omitempty" type:"Struct"`
	Project  *UpdateEventStreamingRequestSinkSinkSLSParametersProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RoleName *UpdateEventStreamingRequestSinkSinkSLSParametersRoleName `json:"RoleName,omitempty" xml:"RoleName,omitempty" type:"Struct"`
	Topic    *UpdateEventStreamingRequestSinkSinkSLSParametersTopic    `json:"Topic,omitempty" xml:"Topic,omitempty" type:"Struct"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	SourceDTSParameters      *UpdateEventStreamingRequestSourceSourceDTSParameters      `json:"SourceDTSParameters,omitempty" xml:"SourceDTSParameters,omitempty" type:"Struct"`
	SourceKafkaParameters    *UpdateEventStreamingRequestSourceSourceKafkaParameters    `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	SourceMNSParameters      *UpdateEventStreamingRequestSourceSourceMNSParameters      `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceMQTTParameters     *UpdateEventStreamingRequestSourceSourceMQTTParameters     `json:"SourceMQTTParameters,omitempty" xml:"SourceMQTTParameters,omitempty" type:"Struct"`
	SourceRabbitMQParameters *UpdateEventStreamingRequestSourceSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	SourceRocketMQParameters *UpdateEventStreamingRequestSourceSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	SourceSLSParameters      *UpdateEventStreamingRequestSourceSourceSLSParameters      `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
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
	BrokerUrl      *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	InitCheckPoint *int64  `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Sid            *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	ConsumerGroup   *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Network         *string `json:"Network,omitempty" xml:"Network,omitempty"`
	OffsetReset     *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *UpdateEventStreamingRequestSourceSourceKafkaParameters) SetVpcId(v string) *UpdateEventStreamingRequestSourceSourceKafkaParameters {
	s.VpcId = &v
	return s
}

type UpdateEventStreamingRequestSourceSourceMNSParameters struct {
	IsBase64Decode *bool   `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	QueueName      *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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

type UpdateEventStreamingRequestSourceSourceRabbitMQParameters struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName       *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	GroupID    *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Offset     *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Timestamp  *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventStreamingRequestSourceSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetGroupID(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *UpdateEventStreamingRequestSourceSourceRocketMQParameters) SetInstanceId(v string) *UpdateEventStreamingRequestSourceSourceRocketMQParameters {
	s.InstanceId = &v
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

type UpdateEventStreamingShrinkRequest struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	FilterPattern      *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RunOptionsShrink   *string `json:"RunOptions,omitempty" xml:"RunOptions,omitempty"`
	SinkShrink         *string `json:"Sink,omitempty" xml:"Sink,omitempty"`
	SourceShrink       *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

type UpdateEventStreamingResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EventBusName  *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	FilterPattern *string `json:"FilterPattern,omitempty" xml:"FilterPattern,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) ListUserDefinedEventSourcesWithOptions(runtime *util.RuntimeOptions) (_result *ListUserDefinedEventSourcesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

func (client *Client) ListUserDefinedEventSources() (_result *ListUserDefinedEventSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDefinedEventSourcesResponse{}
	_body, _err := client.ListUserDefinedEventSourcesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
