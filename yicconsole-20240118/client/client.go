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

type LxPopCmd struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s LxPopCmd) String() string {
	return tea.Prettify(s)
}

func (s LxPopCmd) GoString() string {
	return s.String()
}

func (s *LxPopCmd) SetData(v string) *LxPopCmd {
	s.Data = &v
	return s
}

type BillingProcessMessageRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BillingProcessMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s BillingProcessMessageRequest) GoString() string {
	return s.String()
}

func (s *BillingProcessMessageRequest) SetBody(v *LxPopCmd) *BillingProcessMessageRequest {
	s.Body = v
	return s
}

type BillingProcessMessageResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Synchro *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BillingProcessMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillingProcessMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BillingProcessMessageResponseBody) SetMessage(v string) *BillingProcessMessageResponseBody {
	s.Message = &v
	return s
}

func (s *BillingProcessMessageResponseBody) SetSynchro(v bool) *BillingProcessMessageResponseBody {
	s.Synchro = &v
	return s
}

func (s *BillingProcessMessageResponseBody) SetCode(v string) *BillingProcessMessageResponseBody {
	s.Code = &v
	return s
}

func (s *BillingProcessMessageResponseBody) SetData(v string) *BillingProcessMessageResponseBody {
	s.Data = &v
	return s
}

func (s *BillingProcessMessageResponseBody) SetRequestId(v string) *BillingProcessMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BillingProcessMessageResponseBody) SetSuccess(v bool) *BillingProcessMessageResponseBody {
	s.Success = &v
	return s
}

type BillingProcessMessageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BillingProcessMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillingProcessMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s BillingProcessMessageResponse) GoString() string {
	return s.String()
}

func (s *BillingProcessMessageResponse) SetHeaders(v map[string]*string) *BillingProcessMessageResponse {
	s.Headers = v
	return s
}

func (s *BillingProcessMessageResponse) SetStatusCode(v int32) *BillingProcessMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BillingProcessMessageResponse) SetBody(v *BillingProcessMessageResponseBody) *BillingProcessMessageResponse {
	s.Body = v
	return s
}

type CheckPayOrderRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckPayOrderRequest) GoString() string {
	return s.String()
}

func (s *CheckPayOrderRequest) SetBody(v *LxPopCmd) *CheckPayOrderRequest {
	s.Body = v
	return s
}

type CheckPayOrderResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Synchro   *bool   `json:"synchro,omitempty" xml:"synchro,omitempty"`
}

func (s CheckPayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CheckPayOrderResponseBody) SetCode(v string) *CheckPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CheckPayOrderResponseBody) SetData(v string) *CheckPayOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CheckPayOrderResponseBody) SetMessage(v string) *CheckPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CheckPayOrderResponseBody) SetRequestId(v string) *CheckPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckPayOrderResponseBody) SetSuccess(v bool) *CheckPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CheckPayOrderResponseBody) SetSynchro(v bool) *CheckPayOrderResponseBody {
	s.Synchro = &v
	return s
}

type CheckPayOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckPayOrderResponse) GoString() string {
	return s.String()
}

func (s *CheckPayOrderResponse) SetHeaders(v map[string]*string) *CheckPayOrderResponse {
	s.Headers = v
	return s
}

func (s *CheckPayOrderResponse) SetStatusCode(v int32) *CheckPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckPayOrderResponse) SetBody(v *CheckPayOrderResponseBody) *CheckPayOrderResponse {
	s.Body = v
	return s
}

type CheckRefundRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckRefundRequest) GoString() string {
	return s.String()
}

func (s *CheckRefundRequest) SetBody(v *LxPopCmd) *CheckRefundRequest {
	s.Body = v
	return s
}

type CheckRefundResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Synchro *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckRefundResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRefundResponseBody) SetMessage(v string) *CheckRefundResponseBody {
	s.Message = &v
	return s
}

func (s *CheckRefundResponseBody) SetSynchro(v bool) *CheckRefundResponseBody {
	s.Synchro = &v
	return s
}

func (s *CheckRefundResponseBody) SetCode(v string) *CheckRefundResponseBody {
	s.Code = &v
	return s
}

func (s *CheckRefundResponseBody) SetData(v string) *CheckRefundResponseBody {
	s.Data = &v
	return s
}

func (s *CheckRefundResponseBody) SetRequestId(v string) *CheckRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRefundResponseBody) SetSuccess(v bool) *CheckRefundResponseBody {
	s.Success = &v
	return s
}

type CheckRefundResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckRefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckRefundResponse) GoString() string {
	return s.String()
}

func (s *CheckRefundResponse) SetHeaders(v map[string]*string) *CheckRefundResponse {
	s.Headers = v
	return s
}

func (s *CheckRefundResponse) SetStatusCode(v int32) *CheckRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRefundResponse) SetBody(v *CheckRefundResponseBody) *CheckRefundResponse {
	s.Body = v
	return s
}

type PayOrderRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s PayOrderRequest) GoString() string {
	return s.String()
}

func (s *PayOrderRequest) SetBody(v *LxPopCmd) *PayOrderRequest {
	s.Body = v
	return s
}

type PayOrderResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Synchro *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *PayOrderResponseBody) SetMessage(v string) *PayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *PayOrderResponseBody) SetSynchro(v bool) *PayOrderResponseBody {
	s.Synchro = &v
	return s
}

func (s *PayOrderResponseBody) SetCode(v string) *PayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *PayOrderResponseBody) SetData(v string) *PayOrderResponseBody {
	s.Data = &v
	return s
}

func (s *PayOrderResponseBody) SetRequestId(v string) *PayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *PayOrderResponseBody) SetSuccess(v bool) *PayOrderResponseBody {
	s.Success = &v
	return s
}

type PayOrderResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s PayOrderResponse) GoString() string {
	return s.String()
}

func (s *PayOrderResponse) SetHeaders(v map[string]*string) *PayOrderResponse {
	s.Headers = v
	return s
}

func (s *PayOrderResponse) SetStatusCode(v int32) *PayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *PayOrderResponse) SetBody(v *PayOrderResponseBody) *PayOrderResponse {
	s.Body = v
	return s
}

type PrepaidCeaseRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepaidCeaseRequest) String() string {
	return tea.Prettify(s)
}

func (s PrepaidCeaseRequest) GoString() string {
	return s.String()
}

func (s *PrepaidCeaseRequest) SetBody(v *LxPopCmd) *PrepaidCeaseRequest {
	s.Body = v
	return s
}

type PrepaidCeaseResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Synchro *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PrepaidCeaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PrepaidCeaseResponseBody) GoString() string {
	return s.String()
}

func (s *PrepaidCeaseResponseBody) SetMessage(v string) *PrepaidCeaseResponseBody {
	s.Message = &v
	return s
}

func (s *PrepaidCeaseResponseBody) SetSynchro(v bool) *PrepaidCeaseResponseBody {
	s.Synchro = &v
	return s
}

func (s *PrepaidCeaseResponseBody) SetCode(v string) *PrepaidCeaseResponseBody {
	s.Code = &v
	return s
}

func (s *PrepaidCeaseResponseBody) SetData(v string) *PrepaidCeaseResponseBody {
	s.Data = &v
	return s
}

func (s *PrepaidCeaseResponseBody) SetRequestId(v string) *PrepaidCeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepaidCeaseResponseBody) SetSuccess(v bool) *PrepaidCeaseResponseBody {
	s.Success = &v
	return s
}

type PrepaidCeaseResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PrepaidCeaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PrepaidCeaseResponse) String() string {
	return tea.Prettify(s)
}

func (s PrepaidCeaseResponse) GoString() string {
	return s.String()
}

func (s *PrepaidCeaseResponse) SetHeaders(v map[string]*string) *PrepaidCeaseResponse {
	s.Headers = v
	return s
}

func (s *PrepaidCeaseResponse) SetStatusCode(v int32) *PrepaidCeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepaidCeaseResponse) SetBody(v *PrepaidCeaseResponseBody) *PrepaidCeaseResponse {
	s.Body = v
	return s
}

type RefundRequest struct {
	Body *LxPopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundRequest) GoString() string {
	return s.String()
}

func (s *RefundRequest) SetBody(v *LxPopCmd) *RefundRequest {
	s.Body = v
	return s
}

type RefundResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Synchro *bool   `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundResponseBody) GoString() string {
	return s.String()
}

func (s *RefundResponseBody) SetMessage(v string) *RefundResponseBody {
	s.Message = &v
	return s
}

func (s *RefundResponseBody) SetSynchro(v bool) *RefundResponseBody {
	s.Synchro = &v
	return s
}

func (s *RefundResponseBody) SetCode(v string) *RefundResponseBody {
	s.Code = &v
	return s
}

func (s *RefundResponseBody) SetData(v string) *RefundResponseBody {
	s.Data = &v
	return s
}

func (s *RefundResponseBody) SetRequestId(v string) *RefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundResponseBody) SetSuccess(v bool) *RefundResponseBody {
	s.Success = &v
	return s
}

type RefundResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundResponse) GoString() string {
	return s.String()
}

func (s *RefundResponse) SetHeaders(v map[string]*string) *RefundResponse {
	s.Headers = v
	return s
}

func (s *RefundResponse) SetStatusCode(v int32) *RefundResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundResponse) SetBody(v *RefundResponseBody) *RefundResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("yicconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BillingProcessMessageWithOptions(request *BillingProcessMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BillingProcessMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("BillingProcessMessage"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/lifecycle"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BillingProcessMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillingProcessMessage(request *BillingProcessMessageRequest) (_result *BillingProcessMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BillingProcessMessageResponse{}
	_body, _err := client.BillingProcessMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckPayOrderWithOptions(request *CheckPayOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckPayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckPayOrder"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/verify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckPayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckPayOrder(request *CheckPayOrderRequest) (_result *CheckPayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckPayOrderResponse{}
	_body, _err := client.CheckPayOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckRefundWithOptions(request *CheckRefundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckRefund"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/check-refund"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckRefund(request *CheckRefundRequest) (_result *CheckRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckRefundResponse{}
	_body, _err := client.CheckRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PayOrderWithOptions(request *PayOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PayOrder"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/pay"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PayOrder(request *PayOrderRequest) (_result *PayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PayOrderResponse{}
	_body, _err := client.PayOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PrepaidCeaseWithOptions(request *PrepaidCeaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PrepaidCeaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PrepaidCease"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/prepaid-cease"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PrepaidCeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PrepaidCease(request *PrepaidCeaseRequest) (_result *PrepaidCeaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PrepaidCeaseResponse{}
	_body, _err := client.PrepaidCeaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundWithOptions(request *RefundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("Refund"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/billing/commands/refund"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Refund(request *RefundRequest) (_result *RefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefundResponse{}
	_body, _err := client.RefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
