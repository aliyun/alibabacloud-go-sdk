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
	EnableValidate  *bool
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
	client.Endpoint, _err = client.GetEndpoint(dara.String("schedulerx3"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建应用
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateJobShrinkRequest{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 删除应用分组
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteJobsWithOptions(tmpReq *DeleteJobsRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 批量导出任务信息
//
// @param tmpReq - ExportJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportJobsResponse
func (client *Client) ExportJobsWithOptions(tmpReq *ExportJobsRequest, runtime *dara.RuntimeOptions) (_result *ExportJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExportJobsRequest
//
// @return ExportJobsResponse
func (client *Client) ExportJobs(request *ExportJobsRequest) (_result *ExportJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportJobsResponse{}
	_body, _err := client.ExportJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *dara.RuntimeOptions) (_result *GetAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAppResponse
func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDesigateInfoWithOptions(request *GetDesigateInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDesigateInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDesigateInfoResponse
func (client *Client) GetDesigateInfo(request *GetDesigateInfoRequest) (_result *GetDesigateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDesigateInfoResponse{}
	_body, _err := client.GetDesigateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobExecutionWithOptions(request *GetJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetJobExecutionResponse
func (client *Client) GetJobExecution(request *GetJobExecutionRequest) (_result *GetJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobExecutionResponse{}
	_body, _err := client.GetJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobExecutionProgressWithOptions(request *GetJobExecutionProgressRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionProgressResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetJobExecutionProgressResponse
func (client *Client) GetJobExecutionProgress(request *GetJobExecutionProgressRequest) (_result *GetJobExecutionProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobExecutionProgressResponse{}
	_body, _err := client.GetJobExecutionProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobExecutionThreadDumpWithOptions(request *GetJobExecutionThreadDumpRequest, runtime *dara.RuntimeOptions) (_result *GetJobExecutionThreadDumpResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetJobExecutionThreadDumpResponse
func (client *Client) GetJobExecutionThreadDump(request *GetJobExecutionThreadDumpRequest) (_result *GetJobExecutionThreadDumpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobExecutionThreadDumpResponse{}
	_body, _err := client.GetJobExecutionThreadDumpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetLogWithOptions(request *GetLogRequest, runtime *dara.RuntimeOptions) (_result *GetLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetLogResponse
func (client *Client) GetLog(request *GetLogRequest) (_result *GetLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLogResponse{}
	_body, _err := client.GetLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetLogEventWithOptions(request *GetLogEventRequest, runtime *dara.RuntimeOptions) (_result *GetLogEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetLogEventResponse
func (client *Client) GetLogEvent(request *GetLogEventRequest) (_result *GetLogEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLogEventResponse{}
	_body, _err := client.GetLogEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImportCalendarWithOptions(request *ImportCalendarRequest, runtime *dara.RuntimeOptions) (_result *ImportCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImportCalendarResponse
func (client *Client) ImportCalendar(request *ImportCalendarRequest) (_result *ImportCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportCalendarResponse{}
	_body, _err := client.ImportCalendarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ImportJobsWithOptions(request *ImportJobsRequest, runtime *dara.RuntimeOptions) (_result *ImportJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImportJobsResponse
func (client *Client) ImportJobs(request *ImportJobsRequest) (_result *ImportJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportJobsResponse{}
	_body, _err := client.ImportJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAlarmEventWithOptions(request *ListAlarmEventRequest, runtime *dara.RuntimeOptions) (_result *ListAlarmEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAlarmEventResponse
func (client *Client) ListAlarmEvent(request *ListAlarmEventRequest) (_result *ListAlarmEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlarmEventResponse{}
	_body, _err := client.ListAlarmEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAppNamesWithOptions(request *ListAppNamesRequest, runtime *dara.RuntimeOptions) (_result *ListAppNamesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAppNamesResponse
func (client *Client) ListAppNames(request *ListAppNamesRequest) (_result *ListAppNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppNamesResponse{}
	_body, _err := client.ListAppNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *dara.RuntimeOptions) (_result *ListAppsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCalendarNamesWithOptions(request *ListCalendarNamesRequest, runtime *dara.RuntimeOptions) (_result *ListCalendarNamesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCalendarNamesResponse
func (client *Client) ListCalendarNames(request *ListCalendarNamesRequest) (_result *ListCalendarNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCalendarNamesResponse{}
	_body, _err := client.ListCalendarNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListExecutorsWithOptions(request *ListExecutorsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取任务实例列表
//
// @param request - ListJobExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutionsResponse
func (client *Client) ListJobExecutionsWithOptions(request *ListJobExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListJobExecutionsResponse, _err error) {
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
		Action:      dara.String("ListJobExecutions"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListJobExecutionsResponse
func (client *Client) ListJobExecutions(request *ListJobExecutionsRequest) (_result *ListJobExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobExecutionsResponse{}
	_body, _err := client.ListJobExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListJobScriptHistoryWithOptions(request *ListJobScriptHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListJobScriptHistoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListJobScriptHistoryResponse
func (client *Client) ListJobScriptHistory(request *ListJobScriptHistoryRequest) (_result *ListJobScriptHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobScriptHistoryResponse{}
	_body, _err := client.ListJobScriptHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
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
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// 获取任务列表
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
// 获取executor的label列表
//
// @param request - ListLablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLablesResponse
func (client *Client) ListLablesWithOptions(request *ListLablesRequest, runtime *dara.RuntimeOptions) (_result *ListLablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListLablesResponse
func (client *Client) ListLables(request *ListLablesRequest) (_result *ListLablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLablesResponse{}
	_body, _err := client.ListLablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可用区列表
//
// @param request - ListRegionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionZoneResponse
func (client *Client) ListRegionZoneWithOptions(runtime *dara.RuntimeOptions) (_result *ListRegionZoneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegionZone"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可用区列表
//
// @return ListRegionZoneResponse
func (client *Client) ListRegionZone() (_result *ListRegionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionZoneResponse{}
	_body, _err := client.ListRegionZoneWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取所有region列表
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2024-06-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有region列表
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListScheduleEventWithOptions(request *ListScheduleEventRequest, runtime *dara.RuntimeOptions) (_result *ListScheduleEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListScheduleEventResponse
func (client *Client) ListScheduleEvent(request *ListScheduleEventRequest) (_result *ListScheduleEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScheduleEventResponse{}
	_body, _err := client.ListScheduleEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListScheduleTimesWithOptions(request *ListScheduleTimesRequest, runtime *dara.RuntimeOptions) (_result *ListScheduleTimesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListScheduleTimesResponse
func (client *Client) ListScheduleTimes(request *ListScheduleTimesRequest) (_result *ListScheduleTimesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScheduleTimesResponse{}
	_body, _err := client.ListScheduleTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateDesignateExecutorsWithOptions(tmpReq *OperateDesignateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *OperateDesignateExecutorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateDesignateExecutorsRequest
//
// @return OperateDesignateExecutorsResponse
func (client *Client) OperateDesignateExecutors(request *OperateDesignateExecutorsRequest) (_result *OperateDesignateExecutorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateDesignateExecutorsResponse{}
	_body, _err := client.OperateDesignateExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateDisableJobsWithOptions(tmpReq *OperateDisableJobsRequest, runtime *dara.RuntimeOptions) (_result *OperateDisableJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateDisableJobsRequest
//
// @return OperateDisableJobsResponse
func (client *Client) OperateDisableJobs(request *OperateDisableJobsRequest) (_result *OperateDisableJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateDisableJobsResponse{}
	_body, _err := client.OperateDisableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateEnableJobsWithOptions(tmpReq *OperateEnableJobsRequest, runtime *dara.RuntimeOptions) (_result *OperateEnableJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateEnableJobsRequest
//
// @return OperateEnableJobsResponse
func (client *Client) OperateEnableJobs(request *OperateEnableJobsRequest) (_result *OperateEnableJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateEnableJobsResponse{}
	_body, _err := client.OperateEnableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateExecuteJobWithOptions(request *OperateExecuteJobRequest, runtime *dara.RuntimeOptions) (_result *OperateExecuteJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateExecuteJobResponse
func (client *Client) OperateExecuteJob(request *OperateExecuteJobRequest) (_result *OperateExecuteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateExecuteJobResponse{}
	_body, _err := client.OperateExecuteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateRerunJobWithOptions(request *OperateRerunJobRequest, runtime *dara.RuntimeOptions) (_result *OperateRerunJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateRerunJobResponse
func (client *Client) OperateRerunJob(request *OperateRerunJobRequest) (_result *OperateRerunJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateRerunJobResponse{}
	_body, _err := client.OperateRerunJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateRetryJobExecutionWithOptions(tmpReq *OperateRetryJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateRetryJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateRetryJobExecutionRequest
//
// @return OperateRetryJobExecutionResponse
func (client *Client) OperateRetryJobExecution(request *OperateRetryJobExecutionRequest) (_result *OperateRetryJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateRetryJobExecutionResponse{}
	_body, _err := client.OperateRetryJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止正在运行的任务实例
//
// @param tmpReq - OperateStopJobExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateStopJobExecutionResponse
func (client *Client) OperateStopJobExecutionWithOptions(tmpReq *OperateStopJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateStopJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止正在运行的任务实例
//
// @param request - OperateStopJobExecutionRequest
//
// @return OperateStopJobExecutionResponse
func (client *Client) OperateStopJobExecution(request *OperateStopJobExecutionRequest) (_result *OperateStopJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateStopJobExecutionResponse{}
	_body, _err := client.OperateStopJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAppResponse
func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateClusterWithOptions(request *UpdateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateClusterResponse
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateJobWithOptions(tmpReq *UpdateJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
func (client *Client) UpdateJob(request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateJobScriptWithOptions(request *UpdateJobScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobScriptResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateJobScriptResponse
func (client *Client) UpdateJobScript(request *UpdateJobScriptRequest) (_result *UpdateJobScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateJobScriptResponse{}
	_body, _err := client.UpdateJobScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
