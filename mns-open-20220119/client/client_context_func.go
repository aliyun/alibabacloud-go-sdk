// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// You can call this operation to add one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param tmpReq - AuthorizeEndpointAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeEndpointAclResponse
func (client *Client) AuthorizeEndpointAclWithContext(ctx context.Context, tmpReq *AuthorizeEndpointAclRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeEndpointAclResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AuthorizeEndpointAclShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CidrList) {
		request.CidrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CidrList, dara.String("CidrList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AclStrategy) {
		query["AclStrategy"] = request.AclStrategy
	}

	if !dara.IsNil(request.CidrListShrink) {
		query["CidrList"] = request.CidrListShrink
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeEndpointAcl"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeEndpointAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建事件规则
//
// @param tmpReq - CreateEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventRuleResponse
func (client *Client) CreateEventRuleWithContext(ctx context.Context, tmpReq *CreateEventRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateEventRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateEventRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Endpoint) {
		request.EndpointShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoint, dara.String("Endpoint"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventTypes) {
		request.EventTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventTypes, dara.String("EventTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MatchRules) {
		request.MatchRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchRules, dara.String("MatchRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeliveryMode) {
		query["DeliveryMode"] = request.DeliveryMode
	}

	if !dara.IsNil(request.EndpointShrink) {
		query["Endpoint"] = request.EndpointShrink
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.EventTypesShrink) {
		query["EventTypes"] = request.EventTypesShrink
	}

	if !dara.IsNil(request.MatchRulesShrink) {
		query["MatchRules"] = request.MatchRulesShrink
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventRule"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue.
//
// @param tmpReq - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithContext(ctx context.Context, tmpReq *CreateQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DlqPolicy) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, dara.String("DlqPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantRateLimitPolicy) {
		request.TenantRateLimitPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantRateLimitPolicy, dara.String("TenantRateLimitPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DelaySeconds) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !dara.IsNil(request.DlqPolicyShrink) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !dara.IsNil(request.EnableLogging) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !dara.IsNil(request.MaximumMessageSize) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !dara.IsNil(request.MessageRetentionPeriod) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !dara.IsNil(request.PollingWaitSeconds) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TenantRateLimitPolicyShrink) {
		query["TenantRateLimitPolicy"] = request.TenantRateLimitPolicyShrink
	}

	if !dara.IsNil(request.VisibilityTimeout) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueue"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest, runtime *dara.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableLogging) {
		body["EnableLogging"] = request.EnableLogging
	}

	if !dara.IsNil(request.MaxMessageSize) {
		body["MaxMessageSize"] = request.MaxMessageSize
	}

	if !dara.IsNil(request.TopicName) {
		body["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTopic"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除事件规则
//
// @param request - DeleteEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRuleResponse
func (client *Client) DeleteEventRuleWithContext(ctx context.Context, request *DeleteEventRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRule"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithContext(ctx context.Context, request *DeleteQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQueue"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTopic"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to disenable the endpoint of a type. After the endpoint is disabled, all requests from the endpoint are blocked and an error is returned.
//
// @param request - DisableEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEndpointResponse
func (client *Client) DisableEndpointWithContext(ctx context.Context, request *DisableEndpointRequest, runtime *dara.RuntimeOptions) (_result *DisableEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableEndpoint"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to enable the endpoint of a type. If the endpoint is enabled, requests from the endpoint that are included in the access control lists (ACLs) are not blocked.
//
// @param request - EnableEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEndpointResponse
func (client *Client) EnableEndpointWithContext(ctx context.Context, request *EnableEndpointRequest, runtime *dara.RuntimeOptions) (_result *EnableEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableEndpoint"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetEndpointAttribute
//
// @param request - GetEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEndpointAttributeResponse
func (client *Client) GetEndpointAttributeWithContext(ctx context.Context, request *GetEndpointAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetEndpointAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEndpointAttribute"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEndpointAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an existing queue.
//
// @param request - GetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueAttributesResponse
func (client *Client) GetQueueAttributesWithContext(ctx context.Context, request *GetQueueAttributesRequest, runtime *dara.RuntimeOptions) (_result *GetQueueAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a subscription.
//
// @param request - GetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionAttributesResponse
func (client *Client) GetSubscriptionAttributesWithContext(ctx context.Context, request *GetSubscriptionAttributesRequest, runtime *dara.RuntimeOptions) (_result *GetSubscriptionAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SubscriptionName) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscriptionAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a topic.
//
// @param request - GetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicAttributesResponse
func (client *Client) GetTopicAttributesWithContext(ctx context.Context, request *GetTopicAttributesRequest, runtime *dara.RuntimeOptions) (_result *GetTopicAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all queues that belong to an Alibaba Cloud account. The queues are displayed by page.
//
// @param request - ListQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueResponse
func (client *Client) ListQueueWithContext(ctx context.Context, request *ListQueueRequest, runtime *dara.RuntimeOptions) (_result *ListQueueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueue"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all subscriptions to a topic. The subscriptions are displayed by page.
//
// @param request - ListSubscriptionByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscriptionByTopicResponse
func (client *Client) ListSubscriptionByTopicWithContext(ctx context.Context, request *ListSubscriptionByTopicRequest, runtime *dara.RuntimeOptions) (_result *ListSubscriptionByTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.EndpointValue) {
		query["EndpointValue"] = request.EndpointValue
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SubscriptionName) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubscriptionByTopic"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubscriptionByTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topics that belong to an Alibaba Cloud account. The topics are displayed by page.
//
// @param request - ListTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicResponse
func (client *Client) ListTopicWithContext(ctx context.Context, request *ListTopicRequest, runtime *dara.RuntimeOptions) (_result *ListTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopic"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to delete one or more rules of access control lists (ACLs) for the endpoint of a type.
//
// @param tmpReq - RevokeEndpointAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeEndpointAclResponse
func (client *Client) RevokeEndpointAclWithContext(ctx context.Context, tmpReq *RevokeEndpointAclRequest, runtime *dara.RuntimeOptions) (_result *RevokeEndpointAclResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RevokeEndpointAclShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CidrList) {
		request.CidrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CidrList, dara.String("CidrList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AclStrategy) {
		query["AclStrategy"] = request.AclStrategy
	}

	if !dara.IsNil(request.CidrListShrink) {
		query["CidrList"] = request.CidrListShrink
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeEndpointAcl"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeEndpointAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a queue.
//
// @param tmpReq - SetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetQueueAttributesResponse
func (client *Client) SetQueueAttributesWithContext(ctx context.Context, tmpReq *SetQueueAttributesRequest, runtime *dara.RuntimeOptions) (_result *SetQueueAttributesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetQueueAttributesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DlqPolicy) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, dara.String("DlqPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantRateLimitPolicy) {
		request.TenantRateLimitPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantRateLimitPolicy, dara.String("TenantRateLimitPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DelaySeconds) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !dara.IsNil(request.DlqPolicyShrink) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !dara.IsNil(request.EnableLogging) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !dara.IsNil(request.MaximumMessageSize) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !dara.IsNil(request.MessageRetentionPeriod) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !dara.IsNil(request.PollingWaitSeconds) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.TenantRateLimitPolicyShrink) {
		query["TenantRateLimitPolicy"] = request.TenantRateLimitPolicyShrink
	}

	if !dara.IsNil(request.VisibilityTimeout) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetQueueAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetQueueAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a subscription.
//
// @param tmpReq - SetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSubscriptionAttributesResponse
func (client *Client) SetSubscriptionAttributesWithContext(ctx context.Context, tmpReq *SetSubscriptionAttributesRequest, runtime *dara.RuntimeOptions) (_result *SetSubscriptionAttributesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetSubscriptionAttributesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DlqPolicy) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, dara.String("DlqPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantRateLimitPolicy) {
		request.TenantRateLimitPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantRateLimitPolicy, dara.String("TenantRateLimitPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DlqPolicyShrink) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !dara.IsNil(request.NotifyStrategy) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !dara.IsNil(request.SubscriptionName) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !dara.IsNil(request.TenantRateLimitPolicyShrink) {
		query["TenantRateLimitPolicy"] = request.TenantRateLimitPolicyShrink
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSubscriptionAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSubscriptionAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a topic.
//
// @param request - SetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTopicAttributesResponse
func (client *Client) SetTopicAttributesWithContext(ctx context.Context, request *SetTopicAttributesRequest, runtime *dara.RuntimeOptions) (_result *SetTopicAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableLogging) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !dara.IsNil(request.MaxMessageSize) {
		query["MaxMessageSize"] = request.MaxMessageSize
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTopicAttributes"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTopicAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subscription to a topic.
//
// @param tmpReq - SubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeResponse
func (client *Client) SubscribeWithContext(ctx context.Context, tmpReq *SubscribeRequest, runtime *dara.RuntimeOptions) (_result *SubscribeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubscribeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DlqPolicy) {
		request.DlqPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DlqPolicy, dara.String("DlqPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DmAttributes) {
		request.DmAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DmAttributes, dara.String("DmAttributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DysmsAttributes) {
		request.DysmsAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DysmsAttributes, dara.String("DysmsAttributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KafkaAttributes) {
		request.KafkaAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KafkaAttributes, dara.String("KafkaAttributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantRateLimitPolicy) {
		request.TenantRateLimitPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantRateLimitPolicy, dara.String("TenantRateLimitPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DlqPolicyShrink) {
		query["DlqPolicy"] = request.DlqPolicyShrink
	}

	if !dara.IsNil(request.DmAttributesShrink) {
		query["DmAttributes"] = request.DmAttributesShrink
	}

	if !dara.IsNil(request.DysmsAttributesShrink) {
		query["DysmsAttributes"] = request.DysmsAttributesShrink
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.KafkaAttributesShrink) {
		query["KafkaAttributes"] = request.KafkaAttributesShrink
	}

	if !dara.IsNil(request.MessageTag) {
		query["MessageTag"] = request.MessageTag
	}

	if !dara.IsNil(request.NotifyContentFormat) {
		query["NotifyContentFormat"] = request.NotifyContentFormat
	}

	if !dara.IsNil(request.NotifyStrategy) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.StsRoleArn) {
		query["StsRoleArn"] = request.StsRoleArn
	}

	if !dara.IsNil(request.SubscriptionName) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !dara.IsNil(request.TenantRateLimitPolicyShrink) {
		query["TenantRateLimitPolicy"] = request.TenantRateLimitPolicyShrink
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Subscribe"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subscription.
//
// @param request - UnsubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeResponse
func (client *Client) UnsubscribeWithContext(ctx context.Context, request *UnsubscribeRequest, runtime *dara.RuntimeOptions) (_result *UnsubscribeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SubscriptionName) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Unsubscribe"),
		Version:     dara.String("2022-01-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnsubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
