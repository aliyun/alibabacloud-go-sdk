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
// 创建日历
//
// @param request - CreateCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCalendarResponse
func (client *Client) CreateCalendarWithOptions(request *CreateCalendarRequest, runtime *dara.RuntimeOptions) (_result *CreateCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCalendarResponse
func (client *Client) CreateCalendar(request *CreateCalendarRequest) (_result *CreateCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCalendarResponse{}
	_body, _err := client.CreateCalendarWithOptions(request, runtime)
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
// 添加执行器
//
// @param request - CreateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExecutorsResponse
func (client *Client) CreateExecutorsWithOptions(request *CreateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *CreateExecutorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateExecutorsResponse
func (client *Client) CreateExecutors(request *CreateExecutorsRequest) (_result *CreateExecutorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExecutorsResponse{}
	_body, _err := client.CreateExecutorsWithOptions(request, runtime)
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
// 创建应用
//
// @param request - CreateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithOptions(request *CreateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
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
// @param request - CreateWorkflowRequest
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflow(request *CreateWorkflowRequest) (_result *CreateWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CreateWorkflowWithOptions(request, runtime)
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
// 删除日历
//
// @param request - DeleteCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCalendarResponse
func (client *Client) DeleteCalendarWithOptions(request *DeleteCalendarRequest, runtime *dara.RuntimeOptions) (_result *DeleteCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCalendarResponse
func (client *Client) DeleteCalendar(request *DeleteCalendarRequest) (_result *DeleteCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCalendarResponse{}
	_body, _err := client.DeleteCalendarWithOptions(request, runtime)
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
// 删除工作流
//
// @param request - DeleteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithOptions(request *DeleteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (_result *DeleteWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.DeleteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteWorkflowsWithOptions(tmpReq *DeleteWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteWorkflowsRequest
//
// @return DeleteWorkflowsResponse
func (client *Client) DeleteWorkflows(request *DeleteWorkflowsRequest) (_result *DeleteWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkflowsResponse{}
	_body, _err := client.DeleteWorkflowsWithOptions(request, runtime)
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
// 批量导出工作流信息
//
// @param tmpReq - ExportWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportWorkflowsResponse
func (client *Client) ExportWorkflowsWithOptions(tmpReq *ExportWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ExportWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExportWorkflowsRequest
//
// @return ExportWorkflowsResponse
func (client *Client) ExportWorkflows(request *ExportWorkflowsRequest) (_result *ExportWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportWorkflowsResponse{}
	_body, _err := client.ExportWorkflowsWithOptions(request, runtime)
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
// 获取日历信息
//
// @param request - GetCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalendarResponse
func (client *Client) GetCalendarWithOptions(request *GetCalendarRequest, runtime *dara.RuntimeOptions) (_result *GetCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCalendarResponse
func (client *Client) GetCalendar(request *GetCalendarRequest) (_result *GetCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCalendarResponse{}
	_body, _err := client.GetCalendarWithOptions(request, runtime)
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
// 查询执行器配置信息
//
// @param request - GetExecutorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutorConfigResponse
func (client *Client) GetExecutorConfigWithOptions(request *GetExecutorConfigRequest, runtime *dara.RuntimeOptions) (_result *GetExecutorConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetExecutorConfigResponse
func (client *Client) GetExecutorConfig(request *GetExecutorConfigRequest) (_result *GetExecutorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExecutorConfigResponse{}
	_body, _err := client.GetExecutorConfigWithOptions(request, runtime)
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
// 获取工作流
//
// @param request - GetWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowResponse
func (client *Client) GetWorkflowWithOptions(request *GetWorkflowRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkflowResponse
func (client *Client) GetWorkflow(request *GetWorkflowRequest) (_result *GetWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowResponse{}
	_body, _err := client.GetWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkflowDAGWithOptions(request *GetWorkflowDAGRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDAGResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkflowDAGResponse
func (client *Client) GetWorkflowDAG(request *GetWorkflowDAGRequest) (_result *GetWorkflowDAGResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowDAGResponse{}
	_body, _err := client.GetWorkflowDAGWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkflowDAGPreviewWithOptions(request *GetWorkflowDAGPreviewRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDAGPreviewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkflowDAGPreviewResponse
func (client *Client) GetWorkflowDAGPreview(request *GetWorkflowDAGPreviewRequest) (_result *GetWorkflowDAGPreviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowDAGPreviewResponse{}
	_body, _err := client.GetWorkflowDAGPreviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkflowExecutionDAGWithOptions(request *GetWorkflowExecutionDAGRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowExecutionDAGResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkflowExecutionDAGResponse
func (client *Client) GetWorkflowExecutionDAG(request *GetWorkflowExecutionDAGRequest) (_result *GetWorkflowExecutionDAGResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowExecutionDAGResponse{}
	_body, _err := client.GetWorkflowExecutionDAGWithOptions(request, runtime)
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
// 批量导入工作流
//
// @param request - ImportWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportWorkflowsResponse
func (client *Client) ImportWorkflowsWithOptions(request *ImportWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ImportWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ImportWorkflowsResponse
func (client *Client) ImportWorkflows(request *ImportWorkflowsRequest) (_result *ImportWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportWorkflowsResponse{}
	_body, _err := client.ImportWorkflowsWithOptions(request, runtime)
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
// 查询日历
//
// @param request - ListCalendarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalendarsResponse
func (client *Client) ListCalendarsWithOptions(request *ListCalendarsRequest, runtime *dara.RuntimeOptions) (_result *ListCalendarsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCalendarsResponse
func (client *Client) ListCalendars(request *ListCalendarsRequest) (_result *ListCalendarsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCalendarsResponse{}
	_body, _err := client.ListCalendarsWithOptions(request, runtime)
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
// 获取k8s资源列表
//
// @param request - ListK8sResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sResourceResponse
func (client *Client) ListK8sResourceWithOptions(request *ListK8sResourceRequest, runtime *dara.RuntimeOptions) (_result *ListK8sResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListK8sResourceResponse
func (client *Client) ListK8sResource(request *ListK8sResourceRequest) (_result *ListK8sResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListK8sResourceResponse{}
	_body, _err := client.ListK8sResourceWithOptions(request, runtime)
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
// 获取流程实例列表
//
// @param request - ListWorkflowExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowExecutionsResponse
func (client *Client) ListWorkflowExecutionsWithOptions(request *ListWorkflowExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowExecutionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListWorkflowExecutionsResponse
func (client *Client) ListWorkflowExecutions(request *ListWorkflowExecutionsRequest) (_result *ListWorkflowExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowExecutionsResponse{}
	_body, _err := client.ListWorkflowExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListWorkflowVersionsWithOptions(request *ListWorkflowVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListWorkflowVersionsResponse
func (client *Client) ListWorkflowVersions(request *ListWorkflowVersionsRequest) (_result *ListWorkflowVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowVersionsResponse{}
	_body, _err := client.ListWorkflowVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListWorkflowsWithOptions(request *ListWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListWorkflowsResponse
func (client *Client) ListWorkflows(request *ListWorkflowsRequest) (_result *ListWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.ListWorkflowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateBackfillWorkflowWithOptions(request *OperateBackfillWorkflowRequest, runtime *dara.RuntimeOptions) (_result *OperateBackfillWorkflowResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateBackfillWorkflowResponse
func (client *Client) OperateBackfillWorkflow(request *OperateBackfillWorkflowRequest) (_result *OperateBackfillWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateBackfillWorkflowResponse{}
	_body, _err := client.OperateBackfillWorkflowWithOptions(request, runtime)
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
// 批量禁用工作流
//
// @param tmpReq - OperateDisableWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateDisableWorkflowsResponse
func (client *Client) OperateDisableWorkflowsWithOptions(tmpReq *OperateDisableWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *OperateDisableWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateDisableWorkflowsRequest
//
// @return OperateDisableWorkflowsResponse
func (client *Client) OperateDisableWorkflows(request *OperateDisableWorkflowsRequest) (_result *OperateDisableWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateDisableWorkflowsResponse{}
	_body, _err := client.OperateDisableWorkflowsWithOptions(request, runtime)
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
// 批量启用工作流
//
// @param tmpReq - OperateEnableWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateEnableWorkflowsResponse
func (client *Client) OperateEnableWorkflowsWithOptions(tmpReq *OperateEnableWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *OperateEnableWorkflowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OperateEnableWorkflowsRequest
//
// @return OperateEnableWorkflowsResponse
func (client *Client) OperateEnableWorkflows(request *OperateEnableWorkflowsRequest) (_result *OperateEnableWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateEnableWorkflowsResponse{}
	_body, _err := client.OperateEnableWorkflowsWithOptions(request, runtime)
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
// 运行一次工作流
//
// @param request - OperateExecuteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateExecuteWorkflowResponse
func (client *Client) OperateExecuteWorkflowWithOptions(request *OperateExecuteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *OperateExecuteWorkflowResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateExecuteWorkflowResponse
func (client *Client) OperateExecuteWorkflow(request *OperateExecuteWorkflowRequest) (_result *OperateExecuteWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateExecuteWorkflowResponse{}
	_body, _err := client.OperateExecuteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateHoldJobExecutionWithOptions(request *OperateHoldJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateHoldJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateHoldJobExecutionResponse
func (client *Client) OperateHoldJobExecution(request *OperateHoldJobExecutionRequest) (_result *OperateHoldJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateHoldJobExecutionResponse{}
	_body, _err := client.OperateHoldJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateHoldWorkflowExecutionWithOptions(request *OperateHoldWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateHoldWorkflowExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateHoldWorkflowExecutionResponse
func (client *Client) OperateHoldWorkflowExecution(request *OperateHoldWorkflowExecutionRequest) (_result *OperateHoldWorkflowExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateHoldWorkflowExecutionResponse{}
	_body, _err := client.OperateHoldWorkflowExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateMarkSuccessJobExecutionWithOptions(request *OperateMarkSuccessJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateMarkSuccessJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateMarkSuccessJobExecutionResponse
func (client *Client) OperateMarkSuccessJobExecution(request *OperateMarkSuccessJobExecutionRequest) (_result *OperateMarkSuccessJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateMarkSuccessJobExecutionResponse{}
	_body, _err := client.OperateMarkSuccessJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateMarkSuccessWorkflowExecutionWithOptions(request *OperateMarkSuccessWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateMarkSuccessWorkflowExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateMarkSuccessWorkflowExecutionResponse
func (client *Client) OperateMarkSuccessWorkflowExecution(request *OperateMarkSuccessWorkflowExecutionRequest) (_result *OperateMarkSuccessWorkflowExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateMarkSuccessWorkflowExecutionResponse{}
	_body, _err := client.OperateMarkSuccessWorkflowExecutionWithOptions(request, runtime)
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
// 重跑工作流实例
//
// @param request - OperateRetryWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateRetryWorkflowExecutionResponse
func (client *Client) OperateRetryWorkflowExecutionWithOptions(request *OperateRetryWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateRetryWorkflowExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateRetryWorkflowExecutionResponse
func (client *Client) OperateRetryWorkflowExecution(request *OperateRetryWorkflowExecutionRequest) (_result *OperateRetryWorkflowExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateRetryWorkflowExecutionResponse{}
	_body, _err := client.OperateRetryWorkflowExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateSkipJobExecutionWithOptions(request *OperateSkipJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateSkipJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateSkipJobExecutionResponse
func (client *Client) OperateSkipJobExecution(request *OperateSkipJobExecutionRequest) (_result *OperateSkipJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateSkipJobExecutionResponse{}
	_body, _err := client.OperateSkipJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// 停止运行中的任务实例
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
// 停止正在运行的工作流实例
//
// @param request - OperateStopWorkflowExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateStopWorkflowExecutionResponse
func (client *Client) OperateStopWorkflowExecutionWithOptions(request *OperateStopWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateStopWorkflowExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateStopWorkflowExecutionResponse
func (client *Client) OperateStopWorkflowExecution(request *OperateStopWorkflowExecutionRequest) (_result *OperateStopWorkflowExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateStopWorkflowExecutionResponse{}
	_body, _err := client.OperateStopWorkflowExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateUnholdJobExecutionWithOptions(request *OperateUnholdJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnholdJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateUnholdJobExecutionResponse
func (client *Client) OperateUnholdJobExecution(request *OperateUnholdJobExecutionRequest) (_result *OperateUnholdJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateUnholdJobExecutionResponse{}
	_body, _err := client.OperateUnholdJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateUnholdWorkflowExecutionWithOptions(request *OperateUnholdWorkflowExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnholdWorkflowExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateUnholdWorkflowExecutionResponse
func (client *Client) OperateUnholdWorkflowExecution(request *OperateUnholdWorkflowExecutionRequest) (_result *OperateUnholdWorkflowExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateUnholdWorkflowExecutionResponse{}
	_body, _err := client.OperateUnholdWorkflowExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateUnskipJobExecutionWithOptions(request *OperateUnskipJobExecutionRequest, runtime *dara.RuntimeOptions) (_result *OperateUnskipJobExecutionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateUnskipJobExecutionResponse
func (client *Client) OperateUnskipJobExecution(request *OperateUnskipJobExecutionRequest) (_result *OperateUnskipJobExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateUnskipJobExecutionResponse{}
	_body, _err := client.OperateUnskipJobExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SyncJobsWithOptions(tmpReq *SyncJobsRequest, runtime *dara.RuntimeOptions) (_result *SyncJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SyncJobsRequest
//
// @return SyncJobsResponse
func (client *Client) SyncJobs(request *SyncJobsRequest) (_result *SyncJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncJobsResponse{}
	_body, _err := client.SyncJobsWithOptions(request, runtime)
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
// 更新日历
//
// @param request - UpdateCalendarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCalendarResponse
func (client *Client) UpdateCalendarWithOptions(request *UpdateCalendarRequest, runtime *dara.RuntimeOptions) (_result *UpdateCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCalendarResponse
func (client *Client) UpdateCalendar(request *UpdateCalendarRequest) (_result *UpdateCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCalendarResponse{}
	_body, _err := client.UpdateCalendarWithOptions(request, runtime)
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
// 更新执行器
//
// @param request - UpdateExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExecutorsResponse
func (client *Client) UpdateExecutorsWithOptions(request *UpdateExecutorsRequest, runtime *dara.RuntimeOptions) (_result *UpdateExecutorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateExecutorsResponse
func (client *Client) UpdateExecutors(request *UpdateExecutorsRequest) (_result *UpdateExecutorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExecutorsResponse{}
	_body, _err := client.UpdateExecutorsWithOptions(request, runtime)
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

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithOptions(request *UpdateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (_result *UpdateWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.UpdateWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateWorkflowDAGWithOptions(tmpReq *UpdateWorkflowDAGRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDAGResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateWorkflowDAGRequest
//
// @return UpdateWorkflowDAGResponse
func (client *Client) UpdateWorkflowDAG(request *UpdateWorkflowDAGRequest) (_result *UpdateWorkflowDAGResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowDAGResponse{}
	_body, _err := client.UpdateWorkflowDAGWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateWorkflowDAGVersionWithOptions(request *UpdateWorkflowDAGVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDAGVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateWorkflowDAGVersionResponse
func (client *Client) UpdateWorkflowDAGVersion(request *UpdateWorkflowDAGVersionRequest) (_result *UpdateWorkflowDAGVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowDAGVersionResponse{}
	_body, _err := client.UpdateWorkflowDAGVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
