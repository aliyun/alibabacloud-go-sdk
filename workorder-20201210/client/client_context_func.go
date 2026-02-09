// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - GetTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketTemplateResponse
func (client *Client) GetTicketTemplateWithContext(ctx context.Context, request *GetTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTicketTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicketTemplate"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 是否首次访问BBS工单
//
// @param request - IsFirstBbsTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsFirstBbsTicketResponse
func (client *Client) IsFirstBbsTicketWithContext(ctx context.Context, request *IsFirstBbsTicketRequest, runtime *dara.RuntimeOptions) (_result *IsFirstBbsTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsFirstBbsTicket"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsFirstBbsTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuggestCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuggestCategoryResponse
func (client *Client) SuggestCategoryWithContext(ctx context.Context, request *SuggestCategoryRequest, runtime *dara.RuntimeOptions) (_result *SuggestCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.UseHotSuggestCatchAll) {
		query["UseHotSuggestCatchAll"] = request.UseHotSuggestCatchAll
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuggestCategory"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuggestCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
