// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 检查sql语法
//
// Description:
//
// ## 请求说明
//
// - 该接口支持通过 GET 或 POST 方法调用。
//
// - 必须提供 `InstanceId`、`JobName` 和 `SqlContent` 参数，其中 `SqlContent` 是待校验的 Flink SQL 语句。
//
// - 返回结果中，`Data.Valid` 字段指示 SQL 是否通过校验；若未通过，则错误详情位于 `Data.ErrorList` 中。
//
// - 当前版本要求同时传入实例 ID (`InstanceId`) 和作业名称 (`JobName`) 以构建作业上下文。
//
// - 接口返回成功仅表示校验流程执行完成，并不直接反映 SQL 的有效性，请检查 `Data.Valid` 字段来确定 SQL 是否有效。
//
// - 错误码和异常处理请参考文档中的“错误码”部分。
//
// @param request - CheckSqlContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSqlContentResponse
func (client *Client) CheckSqlContentWithContext(ctx context.Context, request *CheckSqlContentRequest, runtime *dara.RuntimeOptions) (_result *CheckSqlContentResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 创建一个计算实例。接口只完成购买阶段；创建成功后需调用 StartComputeInstance 完成网络配置和部署。
//
// - API 版本：2026-02-02
//
// - Action：CreateComputeInstance
//
// @param request - CreateComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeInstanceResponse
func (client *Client) CreateComputeInstanceWithContext(ctx context.Context, request *CreateComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 该API用于在指定的运行中的计算实例上创建一个新的Flink SQL作业。
//
// - 创建后的作业将处于`INIT`状态。
//
// - 用户可以通过设置`CuLimit`和`CuReserved`来控制作业的资源使用情况。
//
// - `Remark`字段允许用户为作业添加备注信息，便于管理和识别。
//
// - 确保提供的`RegionId`、`InstanceId`以及`JobName`参数准确无误，否则可能导致请求失败。
//
// - 如果尝试创建同名作业，则会返回错误提示。
//
// - 计算实例必须处于运行状态才能成功创建作业。
//
// @param request - CreateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeJobResponse
func (client *Client) CreateComputeJobWithContext(ctx context.Context, request *CreateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeJobResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 删除处于待部署、已停止或已释放状态的计算实例。
//
// - API版本：2026-02-02
//
// - Action：DeleteComputeInstance
//
// @param request - DeleteComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeInstanceResponse
func (client *Client) DeleteComputeInstanceWithContext(ctx context.Context, request *DeleteComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 该接口用于删除一个特定的计算作业。
//
// - 成功调用此接口仅表示删除请求已被系统接受，并非立即完成删除操作。
//
// - 确保提供的`RegionId`、`InstanceId`以及`JobName`参数准确无误，否则可能导致请求失败。
//
// - 如果计算实例或作业处于不允许删除的状态（例如：非运行状态），则会返回相应的错误信息。
//
// - 删除操作不可逆，请谨慎使用。
//
// @param request - DeleteComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeJobResponse
func (client *Client) DeleteComputeJobWithContext(ctx context.Context, request *DeleteComputeJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeInstanceResponse
func (client *Client) GetComputeInstanceWithContext(ctx context.Context, request *GetComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 本接口用于查询指定计算作业的详情。
//
// - 支持使用 GET 或 POST 方法进行请求。
//
// - 所有时间字段以 Unix 时间戳形式返回，单位为毫秒。
//
// - 必须提供 `RegionId`、`InstanceId` 和 `JobName` 参数。
//
// - 授权操作为 `alikafkastreaming:GetComputeJob`，访问级别为读取（Read）。
//
// @param request - GetComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeJobResponse
func (client *Client) GetComputeJobWithContext(ctx context.Context, request *GetComputeJobRequest, runtime *dara.RuntimeOptions) (_result *GetComputeJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDebugDataResponse
func (client *Client) GetJobDebugDataWithContext(ctx context.Context, request *GetJobDebugDataRequest, runtime *dara.RuntimeOptions) (_result *GetJobDebugDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListComputeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesResponse
func (client *Client) ListComputeInstancesWithContext(ctx context.Context, tmpReq *ListComputeInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListComputeInstancesInPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesInPageResponse
func (client *Client) ListComputeInstancesInPageWithContext(ctx context.Context, tmpReq *ListComputeInstancesInPageRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesInPageResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 该接口支持通过 `MaxResults` 和 `NextToken` 参数进行游标分页查询。
//
// - 首次请求时不需要传递 `NextToken`，后续请求需使用上一次响应中返回的 `NextToken` 值。
//
// - 支持按作业名称或备注搜索，并可选择不同的排序字段和方向。
//
// - 返回的时间字段均为 Unix 时间戳（单位：毫秒）。
//
// - 授权操作为 `alikafkastreaming:ListComputeJobs`，访问级别为列出（List），适用于全部资源。
//
// @param request - ListComputeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeJobsResponse
func (client *Client) ListComputeJobsWithContext(ctx context.Context, request *ListComputeJobsRequest, runtime *dara.RuntimeOptions) (_result *ListComputeJobsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedConnectorsResponse
func (client *Client) ListSupportedConnectorsWithContext(ctx context.Context, request *ListSupportedConnectorsRequest, runtime *dara.RuntimeOptions) (_result *ListSupportedConnectorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 重新启用一个已停止的后付费计算实例。接口返回成功表示启用请求已受理。
//
// - API版本：2026-02-02
//
// - Action：ReopenComputeInstance
//
// @param request - ReopenComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenComputeInstanceResponse
func (client *Client) ReopenComputeInstanceWithContext(ctx context.Context, request *ReopenComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReopenComputeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartComputeJobResponse
func (client *Client) RestartComputeJobWithContext(ctx context.Context, request *RestartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *RestartComputeJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 为处于待部署状态的计算实例配置网络并发起部署。
//
// - API 版本：2026-02-02
//
// - Action：StartComputeInstance
//
// @param tmpReq - StartComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeInstanceResponse
func (client *Client) StartComputeInstanceWithContext(ctx context.Context, tmpReq *StartComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartComputeInstanceResponse, _err error) {
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

	if !dara.IsNil(request.VSwitchIdsShrink) {
		query["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - `RecoveryMode` 支持两种模式：`savepoint` 和 `stateless`。如果选择 `savepoint` 模式但没有可用的 savepoint，则会返回错误。
//
// - `CuLimit` 和 `CuReserved` 参数分别用来设定作业的 CU 上限和预留 CU 数量，支持整数或小数形式输入。
//
// - 确保提供的 `RegionId`, `InstanceId`, 和 `JobName` 参数值正确且存在，否则将导致请求失败。
//
// @param request - StartComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeJobResponse
func (client *Client) StartComputeJobWithContext(ctx context.Context, request *StartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StartComputeJobResponse, _err error) {
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

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RecoveryMode) {
		query["RecoveryMode"] = request.RecoveryMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 停止一个正在运行的后付费计算实例。接口返回成功表示停止请求已受理。
//
// - API 版本：2026-02-02
//
// - Action：StopComputeInstance
//
// @param request - StopComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeInstanceResponse
func (client *Client) StopComputeInstanceWithContext(ctx context.Context, request *StopComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopComputeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 该接口用于停止指定的计算作业生产或 Debug 运行实例。
//
// - 接口返回成功表示停止请求已被受理，但并不意味着作业立即停止。
//
// @param request - StopComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeJobResponse
func (client *Client) StopComputeJobWithContext(ctx context.Context, request *StopComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StopComputeJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// 修改计算实例名称。实例需处于部署准备阶段或运行中状态。
//
// - API 版本：2026-02-02
//
// - Action：UpdateComputeInstanceName
//
// @param request - UpdateComputeInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeInstanceNameResponse
func (client *Client) UpdateComputeInstanceNameWithContext(ctx context.Context, request *UpdateComputeInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeInstanceNameResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// - 确保提供的 `InstanceId` 和 `JobName` 是有效的，否则将返回错误。
//
// - 如果实例状态不在运行中，则不允许执行此操作。
//
// - 当前作业状态如果为调试任务正在运行或变更中，则不支持修改。
//
// @param request - UpdateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobResponse
func (client *Client) UpdateComputeJobWithContext(ctx context.Context, request *UpdateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// 本API允许用户修改特定计算作业的计算单元（CU）上限和预留CU数量。在调用此接口前，请确保提供的`InstanceId`和`JobName`正确无误，并且实例处于运行状态。此外，注意检查`CuLimit`与`CuReserved`参数的有效性和合理性，避免因超出限制或不符合业务逻辑导致请求失败。
//
// @param request - UpdateComputeJobCuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobCuResponse
func (client *Client) UpdateComputeJobCuWithContext(ctx context.Context, request *UpdateComputeJobCuRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobCuResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// ## 请求说明
//
// 本接口用于更新特定计算实例下的某个计算作业所保存的Flink SQL草稿内容。请确保提供的`InstanceId`和`JobName`准确无误，并且该作业当前状态支持进行SQL修改操作。
//
// - **注意事项**：
//
//   - 确保目标实例处于运行状态。
//
//   - 当前作业状态需允许修改SQL，即作业不应处于调试或变更过程中。
//
//   - `DraftSql`参数应包含完整的、格式正确的Flink SQL语句。
//
// @param request - UpdateComputeJobDraftSqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobDraftSqlResponse
func (client *Client) UpdateComputeJobDraftSqlWithContext(ctx context.Context, request *UpdateComputeJobDraftSqlRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobDraftSqlResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
