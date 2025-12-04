// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 执行定时计划
//
// @param headers - ApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyScheduledPlanResponse
func (client *Client) ApplyScheduledPlanWithContext(ctx context.Context, namespace *string, scheduledPlanId *string, headers *ApplyScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *ApplyScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId)) + "%3Aapply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消调试
//
// @param request - CancelSqlPreviewRequest
//
// @param headers - CancelSqlPreviewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSqlPreviewResponse
func (client *Client) CancelSqlPreviewWithContext(ctx context.Context, namespace *string, request *CancelSqlPreviewRequest, headers *CancelSqlPreviewHeaders, runtime *dara.RuntimeOptions) (_result *CancelSqlPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["queryId"] = request.QueryId
	}

	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSqlPreview"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/cancel"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSqlPreviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a deployment.
//
// @param request - CreateDeploymentRequest
//
// @param headers - CreateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeploymentWithContext(ctx context.Context, namespace *string, request *CreateDeploymentRequest, headers *CreateDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create a deploymentDraft
//
// @param request - CreateDeploymentDraftRequest
//
// @param headers - CreateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentDraftResponse
func (client *Client) CreateDeploymentDraftWithContext(ctx context.Context, namespace *string, request *CreateDeploymentDraftRequest, headers *CreateDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentDraftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建deploymentTarget
//
// @param request - CreateDeploymentTargetRequest
//
// @param headers - CreateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentTargetResponse
func (client *Client) CreateDeploymentTargetWithContext(ctx context.Context, namespace *string, request *CreateDeploymentTargetRequest, headers *CreateDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeploymentTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentTargetName) {
		query["deploymentTargetName"] = request.DeploymentTargetName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建部署目标V2
//
// @param request - CreateDeploymentTargetV2Request
//
// @param headers - CreateDeploymentTargetV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentTargetV2Response
func (client *Client) CreateDeploymentTargetV2WithContext(ctx context.Context, namespace *string, request *CreateDeploymentTargetV2Request, headers *CreateDeploymentTargetV2Headers, runtime *dara.RuntimeOptions) (_result *CreateDeploymentTargetV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentTargetName) {
		query["deploymentTargetName"] = request.DeploymentTargetName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeploymentTargetV2"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/support-elastic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeploymentTargetV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create a folder
//
// @param request - CreateFolderRequest
//
// @param headers - CreateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithContext(ctx context.Context, namespace *string, request *CreateFolderRequest, headers *CreateFolderHeaders, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a namespace as a member and grants permissions to the user.
//
// @param request - CreateMemberRequest
//
// @param headers - CreateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberResponse
func (client *Client) CreateMemberWithContext(ctx context.Context, namespace *string, request *CreateMemberRequest, headers *CreateMemberHeaders, runtime *dara.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a savepoint.
//
// @param request - CreateSavepointRequest
//
// @param headers - CreateSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavepointResponse
func (client *Client) CreateSavepointWithContext(ctx context.Context, namespace *string, request *CreateSavepointRequest, headers *CreateSavepointHeaders, runtime *dara.RuntimeOptions) (_result *CreateSavepointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		body["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.NativeFormat) {
		body["nativeFormat"] = request.NativeFormat
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSavepointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建定时执行计划
//
// @param request - CreateScheduledPlanRequest
//
// @param headers - CreateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledPlanResponse
func (client *Client) CreateScheduledPlanWithContext(ctx context.Context, namespace *string, request *CreateScheduledPlanRequest, headers *CreateScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *CreateScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建session集群
//
// @param request - CreateSessionClusterRequest
//
// @param headers - CreateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionClusterWithContext(ctx context.Context, namespace *string, request *CreateSessionClusterRequest, headers *CreateSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *CreateSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses all user-defined function (UDF) methods in your JAR or Python file and creates an artifact configuration for a UDF.
//
// @param request - CreateUdfArtifactRequest
//
// @param headers - CreateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfArtifactResponse
func (client *Client) CreateUdfArtifactWithContext(ctx context.Context, namespace *string, request *CreateUdfArtifactRequest, headers *CreateUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *CreateUdfArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a variable.
//
// @param request - CreateVariableRequest
//
// @param headers - CreateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVariableResponse
func (client *Client) CreateVariableWithContext(ctx context.Context, namespace *string, request *CreateVariableRequest, headers *CreateVariableHeaders, runtime *dara.RuntimeOptions) (_result *CreateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a registered custom connector from a workspace.
//
// @param headers - DeleteCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomConnectorResponse
func (client *Client) DeleteCustomConnectorWithContext(ctx context.Context, namespace *string, connectorName *string, headers *DeleteCustomConnectorHeaders, runtime *dara.RuntimeOptions) (_result *DeleteCustomConnectorResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomConnector"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors/" + dara.PercentEncode(dara.StringValue(connectorName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomConnectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a deployment based on the deployment ID.
//
// @param headers - DeleteDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentResponse
func (client *Client) DeleteDeploymentWithContext(ctx context.Context, namespace *string, deploymentId *string, headers *DeleteDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete a deploymentDraft
//
// @param headers - DeleteDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentDraftResponse
func (client *Client) DeleteDeploymentDraftWithContext(ctx context.Context, namespace *string, deploymentDraftId *string, headers *DeleteDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentDraftResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentDraftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除deploymentTarget
//
// @param headers - DeleteDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeploymentTargetResponse
func (client *Client) DeleteDeploymentTargetWithContext(ctx context.Context, namespace *string, deploymentTargetName *string, headers *DeleteDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDeploymentTargetResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeploymentTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete a folder
//
// @param headers - DeleteFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithContext(ctx context.Context, namespace *string, folderId *string, headers *DeleteFolderHeaders, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder/" + dara.PercentEncode(dara.StringValue(folderId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information about a job that is not in the running state in a deployment.
//
// @param headers - DeleteJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithContext(ctx context.Context, namespace *string, jobId *string, headers *DeleteJobHeaders, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions from a member.
//
// @param headers - DeleteMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemberResponse
func (client *Client) DeleteMemberWithContext(ctx context.Context, namespace *string, member *string, headers *DeleteMemberHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMemberResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members/" + dara.PercentEncode(dara.StringValue(member))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a savepoint.
//
// @param headers - DeleteSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavepointResponse
func (client *Client) DeleteSavepointWithContext(ctx context.Context, namespace *string, savepointId *string, headers *DeleteSavepointHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints/" + dara.PercentEncode(dara.StringValue(savepointId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSavepointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除定时执行计划
//
// @param headers - DeleteScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledPlanResponse
func (client *Client) DeleteScheduledPlanWithContext(ctx context.Context, namespace *string, scheduledPlanId *string, headers *DeleteScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *DeleteScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除session集群
//
// @param headers - DeleteSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionClusterResponse
func (client *Client) DeleteSessionClusterWithContext(ctx context.Context, namespace *string, sessionClusterName *string, headers *DeleteSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除UdfArtifact
//
// @param headers - DeleteUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfArtifactResponse
func (client *Client) DeleteUdfArtifactWithContext(ctx context.Context, namespace *string, udfArtifactName *string, headers *DeleteUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUdfArtifactResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/" + dara.PercentEncode(dara.StringValue(udfArtifactName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an existing user-defined function (UDF) from a Realtime Compute for Apache Flink workspace.
//
// @param request - DeleteUdfFunctionRequest
//
// @param headers - DeleteUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfFunctionResponse
func (client *Client) DeleteUdfFunctionWithContext(ctx context.Context, namespace *string, functionName *string, request *DeleteUdfFunctionRequest, headers *DeleteUdfFunctionHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUdfFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		query["className"] = request.ClassName
	}

	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdfFunction"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/function/" + dara.PercentEncode(dara.StringValue(functionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a variable.
//
// @param headers - DeleteVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVariableResponse
func (client *Client) DeleteVariableWithContext(ctx context.Context, namespace *string, name *string, headers *DeleteVariableHeaders, runtime *dara.RuntimeOptions) (_result *DeleteVariableResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// deploy deploymentDraft async
//
// @param request - DeployDeploymentDraftAsyncRequest
//
// @param headers - DeployDeploymentDraftAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployDeploymentDraftAsyncResponse
func (client *Client) DeployDeploymentDraftAsyncWithContext(ctx context.Context, namespace *string, request *DeployDeploymentDraftAsyncRequest, headers *DeployDeploymentDraftAsyncHeaders, runtime *dara.RuntimeOptions) (_result *DeployDeploymentDraftAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployDeploymentDraftAsync"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/async-deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployDeploymentDraftAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行sql语句
//
// @param request - ExecuteSqlStatementRequest
//
// @param headers - ExecuteSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSqlStatementResponse
func (client *Client) ExecuteSqlStatementWithContext(ctx context.Context, namespace *string, request *ExecuteSqlStatementRequest, headers *ExecuteSqlStatementHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSqlStatement"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-statement/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSqlStatementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取调试结果
//
// @param request - FetchSqlPreviewResultsRequest
//
// @param headers - FetchSqlPreviewResultsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchSqlPreviewResultsResponse
func (client *Client) FetchSqlPreviewResultsWithContext(ctx context.Context, namespace *string, request *FetchSqlPreviewResultsRequest, headers *FetchSqlPreviewResultsHeaders, runtime *dara.RuntimeOptions) (_result *FetchSqlPreviewResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["queryId"] = request.QueryId
	}

	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchSqlPreviewResults"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/fetchResults"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchSqlPreviewResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a Flink request proxy.
//
// @param request - FlinkApiProxyRequest
//
// @param headers - FlinkApiProxyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlinkApiProxyResponse
func (client *Client) FlinkApiProxyWithContext(ctx context.Context, request *FlinkApiProxyRequest, headers *FlinkApiProxyHeaders, runtime *dara.RuntimeOptions) (_result *FlinkApiProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlinkApiPath) {
		query["flinkApiPath"] = request.FlinkApiPath
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlinkApiProxy"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/flink-ui/v2/proxy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlinkApiProxyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a ticket that applies for asynchronous generation of the fine-grained resources. This operation returns the ID of the ticket for you to query the asynchronous generation result.
//
// @param request - GenerateResourcePlanWithFlinkConfAsyncRequest
//
// @param headers - GenerateResourcePlanWithFlinkConfAsyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResourcePlanWithFlinkConfAsyncResponse
func (client *Client) GenerateResourcePlanWithFlinkConfAsyncWithContext(ctx context.Context, namespace *string, deploymentId *string, request *GenerateResourcePlanWithFlinkConfAsyncRequest, headers *GenerateResourcePlanWithFlinkConfAsyncHeaders, runtime *dara.RuntimeOptions) (_result *GenerateResourcePlanWithFlinkConfAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateResourcePlanWithFlinkConfAsync"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/resource-plan%3AasyncGenerate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateResourcePlanWithFlinkConfAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用中的执行定时计划
//
// @param request - GetAppliedScheduledPlanRequest
//
// @param headers - GetAppliedScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppliedScheduledPlanResponse
func (client *Client) GetAppliedScheduledPlanWithContext(ctx context.Context, namespace *string, request *GetAppliedScheduledPlanRequest, headers *GetAppliedScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *GetAppliedScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppliedScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans%3AgetExecutedScheduledPlan"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppliedScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取catalog
//
// @param request - GetCatalogsRequest
//
// @param headers - GetCatalogsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogsResponse
func (client *Client) GetCatalogsWithContext(ctx context.Context, namespace *string, request *GetCatalogsRequest, headers *GetCatalogsHeaders, runtime *dara.RuntimeOptions) (_result *GetCatalogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["catalogName"] = request.CatalogName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogs"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取database
//
// @param request - GetDatabasesRequest
//
// @param headers - GetDatabasesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabasesResponse
func (client *Client) GetDatabasesWithContext(ctx context.Context, namespace *string, catalogName *string, request *GetDatabasesRequest, headers *GetDatabasesHeaders, runtime *dara.RuntimeOptions) (_result *GetDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		query["databaseName"] = request.DatabaseName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabases"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs/" + dara.PercentEncode(dara.StringValue(catalogName)) + "/databases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get deploy deploymentDraft result
//
// @param headers - GetDeployDeploymentDraftResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployDeploymentDraftResultResponse
func (client *Client) GetDeployDeploymentDraftResultWithContext(ctx context.Context, namespace *string, ticketId *string, headers *GetDeployDeploymentDraftResultHeaders, runtime *dara.RuntimeOptions) (_result *GetDeployDeploymentDraftResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployDeploymentDraftResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/tickets/" + dara.PercentEncode(dara.StringValue(ticketId)) + "/async-deploy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeployDeploymentDraftResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a deployment.
//
// @param headers - GetDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithContext(ctx context.Context, namespace *string, deploymentId *string, headers *GetDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get a deploymentDraft
//
// @param headers - GetDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftResponse
func (client *Client) GetDeploymentDraftWithContext(ctx context.Context, namespace *string, deploymentDraftId *string, headers *GetDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentDraftResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentDraftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get deploymentDraft lock
//
// @param request - GetDeploymentDraftLockRequest
//
// @param headers - GetDeploymentDraftLockHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentDraftLockResponse
func (client *Client) GetDeploymentDraftLockWithContext(ctx context.Context, namespace *string, request *GetDeploymentDraftLockRequest, headers *GetDeploymentDraftLockHeaders, runtime *dara.RuntimeOptions) (_result *GetDeploymentDraftLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentDraftId) {
		query["deploymentDraftId"] = request.DeploymentDraftId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentDraftLock"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/getLock"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentDraftLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取运行事件
//
// @param request - GetEventsRequest
//
// @param headers - GetEventsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventsResponse
func (client *Client) GetEventsWithContext(ctx context.Context, namespace *string, request *GetEventsRequest, headers *GetEventsHeaders, runtime *dara.RuntimeOptions) (_result *GetEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvents"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get a folder
//
// @param request - GetFolderRequest
//
// @param headers - GetFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithContext(ctx context.Context, namespace *string, request *GetFolderRequest, headers *GetFolderHeaders, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the asynchronous generation result of fine-grained resources based on the ID of the ticket that applies for an asynchronous generation.
//
// @param headers - GetGenerateResourcePlanResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGenerateResourcePlanResultResponse
func (client *Client) GetGenerateResourcePlanResultWithContext(ctx context.Context, namespace *string, ticketId *string, headers *GetGenerateResourcePlanResultHeaders, runtime *dara.RuntimeOptions) (_result *GetGenerateResourcePlanResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGenerateResourcePlanResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/tickets/" + dara.PercentEncode(dara.StringValue(ticketId)) + "/resource-plan%3AasyncGenerate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGenerateResourcePlanResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动态更新结果
//
// @param headers - GetHotUpdateJobResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotUpdateJobResultResponse
func (client *Client) GetHotUpdateJobResultWithContext(ctx context.Context, namespace *string, jobHotUpdateId *string, headers *GetHotUpdateJobResultHeaders, runtime *dara.RuntimeOptions) (_result *GetHotUpdateJobResultResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotUpdateJobResult"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/hot-updates/" + dara.PercentEncode(dara.StringValue(jobHotUpdateId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotUpdateJobResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param headers - GetJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, namespace *string, jobId *string, headers *GetJobHeaders, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业诊断信息
//
// @param headers - GetJobDiagnosisHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDiagnosisResponse
func (client *Client) GetJobDiagnosisWithContext(ctx context.Context, namespace *string, deploymentId *string, jobId *string, headers *GetJobDiagnosisHeaders, runtime *dara.RuntimeOptions) (_result *GetJobDiagnosisResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDiagnosis"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/job-diagnoses/lite"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDiagnosisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the latest startup logs of a job.
//
// @param headers - GetLatestJobStartLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestJobStartLogResponse
func (client *Client) GetLatestJobStartLogWithContext(ctx context.Context, namespace *string, deploymentId *string, headers *GetLatestJobStartLogHeaders, runtime *dara.RuntimeOptions) (_result *GetLatestJobStartLogResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLatestJobStartLog"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId)) + "/latest_jobmanager_start_log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLatestJobStartLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the lineage information of a deployment.
//
// @param request - GetLineageInfoRequest
//
// @param headers - GetLineageInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineageInfoResponse
func (client *Client) GetLineageInfoWithContext(ctx context.Context, request *GetLineageInfoRequest, headers *GetLineageInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetLineageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLineageInfo"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/meta/v2/lineage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLineageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a member.
//
// @param headers - GetMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemberResponse
func (client *Client) GetMemberWithContext(ctx context.Context, namespace *string, member *string, headers *GetMemberHeaders, runtime *dara.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members/" + dara.PercentEncode(dara.StringValue(member))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传文件URL
//
// @param request - GetPreSignedUrlForPutObjectRequest
//
// @param headers - GetPreSignedUrlForPutObjectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPreSignedUrlForPutObjectResponse
func (client *Client) GetPreSignedUrlForPutObjectWithContext(ctx context.Context, namespace *string, request *GetPreSignedUrlForPutObjectRequest, headers *GetPreSignedUrlForPutObjectHeaders, runtime *dara.RuntimeOptions) (_result *GetPreSignedUrlForPutObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPreSignedUrlForPutObject"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/artifacts/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/getPreSignedUrlForPutObject"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPreSignedUrlForPutObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of a savepoint and checkpoint.
//
// @param headers - GetSavepointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavepointResponse
func (client *Client) GetSavepointWithContext(ctx context.Context, namespace *string, savepointId *string, headers *GetSavepointHeaders, runtime *dara.RuntimeOptions) (_result *GetSavepointResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavepoint"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints/" + dara.PercentEncode(dara.StringValue(savepointId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavepointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取session集群
//
// @param headers - GetSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionClusterWithContext(ctx context.Context, namespace *string, sessionClusterName *string, headers *GetSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *GetSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取table
//
// @param request - GetTablesRequest
//
// @param headers - GetTablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTablesResponse
func (client *Client) GetTablesWithContext(ctx context.Context, namespace *string, catalogName *string, databaseName *string, request *GetTablesRequest, headers *GetTablesHeaders, runtime *dara.RuntimeOptions) (_result *GetTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableName) {
		query["tableName"] = request.TableName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTables"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/catalogs/" + dara.PercentEncode(dara.StringValue(catalogName)) + "/databases/" + dara.PercentEncode(dara.StringValue(databaseName)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the JAR or Python file that corresponds to the user-defined function (UDF) that you upload and create.
//
// @param request - GetUdfArtifactsRequest
//
// @param headers - GetUdfArtifactsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfArtifactsResponse
func (client *Client) GetUdfArtifactsWithContext(ctx context.Context, namespace *string, request *GetUdfArtifactsRequest, headers *GetUdfArtifactsHeaders, runtime *dara.RuntimeOptions) (_result *GetUdfArtifactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdfArtifacts"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfArtifactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dynamically updates parameters or resources of a deployment that is running.
//
// @param headers - HotUpdateJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HotUpdateJobResponse
func (client *Client) HotUpdateJobWithContext(ctx context.Context, namespace *string, jobId *string, headers *HotUpdateJobHeaders, runtime *dara.RuntimeOptions) (_result *HotUpdateJobResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("HotUpdateJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "%3AhotUpdate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &HotUpdateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of existing custom connectors.
//
// @param headers - ListCustomConnectorsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomConnectorsResponse
func (client *Client) ListCustomConnectorsWithContext(ctx context.Context, namespace *string, headers *ListCustomConnectorsHeaders, runtime *dara.RuntimeOptions) (_result *ListCustomConnectorsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomConnectors"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomConnectorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list deploymentDrafts
//
// @param request - ListDeploymentDraftsRequest
//
// @param headers - ListDeploymentDraftsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentDraftsResponse
func (client *Client) ListDeploymentDraftsWithContext(ctx context.Context, namespace *string, request *ListDeploymentDraftsRequest, headers *ListDeploymentDraftsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentDraftsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentDrafts"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentDraftsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of clusters in which deployments can be deployed. The cluster can be a session cluster or a per-job cluster.
//
// @param request - ListDeploymentTargetsRequest
//
// @param headers - ListDeploymentTargetsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentTargetsResponse
func (client *Client) ListDeploymentTargetsWithContext(ctx context.Context, namespace *string, request *ListDeploymentTargetsRequest, headers *ListDeploymentTargetsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentTargets"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about all deployments.
//
// @param request - ListDeploymentsRequest
//
// @param headers - ListDeploymentsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithContext(ctx context.Context, namespace *string, request *ListDeploymentsRequest, headers *ListDeploymentsHeaders, runtime *dara.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["creator"] = request.Creator
	}

	if !dara.IsNil(request.ExecutionMode) {
		query["executionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.LabelKey) {
		query["labelKey"] = request.LabelKey
	}

	if !dara.IsNil(request.LabelValueArray) {
		query["labelValueArray"] = request.LabelValueArray
	}

	if !dara.IsNil(request.Modifier) {
		query["modifier"] = request.Modifier
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortName) {
		query["sortName"] = request.SortName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeployments"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出有编辑权限的项目空间。
//
// @param request - ListEditableNamespaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEditableNamespaceResponse
func (client *Client) ListEditableNamespaceWithContext(ctx context.Context, request *ListEditableNamespaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEditableNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEditableNamespace"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/editable"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEditableNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of engine versions that are supported by Realtime Compute for Apache Flink.
//
// @param headers - ListEngineVersionMetadataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineVersionMetadataResponse
func (client *Client) ListEngineVersionMetadataWithContext(ctx context.Context, headers *ListEngineVersionMetadataHeaders, runtime *dara.RuntimeOptions) (_result *ListEngineVersionMetadataResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEngineVersionMetadata"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/engine-version-meta.json"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEngineVersionMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all jobs in a deployment.
//
// @param request - ListJobsRequest
//
// @param headers - ListJobsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, namespace *string, request *ListJobsRequest, headers *ListJobsHeaders, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortName) {
		query["sortName"] = request.SortName
	}

	if !dara.IsNil(request.SortOrder) {
		query["sortOrder"] = request.SortOrder
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mappings between the ID and permissions of a member in a specific namespace.
//
// @param request - ListMembersRequest
//
// @param headers - ListMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
func (client *Client) ListMembersWithContext(ctx context.Context, namespace *string, request *ListMembersRequest, headers *ListMembersHeaders, runtime *dara.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMembers"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of savepoints or checkpoints.
//
// @param request - ListSavepointsRequest
//
// @param headers - ListSavepointsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavepointsResponse
func (client *Client) ListSavepointsWithContext(ctx context.Context, namespace *string, request *ListSavepointsRequest, headers *ListSavepointsHeaders, runtime *dara.RuntimeOptions) (_result *ListSavepointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.JobId) {
		query["jobId"] = request.JobId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSavepoints"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/savepoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSavepointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表定时执行计划
//
// @param request - ListScheduledPlanRequest
//
// @param headers - ListScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanResponse
func (client *Client) ListScheduledPlanWithContext(ctx context.Context, namespace *string, request *ListScheduledPlanRequest, headers *ListScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *ListScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取作业资源变更历史
//
// @param request - ListScheduledPlanExecutedHistoryRequest
//
// @param headers - ListScheduledPlanExecutedHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledPlanExecutedHistoryResponse
func (client *Client) ListScheduledPlanExecutedHistoryWithContext(ctx context.Context, namespace *string, request *ListScheduledPlanExecutedHistoryRequest, headers *ListScheduledPlanExecutedHistoryHeaders, runtime *dara.RuntimeOptions) (_result *ListScheduledPlanExecutedHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		query["deploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledPlanExecutedHistory"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/job-resource-upgradings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledPlanExecutedHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举session集群
//
// @param headers - ListSessionClustersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClustersWithContext(ctx context.Context, namespace *string, headers *ListSessionClustersHeaders, runtime *dara.RuntimeOptions) (_result *ListSessionClustersResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionClusters"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of variables.
//
// @param request - ListVariablesRequest
//
// @param headers - ListVariablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariablesResponse
func (client *Client) ListVariablesWithContext(ctx context.Context, namespace *string, request *ListVariablesRequest, headers *ListVariablesHeaders, runtime *dara.RuntimeOptions) (_result *ListVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["pageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVariables"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVariablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a custom connector in a namespace. The registered custom connector can be used in SQL statements.
//
// @param request - RegisterCustomConnectorRequest
//
// @param headers - RegisterCustomConnectorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomConnectorResponse
func (client *Client) RegisterCustomConnectorWithContext(ctx context.Context, namespace *string, request *RegisterCustomConnectorRequest, headers *RegisterCustomConnectorHeaders, runtime *dara.RuntimeOptions) (_result *RegisterCustomConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JarUrl) {
		query["jarUrl"] = request.JarUrl
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterCustomConnector"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/connectors%3Aregister"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterCustomConnectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers specific or all of the user-defined functions (UDFs) that are parsed from the JAR files. The registered functions can be used in SQL statements.
//
// @param request - RegisterUdfFunctionRequest
//
// @param headers - RegisterUdfFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUdfFunctionResponse
func (client *Client) RegisterUdfFunctionWithContext(ctx context.Context, namespace *string, request *RegisterUdfFunctionRequest, headers *RegisterUdfFunctionHeaders, runtime *dara.RuntimeOptions) (_result *RegisterUdfFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		query["className"] = request.ClassName
	}

	if !dara.IsNil(request.FunctionName) {
		query["functionName"] = request.FunctionName
	}

	if !dara.IsNil(request.UdfArtifactName) {
		query["udfArtifactName"] = request.UdfArtifactName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterUdfFunction"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/function%3AregisterUdfFunction"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterUdfFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI StartJob is deprecated
//
// Summary:
//
// Creates and starts a job.
//
// @param request - StartJobRequest
//
// @param headers - StartJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobResponse
func (client *Client) StartJobWithContext(ctx context.Context, namespace *string, request *StartJobRequest, headers *StartJobHeaders, runtime *dara.RuntimeOptions) (_result *StartJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a job.
//
// @param request - StartJobWithParamsRequest
//
// @param headers - StartJobWithParamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobWithParamsResponse
func (client *Client) StartJobWithParamsWithContext(ctx context.Context, namespace *string, request *StartJobWithParamsRequest, headers *StartJobWithParamsHeaders, runtime *dara.RuntimeOptions) (_result *StartJobWithParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJobWithParams"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs%3Astart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobWithParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动session集群
//
// @param headers - StartSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionClusterWithContext(ctx context.Context, namespace *string, sessionClusterName *string, headers *StartSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *StartSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName)) + "%3Astart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止应用执行定时计划
//
// @param headers - StopApplyScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApplyScheduledPlanResponse
func (client *Client) StopApplyScheduledPlanWithContext(ctx context.Context, namespace *string, scheduledPlanId *string, headers *StopApplyScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *StopApplyScheduledPlanResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApplyScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId)) + "%3Astop"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopApplyScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a job.
//
// @param request - StopJobRequest
//
// @param headers - StopJobHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
func (client *Client) StopJobWithContext(ctx context.Context, namespace *string, jobId *string, request *StopJobRequest, headers *StopJobHeaders, runtime *dara.RuntimeOptions) (_result *StopJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopJob"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "%3Astop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止session集群
//
// @param headers - StopSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionClusterWithContext(ctx context.Context, namespace *string, sessionClusterName *string, headers *StopSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *StopSessionClusterResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName)) + "%3Astop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交sql调试
//
// @param request - SubmitSqlPreviewRequest
//
// @param headers - SubmitSqlPreviewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSqlPreviewResponse
func (client *Client) SubmitSqlPreviewWithContext(ctx context.Context, namespace *string, request *SubmitSqlPreviewRequest, headers *SubmitSqlPreviewHeaders, runtime *dara.RuntimeOptions) (_result *SubmitSqlPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionClusterName) {
		query["sessionClusterName"] = request.SessionClusterName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSqlPreview"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-preview/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSqlPreviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a deployment.
//
// @param request - UpdateDeploymentRequest
//
// @param headers - UpdateDeploymentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentResponse
func (client *Client) UpdateDeploymentWithContext(ctx context.Context, namespace *string, deploymentId *string, request *UpdateDeploymentRequest, headers *UpdateDeploymentHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeployment"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployments/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// update a deploymentDraft
//
// @param request - UpdateDeploymentDraftRequest
//
// @param headers - UpdateDeploymentDraftHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentDraftResponse
func (client *Client) UpdateDeploymentDraftWithContext(ctx context.Context, namespace *string, deploymentDraftId *string, request *UpdateDeploymentDraftRequest, headers *UpdateDeploymentDraftHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentDraft"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-drafts/" + dara.PercentEncode(dara.StringValue(deploymentDraftId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentDraftResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改deploymentTarget
//
// @param request - UpdateDeploymentTargetRequest
//
// @param headers - UpdateDeploymentTargetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentTargetResponse
func (client *Client) UpdateDeploymentTargetWithContext(ctx context.Context, namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetRequest, headers *UpdateDeploymentTargetHeaders, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentTarget"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新部署目标
//
// @param request - UpdateDeploymentTargetV2Request
//
// @param headers - UpdateDeploymentTargetV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeploymentTargetV2Response
func (client *Client) UpdateDeploymentTargetV2WithContext(ctx context.Context, namespace *string, deploymentTargetName *string, request *UpdateDeploymentTargetV2Request, headers *UpdateDeploymentTargetV2Headers, runtime *dara.RuntimeOptions) (_result *UpdateDeploymentTargetV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeploymentTargetV2"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/deployment-targets/support-elastic/" + dara.PercentEncode(dara.StringValue(deploymentTargetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeploymentTargetV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// update a folder
//
// @param request - UpdateFolderRequest
//
// @param headers - UpdateFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithContext(ctx context.Context, namespace *string, folderId *string, request *UpdateFolderRequest, headers *UpdateFolderHeaders, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/folder/" + dara.PercentEncode(dara.StringValue(folderId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions of one or more members in a specific namespace.
//
// @param request - UpdateMemberRequest
//
// @param headers - UpdateMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberResponse
func (client *Client) UpdateMemberWithContext(ctx context.Context, namespace *string, request *UpdateMemberRequest, headers *UpdateMemberHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMember"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gateway/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/members"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新定时执行计划
//
// @param request - UpdateScheduledPlanRequest
//
// @param headers - UpdateScheduledPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledPlanResponse
func (client *Client) UpdateScheduledPlanWithContext(ctx context.Context, namespace *string, scheduledPlanId *string, request *UpdateScheduledPlanRequest, headers *UpdateScheduledPlanHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduledPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduledPlan"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/scheduled-plans/" + dara.PercentEncode(dara.StringValue(scheduledPlanId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduledPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新session集群
//
// @param request - UpdateSessionClusterRequest
//
// @param headers - UpdateSessionClusterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionClusterResponse
func (client *Client) UpdateSessionClusterWithContext(ctx context.Context, namespace *string, sessionClusterName *string, request *UpdateSessionClusterRequest, headers *UpdateSessionClusterHeaders, runtime *dara.RuntimeOptions) (_result *UpdateSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSessionCluster"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sessionclusters/" + dara.PercentEncode(dara.StringValue(sessionClusterName))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSessionClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the JAR file of the user-defined function (UDF) that you create.
//
// @param request - UpdateUdfArtifactRequest
//
// @param headers - UpdateUdfArtifactHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfArtifactResponse
func (client *Client) UpdateUdfArtifactWithContext(ctx context.Context, namespace *string, udfArtifactName *string, request *UpdateUdfArtifactRequest, headers *UpdateUdfArtifactHeaders, runtime *dara.RuntimeOptions) (_result *UpdateUdfArtifactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdfArtifact"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/udfartifacts/" + dara.PercentEncode(dara.StringValue(udfArtifactName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新秘钥
//
// @param request - UpdateVariableRequest
//
// @param headers - UpdateVariableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariableWithContext(ctx context.Context, namespace *string, name *string, request *UpdateVariableRequest, headers *UpdateVariableHeaders, runtime *dara.RuntimeOptions) (_result *UpdateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVariable"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/variables/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the code of an SQL deployment.
//
// @param request - ValidateSqlStatementRequest
//
// @param headers - ValidateSqlStatementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateSqlStatementResponse
func (client *Client) ValidateSqlStatementWithContext(ctx context.Context, namespace *string, request *ValidateSqlStatementRequest, headers *ValidateSqlStatementHeaders, runtime *dara.RuntimeOptions) (_result *ValidateSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Workspace) {
		realHeaders["workspace"] = dara.String(dara.ToString(dara.StringValue(headers.Workspace)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateSqlStatement"),
		Version:     dara.String("2022-07-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/sql-statement/validate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateSqlStatementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
