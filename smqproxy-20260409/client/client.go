// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-mns/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.ProductId = dara.String("SMQProxy")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	return nil
}

// Summary:
//
// Deletes messages from a queue in a batch. A maximum of 16 messages can be deleted at a time.
//
// Description:
//
// ## Request description
//
// - Use this operation together with BatchReceiveMessage.
//
// - The result of a batch operation may contain both successful and failed sub-requests. The HTTP status code is 204 only when all sub-requests succeed.
//
// - A maximum of 16 messages can be deleted per request.
//
// - Provide the list of `ReceiptHandle` values returned from the previous consumption to identify the messages to delete.
//
// - This operation succeeds only when executed before `NextVisibleTime`. After the timeout, the message becomes Active again, the ReceiptHandle becomes invalid, and you must consume the message again to obtain a new handle.
//
// @param request - BatchDeleteMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteMessageResponse
func (client *Client) BatchDeleteMessageWithOptions(queueName *string, request *BatchDeleteMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeleteMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes messages from a queue in a batch. A maximum of 16 messages can be deleted at a time.
//
// Description:
//
// ## Request description
//
// - Use this operation together with BatchReceiveMessage.
//
// - The result of a batch operation may contain both successful and failed sub-requests. The HTTP status code is 204 only when all sub-requests succeed.
//
// - A maximum of 16 messages can be deleted per request.
//
// - Provide the list of `ReceiptHandle` values returned from the previous consumption to identify the messages to delete.
//
// - This operation succeeds only when executed before `NextVisibleTime`. After the timeout, the message becomes Active again, the ReceiptHandle becomes invalid, and you must consume the message again to obtain a new handle.
//
// @param request - BatchDeleteMessageRequest
//
// @return BatchDeleteMessageResponse
func (client *Client) BatchDeleteMessage(queueName *string, request *BatchDeleteMessageRequest) (_result *BatchDeleteMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteMessageResponse{}
	_body, _err := client.BatchDeleteMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch peeks at messages in a queue without changing the message status.
//
// Description:
//
// ## Request description
//
// - This operation batch peeks at messages in a specified queue. You can peek at up to 16 messages at a time.
//
// - This operation does not change the message status. Messages remain in the Active state.
//
// - Long polling is not supported.
//
// - Provide `queueName` as a path parameter, and set the query parameter `peekonly=true` along with `numOfMessages` (range: 1 to 16) to specify the number of messages to peek at.
//
// - A successful response returns an array containing the details of the requested messages, including but not limited to the message ID, body, and enqueue time.
//
// - If the specified queue does not exist or no visible messages are in the queue, the corresponding error code is returned.
//
// @param request - BatchPeekMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPeekMessageResponse
func (client *Client) BatchPeekMessageWithOptions(queueName *string, request *BatchPeekMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchPeekMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch peeks at messages in a queue without changing the message status.
//
// Description:
//
// ## Request description
//
// - This operation batch peeks at messages in a specified queue. You can peek at up to 16 messages at a time.
//
// - This operation does not change the message status. Messages remain in the Active state.
//
// - Long polling is not supported.
//
// - Provide `queueName` as a path parameter, and set the query parameter `peekonly=true` along with `numOfMessages` (range: 1 to 16) to specify the number of messages to peek at.
//
// - A successful response returns an array containing the details of the requested messages, including but not limited to the message ID, body, and enqueue time.
//
// - If the specified queue does not exist or no visible messages are in the queue, the corresponding error code is returned.
//
// @param request - BatchPeekMessageRequest
//
// @return BatchPeekMessageResponse
func (client *Client) BatchPeekMessage(queueName *string, request *BatchPeekMessageRequest) (_result *BatchPeekMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchPeekMessageResponse{}
	_body, _err := client.BatchPeekMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch receives up to 16 messages from a specified queue for consumption.
//
// Description:
//
// ## Operation description
//
// - This operation changes the status of received messages to Inactive. The Inactive duration is determined by the `VisibilityTimeout` attribute of the queue.
//
// - Consumers must call DeleteMessage to delete messages within the VisibilityTimeout period. Otherwise, messages become Active again.
//
// - Long polling is supported: if you set `waitseconds > 0` and the queue is empty, the request waits until a message arrives or the timeout expires.
//
//	Notice: After long polling is enabled, reduce the concurrency of external long polling calls and select an appropriate long polling duration. The server dynamically adjusts the concurrency limit for long polling based on factors such as the number of long polling requests, long polling wait time, and number of access IP addresses. When the queue has no messages, requests that exceed the long polling limit cannot be listened to and directly return 404 MessageNotExist (billed normally by request count). If you need to temporarily increase the long polling limit, submit a ticket promptly.
//
// @param request - BatchReceiveMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchReceiveMessageResponse
func (client *Client) BatchReceiveMessageWithOptions(queueName *string, request *BatchReceiveMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchReceiveMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch receives up to 16 messages from a specified queue for consumption.
//
// Description:
//
// ## Operation description
//
// - This operation changes the status of received messages to Inactive. The Inactive duration is determined by the `VisibilityTimeout` attribute of the queue.
//
// - Consumers must call DeleteMessage to delete messages within the VisibilityTimeout period. Otherwise, messages become Active again.
//
// - Long polling is supported: if you set `waitseconds > 0` and the queue is empty, the request waits until a message arrives or the timeout expires.
//
//	Notice: After long polling is enabled, reduce the concurrency of external long polling calls and select an appropriate long polling duration. The server dynamically adjusts the concurrency limit for long polling based on factors such as the number of long polling requests, long polling wait time, and number of access IP addresses. When the queue has no messages, requests that exceed the long polling limit cannot be listened to and directly return 404 MessageNotExist (billed normally by request count). If you need to temporarily increase the long polling limit, submit a ticket promptly.
//
// @param request - BatchReceiveMessageRequest
//
// @return BatchReceiveMessageResponse
func (client *Client) BatchReceiveMessage(queueName *string, request *BatchReceiveMessageRequest) (_result *BatchReceiveMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchReceiveMessageResponse{}
	_body, _err := client.BatchReceiveMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a batch of up to 16 messages to a specified queue, with a total size not exceeding 64 KB.
//
// Description:
//
// ## Request description
//
// - The total size of messages in a single batch request must not exceed 64 KB.
//
// - `BatchSendMessage` and `SendMessage` use the same URL path. They are distinguished by the body structure: if the body contains a `Messages` array, it is a batch send operation. Otherwise, it is a single message send operation.
//
// - The response of a batch operation may contain both successful and failed sub-messages.
//
// - Each request can contain multiple message objects. Each message object can have a different delay time and priority.
//
// - For FIFO queues, you can use the `MessageGroupId` parameter to specify a message group ID.
//
// @param request - BatchSendMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendMessageResponse
func (client *Client) BatchSendMessageWithOptions(queueName *string, request *BatchSendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchSendMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a batch of up to 16 messages to a specified queue, with a total size not exceeding 64 KB.
//
// Description:
//
// ## Request description
//
// - The total size of messages in a single batch request must not exceed 64 KB.
//
// - `BatchSendMessage` and `SendMessage` use the same URL path. They are distinguished by the body structure: if the body contains a `Messages` array, it is a batch send operation. Otherwise, it is a single message send operation.
//
// - The response of a batch operation may contain both successful and failed sub-messages.
//
// - Each request can contain multiple message objects. Each message object can have a different delay time and priority.
//
// - For FIFO queues, you can use the `MessageGroupId` parameter to specify a message group ID.
//
// @param request - BatchSendMessageRequest
//
// @return BatchSendMessageResponse
func (client *Client) BatchSendMessage(queueName *string, request *BatchSendMessageRequest) (_result *BatchSendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchSendMessageResponse{}
	_body, _err := client.BatchSendMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Extends or shortens the invisibility period of a message that is in the Inactive state.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to modify the VisibilityTimeout of a message that has been consumed and is currently in the Inactive state. The VisibilityTimeout is the countdown from the current moment until the message becomes Active again.
//
// - A successful invocation returns a new `ReceiptHandle`. Use this new handle for subsequent delete or visibility timeout modification operations on this message.
//
// - This operation can only be invoked successfully before `NextVisibleTime`. Otherwise, a fault is returned.
//
// - Common scenarios include when a consumer requires a longer processing time and needs to extend the VisibilityTimeout to prevent the message from being consumed again.
//
// @param request - ChangeMessageVisibilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMessageVisibilityResponse
func (client *Client) ChangeMessageVisibilityWithOptions(queueName *string, request *ChangeMessageVisibilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeMessageVisibilityResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extends or shortens the invisibility period of a message that is in the Inactive state.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to modify the VisibilityTimeout of a message that has been consumed and is currently in the Inactive state. The VisibilityTimeout is the countdown from the current moment until the message becomes Active again.
//
// - A successful invocation returns a new `ReceiptHandle`. Use this new handle for subsequent delete or visibility timeout modification operations on this message.
//
// - This operation can only be invoked successfully before `NextVisibleTime`. Otherwise, a fault is returned.
//
// - Common scenarios include when a consumer requires a longer processing time and needs to extend the VisibilityTimeout to prevent the message from being consumed again.
//
// @param request - ChangeMessageVisibilityRequest
//
// @return ChangeMessageVisibilityResponse
func (client *Client) ChangeMessageVisibility(queueName *string, request *ChangeMessageVisibilityRequest) (_result *ChangeMessageVisibilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeMessageVisibilityResponse{}
	_body, _err := client.ChangeMessageVisibilityWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a message to a specified topic. The message is pushed to all subscribed endpoints.
//
// Description:
//
// ## Operation description
//
// - After a message is published, it is pushed to the endpoints of all subscriptions under the topic.
//
// - You do not need to set `MessageAttributes` when pushing messages to queues or HTTP endpoints.
//
// - You must set the corresponding `MessageAttributes` sub-properties when pushing messages to email, SMS, or mobile push endpoints.
//
// - Base64-encode the message content beforehand to avoid special character issues.
//
// @param request - PublishMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishMessageResponse
func (client *Client) PublishMessageWithOptions(topicName *string, request *PublishMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishMessageResponse, _err error) {
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

	if !dara.IsNil(request.MessageGroupId) {
		body["MessageGroupId"] = request.MessageGroupId
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a message to a specified topic. The message is pushed to all subscribed endpoints.
//
// Description:
//
// ## Operation description
//
// - After a message is published, it is pushed to the endpoints of all subscriptions under the topic.
//
// - You do not need to set `MessageAttributes` when pushing messages to queues or HTTP endpoints.
//
// - You must set the corresponding `MessageAttributes` sub-properties when pushing messages to email, SMS, or mobile push endpoints.
//
// - Base64-encode the message content beforehand to avoid special character issues.
//
// @param request - PublishMessageRequest
//
// @return PublishMessageResponse
func (client *Client) PublishMessage(topicName *string, request *PublishMessageRequest) (_result *PublishMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishMessageResponse{}
	_body, _err := client.PublishMessageWithOptions(topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a message to a specified queue with support for delay and priority settings.
//
// Description:
//
// ## Operation description
//
// - The **SendMessage*	- operation sends a message to a specified queue.
//
// - Messages can be consumed immediately or delayed by setting the `DelaySeconds` parameter.
//
// - You can specify a priority for the message. A smaller value indicates a higher priority.
//
// - For FIFO queues, use `MessageGroupId` to ensure ordered delivery of messages within the same group.
//
// - You can also set custom attributes in `UserProperties` as a JSON-formatted string to provide additional information.
//
// - When `DelaySeconds` is greater than 0, the `ReceiptHandle` returned by the API can be used to delete the delayed message before it becomes Active.
//
// @param request - SendMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(queueName *string, request *SendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to a specified queue with support for delay and priority settings.
//
// Description:
//
// ## Operation description
//
// - The **SendMessage*	- operation sends a message to a specified queue.
//
// - Messages can be consumed immediately or delayed by setting the `DelaySeconds` parameter.
//
// - You can specify a priority for the message. A smaller value indicates a higher priority.
//
// - For FIFO queues, use `MessageGroupId` to ensure ordered delivery of messages within the same group.
//
// - You can also set custom attributes in `UserProperties` as a JSON-formatted string to provide additional information.
//
// - When `DelaySeconds` is greater than 0, the `ReceiptHandle` returned by the API can be used to delete the delayed message before it becomes Active.
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(queueName *string, request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
