// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Binds a custom analyzer to an Elasticsearch instance.
//
// @param request - BindESUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindESUserAnalyzerResponse
func (client *Client) BindESUserAnalyzerWithContext(ctx context.Context, appGroupIdentity *string, esInstanceId *string, request *BindESUserAnalyzerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindESUserAnalyzerResponse, _err error) {
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
		Action:      dara.String("BindESUserAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/es/" + dara.PercentEncode(dara.StringValue(esInstanceId)) + "/actions/bind-analyzer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindESUserAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an Elasticsearch instance.
//
// @param request - BindEsInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEsInstanceResponse
func (client *Client) BindEsInstanceWithContext(ctx context.Context, appGroupIdentity *string, request *BindEsInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindEsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("BindEsInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/actions/bind-es-instance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindEsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Compiles a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompileSortScriptResponse
func (client *Client) CompileSortScriptWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompileSortScriptResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompileSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/actions/compiling"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompileSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an experiment.
//
// @param request - CreateABTestExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestExperimentResponse
func (client *Client) CreateABTestExperimentWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, request *CreateABTestExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABTestExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABTestExperiment"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABTestExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a test group.
//
// @param request - CreateABTestGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestGroupResponse
func (client *Client) CreateABTestGroupWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, request *CreateABTestGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABTestGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABTestGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABTestGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an scenario.
//
// @param request - CreateABTestSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestSceneResponse
func (client *Client) CreateABTestSceneWithContext(ctx context.Context, appGroupIdentity *string, request *CreateABTestSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateABTestSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateABTestScene"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateABTestSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version for an OpenSearch application.
//
// Description:
//
//	  When you create a standard application, a new version of the application is created if the specified application name already exists.
//
//		- When you create a version of an existing application, you must specify the autoSwitch and realtimeShared parameters.
//
//		- When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
//
//		- When you create a version of an existing application, the modification of the value of the quota parameter does not take effect.
//
// @param request - CreateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithContext(ctx context.Context, appGroupIdentity *string, request *CreateAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoSwitch) {
		body["autoSwitch"] = request.AutoSwitch
	}

	if !dara.IsNil(request.Cluster) {
		body["cluster"] = request.Cluster
	}

	if !dara.IsNil(request.ConfigItems) {
		body["configItems"] = request.ConfigItems
	}

	if !dara.IsNil(request.DataSources) {
		body["dataSources"] = request.DataSources
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.FetchFields) {
		body["fetchFields"] = request.FetchFields
	}

	if !dara.IsNil(request.FirstRanks) {
		body["firstRanks"] = request.FirstRanks
	}

	if !dara.IsNil(request.Interpretations) {
		body["interpretations"] = request.Interpretations
	}

	if !dara.IsNil(request.NetworkType) {
		body["networkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Prompts) {
		body["prompts"] = request.Prompts
	}

	if !dara.IsNil(request.QueryProcessors) {
		body["queryProcessors"] = request.QueryProcessors
	}

	if !dara.IsNil(request.RealtimeShared) {
		body["realtimeShared"] = request.RealtimeShared
	}

	if !dara.IsNil(request.Schema) {
		body["schema"] = request.Schema
	}

	if !dara.IsNil(request.Schemas) {
		body["schemas"] = request.Schemas
	}

	if !dara.IsNil(request.SecondRanks) {
		body["secondRanks"] = request.SecondRanks
	}

	if !dara.IsNil(request.Summaries) {
		body["summaries"] = request.Summaries
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an OpenSearch application.
//
// @param request - CreateAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupResponse
func (client *Client) CreateAppGroupWithContext(ctx context.Context, request *CreateAppGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAppGroupCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupCredentialsResponse
func (client *Client) CreateAppGroupCredentialsWithContext(ctx context.Context, appGroupIdentity *string, request *CreateAppGroupCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppGroupCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppGroupCredentials"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/credentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppGroupCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rough sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified rough sort expression. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateFirstRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirstRankResponse
func (client *Client) CreateFirstRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *CreateFirstRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFirstRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFirstRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/first-ranks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFirstRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an algorithm instance of a feature.
//
// Description:
//
// You can call the [GetFunctionCurrentVersion](https://help.aliyun.com/document_detail/421377.html) operation to query the latest version of a feature. The response of the operation includes the createParameters parameter that is used to create an algorithm instance, the usageParameters parameter, and the requirements for setting these parameters.
//
// @param request - CreateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionInstanceResponse
func (client *Client) CreateFunctionInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, request *CreateFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateParameters) {
		body["createParameters"] = request.CreateParameters
	}

	if !dara.IsNil(request.Cron) {
		body["cron"] = request.Cron
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.FunctionType) {
		body["functionType"] = request.FunctionType
	}

	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.UsageParameters) {
		body["usageParameters"] = request.UsageParameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunctionInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an algorithm resource for a specific feature.
//
// @param request - CreateFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResourceResponse
func (client *Client) CreateFunctionResourceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, request *CreateFunctionResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ResourceName) {
		body["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunctionResource"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a training task for an algorithm instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionTaskResponse
func (client *Client) CreateFunctionTaskWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunctionTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an intervention dictionary.
//
// @param request - CreateInterventionDictionaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInterventionDictionaryResponse
func (client *Client) CreateInterventionDictionaryWithContext(ctx context.Context, request *CreateInterventionDictionaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInterventionDictionaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalyzerType) {
		body["analyzerType"] = request.AnalyzerType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
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
		Action:      dara.String("CreateInterventionDictionary"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInterventionDictionaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a query analysis rule. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateQueryProcessorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueryProcessorResponse
func (client *Client) CreateQueryProcessorWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *CreateQueryProcessorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQueryProcessorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueryProcessor"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueryProcessorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled task for an OpenSearch application.
//
// @param request - CreateScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
func (client *Client) CreateScheduledTaskWithContext(ctx context.Context, appGroupIdentity *string, request *CreateScheduledTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scheduled-tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a query policy.
//
// @param request - CreateSearchStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchStrategyResponse
func (client *Client) CreateSearchStrategyWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *CreateSearchStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSearchStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSearchStrategy"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/search-strategies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSearchStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a fine sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified fine sort expression. The default value of dryRun is false if you do not set this parameter.
//
// @param request - CreateSecondRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecondRankResponse
func (client *Client) CreateSecondRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *CreateSecondRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSecondRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecondRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/second-ranks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecondRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sort script.
//
// @param request - CreateSortScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSortScriptResponse
func (client *Client) CreateSortScriptWithContext(ctx context.Context, appGroupIdentity *string, appVersionId *string, request *CreateSortScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSortScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ScriptName) {
		body["scriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom analyzer.
//
// @param request - CreateUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserAnalyzerResponse
func (client *Client) CreateUserAnalyzerWithContext(ctx context.Context, request *CreateUserAnalyzerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserAnalyzerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Business) {
		body["business"] = request.Business
	}

	if !dara.IsNil(request.BusinessAppGroupId) {
		body["businessAppGroupId"] = request.BusinessAppGroupId
	}

	if !dara.IsNil(request.BusinessType) {
		body["businessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
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
		Action:      dara.String("CreateUserAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a test.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestExperimentResponse
func (client *Client) DeleteABTestExperimentWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABTestExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABTestExperiment"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments/" + dara.PercentEncode(dara.StringValue(experimentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABTestExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestGroupResponse
func (client *Client) DeleteABTestGroupWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABTestGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABTestGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABTestGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an A/B test scenario.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestSceneResponse
func (client *Client) DeleteABTestSceneWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteABTestSceneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteABTestScene"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteABTestSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an algorithm instance. Before you delete an instance, make sure that it is not in use to prevent service interruptions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionInstanceResponse
func (client *Client) DeleteFunctionInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunctionInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an algorithm resource.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResourceResponse
func (client *Client) DeleteFunctionResourceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, resourceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunctionResource"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/resources/" + dara.PercentEncode(dara.StringValue(resourceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a training task. The training task in progress cannot be deleted.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionTaskResponse
func (client *Client) DeleteFunctionTaskWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, generation *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunctionTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/tasks/" + dara.PercentEncode(dara.StringValue(generation))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除排序脚本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSortScriptResponse
func (client *Client) DeleteSortScriptWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSortScriptResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a script file.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSortScriptFileResponse
func (client *Client) DeleteSortScriptFileWithContext(ctx context.Context, appGroupIdentity *string, appVersionId *string, scriptName *string, fileName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSortScriptFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSortScriptFile"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/files/src/" + dara.PercentEncode(dara.StringValue(fileName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSortScriptFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestExperimentResponse
func (client *Client) DescribeABTestExperimentWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeABTestExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeABTestExperiment"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments/" + dara.PercentEncode(dara.StringValue(experimentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeABTestExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a test group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestGroupResponse
func (client *Client) DescribeABTestGroupWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeABTestGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeABTestGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeABTestGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an A/B test scenario.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestSceneResponse
func (client *Client) DescribeABTestSceneWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeABTestSceneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeABTestScene"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeABTestSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppResponse
func (client *Client) DescribeAppWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApp"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppGroupResponse
func (client *Client) DescribeAppGroupWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics about a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStatisticsResponse
func (client *Client) DescribeAppStatisticsWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppStatisticsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppStatistics"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/statistics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version list of an OpenSearch application.
//
// Description:
//
//	  When you create a standard application, a new version of the application is created if the specified application name already exists.
//
//		- When you create a version of an existing application, you must specify the autoSwitch and realtimeShared parameters.
//
//		- When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
//
//		- When you create a version of an existing application, the modification of the value of the quota parameter does not take effect.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data collection task of an application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCollctionResponse
func (client *Client) DescribeDataCollctionWithContext(ctx context.Context, appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDataCollctionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataCollction"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/data-collections/" + dara.PercentEncode(dara.StringValue(dataCollectionIdentity))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataCollctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a rough sort expression that is configured for an OpenSearch application version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirstRankResponse
func (client *Client) DescribeFirstRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeFirstRankResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirstRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/first-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirstRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an intervention dictionary.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterventionDictionaryResponse
func (client *Client) DescribeInterventionDictionaryWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInterventionDictionaryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterventionDictionary"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterventionDictionaryResponse{}
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
// @return DescribeQueryProcessorResponse
func (client *Client) DescribeQueryProcessorWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeQueryProcessorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQueryProcessor"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQueryProcessorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of all regions that support OpenSearch.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/regions"),
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

// Summary:
//
// 查看应用定时任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduledTaskResponse
func (client *Client) DescribeScheduledTaskWithContext(ctx context.Context, appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScheduledTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScheduledTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scheduled-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a fine sort expression that is configured for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecondRankResponse
func (client *Client) DescribeSecondRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSecondRankResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecondRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/second-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecondRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取优化大师慢查询开通状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowQueryStatusResponse
func (client *Client) DescribeSlowQueryStatusWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSlowQueryStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowQueryStatus"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowQueryStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义分析器详情
//
// @param request - DescribeUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAnalyzerResponse
func (client *Client) DescribeUserAnalyzerWithContext(ctx context.Context, name *string, request *DescribeUserAnalyzerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeUserAnalyzerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.With) {
		query["with"] = request.With
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用优化大师慢查询服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSlowQueryResponse
func (client *Client) DisableSlowQueryWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableSlowQueryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSlowQuery"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query/actions/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSlowQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables slow query optimization of Optimization Master.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSlowQueryResponse
func (client *Client) EnableSlowQueryWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableSlowQueryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSlowQuery"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query/actions/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSlowQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a wide table that is generated after a JOIN operation is performed on multiple tables.
//
// @param request - GenerateMergedTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateMergedTableResponse
func (client *Client) GenerateMergedTableWithContext(ctx context.Context, request *GenerateMergedTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateMergedTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Spec) {
		query["spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateMergedTable"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/assist/schema/actions/merge"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateMergedTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the type of an industry.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithContext(ctx context.Context, domainName *string, request *GetDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroupIdentity) {
		query["appGroupIdentity"] = request.AppGroupIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version information about the current feature when you create an instance.
//
// @param request - GetFunctionCurrentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionCurrentVersionResponse
func (client *Client) GetFunctionCurrentVersionWithContext(ctx context.Context, functionName *string, request *GetFunctionCurrentVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionCurrentVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.Domain) {
		query["domain"] = request.Domain
	}

	if !dara.IsNil(request.FunctionType) {
		query["functionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionCurrentVersion"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/current-version"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionCurrentVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the algorithm instance that an application uses by default.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionDefaultInstanceResponse
func (client *Client) GetFunctionDefaultInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionDefaultInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionDefaultInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/default-instance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an algorithm instance by instance name.
//
// @param request - GetFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionInstanceResponse
func (client *Client) GetFunctionInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an algorithm resource.
//
// @param request - GetFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResourceResponse
func (client *Client) GetFunctionResourceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, resourceName *string, request *GetFunctionResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionResource"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/resources/" + dara.PercentEncode(dara.StringValue(resourceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a training task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionTaskResponse
func (client *Client) GetFunctionTaskWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, generation *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/tasks/" + dara.PercentEncode(dara.StringValue(generation))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries version information by version ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionVersionResponse
func (client *Client) GetFunctionVersionWithContext(ctx context.Context, functionName *string, versionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionVersion"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/versions/" + dara.PercentEncode(dara.StringValue(versionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionVersionResponse{}
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
// @return GetScriptFileNamesResponse
func (client *Client) GetScriptFileNamesWithContext(ctx context.Context, appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScriptFileNamesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScriptFileNames"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/file-names"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScriptFileNamesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a query policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSearchStrategyResponse
func (client *Client) GetSearchStrategyWithContext(ctx context.Context, appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSearchStrategyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSearchStrategy"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/search-strategies/" + dara.PercentEncode(dara.StringValue(strategyName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSearchStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSortScriptResponse
func (client *Client) GetSortScriptWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSortScriptResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSortScriptFileResponse
func (client *Client) GetSortScriptFileWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSortScriptFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSortScriptFile"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/files/src/" + dara.PercentEncode(dara.StringValue(fileName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSortScriptFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of experiments.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestExperimentsResponse
func (client *Client) ListABTestExperimentsWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABTestExperimentsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABTestExperiments"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABTestExperimentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whitelists.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestFixedFlowDividersResponse
func (client *Client) ListABTestFixedFlowDividersWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABTestFixedFlowDividersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABTestFixedFlowDividers"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments/" + dara.PercentEncode(dara.StringValue(experimentId)) + "/fixed-flow-dividers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验组清单
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestGroupsResponse
func (client *Client) ListABTestGroupsWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABTestGroupsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABTestGroups"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABTestGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries test scenarios.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestScenesResponse
func (client *Client) ListABTestScenesWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListABTestScenesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListABTestScenes"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListABTestScenesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of OpenSearch applications.
//
// Description:
//
//	  This operation allows you to query applications by application name, instance ID, and application type.
//
//		- This operation allows you to sort the applications based on their creation time.
//
//		- This operation supports the parameters for paging.
//
// @param tmpReq - ListAppGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppGroupsResponse
func (client *Client) ListAppGroupsWithContext(ctx context.Context, tmpReq *ListAppGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAppGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
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

	if !dara.IsNil(request.SortBy) {
		query["sortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagsShrink) {
		query["tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppGroups"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data collection tasks of an OpenSearch application.
//
// @param request - ListDataCollectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCollectionsResponse
func (client *Client) ListDataCollectionsWithContext(ctx context.Context, appGroupIdentity *string, request *ListDataCollectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataCollectionsResponse, _err error) {
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
		Action:      dara.String("ListDataCollections"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/data-collections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataCollectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all fields in a table of a data source. This operation is for internal use only.
//
// @param request - ListDataSourceTableFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTableFieldsResponse
func (client *Client) ListDataSourceTableFieldsWithContext(ctx context.Context, dataSourceType *string, request *ListDataSourceTableFieldsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourceTableFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.RawType) {
		query["rawType"] = request.RawType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceTableFields"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/assist/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceType)) + "/fields"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceTableFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains all data from a specified data source.
//
// @param request - ListDataSourceTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTablesResponse
func (client *Client) ListDataSourceTablesWithContext(ctx context.Context, dataSourceType *string, request *ListDataSourceTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourceTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceTables"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/assist/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceType)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rough sort expressions that are configured for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFirstRanksResponse
func (client *Client) ListFirstRanksWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFirstRanksResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFirstRanks"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/first-ranks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFirstRanksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all algorithm instances of a user, which meet specified conditions.
//
// @param request - ListFunctionInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionInstancesResponse
func (client *Client) ListFunctionInstancesWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, request *ListFunctionInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionType) {
		query["functionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionInstances"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries algorithm resources.
//
// @param request - ListFunctionResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionResourcesResponse
func (client *Client) ListFunctionResourcesWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, request *ListFunctionResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionResources"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the training tasks. The returned results are sorted by start time in descending order.
//
// @param request - ListFunctionTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionTasksResponse
func (client *Client) ListFunctionTasksWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, request *ListFunctionTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionTasksResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionTasks"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户的干预词典列表
//
// @param request - ListInterventionDictionariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionariesResponse
func (client *Client) ListInterventionDictionariesWithContext(ctx context.Context, request *ListInterventionDictionariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInterventionDictionariesResponse, _err error) {
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

	if !dara.IsNil(request.Types) {
		query["types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterventionDictionaries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterventionDictionariesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intervention entries in an intervention dictionary.
//
// @param request - ListInterventionDictionaryEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryEntriesResponse
func (client *Client) ListInterventionDictionaryEntriesWithContext(ctx context.Context, name *string, request *ListInterventionDictionaryEntriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInterventionDictionaryEntriesResponse, _err error) {
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

	if !dara.IsNil(request.Word) {
		query["word"] = request.Word
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterventionDictionaryEntries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name)) + "/entries"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实体识别结果
//
// @param request - ListInterventionDictionaryNerResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryNerResultsResponse
func (client *Client) ListInterventionDictionaryNerResultsWithContext(ctx context.Context, name *string, request *ListInterventionDictionaryNerResultsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInterventionDictionaryNerResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterventionDictionaryNerResults"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name)) + "/ner-results"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterventionDictionaryNerResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated with an intervention dictionary. If the intervention dictionary is referenced by query analysis rules, this operation returns all applications that use the intervention dictionary and the information about the query analysis rules.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryRelatedEntitiesResponse
func (client *Client) ListInterventionDictionaryRelatedEntitiesWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInterventionDictionaryRelatedEntitiesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInterventionDictionaryRelatedEntities"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name)) + "/related"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInterventionDictionaryRelatedEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看当前的处理流
//
// @param request - ListProceedingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProceedingsResponse
func (client *Client) ListProceedingsWithContext(ctx context.Context, appGroupIdentity *string, request *ListProceedingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProceedingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterFinished) {
		query["filterFinished"] = request.FilterFinished
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProceedings"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/proceedings"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProceedingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a query analysis test. This API operation is available only to existing applications of OpenSearch Open Source Compatible Edition.
//
// @param request - ListQueryProcessorAnalyzerResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorAnalyzerResultsResponse
func (client *Client) ListQueryProcessorAnalyzerResultsWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, request *ListQueryProcessorAnalyzerResultsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQueryProcessorAnalyzerResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Text) {
		query["text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueryProcessorAnalyzerResults"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors/" + dara.PercentEncode(dara.StringValue(name)) + "/analyze"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueryProcessorAnalyzerResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recommended priority settings of entity types for named entity recognition (NER).
//
// @param request - ListQueryProcessorNersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorNersResponse
func (client *Client) ListQueryProcessorNersWithContext(ctx context.Context, request *ListQueryProcessorNersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQueryProcessorNersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListQueryProcessorNers"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/query-processor/ner/default-priorities"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueryProcessorNersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of query analysis rules that are configured for a version of an OpenSearch application.
//
// @param request - ListQueryProcessorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorsResponse
func (client *Client) ListQueryProcessorsWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *ListQueryProcessorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQueryProcessorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsActive) {
		query["isActive"] = request.IsActive
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQueryProcessors"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQueryProcessorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tickets that are submitted to apply for quotas for an OpenSearch application.
//
// @param request - ListQuotaReviewTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaReviewTasksResponse
func (client *Client) ListQuotaReviewTasksWithContext(ctx context.Context, appGroupIdentity *string, request *ListQuotaReviewTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotaReviewTasksResponse, _err error) {
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
		Action:      dara.String("ListQuotaReviewTasks"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/quota-review-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaReviewTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of scheduled tasks of an OpenSearch application.
//
// @param request - ListScheduledTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledTasksResponse
func (client *Client) ListScheduledTasksWithContext(ctx context.Context, appGroupIdentity *string, request *ListScheduledTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScheduledTasksResponse, _err error) {
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

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheduledTasks"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scheduled-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of query policies.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchStrategiesResponse
func (client *Client) ListSearchStrategiesWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchStrategiesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchStrategies"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/search-strategies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchStrategiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fine sort expressions that are configured for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecondRanksResponse
func (client *Client) ListSecondRanksWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSecondRanksResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecondRanks"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/second-ranks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecondRanksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the suggestions that are provided by Optimization Master for slow queries.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSlowQueryCategoriesResponse
func (client *Client) ListSlowQueryCategoriesWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSlowQueryCategoriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSlowQueryCategories"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query/categories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSlowQueryCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出优化大师慢查询Query清单
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSlowQueryQueriesResponse
func (client *Client) ListSlowQueryQueriesWithContext(ctx context.Context, appGroupIdentity *string, categoryIndex *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSlowQueryQueriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSlowQueryQueries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query/categories/" + dara.PercentEncode(dara.StringValue(categoryIndex)) + "/queries"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSlowQueryQueriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of sort expressions that are configured for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSortExpressionsResponse
func (client *Client) ListSortExpressionsWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSortExpressionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSortExpressions"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/sort-expressions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSortExpressionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all sort scripts of an application version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSortScriptsResponse
func (client *Client) ListSortScriptsWithContext(ctx context.Context, appGroupIdentity *string, appVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSortScriptsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSortScripts"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSortScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries log statistics, such as application error logs, hotword rankings, and slow query logs.
//
// @param request - ListStatisticLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatisticLogsResponse
func (client *Client) ListStatisticLogsWithContext(ctx context.Context, appGroupIdentity *string, moduleName *string, request *ListStatisticLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStatisticLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Columns) {
		query["columns"] = request.Columns
	}

	if !dara.IsNil(request.Distinct) {
		query["distinct"] = request.Distinct
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.SortBy) {
		query["sortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.StopTime) {
		query["stopTime"] = request.StopTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStatisticLogs"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/statistic-logs/" + dara.PercentEncode(dara.StringValue(moduleName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStatisticLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistical reports, such as application reports, drop-down suggestion reports, hotword shading reports, A/B test reports, and data quality reports.
//
// @param request - ListStatisticReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatisticReportResponse
func (client *Client) ListStatisticReportWithContext(ctx context.Context, appGroupIdentity *string, moduleName *string, request *ListStatisticReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListStatisticReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Columns) {
		query["columns"] = request.Columns
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStatisticReport"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/statistic-report/" + dara.PercentEncode(dara.StringValue(moduleName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStatisticReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tagged resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/resource-tags"),
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
// Queries the entries of a custom analyzer.
//
// @param request - ListUserAnalyzerEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAnalyzerEntriesResponse
func (client *Client) ListUserAnalyzerEntriesWithContext(ctx context.Context, name *string, request *ListUserAnalyzerEntriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserAnalyzerEntriesResponse, _err error) {
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

	if !dara.IsNil(request.Word) {
		query["word"] = request.Word
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserAnalyzerEntries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers/" + dara.PercentEncode(dara.StringValue(name)) + "/entries"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom analyzers that belong to the current account.
//
// @param request - ListUserAnalyzersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAnalyzersResponse
func (client *Client) ListUserAnalyzersWithContext(ctx context.Context, request *ListUserAnalyzersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserAnalyzersResponse, _err error) {
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
		Action:      dara.String("ListUserAnalyzers"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAnalyzersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of an OpenSearch application or sets the online version of an OpenSearch application.
//
// @param request - ModifyAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppGroupResponse
func (client *Client) ModifyAppGroupWithContext(ctx context.Context, appGroupIdentity *string, request *ModifyAppGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentVersion) {
		body["currentVersion"] = request.CurrentVersion
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Domain) {
		body["domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the quotas of an OpenSearch application.
//
// @param request - ModifyAppGroupQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppGroupQuotaResponse
func (client *Client) ModifyAppGroupQuotaWithContext(ctx context.Context, appGroupIdentity *string, request *ModifyAppGroupQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyAppGroupQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppGroupQuota"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/quota"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppGroupQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a rough sort expression for an OpenSearch application. If you set dryRun to true, this operation checks the rough sort expression after the expression is modified. If you do not specify this parameter, false is used by default.
//
// @param request - ModifyFirstRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirstRankResponse
func (client *Client) ModifyFirstRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, request *ModifyFirstRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyFirstRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFirstRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/first-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFirstRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a query analysis rule for a specific application version. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifyQueryProcessorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQueryProcessorResponse
func (client *Client) ModifyQueryProcessorWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, request *ModifyQueryProcessorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyQueryProcessorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyQueryProcessor"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQueryProcessorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scheduled task.
//
// @param request - ModifyScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledTaskResponse
func (client *Client) ModifyScheduledTaskWithContext(ctx context.Context, appGroupIdentity *string, taskId *string, request *ModifyScheduledTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
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
		Action:      dara.String("ModifyScheduledTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scheduled-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a fine sort expression that is configured for a specific OpenSearch application version. If you set dryRun to true, the specified fine sort expression is checked after the expression is modified. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifySecondRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecondRankResponse
func (client *Client) ModifySecondRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, request *ModifySecondRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifySecondRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecondRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/second-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecondRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts the changes in intervention entries.
//
// @param request - PushInterventionDictionaryEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushInterventionDictionaryEntriesResponse
func (client *Client) PushInterventionDictionaryEntriesWithContext(ctx context.Context, name *string, request *PushInterventionDictionaryEntriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushInterventionDictionaryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushInterventionDictionaryEntries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name)) + "/entries/actions/bulk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts the changes in the entries of a custom analyzer.
//
// @param request - PushUserAnalyzerEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushUserAnalyzerEntriesResponse
func (client *Client) PushUserAnalyzerEntriesWithContext(ctx context.Context, name *string, request *PushUserAnalyzerEntriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushUserAnalyzerEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Entries) {
		body["entries"] = request.Entries
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushUserAnalyzerEntries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers/" + dara.PercentEncode(dara.StringValue(name)) + "/entries/actions/bulk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布排序脚本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseSortScriptResponse
func (client *Client) ReleaseSortScriptWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseSortScriptResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/actions/release"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppResponse
func (client *Client) RemoveAppWithContext(ctx context.Context, appGroupIdentity *string, appId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveAppResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApp"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an OpenSearch application.
//
// Description:
//
// You can delete only pay-as-you-go applications. You cannot delete subscription applications.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppGroupResponse
func (client *Client) RemoveAppGroupWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveAppGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAppGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables data collection.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataCollectionResponse
func (client *Client) RemoveDataCollectionWithContext(ctx context.Context, appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveDataCollectionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDataCollection"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/data-collections/" + dara.PercentEncode(dara.StringValue(dataCollectionIdentity))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rough sort expression for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFirstRankResponse
func (client *Client) RemoveFirstRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveFirstRankResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFirstRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/first-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFirstRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an intervention dictionary.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInterventionDictionaryResponse
func (client *Client) RemoveInterventionDictionaryWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveInterventionDictionaryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInterventionDictionary"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/intervention-dictionaries/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInterventionDictionaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a query analysis rule for an OpenSearch application version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveQueryProcessorResponse
func (client *Client) RemoveQueryProcessorWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveQueryProcessorResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveQueryProcessor"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/query-processors/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveQueryProcessorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled task of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveScheduledTaskResponse
func (client *Client) RemoveScheduledTaskWithContext(ctx context.Context, appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveScheduledTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveScheduledTask"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scheduled-tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveScheduledTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a query policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSearchStrategyResponse
func (client *Client) RemoveSearchStrategyWithContext(ctx context.Context, appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveSearchStrategyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSearchStrategy"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/search-strategies/" + dara.PercentEncode(dara.StringValue(strategyName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSearchStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a fine sort expression from a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSecondRankResponse
func (client *Client) RemoveSecondRankWithContext(ctx context.Context, appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveSecondRankResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSecondRank"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/second-ranks/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSecondRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom analyzer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserAnalyzerResponse
func (client *Client) RemoveUserAnalyzerWithContext(ctx context.Context, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveUserAnalyzerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/user-analyzers/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews an application. This operation is not available now. You must renew an application in the OpenSearch console.
//
// @param request - RenewAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppGroupResponse
func (client *Client) RenewAppGroupWithContext(ctx context.Context, appGroupIdentity *string, request *RenewAppGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAppGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/actions/renew"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts a service-based application to an instance-based application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceAppGroupCommodityCodeResponse
func (client *Client) ReplaceAppGroupCommodityCodeWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReplaceAppGroupCommodityCodeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceAppGroupCommodityCode"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/actions/to-instance-typed"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceAppGroupCommodityCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a sort script.
//
// @param request - SaveSortScriptFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSortScriptFileResponse
func (client *Client) SaveSortScriptFileWithContext(ctx context.Context, appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, request *SaveSortScriptFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveSortScriptFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSortScriptFile"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName)) + "/files/src/" + dara.PercentEncode(dara.StringValue(fileName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSortScriptFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 立即进行慢查询分析
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSlowQueryAnalyzerResponse
func (client *Client) StartSlowQueryAnalyzerWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartSlowQueryAnalyzerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSlowQueryAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/optimizers/slow-query/actions/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSlowQueryAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/resource-tags"),
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
// Unbinds a custom analyzer from an Elasticsearch instance.
//
// Description:
//
// You can call this operation to unbind a custom analyzer from an Elasticsearch instance.
//
// @param request - UnbindESUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindESUserAnalyzerResponse
func (client *Client) UnbindESUserAnalyzerWithContext(ctx context.Context, appGroupIdentity *string, esInstanceId *string, request *UnbindESUserAnalyzerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindESUserAnalyzerResponse, _err error) {
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
		Action:      dara.String("UnbindESUserAnalyzer"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/es/" + dara.PercentEncode(dara.StringValue(esInstanceId)) + "/actions/unbind-analyzer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindESUserAnalyzerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an Elasticsearch instance from an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEsInstanceResponse
func (client *Client) UnbindEsInstanceWithContext(ctx context.Context, appGroupIdentity *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindEsInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindEsInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/actions/unbind-es-instance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindEsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove tags from resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/resource-tags"),
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
// Modifies the parameters of an A/B test.
//
// @param request - UpdateABTestExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestExperimentResponse
func (client *Client) UpdateABTestExperimentWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABTestExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABTestExperiment"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments/" + dara.PercentEncode(dara.StringValue(experimentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABTestExperimentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies whitelists.
//
// @param request - UpdateABTestFixedFlowDividersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestFixedFlowDividersResponse
func (client *Client) UpdateABTestFixedFlowDividersWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestFixedFlowDividersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABTestFixedFlowDividersResponse, _err error) {
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
		Action:      dara.String("UpdateABTestFixedFlowDividers"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/experiments/" + dara.PercentEncode(dara.StringValue(experimentId)) + "/fixed-flow-dividers"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a test group.
//
// @param request - UpdateABTestGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestGroupResponse
func (client *Client) UpdateABTestGroupWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, groupId *string, request *UpdateABTestGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABTestGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABTestGroup"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABTestGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an A/B test scenario.
//
// @param request - UpdateABTestSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestSceneResponse
func (client *Client) UpdateABTestSceneWithContext(ctx context.Context, appGroupIdentity *string, sceneId *string, request *UpdateABTestSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateABTestSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateABTestScene"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/scenes/" + dara.PercentEncode(dara.StringValue(sceneId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateABTestSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates fetch fields. A dry run is supported.
//
// @param request - UpdateFetchFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFetchFieldsResponse
func (client *Client) UpdateFetchFieldsWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *UpdateFetchFieldsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFetchFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFetchFields"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/fetch-fields"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFetchFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the default algorithm instance used by the specified application. The new algorithm instance automatically overwrites the most recently set default instance. If no instance is set, the default instance is canceled.
//
// @param request - UpdateFunctionDefaultInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionDefaultInstanceResponse
func (client *Client) UpdateFunctionDefaultInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, request *UpdateFunctionDefaultInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFunctionDefaultInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunctionDefaultInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/default-instance"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an algorithm instance.
//
// @param request - UpdateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionInstanceResponse
func (client *Client) UpdateFunctionInstanceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateParameters) {
		body["createParameters"] = request.CreateParameters
	}

	if !dara.IsNil(request.Cron) {
		body["cron"] = request.Cron
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.UsageParameters) {
		body["usageParameters"] = request.UsageParameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunctionInstance"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an algorithm resource.
//
// Description:
//
// You can call this operation to update the information about resources by resource name. You can modify only the values of data and description.
//
// @param request - UpdateFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResourceResponse
func (client *Client) UpdateFunctionResourceWithContext(ctx context.Context, appGroupIdentity *string, functionName *string, resourceName *string, request *UpdateFunctionResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFunctionResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunctionResource"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/resources/" + dara.PercentEncode(dara.StringValue(resourceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a query policy.
//
// @param request - UpdateSearchStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSearchStrategyResponse
func (client *Client) UpdateSearchStrategyWithContext(ctx context.Context, appGroupIdentity *string, appId *string, strategyName *string, request *UpdateSearchStrategyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSearchStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSearchStrategy"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/search-strategies/" + dara.PercentEncode(dara.StringValue(strategyName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSearchStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a sort script.
//
// Description:
//
// You can call this operation to modify the description of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSortScriptResponse
func (client *Client) UpdateSortScriptWithContext(ctx context.Context, appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSortScriptResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSortScript"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appVersionId)) + "/sort-scripts/" + dara.PercentEncode(dara.StringValue(scriptName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSortScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates summaries. A dry run is supported.
//
// @param request - UpdateSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSummariesResponse
func (client *Client) UpdateSummariesWithContext(ctx context.Context, appGroupIdentity *string, appId *string, request *UpdateSummariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSummariesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ToArray(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSummaries"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/app-groups/" + dara.PercentEncode(dara.StringValue(appGroupIdentity)) + "/apps/" + dara.PercentEncode(dara.StringValue(appId)) + "/summaries"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSummariesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies data sources.
//
// @param request - ValidateDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateDataSourcesResponse
func (client *Client) ValidateDataSourcesWithContext(ctx context.Context, request *ValidateDataSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateDataSources"),
		Version:     dara.String("2017-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v4/openapi/assist/data-sources/validations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateDataSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
