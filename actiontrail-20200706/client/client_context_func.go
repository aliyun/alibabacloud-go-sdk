// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an advanced event query history record that saves a custom query conditional statement for reuse and management.
//
// Description:
//
// This topic provides a demo of how to save a conditional statement as an advanced event query history record. The conditional statement is used to query all `AccessKey` access management events in logs.
//
// @param request - CreateAdvancedQueryHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedQueryHistoryResponse
func (client *Client) CreateAdvancedQueryHistoryWithContext(ctx context.Context, request *CreateAdvancedQueryHistoryRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedQueryHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.QuerySql) {
		query["QuerySql"] = request.QuerySql
	}

	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedQueryHistory"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedQueryHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an advanced query template.
//
// @param request - CreateAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedQueryTemplateResponse
func (client *Client) CreateAdvancedQueryTemplateWithContext(ctx context.Context, request *CreateAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateSql) {
		query["TemplateSql"] = request.TemplateSql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data backfill task.
//
// Description:
//
// # Limitations
//
// - You must first call the [CreateTrail](https://help.aliyun.com/document_detail/212313.html) operation to create a single-account trail that delivers events to Simple Log Service (SLS).
//
// - An Alibaba Cloud account can have only one data backfill task running at a time.
//
// This topic provides an example of how to create data backfill task for the trail `trail-name`.
//
// @param request - CreateDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryHistoryJobResponse
func (client *Client) CreateDeliveryHistoryJobWithContext(ctx context.Context, request *CreateDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trail to deliver events to a destination for long-term storage and analysis, such as an Object Storage Service (OSS) bucket, a Simple Log Service (SLS) Logstore, or a MaxCompute project.
//
// Description:
//
// > By default, a trail that you create by using this API is in a **disabled*	- state. You must call the [StartLogging](https://help.aliyun.com/document_detail/432246.html) operation operation to enable the trail. After a trail is enabled, ActionTrail begins delivering events to your specified destination.
//
// ### Prerequisites
//
// Before you create a trail, you must have at least one of the following resources configured as a destination:
//
// - OSS
//
//	You must activate OSS and create a bucket.
//
// - SLS
//
//	You must activate SLS and create a Logstore.
//
//	> When you create a trail with an SLS destination, ActionTrail automatically creates a Logstore named `actiontrail_<trail_name>` in your specified project. To ensure the integrity of your audit data, this Logstore only accepts events delivered by ActionTrail.
//
// - MaxCompute
//
//	You must activate MaxCompute.
//
//	> When you create a trail with a MaxCompute destination, ActionTrail automatically creates a project named `actiontrail_<account_ID>`. To ensure the integrity of your audit data, this project only accepts events delivered by ActionTrail.
//
// ### Usage notes
//
// This example shows how to create a single-account trail named `trail-test` that delivers events to an OSS bucket named `audit-log`.
//
// @param request - CreateTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrailResponse
func (client *Client) CreateTrailWithContext(ctx context.Context, request *CreateTrailRequest, runtime *dara.RuntimeOptions) (_result *CreateTrailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventRW) {
		query["EventRW"] = request.EventRW
	}

	if !dara.IsNil(request.IsOrganizationTrail) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !dara.IsNil(request.MaxComputeProjectArn) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !dara.IsNil(request.MaxComputeWriteRoleArn) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssKeyPrefix) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !dara.IsNil(request.OssWriteRoleArn) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
	}

	if !dara.IsNil(request.SlsProjectArn) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !dara.IsNil(request.SlsWriteRoleArn) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !dara.IsNil(request.TrailRegion) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an advanced query record.
//
// @param request - DeleteAdvancedQueryHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvancedQueryHistoryResponse
func (client *Client) DeleteAdvancedQueryHistoryWithContext(ctx context.Context, request *DeleteAdvancedQueryHistoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdvancedQueryHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdvancedQueryHistory"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdvancedQueryHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an advanced query template.
//
// @param request - DeleteAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvancedQueryTemplateResponse
func (client *Client) DeleteAdvancedQueryTemplateWithContext(ctx context.Context, request *DeleteAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the data event selector for a specified trail.
//
// @param request - DeleteDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataEventSelectorResponse
func (client *Client) DeleteDataEventSelectorWithContext(ctx context.Context, request *DeleteDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataEventSelectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data backfill task.
//
// Description:
//
// This topic describes how to delete a data backfill task whose ID is `16602`.
//
// @param request - DeleteDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryHistoryJobResponse
func (client *Client) DeleteDeliveryHistoryJobWithContext(ctx context.Context, request *DeleteDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trail.
//
// Description:
//
// This topic describes how to delete a sample trail named `trail-test`.
//
// @param request - DeleteTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrailResponse
func (client *Client) DeleteTrailWithContext(ctx context.Context, request *DeleteTrailRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries advanced query templates.
//
// @param request - DescribeAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvancedQueryTemplateResponse
func (client *Client) DescribeAdvancedQueryTemplateWithContext(ctx context.Context, request *DescribeAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are supported by ActionTrail.
//
// Description:
//
// For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lifecycle events of a specified resource.
//
// @param request - DescribeResourceLifeCycleEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceLifeCycleEventsResponse
func (client *Client) DescribeResourceLifeCycleEventsWithContext(ctx context.Context, request *DescribeResourceLifeCycleEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceLifeCycleEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceLifeCycleEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceLifeCycleEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all advanced query scenarios.
//
// @param request - DescribeScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScenesResponse
func (client *Client) DescribeScenesWithContext(ctx context.Context, request *DescribeScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchCode) {
		query["SearchCode"] = request.SearchCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScenes"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries advanced query templates for a specified scenario.
//
// @param request - DescribeSearchTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSearchTemplatesResponse
func (client *Client) DescribeSearchTemplatesWithContext(ctx context.Context, request *DescribeSearchTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSearchTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSearchTemplates"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSearchTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves data for delivery monitoring metrics.
//
// @param request - DescribeTrailDeliveryMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrailDeliveryMetricDataResponse
func (client *Client) DescribeTrailDeliveryMetricDataWithContext(ctx context.Context, request *DescribeTrailDeliveryMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrailDeliveryMetricDataResponse, _err error) {
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
		Action:      dara.String("DescribeTrailDeliveryMetricData"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrailDeliveryMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created trails.
//
// Description:
//
// This topic shows you how to query the information about the single-account trails within an Alibaba Cloud account. In this example, the information about a trail named `test-4` is returned.
//
// @param request - DescribeTrailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrailsResponse
func (client *Client) DescribeTrailsWithContext(ctx context.Context, request *DescribeTrailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeOrganizationTrail) {
		query["IncludeOrganizationTrail"] = request.IncludeOrganizationTrail
	}

	if !dara.IsNil(request.IncludeShadowTrails) {
		query["IncludeShadowTrails"] = request.IncludeShadowTrails
	}

	if !dara.IsNil(request.NameList) {
		query["NameList"] = request.NameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrails"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of daily alerts within a specific time range.
//
// @param request - DescribeUserAlertCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAlertCountResponse
func (client *Client) DescribeUserAlertCountWithContext(ctx context.Context, request *DescribeUserAlertCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAlertCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAlertCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAlertCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of daily logs within a specific time range.
//
// @param request - DescribeUserLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogCountResponse
func (client *Client) DescribeUserLogCountWithContext(ctx context.Context, request *DescribeUserLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserLogCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserLogCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserLogCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of enabled trails, including organization trails.
//
// @param request - DescribeUserTrailCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTrailCountResponse
func (client *Client) DescribeUserTrailCountWithContext(ctx context.Context, request *DescribeUserTrailCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserTrailCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserTrailCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserTrailCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specific type of Insights event.
//
// @param request - DisableInsightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInsightResponse
func (client *Client) DisableInsightWithContext(ctx context.Context, request *DisableInsightRequest, runtime *dara.RuntimeOptions) (_result *DisableInsightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsightType) {
		query["InsightType"] = request.InsightType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInsight"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInsightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Insights feature.
//
// @param request - EnableInsightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInsightResponse
func (client *Client) EnableInsightWithContext(ctx context.Context, request *EnableInsightRequest, runtime *dara.RuntimeOptions) (_result *EnableInsightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsightType) {
		query["InsightType"] = request.InsightType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInsight"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInsightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent events associated with a specified AccessKey pair, including the event name, source, timestamp, and details.
//
// Description:
//
// You can call this operation to query only the information about the most recent events that are generated within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. For more information about supported events, see [Alibaba Cloud services and events that are supported by the AccessKey pair audit feature](https://help.aliyun.com/document_detail/419214.html). Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedEventsResponse
func (client *Client) GetAccessKeyLastUsedEventsWithContext(ctx context.Context, request *GetAccessKeyLastUsedEventsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent usage record of a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about the most recent call of a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedInfoResponse
func (client *Client) GetAccessKeyLastUsedInfoWithContext(ctx context.Context, request *GetAccessKeyLastUsedInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedInfo"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses most recently used by a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about the IP addresses that are most recently used within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedIpsResponse
func (client *Client) GetAccessKeyLastUsedIpsWithContext(ctx context.Context, request *GetAccessKeyLastUsedIpsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedIps"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud services most recently accessed by a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedProductsResponse
func (client *Client) GetAccessKeyLastUsedProductsWithContext(ctx context.Context, request *GetAccessKeyLastUsedProductsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedProducts"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources most recently used by a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about resources that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResourcesResponse
func (client *Client) GetAccessKeyLastUsedResourcesWithContext(ctx context.Context, request *GetAccessKeyLastUsedResourcesRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedResources"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a single advanced template.
//
// @param request - GetAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvancedQueryTemplateResponse
func (client *Client) GetAdvancedQueryTemplateWithContext(ctx context.Context, request *GetAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the data event selector for a specified trail.
//
// @param request - GetDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataEventSelectorResponse
func (client *Client) GetDataEventSelectorWithContext(ctx context.Context, request *GetDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *GetDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataEventSelectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data backfill task.
//
// Description:
//
// This topic provides an example on how to query the details of a data backfill task whose ID is `16602`. The return result shows that historical events for a trail named `trail-name` are delivered to Simple Log Service and the task is complete.
//
// @param request - GetDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryHistoryJobResponse
func (client *Client) GetDeliveryHistoryJobWithContext(ctx context.Context, request *GetDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Insights event types to deliver for a trail.
//
// @param request - GetInsightSelectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInsightSelectorsResponse
func (client *Client) GetInsightSelectorsWithContext(ctx context.Context, request *GetInsightSelectorsRequest, runtime *dara.RuntimeOptions) (_result *GetInsightSelectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInsightSelectors"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInsightSelectorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of Insights events for the current account.
//
// @param request - GetInsightsEventsCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInsightsEventsCountResponse
func (client *Client) GetInsightsEventsCountWithContext(ctx context.Context, request *GetInsightsEventsCountRequest, runtime *dara.RuntimeOptions) (_result *GetInsightsEventsCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInsightsEventsCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInsightsEventsCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a trail.
//
// Description:
//
// This topic describes how to query the status of a sample single-account trail named `trail-test`.
//
// @param request - GetTrailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrailStatusResponse
func (client *Client) GetTrailStatusWithContext(ctx context.Context, request *GetTrailStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTrailStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsOrganizationTrail) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrailStatus"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all data event selectors.
//
// @param request - ListDataEventSelectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataEventSelectorsResponse
func (client *Client) ListDataEventSelectorsWithContext(ctx context.Context, request *ListDataEventSelectorsRequest, runtime *dara.RuntimeOptions) (_result *ListDataEventSelectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataEventSelectors"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataEventSelectorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services that support data events and the names of these events.
//
// @param request - ListDataEventServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataEventServicesResponse
func (client *Client) ListDataEventServicesWithContext(ctx context.Context, request *ListDataEventServicesRequest, runtime *dara.RuntimeOptions) (_result *ListDataEventServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataEventServices"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataEventServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data backfill tasks.
//
// Description:
//
// This topic provides an example of how to query a list of data backfill tasks. The response shows a task with the ID `16602` that delivers historical events from the trail `trail-name` to Simple Log Service (SLS).
//
// @param request - ListDeliveryHistoryJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryHistoryJobsResponse
func (client *Client) ListDeliveryHistoryJobsWithContext(ctx context.Context, request *ListDeliveryHistoryJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListDeliveryHistoryJobs"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed historical events.
//
// Description:
//
// > Do not call this operation frequently. To query events in near-real time, you can create a trail to deliver events to Simple Log Service (SLS) and use its real-time consumption feature. For more information, see [Create a single-account trail](https://help.aliyun.com/document_detail/28810.html), [Create a multi-account trail](https://help.aliyun.com/document_detail/160661.html), and [Real-time consumption](https://help.aliyun.com/document_detail/28997.html).
//
// @param request - LookupEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupEventsResponse
func (client *Client) LookupEventsWithContext(ctx context.Context, request *LookupEventsRequest, runtime *dara.RuntimeOptions) (_result *LookupEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LookupAttribute) {
		query["LookupAttribute"] = request.LookupAttribute
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Insights events.
//
// @param request - LookupInsightEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupInsightEventsResponse
func (client *Client) LookupInsightEventsWithContext(ctx context.Context, request *LookupInsightEventsRequest, runtime *dara.RuntimeOptions) (_result *LookupInsightEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LookupAttribute) {
		query["LookupAttribute"] = request.LookupAttribute
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupInsightEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupInsightEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or configures a data event selector. A trail must exist before you create a data event selector. If a trail does not exist, you can call the CreateTrail operation to create one.
//
// @param request - PutDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDataEventSelectorResponse
func (client *Client) PutDataEventSelectorWithContext(ctx context.Context, request *PutDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *PutDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventSelectors) {
		query["EventSelectors"] = request.EventSelectors
	}

	if !dara.IsNil(request.IsTrailAllRegion) {
		query["IsTrailAllRegion"] = request.IsTrailAllRegion
	}

	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	if !dara.IsNil(request.TrailRegionIds) {
		query["TrailRegionIds"] = request.TrailRegionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDataEventSelectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the types of Insights events to deliver for a trail.
//
// @param request - PutInsightSelectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutInsightSelectorsResponse
func (client *Client) PutInsightSelectorsWithContext(ctx context.Context, request *PutInsightSelectorsRequest, runtime *dara.RuntimeOptions) (_result *PutInsightSelectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsightSelectors) {
		query["InsightSelectors"] = request.InsightSelectors
	}

	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutInsightSelectors"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutInsightSelectorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a trail to start delivering ActionTrail events to Object Storage Service (OSS), Simple Log Service (SLS), or MaxCompute.
//
// Description:
//
// This topic provides an example on how to enable a trail named `trail-test`.
//
// @param request - StartLoggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLoggingResponse
func (client *Client) StartLoggingWithContext(ctx context.Context, request *StartLoggingRequest, runtime *dara.RuntimeOptions) (_result *StartLoggingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLogging"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLoggingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a trail to stop delivering ActionTrail events to Object Storage Service (OSS), Simple Log Service (SLS), or MaxCompute.
//
// Description:
//
// This topic provides an example on how to disable a trail named `trail-test`.
//
// @param request - StopLoggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLoggingResponse
func (client *Client) StopLoggingWithContext(ctx context.Context, request *StopLoggingRequest, runtime *dara.RuntimeOptions) (_result *StopLoggingResponse, _err error) {
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
		Action:      dara.String("StopLogging"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLoggingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an advanced query template.
//
// @param request - UpdateAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdvancedQueryTemplateResponse
func (client *Client) UpdateAdvancedQueryTemplateWithContext(ctx context.Context, request *UpdateAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateSql) {
		query["TemplateSql"] = request.TemplateSql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the region where you want to store global events.
//
// Description:
//
// By default, global events are stored in the Singapore region.
//
//   - To obtain the permissions to call the API operation, you must submit a ticket.
//
//   - Only the China (Hangzhou) region (cn-hangzhou) and the Singapore region (ap-southeast-1) are supported.
//
// @param request - UpdateGlobalEventsStorageRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGlobalEventsStorageRegionResponse
func (client *Client) UpdateGlobalEventsStorageRegionWithContext(ctx context.Context, request *UpdateGlobalEventsStorageRegionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGlobalEventsStorageRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGlobalEventsStorageRegion"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGlobalEventsStorageRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a trail.
//
// Description:
//
// This topic shows you how to change the destination Object Storage Service (OSS) bucket of a sample trail named `trail-test` to `audit-log`.
//
// @param request - UpdateTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrailResponse
func (client *Client) UpdateTrailWithContext(ctx context.Context, request *UpdateTrailRequest, runtime *dara.RuntimeOptions) (_result *UpdateTrailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventRW) {
		query["EventRW"] = request.EventRW
	}

	if !dara.IsNil(request.MaxComputeProjectArn) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !dara.IsNil(request.MaxComputeWriteRoleArn) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssKeyPrefix) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !dara.IsNil(request.OssWriteRoleArn) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
	}

	if !dara.IsNil(request.SlsProjectArn) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !dara.IsNil(request.SlsWriteRoleArn) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !dara.IsNil(request.TrailRegion) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
