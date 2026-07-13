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
// 批量删除消息
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
// 批量删除消息
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
// 批量查看消息
//
// Description:
//
// ## 请求说明
//
// - 该接口用于批量查看指定队列中的消息，一次最多可以查看16条消息。
//
// - 使用此接口不会改变消息的状态，消息仍保持为Active状态。
//
// - 不支持长轮询功能。
//
// - 需要提供`queueName`作为路径参数，并通过查询参数设置`peekonly=true`及指定要查看的消息数量`numOfMessages`（范围在1到16之间）。
//
// - 成功响应将返回一个包含所请求消息详细信息的数组，包括但不限于消息ID、正文、入队时间等。
//
// - 如果指定的队列不存在或队列中没有可见消息，则会返回相应的错误码。
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
// 批量查看消息
//
// Description:
//
// ## 请求说明
//
// - 该接口用于批量查看指定队列中的消息，一次最多可以查看16条消息。
//
// - 使用此接口不会改变消息的状态，消息仍保持为Active状态。
//
// - 不支持长轮询功能。
//
// - 需要提供`queueName`作为路径参数，并通过查询参数设置`peekonly=true`及指定要查看的消息数量`numOfMessages`（范围在1到16之间）。
//
// - 成功响应将返回一个包含所请求消息详细信息的数组，包括但不限于消息ID、正文、入队时间等。
//
// - 如果指定的队列不存在或队列中没有可见消息，则会返回相应的错误码。
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
// 批量消费消息
//
// Description:
//
// ## 请求说明
//
// - 该操作会将取得的消息状态变为 Inactive，Inactive 持续时间由队列属性 `VisibilityTimeout` 决定。
//
// - 消费者需在 VisibilityTimeout 时间内调用 DeleteMessage 删除消息，否则消息会重新变为 Active。
//
// - 支持长轮询（Long Polling）：设置 `waitseconds > 0` 后，若队列为空则等待至有消息到达或超时返回。
//
//	Notice: 进入长轮询后，建议您降低外部调用长轮询的并发数，选择合适的长轮询时间。目前服务端会根据长轮询数量、长轮询等待时间、访问 IP 数量等因素动态调整长轮询防攻击的并发上限值。当队列无消息时，超过长轮询上限值的请求将无法被监听到，并直接返回 404 MessageNotExist（按请求量正常计费）。如果您有临时提升长轮询上限值的需求，请及时提交工单。
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
// 批量消费消息
//
// Description:
//
// ## 请求说明
//
// - 该操作会将取得的消息状态变为 Inactive，Inactive 持续时间由队列属性 `VisibilityTimeout` 决定。
//
// - 消费者需在 VisibilityTimeout 时间内调用 DeleteMessage 删除消息，否则消息会重新变为 Active。
//
// - 支持长轮询（Long Polling）：设置 `waitseconds > 0` 后，若队列为空则等待至有消息到达或超时返回。
//
//	Notice: 进入长轮询后，建议您降低外部调用长轮询的并发数，选择合适的长轮询时间。目前服务端会根据长轮询数量、长轮询等待时间、访问 IP 数量等因素动态调整长轮询防攻击的并发上限值。当队列无消息时，超过长轮询上限值的请求将无法被监听到，并直接返回 404 MessageNotExist（按请求量正常计费）。如果您有临时提升长轮询上限值的需求，请及时提交工单。
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
// 批量发送消息
//
// Description:
//
// ## 请求说明
//
// - 一次 Batch 请求的消息总大小不得超过 64 KB。
//
// - `BatchSendMessage` 与 `SendMessage` 使用相同的 URL 路径，通过 Body 结构区分：包含 `Messages` 数组即为批量发送，否则为单条发送。
//
// - 批量操作的返回结果可能同时包含成功和失败的子消息。
//
// - 每次请求中可以包含多个消息对象，每个消息对象可以设置不同的延迟时间和优先级。
//
// - 对于 FIFO 队列，可以通过 `MessageGroupId` 参数来指定消息分组 ID。
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
// 批量发送消息
//
// Description:
//
// ## 请求说明
//
// - 一次 Batch 请求的消息总大小不得超过 64 KB。
//
// - `BatchSendMessage` 与 `SendMessage` 使用相同的 URL 路径，通过 Body 结构区分：包含 `Messages` 数组即为批量发送，否则为单条发送。
//
// - 批量操作的返回结果可能同时包含成功和失败的子消息。
//
// - 每次请求中可以包含多个消息对象，每个消息对象可以设置不同的延迟时间和优先级。
//
// - 对于 FIFO 队列，可以通过 `MessageGroupId` 参数来指定消息分组 ID。
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
// 修改消息下次可消费时间
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
// 修改消息下次可消费时间
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
// 删除已消费消息
//
// @param request - DeleteMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageResponse
func (client *Client) DeleteMessageWithOptions(queueName *string, request *DeleteMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return DeleteMessageResponse
func (client *Client) DeleteMessage(queueName *string, request *DeleteMessageRequest) (_result *DeleteMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMessageResponse{}
	_body, _err := client.DeleteMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看消息
//
// Description:
//
// ## 请求说明
//
// - **PeekMessage*	- 接口用于查看队列顶部的消息，但不会改变消息的状态，消息仍处于 Active 状态可被正常消费。
//
// - 该接口不支持长轮询，如果队列为空则立即返回 `MessageNotExist` 错误。
//
// - 请求时必须设置 `peekonly=true` 参数以表明仅查看消息。
//
// @param request - PeekMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PeekMessageResponse
func (client *Client) PeekMessageWithOptions(queueName *string, request *PeekMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PeekMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - **PeekMessage*	- 接口用于查看队列顶部的消息，但不会改变消息的状态，消息仍处于 Active 状态可被正常消费。
//
// - 该接口不支持长轮询，如果队列为空则立即返回 `MessageNotExist` 错误。
//
// - 请求时必须设置 `peekonly=true` 参数以表明仅查看消息。
//
// @param request - PeekMessageRequest
//
// @return PeekMessageResponse
func (client *Client) PeekMessage(queueName *string, request *PeekMessageRequest) (_result *PeekMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PeekMessageResponse{}
	_body, _err := client.PeekMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布消息
//
// Description:
//
// ## 请求说明
//
// - 消息发布后会被推送到该 Topic 下所有 Subscription 的 Endpoint。
//
// - 推送到 Queue 和 HTTP Endpoint 时不需要设置 `MessageAttributes`。
//
// - 推送到邮件、短信或移动推送时需要设置对应的 `MessageAttributes` 子属性。
//
// - 消息内容建议事先进行 Base64 编码以避免特殊字符问题。
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
// 发布消息
//
// Description:
//
// ## 请求说明
//
// - 消息发布后会被推送到该 Topic 下所有 Subscription 的 Endpoint。
//
// - 推送到 Queue 和 HTTP Endpoint 时不需要设置 `MessageAttributes`。
//
// - 推送到邮件、短信或移动推送时需要设置对应的 `MessageAttributes` 子属性。
//
// - 消息内容建议事先进行 Base64 编码以避免特殊字符问题。
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
// 消费消息
//
// Description:
//
// ## 请求说明
//
// - `ReceiveMessage` 操作会将取得的消息状态变为 Inactive，Inactive 持续时间由队列属性 `VisibilityTimeout` 决定。
//
// - 消费者需在 VisibilityTimeout 时间内消费成功后调用 `DeleteMessage` 删除该消息，否则消息将重新变为 Active 状态，被再次消费。
//
// - 支持长轮询（Long Polling）：设置 `waitseconds > 0` 后，若队列为空则等待至有消息到达或超时返回。
//
//	Notice: 进入长轮询后，建议您降低外部调用长轮询的并发数，选择合适的长轮询时间。目前服务端会根据长轮询数量、长轮询等待时间、访问 IP 数量等因素动态调整长轮询防攻击的并发上限值。当队列无消息时，超过长轮询上限值的请求将无法被监听到，并直接返回 404 MessageNotExist（按请求量正常计费）。如果您有临时提升长轮询上限值的需求，请及时提交工单。
//
// @param request - ReceiveMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReceiveMessageResponse
func (client *Client) ReceiveMessageWithOptions(queueName *string, request *ReceiveMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReceiveMessageResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - `ReceiveMessage` 操作会将取得的消息状态变为 Inactive，Inactive 持续时间由队列属性 `VisibilityTimeout` 决定。
//
// - 消费者需在 VisibilityTimeout 时间内消费成功后调用 `DeleteMessage` 删除该消息，否则消息将重新变为 Active 状态，被再次消费。
//
// - 支持长轮询（Long Polling）：设置 `waitseconds > 0` 后，若队列为空则等待至有消息到达或超时返回。
//
//	Notice: 进入长轮询后，建议您降低外部调用长轮询的并发数，选择合适的长轮询时间。目前服务端会根据长轮询数量、长轮询等待时间、访问 IP 数量等因素动态调整长轮询防攻击的并发上限值。当队列无消息时，超过长轮询上限值的请求将无法被监听到，并直接返回 404 MessageNotExist（按请求量正常计费）。如果您有临时提升长轮询上限值的需求，请及时提交工单。
//
// @param request - ReceiveMessageRequest
//
// @return ReceiveMessageResponse
func (client *Client) ReceiveMessage(queueName *string, request *ReceiveMessageRequest) (_result *ReceiveMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReceiveMessageResponse{}
	_body, _err := client.ReceiveMessageWithOptions(queueName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送消息
//
// Description:
//
// ## 请求说明
//
// - **SendMessage*	- 接口用于将消息发送至指定队列。
//
// - 消息可以立即被消费或通过设置 `DelaySeconds` 参数来延迟消费。
//
// - 发送的消息可以指定优先级，数值越小表示优先级越高。
//
// - 对于 FIFO 队列，可以通过 `MessageGroupId` 来保证同一分组内消息的顺序投递。
//
// - 用户还可以自定义属性 `UserProperties`，以 JSON 格式字符串形式提供额外信息。
//
// - 当 `DelaySeconds` 大于 0 时，API 返回的 `ReceiptHandle` 可用来在消息变为 Active 状态前删除该延迟消息。
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
// 发送消息
//
// Description:
//
// ## 请求说明
//
// - **SendMessage*	- 接口用于将消息发送至指定队列。
//
// - 消息可以立即被消费或通过设置 `DelaySeconds` 参数来延迟消费。
//
// - 发送的消息可以指定优先级，数值越小表示优先级越高。
//
// - 对于 FIFO 队列，可以通过 `MessageGroupId` 来保证同一分组内消息的顺序投递。
//
// - 用户还可以自定义属性 `UserProperties`，以 JSON 格式字符串形式提供额外信息。
//
// - 当 `DelaySeconds` 大于 0 时，API 返回的 `ReceiptHandle` 可用来在消息变为 Active 状态前删除该延迟消息。
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
