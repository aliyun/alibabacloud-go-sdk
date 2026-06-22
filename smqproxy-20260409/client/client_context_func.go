// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 批量删除消息
//
// @param request - BatchDeleteMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteMessageResponse
func (client *Client) BatchDeleteMessageWithContext(ctx context.Context, queueName *string, request *BatchDeleteMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeleteMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReceiptHandles) {
		body["ReceiptHandles"] = request.ReceiptHandles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查看消息
//
// @param request - BatchPeekMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPeekMessageResponse
func (client *Client) BatchPeekMessageWithContext(ctx context.Context, queueName *string, request *BatchPeekMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchPeekMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NumOfMessages) {
		query["numOfMessages"] = request.NumOfMessages
	}

	if !dara.IsNil(request.Peekonly) {
		query["peekonly"] = request.Peekonly
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchPeekMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPeekMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量消费消息
//
// @param request - BatchReceiveMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchReceiveMessageResponse
func (client *Client) BatchReceiveMessageWithContext(ctx context.Context, queueName *string, request *BatchReceiveMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchReceiveMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NumOfMessages) {
		query["numOfMessages"] = request.NumOfMessages
	}

	if !dara.IsNil(request.Waitseconds) {
		query["waitseconds"] = request.Waitseconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchReceiveMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchReceiveMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量发送消息
//
// @param request - BatchSendMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendMessageResponse
func (client *Client) BatchSendMessageWithContext(ctx context.Context, queueName *string, request *BatchSendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchSendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSendMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSendMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改消息下次可消费时间
//
// @param request - ChangeMessageVisibilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMessageVisibilityResponse
func (client *Client) ChangeMessageVisibilityWithContext(ctx context.Context, queueName *string, request *ChangeMessageVisibilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeMessageVisibilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReceiptHandle) {
		query["receiptHandle"] = request.ReceiptHandle
	}

	if !dara.IsNil(request.VisibilityTimeout) {
		query["visibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMessageVisibility"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMessageVisibilityResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除已消费消息
//
// @param request - DeleteMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageResponse
func (client *Client) DeleteMessageWithContext(ctx context.Context, queueName *string, request *DeleteMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReceiptHandle) {
		query["ReceiptHandle"] = request.ReceiptHandle
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看消息
//
// @param request - PeekMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PeekMessageResponse
func (client *Client) PeekMessageWithContext(ctx context.Context, queueName *string, request *PeekMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PeekMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Peekonly) {
		query["peekonly"] = request.Peekonly
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PeekMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PeekMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布消息
//
// @param request - PublishMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishMessageResponse
func (client *Client) PublishMessageWithContext(ctx context.Context, topicName *string, request *PublishMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MessageAttributes) {
		body["MessageAttributes"] = request.MessageAttributes
	}

	if !dara.IsNil(request.MessageBody) {
		body["MessageBody"] = request.MessageBody
	}

	if !dara.IsNil(request.MessageTag) {
		body["MessageTag"] = request.MessageTag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/topics/" + dara.StringValue(topicName) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消费消息
//
// @param request - ReceiveMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReceiveMessageResponse
func (client *Client) ReceiveMessageWithContext(ctx context.Context, queueName *string, request *ReceiveMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReceiveMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Waitseconds) {
		query["waitseconds"] = request.Waitseconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReceiveMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReceiveMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送消息
//
// @param request - SendMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithContext(ctx context.Context, queueName *string, request *SendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DelaySeconds) {
		body["DelaySeconds"] = request.DelaySeconds
	}

	if !dara.IsNil(request.MessageBody) {
		body["MessageBody"] = request.MessageBody
	}

	if !dara.IsNil(request.MessageGroupId) {
		body["MessageGroupId"] = request.MessageGroupId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.UserProperties) {
		body["UserProperties"] = request.UserProperties
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("2026-04-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queues/" + dara.StringValue(queueName) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
