// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAirflowLoginTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-b3a7f110a6vmvn7xxxxxx
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
}

func (s CreateAirflowLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAirflowLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenRequest) SetAirflowId(v string) *CreateAirflowLoginTokenRequest {
	s.AirflowId = &v
	return s
}

type CreateAirflowLoginTokenResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateAirflowLoginTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4284D079-30F4-5B23-ADC4-28F291622C9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAirflowLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAirflowLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponseBody) SetCode(v string) *CreateAirflowLoginTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetData(v *CreateAirflowLoginTokenResponseBodyData) *CreateAirflowLoginTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetErrorCode(v string) *CreateAirflowLoginTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetHttpStatusCode(v int32) *CreateAirflowLoginTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetMessage(v string) *CreateAirflowLoginTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetRequestId(v string) *CreateAirflowLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBody) SetSuccess(v bool) *CreateAirflowLoginTokenResponseBody {
	s.Success = &v
	return s
}

type CreateAirflowLoginTokenResponseBodyData struct {
	// example:
	//
	// https://data-dms.aliyuncs.com/airflow/x/xxxx/af-ehrmszbxxxxxxx
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// f432d77de03b6b95fc24f91414e29c
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateAirflowLoginTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAirflowLoginTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponseBodyData) SetHost(v string) *CreateAirflowLoginTokenResponseBodyData {
	s.Host = &v
	return s
}

func (s *CreateAirflowLoginTokenResponseBodyData) SetToken(v string) *CreateAirflowLoginTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateAirflowLoginTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAirflowLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAirflowLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAirflowLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponse) SetHeaders(v map[string]*string) *CreateAirflowLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateAirflowLoginTokenResponse) SetStatusCode(v int32) *CreateAirflowLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponse) SetBody(v *CreateAirflowLoginTokenResponseBody) *CreateAirflowLoginTokenResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 获取airflow登录凭证
//
// @param request - CreateAirflowLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAirflowLoginTokenResponse
func (client *Client) CreateAirflowLoginTokenWithOptions(request *CreateAirflowLoginTokenRequest, runtime *util.RuntimeOptions) (_result *CreateAirflowLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AirflowId)) {
		query["AirflowId"] = request.AirflowId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAirflowLoginToken"),
		Version:     tea.String("2025-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAirflowLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取airflow登录凭证
//
// @param request - CreateAirflowLoginTokenRequest
//
// @return CreateAirflowLoginTokenResponse
func (client *Client) CreateAirflowLoginToken(request *CreateAirflowLoginTokenRequest) (_result *CreateAirflowLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAirflowLoginTokenResponse{}
	_body, _err := client.CreateAirflowLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
