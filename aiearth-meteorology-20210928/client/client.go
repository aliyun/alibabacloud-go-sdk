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

type GridQueryRequest struct {
	Element           *string  `json:"element,omitempty" xml:"element,omitempty"`
	ForecastTimestamp *string  `json:"forecastTimestamp,omitempty" xml:"forecastTimestamp,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	PageNo            *int32   `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize          *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Product           *string  `json:"product,omitempty" xml:"product,omitempty"`
	ReportTimestamp   *string  `json:"reportTimestamp,omitempty" xml:"reportTimestamp,omitempty"`
}

func (s GridQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GridQueryRequest) GoString() string {
	return s.String()
}

func (s *GridQueryRequest) SetElement(v string) *GridQueryRequest {
	s.Element = &v
	return s
}

func (s *GridQueryRequest) SetForecastTimestamp(v string) *GridQueryRequest {
	s.ForecastTimestamp = &v
	return s
}

func (s *GridQueryRequest) SetLatitude(v float64) *GridQueryRequest {
	s.Latitude = &v
	return s
}

func (s *GridQueryRequest) SetLongitude(v float64) *GridQueryRequest {
	s.Longitude = &v
	return s
}

func (s *GridQueryRequest) SetPageNo(v int32) *GridQueryRequest {
	s.PageNo = &v
	return s
}

func (s *GridQueryRequest) SetPageSize(v int32) *GridQueryRequest {
	s.PageSize = &v
	return s
}

func (s *GridQueryRequest) SetProduct(v string) *GridQueryRequest {
	s.Product = &v
	return s
}

func (s *GridQueryRequest) SetReportTimestamp(v string) *GridQueryRequest {
	s.ReportTimestamp = &v
	return s
}

type GridQueryResponseBody struct {
	Code      *int64                       `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Module    *GridQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GridQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GridQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GridQueryResponseBody) SetCode(v int64) *GridQueryResponseBody {
	s.Code = &v
	return s
}

func (s *GridQueryResponseBody) SetMessage(v string) *GridQueryResponseBody {
	s.Message = &v
	return s
}

func (s *GridQueryResponseBody) SetModule(v *GridQueryResponseBodyModule) *GridQueryResponseBody {
	s.Module = v
	return s
}

func (s *GridQueryResponseBody) SetRequestId(v string) *GridQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GridQueryResponseBody) SetSuccess(v bool) *GridQueryResponseBody {
	s.Success = &v
	return s
}

type GridQueryResponseBodyModule struct {
	List     []*GridQueryResponseBodyModuleList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNo   *int32                             `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GridQueryResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s GridQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GridQueryResponseBodyModule) SetList(v []*GridQueryResponseBodyModuleList) *GridQueryResponseBodyModule {
	s.List = v
	return s
}

func (s *GridQueryResponseBodyModule) SetPageNo(v int32) *GridQueryResponseBodyModule {
	s.PageNo = &v
	return s
}

func (s *GridQueryResponseBodyModule) SetPageSize(v int32) *GridQueryResponseBodyModule {
	s.PageSize = &v
	return s
}

type GridQueryResponseBodyModuleList struct {
	DataType          *string  `json:"dataType,omitempty" xml:"dataType,omitempty"`
	Element           *string  `json:"element,omitempty" xml:"element,omitempty"`
	ElementUnit       *string  `json:"elementUnit,omitempty" xml:"elementUnit,omitempty"`
	ForecastTimestamp *string  `json:"forecastTimestamp,omitempty" xml:"forecastTimestamp,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	ReportTimestamp   *string  `json:"reportTimestamp,omitempty" xml:"reportTimestamp,omitempty"`
	Value             *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GridQueryResponseBodyModuleList) String() string {
	return tea.Prettify(s)
}

func (s GridQueryResponseBodyModuleList) GoString() string {
	return s.String()
}

func (s *GridQueryResponseBodyModuleList) SetDataType(v string) *GridQueryResponseBodyModuleList {
	s.DataType = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetElement(v string) *GridQueryResponseBodyModuleList {
	s.Element = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetElementUnit(v string) *GridQueryResponseBodyModuleList {
	s.ElementUnit = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetForecastTimestamp(v string) *GridQueryResponseBodyModuleList {
	s.ForecastTimestamp = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetLatitude(v float64) *GridQueryResponseBodyModuleList {
	s.Latitude = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetLongitude(v float64) *GridQueryResponseBodyModuleList {
	s.Longitude = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetReportTimestamp(v string) *GridQueryResponseBodyModuleList {
	s.ReportTimestamp = &v
	return s
}

func (s *GridQueryResponseBodyModuleList) SetValue(v float64) *GridQueryResponseBodyModuleList {
	s.Value = &v
	return s
}

type GridQueryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GridQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GridQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GridQueryResponse) GoString() string {
	return s.String()
}

func (s *GridQueryResponse) SetHeaders(v map[string]*string) *GridQueryResponse {
	s.Headers = v
	return s
}

func (s *GridQueryResponse) SetStatusCode(v int32) *GridQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GridQueryResponse) SetBody(v *GridQueryResponseBody) *GridQueryResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aiearth-meteorology"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GridQueryWithOptions(dataType *string, request *GridQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GridQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Element)) {
		query["element"] = request.Element
	}

	if !tea.BoolValue(util.IsUnset(request.ForecastTimestamp)) {
		query["forecastTimestamp"] = request.ForecastTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["pageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.ReportTimestamp)) {
		query["reportTimestamp"] = request.ReportTimestamp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GridQuery"),
		Version:     tea.String("2021-09-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/grid/" + tea.StringValue(openapiutil.GetEncodeParam(dataType)) + "/v1"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GridQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GridQuery(dataType *string, request *GridQueryRequest) (_result *GridQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GridQueryResponse{}
	_body, _err := client.GridQueryWithOptions(dataType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
