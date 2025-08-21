// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ehpcinstant"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加托管侧用户自定义镜像
//
// @param tmpReq - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithOptions(tmpReq *AddImageRequest, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加托管侧用户自定义镜像
//
// @param request - AddImageRequest
//
// @return AddImageResponse
func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建执行计划创建执行计划
//
// @param tmpReq - CreateActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateActionPlanResponse
func (client *Client) CreateActionPlanWithOptions(tmpReq *CreateActionPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateActionPlanResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建执行计划创建执行计划
//
// @param request - CreateActionPlanRequest
//
// @return CreateActionPlanResponse
func (client *Client) CreateActionPlan(request *CreateActionPlanRequest) (_result *CreateActionPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateActionPlanResponse{}
	_body, _err := client.CreateActionPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交任务
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交任务
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源池
//
// @param tmpReq - CreatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePoolResponse
func (client *Client) CreatePoolWithOptions(tmpReq *CreatePoolRequest, runtime *dara.RuntimeOptions) (_result *CreatePoolResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源池
//
// @param request - CreatePoolRequest
//
// @return CreatePoolResponse
func (client *Client) CreatePool(request *CreatePoolRequest) (_result *CreatePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePoolResponse{}
	_body, _err := client.CreatePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除执行计划
//
// @param request - DeleteActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteActionPlanResponse
func (client *Client) DeleteActionPlanWithOptions(request *DeleteActionPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteActionPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除执行计划
//
// @param request - DeleteActionPlanRequest
//
// @return DeleteActionPlanResponse
func (client *Client) DeleteActionPlan(request *DeleteActionPlanRequest) (_result *DeleteActionPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteActionPlanResponse{}
	_body, _err := client.DeleteActionPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除作业
//
// @param tmpReq - DeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobsResponse
func (client *Client) DeleteJobsWithOptions(tmpReq *DeleteJobsRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除作业
//
// @param request - DeleteJobsRequest
//
// @return DeleteJobsResponse
func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源池
//
// @param request - DeletePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePoolResponse
func (client *Client) DeletePoolWithOptions(request *DeletePoolRequest, runtime *dara.RuntimeOptions) (_result *DeletePoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源池
//
// @param request - DeletePoolRequest
//
// @return DeletePoolResponse
func (client *Client) DeletePool(request *DeletePoolRequest) (_result *DeletePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePoolResponse{}
	_body, _err := client.DeletePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业性能数据
//
// @param tmpReq - DescribeJobMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricDataResponse
func (client *Client) DescribeJobMetricDataWithOptions(tmpReq *DescribeJobMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMetricDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业性能数据
//
// @param request - DescribeJobMetricDataRequest
//
// @return DescribeJobMetricDataResponse
func (client *Client) DescribeJobMetricData(request *DescribeJobMetricDataRequest) (_result *DescribeJobMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJobMetricDataResponse{}
	_body, _err := client.DescribeJobMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业即时监控项
//
// @param tmpReq - DescribeJobMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricLastResponse
func (client *Client) DescribeJobMetricLastWithOptions(tmpReq *DescribeJobMetricLastRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobMetricLastResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业即时监控项
//
// @param request - DescribeJobMetricLastRequest
//
// @return DescribeJobMetricLastResponse
func (client *Client) DescribeJobMetricLast(request *DescribeJobMetricLastRequest) (_result *DescribeJobMetricLastResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeJobMetricLastResponse{}
	_body, _err := client.DescribeJobMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询执行计划详情
//
// @param request - GetActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActionPlanResponse
func (client *Client) GetActionPlanWithOptions(request *GetActionPlanRequest, runtime *dara.RuntimeOptions) (_result *GetActionPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询执行计划详情
//
// @param request - GetActionPlanRequest
//
// @return GetActionPlanResponse
func (client *Client) GetActionPlan(request *GetActionPlanRequest) (_result *GetActionPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetActionPlanResponse{}
	_body, _err := client.GetActionPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看应用版本列表
//
// @param request - GetAppVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppVersionsResponse
func (client *Client) GetAppVersionsWithOptions(request *GetAppVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetAppVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看应用版本列表
//
// @param request - GetAppVersionsRequest
//
// @return GetAppVersionsResponse
func (client *Client) GetAppVersions(request *GetAppVersionsRequest) (_result *GetAppVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppVersionsResponse{}
	_body, _err := client.GetAppVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询托管侧镜像详情。
//
// @param tmpReq - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(tmpReq *GetImageRequest, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询托管侧镜像详情。
//
// @param request - GetImageRequest
//
// @return GetImageResponse
func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业详情
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业详情
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询队列详细信息
//
// @param request - GetPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPoolResponse
func (client *Client) GetPoolWithOptions(request *GetPoolRequest, runtime *dara.RuntimeOptions) (_result *GetPoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询队列详细信息
//
// @param request - GetPoolRequest
//
// @return GetPoolResponse
func (client *Client) GetPool(request *GetPoolRequest) (_result *GetPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPoolResponse{}
	_body, _err := client.GetPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询执行计划的执行情况。
//
// @param request - ListActionPlanActivitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionPlanActivitiesResponse
func (client *Client) ListActionPlanActivitiesWithOptions(request *ListActionPlanActivitiesRequest, runtime *dara.RuntimeOptions) (_result *ListActionPlanActivitiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询执行计划的执行情况。
//
// @param request - ListActionPlanActivitiesRequest
//
// @return ListActionPlanActivitiesResponse
func (client *Client) ListActionPlanActivities(request *ListActionPlanActivitiesRequest) (_result *ListActionPlanActivitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListActionPlanActivitiesResponse{}
	_body, _err := client.ListActionPlanActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询执行计划列表
//
// @param tmpReq - ListActionPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionPlansResponse
func (client *Client) ListActionPlansWithOptions(tmpReq *ListActionPlansRequest, runtime *dara.RuntimeOptions) (_result *ListActionPlansResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询执行计划列表
//
// @param request - ListActionPlansRequest
//
// @return ListActionPlansResponse
func (client *Client) ListActionPlans(request *ListActionPlansRequest) (_result *ListActionPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListActionPlansResponse{}
	_body, _err := client.ListActionPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询全局Executor信息
//
// @param tmpReq - ListExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorsResponse
func (client *Client) ListExecutorsWithOptions(tmpReq *ListExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全局Executor信息
//
// @param request - ListExecutorsRequest
//
// @return ListExecutorsResponse
func (client *Client) ListExecutors(request *ListExecutorsRequest) (_result *ListExecutorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExecutorsResponse{}
	_body, _err := client.ListExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看托管侧镜像列表
//
// @param tmpReq - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(tmpReq *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看托管侧镜像列表
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业Executor信息
//
// @param request - ListJobExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutorsResponse
func (client *Client) ListJobExecutorsWithOptions(request *ListJobExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListJobExecutorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业Executor信息
//
// @param request - ListJobExecutorsRequest
//
// @return ListJobExecutorsResponse
func (client *Client) ListJobExecutors(request *ListJobExecutorsRequest) (_result *ListJobExecutorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobExecutorsResponse{}
	_body, _err := client.ListJobExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业列表
//
// @param tmpReq - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业列表
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资源池列表
//
// @param tmpReq - ListPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoolsResponse
func (client *Client) ListPoolsWithOptions(tmpReq *ListPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListPoolsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源池列表
//
// @param request - ListPoolsRequest
//
// @return ListPoolsResponse
func (client *Client) ListPools(request *ListPoolsRequest) (_result *ListPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPoolsResponse{}
	_body, _err := client.ListPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个或多个资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个或多个资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除托管侧镜像信息。
//
// @param request - RemoveImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageResponse
func (client *Client) RemoveImageWithOptions(request *RemoveImageRequest, runtime *dara.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除托管侧镜像信息。
//
// @param request - RemoveImageRequest
//
// @return RemoveImageResponse
func (client *Client) RemoveImage(request *RemoveImageRequest) (_result *RemoveImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveImageResponse{}
	_body, _err := client.RemoveImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用跨地域同步
//
// @param tmpReq - SynchronizeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeAppResponse
func (client *Client) SynchronizeAppWithOptions(tmpReq *SynchronizeAppRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用跨地域同步
//
// @param request - SynchronizeAppRequest
//
// @return SynchronizeAppResponse
func (client *Client) SynchronizeApp(request *SynchronizeAppRequest) (_result *SynchronizeAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SynchronizeAppResponse{}
	_body, _err := client.SynchronizeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定的资源列表统一创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定的资源列表统一创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定的ECS资源列表统一解绑标签
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定的ECS资源列表统一解绑标签
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新执行计划
//
// @param request - UpdateActionPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateActionPlanResponse
func (client *Client) UpdateActionPlanWithOptions(request *UpdateActionPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateActionPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新执行计划
//
// @param request - UpdateActionPlanRequest
//
// @return UpdateActionPlanResponse
func (client *Client) UpdateActionPlan(request *UpdateActionPlanRequest) (_result *UpdateActionPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateActionPlanResponse{}
	_body, _err := client.UpdateActionPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源池
//
// @param tmpReq - UpdatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePoolResponse
func (client *Client) UpdatePoolWithOptions(tmpReq *UpdatePoolRequest, runtime *dara.RuntimeOptions) (_result *UpdatePoolResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源池
//
// @param request - UpdatePoolRequest
//
// @return UpdatePoolResponse
func (client *Client) UpdatePool(request *UpdatePoolRequest) (_result *UpdatePoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePoolResponse{}
	_body, _err := client.UpdatePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
