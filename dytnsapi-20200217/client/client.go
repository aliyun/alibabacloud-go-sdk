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

type DescribePhoneNumberStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s DescribePhoneNumberStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusRequest) SetOwnerId(v int64) *DescribePhoneNumberStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberStatusRequest) SetPhoneNumber(v string) *DescribePhoneNumberStatusRequest {
	s.PhoneNumber = &v
	return s
}

type DescribePhoneNumberStatusResponseBody struct {
	Code        *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PhoneStatus *DescribePhoneNumberStatusResponseBodyPhoneStatus `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty" type:"Struct"`
}

func (s DescribePhoneNumberStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponseBody) SetCode(v string) *DescribePhoneNumberStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetMessage(v string) *DescribePhoneNumberStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetRequestId(v string) *DescribePhoneNumberStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetPhoneStatus(v *DescribePhoneNumberStatusResponseBodyPhoneStatus) *DescribePhoneNumberStatusResponseBody {
	s.PhoneStatus = v
	return s
}

type DescribePhoneNumberStatusResponseBodyPhoneStatus struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SerialId *string `json:"SerialId,omitempty" xml:"SerialId,omitempty"`
	Carrier  *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetStatus(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Status = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetSerialId(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.SerialId = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetCarrier(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Carrier = &v
	return s
}

type DescribePhoneNumberStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePhoneNumberStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberStatusResponse) SetBody(v *DescribePhoneNumberStatusResponseBody) *DescribePhoneNumberStatusResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s DescribePhoneNumberAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest {
	s.PhoneNumber = &v
	return s
}

type DescribePhoneNumberAttributeResponseBody struct {
	Code                 *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
}

func (s DescribePhoneNumberAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBody) SetCode(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody {
	s.PhoneNumberAttribute = v
	return s
}

type DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetBasicCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetIsNumberPortability(v bool) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetNumberSegment(v int64) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCity(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetProvince(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Province = &v
	return s
}

type DescribePhoneNumberAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePhoneNumberAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetBody(v *DescribePhoneNumberAttributeResponseBody) *DescribePhoneNumberAttributeResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberResaleRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Since                *string `json:"Since,omitempty" xml:"Since,omitempty"`
}

func (s DescribePhoneNumberResaleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleRequest) SetOwnerId(v int64) *DescribePhoneNumberResaleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberResaleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberResaleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetPhoneNumber(v string) *DescribePhoneNumberResaleRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberResaleRequest) SetSince(v string) *DescribePhoneNumberResaleRequest {
	s.Since = &v
	return s
}

type DescribePhoneNumberResaleResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TwiceTelVerify *DescribePhoneNumberResaleResponseBodyTwiceTelVerify `json:"TwiceTelVerify,omitempty" xml:"TwiceTelVerify,omitempty" type:"Struct"`
}

func (s DescribePhoneNumberResaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponseBody) SetCode(v string) *DescribePhoneNumberResaleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetMessage(v string) *DescribePhoneNumberResaleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetRequestId(v string) *DescribePhoneNumberResaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBody) SetTwiceTelVerify(v *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) *DescribePhoneNumberResaleResponseBody {
	s.TwiceTelVerify = v
	return s
}

type DescribePhoneNumberResaleResponseBodyTwiceTelVerify struct {
	Carrier      *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	VerifyResult *int32  `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberResaleResponseBodyTwiceTelVerify) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponseBodyTwiceTelVerify) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) SetCarrier(v string) *DescribePhoneNumberResaleResponseBodyTwiceTelVerify {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberResaleResponseBodyTwiceTelVerify) SetVerifyResult(v int32) *DescribePhoneNumberResaleResponseBodyTwiceTelVerify {
	s.VerifyResult = &v
	return s
}

type DescribePhoneNumberResaleResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePhoneNumberResaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberResaleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberResaleResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberResaleResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberResaleResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberResaleResponse) SetBody(v *DescribePhoneNumberResaleResponseBody) *DescribePhoneNumberResaleResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dytnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribePhoneNumberStatusWithOptions(request *DescribePhoneNumberStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePhoneNumberStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePhoneNumberStatus"), tea.String("2020-02-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberStatus(request *DescribePhoneNumberStatusRequest) (_result *DescribePhoneNumberStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberStatusResponse{}
	_body, _err := client.DescribePhoneNumberStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttributeWithOptions(request *DescribePhoneNumberAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePhoneNumberAttribute"), tea.String("2020-02-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAttribute(request *DescribePhoneNumberAttributeRequest) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.DescribePhoneNumberAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberResaleWithOptions(request *DescribePhoneNumberResaleRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberResaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePhoneNumberResaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePhoneNumberResale"), tea.String("2020-02-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberResale(request *DescribePhoneNumberResaleRequest) (_result *DescribePhoneNumberResaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberResaleResponse{}
	_body, _err := client.DescribePhoneNumberResaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
