// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// sas-绑定授权到机器
//
// @param tmpReq - BindAuthToMachineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAuthToMachineResponse
func (client *Client) BindAuthToMachineWithContext(ctx context.Context, tmpReq *BindAuthToMachineRequest, runtime *dara.RuntimeOptions) (_result *BindAuthToMachineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindAuthToMachineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAuthToMachine"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAuthToMachineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-初始化云安全中心模块规则
//
// @param tmpReq - CreateSasTrialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSasTrialResponse
func (client *Client) CreateSasTrialWithContext(ctx context.Context, tmpReq *CreateSasTrialRequest, runtime *dara.RuntimeOptions) (_result *CreateSasTrialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSasTrialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSasTrial"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSasTrialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-创建服务关联角色
//
// @param tmpReq - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithContext(ctx context.Context, tmpReq *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateServiceLinkedRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建全盘扫描任务
//
// @param request - CreateVirusScanOnceTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirusScanOnceTaskResponse
func (client *Client) CreateVirusScanOnceTaskWithContext(ctx context.Context, request *CreateVirusScanOnceTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVirusScanOnceTaskResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirusScanOnceTask"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirusScanOnceTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-查询云安全中心实例列表
//
// @param tmpReq - DescribeCloudCenterInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudCenterInstancesResponse
func (client *Client) DescribeCloudCenterInstancesWithContext(ctx context.Context, tmpReq *DescribeCloudCenterInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudCenterInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCloudCenterInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudCenterInstances"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudCenterInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-查询服务关联角色状态
//
// @param tmpReq - DescribeServiceLinkedRoleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLinkedRoleStatusResponse
func (client *Client) DescribeServiceLinkedRoleStatusWithContext(ctx context.Context, tmpReq *DescribeServiceLinkedRoleStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceLinkedRoleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeServiceLinkedRoleStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceLinkedRoleStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceLinkedRoleStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-查询安全告警事件
//
// @param tmpReq - DescribeSuspEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEventsWithContext(ctx context.Context, tmpReq *DescribeSuspEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSuspEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSuspEvents"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-查看已购买的云安全中心实例的版本详情
//
// @param tmpReq - DescribeVersionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVersionConfigResponse
func (client *Client) DescribeVersionConfigWithContext(ctx context.Context, tmpReq *DescribeVersionConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVersionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeVersionConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVersionConfig"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVersionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取告警记录分析结果
//
// @param tmpReq - GetAlertRecordAnalysisResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertRecordAnalysisResultResponse
func (client *Client) GetAlertRecordAnalysisResultWithContext(ctx context.Context, tmpReq *GetAlertRecordAnalysisResultRequest, runtime *dara.RuntimeOptions) (_result *GetAlertRecordAnalysisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAlertRecordAnalysisResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UniqueTagList) {
		request.UniqueTagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UniqueTagList, dara.String("UniqueTagList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmUniqueInfo) {
		query["AlarmUniqueInfo"] = request.AlarmUniqueInfo
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.UniqueInfo) {
		query["UniqueInfo"] = request.UniqueInfo
	}

	if !dara.IsNil(request.UniqueTagListShrink) {
		query["UniqueTagList"] = request.UniqueTagListShrink
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertRecordAnalysisResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertRecordAnalysisResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用云安全中心部分接口
//
// @param tmpReq - GetAliYunSafeCenterResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliYunSafeCenterResultResponse
func (client *Client) GetAliYunSafeCenterResultWithContext(ctx context.Context, tmpReq *GetAliYunSafeCenterResultRequest, runtime *dara.RuntimeOptions) (_result *GetAliYunSafeCenterResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAliYunSafeCenterResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateSimilarSecurityEventsQueryTaskRequest) {
		request.CreateSimilarSecurityEventsQueryTaskRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateSimilarSecurityEventsQueryTaskRequest, dara.String("CreateSimilarSecurityEventsQueryTaskRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeInstancesFullStatusRequest) {
		request.DescribeInstancesFullStatusRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeInstancesFullStatusRequest, dara.String("DescribeInstancesFullStatusRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeSecurityEventOperationStatusRequest) {
		request.DescribeSecurityEventOperationStatusRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeSecurityEventOperationStatusRequest, dara.String("DescribeSecurityEventOperationStatusRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DescribeSimilarSecurityEventsRequest) {
		request.DescribeSimilarSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DescribeSimilarSecurityEventsRequest, dara.String("DescribeSimilarSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GetAssetDetailByUuidRequest) {
		request.GetAssetDetailByUuidRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetAssetDetailByUuidRequest, dara.String("GetAssetDetailByUuidRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HandleSecurityEventsRequest) {
		request.HandleSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HandleSecurityEventsRequest, dara.String("HandleSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HandleSimilarSecurityEventsRequest) {
		request.HandleSimilarSecurityEventsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HandleSimilarSecurityEventsRequest, dara.String("HandleSimilarSecurityEventsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ListInstancesRequest) {
		request.ListInstancesRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListInstancesRequest, dara.String("ListInstancesRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSimilarSecurityEventsQueryTaskRequestShrink) {
		query["CreateSimilarSecurityEventsQueryTaskRequest"] = request.CreateSimilarSecurityEventsQueryTaskRequestShrink
	}

	if !dara.IsNil(request.DescribeInstancesFullStatusRequestShrink) {
		query["DescribeInstancesFullStatusRequest"] = request.DescribeInstancesFullStatusRequestShrink
	}

	if !dara.IsNil(request.DescribeSecurityEventOperationStatusRequestShrink) {
		query["DescribeSecurityEventOperationStatusRequest"] = request.DescribeSecurityEventOperationStatusRequestShrink
	}

	if !dara.IsNil(request.DescribeSimilarSecurityEventsRequestShrink) {
		query["DescribeSimilarSecurityEventsRequest"] = request.DescribeSimilarSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.GetAssetDetailByUuidRequestShrink) {
		query["GetAssetDetailByUuidRequest"] = request.GetAssetDetailByUuidRequestShrink
	}

	if !dara.IsNil(request.HandleSecurityEventsRequestShrink) {
		query["HandleSecurityEventsRequest"] = request.HandleSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.HandleSimilarSecurityEventsRequestShrink) {
		query["HandleSimilarSecurityEventsRequest"] = request.HandleSimilarSecurityEventsRequestShrink
	}

	if !dara.IsNil(request.InterfaceCode) {
		query["InterfaceCode"] = request.InterfaceCode
	}

	if !dara.IsNil(request.ListInstancesRequestShrink) {
		query["ListInstancesRequest"] = request.ListInstancesRequestShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAliYunSafeCenterResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAliYunSafeCenterResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-获取能否试用
//
// @param tmpReq - GetCanTrySasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCanTrySasResponse
func (client *Client) GetCanTrySasWithContext(ctx context.Context, tmpReq *GetCanTrySasRequest, runtime *dara.RuntimeOptions) (_result *GetCanTrySasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCanTrySasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCanTrySas"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCanTrySasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全合规包id
//
// @param request - GetCompliancePackIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompliancePackIdResponse
func (client *Client) GetCompliancePackIdWithContext(ctx context.Context, request *GetCompliancePackIdRequest, runtime *dara.RuntimeOptions) (_result *GetCompliancePackIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCompliancePackId"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompliancePackIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一键处置赋权状态
//
// @param request - GetDisposalToolStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisposalToolStatusResponse
func (client *Client) GetDisposalToolStatusWithContext(ctx context.Context, request *GetDisposalToolStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDisposalToolStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDisposalToolStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDisposalToolStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户确认安全联系人记录
//
// @param request - GetNotificationClickRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationClickRecordResponse
func (client *Client) GetNotificationClickRecordWithContext(ctx context.Context, request *GetNotificationClickRecordRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationClickRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationClickRecord"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationClickRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全联系人全部信息
//
// @param request - GetNotificationContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationContactsResponse
func (client *Client) GetNotificationContactsWithContext(ctx context.Context, request *GetNotificationContactsRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationContacts"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全联系人待处理数
//
// @param request - GetNotificationPendNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotificationPendNumberResponse
func (client *Client) GetNotificationPendNumberWithContext(ctx context.Context, request *GetNotificationPendNumberRequest, runtime *dara.RuntimeOptions) (_result *GetNotificationPendNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotificationPendNumber"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotificationPendNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云资源管控事件详情
//
// @param tmpReq - GetResourceControlEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceControlEventResponse
func (client *Client) GetResourceControlEventWithContext(ctx context.Context, tmpReq *GetResourceControlEventRequest, runtime *dara.RuntimeOptions) (_result *GetResourceControlEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourceControlEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceControlEvent"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceControlEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全体检基础信息
//
// @param request - GetSecurityCheckBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityCheckBaseInfoResponse
func (client *Client) GetSecurityCheckBaseInfoWithContext(ctx context.Context, request *GetSecurityCheckBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityCheckBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityCheckBaseInfo"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityCheckBaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全检查结果基础信息
//
// @param request - GetSecurityCheckResultBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityCheckResultBaseInfoResponse
func (client *Client) GetSecurityCheckResultBaseInfoWithContext(ctx context.Context, request *GetSecurityCheckResultBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityCheckResultBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityCheckResultBaseInfo"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityCheckResultBaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全优化建议列表
//
// @param tmpReq - GetSecuritySuggestionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySuggestionListResponse
func (client *Client) GetSecuritySuggestionListWithContext(ctx context.Context, tmpReq *GetSecuritySuggestionListRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySuggestionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSecuritySuggestionListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListConfigRulesRequest) {
		request.ListConfigRulesRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListConfigRulesRequest, dara.String("ListConfigRulesRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ListConfigRulesRequestShrink) {
		query["ListConfigRulesRequest"] = request.ListConfigRulesRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySuggestionList"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySuggestionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全优化建议条数
//
// @param request - GetSecuritySuggestionNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySuggestionNumberResponse
func (client *Client) GetSecuritySuggestionNumberWithContext(ctx context.Context, request *GetSecuritySuggestionNumberRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySuggestionNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySuggestionNumber"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySuggestionNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务关联角色状态
//
// @param request - GetServiceLinkedRoleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLinkedRoleStatusResponse
func (client *Client) GetServiceLinkedRoleStatusWithContext(ctx context.Context, request *GetServiceLinkedRoleStatusRequest, runtime *dara.RuntimeOptions) (_result *GetServiceLinkedRoleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceLinkedRoleStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceLinkedRoleStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-获取有效抵扣实例
//
// @param tmpReq - GetValidDeductInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetValidDeductInstancesResponse
func (client *Client) GetValidDeductInstancesWithContext(ctx context.Context, tmpReq *GetValidDeductInstancesRequest, runtime *dara.RuntimeOptions) (_result *GetValidDeductInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetValidDeductInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetValidDeductInstances"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetValidDeductInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-初始化云安全中心模块规则
//
// @param tmpReq - InitSasModuleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitSasModuleRuleResponse
func (client *Client) InitSasModuleRuleWithContext(ctx context.Context, tmpReq *InitSasModuleRuleRequest, runtime *dara.RuntimeOptions) (_result *InitSasModuleRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InitSasModuleRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBind) {
		query["AutoBind"] = request.AutoBind
	}

	if !dara.IsNil(request.InstancesShrink) {
		query["Instances"] = request.InstancesShrink
	}

	if !dara.IsNil(request.IsTrial) {
		query["IsTrial"] = request.IsTrial
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitSasModuleRule"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitSasModuleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全盘扫描结果
//
// @param request - ListVirusScanMachineEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanMachineEventResponse
func (client *Client) ListVirusScanMachineEventWithContext(ctx context.Context, request *ListVirusScanMachineEventRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanMachineEventResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OperateTaskId) {
		query["OperateTaskId"] = request.OperateTaskId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanMachineEvent"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanMachineEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-开启试用套餐
//
// @param request - OpenTrialPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenTrialPackageResponse
func (client *Client) OpenTrialPackageWithContext(ctx context.Context, request *OpenTrialPackageRequest, runtime *dara.RuntimeOptions) (_result *OpenTrialPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCloseSwitch) {
		query["AutoCloseSwitch"] = request.AutoCloseSwitch
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenTrialPackage"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenTrialPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账号安全事件
//
// @param request - QueryAccountSafetyIncidentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountSafetyIncidentResponse
func (client *Client) QueryAccountSafetyIncidentWithContext(ctx context.Context, request *QueryAccountSafetyIncidentRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountSafetyIncidentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.CaseCode) {
		query["CaseCode"] = request.CaseCode
	}

	if !dara.IsNil(request.Current) {
		query["Current"] = request.Current
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PunishEndTime) {
		query["PunishEndTime"] = request.PunishEndTime
	}

	if !dara.IsNil(request.PunishStartTime) {
		query["PunishStartTime"] = request.PunishStartTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountSafetyIncident"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountSafetyIncidentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云上安全指南的订阅状态
//
// @param request - QueryGuideSubStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGuideSubStatusResponse
func (client *Client) QueryGuideSubStatusWithContext(ctx context.Context, request *QueryGuideSubStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryGuideSubStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGuideSubStatus"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGuideSubStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云资源管控事件
//
// @param tmpReq - QueryResourceControlEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResourceControlEventsResponse
func (client *Client) QueryResourceControlEventsWithContext(ctx context.Context, tmpReq *QueryResourceControlEventsRequest, runtime *dara.RuntimeOptions) (_result *QueryResourceControlEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryResourceControlEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionCodes) {
		request.ActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionCodes, dara.String("ActionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CaseCodesPrefix) {
		request.CaseCodesPrefixShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseCodesPrefix, dara.String("CaseCodesPrefix"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventCodes) {
		request.EventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventCodes, dara.String("EventCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeActionCodes) {
		request.ExcludeActionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeActionCodes, dara.String("ExcludeActionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeEventCodes) {
		request.ExcludeEventCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeEventCodes, dara.String("ExcludeEventCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExcludeReasons) {
		request.ExcludeReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeReasons, dara.String("ExcludeReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeReasons) {
		request.IncludeReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeReasons, dara.String("IncludeReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceCodes) {
		request.SourceCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceCodes, dara.String("SourceCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("StatusList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionCode) {
		query["ActionCode"] = request.ActionCode
	}

	if !dara.IsNil(request.ActionCodesShrink) {
		query["ActionCodes"] = request.ActionCodesShrink
	}

	if !dara.IsNil(request.AliyunLang) {
		query["AliyunLang"] = request.AliyunLang
	}

	if !dara.IsNil(request.BusinessCode) {
		query["BusinessCode"] = request.BusinessCode
	}

	if !dara.IsNil(request.CaseCodesPrefixShrink) {
		query["CaseCodesPrefix"] = request.CaseCodesPrefixShrink
	}

	if !dara.IsNil(request.Current) {
		query["Current"] = request.Current
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EventCode) {
		query["EventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventCodesShrink) {
		query["EventCodes"] = request.EventCodesShrink
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	if !dara.IsNil(request.ExcludeActionCodesShrink) {
		query["ExcludeActionCodes"] = request.ExcludeActionCodesShrink
	}

	if !dara.IsNil(request.ExcludeEventCodesShrink) {
		query["ExcludeEventCodes"] = request.ExcludeEventCodesShrink
	}

	if !dara.IsNil(request.ExcludeReasonsShrink) {
		query["ExcludeReasons"] = request.ExcludeReasonsShrink
	}

	if !dara.IsNil(request.IncludeReasonsShrink) {
		query["IncludeReasons"] = request.IncludeReasonsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PunishEndTime) {
		query["PunishEndTime"] = request.PunishEndTime
	}

	if !dara.IsNil(request.PunishStartTime) {
		query["PunishStartTime"] = request.PunishStartTime
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.SourceCodesShrink) {
		query["SourceCodes"] = request.SourceCodesShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["StatusList"] = request.StatusListShrink
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResourceControlEvents"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResourceControlEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询安全体检简报
//
// @param request - QuerySecurityCheckReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecurityCheckReportResponse
func (client *Client) QuerySecurityCheckReportWithContext(ctx context.Context, request *QuerySecurityCheckReportRequest, runtime *dara.RuntimeOptions) (_result *QuerySecurityCheckReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecurityCheckReport"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecurityCheckReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启处置工具授权
//
// @param request - StartDisposalToolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDisposalToolServiceResponse
func (client *Client) StartDisposalToolServiceWithContext(ctx context.Context, request *StartDisposalToolServiceRequest, runtime *dara.RuntimeOptions) (_result *StartDisposalToolServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDisposalToolService"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDisposalToolServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启安全体检
//
// @param request - StartSecurityCheckServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSecurityCheckServiceResponse
func (client *Client) StartSecurityCheckServiceWithContext(ctx context.Context, request *StartSecurityCheckServiceRequest, runtime *dara.RuntimeOptions) (_result *StartSecurityCheckServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("StartSecurityCheckService"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSecurityCheckServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请解封
//
// @param tmpReq - SubmitApplyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitApplyRecordResponse
func (client *Client) SubmitApplyRecordWithContext(ctx context.Context, tmpReq *SubmitApplyRecordRequest, runtime *dara.RuntimeOptions) (_result *SubmitApplyRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitApplyRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyRequest) {
		query["ApplyRequest"] = request.ApplyRequest
	}

	if !dara.IsNil(request.CommitmentLetter) {
		query["CommitmentLetter"] = request.CommitmentLetter
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	if !dara.IsNil(request.QualificationProof) {
		query["QualificationProof"] = request.QualificationProof
	}

	if !dara.IsNil(request.Trial) {
		query["Trial"] = request.Trial
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitApplyRecord"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitApplyRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sas-更新后付费绑定关系
//
// @param tmpReq - UpdatePostPaidBindRelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostPaidBindRelResponse
func (client *Client) UpdatePostPaidBindRelWithContext(ctx context.Context, tmpReq *UpdatePostPaidBindRelRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostPaidBindRelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePostPaidBindRelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SdkRequest) {
		request.SdkRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdkRequest, dara.String("SdkRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SdkRequestShrink) {
		query["SdkRequest"] = request.SdkRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePostPaidBindRel"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostPaidBindRelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新体检结果
//
// @param request - UpdateSecurityCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityCheckResultResponse
func (client *Client) UpdateSecurityCheckResultWithContext(ctx context.Context, request *UpdateSecurityCheckResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityCheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityCheckResult"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityCheckResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
