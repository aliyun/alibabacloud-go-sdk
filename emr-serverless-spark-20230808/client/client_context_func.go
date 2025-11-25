// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a RAM user or RAM role to a workspace as a member.
//
// @param request - AddMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMembersResponse
func (client *Client) AddMembersWithContext(ctx context.Context, request *AddMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MemberArns) {
		body["memberArns"] = request.MemberArns
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMembers"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/auth/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a Spark job.
//
// @param request - CancelJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelJobRunResponse
func (client *Client) CancelJobRunWithContext(ctx context.Context, workspaceId *string, jobRunId *string, request *CancelJobRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelJobRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelJobRun"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/jobRuns/" + dara.PercentEncode(dara.StringValue(jobRunId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CancelKyuubiSparkApplication
//
// @param request - CancelKyuubiSparkApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelKyuubiSparkApplicationResponse
func (client *Client) CancelKyuubiSparkApplicationWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, applicationId *string, request *CancelKyuubiSparkApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelKyuubiSparkApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelKyuubiSparkApplication"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/application/" + dara.PercentEncode(dara.StringValue(applicationId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelKyuubiSparkApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateKyuubiService
//
// @param request - CreateKyuubiServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKyuubiServiceResponse
func (client *Client) CreateKyuubiServiceWithContext(ctx context.Context, workspaceId *string, request *CreateKyuubiServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKyuubiServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComputeInstance) {
		body["computeInstance"] = request.ComputeInstance
	}

	if !dara.IsNil(request.KyuubiConfigs) {
		body["kyuubiConfigs"] = request.KyuubiConfigs
	}

	if !dara.IsNil(request.KyuubiReleaseVersion) {
		body["kyuubiReleaseVersion"] = request.KyuubiReleaseVersion
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PublicEndpointEnabled) {
		body["publicEndpointEnabled"] = request.PublicEndpointEnabled
	}

	if !dara.IsNil(request.Queue) {
		body["queue"] = request.Queue
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.Replica) {
		body["replica"] = request.Replica
	}

	if !dara.IsNil(request.SparkConfigs) {
		body["sparkConfigs"] = request.SparkConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建kyuubi的token
//
// @param request - CreateKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKyuubiTokenResponse
func (client *Client) CreateKyuubiTokenWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, request *CreateKyuubiTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKyuubiTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoExpireConfiguration) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !dara.IsNil(request.MemberArns) {
		body["memberArns"] = request.MemberArns
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Token) {
		body["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKyuubiToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/kyuubiService/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKyuubiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Livy compute
//
// @param request - CreateLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivyComputeResponse
func (client *Client) CreateLivyComputeWithContext(ctx context.Context, workspaceBizId *string, request *CreateLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		body["authType"] = request.AuthType
	}

	if !dara.IsNil(request.AutoStartConfiguration) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !dara.IsNil(request.AutoStopConfiguration) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !dara.IsNil(request.CpuLimit) {
		body["cpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.DisplayReleaseVersion) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !dara.IsNil(request.EnablePublic) {
		body["enablePublic"] = request.EnablePublic
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Fusion) {
		body["fusion"] = request.Fusion
	}

	if !dara.IsNil(request.LivyServerConf) {
		body["livyServerConf"] = request.LivyServerConf
	}

	if !dara.IsNil(request.LivyVersion) {
		body["livyVersion"] = request.LivyVersion
	}

	if !dara.IsNil(request.MemoryLimit) {
		body["memoryLimit"] = request.MemoryLimit
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NetworkName) {
		body["networkName"] = request.NetworkName
	}

	if !dara.IsNil(request.QueueName) {
		body["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Livy Compute的token
//
// @param request - CreateLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivyComputeTokenResponse
func (client *Client) CreateLivyComputeTokenWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *CreateLivyComputeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLivyComputeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoExpireConfiguration) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Token) {
		body["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivyComputeToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/token"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivyComputeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow.
//
// @param tmpReq - CreateProcessDefinitionWithScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProcessDefinitionWithScheduleResponse
func (client *Client) CreateProcessDefinitionWithScheduleWithContext(ctx context.Context, bizId *string, tmpReq *CreateProcessDefinitionWithScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProcessDefinitionWithScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProcessDefinitionWithScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalParams) {
		request.GlobalParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalParams, dara.String("globalParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("schedule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskDefinitionJson) {
		request.TaskDefinitionJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDefinitionJson, dara.String("taskDefinitionJson"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskRelationJson) {
		request.TaskRelationJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskRelationJson, dara.String("taskRelationJson"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertEmailAddress) {
		query["alertEmailAddress"] = request.AlertEmailAddress
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionType) {
		query["executionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.GlobalParamsShrink) {
		query["globalParams"] = request.GlobalParamsShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ProductNamespace) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !dara.IsNil(request.Publish) {
		query["publish"] = request.Publish
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceQueue) {
		query["resourceQueue"] = request.ResourceQueue
	}

	if !dara.IsNil(request.RetryTimes) {
		query["retryTimes"] = request.RetryTimes
	}

	if !dara.IsNil(request.RunAs) {
		query["runAs"] = request.RunAs
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDefinitionJsonShrink) {
		query["taskDefinitionJson"] = request.TaskDefinitionJsonShrink
	}

	if !dara.IsNil(request.TaskParallelism) {
		query["taskParallelism"] = request.TaskParallelism
	}

	if !dara.IsNil(request.TaskRelationJsonShrink) {
		query["taskRelationJson"] = request.TaskRelationJsonShrink
	}

	if !dara.IsNil(request.Timeout) {
		query["timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProcessDefinitionWithSchedule"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/projects/" + dara.PercentEncode(dara.StringValue(bizId)) + "/process-definition"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// @param request - CreateSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionClusterResponse
func (client *Client) CreateSessionClusterWithContext(ctx context.Context, workspaceId *string, request *CreateSessionClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigs) {
		body["applicationConfigs"] = request.ApplicationConfigs
	}

	if !dara.IsNil(request.AutoStartConfiguration) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !dara.IsNil(request.AutoStopConfiguration) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisplayReleaseVersion) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !dara.IsNil(request.EnvId) {
		body["envId"] = request.EnvId
	}

	if !dara.IsNil(request.Fusion) {
		body["fusion"] = request.Fusion
	}

	if !dara.IsNil(request.Kind) {
		body["kind"] = request.Kind
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PublicEndpointEnabled) {
		body["publicEndpointEnabled"] = request.PublicEndpointEnabled
	}

	if !dara.IsNil(request.QueueName) {
		body["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSessionCluster"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/sessionClusters"),
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
// Creates an SQL query task.
//
// @param request - CreateSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSqlStatementResponse
func (client *Client) CreateSqlStatementWithContext(ctx context.Context, workspaceId *string, request *CreateSqlStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeContent) {
		body["codeContent"] = request.CodeContent
	}

	if !dara.IsNil(request.DefaultCatalog) {
		body["defaultCatalog"] = request.DefaultCatalog
	}

	if !dara.IsNil(request.DefaultDatabase) {
		body["defaultDatabase"] = request.DefaultDatabase
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.SqlComputeId) {
		body["sqlComputeId"] = request.SqlComputeId
	}

	if !dara.IsNil(request.TaskBizId) {
		body["taskBizId"] = request.TaskBizId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSqlStatement"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/statement"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSqlStatementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		body["autoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		body["autoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.AutoRenewPeriodUnit) {
		body["autoRenewPeriodUnit"] = request.AutoRenewPeriodUnit
	}

	if !dara.IsNil(request.AutoStartSessionCluster) {
		body["autoStartSessionCluster"] = request.AutoStartSessionCluster
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DlfCatalogId) {
		body["dlfCatalogId"] = request.DlfCatalogId
	}

	if !dara.IsNil(request.DlfType) {
		body["dlfType"] = request.DlfType
	}

	if !dara.IsNil(request.Duration) {
		body["duration"] = request.Duration
	}

	if !dara.IsNil(request.OssBucket) {
		body["ossBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.PaymentDurationUnit) {
		body["paymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.RamRoleName) {
		body["ramRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.ReleaseType) {
		body["releaseType"] = request.ReleaseType
	}

	if !dara.IsNil(request.ResourceSpec) {
		body["resourceSpec"] = request.ResourceSpec
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	if !dara.IsNil(request.WorkspaceName) {
		body["workspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteKyuubiService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKyuubiServiceResponse
func (client *Client) DeleteKyuubiServiceWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKyuubiServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除compute的token
//
// @param request - DeleteKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKyuubiTokenResponse
func (client *Client) DeleteKyuubiTokenWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, tokenId *string, request *DeleteKyuubiTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKyuubiTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKyuubiToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/kyuubiService/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKyuubiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除livy compute
//
// @param request - DeleteLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivyComputeResponse
func (client *Client) DeleteLivyComputeWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *DeleteLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Livy Compute的token
//
// @param request - DeleteLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivyComputeTokenResponse
func (client *Client) DeleteLivyComputeTokenWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, tokenId *string, request *DeleteLivyComputeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLivyComputeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivyComputeToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivyComputeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the queue of a workspace.
//
// @param request - EditWorkspaceQueueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditWorkspaceQueueResponse
func (client *Client) EditWorkspaceQueueWithContext(ctx context.Context, request *EditWorkspaceQueueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EditWorkspaceQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Environments) {
		body["environments"] = request.Environments
	}

	if !dara.IsNil(request.ResourceSpec) {
		body["resourceSpec"] = request.ResourceSpec
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceQueueName) {
		body["workspaceQueueName"] = request.WorkspaceQueueName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditWorkspaceQueue"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/queues/action/edit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EditWorkspaceQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线工作流及其调度
//
// @param request - GenerateTaskCodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTaskCodesResponse
func (client *Client) GenerateTaskCodesWithContext(ctx context.Context, bizId *string, request *GenerateTaskCodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTaskCodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GenNum) {
		query["genNum"] = request.GenNum
	}

	if !dara.IsNil(request.ProductNamespace) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTaskCodes"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/projects/" + dara.PercentEncode(dara.StringValue(bizId)) + "/task-definition/gen-task-codes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTaskCodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of CU-hours consumed by a queue during a specified cycle.
//
// @param request - GetCuHoursRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCuHoursResponse
func (client *Client) GetCuHoursWithContext(ctx context.Context, workspaceId *string, queue *string, request *GetCuHoursRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCuHoursResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCuHours"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/metric/cuHours/" + dara.PercentEncode(dara.StringValue(queue))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCuHoursResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains job analysis information on E-MapReduce (EMR) Doctor.
//
// @param request - GetDoctorApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorApplicationResponse
func (client *Client) GetDoctorApplicationWithContext(ctx context.Context, workspaceId *string, runId *string, request *GetDoctorApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDoctorApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Locale) {
		query["locale"] = request.Locale
	}

	if !dara.IsNil(request.QueryTime) {
		query["queryTime"] = request.QueryTime
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorApplication"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/runs/" + dara.PercentEncode(dara.StringValue(runId)) + "/action/getDoctorApplication"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the job details.
//
// @param request - GetJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobRunResponse
func (client *Client) GetJobRunWithContext(ctx context.Context, workspaceId *string, jobRunId *string, request *GetJobRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobRun"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/jobRuns/" + dara.PercentEncode(dara.StringValue(jobRunId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetKyuubiService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKyuubiServiceResponse
func (client *Client) GetKyuubiServiceWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKyuubiServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取compute的token
//
// @param request - GetKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKyuubiTokenResponse
func (client *Client) GetKyuubiTokenWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, tokenId *string, request *GetKyuubiTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKyuubiTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKyuubiToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/kyuubiService/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKyuubiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取livy compute
//
// @param request - GetLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivyComputeResponse
func (client *Client) GetLivyComputeWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *GetLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取livy compute token
//
// @param request - GetLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivyComputeTokenResponse
func (client *Client) GetLivyComputeTokenWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, tokenId *string, request *GetLivyComputeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLivyComputeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLivyComputeToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLivyComputeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a session.
//
// @param request - GetSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionClusterResponse
func (client *Client) GetSessionClusterWithContext(ctx context.Context, workspaceId *string, sessionClusterId *string, request *GetSessionClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSessionCluster"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/sessionClusters/" + dara.PercentEncode(dara.StringValue(sessionClusterId))),
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
// Queries the status of an SQL query task.
//
// @param request - GetSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlStatementResponse
func (client *Client) GetSqlStatementWithContext(ctx context.Context, workspaceId *string, statementId *string, request *GetSqlStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlStatement"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/statement/" + dara.PercentEncode(dara.StringValue(statementId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlStatementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries task templates.
//
// @param request - GetTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithContext(ctx context.Context, workspaceBizId *string, request *GetTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateBizId) {
		query["templateBizId"] = request.TemplateBizId
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/template"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a specified role to users.
//
// @param request - GrantRoleToUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsersWithContext(ctx context.Context, request *GrantRoleToUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleArn) {
		body["roleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.UserArns) {
		body["userArns"] = request.UserArns
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantRoleToUsers"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/auth/roles/grant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据目录列表
//
// @param request - ListCatalogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogsWithContext(ctx context.Context, workspaceId *string, request *ListCatalogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["environment"] = request.Environment
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogs"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/catalogs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出作业的executors
//
// @param request - ListJobExecutorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutorsResponse
func (client *Client) ListJobExecutorsWithContext(ctx context.Context, workspaceId *string, jobRunId *string, request *ListJobExecutorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutorType) {
		query["executorType"] = request.ExecutorType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobExecutors"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/jobRuns/" + dara.PercentEncode(dara.StringValue(jobRunId)) + "/executors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobExecutorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Spark jobs.
//
// @param tmpReq - ListJobRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobRunsResponse
func (client *Client) ListJobRunsWithContext(ctx context.Context, workspaceId *string, tmpReq *ListJobRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListJobRunsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EndTime) {
		request.EndTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EndTime, dara.String("endTime"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartTime) {
		request.StartTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTime, dara.String("startTime"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.States) {
		request.StatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.States, dara.String("states"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigs) {
		query["applicationConfigs"] = request.ApplicationConfigs
	}

	if !dara.IsNil(request.Creator) {
		query["creator"] = request.Creator
	}

	if !dara.IsNil(request.EndTimeShrink) {
		query["endTime"] = request.EndTimeShrink
	}

	if !dara.IsNil(request.IsWorkflow) {
		query["isWorkflow"] = request.IsWorkflow
	}

	if !dara.IsNil(request.JobRunDeploymentId) {
		query["jobRunDeploymentId"] = request.JobRunDeploymentId
	}

	if !dara.IsNil(request.JobRunId) {
		query["jobRunId"] = request.JobRunId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MinDuration) {
		query["minDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceQueueId) {
		query["resourceQueueId"] = request.ResourceQueueId
	}

	if !dara.IsNil(request.RuntimeConfigs) {
		query["runtimeConfigs"] = request.RuntimeConfigs
	}

	if !dara.IsNil(request.StartTimeShrink) {
		query["startTime"] = request.StartTimeShrink
	}

	if !dara.IsNil(request.StatesShrink) {
		query["states"] = request.StatesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobRuns"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/jobRuns"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListKyuubiServices
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiServicesResponse
func (client *Client) ListKyuubiServicesWithContext(ctx context.Context, workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKyuubiServicesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKyuubiServices"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKyuubiServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that are submitted by using a Kyuubi gateway.
//
// @param tmpReq - ListKyuubiSparkApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiSparkApplicationsResponse
func (client *Client) ListKyuubiSparkApplicationsWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, tmpReq *ListKyuubiSparkApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKyuubiSparkApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListKyuubiSparkApplicationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderBy) {
		request.OrderByShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderBy, dara.String("orderBy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartTime) {
		request.StartTimeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartTime, dara.String("startTime"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["applicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationName) {
		query["applicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MinDuration) {
		query["minDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderByShrink) {
		query["orderBy"] = request.OrderByShrink
	}

	if !dara.IsNil(request.ResourceQueueId) {
		query["resourceQueueId"] = request.ResourceQueueId
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTimeShrink) {
		query["startTime"] = request.StartTimeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKyuubiSparkApplications"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/applications"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKyuubiSparkApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出compute的token
//
// @param request - ListKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKyuubiTokenResponse
func (client *Client) ListKyuubiTokenWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, request *ListKyuubiTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKyuubiTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKyuubiToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/kyuubiService/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/token"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKyuubiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出livy compute
//
// @param request - ListLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivyComputeResponse
func (client *Client) ListLivyComputeWithContext(ctx context.Context, workspaceBizId *string, request *ListLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出livy compute token
//
// @param request - ListLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivyComputeTokenResponse
func (client *Client) ListLivyComputeTokenWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *ListLivyComputeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLivyComputeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLivyComputeToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/token"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLivyComputeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Log Content
//
// @param request - ListLogContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogContentsResponse
func (client *Client) ListLogContentsWithContext(ctx context.Context, workspaceId *string, request *ListLogContentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogContentsResponse, _err error) {
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

	if !dara.IsNil(request.Length) {
		query["length"] = request.Length
	}

	if !dara.IsNil(request.Offset) {
		query["offset"] = request.Offset
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogContents"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/action/listLogContents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogContentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户列表
//
// @param request - ListMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
func (client *Client) ListMembersWithContext(ctx context.Context, workspaceId *string, request *ListMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMembers"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/auth/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/members"),
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
// Queries the list of published versions of E-MapReduce (EMR) Serverless Spark.
//
// @param request - ListReleaseVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReleaseVersionsResponse
func (client *Client) ListReleaseVersionsWithContext(ctx context.Context, request *ListReleaseVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReleaseVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseType) {
		query["releaseType"] = request.ReleaseType
	}

	if !dara.IsNil(request.ReleaseVersion) {
		query["releaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ReleaseVersionStatus) {
		query["releaseVersionStatus"] = request.ReleaseVersionStatus
	}

	if !dara.IsNil(request.ServiceFilter) {
		query["serviceFilter"] = request.ServiceFilter
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReleaseVersions"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/releaseVersions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReleaseVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of sessions.
//
// @param request - ListSessionClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionClustersResponse
func (client *Client) ListSessionClustersWithContext(ctx context.Context, workspaceId *string, request *ListSessionClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSessionClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Kind) {
		query["kind"] = request.Kind
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.QueueName) {
		query["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionClusterId) {
		query["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionClusters"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/sessionClusters"),
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
// 获取sql statement内容
//
// @param request - ListSqlStatementContentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSqlStatementContentsResponse
func (client *Client) ListSqlStatementContentsWithContext(ctx context.Context, workspaceId *string, request *ListSqlStatementContentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSqlStatementContentsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSqlStatementContents"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/action/listSqlStatementContents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSqlStatementContentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务模板列表
//
// @param request - ListTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateResponse
func (client *Client) ListTemplateWithContext(ctx context.Context, workspaceBizId *string, request *ListTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplate"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/template/listing"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of queues in a Spark workspace.
//
// @param request - ListWorkspaceQueuesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceQueuesResponse
func (client *Client) ListWorkspaceQueuesWithContext(ctx context.Context, workspaceId *string, request *ListWorkspaceQueuesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspaceQueuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["environment"] = request.Environment
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceQueues"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/queues"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceQueuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workspaces.
//
// @param tmpReq - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithContext(ctx context.Context, tmpReq *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Livy Compute的token
//
// @param request - RefreshLivyComputeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLivyComputeTokenResponse
func (client *Client) RefreshLivyComputeTokenWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, tokenId *string, request *RefreshLivyComputeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RefreshLivyComputeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoExpireConfiguration) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Token) {
		body["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshLivyComputeToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshLivyComputeTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a Spark job.
//
// @param request - StartJobRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobRunResponse
func (client *Client) StartJobRunWithContext(ctx context.Context, workspaceId *string, request *StartJobRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartJobRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CodeType) {
		body["codeType"] = request.CodeType
	}

	if !dara.IsNil(request.ConfigurationOverrides) {
		body["configurationOverrides"] = request.ConfigurationOverrides
	}

	if !dara.IsNil(request.DisplayReleaseVersion) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !dara.IsNil(request.ExecutionTimeoutSeconds) {
		body["executionTimeoutSeconds"] = request.ExecutionTimeoutSeconds
	}

	if !dara.IsNil(request.Fusion) {
		body["fusion"] = request.Fusion
	}

	if !dara.IsNil(request.JobDriver) {
		body["jobDriver"] = request.JobDriver
	}

	if !dara.IsNil(request.JobId) {
		body["jobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ResourceQueueId) {
		body["resourceQueueId"] = request.ResourceQueueId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJobRun"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/jobRuns"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartKyuubiService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartKyuubiServiceResponse
func (client *Client) StartKyuubiServiceWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartKyuubiServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动livy compute
//
// @param request - StartLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLivyComputeResponse
func (client *Client) StartLivyComputeWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *StartLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually runs a workflow.
//
// @param request - StartProcessInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartProcessInstanceResponse
func (client *Client) StartProcessInstanceWithContext(ctx context.Context, bizId *string, request *StartProcessInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartProcessInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	if !dara.IsNil(request.Comments) {
		query["comments"] = request.Comments
	}

	if !dara.IsNil(request.Email) {
		query["email"] = request.Email
	}

	if !dara.IsNil(request.Interval) {
		query["interval"] = request.Interval
	}

	if !dara.IsNil(request.IsProd) {
		query["isProd"] = request.IsProd
	}

	if !dara.IsNil(request.ProcessDefinitionCode) {
		query["processDefinitionCode"] = request.ProcessDefinitionCode
	}

	if !dara.IsNil(request.ProductNamespace) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuntimeQueue) {
		query["runtimeQueue"] = request.RuntimeQueue
	}

	if !dara.IsNil(request.VersionHashCode) {
		query["versionHashCode"] = request.VersionHashCode
	}

	if !dara.IsNil(request.VersionNumber) {
		query["versionNumber"] = request.VersionNumber
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartProcessInstance"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/projects/" + dara.PercentEncode(dara.StringValue(bizId)) + "/executors/start-process-instance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a session.
//
// @param request - StartSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSessionClusterResponse
func (client *Client) StartSessionClusterWithContext(ctx context.Context, workspaceId *string, request *StartSessionClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueueName) {
		body["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.SessionClusterId) {
		body["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSessionCluster"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/sessionClusters/action/startSessionCluster"),
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
// # StopKyuubiService
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopKyuubiServiceResponse
func (client *Client) StopKyuubiServiceWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopKyuubiServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止livy compute
//
// @param request - StopLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLivyComputeResponse
func (client *Client) StopLivyComputeWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *StopLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a session.
//
// @param request - StopSessionClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSessionClusterResponse
func (client *Client) StopSessionClusterWithContext(ctx context.Context, workspaceId *string, request *StopSessionClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSessionClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueueName) {
		body["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.SessionClusterId) {
		body["sessionClusterId"] = request.SessionClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSessionCluster"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/sessionClusters/action/stopSessionCluster"),
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
// Terminates an SQL query task.
//
// @param request - TerminateSqlStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateSqlStatementResponse
func (client *Client) TerminateSqlStatementWithContext(ctx context.Context, workspaceId *string, statementId *string, request *TerminateSqlStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TerminateSqlStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateSqlStatement"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/statement/" + dara.PercentEncode(dara.StringValue(statementId)) + "/terminate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateSqlStatementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateKyuubiService
//
// @param request - UpdateKyuubiServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKyuubiServiceResponse
func (client *Client) UpdateKyuubiServiceWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, request *UpdateKyuubiServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKyuubiServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComputeInstance) {
		body["computeInstance"] = request.ComputeInstance
	}

	if !dara.IsNil(request.KyuubiConfigs) {
		body["kyuubiConfigs"] = request.KyuubiConfigs
	}

	if !dara.IsNil(request.KyuubiReleaseVersion) {
		body["kyuubiReleaseVersion"] = request.KyuubiReleaseVersion
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PublicEndpointEnabled) {
		body["publicEndpointEnabled"] = request.PublicEndpointEnabled
	}

	if !dara.IsNil(request.Queue) {
		body["queue"] = request.Queue
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.Replica) {
		body["replica"] = request.Replica
	}

	if !dara.IsNil(request.Restart) {
		body["restart"] = request.Restart
	}

	if !dara.IsNil(request.SparkConfigs) {
		body["sparkConfigs"] = request.SparkConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKyuubiService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/kyuubi/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKyuubiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新kyuubi的token
//
// @param request - UpdateKyuubiTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKyuubiTokenResponse
func (client *Client) UpdateKyuubiTokenWithContext(ctx context.Context, workspaceId *string, kyuubiServiceId *string, tokenId *string, request *UpdateKyuubiTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKyuubiTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoExpireConfiguration) {
		body["autoExpireConfiguration"] = request.AutoExpireConfiguration
	}

	if !dara.IsNil(request.MemberArns) {
		body["memberArns"] = request.MemberArns
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Token) {
		body["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKyuubiToken"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/kyuubiService/" + dara.PercentEncode(dara.StringValue(kyuubiServiceId)) + "/token/" + dara.PercentEncode(dara.StringValue(tokenId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKyuubiTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新livy compute
//
// @param request - UpdateLivyComputeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivyComputeResponse
func (client *Client) UpdateLivyComputeWithContext(ctx context.Context, workspaceBizId *string, livyComputeId *string, request *UpdateLivyComputeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLivyComputeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		body["authType"] = request.AuthType
	}

	if !dara.IsNil(request.AutoStartConfiguration) {
		body["autoStartConfiguration"] = request.AutoStartConfiguration
	}

	if !dara.IsNil(request.AutoStopConfiguration) {
		body["autoStopConfiguration"] = request.AutoStopConfiguration
	}

	if !dara.IsNil(request.CpuLimit) {
		body["cpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.DisplayReleaseVersion) {
		body["displayReleaseVersion"] = request.DisplayReleaseVersion
	}

	if !dara.IsNil(request.EnablePublic) {
		body["enablePublic"] = request.EnablePublic
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Fusion) {
		body["fusion"] = request.Fusion
	}

	if !dara.IsNil(request.LivyServerConf) {
		body["livyServerConf"] = request.LivyServerConf
	}

	if !dara.IsNil(request.LivyVersion) {
		body["livyVersion"] = request.LivyVersion
	}

	if !dara.IsNil(request.MemoryLimit) {
		body["memoryLimit"] = request.MemoryLimit
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NetworkName) {
		body["networkName"] = request.NetworkName
	}

	if !dara.IsNil(request.QueueName) {
		body["queueName"] = request.QueueName
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["releaseVersion"] = request.ReleaseVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivyCompute"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/interactive/v1/workspace/" + dara.PercentEncode(dara.StringValue(workspaceBizId)) + "/livycompute/" + dara.PercentEncode(dara.StringValue(livyComputeId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivyComputeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the workflow and time-based scheduling configurations.
//
// @param tmpReq - UpdateProcessDefinitionWithScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProcessDefinitionWithScheduleResponse
func (client *Client) UpdateProcessDefinitionWithScheduleWithContext(ctx context.Context, bizId *string, code *string, tmpReq *UpdateProcessDefinitionWithScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProcessDefinitionWithScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProcessDefinitionWithScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalParams) {
		request.GlobalParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalParams, dara.String("globalParams"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Schedule) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, dara.String("schedule"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskDefinitionJson) {
		request.TaskDefinitionJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskDefinitionJson, dara.String("taskDefinitionJson"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskRelationJson) {
		request.TaskRelationJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskRelationJson, dara.String("taskRelationJson"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertEmailAddress) {
		query["alertEmailAddress"] = request.AlertEmailAddress
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionType) {
		query["executionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.GlobalParamsShrink) {
		query["globalParams"] = request.GlobalParamsShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ProductNamespace) {
		query["productNamespace"] = request.ProductNamespace
	}

	if !dara.IsNil(request.Publish) {
		query["publish"] = request.Publish
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseState) {
		query["releaseState"] = request.ReleaseState
	}

	if !dara.IsNil(request.ResourceQueue) {
		query["resourceQueue"] = request.ResourceQueue
	}

	if !dara.IsNil(request.RetryTimes) {
		query["retryTimes"] = request.RetryTimes
	}

	if !dara.IsNil(request.RunAs) {
		query["runAs"] = request.RunAs
	}

	if !dara.IsNil(request.ScheduleShrink) {
		query["schedule"] = request.ScheduleShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskDefinitionJsonShrink) {
		query["taskDefinitionJson"] = request.TaskDefinitionJsonShrink
	}

	if !dara.IsNil(request.TaskParallelism) {
		query["taskParallelism"] = request.TaskParallelism
	}

	if !dara.IsNil(request.TaskRelationJsonShrink) {
		query["taskRelationJson"] = request.TaskRelationJsonShrink
	}

	if !dara.IsNil(request.Timeout) {
		query["timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProcessDefinitionWithSchedule"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/projects/" + dara.PercentEncode(dara.StringValue(bizId)) + "/process-definition/" + dara.PercentEncode(dara.StringValue(code))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProcessDefinitionWithScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
