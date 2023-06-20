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

type GetPhoneNumberIdentificationResultRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SessionId            *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionPayload       *string `json:"SessionPayload,omitempty" xml:"SessionPayload,omitempty"`
}

func (s GetPhoneNumberIdentificationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultRequest) SetAuthCode(v string) *GetPhoneNumberIdentificationResultRequest {
	s.AuthCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetOutId(v string) *GetPhoneNumberIdentificationResultRequest {
	s.OutId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetPhoneNumber(v string) *GetPhoneNumberIdentificationResultRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetSessionId(v string) *GetPhoneNumberIdentificationResultRequest {
	s.SessionId = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultRequest) SetSessionPayload(v string) *GetPhoneNumberIdentificationResultRequest {
	s.SessionPayload = &v
	return s
}

type GetPhoneNumberIdentificationResultResponseBody struct {
	Code      *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetPhoneNumberIdentificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneNumberIdentificationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetCode(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetData(v *GetPhoneNumberIdentificationResultResponseBodyData) *GetPhoneNumberIdentificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetMessage(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponseBody) SetRequestId(v string) *GetPhoneNumberIdentificationResultResponseBody {
	s.RequestId = &v
	return s
}

type GetPhoneNumberIdentificationResultResponseBodyData struct {
	IsIdentified *string `json:"IsIdentified,omitempty" xml:"IsIdentified,omitempty"`
}

func (s GetPhoneNumberIdentificationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponseBodyData) SetIsIdentified(v string) *GetPhoneNumberIdentificationResultResponseBodyData {
	s.IsIdentified = &v
	return s
}

type GetPhoneNumberIdentificationResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPhoneNumberIdentificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhoneNumberIdentificationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationResultResponse) SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponse) SetStatusCode(v int32) *GetPhoneNumberIdentificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationResultResponse) SetBody(v *GetPhoneNumberIdentificationResultResponseBody) *GetPhoneNumberIdentificationResultResponse {
	s.Body = v
	return s
}

type GetPhoneNumberIdentificationUrlRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Ip                   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	RememberPhoneNumber  *bool   `json:"RememberPhoneNumber,omitempty" xml:"RememberPhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetAuthCode(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.AuthCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetIp(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.Ip = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetOutId(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.OutId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetPhoneNumber(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetRememberPhoneNumber(v bool) *GetPhoneNumberIdentificationUrlRequest {
	s.RememberPhoneNumber = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetResourceOwnerAccount(v string) *GetPhoneNumberIdentificationUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlRequest) SetResourceOwnerId(v int64) *GetPhoneNumberIdentificationUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetPhoneNumberIdentificationUrlResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetPhoneNumberIdentificationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetCode(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetData(v *GetPhoneNumberIdentificationUrlResponseBodyData) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetMessage(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBody) SetRequestId(v string) *GetPhoneNumberIdentificationUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetPhoneNumberIdentificationUrlResponseBodyData struct {
	IdentificationUrl *string `json:"IdentificationUrl,omitempty" xml:"IdentificationUrl,omitempty"`
	SessionId         *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetPhoneNumberIdentificationUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) SetIdentificationUrl(v string) *GetPhoneNumberIdentificationUrlResponseBodyData {
	s.IdentificationUrl = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponseBodyData) SetSessionId(v string) *GetPhoneNumberIdentificationUrlResponseBodyData {
	s.SessionId = &v
	return s
}

type GetPhoneNumberIdentificationUrlResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPhoneNumberIdentificationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhoneNumberIdentificationUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberIdentificationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetHeaders(v map[string]*string) *GetPhoneNumberIdentificationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetStatusCode(v int32) *GetPhoneNumberIdentificationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberIdentificationUrlResponse) SetBody(v *GetPhoneNumberIdentificationUrlResponseBody) *GetPhoneNumberIdentificationUrlResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dytnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetPhoneNumberIdentificationResultWithOptions(request *GetPhoneNumberIdentificationResultRequest, runtime *util.RuntimeOptions) (_result *GetPhoneNumberIdentificationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPayload)) {
		query["SessionPayload"] = request.SessionPayload
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneNumberIdentificationResult"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneNumberIdentificationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhoneNumberIdentificationResult(request *GetPhoneNumberIdentificationResultRequest) (_result *GetPhoneNumberIdentificationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhoneNumberIdentificationResultResponse{}
	_body, _err := client.GetPhoneNumberIdentificationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhoneNumberIdentificationUrlWithOptions(request *GetPhoneNumberIdentificationUrlRequest, runtime *util.RuntimeOptions) (_result *GetPhoneNumberIdentificationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RememberPhoneNumber)) {
		query["RememberPhoneNumber"] = request.RememberPhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneNumberIdentificationUrl"),
		Version:     tea.String("2023-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneNumberIdentificationUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhoneNumberIdentificationUrl(request *GetPhoneNumberIdentificationUrlRequest) (_result *GetPhoneNumberIdentificationUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhoneNumberIdentificationUrlResponse{}
	_body, _err := client.GetPhoneNumberIdentificationUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
