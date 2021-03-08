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

type ActivateLicenseRequest struct {
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	LicensePublisher *string `json:"LicensePublisher,omitempty" xml:"LicensePublisher,omitempty"`
	LicenseCode      *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) SetBizType(v string) *ActivateLicenseRequest {
	s.BizType = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicensePublisher(v string) *ActivateLicenseRequest {
	s.LicensePublisher = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

type ActivateLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetData(v bool) *ActivateLicenseResponseBody {
	s.Data = &v
	return s
}

type ActivateLicenseResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mseap"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *util.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ActivateLicense"), tea.String("2021-01-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
