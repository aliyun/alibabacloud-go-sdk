// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetOpenStatusResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"gmtModified":"2021-07-27T04:00:00.000+00:00","mpStatus":1,"id":11,"pk":"1084126944995576","gmtCreate":"2021-07-27T04:00:00.000+00:00","parentPk":"1084126944995576"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EB-FCAC-1B78-BBB8-500ED951E9EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenStatusResponseBody) SetCode(v int32) *GetOpenStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenStatusResponseBody) SetData(v map[string]interface{}) *GetOpenStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetOpenStatusResponseBody) SetMessage(v string) *GetOpenStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenStatusResponseBody) SetRequestId(v string) *GetOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenStatusResponseBody) SetSuccess(v bool) *GetOpenStatusResponseBody {
	s.Success = &v
	return s
}

type GetOpenStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOpenStatusResponse) SetHeaders(v map[string]*string) *GetOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOpenStatusResponse) SetStatusCode(v int32) *GetOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenStatusResponse) SetBody(v *GetOpenStatusResponseBody) *GetOpenStatusResponse {
	s.Body = v
	return s
}

type GetOrderInfoRequest struct {
	ListReleased *bool `json:"ListReleased,omitempty" xml:"ListReleased,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MP
	RelService *string `json:"RelService,omitempty" xml:"RelService,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOrderInfoRequest) SetListReleased(v bool) *GetOrderInfoRequest {
	s.ListReleased = &v
	return s
}

func (s *GetOrderInfoRequest) SetRelService(v string) *GetOrderInfoRequest {
	s.RelService = &v
	return s
}

func (s *GetOrderInfoRequest) SetResourceType(v int32) *GetOrderInfoRequest {
	s.ResourceType = &v
	return s
}

type GetOrderInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"licenseKey":"eems7ri3b1u5nui*****","currentConcurrency":2,"instanceId":"opt_mplicense_public_cn-****","totalDays":124,"currentDays":103}
	Data []*GetOrderInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5D6653C5-CA2B-14EC-9CF0-50AA0FF49C31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderInfoResponseBody) SetCode(v string) *GetOrderInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrderInfoResponseBody) SetData(v []*GetOrderInfoResponseBodyData) *GetOrderInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderInfoResponseBody) SetMessage(v string) *GetOrderInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderInfoResponseBody) SetRequestId(v string) *GetOrderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderInfoResponseBody) SetSuccess(v bool) *GetOrderInfoResponseBody {
	s.Success = &v
	return s
}

type GetOrderInfoResponseBodyData struct {
	BizType            *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CurrentConcurrency *int32  `json:"currentConcurrency,omitempty" xml:"currentConcurrency,omitempty"`
	CurrentDays        *int32  `json:"currentDays,omitempty" xml:"currentDays,omitempty"`
	InstanceId         *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsReleased         *bool   `json:"isReleased,omitempty" xml:"isReleased,omitempty"`
	LicenseKey         *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
	Remark             *string `json:"remark,omitempty" xml:"remark,omitempty"`
	TotalDays          *int32  `json:"totalDays,omitempty" xml:"totalDays,omitempty"`
}

func (s GetOrderInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrderInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrderInfoResponseBodyData) SetBizType(v string) *GetOrderInfoResponseBodyData {
	s.BizType = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetCurrentConcurrency(v int32) *GetOrderInfoResponseBodyData {
	s.CurrentConcurrency = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetCurrentDays(v int32) *GetOrderInfoResponseBodyData {
	s.CurrentDays = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetInstanceId(v string) *GetOrderInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetIsReleased(v bool) *GetOrderInfoResponseBodyData {
	s.IsReleased = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetLicenseKey(v string) *GetOrderInfoResponseBodyData {
	s.LicenseKey = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetRemark(v string) *GetOrderInfoResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetOrderInfoResponseBodyData) SetTotalDays(v int32) *GetOrderInfoResponseBodyData {
	s.TotalDays = &v
	return s
}

type GetOrderInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOrderInfoResponse) SetHeaders(v map[string]*string) *GetOrderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOrderInfoResponse) SetStatusCode(v int32) *GetOrderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderInfoResponse) SetBody(v *GetOrderInfoResponseBody) *GetOrderInfoResponse {
	s.Body = v
	return s
}

type GetOrderUsageRequest struct {
	// This parameter is required.
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MP
	RelService *string `json:"RelService,omitempty" xml:"RelService,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1
	TimeRange *int32 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s GetOrderUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderUsageRequest) GoString() string {
	return s.String()
}

func (s *GetOrderUsageRequest) SetLicenseKey(v string) *GetOrderUsageRequest {
	s.LicenseKey = &v
	return s
}

func (s *GetOrderUsageRequest) SetRelService(v string) *GetOrderUsageRequest {
	s.RelService = &v
	return s
}

func (s *GetOrderUsageRequest) SetResourceType(v int32) *GetOrderUsageRequest {
	s.ResourceType = &v
	return s
}

func (s *GetOrderUsageRequest) SetTimeRange(v int32) *GetOrderUsageRequest {
	s.TimeRange = &v
	return s
}

type GetOrderUsageResponseBody struct {
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5D6653C5-CA2B-14EC-9CF0-50AA0FF49C31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderUsageResponseBody) SetData(v []map[string]interface{}) *GetOrderUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderUsageResponseBody) SetMessage(v string) *GetOrderUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderUsageResponseBody) SetRequestId(v string) *GetOrderUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderUsageResponseBody) SetSuccess(v bool) *GetOrderUsageResponseBody {
	s.Success = &v
	return s
}

type GetOrderUsageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderUsageResponse) GoString() string {
	return s.String()
}

func (s *GetOrderUsageResponse) SetHeaders(v map[string]*string) *GetOrderUsageResponse {
	s.Headers = v
	return s
}

func (s *GetOrderUsageResponse) SetStatusCode(v int32) *GetOrderUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderUsageResponse) SetBody(v *GetOrderUsageResponseBody) *GetOrderUsageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("opt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - GetOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenStatusResponse
func (client *Client) GetOpenStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetOpenStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetOpenStatus"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOpenStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOpenStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @return GetOpenStatusResponse
func (client *Client) GetOpenStatus() (_result *GetOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenStatusResponse{}
	_body, _err := client.GetOpenStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取在线license列表
//
// @param request - GetOrderInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderInfoResponse
func (client *Client) GetOrderInfoWithOptions(request *GetOrderInfoRequest, runtime *util.RuntimeOptions) (_result *GetOrderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListReleased)) {
		query["ListReleased"] = request.ListReleased
	}

	if !tea.BoolValue(util.IsUnset(request.RelService)) {
		query["RelService"] = request.RelService
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrderInfo"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOrderInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOrderInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取在线license列表
//
// @param request - GetOrderInfoRequest
//
// @return GetOrderInfoResponse
func (client *Client) GetOrderInfo(request *GetOrderInfoRequest) (_result *GetOrderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderInfoResponse{}
	_body, _err := client.GetOrderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetOrderUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderUsageResponse
func (client *Client) GetOrderUsageWithOptions(request *GetOrderUsageRequest, runtime *util.RuntimeOptions) (_result *GetOrderUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LicenseKey)) {
		query["LicenseKey"] = request.LicenseKey
	}

	if !tea.BoolValue(util.IsUnset(request.RelService)) {
		query["RelService"] = request.RelService
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrderUsage"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOrderUsageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOrderUsageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetOrderUsageRequest
//
// @return GetOrderUsageResponse
func (client *Client) GetOrderUsage(request *GetOrderUsageRequest) (_result *GetOrderUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderUsageResponse{}
	_body, _err := client.GetOrderUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
