// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryTokenForMnsQueueRequest struct {
	// This parameter is required.
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTokenForMnsQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenForMnsQueueRequest) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueRequest) SetMessageType(v string) *QueryTokenForMnsQueueRequest {
	s.MessageType = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetOwnerId(v int64) *QueryTokenForMnsQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetQueueName(v string) *QueryTokenForMnsQueueRequest {
	s.QueueName = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetResourceOwnerAccount(v string) *QueryTokenForMnsQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetResourceOwnerId(v int64) *QueryTokenForMnsQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTokenForMnsQueueResponseBody struct {
	Code            *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageTokenDTO *QueryTokenForMnsQueueResponseBodyMessageTokenDTO `json:"MessageTokenDTO,omitempty" xml:"MessageTokenDTO,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTokenForMnsQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenForMnsQueueResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponseBody) SetCode(v string) *QueryTokenForMnsQueueResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetMessage(v string) *QueryTokenForMnsQueueResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetMessageTokenDTO(v *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) *QueryTokenForMnsQueueResponseBody {
	s.MessageTokenDTO = v
	return s
}

func (s *QueryTokenForMnsQueueResponseBody) SetRequestId(v string) *QueryTokenForMnsQueueResponseBody {
	s.RequestId = &v
	return s
}

type QueryTokenForMnsQueueResponseBodyMessageTokenDTO struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s QueryTokenForMnsQueueResponseBodyMessageTokenDTO) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenForMnsQueueResponseBodyMessageTokenDTO) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetAccessKeyId(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.AccessKeyId = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetAccessKeySecret(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.AccessKeySecret = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetCreateTime(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.CreateTime = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetExpireTime(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.ExpireTime = &v
	return s
}

func (s *QueryTokenForMnsQueueResponseBodyMessageTokenDTO) SetSecurityToken(v string) *QueryTokenForMnsQueueResponseBodyMessageTokenDTO {
	s.SecurityToken = &v
	return s
}

type QueryTokenForMnsQueueResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTokenForMnsQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTokenForMnsQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenForMnsQueueResponse) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueResponse) SetHeaders(v map[string]*string) *QueryTokenForMnsQueueResponse {
	s.Headers = v
	return s
}

func (s *QueryTokenForMnsQueueResponse) SetStatusCode(v int32) *QueryTokenForMnsQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTokenForMnsQueueResponse) SetBody(v *QueryTokenForMnsQueueResponseBody) *QueryTokenForMnsQueueResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dybaseapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - QueryTokenForMnsQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTokenForMnsQueueResponse
func (client *Client) QueryTokenForMnsQueueWithOptions(request *QueryTokenForMnsQueueRequest, runtime *util.RuntimeOptions) (_result *QueryTokenForMnsQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
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
		Action:      tea.String("QueryTokenForMnsQueue"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTokenForMnsQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTokenForMnsQueueRequest
//
// @return QueryTokenForMnsQueueResponse
func (client *Client) QueryTokenForMnsQueue(request *QueryTokenForMnsQueueRequest) (_result *QueryTokenForMnsQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTokenForMnsQueueResponse{}
	_body, _err := client.QueryTokenForMnsQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
