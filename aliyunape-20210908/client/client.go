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

type ExecuteRequest struct {
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// source
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// extendParam
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// serviceParam
	ServiceParam map[string]*string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// aliyunPk
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ExecuteRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequest) SetAppName(v string) *ExecuteRequest {
	s.AppName = &v
	return s
}

func (s *ExecuteRequest) SetChannel(v string) *ExecuteRequest {
	s.Channel = &v
	return s
}

func (s *ExecuteRequest) SetExtendParam(v map[string]*string) *ExecuteRequest {
	s.ExtendParam = v
	return s
}

func (s *ExecuteRequest) SetOrderId(v string) *ExecuteRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteRequest) SetRequestId(v string) *ExecuteRequest {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequest) SetServiceParam(v map[string]*string) *ExecuteRequest {
	s.ServiceParam = v
	return s
}

func (s *ExecuteRequest) SetUserId(v int64) *ExecuteRequest {
	s.UserId = &v
	return s
}

type ExecuteShrinkRequest struct {
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// source
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// extendParam
	ExtendParamShrink *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// serviceParam
	ServiceParamShrink *string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// aliyunPk
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ExecuteShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteShrinkRequest) SetAppName(v string) *ExecuteShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ExecuteShrinkRequest) SetChannel(v string) *ExecuteShrinkRequest {
	s.Channel = &v
	return s
}

func (s *ExecuteShrinkRequest) SetExtendParamShrink(v string) *ExecuteShrinkRequest {
	s.ExtendParamShrink = &v
	return s
}

func (s *ExecuteShrinkRequest) SetOrderId(v string) *ExecuteShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteShrinkRequest) SetRequestId(v string) *ExecuteShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *ExecuteShrinkRequest) SetServiceParamShrink(v string) *ExecuteShrinkRequest {
	s.ServiceParamShrink = &v
	return s
}

func (s *ExecuteShrinkRequest) SetUserId(v int64) *ExecuteShrinkRequest {
	s.UserId = &v
	return s
}

type ExecuteResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteResponseBody) SetCode(v string) *ExecuteResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteResponseBody) SetData(v []map[string]interface{}) *ExecuteResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteResponseBody) SetMessage(v string) *ExecuteResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteResponseBody) SetRequestId(v string) *ExecuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteResponseBody) SetRt(v int64) *ExecuteResponseBody {
	s.Rt = &v
	return s
}

func (s *ExecuteResponseBody) SetSuccess(v bool) *ExecuteResponseBody {
	s.Success = &v
	return s
}

type ExecuteResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResponse) GoString() string {
	return s.String()
}

func (s *ExecuteResponse) SetHeaders(v map[string]*string) *ExecuteResponse {
	s.Headers = v
	return s
}

func (s *ExecuteResponse) SetBody(v *ExecuteResponseBody) *ExecuteResponse {
	s.Body = v
	return s
}

type WeathermonitorProvinceHourRequest struct {
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 渠道名称
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// 扩展参数
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求参数
	ServiceParam map[string]*string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s WeathermonitorProvinceHourRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourRequest) SetAppName(v string) *WeathermonitorProvinceHourRequest {
	s.AppName = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetChannel(v string) *WeathermonitorProvinceHourRequest {
	s.Channel = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetExtendParam(v map[string]*string) *WeathermonitorProvinceHourRequest {
	s.ExtendParam = v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetOrderId(v string) *WeathermonitorProvinceHourRequest {
	s.OrderId = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetRequestId(v string) *WeathermonitorProvinceHourRequest {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetServiceParam(v map[string]*string) *WeathermonitorProvinceHourRequest {
	s.ServiceParam = v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetUserId(v int64) *WeathermonitorProvinceHourRequest {
	s.UserId = &v
	return s
}

type WeathermonitorProvinceHourShrinkRequest struct {
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 渠道名称
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// 扩展参数
	ExtendParamShrink *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求参数
	ServiceParamShrink *string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s WeathermonitorProvinceHourShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourShrinkRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetAppName(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.AppName = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetChannel(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.Channel = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetExtendParamShrink(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.ExtendParamShrink = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetOrderId(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetRequestId(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetServiceParamShrink(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.ServiceParamShrink = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetUserId(v int64) *WeathermonitorProvinceHourShrinkRequest {
	s.UserId = &v
	return s
}

type WeathermonitorProvinceHourResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeathermonitorProvinceHourResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourResponseBody) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourResponseBody) SetCode(v string) *WeathermonitorProvinceHourResponseBody {
	s.Code = &v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetData(v []map[string]interface{}) *WeathermonitorProvinceHourResponseBody {
	s.Data = v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetMessage(v string) *WeathermonitorProvinceHourResponseBody {
	s.Message = &v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetRequestId(v string) *WeathermonitorProvinceHourResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetRt(v int64) *WeathermonitorProvinceHourResponseBody {
	s.Rt = &v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetSuccess(v bool) *WeathermonitorProvinceHourResponseBody {
	s.Success = &v
	return s
}

type WeathermonitorProvinceHourResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeathermonitorProvinceHourResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeathermonitorProvinceHourResponse) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourResponse) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourResponse) SetHeaders(v map[string]*string) *WeathermonitorProvinceHourResponse {
	s.Headers = v
	return s
}

func (s *WeathermonitorProvinceHourResponse) SetBody(v *WeathermonitorProvinceHourResponseBody) *WeathermonitorProvinceHourResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliyunape"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ExecuteWithOptions(tmpReq *ExecuteRequest, runtime *util.RuntimeOptions) (_result *ExecuteResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendParam)) {
		request.ExtendParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendParam, tea.String("ExtendParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ServiceParam)) {
		request.ServiceParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceParam, tea.String("ServiceParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["Channel"] = request.Channel
	query["ExtendParam"] = request.ExtendParamShrink
	query["OrderId"] = request.OrderId
	query["RequestId"] = request.RequestId
	query["ServiceParam"] = request.ServiceParamShrink
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("Execute"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Execute(request *ExecuteRequest) (_result *ExecuteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteResponse{}
	_body, _err := client.ExecuteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeathermonitorProvinceHourWithOptions(tmpReq *WeathermonitorProvinceHourRequest, runtime *util.RuntimeOptions) (_result *WeathermonitorProvinceHourResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &WeathermonitorProvinceHourShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendParam)) {
		request.ExtendParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendParam, tea.String("ExtendParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ServiceParam)) {
		request.ServiceParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceParam, tea.String("ServiceParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["Channel"] = request.Channel
	query["ExtendParam"] = request.ExtendParamShrink
	query["OrderId"] = request.OrderId
	query["RequestId"] = request.RequestId
	query["ServiceParam"] = request.ServiceParamShrink
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("WeathermonitorProvinceHour"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WeathermonitorProvinceHourResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WeathermonitorProvinceHour(request *WeathermonitorProvinceHourRequest) (_result *WeathermonitorProvinceHourResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeathermonitorProvinceHourResponse{}
	_body, _err := client.WeathermonitorProvinceHourWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
