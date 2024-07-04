// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PredictModelRequest struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
}

func (s PredictModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictModelRequest) GoString() string {
	return s.String()
}

func (s *PredictModelRequest) SetContent(v string) *PredictModelRequest {
	s.Content = &v
	return s
}

func (s *PredictModelRequest) SetModelId(v int64) *PredictModelRequest {
	s.ModelId = &v
	return s
}

type PredictModelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PredictModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictModelResponseBody) GoString() string {
	return s.String()
}

func (s *PredictModelResponseBody) SetCode(v string) *PredictModelResponseBody {
	s.Code = &v
	return s
}

func (s *PredictModelResponseBody) SetData(v string) *PredictModelResponseBody {
	s.Data = &v
	return s
}

func (s *PredictModelResponseBody) SetMessage(v string) *PredictModelResponseBody {
	s.Message = &v
	return s
}

func (s *PredictModelResponseBody) SetRequestId(v string) *PredictModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictModelResponseBody) SetSuccess(v string) *PredictModelResponseBody {
	s.Success = &v
	return s
}

type PredictModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredictModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredictModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictModelResponse) GoString() string {
	return s.String()
}

func (s *PredictModelResponse) SetHeaders(v map[string]*string) *PredictModelResponse {
	s.Headers = v
	return s
}

func (s *PredictModelResponse) SetStatusCode(v int32) *PredictModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictModelResponse) SetBody(v *PredictModelResponseBody) *PredictModelResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nlp-automl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - PredictModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PredictModelResponse
func (client *Client) PredictModelWithOptions(request *PredictModelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PredictModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictModel"),
		Version:     tea.String("2019-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/automl/predict"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PredictModelRequest
//
// @return PredictModelResponse
func (client *Client) PredictModel(request *PredictModelRequest) (_result *PredictModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PredictModelResponse{}
	_body, _err := client.PredictModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
