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

type QueryMessageRequest struct {
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageRequest) SetMessageId(v string) *QueryMessageRequest {
	s.MessageId = &v
	return s
}

type QueryMessageResponseBody struct {
	Status              *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorDescription    *string                               `json:"ErrorDescription,omitempty" xml:"ErrorDescription,omitempty"`
	ResponseCode        *string                               `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	ReceiveDate         *string                               `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	NumberDetail        *QueryMessageResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	Message             *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ResponseDescription *string                               `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	ErrorCode           *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	SendDate            *string                               `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	To                  *string                               `json:"To,omitempty" xml:"To,omitempty"`
	MessageId           *string                               `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBody) SetStatus(v string) *QueryMessageResponseBody {
	s.Status = &v
	return s
}

func (s *QueryMessageResponseBody) SetErrorDescription(v string) *QueryMessageResponseBody {
	s.ErrorDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseCode(v string) *QueryMessageResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetReceiveDate(v string) *QueryMessageResponseBody {
	s.ReceiveDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetNumberDetail(v *QueryMessageResponseBodyNumberDetail) *QueryMessageResponseBody {
	s.NumberDetail = v
	return s
}

func (s *QueryMessageResponseBody) SetMessage(v string) *QueryMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseDescription(v string) *QueryMessageResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetErrorCode(v string) *QueryMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetSendDate(v string) *QueryMessageResponseBody {
	s.SendDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetTo(v string) *QueryMessageResponseBody {
	s.To = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessageId(v string) *QueryMessageResponseBody {
	s.MessageId = &v
	return s
}

type QueryMessageResponseBodyNumberDetail struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s QueryMessageResponseBodyNumberDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBodyNumberDetail) SetCarrier(v string) *QueryMessageResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) SetRegion(v string) *QueryMessageResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) SetCountry(v string) *QueryMessageResponseBodyNumberDetail {
	s.Country = &v
	return s
}

type QueryMessageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageResponse) SetHeaders(v map[string]*string) *QueryMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageResponse) SetBody(v *QueryMessageResponseBody) *QueryMessageResponse {
	s.Body = v
	return s
}

type BatchSendMessageToGlobeRequest struct {
	To      *string `json:"To,omitempty" xml:"To,omitempty"`
	From    *string `json:"From,omitempty" xml:"From,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchSendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeRequest) SetTo(v string) *BatchSendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetFrom(v string) *BatchSendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetMessage(v string) *BatchSendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetType(v string) *BatchSendMessageToGlobeRequest {
	s.Type = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetTaskId(v string) *BatchSendMessageToGlobeRequest {
	s.TaskId = &v
	return s
}

type BatchSendMessageToGlobeResponseBody struct {
	ResponseCode        *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedList          *string `json:"FailedList,omitempty" xml:"FailedList,omitempty"`
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	From                *string `json:"From,omitempty" xml:"From,omitempty"`
	MessageIdList       *string `json:"MessageIdList,omitempty" xml:"MessageIdList,omitempty"`
	SuccessCount        *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BatchSendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseCode(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetRequestId(v string) *BatchSendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetFailedList(v string) *BatchSendMessageToGlobeResponseBody {
	s.FailedList = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseDescription(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetFrom(v string) *BatchSendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetMessageIdList(v string) *BatchSendMessageToGlobeResponseBody {
	s.MessageIdList = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetSuccessCount(v string) *BatchSendMessageToGlobeResponseBody {
	s.SuccessCount = &v
	return s
}

type BatchSendMessageToGlobeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSendMessageToGlobeResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessageToGlobeResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeResponse) SetHeaders(v map[string]*string) *BatchSendMessageToGlobeResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMessageToGlobeResponse) SetBody(v *BatchSendMessageToGlobeResponseBody) *BatchSendMessageToGlobeResponse {
	s.Body = v
	return s
}

type SmsConversionRequest struct {
	MessageId      *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Delivered      *bool   `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	ConversionTime *int64  `json:"ConversionTime,omitempty" xml:"ConversionTime,omitempty"`
}

func (s SmsConversionRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionRequest) GoString() string {
	return s.String()
}

func (s *SmsConversionRequest) SetMessageId(v string) *SmsConversionRequest {
	s.MessageId = &v
	return s
}

func (s *SmsConversionRequest) SetDelivered(v bool) *SmsConversionRequest {
	s.Delivered = &v
	return s
}

func (s *SmsConversionRequest) SetConversionTime(v int64) *SmsConversionRequest {
	s.ConversionTime = &v
	return s
}

type SmsConversionResponseBody struct {
	ResponseCode        *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SmsConversionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionResponseBody) GoString() string {
	return s.String()
}

func (s *SmsConversionResponseBody) SetResponseCode(v string) *SmsConversionResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SmsConversionResponseBody) SetResponseDescription(v string) *SmsConversionResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SmsConversionResponseBody) SetRequestId(v string) *SmsConversionResponseBody {
	s.RequestId = &v
	return s
}

type SmsConversionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SmsConversionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SmsConversionResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionResponse) GoString() string {
	return s.String()
}

func (s *SmsConversionResponse) SetHeaders(v map[string]*string) *SmsConversionResponse {
	s.Headers = v
	return s
}

func (s *SmsConversionResponse) SetBody(v *SmsConversionResponseBody) *SmsConversionResponse {
	s.Body = v
	return s
}

type SendMessageToGlobeRequest struct {
	To      *string `json:"To,omitempty" xml:"To,omitempty"`
	From    *string `json:"From,omitempty" xml:"From,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeRequest) SetTo(v string) *SendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetFrom(v string) *SendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetMessage(v string) *SendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetTaskId(v string) *SendMessageToGlobeRequest {
	s.TaskId = &v
	return s
}

type SendMessageToGlobeResponseBody struct {
	ResponseCode        *string                                     `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	NumberDetail        *SendMessageToGlobeResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	RequestId           *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Segments            *string                                     `json:"Segments,omitempty" xml:"Segments,omitempty"`
	ResponseDescription *string                                     `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	From                *string                                     `json:"From,omitempty" xml:"From,omitempty"`
	To                  *string                                     `json:"To,omitempty" xml:"To,omitempty"`
	MessageId           *string                                     `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBody) SetResponseCode(v string) *SendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetNumberDetail(v *SendMessageToGlobeResponseBodyNumberDetail) *SendMessageToGlobeResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetRequestId(v string) *SendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetSegments(v string) *SendMessageToGlobeResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetResponseDescription(v string) *SendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetFrom(v string) *SendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetTo(v string) *SendMessageToGlobeResponseBody {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetMessageId(v string) *SendMessageToGlobeResponseBody {
	s.MessageId = &v
	return s
}

type SendMessageToGlobeResponseBodyNumberDetail struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s SendMessageToGlobeResponseBodyNumberDetail) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCarrier(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetRegion(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCountry(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Country = &v
	return s
}

type SendMessageToGlobeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageToGlobeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponse) SetHeaders(v map[string]*string) *SendMessageToGlobeResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToGlobeResponse) SetBody(v *SendMessageToGlobeResponseBody) *SendMessageToGlobeResponse {
	s.Body = v
	return s
}

type ConversionDataRequest struct {
	ReportTime     *int64  `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
}

func (s ConversionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataRequest) GoString() string {
	return s.String()
}

func (s *ConversionDataRequest) SetReportTime(v int64) *ConversionDataRequest {
	s.ReportTime = &v
	return s
}

func (s *ConversionDataRequest) SetConversionRate(v string) *ConversionDataRequest {
	s.ConversionRate = &v
	return s
}

type ConversionDataResponseBody struct {
	ResponseCode        *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConversionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataResponseBody) GoString() string {
	return s.String()
}

func (s *ConversionDataResponseBody) SetResponseCode(v string) *ConversionDataResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *ConversionDataResponseBody) SetResponseDescription(v string) *ConversionDataResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *ConversionDataResponseBody) SetRequestId(v string) *ConversionDataResponseBody {
	s.RequestId = &v
	return s
}

type ConversionDataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConversionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConversionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataResponse) GoString() string {
	return s.String()
}

func (s *ConversionDataResponse) SetHeaders(v map[string]*string) *ConversionDataResponse {
	s.Headers = v
	return s
}

func (s *ConversionDataResponse) SetBody(v *ConversionDataResponseBody) *ConversionDataResponse {
	s.Body = v
	return s
}

type SendMessageWithTemplateRequest struct {
	To              *string `json:"To,omitempty" xml:"To,omitempty"`
	From            *string `json:"From,omitempty" xml:"From,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam   *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
}

func (s SendMessageWithTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateRequest) SetTo(v string) *SendMessageWithTemplateRequest {
	s.To = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetFrom(v string) *SendMessageWithTemplateRequest {
	s.From = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetTemplateCode(v string) *SendMessageWithTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetTemplateParam(v string) *SendMessageWithTemplateRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetSmsUpExtendCode(v string) *SendMessageWithTemplateRequest {
	s.SmsUpExtendCode = &v
	return s
}

type SendMessageWithTemplateResponseBody struct {
	ResponseCode        *string                                          `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	NumberDetail        *SendMessageWithTemplateResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	ResponseDescription *string                                          `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	Segments            *string                                          `json:"Segments,omitempty" xml:"Segments,omitempty"`
	To                  *string                                          `json:"To,omitempty" xml:"To,omitempty"`
	MessageId           *string                                          `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendMessageWithTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponseBody) SetResponseCode(v string) *SendMessageWithTemplateResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetNumberDetail(v *SendMessageWithTemplateResponseBodyNumberDetail) *SendMessageWithTemplateResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetResponseDescription(v string) *SendMessageWithTemplateResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetSegments(v string) *SendMessageWithTemplateResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetTo(v string) *SendMessageWithTemplateResponseBody {
	s.To = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetMessageId(v string) *SendMessageWithTemplateResponseBody {
	s.MessageId = &v
	return s
}

type SendMessageWithTemplateResponseBodyNumberDetail struct {
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s SendMessageWithTemplateResponseBodyNumberDetail) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetCarrier(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetRegion(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetCountry(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Country = &v
	return s
}

type SendMessageWithTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageWithTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageWithTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateResponse) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponse) SetHeaders(v map[string]*string) *SendMessageWithTemplateResponse {
	s.Headers = v
	return s
}

func (s *SendMessageWithTemplateResponse) SetBody(v *SendMessageWithTemplateResponseBody) *SendMessageWithTemplateResponse {
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
	client.EndpointMap = map[string]*string{
		"ap-southeast-1": tea.String("dysmsapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": tea.String("dysmsapi-xman.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     tea.String("dysmsapi-proxy.cn-beijing.aliyuncs.com"),
		"cn-hongkong":    tea.String("dysmsapi-xman.cn-hongkong.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dysmsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryMessageWithOptions(request *QueryMessageRequest, runtime *util.RuntimeOptions) (_result *QueryMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMessage"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMessage(request *QueryMessageRequest) (_result *QueryMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMessageResponse{}
	_body, _err := client.QueryMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSendMessageToGlobeWithOptions(request *BatchSendMessageToGlobeRequest, runtime *util.RuntimeOptions) (_result *BatchSendMessageToGlobeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchSendMessageToGlobeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchSendMessageToGlobe"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSendMessageToGlobe(request *BatchSendMessageToGlobeRequest) (_result *BatchSendMessageToGlobeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSendMessageToGlobeResponse{}
	_body, _err := client.BatchSendMessageToGlobeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SmsConversionWithOptions(request *SmsConversionRequest, runtime *util.RuntimeOptions) (_result *SmsConversionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SmsConversionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SmsConversion"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SmsConversion(request *SmsConversionRequest) (_result *SmsConversionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmsConversionResponse{}
	_body, _err := client.SmsConversionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageToGlobeWithOptions(request *SendMessageToGlobeRequest, runtime *util.RuntimeOptions) (_result *SendMessageToGlobeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMessageToGlobe"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessageToGlobe(request *SendMessageToGlobeRequest) (_result *SendMessageToGlobeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.SendMessageToGlobeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConversionDataWithOptions(request *ConversionDataRequest, runtime *util.RuntimeOptions) (_result *ConversionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConversionDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConversionData"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConversionData(request *ConversionDataRequest) (_result *ConversionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConversionDataResponse{}
	_body, _err := client.ConversionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithTemplateWithOptions(request *SendMessageWithTemplateRequest, runtime *util.RuntimeOptions) (_result *SendMessageWithTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMessageWithTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMessageWithTemplate"), tea.String("2018-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessageWithTemplate(request *SendMessageWithTemplateRequest) (_result *SendMessageWithTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageWithTemplateResponse{}
	_body, _err := client.SendMessageWithTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
