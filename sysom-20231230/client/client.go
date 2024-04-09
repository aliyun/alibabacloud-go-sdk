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

type AuthDiagnosisRequest struct {
	Instances []*AuthDiagnosisRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s AuthDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequest) SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest {
	s.Instances = v
	return s
}

type AuthDiagnosisRequestInstances struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s AuthDiagnosisRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequestInstances) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequestInstances) SetInstance(v string) *AuthDiagnosisRequestInstances {
	s.Instance = &v
	return s
}

func (s *AuthDiagnosisRequestInstances) SetRegion(v string) *AuthDiagnosisRequestInstances {
	s.Region = &v
	return s
}

type AuthDiagnosisResponseBody struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string     `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s AuthDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponseBody) SetCode(v string) *AuthDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetData(v interface{}) *AuthDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *AuthDiagnosisResponseBody) SetMessage(v string) *AuthDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetRequestId(v string) *AuthDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type AuthDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponse) SetHeaders(v map[string]*string) *AuthDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *AuthDiagnosisResponse) SetStatusCode(v int32) *AuthDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthDiagnosisResponse) SetBody(v *AuthDiagnosisResponseBody) *AuthDiagnosisResponse {
	s.Body = v
	return s
}

type GenerateCopilotResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseRequest) SetLlmParamString(v string) *GenerateCopilotResponseRequest {
	s.LlmParamString = &v
	return s
}

type GenerateCopilotResponseResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	Massage *string `json:"massage,omitempty" xml:"massage,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponseBody) SetCode(v string) *GenerateCopilotResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetData(v string) *GenerateCopilotResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetMassage(v string) *GenerateCopilotResponseResponseBody {
	s.Massage = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetRequestId(v string) *GenerateCopilotResponseResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCopilotResponseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotResponseResponse) SetStatusCode(v int32) *GenerateCopilotResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotResponseResponse) SetBody(v *GenerateCopilotResponseResponseBody) *GenerateCopilotResponseResponse {
	s.Body = v
	return s
}

type GetDiagnosisResultRequest struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) SetTaskId(v string) *GetDiagnosisResultRequest {
	s.TaskId = &v
	return s
}

type GetDiagnosisResultResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data      *GetDiagnosisResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetDiagnosisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBody) SetCode(v string) *GetDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetData(v *GetDiagnosisResultResponseBodyData) *GetDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetMessage(v string) *GetDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetRequestId(v string) *GetDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDiagnosisResultResponseBodyData struct {
	Code        *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Command     interface{} `json:"command,omitempty" xml:"command,omitempty"`
	ErrMsg      *string     `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	Params      interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Result      interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ServiceName *string     `json:"service_name,omitempty" xml:"service_name,omitempty"`
	Status      *string     `json:"status,omitempty" xml:"status,omitempty"`
	TaskId      *string     `json:"task_id,omitempty" xml:"task_id,omitempty"`
	Url         *string     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDiagnosisResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBodyData) SetCode(v int32) *GetDiagnosisResultResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetCommand(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Command = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetErrMsg(v string) *GetDiagnosisResultResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetParams(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Params = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetResult(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetServiceName(v string) *GetDiagnosisResultResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetStatus(v string) *GetDiagnosisResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetTaskId(v string) *GetDiagnosisResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetUrl(v string) *GetDiagnosisResultResponseBodyData {
	s.Url = &v
	return s
}

type GetDiagnosisResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisResultResponse) SetStatusCode(v int32) *GetDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosisResultResponse) SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse {
	s.Body = v
	return s
}

type InvokeDiagnosisRequest struct {
	Channel     *string `json:"channel,omitempty" xml:"channel,omitempty"`
	Params      *string `json:"params,omitempty" xml:"params,omitempty"`
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s InvokeDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisRequest) SetChannel(v string) *InvokeDiagnosisRequest {
	s.Channel = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetParams(v string) *InvokeDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetServiceName(v string) *InvokeDiagnosisRequest {
	s.ServiceName = &v
	return s
}

type InvokeDiagnosisResponseBody struct {
	Code      *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data      *InvokeDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                          `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                          `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s InvokeDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBody) SetCode(v string) *InvokeDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetData(v *InvokeDiagnosisResponseBodyData) *InvokeDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetMessage(v string) *InvokeDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetRequestId(v string) *InvokeDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type InvokeDiagnosisResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InvokeDiagnosisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBodyData) SetTaskId(v string) *InvokeDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

type InvokeDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeDiagnosisResponse) SetStatusCode(v int32) *InvokeDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDiagnosisResponse) SetBody(v *InvokeDiagnosisResponseBody) *InvokeDiagnosisResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("sysom"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AuthDiagnosisWithOptions(request *AuthDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AuthDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/auth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthDiagnosis(request *AuthDiagnosisRequest) (_result *AuthDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.AuthDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateCopilotResponseWithOptions(request *GenerateCopilotResponseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateCopilotResponseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LlmParamString)) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCopilotResponse"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/copilot/generate_copilot_response"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateCopilotResponse(request *GenerateCopilotResponseRequest) (_result *GenerateCopilotResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.GenerateCopilotResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnosisResultWithOptions(request *GetDiagnosisResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisResult"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/get_diagnosis_results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnosisResult(request *GetDiagnosisResultRequest) (_result *GetDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.GetDiagnosisResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeDiagnosisWithOptions(request *InvokeDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokeDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["service_name"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/invoke_diagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeDiagnosis(request *InvokeDiagnosisRequest) (_result *InvokeDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.InvokeDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
