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

type CreateAsyncPredictRequest struct {
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DetailTag      *string `json:"DetailTag,omitempty" xml:"DetailTag,omitempty"`
	FetchContent   *string `json:"FetchContent,omitempty" xml:"FetchContent,omitempty"`
	FileContent    *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	FileType       *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileUrl        *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ModelId        *int32  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelVersion   *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	TopK           *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s CreateAsyncPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictRequest) SetContent(v string) *CreateAsyncPredictRequest {
	s.Content = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetDetailTag(v string) *CreateAsyncPredictRequest {
	s.DetailTag = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFetchContent(v string) *CreateAsyncPredictRequest {
	s.FetchContent = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileContent(v string) *CreateAsyncPredictRequest {
	s.FileContent = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileType(v string) *CreateAsyncPredictRequest {
	s.FileType = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileUrl(v string) *CreateAsyncPredictRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetModelId(v int32) *CreateAsyncPredictRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetModelVersion(v string) *CreateAsyncPredictRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetServiceName(v string) *CreateAsyncPredictRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetServiceVersion(v string) *CreateAsyncPredictRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetTopK(v int32) *CreateAsyncPredictRequest {
	s.TopK = &v
	return s
}

type CreateAsyncPredictResponseBody struct {
	AsyncPredictId *int64  `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictResponseBody) SetAsyncPredictId(v int64) *CreateAsyncPredictResponseBody {
	s.AsyncPredictId = &v
	return s
}

func (s *CreateAsyncPredictResponseBody) SetRequestId(v string) *CreateAsyncPredictResponseBody {
	s.RequestId = &v
	return s
}

type CreateAsyncPredictResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAsyncPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsyncPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictResponse) SetHeaders(v map[string]*string) *CreateAsyncPredictResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncPredictResponse) SetStatusCode(v int32) *CreateAsyncPredictResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsyncPredictResponse) SetBody(v *CreateAsyncPredictResponseBody) *CreateAsyncPredictResponse {
	s.Body = v
	return s
}

type GetAsyncPredictRequest struct {
	AsyncPredictId *int32 `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
}

func (s GetAsyncPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictRequest) SetAsyncPredictId(v int32) *GetAsyncPredictRequest {
	s.AsyncPredictId = &v
	return s
}

type GetAsyncPredictResponseBody struct {
	AsyncPredictId *int32  `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictResponseBody) SetAsyncPredictId(v int32) *GetAsyncPredictResponseBody {
	s.AsyncPredictId = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetContent(v string) *GetAsyncPredictResponseBody {
	s.Content = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetRequestId(v string) *GetAsyncPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetStatus(v int32) *GetAsyncPredictResponseBody {
	s.Status = &v
	return s
}

type GetAsyncPredictResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictResponse) SetHeaders(v map[string]*string) *GetAsyncPredictResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncPredictResponse) SetStatusCode(v int32) *GetAsyncPredictResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncPredictResponse) SetBody(v *GetAsyncPredictResponseBody) *GetAsyncPredictResponse {
	s.Body = v
	return s
}

type GetPredictResultRequest struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DetailTag    *string `json:"DetailTag,omitempty" xml:"DetailTag,omitempty"`
	ModelId      *int32  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	TopK         *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s GetPredictResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultRequest) GoString() string {
	return s.String()
}

func (s *GetPredictResultRequest) SetContent(v string) *GetPredictResultRequest {
	s.Content = &v
	return s
}

func (s *GetPredictResultRequest) SetDetailTag(v string) *GetPredictResultRequest {
	s.DetailTag = &v
	return s
}

func (s *GetPredictResultRequest) SetModelId(v int32) *GetPredictResultRequest {
	s.ModelId = &v
	return s
}

func (s *GetPredictResultRequest) SetModelVersion(v string) *GetPredictResultRequest {
	s.ModelVersion = &v
	return s
}

func (s *GetPredictResultRequest) SetTopK(v int32) *GetPredictResultRequest {
	s.TopK = &v
	return s
}

type GetPredictResultResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPredictResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponseBody) SetContent(v string) *GetPredictResultResponseBody {
	s.Content = &v
	return s
}

func (s *GetPredictResultResponseBody) SetRequestId(v string) *GetPredictResultResponseBody {
	s.RequestId = &v
	return s
}

type GetPredictResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPredictResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPredictResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponse) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponse) SetHeaders(v map[string]*string) *GetPredictResultResponse {
	s.Headers = v
	return s
}

func (s *GetPredictResultResponse) SetStatusCode(v int32) *GetPredictResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPredictResultResponse) SetBody(v *GetPredictResultResponseBody) *GetPredictResultResponse {
	s.Body = v
	return s
}

type RunPreTrainServiceRequest struct {
	PredictContent *string `json:"PredictContent,omitempty" xml:"PredictContent,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s RunPreTrainServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceRequest) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceRequest) SetPredictContent(v string) *RunPreTrainServiceRequest {
	s.PredictContent = &v
	return s
}

func (s *RunPreTrainServiceRequest) SetServiceName(v string) *RunPreTrainServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *RunPreTrainServiceRequest) SetServiceVersion(v string) *RunPreTrainServiceRequest {
	s.ServiceVersion = &v
	return s
}

type RunPreTrainServiceResponseBody struct {
	PredictResult *string `json:"PredictResult,omitempty" xml:"PredictResult,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunPreTrainServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceResponseBody) SetPredictResult(v string) *RunPreTrainServiceResponseBody {
	s.PredictResult = &v
	return s
}

func (s *RunPreTrainServiceResponseBody) SetRequestId(v string) *RunPreTrainServiceResponseBody {
	s.RequestId = &v
	return s
}

type RunPreTrainServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunPreTrainServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunPreTrainServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceResponse) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceResponse) SetHeaders(v map[string]*string) *RunPreTrainServiceResponse {
	s.Headers = v
	return s
}

func (s *RunPreTrainServiceResponse) SetStatusCode(v int32) *RunPreTrainServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPreTrainServiceResponse) SetBody(v *RunPreTrainServiceResponseBody) *RunPreTrainServiceResponse {
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

func (client *Client) CreateAsyncPredictWithOptions(request *CreateAsyncPredictRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DetailTag)) {
		body["DetailTag"] = request.DetailTag
	}

	if !tea.BoolValue(util.IsUnset(request.FetchContent)) {
		body["FetchContent"] = request.FetchContent
	}

	if !tea.BoolValue(util.IsUnset(request.FileContent)) {
		body["FileContent"] = request.FileContent
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		body["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["TopK"] = request.TopK
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAsyncPredict"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAsyncPredictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncPredict(request *CreateAsyncPredictRequest) (_result *CreateAsyncPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncPredictResponse{}
	_body, _err := client.CreateAsyncPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncPredictWithOptions(request *GetAsyncPredictRequest, runtime *util.RuntimeOptions) (_result *GetAsyncPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncPredict"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncPredictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncPredict(request *GetAsyncPredictRequest) (_result *GetAsyncPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncPredictResponse{}
	_body, _err := client.GetAsyncPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPredictResultWithOptions(request *GetPredictResultRequest, runtime *util.RuntimeOptions) (_result *GetPredictResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DetailTag)) {
		body["DetailTag"] = request.DetailTag
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["TopK"] = request.TopK
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPredictResult"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPredictResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPredictResult(request *GetPredictResultRequest) (_result *GetPredictResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPredictResultResponse{}
	_body, _err := client.GetPredictResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunPreTrainServiceWithOptions(request *RunPreTrainServiceRequest, runtime *util.RuntimeOptions) (_result *RunPreTrainServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PredictContent)) {
		body["PredictContent"] = request.PredictContent
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		body["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunPreTrainService"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunPreTrainServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunPreTrainService(request *RunPreTrainServiceRequest) (_result *RunPreTrainServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunPreTrainServiceResponse{}
	_body, _err := client.RunPreTrainServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
