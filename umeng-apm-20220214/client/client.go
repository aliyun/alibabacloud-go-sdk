// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetStatTrendRequest struct {
	AppVersion   *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	EndDate      *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	StartDate    *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Type         *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetStatTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendRequest) GoString() string {
	return s.String()
}

func (s *GetStatTrendRequest) SetAppVersion(v string) *GetStatTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetStatTrendRequest) SetDataSourceId(v string) *GetStatTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetStatTrendRequest) SetEndDate(v string) *GetStatTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetStatTrendRequest) SetStartDate(v string) *GetStatTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetStatTrendRequest) SetType(v int32) *GetStatTrendRequest {
	s.Type = &v
	return s
}

type GetStatTrendResponseBody struct {
	Code    *int64                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetStatTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg     *string                         `json:"msg,omitempty" xml:"msg,omitempty"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetStatTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponseBody) SetCode(v int64) *GetStatTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetStatTrendResponseBody) SetData(v []*GetStatTrendResponseBodyData) *GetStatTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetStatTrendResponseBody) SetMsg(v string) *GetStatTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetStatTrendResponseBody) SetSuccess(v bool) *GetStatTrendResponseBody {
	s.Success = &v
	return s
}

type GetStatTrendResponseBodyData struct {
	AffectedUserCount *int64  `json:"affectedUserCount,omitempty" xml:"affectedUserCount,omitempty"`
	AffectedUserRate  *int64  `json:"affectedUserRate,omitempty" xml:"affectedUserRate,omitempty"`
	ErrorCount        *int64  `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	ErrorRate         *int64  `json:"errorRate,omitempty" xml:"errorRate,omitempty"`
	TimePoint         *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetStatTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponseBodyData) SetAffectedUserCount(v int64) *GetStatTrendResponseBodyData {
	s.AffectedUserCount = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetAffectedUserRate(v int64) *GetStatTrendResponseBodyData {
	s.AffectedUserRate = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetErrorCount(v int64) *GetStatTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetErrorRate(v int64) *GetStatTrendResponseBodyData {
	s.ErrorRate = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetTimePoint(v string) *GetStatTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetStatTrendResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStatTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStatTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponse) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponse) SetHeaders(v map[string]*string) *GetStatTrendResponse {
	s.Headers = v
	return s
}

func (s *GetStatTrendResponse) SetStatusCode(v int32) *GetStatTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatTrendResponse) SetBody(v *GetStatTrendResponseBody) *GetStatTrendResponse {
	s.Body = v
	return s
}

type GetSymUploadParamRequest struct {
	AppVersion   *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	FileName     *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType     *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s GetSymUploadParamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamRequest) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamRequest) SetAppVersion(v string) *GetSymUploadParamRequest {
	s.AppVersion = &v
	return s
}

func (s *GetSymUploadParamRequest) SetDataSourceId(v string) *GetSymUploadParamRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetSymUploadParamRequest) SetFileName(v string) *GetSymUploadParamRequest {
	s.FileName = &v
	return s
}

func (s *GetSymUploadParamRequest) SetFileType(v int32) *GetSymUploadParamRequest {
	s.FileType = &v
	return s
}

type GetSymUploadParamResponseBody struct {
	Code    *int64                             `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetSymUploadParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Msg     *string                            `json:"msg,omitempty" xml:"msg,omitempty"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetSymUploadParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponseBody) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponseBody) SetCode(v int64) *GetSymUploadParamResponseBody {
	s.Code = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetData(v *GetSymUploadParamResponseBodyData) *GetSymUploadParamResponseBody {
	s.Data = v
	return s
}

func (s *GetSymUploadParamResponseBody) SetMsg(v string) *GetSymUploadParamResponseBody {
	s.Msg = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetSuccess(v bool) *GetSymUploadParamResponseBody {
	s.Success = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetTraceId(v string) *GetSymUploadParamResponseBody {
	s.TraceId = &v
	return s
}

type GetSymUploadParamResponseBodyData struct {
	AccessKeyId   *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	Callback      *string `json:"callback,omitempty" xml:"callback,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
	Policy        *string `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature     *string `json:"signature,omitempty" xml:"signature,omitempty"`
	UploadAddress *string `json:"uploadAddress,omitempty" xml:"uploadAddress,omitempty"`
}

func (s GetSymUploadParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponseBodyData) SetAccessKeyId(v string) *GetSymUploadParamResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetCallback(v string) *GetSymUploadParamResponseBodyData {
	s.Callback = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetKey(v string) *GetSymUploadParamResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetPolicy(v string) *GetSymUploadParamResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetSignature(v string) *GetSymUploadParamResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetUploadAddress(v string) *GetSymUploadParamResponseBodyData {
	s.UploadAddress = &v
	return s
}

type GetSymUploadParamResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSymUploadParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSymUploadParamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponse) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponse) SetHeaders(v map[string]*string) *GetSymUploadParamResponse {
	s.Headers = v
	return s
}

func (s *GetSymUploadParamResponse) SetStatusCode(v int32) *GetSymUploadParamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSymUploadParamResponse) SetBody(v *GetSymUploadParamResponseBody) *GetSymUploadParamResponse {
	s.Body = v
	return s
}

type GetTodayStatTrendRequest struct {
	AppVersion   *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	Type         *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTodayStatTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendRequest) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendRequest) SetAppVersion(v string) *GetTodayStatTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetTodayStatTrendRequest) SetDataSourceId(v string) *GetTodayStatTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetTodayStatTrendRequest) SetType(v int32) *GetTodayStatTrendRequest {
	s.Type = &v
	return s
}

type GetTodayStatTrendResponseBody struct {
	Code    *int64                               `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetTodayStatTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg     *string                              `json:"msg,omitempty" xml:"msg,omitempty"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTodayStatTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponseBody) SetCode(v int64) *GetTodayStatTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetData(v []*GetTodayStatTrendResponseBodyData) *GetTodayStatTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetMsg(v string) *GetTodayStatTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetSuccess(v bool) *GetTodayStatTrendResponseBody {
	s.Success = &v
	return s
}

type GetTodayStatTrendResponseBodyData struct {
	AffectedUserCount *int64  `json:"affectedUserCount,omitempty" xml:"affectedUserCount,omitempty"`
	AffectedUserRate  *int64  `json:"affectedUserRate,omitempty" xml:"affectedUserRate,omitempty"`
	ErrorCount        *int64  `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	ErrorRate         *int64  `json:"errorRate,omitempty" xml:"errorRate,omitempty"`
	TimePoint         *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetTodayStatTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponseBodyData) SetAffectedUserCount(v int64) *GetTodayStatTrendResponseBodyData {
	s.AffectedUserCount = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetAffectedUserRate(v int64) *GetTodayStatTrendResponseBodyData {
	s.AffectedUserRate = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetErrorCount(v int64) *GetTodayStatTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetErrorRate(v int64) *GetTodayStatTrendResponseBodyData {
	s.ErrorRate = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetTimePoint(v string) *GetTodayStatTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetTodayStatTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTodayStatTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTodayStatTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponse) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponse) SetHeaders(v map[string]*string) *GetTodayStatTrendResponse {
	s.Headers = v
	return s
}

func (s *GetTodayStatTrendResponse) SetStatusCode(v int32) *GetTodayStatTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodayStatTrendResponse) SetBody(v *GetTodayStatTrendResponseBody) *GetTodayStatTrendResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-apm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetStatTrend(request *GetStatTrendRequest) (_result *GetStatTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetStatTrendResponse{}
	_body, _err := client.GetStatTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStatTrendWithOptions(request *GetStatTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetStatTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStatTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getStatTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStatTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSymUploadParam(request *GetSymUploadParamRequest) (_result *GetSymUploadParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSymUploadParamResponse{}
	_body, _err := client.GetSymUploadParamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSymUploadParamWithOptions(request *GetSymUploadParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSymUploadParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["fileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSymUploadParam"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/getSymUploadParam"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSymUploadParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTodayStatTrend(request *GetTodayStatTrendRequest) (_result *GetTodayStatTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTodayStatTrendResponse{}
	_body, _err := client.GetTodayStatTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTodayStatTrendWithOptions(request *GetTodayStatTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTodayStatTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTodayStatTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getTodayStatTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTodayStatTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
