// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Triggers reindexing.
//
// Description:
//
// ## Method
//
//	POST
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/actions/build-index
//
// @param request - BuildIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildIndexResponse
func (client *Client) BuildIndexWithContext(ctx context.Context, instanceId *string, request *BuildIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BuildIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildMode) {
		body["buildMode"] = request.BuildMode
	}

	if !dara.IsNil(request.DataSourceName) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceType) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DataTimeSec) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.Generation) {
		body["generation"] = request.Generation
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/build-index"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更换实例资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, instanceId *string, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		body["newResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/actions/change-resource-group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloneSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneSqlInstanceResponse
func (client *Client) CloneSqlInstanceWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *CloneSqlInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloneSqlInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.TargetFolderId) {
		body["targetFolderId"] = request.TargetFolderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneSqlInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId)) + "/actions/clone"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneSqlInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAliasResponse
func (client *Client) CreateAliasWithContext(ctx context.Context, instanceId *string, request *CreateAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewMode) {
		query["newMode"] = request.NewMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Index) {
		body["index"] = request.Index
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlias"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/aliases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// Description:
//
// ### [](#method)Method
//
// `POST`
//
// ### [](#uri)URI
//
// `/openapi/ha3/instances/{instanceId}/clusters`
//
// @param request - CreateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, instanceId *string, request *CreateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoLoad) {
		body["autoLoad"] = request.AutoLoad
	}

	if !dara.IsNil(request.DataNode) {
		body["dataNode"] = request.DataNode
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.QueryNode) {
		body["queryNode"] = request.QueryNode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigDirResponse
func (client *Client) CreateConfigDirWithContext(ctx context.Context, instanceId *string, configName *string, request *CreateConfigDirRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConfigDirResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DirName) {
		body["dirName"] = request.DirName
	}

	if !dara.IsNil(request.ParentFullPath) {
		body["parentFullPath"] = request.ParentFullPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigDir"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/dir"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigDirResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigFileResponse
func (client *Client) CreateConfigFileWithContext(ctx context.Context, instanceId *string, configName *string, request *CreateConfigFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConfigFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.OssPath) {
		body["ossPath"] = request.OssPath
	}

	if !dara.IsNil(request.ParentFullPath) {
		body["parentFullPath"] = request.ParentFullPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/file"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates data sources.
//
// @param request - CreateDataSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithContext(ctx context.Context, instanceId *string, request *CreateDataSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuildIndex) {
		body["autoBuildIndex"] = request.AutoBuildIndex
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SaroConfig) {
		body["saroConfig"] = request.SaroConfig
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
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithContext(ctx context.Context, instanceId *string, database *string, request *CreateFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Parent) {
		body["parent"] = request.Parent
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/folders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an index.
//
// Description:
//
// ### Method
//
// ```java
//
// # POST
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/indexes
//
// ```
//
// @param request - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithContext(ctx context.Context, instanceId *string, request *CreateIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildParallelNum) {
		body["buildParallelNum"] = request.BuildParallelNum
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.DataSource) {
		body["dataSource"] = request.DataSource
	}

	if !dara.IsNil(request.DataSourceInfo) {
		body["dataSourceInfo"] = request.DataSourceInfo
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.Extend) {
		body["extend"] = request.Extend
	}

	if !dara.IsNil(request.MergeParallelNum) {
		body["mergeParallelNum"] = request.MergeParallelNum
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Havenask instance.
//
// Description:
//
// ### [](#)Method
//
// `POST`
//
// ### [](#uri)URI
//
// `/api/instances?dryRun=false`
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Components) {
		body["components"] = request.Components
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型信息
//
// @param request - CreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithContext(ctx context.Context, instanceId *string, request *CreateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public endpoint.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublicUrlResponse
func (client *Client) CreatePublicUrlWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePublicUrlResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePublicUrl"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/public-url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePublicUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSqlInstanceResponse
func (client *Client) CreateSqlInstanceWithContext(ctx context.Context, instanceId *string, database *string, request *CreateSqlInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSqlInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Parent) {
		body["parent"] = request.Parent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSqlInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSqlInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an index table.
//
// @param request - CreateTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableResponse
func (client *Client) CreateTableWithContext(ctx context.Context, instanceId *string, request *CreateTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataProcessConfig) {
		body["dataProcessConfig"] = request.DataProcessConfig
	}

	if !dara.IsNil(request.DataProcessorCount) {
		body["dataProcessorCount"] = request.DataProcessorCount
	}

	if !dara.IsNil(request.DataSource) {
		body["dataSource"] = request.DataSource
	}

	if !dara.IsNil(request.FieldSchema) {
		body["fieldSchema"] = request.FieldSchema
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PartitionCount) {
		body["partitionCount"] = request.PartitionCount
	}

	if !dara.IsNil(request.PrimaryKey) {
		body["primaryKey"] = request.PrimaryKey
	}

	if !dara.IsNil(request.RawSchema) {
		body["rawSchema"] = request.RawSchema
	}

	if !dara.IsNil(request.Scene) {
		body["scene"] = request.Scene
	}

	if !dara.IsNil(request.VectorIndex) {
		body["vectorIndex"] = request.VectorIndex
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTable"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调试模型
//
// @param request - DebugModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugModelResponse
func (client *Client) DebugModelWithContext(ctx context.Context, instanceId *string, modelName *string, request *DebugModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DebugModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsOnline) {
		query["isOnline"] = request.IsOnline
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugModel"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelName)) + "/actions/debug"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the details about advanced configurations.
//
// Description:
//
// ## Method
//
//	DELETE
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvanceConfigResponse
func (client *Client) DeleteAdvanceConfigWithContext(ctx context.Context, instanceId *string, configName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAdvanceConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdvanceConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdvanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithContext(ctx context.Context, instanceId *string, alias *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlias"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/aliases/" + dara.PercentEncode(dara.StringValue(alias))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigDirResponse
func (client *Client) DeleteConfigDirWithContext(ctx context.Context, instanceId *string, configName *string, request *DeleteConfigDirRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigDirResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirName) {
		query["dirName"] = request.DirName
	}

	if !dara.IsNil(request.ParentFullPath) {
		query["parentFullPath"] = request.ParentFullPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigDir"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/dir"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigDirResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigFileResponse
func (client *Client) DeleteConfigFileWithContext(ctx context.Context, instanceId *string, configName *string, request *DeleteConfigFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.ParentFullPath) {
		query["parentFullPath"] = request.ParentFullPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/file"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified data source.
//
// Description:
//
// ## Method
//
// `DELETE`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithContext(ctx context.Context, instanceId *string, dataSourceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithContext(ctx context.Context, instanceId *string, database *string, folderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/folders/" + dara.PercentEncode(dara.StringValue(folderId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an index.
//
// Description:
//
// ## Method
//
//	DELETE
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}?dataSource=xxx
//
// @param request - DeleteIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndexWithContext(ctx context.Context, instanceId *string, indexName *string, request *DeleteIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSource) {
		query["dataSource"] = request.DataSource
	}

	if !dara.IsNil(request.DeleteDataSource) {
		query["deleteDataSource"] = request.DeleteDataSource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the version of an index.
//
// Description:
//
// ## Method
//
//	DELETE
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexVersionResponse
func (client *Client) DeleteIndexVersionWithContext(ctx context.Context, instanceId *string, indexName *string, versionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndexVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndexVersion"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/versions/" + dara.PercentEncode(dara.StringValue(versionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndexVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified instance.
//
// Description:
//
// ### Method
//
// `DELETE`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithContext(ctx context.Context, instanceId *string, modelName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除公网域名
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicUrlResponse
func (client *Client) DeletePublicUrlWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePublicUrlResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePublicUrl"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/public-url"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePublicUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSqlInstanceResponse
func (client *Client) DeleteSqlInstanceWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSqlInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSqlInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSqlInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an index table.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableResponse
func (client *Client) DeleteTableWithContext(ctx context.Context, instanceId *string, tableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTable"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions.
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// @param request - ExecuteSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSqlInstanceResponse
func (client *Client) ExecuteSqlInstanceWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *ExecuteSqlInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteSqlInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CombineParam) {
		body["combineParam"] = request.CombineParam
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.DynamicParam) {
		body["dynamicParam"] = request.DynamicParam
	}

	if !dara.IsNil(request.Kvpair) {
		body["kvpair"] = request.Kvpair
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.StaticParam) {
		body["staticParam"] = request.StaticParam
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSqlInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId)) + "/actions/execution"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSqlInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a forced switchover.
//
// Description:
//
// ### [](#)Method
//
// ```java
//
// # PUT
//
// ```
//
// ### [](#uri)URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/force-switch/{fsmId}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceSwitchResponse
func (client *Client) ForceSwitchWithContext(ctx context.Context, instanceId *string, fsmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ForceSwitchResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForceSwitch"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/force-switch/" + dara.PercentEncode(dara.StringValue(fsmId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ForceSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an advanced configuration.
//
// Description:
//
// ## Method
//
//	GET
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}
//
// @param request - GetAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvanceConfigResponse
func (client *Client) GetAdvanceConfigWithContext(ctx context.Context, instanceId *string, configName *string, request *GetAdvanceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAdvanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdvanceConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an advanced configuration file.
//
// Description:
//
// ## Method
//
//	GET
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - GetAdvanceConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvanceConfigFileResponse
func (client *Client) GetAdvanceConfigFileWithContext(ctx context.Context, instanceId *string, configName *string, request *GetAdvanceConfigFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAdvanceConfigFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdvanceConfigFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/file"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvanceConfigFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cluster.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instance/{instanceId}/clusters/{clusterName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithContext(ctx context.Context, instanceId *string, clusterName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters/" + dara.PercentEncode(dara.StringValue(clusterName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the runtime information about a specified cluster.
//
// Description:
//
// ### Method
//
// # GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-run-time-info
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterRunTimeInfoResponse
func (client *Client) GetClusterRunTimeInfoWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterRunTimeInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterRunTimeInfo"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/cluster-run-time-info"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterRunTimeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithContext(ctx context.Context, instanceId *string, dataSourceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSource"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源部署信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDeployResponse
func (client *Client) GetDataSourceDeployWithContext(ctx context.Context, instanceId *string, deployName *string, dataSourceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataSourceDeployResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceDeploy"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/deploys/" + dara.PercentEncode(dara.StringValue(deployName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceDeployResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseSchemaResponse
func (client *Client) GetDatabaseSchemaWithContext(ctx context.Context, instanceId *string, database *string, tableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatabaseSchemaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseSchema"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName)) + "/schema"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the overview of the deployment.
//
// Description:
//
// ## Method
//
// # GET
//
// ## URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/deploy-graph
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployGraphResponse
func (client *Client) GetDeployGraphWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeployGraphResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployGraph"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/deploy-graph"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeployGraphResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an index table version.
//
// Description:
//
// ## [](#)Method
//
//	GET
//
// ## [](#uri)URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file
//
// @param request - GetFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithContext(ctx context.Context, instanceId *string, indexName *string, versionName *string, request *GetFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/versions/" + dara.PercentEncode(dara.StringValue(versionName)) + "/file"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an index version.
//
// Description:
//
// ## [](#)Method
//
//	GET
//
// ## [](#uri)URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexResponse
func (client *Client) GetIndexWithContext(ctx context.Context, instanceId *string, indexName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIndexResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the online effective policy of an index.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexOnlineStrategyResponse
func (client *Client) GetIndexOnlineStrategyWithContext(ctx context.Context, instanceId *string, dataSourceName *string, deployName *string, indexName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIndexOnlineStrategyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndexOnlineStrategy"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/deploys/" + dara.PercentEncode(dara.StringValue(deployName)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/online-strategy"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexOnlineStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about index versions that the current index version can be rolled back to.
//
// Description:
//
// ## Method
//
//	GET
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexVersionResponse
func (client *Client) GetIndexVersionWithContext(ctx context.Context, instanceId *string, clusterName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIndexVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIndexVersion"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters/" + dara.PercentEncode(dara.StringValue(clusterName)) + "/index-version"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIndexVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on the instance ID.
//
// Description:
//
// ### [](#)Method
//
// ```java
//
// # GET
//
// ```
//
// ### [](#uri)URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过实例ID和模型名称查询特定模型的详细配置信息。
//
// Description:
//
// ## 请求说明
//
// - 该API用于获取指定实例下的特定模型的详细信息，包括模型类型、URL、状态等。
//
// - 确保提供的`instanceId`和`modelName`是有效的，否则可能返回错误或找不到资源。
//
// - 返回的数据结构中包含了模型的内容（如请求头、参数等）以及创建和更新时间，有助于了解模型的具体配置及其最新状态。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelResponse
func (client *Client) GetModelWithContext(ctx context.Context, instanceId *string, modelName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModel"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the configuration information of a node.
//
// @param request - GetNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeConfigResponse
func (client *Client) GetNodeConfigWithContext(ctx context.Context, instanceId *string, request *GetNodeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["clusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/node-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSqlInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlInstanceResponse
func (client *Client) GetSqlInstanceWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *GetSqlInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSqlInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an index table.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithContext(ctx context.Context, instanceId *string, tableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTable"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an index version based on the ID of the full index version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableGenerationResponse
func (client *Client) GetTableGenerationWithContext(ctx context.Context, instanceId *string, tableName *string, generationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableGenerationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableGeneration"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName)) + "/index_versions/" + dara.PercentEncode(dara.StringValue(generationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the files in an advanced configuration directory.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/dir?dirName={dirName}`
//
// @param request - ListAdvanceConfigDirRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdvanceConfigDirResponse
func (client *Client) ListAdvanceConfigDirWithContext(ctx context.Context, instanceId *string, configName *string, request *ListAdvanceConfigDirRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAdvanceConfigDirResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirName) {
		query["dirName"] = request.DirName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAdvanceConfigDir"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/dir"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAdvanceConfigDirResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of advanced configurations.
//
// Description:
//
// ## Sample requests
//
// `GET /openapi/ha3/instances/ose-test1/advanced-configs`
//
// @param request - ListAdvanceConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdvanceConfigsResponse
func (client *Client) ListAdvanceConfigsWithContext(ctx context.Context, instanceId *string, request *ListAdvanceConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAdvanceConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceName) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.IndexName) {
		query["indexName"] = request.IndexName
	}

	if !dara.IsNil(request.NewMode) {
		query["newMode"] = request.NewMode
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAdvanceConfigs"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAdvanceConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesResponse
func (client *Client) ListAliasesWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAliases"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/aliases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster names.
//
// Description:
//
// ### Method
//
// # GET
//
// ### URI
//
// /openapi/ha3/instances/{instanceId}/cluster-names
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterNamesResponse
func (client *Client) ListClusterNamesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterNamesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterNames"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/cluster-names"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterNamesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cluster tasks.
//
// Description:
//
// ### Method
//
// ```java
//
// # GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/cluster-tasks
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterTasksResponse
func (client *Client) ListClusterTasksWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterTasksResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterTasks"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/cluster-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries clusters.
//
// Description:
//
// ### Method
//
// ```java
//
// # GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the schema information about a data source.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/schemas`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSchemasResponse
func (client *Client) ListDataSourceSchemasWithContext(ctx context.Context, instanceId *string, dataSourceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourceSchemasResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceSchemas"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays data source tasks.
//
// Description:
//
// ### [](#)Method
//
// ```java
//
// # GET
//
// ```
//
// ### [](#uri)URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/data-source-tasks
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTasksResponse
func (client *Client) ListDataSourceTasksWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourceTasksResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceTasks"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-source-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of data sources.
//
// Description:
//
// ## Method
//
// `GET`
//
// ## URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources`
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources"),
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical index versions of a data source.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/data-sources/{dataSourceName}/generations?domainName={domainName}`
//
// @param request - ListDateSourceGenerationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDateSourceGenerationsResponse
func (client *Client) ListDateSourceGenerationsWithContext(ctx context.Context, instanceId *string, dataSourceName *string, request *ListDateSourceGenerationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDateSourceGenerationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["domainName"] = request.DomainName
	}

	if !dara.IsNil(request.ValidStatus) {
		query["validStatus"] = request.ValidStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDateSourceGenerations"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/generations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDateSourceGenerationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexRecoverRecordsResponse
func (client *Client) ListIndexRecoverRecordsWithContext(ctx context.Context, indexName *string, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexRecoverRecordsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndexRecoverRecords"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/actions/list-recover-records"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexRecoverRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the index list.
//
// Description:
//
// ## Method
//
//	GET
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/indexes
//
// @param request - ListIndexesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexesResponse
func (client *Client) ListIndexesWithContext(ctx context.Context, instanceId *string, request *ListIndexesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIndexesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["catalog"] = request.Catalog
	}

	if !dara.IsNil(request.Database) {
		query["database"] = request.Database
	}

	if !dara.IsNil(request.NewMode) {
		query["newMode"] = request.NewMode
	}

	if !dara.IsNil(request.Table) {
		query["table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndexes"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of an instance.
//
// Description:
//
// ### Method
//
// `GET`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/specs?type=qrs`
//
// @param request - ListInstanceSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceSpecsResponse
func (client *Client) ListInstanceSpecsWithContext(ctx context.Context, instanceId *string, request *ListInstanceSpecsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceSpecsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceSpecs"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/specs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceSpecsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// Description:
//
// ### [](#)Method
//
// `GET`
//
// ### [](#uri)URI
//
// `/openapi/ha3/instances`
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, tmpReq *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["catalog"] = request.Catalog
	}

	if !dara.IsNil(request.DataSourceType) {
		query["dataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.Database) {
		query["database"] = request.Database
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.Edition) {
		query["edition"] = request.Edition
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Table) {
		query["table"] = request.Table
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogsResponse
func (client *Client) ListLogsWithContext(ctx context.Context, instanceId *string, request *ListLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogs"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过实例ID查询指定条件下的模型列表。
//
// Description:
//
// ## 请求说明
//
// 本API用于从指定实例中获取模型列表，支持通过模型名称、类型以及分页参数进行筛选。请求时需提供实例ID作为路径参数，其他筛选条件为可选的查询参数。
//
// @param request - ListModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
func (client *Client) ListModelsWithContext(ctx context.Context, instanceId *string, request *ListModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModels"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an online configuration.
//
// Description:
//
// ### Method
//
// ```java
//
// # GET
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs?domain={domain}
//
// ```
//
// @param request - ListOnlineConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineConfigsResponse
func (client *Client) ListOnlineConfigsWithContext(ctx context.Context, instanceId *string, nodeName *string, request *ListOnlineConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOnlineConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOnlineConfigs"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/node/" + dara.PercentEncode(dara.StringValue(nodeName)) + "/online-configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOnlineConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPausePolicysResponse
func (client *Client) ListPausePolicysWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPausePolicysResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPausePolicys"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/pause-policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPausePolicysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 召回引擎版使用POST请求获取搜索测试结果
//
// @param request - ListPostQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPostQueryResultResponse
func (client *Client) ListPostQueryResultWithContext(ctx context.Context, instanceId *string, request *ListPostQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPostQueryResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPostQueryResult"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPostQueryResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the query result.
//
// Description:
//
// ### [](#)Method
//
// `GET`
//
// ### [](#uri)URI
//
// `/openapi/ha3/instances/{instanceId}/query?query=xxxx`
//
// @param request - ListQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryResultResponse
func (client *Client) ListQueryResultWithContext(ctx context.Context, instanceId *string, request *ListQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQueryResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.Sql) {
		query["sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueryResult"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueryResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 召回引擎版获取rest查询搜索测试结果
//
// @param request - ListRestQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRestQueryResultResponse
func (client *Client) ListRestQueryResultWithContext(ctx context.Context, instanceId *string, request *ListRestQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRestQueryResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IndexName) {
		body["indexName"] = request.IndexName
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRestQueryResult"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/rest-query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRestQueryResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过数据源配置获取schema信息
//
// @param request - ListSchemasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemasResponse
func (client *Client) ListSchemasWithContext(ctx context.Context, instanceId *string, request *ListSchemasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["accessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.AccessSecret) {
		query["accessSecret"] = request.AccessSecret
	}

	if !dara.IsNil(request.Endpoint) {
		query["endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Partition) {
		query["partition"] = request.Partition
	}

	if !dara.IsNil(request.Project) {
		query["project"] = request.Project
	}

	if !dara.IsNil(request.Table) {
		query["table"] = request.Table
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemas"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of index versions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableGenerationsResponse
func (client *Client) ListTableGenerationsWithContext(ctx context.Context, instanceId *string, tableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTableGenerationsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableGenerations"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName)) + "/index_versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableGenerationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of index tables.
//
// @param request - ListTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithContext(ctx context.Context, instanceId *string, request *ListTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewMode) {
		query["newMode"] = request.NewMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables"),
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
// 查标签接口
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("resourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/resource-tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群任务列表（数据源+集群）
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, instanceId *string, request *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向量检索版获取向量查询搜索测试结果
//
// @param request - ListVectorQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVectorQueryResultResponse
func (client *Client) ListVectorQueryResultWithContext(ctx context.Context, instanceId *string, request *ListVectorQueryResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVectorQueryResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["path"] = request.Path
	}

	if !dara.IsNil(request.QueryType) {
		query["queryType"] = request.QueryType
	}

	if !dara.IsNil(request.VectorQueryType) {
		query["vectorQueryType"] = request.VectorQueryType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVectorQueryResult"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/vector-query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVectorQueryResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAdvanceConfigResponse
func (client *Client) ModifyAdvanceConfigWithContext(ctx context.Context, instanceId *string, configName *string, request *ModifyAdvanceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAdvanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.Desc) {
		body["desc"] = request.Desc
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.UpdateTime) {
		body["updateTime"] = request.UpdateTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAdvanceConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAdvanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the advanced configurations.
//
// Description:
//
// ## Method
//
//	put
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/file?fileName={fileName}
//
// @param request - ModifyAdvanceConfigFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAdvanceConfigFileResponse
func (client *Client) ModifyAdvanceConfigFileWithContext(ctx context.Context, instanceId *string, configName *string, request *ModifyAdvanceConfigFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAdvanceConfigFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAdvanceConfigFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/file"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAdvanceConfigFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAliasResponse
func (client *Client) ModifyAliasWithContext(ctx context.Context, instanceId *string, alias *string, request *ModifyAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Index) {
		body["index"] = request.Index
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAlias"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/aliases/" + dara.PercentEncode(dara.StringValue(alias))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a specified cluster.
//
// Description:
//
// ### [](#)Method
//
// `PUT`
//
// ### [](#uri)URI
//
// `/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/desc`
//
// @param request - ModifyClusterDescRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterDescResponse
func (client *Client) ModifyClusterDescWithContext(ctx context.Context, instanceId *string, clusterName *string, request *ModifyClusterDescRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterDescResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterDesc"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters/" + dara.PercentEncode(dara.StringValue(clusterName)) + "/desc"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterDescResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration information of a cluster.
//
// Description:
//
// ## Request syntax
//
//	PUT /openapi/ha3/instances/{instanceId}/cluster-offline-config
//
// @param request - ModifyClusterOfflineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterOfflineConfigResponse
func (client *Client) ModifyClusterOfflineConfigWithContext(ctx context.Context, instanceId *string, request *ModifyClusterOfflineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterOfflineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildMode) {
		body["buildMode"] = request.BuildMode
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.DataSourceName) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceType) {
		body["dataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DataTimeSec) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.Generation) {
		body["generation"] = request.Generation
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	if !dara.IsNil(request.PushMode) {
		body["pushMode"] = request.PushMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterOfflineConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/cluster-offline-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterOfflineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the online configuration of a cluster.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/cluster-online-config`
//
// @param request - ModifyClusterOnlineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterOnlineConfigResponse
func (client *Client) ModifyClusterOnlineConfigWithContext(ctx context.Context, instanceId *string, request *ModifyClusterOnlineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterOnlineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Clusters) {
		body["clusters"] = request.Clusters
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterOnlineConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/cluster-online-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterOnlineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改数据源部署信息
//
// @param request - ModifyDataSourceDeployRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataSourceDeployResponse
func (client *Client) ModifyDataSourceDeployWithContext(ctx context.Context, instanceId *string, deployName *string, dataSourceName *string, request *ModifyDataSourceDeployRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDataSourceDeployResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GenerationId) {
		query["generationId"] = request.GenerationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuildIndex) {
		body["autoBuildIndex"] = request.AutoBuildIndex
	}

	if !dara.IsNil(request.Extend) {
		body["extend"] = request.Extend
	}

	if !dara.IsNil(request.Processor) {
		body["processor"] = request.Processor
	}

	if !dara.IsNil(request.Storage) {
		body["storage"] = request.Storage
	}

	if !dara.IsNil(request.Swift) {
		body["swift"] = request.Swift
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataSourceDeploy"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/deploys/" + dara.PercentEncode(dara.StringValue(deployName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataSourceDeployResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a file.
//
// Description:
//
// ## Method
//
//	PUT
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}/versions/{versionName}/file?fileName=/root/test.txt
//
// @param request - ModifyFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileResponse
func (client *Client) ModifyFileWithContext(ctx context.Context, instanceId *string, indexName *string, versionName *string, request *ModifyFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFile"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/versions/" + dara.PercentEncode(dara.StringValue(versionName)) + "/file"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑索引表
//
// @param request - ModifyIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexResponse
func (client *Client) ModifyIndexWithContext(ctx context.Context, instanceId *string, indexName *string, request *ModifyIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildParallelNum) {
		body["buildParallelNum"] = request.BuildParallelNum
	}

	if !dara.IsNil(request.Cluster) {
		body["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.ClusterConfigName) {
		body["clusterConfigName"] = request.ClusterConfigName
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.DataSource) {
		body["dataSource"] = request.DataSource
	}

	if !dara.IsNil(request.DataSourceInfo) {
		body["dataSourceInfo"] = request.DataSourceInfo
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.Extend) {
		body["extend"] = request.Extend
	}

	if !dara.IsNil(request.MergeParallelNum) {
		body["mergeParallelNum"] = request.MergeParallelNum
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	if !dara.IsNil(request.PushMode) {
		body["pushMode"] = request.PushMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the online policy of an index.
//
// @param request - ModifyIndexOnlineStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexOnlineStrategyResponse
func (client *Client) ModifyIndexOnlineStrategyWithContext(ctx context.Context, instanceId *string, dataSourceName *string, deployName *string, indexName *string, request *ModifyIndexOnlineStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyIndexOnlineStrategyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeRate) {
		body["changeRate"] = request.ChangeRate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIndexOnlineStrategy"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/deploys/" + dara.PercentEncode(dara.StringValue(deployName)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/online-strategy"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIndexOnlineStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about index partitions.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/index-partition`
//
// @param request - ModifyIndexPartitionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexPartitionResponse
func (client *Client) ModifyIndexPartitionWithContext(ctx context.Context, instanceId *string, request *ModifyIndexPartitionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyIndexPartitionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceName) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DomainName) {
		body["domainName"] = request.DomainName
	}

	if !dara.IsNil(request.Generation) {
		body["generation"] = request.Generation
	}

	if !dara.IsNil(request.IndexInfos) {
		body["indexInfos"] = request.IndexInfos
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIndexPartition"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/index-partition"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIndexPartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the index version of a cluster (an index version rollback).
//
// Description:
//
// ## [](#)Method
//
//	PUT
//
// ## [](#uri)URI
//
//	/openapi/ha3/instances/{instanceId}/clusters/{clusterName}/index-version
//
// @param request - ModifyIndexVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIndexVersionResponse
func (client *Client) ModifyIndexVersionWithContext(ctx context.Context, instanceId *string, clusterName *string, request *ModifyIndexVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyIndexVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIndexVersion"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters/" + dara.PercentEncode(dara.StringValue(clusterName)) + "/index-version"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIndexVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改模型详情，修改模型状态
//
// @param request - ModifyModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyModelResponse
func (client *Client) ModifyModelWithContext(ctx context.Context, instanceId *string, modelName *string, request *ModifyModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyModel"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/models/" + dara.PercentEncode(dara.StringValue(modelName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a node.
//
// Description:
//
// ### Method
//
// ```java
//
// # PUT
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node-config?type=qrs&name=test
//
// ```
//
// @param request - ModifyNodeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeConfigResponse
func (client *Client) ModifyNodeConfigWithContext(ctx context.Context, instanceId *string, request *ModifyNodeConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyNodeConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["clusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DataSourceName) {
		query["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		body["active"] = request.Active
	}

	if !dara.IsNil(request.DataDuplicateNumber) {
		body["dataDuplicateNumber"] = request.DataDuplicateNumber
	}

	if !dara.IsNil(request.DataFragmentNumber) {
		body["dataFragmentNumber"] = request.DataFragmentNumber
	}

	if !dara.IsNil(request.FlowRatio) {
		body["flowRatio"] = request.FlowRatio
	}

	if !dara.IsNil(request.MinServicePercent) {
		body["minServicePercent"] = request.MinServicePercent
	}

	if !dara.IsNil(request.Published) {
		body["published"] = request.Published
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodeConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/node-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies online configurations.
//
// Description:
//
// ### Method
//
// ```java
//
// put
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/node/{nodeName}/online-configs/{indexName}
//
// ```
//
// @param request - ModifyOnlineConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOnlineConfigResponse
func (client *Client) ModifyOnlineConfigWithContext(ctx context.Context, instanceId *string, nodeName *string, indexName *string, request *ModifyOnlineConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyOnlineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOnlineConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/node/" + dara.PercentEncode(dara.StringValue(nodeName)) + "/online-configs/" + dara.PercentEncode(dara.StringValue(indexName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOnlineConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例的密码
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/password`
//
// @param request - ModifyPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPasswordResponse
func (client *Client) ModifyPasswordWithContext(ctx context.Context, instanceId *string, request *ModifyPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPassword"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/password"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyPausePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPausePolicyResponse
func (client *Client) ModifyPausePolicyWithContext(ctx context.Context, instanceId *string, request *ModifyPausePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyPausePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPausePolicy"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/pause-policies"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPausePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改公网域名访问白名单
//
// @param request - ModifyPublicUrlIpListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPublicUrlIpListResponse
func (client *Client) ModifyPublicUrlIpListWithContext(ctx context.Context, instanceId *string, request *ModifyPublicUrlIpListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyPublicUrlIpListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPublicUrlIpList"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/public-url-ip-list"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPublicUrlIpListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过指定实例ID来修改数据节点的副本或分片数量。
//
// Description:
//
// ## 请求说明
//
// 本API允许用户修改特定实例下的数据节点副本数或分片数。请求时，需提供实例ID，并在请求体中指定要修改的`replica`（副本数）或`partition`（分片数）。请注意，这两个参数都是可选的，但至少需要提供其中一个以执行更新操作。
//
// @param request - ModifySearcherReplicaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySearcherReplicaResponse
func (client *Client) ModifySearcherReplicaWithContext(ctx context.Context, instanceId *string, request *ModifySearcherReplicaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifySearcherReplicaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	if !dara.IsNil(request.Replica) {
		body["replica"] = request.Replica
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySearcherReplica"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/replica"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySearcherReplicaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an index table.
//
// @param request - ModifyTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTableResponse
func (client *Client) ModifyTableWithContext(ctx context.Context, instanceId *string, tableName *string, request *ModifyTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataProcessConfig) {
		body["dataProcessConfig"] = request.DataProcessConfig
	}

	if !dara.IsNil(request.DataSource) {
		body["dataSource"] = request.DataSource
	}

	if !dara.IsNil(request.FieldSchema) {
		body["fieldSchema"] = request.FieldSchema
	}

	if !dara.IsNil(request.PartitionCount) {
		body["partitionCount"] = request.PartitionCount
	}

	if !dara.IsNil(request.PrimaryKey) {
		body["primaryKey"] = request.PrimaryKey
	}

	if !dara.IsNil(request.RawSchema) {
		body["rawSchema"] = request.RawSchema
	}

	if !dara.IsNil(request.VectorIndex) {
		body["vectorIndex"] = request.VectorIndex
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTable"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a version of advanced configurations.
//
// Description:
//
// ## Method
//
// ~~~
//
// # POST
//
// ~~~
//
// ## URI
//
// ~~~
//
// /openapi/ha3/instances/{instanceId}/advanced-configs/{configName}/actions/publish
//
// ~~~
//
// @param request - PublishAdvanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishAdvanceConfigResponse
func (client *Client) PublishAdvanceConfigWithContext(ctx context.Context, instanceId *string, configName *string, request *PublishAdvanceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishAdvanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Desc) {
		body["desc"] = request.Desc
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishAdvanceConfig"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/advanced-configs/" + dara.PercentEncode(dara.StringValue(configName)) + "/actions/publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishAdvanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified index version.
//
// Description:
//
// ## Method
//
//	POST
//
// ## URI
//
//	/openapi/ha3/instances/{instanceId}/indexes/{indexName}/actions/publish
//
// @param request - PublishIndexVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishIndexVersionResponse
func (client *Client) PublishIndexVersionWithContext(ctx context.Context, instanceId *string, indexName *string, request *PublishIndexVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishIndexVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishIndexVersion"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/actions/publish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishIndexVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PushDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDocumentsResponse
func (client *Client) PushDocumentsWithContext(ctx context.Context, instanceId *string, dataSourceName *string, request *PushDocumentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushDocumentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PkField) {
		query["pkField"] = request.PkField
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushDocuments"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceName)) + "/actions/bulk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores data from an index.
//
// Description:
//
// ### Method
//
// `POST`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}/recover-index`
//
// @param request - RecoverIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverIndexResponse
func (client *Client) RecoverIndexWithContext(ctx context.Context, instanceId *string, request *RecoverIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecoverIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BuildDeployId) {
		body["buildDeployId"] = request.BuildDeployId
	}

	if !dara.IsNil(request.DataSourceName) {
		body["dataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.Generation) {
		body["generation"] = request.Generation
	}

	if !dara.IsNil(request.IndexName) {
		body["indexName"] = request.IndexName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/recover-index"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds an index.
//
// @param request - ReindexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReindexResponse
func (client *Client) ReindexWithContext(ctx context.Context, instanceId *string, tableName *string, request *ReindexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReindexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataTimeSec) {
		body["dataTimeSec"] = request.DataTimeSec
	}

	if !dara.IsNil(request.OssDataPath) {
		body["ossDataPath"] = request.OssDataPath
	}

	if !dara.IsNil(request.Partition) {
		body["partition"] = request.Partition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Reindex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(tableName)) + "/reindex"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReindexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cluster.
//
// Description:
//
// ### Method
//
// ```java
//
// # DELETE
//
// ```
//
// ### URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/clusters/{clusterName}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClusterResponse
func (client *Client) RemoveClusterWithContext(ctx context.Context, instanceId *string, clusterName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveClusterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveCluster"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/clusters/" + dara.PercentEncode(dara.StringValue(clusterName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RenameFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFolderResponse
func (client *Client) RenameFolderWithContext(ctx context.Context, instanceId *string, database *string, folderId *string, request *RenameFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenameFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameFolder"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/folders/" + dara.PercentEncode(dara.StringValue(folderId)) + "/name"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartIndexResponse
func (client *Client) StartIndexWithContext(ctx context.Context, instanceId *string, indexName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartIndexResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/startIndex"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopIndexResponse
func (client *Client) StopIndexWithContext(ctx context.Context, instanceId *string, indexName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopIndexResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopIndex"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/indexes/" + dara.PercentEncode(dara.StringValue(indexName)) + "/stopIndex"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an FSM task.
//
// Description:
//
// ### [](#)Method
//
// ```java
//
// # PUT
//
// ```
//
// ### [](#uri)URI
//
// ```java
//
// /openapi/ha3/instances/{instanceId}/stop-task/{fsmId}
//
// ```
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskResponse
func (client *Client) StopTaskWithContext(ctx context.Context, instanceId *string, fsmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTask"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/stop-task/" + dara.PercentEncode(dara.StringValue(fsmId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/resource-tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("resourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKey) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, dara.String("tagKey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeyShrink) {
		query["tagKey"] = request.TagKeyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/resource-tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified instance.
//
// Description:
//
// ### Method
//
// `PUT`
//
// ### URI
//
// `/openapi/ha3/instances/{instanceId}`
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Components) {
		body["components"] = request.Components
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.OrderType) {
		body["orderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSqlInstanceContentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSqlInstanceContentResponse
func (client *Client) UpdateSqlInstanceContentWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *UpdateSqlInstanceContentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSqlInstanceContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSqlInstanceContent"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId)) + "/content"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSqlInstanceContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSqlInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSqlInstanceNameResponse
func (client *Client) UpdateSqlInstanceNameWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *UpdateSqlInstanceNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSqlInstanceNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSqlInstanceName"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId)) + "/name"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSqlInstanceNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSqlInstanceParamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSqlInstanceParamsResponse
func (client *Client) UpdateSqlInstanceParamsWithContext(ctx context.Context, instanceId *string, database *string, sqlInstanceId *string, request *UpdateSqlInstanceParamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSqlInstanceParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CombineParam) {
		body["combineParam"] = request.CombineParam
	}

	if !dara.IsNil(request.DynamicParam) {
		body["dynamicParam"] = request.DynamicParam
	}

	if !dara.IsNil(request.Kvpair) {
		body["kvpair"] = request.Kvpair
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.StaticParam) {
		body["staticParam"] = request.StaticParam
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSqlInstanceParams"),
		Version:     dara.String("2021-10-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/ha3/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sql-studio/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/sql-instances/" + dara.PercentEncode(dara.StringValue(sqlInstanceId)) + "/params"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSqlInstanceParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
