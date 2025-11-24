// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activate a Quota Plan Immediately.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyComputeQuotaPlanResponse
func (client *Client) ApplyComputeQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyComputeQuotaPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan/" + dara.PercentEncode(dara.StringValue(planName)) + "/apply"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create MaxCompute ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - CreateComputeQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeQuotaPlanResponse
func (client *Client) CreateComputeQuotaPlanWithContext(ctx context.Context, nickname *string, request *CreateComputeQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComputeQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateMmsDataSource
//
// @param request - CreateMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsDataSourceResponse
func (client *Client) CreateMmsDataSourceWithContext(ctx context.Context, request *CreateMmsDataSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMmsDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Networklink) {
		body["networklink"] = request.Networklink
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMmsDataSource"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMmsDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源的更新元数据操作
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsFetchMetadataJobResponse
func (client *Client) CreateMmsFetchMetadataJobWithContext(ctx context.Context, sourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMmsFetchMetadataJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMmsFetchMetadataJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/scans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMmsFetchMetadataJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建迁移任务
//
// @param request - CreateMmsJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMmsJobResponse
func (client *Client) CreateMmsJobWithContext(ctx context.Context, sourceId *string, request *CreateMmsJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMmsJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ColumnMapping) {
		body["columnMapping"] = request.ColumnMapping
	}

	if !dara.IsNil(request.DstDbName) {
		body["dstDbName"] = request.DstDbName
	}

	if !dara.IsNil(request.DstSchemaName) {
		body["dstSchemaName"] = request.DstSchemaName
	}

	if !dara.IsNil(request.EnableDataMigration) {
		body["enableDataMigration"] = request.EnableDataMigration
	}

	if !dara.IsNil(request.EnableSchemaMigration) {
		body["enableSchemaMigration"] = request.EnableSchemaMigration
	}

	if !dara.IsNil(request.EnableVerification) {
		body["enableVerification"] = request.EnableVerification
	}

	if !dara.IsNil(request.Eta) {
		body["eta"] = request.Eta
	}

	if !dara.IsNil(request.Increment) {
		body["increment"] = request.Increment
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Others) {
		body["others"] = request.Others
	}

	if !dara.IsNil(request.PartitionFilters) {
		body["partitionFilters"] = request.PartitionFilters
	}

	if !dara.IsNil(request.Partitions) {
		body["partitions"] = request.Partitions
	}

	if !dara.IsNil(request.SchemaOnly) {
		body["schemaOnly"] = request.SchemaOnly
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceName) {
		body["sourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SrcDbName) {
		body["srcDbName"] = request.SrcDbName
	}

	if !dara.IsNil(request.SrcSchemaName) {
		body["srcSchemaName"] = request.SrcSchemaName
	}

	if !dara.IsNil(request.TableBlackList) {
		body["tableBlackList"] = request.TableBlackList
	}

	if !dara.IsNil(request.TableMapping) {
		body["tableMapping"] = request.TableMapping
	}

	if !dara.IsNil(request.TableWhiteList) {
		body["tableWhiteList"] = request.TableWhiteList
	}

	if !dara.IsNil(request.Tables) {
		body["tables"] = request.Tables
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a package.
//
// @param request - CreatePackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePackageResponse
func (client *Client) CreatePackageWithContext(ctx context.Context, projectName *string, request *CreatePackageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsInstall) {
		query["isInstall"] = request.IsInstall
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePackage"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/packages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a MaxCompute project.
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quota plan.
//
// @param request - CreateQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaPlanResponse
func (client *Client) CreateQuotaPlanWithContext(ctx context.Context, nickname *string, request *CreateQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/plans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a role at the MaxCompute project level.
//
// @param request - CreateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithContext(ctx context.Context, projectName *string, request *CreateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRole"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete MaxCompute compute quota plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeQuotaPlanResponse
func (client *Client) DeleteComputeQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteComputeQuotaPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan/" + dara.PercentEncode(dara.StringValue(planName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmsDataSourceResponse
func (client *Client) DeleteMmsDataSourceWithContext(ctx context.Context, sourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMmsDataSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMmsDataSource"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMmsDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除迁移计划
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMmsJobResponse
func (client *Client) DeleteMmsJobWithContext(ctx context.Context, sourceId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMmsJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a MaxCompute project.
//
// @param request - DeleteProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, projectName *string, request *DeleteProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsLogical) {
		query["isLogical"] = request.IsLogical
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota plan.
//
// @param request - DeleteQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaPlanResponse
func (client *Client) DeleteQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, request *DeleteQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/plans/" + dara.PercentEncode(dara.StringValue(planName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetComputeEffectivePlan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeEffectivePlanResponse
func (client *Client) GetComputeEffectivePlanWithContext(ctx context.Context, nickname *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetComputeEffectivePlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeEffectivePlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeEffectivePlan/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeEffectivePlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get detailed information of a single compute quota plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeQuotaPlanResponse
func (client *Client) GetComputeQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetComputeQuotaPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan/" + dara.PercentEncode(dara.StringValue(planName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the time-specific configuration of compute quota.
//
// @param request - GetComputeQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeQuotaScheduleResponse
func (client *Client) GetComputeQuotaScheduleWithContext(ctx context.Context, nickname *string, request *GetComputeQuotaScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetComputeQuotaScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayTimezone) {
		query["displayTimezone"] = request.DisplayTimezone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeQuotaSchedule"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaSchedule"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeQuotaScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about a job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInfoResponse
func (client *Client) GetJobInfoWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobInfo"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/info"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs statistics on all jobs that are complete on a specified day and obtains the total resource usage of each job executor on a daily basis.
//
// @param tmpReq - GetJobResourceUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResourceUsageResponse
func (client *Client) GetJobResourceUsageWithContext(ctx context.Context, tmpReq *GetJobResourceUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResourceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetJobResourceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobOwnerList) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, dara.String("jobOwnerList"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QuotaNicknameList) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, dara.String("quotaNicknameList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.JobOwnerListShrink) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QuotaNicknameListShrink) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobResourceUsage"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/resourceUsage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResourceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsAsyncTask
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsAsyncTaskResponse
func (client *Client) GetMmsAsyncTaskWithContext(ctx context.Context, sourceId *string, asyncTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsAsyncTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsAsyncTask"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/asyncTasks/" + dara.PercentEncode(dara.StringValue(asyncTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源
//
// @param request - GetMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsDataSourceResponse
func (client *Client) GetMmsDataSourceWithContext(ctx context.Context, sourceId *string, request *GetMmsDataSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	if !dara.IsNil(request.WithConfig) {
		query["withConfig"] = request.WithConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsDataSource"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsDb
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsDbResponse
func (client *Client) GetMmsDbWithContext(ctx context.Context, sourceId *string, dbId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsDbResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsDb"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/dbs/" + dara.PercentEncode(dara.StringValue(dbId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsDbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsFetchMetadataJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsFetchMetadataJobResponse
func (client *Client) GetMmsFetchMetadataJobWithContext(ctx context.Context, sourceId *string, scanId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsFetchMetadataJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsFetchMetadataJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/scans/" + dara.PercentEncode(dara.StringValue(scanId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsFetchMetadataJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取迁移计划
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsJobResponse
func (client *Client) GetMmsJobWithContext(ctx context.Context, sourceId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsPartition
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsPartitionResponse
func (client *Client) GetMmsPartitionWithContext(ctx context.Context, sourceId *string, partitionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsPartitionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsPartition"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/partitions/" + dara.PercentEncode(dara.StringValue(partitionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsPartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsTable
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsTableResponse
func (client *Client) GetMmsTableWithContext(ctx context.Context, sourceId *string, tableId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsTable"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMmsTask
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMmsTaskResponse
func (client *Client) GetMmsTaskWithContext(ctx context.Context, sourceId *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMmsTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMmsTask"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMmsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a package.
//
// @param request - GetPackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPackageResponse
func (client *Client) GetPackageWithContext(ctx context.Context, projectName *string, packageName *string, request *GetPackageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceProject) {
		query["sourceProject"] = request.SourceProject
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPackage"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/packages/" + dara.PercentEncode(dara.StringValue(packageName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a MaxCompute project.
//
// @param request - GetProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, projectName *string, request *GetProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WithQuotaProductType) {
		query["withQuotaProductType"] = request.WithQuotaProductType
	}

	if !dara.IsNil(request.WithStorageTierInfo) {
		query["withStorageTierInfo"] = request.WithStorageTierInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a specified level-1 quota.
//
// @param request - GetQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithContext(ctx context.Context, nickname *string, request *GetQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AkProven) {
		query["AkProven"] = request.AkProven
	}

	if !dara.IsNil(request.Mock) {
		query["mock"] = request.Mock
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuota"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information of a quota plan.
//
// @param request - GetQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaPlanResponse
func (client *Client) GetQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, request *GetQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/plans/" + dara.PercentEncode(dara.StringValue(planName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the scheduling plan for a quota plan.
//
// @param request - GetQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaScheduleResponse
func (client *Client) GetQuotaScheduleWithContext(ctx context.Context, nickname *string, request *GetQuotaScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayTimezone) {
		query["displayTimezone"] = request.DisplayTimezone
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaSchedule"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/schedule"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota resource consumption information.
//
// @param tmpReq - GetQuotaUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaUsageResponse
func (client *Client) GetQuotaUsageWithContext(ctx context.Context, nickname *string, tmpReq *GetQuotaUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetQuotaUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PlotTypes) {
		request.PlotTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlotTypes, dara.String("plotTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.YAxisTypes) {
		request.YAxisTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.YAxisTypes, dara.String("yAxisTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggMethod) {
		query["aggMethod"] = request.AggMethod
	}

	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.PlotTypesShrink) {
		query["plotTypes"] = request.PlotTypesShrink
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.SubQuotaNickname) {
		query["subQuotaNickname"] = request.SubQuotaNickname
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	if !dara.IsNil(request.YAxisTypesShrink) {
		query["yAxisTypes"] = request.YAxisTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaUsage"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/usage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the ACL-based permissions that is granted to a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleAclResponse
func (client *Client) GetRoleAclWithContext(ctx context.Context, projectName *string, roleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRoleAclResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoleAcl"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles/" + dara.PercentEncode(dara.StringValue(roleName)) + "/roleAcl"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains ACL-based permissions on an object that are granted to a project-level role.
//
// @param request - GetRoleAclOnObjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleAclOnObjectResponse
func (client *Client) GetRoleAclOnObjectWithContext(ctx context.Context, projectName *string, roleName *string, request *GetRoleAclOnObjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRoleAclOnObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectName) {
		query["objectName"] = request.ObjectName
	}

	if !dara.IsNil(request.ObjectType) {
		query["objectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRoleAclOnObject"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles/" + dara.PercentEncode(dara.StringValue(roleName)) + "/acl"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleAclOnObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the policy that is attached to a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRolePolicyResponse
func (client *Client) GetRolePolicyWithContext(ctx context.Context, projectName *string, roleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRolePolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRolePolicy"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles/" + dara.PercentEncode(dara.StringValue(roleName)) + "/policy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRolePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the running state data of jobs that are in the running state in a specified period of time.
//
// @param tmpReq - GetRunningJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunningJobsResponse
func (client *Client) GetRunningJobsWithContext(ctx context.Context, tmpReq *GetRunningJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRunningJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetRunningJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobOwnerList) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, dara.String("jobOwnerList"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QuotaNicknameList) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, dara.String("quotaNicknameList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.JobOwnerListShrink) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QuotaNicknameListShrink) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRunningJobs"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/runningJobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRunningJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetStorageAmountSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageAmountSummaryResponse
func (client *Client) GetStorageAmountSummaryWithContext(ctx context.Context, request *GetStorageAmountSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStorageAmountSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageAmountSummary"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/amount"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageAmountSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetStorageSizeSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageSizeSummaryResponse
func (client *Client) GetStorageSizeSummaryWithContext(ctx context.Context, request *GetStorageSizeSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStorageSizeSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageSizeSummary"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/size"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageSizeSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - GetStorageSummaryComparedRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageSummaryComparedResponse
func (client *Client) GetStorageSummaryComparedWithContext(ctx context.Context, _type *string, tmpReq *GetStorageSummaryComparedRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStorageSummaryComparedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStorageSummaryComparedShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Projects) {
		request.ProjectsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Projects, dara.String("projects"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginDate) {
		query["beginDate"] = request.BeginDate
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProjectsShrink) {
		query["projects"] = request.ProjectsShrink
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageSummaryCompared"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/" + dara.PercentEncode(dara.StringValue(_type)) + "/compared"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageSummaryComparedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views the information about MaxCompute internal tables, views, external tables, clustered tables, or transactional tables.
//
// @param request - GetTableInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableInfoResponse
func (client *Client) GetTableInfoWithContext(ctx context.Context, projectName *string, tableName *string, request *GetTableInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SchemaName) {
		query["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableInfo"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the trusted projects of the current project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrustedProjectsResponse
func (client *Client) GetTrustedProjectsWithContext(ctx context.Context, projectName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrustedProjectsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrustedProjects"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/trustedProjects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrustedProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a running job.
//
// @param request - KillJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillJobsResponse
func (client *Client) KillJobsWithContext(ctx context.Context, request *KillJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *KillJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillJobs"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/kill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &KillJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get compute usage of pay-as-you-go jobs.
//
// @param request - ListComputeMetricsByInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeMetricsByInstanceResponse
func (client *Client) ListComputeMetricsByInstanceWithContext(ctx context.Context, request *ListComputeMetricsByInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComputeMetricsByInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobOwner) {
		body["jobOwner"] = request.JobOwner
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectNames) {
		body["projectNames"] = request.ProjectNames
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.Signature) {
		body["signature"] = request.Signature
	}

	if !dara.IsNil(request.SpecCodes) {
		body["specCodes"] = request.SpecCodes
	}

	if !dara.IsNil(request.StartDate) {
		body["startDate"] = request.StartDate
	}

	if !dara.IsNil(request.Types) {
		body["types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeMetricsByInstance"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/computeMetrics/listByInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeMetricsByInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get computeQuotaPlan list.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeQuotaPlanResponse
func (client *Client) ListComputeQuotaPlanWithContext(ctx context.Context, nickname *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComputeQuotaPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains functions in a MaxCompute project.
//
// @param request - ListFunctionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithContext(ctx context.Context, projectName *string, request *ListFunctionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItem) {
		query["maxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.SchemaName) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctions"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/functions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views a list of jobs.
//
// @param request - ListJobInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobInfosResponse
func (client *Client) ListJobInfosWithContext(ctx context.Context, request *ListJobInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		query["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.OrderColumn) {
		query["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtNodeIdList) {
		body["extNodeIdList"] = request.ExtNodeIdList
	}

	if !dara.IsNil(request.From) {
		body["from"] = request.From
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.JobOwnerList) {
		body["jobOwnerList"] = request.JobOwnerList
	}

	if !dara.IsNil(request.PriorityList) {
		body["priorityList"] = request.PriorityList
	}

	if !dara.IsNil(request.ProjectList) {
		body["projectList"] = request.ProjectList
	}

	if !dara.IsNil(request.QuotaNickname) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !dara.IsNil(request.SceneTagList) {
		body["sceneTagList"] = request.SceneTagList
	}

	if !dara.IsNil(request.SignatureList) {
		body["signatureList"] = request.SignatureList
	}

	if !dara.IsNil(request.SortByList) {
		body["sortByList"] = request.SortByList
	}

	if !dara.IsNil(request.SortOrderList) {
		body["sortOrderList"] = request.SortOrderList
	}

	if !dara.IsNil(request.StatusList) {
		body["statusList"] = request.StatusList
	}

	if !dara.IsNil(request.To) {
		body["to"] = request.To
	}

	if !dara.IsNil(request.TypeList) {
		body["typeList"] = request.TypeList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobInfos"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve performance metrics for completed jobs.
//
// @param request - ListJobMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobMetricResponse
func (client *Client) ListJobMetricWithContext(ctx context.Context, request *ListJobMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Group) {
		body["group"] = request.Group
	}

	if !dara.IsNil(request.Metric) {
		body["metric"] = request.Metric
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.Project) {
		body["project"] = request.Project
	}

	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobMetric"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/metric"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views a list of job snapshot data at a specific point in time.
//
// @param request - ListJobSnapshotInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobSnapshotInfosResponse
func (client *Client) ListJobSnapshotInfosWithContext(ctx context.Context, request *ListJobSnapshotInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobSnapshotInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		query["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.OrderColumn) {
		query["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtNodeIdList) {
		body["extNodeIdList"] = request.ExtNodeIdList
	}

	if !dara.IsNil(request.From) {
		body["from"] = request.From
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["instanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.JobOwnerList) {
		body["jobOwnerList"] = request.JobOwnerList
	}

	if !dara.IsNil(request.PriorityList) {
		body["priorityList"] = request.PriorityList
	}

	if !dara.IsNil(request.ProjectList) {
		body["projectList"] = request.ProjectList
	}

	if !dara.IsNil(request.QuotaNickname) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !dara.IsNil(request.SignatureList) {
		body["signatureList"] = request.SignatureList
	}

	if !dara.IsNil(request.SortByList) {
		body["sortByList"] = request.SortByList
	}

	if !dara.IsNil(request.SortOrderList) {
		body["sortOrderList"] = request.SortOrderList
	}

	if !dara.IsNil(request.StatusList) {
		body["statusList"] = request.StatusList
	}

	if !dara.IsNil(request.To) {
		body["to"] = request.To
	}

	if !dara.IsNil(request.TypeList) {
		body["typeList"] = request.TypeList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobSnapshotInfos"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/snapshot"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobSnapshotInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMmsDataSources
//
// @param request - ListMmsDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsDataSourcesResponse
func (client *Client) ListMmsDataSourcesWithContext(ctx context.Context, request *ListMmsDataSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsDataSources"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsDataSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个数据源内“库”列表
//
// @param tmpReq - ListMmsDbsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsDbsResponse
func (client *Client) ListMmsDbsWithContext(ctx context.Context, sourceId *string, tmpReq *ListMmsDbsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsDbsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMmsDbsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sorter) {
		request.SorterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sorter, dara.String("sorter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SorterShrink) {
		query["sorter"] = request.SorterShrink
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsDbs"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/dbs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsDbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMmsJobs
//
// @param request - ListMmsJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsJobsResponse
func (client *Client) ListMmsJobsWithContext(ctx context.Context, sourceId *string, request *ListMmsJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstDbName) {
		query["dstDbName"] = request.DstDbName
	}

	if !dara.IsNil(request.DstTableName) {
		query["dstTableName"] = request.DstTableName
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcDbName) {
		query["srcDbName"] = request.SrcDbName
	}

	if !dara.IsNil(request.SrcTableName) {
		query["srcTableName"] = request.SrcTableName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Stopped) {
		query["stopped"] = request.Stopped
	}

	if !dara.IsNil(request.TimerId) {
		query["timerId"] = request.TimerId
	}

	if !dara.IsNil(request.Sorter) {
		query["sorter"] = request.Sorter
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsJobs"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取元数据-分区
//
// @param tmpReq - ListMmsPartitionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsPartitionsResponse
func (client *Client) ListMmsPartitionsWithContext(ctx context.Context, sourceId *string, tmpReq *ListMmsPartitionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsPartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMmsPartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Status) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, dara.String("status"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["dbId"] = request.DbId
	}

	if !dara.IsNil(request.DbName) {
		query["dbName"] = request.DbName
	}

	if !dara.IsNil(request.LastDdlTimeEnd) {
		query["lastDdlTimeEnd"] = request.LastDdlTimeEnd
	}

	if !dara.IsNil(request.LastDdlTimeStart) {
		query["lastDdlTimeStart"] = request.LastDdlTimeStart
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusShrink) {
		query["status"] = request.StatusShrink
	}

	if !dara.IsNil(request.TableId) {
		query["tableId"] = request.TableId
	}

	if !dara.IsNil(request.TableName) {
		query["tableName"] = request.TableName
	}

	if !dara.IsNil(request.Updated) {
		query["updated"] = request.Updated
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	if !dara.IsNil(request.Sorter) {
		query["sorter"] = request.Sorter
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsPartitions"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/partitions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsPartitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMmsTables
//
// @param tmpReq - ListMmsTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTablesResponse
func (client *Client) ListMmsTablesWithContext(ctx context.Context, sourceId *string, tmpReq *ListMmsTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMmsTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Status) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, dara.String("status"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["dbId"] = request.DbId
	}

	if !dara.IsNil(request.DbName) {
		query["dbName"] = request.DbName
	}

	if !dara.IsNil(request.DstName) {
		query["dstName"] = request.DstName
	}

	if !dara.IsNil(request.DstProjectName) {
		query["dstProjectName"] = request.DstProjectName
	}

	if !dara.IsNil(request.DstSchemaName) {
		query["dstSchemaName"] = request.DstSchemaName
	}

	if !dara.IsNil(request.HasPartitions) {
		query["hasPartitions"] = request.HasPartitions
	}

	if !dara.IsNil(request.LastDdlTimeEnd) {
		query["lastDdlTimeEnd"] = request.LastDdlTimeEnd
	}

	if !dara.IsNil(request.LastDdlTimeStart) {
		query["lastDdlTimeStart"] = request.LastDdlTimeStart
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OnlyName) {
		query["onlyName"] = request.OnlyName
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusShrink) {
		query["status"] = request.StatusShrink
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.Sorter) {
		query["sorter"] = request.Sorter
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsTables"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMmsTaskLogs
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTaskLogsResponse
func (client *Client) ListMmsTaskLogsWithContext(ctx context.Context, sourceId *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsTaskLogsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsTaskLogs"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsTaskLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMmsTasks
//
// @param request - ListMmsTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMmsTasksResponse
func (client *Client) ListMmsTasksWithContext(ctx context.Context, sourceId *string, request *ListMmsTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMmsTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstDbName) {
		query["dstDbName"] = request.DstDbName
	}

	if !dara.IsNil(request.DstTableName) {
		query["dstTableName"] = request.DstTableName
	}

	if !dara.IsNil(request.JobId) {
		query["jobId"] = request.JobId
	}

	if !dara.IsNil(request.JobName) {
		query["jobName"] = request.JobName
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Partition) {
		query["partition"] = request.Partition
	}

	if !dara.IsNil(request.SrcDbName) {
		query["srcDbName"] = request.SrcDbName
	}

	if !dara.IsNil(request.SrcTableName) {
		query["srcTableName"] = request.SrcTableName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Sorter) {
		query["sorter"] = request.Sorter
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMmsTasks"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMmsTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the packages in a MaxCompute project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPackagesResponse
func (client *Client) ListPackagesWithContext(ctx context.Context, projectName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPackagesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPackages"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/packages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users in a project.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectUsersResponse
func (client *Client) ListProjectUsersWithContext(ctx context.Context, projectName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectUsersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectUsers"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of MaxCompute projects.
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, request *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListSystemCatalog) {
		query["listSystemCatalog"] = request.ListSystemCatalog
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItem) {
		query["maxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.QuotaName) {
		query["quotaName"] = request.QuotaName
	}

	if !dara.IsNil(request.QuotaNickName) {
		query["quotaNickName"] = request.QuotaNickName
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.SaleTags) {
		query["saleTags"] = request.SaleTags
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quotas.
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithContext(ctx context.Context, request *ListQuotasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingType) {
		query["billingType"] = request.BillingType
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItem) {
		query["maxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.SaleTags) {
		query["saleTags"] = request.SaleTags
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotas"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains quota plans.
//
// @param request - ListQuotasPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasPlansResponse
func (client *Client) ListQuotasPlansWithContext(ctx context.Context, nickname *string, request *ListQuotasPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotasPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotasPlans"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/plans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotasPlansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains resources in a MaxCompute project.
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, projectName *string, request *ListResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItem) {
		query["maxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.SchemaName) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains MaxCompute project-level roles.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, projectName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage details of a specific partition in a partitioned table in a MaxCompute project.
//
// @param tmpReq - ListStoragePartitionsInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStoragePartitionsInfoResponse
func (client *Client) ListStoragePartitionsInfoWithContext(ctx context.Context, project *string, table *string, tmpReq *ListStoragePartitionsInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStoragePartitionsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListStoragePartitionsInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		query["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.OrderColumn) {
		query["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PartitionPrefix) {
		query["partitionPrefix"] = request.PartitionPrefix
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Schema) {
		query["schema"] = request.Schema
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.TypesShrink) {
		query["types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStoragePartitionsInfo"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/projects/" + dara.PercentEncode(dara.StringValue(project)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/partitionsInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStoragePartitionsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListStorageProjectsInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStorageProjectsInfoResponse
func (client *Client) ListStorageProjectsInfoWithContext(ctx context.Context, request *ListStorageProjectsInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStorageProjectsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		query["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.OrderColumn) {
		query["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectPrefix) {
		query["projectPrefix"] = request.ProjectPrefix
	}

	if !dara.IsNil(request.RecentDays) {
		query["recentDays"] = request.RecentDays
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStorageProjectsInfo"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/projectsInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStorageProjectsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the table storage details of a MaxCompute project.
//
// @param tmpReq - ListStorageTablesInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStorageTablesInfoResponse
func (client *Client) ListStorageTablesInfoWithContext(ctx context.Context, project *string, tmpReq *ListStorageTablesInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStorageTablesInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListStorageTablesInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("types"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		query["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	if !dara.IsNil(request.OrderColumn) {
		query["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecentDays) {
		query["recentDays"] = request.RecentDays
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Schema) {
		query["schema"] = request.Schema
	}

	if !dara.IsNil(request.TablePrefix) {
		query["tablePrefix"] = request.TablePrefix
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	if !dara.IsNil(request.TypesShrink) {
		query["types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStorageTablesInfo"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/analysis/storage/projects/" + dara.PercentEncode(dara.StringValue(project)) + "/tablesInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStorageTablesInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains tables in a MaxCompute project.
//
// @param request - ListTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithContext(ctx context.Context, projectName *string, request *ListTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItem) {
		query["maxItem"] = request.MaxItem
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	if !dara.IsNil(request.SchemaName) {
		query["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTunnelQuotaTimerResponse
func (client *Client) ListTunnelQuotaTimerWithContext(ctx context.Context, nickname *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTunnelQuotaTimerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTunnelQuotaTimer"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tunnel/" + dara.PercentEncode(dara.StringValue(nickname)) + "/timers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTunnelQuotaTimerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of all users under a tenant.
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about the users who are assigned a project-level role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersByRoleResponse
func (client *Client) ListUsersByRoleWithContext(ctx context.Context, projectName *string, roleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersByRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersByRole"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles/" + dara.PercentEncode(dara.StringValue(roleName)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersByRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specified level-1 quota group.
//
// @param request - QueryQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQuotaResponse
func (client *Client) QueryQuotaWithContext(ctx context.Context, nickname *string, request *QueryQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AkProven) {
		query["AkProven"] = request.AkProven
	}

	if !dara.IsNil(request.Mock) {
		query["mock"] = request.Mock
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryQuota"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询quota的资源使用信息
//
// @param request - QueryQuotaMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQuotaMetricResponse
func (client *Client) QueryQuotaMetricWithContext(ctx context.Context, metric *string, request *QueryQuotaMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryQuotaMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Strategy) {
		query["strategy"] = request.Strategy
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Interval) {
		body["interval"] = request.Interval
	}

	if !dara.IsNil(request.Nickname) {
		body["nickname"] = request.Nickname
	}

	if !dara.IsNil(request.SubQuotaNickname) {
		body["subQuotaNickname"] = request.SubQuotaNickname
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryQuotaMetric"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/quota/" + dara.PercentEncode(dara.StringValue(metric))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryQuotaMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看存储数据的时序指标
//
// @param request - QueryStorageMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStorageMetricResponse
func (client *Client) QueryStorageMetricWithContext(ctx context.Context, metric *string, request *QueryStorageMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryStorageMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectList) {
		body["projectList"] = request.ProjectList
	}

	if !dara.IsNil(request.TypeList) {
		body["typeList"] = request.TypeList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryStorageMetric"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/storage/" + dara.PercentEncode(dara.StringValue(metric))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryStorageMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询tunnel资源使用信息
//
// @param request - QueryTunnelMetricRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTunnelMetricResponse
func (client *Client) QueryTunnelMetricWithContext(ctx context.Context, metric *string, request *QueryTunnelMetricRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTunnelMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Strategy) {
		query["strategy"] = request.Strategy
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeList) {
		body["codeList"] = request.CodeList
	}

	if !dara.IsNil(request.GroupList) {
		body["groupList"] = request.GroupList
	}

	if !dara.IsNil(request.OperationList) {
		body["operationList"] = request.OperationList
	}

	if !dara.IsNil(request.Project) {
		body["project"] = request.Project
	}

	if !dara.IsNil(request.QuotaNickname) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !dara.IsNil(request.TableList) {
		body["tableList"] = request.TableList
	}

	if !dara.IsNil(request.TopN) {
		body["topN"] = request.TopN
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTunnelMetric"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/tunnel/" + dara.PercentEncode(dara.StringValue(metric))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTunnelMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询tunnel资源使用详情
//
// @param request - QueryTunnelMetricDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTunnelMetricDetailResponse
func (client *Client) QueryTunnelMetricDetailWithContext(ctx context.Context, metric *string, request *QueryTunnelMetricDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTunnelMetricDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AscOrder) {
		body["ascOrder"] = request.AscOrder
	}

	if !dara.IsNil(request.GroupList) {
		body["groupList"] = request.GroupList
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.OperationList) {
		body["operationList"] = request.OperationList
	}

	if !dara.IsNil(request.OrderColumn) {
		body["orderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.Project) {
		body["project"] = request.Project
	}

	if !dara.IsNil(request.QuotaNickname) {
		body["quotaNickname"] = request.QuotaNickname
	}

	if !dara.IsNil(request.TableList) {
		body["tableList"] = request.TableList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTunnelMetricDetail"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/observations/tunnel/" + dara.PercentEncode(dara.StringValue(metric)) + "/detail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTunnelMetricDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RetryMmsJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryMmsJobResponse
func (client *Client) RetryMmsJobWithContext(ctx context.Context, sourceId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryMmsJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/retry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartMmsJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMmsJobResponse
func (client *Client) StartMmsJobWithContext(ctx context.Context, sourceId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartMmsJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StopMmsJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMmsJobResponse
func (client *Client) StopMmsJobWithContext(ctx context.Context, sourceId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopMmsJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMmsJob"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMmsJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SumStorageMetricsByDateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SumStorageMetricsByDateResponse
func (client *Client) SumStorageMetricsByDateWithContext(ctx context.Context, request *SumStorageMetricsByDateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SumStorageMetricsByDateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProjectNames) {
		body["projectNames"] = request.ProjectNames
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.StartDate) {
		body["startDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatsType) {
		body["statsType"] = request.StatsType
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SumStorageMetricsByDate"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/storageMetrics/sumByDate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SumStorageMetricsByDateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the ComputeQuotaPlan.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the <props="china">[Pricing and Charges](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Charges](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeQuotaPlanResponse
func (client *Client) UpdateComputeQuotaPlanWithContext(ctx context.Context, nickname *string, request *UpdateComputeQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComputeQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaPlan"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the time-based plan for computing quota.
//
// Description:
//
// Please ensure that before using this interface, you have fully understood the<props="china">[Pricing and Billing](https://help.aliyun.com/zh/maxcompute/product-overview/computing-pricing-1)
//
// <props="intl">[Pricing and Billing](https://www.alibabacloud.com/help/maxcompute/product-overview/computing-pricing-1) of MaxCompute Elastic Reserved CU.
//
// @param request - UpdateComputeQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeQuotaScheduleResponse
func (client *Client) UpdateComputeQuotaScheduleWithContext(ctx context.Context, nickname *string, request *UpdateComputeQuotaScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComputeQuotaScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScheduleTimezone) {
		query["scheduleTimezone"] = request.ScheduleTimezone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeQuotaSchedule"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeQuotaSchedule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeQuotaScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the basic configuration of the calculation quota, including adding or deleting the sub quota, modifying the basic properties of the secondary quota, and the CU configuration of the effective quota plan.
//
// @param request - UpdateComputeSubQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSubQuotaResponse
func (client *Client) UpdateComputeSubQuotaWithContext(ctx context.Context, nickname *string, request *UpdateComputeSubQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComputeSubQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SubQuotaInfoList) {
		body["subQuotaInfoList"] = request.SubQuotaInfoList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeSubQuota"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/computeSubQuota"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeSubQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据源配置、名称，启/停数据源实例
//
// @param request - UpdateMmsDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMmsDataSourceResponse
func (client *Client) UpdateMmsDataSourceWithContext(ctx context.Context, sourceId *string, request *UpdateMmsDataSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMmsDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		body["action"] = request.Action
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Test) {
		body["test"] = request.Test
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMmsDataSource"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/mms/datasources/" + dara.PercentEncode(dara.StringValue(sourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMmsDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the objects in a package and projects in which the package can be installed.
//
// @param request - UpdatePackageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePackageResponse
func (client *Client) UpdatePackageWithContext(ctx context.Context, projectName *string, packageName *string, request *UpdatePackageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePackage"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/packages/" + dara.PercentEncode(dara.StringValue(packageName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Project Basic Information
//
// @param request - UpdateProjectBasicMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectBasicMetaResponse
func (client *Client) UpdateProjectBasicMetaWithContext(ctx context.Context, projectName *string, request *UpdateProjectBasicMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectBasicMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.Properties) {
		body["properties"] = request.Properties
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectBasicMeta"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/meta"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectBasicMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Default Project Compute Quota
//
// @param request - UpdateProjectDefaultQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectDefaultQuotaResponse
func (client *Client) UpdateProjectDefaultQuotaWithContext(ctx context.Context, projectName *string, request *UpdateProjectDefaultQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectDefaultQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectDefaultQuota"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/quota"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectDefaultQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of a MaxCompute project.
//
// @param request - UpdateProjectIpWhiteListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectIpWhiteListResponse
func (client *Client) UpdateProjectIpWhiteListWithContext(ctx context.Context, projectName *string, request *UpdateProjectIpWhiteListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectIpWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectIpWhiteList"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/ipWhiteList"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectIpWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将project的二层模型升级为三层模型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectModelTierResponse
func (client *Client) UpdateProjectModelTierWithContext(ctx context.Context, projectName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectModelTierResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectModelTier"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/modelTier"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectModelTierResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a quota plan.
//
// @param request - UpdateQuotaPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaPlanResponse
func (client *Client) UpdateQuotaPlanWithContext(ctx context.Context, nickname *string, planName *string, request *UpdateQuotaPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQuotaPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQuotaPlan"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/plans/" + dara.PercentEncode(dara.StringValue(planName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQuotaPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the scheduling plan for a quota plan.
//
// @param request - UpdateQuotaScheduleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaScheduleResponse
func (client *Client) UpdateQuotaScheduleWithContext(ctx context.Context, nickname *string, request *UpdateQuotaScheduleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQuotaScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.TenantId) {
		query["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQuotaSchedule"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(nickname)) + "/schedule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQuotaScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the time-specific configuration of an exclusive resource group for Tunnel (referred to as Tunnel quota).
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing and prices](https://www.alibabacloud.com/help/maxcompute/product-overview/data-transfer-fees-hourly-billing) of Tunnel quotas and elastically reserved computing resources.
//
// @param request - UpdateTunnelQuotaTimerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTunnelQuotaTimerResponse
func (client *Client) UpdateTunnelQuotaTimerWithContext(ctx context.Context, nickname *string, request *UpdateTunnelQuotaTimerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTunnelQuotaTimerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Timezone) {
		query["timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTunnelQuotaTimer"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tunnel/" + dara.PercentEncode(dara.StringValue(nickname)) + "/timers"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTunnelQuotaTimerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add or remove users from a project role.
//
// @param request - UpdateUsersToRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUsersToRoleResponse
func (client *Client) UpdateUsersToRoleWithContext(ctx context.Context, projectName *string, roleName *string, request *UpdateUsersToRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUsersToRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Add) {
		body["add"] = request.Add
	}

	if !dara.IsNil(request.Remove) {
		body["remove"] = request.Remove
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUsersToRole"),
		Version:     dara.String("2022-01-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/projects/" + dara.PercentEncode(dara.StringValue(projectName)) + "/roles/" + dara.PercentEncode(dara.StringValue(roleName)) + "/users"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUsersToRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
