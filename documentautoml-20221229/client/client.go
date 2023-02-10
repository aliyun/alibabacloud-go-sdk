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

type PredictClassifierModelRequest struct {
	AutoPrediction *bool   `json:"AutoPrediction,omitempty" xml:"AutoPrediction,omitempty"`
	ClassifierId   *string `json:"ClassifierId,omitempty" xml:"ClassifierId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s PredictClassifierModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictClassifierModelRequest) GoString() string {
	return s.String()
}

func (s *PredictClassifierModelRequest) SetAutoPrediction(v bool) *PredictClassifierModelRequest {
	s.AutoPrediction = &v
	return s
}

func (s *PredictClassifierModelRequest) SetClassifierId(v string) *PredictClassifierModelRequest {
	s.ClassifierId = &v
	return s
}

func (s *PredictClassifierModelRequest) SetContent(v string) *PredictClassifierModelRequest {
	s.Content = &v
	return s
}

type PredictClassifierModelResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PredictClassifierModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictClassifierModelResponseBody) GoString() string {
	return s.String()
}

func (s *PredictClassifierModelResponseBody) SetCode(v int32) *PredictClassifierModelResponseBody {
	s.Code = &v
	return s
}

func (s *PredictClassifierModelResponseBody) SetData(v string) *PredictClassifierModelResponseBody {
	s.Data = &v
	return s
}

func (s *PredictClassifierModelResponseBody) SetMessage(v string) *PredictClassifierModelResponseBody {
	s.Message = &v
	return s
}

func (s *PredictClassifierModelResponseBody) SetRequestId(v string) *PredictClassifierModelResponseBody {
	s.RequestId = &v
	return s
}

type PredictClassifierModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PredictClassifierModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictClassifierModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictClassifierModelResponse) GoString() string {
	return s.String()
}

func (s *PredictClassifierModelResponse) SetHeaders(v map[string]*string) *PredictClassifierModelResponse {
	s.Headers = v
	return s
}

func (s *PredictClassifierModelResponse) SetStatusCode(v int32) *PredictClassifierModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictClassifierModelResponse) SetBody(v *PredictClassifierModelResponseBody) *PredictClassifierModelResponse {
	s.Body = v
	return s
}

type PredictModelRequest struct {
	BinaryToText *bool   `json:"BinaryToText,omitempty" xml:"BinaryToText,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ModelId      *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
}

func (s PredictModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictModelRequest) GoString() string {
	return s.String()
}

func (s *PredictModelRequest) SetBinaryToText(v bool) *PredictModelRequest {
	s.BinaryToText = &v
	return s
}

func (s *PredictModelRequest) SetContent(v string) *PredictModelRequest {
	s.Content = &v
	return s
}

func (s *PredictModelRequest) SetModelId(v int64) *PredictModelRequest {
	s.ModelId = &v
	return s
}

func (s *PredictModelRequest) SetModelVersion(v string) *PredictModelRequest {
	s.ModelVersion = &v
	return s
}

type PredictModelResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PredictModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictModelResponseBody) GoString() string {
	return s.String()
}

func (s *PredictModelResponseBody) SetCode(v int32) *PredictModelResponseBody {
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

type PredictModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PredictModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type PredictTemplateModelRequest struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PredictTemplateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictTemplateModelRequest) GoString() string {
	return s.String()
}

func (s *PredictTemplateModelRequest) SetContent(v string) *PredictTemplateModelRequest {
	s.Content = &v
	return s
}

func (s *PredictTemplateModelRequest) SetProjectId(v int64) *PredictTemplateModelRequest {
	s.ProjectId = &v
	return s
}

type PredictTemplateModelResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PredictTemplateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictTemplateModelResponseBody) GoString() string {
	return s.String()
}

func (s *PredictTemplateModelResponseBody) SetCode(v string) *PredictTemplateModelResponseBody {
	s.Code = &v
	return s
}

func (s *PredictTemplateModelResponseBody) SetData(v string) *PredictTemplateModelResponseBody {
	s.Data = &v
	return s
}

func (s *PredictTemplateModelResponseBody) SetMessage(v string) *PredictTemplateModelResponseBody {
	s.Message = &v
	return s
}

func (s *PredictTemplateModelResponseBody) SetRequestId(v string) *PredictTemplateModelResponseBody {
	s.RequestId = &v
	return s
}

type PredictTemplateModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PredictTemplateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PredictTemplateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictTemplateModelResponse) GoString() string {
	return s.String()
}

func (s *PredictTemplateModelResponse) SetHeaders(v map[string]*string) *PredictTemplateModelResponse {
	s.Headers = v
	return s
}

func (s *PredictTemplateModelResponse) SetStatusCode(v int32) *PredictTemplateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictTemplateModelResponse) SetBody(v *PredictTemplateModelResponseBody) *PredictTemplateModelResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("documentautoml"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) PredictClassifierModelWithOptions(request *PredictClassifierModelRequest, runtime *util.RuntimeOptions) (_result *PredictClassifierModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPrediction)) {
		query["AutoPrediction"] = request.AutoPrediction
	}

	if !tea.BoolValue(util.IsUnset(request.ClassifierId)) {
		query["ClassifierId"] = request.ClassifierId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictClassifierModel"),
		Version:     tea.String("2022-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictClassifierModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictClassifierModel(request *PredictClassifierModelRequest) (_result *PredictClassifierModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictClassifierModelResponse{}
	_body, _err := client.PredictClassifierModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredictModelWithOptions(request *PredictModelRequest, runtime *util.RuntimeOptions) (_result *PredictModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BinaryToText)) {
		query["BinaryToText"] = request.BinaryToText
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		query["ModelVersion"] = request.ModelVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictModel"),
		Version:     tea.String("2022-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
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

func (client *Client) PredictModel(request *PredictModelRequest) (_result *PredictModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictModelResponse{}
	_body, _err := client.PredictModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PredictTemplateModelWithOptions(request *PredictTemplateModelRequest, runtime *util.RuntimeOptions) (_result *PredictTemplateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictTemplateModel"),
		Version:     tea.String("2022-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictTemplateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictTemplateModel(request *PredictTemplateModelRequest) (_result *PredictTemplateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictTemplateModelResponse{}
	_body, _err := client.PredictTemplateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
