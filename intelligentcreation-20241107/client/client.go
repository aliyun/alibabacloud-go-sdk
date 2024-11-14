// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateLabelTaskRequest struct {
	// example:
	//
	// oss://test/test
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// example:
	//
	// 7caa09aa60204086909ba3958810a567
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	// example:
	//
	// 1_for_7_label_v1
	LabelTemplateId *string `json:"labelTemplateId,omitempty" xml:"labelTemplateId,omitempty"`
}

func (s CreateLabelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskRequest) SetFileUrl(v string) *CreateLabelTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateLabelTaskRequest) SetIdempotentId(v string) *CreateLabelTaskRequest {
	s.IdempotentId = &v
	return s
}

func (s *CreateLabelTaskRequest) SetLabelTemplateId(v string) *CreateLabelTaskRequest {
	s.LabelTemplateId = &v
	return s
}

type CreateLabelTaskResponseBody struct {
	// example:
	//
	// PARAM_ERROR
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 56AC346B-AF40-5E4F-AFFE-FD8BA5E6FB3A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 829682927337963520
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateLabelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskResponseBody) SetErrorCode(v string) *CreateLabelTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLabelTaskResponseBody) SetErrorMessage(v string) *CreateLabelTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLabelTaskResponseBody) SetRequestId(v string) *CreateLabelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLabelTaskResponseBody) SetSuccess(v bool) *CreateLabelTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLabelTaskResponseBody) SetTaskId(v string) *CreateLabelTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateLabelTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLabelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLabelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskResponse) SetHeaders(v map[string]*string) *CreateLabelTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLabelTaskResponse) SetStatusCode(v int32) *CreateLabelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLabelTaskResponse) SetBody(v *CreateLabelTaskResponseBody) *CreateLabelTaskResponse {
	s.Body = v
	return s
}

type GetLabelTaskResultRequest struct {
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetLabelTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetLabelTaskResultRequest) SetTaskId(v string) *GetLabelTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetLabelTaskResultResponseBody struct {
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// https://yic-pre.oss-cn-hangzhou.aliyuncs.com/ai_tag_algo_log/1539704706413278/829593441691238400/merge_label_final_result?Expires=1732083305&OSSAccessKeyId=LTAI5tPHLyFPhh4UoRias4Zg&Signature=Am3cBxlc6hYFKtdGlw9o1R26Vsk%3D
	ResultDataUrl *string `json:"resultDataUrl,omitempty" xml:"resultDataUrl,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 20023
	Tokens *int64 `json:"tokens,omitempty" xml:"tokens,omitempty"`
}

func (s GetLabelTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetLabelTaskResultResponseBody) SetMessage(v string) *GetLabelTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetLabelTaskResultResponseBody) SetRequestId(v string) *GetLabelTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLabelTaskResultResponseBody) SetResultDataUrl(v string) *GetLabelTaskResultResponseBody {
	s.ResultDataUrl = &v
	return s
}

func (s *GetLabelTaskResultResponseBody) SetStatus(v string) *GetLabelTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetLabelTaskResultResponseBody) SetTokens(v int64) *GetLabelTaskResultResponseBody {
	s.Tokens = &v
	return s
}

type GetLabelTaskResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLabelTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLabelTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetLabelTaskResultResponse) SetHeaders(v map[string]*string) *GetLabelTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetLabelTaskResultResponse) SetStatusCode(v int32) *GetLabelTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLabelTaskResultResponse) SetBody(v *GetLabelTaskResultResponseBody) *GetLabelTaskResultResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建模型标注任务
//
// @param request - CreateLabelTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLabelTaskResponse
func (client *Client) CreateLabelTaskWithOptions(request *CreateLabelTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLabelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		body["idempotentId"] = request.IdempotentId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelTemplateId)) {
		body["labelTemplateId"] = request.LabelTemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabelTask"),
		Version:     tea.String("2024-11-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aitag/createLabelTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLabelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型标注任务
//
// @param request - CreateLabelTaskRequest
//
// @return CreateLabelTaskResponse
func (client *Client) CreateLabelTask(request *CreateLabelTaskRequest) (_result *CreateLabelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLabelTaskResponse{}
	_body, _err := client.CreateLabelTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型标注任务结果
//
// @param request - GetLabelTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLabelTaskResultResponse
func (client *Client) GetLabelTaskResultWithOptions(request *GetLabelTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLabelTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLabelTaskResult"),
		Version:     tea.String("2024-11-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/aitag/getLabelTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLabelTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型标注任务结果
//
// @param request - GetLabelTaskResultRequest
//
// @return GetLabelTaskResultResponse
func (client *Client) GetLabelTaskResult(request *GetLabelTaskResultRequest) (_result *GetLabelTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLabelTaskResultResponse{}
	_body, _err := client.GetLabelTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
