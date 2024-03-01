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

type ActualDeductResourceCmd struct {
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ActualDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceCmd) SetCost(v int64) *ActualDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ActualDeductResourceCmd) SetExtraInfo(v string) *ActualDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ActualDeductResourceCmd) SetIdempotentId(v string) *ActualDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ActualDeductResourceCmd) SetTaskId(v string) *ActualDeductResourceCmd {
	s.TaskId = &v
	return s
}

type ActualDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ActualDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceResult) SetErrorCode(v string) *ActualDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ActualDeductResourceResult) SetErrorMessage(v string) *ActualDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ActualDeductResourceResult) SetRequestId(v string) *ActualDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ActualDeductResourceResult) SetSuccess(v bool) *ActualDeductResourceResult {
	s.Success = &v
	return s
}

type DirectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	DeductScene  *string `json:"deductScene,omitempty" xml:"deductScene,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s DirectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceCmd) SetAccountId(v string) *DirectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetCost(v int64) *DirectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *DirectDeductResourceCmd) SetDeductScene(v string) *DirectDeductResourceCmd {
	s.DeductScene = &v
	return s
}

func (s *DirectDeductResourceCmd) SetExtraInfo(v string) *DirectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *DirectDeductResourceCmd) SetIdempotentId(v string) *DirectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetResourceType(v int64) *DirectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *DirectDeductResourceCmd) SetSubAccountId(v string) *DirectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetToken(v string) *DirectDeductResourceCmd {
	s.Token = &v
	return s
}

type DirectDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DirectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceResult) SetErrorCode(v string) *DirectDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *DirectDeductResourceResult) SetErrorMessage(v string) *DirectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *DirectDeductResourceResult) SetRequestId(v string) *DirectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *DirectDeductResourceResult) SetSuccess(v bool) *DirectDeductResourceResult {
	s.Success = &v
	return s
}

type ExpectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ExpectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceCmd) SetAccountId(v string) *ExpectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetCost(v int64) *ExpectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetExtraInfo(v string) *ExpectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetIdempotentId(v string) *ExpectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetResourceType(v int64) *ExpectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetSubAccountId(v string) *ExpectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetToken(v string) *ExpectDeductResourceCmd {
	s.Token = &v
	return s
}

type ExpectDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExpectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceResult) SetErrorCode(v string) *ExpectDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ExpectDeductResourceResult) SetErrorMessage(v string) *ExpectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ExpectDeductResourceResult) SetRequestId(v string) *ExpectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ExpectDeductResourceResult) SetSuccess(v bool) *ExpectDeductResourceResult {
	s.Success = &v
	return s
}

func (s *ExpectDeductResourceResult) SetTaskId(v string) *ExpectDeductResourceResult {
	s.TaskId = &v
	return s
}

type ActualDeductResourceRequest struct {
	Body *ActualDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceRequest) SetBody(v *ActualDeductResourceCmd) *ActualDeductResourceRequest {
	s.Body = v
	return s
}

type ActualDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActualDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceResponse) SetHeaders(v map[string]*string) *ActualDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *ActualDeductResourceResponse) SetStatusCode(v int32) *ActualDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActualDeductResourceResponse) SetBody(v *ActualDeductResourceResult) *ActualDeductResourceResponse {
	s.Body = v
	return s
}

type ActualDeductResourcesRequest struct {
	Body *ActualDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *ActualDeductResourcesRequest) SetBody(v *ActualDeductResourceCmd) *ActualDeductResourcesRequest {
	s.Body = v
	return s
}

type ActualDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActualDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *ActualDeductResourcesResponse) SetHeaders(v map[string]*string) *ActualDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *ActualDeductResourcesResponse) SetStatusCode(v int32) *ActualDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActualDeductResourcesResponse) SetBody(v *ActualDeductResourceResult) *ActualDeductResourcesResponse {
	s.Body = v
	return s
}

type DirectDeductResourceRequest struct {
	Body *DirectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceRequest) SetBody(v *DirectDeductResourceCmd) *DirectDeductResourceRequest {
	s.Body = v
	return s
}

type DirectDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DirectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceResponse) SetHeaders(v map[string]*string) *DirectDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *DirectDeductResourceResponse) SetStatusCode(v int32) *DirectDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DirectDeductResourceResponse) SetBody(v *DirectDeductResourceResult) *DirectDeductResourceResponse {
	s.Body = v
	return s
}

type DirectDeductResourcesRequest struct {
	Body *DirectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *DirectDeductResourcesRequest) SetBody(v *DirectDeductResourceCmd) *DirectDeductResourcesRequest {
	s.Body = v
	return s
}

type DirectDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DirectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *DirectDeductResourcesResponse) SetHeaders(v map[string]*string) *DirectDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *DirectDeductResourcesResponse) SetStatusCode(v int32) *DirectDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DirectDeductResourcesResponse) SetBody(v *DirectDeductResourceResult) *DirectDeductResourcesResponse {
	s.Body = v
	return s
}

type ExpectDeductResourceRequest struct {
	Body *ExpectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceRequest) SetBody(v *ExpectDeductResourceCmd) *ExpectDeductResourceRequest {
	s.Body = v
	return s
}

type ExpectDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExpectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceResponse) SetHeaders(v map[string]*string) *ExpectDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *ExpectDeductResourceResponse) SetStatusCode(v int32) *ExpectDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpectDeductResourceResponse) SetBody(v *ExpectDeductResourceResult) *ExpectDeductResourceResponse {
	s.Body = v
	return s
}

type ExpectDeductResourcesRequest struct {
	Body *ExpectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourcesRequest) SetBody(v *ExpectDeductResourceCmd) *ExpectDeductResourcesRequest {
	s.Body = v
	return s
}

type ExpectDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExpectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourcesResponse) SetHeaders(v map[string]*string) *ExpectDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *ExpectDeductResourcesResponse) SetStatusCode(v int32) *ExpectDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpectDeductResourcesResponse) SetBody(v *ExpectDeductResourceResult) *ExpectDeductResourcesResponse {
	s.Body = v
	return s
}

type GetRemainResourceRequest struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s GetRemainResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceRequest) GoString() string {
	return s.String()
}

func (s *GetRemainResourceRequest) SetAccountId(v string) *GetRemainResourceRequest {
	s.AccountId = &v
	return s
}

func (s *GetRemainResourceRequest) SetResourceType(v string) *GetRemainResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetRemainResourceRequest) SetSubAccountId(v string) *GetRemainResourceRequest {
	s.SubAccountId = &v
	return s
}

type GetRemainResourceResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RemainCount  *int32  `json:"remainCount,omitempty" xml:"remainCount,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRemainResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemainResourceResponseBody) SetErrorCode(v string) *GetRemainResourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetErrorMessage(v string) *GetRemainResourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetRemainCount(v int32) *GetRemainResourceResponseBody {
	s.RemainCount = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetRequestId(v string) *GetRemainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetSuccess(v bool) *GetRemainResourceResponseBody {
	s.Success = &v
	return s
}

type GetRemainResourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRemainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRemainResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceResponse) GoString() string {
	return s.String()
}

func (s *GetRemainResourceResponse) SetHeaders(v map[string]*string) *GetRemainResourceResponse {
	s.Headers = v
	return s
}

func (s *GetRemainResourceResponse) SetStatusCode(v int32) *GetRemainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemainResourceResponse) SetBody(v *GetRemainResourceResponseBody) *GetRemainResourceResponse {
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

func (client *Client) ActualDeductResourceWithOptions(request *ActualDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActualDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActualDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/actual-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActualDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActualDeductResource(request *ActualDeductResourceRequest) (_result *ActualDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActualDeductResourceResponse{}
	_body, _err := client.ActualDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActualDeductResourcesWithOptions(request *ActualDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActualDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActualDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/actualDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActualDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActualDeductResources(request *ActualDeductResourcesRequest) (_result *ActualDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActualDeductResourcesResponse{}
	_body, _err := client.ActualDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DirectDeductResourceWithOptions(request *DirectDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DirectDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("DirectDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/direct-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DirectDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DirectDeductResource(request *DirectDeductResourceRequest) (_result *DirectDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DirectDeductResourceResponse{}
	_body, _err := client.DirectDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DirectDeductResourcesWithOptions(request *DirectDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DirectDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("DirectDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/directDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DirectDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DirectDeductResources(request *DirectDeductResourcesRequest) (_result *DirectDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DirectDeductResourcesResponse{}
	_body, _err := client.DirectDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpectDeductResourceWithOptions(request *ExpectDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExpectDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpectDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/expect-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpectDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpectDeductResource(request *ExpectDeductResourceRequest) (_result *ExpectDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpectDeductResourceResponse{}
	_body, _err := client.ExpectDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpectDeductResourcesWithOptions(request *ExpectDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExpectDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpectDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/expectDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpectDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpectDeductResources(request *ExpectDeductResourcesRequest) (_result *ExpectDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpectDeductResourcesResponse{}
	_body, _err := client.ExpectDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRemainResourceWithOptions(request *GetRemainResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRemainResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountId)) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRemainResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/getRemainResource"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRemainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRemainResource(request *GetRemainResourceRequest) (_result *GetRemainResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRemainResourceResponse{}
	_body, _err := client.GetRemainResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
