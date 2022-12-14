// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDesktopPoolResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDesktopPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopPoolResponseBody) SetCode(v string) *CreateDesktopPoolResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDesktopPoolResponseBody) SetData(v string) *CreateDesktopPoolResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDesktopPoolResponseBody) SetHttpStatusCode(v int32) *CreateDesktopPoolResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDesktopPoolResponseBody) SetMessage(v string) *CreateDesktopPoolResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDesktopPoolResponseBody) SetRequestId(v string) *CreateDesktopPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopPoolResponseBody) SetSuccess(v bool) *CreateDesktopPoolResponseBody {
	s.Success = &v
	return s
}

type CreateDesktopPoolResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDesktopPoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopPoolResponse) SetHeaders(v map[string]*string) *CreateDesktopPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopPoolResponse) SetStatusCode(v int32) *CreateDesktopPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDesktopPoolResponse) SetBody(v *CreateDesktopPoolResponseBody) *CreateDesktopPoolResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("edsofficeservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDesktopPoolWithOptions(runtime *util.RuntimeOptions) (_result *CreateDesktopPoolResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CreateDesktopPool"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDesktopPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopPool() (_result *CreateDesktopPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopPoolResponse{}
	_body, _err := client.CreateDesktopPoolWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
