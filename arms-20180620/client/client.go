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

type ARMSQueryDataSetRequest struct {
	DatasetId     *int64                                 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DateStr       *string                                `json:"DateStr,omitempty" xml:"DateStr,omitempty"`
	Dimensions    []*ARMSQueryDataSetRequestDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	HungryMode    *bool                                  `json:"HungryMode,omitempty" xml:"HungryMode,omitempty"`
	IntervalInSec *int32                                 `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	IsDrillDown   *bool                                  `json:"IsDrillDown,omitempty" xml:"IsDrillDown,omitempty"`
	Limit         *int32                                 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxTime       *int64                                 `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	Measures      []*string                              `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	MinTime       *int64                                 `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	OptionalDims  []*ARMSQueryDataSetRequestOptionalDims `json:"OptionalDims,omitempty" xml:"OptionalDims,omitempty" type:"Repeated"`
	OrderByKey    *string                                `json:"OrderByKey,omitempty" xml:"OrderByKey,omitempty"`
	ReduceTail    *bool                                  `json:"ReduceTail,omitempty" xml:"ReduceTail,omitempty"`
	RequiredDims  []*ARMSQueryDataSetRequestRequiredDims `json:"RequiredDims,omitempty" xml:"RequiredDims,omitempty" type:"Repeated"`
	SecurityToken *string                                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ARMSQueryDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequest) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequest) SetDatasetId(v int64) *ARMSQueryDataSetRequest {
	s.DatasetId = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDateStr(v string) *ARMSQueryDataSetRequest {
	s.DateStr = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetDimensions(v []*ARMSQueryDataSetRequestDimensions) *ARMSQueryDataSetRequest {
	s.Dimensions = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetHungryMode(v bool) *ARMSQueryDataSetRequest {
	s.HungryMode = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIntervalInSec(v int32) *ARMSQueryDataSetRequest {
	s.IntervalInSec = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetIsDrillDown(v bool) *ARMSQueryDataSetRequest {
	s.IsDrillDown = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetLimit(v int32) *ARMSQueryDataSetRequest {
	s.Limit = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMaxTime(v int64) *ARMSQueryDataSetRequest {
	s.MaxTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMeasures(v []*string) *ARMSQueryDataSetRequest {
	s.Measures = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetMinTime(v int64) *ARMSQueryDataSetRequest {
	s.MinTime = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOptionalDims(v []*ARMSQueryDataSetRequestOptionalDims) *ARMSQueryDataSetRequest {
	s.OptionalDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetOrderByKey(v string) *ARMSQueryDataSetRequest {
	s.OrderByKey = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetReduceTail(v bool) *ARMSQueryDataSetRequest {
	s.ReduceTail = &v
	return s
}

func (s *ARMSQueryDataSetRequest) SetRequiredDims(v []*ARMSQueryDataSetRequestRequiredDims) *ARMSQueryDataSetRequest {
	s.RequiredDims = v
	return s
}

func (s *ARMSQueryDataSetRequest) SetSecurityToken(v string) *ARMSQueryDataSetRequest {
	s.SecurityToken = &v
	return s
}

type ARMSQueryDataSetRequestDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestDimensions) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestDimensions) SetKey(v string) *ARMSQueryDataSetRequestDimensions {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetType(v string) *ARMSQueryDataSetRequestDimensions {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestDimensions) SetValue(v string) *ARMSQueryDataSetRequestDimensions {
	s.Value = &v
	return s
}

type ARMSQueryDataSetRequestOptionalDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestOptionalDims) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestOptionalDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetKey(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetType(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestOptionalDims) SetValue(v string) *ARMSQueryDataSetRequestOptionalDims {
	s.Value = &v
	return s
}

type ARMSQueryDataSetRequestRequiredDims struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ARMSQueryDataSetRequestRequiredDims) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetRequestRequiredDims) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetKey(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Key = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetType(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Type = &v
	return s
}

func (s *ARMSQueryDataSetRequestRequiredDims) SetValue(v string) *ARMSQueryDataSetRequestRequiredDims {
	s.Value = &v
	return s
}

type ARMSQueryDataSetResponseBody struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ARMSQueryDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponseBody) SetData(v string) *ARMSQueryDataSetResponseBody {
	s.Data = &v
	return s
}

type ARMSQueryDataSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ARMSQueryDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ARMSQueryDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ARMSQueryDataSetResponse) GoString() string {
	return s.String()
}

func (s *ARMSQueryDataSetResponse) SetHeaders(v map[string]*string) *ARMSQueryDataSetResponse {
	s.Headers = v
	return s
}

func (s *ARMSQueryDataSetResponse) SetStatusCode(v int32) *ARMSQueryDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ARMSQueryDataSetResponse) SetBody(v *ARMSQueryDataSetResponseBody) *ARMSQueryDataSetResponse {
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("arms.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("arms.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("arms.aliyuncs.com"),
		"cn-edge-1":                   tea.String("arms.aliyuncs.com"),
		"cn-fujian":                   tea.String("arms.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("arms.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("arms.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("arms.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("arms.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("arms.aliyuncs.com"),
		"cn-wuhan":                    tea.String("arms.aliyuncs.com"),
		"cn-yushanfang":               tea.String("arms.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("arms.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("arms.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("arms.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("arms.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("arms.aliyuncs.com"),
		"me-east-1":                   tea.String("arms.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("arms.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("arms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ARMSQueryDataSetWithOptions(request *ARMSQueryDataSetRequest, runtime *util.RuntimeOptions) (_result *ARMSQueryDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.DateStr)) {
		query["DateStr"] = request.DateStr
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.HungryMode)) {
		query["HungryMode"] = request.HungryMode
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.IsDrillDown)) {
		query["IsDrillDown"] = request.IsDrillDown
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		query["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		query["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.OptionalDims)) {
		query["OptionalDims"] = request.OptionalDims
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByKey)) {
		query["OrderByKey"] = request.OrderByKey
	}

	if !tea.BoolValue(util.IsUnset(request.ReduceTail)) {
		query["ReduceTail"] = request.ReduceTail
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredDims)) {
		query["RequiredDims"] = request.RequiredDims
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ARMSQueryDataSet"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ARMSQueryDataSet(request *ARMSQueryDataSetRequest) (_result *ARMSQueryDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.ARMSQueryDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
