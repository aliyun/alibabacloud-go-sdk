// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # List projects for tenant
//
// @param request - DescribeProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectsResponse
func (client *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjects"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List findings for one engine (SAST / SCA)
//
// @param request - DescribeScanResultsByEngineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScanResultsByEngineResponse
func (client *Client) DescribeScanResultsByEngineWithContext(ctx context.Context, projectId *string, scanId *string, engine *string, request *DescribeScanResultsByEngineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScanResultsByEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineState) {
		query["baselineState"] = request.BaselineState
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PackageName) {
		query["packageName"] = request.PackageName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScanResultsByEngine"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans/" + dara.PercentEncode(dara.StringValue(scanId)) + "/results/" + dara.PercentEncode(dara.StringValue(engine))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScanResultsByEngineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List scans for project
//
// @param request - DescribeScansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScansResponse
func (client *Client) DescribeScansWithContext(ctx context.Context, projectId *string, request *DescribeScansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScans"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
