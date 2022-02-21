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

type SendMessageRequest struct {
	// 人群ID，用于关联人群。
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 外部拓展字段。
	OutIds []*string `json:"OutIds,omitempty" xml:"OutIds,omitempty" type:"Repeated"`
	// 手机号，每个手机号对应一个模板变量、上行拓展码和外部拓展字段。
	PhoneNumbers []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// 发送计划ID，用于关联发送计划。
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// 签名名称。
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名ID，同时只能指定签名名称或签名ID其中之一。
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// 短信上行拓展码。
	SmsUpExtendCodes []*string `json:"SmsUpExtendCodes,omitempty" xml:"SmsUpExtendCodes,omitempty" type:"Repeated"`
	// 模板Code。
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板ID，同时只能指定模板Code或模板ID其中之一。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 短信模板变量对应的实际值，JSON格式。支持传入多个参数，示例：{"name":"张三","number":"15038****76"}。
	TemplateParams []*string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty" type:"Repeated"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetGroupId(v string) *SendMessageRequest {
	s.GroupId = &v
	return s
}

func (s *SendMessageRequest) SetOutIds(v []*string) *SendMessageRequest {
	s.OutIds = v
	return s
}

func (s *SendMessageRequest) SetPhoneNumbers(v []*string) *SendMessageRequest {
	s.PhoneNumbers = v
	return s
}

func (s *SendMessageRequest) SetScheduleId(v string) *SendMessageRequest {
	s.ScheduleId = &v
	return s
}

func (s *SendMessageRequest) SetSignName(v string) *SendMessageRequest {
	s.SignName = &v
	return s
}

func (s *SendMessageRequest) SetSignatureId(v string) *SendMessageRequest {
	s.SignatureId = &v
	return s
}

func (s *SendMessageRequest) SetSmsUpExtendCodes(v []*string) *SendMessageRequest {
	s.SmsUpExtendCodes = v
	return s
}

func (s *SendMessageRequest) SetTemplateCode(v string) *SendMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageRequest) SetTemplateId(v string) *SendMessageRequest {
	s.TemplateId = &v
	return s
}

func (s *SendMessageRequest) SetTemplateParams(v []*string) *SendMessageRequest {
	s.TemplateParams = v
	return s
}

type SendMessageResponseBody struct {
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetData(v string) *SendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorCode(v int32) *SendMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorMessage(v string) *SendMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("paiplugin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OutIds)) {
		body["OutIds"] = request.OutIds
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumbers)) {
		body["PhoneNumbers"] = request.PhoneNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleId)) {
		body["ScheduleId"] = request.ScheduleId
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		body["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureId)) {
		body["SignatureId"] = request.SignatureId
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCodes)) {
		body["SmsUpExtendCodes"] = request.SmsUpExtendCodes
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParams)) {
		body["TemplateParams"] = request.TemplateParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
