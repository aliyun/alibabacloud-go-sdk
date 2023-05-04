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

type VerifyCaptchaRequest struct {
	CaptchaVerifyParam *string `json:"CaptchaVerifyParam,omitempty" xml:"CaptchaVerifyParam,omitempty"`
}

func (s VerifyCaptchaRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyCaptchaRequest) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaRequest) SetCaptchaVerifyParam(v string) *VerifyCaptchaRequest {
	s.CaptchaVerifyParam = &v
	return s
}

type VerifyCaptchaResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *VerifyCaptchaResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyCaptchaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponseBody) SetCode(v string) *VerifyCaptchaResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetMessage(v string) *VerifyCaptchaResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetRequestId(v string) *VerifyCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCaptchaResponseBody) SetResult(v *VerifyCaptchaResponseBodyResult) *VerifyCaptchaResponseBody {
	s.Result = v
	return s
}

func (s *VerifyCaptchaResponseBody) SetSuccess(v bool) *VerifyCaptchaResponseBody {
	s.Success = &v
	return s
}

type VerifyCaptchaResponseBodyResult struct {
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyCaptchaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyCaptchaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponseBodyResult) SetVerifyResult(v bool) *VerifyCaptchaResponseBodyResult {
	s.VerifyResult = &v
	return s
}

type VerifyCaptchaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyCaptchaResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyCaptchaResponse) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponse) SetHeaders(v map[string]*string) *VerifyCaptchaResponse {
	s.Headers = v
	return s
}

func (s *VerifyCaptchaResponse) SetStatusCode(v int32) *VerifyCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCaptchaResponse) SetBody(v *VerifyCaptchaResponseBody) *VerifyCaptchaResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("captcha"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) VerifyCaptchaWithOptions(request *VerifyCaptchaRequest, runtime *util.RuntimeOptions) (_result *VerifyCaptchaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaptchaVerifyParam)) {
		query["CaptchaVerifyParam"] = request.CaptchaVerifyParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyCaptcha"),
		Version:     tea.String("2023-03-05"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyCaptchaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyCaptcha(request *VerifyCaptchaRequest) (_result *VerifyCaptchaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyCaptchaResponse{}
	_body, _err := client.VerifyCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
