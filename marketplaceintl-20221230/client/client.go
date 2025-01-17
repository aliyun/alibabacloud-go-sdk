// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PushMeteringDataRequest struct {
	// example:
	//
	// 2023-01-11 10:31:00
	GmtCreate    *string                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MeteringData []*PushMeteringDataRequestMeteringData `json:"MeteringData,omitempty" xml:"MeteringData,omitempty" type:"Repeated"`
}

func (s PushMeteringDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) SetGmtCreate(v string) *PushMeteringDataRequest {
	s.GmtCreate = &v
	return s
}

func (s *PushMeteringDataRequest) SetMeteringData(v []*PushMeteringDataRequestMeteringData) *PushMeteringDataRequest {
	s.MeteringData = v
	return s
}

type PushMeteringDataRequestMeteringData struct {
	// example:
	//
	// 1666854480406
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// gtm-cn-20p314k5h05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test001
	MeteringAssist *string `json:"MeteringAssist,omitempty" xml:"MeteringAssist,omitempty"`
	// example:
	//
	// {"VirtualCpu":10}
	MeteringEntity *string `json:"MeteringEntity,omitempty" xml:"MeteringEntity,omitempty"`
	// example:
	//
	// 1662284820000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PushMeteringDataRequestMeteringData) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequestMeteringData) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequestMeteringData) SetEndTime(v int64) *PushMeteringDataRequestMeteringData {
	s.EndTime = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetInstanceId(v string) *PushMeteringDataRequestMeteringData {
	s.InstanceId = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringAssist(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringAssist = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringEntity(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringEntity = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetStartTime(v int64) *PushMeteringDataRequestMeteringData {
	s.StartTime = &v
	return s
}

type PushMeteringDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// parameter \\"service\\" can not be blank.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// false
	ForceFatal *bool `json:"ForceFatal,omitempty" xml:"ForceFatal,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) SetCode(v string) *PushMeteringDataResponseBody {
	s.Code = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetDynamicMessage(v string) *PushMeteringDataResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetForceFatal(v bool) *PushMeteringDataResponseBody {
	s.ForceFatal = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetMessage(v string) *PushMeteringDataResponseBody {
	s.Message = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetResult(v bool) *PushMeteringDataResponseBody {
	s.Result = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetSuccess(v bool) *PushMeteringDataResponseBody {
	s.Success = &v
	return s
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMeteringDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("marketplaceintl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 国际云市场推送计量数据
//
// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *util.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtCreate)) {
		body["GmtCreate"] = request.GmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.MeteringData)) {
		body["MeteringData"] = request.MeteringData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMeteringData"),
		Version:     tea.String("2022-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际云市场推送计量数据
//
// @param request - PushMeteringDataRequest
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
