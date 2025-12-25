// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建应用
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		body["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EnableLog) {
		body["EnableLog"] = request.EnableLog
	}

	if !dara.IsNil(request.LabelRouteStrategy) {
		body["LabelRouteStrategy"] = request.LabelRouteStrategy
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建日历
//
// @param request - CreateCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCalendarResponse
func (client *Client) CreateCalendarWithContext(ctx context.Context, request *CreateCalendarRequest, runtime *dara.RuntimeOptions) (_result *CreateCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarName) {
		body["CalendarName"] = request.CalendarName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Months) {
		body["Months"] = request.Months
	}

	if !dara.IsNil(request.Year) {
		body["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCalendar"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建集群
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, tmpReq *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VSwitches) {
		request.VSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitches, dara.String("VSwitches"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EngineType) {
		body["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.VSwitchesShrink) {
		body["VSwitches"] = request.VSwitchesShrink
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加执行器
//
// @param request - CreateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExecutorsResponse
func (client *Client) CreateExecutorsWithContext(ctx context.Context, request *CreateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *CreateExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkerType) {
		body["WorkerType"] = request.WorkerType
	}

	if !dara.IsNil(request.Workers) {
		body["Workers"] = request.Workers
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExecutors"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExecutorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithContext(ctx context.Context, tmpReq *CreateJobRequest, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Coordinate) {
		request.CoordinateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Coordinate, dara.String("Coordinate"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NoticeConfig) {
		request.NoticeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeConfig, dara.String("NoticeConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NoticeContacts) {
		request.NoticeContactsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeContacts, dara.String("NoticeContacts"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AttemptInterval) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ChildJobId) {
		body["ChildJobId"] = request.ChildJobId
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CoordinateShrink) {
		body["Coordinate"] = request.CoordinateShrink
	}

	if !dara.IsNil(request.DependentStrategy) {
		body["DependentStrategy"] = request.DependentStrategy
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutorBlockStrategy) {
		body["ExecutorBlockStrategy"] = request.ExecutorBlockStrategy
	}

	if !dara.IsNil(request.JobHandler) {
		body["JobHandler"] = request.JobHandler
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MaxAttempt) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NoticeConfigShrink) {
		body["NoticeConfig"] = request.NoticeConfigShrink
	}

	if !dara.IsNil(request.NoticeContactsShrink) {
		body["NoticeContacts"] = request.NoticeContactsShrink
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteStrategy) {
		body["RouteStrategy"] = request.RouteStrategy
	}

	if !dara.IsNil(request.Script) {
		body["Script"] = request.Script
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeType) {
		body["StartTimeType"] = request.StartTimeType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.Weight) {
		body["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - CreateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 删除应用分组
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除日历
//
// @param request - DeleteCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCalendarResponse
func (client *Client) DeleteCalendarWithContext(ctx context.Context, request *DeleteCalendarRequest, runtime *dara.RuntimeOptions) (_result *DeleteCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarName) {
		body["CalendarName"] = request.CalendarName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Year) {
		body["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCalendar"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 释放删除集群
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除任务
//
// @param tmpReq - DeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobsResponse
func (client *Client) DeleteJobsWithContext(ctx context.Context, tmpReq *DeleteJobsRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobIdsShrink) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobsResponse{}
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DeleteJobs) {
		body["DeleteJobs"] = request.DeleteJobs
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 批量删除工作流
//
// @param tmpReq - DeleteWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowsResponse
func (client *Client) DeleteWorkflowsWithContext(ctx context.Context, tmpReq *DeleteWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowIds) {
		request.WorkflowIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowIds, dara.String("WorkflowIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DeleteJobs) {
		body["DeleteJobs"] = request.DeleteJobs
	}

	if !dara.IsNil(request.WorkflowIdsShrink) {
		body["WorkflowIds"] = request.WorkflowIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导出任务信息
//
// @param tmpReq - ExportJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportJobsResponse
func (client *Client) ExportJobsWithContext(ctx context.Context, tmpReq *ExportJobsRequest, runtime *dara.RuntimeOptions) (_result *ExportJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ExportJobType) {
		body["ExportJobType"] = request.ExportJobType
	}

	if !dara.IsNil(request.JobIdsShrink) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("byte"),
	}
	_result = &ExportJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导出工作流信息
//
// @param tmpReq - ExportWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportWorkflowsResponse
func (client *Client) ExportWorkflowsWithContext(ctx context.Context, tmpReq *ExportWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ExportWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowId) {
		request.WorkflowIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowId, dara.String("WorkflowId"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowIdShrink) {
		body["WorkflowId"] = request.WorkflowIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("byte"),
	}
	_result = &ExportWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定应用
//
// @param request - GetAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResponse
func (client *Client) GetAppWithContext(ctx context.Context, request *GetAppRequest, runtime *dara.RuntimeOptions) (_result *GetAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApp"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日历信息
//
// @param request - GetCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalendarResponse
func (client *Client) GetCalendarWithContext(ctx context.Context, request *GetCalendarRequest, runtime *dara.RuntimeOptions) (_result *GetCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalendarName) {
		query["CalendarName"] = request.CalendarName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Year) {
		query["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCalendar"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群详细信息
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithContext(ctx context.Context, request *GetClusterRequest, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定机器信息
//
// @param request - GetDesigateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDesigateInfoResponse
func (client *Client) GetDesigateInfoWithContext(ctx context.Context, request *GetDesigateInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDesigateInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDesigateInfo"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDesigateInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询执行器配置信息
//
// @param request - GetExecutorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutorConfigResponse
func (client *Client) GetExecutorConfigWithContext(ctx context.Context, request *GetExecutorConfigRequest, runtime *dara.RuntimeOptions) (_result *GetExecutorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecutorConfig"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecutorConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务执行的详细信息
//
// @param request - GetJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobExecutionResponse
func (client *Client) GetJobExecutionWithContext(ctx context.Context, request *GetJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !dara.IsNil(request.MseSessionId) {
		query["MseSessionId"] = request.MseSessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务执行的详情
//
// @param request - GetJobExecutionProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobExecutionProgressResponse
func (client *Client) GetJobExecutionProgressWithContext(ctx context.Context, request *GetJobExecutionProgressRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobExecutionProgress"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobExecutionProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务的线程堆栈
//
// @param request - GetJobExecutionThreadDumpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobExecutionThreadDumpResponse
func (client *Client) GetJobExecutionThreadDumpWithContext(ctx context.Context, request *GetJobExecutionThreadDumpRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionThreadDumpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobExecutionThreadDump"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobExecutionThreadDumpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日志
//
// @param request - GetLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogResponse
func (client *Client) GetLogWithContext(ctx context.Context, request *GetLogRequest, runtime *dara.RuntimeOptions) (_result *GetLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLog"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询事件
//
// @param request - GetLogEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogEventResponse
func (client *Client) GetLogEventWithContext(ctx context.Context, request *GetLogEventRequest, runtime *dara.RuntimeOptions) (_result *GetLogEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLogEvent"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流
//
// @param request - GetWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowResponse
func (client *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流的DAG信息
//
// @param request - GetWorkflowDAGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDAGResponse
func (client *Client) GetWorkflowDAGWithContext(ctx context.Context, request *GetWorkflowDAGRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDAGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowDAG"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowDAGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流的DAG信息
//
// @param request - GetWorkflowDAGPreviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDAGPreviewResponse
func (client *Client) GetWorkflowDAGPreviewWithContext(ctx context.Context, request *GetWorkflowDAGPreviewRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDAGPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DagVersion) {
		query["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowDAGPreview"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowDAGPreviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流实例DAG信息
//
// @param request - GetWorkflowExecutionDAGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowExecutionDAGResponse
func (client *Client) GetWorkflowExecutionDAGWithContext(ctx context.Context, request *GetWorkflowExecutionDAGRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowExecutionDAGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		query["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowExecutionDAG"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowExecutionDAGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入日历
//
// @param request - ImportCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCalendarResponse
func (client *Client) ImportCalendarWithContext(ctx context.Context, request *ImportCalendarRequest, runtime *dara.RuntimeOptions) (_result *ImportCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Months) {
		body["Months"] = request.Months
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Year) {
		body["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportCalendar"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入任务
//
// @param request - ImportJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportJobsResponse
func (client *Client) ImportJobsWithContext(ctx context.Context, request *ImportJobsRequest, runtime *dara.RuntimeOptions) (_result *ImportJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateApp) {
		body["AutoCreateApp"] = request.AutoCreateApp
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Overwrite) {
		body["Overwrite"] = request.Overwrite
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入工作流
//
// @param request - ImportWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportWorkflowsResponse
func (client *Client) ImportWorkflowsWithContext(ctx context.Context, request *ImportWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ImportWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateApp) {
		body["AutoCreateApp"] = request.AutoCreateApp
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Overwrite) {
		body["Overwrite"] = request.Overwrite
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取报警事件
//
// @param request - ListAlarmEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmEventResponse
func (client *Client) ListAlarmEventWithContext(ctx context.Context, request *ListAlarmEventRequest, runtime *dara.RuntimeOptions) (_result *ListAlarmEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlarmEvent"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlarmEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用名字列表
//
// @param request - ListAppNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppNamesResponse
func (client *Client) ListAppNamesWithContext(ctx context.Context, request *ListAppNamesRequest, runtime *dara.RuntimeOptions) (_result *ListAppNamesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppNames"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppNamesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用列表
//
// @param request - ListAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithContext(ctx context.Context, request *ListAppsRequest, runtime *dara.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApps"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日历名字列表
//
// @param request - ListCalendarNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalendarNamesResponse
func (client *Client) ListCalendarNamesWithContext(ctx context.Context, request *ListCalendarNamesRequest, runtime *dara.RuntimeOptions) (_result *ListCalendarNamesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalendarNames"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalendarNamesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日历
//
// @param request - ListCalendarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalendarsResponse
func (client *Client) ListCalendarsWithContext(ctx context.Context, request *ListCalendarsRequest, runtime *dara.RuntimeOptions) (_result *ListCalendarsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalendarName) {
		query["CalendarName"] = request.CalendarName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.FetchCalendarDetail) {
		query["FetchCalendarDetail"] = request.FetchCalendarDetail
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Year) {
		query["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalendars"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalendarsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithContext(ctx context.Context, request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Executor列表
//
// @param request - ListExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorsResponse
func (client *Client) ListExecutorsWithContext(ctx context.Context, request *ListExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutors"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutorsResponse{}
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
// @param request - ListJobExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutionsResponse
func (client *Client) ListJobExecutionsWithContext(ctx context.Context, request *ListJobExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListJobExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		query["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobExecutions"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobExecutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务脚本历史列表
//
// @param request - ListJobScriptHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobScriptHistoryResponse
func (client *Client) ListJobScriptHistoryWithContext(ctx context.Context, request *ListJobScriptHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListJobScriptHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("ListJobScriptHistory"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobScriptHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务列表
//
// @param request - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.JobHandler) {
		query["JobHandler"] = request.JobHandler
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 获取k8s资源列表
//
// @param request - ListK8sResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sResourceResponse
func (client *Client) ListK8sResourceWithContext(ctx context.Context, request *ListK8sResourceRequest, runtime *dara.RuntimeOptions) (_result *ListK8sResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.K8sClusterId) {
		query["K8sClusterId"] = request.K8sClusterId
	}

	if !dara.IsNil(request.K8sNamespace) {
		query["K8sNamespace"] = request.K8sNamespace
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListK8sResource"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListK8sResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取executor的label列表
//
// @param request - ListLablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLablesResponse
func (client *Client) ListLablesWithContext(ctx context.Context, request *ListLablesRequest, runtime *dara.RuntimeOptions) (_result *ListLablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLables"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度事件
//
// @param request - ListScheduleEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduleEventResponse
func (client *Client) ListScheduleEventWithContext(ctx context.Context, request *ListScheduleEventRequest, runtime *dara.RuntimeOptions) (_result *ListScheduleEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduleEvent"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduleEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定时间类型和表达式未来5次调度时间
//
// @param request - ListScheduleTimesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduleTimesResponse
func (client *Client) ListScheduleTimesWithContext(ctx context.Context, request *ListScheduleTimesRequest, runtime *dara.RuntimeOptions) (_result *ListScheduleTimesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduleTimes"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduleTimesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程实例列表
//
// @param request - ListWorkflowExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowExecutionsResponse
func (client *Client) ListWorkflowExecutionsWithContext(ctx context.Context, request *ListWorkflowExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		query["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowName) {
		query["WorkflowName"] = request.WorkflowName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowExecutions"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowExecutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前工作流版本列表
//
// @param request - ListWorkflowVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowVersionsResponse
func (client *Client) ListWorkflowVersionsWithContext(ctx context.Context, request *ListWorkflowVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowVersions"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowVersionsResponse{}
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflowsWithContext(ctx context.Context, request *ListWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 补数工作流
//
// @param request - OperateBackfillWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBackfillWorkflowResponse
func (client *Client) OperateBackfillWorkflowWithContext(ctx context.Context, request *OperateBackfillWorkflowRequest, runtime *dara.RuntimeOptions) (_result *OperateBackfillWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateBackfillWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateBackfillWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定执行器
//
// @param tmpReq - OperateDesignateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDesignateExecutorsResponse
func (client *Client) OperateDesignateExecutorsWithContext(ctx context.Context, tmpReq *OperateDesignateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *OperateDesignateExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateDesignateExecutorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddressList) {
		request.AddressListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddressList, dara.String("AddressList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressListShrink) {
		body["AddressList"] = request.AddressListShrink
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DesignateType) {
		body["DesignateType"] = request.DesignateType
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Transferable) {
		body["Transferable"] = request.Transferable
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateDesignateExecutors"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateDesignateExecutorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量禁用任务
//
// @param tmpReq - OperateDisableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDisableJobsResponse
func (client *Client) OperateDisableJobsWithContext(ctx context.Context, tmpReq *OperateDisableJobsRequest, runtime *dara.RuntimeOptions) (_result *OperateDisableJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateDisableJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobIdsShrink) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateDisableJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateDisableJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量禁用工作流
//
// @param tmpReq - OperateDisableWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDisableWorkflowsResponse
func (client *Client) OperateDisableWorkflowsWithContext(ctx context.Context, tmpReq *OperateDisableWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *OperateDisableWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateDisableWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowIds) {
		request.WorkflowIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowIds, dara.String("WorkflowIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowIdsShrink) {
		body["WorkflowIds"] = request.WorkflowIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateDisableWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateDisableWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量启用任务
//
// @param tmpReq - OperateEnableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateEnableJobsResponse
func (client *Client) OperateEnableJobsWithContext(ctx context.Context, tmpReq *OperateEnableJobsRequest, runtime *dara.RuntimeOptions) (_result *OperateEnableJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateEnableJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobIdsShrink) {
		body["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateEnableJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateEnableJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量启用工作流
//
// @param tmpReq - OperateEnableWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateEnableWorkflowsResponse
func (client *Client) OperateEnableWorkflowsWithContext(ctx context.Context, tmpReq *OperateEnableWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *OperateEnableWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateEnableWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WorkflowIds) {
		request.WorkflowIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkflowIds, dara.String("WorkflowIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowIdsShrink) {
		body["WorkflowIds"] = request.WorkflowIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateEnableWorkflows"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateEnableWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行一次任务
//
// @param request - OperateExecuteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateExecuteJobResponse
func (client *Client) OperateExecuteJobWithContext(ctx context.Context, request *OperateExecuteJobRequest, runtime *dara.RuntimeOptions) (_result *OperateExecuteJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceParameters) {
		body["InstanceParameters"] = request.InstanceParameters
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.Worker) {
		body["Worker"] = request.Worker
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateExecuteJob"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateExecuteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行一次工作流
//
// @param request - OperateExecuteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateExecuteWorkflowResponse
func (client *Client) OperateExecuteWorkflowWithContext(ctx context.Context, request *OperateExecuteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *OperateExecuteWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateExecuteWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateExecuteWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Hold住任务实例
//
// @param request - OperateHoldJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateHoldJobExecutionResponse
func (client *Client) OperateHoldJobExecutionWithContext(ctx context.Context, request *OperateHoldJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateHoldJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateHoldJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateHoldJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将工作流中未开始的节点置为Held状态
//
// @param request - OperateHoldWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateHoldWorkflowExecutionResponse
func (client *Client) OperateHoldWorkflowExecutionWithContext(ctx context.Context, request *OperateHoldWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateHoldWorkflowExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		body["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateHoldWorkflowExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateHoldWorkflowExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 标记任务实例为成功状态
//
// @param request - OperateMarkSuccessJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateMarkSuccessJobExecutionResponse
func (client *Client) OperateMarkSuccessJobExecutionWithContext(ctx context.Context, request *OperateMarkSuccessJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateMarkSuccessJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateMarkSuccessJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateMarkSuccessJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将工作流实例标记为成功
//
// @param request - OperateMarkSuccessWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateMarkSuccessWorkflowExecutionResponse
func (client *Client) OperateMarkSuccessWorkflowExecutionWithContext(ctx context.Context, request *OperateMarkSuccessWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateMarkSuccessWorkflowExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		body["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateMarkSuccessWorkflowExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateMarkSuccessWorkflowExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重刷任务历史数据
//
// @param request - OperateRerunJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRerunJobResponse
func (client *Client) OperateRerunJobWithContext(ctx context.Context, request *OperateRerunJobRequest, runtime *dara.RuntimeOptions) (_result *OperateRerunJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataTime) {
		query["DataTime"] = request.DataTime
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateRerunJob"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateRerunJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重跑失败的任务实例
//
// @param tmpReq - OperateRetryJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRetryJobExecutionResponse
func (client *Client) OperateRetryJobExecutionWithContext(ctx context.Context, tmpReq *OperateRetryJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateRetryJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateRetryJobExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskList) {
		request.TaskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskList, dara.String("TaskList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !dara.IsNil(request.TaskListShrink) {
		query["TaskList"] = request.TaskListShrink
	}

	if !dara.IsNil(request.TriggerChild) {
		query["TriggerChild"] = request.TriggerChild
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateRetryJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateRetryJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重跑工作流实例
//
// @param request - OperateRetryWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRetryWorkflowExecutionResponse
func (client *Client) OperateRetryWorkflowExecutionWithContext(ctx context.Context, request *OperateRetryWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateRetryWorkflowExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.OnlyFailed) {
		body["OnlyFailed"] = request.OnlyFailed
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		body["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateRetryWorkflowExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateRetryWorkflowExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跳过任务实例
//
// @param request - OperateSkipJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateSkipJobExecutionResponse
func (client *Client) OperateSkipJobExecutionWithContext(ctx context.Context, request *OperateSkipJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateSkipJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateSkipJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateSkipJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止运行中的任务实例
//
// @param tmpReq - OperateStopJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateStopJobExecutionResponse
func (client *Client) OperateStopJobExecutionWithContext(ctx context.Context, tmpReq *OperateStopJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateStopJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateStopJobExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskList) {
		request.TaskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskList, dara.String("TaskList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	if !dara.IsNil(request.TaskListShrink) {
		query["TaskList"] = request.TaskListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateStopJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateStopJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止正在运行的工作流实例
//
// @param request - OperateStopWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateStopWorkflowExecutionResponse
func (client *Client) OperateStopWorkflowExecutionWithContext(ctx context.Context, request *OperateStopWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateStopWorkflowExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		body["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateStopWorkflowExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateStopWorkflowExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将held状态的任务恢复
//
// @param request - OperateUnholdJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateUnholdJobExecutionResponse
func (client *Client) OperateUnholdJobExecutionWithContext(ctx context.Context, request *OperateUnholdJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnholdJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateUnholdJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateUnholdJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将工作流中held状态的节点恢复
//
// @param request - OperateUnholdWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateUnholdWorkflowExecutionResponse
func (client *Client) OperateUnholdWorkflowExecutionWithContext(ctx context.Context, request *OperateUnholdWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnholdWorkflowExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkflowExecutionId) {
		body["WorkflowExecutionId"] = request.WorkflowExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateUnholdWorkflowExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateUnholdWorkflowExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将skipped状态的任务恢复
//
// @param request - OperateUnskipJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateUnskipJobExecutionResponse
func (client *Client) OperateUnskipJobExecutionWithContext(ctx context.Context, request *OperateUnskipJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnskipJobExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobExecutionId) {
		query["JobExecutionId"] = request.JobExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateUnskipJobExecution"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateUnskipJobExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步任务
//
// @param tmpReq - SyncJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncJobsResponse
func (client *Client) SyncJobsWithContext(ctx context.Context, tmpReq *SyncJobsRequest, runtime *dara.RuntimeOptions) (_result *SyncJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdsShrink) {
		body["JobIds"] = request.JobIdsShrink
	}

	if !dara.IsNil(request.OriginalAppName) {
		body["OriginalAppName"] = request.OriginalAppName
	}

	if !dara.IsNil(request.OriginalClusterId) {
		body["OriginalClusterId"] = request.OriginalClusterId
	}

	if !dara.IsNil(request.TargetAppName) {
		body["TargetAppName"] = request.TargetAppName
	}

	if !dara.IsNil(request.TargetClusterId) {
		body["TargetClusterId"] = request.TargetClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用分组
//
// @param request - UpdateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppResponse
func (client *Client) UpdateAppWithContext(ctx context.Context, request *UpdateAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		body["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EnableLog) {
		body["EnableLog"] = request.EnableLog
	}

	if !dara.IsNil(request.LabelRouteStrategy) {
		body["LabelRouteStrategy"] = request.LabelRouteStrategy
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApp"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新日历
//
// @param request - UpdateCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCalendarResponse
func (client *Client) UpdateCalendarWithContext(ctx context.Context, request *UpdateCalendarRequest, runtime *dara.RuntimeOptions) (_result *UpdateCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarName) {
		body["CalendarName"] = request.CalendarName
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Incremental) {
		body["Incremental"] = request.Incremental
	}

	if !dara.IsNil(request.Months) {
		body["Months"] = request.Months
	}

	if !dara.IsNil(request.Year) {
		body["Year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCalendar"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新集群
//
// @param request - UpdateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterResponse
func (client *Client) UpdateClusterWithContext(ctx context.Context, request *UpdateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCluster"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新执行器
//
// @param request - UpdateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExecutorsResponse
func (client *Client) UpdateExecutorsWithContext(ctx context.Context, request *UpdateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *UpdateExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WorkerType) {
		body["WorkerType"] = request.WorkerType
	}

	if !dara.IsNil(request.Workers) {
		body["Workers"] = request.Workers
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExecutors"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExecutorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新任务信息
//
// @param tmpReq - UpdateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithContext(ctx context.Context, tmpReq *UpdateJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NoticeConfig) {
		request.NoticeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeConfig, dara.String("NoticeConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NoticeContacts) {
		request.NoticeContactsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NoticeContacts, dara.String("NoticeContacts"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AttemptInterval) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ChildJobId) {
		body["ChildJobId"] = request.ChildJobId
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DependentStrategy) {
		body["DependentStrategy"] = request.DependentStrategy
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutorBlockStrategy) {
		body["ExecutorBlockStrategy"] = request.ExecutorBlockStrategy
	}

	if !dara.IsNil(request.JobHandler) {
		body["JobHandler"] = request.JobHandler
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaxAttempt) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NoticeConfigShrink) {
		body["NoticeConfig"] = request.NoticeConfigShrink
	}

	if !dara.IsNil(request.NoticeContactsShrink) {
		body["NoticeContacts"] = request.NoticeContactsShrink
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteStrategy) {
		body["RouteStrategy"] = request.RouteStrategy
	}

	if !dara.IsNil(request.Script) {
		body["Script"] = request.Script
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeType) {
		body["StartTimeType"] = request.StartTimeType
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.Weight) {
		body["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJob"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新任务脚本内容
//
// @param request - UpdateJobScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobScriptResponse
func (client *Client) UpdateJobScriptWithContext(ctx context.Context, request *UpdateJobScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.ScriptContent) {
		body["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.VersionDescription) {
		body["VersionDescription"] = request.VersionDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJobScript"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobScriptResponse{}
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithContext(ctx context.Context, request *UpdateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflow"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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

// Summary:
//
// 更新工作流DAG
//
// @param tmpReq - UpdateWorkflowDAGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDAGResponse
func (client *Client) UpdateWorkflowDAGWithContext(ctx context.Context, tmpReq *UpdateWorkflowDAGRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDAGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkflowDAGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Dag) {
		request.DagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dag, dara.String("Dag"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DagShrink) {
		body["Dag"] = request.DagShrink
	}

	if !dara.IsNil(request.DagVersion) {
		body["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflowDAG"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowDAGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 切换工作流DAG版本
//
// @param request - UpdateWorkflowDAGVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDAGVersionResponse
func (client *Client) UpdateWorkflowDAGVersionWithContext(ctx context.Context, request *UpdateWorkflowDAGVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDAGVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DagVersion) {
		body["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflowDAGVersion"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowDAGVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
