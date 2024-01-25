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

type GetMessageStatusRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMessageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMessageStatusRequest) SetTaskId(v string) *GetMessageStatusRequest {
	s.TaskId = &v
	return s
}

type GetMessageStatusResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMessageStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageStatusResponseBody) SetCode(v string) *GetMessageStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetMessageStatusResponseBody) SetData(v *GetMessageStatusResponseBodyData) *GetMessageStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageStatusResponseBody) SetMessage(v string) *GetMessageStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetMessageStatusResponseBody) SetRequestId(v string) *GetMessageStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageStatusResponseBodyData struct {
	Mobile    *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMessageStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMessageStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageStatusResponseBodyData) SetMobile(v string) *GetMessageStatusResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetMessageStatusResponseBodyData) SetRequestId(v string) *GetMessageStatusResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetMessageStatusResponseBodyData) SetStatus(v string) *GetMessageStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetMessageStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMessageStatusResponse) SetHeaders(v map[string]*string) *GetMessageStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMessageStatusResponse) SetStatusCode(v int32) *GetMessageStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageStatusResponse) SetBody(v *GetMessageStatusResponseBody) *GetMessageStatusResponse {
	s.Body = v
	return s
}

type SendBatchMessageRequest struct {
	BatchFlag         *string `json:"BatchFlag,omitempty" xml:"BatchFlag,omitempty"`
	ExtendInfo        *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	IdType            *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	PhoneNumberJson   *string `json:"PhoneNumberJson,omitempty" xml:"PhoneNumberJson,omitempty"`
	SignNameJson      *string `json:"SignNameJson,omitempty" xml:"SignNameJson,omitempty"`
	SpecificChannel   *string `json:"SpecificChannel,omitempty" xml:"SpecificChannel,omitempty"`
	TemplateCode      *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParamJson *string `json:"TemplateParamJson,omitempty" xml:"TemplateParamJson,omitempty"`
}

func (s SendBatchMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBatchMessageRequest) GoString() string {
	return s.String()
}

func (s *SendBatchMessageRequest) SetBatchFlag(v string) *SendBatchMessageRequest {
	s.BatchFlag = &v
	return s
}

func (s *SendBatchMessageRequest) SetExtendInfo(v string) *SendBatchMessageRequest {
	s.ExtendInfo = &v
	return s
}

func (s *SendBatchMessageRequest) SetIdType(v string) *SendBatchMessageRequest {
	s.IdType = &v
	return s
}

func (s *SendBatchMessageRequest) SetPhoneNumberJson(v string) *SendBatchMessageRequest {
	s.PhoneNumberJson = &v
	return s
}

func (s *SendBatchMessageRequest) SetSignNameJson(v string) *SendBatchMessageRequest {
	s.SignNameJson = &v
	return s
}

func (s *SendBatchMessageRequest) SetSpecificChannel(v string) *SendBatchMessageRequest {
	s.SpecificChannel = &v
	return s
}

func (s *SendBatchMessageRequest) SetTemplateCode(v string) *SendBatchMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendBatchMessageRequest) SetTemplateParamJson(v string) *SendBatchMessageRequest {
	s.TemplateParamJson = &v
	return s
}

type SendBatchMessageResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*SendBatchMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendBatchMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBatchMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendBatchMessageResponseBody) SetCode(v string) *SendBatchMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendBatchMessageResponseBody) SetData(v []*SendBatchMessageResponseBodyData) *SendBatchMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendBatchMessageResponseBody) SetMessage(v string) *SendBatchMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendBatchMessageResponseBody) SetRequestId(v string) *SendBatchMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendBatchMessageResponseBodyData struct {
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SendBatchMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendBatchMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendBatchMessageResponseBodyData) SetMobile(v string) *SendBatchMessageResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *SendBatchMessageResponseBodyData) SetTaskId(v string) *SendBatchMessageResponseBodyData {
	s.TaskId = &v
	return s
}

type SendBatchMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBatchMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBatchMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendBatchMessageResponse) GoString() string {
	return s.String()
}

func (s *SendBatchMessageResponse) SetHeaders(v map[string]*string) *SendBatchMessageResponse {
	s.Headers = v
	return s
}

func (s *SendBatchMessageResponse) SetStatusCode(v int32) *SendBatchMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBatchMessageResponse) SetBody(v *SendBatchMessageResponseBody) *SendBatchMessageResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-finplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetMessageStatusWithOptions(request *GetMessageStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMessageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMessageStatus"),
		Version:     tea.String("2021-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sms/message/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMessageStatus(request *GetMessageStatusRequest) (_result *GetMessageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMessageStatusResponse{}
	_body, _err := client.GetMessageStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendBatchMessageWithOptions(request *SendBatchMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendBatchMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchFlag)) {
		body["BatchFlag"] = request.BatchFlag
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendInfo)) {
		body["ExtendInfo"] = request.ExtendInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		body["IdType"] = request.IdType
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberJson)) {
		body["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !tea.BoolValue(util.IsUnset(request.SignNameJson)) {
		body["SignNameJson"] = request.SignNameJson
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificChannel)) {
		body["SpecificChannel"] = request.SpecificChannel
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParamJson)) {
		body["TemplateParamJson"] = request.TemplateParamJson
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendBatchMessage"),
		Version:     tea.String("2021-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sms/message/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendBatchMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendBatchMessage(request *SendBatchMessageRequest) (_result *SendBatchMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendBatchMessageResponse{}
	_body, _err := client.SendBatchMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
