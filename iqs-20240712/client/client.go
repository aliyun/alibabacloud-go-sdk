// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AgentBaseQuery struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AgentBaseQuery) String() string {
	return tea.Prettify(s)
}

func (s AgentBaseQuery) GoString() string {
	return s.String()
}

func (s *AgentBaseQuery) SetQuery(v string) *AgentBaseQuery {
	s.Query = &v
	return s
}

type CommonAgentQuery struct {
	Query              *string `json:"query,omitempty" xml:"query,omitempty"`
	QuerySceneEnumCode *string `json:"querySceneEnumCode,omitempty" xml:"querySceneEnumCode,omitempty"`
}

func (s CommonAgentQuery) String() string {
	return tea.Prettify(s)
}

func (s CommonAgentQuery) GoString() string {
	return s.String()
}

func (s *CommonAgentQuery) SetQuery(v string) *CommonAgentQuery {
	s.Query = &v
	return s
}

func (s *CommonAgentQuery) SetQuerySceneEnumCode(v string) *CommonAgentQuery {
	s.QuerySceneEnumCode = &v
	return s
}

type QueryResult struct {
	Data []*QueryResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryResult) String() string {
	return tea.Prettify(s)
}

func (s QueryResult) GoString() string {
	return s.String()
}

func (s *QueryResult) SetData(v []*QueryResultData) *QueryResult {
	s.Data = v
	return s
}

type QueryResultData struct {
	Address      *string                  `json:"address,omitempty" xml:"address,omitempty"`
	CityCode     *string                  `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	CityName     *string                  `json:"cityName,omitempty" xml:"cityName,omitempty"`
	DistrictCode *string                  `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	DistrictName *string                  `json:"districtName,omitempty" xml:"districtName,omitempty"`
	Id           *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Images       []*QueryResultDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Latitude     *string                  `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude    *string                  `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Metadata     *QueryResultDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Name         *string                  `json:"name,omitempty" xml:"name,omitempty"`
	ProvinceCode *string                  `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	ProvinceName *string                  `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	TypeCode     *string                  `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	Types        *string                  `json:"types,omitempty" xml:"types,omitempty"`
}

func (s QueryResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryResultData) GoString() string {
	return s.String()
}

func (s *QueryResultData) SetAddress(v string) *QueryResultData {
	s.Address = &v
	return s
}

func (s *QueryResultData) SetCityCode(v string) *QueryResultData {
	s.CityCode = &v
	return s
}

func (s *QueryResultData) SetCityName(v string) *QueryResultData {
	s.CityName = &v
	return s
}

func (s *QueryResultData) SetDistrictCode(v string) *QueryResultData {
	s.DistrictCode = &v
	return s
}

func (s *QueryResultData) SetDistrictName(v string) *QueryResultData {
	s.DistrictName = &v
	return s
}

func (s *QueryResultData) SetId(v string) *QueryResultData {
	s.Id = &v
	return s
}

func (s *QueryResultData) SetImages(v []*QueryResultDataImages) *QueryResultData {
	s.Images = v
	return s
}

func (s *QueryResultData) SetLatitude(v string) *QueryResultData {
	s.Latitude = &v
	return s
}

func (s *QueryResultData) SetLongitude(v string) *QueryResultData {
	s.Longitude = &v
	return s
}

func (s *QueryResultData) SetMetadata(v *QueryResultDataMetadata) *QueryResultData {
	s.Metadata = v
	return s
}

func (s *QueryResultData) SetName(v string) *QueryResultData {
	s.Name = &v
	return s
}

func (s *QueryResultData) SetProvinceCode(v string) *QueryResultData {
	s.ProvinceCode = &v
	return s
}

func (s *QueryResultData) SetProvinceName(v string) *QueryResultData {
	s.ProvinceName = &v
	return s
}

func (s *QueryResultData) SetTypeCode(v string) *QueryResultData {
	s.TypeCode = &v
	return s
}

func (s *QueryResultData) SetTypes(v string) *QueryResultData {
	s.Types = &v
	return s
}

type QueryResultDataImages struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	Url   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryResultDataImages) String() string {
	return tea.Prettify(s)
}

func (s QueryResultDataImages) GoString() string {
	return s.String()
}

func (s *QueryResultDataImages) SetTitle(v string) *QueryResultDataImages {
	s.Title = &v
	return s
}

func (s *QueryResultDataImages) SetUrl(v string) *QueryResultDataImages {
	s.Url = &v
	return s
}

type QueryResultDataMetadata struct {
	BusinessArea      *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	MainTag           *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	Phone             *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Score             *string `json:"score,omitempty" xml:"score,omitempty"`
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s QueryResultDataMetadata) String() string {
	return tea.Prettify(s)
}

func (s QueryResultDataMetadata) GoString() string {
	return s.String()
}

func (s *QueryResultDataMetadata) SetBusinessArea(v string) *QueryResultDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *QueryResultDataMetadata) SetDailyOpeningHours(v string) *QueryResultDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *QueryResultDataMetadata) SetMainTag(v string) *QueryResultDataMetadata {
	s.MainTag = &v
	return s
}

func (s *QueryResultDataMetadata) SetPhone(v string) *QueryResultDataMetadata {
	s.Phone = &v
	return s
}

func (s *QueryResultDataMetadata) SetScore(v string) *QueryResultDataMetadata {
	s.Score = &v
	return s
}

func (s *QueryResultDataMetadata) SetWeeklyOpeningDays(v string) *QueryResultDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

type CommonQueryBySceneRequest struct {
	Body *CommonAgentQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonQueryBySceneRequest) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneRequest) SetBody(v *CommonAgentQuery) *CommonQueryBySceneRequest {
	s.Body = v
	return s
}

type CommonQueryBySceneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonQueryBySceneResponse) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneResponse) SetHeaders(v map[string]*string) *CommonQueryBySceneResponse {
	s.Headers = v
	return s
}

func (s *CommonQueryBySceneResponse) SetStatusCode(v int32) *CommonQueryBySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonQueryBySceneResponse) SetBody(v *QueryResult) *CommonQueryBySceneResponse {
	s.Body = v
	return s
}

type QueryAttractionsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttractionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsRequest) GoString() string {
	return s.String()
}

func (s *QueryAttractionsRequest) SetBody(v *AgentBaseQuery) *QueryAttractionsRequest {
	s.Body = v
	return s
}

type QueryAttractionsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryAttractionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAttractionsResponseBody) SetQueryResult(v *QueryResult) *QueryAttractionsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryAttractionsResponseBody) SetRequestId(v string) *QueryAttractionsResponseBody {
	s.RequestId = &v
	return s
}

type QueryAttractionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAttractionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAttractionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAttractionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAttractionsResponse) SetHeaders(v map[string]*string) *QueryAttractionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAttractionsResponse) SetStatusCode(v int32) *QueryAttractionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAttractionsResponse) SetBody(v *QueryAttractionsResponseBody) *QueryAttractionsResponse {
	s.Body = v
	return s
}

type QueryHotelsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelsRequest) SetBody(v *AgentBaseQuery) *QueryHotelsRequest {
	s.Body = v
	return s
}

type QueryHotelsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 78032F8B-0157-53F9-B1A8-3BD86ADE38D0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryHotelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotelsResponseBody) SetQueryResult(v *QueryResult) *QueryHotelsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryHotelsResponseBody) SetRequestId(v string) *QueryHotelsResponseBody {
	s.RequestId = &v
	return s
}

type QueryHotelsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotelsResponse) GoString() string {
	return s.String()
}

func (s *QueryHotelsResponse) SetHeaders(v map[string]*string) *QueryHotelsResponse {
	s.Headers = v
	return s
}

func (s *QueryHotelsResponse) SetStatusCode(v int32) *QueryHotelsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelsResponse) SetBody(v *QueryHotelsResponseBody) *QueryHotelsResponse {
	s.Body = v
	return s
}

type QueryRestaurantsRequest struct {
	Body *AgentBaseQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRestaurantsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsRequest) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsRequest) SetBody(v *AgentBaseQuery) *QueryRestaurantsRequest {
	s.Body = v
	return s
}

type QueryRestaurantsResponseBody struct {
	QueryResult *QueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D78E16C0-4D44-5065-BFF7-127F7859F687
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryRestaurantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsResponseBody) SetQueryResult(v *QueryResult) *QueryRestaurantsResponseBody {
	s.QueryResult = v
	return s
}

func (s *QueryRestaurantsResponseBody) SetRequestId(v string) *QueryRestaurantsResponseBody {
	s.RequestId = &v
	return s
}

type QueryRestaurantsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRestaurantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRestaurantsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRestaurantsResponse) GoString() string {
	return s.String()
}

func (s *QueryRestaurantsResponse) SetHeaders(v map[string]*string) *QueryRestaurantsResponse {
	s.Headers = v
	return s
}

func (s *QueryRestaurantsResponse) SetStatusCode(v int32) *QueryRestaurantsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRestaurantsResponse) SetBody(v *QueryRestaurantsResponseBody) *QueryRestaurantsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 自然语言通用查询
//
// @param request - CommonQueryBySceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryBySceneWithOptions(request *CommonQueryBySceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CommonQueryBySceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CommonQueryByScene"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v2/nl/common"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自然语言通用查询
//
// @param request - CommonQueryBySceneRequest
//
// @return CommonQueryBySceneResponse
func (client *Client) CommonQueryByScene(request *CommonQueryBySceneRequest) (_result *CommonQueryBySceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CommonQueryBySceneResponse{}
	_body, _err := client.CommonQueryBySceneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 景点查询
//
// @param request - QueryAttractionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAttractionsResponse
func (client *Client) QueryAttractionsWithOptions(request *QueryAttractionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAttractionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAttractions"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/attractions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAttractionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 景点查询
//
// @param request - QueryAttractionsRequest
//
// @return QueryAttractionsResponse
func (client *Client) QueryAttractions(request *QueryAttractionsRequest) (_result *QueryAttractionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAttractionsResponse{}
	_body, _err := client.QueryAttractionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店查询
//
// @param request - QueryHotelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHotelsResponse
func (client *Client) QueryHotelsWithOptions(request *QueryHotelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryHotelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotels"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/hotels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店查询
//
// @param request - QueryHotelsRequest
//
// @return QueryHotelsResponse
func (client *Client) QueryHotels(request *QueryHotelsRequest) (_result *QueryHotelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryHotelsResponse{}
	_body, _err := client.QueryHotelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 餐厅查询
//
// @param request - QueryRestaurantsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRestaurantsResponse
func (client *Client) QueryRestaurantsWithOptions(request *QueryRestaurantsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryRestaurantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRestaurants"),
		Version:     tea.String("2024-07-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/amap-function-call-agent/iqs-agent-service/v1/nl/restaurants"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRestaurantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 餐厅查询
//
// @param request - QueryRestaurantsRequest
//
// @return QueryRestaurantsResponse
func (client *Client) QueryRestaurants(request *QueryRestaurantsRequest) (_result *QueryRestaurantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRestaurantsResponse{}
	_body, _err := client.QueryRestaurantsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
