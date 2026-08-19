// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomerLabelResponse
func (client *Client) AddCustomerLabelWithContext(ctx context.Context, request *AddCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *AddCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endtime) {
		query["Endtime"] = request.Endtime
	}

	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomerLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 传入多个labelSeries查询标签
//
// @param tmpReq - BatchFetchAccountLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchFetchAccountLabelResponse
func (client *Client) BatchFetchAccountLabelWithContext(ctx context.Context, tmpReq *BatchFetchAccountLabelRequest, runtime *dara.RuntimeOptions) (_result *BatchFetchAccountLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchFetchAccountLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSeriesList) {
		request.LabelSeriesListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSeriesList, dara.String("LabelSeriesList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Instant) {
		query["Instant"] = request.Instant
	}

	if !dara.IsNil(request.LabelSeriesListShrink) {
		query["LabelSeriesList"] = request.LabelSeriesListShrink
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchFetchAccountLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchFetchAccountLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomerLabelResponse
func (client *Client) DeleteCustomerLabelWithContext(ctx context.Context, request *DeleteCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomerLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerLabelResponse
func (client *Client) QueryCustomerLabelWithContext(ctx context.Context, request *QueryCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Instant) {
		query["Instant"] = request.Instant
	}

	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomerLabelByConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerLabelByConfigGroupResponse
func (client *Client) QueryCustomerLabelByConfigGroupWithContext(ctx context.Context, request *QueryCustomerLabelByConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerLabelByConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerLabelByConfigGroup"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerLabelByConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
