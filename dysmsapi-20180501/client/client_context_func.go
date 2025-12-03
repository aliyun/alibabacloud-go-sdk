// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sends a message to multiple recipients in countries or regions outside the Chinese mainland.
//
// Description:
//
// ## Usage notes
//
//   - You cannot call the BatchSendMessageToGlobe operation to send messages to the Chinese mainland.
//
//   - You can call the BatchSendMessageToGlobe operation to send notifications and promotional messages to a small number of mobile phone numbers at a time. To send messages to a large number of mobile phone numbers at a time, use the mass messaging feature of the SMS console.
//
//   - To ensure that messages can be sent on time, call the [SendMessageToGlobe](https://help.aliyun.com/document_detail/406238.html) operation.
//
//   - In each request, you can send messages to up to 1,000 mobile phone numbers.
//
// @param request - BatchSendMessageToGlobeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendMessageToGlobeResponse
func (client *Client) BatchSendMessageToGlobeWithContext(ctx context.Context, request *BatchSendMessageToGlobeRequest, runtime *dara.RuntimeOptions) (_result *BatchSendMessageToGlobeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSendMessageToGlobe"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSendMessageToGlobeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API, sends conversion data to the Alibaba SMS service.
//
// Description:
//
// Metrics:
//
//   - Requested OTP messages
//
//   - Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// @param request - ConversionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversionDataResponse
func (client *Client) ConversionDataWithContext(ctx context.Context, request *ConversionDataRequest, runtime *dara.RuntimeOptions) (_result *ConversionDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversionRate) {
		body["ConversionRate"] = request.ConversionRate
	}

	if !dara.IsNil(request.ReportTime) {
		body["ReportTime"] = request.ReportTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConversionData"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConversionDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the delivery report of a message.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
//
// @param request - QueryMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageResponse
func (client *Client) QueryMessageWithContext(ctx context.Context, request *QueryMessageRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMessage"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to regions outside the Chinese mainland.
//
// Description:
//
// ### [](#)
//
// The SendMessageToGlobe API operation does not support message delivery to the Chinese mainland.
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to 2,000 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - SendMessageToGlobeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageToGlobeResponse
func (client *Client) SendMessageToGlobeWithContext(ctx context.Context, request *SendMessageToGlobeRequest, runtime *dara.RuntimeOptions) (_result *SendMessageToGlobeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageToGlobe"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageToGlobeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to the Chinese mainland by using a message template.
//
// Description:
//
// ### Usage notes
//
// You can call the SendMessageWithTemplate operation to send messages only to the Chinese mainland.
//
// @param request - SendMessageWithTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageWithTemplateResponse
func (client *Client) SendMessageWithTemplateWithContext(ctx context.Context, request *SendMessageWithTemplateRequest, runtime *dara.RuntimeOptions) (_result *SendMessageWithTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.SmsUpExtendCode) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateParam) {
		query["TemplateParam"] = request.TemplateParam
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageWithTemplate"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageWithTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delivers one-time password (OTP) message statuses to Alibaba Cloud, which calculates and monitors OTP conversion rates.
//
// Description:
//
// Metrics:
//
//   - Requested OTP messages
//
//   - Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations:
//
// >	- Call the SmsConversion operation in an asynchronous manner by configuring queues or events.
//
// >	- Manually degrade your services or use a circuit breaker to automatically degrade services.
//
// @param request - SmsConversionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsConversionResponse
func (client *Client) SmsConversionWithContext(ctx context.Context, request *SmsConversionRequest, runtime *dara.RuntimeOptions) (_result *SmsConversionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversionTime) {
		query["ConversionTime"] = request.ConversionTime
	}

	if !dara.IsNil(request.Delivered) {
		query["Delivered"] = request.Delivered
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmsConversion"),
		Version:     dara.String("2018-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmsConversionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
