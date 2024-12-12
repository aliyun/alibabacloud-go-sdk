// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OpenApiModelsPredictCmd struct {
}

func (s OpenApiModelsPredictCmd) String() string {
	return tea.Prettify(s)
}

func (s OpenApiModelsPredictCmd) GoString() string {
	return s.String()
}

type PredictRequest struct {
	// example:
	//
	// db_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	Input  *string `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// _polar4ai_tongyi
	ModelClass *string                `json:"modelClass,omitempty" xml:"modelClass,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s PredictRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictRequest) GoString() string {
	return s.String()
}

func (s *PredictRequest) SetDbName(v string) *PredictRequest {
	s.DbName = &v
	return s
}

func (s *PredictRequest) SetInput(v string) *PredictRequest {
	s.Input = &v
	return s
}

func (s *PredictRequest) SetInstanceName(v string) *PredictRequest {
	s.InstanceName = &v
	return s
}

func (s *PredictRequest) SetModelClass(v string) *PredictRequest {
	s.ModelClass = &v
	return s
}

func (s *PredictRequest) SetParameters(v map[string]interface{}) *PredictRequest {
	s.Parameters = v
	return s
}

type PredictResponseBody struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ER_ILLEGAL_MODEL_CLASS
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// Illegal model class.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictResponseBody) GoString() string {
	return s.String()
}

func (s *PredictResponseBody) SetData(v interface{}) *PredictResponseBody {
	s.Data = v
	return s
}

func (s *PredictResponseBody) SetErrCode(v string) *PredictResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PredictResponseBody) SetErrMessage(v string) *PredictResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PredictResponseBody) SetRequestId(v string) *PredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictResponseBody) SetSuccess(v bool) *PredictResponseBody {
	s.Success = &v
	return s
}

type PredictResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredictResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictResponse) GoString() string {
	return s.String()
}

func (s *PredictResponse) SetHeaders(v map[string]*string) *PredictResponse {
	s.Headers = v
	return s
}

func (s *PredictResponse) SetStatusCode(v int32) *PredictResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictResponse) SetBody(v *PredictResponseBody) *PredictResponse {
	s.Body = v
	return s
}

type PredictSseRequest struct {
	// example:
	//
	// db_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	Input  *string `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// _polar4ai_tongyi
	ModelClass *string `json:"modelClass,omitempty" xml:"modelClass,omitempty"`
	// example:
	//
	// {"basic_index_name":"index_table"}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s PredictSseRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictSseRequest) GoString() string {
	return s.String()
}

func (s *PredictSseRequest) SetDbName(v string) *PredictSseRequest {
	s.DbName = &v
	return s
}

func (s *PredictSseRequest) SetInput(v string) *PredictSseRequest {
	s.Input = &v
	return s
}

func (s *PredictSseRequest) SetInstanceName(v string) *PredictSseRequest {
	s.InstanceName = &v
	return s
}

func (s *PredictSseRequest) SetModelClass(v string) *PredictSseRequest {
	s.ModelClass = &v
	return s
}

func (s *PredictSseRequest) SetParameters(v map[string]interface{}) *PredictSseRequest {
	s.Parameters = v
	return s
}

type PredictSseResponseBody struct {
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ER_ILLEGAL_MODEL_CLASS
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// Illegal model class.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PredictSseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictSseResponseBody) GoString() string {
	return s.String()
}

func (s *PredictSseResponseBody) SetData(v interface{}) *PredictSseResponseBody {
	s.Data = v
	return s
}

func (s *PredictSseResponseBody) SetErrCode(v string) *PredictSseResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PredictSseResponseBody) SetErrMessage(v string) *PredictSseResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PredictSseResponseBody) SetRequestId(v string) *PredictSseResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictSseResponseBody) SetSuccess(v bool) *PredictSseResponseBody {
	s.Success = &v
	return s
}

type PredictSseResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredictSseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredictSseResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictSseResponse) GoString() string {
	return s.String()
}

func (s *PredictSseResponse) SetHeaders(v map[string]*string) *PredictSseResponse {
	s.Headers = v
	return s
}

func (s *PredictSseResponse) SetStatusCode(v int32) *PredictSseResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictSseResponse) SetBody(v *PredictSseResponseBody) *PredictSseResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("polardbai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 模型预测
//
// @param request - PredictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PredictResponse
func (client *Client) PredictWithOptions(request *PredictRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["dbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelClass)) {
		body["modelClass"] = request.ModelClass
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Predict"),
		Version:     tea.String("2024-08-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/openapi/models/predict"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型预测
//
// @param request - PredictRequest
//
// @return PredictResponse
func (client *Client) Predict(request *PredictRequest) (_result *PredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PredictResponse{}
	_body, _err := client.PredictWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模型推理（在线，离线）
//
// @param request - PredictSseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PredictSseResponse
func (client *Client) PredictSseWithOptions(request *PredictSseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PredictSseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["dbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelClass)) {
		body["modelClass"] = request.ModelClass
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictSse"),
		Version:     tea.String("2024-08-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1/openapi/models/predictSse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictSseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型推理（在线，离线）
//
// @param request - PredictSseRequest
//
// @return PredictSseResponse
func (client *Client) PredictSse(request *PredictSseRequest) (_result *PredictSseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PredictSseResponse{}
	_body, _err := client.PredictSseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
