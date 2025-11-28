// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Cancels an artifact building task.
//
// @param request - CancelArtifactBuildTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelArtifactBuildTaskResponse
func (client *Client) CancelArtifactBuildTaskWithContext(ctx context.Context, request *CancelArtifactBuildTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelArtifactBuildTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildTaskId) {
		query["BuildTaskId"] = request.BuildTaskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelArtifactBuildTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelArtifactBuildTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an image building task of a repository.
//
// @param request - CancelRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRepoBuildRecordResponse
func (client *Client) CancelRepoBuildRecordWithContext(ctx context.Context, request *CancelRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *CancelRepoBuildRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRecordId) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRepoBuildRecord"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRepoBuildRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a single replication task.
//
// @param request - CancelRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRepoSyncTaskResponse
func (client *Client) CancelRepoSyncTaskWithContext(ctx context.Context, request *CancelRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelRepoSyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SyncTaskId) {
		query["SyncTaskId"] = request.SyncTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRepoSyncTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRepoSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a resource belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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

// Summary:
//
// Create image repository acceleration rules for image building.
//
// Description:
//
// You can create building rules of accelerated images only for image repositories of Container Registry Advanced Edition instances. You cannot create building rules of accelerated images for image repositories of Container Registry Basic Edition instances. For more information, see [Specifications of different editions](https://www.alibabacloud.com/help/en/acr/product-overview/what-is-container-registry?spm=openapi-amp.newDocPublishment.0.0.bf82281fRj7rmV#section-n3q-ps7-x6k).
//
// Accelerated images are not supported in Alibaba Finance Cloud regions or Alibaba Gov Cloud regions.
//
// @param tmpReq - CreateArtifactBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactBuildRuleResponse
func (client *Client) CreateArtifactBuildRuleWithContext(ctx context.Context, tmpReq *CreateArtifactBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactBuildRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateArtifactBuildRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ScopeId) {
		query["ScopeId"] = request.ScopeId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifactBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lifecycle management rule for an artifact.
//
// @param request - CreateArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactLifecycleRuleResponse
func (client *Client) CreateArtifactLifecycleRuleWithContext(ctx context.Context, request *CreateArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactLifecycleRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auto) {
		query["Auto"] = request.Auto
	}

	if !dara.IsNil(request.EnableDeleteTag) {
		query["EnableDeleteTag"] = request.EnableDeleteTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RetentionTagCount) {
		query["RetentionTagCount"] = request.RetentionTagCount
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.TagRegexp) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifactLifecycleRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription rule.
//
// @param request - CreateArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactSubscriptionRuleResponse
func (client *Client) CreateArtifactSubscriptionRuleWithContext(ctx context.Context, request *CreateArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactSubscriptionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accelerate) {
		query["Accelerate"] = request.Accelerate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.Override) {
		query["Override"] = request.Override
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.SourceNamespaceName) {
		query["SourceNamespaceName"] = request.SourceNamespaceName
	}

	if !dara.IsNil(request.SourceProvider) {
		query["SourceProvider"] = request.SourceProvider
	}

	if !dara.IsNil(request.SourceRepoName) {
		query["SourceRepoName"] = request.SourceRepoName
	}

	if !dara.IsNil(request.TagCount) {
		query["TagCount"] = request.TagCount
	}

	if !dara.IsNil(request.TagRegexp) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifactSubscriptionRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription task.
//
// @param request - CreateArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactSubscriptionTaskResponse
func (client *Client) CreateArtifactSubscriptionTaskWithContext(ctx context.Context, request *CreateArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactSubscriptionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifactSubscriptionTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building record based on an existing record.
//
// @param request - CreateBuildRecordByRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBuildRecordByRecordResponse
func (client *Client) CreateBuildRecordByRecordWithContext(ctx context.Context, request *CreateBuildRecordByRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateBuildRecordByRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRecordId) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBuildRecordByRecord"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBuildRecordByRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building record based on a rule.
//
// @param request - CreateBuildRecordByRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBuildRecordByRuleResponse
func (client *Client) CreateBuildRecordByRuleWithContext(ctx context.Context, request *CreateBuildRecordByRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateBuildRecordByRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRuleId) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBuildRecordByRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBuildRecordByRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery chain.
//
// @param request - CreateChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChainResponse
func (client *Client) CreateChainWithContext(ctx context.Context, request *CreateChainRequest, runtime *dara.RuntimeOptions) (_result *CreateChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChainConfig) {
		query["ChainConfig"] = request.ChainConfig
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.ScopeExclude) {
		query["ScopeExclude"] = request.ScopeExclude
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChain"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a chart namespace in an instance.
//
// @param request - CreateChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChartNamespaceResponse
func (client *Client) CreateChartNamespaceWithContext(ctx context.Context, request *CreateChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateChartNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateRepo) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !dara.IsNil(request.DefaultRepoType) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChartNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChartNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a chart repository.
//
// @param request - CreateChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChartRepositoryResponse
func (client *Client) CreateChartRepositoryWithContext(ctx context.Context, request *CreateChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *CreateChartRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoType) {
		query["RepoType"] = request.RepoType
	}

	if !dara.IsNil(request.Summary) {
		query["Summary"] = request.Summary
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChartRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChartRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a whitelist policy for the public endpoint of the instance.
//
// @param request - CreateInstanceEndpointAclPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceEndpointAclPolicyResponse
func (client *Client) CreateInstanceEndpointAclPolicyWithContext(ctx context.Context, request *CreateInstanceEndpointAclPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceEndpointAclPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.Entry) {
		query["Entry"] = request.Entry
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceEndpointAclPolicy"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) with a Container Registry instance.
//
// Description:
//
// The VPC quota must be purchased separately.
//
// @param request - CreateInstanceVpcEndpointLinkedVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceVpcEndpointLinkedVpcResponse
func (client *Client) CreateInstanceVpcEndpointLinkedVpcWithContext(ctx context.Context, request *CreateInstanceVpcEndpointLinkedVpcRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceVpcEndpointLinkedVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableCreateDNSRecordInPvzt) {
		query["EnableCreateDNSRecordInPvzt"] = request.EnableCreateDNSRecordInPvzt
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceVpcEndpointLinkedVpc"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace of image repositories.
//
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithContext(ctx context.Context, tmpReq *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DefaultRepoConfiguration) {
		request.DefaultRepoConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRepoConfiguration, dara.String("DefaultRepoConfiguration"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateRepo) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !dara.IsNil(request.DefaultRepoConfigurationShrink) {
		query["DefaultRepoConfiguration"] = request.DefaultRepoConfigurationShrink
	}

	if !dara.IsNil(request.DefaultRepoType) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building rule for a repository.
//
// @param request - CreateRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoBuildRuleResponse
func (client *Client) CreateRepoBuildRuleWithContext(ctx context.Context, request *CreateRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoBuildRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildArgs) {
		query["BuildArgs"] = request.BuildArgs
	}

	if !dara.IsNil(request.DockerfileLocation) {
		query["DockerfileLocation"] = request.DockerfileLocation
	}

	if !dara.IsNil(request.DockerfileName) {
		query["DockerfileName"] = request.DockerfileName
	}

	if !dara.IsNil(request.ImageTag) {
		query["ImageTag"] = request.ImageTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Platforms) {
		query["Platforms"] = request.Platforms
	}

	if !dara.IsNil(request.PushName) {
		query["PushName"] = request.PushName
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a source code repository to an image repository.
//
// @param request - CreateRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSourceCodeRepoResponse
func (client *Client) CreateRepoSourceCodeRepoWithContext(ctx context.Context, request *CreateRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSourceCodeRepoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.CodeRepoName) {
		query["CodeRepoName"] = request.CodeRepoName
	}

	if !dara.IsNil(request.CodeRepoNamespaceName) {
		query["CodeRepoNamespaceName"] = request.CodeRepoNamespaceName
	}

	if !dara.IsNil(request.CodeRepoType) {
		query["CodeRepoType"] = request.CodeRepoType
	}

	if !dara.IsNil(request.DisableCacheBuild) {
		query["DisableCacheBuild"] = request.DisableCacheBuild
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OverseaBuild) {
		query["OverseaBuild"] = request.OverseaBuild
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoSourceCodeRepo"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image synchronization rule for an image repository.
//
// @param request - CreateRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncRuleResponse
func (client *Client) CreateRepoSyncRuleWithContext(ctx context.Context, request *CreateRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNameFilter) {
		query["RepoNameFilter"] = request.RepoNameFilter
	}

	if !dara.IsNil(request.SyncRuleName) {
		query["SyncRuleName"] = request.SyncRuleName
	}

	if !dara.IsNil(request.SyncScope) {
		query["SyncScope"] = request.SyncScope
	}

	if !dara.IsNil(request.SyncTrigger) {
		query["SyncTrigger"] = request.SyncTrigger
	}

	if !dara.IsNil(request.TagFilter) {
		query["TagFilter"] = request.TagFilter
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !dara.IsNil(request.TargetNamespaceName) {
		query["TargetNamespaceName"] = request.TargetNamespaceName
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	if !dara.IsNil(request.TargetRepoName) {
		query["TargetRepoName"] = request.TargetRepoName
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoSyncRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoSyncRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncTaskResponse
func (client *Client) CreateRepoSyncTaskWithContext(ctx context.Context, request *CreateRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Override) {
		query["Override"] = request.Override
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !dara.IsNil(request.TargetNamespace) {
		query["TargetNamespace"] = request.TargetNamespace
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	if !dara.IsNil(request.TargetRepoName) {
		query["TargetRepoName"] = request.TargetRepoName
	}

	if !dara.IsNil(request.TargetTag) {
		query["TargetTag"] = request.TargetTag
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoSyncTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image replication task based on a manual replication rule.
//
// @param request - CreateRepoSyncTaskByRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncTaskByRuleResponse
func (client *Client) CreateRepoSyncTaskByRuleWithContext(ctx context.Context, request *CreateRepoSyncTaskByRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncTaskByRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.SyncRuleId) {
		query["SyncRuleId"] = request.SyncRuleId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoSyncTaskByRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoSyncTaskByRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image tag based on an existing image tag in an image repository.
//
// @param request - CreateRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTagResponse
func (client *Client) CreateRepoTagWithContext(ctx context.Context, request *CreateRepoTagRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromTag) {
		query["FromTag"] = request.FromTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.ToTag) {
		query["ToTag"] = request.ToTag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoTag"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image scan task.
//
// @param request - CreateRepoTagScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTagScanTaskResponse
func (client *Client) CreateRepoTagScanTaskWithContext(ctx context.Context, request *CreateRepoTagScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTagScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.ScanService) {
		query["ScanService"] = request.ScanService
	}

	if !dara.IsNil(request.ScanType) {
		query["ScanType"] = request.ScanType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoTagScanTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoTagScanTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trigger for a repository.
//
// @param request - CreateRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTriggerResponse
func (client *Client) CreateRepoTriggerWithContext(ctx context.Context, request *CreateRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.TriggerName) {
		query["TriggerName"] = request.TriggerName
	}

	if !dara.IsNil(request.TriggerTag) {
		query["TriggerTag"] = request.TriggerTag
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.TriggerUrl) {
		query["TriggerUrl"] = request.TriggerUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoTrigger"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image repository.
//
// @param request - CreateRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest, runtime *dara.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Detail) {
		query["Detail"] = request.Detail
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoType) {
		query["RepoType"] = request.RepoType
	}

	if !dara.IsNil(request.Summary) {
		query["Summary"] = request.Summary
	}

	if !dara.IsNil(request.TagImmutability) {
		query["TagImmutability"] = request.TagImmutability
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建扫描规则
//
// @param tmpReq - CreateScanRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanRuleResponse
func (client *Client) CreateScanRuleWithContext(ctx context.Context, tmpReq *CreateScanRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateScanRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScanRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Namespaces) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, dara.String("Namespaces"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepoNames) {
		request.RepoNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepoNames, dara.String("RepoNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespacesShrink) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !dara.IsNil(request.RepoNamesShrink) {
		query["RepoNames"] = request.RepoNamesShrink
	}

	if !dara.IsNil(request.RepoTagFilterPattern) {
		query["RepoTagFilterPattern"] = request.RepoTagFilterPattern
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.ScanScope) {
		query["ScanScope"] = request.ScanScope
	}

	if !dara.IsNil(request.ScanType) {
		query["ScanType"] = request.ScanType
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScanRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScanRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例域名路由规则
//
// @param tmpReq - CreateStorageDomainRoutingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStorageDomainRoutingRuleResponse
func (client *Client) CreateStorageDomainRoutingRuleWithContext(ctx context.Context, tmpReq *CreateStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateStorageDomainRoutingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStorageDomainRoutingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Routes) {
		request.RoutesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Routes, dara.String("Routes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RoutesShrink) {
		query["Routes"] = request.RoutesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStorageDomainRoutingRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStorageDomainRoutingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an artifact lifecycle management rule.
//
// @param request - DeleteArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactLifecycleRuleResponse
func (client *Client) DeleteArtifactLifecycleRuleWithContext(ctx context.Context, request *DeleteArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactLifecycleRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteArtifactLifecycleRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an artifact subscription rule.
//
// @param request - DeleteArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactSubscriptionRuleResponse
func (client *Client) DeleteArtifactSubscriptionRuleWithContext(ctx context.Context, request *DeleteArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactSubscriptionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteArtifactSubscriptionRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a delivery pipeline.
//
// @param request - DeleteChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChainResponse
func (client *Client) DeleteChainWithContext(ctx context.Context, request *DeleteChainRequest, runtime *dara.RuntimeOptions) (_result *DeleteChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChainId) {
		query["ChainId"] = request.ChainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChain"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart namespace from an instance.
//
// Description:
//
// >  If you delete a chart namespace, all repositories in the namespace and the charts in all repositories are deleted.
//
// @param request - DeleteChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartNamespaceResponse
func (client *Client) DeleteChartNamespaceWithContext(ctx context.Context, request *DeleteChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChartNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChartNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart version from a chart repository.
//
// @param request - DeleteChartReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartReleaseResponse
func (client *Client) DeleteChartReleaseWithContext(ctx context.Context, request *DeleteChartReleaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Chart) {
		query["Chart"] = request.Chart
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChartRelease"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChartReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart repository from an instance.
//
// @param request - DeleteChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartRepositoryResponse
func (client *Client) DeleteChartRepositoryWithContext(ctx context.Context, request *DeleteChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChartRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChartRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event notification rule.
//
// @param request - DeleteEventCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventCenterRuleResponse
func (client *Client) DeleteEventCenterRuleWithContext(ctx context.Context, request *DeleteEventCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventCenterRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventCenterRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a whitelist policy for the public endpoint of an instance.
//
// @param request - DeleteInstanceEndpointAclPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceEndpointAclPolicyResponse
func (client *Client) DeleteInstanceEndpointAclPolicyWithContext(ctx context.Context, request *DeleteInstanceEndpointAclPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceEndpointAclPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.Entry) {
		query["Entry"] = request.Entry
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceEndpointAclPolicy"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a virtual private cloud (VPC) from an instance.
//
// @param request - DeleteInstanceVpcEndpointLinkedVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceVpcEndpointLinkedVpcResponse
func (client *Client) DeleteInstanceVpcEndpointLinkedVpcWithContext(ctx context.Context, request *DeleteInstanceVpcEndpointLinkedVpcRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceVpcEndpointLinkedVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceVpcEndpointLinkedVpc"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > After you delete a namespace, all repositories in the namespace and all images in these repositories are deleted as well.
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image building rule of a repository.
//
// @param request - DeleteRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoBuildRuleResponse
func (client *Client) DeleteRepoBuildRuleWithContext(ctx context.Context, request *DeleteRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoBuildRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRuleId) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepoBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepoBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image replication rule of an image repository.
//
// @param request - DeleteRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoSyncRuleResponse
func (client *Client) DeleteRepoSyncRuleWithContext(ctx context.Context, request *DeleteRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoSyncRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SyncRuleId) {
		query["SyncRuleId"] = request.SyncRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepoSyncRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepoSyncRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image tag.
//
// @param request - DeleteRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoTagResponse
func (client *Client) DeleteRepoTagWithContext(ctx context.Context, request *DeleteRepoTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepoTag"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepoTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trigger of an image repository.
//
// @param request - DeleteRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoTriggerResponse
func (client *Client) DeleteRepoTriggerWithContext(ctx context.Context, request *DeleteRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.TriggerId) {
		query["TriggerId"] = request.TriggerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepoTrigger"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepoTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image repository.
//
// Description:
//
// If you delete a repository, all images in the repository are also deleted.
//
// @param request - DeleteRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除扫描规则
//
// @param request - DeleteScanRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScanRuleResponse
func (client *Client) DeleteScanRuleWithContext(ctx context.Context, request *DeleteScanRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteScanRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScanRuleId) {
		query["ScanRuleId"] = request.ScanRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScanRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScanRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实例存储域名映射规则
//
// @param request - DeleteStorageDomainRoutingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStorageDomainRoutingRuleResponse
func (client *Client) DeleteStorageDomainRoutingRuleWithContext(ctx context.Context, request *DeleteStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteStorageDomainRoutingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStorageDomainRoutingRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStorageDomainRoutingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArtifactBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactBuildRuleResponse
func (client *Client) GetArtifactBuildRuleWithContext(ctx context.Context, request *GetArtifactBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactBuildRuleResponse, _err error) {
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
		Action:      dara.String("GetArtifactBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact building task.
//
// @param request - GetArtifactBuildTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactBuildTaskResponse
func (client *Client) GetArtifactBuildTaskWithContext(ctx context.Context, request *GetArtifactBuildTaskRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactBuildTaskResponse, _err error) {
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
		Action:      dara.String("GetArtifactBuildTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactBuildTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - GetArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactLifecycleRuleResponse
func (client *Client) GetArtifactLifecycleRuleWithContext(ctx context.Context, request *GetArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactLifecycleRuleResponse, _err error) {
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
		Action:      dara.String("GetArtifactLifecycleRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an artifact subscription rule.
//
// @param request - GetArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionRuleResponse
func (client *Client) GetArtifactSubscriptionRuleWithContext(ctx context.Context, request *GetArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionRuleResponse, _err error) {
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
		Action:      dara.String("GetArtifactSubscriptionRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionTaskResponse
func (client *Client) GetArtifactSubscriptionTaskWithContext(ctx context.Context, request *GetArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResponse, _err error) {
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
		Action:      dara.String("GetArtifactSubscriptionTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionTaskResultResponse
func (client *Client) GetArtifactSubscriptionTaskResultWithContext(ctx context.Context, request *GetArtifactSubscriptionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResultResponse, _err error) {
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
		Action:      dara.String("GetArtifactSubscriptionTaskResult"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactSubscriptionTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a temporary username and a token that you can use to log on to a Container Registry instance.
//
// Description:
//
// The ID of the Container Registry instance.
//
// @param request - GetAuthorizationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationTokenResponse
func (client *Client) GetAuthorizationTokenWithContext(ctx context.Context, request *GetAuthorizationTokenRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorizationToken"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorizationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取交付链
//
// @param request - GetChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChainResponse
func (client *Client) GetChainWithContext(ctx context.Context, request *GetChainRequest, runtime *dara.RuntimeOptions) (_result *GetChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChainId) {
		query["ChainId"] = request.ChainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChain"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a chart namespace in an instance.
//
// @param request - GetChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChartNamespaceResponse
func (client *Client) GetChartNamespaceWithContext(ctx context.Context, request *GetChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetChartNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChartNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChartNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a chart repository.
//
// @param request - GetChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChartRepositoryResponse
func (client *Client) GetChartRepositoryWithContext(ctx context.Context, request *GetChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *GetChartRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChartRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChartRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the resource group to which the instance belongs.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the endpoint of an instance.
//
// @param request - GetInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceEndpointResponse
func (client *Client) GetInstanceEndpointWithContext(ctx context.Context, request *GetInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceEndpoint"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota usage of an instance.
//
// @param request - GetInstanceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUsageResponse
func (client *Client) GetInstanceUsageWithContext(ctx context.Context, request *GetInstanceUsageRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceUsage"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of the virtual private clouds (VPCs) in which a Container Registry instance is deployed.
//
// @param request - GetInstanceVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceVpcEndpointResponse
func (client *Client) GetInstanceVpcEndpointWithContext(ctx context.Context, request *GetInstanceVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceVpcEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceVpcEndpoint"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - GetNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNamespaceResponse
func (client *Client) GetNamespaceWithContext(ctx context.Context, request *GetNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about image building records of a repository.
//
// Description:
//
// ***
//
// @param request - GetRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoBuildRecordResponse
func (client *Client) GetRepoBuildRecordWithContext(ctx context.Context, request *GetRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *GetRepoBuildRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRecordId) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoBuildRecord"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoBuildRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an image building task.
//
// @param request - GetRepoBuildRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoBuildRecordStatusResponse
func (client *Client) GetRepoBuildRecordStatusWithContext(ctx context.Context, request *GetRepoBuildRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRepoBuildRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRecordId) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoBuildRecordStatus"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoBuildRecordStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the source code repository that is bound to an image repository.
//
// @param request - GetRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoSourceCodeRepoResponse
func (client *Client) GetRepoSourceCodeRepoWithContext(ctx context.Context, request *GetRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *GetRepoSourceCodeRepoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoSourceCodeRepo"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an image synchronization task in an instance.
//
// @param request - GetRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoSyncTaskResponse
func (client *Client) GetRepoSyncTaskWithContext(ctx context.Context, request *GetRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetRepoSyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SyncTaskId) {
		query["SyncTaskId"] = request.SyncTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoSyncTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The version of the repository.
//
// @param request - GetRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagResponse
func (client *Client) GetRepoTagWithContext(ctx context.Context, request *GetRepoTagRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagResponse, _err error) {
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
		Action:      dara.String("GetRepoTag"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scanning status of an image tag.
//
// @param request - GetRepoTagScanStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagScanStatusResponse
func (client *Client) GetRepoTagScanStatusWithContext(ctx context.Context, request *GetRepoTagScanStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagScanStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.ScanTaskId) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !dara.IsNil(request.ScanType) {
		query["ScanType"] = request.ScanType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoTagScanStatus"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoTagScanStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRepoTagScanSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagScanSummaryResponse
func (client *Client) GetRepoTagScanSummaryWithContext(ctx context.Context, request *GetRepoTagScanSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagScanSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.ScanTaskId) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepoTagScanSummary"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepoTagScanSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about an image repository.
//
// @param request - GetRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryResponse
func (client *Client) GetRepositoryWithContext(ctx context.Context, request *GetRepositoryRequest, runtime *dara.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询扫描规则
//
// @param request - GetScanRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScanRuleResponse
func (client *Client) GetScanRuleWithContext(ctx context.Context, request *GetScanRuleRequest, runtime *dara.RuntimeOptions) (_result *GetScanRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScanRuleId) {
		query["ScanRuleId"] = request.ScanRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScanRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScanRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例存储域名路由规则
//
// @param request - GetStorageDomainRoutingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageDomainRoutingRuleResponse
func (client *Client) GetStorageDomainRoutingRuleWithContext(ctx context.Context, request *GetStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetStorageDomainRoutingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageDomainRoutingRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageDomainRoutingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log entries of an artifact building task.
//
// @param request - ListArtifactBuildTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactBuildTaskLogResponse
func (client *Client) ListArtifactBuildTaskLogWithContext(ctx context.Context, request *ListArtifactBuildTaskLogRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactBuildTaskLogResponse, _err error) {
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
		Action:      dara.String("ListArtifactBuildTaskLog"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactBuildTaskLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - ListArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactLifecycleRuleResponse
func (client *Client) ListArtifactLifecycleRuleWithContext(ctx context.Context, request *ListArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactLifecycleRuleResponse, _err error) {
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
		Action:      dara.String("ListArtifactLifecycleRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the subscription rules of artifacts.
//
// @param request - ListArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactSubscriptionRuleResponse
func (client *Client) ListArtifactSubscriptionRuleWithContext(ctx context.Context, request *ListArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactSubscriptionRuleResponse, _err error) {
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
		Action:      dara.String("ListArtifactSubscriptionRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists artifact subscription tasks.
//
// @param request - ListArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactSubscriptionTaskResponse
func (client *Client) ListArtifactSubscriptionTaskWithContext(ctx context.Context, request *ListArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactSubscriptionTaskResponse, _err error) {
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
		Action:      dara.String("ListArtifactSubscriptionTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries delivery chains.
//
// @param request - ListChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChainResponse
func (client *Client) ListChainWithContext(ctx context.Context, request *ListChainRequest, runtime *dara.RuntimeOptions) (_result *ListChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChain"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The response code.
//
// @param request - ListChainInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChainInstanceResponse
func (client *Client) ListChainInstanceWithContext(ctx context.Context, request *ListChainInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListChainInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChainInstance"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChainInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart namespaces of a Container Registry instance.
//
// @param request - ListChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartNamespaceResponse
func (client *Client) ListChartNamespaceWithContext(ctx context.Context, request *ListChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ListChartNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NamespaceStatus) {
		query["NamespaceStatus"] = request.NamespaceStatus
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChartNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChartNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a chart in a chart repository.
//
// @param request - ListChartReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartReleaseResponse
func (client *Client) ListChartReleaseWithContext(ctx context.Context, request *ListChartReleaseRequest, runtime *dara.RuntimeOptions) (_result *ListChartReleaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Chart) {
		query["Chart"] = request.Chart
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChartRelease"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChartReleaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart repositories of a Container Registry instance.
//
// @param request - ListChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartRepositoryResponse
func (client *Client) ListChartRepositoryWithContext(ctx context.Context, request *ListChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *ListChartRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoStatus) {
		query["RepoStatus"] = request.RepoStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChartRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChartRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical events of an event rule.
//
// @param request - ListEventCenterRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventCenterRecordResponse
func (client *Client) ListEventCenterRecordWithContext(ctx context.Context, request *ListEventCenterRecordRequest, runtime *dara.RuntimeOptions) (_result *ListEventCenterRecordResponse, _err error) {
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
		Action:      dara.String("ListEventCenterRecord"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventCenterRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of event notification rules.
//
// @param request - ListEventCenterRuleNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventCenterRuleNameResponse
func (client *Client) ListEventCenterRuleNameWithContext(ctx context.Context, request *ListEventCenterRuleNameRequest, runtime *dara.RuntimeOptions) (_result *ListEventCenterRuleNameResponse, _err error) {
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
		Action:      dara.String("ListEventCenterRuleName"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventCenterRuleNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Container Registry instances.
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstance"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of a Container Registry instance.
//
// @param request - ListInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceEndpointResponse
func (client *Client) ListInstanceEndpointWithContext(ctx context.Context, request *ListInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	if !dara.IsNil(request.Summary) {
		query["Summary"] = request.Summary
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceEndpoint"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions in which you can create Container Registry instances.
//
// @param request - ListInstanceRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRegionResponse
func (client *Client) ListInstanceRegionWithContext(ctx context.Context, request *ListInstanceRegionRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceRegion"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries namespaces in a Container Registry instance.
//
// @param request - ListNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespaceResponse
func (client *Client) ListNamespaceWithContext(ctx context.Context, request *ListNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ListNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.NamespaceStatus) {
		query["NamespaceStatus"] = request.NamespaceStatus
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image building records of an image repository.
//
// @param request - ListRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRecordResponse
func (client *Client) ListRepoBuildRecordWithContext(ctx context.Context, request *ListRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoBuildRecord"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoBuildRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of an image building record.
//
// @param request - ListRepoBuildRecordLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRecordLogResponse
func (client *Client) ListRepoBuildRecordLogWithContext(ctx context.Context, request *ListRepoBuildRecordLogRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRecordLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildRecordId) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoBuildRecordLog"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoBuildRecordLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image building rules of a repository.
//
// @param request - ListRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRuleResponse
func (client *Client) ListRepoBuildRuleWithContext(ctx context.Context, request *ListRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image synchronization rules of a repository.
//
// @param request - ListRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoSyncRuleResponse
func (client *Client) ListRepoSyncRuleWithContext(ctx context.Context, request *ListRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *ListRepoSyncRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoSyncRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoSyncRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image synchronization tasks in an image repository.
//
// @param request - ListRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoSyncTaskResponse
func (client *Client) ListRepoSyncTaskWithContext(ctx context.Context, request *ListRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *ListRepoSyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.SyncRecordId) {
		query["SyncRecordId"] = request.SyncRecordId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoSyncTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image tags in a repository.
//
// @param request - ListRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTagResponse
func (client *Client) ListRepoTagWithContext(ctx context.Context, request *ListRepoTagRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoTag"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a security scan that is created for an image tag.
//
// @param request - ListRepoTagScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTagScanResultResponse
func (client *Client) ListRepoTagScanResultWithContext(ctx context.Context, request *ListRepoTagScanResultRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTagScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.FilterValue) {
		query["FilterValue"] = request.FilterValue
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.ScanTaskId) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !dara.IsNil(request.ScanType) {
		query["ScanType"] = request.ScanType
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VulQueryKey) {
		query["VulQueryKey"] = request.VulQueryKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoTagScanResult"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoTagScanResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the triggers of a repository.
//
// @param request - ListRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTriggerResponse
func (client *Client) ListRepoTriggerWithContext(ctx context.Context, request *ListRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepoTrigger"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepoTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image repositories.
//
// @param request - ListRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryResponse
func (client *Client) ListRepositoryWithContext(ctx context.Context, request *ListRepositoryRequest, runtime *dara.RuntimeOptions) (_result *ListRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoStatus) {
		query["RepoStatus"] = request.RepoStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the baseline risks of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanBaselineByTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanBaselineByTaskResponse
func (client *Client) ListScanBaselineByTaskWithContext(ctx context.Context, request *ListScanBaselineByTaskRequest, runtime *dara.RuntimeOptions) (_result *ListScanBaselineByTaskResponse, _err error) {
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
		Action:      dara.String("ListScanBaselineByTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScanBaselineByTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the malicious files of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanMaliciousFileByTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanMaliciousFileByTaskResponse
func (client *Client) ListScanMaliciousFileByTaskWithContext(ctx context.Context, request *ListScanMaliciousFileByTaskRequest, runtime *dara.RuntimeOptions) (_result *ListScanMaliciousFileByTaskResponse, _err error) {
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
		Action:      dara.String("ListScanMaliciousFileByTask"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScanMaliciousFileByTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询扫描规则
//
// @param request - ListScanRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanRuleResponse
func (client *Client) ListScanRuleWithContext(ctx context.Context, request *ListScanRuleRequest, runtime *dara.RuntimeOptions) (_result *ListScanRuleResponse, _err error) {
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
		Action:      dara.String("ListScanRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScanRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to cloud resources. Instance resources are supported.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Resets the logon password of a Container Registry instance.
//
// @param request - ResetLoginPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetLoginPasswordResponse
func (client *Client) ResetLoginPasswordWithContext(ctx context.Context, request *ResetLoginPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetLoginPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetLoginPassword"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetLoginPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources. Instance resources are supported.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Removes tags from resources. Instance resources are supported.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Updates a lifecycle management rule of an artifact.
//
// @param request - UpdateArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactLifecycleRuleResponse
func (client *Client) UpdateArtifactLifecycleRuleWithContext(ctx context.Context, request *UpdateArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactLifecycleRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auto) {
		query["Auto"] = request.Auto
	}

	if !dara.IsNil(request.EnableDeleteTag) {
		query["EnableDeleteTag"] = request.EnableDeleteTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RetentionTagCount) {
		query["RetentionTagCount"] = request.RetentionTagCount
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.TagRegexp) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateArtifactLifecycleRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an artifact subscription rule.
//
// @param request - UpdateArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactSubscriptionRuleResponse
func (client *Client) UpdateArtifactSubscriptionRuleWithContext(ctx context.Context, request *UpdateArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactSubscriptionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accelerate) {
		query["Accelerate"] = request.Accelerate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.Override) {
		query["Override"] = request.Override
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.SourceNamespaceName) {
		query["SourceNamespaceName"] = request.SourceNamespaceName
	}

	if !dara.IsNil(request.SourceProvider) {
		query["SourceProvider"] = request.SourceProvider
	}

	if !dara.IsNil(request.SourceRepoName) {
		query["SourceRepoName"] = request.SourceRepoName
	}

	if !dara.IsNil(request.TagCount) {
		query["TagCount"] = request.TagCount
	}

	if !dara.IsNil(request.TagRegexp) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateArtifactSubscriptionRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a delivery chain, such as the node execution sequence of the delivery chain.
//
// @param request - UpdateChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChainResponse
func (client *Client) UpdateChainWithContext(ctx context.Context, request *UpdateChainRequest, runtime *dara.RuntimeOptions) (_result *UpdateChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChainConfig) {
		query["ChainConfig"] = request.ChainConfig
	}

	if !dara.IsNil(request.ChainId) {
		query["ChainId"] = request.ChainId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ScopeExclude) {
		query["ScopeExclude"] = request.ScopeExclude
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChain"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a chart namespace.
//
// @param request - UpdateChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChartNamespaceResponse
func (client *Client) UpdateChartNamespaceWithContext(ctx context.Context, request *UpdateChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateChartNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateRepo) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !dara.IsNil(request.DefaultRepoType) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChartNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChartNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a chart repository of a Container Registry instance.
//
// @param request - UpdateChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChartRepositoryResponse
func (client *Client) UpdateChartRepositoryWithContext(ctx context.Context, request *UpdateChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateChartRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoType) {
		query["RepoType"] = request.RepoType
	}

	if !dara.IsNil(request.Summary) {
		query["Summary"] = request.Summary
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChartRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChartRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an event notification rule.
//
// @param tmpReq - UpdateEventCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventCenterRuleResponse
func (client *Client) UpdateEventCenterRuleWithContext(ctx context.Context, tmpReq *UpdateEventCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEventCenterRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Namespaces) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, dara.String("Namespaces"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepoNames) {
		request.RepoNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepoNames, dara.String("RepoNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventChannel) {
		query["EventChannel"] = request.EventChannel
	}

	if !dara.IsNil(request.EventConfig) {
		query["EventConfig"] = request.EventConfig
	}

	if !dara.IsNil(request.EventScope) {
		query["EventScope"] = request.EventScope
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespacesShrink) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !dara.IsNil(request.RepoNamesShrink) {
		query["RepoNames"] = request.RepoNamesShrink
	}

	if !dara.IsNil(request.RepoTagFilterPattern) {
		query["RepoTagFilterPattern"] = request.RepoTagFilterPattern
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEventCenterRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEventCenterRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of an instance endpoint.
//
// @param request - UpdateInstanceEndpointStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceEndpointStatusResponse
func (client *Client) UpdateInstanceEndpointStatusWithContext(ctx context.Context, request *UpdateInstanceEndpointStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceEndpointStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceEndpointStatus"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceEndpointStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a namespace.
//
// @param tmpReq - UpdateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespaceWithContext(ctx context.Context, tmpReq *UpdateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DefaultRepoConfiguration) {
		request.DefaultRepoConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRepoConfiguration, dara.String("DefaultRepoConfiguration"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateRepo) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !dara.IsNil(request.DefaultRepoConfigurationShrink) {
		query["DefaultRepoConfiguration"] = request.DefaultRepoConfigurationShrink
	}

	if !dara.IsNil(request.DefaultRepoType) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespace"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an image building rule for a repository.
//
// @param request - UpdateRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoBuildRuleResponse
func (client *Client) UpdateRepoBuildRuleWithContext(ctx context.Context, request *UpdateRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoBuildRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuildArgs) {
		query["BuildArgs"] = request.BuildArgs
	}

	if !dara.IsNil(request.BuildRuleId) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !dara.IsNil(request.DockerfileLocation) {
		query["DockerfileLocation"] = request.DockerfileLocation
	}

	if !dara.IsNil(request.DockerfileName) {
		query["DockerfileName"] = request.DockerfileName
	}

	if !dara.IsNil(request.ImageTag) {
		query["ImageTag"] = request.ImageTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Platforms) {
		query["Platforms"] = request.Platforms
	}

	if !dara.IsNil(request.PushName) {
		query["PushName"] = request.PushName
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepoBuildRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepoBuildRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the URL of the source code repository that is bound to an image repository.
//
// @param request - UpdateRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoSourceCodeRepoResponse
func (client *Client) UpdateRepoSourceCodeRepoWithContext(ctx context.Context, request *UpdateRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoSourceCodeRepoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.CodeRepoId) {
		query["CodeRepoId"] = request.CodeRepoId
	}

	if !dara.IsNil(request.CodeRepoName) {
		query["CodeRepoName"] = request.CodeRepoName
	}

	if !dara.IsNil(request.CodeRepoNamespaceName) {
		query["CodeRepoNamespaceName"] = request.CodeRepoNamespaceName
	}

	if !dara.IsNil(request.CodeRepoType) {
		query["CodeRepoType"] = request.CodeRepoType
	}

	if !dara.IsNil(request.DisableCacheBuild) {
		query["DisableCacheBuild"] = request.DisableCacheBuild
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OverseaBuild) {
		query["OverseaBuild"] = request.OverseaBuild
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepoSourceCodeRepo"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a trigger of an image repository.
//
// @param request - UpdateRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoTriggerResponse
func (client *Client) UpdateRepoTriggerWithContext(ctx context.Context, request *UpdateRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoTriggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.TriggerId) {
		query["TriggerId"] = request.TriggerId
	}

	if !dara.IsNil(request.TriggerName) {
		query["TriggerName"] = request.TriggerName
	}

	if !dara.IsNil(request.TriggerTag) {
		query["TriggerTag"] = request.TriggerTag
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.TriggerUrl) {
		query["TriggerUrl"] = request.TriggerUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepoTrigger"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepoTriggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// @param request - UpdateRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepositoryResponse
func (client *Client) UpdateRepositoryWithContext(ctx context.Context, request *UpdateRepositoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Detail) {
		query["Detail"] = request.Detail
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	if !dara.IsNil(request.RepoNamespaceName) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !dara.IsNil(request.RepoType) {
		query["RepoType"] = request.RepoType
	}

	if !dara.IsNil(request.Summary) {
		query["Summary"] = request.Summary
	}

	if !dara.IsNil(request.TagImmutability) {
		query["TagImmutability"] = request.TagImmutability
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepository"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新扫描规则
//
// @param tmpReq - UpdateScanRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScanRuleResponse
func (client *Client) UpdateScanRuleWithContext(ctx context.Context, tmpReq *UpdateScanRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateScanRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateScanRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Namespaces) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, dara.String("Namespaces"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepoNames) {
		request.RepoNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepoNames, dara.String("RepoNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespacesShrink) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !dara.IsNil(request.RepoNamesShrink) {
		query["RepoNames"] = request.RepoNamesShrink
	}

	if !dara.IsNil(request.RepoTagFilterPattern) {
		query["RepoTagFilterPattern"] = request.RepoTagFilterPattern
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.ScanRuleId) {
		query["ScanRuleId"] = request.ScanRuleId
	}

	if !dara.IsNil(request.ScanScope) {
		query["ScanScope"] = request.ScanScope
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScanRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScanRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例域名路由规则
//
// @param tmpReq - UpdateStorageDomainRoutingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStorageDomainRoutingRuleResponse
func (client *Client) UpdateStorageDomainRoutingRuleWithContext(ctx context.Context, tmpReq *UpdateStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateStorageDomainRoutingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStorageDomainRoutingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Routes) {
		request.RoutesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Routes, dara.String("Routes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RoutesShrink) {
		query["Routes"] = request.RoutesShrink
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStorageDomainRoutingRule"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStorageDomainRoutingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
