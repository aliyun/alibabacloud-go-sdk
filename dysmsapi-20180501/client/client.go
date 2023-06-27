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

type BatchSendMessageToGlobeRequest struct {
	// The mobile phone number of the sender. You can also specify a sender ID. The sender ID can contain both letters and digits. If it does, the ID must be between 1 to 11 characters in length. If the sender ID contains only digits, it must be 1 to 15 characters in length.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The content of the message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the messaging campaign. It must be 1 to 255 characters in length. The ID is the value of the TaskId field in the delivery receipt of the message.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mobile phone numbers to which the message is sent. You must add the dialing code to the beginning of each mobile phone number.
	//
	// For more information, see [Dialing codes](https://www.alibabacloud.com/help/zh/short-message-service/latest/dialing-codes).
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The type of the message. Valid values:
	//
	// *   **NOTIFY**: notification
	// *   **MKT**: promotional message
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ValidityPeriod *int64  `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s BatchSendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeRequest) SetFrom(v string) *BatchSendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetMessage(v string) *BatchSendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetTaskId(v string) *BatchSendMessageToGlobeRequest {
	s.TaskId = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetTo(v string) *BatchSendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetType(v string) *BatchSendMessageToGlobeRequest {
	s.Type = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetValidityPeriod(v int64) *BatchSendMessageToGlobeRequest {
	s.ValidityPeriod = &v
	return s
}

type BatchSendMessageToGlobeResponseBody struct {
	// The list of mobile phone numbers that failed to receive the message.
	FailedList *string `json:"FailedList,omitempty" xml:"FailedList,omitempty"`
	// The sender ID returned.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the message.
	MessageIdList *string `json:"MessageIdList,omitempty" xml:"MessageIdList,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. If OK is returned, the request is successful. For more information, see [Error codes](https://www.alibabacloud.com/help/zh/short-message-service/latest/error-codes).
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the status code.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The number of mobile phone numbers that received the message.
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s BatchSendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeResponseBody) SetFailedList(v string) *BatchSendMessageToGlobeResponseBody {
	s.FailedList = &v
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

func (s *BatchSendMessageToGlobeResponseBody) SetRequestId(v string) *BatchSendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseCode(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetResponseDescription(v string) *BatchSendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *BatchSendMessageToGlobeResponseBody) SetSuccessCount(v string) *BatchSendMessageToGlobeResponseBody {
	s.SuccessCount = &v
	return s
}

type BatchSendMessageToGlobeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchSendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchSendMessageToGlobeResponse) SetStatusCode(v int32) *BatchSendMessageToGlobeResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMessageToGlobeResponse) SetBody(v *BatchSendMessageToGlobeResponseBody) *BatchSendMessageToGlobeResponse {
	s.Body = v
	return s
}

type ConversionDataRequest struct {
	// Conversion rate monitoring return value.
	//
	// >  The value of this parameter is of type double, and the value is between \[0,1].
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
	// Timestamp of the conversion rate observation should be a Unix timestamp, millisecond-level long integer.
	//
	// >  If this field is not specified: the current timestamp is the default.
	ReportTime *int64 `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
}

func (s ConversionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataRequest) GoString() string {
	return s.String()
}

func (s *ConversionDataRequest) SetConversionRate(v string) *ConversionDataRequest {
	s.ConversionRate = &v
	return s
}

func (s *ConversionDataRequest) SetReportTime(v int64) *ConversionDataRequest {
	s.ReportTime = &v
	return s
}

type ConversionDataResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code. Returning OK means the request was successful. For other error codes, please refer to the [Error codes](~~180674~~) list.
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the status code.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
}

func (s ConversionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConversionDataResponseBody) GoString() string {
	return s.String()
}

func (s *ConversionDataResponseBody) SetRequestId(v string) *ConversionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConversionDataResponseBody) SetResponseCode(v string) *ConversionDataResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *ConversionDataResponseBody) SetResponseDescription(v string) *ConversionDataResponseBody {
	s.ResponseDescription = &v
	return s
}

type ConversionDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConversionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConversionDataResponse) SetStatusCode(v int32) *ConversionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversionDataResponse) SetBody(v *ConversionDataResponseBody) *ConversionDataResponse {
	s.Body = v
	return s
}

type QueryMessageRequest struct {
	// The ID of the message.
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
	// The status code of the message.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The description of the status code.
	ErrorDescription *string `json:"ErrorDescription,omitempty" xml:"ErrorDescription,omitempty"`
	// The content of the message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number.
	NumberDetail *QueryMessageResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The time when the delivery receipt was received from the carrier.
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code of the delivery request.
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the delivery request status.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The time when the message was sent to the carrier.
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The delivery status of the message.
	//
	// *   1: The message was sent.
	// *   2: The message failed to be sent.
	// *   3: The message is being sent.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mobile phone number to which the message was sent.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s QueryMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBody) SetErrorCode(v string) *QueryMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetErrorDescription(v string) *QueryMessageResponseBody {
	s.ErrorDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessage(v string) *QueryMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessageId(v string) *QueryMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *QueryMessageResponseBody) SetNumberDetail(v *QueryMessageResponseBodyNumberDetail) *QueryMessageResponseBody {
	s.NumberDetail = v
	return s
}

func (s *QueryMessageResponseBody) SetReceiveDate(v string) *QueryMessageResponseBody {
	s.ReceiveDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetRequestId(v string) *QueryMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseCode(v string) *QueryMessageResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseDescription(v string) *QueryMessageResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetSendDate(v string) *QueryMessageResponseBody {
	s.SendDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetStatus(v string) *QueryMessageResponseBody {
	s.Status = &v
	return s
}

func (s *QueryMessageResponseBody) SetTo(v string) *QueryMessageResponseBody {
	s.To = &v
	return s
}

type QueryMessageResponseBodyNumberDetail struct {
	// The carrier that owns the mobile phone number.
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The country to which the mobile phone number belongs.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The region to which the mobile phone number belongs.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s *QueryMessageResponseBodyNumberDetail) SetCountry(v string) *QueryMessageResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) SetRegion(v string) *QueryMessageResponseBodyNumberDetail {
	s.Region = &v
	return s
}

type QueryMessageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMessageResponse) SetStatusCode(v int32) *QueryMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageResponse) SetBody(v *QueryMessageResponseBody) *QueryMessageResponse {
	s.Body = v
	return s
}

type SendMessageToGlobeRequest struct {
	// The mobile phone number of the sender. You can also specify a sender ID. The sender ID can contain both letters and digits. If it does, the ID must be between 1 to 11 characters in length. If the sender ID contains only digits, it must be 1 to 15 characters in length.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The content of the message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the messaging campaign. It must be 1 to 255 characters in length. The ID is the value of the TaskId field in the delivery receipt of the message.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mobile phone number to which the message is sent. You must add the dialing code to the beginning of the mobile phone number. Example: 8521245567\*\*\*\*.
	//
	// For more information, see [Dialing codes](https://www.alibabacloud.com/help/zh/short-message-service/latest/dialing-codes).
	To             *string `json:"To,omitempty" xml:"To,omitempty"`
	ValidityPeriod *int64  `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s SendMessageToGlobeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeRequest) GoString() string {
	return s.String()
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

func (s *SendMessageToGlobeRequest) SetTo(v string) *SendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetValidityPeriod(v int64) *SendMessageToGlobeRequest {
	s.ValidityPeriod = &v
	return s
}

type SendMessageToGlobeResponseBody struct {
	// The sender ID returned.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number of the recipient.
	NumberDetail *SendMessageToGlobeResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code of the delivery request.
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the delivery request status.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The number of messages that incurred fees.
	Segments *string `json:"Segments,omitempty" xml:"Segments,omitempty"`
	// The mobile phone number to which the message was sent.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendMessageToGlobeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBody) SetFrom(v string) *SendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetMessageId(v string) *SendMessageToGlobeResponseBody {
	s.MessageId = &v
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

func (s *SendMessageToGlobeResponseBody) SetResponseCode(v string) *SendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetResponseDescription(v string) *SendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetSegments(v string) *SendMessageToGlobeResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetTo(v string) *SendMessageToGlobeResponseBody {
	s.To = &v
	return s
}

type SendMessageToGlobeResponseBodyNumberDetail struct {
	// The carrier that owns the mobile phone number.
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The country to which the mobile phone number belongs.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The region to which the mobile phone number belongs.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCountry(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetRegion(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Region = &v
	return s
}

type SendMessageToGlobeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendMessageToGlobeResponse) SetStatusCode(v int32) *SendMessageToGlobeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageToGlobeResponse) SetBody(v *SendMessageToGlobeResponseBody) *SendMessageToGlobeResponse {
	s.Body = v
	return s
}

type SendMessageWithTemplateRequest struct {
	// The signature. To query the signature, log on to the [Short Message Service (SMS) console](https://sms-intl.console.aliyun.com/overview) and navigate to the **Signatures** tab of the **Go China** page.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The extension code of the MO message.
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the message template. To query the code, log on to the [SMS console](https://sms-intl.console.aliyun.com/overview) and navigate to the **Templates** tab of the **Go China** page.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the message template. If a variable exists in the template, the parameter is required.
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	// The mobile phone number to which the message is sent. You must add the country code to the beginning of the mobile phone number. Example: 861503871\*\*\*\*.
	//
	// For more information, see [Dialing codes](https://www.alibabacloud.com/help/zh/short-message-service/latest/dialing-codes).
	To             *string `json:"To,omitempty" xml:"To,omitempty"`
	ValidityPeriod *int64  `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s SendMessageWithTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateRequest) SetFrom(v string) *SendMessageWithTemplateRequest {
	s.From = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetSmsUpExtendCode(v string) *SendMessageWithTemplateRequest {
	s.SmsUpExtendCode = &v
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

func (s *SendMessageWithTemplateRequest) SetTo(v string) *SendMessageWithTemplateRequest {
	s.To = &v
	return s
}

func (s *SendMessageWithTemplateRequest) SetValidityPeriod(v int64) *SendMessageWithTemplateRequest {
	s.ValidityPeriod = &v
	return s
}

type SendMessageWithTemplateResponseBody struct {
	// The ID of the message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number of the recipient.
	NumberDetail *SendMessageWithTemplateResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code of the delivery request.
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the delivery request status.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The number of messages that incurred fees.
	Segments *string `json:"Segments,omitempty" xml:"Segments,omitempty"`
	// The mobile phone number to which the message was sent. The dialing code was added to the beginning of the mobile phone number. Example: 861503871\*\*\*\*.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendMessageWithTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageWithTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponseBody) SetMessageId(v string) *SendMessageWithTemplateResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetNumberDetail(v *SendMessageWithTemplateResponseBodyNumberDetail) *SendMessageWithTemplateResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetRequestId(v string) *SendMessageWithTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetResponseCode(v string) *SendMessageWithTemplateResponseBody {
	s.ResponseCode = &v
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

type SendMessageWithTemplateResponseBodyNumberDetail struct {
	// The carrier that owns the mobile phone number.
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The country to which the mobile phone number belongs.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The region to which the mobile phone number belongs.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetCountry(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetRegion(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Region = &v
	return s
}

type SendMessageWithTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageWithTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendMessageWithTemplateResponse) SetStatusCode(v int32) *SendMessageWithTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageWithTemplateResponse) SetBody(v *SendMessageWithTemplateResponseBody) *SendMessageWithTemplateResponse {
	s.Body = v
	return s
}

type SmsConversionRequest struct {
	// The time when the OTP message was delivered. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// *   If you leave the parameter empty, the current timestamp is specified by default.
	// *   If you specify the parameter, the timestamp must be greater than the message sending time and less than the current timestamp.
	ConversionTime *int64 `json:"ConversionTime,omitempty" xml:"ConversionTime,omitempty"`
	// Specifies whether customers replied to the OTP message. Valid values: true and false.
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// The ID of the OTP message.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SmsConversionRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionRequest) GoString() string {
	return s.String()
}

func (s *SmsConversionRequest) SetConversionTime(v int64) *SmsConversionRequest {
	s.ConversionTime = &v
	return s
}

func (s *SmsConversionRequest) SetDelivered(v bool) *SmsConversionRequest {
	s.Delivered = &v
	return s
}

func (s *SmsConversionRequest) SetMessageId(v string) *SmsConversionRequest {
	s.MessageId = &v
	return s
}

type SmsConversionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. If OK is returned, the request is successful. For more information, see [Error codes](~~180674~~).
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the status code.
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
}

func (s SmsConversionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsConversionResponseBody) GoString() string {
	return s.String()
}

func (s *SmsConversionResponseBody) SetRequestId(v string) *SmsConversionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsConversionResponseBody) SetResponseCode(v string) *SmsConversionResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SmsConversionResponseBody) SetResponseDescription(v string) *SmsConversionResponseBody {
	s.ResponseDescription = &v
	return s
}

type SmsConversionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SmsConversionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SmsConversionResponse) SetStatusCode(v int32) *SmsConversionResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsConversionResponse) SetBody(v *SmsConversionResponseBody) *SmsConversionResponse {
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
		"ap-southeast-5": tea.String("dysmsapi.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     tea.String("dysmsapi-proxy.cn-beijing.aliyuncs.com"),
		"cn-hongkong":    tea.String("dysmsapi-xman.cn-hongkong.aliyuncs.com"),
		"eu-central-1":   tea.String("dysmsapi.eu-central-1.aliyuncs.com"),
		"us-east-1":      tea.String("dysmsapi.us-east-1.aliyuncs.com"),
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

/**
 * *   You cannot call the BatchSendMessageToGlobe operation to send messages to the Chinese mainland.
 * *   You can call the BatchSendMessageToGlobe operation to send notifications and promotional messages to a limited number of mobile phone numbers at a time. To send messages exceeding more than 1,000 mobile number per request, you can choose to use the broadcast messaging feature available in the Alibaba Cloud SMS console.
 * *   For time-sensitive related messages, we recommend that you use the [SendMessageToGlobe](~~SendMessageToGlobe~~) operation to ensure that messages are delivered on time.
 * *   In each request, you can send messages to up to 1,000 mobile phone numbers.
 * ### QPS limit
 * You may call this operation only once per second. If the number of calls per second exceeds this limit, throttling will be triggered. This can potentially impact your business operations. Therefore, we recommend that you take note of this limit when making calls to this operation.
 *
 * @param request BatchSendMessageToGlobeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BatchSendMessageToGlobeResponse
 */
func (client *Client) BatchSendMessageToGlobeWithOptions(request *BatchSendMessageToGlobeRequest, runtime *util.RuntimeOptions) (_result *BatchSendMessageToGlobeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityPeriod)) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSendMessageToGlobe"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendMessageToGlobeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   You cannot call the BatchSendMessageToGlobe operation to send messages to the Chinese mainland.
 * *   You can call the BatchSendMessageToGlobe operation to send notifications and promotional messages to a limited number of mobile phone numbers at a time. To send messages exceeding more than 1,000 mobile number per request, you can choose to use the broadcast messaging feature available in the Alibaba Cloud SMS console.
 * *   For time-sensitive related messages, we recommend that you use the [SendMessageToGlobe](~~SendMessageToGlobe~~) operation to ensure that messages are delivered on time.
 * *   In each request, you can send messages to up to 1,000 mobile phone numbers.
 * ### QPS limit
 * You may call this operation only once per second. If the number of calls per second exceeds this limit, throttling will be triggered. This can potentially impact your business operations. Therefore, we recommend that you take note of this limit when making calls to this operation.
 *
 * @param request BatchSendMessageToGlobeRequest
 * @return BatchSendMessageToGlobeResponse
 */
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

/**
 * Metrics:
 * *   Requested OTP messages
 * *   Verified OTP messages
 * An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
 *
 * @param request ConversionDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ConversionDataResponse
 */
func (client *Client) ConversionDataWithOptions(request *ConversionDataRequest, runtime *util.RuntimeOptions) (_result *ConversionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversionRate)) {
		body["ConversionRate"] = request.ConversionRate
	}

	if !tea.BoolValue(util.IsUnset(request.ReportTime)) {
		body["ReportTime"] = request.ReportTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConversionData"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConversionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Metrics:
 * *   Requested OTP messages
 * *   Verified OTP messages
 * An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
 *
 * @param request ConversionDataRequest
 * @return ConversionDataResponse
 */
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

/**
 * ### QPS limit
 * You can call this operation up to 300 times per second. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request QueryMessageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMessageResponse
 */
func (client *Client) QueryMessageWithOptions(request *QueryMessageRequest, runtime *util.RuntimeOptions) (_result *QueryMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMessage"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### QPS limit
 * You can call this operation up to 300 times per second. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request QueryMessageRequest
 * @return QueryMessageResponse
 */
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

/**
 * ### Usage notes
 * You cannot call the SendMessageToGlobe operation to send messages to the Chinese mainland.
 * ### QPS limit
 * You may call this operation up to 300 times per second. If the number of calls per second exceeds this limit, throttling will be triggered. This can potentially impact your business operations. Therefore, we recommend that you take note of this limit when making calls to this operation.
 *
 * @param request SendMessageToGlobeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendMessageToGlobeResponse
 */
func (client *Client) SendMessageToGlobeWithOptions(request *SendMessageToGlobeRequest, runtime *util.RuntimeOptions) (_result *SendMessageToGlobeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityPeriod)) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessageToGlobe"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Usage notes
 * You cannot call the SendMessageToGlobe operation to send messages to the Chinese mainland.
 * ### QPS limit
 * You may call this operation up to 300 times per second. If the number of calls per second exceeds this limit, throttling will be triggered. This can potentially impact your business operations. Therefore, we recommend that you take note of this limit when making calls to this operation.
 *
 * @param request SendMessageToGlobeRequest
 * @return SendMessageToGlobeResponse
 */
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

/**
 * ### Usage notes
 * You can call the SendMessageWithTemplate operation to send messages only to the Chinese mainland.
 *
 * @param request SendMessageWithTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendMessageWithTemplateResponse
 */
func (client *Client) SendMessageWithTemplateWithOptions(request *SendMessageWithTemplateRequest, runtime *util.RuntimeOptions) (_result *SendMessageWithTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCode)) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParam)) {
		query["TemplateParam"] = request.TemplateParam
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.ValidityPeriod)) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessageWithTemplate"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageWithTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Usage notes
 * You can call the SendMessageWithTemplate operation to send messages only to the Chinese mainland.
 *
 * @param request SendMessageWithTemplateRequest
 * @return SendMessageWithTemplateResponse
 */
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

/**
 * Metrics:
 * *   Requested OTP messages
 * *   Verified OTP messages
 * An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
 * > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations:
 * > * Call the SmsConversion operation in an asynchronous manner by configuring queues or events.
 * > * Manually degrade your services or use a circuit breaker to automatically degrade services.
 *
 * @param request SmsConversionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SmsConversionResponse
 */
func (client *Client) SmsConversionWithOptions(request *SmsConversionRequest, runtime *util.RuntimeOptions) (_result *SmsConversionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversionTime)) {
		query["ConversionTime"] = request.ConversionTime
	}

	if !tea.BoolValue(util.IsUnset(request.Delivered)) {
		query["Delivered"] = request.Delivered
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["MessageId"] = request.MessageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsConversion"),
		Version:     tea.String("2018-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsConversionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Metrics:
 * *   Requested OTP messages
 * *   Verified OTP messages
 * An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
 * > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations:
 * > * Call the SmsConversion operation in an asynchronous manner by configuring queues or events.
 * > * Manually degrade your services or use a circuit breaker to automatically degrade services.
 *
 * @param request SmsConversionRequest
 * @return SmsConversionResponse
 */
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
