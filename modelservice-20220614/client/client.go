// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetMonthAmountRequest struct {
	// This parameter is required.
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetMonthAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonthAmountRequest) GoString() string {
	return s.String()
}

func (s *GetMonthAmountRequest) SetSceneType(v string) *GetMonthAmountRequest {
	s.SceneType = &v
	return s
}

type GetMonthAmountResponseBody struct {
	BackgroundAmount *int32  `json:"BackgroundAmount,omitempty" xml:"BackgroundAmount,omitempty"`
	PostpayAmount    *int32  `json:"PostpayAmount,omitempty" xml:"PostpayAmount,omitempty"`
	PrepayAmount     *int32  `json:"PrepayAmount,omitempty" xml:"PrepayAmount,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalAmount      *int32  `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s GetMonthAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonthAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonthAmountResponseBody) SetBackgroundAmount(v int32) *GetMonthAmountResponseBody {
	s.BackgroundAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetPostpayAmount(v int32) *GetMonthAmountResponseBody {
	s.PostpayAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetPrepayAmount(v int32) *GetMonthAmountResponseBody {
	s.PrepayAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetRequestId(v string) *GetMonthAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetTotalAmount(v int32) *GetMonthAmountResponseBody {
	s.TotalAmount = &v
	return s
}

type GetMonthAmountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonthAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonthAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonthAmountResponse) GoString() string {
	return s.String()
}

func (s *GetMonthAmountResponse) SetHeaders(v map[string]*string) *GetMonthAmountResponse {
	s.Headers = v
	return s
}

func (s *GetMonthAmountResponse) SetStatusCode(v int32) *GetMonthAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonthAmountResponse) SetBody(v *GetMonthAmountResponseBody) *GetMonthAmountResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// example:
	//
	// image_label
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetSceneType(v string) *GetUserRequest {
	s.SceneType = &v
	return s
}

type GetUserResponseBody struct {
	// example:
	//
	// 20200425******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ai-service.******.alicontainer.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 897658D5-1FB8-5CFA-A030-727CCAE59EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Njk0Njk******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetAppId(v string) *GetUserResponseBody {
	s.AppId = &v
	return s
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetHost(v string) *GetUserResponseBody {
	s.Host = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetToken(v string) *GetUserResponseBody {
	s.Token = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListDayAmountRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// sales_pick
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDayAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDayAmountRequest) GoString() string {
	return s.String()
}

func (s *ListDayAmountRequest) SetEndTime(v string) *ListDayAmountRequest {
	s.EndTime = &v
	return s
}

func (s *ListDayAmountRequest) SetSceneType(v string) *ListDayAmountRequest {
	s.SceneType = &v
	return s
}

func (s *ListDayAmountRequest) SetStartTime(v string) *ListDayAmountRequest {
	s.StartTime = &v
	return s
}

type ListDayAmountResponseBody struct {
	DayAmounts []*ListDayAmountResponseBodyDayAmounts `json:"DayAmounts,omitempty" xml:"DayAmounts,omitempty" type:"Repeated"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDayAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDayAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponseBody) SetDayAmounts(v []*ListDayAmountResponseBodyDayAmounts) *ListDayAmountResponseBody {
	s.DayAmounts = v
	return s
}

func (s *ListDayAmountResponseBody) SetRequestId(v string) *ListDayAmountResponseBody {
	s.RequestId = &v
	return s
}

type ListDayAmountResponseBodyDayAmounts struct {
	Amount *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Date   *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s ListDayAmountResponseBodyDayAmounts) String() string {
	return tea.Prettify(s)
}

func (s ListDayAmountResponseBodyDayAmounts) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponseBodyDayAmounts) SetAmount(v int32) *ListDayAmountResponseBodyDayAmounts {
	s.Amount = &v
	return s
}

func (s *ListDayAmountResponseBodyDayAmounts) SetDate(v string) *ListDayAmountResponseBodyDayAmounts {
	s.Date = &v
	return s
}

type ListDayAmountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDayAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDayAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDayAmountResponse) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponse) SetHeaders(v map[string]*string) *ListDayAmountResponse {
	s.Headers = v
	return s
}

func (s *ListDayAmountResponse) SetStatusCode(v int32) *ListDayAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDayAmountResponse) SetBody(v *ListDayAmountResponseBody) *ListDayAmountResponse {
	s.Body = v
	return s
}

type ListRechargeBillsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// sales_pick
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s ListRechargeBillsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRechargeBillsRequest) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsRequest) SetPageNumber(v int32) *ListRechargeBillsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRechargeBillsRequest) SetPageSize(v int32) *ListRechargeBillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRechargeBillsRequest) SetSceneType(v string) *ListRechargeBillsRequest {
	s.SceneType = &v
	return s
}

type ListRechargeBillsResponseBody struct {
	InstanseInfos []*ListRechargeBillsResponseBodyInstanseInfos `json:"InstanseInfos,omitempty" xml:"InstanseInfos,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResidueAmount *int32                                        `json:"ResidueAmount,omitempty" xml:"ResidueAmount,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRechargeBillsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRechargeBillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponseBody) SetInstanseInfos(v []*ListRechargeBillsResponseBodyInstanseInfos) *ListRechargeBillsResponseBody {
	s.InstanseInfos = v
	return s
}

func (s *ListRechargeBillsResponseBody) SetRequestId(v string) *ListRechargeBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRechargeBillsResponseBody) SetResidueAmount(v int32) *ListRechargeBillsResponseBody {
	s.ResidueAmount = &v
	return s
}

func (s *ListRechargeBillsResponseBody) SetTotalCount(v int32) *ListRechargeBillsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRechargeBillsResponseBodyInstanseInfos struct {
	CurrTimes  *int32  `json:"CurrTimes,omitempty" xml:"CurrTimes,omitempty"`
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	InitTimes  *int32  `json:"InitTimes,omitempty" xml:"InitTimes,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListRechargeBillsResponseBodyInstanseInfos) String() string {
	return tea.Prettify(s)
}

func (s ListRechargeBillsResponseBodyInstanseInfos) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetCurrTimes(v int32) *ListRechargeBillsResponseBodyInstanseInfos {
	s.CurrTimes = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetGmtEndTime(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.GmtEndTime = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetInitTimes(v int32) *ListRechargeBillsResponseBodyInstanseInfos {
	s.InitTimes = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetInstanceId(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.InstanceId = &v
	return s
}

func (s *ListRechargeBillsResponseBodyInstanseInfos) SetSource(v string) *ListRechargeBillsResponseBodyInstanseInfos {
	s.Source = &v
	return s
}

type ListRechargeBillsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRechargeBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRechargeBillsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRechargeBillsResponse) GoString() string {
	return s.String()
}

func (s *ListRechargeBillsResponse) SetHeaders(v map[string]*string) *ListRechargeBillsResponse {
	s.Headers = v
	return s
}

func (s *ListRechargeBillsResponse) SetStatusCode(v int32) *ListRechargeBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRechargeBillsResponse) SetBody(v *ListRechargeBillsResponseBody) *ListRechargeBillsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("modelservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取当月的使用量
//
// @param request - GetMonthAmountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonthAmountResponse
func (client *Client) GetMonthAmountWithOptions(request *GetMonthAmountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMonthAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMonthAmount"),
		Version:     tea.String("2022-06-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/month/amount"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMonthAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当月的使用量
//
// @param request - GetMonthAmountRequest
//
// @return GetMonthAmountResponse
func (client *Client) GetMonthAmount(request *GetMonthAmountRequest) (_result *GetMonthAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMonthAmountResponse{}
	_body, _err := client.GetMonthAmountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取user相关的appid, token等信息
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-06-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/user/info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取user相关的appid, token等信息
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 每天的调用量列表
//
// @param request - ListDayAmountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDayAmountResponse
func (client *Client) ListDayAmountWithOptions(request *ListDayAmountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDayAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDayAmount"),
		Version:     tea.String("2022-06-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/day/amount"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDayAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 每天的调用量列表
//
// @param request - ListDayAmountRequest
//
// @return ListDayAmountResponse
func (client *Client) ListDayAmount(request *ListDayAmountRequest) (_result *ListDayAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDayAmountResponse{}
	_body, _err := client.ListDayAmountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户充值列表
//
// @param request - ListRechargeBillsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRechargeBillsResponse
func (client *Client) ListRechargeBillsWithOptions(request *ListRechargeBillsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRechargeBillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRechargeBills"),
		Version:     tea.String("2022-06-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/rechargebills"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRechargeBillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户充值列表
//
// @param request - ListRechargeBillsRequest
//
// @return ListRechargeBillsResponse
func (client *Client) ListRechargeBills(request *ListRechargeBillsRequest) (_result *ListRechargeBillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRechargeBillsResponse{}
	_body, _err := client.ListRechargeBillsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
