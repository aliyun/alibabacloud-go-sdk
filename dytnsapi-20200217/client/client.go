// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeEmptyNumberDetectRequest struct {
	EncryptType          *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeEmptyNumberDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectRequest) SetEncryptType(v string) *DescribeEmptyNumberDetectRequest {
	s.EncryptType = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetOwnerId(v int64) *DescribeEmptyNumberDetectRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetPhone(v string) *DescribeEmptyNumberDetectRequest {
	s.Phone = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetResourceOwnerAccount(v string) *DescribeEmptyNumberDetectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEmptyNumberDetectRequest) SetResourceOwnerId(v int64) *DescribeEmptyNumberDetectRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeEmptyNumberDetectResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeEmptyNumberDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmptyNumberDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponseBody) SetCode(v string) *DescribeEmptyNumberDetectResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetData(v []*DescribeEmptyNumberDetectResponseBodyData) *DescribeEmptyNumberDetectResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetMessage(v string) *DescribeEmptyNumberDetectResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBody) SetRequestId(v string) *DescribeEmptyNumberDetectResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEmptyNumberDetectResponseBodyData struct {
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEmptyNumberDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponseBodyData) SetNumber(v string) *DescribeEmptyNumberDetectResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribeEmptyNumberDetectResponseBodyData) SetStatus(v string) *DescribeEmptyNumberDetectResponseBodyData {
	s.Status = &v
	return s
}

type DescribeEmptyNumberDetectResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEmptyNumberDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEmptyNumberDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmptyNumberDetectResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberDetectResponse) SetHeaders(v map[string]*string) *DescribeEmptyNumberDetectResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmptyNumberDetectResponse) SetBody(v *DescribeEmptyNumberDetectResponseBody) *DescribeEmptyNumberDetectResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAnalysisRequest struct {
	AuthCode             *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	InputNumber          *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NumberType           *int64  `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetMask(v string) *DescribePhoneNumberAnalysisRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetNumberType(v int64) *DescribePhoneNumberAnalysisRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetRate(v int64) *DescribePhoneNumberAnalysisRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribePhoneNumberAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetCode(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetData(v []*DescribePhoneNumberAnalysisResponseBodyData) *DescribePhoneNumberAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetMessage(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBody) SetRequestId(v string) *DescribePhoneNumberAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAnalysisResponseBodyData struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetCode(v string) *DescribePhoneNumberAnalysisResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponseBodyData) SetNumber(v string) *DescribePhoneNumberAnalysisResponseBodyData {
	s.Number = &v
	return s
}

type DescribePhoneNumberAnalysisResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePhoneNumberAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePhoneNumberAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetBody(v *DescribePhoneNumberAnalysisResponseBody) *DescribePhoneNumberAnalysisResponse {
	s.Body = v
	return s
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribePhoneNumberAttributeRequest) SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest {
	s.PhoneNumber = &v
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

type DescribePhoneNumberAttributeResponseBody struct {
	Code                 *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribePhoneNumberAttributeResponseBody) SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody {
	s.PhoneNumberAttribute = v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
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

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCity(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.City = &v
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
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribePhoneNumberResaleRequest) SetPhoneNumber(v string) *DescribePhoneNumberResaleRequest {
	s.PhoneNumber = &v
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

type DescribePhoneNumberStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribePhoneNumberStatusRequest) SetPhoneNumber(v string) *DescribePhoneNumberStatusRequest {
	s.PhoneNumber = &v
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

type DescribePhoneNumberStatusResponseBody struct {
	Code        *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneStatus *DescribePhoneNumberStatusResponseBodyPhoneStatus `json:"PhoneStatus,omitempty" xml:"PhoneStatus,omitempty" type:"Struct"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribePhoneNumberStatusResponseBody) SetPhoneStatus(v *DescribePhoneNumberStatusResponseBodyPhoneStatus) *DescribePhoneNumberStatusResponseBody {
	s.PhoneStatus = v
	return s
}

func (s *DescribePhoneNumberStatusResponseBody) SetRequestId(v string) *DescribePhoneNumberStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribePhoneNumberStatusResponseBodyPhoneStatus struct {
	Carrier  *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	SerialId *string `json:"SerialId,omitempty" xml:"SerialId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribePhoneNumberStatusResponseBodyPhoneStatus) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetCarrier(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetSerialId(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.SerialId = &v
	return s
}

func (s *DescribePhoneNumberStatusResponseBodyPhoneStatus) SetStatus(v string) *DescribePhoneNumberStatusResponseBodyPhoneStatus {
	s.Status = &v
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

func (client *Client) DescribeEmptyNumberDetectWithOptions(request *DescribeEmptyNumberDetectRequest, runtime *util.RuntimeOptions) (_result *DescribeEmptyNumberDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EncryptType"] = request.EncryptType
	query["OwnerId"] = request.OwnerId
	query["Phone"] = request.Phone
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEmptyNumberDetect"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEmptyNumberDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEmptyNumberDetect(request *DescribeEmptyNumberDetectRequest) (_result *DescribeEmptyNumberDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEmptyNumberDetectResponse{}
	_body, _err := client.DescribeEmptyNumberDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysisWithOptions(request *DescribePhoneNumberAnalysisRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AuthCode"] = request.AuthCode
	query["InputNumber"] = request.InputNumber
	query["Mask"] = request.Mask
	query["NumberType"] = request.NumberType
	query["OwnerId"] = request.OwnerId
	query["Rate"] = request.Rate
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAnalysis"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePhoneNumberAnalysis(request *DescribePhoneNumberAnalysisRequest) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PhoneNumber"] = request.PhoneNumber
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberAttribute"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PhoneNumber"] = request.PhoneNumber
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Since"] = request.Since
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberResale"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberResaleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribePhoneNumberStatusWithOptions(request *DescribePhoneNumberStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePhoneNumberStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PhoneNumber"] = request.PhoneNumber
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePhoneNumberStatus"),
		Version:     tea.String("2020-02-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePhoneNumberStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
