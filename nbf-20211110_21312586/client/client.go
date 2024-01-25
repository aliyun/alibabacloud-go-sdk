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

type ApiTestRequest struct {
	TestCmd *string `json:"testCmd,omitempty" xml:"testCmd,omitempty"`
}

func (s ApiTestRequest) String() string {
	return tea.Prettify(s)
}

func (s ApiTestRequest) GoString() string {
	return s.String()
}

func (s *ApiTestRequest) SetTestCmd(v string) *ApiTestRequest {
	s.TestCmd = &v
	return s
}

type ApiTestResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApiTestResponse) String() string {
	return tea.Prettify(s)
}

func (s ApiTestResponse) GoString() string {
	return s.String()
}

func (s *ApiTestResponse) SetHeaders(v map[string]*string) *ApiTestResponse {
	s.Headers = v
	return s
}

func (s *ApiTestResponse) SetStatusCode(v int32) *ApiTestResponse {
	s.StatusCode = &v
	return s
}

func (s *ApiTestResponse) SetBody(v string) *ApiTestResponse {
	s.Body = &v
	return s
}

type BuildSdkRequest struct {
	BuildCmd *string `json:"buildCmd,omitempty" xml:"buildCmd,omitempty"`
}

func (s BuildSdkRequest) String() string {
	return tea.Prettify(s)
}

func (s BuildSdkRequest) GoString() string {
	return s.String()
}

func (s *BuildSdkRequest) SetBuildCmd(v string) *BuildSdkRequest {
	s.BuildCmd = &v
	return s
}

type BuildSdkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildSdkResponse) String() string {
	return tea.Prettify(s)
}

func (s BuildSdkResponse) GoString() string {
	return s.String()
}

func (s *BuildSdkResponse) SetHeaders(v map[string]*string) *BuildSdkResponse {
	s.Headers = v
	return s
}

func (s *BuildSdkResponse) SetStatusCode(v int32) *BuildSdkResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildSdkResponse) SetBody(v string) *BuildSdkResponse {
	s.Body = &v
	return s
}

type CreateAndReleaseApiRequest struct {
	CreatApiCmd *string `json:"creatApiCmd,omitempty" xml:"creatApiCmd,omitempty"`
}

func (s CreateAndReleaseApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndReleaseApiRequest) GoString() string {
	return s.String()
}

func (s *CreateAndReleaseApiRequest) SetCreatApiCmd(v string) *CreateAndReleaseApiRequest {
	s.CreatApiCmd = &v
	return s
}

type CreateAndReleaseApiResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndReleaseApiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndReleaseApiResponse) GoString() string {
	return s.String()
}

func (s *CreateAndReleaseApiResponse) SetHeaders(v map[string]*string) *CreateAndReleaseApiResponse {
	s.Headers = v
	return s
}

func (s *CreateAndReleaseApiResponse) SetStatusCode(v int32) *CreateAndReleaseApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndReleaseApiResponse) SetBody(v string) *CreateAndReleaseApiResponse {
	s.Body = &v
	return s
}

type CreateSdkVersionRequest struct {
	CreateSdkCmd *string `json:"createSdkCmd,omitempty" xml:"createSdkCmd,omitempty"`
}

func (s CreateSdkVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSdkVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateSdkVersionRequest) SetCreateSdkCmd(v string) *CreateSdkVersionRequest {
	s.CreateSdkCmd = &v
	return s
}

type CreateSdkVersionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSdkVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSdkVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateSdkVersionResponse) SetHeaders(v map[string]*string) *CreateSdkVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateSdkVersionResponse) SetStatusCode(v int32) *CreateSdkVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSdkVersionResponse) SetBody(v string) *CreateSdkVersionResponse {
	s.Body = &v
	return s
}

type DeleteApiRequest struct {
	ApiExternalId *string `json:"apiExternalId,omitempty" xml:"apiExternalId,omitempty"`
}

func (s DeleteApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiRequest) SetApiExternalId(v string) *DeleteApiRequest {
	s.ApiExternalId = &v
	return s
}

type DeleteApiResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiResponse) SetHeaders(v map[string]*string) *DeleteApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiResponse) SetStatusCode(v int32) *DeleteApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiResponse) SetBody(v string) *DeleteApiResponse {
	s.Body = &v
	return s
}

type GetResultRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResultRequest) GoString() string {
	return s.String()
}

func (s *GetResultRequest) SetTaskId(v string) *GetResultRequest {
	s.TaskId = &v
	return s
}

type GetResultResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResultResponse) GoString() string {
	return s.String()
}

func (s *GetResultResponse) SetHeaders(v map[string]*string) *GetResultResponse {
	s.Headers = v
	return s
}

func (s *GetResultResponse) SetStatusCode(v int32) *GetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResultResponse) SetBody(v string) *GetResultResponse {
	s.Body = &v
	return s
}

type OpenApiGenericProxyResponseBody struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OpenApiGenericProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenApiGenericProxyResponseBody) GoString() string {
	return s.String()
}

func (s *OpenApiGenericProxyResponseBody) SetData(v string) *OpenApiGenericProxyResponseBody {
	s.Data = &v
	return s
}

type OpenApiGenericProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenApiGenericProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenApiGenericProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenApiGenericProxyResponse) GoString() string {
	return s.String()
}

func (s *OpenApiGenericProxyResponse) SetHeaders(v map[string]*string) *OpenApiGenericProxyResponse {
	s.Headers = v
	return s
}

func (s *OpenApiGenericProxyResponse) SetStatusCode(v int32) *OpenApiGenericProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenApiGenericProxyResponse) SetBody(v *OpenApiGenericProxyResponseBody) *OpenApiGenericProxyResponse {
	s.Body = v
	return s
}

type PreCheckRequest struct {
	ApiSchemaDTO          *string `json:"apiSchemaDTO,omitempty" xml:"apiSchemaDTO,omitempty"`
	GroupVersionExtraInfo *string `json:"groupVersionExtraInfo,omitempty" xml:"groupVersionExtraInfo,omitempty"`
	NamespaceExternalId   *string `json:"namespaceExternalId,omitempty" xml:"namespaceExternalId,omitempty"`
}

func (s PreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s PreCheckRequest) GoString() string {
	return s.String()
}

func (s *PreCheckRequest) SetApiSchemaDTO(v string) *PreCheckRequest {
	s.ApiSchemaDTO = &v
	return s
}

func (s *PreCheckRequest) SetGroupVersionExtraInfo(v string) *PreCheckRequest {
	s.GroupVersionExtraInfo = &v
	return s
}

func (s *PreCheckRequest) SetNamespaceExternalId(v string) *PreCheckRequest {
	s.NamespaceExternalId = &v
	return s
}

type PreCheckResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s PreCheckResponse) GoString() string {
	return s.String()
}

func (s *PreCheckResponse) SetHeaders(v map[string]*string) *PreCheckResponse {
	s.Headers = v
	return s
}

func (s *PreCheckResponse) SetStatusCode(v int32) *PreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *PreCheckResponse) SetBody(v string) *PreCheckResponse {
	s.Body = &v
	return s
}

type PublishSdkRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s PublishSdkRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishSdkRequest) GoString() string {
	return s.String()
}

func (s *PublishSdkRequest) SetTaskId(v string) *PublishSdkRequest {
	s.TaskId = &v
	return s
}

type PublishSdkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishSdkResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishSdkResponse) GoString() string {
	return s.String()
}

func (s *PublishSdkResponse) SetHeaders(v map[string]*string) *PublishSdkResponse {
	s.Headers = v
	return s
}

func (s *PublishSdkResponse) SetStatusCode(v int32) *PublishSdkResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishSdkResponse) SetBody(v string) *PublishSdkResponse {
	s.Body = &v
	return s
}

type SerializeApiRequest struct {
	ApiSchemaDTO *string `json:"apiSchemaDTO,omitempty" xml:"apiSchemaDTO,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SerializeApiRequest) String() string {
	return tea.Prettify(s)
}

func (s SerializeApiRequest) GoString() string {
	return s.String()
}

func (s *SerializeApiRequest) SetApiSchemaDTO(v string) *SerializeApiRequest {
	s.ApiSchemaDTO = &v
	return s
}

func (s *SerializeApiRequest) SetType(v string) *SerializeApiRequest {
	s.Type = &v
	return s
}

type SerializeApiResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SerializeApiResponse) String() string {
	return tea.Prettify(s)
}

func (s SerializeApiResponse) GoString() string {
	return s.String()
}

func (s *SerializeApiResponse) SetHeaders(v map[string]*string) *SerializeApiResponse {
	s.Headers = v
	return s
}

func (s *SerializeApiResponse) SetStatusCode(v int32) *SerializeApiResponse {
	s.StatusCode = &v
	return s
}

func (s *SerializeApiResponse) SetBody(v string) *SerializeApiResponse {
	s.Body = &v
	return s
}

type UpdateAndReleaseApiRequest struct {
	UpdateApiCmd *string `json:"updateApiCmd,omitempty" xml:"updateApiCmd,omitempty"`
}

func (s UpdateAndReleaseApiRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndReleaseApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndReleaseApiRequest) SetUpdateApiCmd(v string) *UpdateAndReleaseApiRequest {
	s.UpdateApiCmd = &v
	return s
}

type UpdateAndReleaseApiResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAndReleaseApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndReleaseApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateAndReleaseApiResponse) SetHeaders(v map[string]*string) *UpdateAndReleaseApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateAndReleaseApiResponse) SetStatusCode(v int32) *UpdateAndReleaseApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAndReleaseApiResponse) SetBody(v string) *UpdateAndReleaseApiResponse {
	s.Body = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("nbf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ApiTestWithOptions(request *ApiTestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApiTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TestCmd)) {
		query["testCmd"] = request.TestCmd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApiTest"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/apiTest"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &ApiTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApiTest(request *ApiTestRequest) (_result *ApiTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApiTestResponse{}
	_body, _err := client.ApiTestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BuildSdkWithOptions(request *BuildSdkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BuildSdkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildCmd)) {
		query["buildCmd"] = request.BuildCmd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BuildSdk"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/buildSdk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &BuildSdkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BuildSdk(request *BuildSdkRequest) (_result *BuildSdkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildSdkResponse{}
	_body, _err := client.BuildSdkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAndReleaseApiWithOptions(request *CreateAndReleaseApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAndReleaseApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatApiCmd)) {
		query["creatApiCmd"] = request.CreatApiCmd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAndReleaseApi"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/createAndReleaseApi"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &CreateAndReleaseApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAndReleaseApi(request *CreateAndReleaseApiRequest) (_result *CreateAndReleaseApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAndReleaseApiResponse{}
	_body, _err := client.CreateAndReleaseApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSdkVersionWithOptions(request *CreateSdkVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSdkVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateSdkCmd)) {
		query["createSdkCmd"] = request.CreateSdkCmd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSdkVersion"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/createSdkVersion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &CreateSdkVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSdkVersion(request *CreateSdkVersionRequest) (_result *CreateSdkVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSdkVersionResponse{}
	_body, _err := client.CreateSdkVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApiWithOptions(request *DeleteApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiExternalId)) {
		query["apiExternalId"] = request.ApiExternalId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApi"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/deleteApi"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &DeleteApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApi(request *DeleteApiRequest) (_result *DeleteApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApiResponse{}
	_body, _err := client.DeleteApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResultWithOptions(request *GetResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResultResponse, _err error) {
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
		Action:      tea.String("GetResult"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/getResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResult(request *GetResultRequest) (_result *GetResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResultResponse{}
	_body, _err := client.GetResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenApiGenericProxyWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenApiGenericProxyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("OpenApiGenericProxy"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/openApiGenericProxy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenApiGenericProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenApiGenericProxy() (_result *OpenApiGenericProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenApiGenericProxyResponse{}
	_body, _err := client.OpenApiGenericProxyWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreCheckWithOptions(request *PreCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSchemaDTO)) {
		query["apiSchemaDTO"] = request.ApiSchemaDTO
	}

	if !tea.BoolValue(util.IsUnset(request.GroupVersionExtraInfo)) {
		query["groupVersionExtraInfo"] = request.GroupVersionExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceExternalId)) {
		query["namespaceExternalId"] = request.NamespaceExternalId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreCheck"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/preCheck"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &PreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreCheck(request *PreCheckRequest) (_result *PreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreCheckResponse{}
	_body, _err := client.PreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishSdkWithOptions(request *PublishSdkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishSdkResponse, _err error) {
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
		Action:      tea.String("PublishSdk"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/publishSdk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &PublishSdkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishSdk(request *PublishSdkRequest) (_result *PublishSdkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishSdkResponse{}
	_body, _err := client.PublishSdkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SerializeApiWithOptions(request *SerializeApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SerializeApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSchemaDTO)) {
		query["apiSchemaDTO"] = request.ApiSchemaDTO
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SerializeApi"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/serializeApi"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &SerializeApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SerializeApi(request *SerializeApiRequest) (_result *SerializeApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SerializeApiResponse{}
	_body, _err := client.SerializeApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAndReleaseApiWithOptions(request *UpdateAndReleaseApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAndReleaseApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateApiCmd)) {
		query["updateApiCmd"] = request.UpdateApiCmd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAndReleaseApi"),
		Version:     tea.String("2021-11-10_21-31-25-86"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/nbf_gateway_inner/1_0_0/updateAndReleaseApi"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &UpdateAndReleaseApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAndReleaseApi(request *UpdateAndReleaseApiRequest) (_result *UpdateAndReleaseApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAndReleaseApiResponse{}
	_body, _err := client.UpdateAndReleaseApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
