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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels an artifact building task.
//
// @param request - CancelArtifactBuildTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelArtifactBuildTaskResponse
func (client *Client) CancelArtifactBuildTaskWithOptions(request *CancelArtifactBuildTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelArtifactBuildTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an artifact building task.
//
// @param request - CancelArtifactBuildTaskRequest
//
// @return CancelArtifactBuildTaskResponse
func (client *Client) CancelArtifactBuildTask(request *CancelArtifactBuildTaskRequest) (_result *CancelArtifactBuildTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelArtifactBuildTaskResponse{}
	_body, _err := client.CancelArtifactBuildTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelRepoBuildRecordWithOptions(request *CancelRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *CancelRepoBuildRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelRepoBuildRecordResponse
func (client *Client) CancelRepoBuildRecord(request *CancelRepoBuildRecordRequest) (_result *CancelRepoBuildRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelRepoBuildRecordResponse{}
	_body, _err := client.CancelRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelRepoSyncTaskWithOptions(request *CancelRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelRepoSyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelRepoSyncTaskResponse
func (client *Client) CancelRepoSyncTask(request *CancelRepoSyncTaskRequest) (_result *CancelRepoSyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelRepoSyncTaskResponse{}
	_body, _err := client.CancelRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateArtifactBuildRuleWithOptions(tmpReq *CreateArtifactBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateArtifactBuildRuleRequest
//
// @return CreateArtifactBuildRuleResponse
func (client *Client) CreateArtifactBuildRule(request *CreateArtifactBuildRuleRequest) (_result *CreateArtifactBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArtifactBuildRuleResponse{}
	_body, _err := client.CreateArtifactBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateArtifactLifecycleRuleWithOptions(request *CreateArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactLifecycleRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateArtifactLifecycleRuleResponse
func (client *Client) CreateArtifactLifecycleRule(request *CreateArtifactLifecycleRuleRequest) (_result *CreateArtifactLifecycleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArtifactLifecycleRuleResponse{}
	_body, _err := client.CreateArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateArtifactSubscriptionRuleWithOptions(request *CreateArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactSubscriptionRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateArtifactSubscriptionRuleResponse
func (client *Client) CreateArtifactSubscriptionRule(request *CreateArtifactSubscriptionRuleRequest) (_result *CreateArtifactSubscriptionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CreateArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateArtifactSubscriptionTaskWithOptions(request *CreateArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactSubscriptionTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateArtifactSubscriptionTaskResponse
func (client *Client) CreateArtifactSubscriptionTask(request *CreateArtifactSubscriptionTaskRequest) (_result *CreateArtifactSubscriptionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArtifactSubscriptionTaskResponse{}
	_body, _err := client.CreateArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBuildRecordByRecordWithOptions(request *CreateBuildRecordByRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateBuildRecordByRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBuildRecordByRecordResponse
func (client *Client) CreateBuildRecordByRecord(request *CreateBuildRecordByRecordRequest) (_result *CreateBuildRecordByRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBuildRecordByRecordResponse{}
	_body, _err := client.CreateBuildRecordByRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBuildRecordByRuleWithOptions(request *CreateBuildRecordByRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateBuildRecordByRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBuildRecordByRuleResponse
func (client *Client) CreateBuildRecordByRule(request *CreateBuildRecordByRuleRequest) (_result *CreateBuildRecordByRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBuildRecordByRuleResponse{}
	_body, _err := client.CreateBuildRecordByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChainWithOptions(request *CreateChainRequest, runtime *dara.RuntimeOptions) (_result *CreateChainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChainResponse
func (client *Client) CreateChain(request *CreateChainRequest) (_result *CreateChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChainResponse{}
	_body, _err := client.CreateChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChartNamespaceWithOptions(request *CreateChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateChartNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChartNamespaceResponse
func (client *Client) CreateChartNamespace(request *CreateChartNamespaceRequest) (_result *CreateChartNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChartNamespaceResponse{}
	_body, _err := client.CreateChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChartRepositoryWithOptions(request *CreateChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *CreateChartRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChartRepositoryResponse
func (client *Client) CreateChartRepository(request *CreateChartRepositoryRequest) (_result *CreateChartRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChartRepositoryResponse{}
	_body, _err := client.CreateChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateInstanceEndpointAclPolicyWithOptions(request *CreateInstanceEndpointAclPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceEndpointAclPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateInstanceEndpointAclPolicyResponse
func (client *Client) CreateInstanceEndpointAclPolicy(request *CreateInstanceEndpointAclPolicyRequest) (_result *CreateInstanceEndpointAclPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CreateInstanceEndpointAclPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateInstanceVpcEndpointLinkedVpcWithOptions(request *CreateInstanceVpcEndpointLinkedVpcRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceVpcEndpointLinkedVpcResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateInstanceVpcEndpointLinkedVpcResponse
func (client *Client) CreateInstanceVpcEndpointLinkedVpc(request *CreateInstanceVpcEndpointLinkedVpcRequest) (_result *CreateInstanceVpcEndpointLinkedVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CreateInstanceVpcEndpointLinkedVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateNamespaceWithOptions(tmpReq *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoBuildRuleWithOptions(request *CreateRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoBuildRuleResponse
func (client *Client) CreateRepoBuildRule(request *CreateRepoBuildRuleRequest) (_result *CreateRepoBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoBuildRuleResponse{}
	_body, _err := client.CreateRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoSourceCodeRepoWithOptions(request *CreateRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSourceCodeRepoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoSourceCodeRepoResponse
func (client *Client) CreateRepoSourceCodeRepo(request *CreateRepoSourceCodeRepoRequest) (_result *CreateRepoSourceCodeRepoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoSourceCodeRepoResponse{}
	_body, _err := client.CreateRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoSyncRuleWithOptions(request *CreateRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoSyncRuleResponse
func (client *Client) CreateRepoSyncRule(request *CreateRepoSyncRuleRequest) (_result *CreateRepoSyncRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoSyncRuleResponse{}
	_body, _err := client.CreateRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncTaskResponse
func (client *Client) CreateRepoSyncTaskWithOptions(request *CreateRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRepoSyncTaskRequest
//
// @return CreateRepoSyncTaskResponse
func (client *Client) CreateRepoSyncTask(request *CreateRepoSyncTaskRequest) (_result *CreateRepoSyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoSyncTaskResponse{}
	_body, _err := client.CreateRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoSyncTaskByRuleWithOptions(request *CreateRepoSyncTaskByRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoSyncTaskByRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoSyncTaskByRuleResponse
func (client *Client) CreateRepoSyncTaskByRule(request *CreateRepoSyncTaskByRuleRequest) (_result *CreateRepoSyncTaskByRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoSyncTaskByRuleResponse{}
	_body, _err := client.CreateRepoSyncTaskByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoTagWithOptions(request *CreateRepoTagRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTagResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoTagResponse
func (client *Client) CreateRepoTag(request *CreateRepoTagRequest) (_result *CreateRepoTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoTagResponse{}
	_body, _err := client.CreateRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoTagScanTaskWithOptions(request *CreateRepoTagScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTagScanTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoTagScanTaskResponse
func (client *Client) CreateRepoTagScanTask(request *CreateRepoTagScanTaskRequest) (_result *CreateRepoTagScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoTagScanTaskResponse{}
	_body, _err := client.CreateRepoTagScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepoTriggerWithOptions(request *CreateRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *CreateRepoTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepoTriggerResponse
func (client *Client) CreateRepoTrigger(request *CreateRepoTriggerRequest) (_result *CreateRepoTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepoTriggerResponse{}
	_body, _err := client.CreateRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRepositoryWithOptions(request *CreateRepositoryRequest, runtime *dara.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRepositoryResponse
func (client *Client) CreateRepository(request *CreateRepositoryRequest) (_result *CreateRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CreateRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateScanRuleWithOptions(tmpReq *CreateScanRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateScanRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateScanRuleRequest
//
// @return CreateScanRuleResponse
func (client *Client) CreateScanRule(request *CreateScanRuleRequest) (_result *CreateScanRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScanRuleResponse{}
	_body, _err := client.CreateScanRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateStorageDomainRoutingRuleWithOptions(tmpReq *CreateStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateStorageDomainRoutingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateStorageDomainRoutingRuleRequest
//
// @return CreateStorageDomainRoutingRuleResponse
func (client *Client) CreateStorageDomainRoutingRule(request *CreateStorageDomainRoutingRuleRequest) (_result *CreateStorageDomainRoutingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStorageDomainRoutingRuleResponse{}
	_body, _err := client.CreateStorageDomainRoutingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteArtifactLifecycleRuleWithOptions(request *DeleteArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactLifecycleRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteArtifactLifecycleRuleResponse
func (client *Client) DeleteArtifactLifecycleRule(request *DeleteArtifactLifecycleRuleRequest) (_result *DeleteArtifactLifecycleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteArtifactLifecycleRuleResponse{}
	_body, _err := client.DeleteArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteArtifactSubscriptionRuleWithOptions(request *DeleteArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactSubscriptionRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteArtifactSubscriptionRuleResponse
func (client *Client) DeleteArtifactSubscriptionRule(request *DeleteArtifactSubscriptionRuleRequest) (_result *DeleteArtifactSubscriptionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteArtifactSubscriptionRuleResponse{}
	_body, _err := client.DeleteArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChainWithOptions(request *DeleteChainRequest, runtime *dara.RuntimeOptions) (_result *DeleteChainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChainResponse
func (client *Client) DeleteChain(request *DeleteChainRequest) (_result *DeleteChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChainResponse{}
	_body, _err := client.DeleteChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChartNamespaceWithOptions(request *DeleteChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChartNamespaceResponse
func (client *Client) DeleteChartNamespace(request *DeleteChartNamespaceRequest) (_result *DeleteChartNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChartNamespaceResponse{}
	_body, _err := client.DeleteChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChartReleaseWithOptions(request *DeleteChartReleaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartReleaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChartReleaseResponse
func (client *Client) DeleteChartRelease(request *DeleteChartReleaseRequest) (_result *DeleteChartReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChartReleaseResponse{}
	_body, _err := client.DeleteChartReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChartRepositoryWithOptions(request *DeleteChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteChartRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChartRepositoryResponse
func (client *Client) DeleteChartRepository(request *DeleteChartRepositoryRequest) (_result *DeleteChartRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChartRepositoryResponse{}
	_body, _err := client.DeleteChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEventCenterRuleWithOptions(request *DeleteEventCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventCenterRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEventCenterRuleResponse
func (client *Client) DeleteEventCenterRule(request *DeleteEventCenterRuleRequest) (_result *DeleteEventCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventCenterRuleResponse{}
	_body, _err := client.DeleteEventCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteInstanceEndpointAclPolicyWithOptions(request *DeleteInstanceEndpointAclPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceEndpointAclPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteInstanceEndpointAclPolicyResponse
func (client *Client) DeleteInstanceEndpointAclPolicy(request *DeleteInstanceEndpointAclPolicyRequest) (_result *DeleteInstanceEndpointAclPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceEndpointAclPolicyResponse{}
	_body, _err := client.DeleteInstanceEndpointAclPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteInstanceVpcEndpointLinkedVpcWithOptions(request *DeleteInstanceVpcEndpointLinkedVpcRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceVpcEndpointLinkedVpcResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteInstanceVpcEndpointLinkedVpcResponse
func (client *Client) DeleteInstanceVpcEndpointLinkedVpc(request *DeleteInstanceVpcEndpointLinkedVpcRequest) (_result *DeleteInstanceVpcEndpointLinkedVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.DeleteInstanceVpcEndpointLinkedVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRepoBuildRuleWithOptions(request *DeleteRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRepoBuildRuleResponse
func (client *Client) DeleteRepoBuildRule(request *DeleteRepoBuildRuleRequest) (_result *DeleteRepoBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRepoBuildRuleResponse{}
	_body, _err := client.DeleteRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRepoSyncRuleWithOptions(request *DeleteRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoSyncRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRepoSyncRuleResponse
func (client *Client) DeleteRepoSyncRule(request *DeleteRepoSyncRuleRequest) (_result *DeleteRepoSyncRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRepoSyncRuleResponse{}
	_body, _err := client.DeleteRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRepoTagWithOptions(request *DeleteRepoTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoTagResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRepoTagResponse
func (client *Client) DeleteRepoTag(request *DeleteRepoTagRequest) (_result *DeleteRepoTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRepoTagResponse{}
	_body, _err := client.DeleteRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRepoTriggerWithOptions(request *DeleteRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepoTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRepoTriggerResponse
func (client *Client) DeleteRepoTrigger(request *DeleteRepoTriggerRequest) (_result *DeleteRepoTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRepoTriggerResponse{}
	_body, _err := client.DeleteRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRepositoryWithOptions(request *DeleteRepositoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepository(request *DeleteRepositoryRequest) (_result *DeleteRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DeleteRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteScanRuleWithOptions(request *DeleteScanRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteScanRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteScanRuleResponse
func (client *Client) DeleteScanRule(request *DeleteScanRuleRequest) (_result *DeleteScanRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScanRuleResponse{}
	_body, _err := client.DeleteScanRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteStorageDomainRoutingRuleWithOptions(request *DeleteStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteStorageDomainRoutingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteStorageDomainRoutingRuleResponse
func (client *Client) DeleteStorageDomainRoutingRule(request *DeleteStorageDomainRoutingRuleRequest) (_result *DeleteStorageDomainRoutingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStorageDomainRoutingRuleResponse{}
	_body, _err := client.DeleteStorageDomainRoutingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetArtifactBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactBuildRuleResponse
func (client *Client) GetArtifactBuildRuleWithOptions(request *GetArtifactBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArtifactBuildRuleRequest
//
// @return GetArtifactBuildRuleResponse
func (client *Client) GetArtifactBuildRule(request *GetArtifactBuildRuleRequest) (_result *GetArtifactBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactBuildRuleResponse{}
	_body, _err := client.GetArtifactBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactBuildTaskWithOptions(request *GetArtifactBuildTaskRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactBuildTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactBuildTaskResponse
func (client *Client) GetArtifactBuildTask(request *GetArtifactBuildTaskRequest) (_result *GetArtifactBuildTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactBuildTaskResponse{}
	_body, _err := client.GetArtifactBuildTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactLifecycleRuleWithOptions(request *GetArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactLifecycleRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactLifecycleRuleResponse
func (client *Client) GetArtifactLifecycleRule(request *GetArtifactLifecycleRuleRequest) (_result *GetArtifactLifecycleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactLifecycleRuleResponse{}
	_body, _err := client.GetArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactSubscriptionRuleWithOptions(request *GetArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactSubscriptionRuleResponse
func (client *Client) GetArtifactSubscriptionRule(request *GetArtifactSubscriptionRuleRequest) (_result *GetArtifactSubscriptionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactSubscriptionRuleResponse{}
	_body, _err := client.GetArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactSubscriptionTaskWithOptions(request *GetArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactSubscriptionTaskResponse
func (client *Client) GetArtifactSubscriptionTask(request *GetArtifactSubscriptionTaskRequest) (_result *GetArtifactSubscriptionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactSubscriptionTaskResponse{}
	_body, _err := client.GetArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactSubscriptionTaskResultWithOptions(request *GetArtifactSubscriptionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactSubscriptionTaskResultResponse
func (client *Client) GetArtifactSubscriptionTaskResult(request *GetArtifactSubscriptionTaskResultRequest) (_result *GetArtifactSubscriptionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactSubscriptionTaskResultResponse{}
	_body, _err := client.GetArtifactSubscriptionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAuthorizationTokenWithOptions(request *GetAuthorizationTokenRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationTokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAuthorizationTokenResponse
func (client *Client) GetAuthorizationToken(request *GetAuthorizationTokenRequest) (_result *GetAuthorizationTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuthorizationTokenResponse{}
	_body, _err := client.GetAuthorizationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChainWithOptions(request *GetChainRequest, runtime *dara.RuntimeOptions) (_result *GetChainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChainResponse
func (client *Client) GetChain(request *GetChainRequest) (_result *GetChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChainResponse{}
	_body, _err := client.GetChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChartNamespaceWithOptions(request *GetChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetChartNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChartNamespaceResponse
func (client *Client) GetChartNamespace(request *GetChartNamespaceRequest) (_result *GetChartNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChartNamespaceResponse{}
	_body, _err := client.GetChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChartRepositoryWithOptions(request *GetChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *GetChartRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChartRepositoryResponse
func (client *Client) GetChartRepository(request *GetChartRepositoryRequest) (_result *GetChartRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChartRepositoryResponse{}
	_body, _err := client.GetChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of instances.
//
// @param request - GetInstanceCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceCountResponse
func (client *Client) GetInstanceCountWithOptions(runtime *dara.RuntimeOptions) (_result *GetInstanceCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceCount"),
		Version:     dara.String("2018-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of instances.
//
// @return GetInstanceCountResponse
func (client *Client) GetInstanceCount() (_result *GetInstanceCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceCountResponse{}
	_body, _err := client.GetInstanceCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceEndpointWithOptions(request *GetInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceEndpointResponse
func (client *Client) GetInstanceEndpoint(request *GetInstanceEndpointRequest) (_result *GetInstanceEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceEndpointResponse{}
	_body, _err := client.GetInstanceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceUsageWithOptions(request *GetInstanceUsageRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUsageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceUsageResponse
func (client *Client) GetInstanceUsage(request *GetInstanceUsageRequest) (_result *GetInstanceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceUsageResponse{}
	_body, _err := client.GetInstanceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceVpcEndpointWithOptions(request *GetInstanceVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceVpcEndpointResponse
func (client *Client) GetInstanceVpcEndpoint(request *GetInstanceVpcEndpointRequest) (_result *GetInstanceVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceVpcEndpointResponse{}
	_body, _err := client.GetInstanceVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNamespaceWithOptions(request *GetNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetNamespaceResponse
func (client *Client) GetNamespace(request *GetNamespaceRequest) (_result *GetNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNamespaceResponse{}
	_body, _err := client.GetNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoBuildRecordWithOptions(request *GetRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *GetRepoBuildRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoBuildRecordResponse
func (client *Client) GetRepoBuildRecord(request *GetRepoBuildRecordRequest) (_result *GetRepoBuildRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoBuildRecordResponse{}
	_body, _err := client.GetRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoBuildRecordStatusWithOptions(request *GetRepoBuildRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRepoBuildRecordStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoBuildRecordStatusResponse
func (client *Client) GetRepoBuildRecordStatus(request *GetRepoBuildRecordStatusRequest) (_result *GetRepoBuildRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoBuildRecordStatusResponse{}
	_body, _err := client.GetRepoBuildRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoSourceCodeRepoWithOptions(request *GetRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *GetRepoSourceCodeRepoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoSourceCodeRepoResponse
func (client *Client) GetRepoSourceCodeRepo(request *GetRepoSourceCodeRepoRequest) (_result *GetRepoSourceCodeRepoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoSourceCodeRepoResponse{}
	_body, _err := client.GetRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoSyncTaskWithOptions(request *GetRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetRepoSyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoSyncTaskResponse
func (client *Client) GetRepoSyncTask(request *GetRepoSyncTaskRequest) (_result *GetRepoSyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoSyncTaskResponse{}
	_body, _err := client.GetRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoTagWithOptions(request *GetRepoTagRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoTagResponse
func (client *Client) GetRepoTag(request *GetRepoTagRequest) (_result *GetRepoTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoTagResponse{}
	_body, _err := client.GetRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepoTagScanStatusWithOptions(request *GetRepoTagScanStatusRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagScanStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepoTagScanStatusResponse
func (client *Client) GetRepoTagScanStatus(request *GetRepoTagScanStatusRequest) (_result *GetRepoTagScanStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoTagScanStatusResponse{}
	_body, _err := client.GetRepoTagScanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRepoTagScanSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagScanSummaryResponse
func (client *Client) GetRepoTagScanSummaryWithOptions(request *GetRepoTagScanSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetRepoTagScanSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRepoTagScanSummaryRequest
//
// @return GetRepoTagScanSummaryResponse
func (client *Client) GetRepoTagScanSummary(request *GetRepoTagScanSummaryRequest) (_result *GetRepoTagScanSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepoTagScanSummaryResponse{}
	_body, _err := client.GetRepoTagScanSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRepositoryWithOptions(request *GetRepositoryRequest, runtime *dara.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRepositoryResponse
func (client *Client) GetRepository(request *GetRepositoryRequest) (_result *GetRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRepositoryResponse{}
	_body, _err := client.GetRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetScanRuleWithOptions(request *GetScanRuleRequest, runtime *dara.RuntimeOptions) (_result *GetScanRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetScanRuleResponse
func (client *Client) GetScanRule(request *GetScanRuleRequest) (_result *GetScanRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScanRuleResponse{}
	_body, _err := client.GetScanRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetStorageDomainRoutingRuleWithOptions(request *GetStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetStorageDomainRoutingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetStorageDomainRoutingRuleResponse
func (client *Client) GetStorageDomainRoutingRule(request *GetStorageDomainRoutingRuleRequest) (_result *GetStorageDomainRoutingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStorageDomainRoutingRuleResponse{}
	_body, _err := client.GetStorageDomainRoutingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactBuildTaskLogWithOptions(request *ListArtifactBuildTaskLogRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactBuildTaskLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactBuildTaskLogResponse
func (client *Client) ListArtifactBuildTaskLog(request *ListArtifactBuildTaskLogRequest) (_result *ListArtifactBuildTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactBuildTaskLogResponse{}
	_body, _err := client.ListArtifactBuildTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactLifecycleRuleWithOptions(request *ListArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactLifecycleRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactLifecycleRuleResponse
func (client *Client) ListArtifactLifecycleRule(request *ListArtifactLifecycleRuleRequest) (_result *ListArtifactLifecycleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactLifecycleRuleResponse{}
	_body, _err := client.ListArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactSubscriptionRuleWithOptions(request *ListArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactSubscriptionRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactSubscriptionRuleResponse
func (client *Client) ListArtifactSubscriptionRule(request *ListArtifactSubscriptionRuleRequest) (_result *ListArtifactSubscriptionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactSubscriptionRuleResponse{}
	_body, _err := client.ListArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactSubscriptionTaskWithOptions(request *ListArtifactSubscriptionTaskRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactSubscriptionTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactSubscriptionTaskResponse
func (client *Client) ListArtifactSubscriptionTask(request *ListArtifactSubscriptionTaskRequest) (_result *ListArtifactSubscriptionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactSubscriptionTaskResponse{}
	_body, _err := client.ListArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChainWithOptions(request *ListChainRequest, runtime *dara.RuntimeOptions) (_result *ListChainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChainResponse
func (client *Client) ListChain(request *ListChainRequest) (_result *ListChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChainResponse{}
	_body, _err := client.ListChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChainInstanceWithOptions(request *ListChainInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListChainInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChainInstanceResponse
func (client *Client) ListChainInstance(request *ListChainInstanceRequest) (_result *ListChainInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChainInstanceResponse{}
	_body, _err := client.ListChainInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChartNamespaceWithOptions(request *ListChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ListChartNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChartNamespaceResponse
func (client *Client) ListChartNamespace(request *ListChartNamespaceRequest) (_result *ListChartNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChartNamespaceResponse{}
	_body, _err := client.ListChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChartReleaseWithOptions(request *ListChartReleaseRequest, runtime *dara.RuntimeOptions) (_result *ListChartReleaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChartReleaseResponse
func (client *Client) ListChartRelease(request *ListChartReleaseRequest) (_result *ListChartReleaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChartReleaseResponse{}
	_body, _err := client.ListChartReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChartRepositoryWithOptions(request *ListChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *ListChartRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChartRepositoryResponse
func (client *Client) ListChartRepository(request *ListChartRepositoryRequest) (_result *ListChartRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChartRepositoryResponse{}
	_body, _err := client.ListChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEventCenterRecordWithOptions(request *ListEventCenterRecordRequest, runtime *dara.RuntimeOptions) (_result *ListEventCenterRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEventCenterRecordResponse
func (client *Client) ListEventCenterRecord(request *ListEventCenterRecordRequest) (_result *ListEventCenterRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEventCenterRecordResponse{}
	_body, _err := client.ListEventCenterRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEventCenterRuleNameWithOptions(request *ListEventCenterRuleNameRequest, runtime *dara.RuntimeOptions) (_result *ListEventCenterRuleNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEventCenterRuleNameResponse
func (client *Client) ListEventCenterRuleName(request *ListEventCenterRuleNameRequest) (_result *ListEventCenterRuleNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEventCenterRuleNameResponse{}
	_body, _err := client.ListEventCenterRuleNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInstanceEndpointWithOptions(request *ListInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInstanceEndpointResponse
func (client *Client) ListInstanceEndpoint(request *ListInstanceEndpointRequest) (_result *ListInstanceEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceEndpointResponse{}
	_body, _err := client.ListInstanceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListInstanceRegionWithOptions(request *ListInstanceRegionRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceRegionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListInstanceRegionResponse
func (client *Client) ListInstanceRegion(request *ListInstanceRegionRequest) (_result *ListInstanceRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceRegionResponse{}
	_body, _err := client.ListInstanceRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListNamespaceWithOptions(request *ListNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ListNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListNamespaceResponse
func (client *Client) ListNamespace(request *ListNamespaceRequest) (_result *ListNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamespaceResponse{}
	_body, _err := client.ListNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoBuildRecordWithOptions(request *ListRepoBuildRecordRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoBuildRecordResponse
func (client *Client) ListRepoBuildRecord(request *ListRepoBuildRecordRequest) (_result *ListRepoBuildRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoBuildRecordResponse{}
	_body, _err := client.ListRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoBuildRecordLogWithOptions(request *ListRepoBuildRecordLogRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRecordLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoBuildRecordLogResponse
func (client *Client) ListRepoBuildRecordLog(request *ListRepoBuildRecordLogRequest) (_result *ListRepoBuildRecordLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoBuildRecordLogResponse{}
	_body, _err := client.ListRepoBuildRecordLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoBuildRuleWithOptions(request *ListRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *ListRepoBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoBuildRuleResponse
func (client *Client) ListRepoBuildRule(request *ListRepoBuildRuleRequest) (_result *ListRepoBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoBuildRuleResponse{}
	_body, _err := client.ListRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoSyncRuleWithOptions(request *ListRepoSyncRuleRequest, runtime *dara.RuntimeOptions) (_result *ListRepoSyncRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoSyncRuleResponse
func (client *Client) ListRepoSyncRule(request *ListRepoSyncRuleRequest) (_result *ListRepoSyncRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoSyncRuleResponse{}
	_body, _err := client.ListRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoSyncTaskWithOptions(request *ListRepoSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *ListRepoSyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoSyncTaskResponse
func (client *Client) ListRepoSyncTask(request *ListRepoSyncTaskRequest) (_result *ListRepoSyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoSyncTaskResponse{}
	_body, _err := client.ListRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoTagWithOptions(request *ListRepoTagRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTagResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoTagResponse
func (client *Client) ListRepoTag(request *ListRepoTagRequest) (_result *ListRepoTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoTagResponse{}
	_body, _err := client.ListRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoTagScanResultWithOptions(request *ListRepoTagScanResultRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTagScanResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoTagScanResultResponse
func (client *Client) ListRepoTagScanResult(request *ListRepoTagScanResultRequest) (_result *ListRepoTagScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoTagScanResultResponse{}
	_body, _err := client.ListRepoTagScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepoTriggerWithOptions(request *ListRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *ListRepoTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepoTriggerResponse
func (client *Client) ListRepoTrigger(request *ListRepoTriggerRequest) (_result *ListRepoTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepoTriggerResponse{}
	_body, _err := client.ListRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRepositoryWithOptions(request *ListRepositoryRequest, runtime *dara.RuntimeOptions) (_result *ListRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRepositoryResponse
func (client *Client) ListRepository(request *ListRepositoryRequest) (_result *ListRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRepositoryResponse{}
	_body, _err := client.ListRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListScanBaselineByTaskWithOptions(request *ListScanBaselineByTaskRequest, runtime *dara.RuntimeOptions) (_result *ListScanBaselineByTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListScanBaselineByTaskResponse
func (client *Client) ListScanBaselineByTask(request *ListScanBaselineByTaskRequest) (_result *ListScanBaselineByTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScanBaselineByTaskResponse{}
	_body, _err := client.ListScanBaselineByTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListScanMaliciousFileByTaskWithOptions(request *ListScanMaliciousFileByTaskRequest, runtime *dara.RuntimeOptions) (_result *ListScanMaliciousFileByTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListScanMaliciousFileByTaskResponse
func (client *Client) ListScanMaliciousFileByTask(request *ListScanMaliciousFileByTaskRequest) (_result *ListScanMaliciousFileByTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScanMaliciousFileByTaskResponse{}
	_body, _err := client.ListScanMaliciousFileByTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListScanRuleWithOptions(request *ListScanRuleRequest, runtime *dara.RuntimeOptions) (_result *ListScanRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListScanRuleResponse
func (client *Client) ListScanRule(request *ListScanRuleRequest) (_result *ListScanRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScanRuleResponse{}
	_body, _err := client.ListScanRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetLoginPasswordWithOptions(request *ResetLoginPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetLoginPasswordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ResetLoginPasswordResponse
func (client *Client) ResetLoginPassword(request *ResetLoginPasswordRequest) (_result *ResetLoginPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetLoginPasswordResponse{}
	_body, _err := client.ResetLoginPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateArtifactLifecycleRuleWithOptions(request *UpdateArtifactLifecycleRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactLifecycleRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateArtifactLifecycleRuleResponse
func (client *Client) UpdateArtifactLifecycleRule(request *UpdateArtifactLifecycleRuleRequest) (_result *UpdateArtifactLifecycleRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateArtifactLifecycleRuleResponse{}
	_body, _err := client.UpdateArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateArtifactSubscriptionRuleWithOptions(request *UpdateArtifactSubscriptionRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactSubscriptionRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateArtifactSubscriptionRuleResponse
func (client *Client) UpdateArtifactSubscriptionRule(request *UpdateArtifactSubscriptionRuleRequest) (_result *UpdateArtifactSubscriptionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateArtifactSubscriptionRuleResponse{}
	_body, _err := client.UpdateArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChainWithOptions(request *UpdateChainRequest, runtime *dara.RuntimeOptions) (_result *UpdateChainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateChainResponse
func (client *Client) UpdateChain(request *UpdateChainRequest) (_result *UpdateChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChainResponse{}
	_body, _err := client.UpdateChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChartNamespaceWithOptions(request *UpdateChartNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateChartNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateChartNamespaceResponse
func (client *Client) UpdateChartNamespace(request *UpdateChartNamespaceRequest) (_result *UpdateChartNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChartNamespaceResponse{}
	_body, _err := client.UpdateChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChartRepositoryWithOptions(request *UpdateChartRepositoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateChartRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateChartRepositoryResponse
func (client *Client) UpdateChartRepository(request *UpdateChartRepositoryRequest) (_result *UpdateChartRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChartRepositoryResponse{}
	_body, _err := client.UpdateChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateEventCenterRuleWithOptions(tmpReq *UpdateEventCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateEventCenterRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateEventCenterRuleRequest
//
// @return UpdateEventCenterRuleResponse
func (client *Client) UpdateEventCenterRule(request *UpdateEventCenterRuleRequest) (_result *UpdateEventCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEventCenterRuleResponse{}
	_body, _err := client.UpdateEventCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateInstanceEndpointStatusWithOptions(request *UpdateInstanceEndpointStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceEndpointStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateInstanceEndpointStatusResponse
func (client *Client) UpdateInstanceEndpointStatus(request *UpdateInstanceEndpointStatusRequest) (_result *UpdateInstanceEndpointStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceEndpointStatusResponse{}
	_body, _err := client.UpdateInstanceEndpointStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateNamespaceWithOptions(tmpReq *UpdateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateNamespaceRequest
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRepoBuildRuleWithOptions(request *UpdateRepoBuildRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoBuildRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRepoBuildRuleResponse
func (client *Client) UpdateRepoBuildRule(request *UpdateRepoBuildRuleRequest) (_result *UpdateRepoBuildRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRepoBuildRuleResponse{}
	_body, _err := client.UpdateRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRepoSourceCodeRepoWithOptions(request *UpdateRepoSourceCodeRepoRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoSourceCodeRepoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRepoSourceCodeRepoResponse
func (client *Client) UpdateRepoSourceCodeRepo(request *UpdateRepoSourceCodeRepoRequest) (_result *UpdateRepoSourceCodeRepoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRepoSourceCodeRepoResponse{}
	_body, _err := client.UpdateRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRepoTriggerWithOptions(request *UpdateRepoTriggerRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepoTriggerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRepoTriggerResponse
func (client *Client) UpdateRepoTrigger(request *UpdateRepoTriggerRequest) (_result *UpdateRepoTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRepoTriggerResponse{}
	_body, _err := client.UpdateRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRepositoryWithOptions(request *UpdateRepositoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRepositoryResponse
func (client *Client) UpdateRepository(request *UpdateRepositoryRequest) (_result *UpdateRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.UpdateRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateScanRuleWithOptions(tmpReq *UpdateScanRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateScanRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateScanRuleRequest
//
// @return UpdateScanRuleResponse
func (client *Client) UpdateScanRule(request *UpdateScanRuleRequest) (_result *UpdateScanRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScanRuleResponse{}
	_body, _err := client.UpdateScanRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateStorageDomainRoutingRuleWithOptions(tmpReq *UpdateStorageDomainRoutingRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateStorageDomainRoutingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateStorageDomainRoutingRuleRequest
//
// @return UpdateStorageDomainRoutingRuleResponse
func (client *Client) UpdateStorageDomainRoutingRule(request *UpdateStorageDomainRoutingRuleRequest) (_result *UpdateStorageDomainRoutingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStorageDomainRoutingRuleResponse{}
	_body, _err := client.UpdateStorageDomainRoutingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
