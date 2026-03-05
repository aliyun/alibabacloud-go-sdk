// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # CreateProxyBrandUser
//
// @param request - CreateProxyBrandUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProxyBrandUserResponse
func (client *Client) CreateProxyBrandUserWithContext(ctx context.Context, request *CreateProxyBrandUserRequest, runtime *dara.RuntimeOptions) (_result *CreateProxyBrandUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Company) {
		query["Company"] = request.Company
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPhone) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProxyBrandUser"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProxyBrandUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateUnionTask
//
// @param tmpReq - CreateUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUnionTaskResponse
func (client *Client) CreateUnionTaskWithContext(ctx context.Context, tmpReq *CreateUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUnionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MediaIdBlackList) {
		request.MediaIdBlackListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaIdBlackList, dara.String("MediaIdBlackList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MediaIdWhiteList) {
		request.MediaIdWhiteListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaIdWhiteList, dara.String("MediaIdWhiteList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AnchorId) {
		query["AnchorId"] = request.AnchorId
	}

	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChargePloy) {
		query["ChargePloy"] = request.ChargePloy
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContentId) {
		query["ContentId"] = request.ContentId
	}

	if !dara.IsNil(request.ContentUrl) {
		query["ContentUrl"] = request.ContentUrl
	}

	if !dara.IsNil(request.CustomCreativeType) {
		query["CustomCreativeType"] = request.CustomCreativeType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IndustryLabelBagId) {
		query["IndustryLabelBagId"] = request.IndustryLabelBagId
	}

	if !dara.IsNil(request.MediaIdBlackListShrink) {
		query["MediaIdBlackList"] = request.MediaIdBlackListShrink
	}

	if !dara.IsNil(request.MediaIdWhiteListShrink) {
		query["MediaIdWhiteList"] = request.MediaIdWhiteListShrink
	}

	if !dara.IsNil(request.MediaIndustry) {
		query["MediaIndustry"] = request.MediaIndustry
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OptimizationSwitch) {
		query["OptimizationSwitch"] = request.OptimizationSwitch
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaDay) {
		query["QuotaDay"] = request.QuotaDay
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskBizType) {
		query["TaskBizType"] = request.TaskBizType
	}

	if !dara.IsNil(request.TaskRuleType) {
		query["TaskRuleType"] = request.TaskRuleType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUnionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除聚合拉新品牌
//
// @param request - DeleteUnionBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUnionBrandResponse
func (client *Client) DeleteUnionBrandWithContext(ctx context.Context, request *DeleteUnionBrandRequest, runtime *dara.RuntimeOptions) (_result *DeleteUnionBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandMainId) {
		query["BrandMainId"] = request.BrandMainId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUnionBrand"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUnionBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # EndUnionTask
//
// @param request - EndUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndUnionTaskResponse
func (client *Client) EndUnionTaskWithContext(ctx context.Context, request *EndUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *EndUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndUnionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 聚合拉新开放接口查询媒体行业
//
// @param request - ListUnionMediaIndustryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUnionMediaIndustryResponse
func (client *Client) ListUnionMediaIndustryWithContext(ctx context.Context, request *ListUnionMediaIndustryRequest, runtime *dara.RuntimeOptions) (_result *ListUnionMediaIndustryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUnionMediaIndustry"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUnionMediaIndustryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAvailableBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableBalanceResponse
func (client *Client) QueryAvailableBalanceWithContext(ctx context.Context, request *QueryAvailableBalanceRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableBalanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableBalance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableBalanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryContentList
//
// @param request - QueryContentListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContentListResponse
func (client *Client) QueryContentListWithContext(ctx context.Context, request *QueryContentListRequest, runtime *dara.RuntimeOptions) (_result *QueryContentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskBizType) {
		query["TaskBizType"] = request.TaskBizType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContentList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取行业标签包
//
// @param request - QueryIndustryLabelBagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndustryLabelBagResponse
func (client *Client) QueryIndustryLabelBagWithContext(ctx context.Context, request *QueryIndustryLabelBagRequest, runtime *dara.RuntimeOptions) (_result *QueryIndustryLabelBagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryIndustryLabelBag"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIndustryLabelBagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskBizTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskBizTypeResponse
func (client *Client) QueryTaskBizTypeWithContext(ctx context.Context, request *QueryTaskBizTypeRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskBizTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskBizType"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskBizTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryTaskRuleLimit
//
// @param request - QueryTaskRuleLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskRuleLimitResponse
func (client *Client) QueryTaskRuleLimitWithContext(ctx context.Context, request *QueryTaskRuleLimitRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskRuleLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskRuleLimit"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskRuleLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionChannelResponse
func (client *Client) QueryUnionChannelWithContext(ctx context.Context, request *QueryUnionChannelRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionChannel"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryUnionContentInfo
//
// @param request - QueryUnionContentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionContentInfoResponse
func (client *Client) QueryUnionContentInfoWithContext(ctx context.Context, request *QueryUnionContentInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionContentInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ContentId) {
		query["ContentId"] = request.ContentId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionContentInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionContentInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionTaskInfoResponse
func (client *Client) QueryUnionTaskInfoWithContext(ctx context.Context, request *QueryUnionTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionTaskInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionTaskListResponse
func (client *Client) QueryUnionTaskListWithContext(ctx context.Context, request *QueryUnionTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionTaskList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启聚合拉新计划
//
// @param request - StartUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartUnionTaskResponse
func (client *Client) StartUnionTaskWithContext(ctx context.Context, request *StartUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *StartUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartUnionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停聚合拉新计划
//
// @param request - StopUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopUnionTaskResponse
func (client *Client) StopUnionTaskWithContext(ctx context.Context, request *StopUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *StopUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopUnionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 聚合拉新对外接口更新计划
//
// @param request - UpdateUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUnionTaskResponse
func (client *Client) UpdateUnionTaskWithContext(ctx context.Context, request *UpdateUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChargePloy) {
		query["ChargePloy"] = request.ChargePloy
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OptimizationSwitch) {
		query["OptimizationSwitch"] = request.OptimizationSwitch
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaDay) {
		query["QuotaDay"] = request.QuotaDay
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUnionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
