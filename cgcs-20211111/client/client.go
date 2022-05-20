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

type CreateAppSessionRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端ip
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	CustomUserId   *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	EnablePostpaid *bool   `json:"EnablePostpaid,omitempty" xml:"EnablePostpaid,omitempty"`
	// 启动参数
	StartParameters []*CreateAppSessionRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Timeout    *int64                               `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequest) SetAppId(v string) *CreateAppSessionRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionRequest) SetAppVersion(v string) *CreateAppSessionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionRequest) SetClientIp(v string) *CreateAppSessionRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomSessionId(v string) *CreateAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomUserId(v string) *CreateAppSessionRequest {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionRequest) SetEnablePostpaid(v bool) *CreateAppSessionRequest {
	s.EnablePostpaid = &v
	return s
}

func (s *CreateAppSessionRequest) SetStartParameters(v []*CreateAppSessionRequestStartParameters) *CreateAppSessionRequest {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionRequest) SetSystemInfo(v []*CreateAppSessionRequestSystemInfo) *CreateAppSessionRequest {
	s.SystemInfo = v
	return s
}

func (s *CreateAppSessionRequest) SetTimeout(v int64) *CreateAppSessionRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionRequestStartParameters struct {
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestStartParameters) SetKey(v string) *CreateAppSessionRequestStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestStartParameters) SetValue(v string) *CreateAppSessionRequestStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionRequestSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestSystemInfo) SetKey(v string) *CreateAppSessionRequestSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestSystemInfo) SetValue(v string) *CreateAppSessionRequestSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponseBody) SetAppId(v string) *CreateAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetAppVersion(v string) *CreateAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetCustomSessionId(v string) *CreateAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetPlatformSessionId(v string) *CreateAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetRequestId(v string) *CreateAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponse) SetHeaders(v map[string]*string) *CreateAppSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionResponse) SetStatusCode(v int32) *CreateAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSessionResponse) SetBody(v *CreateAppSessionResponseBody) *CreateAppSessionResponse {
	s.Body = v
	return s
}

type GetAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s GetAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAppSessionRequest) SetCustomSessionId(v string) *GetAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionRequest) SetPlatformSessionId(v string) *GetAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

type GetAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBody) SetAppId(v string) *GetAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetAppVersion(v string) *GetAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetAppSessionResponseBody) SetCustomSessionId(v string) *GetAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetPlatformSessionId(v string) *GetAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetRequestId(v string) *GetAppSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetStatus(v string) *GetAppSessionResponseBody {
	s.Status = &v
	return s
}

type GetAppSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponse) SetHeaders(v map[string]*string) *GetAppSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAppSessionResponse) SetStatusCode(v int32) *GetAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSessionResponse) SetBody(v *GetAppSessionResponseBody) *GetAppSessionResponse {
	s.Body = v
	return s
}

type ListAppSessionsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 自定义会话id
	CustomSessionIds []*string `json:"CustomSessionIds,omitempty" xml:"CustomSessionIds,omitempty" type:"Repeated"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 自定义用户id
	PlatformSessionIds []*string `json:"PlatformSessionIds,omitempty" xml:"PlatformSessionIds,omitempty" type:"Repeated"`
}

func (s ListAppSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSessionsRequest) SetAppId(v string) *ListAppSessionsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsRequest) SetCustomSessionIds(v []*string) *ListAppSessionsRequest {
	s.CustomSessionIds = v
	return s
}

func (s *ListAppSessionsRequest) SetPageNumber(v int32) *ListAppSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsRequest) SetPageSize(v int32) *ListAppSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsRequest) SetPlatformSessionIds(v []*string) *ListAppSessionsRequest {
	s.PlatformSessionIds = v
	return s
}

type ListAppSessionsResponseBody struct {
	AppSessions []*ListAppSessionsResponseBodyAppSessions `json:"AppSessions,omitempty" xml:"AppSessions,omitempty" type:"Repeated"`
	PageNumber  *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBody) SetAppSessions(v []*ListAppSessionsResponseBodyAppSessions) *ListAppSessionsResponseBody {
	s.AppSessions = v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageNumber(v int32) *ListAppSessionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageSize(v int32) *ListAppSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetRequestId(v string) *ListAppSessionsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppSessionsResponseBodyAppSessions struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppSessionsResponseBodyAppSessions) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBodyAppSessions) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppVersion(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppVersion = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetCustomSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.CustomSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetPlatformSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.PlatformSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetStatus(v string) *ListAppSessionsResponseBodyAppSessions {
	s.Status = &v
	return s
}

type ListAppSessionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponse) SetHeaders(v map[string]*string) *ListAppSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppSessionsResponse) SetStatusCode(v int32) *ListAppSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppSessionsResponse) SetBody(v *ListAppSessionsResponseBody) *ListAppSessionsResponse {
	s.Body = v
	return s
}

type StopAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s StopAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionRequest) SetCustomSessionId(v string) *StopAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionRequest) SetPlatformSessionId(v string) *StopAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

type StopAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponseBody) SetAppId(v string) *StopAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetAppVersion(v string) *StopAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *StopAppSessionResponseBody) SetCustomSessionId(v string) *StopAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetPlatformSessionId(v string) *StopAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetRequestId(v string) *StopAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type StopAppSessionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponse) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponse) SetHeaders(v map[string]*string) *StopAppSessionResponse {
	s.Headers = v
	return s
}

func (s *StopAppSessionResponse) SetStatusCode(v int32) *StopAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppSessionResponse) SetBody(v *StopAppSessionResponseBody) *StopAppSessionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cgcs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAppSessionWithOptions(request *CreateAppSessionRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePostpaid)) {
		query["EnablePostpaid"] = request.EnablePostpaid
	}

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		query["SystemInfo"] = request.SystemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSession(request *CreateAppSessionRequest) (_result *CreateAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CreateAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppSessionWithOptions(request *GetAppSessionRequest, runtime *util.RuntimeOptions) (_result *GetAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppSession(request *GetAppSessionRequest) (_result *GetAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSessionResponse{}
	_body, _err := client.GetAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppSessionsWithOptions(request *ListAppSessionsRequest, runtime *util.RuntimeOptions) (_result *ListAppSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionIds)) {
		query["CustomSessionIds"] = request.CustomSessionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionIds)) {
		query["PlatformSessionIds"] = request.PlatformSessionIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppSessions"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppSessions(request *ListAppSessionsRequest) (_result *ListAppSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.ListAppSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppSessionWithOptions(request *StopAppSessionRequest, runtime *util.RuntimeOptions) (_result *StopAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppSession(request *StopAppSessionRequest) (_result *StopAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppSessionResponse{}
	_body, _err := client.StopAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
