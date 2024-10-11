// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type NlpAddressNormalizationRequest struct {
	CityStdManual *string `json:"CityStdManual,omitempty" xml:"CityStdManual,omitempty"`
	// example:
	//
	// BACOLOR
	CityStr           *string `json:"CityStr,omitempty" xml:"CityStr,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DistrictStdManual *string `json:"DistrictStdManual,omitempty" xml:"DistrictStdManual,omitempty"`
	// example:
	//
	// 10671
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProvinceStdManual *string `json:"ProvinceStdManual,omitempty" xml:"ProvinceStdManual,omitempty"`
	// example:
	//
	// PAMPANGA
	ProvinceStr *string `json:"ProvinceStr,omitempty" xml:"ProvinceStr,omitempty"`
	// example:
	//
	// 350 Magliman, Bacolor, Pampanga AQUA MALTA Refilling Station
	QueryStr *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
}

func (s NlpAddressNormalizationRequest) String() string {
	return tea.Prettify(s)
}

func (s NlpAddressNormalizationRequest) GoString() string {
	return s.String()
}

func (s *NlpAddressNormalizationRequest) SetCityStdManual(v string) *NlpAddressNormalizationRequest {
	s.CityStdManual = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetCityStr(v string) *NlpAddressNormalizationRequest {
	s.CityStr = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetClientToken(v string) *NlpAddressNormalizationRequest {
	s.ClientToken = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetDistrictStdManual(v string) *NlpAddressNormalizationRequest {
	s.DistrictStdManual = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetInstanceId(v string) *NlpAddressNormalizationRequest {
	s.InstanceId = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetProvinceStdManual(v string) *NlpAddressNormalizationRequest {
	s.ProvinceStdManual = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetProvinceStr(v string) *NlpAddressNormalizationRequest {
	s.ProvinceStr = &v
	return s
}

func (s *NlpAddressNormalizationRequest) SetQueryStr(v string) *NlpAddressNormalizationRequest {
	s.QueryStr = &v
	return s
}

type NlpAddressNormalizationResponseBody struct {
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *NlpAddressNormalizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 99540D1A-9112-56E5-8DCC-1A2603F4C500
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1701051540592
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s NlpAddressNormalizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NlpAddressNormalizationResponseBody) GoString() string {
	return s.String()
}

func (s *NlpAddressNormalizationResponseBody) SetCode(v string) *NlpAddressNormalizationResponseBody {
	s.Code = &v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetData(v *NlpAddressNormalizationResponseBodyData) *NlpAddressNormalizationResponseBody {
	s.Data = v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetMessage(v string) *NlpAddressNormalizationResponseBody {
	s.Message = &v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetRequestId(v string) *NlpAddressNormalizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetResult(v string) *NlpAddressNormalizationResponseBody {
	s.Result = &v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetSuccess(v bool) *NlpAddressNormalizationResponseBody {
	s.Success = &v
	return s
}

func (s *NlpAddressNormalizationResponseBody) SetTimestamp(v string) *NlpAddressNormalizationResponseBody {
	s.Timestamp = &v
	return s
}

type NlpAddressNormalizationResponseBodyData struct {
	CityStd     *string   `json:"CityStd,omitempty" xml:"CityStd,omitempty"`
	DistrictStd *string   `json:"DistrictStd,omitempty" xml:"DistrictStd,omitempty"`
	ProvinceStd *string   `json:"ProvinceStd,omitempty" xml:"ProvinceStd,omitempty"`
	Results     []*string `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	StatusEs    *string   `json:"StatusEs,omitempty" xml:"StatusEs,omitempty"`
}

func (s NlpAddressNormalizationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s NlpAddressNormalizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *NlpAddressNormalizationResponseBodyData) SetCityStd(v string) *NlpAddressNormalizationResponseBodyData {
	s.CityStd = &v
	return s
}

func (s *NlpAddressNormalizationResponseBodyData) SetDistrictStd(v string) *NlpAddressNormalizationResponseBodyData {
	s.DistrictStd = &v
	return s
}

func (s *NlpAddressNormalizationResponseBodyData) SetProvinceStd(v string) *NlpAddressNormalizationResponseBodyData {
	s.ProvinceStd = &v
	return s
}

func (s *NlpAddressNormalizationResponseBodyData) SetResults(v []*string) *NlpAddressNormalizationResponseBodyData {
	s.Results = v
	return s
}

func (s *NlpAddressNormalizationResponseBodyData) SetStatusEs(v string) *NlpAddressNormalizationResponseBodyData {
	s.StatusEs = &v
	return s
}

type NlpAddressNormalizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NlpAddressNormalizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NlpAddressNormalizationResponse) String() string {
	return tea.Prettify(s)
}

func (s NlpAddressNormalizationResponse) GoString() string {
	return s.String()
}

func (s *NlpAddressNormalizationResponse) SetHeaders(v map[string]*string) *NlpAddressNormalizationResponse {
	s.Headers = v
	return s
}

func (s *NlpAddressNormalizationResponse) SetStatusCode(v int32) *NlpAddressNormalizationResponse {
	s.StatusCode = &v
	return s
}

func (s *NlpAddressNormalizationResponse) SetBody(v *NlpAddressNormalizationResponseBody) *NlpAddressNormalizationResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("superappnlp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 地址解析
//
// @param request - NlpAddressNormalizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NlpAddressNormalizationResponse
func (client *Client) NlpAddressNormalizationWithOptions(request *NlpAddressNormalizationRequest, runtime *util.RuntimeOptions) (_result *NlpAddressNormalizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CityStdManual)) {
		query["CityStdManual"] = request.CityStdManual
	}

	if !tea.BoolValue(util.IsUnset(request.CityStr)) {
		query["CityStr"] = request.CityStr
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictStdManual)) {
		query["DistrictStdManual"] = request.DistrictStdManual
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvinceStdManual)) {
		query["ProvinceStdManual"] = request.ProvinceStdManual
	}

	if !tea.BoolValue(util.IsUnset(request.ProvinceStr)) {
		query["ProvinceStr"] = request.ProvinceStr
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStr)) {
		query["QueryStr"] = request.QueryStr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NlpAddressNormalization"),
		Version:     tea.String("2024-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NlpAddressNormalizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 地址解析
//
// @param request - NlpAddressNormalizationRequest
//
// @return NlpAddressNormalizationResponse
func (client *Client) NlpAddressNormalization(request *NlpAddressNormalizationRequest) (_result *NlpAddressNormalizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NlpAddressNormalizationResponse{}
	_body, _err := client.NlpAddressNormalizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
