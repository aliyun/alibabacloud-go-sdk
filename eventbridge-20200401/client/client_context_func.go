// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 检查账号是否存在ServiceLinkedRole授权
//
// @param request - CheckServiceLinkedRoleForProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleForProductResponse
func (client *Client) CheckServiceLinkedRoleForProductWithContext(ctx context.Context, request *CheckServiceLinkedRoleForProductRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleForProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRoleForProduct"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleForProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API destination.
//
// Description:
//
// You can call this API operation to create an API destination.
//
// @param tmpReq - CreateApiDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiDestinationResponse
func (client *Client) CreateApiDestinationWithContext(ctx context.Context, tmpReq *CreateApiDestinationRequest, runtime *dara.RuntimeOptions) (_result *CreateApiDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApiDestinationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HttpApiParameters) {
		request.HttpApiParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpApiParameters, dara.String("HttpApiParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDestinationName) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HttpApiParametersShrink) {
		query["HttpApiParameters"] = request.HttpApiParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiDestination"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a connection.
//
// Description:
//
// You can call this API operation to create a connection.
//
// @param tmpReq - CreateConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnectionResponse
func (client *Client) CreateConnectionWithContext(ctx context.Context, tmpReq *CreateConnectionRequest, runtime *dara.RuntimeOptions) (_result *CreateConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthParameters) {
		request.AuthParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthParameters, dara.String("AuthParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NetworkParameters) {
		request.NetworkParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkParameters, dara.String("NetworkParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthParametersShrink) {
		query["AuthParameters"] = request.AuthParametersShrink
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkParametersShrink) {
		query["NetworkParameters"] = request.NetworkParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConnection"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an event bus.
//
// Description:
//
// Creates an event bus.
//
// @param request - CreateEventBusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventBusResponse
func (client *Client) CreateEventBusWithContext(ctx context.Context, request *CreateEventBusRequest, runtime *dara.RuntimeOptions) (_result *CreateEventBusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventBus"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventBusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an event source.
//
// Description:
//
// You can call this operation to create an event source.
//
// @param tmpReq - CreateEventSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSourceResponse
func (client *Client) CreateEventSourceWithContext(ctx context.Context, tmpReq *CreateEventSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateEventSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEventSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExternalSourceConfig) {
		request.ExternalSourceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalSourceConfig, dara.String("ExternalSourceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceHttpEventParameters) {
		request.SourceHttpEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceHttpEventParameters, dara.String("SourceHttpEventParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceKafkaParameters) {
		request.SourceKafkaParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceKafkaParameters, dara.String("SourceKafkaParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceMNSParameters) {
		request.SourceMNSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMNSParameters, dara.String("SourceMNSParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceOSSEventParameters) {
		request.SourceOSSEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceOSSEventParameters, dara.String("SourceOSSEventParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceRabbitMQParameters) {
		request.SourceRabbitMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRabbitMQParameters, dara.String("SourceRabbitMQParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceRocketMQParameters) {
		request.SourceRocketMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRocketMQParameters, dara.String("SourceRocketMQParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceSLSParameters) {
		request.SourceSLSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSLSParameters, dara.String("SourceSLSParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceScheduledEventParameters) {
		request.SourceScheduledEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceScheduledEventParameters, dara.String("SourceScheduledEventParameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		body["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventSourceName) {
		body["EventSourceName"] = request.EventSourceName
	}

	if !dara.IsNil(request.ExternalSourceConfigShrink) {
		body["ExternalSourceConfig"] = request.ExternalSourceConfigShrink
	}

	if !dara.IsNil(request.ExternalSourceType) {
		body["ExternalSourceType"] = request.ExternalSourceType
	}

	if !dara.IsNil(request.LinkedExternalSource) {
		body["LinkedExternalSource"] = request.LinkedExternalSource
	}

	if !dara.IsNil(request.SourceHttpEventParametersShrink) {
		body["SourceHttpEventParameters"] = request.SourceHttpEventParametersShrink
	}

	if !dara.IsNil(request.SourceKafkaParametersShrink) {
		body["SourceKafkaParameters"] = request.SourceKafkaParametersShrink
	}

	if !dara.IsNil(request.SourceMNSParametersShrink) {
		body["SourceMNSParameters"] = request.SourceMNSParametersShrink
	}

	if !dara.IsNil(request.SourceOSSEventParametersShrink) {
		body["SourceOSSEventParameters"] = request.SourceOSSEventParametersShrink
	}

	if !dara.IsNil(request.SourceRabbitMQParametersShrink) {
		body["SourceRabbitMQParameters"] = request.SourceRabbitMQParametersShrink
	}

	if !dara.IsNil(request.SourceRocketMQParametersShrink) {
		body["SourceRocketMQParameters"] = request.SourceRocketMQParametersShrink
	}

	if !dara.IsNil(request.SourceSLSParametersShrink) {
		body["SourceSLSParameters"] = request.SourceSLSParametersShrink
	}

	if !dara.IsNil(request.SourceScheduledEventParametersShrink) {
		body["SourceScheduledEventParameters"] = request.SourceScheduledEventParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventSource"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an event stream.
//
// Description:
//
// You can call this API operation to create an event stream.
//
// @param tmpReq - CreateEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventStreamingResponse
func (client *Client) CreateEventStreamingWithContext(ctx context.Context, tmpReq *CreateEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *CreateEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEventStreamingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RunOptions) {
		request.RunOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RunOptions, dara.String("RunOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sink) {
		request.SinkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sink, dara.String("Sink"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Source) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, dara.String("Source"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Transforms) {
		request.TransformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transforms, dara.String("Transforms"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	if !dara.IsNil(request.FilterPattern) {
		body["FilterPattern"] = request.FilterPattern
	}

	if !dara.IsNil(request.RunOptionsShrink) {
		body["RunOptions"] = request.RunOptionsShrink
	}

	if !dara.IsNil(request.SinkShrink) {
		body["Sink"] = request.SinkShrink
	}

	if !dara.IsNil(request.SourceShrink) {
		body["Source"] = request.SourceShrink
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TransformsShrink) {
		body["Transforms"] = request.TransformsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an event rule.
//
// Description:
//
// You can call this API operation to create an event rule.
//
// @param tmpReq - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithContext(ctx context.Context, tmpReq *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventTargets) {
		request.EventTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventTargets, dara.String("EventTargets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventTargetsShrink) {
		query["EventTargets"] = request.EventTargetsShrink
	}

	if !dara.IsNil(request.FilterPattern) {
		query["FilterPattern"] = request.FilterPattern
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role for your cloud service.
//
// Description:
//
// You can call this API operation to create a service-linked role for your cloud service.
//
// @param request - CreateServiceLinkedRoleForProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleForProductResponse
func (client *Client) CreateServiceLinkedRoleForProductWithContext(ctx context.Context, request *CreateServiceLinkedRoleForProductRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleForProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRoleForProduct"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleForProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API destination.
//
// Description:
//
// You can call this API operation to delete an API destination.
//
// @param request - DeleteApiDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiDestinationResponse
func (client *Client) DeleteApiDestinationWithContext(ctx context.Context, request *DeleteApiDestinationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDestinationName) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiDestination"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection.
//
// Description:
//
// You can call this API operation to delete a connection.
//
// @param request - DeleteConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnectionWithContext(ctx context.Context, request *DeleteConnectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnection"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event bus.
//
// Description:
//
// You can call this API operation to delete an event bus.
//
// @param request - DeleteEventBusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventBusResponse
func (client *Client) DeleteEventBusWithContext(ctx context.Context, request *DeleteEventBusRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventBusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventBus"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventBusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event source.
//
// Description:
//
// You can call this API operation to delete an event source.
//
// @param request - DeleteEventSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSourceResponse
func (client *Client) DeleteEventSourceWithContext(ctx context.Context, request *DeleteEventSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		body["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventSourceName) {
		body["EventSourceName"] = request.EventSourceName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventSource"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event stream.
//
// Description:
//
// You can call this API operation to delete an event stream.
//
// @param request - DeleteEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventStreamingResponse
func (client *Client) DeleteEventStreamingWithContext(ctx context.Context, request *DeleteEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event rule.
//
// Description:
//
// You can call this API operation to delete an event rule.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more event targets of an event rule.
//
// Description:
//
// You can call this API operation to delete one or more event targets of an event rule.
//
// @param tmpReq - DeleteTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTargetsResponse
func (client *Client) DeleteTargetsWithContext(ctx context.Context, tmpReq *DeleteTargetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteTargetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TargetIds) {
		request.TargetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetIds, dara.String("TargetIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.TargetIdsShrink) {
		query["TargetIds"] = request.TargetIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTargets"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an event rule.
//
// Description:
//
// You can call this API operation to disable an event rule.
//
// @param request - DisableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRuleResponse
func (client *Client) DisableRuleWithContext(ctx context.Context, request *DisableRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发现EventSource(例如：Mysql)的Schema和SimpleData
//
// @param tmpReq - DiscoverEventSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiscoverEventSourceResponse
func (client *Client) DiscoverEventSourceWithContext(ctx context.Context, tmpReq *DiscoverEventSourceRequest, runtime *dara.RuntimeOptions) (_result *DiscoverEventSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DiscoverEventSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceMySQLParameters) {
		request.SourceMySQLParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMySQLParameters, dara.String("SourceMySQLParameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceMySQLParametersShrink) {
		body["SourceMySQLParameters"] = request.SourceMySQLParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DiscoverEventSource"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DiscoverEventSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an event rule.
//
// Description:
//
// You can call this API operation to enable an event rule.
//
// @param request - EnableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRuleResponse
func (client *Client) EnableRuleWithContext(ctx context.Context, request *EnableRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # EventCenterQueryEvents
//
// @param tmpReq - EventCenterQueryEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EventCenterQueryEventsResponse
func (client *Client) EventCenterQueryEventsWithContext(ctx context.Context, tmpReq *EventCenterQueryEventsRequest, runtime *dara.RuntimeOptions) (_result *EventCenterQueryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EventCenterQueryEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusName) {
		query["BusName"] = request.BusName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["Body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EventCenterQueryEvents"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EventCenterQueryEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an API destination.
//
// Description:
//
// You can call this API operation to query the information about an API destination.
//
// @param request - GetApiDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiDestinationResponse
func (client *Client) GetApiDestinationWithContext(ctx context.Context, request *GetApiDestinationRequest, runtime *dara.RuntimeOptions) (_result *GetApiDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDestinationName) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiDestination"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a connection.
//
// Description:
//
// You can call this API operation to query the configurations of a connection.
//
// @param request - GetConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionResponse
func (client *Client) GetConnectionWithContext(ctx context.Context, request *GetConnectionRequest, runtime *dara.RuntimeOptions) (_result *GetConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnection"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an event bus.
//
// Description:
//
// You can call this API operation to query the detailed information about an event bus.
//
// @param request - GetEventBusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventBusResponse
func (client *Client) GetEventBusWithContext(ctx context.Context, request *GetEventBusRequest, runtime *dara.RuntimeOptions) (_result *GetEventBusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEventBus"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventBusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an event stream.
//
// Description:
//
// You can call this API operation to query the details of an event stream.
//
// @param request - GetEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventStreamingResponse
func (client *Client) GetEventStreamingWithContext(ctx context.Context, request *GetEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *GetEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an event rule.
//
// Description:
//
// You can call this API operation to query the details of an event rule.
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithContext(ctx context.Context, request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of API destinations.
//
// Description:
//
// You can use this API operation to query a list of API destinations.
//
// @param request - ListApiDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiDestinationsResponse
func (client *Client) ListApiDestinationsWithContext(ctx context.Context, request *ListApiDestinationsRequest, runtime *dara.RuntimeOptions) (_result *ListApiDestinationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDestinationNamePrefix) {
		query["ApiDestinationNamePrefix"] = request.ApiDestinationNamePrefix
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiDestinations"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries connections.
//
// Description:
//
// You can call this API operation to query connections.
//
// @param request - ListConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectionsResponse
func (client *Client) ListConnectionsWithContext(ctx context.Context, request *ListConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionNamePrefix) {
		body["ConnectionNamePrefix"] = request.ConnectionNamePrefix
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnections"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all event buses.
//
// Description:
//
// You can call this API operation to query all event buses.
//
// @param request - ListEventBusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventBusesResponse
func (client *Client) ListEventBusesWithContext(ctx context.Context, request *ListEventBusesRequest, runtime *dara.RuntimeOptions) (_result *ListEventBusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamePrefix) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEventBuses"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventBusesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event streams.
//
// Description:
//
// You can call this API operation to query event streams.
//
// @param request - ListEventStreamingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventStreamingsResponse
func (client *Client) ListEventStreamingsWithContext(ctx context.Context, request *ListEventStreamingsRequest, runtime *dara.RuntimeOptions) (_result *ListEventStreamingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamePrefix) {
		body["NamePrefix"] = request.NamePrefix
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SinkArn) {
		body["SinkArn"] = request.SinkArn
	}

	if !dara.IsNil(request.SourceArn) {
		body["SourceArn"] = request.SourceArn
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEventStreamings"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventStreamingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all rules of an event bus.
//
// Description:
//
// You can call this API operation to query all rules of an event bus.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithContext(ctx context.Context, request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RuleNamePrefix) {
		query["RuleNamePrefix"] = request.RuleNamePrefix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all event targets of an event rule.
//
// @param request - ListTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTargetsResponse
func (client *Client) ListTargetsWithContext(ctx context.Context, request *ListTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Arn) {
		query["Arn"] = request.Arn
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTargets"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom event sources.
//
// Description:
//
// You can call this API operation to query custom event sources.
//
// @param request - ListUserDefinedEventSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDefinedEventSourcesResponse
func (client *Client) ListUserDefinedEventSourcesWithContext(ctx context.Context, request *ListUserDefinedEventSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListUserDefinedEventSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NamePrefix) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDefinedEventSources"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDefinedEventSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an event stream that is running.
//
// Description:
//
// You can call this API operation to stop an event stream that is running.
//
// @param request - PauseEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseEventStreamingResponse
func (client *Client) PauseEventStreamingWithContext(ctx context.Context, request *PauseEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *PauseEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates event targets under a rule.
//
// Description:
//
// You can call this API operation to create or update event targets under a rule.
//
// @param tmpReq - PutTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutTargetsResponse
func (client *Client) PutTargetsWithContext(ctx context.Context, tmpReq *PutTargetsRequest, runtime *dara.RuntimeOptions) (_result *PutTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PutTargetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Targets) {
		request.TargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Targets, dara.String("Targets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.TargetsShrink) {
		query["Targets"] = request.TargetsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutTargets"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of an event.
//
// Description:
//
// You can call this API operation to query the content of an event.
//
// @param request - QueryEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventResponse
func (client *Client) QueryEventWithContext(ctx context.Context, request *QueryEventRequest, runtime *dara.RuntimeOptions) (_result *QueryEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventSource) {
		query["EventSource"] = request.EventSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEvent"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event traces.
//
// Description:
//
// You can call this API operation to query event traces.
//
// @param request - QueryEventTracesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEventTracesResponse
func (client *Client) QueryEventTracesWithContext(ctx context.Context, request *QueryEventTracesRequest, runtime *dara.RuntimeOptions) (_result *QueryEventTracesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEventTraces"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEventTracesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event traces by event ID.
//
// Description:
//
// You can call this API operation to query event traces by event ID.
//
// @param request - QueryTracedEventByEventIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTracedEventByEventIdResponse
func (client *Client) QueryTracedEventByEventIdWithContext(ctx context.Context, request *QueryTracedEventByEventIdRequest, runtime *dara.RuntimeOptions) (_result *QueryTracedEventByEventIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventSource) {
		query["EventSource"] = request.EventSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTracedEventByEventId"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTracedEventByEventIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event traces by time range.
//
// Description:
//
// You can call this API operation to query event traces by time range.
//
// @param request - QueryTracedEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTracedEventsResponse
func (client *Client) QueryTracedEventsWithContext(ctx context.Context, request *QueryTracedEventsRequest, runtime *dara.RuntimeOptions) (_result *QueryTracedEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventSource) {
		query["EventSource"] = request.EventSource
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MatchedRule) {
		query["MatchedRule"] = request.MatchedRule
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTracedEvents"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTracedEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a created or deactivated event stream.
//
// Description:
//
// You can call this API operation to enable a created or deactivated event stream.
//
// @param request - StartEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEventStreamingResponse
func (client *Client) StartEventStreamingWithContext(ctx context.Context, request *StartEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *StartEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the event pattern matches the provided JSON format.
//
// Description:
//
// You can call this API operation to check whether the event pattern matches the provided JSON format.
//
// @param request - TestEventPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestEventPatternResponse
func (client *Client) TestEventPatternWithContext(ctx context.Context, request *TestEventPatternRequest, runtime *dara.RuntimeOptions) (_result *TestEventPatternResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Event) {
		body["Event"] = request.Event
	}

	if !dara.IsNil(request.EventPattern) {
		body["EventPattern"] = request.EventPattern
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestEventPattern"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestEventPatternResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether event source configurations are available.
//
// Description:
//
// You can call this API operation to query all custom event sources.
//
// @param tmpReq - TestEventSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestEventSourceConfigResponse
func (client *Client) TestEventSourceConfigWithContext(ctx context.Context, tmpReq *TestEventSourceConfigRequest, runtime *dara.RuntimeOptions) (_result *TestEventSourceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TestEventSourceConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SourceMySQLParameters) {
		request.SourceMySQLParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMySQLParameters, dara.String("SourceMySQLParameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceMySQLParametersShrink) {
		body["SourceMySQLParameters"] = request.SourceMySQLParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestEventSourceConfig"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestEventSourceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an API destination.
//
// Description:
//
// You can call this API operation to update an API destination.
//
// @param tmpReq - UpdateApiDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiDestinationResponse
func (client *Client) UpdateApiDestinationWithContext(ctx context.Context, tmpReq *UpdateApiDestinationRequest, runtime *dara.RuntimeOptions) (_result *UpdateApiDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApiDestinationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HttpApiParameters) {
		request.HttpApiParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HttpApiParameters, dara.String("HttpApiParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDestinationName) {
		query["ApiDestinationName"] = request.ApiDestinationName
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HttpApiParametersShrink) {
		query["HttpApiParameters"] = request.HttpApiParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiDestination"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a connection.
//
// Description:
//
// You can call this API operation to update a connection.
//
// @param tmpReq - UpdateConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectionResponse
func (client *Client) UpdateConnectionWithContext(ctx context.Context, tmpReq *UpdateConnectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateConnectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthParameters) {
		request.AuthParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthParameters, dara.String("AuthParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NetworkParameters) {
		request.NetworkParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkParameters, dara.String("NetworkParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthParametersShrink) {
		query["AuthParameters"] = request.AuthParametersShrink
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkParametersShrink) {
		query["NetworkParameters"] = request.NetworkParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnection"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an event bus.
//
// Description:
//
// You can call this operation to update an event bus.
//
// @param request - UpdateEventBusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventBusResponse
func (client *Client) UpdateEventBusWithContext(ctx context.Context, request *UpdateEventBusRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventBusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventBus"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventBusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an event source.
//
// Description:
//
// You can call this API operation to update an event source.
//
// @param tmpReq - UpdateEventSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventSourceResponse
func (client *Client) UpdateEventSourceWithContext(ctx context.Context, tmpReq *UpdateEventSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEventSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExternalSourceConfig) {
		request.ExternalSourceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalSourceConfig, dara.String("ExternalSourceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceHttpEventParameters) {
		request.SourceHttpEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceHttpEventParameters, dara.String("SourceHttpEventParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceKafkaParameters) {
		request.SourceKafkaParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceKafkaParameters, dara.String("SourceKafkaParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceMNSParameters) {
		request.SourceMNSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceMNSParameters, dara.String("SourceMNSParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceOSSEventParameters) {
		request.SourceOSSEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceOSSEventParameters, dara.String("SourceOSSEventParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceRabbitMQParameters) {
		request.SourceRabbitMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRabbitMQParameters, dara.String("SourceRabbitMQParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceRocketMQParameters) {
		request.SourceRocketMQParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceRocketMQParameters, dara.String("SourceRocketMQParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceSLSParameters) {
		request.SourceSLSParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceSLSParameters, dara.String("SourceSLSParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceScheduledEventParameters) {
		request.SourceScheduledEventParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceScheduledEventParameters, dara.String("SourceScheduledEventParameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		body["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.EventSourceName) {
		body["EventSourceName"] = request.EventSourceName
	}

	if !dara.IsNil(request.ExternalSourceConfigShrink) {
		body["ExternalSourceConfig"] = request.ExternalSourceConfigShrink
	}

	if !dara.IsNil(request.ExternalSourceType) {
		body["ExternalSourceType"] = request.ExternalSourceType
	}

	if !dara.IsNil(request.LinkedExternalSource) {
		body["LinkedExternalSource"] = request.LinkedExternalSource
	}

	if !dara.IsNil(request.SourceHttpEventParametersShrink) {
		body["SourceHttpEventParameters"] = request.SourceHttpEventParametersShrink
	}

	if !dara.IsNil(request.SourceKafkaParametersShrink) {
		body["SourceKafkaParameters"] = request.SourceKafkaParametersShrink
	}

	if !dara.IsNil(request.SourceMNSParametersShrink) {
		body["SourceMNSParameters"] = request.SourceMNSParametersShrink
	}

	if !dara.IsNil(request.SourceOSSEventParametersShrink) {
		body["SourceOSSEventParameters"] = request.SourceOSSEventParametersShrink
	}

	if !dara.IsNil(request.SourceRabbitMQParametersShrink) {
		body["SourceRabbitMQParameters"] = request.SourceRabbitMQParametersShrink
	}

	if !dara.IsNil(request.SourceRocketMQParametersShrink) {
		body["SourceRocketMQParameters"] = request.SourceRocketMQParametersShrink
	}

	if !dara.IsNil(request.SourceSLSParametersShrink) {
		body["SourceSLSParameters"] = request.SourceSLSParametersShrink
	}

	if !dara.IsNil(request.SourceScheduledEventParametersShrink) {
		body["SourceScheduledEventParameters"] = request.SourceScheduledEventParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventSource"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an event stream, such as the basic information and the information about the event source, event filtering rule, and event target.
//
// Description:
//
// You can call this API operation to modify the information about an event stream, such as the basic information and the information about the event source, event filtering rule, and event target.
//
// @param tmpReq - UpdateEventStreamingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventStreamingResponse
func (client *Client) UpdateEventStreamingWithContext(ctx context.Context, tmpReq *UpdateEventStreamingRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventStreamingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEventStreamingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RunOptions) {
		request.RunOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RunOptions, dara.String("RunOptions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sink) {
		request.SinkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sink, dara.String("Sink"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Source) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, dara.String("Source"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Transforms) {
		request.TransformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transforms, dara.String("Transforms"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	if !dara.IsNil(request.FilterPattern) {
		body["FilterPattern"] = request.FilterPattern
	}

	if !dara.IsNil(request.RunOptionsShrink) {
		body["RunOptions"] = request.RunOptionsShrink
	}

	if !dara.IsNil(request.SinkShrink) {
		body["Sink"] = request.SinkShrink
	}

	if !dara.IsNil(request.SourceShrink) {
		body["Source"] = request.SourceShrink
	}

	if !dara.IsNil(request.TransformsShrink) {
		body["Transforms"] = request.TransformsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventStreaming"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventStreamingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询事件流
//
// @param request - UpdateEventStreamingBusinessOptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventStreamingBusinessOptionResponse
func (client *Client) UpdateEventStreamingBusinessOptionWithContext(ctx context.Context, request *UpdateEventStreamingBusinessOptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventStreamingBusinessOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessMode) {
		body["BusinessMode"] = request.BusinessMode
	}

	if !dara.IsNil(request.EventStreamingName) {
		body["EventStreamingName"] = request.EventStreamingName
	}

	if !dara.IsNil(request.MaxCapacityUnitCount) {
		body["MaxCapacityUnitCount"] = request.MaxCapacityUnitCount
	}

	if !dara.IsNil(request.MinCapacityUnitCount) {
		body["MinCapacityUnitCount"] = request.MinCapacityUnitCount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventStreamingBusinessOption"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventStreamingBusinessOptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an event rule.
//
// Description:
//
// You can call this API operation to update the configurations of an event rule.
//
// @param request - UpdateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
func (client *Client) UpdateRuleWithContext(ctx context.Context, request *UpdateRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventBusName) {
		query["EventBusName"] = request.EventBusName
	}

	if !dara.IsNil(request.FilterPattern) {
		query["FilterPattern"] = request.FilterPattern
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRule"),
		Version:     dara.String("2020-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
