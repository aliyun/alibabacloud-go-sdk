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
	// https://example.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://bucket-name.oss-zhangjiakou.aliyuncs.com/path/filename.jsonl
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// abcdefg1234567
	IdempotentId *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	LabelTemplateId *string `json:"LabelTemplateId,omitempty" xml:"LabelTemplateId,omitempty"`
}

func (s CreateLabelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskRequest) SetCallbackUrl(v string) *CreateLabelTaskRequest {
	s.CallbackUrl = &v
	return s
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
	// ABCDEFGH-1234-5678-90AB-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 03s1Xmp791KUMYuZVEb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLabelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLabelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelTaskResponseBody) SetRequestId(v string) *CreateLabelTaskResponseBody {
	s.RequestId = &v
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

type CreateOssUploadTokenRequest struct {
	// example:
	//
	// test-file
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// JSONL
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s CreateOssUploadTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOssUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateOssUploadTokenRequest) SetFileName(v string) *CreateOssUploadTokenRequest {
	s.FileName = &v
	return s
}

func (s *CreateOssUploadTokenRequest) SetFileType(v string) *CreateOssUploadTokenRequest {
	s.FileType = &v
	return s
}

type CreateOssUploadTokenResponseBody struct {
	// example:
	//
	// LTAI5tQPEXqDVu7794Bvw2xM
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 1740758400000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// https://bucket-name.oss-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// path/filename.jsonl
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// vyAnsgE0fgjYGF0W79ryhhhQmec=
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// ABCDEFGH-1234-5678-90AB-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vyAnsgE0fgjYGF0W79ryhhhQmec=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// https://bucket-name.oss-zhangjiakou.aliyuncs.com/path/filename.jsonl
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateOssUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOssUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssUploadTokenResponseBody) SetAccessId(v string) *CreateOssUploadTokenResponseBody {
	s.AccessId = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetExpireTime(v string) *CreateOssUploadTokenResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetHost(v string) *CreateOssUploadTokenResponseBody {
	s.Host = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetKey(v string) *CreateOssUploadTokenResponseBody {
	s.Key = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetPolicy(v string) *CreateOssUploadTokenResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetRequestId(v string) *CreateOssUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetSignature(v string) *CreateOssUploadTokenResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateOssUploadTokenResponseBody) SetUrl(v string) *CreateOssUploadTokenResponseBody {
	s.Url = &v
	return s
}

type CreateOssUploadTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateOssUploadTokenResponse) SetHeaders(v map[string]*string) *CreateOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateOssUploadTokenResponse) SetStatusCode(v int32) *CreateOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssUploadTokenResponse) SetBody(v *CreateOssUploadTokenResponseBody) *CreateOssUploadTokenResponse {
	s.Body = v
	return s
}

type CreateTextLabelRequest struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// abc
	LabelTemplateId *string `json:"LabelTemplateId,omitempty" xml:"LabelTemplateId,omitempty"`
}

func (s CreateTextLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextLabelRequest) GoString() string {
	return s.String()
}

func (s *CreateTextLabelRequest) SetData(v string) *CreateTextLabelRequest {
	s.Data = &v
	return s
}

func (s *CreateTextLabelRequest) SetLabelTemplateId(v string) *CreateTextLabelRequest {
	s.LabelTemplateId = &v
	return s
}

type CreateTextLabelResponseBody struct {
	// example:
	//
	// {\\"tag\\": \\"tag1\\", \\"Emotion\\": \\"1\\"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// ABCDEFGH-1234-5678-90AB-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	Tokens *string `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s CreateTextLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTextLabelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTextLabelResponseBody) SetData(v string) *CreateTextLabelResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTextLabelResponseBody) SetRequestId(v string) *CreateTextLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTextLabelResponseBody) SetTokens(v string) *CreateTextLabelResponseBody {
	s.Tokens = &v
	return s
}

type CreateTextLabelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTextLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTextLabelResponse) GoString() string {
	return s.String()
}

func (s *CreateTextLabelResponse) SetHeaders(v map[string]*string) *CreateTextLabelResponse {
	s.Headers = v
	return s
}

func (s *CreateTextLabelResponse) SetStatusCode(v int32) *CreateTextLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextLabelResponse) SetBody(v *CreateTextLabelResponseBody) *CreateTextLabelResponse {
	s.Body = v
	return s
}

type GetLabelTaskResultRequest struct {
	// example:
	//
	// 03s1Xmp791KUMYuZVEb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// ABCDEFGH-1234-5678-90AB-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http://bucket.oss.cn-zhangjiakou.aliyuncs.com/key
	ResultDataUrl *string `json:"ResultDataUrl,omitempty" xml:"ResultDataUrl,omitempty"`
	// example:
	//
	// IN_PROGRESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	Tokens *string `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
}

func (s GetLabelTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLabelTaskResultResponseBody) GoString() string {
	return s.String()
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

func (s *GetLabelTaskResultResponseBody) SetTokens(v string) *GetLabelTaskResultResponseBody {
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
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		body["IdempotentId"] = request.IdempotentId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelTemplateId)) {
		body["LabelTemplateId"] = request.LabelTemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLabelTask"),
		Version:     tea.String("2025-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tongjian/yic-tongjian/openService/v2/aitag/createLabelTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLabelTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLabelTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 创建OSS上传token
//
// @param request - CreateOssUploadTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOssUploadTokenResponse
func (client *Client) CreateOssUploadTokenWithOptions(request *CreateOssUploadTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOssUploadTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOssUploadToken"),
		Version:     tea.String("2025-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tongjian/yic-tongjian/openService/v2/base/createOssUploadToken"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateOssUploadTokenResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateOssUploadTokenResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建OSS上传token
//
// @param request - CreateOssUploadTokenRequest
//
// @return CreateOssUploadTokenResponse
func (client *Client) CreateOssUploadToken(request *CreateOssUploadTokenRequest) (_result *CreateOssUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOssUploadTokenResponse{}
	_body, _err := client.CreateOssUploadTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 单条文本标注
//
// @param request - CreateTextLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextLabelResponse
func (client *Client) CreateTextLabelWithOptions(request *CreateTextLabelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.LabelTemplateId)) {
		body["LabelTemplateId"] = request.LabelTemplateId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTextLabel"),
		Version:     tea.String("2025-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tongjian/yic-tongjian/openService/v2/aitag/createTextLabel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTextLabelResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTextLabelResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 单条文本标注
//
// @param request - CreateTextLabelRequest
//
// @return CreateTextLabelResponse
func (client *Client) CreateTextLabel(request *CreateTextLabelRequest) (_result *CreateTextLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextLabelResponse{}
	_body, _err := client.CreateTextLabelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型标注任务
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
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLabelTaskResult"),
		Version:     tea.String("2025-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tongjian/yic-tongjian/openService/v2/aitag/getLabelTaskResult"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLabelTaskResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLabelTaskResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询模型标注任务
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
