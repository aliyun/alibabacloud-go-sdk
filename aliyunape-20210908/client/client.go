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
	// aliyunPk
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// serviceParam
	ServiceParam map[string]*string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// extendParam
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// source
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s ExecuteRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequest) SetUserId(v int64) *ExecuteRequest {
	s.UserId = &v
	return s
}

func (s *ExecuteRequest) SetServiceParam(v map[string]*string) *ExecuteRequest {
	s.ServiceParam = v
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

func (s *ExecuteRequest) SetAppName(v string) *ExecuteRequest {
	s.AppName = &v
	return s
}

func (s *ExecuteRequest) SetRequestId(v string) *ExecuteRequest {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequest) SetChannel(v string) *ExecuteRequest {
	s.Channel = &v
	return s
}

type ExecuteShrinkRequest struct {
	// aliyunPk
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// serviceParam
	ServiceParamShrink *string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// extendParam
	ExtendParamShrink *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// source
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s ExecuteShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteShrinkRequest) SetUserId(v int64) *ExecuteShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ExecuteShrinkRequest) SetServiceParamShrink(v string) *ExecuteShrinkRequest {
	s.ServiceParamShrink = &v
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

func (s *ExecuteShrinkRequest) SetAppName(v string) *ExecuteShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ExecuteShrinkRequest) SetRequestId(v string) *ExecuteShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *ExecuteShrinkRequest) SetChannel(v string) *ExecuteShrinkRequest {
	s.Channel = &v
	return s
}

type ExecuteResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteResponseBody) SetRt(v int64) *ExecuteResponseBody {
	s.Rt = &v
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

func (s *ExecuteResponseBody) SetData(v []map[string]interface{}) *ExecuteResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteResponseBody) SetCode(v string) *ExecuteResponseBody {
	s.Code = &v
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
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 请求参数
	ServiceParam map[string]*string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// 扩展参数
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 渠道名称
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s WeathermonitorProvinceHourRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourRequest) SetUserId(v int64) *WeathermonitorProvinceHourRequest {
	s.UserId = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetServiceParam(v map[string]*string) *WeathermonitorProvinceHourRequest {
	s.ServiceParam = v
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

func (s *WeathermonitorProvinceHourRequest) SetAppName(v string) *WeathermonitorProvinceHourRequest {
	s.AppName = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetRequestId(v string) *WeathermonitorProvinceHourRequest {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorProvinceHourRequest) SetChannel(v string) *WeathermonitorProvinceHourRequest {
	s.Channel = &v
	return s
}

type WeathermonitorProvinceHourShrinkRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 请求参数
	ServiceParamShrink *string `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty"`
	// 扩展参数
	ExtendParamShrink *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// orderId
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 渠道名称
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s WeathermonitorProvinceHourShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourShrinkRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetUserId(v int64) *WeathermonitorProvinceHourShrinkRequest {
	s.UserId = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetServiceParamShrink(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.ServiceParamShrink = &v
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

func (s *WeathermonitorProvinceHourShrinkRequest) SetAppName(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.AppName = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetRequestId(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorProvinceHourShrinkRequest) SetChannel(v string) *WeathermonitorProvinceHourShrinkRequest {
	s.Channel = &v
	return s
}

type WeathermonitorProvinceHourResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeathermonitorProvinceHourResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorProvinceHourResponseBody) GoString() string {
	return s.String()
}

func (s *WeathermonitorProvinceHourResponseBody) SetRt(v int64) *WeathermonitorProvinceHourResponseBody {
	s.Rt = &v
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

func (s *WeathermonitorProvinceHourResponseBody) SetData(v []map[string]interface{}) *WeathermonitorProvinceHourResponseBody {
	s.Data = v
	return s
}

func (s *WeathermonitorProvinceHourResponseBody) SetCode(v string) *WeathermonitorProvinceHourResponseBody {
	s.Code = &v
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

type WeathermonitorRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户中心--我的订单--订单请求--实例名称：aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 页面条数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 气象实况时间 yyyymmddhh0000 （数据最小时间2021-08-16）（小时）	20210817120000
	CurHour *string `json:"CurHour,omitempty" xml:"CurHour,omitempty"`
	// 页码
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s WeathermonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorRequest) SetUserId(v int64) *WeathermonitorRequest {
	s.UserId = &v
	return s
}

func (s *WeathermonitorRequest) SetOrderId(v string) *WeathermonitorRequest {
	s.OrderId = &v
	return s
}

func (s *WeathermonitorRequest) SetRequestId(v string) *WeathermonitorRequest {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorRequest) SetPageSize(v int32) *WeathermonitorRequest {
	s.PageSize = &v
	return s
}

func (s *WeathermonitorRequest) SetCurHour(v string) *WeathermonitorRequest {
	s.CurHour = &v
	return s
}

func (s *WeathermonitorRequest) SetPageNum(v int32) *WeathermonitorRequest {
	s.PageNum = &v
	return s
}

type WeathermonitorResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeathermonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorResponseBody) GoString() string {
	return s.String()
}

func (s *WeathermonitorResponseBody) SetRt(v int64) *WeathermonitorResponseBody {
	s.Rt = &v
	return s
}

func (s *WeathermonitorResponseBody) SetMessage(v string) *WeathermonitorResponseBody {
	s.Message = &v
	return s
}

func (s *WeathermonitorResponseBody) SetRequestId(v string) *WeathermonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorResponseBody) SetData(v []map[string]interface{}) *WeathermonitorResponseBody {
	s.Data = v
	return s
}

func (s *WeathermonitorResponseBody) SetCode(v string) *WeathermonitorResponseBody {
	s.Code = &v
	return s
}

func (s *WeathermonitorResponseBody) SetSuccess(v bool) *WeathermonitorResponseBody {
	s.Success = &v
	return s
}

type WeathermonitorResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeathermonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeathermonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorResponse) GoString() string {
	return s.String()
}

func (s *WeathermonitorResponse) SetHeaders(v map[string]*string) *WeathermonitorResponse {
	s.Headers = v
	return s
}

func (s *WeathermonitorResponse) SetBody(v *WeathermonitorResponseBody) *WeathermonitorResponse {
	s.Body = v
	return s
}

type WeatherforecastTimeRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户中心--我的订单--订单请求--实例名称：aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 经度，范围为（70°E~139.96°E）
	Lon *string `json:"Lon,omitempty" xml:"Lon,omitempty"`
	// 20210809090000
	CurHour *string `json:"CurHour,omitempty" xml:"CurHour,omitempty"`
	// 纬度，范围为（15°N~59.95°N
	Lat *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
}

func (s WeatherforecastTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeRequest) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeRequest) SetUserId(v int64) *WeatherforecastTimeRequest {
	s.UserId = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetOrderId(v string) *WeatherforecastTimeRequest {
	s.OrderId = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetRequestId(v string) *WeatherforecastTimeRequest {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetLon(v string) *WeatherforecastTimeRequest {
	s.Lon = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetCurHour(v string) *WeatherforecastTimeRequest {
	s.CurHour = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetLat(v string) *WeatherforecastTimeRequest {
	s.Lat = &v
	return s
}

type WeatherforecastTimeResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeatherforecastTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeResponseBody) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeResponseBody) SetRt(v int64) *WeatherforecastTimeResponseBody {
	s.Rt = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetMessage(v string) *WeatherforecastTimeResponseBody {
	s.Message = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetRequestId(v string) *WeatherforecastTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetData(v []map[string]interface{}) *WeatherforecastTimeResponseBody {
	s.Data = v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetCode(v string) *WeatherforecastTimeResponseBody {
	s.Code = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetSuccess(v bool) *WeatherforecastTimeResponseBody {
	s.Success = &v
	return s
}

type WeatherforecastTimeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeatherforecastTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeatherforecastTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeResponse) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeResponse) SetHeaders(v map[string]*string) *WeatherforecastTimeResponse {
	s.Headers = v
	return s
}

func (s *WeatherforecastTimeResponse) SetBody(v *WeatherforecastTimeResponseBody) *WeatherforecastTimeResponse {
	s.Body = v
	return s
}

type StationDayRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户中心--我的订单--订单请求--实例名称：aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 气象预测开始时间
	StartForecast *string `json:"StartForecast,omitempty" xml:"StartForecast,omitempty"`
	// 全国站点（入参单一站点）
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
}

func (s StationDayRequest) String() string {
	return tea.Prettify(s)
}

func (s StationDayRequest) GoString() string {
	return s.String()
}

func (s *StationDayRequest) SetUserId(v int64) *StationDayRequest {
	s.UserId = &v
	return s
}

func (s *StationDayRequest) SetOrderId(v string) *StationDayRequest {
	s.OrderId = &v
	return s
}

func (s *StationDayRequest) SetRequestId(v string) *StationDayRequest {
	s.RequestId = &v
	return s
}

func (s *StationDayRequest) SetStartForecast(v string) *StationDayRequest {
	s.StartForecast = &v
	return s
}

func (s *StationDayRequest) SetStation(v string) *StationDayRequest {
	s.Station = &v
	return s
}

type StationDayResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StationDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StationDayResponseBody) GoString() string {
	return s.String()
}

func (s *StationDayResponseBody) SetRt(v int64) *StationDayResponseBody {
	s.Rt = &v
	return s
}

func (s *StationDayResponseBody) SetMessage(v string) *StationDayResponseBody {
	s.Message = &v
	return s
}

func (s *StationDayResponseBody) SetRequestId(v string) *StationDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *StationDayResponseBody) SetData(v []map[string]interface{}) *StationDayResponseBody {
	s.Data = v
	return s
}

func (s *StationDayResponseBody) SetCode(v string) *StationDayResponseBody {
	s.Code = &v
	return s
}

func (s *StationDayResponseBody) SetSuccess(v bool) *StationDayResponseBody {
	s.Success = &v
	return s
}

type StationDayResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StationDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StationDayResponse) String() string {
	return tea.Prettify(s)
}

func (s StationDayResponse) GoString() string {
	return s.String()
}

func (s *StationDayResponse) SetHeaders(v map[string]*string) *StationDayResponse {
	s.Headers = v
	return s
}

func (s *StationDayResponse) SetBody(v *StationDayResponseBody) *StationDayResponse {
	s.Body = v
	return s
}

type WeatherforecastRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户中心--我的订单--订单请求--实例名称：aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// yyyymmdd080000或yyyymmdd200000
	StartForecast *string `json:"StartForecast,omitempty" xml:"StartForecast,omitempty"`
	// 经度，范围为（70°E~139.96°E）
	Lon *string `json:"Lon,omitempty" xml:"Lon,omitempty"`
	// 纬度，范围为（15°N~59.95°N）
	Lat *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
}

func (s WeatherforecastRequest) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastRequest) GoString() string {
	return s.String()
}

func (s *WeatherforecastRequest) SetUserId(v int64) *WeatherforecastRequest {
	s.UserId = &v
	return s
}

func (s *WeatherforecastRequest) SetOrderId(v string) *WeatherforecastRequest {
	s.OrderId = &v
	return s
}

func (s *WeatherforecastRequest) SetRequestId(v string) *WeatherforecastRequest {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastRequest) SetStartForecast(v string) *WeatherforecastRequest {
	s.StartForecast = &v
	return s
}

func (s *WeatherforecastRequest) SetLon(v string) *WeatherforecastRequest {
	s.Lon = &v
	return s
}

func (s *WeatherforecastRequest) SetLat(v string) *WeatherforecastRequest {
	s.Lat = &v
	return s
}

type WeatherforecastResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeatherforecastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastResponseBody) GoString() string {
	return s.String()
}

func (s *WeatherforecastResponseBody) SetRt(v int64) *WeatherforecastResponseBody {
	s.Rt = &v
	return s
}

func (s *WeatherforecastResponseBody) SetMessage(v string) *WeatherforecastResponseBody {
	s.Message = &v
	return s
}

func (s *WeatherforecastResponseBody) SetRequestId(v string) *WeatherforecastResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastResponseBody) SetData(v []map[string]interface{}) *WeatherforecastResponseBody {
	s.Data = v
	return s
}

func (s *WeatherforecastResponseBody) SetCode(v string) *WeatherforecastResponseBody {
	s.Code = &v
	return s
}

func (s *WeatherforecastResponseBody) SetSuccess(v bool) *WeatherforecastResponseBody {
	s.Success = &v
	return s
}

type WeatherforecastResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeatherforecastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeatherforecastResponse) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastResponse) GoString() string {
	return s.String()
}

func (s *WeatherforecastResponse) SetHeaders(v map[string]*string) *WeatherforecastResponse {
	s.Headers = v
	return s
}

func (s *WeatherforecastResponse) SetBody(v *WeatherforecastResponseBody) *WeatherforecastResponse {
	s.Body = v
	return s
}

type HistoricalRequest struct {
	// UserId
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户中心--我的订单--订单请求--实例名称：aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 全国（入参单一站点）
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
	// pageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// startTime
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// endTime
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// pageNum
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s HistoricalRequest) String() string {
	return tea.Prettify(s)
}

func (s HistoricalRequest) GoString() string {
	return s.String()
}

func (s *HistoricalRequest) SetUserId(v int64) *HistoricalRequest {
	s.UserId = &v
	return s
}

func (s *HistoricalRequest) SetOrderId(v string) *HistoricalRequest {
	s.OrderId = &v
	return s
}

func (s *HistoricalRequest) SetRequestId(v string) *HistoricalRequest {
	s.RequestId = &v
	return s
}

func (s *HistoricalRequest) SetStation(v string) *HistoricalRequest {
	s.Station = &v
	return s
}

func (s *HistoricalRequest) SetPageSize(v int32) *HistoricalRequest {
	s.PageSize = &v
	return s
}

func (s *HistoricalRequest) SetStartTime(v string) *HistoricalRequest {
	s.StartTime = &v
	return s
}

func (s *HistoricalRequest) SetEndTime(v string) *HistoricalRequest {
	s.EndTime = &v
	return s
}

func (s *HistoricalRequest) SetPageNum(v int32) *HistoricalRequest {
	s.PageNum = &v
	return s
}

type HistoricalResponseBody struct {
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HistoricalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HistoricalResponseBody) GoString() string {
	return s.String()
}

func (s *HistoricalResponseBody) SetRt(v int64) *HistoricalResponseBody {
	s.Rt = &v
	return s
}

func (s *HistoricalResponseBody) SetMessage(v string) *HistoricalResponseBody {
	s.Message = &v
	return s
}

func (s *HistoricalResponseBody) SetRequestId(v string) *HistoricalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HistoricalResponseBody) SetData(v []map[string]interface{}) *HistoricalResponseBody {
	s.Data = v
	return s
}

func (s *HistoricalResponseBody) SetCode(v string) *HistoricalResponseBody {
	s.Code = &v
	return s
}

func (s *HistoricalResponseBody) SetSuccess(v bool) *HistoricalResponseBody {
	s.Success = &v
	return s
}

type HistoricalResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HistoricalResponse) String() string {
	return tea.Prettify(s)
}

func (s HistoricalResponse) GoString() string {
	return s.String()
}

func (s *HistoricalResponse) SetHeaders(v map[string]*string) *HistoricalResponse {
	s.Headers = v
	return s
}

func (s *HistoricalResponse) SetBody(v *HistoricalResponseBody) *HistoricalResponse {
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
	if !tea.BoolValue(util.IsUnset(tmpReq.ServiceParam)) {
		request.ServiceParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceParam, tea.String("ServiceParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendParam)) {
		request.ExtendParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendParam, tea.String("ExtendParam"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Execute"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(tmpReq.ServiceParam)) {
		request.ServiceParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceParam, tea.String("ServiceParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendParam)) {
		request.ExtendParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendParam, tea.String("ExtendParam"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WeathermonitorProvinceHourResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WeathermonitorProvinceHour"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) WeathermonitorWithOptions(request *WeathermonitorRequest, runtime *util.RuntimeOptions) (_result *WeathermonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WeathermonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Weathermonitor"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Weathermonitor(request *WeathermonitorRequest) (_result *WeathermonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeathermonitorResponse{}
	_body, _err := client.WeathermonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeatherforecastTimeWithOptions(request *WeatherforecastTimeRequest, runtime *util.RuntimeOptions) (_result *WeatherforecastTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WeatherforecastTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WeatherforecastTime"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WeatherforecastTime(request *WeatherforecastTimeRequest) (_result *WeatherforecastTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeatherforecastTimeResponse{}
	_body, _err := client.WeatherforecastTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StationDayWithOptions(request *StationDayRequest, runtime *util.RuntimeOptions) (_result *StationDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StationDayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StationDay"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StationDay(request *StationDayRequest) (_result *StationDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StationDayResponse{}
	_body, _err := client.StationDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeatherforecastWithOptions(request *WeatherforecastRequest, runtime *util.RuntimeOptions) (_result *WeatherforecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WeatherforecastResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Weatherforecast"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Weatherforecast(request *WeatherforecastRequest) (_result *WeatherforecastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeatherforecastResponse{}
	_body, _err := client.WeatherforecastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HistoricalWithOptions(request *HistoricalRequest, runtime *util.RuntimeOptions) (_result *HistoricalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HistoricalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("Historical"), tea.String("2021-09-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Historical(request *HistoricalRequest) (_result *HistoricalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HistoricalResponse{}
	_body, _err := client.HistoricalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
