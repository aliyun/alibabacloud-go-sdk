// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Calls this operation to cancel a specified asynchronous task.
//
// Description:
//
// Only tasks that are in the Pending or Running state can be canceled.
//
// You can call the CancelTask operation to cancel a Cloud Control API task, but the tasks that have been started in the downstream Alibaba Cloud services cannot be canceled.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/operation/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to create resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, requestPath *string, request *CreateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to delete resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - DeleteResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, requestPath *string, tmpReq *DeleteResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An RFQ interface through which users can query resource prices.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPriceResponse
func (client *Client) GetPriceWithContext(ctx context.Context, requestPath *string, tmpReq *GetPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPriceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceAttributes) {
		request.ResourceAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceAttributes, dara.String("resourceAttributes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceAttributesShrink) {
		query["resourceAttributes"] = request.ResourceAttributesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrice"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the operation to obtain resource metadata.
//
// @param requestPath - the whole path of resource string
//
// @param headers - GetResourceTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceTypeWithContext(ctx context.Context, requestPath *string, headers *GetResourceTypeHeaders, runtime *dara.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceType"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the operation to query resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out CloudControl API.
//
// You can call this operation to query resources List and Get based on different request paths.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithContext(ctx context.Context, requestPath *string, tmpReq *GetResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResources"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to query a specified asynchronous task.
//
// Description:
//
// GET /api/v1/tasks/{taskId}.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the operation to query the valid values of resource attributes, such as RegionID and ZoneId.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - ListDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithContext(ctx context.Context, requestPath *string, tmpReq *ListDataSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributeName) {
		query["attributeName"] = request.AttributeName
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to list the supported services.
//
// Description:
//
// GET /api/v1/providers/{provider}/products.
//
// @param request - ListProductsRequest
//
// @param headers - ListProductsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithContext(ctx context.Context, provider *string, request *ListProductsRequest, headers *ListProductsHeaders, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/providers/" + dara.PercentEncode(dara.StringValue(provider)) + "/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to list the resource types of a service.
//
// Description:
//
// GET /api/v1/providers/{provider}/products/{product}/resourceTypes.
//
// @param tmpReq - ListResourceTypesRequest
//
// @param headers - ListResourceTypesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithContext(ctx context.Context, provider *string, product *string, tmpReq *ListResourceTypesRequest, headers *ListResourceTypesHeaders, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListResourceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceTypes) {
		request.ResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypes, dara.String("resourceTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceTypesShrink) {
		query["resourceTypes"] = request.ResourceTypesShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypes"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/providers/" + dara.PercentEncode(dara.StringValue(provider)) + "/products/" + dara.PercentEncode(dara.StringValue(product)) + "/resourceTypes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to update resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// If resources fail to be updated at any time, the Cloud Control API does not roll the resource back to the original status.
//
// The resource APIs cannot be rolled back. If the API operation is partially failed to be called, you can call the GetResource operation to view the latest status of the resource. If necessary, you can call the UpdateResource or DeleteResource operation to manually compensate for the failure.
//
// @param requestPath - the whole path of resource string
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, requestPath *string, request *UpdateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
