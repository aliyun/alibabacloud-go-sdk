// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 获取我的企业支持计划
//
// @param request - ListEnterpriseSupportPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseSupportPlanResponse
func (client *Client) ListEnterpriseSupportPlanWithContext(ctx context.Context, request *ListEnterpriseSupportPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnterpriseSupportPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseSupportPlan"),
		Version:     dara.String("2023-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/customerWorkbench/pop/api/enterpriseSupport/listEnterpriseSupportPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseSupportPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取我的企业支持计划(下拉)
//
// @param request - ListEnterpriseSupportPlanSimpleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseSupportPlanSimpleResponse
func (client *Client) ListEnterpriseSupportPlanSimpleWithContext(ctx context.Context, request *ListEnterpriseSupportPlanSimpleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnterpriseSupportPlanSimpleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseSupportPlanSimple"),
		Version:     dara.String("2023-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/customerWorkbench/pop/api/enterpriseSupport/listEnterpriseSupportPlanSimple"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseSupportPlanSimpleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专家服务列表
//
// @param request - ListServiceApplyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceApplyResponse
func (client *Client) ListServiceApplyWithContext(ctx context.Context, request *ListServiceApplyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyType) {
		body["applyType"] = request.ApplyType
	}

	if !dara.IsNil(request.EndDate) {
		body["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		body["productCode"] = request.ProductCode
	}

	if !dara.IsNil(request.StartDate) {
		body["startDate"] = request.StartDate
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceApply"),
		Version:     dara.String("2023-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/customerWorkbench/pop/api/expert/service/listServiceApply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过UID分页获取云企任务单
//
// @param request - ListYunQiTaskByUidRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListYunQiTaskByUidResponse
func (client *Client) ListYunQiTaskByUidWithContext(ctx context.Context, request *ListYunQiTaskByUidRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListYunQiTaskByUidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		body["createTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["createTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.FreeOrderApplyCodes) {
		body["freeOrderApplyCodes"] = request.FreeOrderApplyCodes
	}

	if !dara.IsNil(request.FreeOrderApplyIds) {
		body["freeOrderApplyIds"] = request.FreeOrderApplyIds
	}

	if !dara.IsNil(request.OrderIds) {
		body["orderIds"] = request.OrderIds
	}

	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusList) {
		body["statusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListYunQiTaskByUid"),
		Version:     dara.String("2023-12-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/customerWorkbench/pop/api/record/listYunQiTaskByUid"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListYunQiTaskByUidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
