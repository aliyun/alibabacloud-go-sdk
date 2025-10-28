// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 修改实例所在资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/resource-groups"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建函数别名。
//
// @param request - CreateAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAliasResponse
func (client *Client) CreateAliasWithContext(ctx context.Context, functionName *string, request *CreateAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlias"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/aliases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom domain name.
//
// Description:
//
// If you want to use a fixed domain name to access an application or function in a production environment of Function Compute, or to resolve the issue of forced downloads when accessing an HTTP trigger, you can bind a custom domain name to the application or function.
//
// @param request - CreateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomainWithContext(ctx context.Context, request *CreateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a function.
//
// Description:
//
// Resources of Function Compute are scheduled and run based on functions. A function usually refers to a code snippet that is written by a user and can be independently executed to respond to events and requests.
//
// @param request - CreateFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithContext(ctx context.Context, request *CreateFunctionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunction"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建层版本。
//
// @param request - CreateLayerVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayerVersionResponse
func (client *Client) CreateLayerVersionWithContext(ctx context.Context, layerName *string, request *CreateLayerVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLayerVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLayerVersion"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers/" + dara.PercentEncode(dara.StringValue(layerName)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLayerVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会话资源
//
// @param request - CreateSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithContext(ctx context.Context, functionName *string, request *CreateSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSession"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/sessions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建函数触发器。
//
// @param request - CreateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithContext(ctx context.Context, functionName *string, request *CreateTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrigger"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/triggers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VPC connection.
//
// @param request - CreateVpcBindingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcBindingResponse
func (client *Client) CreateVpcBindingWithContext(ctx context.Context, functionName *string, request *CreateVpcBindingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVpcBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcBinding"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/vpc-bindings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateVpcBindingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alias.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithContext(ctx context.Context, functionName *string, aliasName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlias"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/aliases/" + dara.PercentEncode(dara.StringValue(aliasName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an asynchronous invocation configuration.
//
// @param request - DeleteAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsyncInvokeConfigResponse
func (client *Client) DeleteAsyncInvokeConfigWithContext(ctx context.Context, functionName *string, request *DeleteAsyncInvokeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAsyncInvokeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAsyncInvokeConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-invoke-config"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteAsyncInvokeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a concurrency configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConcurrencyConfigResponse
func (client *Client) DeleteConcurrencyConfigWithContext(ctx context.Context, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConcurrencyConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConcurrencyConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/concurrency"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteConcurrencyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomainWithContext(ctx context.Context, domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a function.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithContext(ctx context.Context, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunction"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// http://pre.hhht/#vpc
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionVersionResponse
func (client *Client) DeleteFunctionVersionWithContext(ctx context.Context, functionName *string, versionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunctionVersion"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/versions/" + dara.PercentEncode(dara.StringValue(versionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteFunctionVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a layer version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerVersionResponse
func (client *Client) DeleteLayerVersionWithContext(ctx context.Context, layerName *string, version *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLayerVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLayerVersion"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers/" + dara.PercentEncode(dara.StringValue(layerName)) + "/versions/" + dara.PercentEncode(dara.StringValue(version))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteLayerVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a provisioned configuration.
//
// @param request - DeleteProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProvisionConfigResponse
func (client *Client) DeleteProvisionConfigWithContext(ctx context.Context, functionName *string, request *DeleteProvisionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProvisionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProvisionConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/provision-config"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteProvisionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除弹性配置
//
// @param request - DeleteScalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingConfigResponse
func (client *Client) DeleteScalingConfigWithContext(ctx context.Context, functionName *string, request *DeleteScalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScalingConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/scaling-config"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteScalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会话资源
//
// @param request - DeleteSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionResponse
func (client *Client) DeleteSessionWithContext(ctx context.Context, functionName *string, sessionId *string, request *DeleteSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSession"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trigger.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTriggerWithContext(ctx context.Context, functionName *string, triggerName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrigger"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/triggers/" + dara.PercentEncode(dara.StringValue(triggerName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specified policy group for a VPC firewall.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcBindingResponse
func (client *Client) DeleteVpcBindingWithContext(ctx context.Context, functionName *string, vpcId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVpcBindingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcBinding"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/vpc-bindings/" + dara.PercentEncode(dara.StringValue(vpcId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteVpcBindingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品的地域信息列表
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DisableFunctionInvocation operation prevents a function from being invoked and optionally terminates all requests that are being processed. Once a function\\"s invocation is disabled, no new instances can be created, and the existing provisioned instances are destroyed. This operation is currently in private preview.
//
// Description:
//
// Exercise caution when you call this operation on a function in a production environment, as improper deactivation may lead to business disruptions.
//
// @param request - DisableFunctionInvocationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableFunctionInvocationResponse
func (client *Client) DisableFunctionInvocationWithContext(ctx context.Context, functionName *string, request *DisableFunctionInvocationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableFunctionInvocationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AbortOngoingRequest) {
		body["abortOngoingRequest"] = request.AbortOngoingRequest
	}

	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableFunctionInvocation"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/invoke/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableFunctionInvocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 允许函数调用
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFunctionInvocationResponse
func (client *Client) EnableFunctionInvocationWithContext(ctx context.Context, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableFunctionInvocationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableFunctionInvocation"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/invoke/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableFunctionInvocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an alias.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliasResponse
func (client *Client) GetAliasWithContext(ctx context.Context, functionName *string, aliasName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAliasResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlias"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/aliases/" + dara.PercentEncode(dara.StringValue(aliasName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets asynchronous invocation configurations of a function.
//
// @param request - GetAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncInvokeConfigResponse
func (client *Client) GetAsyncInvokeConfigWithContext(ctx context.Context, functionName *string, request *GetAsyncInvokeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAsyncInvokeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncInvokeConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-invoke-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncInvokeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithContext(ctx context.Context, functionName *string, taskId *string, request *GetAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTask"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a concurrency configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConcurrencyConfigResponse
func (client *Client) GetConcurrencyConfigWithContext(ctx context.Context, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConcurrencyConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConcurrencyConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/concurrency"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConcurrencyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a custom domain name.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomainWithContext(ctx context.Context, domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// http://pre.hhht/#vpc
//
// @param request - GetFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithContext(ctx context.Context, functionName *string, request *GetFunctionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunction"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a code package of a function.
//
// @param request - GetFunctionCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionCodeResponse
func (client *Client) GetFunctionCodeWithContext(ctx context.Context, functionName *string, request *GetFunctionCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionCode"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/code"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of a layer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerVersionResponse
func (client *Client) GetLayerVersionWithContext(ctx context.Context, layerName *string, version *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLayerVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLayerVersion"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers/" + dara.PercentEncode(dara.StringValue(layerName)) + "/versions/" + dara.PercentEncode(dara.StringValue(version))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLayerVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain version information of a layer by using ARNs.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerVersionByArnResponse
func (client *Client) GetLayerVersionByArnWithContext(ctx context.Context, arn *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLayerVersionByArnResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLayerVersionByArn"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layerarn/" + dara.PercentEncode(dara.StringValue(arn))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLayerVersionByArnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries provisioned configurations.
//
// @param request - GetProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProvisionConfigResponse
func (client *Client) GetProvisionConfigWithContext(ctx context.Context, functionName *string, request *GetProvisionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProvisionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProvisionConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/provision-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProvisionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取弹性配置
//
// @param request - GetScalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScalingConfigResponse
func (client *Client) GetScalingConfigWithContext(ctx context.Context, functionName *string, request *GetScalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScalingConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/scaling-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取函数会话信息。
//
// @param request - GetSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionResponse
func (client *Client) GetSessionWithContext(ctx context.Context, functionName *string, sessionId *string, request *GetSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSession"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a trigger.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTriggerResponse
func (client *Client) GetTriggerWithContext(ctx context.Context, functionName *string, triggerName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrigger"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/triggers/" + dara.PercentEncode(dara.StringValue(triggerName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes a function.
//
// @param request - InvokeFunctionRequest
//
// @param headers - InvokeFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeFunctionResponse
func (client *Client) InvokeFunctionWithContext(ctx context.Context, functionName *string, request *InvokeFunctionRequest, headers *InvokeFunctionHeaders, runtime *dara.RuntimeOptions) (_result *InvokeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XFcAsyncTaskId) {
		realHeaders["x-fc-async-task-id"] = dara.String(dara.ToString(dara.StringValue(headers.XFcAsyncTaskId)))
	}

	if !dara.IsNil(headers.XFcInvocationType) {
		realHeaders["x-fc-invocation-type"] = dara.String(dara.ToString(dara.StringValue(headers.XFcInvocationType)))
	}

	if !dara.IsNil(headers.XFcLogType) {
		realHeaders["x-fc-log-type"] = dara.String(dara.ToString(dara.StringValue(headers.XFcLogType)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
		Stream:  request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeFunction"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/invocations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("binary"),
	}
	res := &InvokeFunctionResponse{}
	callApiTmp, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp := dara.ToMap(callApiTmp)
	if !dara.IsNil(tmp["body"]) {
		respBody := dara.ToReader(tmp["body"])
		res.Body = respBody
	}

	if !dara.IsNil(tmp["headers"]) {
		respHeaders := dara.ToMap(tmp["headers"])
		res.Headers = openapiutil.StringifyMapValue(respHeaders)
	}

	if !dara.IsNil(tmp["statusCode"]) {
		statusCode := dara.ForceInt(tmp["statusCode"])
		res.StatusCode = dara.ToInt32(dara.Int(statusCode))
	}

	_result = res
	return _result, _err
}

// Summary:
//
// Queries aliases.
//
// @param request - ListAliasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesResponse
func (client *Client) ListAliasesWithContext(ctx context.Context, functionName *string, request *ListAliasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAliases"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/aliases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all asynchronous configurations of a function.
//
// @param request - ListAsyncInvokeConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncInvokeConfigsResponse
func (client *Client) ListAsyncInvokeConfigsWithContext(ctx context.Context, request *ListAsyncInvokeConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAsyncInvokeConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsyncInvokeConfigs"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/async-invoke-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsyncInvokeConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists asynchronous tasks.
//
// @param request - ListAsyncTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasksWithContext(ctx context.Context, functionName *string, request *ListAsyncTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludePayload) {
		query["includePayload"] = request.IncludePayload
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	if !dara.IsNil(request.SortOrderByTime) {
		query["sortOrderByTime"] = request.SortOrderByTime
	}

	if !dara.IsNil(request.StartedTimeBegin) {
		query["startedTimeBegin"] = request.StartedTimeBegin
	}

	if !dara.IsNil(request.StartedTimeEnd) {
		query["startedTimeEnd"] = request.StartedTimeEnd
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsyncTasks"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出函数并发度配置。
//
// @param request - ListConcurrencyConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConcurrencyConfigsResponse
func (client *Client) ListConcurrencyConfigsWithContext(ctx context.Context, request *ListConcurrencyConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConcurrencyConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConcurrencyConfigs"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/concurrency-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConcurrencyConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom domain names.
//
// @param request - ListCustomDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomainsWithContext(ctx context.Context, request *ListCustomDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCustomDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomDomains"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of a function.
//
// @param request - ListFunctionVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionVersionsResponse
func (client *Client) ListFunctionVersionsWithContext(ctx context.Context, functionName *string, request *ListFunctionVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionVersions"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出函数。
//
// Description:
//
// ListFunctions returns only a subset of a function\\"s attribute fields. To obtain the additional fields, which include state, stateReasonCode, stateReason, lastUpdateStatus, lastUpdateStatusReasonCode, and lastUpdateStatusReason, use [GetFunction](https://help.aliyun.com/document_detail/2618610.html).
//
// @param tmpReq - ListFunctionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithContext(ctx context.Context, tmpReq *ListFunctionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFunctionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.FcVersion) {
		query["fcVersion"] = request.FcVersion
	}

	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.GpuType) {
		query["gpuType"] = request.GpuType
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Runtime) {
		query["runtime"] = request.Runtime
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctions"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of function instances.
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, functionName *string, tmpReq *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("instanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceStatus) {
		request.InstanceStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceStatus, dara.String("instanceStatus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimeMs) {
		query["endTimeMs"] = request.EndTimeMs
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["instanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.InstanceStatusShrink) {
		query["instanceStatus"] = request.InstanceStatusShrink
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	if !dara.IsNil(request.StartKey) {
		query["startKey"] = request.StartKey
	}

	if !dara.IsNil(request.StartTimeMs) {
		query["startTimeMs"] = request.StartTimeMs
	}

	if !dara.IsNil(request.WithAllActive) {
		query["withAllActive"] = request.WithAllActive
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a list of layer versions.
//
// @param request - ListLayerVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayerVersionsResponse
func (client *Client) ListLayerVersionsWithContext(ctx context.Context, layerName *string, request *ListLayerVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLayerVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.StartVersion) {
		query["startVersion"] = request.StartVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLayerVersions"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers/" + dara.PercentEncode(dara.StringValue(layerName)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLayerVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a list of layers.
//
// @param request - ListLayersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayersResponse
func (client *Client) ListLayersWithContext(ctx context.Context, request *ListLayersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLayersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Official) {
		query["official"] = request.Official
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.Public) {
		query["public"] = request.Public
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLayers"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLayersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of provisioned configurations.
//
// @param request - ListProvisionConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProvisionConfigsResponse
func (client *Client) ListProvisionConfigsWithContext(ctx context.Context, request *ListProvisionConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProvisionConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProvisionConfigs"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/provision-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProvisionConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取弹性配置列表
//
// @param request - ListScalingConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScalingConfigsResponse
func (client *Client) ListScalingConfigsWithContext(ctx context.Context, request *ListScalingConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScalingConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScalingConfigs"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/scaling-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScalingConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出函数会话信息
//
// @param request - ListSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionsResponse
func (client *Client) ListSessionsWithContext(ctx context.Context, functionName *string, request *ListSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionStatus) {
		query["sessionStatus"] = request.SessionStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessions"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all tagged resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/tags-v2"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the triggers of a function.
//
// @param request - ListTriggersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTriggersResponse
func (client *Client) ListTriggersWithContext(ctx context.Context, functionName *string, request *ListTriggersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTriggers"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/triggers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of existing VPC connections.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcBindingsResponse
func (client *Client) ListVpcBindingsWithContext(ctx context.Context, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVpcBindingsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcBindings"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/vpc-bindings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcBindingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a function version.
//
// @param request - PublishFunctionVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFunctionVersionResponse
func (client *Client) PublishFunctionVersionWithContext(ctx context.Context, functionName *string, request *PublishFunctionVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishFunctionVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishFunctionVersion"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishFunctionVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an asynchronous invocation configuration for a function.
//
// @param request - PutAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutAsyncInvokeConfigResponse
func (client *Client) PutAsyncInvokeConfigWithContext(ctx context.Context, functionName *string, request *PutAsyncInvokeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutAsyncInvokeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutAsyncInvokeConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-invoke-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutAsyncInvokeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures concurrency of a function.
//
// @param request - PutConcurrencyConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutConcurrencyConfigResponse
func (client *Client) PutConcurrencyConfigWithContext(ctx context.Context, functionName *string, request *PutConcurrencyConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutConcurrencyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutConcurrencyConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/concurrency"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutConcurrencyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies permissions of a layer.
//
// @param request - PutLayerACLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutLayerACLResponse
func (client *Client) PutLayerACLWithContext(ctx context.Context, layerName *string, request *PutLayerACLRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutLayerACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Acl) {
		query["acl"] = request.Acl
	}

	if !dara.IsNil(request.Public) {
		query["public"] = request.Public
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutLayerACL"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/layers/" + dara.PercentEncode(dara.StringValue(layerName)) + "/acl"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PutLayerACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates provisioned configurations.
//
// @param request - PutProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutProvisionConfigResponse
func (client *Client) PutProvisionConfigWithContext(ctx context.Context, functionName *string, request *PutProvisionConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutProvisionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutProvisionConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/provision-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutProvisionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置弹性配置
//
// @param request - PutScalingConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutScalingConfigResponse
func (client *Client) PutScalingConfigWithContext(ctx context.Context, functionName *string, request *PutScalingConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PutScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutScalingConfig"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/scaling-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PutScalingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an asynchronous task.
//
// @param request - StopAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAsyncTaskResponse
func (client *Client) StopAsyncTaskWithContext(ctx context.Context, functionName *string, taskId *string, request *StopAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAsyncTask"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/async-tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &StopAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to a resource.
//
// Description:
//
// Tags are used to identify resources. Tags allow you to categorize, search for, and aggregate resources that have the same characteristics from different dimensions. This facilitates resource management. For more information, see [Tag overview](https://help.aliyun.com/document_detail/156983.html).
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/tags-v2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKey) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, dara.String("TagKey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeyShrink) {
		query["TagKey"] = request.TagKeyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/tags-v2"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alias.
//
// @param request - UpdateAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithContext(ctx context.Context, functionName *string, aliasName *string, request *UpdateAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlias"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/aliases/" + dara.PercentEncode(dara.StringValue(aliasName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a custom domain name.
//
// @param request - UpdateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomainWithContext(ctx context.Context, domainName *string, request *UpdateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCustomDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a function.
//
// @param request - UpdateFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithContext(ctx context.Context, functionName *string, request *UpdateFunctionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunction"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会话配置
//
// @param request - UpdateSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionResponse
func (client *Client) UpdateSessionWithContext(ctx context.Context, functionName *string, sessionId *string, request *UpdateSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Qualifier) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSession"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a trigger.
//
// @param request - UpdateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTriggerWithContext(ctx context.Context, functionName *string, triggerName *string, request *UpdateTriggerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrigger"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/triggers/" + dara.PercentEncode(dara.StringValue(triggerName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
