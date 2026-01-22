// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Add a custom image.
//
// @param tmpReq - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithContext(ctx context.Context, tmpReq *AddImageRequest, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContainerImageSpec) {
		request.ContainerImageSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContainerImageSpec, dara.String("ContainerImageSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VMImageSpec) {
		request.VMImageSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VMImageSpec, dara.String("VMImageSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerImageSpecShrink) {
		query["ContainerImageSpec"] = request.ContainerImageSpecShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ImageVersion) {
		query["ImageVersion"] = request.ImageVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.VMImageSpecShrink) {
		query["VMImageSpec"] = request.VMImageSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImage"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a E-HPC execution plan.
//
// Description:
//
// *Make sure that you fully understand E-HPC Instnat billing methods and [prices](https://help.aliyun.com/zh/e-hpc/e-hpc-instant/product-overview/billing-overview?spm=a2c4g.11186623.help-menu-57664.d_0_2_0.5fdd28422y6UvO).
//
// @param tmpReq - CreateActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateActionPlanResponse
func (client *Client) CreateActionPlanWithContext(ctx context.Context, tmpReq *CreateActionPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateActionPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateActionPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Regions) {
		request.RegionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Regions, dara.String("Regions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Resources) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, dara.String("Resources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanName) {
		query["ActionPlanName"] = request.ActionPlanName
	}

	if !dara.IsNil(request.AllocationSpec) {
		query["AllocationSpec"] = request.AllocationSpec
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.IntervalMinutes) {
		query["IntervalMinutes"] = request.IntervalMinutes
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.PrologScript) {
		query["PrologScript"] = request.PrologScript
	}

	if !dara.IsNil(request.RegionsShrink) {
		query["Regions"] = request.RegionsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourcesShrink) {
		query["Resources"] = request.ResourcesShrink
	}

	if !dara.IsNil(request.Script) {
		query["Script"] = request.Script
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateActionPlan"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateActionPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a E-HPC Instant job.
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
	if !dara.IsNil(tmpReq.DependencyPolicy) {
		request.DependencyPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DependencyPolicy, dara.String("DependencyPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeploymentPolicy) {
		request.DeploymentPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentPolicy, dara.String("DeploymentPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityPolicy) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, dara.String("SecurityPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DependencyPolicyShrink) {
		query["DependencyPolicy"] = request.DependencyPolicyShrink
	}

	if !dara.IsNil(request.DeploymentPolicyShrink) {
		query["DeploymentPolicy"] = request.DeploymentPolicyShrink
	}

	if !dara.IsNil(request.JobDescription) {
		query["JobDescription"] = request.JobDescription
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobScheduler) {
		query["JobScheduler"] = request.JobScheduler
	}

	if !dara.IsNil(request.SecurityPolicyShrink) {
		query["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !dara.IsNil(request.TasksShrink) {
		query["Tasks"] = request.TasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2023-07-01"),
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
// Creates a resource pool.
//
// @param tmpReq - CreatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePoolResponse
func (client *Client) CreatePoolWithContext(ctx context.Context, tmpReq *CreatePoolRequest, runtime *dara.RuntimeOptions) (_result *CreatePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceLimits) {
		request.ResourceLimitsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceLimits, dara.String("ResourceLimits"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PoolName) {
		query["PoolName"] = request.PoolName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceLimitsShrink) {
		query["ResourceLimits"] = request.ResourceLimitsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePool"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete an execution plan
//
// Description:
//
// *Make sure that you fully understand E-HPC Instnat billing methods and [prices](https://help.aliyun.com/zh/e-hpc/e-hpc-instant/product-overview/billing-overview?spm=a2c4g.11186623.help-menu-57664.d_0_2_0.5fdd28422y6UvO).**
//
// This operation stops all Instant jobs that are managed by ActionPlanId.
//
// @param request - DeleteActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteActionPlanResponse
func (client *Client) DeleteActionPlanWithContext(ctx context.Context, request *DeleteActionPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteActionPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanId) {
		query["ActionPlanId"] = request.ActionPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteActionPlan"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteActionPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more job records that are in the final state from a specified cluster.
//
// @param tmpReq - DeleteJobRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobRecordsResponse
func (client *Client) DeleteJobRecordsWithContext(ctx context.Context, tmpReq *DeleteJobRecordsRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteJobRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIds) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, dara.String("JobIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIdsShrink) {
		query["JobIds"] = request.JobIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJobRecords"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more jobs from a specified cluster.
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
	if !dara.IsNil(tmpReq.ExecutorIds) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, dara.String("ExecutorIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.JobSpec) {
		request.JobSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSpec, dara.String("JobSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutorIdsShrink) {
		query["ExecutorIds"] = request.ExecutorIdsShrink
	}

	if !dara.IsNil(request.JobScheduler) {
		query["JobScheduler"] = request.JobScheduler
	}

	if !dara.IsNil(request.JobSpecShrink) {
		query["JobSpec"] = request.JobSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJobs"),
		Version:     dara.String("2023-07-01"),
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
// You can execute this statement to delete a resource pool.
//
// @param request - DeletePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePoolResponse
func (client *Client) DeletePoolWithContext(ctx context.Context, request *DeletePoolRequest, runtime *dara.RuntimeOptions) (_result *DeletePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PoolName) {
		query["PoolName"] = request.PoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePool"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can query the monitoring time series dataset of a job by specifying the job array index and query metric parameters.
//
// @param tmpReq - DescribeJobMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricDataResponse
func (client *Client) DescribeJobMetricDataWithContext(ctx context.Context, tmpReq *DescribeJobMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeJobMetricDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ArrayIndex) {
		request.ArrayIndexShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArrayIndex, dara.String("ArrayIndex"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrayIndexShrink) {
		query["ArrayIndex"] = request.ArrayIndexShrink
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobMetricData"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all instant monitoring metrics in the job array list by specifying a specific job array index list.
//
// @param tmpReq - DescribeJobMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricLastResponse
func (client *Client) DescribeJobMetricLastWithContext(ctx context.Context, tmpReq *DescribeJobMetricLastRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMetricLastResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeJobMetricLastShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ArrayIndex) {
		request.ArrayIndexShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArrayIndex, dara.String("ArrayIndex"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrayIndexShrink) {
		query["ArrayIndex"] = request.ArrayIndexShrink
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobMetricLast"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobMetricLastResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业输出
//
// @param request - DescribeJobResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobResultsResponse
func (client *Client) DescribeJobResultsWithContext(ctx context.Context, request *DescribeJobResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArrayIndex) {
		query["ArrayIndex"] = request.ArrayIndex
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.LimitBytes) {
		query["LimitBytes"] = request.LimitBytes
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobResults"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Querying Execution Plan Details
//
// @param request - GetActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActionPlanResponse
func (client *Client) GetActionPlanWithContext(ctx context.Context, request *GetActionPlanRequest, runtime *dara.RuntimeOptions) (_result *GetActionPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanId) {
		query["ActionPlanId"] = request.ActionPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetActionPlan"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetActionPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the application version list.
//
// @param request - GetAppVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppVersionsResponse
func (client *Client) GetAppVersionsWithContext(ctx context.Context, request *GetAppVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetAppVersionsResponse, _err error) {
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

	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppVersions"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about an image.
//
// @param tmpReq - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithContext(ctx context.Context, tmpReq *GetImageRequest, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdditionalRegionIds) {
		request.AdditionalRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalRegionIds, dara.String("AdditionalRegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalRegionIdsShrink) {
		query["AdditionalRegionIds"] = request.AdditionalRegionIdsShrink
	}

	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImage"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of an execution job.
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, request *GetJobRequest, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Obtains the details of a resource pool.
//
// @param request - GetPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPoolResponse
func (client *Client) GetPoolWithContext(ctx context.Context, request *GetPoolRequest, runtime *dara.RuntimeOptions) (_result *GetPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PoolName) {
		query["PoolName"] = request.PoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPool"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of an execution plan.
//
// @param request - ListActionPlanActivitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionPlanActivitiesResponse
func (client *Client) ListActionPlanActivitiesWithContext(ctx context.Context, request *ListActionPlanActivitiesRequest, runtime *dara.RuntimeOptions) (_result *ListActionPlanActivitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanId) {
		query["ActionPlanId"] = request.ActionPlanId
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
		Action:      dara.String("ListActionPlanActivities"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActionPlanActivitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of execution plans.
//
// @param tmpReq - ListActionPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionPlansResponse
func (client *Client) ListActionPlansWithContext(ctx context.Context, tmpReq *ListActionPlansRequest, runtime *dara.RuntimeOptions) (_result *ListActionPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListActionPlansShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionPlanIds) {
		request.ActionPlanIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionPlanIds, dara.String("ActionPlanIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanIdsShrink) {
		query["ActionPlanIds"] = request.ActionPlanIdsShrink
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
		Action:      dara.String("ListActionPlans"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActionPlansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running event list of one or more executers.
//
// Description:
//
// Queries job executor information.
//
// @param tmpReq - ListExecutorEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorEventsResponse
func (client *Client) ListExecutorEventsWithContext(ctx context.Context, tmpReq *ListExecutorEventsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutorEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListExecutorEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutorEvents"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutorEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Querying Global Executor Information
//
// @param tmpReq - ListExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorsResponse
func (client *Client) ListExecutorsWithContext(ctx context.Context, tmpReq *ListExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListExecutorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutors"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries the image list.
//
// @param tmpReq - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithContext(ctx context.Context, tmpReq *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageIds) {
		request.ImageIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageIds, dara.String("ImageIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !dara.IsNil(request.ImageIdsShrink) {
		query["ImageIds"] = request.ImageIdsShrink
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries job executor information.
//
// Description:
//
// Queries job executor information.
//
// @param request - ListJobExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutorsResponse
func (client *Client) ListJobExecutorsWithContext(ctx context.Context, request *ListJobExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListJobExecutorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobExecutors"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the jobs in a cluster.
//
// @param tmpReq - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, tmpReq *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SortBy) {
		request.SortByShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortBy, dara.String("SortBy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortByShrink) {
		query["SortBy"] = request.SortByShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2023-07-01"),
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
// Queries the resource pool list.
//
// @param tmpReq - ListPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoolsResponse
func (client *Client) ListPoolsWithContext(ctx context.Context, tmpReq *ListPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPoolsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPools"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are bound to one or more Instant resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResult) {
		query["MaxResult"] = request.MaxResult
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// # Remove a custom image
//
// @param request - RemoveImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageResponse
func (client *Client) RemoveImageWithContext(ctx context.Context, request *RemoveImageRequest, runtime *dara.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveImage"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Application cross-region synchronization
//
// @param tmpReq - SynchronizeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeAppResponse
func (client *Client) SynchronizeAppWithContext(ctx context.Context, tmpReq *SynchronizeAppRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SynchronizeAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TargetRegionIds) {
		request.TargetRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetRegionIds, dara.String("TargetRegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TargetRegionIdsShrink) {
		query["TargetRegionIds"] = request.TargetRegionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SynchronizeApp"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SynchronizeAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create and bind tags to Instant resource list
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
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
// Unbind tags from Instant resource list. If the tag is not bound to other resources, the tag is automatically deleted.
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithContext(ctx context.Context, request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Adjust the resource scale of the execution plan or modify the execution status
//
// @param request - UpdateActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateActionPlanResponse
func (client *Client) UpdateActionPlanWithContext(ctx context.Context, request *UpdateActionPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateActionPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionPlanId) {
		query["ActionPlanId"] = request.ActionPlanId
	}

	if !dara.IsNil(request.DesiredCapacity) {
		query["DesiredCapacity"] = request.DesiredCapacity
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.IntervalMinutes) {
		query["IntervalMinutes"] = request.IntervalMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateActionPlan"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateActionPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the resource pool configuration.
//
// @param tmpReq - UpdatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePoolResponse
func (client *Client) UpdatePoolWithContext(ctx context.Context, tmpReq *UpdatePoolRequest, runtime *dara.RuntimeOptions) (_result *UpdatePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceLimits) {
		request.ResourceLimitsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceLimits, dara.String("ResourceLimits"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PoolName) {
		query["PoolName"] = request.PoolName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceLimitsShrink) {
		query["ResourceLimits"] = request.ResourceLimitsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePool"),
		Version:     dara.String("2023-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
