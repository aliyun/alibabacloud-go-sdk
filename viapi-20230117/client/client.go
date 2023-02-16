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

type CancelWaitingAsyncJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelWaitingAsyncJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelWaitingAsyncJobRequest) GoString() string {
	return s.String()
}

func (s *CancelWaitingAsyncJobRequest) SetJobId(v string) *CancelWaitingAsyncJobRequest {
	s.JobId = &v
	return s
}

type CancelWaitingAsyncJobResponseBody struct {
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelWaitingAsyncJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelWaitingAsyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelWaitingAsyncJobResponseBody) SetHttpCode(v string) *CancelWaitingAsyncJobResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CancelWaitingAsyncJobResponseBody) SetMessage(v string) *CancelWaitingAsyncJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelWaitingAsyncJobResponseBody) SetRequestId(v string) *CancelWaitingAsyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelWaitingAsyncJobResponseBody) SetSuccess(v bool) *CancelWaitingAsyncJobResponseBody {
	s.Success = &v
	return s
}

type CancelWaitingAsyncJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelWaitingAsyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelWaitingAsyncJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelWaitingAsyncJobResponse) GoString() string {
	return s.String()
}

func (s *CancelWaitingAsyncJobResponse) SetHeaders(v map[string]*string) *CancelWaitingAsyncJobResponse {
	s.Headers = v
	return s
}

func (s *CancelWaitingAsyncJobResponse) SetStatusCode(v int32) *CancelWaitingAsyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelWaitingAsyncJobResponse) SetBody(v *CancelWaitingAsyncJobResponseBody) *CancelWaitingAsyncJobResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type QueryAsyncJobListRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNum    *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PopApiName *string `json:"PopApiName,omitempty" xml:"PopApiName,omitempty"`
	PopProduct *string `json:"PopProduct,omitempty" xml:"PopProduct,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAsyncJobListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryAsyncJobListRequest) SetEndTime(v string) *QueryAsyncJobListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetJobId(v string) *QueryAsyncJobListRequest {
	s.JobId = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetPageNum(v string) *QueryAsyncJobListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetPageSize(v string) *QueryAsyncJobListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetPopApiName(v string) *QueryAsyncJobListRequest {
	s.PopApiName = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetPopProduct(v string) *QueryAsyncJobListRequest {
	s.PopProduct = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetStartTime(v string) *QueryAsyncJobListRequest {
	s.StartTime = &v
	return s
}

func (s *QueryAsyncJobListRequest) SetStatus(v string) *QueryAsyncJobListRequest {
	s.Status = &v
	return s
}

type QueryAsyncJobListResponseBody struct {
	Data      *QueryAsyncJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode  *string                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAsyncJobListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsyncJobListResponseBody) SetData(v *QueryAsyncJobListResponseBodyData) *QueryAsyncJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAsyncJobListResponseBody) SetHttpCode(v string) *QueryAsyncJobListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryAsyncJobListResponseBody) SetMessage(v string) *QueryAsyncJobListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAsyncJobListResponseBody) SetRequestId(v string) *QueryAsyncJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAsyncJobListResponseBody) SetSuccess(v bool) *QueryAsyncJobListResponseBody {
	s.Success = &v
	return s
}

type QueryAsyncJobListResponseBodyData struct {
	CurrentPage *int32                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result      []*QueryAsyncJobListResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalPage   *int32                                     `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalRecord *int32                                     `json:"TotalRecord,omitempty" xml:"TotalRecord,omitempty"`
}

func (s QueryAsyncJobListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAsyncJobListResponseBodyData) SetCurrentPage(v int32) *QueryAsyncJobListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyData) SetPageSize(v int32) *QueryAsyncJobListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyData) SetResult(v []*QueryAsyncJobListResponseBodyDataResult) *QueryAsyncJobListResponseBodyData {
	s.Result = v
	return s
}

func (s *QueryAsyncJobListResponseBodyData) SetTotalPage(v int32) *QueryAsyncJobListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyData) SetTotalRecord(v int32) *QueryAsyncJobListResponseBodyData {
	s.TotalRecord = &v
	return s
}

type QueryAsyncJobListResponseBodyDataResult struct {
	CallerParentId *string `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PopApiName     *string `json:"PopApiName,omitempty" xml:"PopApiName,omitempty"`
	PopProduct     *string `json:"PopProduct,omitempty" xml:"PopProduct,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryAsyncJobListResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncJobListResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetCallerParentId(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.CallerParentId = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetEndTime(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.EndTime = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetGmtCreate(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetJobId(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.JobId = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetPopApiName(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.PopApiName = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetPopProduct(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.PopProduct = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetStartTime(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.StartTime = &v
	return s
}

func (s *QueryAsyncJobListResponseBodyDataResult) SetStatus(v string) *QueryAsyncJobListResponseBodyDataResult {
	s.Status = &v
	return s
}

type QueryAsyncJobListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAsyncJobListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAsyncJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryAsyncJobListResponse) SetHeaders(v map[string]*string) *QueryAsyncJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryAsyncJobListResponse) SetStatusCode(v int32) *QueryAsyncJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsyncJobListResponse) SetBody(v *QueryAsyncJobListResponseBody) *QueryAsyncJobListResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("viapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelWaitingAsyncJobWithOptions(request *CancelWaitingAsyncJobRequest, runtime *util.RuntimeOptions) (_result *CancelWaitingAsyncJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelWaitingAsyncJob"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelWaitingAsyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelWaitingAsyncJob(request *CancelWaitingAsyncJobRequest) (_result *CancelWaitingAsyncJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelWaitingAsyncJobResponse{}
	_body, _err := client.CancelWaitingAsyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAsyncJobListWithOptions(request *QueryAsyncJobListRequest, runtime *util.RuntimeOptions) (_result *QueryAsyncJobListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PopApiName)) {
		body["PopApiName"] = request.PopApiName
	}

	if !tea.BoolValue(util.IsUnset(request.PopProduct)) {
		body["PopProduct"] = request.PopProduct
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAsyncJobList"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAsyncJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAsyncJobList(request *QueryAsyncJobListRequest) (_result *QueryAsyncJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAsyncJobListResponse{}
	_body, _err := client.QueryAsyncJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
