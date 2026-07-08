// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a flow.
//
// Description:
//
// Creates a flow or a flow version.
//
// @param request - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithContext(ctx context.Context, request *CreateFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.LaunchStatus) {
		query["LaunchStatus"] = request.LaunchStatus
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - CreateUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserAuthConfigResponse
func (client *Client) CreateUserAuthConfigWithContext(ctx context.Context, request *CreateUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfig) {
		query["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthConfigName) {
		query["AuthConfigName"] = request.AuthConfigName
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserAuthConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection flow.
//
// @param request - DeleteFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithContext(ctx context.Context, request *DeleteFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user authentication credential.
//
// Description:
//
// Creates a connection flow or a connection flow version.
//
// @param request - DeleteUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserAuthConfigResponse
func (client *Client) DeleteUserAuthConfigWithContext(ctx context.Context, request *DeleteUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserAuthConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a flow.
//
// @param request - DisableFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableFlowResponse
func (client *Client) DisableFlowWithContext(ctx context.Context, request *DisableFlowRequest, runtime *dara.RuntimeOptions) (_result *DisableFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a flow.
//
// @param request - EnableFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFlowResponse
func (client *Client) EnableFlowWithContext(ctx context.Context, request *EnableFlowRequest, runtime *dara.RuntimeOptions) (_result *EnableFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a logon session token.
//
// @param request - GenerateUserSessionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionTokenWithContext(ctx context.Context, request *GenerateUserSessionTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateUserSessionTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.ExpireSecond) {
		query["ExpireSecond"] = request.ExpireSecond
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.IntegrateId) {
		query["IntegrateId"] = request.IntegrateId
	}

	if !dara.IsNil(request.UserAvatar) {
		query["UserAvatar"] = request.UserAvatar
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUserSessionToken"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a flow.
//
// @param request - GetFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowResponse
func (client *Client) GetFlowWithContext(ctx context.Context, request *GetFlowRequest, runtime *dara.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a user authentication credential.
//
// Description:
//
// This operation gets the details of a specified credential.
//
// @param request - GetUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAuthConfigResponse
func (client *Client) GetUserAuthConfigWithContext(ctx context.Context, request *GetUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *GetUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserAuthConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes a connector action.
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithSSECtx(ctx context.Context, tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions, _yield chan *InvokeActionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeActionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// Invokes a connector action.
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithContext(ctx context.Context, tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) (_result *InvokeActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeActionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Launches a flow.
//
// @param request - LaunchFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchFlowResponse
func (client *Client) LaunchFlowWithContext(ctx context.Context, request *LaunchFlowRequest, runtime *dara.RuntimeOptions) (_result *LaunchFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LaunchFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LaunchFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of connector flows.
//
// Description:
//
// Creates a connector flow or a connector flow version.
//
// @param request - ListFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowsResponse
func (client *Client) ListFlowsWithContext(ctx context.Context, request *ListFlowsRequest, runtime *dara.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlows"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists user authentication credentials.
//
// Description:
//
// This operation retrieves user auth configs that match specified filters.
//
// @param request - ListUserAuthConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAuthConfigsResponse
func (client *Client) ListUserAuthConfigsWithContext(ctx context.Context, request *ListUserAuthConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListUserAuthConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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
		Action:      dara.String("ListUserAuthConfigs"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAuthConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a connection flow.
//
// @param request - UpdateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlowWithContext(ctx context.Context, request *UpdateFlowRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.FlowDesc) {
		query["FlowDesc"] = request.FlowDesc
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowTemplate) {
		query["FlowTemplate"] = request.FlowTemplate
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a user authentication credential.
//
// Description:
//
// Updates the configuration of a specific user authentication credential.
//
// @param request - UpdateUserAuthConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserAuthConfigResponse
func (client *Client) UpdateUserAuthConfigWithContext(ctx context.Context, request *UpdateUserAuthConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserAuthConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthConfig) {
		query["AuthConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.AuthConfigId) {
		query["AuthConfigId"] = request.AuthConfigId
	}

	if !dara.IsNil(request.AuthConfigName) {
		query["AuthConfigName"] = request.AuthConfigName
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserAuthConfig"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserAuthConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Withdraws a connection flow.
//
// @param request - WithdrawFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawFlowResponse
func (client *Client) WithdrawFlowWithContext(ctx context.Context, request *WithdrawFlowRequest, runtime *dara.RuntimeOptions) (_result *WithdrawFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WithdrawFlow"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) invokeActionWithSSECtx_opYieldFunc(_yield chan *InvokeActionResponse, _yieldErr chan error, ctx context.Context, tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
