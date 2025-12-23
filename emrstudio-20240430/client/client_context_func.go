// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithContext(ctx context.Context, projectId *string, request *CreateWorkflowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertGroupId) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !dara.IsNil(request.AlertStrategy) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !dara.IsNil(request.CronExpr) {
		query["cronExpr"] = request.CronExpr
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionType) {
		query["executionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.FailureStrategy) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		query["parentDirectoryId"] = request.ParentDirectoryId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleEndTime) {
		query["scheduleEndTime"] = request.ScheduleEndTime
	}

	if !dara.IsNil(request.ScheduleStartTime) {
		query["scheduleStartTime"] = request.ScheduleStartTime
	}

	if !dara.IsNil(request.ScheduleState) {
		query["scheduleState"] = request.ScheduleState
	}

	if !dara.IsNil(request.TaskDefinitionJson) {
		query["taskDefinitionJson"] = request.TaskDefinitionJson
	}

	if !dara.IsNil(request.TaskRelationJson) {
		query["taskRelationJson"] = request.TaskRelationJson
	}

	if !dara.IsNil(request.TimeZone) {
		query["timeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.Timeout) {
		query["timeout"] = request.Timeout
	}

	if !dara.IsNil(request.WorkflowInstancePriority) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !dara.IsNil(request.WorkflowParams) {
		query["workflowParams"] = request.WorkflowParams
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskDefinitionJsonValue) {
		body["taskDefinitionJsonValue"] = request.TaskDefinitionJsonValue
	}

	if !dara.IsNil(request.TaskRelationJsonValue) {
		body["taskRelationJsonValue"] = request.TaskRelationJsonValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflow"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithContext(ctx context.Context, projectId *string, workflowId *string, request *DeleteWorkflowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflow"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows/" + dara.PercentEncode(dara.StringValue(workflowId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取id关联信息
//
// @param request - DescribeIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIdResponse
func (client *Client) DescribeIdWithContext(ctx context.Context, request *DescribeIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeId"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/relatedIds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务定义
//
// @param request - DescribeManualTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeManualTaskResponse
func (client *Client) DescribeManualTaskWithContext(ctx context.Context, projectId *string, manualTaskId *string, request *DescribeManualTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeManualTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeManualTask"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/manualTasks/" + dara.PercentEncode(dara.StringValue(manualTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeManualTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务实例
//
// @param request - DescribeManualTaskInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeManualTaskInstanceResponse
func (client *Client) DescribeManualTaskInstanceWithContext(ctx context.Context, manualTaskInstanceId *string, projectId *string, request *DescribeManualTaskInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeManualTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeManualTaskInstance"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/manualTaskInstances/" + dara.PercentEncode(dara.StringValue(manualTaskInstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeManualTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - DescribeProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectResponse
func (client *Client) DescribeProjectWithContext(ctx context.Context, projectId *string, request *DescribeProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProject"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务定义
//
// @param request - DescribeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithContext(ctx context.Context, workflowId *string, projectId *string, taskId *string, request *DescribeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTask"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows/" + dara.PercentEncode(dara.StringValue(workflowId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例
//
// @param request - DescribeTaskInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskInstanceResponse
func (client *Client) DescribeTaskInstanceWithContext(ctx context.Context, projectId *string, workflowInstanceId *string, taskInstanceId *string, request *DescribeTaskInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTaskInstance"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows/" + dara.PercentEncode(dara.StringValue(workflowInstanceId)) + "/taskInstances/" + dara.PercentEncode(dara.StringValue(taskInstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - DescribeWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkflowResponse
func (client *Client) DescribeWorkflowWithContext(ctx context.Context, projectId *string, workflowId *string, request *DescribeWorkflowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWorkflow"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows/" + dara.PercentEncode(dara.StringValue(workflowId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流实例详情
//
// @param request - DescribeWorkflowInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkflowInstanceResponse
func (client *Client) DescribeWorkflowInstanceWithContext(ctx context.Context, projectId *string, workflowInstanceId *string, request *DescribeWorkflowInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeWorkflowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWorkflowInstance"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflowInstances/" + dara.PercentEncode(dara.StringValue(workflowInstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWorkflowInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例的日志
//
// @param request - GetInstanceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLogResponse
func (client *Client) GetInstanceLogWithContext(ctx context.Context, projectId *string, instanceId *string, request *GetInstanceLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceLogResponse, _err error) {
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

	if !dara.IsNil(request.SkipLineNum) {
		query["skipLineNum"] = request.SkipLineNum
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceLog"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询告警组列表
//
// @param request - ListAlertGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertGroupsResponse
func (client *Client) ListAlertGroupsWithContext(ctx context.Context, projectId *string, request *ListAlertGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlertGroupsResponse, _err error) {
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

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertGroups"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/alert-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手动任务实例列表
//
// @param request - ListManualTaskInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManualTaskInstancesResponse
func (client *Client) ListManualTaskInstancesWithContext(ctx context.Context, projectId *string, request *ListManualTaskInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListManualTaskInstancesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListManualTaskInstances"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/manualTaskInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManualTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询手动任务定义列表
//
// @param request - ListManualTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManualTasksResponse
func (client *Client) ListManualTasksWithContext(ctx context.Context, projectId *string, request *ListManualTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListManualTasksResponse, _err error) {
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

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListManualTasks"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/manualTasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManualTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目详情
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, request *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
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

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度资源组列表
//
// @param request - ListResourceGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithContext(ctx context.Context, request *ListResourceGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupName) {
		query["resourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.ResourceGroupType) {
		query["resourceGroupType"] = request.ResourceGroupType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/resourcegroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务实例列表
//
// @param request - ListTaskInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstancesWithContext(ctx context.Context, projectId *string, request *ListTaskInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTaskInstancesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		query["workflowInstanceId"] = request.WorkflowInstanceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskInstances"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/taskInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务定义列表
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, projectId *string, request *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
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

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.TaskType) {
		query["taskType"] = request.TaskType
	}

	if !dara.IsNil(request.WorkflowId) {
		query["workflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作流目录列表
//
// @param request - ListWorkflowDirectoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDirectoriesResponse
func (client *Client) ListWorkflowDirectoriesWithContext(ctx context.Context, projectId *string, request *ListWorkflowDirectoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkflowDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["directoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		query["parentDirectoryId"] = request.ParentDirectoryId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowDirectories"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/directories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流实例列表
//
// @param request - ListWorkflowInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstancesWithContext(ctx context.Context, projectId *string, request *ListWorkflowInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkflowInstancesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowId) {
		query["workflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowInstances"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflowInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流列表
//
// @param request - ListWorkflowsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflowsWithContext(ctx context.Context, projectId *string, request *ListWorkflowsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
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

	if !dara.IsNil(request.SearchVal) {
		query["searchVal"] = request.SearchVal
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflows"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作工作流实例
//
// @param request - OperateWorkflowInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateWorkflowInstanceResponse
func (client *Client) OperateWorkflowInstanceWithContext(ctx context.Context, projectId *string, request *OperateWorkflowInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateWorkflowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecType) {
		body["execType"] = request.ExecType
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		body["workflowInstanceId"] = request.WorkflowInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateWorkflowInstance"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/executors/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateWorkflowInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行工作流
//
// @param request - RunWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWorkflowResponse
func (client *Client) RunWorkflowWithContext(ctx context.Context, projectId *string, request *RunWorkflowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertGroupId) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !dara.IsNil(request.AlertStrategy) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !dara.IsNil(request.ComplementDependentMode) {
		query["complementDependentMode"] = request.ComplementDependentMode
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExecType) {
		query["execType"] = request.ExecType
	}

	if !dara.IsNil(request.ExpectedParallelismNumber) {
		query["expectedParallelismNumber"] = request.ExpectedParallelismNumber
	}

	if !dara.IsNil(request.FailureStrategy) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RunMode) {
		query["runMode"] = request.RunMode
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["scheduleTime"] = request.ScheduleTime
	}

	if !dara.IsNil(request.StartParams) {
		query["startParams"] = request.StartParams
	}

	if !dara.IsNil(request.WorkflowId) {
		query["workflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowInstancePriority) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunWorkflow"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/executors/run-workflow"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithContext(ctx context.Context, projectId *string, workflowId *string, request *UpdateWorkflowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertGroupId) {
		query["alertGroupId"] = request.AlertGroupId
	}

	if !dara.IsNil(request.AlertStrategy) {
		query["alertStrategy"] = request.AlertStrategy
	}

	if !dara.IsNil(request.CronExpr) {
		query["cronExpr"] = request.CronExpr
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionType) {
		query["executionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.FailureStrategy) {
		query["failureStrategy"] = request.FailureStrategy
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ParentDirectoryId) {
		query["parentDirectoryId"] = request.ParentDirectoryId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleEndTime) {
		query["scheduleEndTime"] = request.ScheduleEndTime
	}

	if !dara.IsNil(request.ScheduleStartTime) {
		query["scheduleStartTime"] = request.ScheduleStartTime
	}

	if !dara.IsNil(request.ScheduleState) {
		query["scheduleState"] = request.ScheduleState
	}

	if !dara.IsNil(request.TaskDefinitionJson) {
		query["taskDefinitionJson"] = request.TaskDefinitionJson
	}

	if !dara.IsNil(request.TaskRelationJson) {
		query["taskRelationJson"] = request.TaskRelationJson
	}

	if !dara.IsNil(request.TimeZone) {
		query["timeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.Timeout) {
		query["timeout"] = request.Timeout
	}

	if !dara.IsNil(request.WorkflowInstancePriority) {
		query["workflowInstancePriority"] = request.WorkflowInstancePriority
	}

	if !dara.IsNil(request.WorkflowParams) {
		query["workflowParams"] = request.WorkflowParams
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskDefinitionJsonValue) {
		body["taskDefinitionJsonValue"] = request.TaskDefinitionJsonValue
	}

	if !dara.IsNil(request.TaskRelationJsonValue) {
		body["taskRelationJsonValue"] = request.TaskRelationJsonValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflow"),
		Version:     dara.String("2024-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dolphinscheduler/v3/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/workflows/" + dara.PercentEncode(dara.StringValue(workflowId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
