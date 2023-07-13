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

type GetDeviceInfoRequest struct {
	DeviceId  *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	Ds        *string `json:"ds,omitempty" xml:"ds,omitempty"`
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoRequest) SetDeviceId(v string) *GetDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoRequest) SetDs(v string) *GetDeviceInfoRequest {
	s.Ds = &v
	return s
}

func (s *GetDeviceInfoRequest) SetFactoryId(v string) *GetDeviceInfoRequest {
	s.FactoryId = &v
	return s
}

type GetDeviceInfoResponseBody struct {
	Code      *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *GetDeviceInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpCode  *int32                         `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBody) SetCode(v string) *GetDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetData(v *GetDeviceInfoResponseBodyData) *GetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceInfoResponseBody) SetHttpCode(v int32) *GetDeviceInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetRequestId(v string) *GetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceInfoResponseBody) SetSuccess(v bool) *GetDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type GetDeviceInfoResponseBodyData struct {
	DeviceId       *string                                    `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName     *string                                    `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	FirstTypeName  *string                                    `json:"firstTypeName,omitempty" xml:"firstTypeName,omitempty"`
	RecordList     []*GetDeviceInfoResponseBodyDataRecordList `json:"recordList,omitempty" xml:"recordList,omitempty" type:"Repeated"`
	SecondTypeName *string                                    `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceId(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetDeviceName(v string) *GetDeviceInfoResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetFirstTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetRecordList(v []*GetDeviceInfoResponseBodyDataRecordList) *GetDeviceInfoResponseBodyData {
	s.RecordList = v
	return s
}

func (s *GetDeviceInfoResponseBodyData) SetSecondTypeName(v string) *GetDeviceInfoResponseBodyData {
	s.SecondTypeName = &v
	return s
}

type GetDeviceInfoResponseBodyDataRecordList struct {
	Identifier     *string  `json:"identifier,omitempty" xml:"identifier,omitempty"`
	ParamName      *string  `json:"paramName,omitempty" xml:"paramName,omitempty"`
	StatisticsDate *string  `json:"statisticsDate,omitempty" xml:"statisticsDate,omitempty"`
	Type           *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit           *string  `json:"unit,omitempty" xml:"unit,omitempty"`
	Value          *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDeviceInfoResponseBodyDataRecordList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponseBodyDataRecordList) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetIdentifier(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Identifier = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetParamName(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.ParamName = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetStatisticsDate(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.StatisticsDate = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetType(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Type = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetUnit(v string) *GetDeviceInfoResponseBodyDataRecordList {
	s.Unit = &v
	return s
}

func (s *GetDeviceInfoResponseBodyDataRecordList) SetValue(v float64) *GetDeviceInfoResponseBodyDataRecordList {
	s.Value = &v
	return s
}

type GetDeviceInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponse) SetHeaders(v map[string]*string) *GetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceInfoResponse) SetStatusCode(v int32) *GetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceInfoResponse) SetBody(v *GetDeviceInfoResponseBody) *GetDeviceInfoResponse {
	s.Body = v
	return s
}

type GetDeviceListRequest struct {
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceListRequest) SetFactoryId(v string) *GetDeviceListRequest {
	s.FactoryId = &v
	return s
}

type GetDeviceListResponseBody struct {
	Code      *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *GetDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpCode  *int32                         `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBody) SetCode(v string) *GetDeviceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBody) SetData(v *GetDeviceListResponseBodyData) *GetDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceListResponseBody) SetHttpCode(v int32) *GetDeviceListResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBody) SetRequestId(v string) *GetDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceListResponseBody) SetSuccess(v bool) *GetDeviceListResponseBody {
	s.Success = &v
	return s
}

type GetDeviceListResponseBodyData struct {
	Code       *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	DeviceList []*GetDeviceListResponseBodyDataDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	FactoryId  *string                                    `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	HttpCode   *int32                                     `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	Success    *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyData) SetCode(v string) *GetDeviceListResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetDeviceList(v []*GetDeviceListResponseBodyDataDeviceList) *GetDeviceListResponseBodyData {
	s.DeviceList = v
	return s
}

func (s *GetDeviceListResponseBodyData) SetFactoryId(v string) *GetDeviceListResponseBodyData {
	s.FactoryId = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetHttpCode(v int32) *GetDeviceListResponseBodyData {
	s.HttpCode = &v
	return s
}

func (s *GetDeviceListResponseBodyData) SetSuccess(v bool) *GetDeviceListResponseBodyData {
	s.Success = &v
	return s
}

type GetDeviceListResponseBodyDataDeviceList struct {
	DeviceId       *string                                      `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName     *string                                      `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	FirstTypeName  *string                                      `json:"firstTypeName,omitempty" xml:"firstTypeName,omitempty"`
	Info           *GetDeviceListResponseBodyDataDeviceListInfo `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	ParentDevice   *string                                      `json:"parentDevice,omitempty" xml:"parentDevice,omitempty"`
	SecondTypeName *string                                      `json:"secondTypeName,omitempty" xml:"secondTypeName,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceList) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceList) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceId(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetDeviceName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetFirstTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.FirstTypeName = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetInfo(v *GetDeviceListResponseBodyDataDeviceListInfo) *GetDeviceListResponseBodyDataDeviceList {
	s.Info = v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetParentDevice(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.ParentDevice = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceList) SetSecondTypeName(v string) *GetDeviceListResponseBodyDataDeviceList {
	s.SecondTypeName = &v
	return s
}

type GetDeviceListResponseBodyDataDeviceListInfo struct {
	ConstKva      *int32 `json:"constKva,omitempty" xml:"constKva,omitempty"`
	Ct            *int32 `json:"ct,omitempty" xml:"ct,omitempty"`
	Magnification *int32 `json:"magnification,omitempty" xml:"magnification,omitempty"`
	Pressure      *int32 `json:"pressure,omitempty" xml:"pressure,omitempty"`
	Pt            *int32 `json:"pt,omitempty" xml:"pt,omitempty"`
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponseBodyDataDeviceListInfo) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetConstKva(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.ConstKva = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetCt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Ct = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetMagnification(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Magnification = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPressure(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pressure = &v
	return s
}

func (s *GetDeviceListResponseBodyDataDeviceListInfo) SetPt(v int32) *GetDeviceListResponseBodyDataDeviceListInfo {
	s.Pt = &v
	return s
}

type GetDeviceListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponse) SetHeaders(v map[string]*string) *GetDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceListResponse) SetStatusCode(v int32) *GetDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceListResponse) SetBody(v *GetDeviceListResponseBody) *GetDeviceListResponse {
	s.Body = v
	return s
}

type GetOrgAndFactoryResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data      []*GetOrgAndFactoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HttpCode  *int32                              `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetOrgAndFactoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBody) SetCode(v string) *GetOrgAndFactoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetData(v []*GetOrgAndFactoryResponseBodyData) *GetOrgAndFactoryResponseBody {
	s.Data = v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetHttpCode(v int32) *GetOrgAndFactoryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetRequestId(v string) *GetOrgAndFactoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetSuccess(v bool) *GetOrgAndFactoryResponseBody {
	s.Success = &v
	return s
}

type GetOrgAndFactoryResponseBodyData struct {
	AliyunPk         *string                                        `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	FactoryList      []*GetOrgAndFactoryResponseBodyDataFactoryList `json:"factoryList,omitempty" xml:"factoryList,omitempty" type:"Repeated"`
	OrganizationId   *string                                        `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	OrganizationName *string                                        `json:"organizationName,omitempty" xml:"organizationName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyData) SetAliyunPk(v string) *GetOrgAndFactoryResponseBodyData {
	s.AliyunPk = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetFactoryList(v []*GetOrgAndFactoryResponseBodyDataFactoryList) *GetOrgAndFactoryResponseBodyData {
	s.FactoryList = v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationId(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationName(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationName = &v
	return s
}

type GetOrgAndFactoryResponseBodyDataFactoryList struct {
	FactoryId   *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	FactoryName *string `json:"factoryName,omitempty" xml:"factoryName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryId(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryName(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryName = &v
	return s
}

type GetOrgAndFactoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrgAndFactoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrgAndFactoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgAndFactoryResponse) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponse) SetHeaders(v map[string]*string) *GetOrgAndFactoryResponse {
	s.Headers = v
	return s
}

func (s *GetOrgAndFactoryResponse) SetStatusCode(v int32) *GetOrgAndFactoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgAndFactoryResponse) SetBody(v *GetOrgAndFactoryResponseBody) *GetOrgAndFactoryResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("energyexpertexternal"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetDeviceInfoWithOptions(request *GetDeviceInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeviceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ds)) {
		query["ds"] = request.Ds
	}

	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceInfo"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getDeviceInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceInfo(request *GetDeviceInfoRequest) (_result *GetDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.GetDeviceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceListWithOptions(request *GetDeviceListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FactoryId)) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceList"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getDeviceList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceList(request *GetDeviceListRequest) (_result *GetDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceListResponse{}
	_body, _err := client.GetDeviceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrgAndFactoryWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrgAndFactoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrgAndFactory"),
		Version:     tea.String("2022-09-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/external/getOrgAndFactory"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrgAndFactory() (_result *GetOrgAndFactoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.GetOrgAndFactoryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
