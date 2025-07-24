// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 检测资源连接状态。
//
// @param request - CheckInstanceDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceDatasourceResponse
func (client *Client) CheckInstanceDatasourceWithContext(ctx context.Context, InstanceId *string, request *CheckInstanceDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceDatasourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceDatasource"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/action/checkdatasource"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceDatasourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查FG配置内容是否正确，是否满足所有规则。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckModelFeatureFGFeatureResponse
func (client *Client) CheckModelFeatureFGFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckModelFeatureFGFeatureResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckModelFeatureFGFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/action/checkfgfeature"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckModelFeatureFGFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源。
//
// @param request - CreateDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasourceResponse
func (client *Client) CreateDatasourceWithContext(ctx context.Context, InstanceId *string, request *CreateDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasource"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征实体
//
// @param request - CreateFeatureEntityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureEntityResponse
func (client *Client) CreateFeatureEntityWithContext(ctx context.Context, InstanceId *string, request *CreateFeatureEntityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JoinId) {
		body["JoinId"] = request.JoinId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureEntity"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureentities"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征视图。
//
// @param request - CreateFeatureViewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeatureViewResponse
func (client *Client) CreateFeatureViewWithContext(ctx context.Context, InstanceId *string, request *CreateFeatureViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureViewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.FeatureEntityId) {
		body["FeatureEntityId"] = request.FeatureEntityId
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegisterDatasourceId) {
		body["RegisterDatasourceId"] = request.RegisterDatasourceId
	}

	if !dara.IsNil(request.RegisterTable) {
		body["RegisterTable"] = request.RegisterTable
	}

	if !dara.IsNil(request.SyncOnlineTable) {
		body["SyncOnlineTable"] = request.SyncOnlineTable
	}

	if !dara.IsNil(request.TTL) {
		body["TTL"] = request.TTL
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WriteMethod) {
		body["WriteMethod"] = request.WriteMethod
	}

	if !dara.IsNil(request.WriteToFeatureDB) {
		body["WriteToFeatureDB"] = request.WriteToFeatureDB
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFeatureView"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFeatureViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Feature Store实例。
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
	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances"),
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
// 创建大模型调用信息配置
//
// @param request - CreateLLMConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLLMConfigResponse
func (client *Client) CreateLLMConfigWithContext(ctx context.Context, InstanceId *string, request *CreateLLMConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLLMConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		body["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.BatchSize) {
		body["BatchSize"] = request.BatchSize
	}

	if !dara.IsNil(request.MaxTokens) {
		body["MaxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Rps) {
		body["Rps"] = request.Rps
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLLMConfig"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/llmconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLLMConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建label表
//
// @param request - CreateLabelTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLabelTableResponse
func (client *Client) CreateLabelTableWithContext(ctx context.Context, InstanceId *string, request *CreateLabelTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLabelTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLabelTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labeltables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLabelTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型特征。
//
// @param request - CreateModelFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelFeatureResponse
func (client *Client) CreateModelFeatureWithContext(ctx context.Context, InstanceId *string, request *CreateModelFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Features) {
		body["Features"] = request.Features
	}

	if !dara.IsNil(request.LabelPriorityLevel) {
		body["LabelPriorityLevel"] = request.LabelPriorityLevel
	}

	if !dara.IsNil(request.LabelTableId) {
		body["LabelTableId"] = request.LabelTableId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SequenceFeatureViewIds) {
		body["SequenceFeatureViewIds"] = request.SequenceFeatureViewIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建FeatureStore项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, InstanceId *string, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OfflineDatasourceId) {
		body["OfflineDatasourceId"] = request.OfflineDatasourceId
	}

	if !dara.IsNil(request.OfflineLifeCycle) {
		body["OfflineLifeCycle"] = request.OfflineLifeCycle
	}

	if !dara.IsNil(request.OnlineDatasourceId) {
		body["OnlineDatasourceId"] = request.OnlineDatasourceId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects"),
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
// 创建feature store服务账户角色
//
// @param request - CreateServiceIdentityRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRoleWithContext(ctx context.Context, request *CreateServiceIdentityRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceIdentityRole"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/serviceidentityroles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定数据源。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasourceResponse
func (client *Client) DeleteDatasourceWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasource"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定特征实体
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureEntityResponse
func (client *Client) DeleteFeatureEntityWithContext(ctx context.Context, InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFeatureEntityResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFeatureEntity"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureentities/" + dara.PercentEncode(dara.StringValue(FeatureEntityId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFeatureEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定特征视图。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureViewResponse
func (client *Client) DeleteFeatureViewWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFeatureViewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFeatureView"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFeatureViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除大模型调用信息配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLLMConfigResponse
func (client *Client) DeleteLLMConfigWithContext(ctx context.Context, InstanceId *string, LLMConfigId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLLMConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLLMConfig"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/llmconfigs/" + dara.PercentEncode(dara.StringValue(LLMConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLLMConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除label表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLabelTableResponse
func (client *Client) DeleteLabelTableWithContext(ctx context.Context, InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLabelTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLabelTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labeltables/" + dara.PercentEncode(dara.StringValue(LabelTableId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLabelTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定模型特征。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelFeatureResponse
func (client *Client) DeleteModelFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelFeatureResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定Feature Store项目。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId))),
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
// 导出训练集表。
//
// @param request - ExportModelFeatureTrainingSetTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportModelFeatureTrainingSetTableResponse
func (client *Client) ExportModelFeatureTrainingSetTableWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, request *ExportModelFeatureTrainingSetTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportModelFeatureTrainingSetTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FeatureViewConfig) {
		body["FeatureViewConfig"] = request.FeatureViewConfig
	}

	if !dara.IsNil(request.LabelInputConfig) {
		body["LabelInputConfig"] = request.LabelInputConfig
	}

	if !dara.IsNil(request.RealTimeIterateInterval) {
		body["RealTimeIterateInterval"] = request.RealTimeIterateInterval
	}

	if !dara.IsNil(request.RealTimePartitionCountValue) {
		body["RealTimePartitionCountValue"] = request.RealTimePartitionCountValue
	}

	if !dara.IsNil(request.TrainingSetConfig) {
		body["TrainingSetConfig"] = request.TrainingSetConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportModelFeatureTrainingSetTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/action/exporttrainingsettable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportModelFeatureTrainingSetTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasourceResponse
func (client *Client) GetDatasourceWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasource"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源下指定表的详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasourceTableResponse
func (client *Client) GetDatasourceTableWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, TableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasourceTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasourceTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId)) + "/tables/" + dara.PercentEncode(dara.StringValue(TableName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasourceTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征实体详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureEntityResponse
func (client *Client) GetFeatureEntityWithContext(ctx context.Context, InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureEntityResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureEntity"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureentities/" + dara.PercentEncode(dara.StringValue(FeatureEntityId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureViewResponse
func (client *Client) GetFeatureViewWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureViewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureView"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
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
// 获取 LLMConfig 信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLLMConfigResponse
func (client *Client) GetLLMConfigWithContext(ctx context.Context, InstanceId *string, LLMConfigId *string, RegionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLLMConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLLMConfig"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/llmconfigs/" + dara.PercentEncode(dara.StringValue(LLMConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLLMConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Label表详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLabelTableResponse
func (client *Client) GetLabelTableWithContext(ctx context.Context, InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLabelTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLabelTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labeltables/" + dara.PercentEncode(dara.StringValue(LabelTableId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLabelTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征详情。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureResponse
func (client *Client) GetModelFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征的FG特征配置信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureFGFeatureResponse
func (client *Client) GetModelFeatureFGFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureFGFeatureResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelFeatureFGFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/fgfeature"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelFeatureFGFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征的fg.json文件配置信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelFeatureFGInfoResponse
func (client *Client) GetModelFeatureFGInfoWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureFGInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelFeatureFGInfo"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/fginfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelFeatureFGInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定Feature Store项目详细信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId))),
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
// 获取项目下特征实体详细信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectFeatureEntityResponse
func (client *Client) GetProjectFeatureEntityWithContext(ctx context.Context, InstanceId *string, ProjectId *string, FeatureEntityName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectFeatureEntityResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectFeatureEntity"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/featureentities/" + dara.PercentEncode(dara.StringValue(FeatureEntityName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectFeatureEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取feature store服务账户角色。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRoleWithContext(ctx context.Context, RoleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceIdentityRole"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/serviceidentityroles/" + dara.PercentEncode(dara.StringValue(RoleName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, InstanceId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
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
// 获取数据源下所有特征视图信息。
//
// @param request - ListDatasourceFeatureViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasourceFeatureViewsResponse
func (client *Client) ListDatasourceFeatureViewsWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, request *ListDatasourceFeatureViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourceFeatureViewsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ShowStorageUsage) {
		query["ShowStorageUsage"] = request.ShowStorageUsage
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasourceFeatureViews"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId)) + "/featureviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasourceFeatureViewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源下所有表。
//
// @param request - ListDatasourceTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasourceTablesResponse
func (client *Client) ListDatasourceTablesWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, request *ListDatasourceTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourceTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasourceTables"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasourceTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据源列表。
//
// @param request - ListDatasourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasourcesResponse
func (client *Client) ListDatasourcesWithContext(ctx context.Context, InstanceId *string, request *ListDatasourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasources"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建特征实体列表
//
// @param tmpReq - ListFeatureEntitiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureEntitiesResponse
func (client *Client) ListFeatureEntitiesWithContext(ctx context.Context, InstanceId *string, tmpReq *ListFeatureEntitiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureEntitiesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FeatureEntityIds) {
		request.FeatureEntityIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureEntityIds, dara.String("FeatureEntityIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureEntityIdsShrink) {
		query["FeatureEntityIds"] = request.FeatureEntityIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureEntities"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureentities"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征字段血缘关系。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewFieldRelationshipsResponse
func (client *Client) ListFeatureViewFieldRelationshipsWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, FieldName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewFieldRelationshipsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureViewFieldRelationships"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId)) + "/fields/" + dara.PercentEncode(dara.StringValue(FieldName)) + "/relationships"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureViewFieldRelationshipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图下的在线特征数据。
//
// @param tmpReq - ListFeatureViewOnlineFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewOnlineFeaturesResponse
func (client *Client) ListFeatureViewOnlineFeaturesWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, tmpReq *ListFeatureViewOnlineFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewOnlineFeaturesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureViewOnlineFeaturesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JoinIds) {
		request.JoinIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JoinIds, dara.String("JoinIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.JoinIdsShrink) {
		query["JoinIds"] = request.JoinIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureViewOnlineFeatures"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId)) + "/onlinefeatures"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureViewOnlineFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewRelationshipsResponse
func (client *Client) ListFeatureViewRelationshipsWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewRelationshipsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureViewRelationships"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId)) + "/relationships"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureViewRelationshipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图列表。
//
// @param tmpReq - ListFeatureViewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeatureViewsResponse
func (client *Client) ListFeatureViewsWithContext(ctx context.Context, InstanceId *string, tmpReq *ListFeatureViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListFeatureViewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FeatureViewIds) {
		request.FeatureViewIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureViewIds, dara.String("FeatureViewIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	if !dara.IsNil(request.FeatureViewIdsShrink) {
		query["FeatureViewIds"] = request.FeatureViewIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatureViews"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeatureViewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Feature Store实例列表。
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances"),
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

// Summary:
//
// 获取大模型调用信息配置
//
// @param request - ListLLMConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLLMConfigsResponse
func (client *Client) ListLLMConfigsWithContext(ctx context.Context, InstanceId *string, request *ListLLMConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLLMConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLLMConfigs"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/llmconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLLMConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Label表列表。
//
// @param tmpReq - ListLabelTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLabelTablesResponse
func (client *Client) ListLabelTablesWithContext(ctx context.Context, InstanceId *string, tmpReq *ListLabelTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLabelTablesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListLabelTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelTableIds) {
		request.LabelTableIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelTableIds, dara.String("LabelTableIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelTableIdsShrink) {
		query["LabelTableIds"] = request.LabelTableIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLabelTables"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labeltables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLabelTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取注册FG特征时模型特征下可选的所有特征。
//
// @param request - ListModelFeatureAvailableFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelFeatureAvailableFeaturesResponse
func (client *Client) ListModelFeatureAvailableFeaturesWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, request *ListModelFeatureAvailableFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelFeatureAvailableFeaturesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureName) {
		query["FeatureName"] = request.FeatureName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelFeatureAvailableFeatures"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/availablefeatures"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelFeatureAvailableFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型特征列表。
//
// @param tmpReq - ListModelFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelFeaturesResponse
func (client *Client) ListModelFeaturesWithContext(ctx context.Context, InstanceId *string, tmpReq *ListModelFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelFeaturesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListModelFeaturesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModelFeatureIds) {
		request.ModelFeatureIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelFeatureIds, dara.String("ModelFeatureIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelFeatureIdsShrink) {
		query["ModelFeatureIds"] = request.ModelFeatureIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelFeatures"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下的所有特征视图、特征信息。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectFeatureViewsResponse
func (client *Client) ListProjectFeatureViewsWithContext(ctx context.Context, InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectFeatureViewsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectFeatureViews"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/featureviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectFeatureViewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下所有特征信息
//
// @param request - ListProjectFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectFeaturesResponse
func (client *Client) ListProjectFeaturesWithContext(ctx context.Context, InstanceId *string, ProjectId *string, request *ListProjectFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectFeaturesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectFeatures"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId)) + "/features"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectFeaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Feature Store项目列表。
//
// @param tmpReq - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, InstanceId *string, tmpReq *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProjectIds) {
		request.ProjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProjectIds, dara.String("ProjectIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectIdsShrink) {
		query["ProjectIds"] = request.ProjectIdsShrink
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects"),
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
// 获取任务日志列表
//
// @param request - ListTaskLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskLogsResponse
func (client *Client) ListTaskLogsWithContext(ctx context.Context, InstanceId *string, TaskId *string, request *ListTaskLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTaskLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskLogs"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务列表
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, InstanceId *string, tmpReq *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/tasks"),
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
// 将特征视图的离线数据发布/同步到线上。
//
// @param request - PublishFeatureViewTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFeatureViewTableResponse
func (client *Client) PublishFeatureViewTableWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, request *PublishFeatureViewTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishFeatureViewTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.EventTime) {
		body["EventTime"] = request.EventTime
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OfflineToOnline) {
		body["OfflineToOnline"] = request.OfflineToOnline
	}

	if !dara.IsNil(request.Partitions) {
		body["Partitions"] = request.Partitions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishFeatureViewTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId)) + "/action/publishtable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishFeatureViewTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止任务。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskResponse
func (client *Client) StopTaskWithContext(ctx context.Context, InstanceId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTask"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/tasks/" + dara.PercentEncode(dara.StringValue(TaskId)) + "/action/stop"),
		Method:      dara.String("POST"),
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
// 更新数据源信息。
//
// @param request - UpdateDatasourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasourceResponse
func (client *Client) UpdateDatasourceWithContext(ctx context.Context, InstanceId *string, DatasourceId *string, request *UpdateDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasource"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/datasources/" + dara.PercentEncode(dara.StringValue(DatasourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新大模型调用信息配置
//
// @param request - UpdateLLMConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLLMConfigResponse
func (client *Client) UpdateLLMConfigWithContext(ctx context.Context, InstanceId *string, LLMConfigId *string, request *UpdateLLMConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLLMConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BaseUrl) {
		body["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.BatchSize) {
		body["BatchSize"] = request.BatchSize
	}

	if !dara.IsNil(request.MaxTokens) {
		body["MaxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Rps) {
		body["Rps"] = request.Rps
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLLMConfig"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/llmconfigs/" + dara.PercentEncode(dara.StringValue(LLMConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLLMConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新label表。
//
// @param request - UpdateLabelTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLabelTableResponse
func (client *Client) UpdateLabelTableWithContext(ctx context.Context, InstanceId *string, LabelTableId *string, request *UpdateLabelTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLabelTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceId) {
		body["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLabelTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labeltables/" + dara.PercentEncode(dara.StringValue(LabelTableId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLabelTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型特征。
//
// @param request - UpdateModelFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelFeatureResponse
func (client *Client) UpdateModelFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Features) {
		body["Features"] = request.Features
	}

	if !dara.IsNil(request.LabelPriorityLevel) {
		body["LabelPriorityLevel"] = request.LabelPriorityLevel
	}

	if !dara.IsNil(request.LabelTableId) {
		body["LabelTableId"] = request.LabelTableId
	}

	if !dara.IsNil(request.SequenceFeatureViewIds) {
		body["SequenceFeatureViewIds"] = request.SequenceFeatureViewIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模型特征的FG特征配置信息。
//
// @param request - UpdateModelFeatureFGFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelFeatureFGFeatureResponse
func (client *Client) UpdateModelFeatureFGFeatureWithContext(ctx context.Context, InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureFGFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelFeatureFGFeatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LookupFeatures) {
		body["LookupFeatures"] = request.LookupFeatures
	}

	if !dara.IsNil(request.RawFeatures) {
		body["RawFeatures"] = request.RawFeatures
	}

	if !dara.IsNil(request.Reserves) {
		body["Reserves"] = request.Reserves
	}

	if !dara.IsNil(request.SequenceFeatures) {
		body["SequenceFeatures"] = request.SequenceFeatures
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelFeatureFGFeature"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/modelfeatures/" + dara.PercentEncode(dara.StringValue(ModelFeatureId)) + "/fgfeature"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelFeatureFGFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定Feature Store项目信息。
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, InstanceId *string, ProjectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/projects/" + dara.PercentEncode(dara.StringValue(ProjectId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取特征视图血缘关系。
//
// @param request - WriteFeatureViewTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WriteFeatureViewTableResponse
func (client *Client) WriteFeatureViewTableWithContext(ctx context.Context, InstanceId *string, FeatureViewId *string, request *WriteFeatureViewTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WriteFeatureViewTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Partitions) {
		body["Partitions"] = request.Partitions
	}

	if !dara.IsNil(request.UrlDatasource) {
		body["UrlDatasource"] = request.UrlDatasource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WriteFeatureViewTable"),
		Version:     dara.String("2023-06-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/featureviews/" + dara.PercentEncode(dara.StringValue(FeatureViewId)) + "/action/writetable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WriteFeatureViewTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
