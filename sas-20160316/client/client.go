// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type IsSasServiceOpeningRequest struct {
}

func (s IsSasServiceOpeningRequest) String() string {
	return tea.Prettify(s)
}

func (s IsSasServiceOpeningRequest) GoString() string {
	return s.String()
}

type IsSasServiceOpeningResponse struct {
	Data    *bool   `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty" require:"true"`
}

func (s IsSasServiceOpeningResponse) String() string {
	return tea.Prettify(s)
}

func (s IsSasServiceOpeningResponse) GoString() string {
	return s.String()
}

func (s *IsSasServiceOpeningResponse) SetData(v bool) *IsSasServiceOpeningResponse {
	s.Data = &v
	return s
}

func (s *IsSasServiceOpeningResponse) SetCode(v string) *IsSasServiceOpeningResponse {
	s.Code = &v
	return s
}

func (s *IsSasServiceOpeningResponse) SetMessage(v string) *IsSasServiceOpeningResponse {
	s.Message = &v
	return s
}

func (s *IsSasServiceOpeningResponse) SetSuccess(v bool) *IsSasServiceOpeningResponse {
	s.Success = &v
	return s
}

type GetEventsCountRequest struct {
}

func (s GetEventsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEventsCountRequest) GoString() string {
	return s.String()
}

type GetEventsCountResponse struct {
	Code       *string                           `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message    *string                           `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Success    *bool                             `json:"success,omitempty" xml:"success,omitempty" require:"true"`
	EventCount *GetEventsCountResponseEventCount `json:"EventCount,omitempty" xml:"EventCount,omitempty" require:"true" type:"Struct"`
}

func (s GetEventsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventsCountResponse) GoString() string {
	return s.String()
}

func (s *GetEventsCountResponse) SetCode(v string) *GetEventsCountResponse {
	s.Code = &v
	return s
}

func (s *GetEventsCountResponse) SetMessage(v string) *GetEventsCountResponse {
	s.Message = &v
	return s
}

func (s *GetEventsCountResponse) SetSuccess(v bool) *GetEventsCountResponse {
	s.Success = &v
	return s
}

func (s *GetEventsCountResponse) SetEventCount(v *GetEventsCountResponseEventCount) *GetEventsCountResponse {
	s.EventCount = v
	return s
}

type GetEventsCountResponseEventCount struct {
	Event  *int64 `json:"event,omitempty" xml:"event,omitempty" require:"true"`
	Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
	Vul    *int64 `json:"vul,omitempty" xml:"vul,omitempty" require:"true"`
}

func (s GetEventsCountResponseEventCount) String() string {
	return tea.Prettify(s)
}

func (s GetEventsCountResponseEventCount) GoString() string {
	return s.String()
}

func (s *GetEventsCountResponseEventCount) SetEvent(v int64) *GetEventsCountResponseEventCount {
	s.Event = &v
	return s
}

func (s *GetEventsCountResponseEventCount) SetAttack(v int64) *GetEventsCountResponseEventCount {
	s.Attack = &v
	return s
}

func (s *GetEventsCountResponseEventCount) SetVul(v int64) *GetEventsCountResponseEventCount {
	s.Vul = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":         tea.String("tds.aliyuncs.com"),
		"ap-southeast-3":      tea.String("tds.ap-southeast-3.aliyuncs.com"),
		"cn-shanghai-et2-b01": tea.String("tds.cn-shanghai-et2-b01.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) IsSasServiceOpeningWithOptions(request *IsSasServiceOpeningRequest, runtime *util.RuntimeOptions) (_result *IsSasServiceOpeningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &IsSasServiceOpeningResponse{}
	_body, _err := client.DoRequest(tea.String("IsSasServiceOpening"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-03-16"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsSasServiceOpening(request *IsSasServiceOpeningRequest) (_result *IsSasServiceOpeningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsSasServiceOpeningResponse{}
	_body, _err := client.IsSasServiceOpeningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEventsCountWithOptions(request *GetEventsCountRequest, runtime *util.RuntimeOptions) (_result *GetEventsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetEventsCountResponse{}
	_body, _err := client.DoRequest(tea.String("GetEventsCount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-03-16"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEventsCount(request *GetEventsCountRequest) (_result *GetEventsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEventsCountResponse{}
	_body, _err := client.GetEventsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
