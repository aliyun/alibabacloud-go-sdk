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
