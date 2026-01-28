// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - ARMSQueryDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ARMSQueryDataSetResponse
func (client *Client) ARMSQueryDataSetWithContext(ctx context.Context, request *ARMSQueryDataSetRequest, runtime *dara.RuntimeOptions) (_result *ARMSQueryDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DateStr) {
		query["DateStr"] = request.DateStr
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.HackerUserId) {
		query["HackerUserId"] = request.HackerUserId
	}

	if !dara.IsNil(request.HungryMode) {
		query["HungryMode"] = request.HungryMode
	}

	if !dara.IsNil(request.IntervalInSec) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !dara.IsNil(request.IsDrillDown) {
		query["IsDrillDown"] = request.IsDrillDown
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MaxTime) {
		query["MaxTime"] = request.MaxTime
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.MinTime) {
		query["MinTime"] = request.MinTime
	}

	if !dara.IsNil(request.OptionalDims) {
		query["OptionalDims"] = request.OptionalDims
	}

	if !dara.IsNil(request.OrderByKey) {
		query["OrderByKey"] = request.OrderByKey
	}

	if !dara.IsNil(request.ReduceTail) {
		query["ReduceTail"] = request.ReduceTail
	}

	if !dara.IsNil(request.RequiredDims) {
		query["RequiredDims"] = request.RequiredDims
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ARMSQueryDataSet"),
		Version:     dara.String("2018-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ARMSQueryDataSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MetricQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MetricQueryResponse
func (client *Client) MetricQueryWithContext(ctx context.Context, request *MetricQueryRequest, runtime *dara.RuntimeOptions) (_result *MetricQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.HackerUserId) {
		query["HackerUserId"] = request.HackerUserId
	}

	if !dara.IsNil(request.IintervalInSec) {
		query["IintervalInSec"] = request.IintervalInSec
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Measures) {
		query["Measures"] = request.Measures
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MetricQuery"),
		Version:     dara.String("2018-10-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MetricQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
