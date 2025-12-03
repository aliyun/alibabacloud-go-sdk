// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 商品授权码激活
//
// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithContext(ctx context.Context, request *ActivateLicenseRequest, runtime *dara.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.LicenseCode) {
		query["LicenseCode"] = request.LicenseCode
	}

	if !dara.IsNil(request.LicenseNo) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !dara.IsNil(request.LicensePublisher) {
		query["LicensePublisher"] = request.LicensePublisher
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateLicense"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务回调
//
// @param request - CallbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallbackTaskResponse
func (client *Client) CallbackTaskWithContext(ctx context.Context, request *CallbackTaskRequest, runtime *dara.RuntimeOptions) (_result *CallbackTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.OutTaskId) {
		query["OutTaskId"] = request.OutTaskId
	}

	if !dara.IsNil(request.PrincipalKey) {
		query["PrincipalKey"] = request.PrincipalKey
	}

	if !dara.IsNil(request.TaskData) {
		query["TaskData"] = request.TaskData
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CallbackTask"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CallbackTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询协议状态
//
// @param request - DescribeAgreementStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAgreementStatusResponse
func (client *Client) DescribeAgreementStatusWithContext(ctx context.Context, request *DescribeAgreementStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAgreementStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgreementCode) {
		query["AgreementCode"] = request.AgreementCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAgreementStatus"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAgreementStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴生成上传文件策略
//
// @param request - GenerateUploadFilePolicyForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadFilePolicyForPartnerResponse
func (client *Client) GenerateUploadFilePolicyForPartnerWithContext(ctx context.Context, request *GenerateUploadFilePolicyForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadFilePolicyForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUploadFilePolicyForPartner"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadFilePolicyForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取node节点通过流程id
//
// @param request - GetNodeByFlowIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeByFlowIdResponse
func (client *Client) GetNodeByFlowIdWithContext(ctx context.Context, request *GetNodeByFlowIdRequest, runtime *dara.RuntimeOptions) (_result *GetNodeByFlowIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeByFlowId"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeByFlowIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取node节点通过模版id
//
// @param request - GetNodeByTemplateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeByTemplateIdResponse
func (client *Client) GetNodeByTemplateIdWithContext(ctx context.Context, request *GetNodeByTemplateIdRequest, runtime *dara.RuntimeOptions) (_result *GetNodeByTemplateIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeByTemplateId"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeByTemplateIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴获取订单概要信息
//
// @param request - GetOrderSummaryForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderSummaryForPartnerResponse
func (client *Client) GetOrderSummaryForPartnerWithContext(ctx context.Context, request *GetOrderSummaryForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetOrderSummaryForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrderSummaryForPartner"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrderSummaryForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴获取用户跨平台信息
//
// @param request - GetPlatformUserInfoForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlatformUserInfoForPartnerResponse
func (client *Client) GetPlatformUserInfoForPartnerWithContext(ctx context.Context, request *GetPlatformUserInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetPlatformUserInfoForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PlatformType) {
		query["PlatformType"] = request.PlatformType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlatformUserInfoForPartner"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlatformUserInfoForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代理
//
// @param request - GetProxyByTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProxyByTypeResponse
func (client *Client) GetProxyByTypeWithContext(ctx context.Context, request *GetProxyByTypeRequest, runtime *dara.RuntimeOptions) (_result *GetProxyByTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProxyByType"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProxyByTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取redis值
//
// @param request - GetRedisValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRedisValueResponse
func (client *Client) GetRedisValueWithContext(ctx context.Context, request *GetRedisValueRequest, runtime *dara.RuntimeOptions) (_result *GetRedisValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRedisValue"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRedisValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取变量
//
// @param request - GetVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVariableResponse
func (client *Client) GetVariableWithContext(ctx context.Context, request *GetVariableRequest, runtime *dara.RuntimeOptions) (_result *GetVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVariable"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 识别验证码
//
// @param request - IdentifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IdentifyCodeResponse
func (client *Client) IdentifyCodeWithContext(ctx context.Context, request *IdentifyCodeRequest, runtime *dara.RuntimeOptions) (_result *IdentifyCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IdentifyCode"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IdentifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取协议变更识别模型
//
// @param request - PullRpaModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullRpaModelResponse
func (client *Client) PullRpaModelWithContext(ctx context.Context, request *PullRpaModelRequest, runtime *dara.RuntimeOptions) (_result *PullRpaModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PullRpaModel"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PullRpaModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RPA拉取任务
//
// @param request - PullTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullTaskResponse
func (client *Client) PullTaskWithContext(ctx context.Context, request *PullTaskRequest, runtime *dara.RuntimeOptions) (_result *PullTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.PrincipalKey) {
		query["PrincipalKey"] = request.PrincipalKey
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerSecurityTransport) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PullTask"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PullTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送RPA运行时任务
//
// @param request - PushRpaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushRpaTaskResponse
func (client *Client) PushRpaTaskWithContext(ctx context.Context, request *PushRpaTaskRequest, runtime *dara.RuntimeOptions) (_result *PushRpaTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.Request) {
		query["Request"] = request.Request
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushRpaTask"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushRpaTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送运行时任务明细
//
// @param request - PushRpaTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushRpaTaskDetailResponse
func (client *Client) PushRpaTaskDetailWithContext(ctx context.Context, request *PushRpaTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *PushRpaTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.InputData) {
		query["InputData"] = request.InputData
	}

	if !dara.IsNil(request.InputHtml) {
		query["InputHtml"] = request.InputHtml
	}

	if !dara.IsNil(request.InputScreenshot) {
		query["InputScreenshot"] = request.InputScreenshot
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ModelDetailId) {
		query["ModelDetailId"] = request.ModelDetailId
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.OutputData) {
		query["OutputData"] = request.OutputData
	}

	if !dara.IsNil(request.OutputHtml) {
		query["OutputHtml"] = request.OutputHtml
	}

	if !dara.IsNil(request.OutputScreenshot) {
		query["OutputScreenshot"] = request.OutputScreenshot
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskDetailId) {
		query["TaskDetailId"] = request.TaskDetailId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushRpaTaskDetail"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushRpaTaskDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴发送消息通知
//
// @param tmpReq - SendNotificationForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendNotificationForPartnerResponse
func (client *Client) SendNotificationForPartnerWithContext(ctx context.Context, tmpReq *SendNotificationForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SendNotificationForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendNotificationForPartnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParamMap) {
		request.ParamMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParamMap, dara.String("ParamMap"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SmartSubChannels) {
		request.SmartSubChannelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SmartSubChannels, dara.String("SmartSubChannels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.NotifyType) {
		query["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.NotifycationEventType) {
		query["NotifycationEventType"] = request.NotifycationEventType
	}

	if !dara.IsNil(request.ParamMapShrink) {
		query["ParamMap"] = request.ParamMapShrink
	}

	if !dara.IsNil(request.SendTarget) {
		query["SendTarget"] = request.SendTarget
	}

	if !dara.IsNil(request.SmartSubChannelsShrink) {
		query["SmartSubChannels"] = request.SmartSubChannelsShrink
	}

	if !dara.IsNil(request.TrackId) {
		query["TrackId"] = request.TrackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendNotificationForPartner"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendNotificationForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// redis设置
//
// @param request - SetRedisValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRedisValueResponse
func (client *Client) SetRedisValueWithContext(ctx context.Context, request *SetRedisValueRequest, runtime *dara.RuntimeOptions) (_result *SetRedisValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunKp) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OriginalRequest) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserBid) {
		query["UserBid"] = request.UserBid
	}

	if !dara.IsNil(request.UserCallerParentId) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !dara.IsNil(request.UserCallerType) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.UserKp) {
		query["UserKp"] = request.UserKp
	}

	if !dara.IsNil(request.UserMfaPresent) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !dara.IsNil(request.UserSecurityToken) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRedisValue"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRedisValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新协议状态
//
// @param request - UpdateAgreementStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgreementStatusResponse
func (client *Client) UpdateAgreementStatusWithContext(ctx context.Context, request *UpdateAgreementStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgreementStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgreementCode) {
		query["AgreementCode"] = request.AgreementCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgreementStatus"),
		Version:     dara.String("2021-01-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgreementStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
