// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("paifeaturestore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
func (client *Client) CheckInstanceDatasourceWithOptions(InstanceId *string, request *CheckInstanceDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceDatasourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测资源连接状态。
//
// @param request - CheckInstanceDatasourceRequest
//
// @return CheckInstanceDatasourceResponse
func (client *Client) CheckInstanceDatasource(InstanceId *string, request *CheckInstanceDatasourceRequest) (_result *CheckInstanceDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceDatasourceResponse{}
	_body, _err := client.CheckInstanceDatasourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckModelFeatureFGFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckModelFeatureFGFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckModelFeatureFGFeatureResponse
func (client *Client) CheckModelFeatureFGFeature(InstanceId *string, ModelFeatureId *string) (_result *CheckModelFeatureFGFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckModelFeatureFGFeatureResponse{}
	_body, _err := client.CheckModelFeatureFGFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDatasourceWithOptions(InstanceId *string, request *CreateDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDatasourceResponse
func (client *Client) CreateDatasource(InstanceId *string, request *CreateDatasourceRequest) (_result *CreateDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasourceResponse{}
	_body, _err := client.CreateDatasourceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateFeatureEntityWithOptions(InstanceId *string, request *CreateFeatureEntityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureEntityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateFeatureEntityResponse
func (client *Client) CreateFeatureEntity(InstanceId *string, request *CreateFeatureEntityRequest) (_result *CreateFeatureEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureEntityResponse{}
	_body, _err := client.CreateFeatureEntityWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateFeatureViewWithOptions(InstanceId *string, request *CreateFeatureViewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFeatureViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateFeatureViewResponse
func (client *Client) CreateFeatureView(InstanceId *string, request *CreateFeatureViewRequest) (_result *CreateFeatureViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFeatureViewResponse{}
	_body, _err := client.CreateFeatureViewWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLLMConfigWithOptions(InstanceId *string, request *CreateLLMConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLLMConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateLLMConfigResponse
func (client *Client) CreateLLMConfig(InstanceId *string, request *CreateLLMConfigRequest) (_result *CreateLLMConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLLMConfigResponse{}
	_body, _err := client.CreateLLMConfigWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLabelTableWithOptions(InstanceId *string, request *CreateLabelTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLabelTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateLabelTableResponse
func (client *Client) CreateLabelTable(InstanceId *string, request *CreateLabelTableRequest) (_result *CreateLabelTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLabelTableResponse{}
	_body, _err := client.CreateLabelTableWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateModelFeatureWithOptions(InstanceId *string, request *CreateModelFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateModelFeatureResponse
func (client *Client) CreateModelFeature(InstanceId *string, request *CreateModelFeatureRequest) (_result *CreateModelFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelFeatureResponse{}
	_body, _err := client.CreateModelFeatureWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateProjectWithOptions(InstanceId *string, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateProjectResponse
func (client *Client) CreateProject(InstanceId *string, request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceIdentityRoleWithOptions(request *CreateServiceIdentityRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRole(request *CreateServiceIdentityRoleRequest) (_result *CreateServiceIdentityRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CreateServiceIdentityRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDatasourceWithOptions(InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDatasourceResponse
func (client *Client) DeleteDatasource(InstanceId *string, DatasourceId *string) (_result *DeleteDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasourceResponse{}
	_body, _err := client.DeleteDatasourceWithOptions(InstanceId, DatasourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFeatureEntityWithOptions(InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFeatureEntityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFeatureEntityResponse
func (client *Client) DeleteFeatureEntity(InstanceId *string, FeatureEntityId *string) (_result *DeleteFeatureEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFeatureEntityResponse{}
	_body, _err := client.DeleteFeatureEntityWithOptions(InstanceId, FeatureEntityId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFeatureViewWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFeatureViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFeatureViewResponse
func (client *Client) DeleteFeatureView(InstanceId *string, FeatureViewId *string) (_result *DeleteFeatureViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFeatureViewResponse{}
	_body, _err := client.DeleteFeatureViewWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteLLMConfigWithOptions(InstanceId *string, LLMConfigId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLLMConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteLLMConfigResponse
func (client *Client) DeleteLLMConfig(InstanceId *string, LLMConfigId *string) (_result *DeleteLLMConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLLMConfigResponse{}
	_body, _err := client.DeleteLLMConfigWithOptions(InstanceId, LLMConfigId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteLabelTableWithOptions(InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLabelTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteLabelTableResponse
func (client *Client) DeleteLabelTable(InstanceId *string, LabelTableId *string) (_result *DeleteLabelTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLabelTableResponse{}
	_body, _err := client.DeleteLabelTableWithOptions(InstanceId, LabelTableId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteModelFeatureResponse
func (client *Client) DeleteModelFeature(InstanceId *string, ModelFeatureId *string) (_result *DeleteModelFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelFeatureResponse{}
	_body, _err := client.DeleteModelFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteProjectWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteProjectResponse
func (client *Client) DeleteProject(InstanceId *string, ProjectId *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportModelFeatureTrainingSetTableWithOptions(InstanceId *string, ModelFeatureId *string, request *ExportModelFeatureTrainingSetTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportModelFeatureTrainingSetTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportModelFeatureTrainingSetTableResponse
func (client *Client) ExportModelFeatureTrainingSetTable(InstanceId *string, ModelFeatureId *string, request *ExportModelFeatureTrainingSetTableRequest) (_result *ExportModelFeatureTrainingSetTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportModelFeatureTrainingSetTableResponse{}
	_body, _err := client.ExportModelFeatureTrainingSetTableWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatasourceWithOptions(InstanceId *string, DatasourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatasourceResponse
func (client *Client) GetDatasource(InstanceId *string, DatasourceId *string) (_result *GetDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasourceResponse{}
	_body, _err := client.GetDatasourceWithOptions(InstanceId, DatasourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatasourceTableWithOptions(InstanceId *string, DatasourceId *string, TableName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasourceTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatasourceTableResponse
func (client *Client) GetDatasourceTable(InstanceId *string, DatasourceId *string, TableName *string) (_result *GetDatasourceTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasourceTableResponse{}
	_body, _err := client.GetDatasourceTableWithOptions(InstanceId, DatasourceId, TableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFeatureEntityWithOptions(InstanceId *string, FeatureEntityId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureEntityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFeatureEntityResponse
func (client *Client) GetFeatureEntity(InstanceId *string, FeatureEntityId *string) (_result *GetFeatureEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureEntityResponse{}
	_body, _err := client.GetFeatureEntityWithOptions(InstanceId, FeatureEntityId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFeatureViewWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFeatureViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFeatureViewResponse
func (client *Client) GetFeatureView(InstanceId *string, FeatureViewId *string) (_result *GetFeatureViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFeatureViewResponse{}
	_body, _err := client.GetFeatureViewWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceResponse
func (client *Client) GetInstance(InstanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetLLMConfigWithOptions(InstanceId *string, LLMConfigId *string, RegionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLLMConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetLLMConfigResponse
func (client *Client) GetLLMConfig(InstanceId *string, LLMConfigId *string, RegionId *string) (_result *GetLLMConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLLMConfigResponse{}
	_body, _err := client.GetLLMConfigWithOptions(InstanceId, LLMConfigId, RegionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetLabelTableWithOptions(InstanceId *string, LabelTableId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLabelTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetLabelTableResponse
func (client *Client) GetLabelTable(InstanceId *string, LabelTableId *string) (_result *GetLabelTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLabelTableResponse{}
	_body, _err := client.GetLabelTableWithOptions(InstanceId, LabelTableId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetModelFeatureResponse
func (client *Client) GetModelFeature(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureResponse{}
	_body, _err := client.GetModelFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetModelFeatureFGFeatureWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureFGFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetModelFeatureFGFeatureResponse
func (client *Client) GetModelFeatureFGFeature(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureFGFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureFGFeatureResponse{}
	_body, _err := client.GetModelFeatureFGFeatureWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetModelFeatureFGInfoWithOptions(InstanceId *string, ModelFeatureId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelFeatureFGInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetModelFeatureFGInfoResponse
func (client *Client) GetModelFeatureFGInfo(InstanceId *string, ModelFeatureId *string) (_result *GetModelFeatureFGInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelFeatureFGInfoResponse{}
	_body, _err := client.GetModelFeatureFGInfoWithOptions(InstanceId, ModelFeatureId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProjectWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProjectResponse
func (client *Client) GetProject(InstanceId *string, ProjectId *string) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProjectFeatureEntityWithOptions(InstanceId *string, ProjectId *string, FeatureEntityName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectFeatureEntityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProjectFeatureEntityResponse
func (client *Client) GetProjectFeatureEntity(InstanceId *string, ProjectId *string, FeatureEntityName *string) (_result *GetProjectFeatureEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectFeatureEntityResponse{}
	_body, _err := client.GetProjectFeatureEntityWithOptions(InstanceId, ProjectId, FeatureEntityName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceIdentityRoleWithOptions(RoleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRole(RoleName *string) (_result *GetServiceIdentityRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.GetServiceIdentityRoleWithOptions(RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTaskWithOptions(InstanceId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTaskResponse
func (client *Client) GetTask(InstanceId *string, TaskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(InstanceId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasourceFeatureViewsWithOptions(InstanceId *string, DatasourceId *string, request *ListDatasourceFeatureViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourceFeatureViewsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatasourceFeatureViewsResponse
func (client *Client) ListDatasourceFeatureViews(InstanceId *string, DatasourceId *string, request *ListDatasourceFeatureViewsRequest) (_result *ListDatasourceFeatureViewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasourceFeatureViewsResponse{}
	_body, _err := client.ListDatasourceFeatureViewsWithOptions(InstanceId, DatasourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasourceTablesWithOptions(InstanceId *string, DatasourceId *string, request *ListDatasourceTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourceTablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatasourceTablesResponse
func (client *Client) ListDatasourceTables(InstanceId *string, DatasourceId *string, request *ListDatasourceTablesRequest) (_result *ListDatasourceTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasourceTablesResponse{}
	_body, _err := client.ListDatasourceTablesWithOptions(InstanceId, DatasourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasourcesWithOptions(InstanceId *string, request *ListDatasourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatasourcesResponse
func (client *Client) ListDatasources(InstanceId *string, request *ListDatasourcesRequest) (_result *ListDatasourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasourcesResponse{}
	_body, _err := client.ListDatasourcesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFeatureEntitiesWithOptions(InstanceId *string, tmpReq *ListFeatureEntitiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureEntitiesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFeatureEntitiesRequest
//
// @return ListFeatureEntitiesResponse
func (client *Client) ListFeatureEntities(InstanceId *string, request *ListFeatureEntitiesRequest) (_result *ListFeatureEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureEntitiesResponse{}
	_body, _err := client.ListFeatureEntitiesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFeatureViewFieldRelationshipsWithOptions(InstanceId *string, FeatureViewId *string, FieldName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewFieldRelationshipsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListFeatureViewFieldRelationshipsResponse
func (client *Client) ListFeatureViewFieldRelationships(InstanceId *string, FeatureViewId *string, FieldName *string) (_result *ListFeatureViewFieldRelationshipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewFieldRelationshipsResponse{}
	_body, _err := client.ListFeatureViewFieldRelationshipsWithOptions(InstanceId, FeatureViewId, FieldName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFeatureViewOnlineFeaturesWithOptions(InstanceId *string, FeatureViewId *string, tmpReq *ListFeatureViewOnlineFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewOnlineFeaturesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFeatureViewOnlineFeaturesRequest
//
// @return ListFeatureViewOnlineFeaturesResponse
func (client *Client) ListFeatureViewOnlineFeatures(InstanceId *string, FeatureViewId *string, request *ListFeatureViewOnlineFeaturesRequest) (_result *ListFeatureViewOnlineFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewOnlineFeaturesResponse{}
	_body, _err := client.ListFeatureViewOnlineFeaturesWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFeatureViewRelationshipsWithOptions(InstanceId *string, FeatureViewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewRelationshipsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListFeatureViewRelationshipsResponse
func (client *Client) ListFeatureViewRelationships(InstanceId *string, FeatureViewId *string) (_result *ListFeatureViewRelationshipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewRelationshipsResponse{}
	_body, _err := client.ListFeatureViewRelationshipsWithOptions(InstanceId, FeatureViewId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFeatureViewsWithOptions(InstanceId *string, tmpReq *ListFeatureViewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeatureViewsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFeatureViewsRequest
//
// @return ListFeatureViewsResponse
func (client *Client) ListFeatureViews(InstanceId *string, request *ListFeatureViewsRequest) (_result *ListFeatureViewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeatureViewsResponse{}
	_body, _err := client.ListFeatureViewsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListLLMConfigsWithOptions(InstanceId *string, request *ListLLMConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLLMConfigsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListLLMConfigsResponse
func (client *Client) ListLLMConfigs(InstanceId *string, request *ListLLMConfigsRequest) (_result *ListLLMConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLLMConfigsResponse{}
	_body, _err := client.ListLLMConfigsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListLabelTablesWithOptions(InstanceId *string, tmpReq *ListLabelTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLabelTablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListLabelTablesRequest
//
// @return ListLabelTablesResponse
func (client *Client) ListLabelTables(InstanceId *string, request *ListLabelTablesRequest) (_result *ListLabelTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLabelTablesResponse{}
	_body, _err := client.ListLabelTablesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListModelFeatureAvailableFeaturesWithOptions(InstanceId *string, ModelFeatureId *string, request *ListModelFeatureAvailableFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelFeatureAvailableFeaturesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListModelFeatureAvailableFeaturesResponse
func (client *Client) ListModelFeatureAvailableFeatures(InstanceId *string, ModelFeatureId *string, request *ListModelFeatureAvailableFeaturesRequest) (_result *ListModelFeatureAvailableFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelFeatureAvailableFeaturesResponse{}
	_body, _err := client.ListModelFeatureAvailableFeaturesWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListModelFeaturesWithOptions(InstanceId *string, tmpReq *ListModelFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelFeaturesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListModelFeaturesRequest
//
// @return ListModelFeaturesResponse
func (client *Client) ListModelFeatures(InstanceId *string, request *ListModelFeaturesRequest) (_result *ListModelFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelFeaturesResponse{}
	_body, _err := client.ListModelFeaturesWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProjectFeatureViewsWithOptions(InstanceId *string, ProjectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectFeatureViewsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProjectFeatureViewsResponse
func (client *Client) ListProjectFeatureViews(InstanceId *string, ProjectId *string) (_result *ListProjectFeatureViewsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectFeatureViewsResponse{}
	_body, _err := client.ListProjectFeatureViewsWithOptions(InstanceId, ProjectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProjectFeaturesWithOptions(InstanceId *string, ProjectId *string, request *ListProjectFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectFeaturesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProjectFeaturesResponse
func (client *Client) ListProjectFeatures(InstanceId *string, ProjectId *string, request *ListProjectFeaturesRequest) (_result *ListProjectFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectFeaturesResponse{}
	_body, _err := client.ListProjectFeaturesWithOptions(InstanceId, ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProjectsWithOptions(InstanceId *string, tmpReq *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(InstanceId *string, request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTaskLogsWithOptions(InstanceId *string, TaskId *string, request *ListTaskLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTaskLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTaskLogsResponse
func (client *Client) ListTaskLogs(InstanceId *string, TaskId *string, request *ListTaskLogsRequest) (_result *ListTaskLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTaskLogsResponse{}
	_body, _err := client.ListTaskLogsWithOptions(InstanceId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTasksWithOptions(InstanceId *string, tmpReq *ListTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(InstanceId *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PublishFeatureViewTableWithOptions(InstanceId *string, FeatureViewId *string, request *PublishFeatureViewTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishFeatureViewTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PublishFeatureViewTableResponse
func (client *Client) PublishFeatureViewTable(InstanceId *string, FeatureViewId *string, request *PublishFeatureViewTableRequest) (_result *PublishFeatureViewTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishFeatureViewTableResponse{}
	_body, _err := client.PublishFeatureViewTableWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopTaskWithOptions(InstanceId *string, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopTaskResponse
func (client *Client) StopTask(InstanceId *string, TaskId *string) (_result *StopTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTaskResponse{}
	_body, _err := client.StopTaskWithOptions(InstanceId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateDatasourceWithOptions(InstanceId *string, DatasourceId *string, request *UpdateDatasourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateDatasourceResponse
func (client *Client) UpdateDatasource(InstanceId *string, DatasourceId *string, request *UpdateDatasourceRequest) (_result *UpdateDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasourceResponse{}
	_body, _err := client.UpdateDatasourceWithOptions(InstanceId, DatasourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateLLMConfigWithOptions(InstanceId *string, LLMConfigId *string, request *UpdateLLMConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLLMConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateLLMConfigResponse
func (client *Client) UpdateLLMConfig(InstanceId *string, LLMConfigId *string, request *UpdateLLMConfigRequest) (_result *UpdateLLMConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLLMConfigResponse{}
	_body, _err := client.UpdateLLMConfigWithOptions(InstanceId, LLMConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateLabelTableWithOptions(InstanceId *string, LabelTableId *string, request *UpdateLabelTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLabelTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateLabelTableResponse
func (client *Client) UpdateLabelTable(InstanceId *string, LabelTableId *string, request *UpdateLabelTableRequest) (_result *UpdateLabelTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLabelTableResponse{}
	_body, _err := client.UpdateLabelTableWithOptions(InstanceId, LabelTableId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateModelFeatureWithOptions(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateModelFeatureResponse
func (client *Client) UpdateModelFeature(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureRequest) (_result *UpdateModelFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelFeatureResponse{}
	_body, _err := client.UpdateModelFeatureWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateModelFeatureFGFeatureWithOptions(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureFGFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelFeatureFGFeatureResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateModelFeatureFGFeatureResponse
func (client *Client) UpdateModelFeatureFGFeature(InstanceId *string, ModelFeatureId *string, request *UpdateModelFeatureFGFeatureRequest) (_result *UpdateModelFeatureFGFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelFeatureFGFeatureResponse{}
	_body, _err := client.UpdateModelFeatureFGFeatureWithOptions(InstanceId, ModelFeatureId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateProjectWithOptions(InstanceId *string, ProjectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateProjectResponse
func (client *Client) UpdateProject(InstanceId *string, ProjectId *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(InstanceId, ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) WriteFeatureViewTableWithOptions(InstanceId *string, FeatureViewId *string, request *WriteFeatureViewTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WriteFeatureViewTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return WriteFeatureViewTableResponse
func (client *Client) WriteFeatureViewTable(InstanceId *string, FeatureViewId *string, request *WriteFeatureViewTableRequest) (_result *WriteFeatureViewTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WriteFeatureViewTableResponse{}
	_body, _err := client.WriteFeatureViewTableWithOptions(InstanceId, FeatureViewId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
