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
	client.Endpoint, _err = client.GetEndpoint(dara.String("alikafkastreaming"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检查sql语法
//
// @param request - CheckSqlContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSqlContentResponse
func (client *Client) CheckSqlContentWithOptions(request *CheckSqlContentRequest, runtime *dara.RuntimeOptions) (_result *CheckSqlContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SqlContent) {
		query["SqlContent"] = request.SqlContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSqlContent"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSqlContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查sql语法
//
// @param request - CheckSqlContentRequest
//
// @return CheckSqlContentResponse
func (client *Client) CheckSqlContent(request *CheckSqlContentRequest) (_result *CheckSqlContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSqlContentResponse{}
	_body, _err := client.CheckSqlContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 流计算实例
//
// @param request - CreateComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeInstanceResponse
func (client *Client) CreateComputeInstanceWithOptions(request *CreateComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 流计算实例
//
// @param request - CreateComputeInstanceRequest
//
// @return CreateComputeInstanceResponse
func (client *Client) CreateComputeInstance(request *CreateComputeInstanceRequest) (_result *CreateComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComputeInstanceResponse{}
	_body, _err := client.CreateComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 JOB
//
// @param request - CreateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeJobResponse
func (client *Client) CreateComputeJobWithOptions(request *CreateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.DraftSql) {
		query["DraftSql"] = request.DraftSql
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobConfig) {
		query["JobConfig"] = request.JobConfig
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UpgradeMode) {
		query["UpgradeMode"] = request.UpgradeMode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 JOB
//
// @param request - CreateComputeJobRequest
//
// @return CreateComputeJobResponse
func (client *Client) CreateComputeJob(request *CreateComputeJobRequest) (_result *CreateComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComputeJobResponse{}
	_body, _err := client.CreateComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeInstanceResponse
func (client *Client) DeleteComputeInstanceWithOptions(request *DeleteComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteComputeInstanceRequest
//
// @return DeleteComputeInstanceResponse
func (client *Client) DeleteComputeInstance(request *DeleteComputeInstanceRequest) (_result *DeleteComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComputeInstanceResponse{}
	_body, _err := client.DeleteComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 JOB
//
// @param request - DeleteComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeJobResponse
func (client *Client) DeleteComputeJobWithOptions(request *DeleteComputeJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 JOB
//
// @param request - DeleteComputeJobRequest
//
// @return DeleteComputeJobResponse
func (client *Client) DeleteComputeJob(request *DeleteComputeJobRequest) (_result *DeleteComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComputeJobResponse{}
	_body, _err := client.DeleteComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个实例
//
// @param request - GetComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeInstanceResponse
func (client *Client) GetComputeInstanceWithOptions(request *GetComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个实例
//
// @param request - GetComputeInstanceRequest
//
// @return GetComputeInstanceResponse
func (client *Client) GetComputeInstance(request *GetComputeInstanceRequest) (_result *GetComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComputeInstanceResponse{}
	_body, _err := client.GetComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 JOB 详情
//
// @param request - GetComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeJobResponse
func (client *Client) GetComputeJobWithOptions(request *GetComputeJobRequest, runtime *dara.RuntimeOptions) (_result *GetComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 JOB 详情
//
// @param request - GetComputeJobRequest
//
// @return GetComputeJobResponse
func (client *Client) GetComputeJob(request *GetComputeJobRequest) (_result *GetComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComputeJobResponse{}
	_body, _err := client.GetComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取debug信息
//
// @param request - GetJobDebugDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDebugDataResponse
func (client *Client) GetJobDebugDataWithOptions(request *GetJobDebugDataRequest, runtime *dara.RuntimeOptions) (_result *GetJobDebugDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDebugData"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDebugDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取debug信息
//
// @param request - GetJobDebugDataRequest
//
// @return GetJobDebugDataResponse
func (client *Client) GetJobDebugData(request *GetJobDebugDataRequest) (_result *GetJobDebugDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobDebugDataResponse{}
	_body, _err := client.GetJobDebugDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param tmpReq - ListComputeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesResponse
func (client *Client) ListComputeInstancesWithOptions(tmpReq *ListComputeInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeInstances"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param request - ListComputeInstancesRequest
//
// @return ListComputeInstancesResponse
func (client *Client) ListComputeInstances(request *ListComputeInstancesRequest) (_result *ListComputeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeInstancesResponse{}
	_body, _err := client.ListComputeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param tmpReq - ListComputeInstancesInPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesInPageResponse
func (client *Client) ListComputeInstancesInPageWithOptions(tmpReq *ListComputeInstancesInPageRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesInPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeInstancesInPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeInstancesInPage"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeInstancesInPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param request - ListComputeInstancesInPageRequest
//
// @return ListComputeInstancesInPageResponse
func (client *Client) ListComputeInstancesInPage(request *ListComputeInstancesInPageRequest) (_result *ListComputeInstancesInPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeInstancesInPageResponse{}
	_body, _err := client.ListComputeInstancesInPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询 JOB 列表
//
// @param request - ListComputeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeJobsResponse
func (client *Client) ListComputeJobsWithOptions(request *ListComputeJobsRequest, runtime *dara.RuntimeOptions) (_result *ListComputeJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeJobs"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询 JOB 列表
//
// @param request - ListComputeJobsRequest
//
// @return ListComputeJobsResponse
func (client *Client) ListComputeJobs(request *ListComputeJobsRequest) (_result *ListComputeJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeJobsResponse{}
	_body, _err := client.ListComputeJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询创建 SQL 任务时支持的连接器列表
//
// @param request - ListSupportedConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedConnectorsResponse
func (client *Client) ListSupportedConnectorsWithOptions(request *ListSupportedConnectorsRequest, runtime *dara.RuntimeOptions) (_result *ListSupportedConnectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupportedConnectors"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupportedConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询创建 SQL 任务时支持的连接器列表
//
// @param request - ListSupportedConnectorsRequest
//
// @return ListSupportedConnectorsResponse
func (client *Client) ListSupportedConnectors(request *ListSupportedConnectorsRequest) (_result *ListSupportedConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupportedConnectorsResponse{}
	_body, _err := client.ListSupportedConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新启动后付费实例
//
// @param request - ReopenComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenComputeInstanceResponse
func (client *Client) ReopenComputeInstanceWithOptions(request *ReopenComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReopenComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新启动后付费实例
//
// @param request - ReopenComputeInstanceRequest
//
// @return ReopenComputeInstanceResponse
func (client *Client) ReopenComputeInstance(request *ReopenComputeInstanceRequest) (_result *ReopenComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReopenComputeInstanceResponse{}
	_body, _err := client.ReopenComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启 JOB
//
// @param request - RestartComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartComputeJobResponse
func (client *Client) RestartComputeJobWithOptions(request *RestartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *RestartComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启 JOB
//
// @param request - RestartComputeJobRequest
//
// @return RestartComputeJobResponse
func (client *Client) RestartComputeJob(request *RestartComputeJobRequest) (_result *RestartComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartComputeJobResponse{}
	_body, _err := client.RestartComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 部署实例
//
// @param tmpReq - StartComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeInstanceResponse
func (client *Client) StartComputeInstanceWithOptions(tmpReq *StartComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartComputeInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SelectedZones) {
		query["SelectedZones"] = request.SelectedZones
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		query["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 部署实例
//
// @param request - StartComputeInstanceRequest
//
// @return StartComputeInstanceResponse
func (client *Client) StartComputeInstance(request *StartComputeInstanceRequest) (_result *StartComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartComputeInstanceResponse{}
	_body, _err := client.StartComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 JOB
//
// @param request - StartComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeJobResponse
func (client *Client) StartComputeJobWithOptions(request *StartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StartComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.DraftSql) {
		query["DraftSql"] = request.DraftSql
	}

	if !dara.IsNil(request.DraftSqlStart) {
		query["DraftSqlStart"] = request.DraftSqlStart
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RecoveryMode) {
		query["RecoveryMode"] = request.RecoveryMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 JOB
//
// @param request - StartComputeJobRequest
//
// @return StartComputeJobResponse
func (client *Client) StartComputeJob(request *StartComputeJobRequest) (_result *StartComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartComputeJobResponse{}
	_body, _err := client.StartComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停用/释放后付费实例
//
// @param request - StopComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeInstanceResponse
func (client *Client) StopComputeInstanceWithOptions(request *StopComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停用/释放后付费实例
//
// @param request - StopComputeInstanceRequest
//
// @return StopComputeInstanceResponse
func (client *Client) StopComputeInstance(request *StopComputeInstanceRequest) (_result *StopComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopComputeInstanceResponse{}
	_body, _err := client.StopComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止 JOB
//
// @param request - StopComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeJobResponse
func (client *Client) StopComputeJobWithOptions(request *StopComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StopComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止 JOB
//
// @param request - StopComputeJobRequest
//
// @return StopComputeJobResponse
func (client *Client) StopComputeJob(request *StopComputeJobRequest) (_result *StopComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopComputeJobResponse{}
	_body, _err := client.StopComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例名称
//
// @param request - UpdateComputeInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeInstanceNameResponse
func (client *Client) UpdateComputeInstanceNameWithOptions(request *UpdateComputeInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeInstanceName"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例名称
//
// @param request - UpdateComputeInstanceNameRequest
//
// @return UpdateComputeInstanceNameResponse
func (client *Client) UpdateComputeInstanceName(request *UpdateComputeInstanceNameRequest) (_result *UpdateComputeInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeInstanceNameResponse{}
	_body, _err := client.UpdateComputeInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 JOB
//
// @param request - UpdateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobResponse
func (client *Client) UpdateComputeJobWithOptions(request *UpdateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UpgradeMode) {
		query["UpgradeMode"] = request.UpgradeMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 JOB
//
// @param request - UpdateComputeJobRequest
//
// @return UpdateComputeJobResponse
func (client *Client) UpdateComputeJob(request *UpdateComputeJobRequest) (_result *UpdateComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobResponse{}
	_body, _err := client.UpdateComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 JOB 的 CU 配额
//
// @param request - UpdateComputeJobCuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobCuResponse
func (client *Client) UpdateComputeJobCuWithOptions(request *UpdateComputeJobCuRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobCuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJobCu"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobCuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 JOB 的 CU 配额
//
// @param request - UpdateComputeJobCuRequest
//
// @return UpdateComputeJobCuResponse
func (client *Client) UpdateComputeJobCu(request *UpdateComputeJobCuRequest) (_result *UpdateComputeJobCuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobCuResponse{}
	_body, _err := client.UpdateComputeJobCuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 JOB 的 SQL
//
// @param request - UpdateComputeJobDraftSqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobDraftSqlResponse
func (client *Client) UpdateComputeJobDraftSqlWithOptions(request *UpdateComputeJobDraftSqlRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobDraftSqlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftSql) {
		query["DraftSql"] = request.DraftSql
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJobDraftSql"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobDraftSqlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 JOB 的 SQL
//
// @param request - UpdateComputeJobDraftSqlRequest
//
// @return UpdateComputeJobDraftSqlResponse
func (client *Client) UpdateComputeJobDraftSql(request *UpdateComputeJobDraftSqlRequest) (_result *UpdateComputeJobDraftSqlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobDraftSqlResponse{}
	_body, _err := client.UpdateComputeJobDraftSqlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
