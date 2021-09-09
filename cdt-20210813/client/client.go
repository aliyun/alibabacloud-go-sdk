// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetCdtServiceStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetCdtServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCdtServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCdtServiceStatusRequest) SetOwnerId(v int64) *GetCdtServiceStatusRequest {
	s.OwnerId = &v
	return s
}

type GetCdtServiceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetCdtServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCdtServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCdtServiceStatusResponseBody) SetRequestId(v string) *GetCdtServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCdtServiceStatusResponseBody) SetEnabled(v bool) *GetCdtServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

type GetCdtServiceStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCdtServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCdtServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCdtServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCdtServiceStatusResponse) SetHeaders(v map[string]*string) *GetCdtServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCdtServiceStatusResponse) SetBody(v *GetCdtServiceStatusResponseBody) *GetCdtServiceStatusResponse {
	s.Body = v
	return s
}

type OpenCdtServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenCdtServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCdtServiceRequest) SetOwnerId(v int64) *OpenCdtServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenCdtServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenCdtServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCdtServiceResponseBody) SetRequestId(v string) *OpenCdtServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCdtServiceResponseBody) SetOrderId(v string) *OpenCdtServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenCdtServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenCdtServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCdtServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCdtServiceResponse) SetHeaders(v map[string]*string) *OpenCdtServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCdtServiceResponse) SetBody(v *OpenCdtServiceResponseBody) *OpenCdtServiceResponse {
	s.Body = v
	return s
}

type GetCdtCbServiceStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetCdtCbServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCdtCbServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCdtCbServiceStatusRequest) SetOwnerId(v int64) *GetCdtCbServiceStatusRequest {
	s.OwnerId = &v
	return s
}

type GetCdtCbServiceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetCdtCbServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCdtCbServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCdtCbServiceStatusResponseBody) SetRequestId(v string) *GetCdtCbServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCdtCbServiceStatusResponseBody) SetEnabled(v bool) *GetCdtCbServiceStatusResponseBody {
	s.Enabled = &v
	return s
}

type GetCdtCbServiceStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCdtCbServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCdtCbServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCdtCbServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCdtCbServiceStatusResponse) SetHeaders(v map[string]*string) *GetCdtCbServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCdtCbServiceStatusResponse) SetBody(v *GetCdtCbServiceStatusResponseBody) *GetCdtCbServiceStatusResponse {
	s.Body = v
	return s
}

type OpenCdtCbServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenCdtCbServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtCbServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCdtCbServiceRequest) SetOwnerId(v int64) *OpenCdtCbServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenCdtCbServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenCdtCbServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtCbServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCdtCbServiceResponseBody) SetRequestId(v string) *OpenCdtCbServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCdtCbServiceResponseBody) SetOrderId(v string) *OpenCdtCbServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenCdtCbServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenCdtCbServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCdtCbServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCdtCbServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCdtCbServiceResponse) SetHeaders(v map[string]*string) *OpenCdtCbServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCdtCbServiceResponse) SetBody(v *OpenCdtCbServiceResponseBody) *OpenCdtCbServiceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cdt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetCdtServiceStatusWithOptions(request *GetCdtServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetCdtServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCdtServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCdtServiceStatus"), tea.String("2021-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCdtServiceStatus(request *GetCdtServiceStatusRequest) (_result *GetCdtServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCdtServiceStatusResponse{}
	_body, _err := client.GetCdtServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCdtServiceWithOptions(request *OpenCdtServiceRequest, runtime *util.RuntimeOptions) (_result *OpenCdtServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenCdtServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenCdtService"), tea.String("2021-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCdtService(request *OpenCdtServiceRequest) (_result *OpenCdtServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCdtServiceResponse{}
	_body, _err := client.OpenCdtServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCdtCbServiceStatusWithOptions(request *GetCdtCbServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetCdtCbServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCdtCbServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCdtCbServiceStatus"), tea.String("2021-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCdtCbServiceStatus(request *GetCdtCbServiceStatusRequest) (_result *GetCdtCbServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCdtCbServiceStatusResponse{}
	_body, _err := client.GetCdtCbServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCdtCbServiceWithOptions(request *OpenCdtCbServiceRequest, runtime *util.RuntimeOptions) (_result *OpenCdtCbServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenCdtCbServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenCdtCbService"), tea.String("2021-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCdtCbService(request *OpenCdtCbServiceRequest) (_result *OpenCdtCbServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCdtCbServiceResponse{}
	_body, _err := client.OpenCdtCbServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
