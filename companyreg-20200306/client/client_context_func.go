// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - BindProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindProduceAuthorizationResponse
func (client *Client) BindProduceAuthorizationWithContext(ctx context.Context, request *BindProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *BindProduceAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedUserIds) {
		body["AuthorizedUserIds"] = request.AuthorizedUserIds
	}

	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindProduceAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseIntentionForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseIntentionForPartnerResponse
func (client *Client) CloseIntentionForPartnerWithContext(ctx context.Context, request *CloseIntentionForPartnerRequest, runtime *dara.RuntimeOptions) (_result *CloseIntentionForPartnerResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseIntentionForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseIntentionForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseUserIntentionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseUserIntentionResponse
func (client *Client) CloseUserIntentionWithContext(ctx context.Context, request *CloseUserIntentionRequest, runtime *dara.RuntimeOptions) (_result *CloseUserIntentionResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseUserIntention"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseUserIntentionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateBusinessOpportunity
//
// @param request - CreateBusinessOpportunityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessOpportunityResponse
func (client *Client) CreateBusinessOpportunityWithContext(ctx context.Context, request *CreateBusinessOpportunityRequest, runtime *dara.RuntimeOptions) (_result *CreateBusinessOpportunityResponse, _err error) {
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

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.VCode) {
		query["VCode"] = request.VCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBusinessOpportunity"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBusinessOpportunityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateProduceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProduceForPartnerResponse
func (client *Client) CreateProduceForPartnerWithContext(ctx context.Context, request *CreateProduceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *CreateProduceForPartnerResponse, _err error) {
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

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProduceForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProduceForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePartnerConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePartnerConfigResponse
func (client *Client) DescribePartnerConfigWithContext(ctx context.Context, request *DescribePartnerConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribePartnerConfigResponse, _err error) {
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

	if !dara.IsNil(request.PartnerCode) {
		query["PartnerCode"] = request.PartnerCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePartnerConfig"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePartnerConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GenerateUploadFilePolicy
//
// @param request - GenerateUploadFilePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadFilePolicyResponse
func (client *Client) GenerateUploadFilePolicyWithContext(ctx context.Context, request *GenerateUploadFilePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadFilePolicyResponse, _err error) {
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
		Action:      dara.String("GenerateUploadFilePolicy"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAlipayUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlipayUrlResponse
func (client *Client) GetAlipayUrlWithContext(ctx context.Context, request *GetAlipayUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAlipayUrlResponse, _err error) {
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
		Action:      dara.String("GetAlipayUrl"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIntentionNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentionNoteResponse
func (client *Client) ListIntentionNoteWithContext(ctx context.Context, request *ListIntentionNoteRequest, runtime *dara.RuntimeOptions) (_result *ListIntentionNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntentionNote"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntentionNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProduceAuthorizationResponse
func (client *Client) ListProduceAuthorizationWithContext(ctx context.Context, request *ListProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ListProduceAuthorizationResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProduceAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserDetailSolutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDetailSolutionsResponse
func (client *Client) ListUserDetailSolutionsWithContext(ctx context.Context, request *ListUserDetailSolutionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserDetailSolutionsResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDetailSolutions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDetailSolutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserIntentionNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserIntentionNotesResponse
func (client *Client) ListUserIntentionNotesWithContext(ctx context.Context, request *ListUserIntentionNotesRequest, runtime *dara.RuntimeOptions) (_result *ListUserIntentionNotesResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserIntentionNotes"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserIntentionNotesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户控制天需求列表查询
//
// @param request - ListUserIntentionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserIntentionsResponse
func (client *Client) ListUserIntentionsWithContext(ctx context.Context, request *ListUserIntentionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserIntentionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BizTypes) {
		query["BizTypes"] = request.BizTypes
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortFiled) {
		query["SortFiled"] = request.SortFiled
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WithExtInfo) {
		query["WithExtInfo"] = request.WithExtInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserIntentions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserIntentionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserProduceOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserProduceOperateLogsResponse
func (client *Client) ListUserProduceOperateLogsWithContext(ctx context.Context, request *ListUserProduceOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *ListUserProduceOperateLogsResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserProduceOperateLogs"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserProduceOperateLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - ListUserSolutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSolutionsResponse
func (client *Client) ListUserSolutionsWithContext(ctx context.Context, tmpReq *ListUserSolutionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserSolutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUserSolutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExistStatus) {
		request.ExistStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExistStatus, dara.String("ExistStatus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExistStatusShrink) {
		query["ExistStatus"] = request.ExistStatusShrink
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserSolutions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserSolutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商玄坛呼叫中心操作
//
// @param request - OperateCallCenterForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateCallCenterForPartnerResponse
func (client *Client) OperateCallCenterForPartnerWithContext(ctx context.Context, request *OperateCallCenterForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateCallCenterForPartnerResponse, _err error) {
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

	if !dara.IsNil(request.CallAction) {
		query["CallAction"] = request.CallAction
	}

	if !dara.IsNil(request.EmployeeCode) {
		query["EmployeeCode"] = request.EmployeeCode
	}

	if !dara.IsNil(request.Request) {
		query["Request"] = request.Request
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateCallCenterForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateCallCenterForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - OperateProduceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateProduceForPartnerResponse
func (client *Client) OperateProduceForPartnerWithContext(ctx context.Context, request *OperateProduceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateProduceForPartnerResponse, _err error) {
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

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateProduceForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateProduceForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutMeasureDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMeasureDataResponse
func (client *Client) PutMeasureDataWithContext(ctx context.Context, request *PutMeasureDataRequest, runtime *dara.RuntimeOptions) (_result *PutMeasureDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMeasureData"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutMeasureReadyFlagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMeasureReadyFlagResponse
func (client *Client) PutMeasureReadyFlagWithContext(ctx context.Context, request *PutMeasureReadyFlagRequest, runtime *dara.RuntimeOptions) (_result *PutMeasureReadyFlagResponse, _err error) {
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

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReadyFlag) {
		query["ReadyFlag"] = request.ReadyFlag
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMeasureReadyFlag"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取玄坛合作伙伴双呼时可用外呼号码
//
// @param request - QueryAvailableNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableNumbersResponse
func (client *Client) QueryAvailableNumbersWithContext(ctx context.Context, request *QueryAvailableNumbersRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableNumbersResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableNumbers"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableNumbersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBagRemainingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBagRemainingResponse
func (client *Client) QueryBagRemainingWithContext(ctx context.Context, request *QueryBagRemainingRequest, runtime *dara.RuntimeOptions) (_result *QueryBagRemainingResponse, _err error) {
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
		Action:      dara.String("QueryBagRemaining"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBagRemainingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询玄坛外呼语音文件
//
// @param request - QueryCallRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallRecordListResponse
func (client *Client) QueryCallRecordListWithContext(ctx context.Context, request *QueryCallRecordListRequest, runtime *dara.RuntimeOptions) (_result *QueryCallRecordListResponse, _err error) {
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

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallRecordList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallRecordListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceResponse
func (client *Client) QueryInstanceWithContext(ctx context.Context, request *QueryInstanceRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstance"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryPartnerIntentionList
//
// @param request - QueryPartnerIntentionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPartnerIntentionListResponse
func (client *Client) QueryPartnerIntentionListWithContext(ctx context.Context, request *QueryPartnerIntentionListRequest, runtime *dara.RuntimeOptions) (_result *QueryPartnerIntentionListResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPartnerIntentionList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPartnerIntentionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryPartnerProduceList
//
// @param request - QueryPartnerProduceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPartnerProduceListResponse
func (client *Client) QueryPartnerProduceListWithContext(ctx context.Context, request *QueryPartnerProduceListRequest, runtime *dara.RuntimeOptions) (_result *QueryPartnerProduceListResponse, _err error) {
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
		Action:      dara.String("QueryPartnerProduceList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPartnerProduceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商玄坛外呼呼叫中心事件回传
//
// @param request - RecordCallCenterEventForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordCallCenterEventForPartnerResponse
func (client *Client) RecordCallCenterEventForPartnerWithContext(ctx context.Context, request *RecordCallCenterEventForPartnerRequest, runtime *dara.RuntimeOptions) (_result *RecordCallCenterEventForPartnerResponse, _err error) {
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

	if !dara.IsNil(request.CallAction) {
		query["CallAction"] = request.CallAction
	}

	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.ConnId) {
		query["ConnId"] = request.ConnId
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.EmployeeCode) {
		query["EmployeeCode"] = request.EmployeeCode
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RelatedId) {
		query["RelatedId"] = request.RelatedId
	}

	if !dara.IsNil(request.SecretMobile) {
		query["SecretMobile"] = request.SecretMobile
	}

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordCallCenterEventForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordCallCenterEventForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RecordPostBack
//
// @param request - RecordPostBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordPostBackResponse
func (client *Client) RecordPostBackWithContext(ctx context.Context, request *RecordPostBackRequest, runtime *dara.RuntimeOptions) (_result *RecordPostBackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["bizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Contactor) {
		query["contactor"] = request.Contactor
	}

	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.EntityKey) {
		query["entityKey"] = request.EntityKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordPostBack"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordPostBackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectSolutionResponse
func (client *Client) RejectSolutionWithContext(ctx context.Context, request *RejectSolutionRequest, runtime *dara.RuntimeOptions) (_result *RejectSolutionResponse, _err error) {
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

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.SolutionBizId) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectSolutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectUserSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectUserSolutionResponse
func (client *Client) RejectUserSolutionWithContext(ctx context.Context, request *RejectUserSolutionRequest, runtime *dara.RuntimeOptions) (_result *RejectUserSolutionResponse, _err error) {
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

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.SolutionBizId) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectUserSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectUserSolutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseProduceAuthorizationResponse
func (client *Client) ReleaseProduceAuthorizationWithContext(ctx context.Context, request *ReleaseProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ReleaseProduceAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedUserId) {
		body["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseProduceAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛双呼外呼发起
//
// @param request - StartBackToBackCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBackToBackCallResponse
func (client *Client) StartBackToBackCallWithContext(ctx context.Context, request *StartBackToBackCallRequest, runtime *dara.RuntimeOptions) (_result *StartBackToBackCallResponse, _err error) {
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

	if !dara.IsNil(request.CallCenterNumber) {
		query["CallCenterNumber"] = request.CallCenterNumber
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.MobileKey) {
		query["MobileKey"] = request.MobileKey
	}

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBackToBackCall"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBackToBackCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴提交需求单
//
// @param request - SubmitIntentionForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIntentionForPartnerResponse
func (client *Client) SubmitIntentionForPartnerWithContext(ctx context.Context, request *SubmitIntentionForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SubmitIntentionForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.CommodityType) {
		query["CommodityType"] = request.CommodityType
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.Grade) {
		query["Grade"] = request.Grade
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIntentionForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIntentionForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitIntentionNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIntentionNoteResponse
func (client *Client) SubmitIntentionNoteWithContext(ctx context.Context, request *SubmitIntentionNoteRequest, runtime *dara.RuntimeOptions) (_result *SubmitIntentionNoteResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIntentionNote"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIntentionNoteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSolutionResponse
func (client *Client) SubmitSolutionWithContext(ctx context.Context, request *SubmitSolutionRequest, runtime *dara.RuntimeOptions) (_result *SubmitSolutionResponse, _err error) {
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

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Solution) {
		query["Solution"] = request.Solution
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSolutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛需求单转派小二
//
// @param request - TransferIntentionOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferIntentionOwnerResponse
func (client *Client) TransferIntentionOwnerWithContext(ctx context.Context, request *TransferIntentionOwnerRequest, runtime *dara.RuntimeOptions) (_result *TransferIntentionOwnerResponse, _err error) {
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

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferIntentionOwner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferIntentionOwnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛服务单转派小二
//
// @param request - TransferProduceOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferProduceOwnerResponse
func (client *Client) TransferProduceOwnerWithContext(ctx context.Context, request *TransferProduceOwnerRequest, runtime *dara.RuntimeOptions) (_result *TransferProduceOwnerResponse, _err error) {
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

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferProduceOwner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferProduceOwnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
