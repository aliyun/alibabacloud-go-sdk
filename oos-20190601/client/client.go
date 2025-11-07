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
	client.Endpoint, _err = client.GetEndpoint(dara.String("oos"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 分析仓库
//
// @param request - AnalyzeGitRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeGitRepositoryResponse
func (client *Client) AnalyzeGitRepositoryWithOptions(request *AnalyzeGitRepositoryRequest, runtime *dara.RuntimeOptions) (_result *AnalyzeGitRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Branch) {
		query["Branch"] = request.Branch
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeGitRepository"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeGitRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分析仓库
//
// @param request - AnalyzeGitRepositoryRequest
//
// @return AnalyzeGitRepositoryResponse
func (client *Client) AnalyzeGitRepository(request *AnalyzeGitRepositoryRequest) (_result *AnalyzeGitRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnalyzeGitRepositoryResponse{}
	_body, _err := client.AnalyzeGitRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an execution.
//
// @param request - CancelExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelExecutionResponse
func (client *Client) CancelExecutionWithOptions(request *CancelExecutionRequest, runtime *dara.RuntimeOptions) (_result *CancelExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelExecution"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an execution.
//
// @param request - CancelExecutionRequest
//
// @return CancelExecutionResponse
func (client *Client) CancelExecution(request *CancelExecutionRequest) (_result *CancelExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelExecutionResponse{}
	_body, _err := client.CancelExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a cloud resource belongs.
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
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2019-06-01"),
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
// Modifies the resource group to which a cloud resource belongs.
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
// 检测仓库中文件是否存在
//
// @param request - CheckGitRepoFileExistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckGitRepoFileExistsResponse
func (client *Client) CheckGitRepoFileExistsWithOptions(request *CheckGitRepoFileExistsRequest, runtime *dara.RuntimeOptions) (_result *CheckGitRepoFileExistsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Branch) {
		query["Branch"] = request.Branch
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckGitRepoFileExists"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckGitRepoFileExistsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测仓库中文件是否存在
//
// @param request - CheckGitRepoFileExistsRequest
//
// @return CheckGitRepoFileExistsResponse
func (client *Client) CheckGitRepoFileExists(request *CheckGitRepoFileExistsRequest) (_result *CheckGitRepoFileExistsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckGitRepoFileExistsResponse{}
	_body, _err := client.CheckGitRepoFileExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检测仓库是否存在
//
// @param request - CheckGitRepositoryExistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckGitRepositoryExistsResponse
func (client *Client) CheckGitRepositoryExistsWithOptions(request *CheckGitRepositoryExistsRequest, runtime *dara.RuntimeOptions) (_result *CheckGitRepositoryExistsResponse, _err error) {
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

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckGitRepositoryExists"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckGitRepositoryExistsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测仓库是否存在
//
// @param request - CheckGitRepositoryExistsRequest
//
// @return CheckGitRepositoryExistsResponse
func (client *Client) CheckGitRepositoryExists(request *CheckGitRepositoryExistsRequest) (_result *CheckGitRepositoryExistsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckGitRepositoryExistsResponse{}
	_body, _err := client.CheckGitRepositoryExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Continues deploying an application group when an error occurs for calling the DeployApplicationGroup operation. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ContinueDeployApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployApplicationGroupResponse
func (client *Client) ContinueDeployApplicationGroupWithOptions(request *ContinueDeployApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *ContinueDeployApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.DeployParameters) {
		query["DeployParameters"] = request.DeployParameters
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueDeployApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueDeployApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continues deploying an application group when an error occurs for calling the DeployApplicationGroup operation. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ContinueDeployApplicationGroupRequest
//
// @return ContinueDeployApplicationGroupResponse
func (client *Client) ContinueDeployApplicationGroup(request *ContinueDeployApplicationGroupRequest) (_result *ContinueDeployApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinueDeployApplicationGroupResponse{}
	_body, _err := client.ContinueDeployApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlarmConfig) {
		request.AlarmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlarmConfig, dara.String("AlarmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmConfigShrink) {
		query["AlarmConfig"] = request.AlarmConfigShrink
	}

	if !dara.IsNil(request.ApplicationSource) {
		query["ApplicationSource"] = request.ApplicationSource
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationGroupResponse
func (client *Client) CreateApplicationGroupWithOptions(request *CreateApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CmsGroupId) {
		query["CmsGroupId"] = request.CmsGroupId
	}

	if !dara.IsNil(request.DeployRegionId) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImportTagKey) {
		query["ImportTagKey"] = request.ImportTagKey
	}

	if !dara.IsNil(request.ImportTagValue) {
		query["ImportTagValue"] = request.ImportTagValue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationGroupRequest
//
// @return CreateApplicationGroupResponse
func (client *Client) CreateApplicationGroup(request *CreateApplicationGroupRequest) (_result *CreateApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.CreateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建仓库
//
// @param request - CreateGitRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGitRepositoryResponse
func (client *Client) CreateGitRepositoryWithOptions(request *CreateGitRepositoryRequest, runtime *dara.RuntimeOptions) (_result *CreateGitRepositoryResponse, _err error) {
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

	if !dara.IsNil(request.IsPrivate) {
		query["IsPrivate"] = request.IsPrivate
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceRepoBranch) {
		query["SourceRepoBranch"] = request.SourceRepoBranch
	}

	if !dara.IsNil(request.SourceRepoName) {
		query["SourceRepoName"] = request.SourceRepoName
	}

	if !dara.IsNil(request.SourceRepoOwner) {
		query["SourceRepoOwner"] = request.SourceRepoOwner
	}

	if !dara.IsNil(request.TargetRepoName) {
		query["TargetRepoName"] = request.TargetRepoName
	}

	if !dara.IsNil(request.TargetRepoOwner) {
		query["TargetRepoOwner"] = request.TargetRepoOwner
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGitRepository"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGitRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建仓库
//
// @param request - CreateGitRepositoryRequest
//
// @return CreateGitRepositoryResponse
func (client *Client) CreateGitRepository(request *CreateGitRepositoryRequest) (_result *CreateGitRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGitRepositoryResponse{}
	_body, _err := client.CreateGitRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an O\\\\\\\\\\\\&M Item.
//
// @param tmpReq - CreateOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpsItemResponse
func (client *Client) CreateOpsItemWithOptions(tmpReq *CreateOpsItemRequest, runtime *dara.RuntimeOptions) (_result *CreateOpsItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOpsItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DedupString) {
		query["DedupString"] = request.DedupString
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Solutions) {
		query["Solutions"] = request.Solutions
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpsItem"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an O\\\\\\\\\\\\&M Item.
//
// @param request - CreateOpsItemRequest
//
// @return CreateOpsItemResponse
func (client *Client) CreateOpsItem(request *CreateOpsItemRequest) (_result *CreateOpsItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOpsItemResponse{}
	_body, _err := client.CreateOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a common parameter.
//
// @param tmpReq - CreateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterResponse
func (client *Client) CreateParameterWithOptions(tmpReq *CreateParameterRequest, runtime *dara.RuntimeOptions) (_result *CreateParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Constraints) {
		query["Constraints"] = request.Constraints
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a common parameter.
//
// @param request - CreateParameterRequest
//
// @return CreateParameterResponse
func (client *Client) CreateParameter(request *CreateParameterRequest) (_result *CreateParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateParameterResponse{}
	_body, _err := client.CreateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a patch baseline.
//
// @param tmpReq - CreatePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePatchBaselineResponse
func (client *Client) CreatePatchBaselineWithOptions(tmpReq *CreatePatchBaselineRequest, runtime *dara.RuntimeOptions) (_result *CreatePatchBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePatchBaselineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovedPatches) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, dara.String("ApprovedPatches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RejectedPatches) {
		request.RejectedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RejectedPatches, dara.String("RejectedPatches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalRules) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !dara.IsNil(request.ApprovedPatchesShrink) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !dara.IsNil(request.ApprovedPatchesEnableNonSecurity) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OperationSystem) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RejectedPatchesShrink) {
		query["RejectedPatches"] = request.RejectedPatchesShrink
	}

	if !dara.IsNil(request.RejectedPatchesAction) {
		query["RejectedPatchesAction"] = request.RejectedPatchesAction
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePatchBaseline"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a patch baseline.
//
// @param request - CreatePatchBaselineRequest
//
// @return CreatePatchBaselineResponse
func (client *Client) CreatePatchBaseline(request *CreatePatchBaselineRequest) (_result *CreatePatchBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.CreatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an encryption parameter. Make sure that you have the permissions to call this operation.
//
// @param tmpReq - CreateSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretParameterResponse
func (client *Client) CreateSecretParameterWithOptions(tmpReq *CreateSecretParameterRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecretParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Constraints) {
		query["Constraints"] = request.Constraints
	}

	if !dara.IsNil(request.DKMSInstanceId) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecretParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an encryption parameter. Make sure that you have the permissions to call this operation.
//
// @param request - CreateSecretParameterRequest
//
// @return CreateSecretParameterResponse
func (client *Client) CreateSecretParameter(request *CreateSecretParameterRequest) (_result *CreateSecretParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.CreateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a desired-state configuration.
//
// @param tmpReq - CreateStateConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStateConfigurationResponse
func (client *Client) CreateStateConfigurationWithOptions(tmpReq *CreateStateConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateStateConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigureMode) {
		query["ConfigureMode"] = request.ConfigureMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleExpression) {
		query["ScheduleExpression"] = request.ScheduleExpression
	}

	if !dara.IsNil(request.ScheduleType) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStateConfiguration"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a desired-state configuration.
//
// @param request - CreateStateConfigurationRequest
//
// @return CreateStateConfigurationResponse
func (client *Client) CreateStateConfiguration(request *CreateStateConfigurationRequest) (_result *CreateStateConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.CreateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// @param tmpReq - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(tmpReq *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetainResource) {
		query["RetainResource"] = request.RetainResource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application group. You can call this operation only for the application groups which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationGroupResponse
func (client *Client) DeleteApplicationGroupWithOptions(request *DeleteApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetainResource) {
		query["RetainResource"] = request.RetainResource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application group. You can call this operation only for the application groups which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationGroupRequest
//
// @return DeleteApplicationGroupResponse
func (client *Client) DeleteApplicationGroup(request *DeleteApplicationGroupRequest) (_result *DeleteApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.DeleteApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple executions.
//
// @param request - DeleteExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExecutionsResponse
func (client *Client) DeleteExecutionsWithOptions(request *DeleteExecutionsRequest, runtime *dara.RuntimeOptions) (_result *DeleteExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionIds) {
		query["ExecutionIds"] = request.ExecutionIds
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExecutions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple executions.
//
// @param request - DeleteExecutionsRequest
//
// @return DeleteExecutionsResponse
func (client *Client) DeleteExecutions(request *DeleteExecutionsRequest) (_result *DeleteExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.DeleteExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes O\\\\\\&M Items.
//
// @param request - DeleteOpsItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOpsItemsResponse
func (client *Client) DeleteOpsItemsWithOptions(request *DeleteOpsItemsRequest, runtime *dara.RuntimeOptions) (_result *DeleteOpsItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpsItemIds) {
		query["OpsItemIds"] = request.OpsItemIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOpsItems"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOpsItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes O\\\\\\&M Items.
//
// @param request - DeleteOpsItemsRequest
//
// @return DeleteOpsItemsResponse
func (client *Client) DeleteOpsItems(request *DeleteOpsItemsRequest) (_result *DeleteOpsItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOpsItemsResponse{}
	_body, _err := client.DeleteOpsItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a common parameter.
//
// @param request - DeleteParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterResponse
func (client *Client) DeleteParameterWithOptions(request *DeleteParameterRequest, runtime *dara.RuntimeOptions) (_result *DeleteParameterResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a common parameter.
//
// @param request - DeleteParameterRequest
//
// @return DeleteParameterResponse
func (client *Client) DeleteParameter(request *DeleteParameterRequest) (_result *DeleteParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteParameterResponse{}
	_body, _err := client.DeleteParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a patch baseline.
//
// @param request - DeletePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePatchBaselineResponse
func (client *Client) DeletePatchBaselineWithOptions(request *DeletePatchBaselineRequest, runtime *dara.RuntimeOptions) (_result *DeletePatchBaselineResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePatchBaseline"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a patch baseline.
//
// @param request - DeletePatchBaselineRequest
//
// @return DeletePatchBaselineResponse
func (client *Client) DeletePatchBaseline(request *DeletePatchBaselineRequest) (_result *DeletePatchBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.DeletePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an encryption parameter. Make sure that you have the permissions to call the DeleteSecret operation before you call this operation.
//
// @param request - DeleteSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretParameterResponse
func (client *Client) DeleteSecretParameterWithOptions(request *DeleteSecretParameterRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretParameterResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecretParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an encryption parameter. Make sure that you have the permissions to call the DeleteSecret operation before you call this operation.
//
// @param request - DeleteSecretParameterRequest
//
// @return DeleteSecretParameterResponse
func (client *Client) DeleteSecretParameter(request *DeleteSecretParameterRequest) (_result *DeleteSecretParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.DeleteSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple desired-state configurations at a time.
//
// @param request - DeleteStateConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStateConfigurationsResponse
func (client *Client) DeleteStateConfigurationsWithOptions(request *DeleteStateConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteStateConfigurationsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StateConfigurationIds) {
		query["StateConfigurationIds"] = request.StateConfigurationIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStateConfigurations"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple desired-state configurations at a time.
//
// @param request - DeleteStateConfigurationsRequest
//
// @return DeleteStateConfigurationsResponse
func (client *Client) DeleteStateConfigurations(request *DeleteStateConfigurationsRequest) (_result *DeleteStateConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.DeleteStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoDeleteExecutions) {
		query["AutoDeleteExecutions"] = request.AutoDeleteExecutions
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple templates.
//
// @param request - DeleteTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplatesResponse
func (client *Client) DeleteTemplatesWithOptions(request *DeleteTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoDeleteExecutions) {
		query["AutoDeleteExecutions"] = request.AutoDeleteExecutions
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateNames) {
		query["TemplateNames"] = request.TemplateNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplates"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple templates.
//
// @param request - DeleteTemplatesRequest
//
// @return DeleteTemplatesResponse
func (client *Client) DeleteTemplates(request *DeleteTemplatesRequest) (_result *DeleteTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.DeleteTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys an application group. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeployApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationGroupResponse
func (client *Client) DeployApplicationGroupWithOptions(request *DeployApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *DeployApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.DeployParameters) {
		query["DeployParameters"] = request.DeployParameters
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RevisionId) {
		query["RevisionId"] = request.RevisionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an application group. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeployApplicationGroupRequest
//
// @return DeployApplicationGroupResponse
func (client *Client) DeployApplicationGroup(request *DeployApplicationGroupRequest) (_result *DeployApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployApplicationGroupResponse{}
	_body, _err := client.DeployApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用分组资源成本
//
// @param request - DescribeApplicationGroupBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationGroupBillResponse
func (client *Client) DescribeApplicationGroupBillWithOptions(request *DescribeApplicationGroupBillRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationGroupBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationGroupBill"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationGroupBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用分组资源成本
//
// @param request - DescribeApplicationGroupBillRequest
//
// @return DescribeApplicationGroupBillResponse
func (client *Client) DescribeApplicationGroupBill(request *DescribeApplicationGroupBillRequest) (_result *DescribeApplicationGroupBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApplicationGroupBillResponse{}
	_body, _err := client.DescribeApplicationGroupBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) policy required for template execution.
//
// @param request - GenerateExecutionPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateExecutionPolicyResponse
func (client *Client) GenerateExecutionPolicyWithOptions(request *GenerateExecutionPolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateExecutionPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RamRole) {
		query["RamRole"] = request.RamRole
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateExecutionPolicy"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) policy required for template execution.
//
// @param request - GenerateExecutionPolicyRequest
//
// @return GenerateExecutionPolicyResponse
func (client *Client) GenerateExecutionPolicy(request *GenerateExecutionPolicyRequest) (_result *GenerateExecutionPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.GenerateExecutionPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an O\\\\\\&M item.
//
// @param request - GenerateOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOpsItemResponse
func (client *Client) GenerateOpsItemWithOptions(request *GenerateOpsItemRequest, runtime *dara.RuntimeOptions) (_result *GenerateOpsItemResponse, _err error) {
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

	if !dara.IsNil(request.ConfigurationId) {
		query["ConfigurationId"] = request.ConfigurationId
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.DataSource) {
		query["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateOpsItem"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an O\\\\\\&M item.
//
// @param request - GenerateOpsItemRequest
//
// @return GenerateOpsItemResponse
func (client *Client) GenerateOpsItem(request *GenerateOpsItemRequest) (_result *GenerateOpsItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateOpsItemResponse{}
	_body, _err := client.GenerateOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationGroupResponse
func (client *Client) GetApplicationGroupWithOptions(request *GetApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationGroupRequest
//
// @return GetApplicationGroupResponse
func (client *Client) GetApplicationGroup(request *GetApplicationGroupRequest) (_result *GetApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.GetApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the template of an execution, including the content of the template.
//
// @param request - GetExecutionTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutionTemplateResponse
func (client *Client) GetExecutionTemplateWithOptions(request *GetExecutionTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetExecutionTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecutionTemplate"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the template of an execution, including the content of the template.
//
// @param request - GetExecutionTemplateRequest
//
// @return GetExecutionTemplateResponse
func (client *Client) GetExecutionTemplate(request *GetExecutionTemplateRequest) (_result *GetExecutionTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.GetExecutionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Git分支详情
//
// @param request - GetGitBranchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGitBranchResponse
func (client *Client) GetGitBranchWithOptions(request *GetGitBranchRequest, runtime *dara.RuntimeOptions) (_result *GetGitBranchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Branch) {
		query["Branch"] = request.Branch
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGitBranch"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGitBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Git分支详情
//
// @param request - GetGitBranchRequest
//
// @return GetGitBranchResponse
func (client *Client) GetGitBranch(request *GetGitBranchRequest) (_result *GetGitBranchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGitBranchResponse{}
	_body, _err := client.GetGitBranchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Git仓库详情
//
// @param request - GetGitRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGitRepositoryResponse
func (client *Client) GetGitRepositoryWithOptions(request *GetGitRepositoryRequest, runtime *dara.RuntimeOptions) (_result *GetGitRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGitRepository"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGitRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Git仓库详情
//
// @param request - GetGitRepositoryRequest
//
// @return GetGitRepositoryResponse
func (client *Client) GetGitRepository(request *GetGitRepositoryRequest) (_result *GetGitRepositoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGitRepositoryResponse{}
	_body, _err := client.GetGitRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the properties of a configuration list.
//
// @param request - GetInventorySchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInventorySchemaResponse
func (client *Client) GetInventorySchemaWithOptions(request *GetInventorySchemaRequest, runtime *dara.RuntimeOptions) (_result *GetInventorySchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aggregator) {
		query["Aggregator"] = request.Aggregator
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TypeName) {
		query["TypeName"] = request.TypeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInventorySchema"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the properties of a configuration list.
//
// @param request - GetInventorySchemaRequest
//
// @return GetInventorySchemaResponse
func (client *Client) GetInventorySchema(request *GetInventorySchemaRequest) (_result *GetInventorySchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.GetInventorySchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an O\\\\\\\\\\\\&M item.
//
// @param request - GetOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpsItemResponse
func (client *Client) GetOpsItemWithOptions(request *GetOpsItemRequest, runtime *dara.RuntimeOptions) (_result *GetOpsItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpsItemId) {
		query["OpsItemId"] = request.OpsItemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpsItem"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an O\\\\\\\\\\\\&M item.
//
// @param request - GetOpsItemRequest
//
// @return GetOpsItemResponse
func (client *Client) GetOpsItem(request *GetOpsItemRequest) (_result *GetOpsItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpsItemResponse{}
	_body, _err := client.GetOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a common parameter and its value.
//
// @param request - GetParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParameterResponse
func (client *Client) GetParameterWithOptions(request *GetParameterRequest, runtime *dara.RuntimeOptions) (_result *GetParameterResponse, _err error) {
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

	if !dara.IsNil(request.ParameterVersion) {
		query["ParameterVersion"] = request.ParameterVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a common parameter and its value.
//
// @param request - GetParameterRequest
//
// @return GetParameterResponse
func (client *Client) GetParameter(request *GetParameterRequest) (_result *GetParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetParameterResponse{}
	_body, _err := client.GetParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more parameters.
//
// @param request - GetParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParametersResponse
func (client *Client) GetParametersWithOptions(request *GetParametersRequest, runtime *dara.RuntimeOptions) (_result *GetParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParameters"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more parameters.
//
// @param request - GetParametersRequest
//
// @return GetParametersResponse
func (client *Client) GetParameters(request *GetParametersRequest) (_result *GetParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetParametersResponse{}
	_body, _err := client.GetParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more parameters by path.
//
// @param request - GetParametersByPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParametersByPathResponse
func (client *Client) GetParametersByPathWithOptions(request *GetParametersByPathRequest, runtime *dara.RuntimeOptions) (_result *GetParametersByPathResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Recursive) {
		query["Recursive"] = request.Recursive
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParametersByPath"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more parameters by path.
//
// @param request - GetParametersByPathRequest
//
// @return GetParametersByPathResponse
func (client *Client) GetParametersByPath(request *GetParametersByPathRequest) (_result *GetParametersByPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.GetParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a patch baseline.
//
// @param request - GetPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPatchBaselineResponse
func (client *Client) GetPatchBaselineWithOptions(request *GetPatchBaselineRequest, runtime *dara.RuntimeOptions) (_result *GetPatchBaselineResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPatchBaseline"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a patch baseline.
//
// @param request - GetPatchBaselineRequest
//
// @return GetPatchBaselineResponse
func (client *Client) GetPatchBaseline(request *GetPatchBaselineRequest) (_result *GetPatchBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.GetPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an encryption parameter, including the parameter value. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParameterResponse
func (client *Client) GetSecretParameterWithOptions(request *GetSecretParameterRequest, runtime *dara.RuntimeOptions) (_result *GetSecretParameterResponse, _err error) {
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

	if !dara.IsNil(request.ParameterVersion) {
		query["ParameterVersion"] = request.ParameterVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WithDecryption) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an encryption parameter, including the parameter value. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParameterRequest
//
// @return GetSecretParameterResponse
func (client *Client) GetSecretParameter(request *GetSecretParameterRequest) (_result *GetSecretParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.GetSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more encryption parameters. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParametersResponse
func (client *Client) GetSecretParametersWithOptions(request *GetSecretParametersRequest, runtime *dara.RuntimeOptions) (_result *GetSecretParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WithDecryption) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretParameters"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more encryption parameters. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersRequest
//
// @return GetSecretParametersResponse
func (client *Client) GetSecretParameters(request *GetSecretParametersRequest) (_result *GetSecretParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.GetSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries encryption parameters by path. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersByPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParametersByPathResponse
func (client *Client) GetSecretParametersByPathWithOptions(request *GetSecretParametersByPathRequest, runtime *dara.RuntimeOptions) (_result *GetSecretParametersByPathResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Recursive) {
		query["Recursive"] = request.Recursive
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WithDecryption) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretParametersByPath"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries encryption parameters by path. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersByPathRequest
//
// @return GetSecretParametersByPathResponse
func (client *Client) GetSecretParametersByPath(request *GetSecretParametersByPathRequest) (_result *GetSecretParametersByPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.GetSecretParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of the delivery feature.
//
// @param request - GetServiceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceSettingsResponse
func (client *Client) GetServiceSettingsWithOptions(request *GetServiceSettingsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceSettingsResponse, _err error) {
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
		Action:      dara.String("GetServiceSettings"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the delivery feature.
//
// @param request - GetServiceSettingsRequest
//
// @return GetServiceSettingsResponse
func (client *Client) GetServiceSettings(request *GetServiceSettingsRequest) (_result *GetServiceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.GetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a template, including the content of the template.
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a template, including the content of the template.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取参数可用值
//
// @param request - GetTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateParameterConstraintsResponse
func (client *Client) GetTemplateParameterConstraintsWithOptions(request *GetTemplateParameterConstraintsRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateParameterConstraints"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取参数可用值
//
// @param request - GetTemplateParameterConstraintsRequest
//
// @return GetTemplateParameterConstraintsResponse
func (client *Client) GetTemplateParameterConstraints(request *GetTemplateParameterConstraintsRequest) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.GetTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available actions, including atomic actions and cloud product actions.
//
// @param request - ListActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionsResponse
func (client *Client) ListActionsWithOptions(request *ListActionsRequest, runtime *dara.RuntimeOptions) (_result *ListActionsResponse, _err error) {
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

	if !dara.IsNil(request.OOSActionName) {
		query["OOSActionName"] = request.OOSActionName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListActions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available actions, including atomic actions and cloud product actions.
//
// @param request - ListActionsRequest
//
// @return ListActionsResponse
func (client *Client) ListActions(request *ListActionsRequest) (_result *ListActionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListActionsResponse{}
	_body, _err := client.ListActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of application groups. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationGroupsResponse
func (client *Client) ListApplicationGroupsWithOptions(request *ListApplicationGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.DeployRegionId) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationGroups"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of application groups. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationGroupsRequest
//
// @return ListApplicationGroupsResponse
func (client *Client) ListApplicationGroups(request *ListApplicationGroupsRequest) (_result *ListApplicationGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.ListApplicationGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of applications. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(tmpReq *ListApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListApplicationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationType) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of applications. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationsRequest
//
// @return ListApplicationsResponse
func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of an execution.
//
// Description:
//
// ***
//
// @param request - ListExecutionLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionLogsResponse
func (client *Client) ListExecutionLogsWithOptions(request *ListExecutionLogsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutionLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskExecutionId) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutionLogs"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of an execution.
//
// Description:
//
// ***
//
// @param request - ListExecutionLogsRequest
//
// @return ListExecutionLogsResponse
func (client *Client) ListExecutionLogs(request *ListExecutionLogsRequest) (_result *ListExecutionLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.ListExecutionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries high-risk tasks in the execution of a template.
//
// @param request - ListExecutionRiskyTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionRiskyTasksResponse
func (client *Client) ListExecutionRiskyTasksWithOptions(request *ListExecutionRiskyTasksRequest, runtime *dara.RuntimeOptions) (_result *ListExecutionRiskyTasksResponse, _err error) {
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutionRiskyTasks"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries high-risk tasks in the execution of a template.
//
// @param request - ListExecutionRiskyTasksRequest
//
// @return ListExecutionRiskyTasksResponse
func (client *Client) ListExecutionRiskyTasks(request *ListExecutionRiskyTasksRequest) (_result *ListExecutionRiskyTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.ListExecutionRiskyTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries executions. Multiple methods are supported to filter executions.
//
// @param tmpReq - ListExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionsResponse
func (client *Client) ListExecutionsWithOptions(tmpReq *ListExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListExecutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Categories) {
		query["Categories"] = request.Categories
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Depth) {
		query["Depth"] = request.Depth
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndDateAfter) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !dara.IsNil(request.EndDateBefore) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !dara.IsNil(request.ExecutedBy) {
		query["ExecutedBy"] = request.ExecutedBy
	}

	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.IncludeChildExecution) {
		query["IncludeChildExecution"] = request.IncludeChildExecution
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentExecutionId) {
		query["ParentExecutionId"] = request.ParentExecutionId
	}

	if !dara.IsNil(request.RamRole) {
		query["RamRole"] = request.RamRole
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceTemplateName) {
		query["ResourceTemplateName"] = request.ResourceTemplateName
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StartDateAfter) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !dara.IsNil(request.StartDateBefore) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExecutions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries executions. Multiple methods are supported to filter executions.
//
// @param request - ListExecutionsRequest
//
// @return ListExecutionsResponse
func (client *Client) ListExecutions(request *ListExecutionsRequest) (_result *ListExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExecutionsResponse{}
	_body, _err := client.ListExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前阿里云用户已授权在应用管理的托管平台账户
//
// @param request - ListGitAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitAccountsResponse
func (client *Client) ListGitAccountsWithOptions(request *ListGitAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListGitAccountsResponse, _err error) {
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

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGitAccounts"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGitAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前阿里云用户已授权在应用管理的托管平台账户
//
// @param request - ListGitAccountsRequest
//
// @return ListGitAccountsResponse
func (client *Client) ListGitAccounts(request *ListGitAccountsRequest) (_result *ListGitAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGitAccountsResponse{}
	_body, _err := client.ListGitAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定Git仓库的所有分支
//
// @param request - ListGitBranchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitBranchesResponse
func (client *Client) ListGitBranchesWithOptions(request *ListGitBranchesRequest, runtime *dara.RuntimeOptions) (_result *ListGitBranchesResponse, _err error) {
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

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGitBranches"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGitBranchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定Git仓库的所有分支
//
// @param request - ListGitBranchesRequest
//
// @return ListGitBranchesResponse
func (client *Client) ListGitBranches(request *ListGitBranchesRequest) (_result *ListGitBranchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGitBranchesResponse{}
	_body, _err := client.ListGitBranchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取已授权用户的组织
//
// @param request - ListGitOrganizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitOrganizationsResponse
func (client *Client) ListGitOrganizationsWithOptions(request *ListGitOrganizationsRequest, runtime *dara.RuntimeOptions) (_result *ListGitOrganizationsResponse, _err error) {
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

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGitOrganizations"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGitOrganizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取已授权用户的组织
//
// @param request - ListGitOrganizationsRequest
//
// @return ListGitOrganizationsResponse
func (client *Client) ListGitOrganizations(request *ListGitOrganizationsRequest) (_result *ListGitOrganizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGitOrganizationsResponse{}
	_body, _err := client.ListGitOrganizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Git仓库列表
//
// @param request - ListGitRepositoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitRepositoriesResponse
func (client *Client) ListGitRepositoriesWithOptions(request *ListGitRepositoriesRequest, runtime *dara.RuntimeOptions) (_result *ListGitRepositoriesResponse, _err error) {
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

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.OrgName) {
		query["OrgName"] = request.OrgName
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

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGitRepositories"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGitRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Git仓库列表
//
// @param request - ListGitRepositoriesRequest
//
// @return ListGitRepositoriesResponse
func (client *Client) ListGitRepositories(request *ListGitRepositoriesRequest) (_result *ListGitRepositoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGitRepositoriesResponse{}
	_body, _err := client.ListGitRepositoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取仓库文件与目录信息
//
// @param request - ListGitRepositoryContentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitRepositoryContentsResponse
func (client *Client) ListGitRepositoryContentsWithOptions(request *ListGitRepositoryContentsRequest, runtime *dara.RuntimeOptions) (_result *ListGitRepositoryContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Branch) {
		query["Branch"] = request.Branch
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoFullName) {
		query["RepoFullName"] = request.RepoFullName
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGitRepositoryContents"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGitRepositoryContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取仓库文件与目录信息
//
// @param request - ListGitRepositoryContentsRequest
//
// @return ListGitRepositoryContentsResponse
func (client *Client) ListGitRepositoryContents(request *ListGitRepositoryContentsRequest) (_result *ListGitRepositoryContentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGitRepositoryContentsResponse{}
	_body, _err := client.ListGitRepositoryContentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Instance Package States
//
// @param request - ListInstancePackageStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePackageStatesResponse
func (client *Client) ListInstancePackageStatesWithOptions(request *ListInstancePackageStatesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancePackageStatesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateNames) {
		query["TemplateNames"] = request.TemplateNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancePackageStates"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancePackageStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Instance Package States
//
// @param request - ListInstancePackageStatesRequest
//
// @return ListInstancePackageStatesResponse
func (client *Client) ListInstancePackageStates(request *ListInstancePackageStatesRequest) (_result *ListInstancePackageStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancePackageStatesResponse{}
	_body, _err := client.ListInstancePackageStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePatchStatesResponse
func (client *Client) ListInstancePatchStatesWithOptions(request *ListInstancePatchStatesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancePatchStatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancePatchStates"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchStatesRequest
//
// @return ListInstancePatchStatesResponse
func (client *Client) ListInstancePatchStates(request *ListInstancePatchStatesRequest) (_result *ListInstancePatchStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.ListInstancePatchStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePatchesResponse
func (client *Client) ListInstancePatchesWithOptions(request *ListInstancePatchesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancePatchesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PatchStatuses) {
		query["PatchStatuses"] = request.PatchStatuses
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancePatches"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchesRequest
//
// @return ListInstancePatchesResponse
func (client *Client) ListInstancePatches(request *ListInstancePatchesRequest) (_result *ListInstancePatchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.ListInstancePatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an Elastic Compute Service (ECS) instance.
//
// @param request - ListInventoryEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInventoryEntriesResponse
func (client *Client) ListInventoryEntriesWithOptions(request *ListInventoryEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListInventoryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TypeName) {
		query["TypeName"] = request.TypeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInventoryEntries"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an Elastic Compute Service (ECS) instance.
//
// @param request - ListInventoryEntriesRequest
//
// @return ListInventoryEntriesResponse
func (client *Client) ListInventoryEntries(request *ListInventoryEntriesRequest) (_result *ListInventoryEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.ListInventoryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries O\\&M items.
//
// @param tmpReq - ListOpsItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpsItemsResponse
func (client *Client) ListOpsItemsWithOptions(tmpReq *ListOpsItemsRequest, runtime *dara.RuntimeOptions) (_result *ListOpsItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListOpsItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceTags) {
		request.ResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTags, dara.String("ResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceTagsShrink) {
		query["ResourceTags"] = request.ResourceTagsShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOpsItems"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpsItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries O\\&M items.
//
// @param request - ListOpsItemsRequest
//
// @return ListOpsItemsResponse
func (client *Client) ListOpsItems(request *ListOpsItemsRequest) (_result *ListOpsItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOpsItemsResponse{}
	_body, _err := client.ListOpsItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a common parameter.
//
// @param request - ListParameterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterVersionsResponse
func (client *Client) ListParameterVersionsWithOptions(request *ListParameterVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListParameterVersionsResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParameterVersions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a common parameter.
//
// @param request - ListParameterVersionsRequest
//
// @return ListParameterVersionsResponse
func (client *Client) ListParameterVersions(request *ListParameterVersionsRequest) (_result *ListParameterVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.ListParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple methods are supported to filter common parameters.
//
// @param tmpReq - ListParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParametersResponse
func (client *Client) ListParametersWithOptions(tmpReq *ListParametersRequest, runtime *dara.RuntimeOptions) (_result *ListParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
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

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Recursive) {
		query["Recursive"] = request.Recursive
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParameters"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple methods are supported to filter common parameters.
//
// @param request - ListParametersRequest
//
// @return ListParametersResponse
func (client *Client) ListParameters(request *ListParametersRequest) (_result *ListParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListParametersResponse{}
	_body, _err := client.ListParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of patch baselines.
//
// @param tmpReq - ListPatchBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPatchBaselinesResponse
func (client *Client) ListPatchBaselinesWithOptions(tmpReq *ListPatchBaselinesRequest, runtime *dara.RuntimeOptions) (_result *ListPatchBaselinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPatchBaselinesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovedPatches) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, dara.String("ApprovedPatches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovedPatchesShrink) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !dara.IsNil(request.ApprovedPatchesEnableNonSecurity) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OperationSystem) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPatchBaselines"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of patch baselines.
//
// @param request - ListPatchBaselinesRequest
//
// @return ListPatchBaselinesResponse
func (client *Client) ListPatchBaselines(request *ListPatchBaselinesRequest) (_result *ListPatchBaselinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.ListPatchBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduled execution that involves O&M operations on Elastic Compute Service (ECS) instances.
//
// @param request - ListResourceExecutionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExecutionStatusResponse
func (client *Client) ListResourceExecutionStatusWithOptions(request *ListResourceExecutionStatusRequest, runtime *dara.RuntimeOptions) (_result *ListResourceExecutionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceExecutionStatus"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduled execution that involves O&M operations on Elastic Compute Service (ECS) instances.
//
// @param request - ListResourceExecutionStatusRequest
//
// @return ListResourceExecutionStatusResponse
func (client *Client) ListResourceExecutionStatus(request *ListResourceExecutionStatusRequest) (_result *ListResourceExecutionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.ListResourceExecutionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries versions of an encryption parameter.
//
// @param request - ListSecretParameterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretParameterVersionsResponse
func (client *Client) ListSecretParameterVersionsWithOptions(request *ListSecretParameterVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListSecretParameterVersionsResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.WithDecryption) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecretParameterVersions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of an encryption parameter.
//
// @param request - ListSecretParameterVersionsRequest
//
// @return ListSecretParameterVersionsResponse
func (client *Client) ListSecretParameterVersions(request *ListSecretParameterVersionsRequest) (_result *ListSecretParameterVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.ListSecretParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple types of queries are supported.
//
// Description:
//
// Before you call this operation, make sure that you have the permission to manage Key Management Service (KMS) secrets.
//
// @param tmpReq - ListSecretParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretParametersResponse
func (client *Client) ListSecretParametersWithOptions(tmpReq *ListSecretParametersRequest, runtime *dara.RuntimeOptions) (_result *ListSecretParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSecretParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
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

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Recursive) {
		query["Recursive"] = request.Recursive
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecretParameters"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple types of queries are supported.
//
// Description:
//
// Before you call this operation, make sure that you have the permission to manage Key Management Service (KMS) secrets.
//
// @param request - ListSecretParametersRequest
//
// @return ListSecretParametersResponse
func (client *Client) ListSecretParameters(request *ListSecretParametersRequest) (_result *ListSecretParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.ListSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries desired-state configurations.
//
// @param tmpReq - ListStateConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStateConfigurationsResponse
func (client *Client) ListStateConfigurationsWithOptions(tmpReq *ListStateConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListStateConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListStateConfigurationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StateConfigurationIds) {
		query["StateConfigurationIds"] = request.StateConfigurationIds
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStateConfigurations"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries desired-state configurations.
//
// @param request - ListStateConfigurationsRequest
//
// @return ListStateConfigurationsResponse
func (client *Client) ListStateConfigurations(request *ListStateConfigurationsRequest) (_result *ListStateConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.ListStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-06-01"),
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
// Queries the tags that are added to one or more resources.
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
// Queries the values of created tags.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the values of created tags.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries task executions. Multiple methods are supported to filter task executions.
//
// @param request - ListTaskExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskExecutionsResponse
func (client *Client) ListTaskExecutionsWithOptions(request *ListTaskExecutionsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDateAfter) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !dara.IsNil(request.EndDateBefore) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.IncludeChildTaskExecution) {
		query["IncludeChildTaskExecution"] = request.IncludeChildTaskExecution
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentTaskExecutionId) {
		query["ParentTaskExecutionId"] = request.ParentTaskExecutionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StartDateAfter) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !dara.IsNil(request.StartDateBefore) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	if !dara.IsNil(request.TaskExecutionId) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskExecutions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries task executions. Multiple methods are supported to filter task executions.
//
// @param request - ListTaskExecutionsRequest
//
// @return ListTaskExecutionsResponse
func (client *Client) ListTaskExecutions(request *ListTaskExecutionsRequest) (_result *ListTaskExecutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.ListTaskExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplateVersions"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries templates. Multiple methods are supported to filter templates.
//
// @param tmpReq - ListTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(tmpReq *ListTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.CreatedBy) {
		query["CreatedBy"] = request.CreatedBy
	}

	if !dara.IsNil(request.CreatedDateAfter) {
		query["CreatedDateAfter"] = request.CreatedDateAfter
	}

	if !dara.IsNil(request.CreatedDateBefore) {
		query["CreatedDateBefore"] = request.CreatedDateBefore
	}

	if !dara.IsNil(request.HasTrigger) {
		query["HasTrigger"] = request.HasTrigger
	}

	if !dara.IsNil(request.IsExample) {
		query["IsExample"] = request.IsExample
	}

	if !dara.IsNil(request.IsFavorite) {
		query["IsFavorite"] = request.IsFavorite
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateFormat) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries templates. Multiple methods are supported to filter templates.
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Notifies an execution in the Waiting state of the subsequent operations.
//
// Description:
//
// You can call this operation to notify an execution in the following scenarios:
//
//   - If a template contains a special task, such as an approval task, the Operation Orchestration Service (OOS) execution engine sets the execution state to Waiting when the approval task is being run. You can call this operation to specify whether to continue the execution.
//
//   - If you perform debugging in the debug mode, you can call this operation to notify the execution of the subsequent operations after the execution is created or a task is complete.
//
//   - If a high-risk operation task waits for approval, you can call this operation to specify whether to continue the execution.
//
// @param request - NotifyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyExecutionResponse
func (client *Client) NotifyExecutionWithOptions(request *NotifyExecutionRequest, runtime *dara.RuntimeOptions) (_result *NotifyExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.ExecutionStatus) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !dara.IsNil(request.LoopItem) {
		query["LoopItem"] = request.LoopItem
	}

	if !dara.IsNil(request.NotifyNote) {
		query["NotifyNote"] = request.NotifyNote
	}

	if !dara.IsNil(request.NotifyType) {
		query["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskExecutionId) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	if !dara.IsNil(request.TaskExecutionIds) {
		query["TaskExecutionIds"] = request.TaskExecutionIds
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NotifyExecution"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Notifies an execution in the Waiting state of the subsequent operations.
//
// Description:
//
// You can call this operation to notify an execution in the following scenarios:
//
//   - If a template contains a special task, such as an approval task, the Operation Orchestration Service (OOS) execution engine sets the execution state to Waiting when the approval task is being run. You can call this operation to specify whether to continue the execution.
//
//   - If you perform debugging in the debug mode, you can call this operation to notify the execution of the subsequent operations after the execution is created or a task is complete.
//
//   - If a high-risk operation task waits for approval, you can call this operation to specify whether to continue the execution.
//
// @param request - NotifyExecutionRequest
//
// @return NotifyExecutionResponse
func (client *Client) NotifyExecution(request *NotifyExecutionRequest) (_result *NotifyExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.NotifyExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers the default patch baseline.
//
// @param request - RegisterDefaultPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDefaultPatchBaselineResponse
func (client *Client) RegisterDefaultPatchBaselineWithOptions(request *RegisterDefaultPatchBaselineRequest, runtime *dara.RuntimeOptions) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDefaultPatchBaseline"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers the default patch baseline.
//
// @param request - RegisterDefaultPatchBaselineRequest
//
// @return RegisterDefaultPatchBaselineResponse
func (client *Client) RegisterDefaultPatchBaseline(request *RegisterDefaultPatchBaselineRequest) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.RegisterDefaultPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details or aggregate information of a configuration inventory.
//
// @param request - SearchInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInventoryResponse
func (client *Client) SearchInventoryWithOptions(request *SearchInventoryRequest, runtime *dara.RuntimeOptions) (_result *SearchInventoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aggregator) {
		query["Aggregator"] = request.Aggregator
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchInventory"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details or aggregate information of a configuration inventory.
//
// @param request - SearchInventoryRequest
//
// @return SearchInventoryResponse
func (client *Client) SearchInventory(request *SearchInventoryRequest) (_result *SearchInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchInventoryResponse{}
	_body, _err := client.SearchInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the feature of delivering template execution records and sets the storage location.
//
// @param request - SetServiceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetServiceSettingsResponse
func (client *Client) SetServiceSettingsWithOptions(request *SetServiceSettingsRequest, runtime *dara.RuntimeOptions) (_result *SetServiceSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryOssBucketName) {
		query["DeliveryOssBucketName"] = request.DeliveryOssBucketName
	}

	if !dara.IsNil(request.DeliveryOssEnabled) {
		query["DeliveryOssEnabled"] = request.DeliveryOssEnabled
	}

	if !dara.IsNil(request.DeliveryOssKeyPrefix) {
		query["DeliveryOssKeyPrefix"] = request.DeliveryOssKeyPrefix
	}

	if !dara.IsNil(request.DeliverySlsEnabled) {
		query["DeliverySlsEnabled"] = request.DeliverySlsEnabled
	}

	if !dara.IsNil(request.DeliverySlsProjectName) {
		query["DeliverySlsProjectName"] = request.DeliverySlsProjectName
	}

	if !dara.IsNil(request.RdcEnterpriseId) {
		query["RdcEnterpriseId"] = request.RdcEnterpriseId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetServiceSettings"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the feature of delivering template execution records and sets the storage location.
//
// @param request - SetServiceSettingsRequest
//
// @return SetServiceSettingsResponse
func (client *Client) SetServiceSettings(request *SetServiceSettingsRequest) (_result *SetServiceSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.SetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an execution.
//
// @param tmpReq - StartExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartExecutionResponse
func (client *Client) StartExecutionWithOptions(tmpReq *StartExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.LoopMode) {
		query["LoopMode"] = request.LoopMode
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.ParentExecutionId) {
		query["ParentExecutionId"] = request.ParentExecutionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SafetyCheck) {
		query["SafetyCheck"] = request.SafetyCheck
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartExecution"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an execution.
//
// @param request - StartExecutionRequest
//
// @return StartExecutionResponse
func (client *Client) StartExecution(request *StartExecutionRequest) (_result *StartExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartExecutionResponse{}
	_body, _err := client.StartExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to one or more resources.
//
// @param tmpReq - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-06-01"),
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
// Adds tags to one or more resources.
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
// Debugs a started execution that contains an event trigger task or alert trigger task. If the operation is called, a message body is sent to the event trigger task or alert trigger task. After the trigger task receives the message body, the trigger task generates a new child execution.
//
// @param request - TriggerExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerExecutionResponse
func (client *Client) TriggerExecutionWithOptions(request *TriggerExecutionRequest, runtime *dara.RuntimeOptions) (_result *TriggerExecutionResponse, _err error) {
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

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerExecution"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a started execution that contains an event trigger task or alert trigger task. If the operation is called, a message body is sent to the event trigger task or alert trigger task. After the trigger task receives the message body, the trigger task generates a new child execution.
//
// @param request - TriggerExecutionRequest
//
// @return TriggerExecutionResponse
func (client *Client) TriggerExecution(request *TriggerExecutionRequest) (_result *TriggerExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.TriggerExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from one or more resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("TagKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeysShrink) {
		query["TagKeys"] = request.TagKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-06-01"),
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
// Removes tags from one or more resources.
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
// Updates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - UpdateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithOptions(tmpReq *UpdateApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlarmConfig) {
		request.AlarmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlarmConfig, dara.String("AlarmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmConfigShrink) {
		query["AlarmConfig"] = request.AlarmConfigShrink
	}

	if !dara.IsNil(request.DeleteAlarmRulesBeforeUpdate) {
		query["DeleteAlarmRulesBeforeUpdate"] = request.DeleteAlarmRulesBeforeUpdate
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplication"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - UpdateApplicationRequest
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - UpdateApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationGroupResponse
func (client *Client) UpdateApplicationGroupWithOptions(tmpReq *UpdateApplicationGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApplicationGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.DeployedRevisionId) {
		query["DeployedRevisionId"] = request.DeployedRevisionId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewName) {
		query["NewName"] = request.NewName
	}

	if !dara.IsNil(request.OperationName) {
		query["OperationName"] = request.OperationName
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationGroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - UpdateApplicationGroupRequest
//
// @return UpdateApplicationGroupResponse
func (client *Client) UpdateApplicationGroup(request *UpdateApplicationGroupRequest) (_result *UpdateApplicationGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.UpdateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update executions that are in Running or Waiting status.
//
// @param request - UpdateExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExecutionResponse
func (client *Client) UpdateExecutionWithOptions(request *UpdateExecutionRequest, runtime *dara.RuntimeOptions) (_result *UpdateExecutionResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionId) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExecution"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update executions that are in Running or Waiting status.
//
// @param request - UpdateExecutionRequest
//
// @return UpdateExecutionResponse
func (client *Client) UpdateExecution(request *UpdateExecutionRequest) (_result *UpdateExecutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.UpdateExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the installed extensions of an instance.
//
// @param tmpReq - UpdateInstancePackageStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstancePackageStateResponse
func (client *Client) UpdateInstancePackageStateWithOptions(tmpReq *UpdateInstancePackageStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstancePackageStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateInstancePackageStateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationInfo) {
		query["ConfigurationInfo"] = request.ConfigurationInfo
	}

	if !dara.IsNil(request.ConfigureAction) {
		query["ConfigureAction"] = request.ConfigureAction
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstancePackageState"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstancePackageStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the installed extensions of an instance.
//
// @param request - UpdateInstancePackageStateRequest
//
// @return UpdateInstancePackageStateResponse
func (client *Client) UpdateInstancePackageState(request *UpdateInstancePackageStateRequest) (_result *UpdateInstancePackageStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstancePackageStateResponse{}
	_body, _err := client.UpdateInstancePackageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an O\\\\\\\\\\\\&M item.
//
// @param tmpReq - UpdateOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOpsItemResponse
func (client *Client) UpdateOpsItemWithOptions(tmpReq *UpdateOpsItemRequest, runtime *dara.RuntimeOptions) (_result *UpdateOpsItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateOpsItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DedupString) {
		query["DedupString"] = request.DedupString
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OpsItemId) {
		query["OpsItemId"] = request.OpsItemId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Solutions) {
		query["Solutions"] = request.Solutions
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOpsItem"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an O\\\\\\\\\\\\&M item.
//
// @param request - UpdateOpsItemRequest
//
// @return UpdateOpsItemResponse
func (client *Client) UpdateOpsItem(request *UpdateOpsItemRequest) (_result *UpdateOpsItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOpsItemResponse{}
	_body, _err := client.UpdateOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a common parameter.
//
// @param request - UpdateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParameterResponse
func (client *Client) UpdateParameterWithOptions(request *UpdateParameterRequest, runtime *dara.RuntimeOptions) (_result *UpdateParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a common parameter.
//
// @param request - UpdateParameterRequest
//
// @return UpdateParameterResponse
func (client *Client) UpdateParameter(request *UpdateParameterRequest) (_result *UpdateParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateParameterResponse{}
	_body, _err := client.UpdateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a patch baseline.
//
// @param tmpReq - UpdatePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePatchBaselineResponse
func (client *Client) UpdatePatchBaselineWithOptions(tmpReq *UpdatePatchBaselineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePatchBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePatchBaselineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovedPatches) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, dara.String("ApprovedPatches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RejectedPatches) {
		request.RejectedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RejectedPatches, dara.String("RejectedPatches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sources) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, dara.String("Sources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalRules) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !dara.IsNil(request.ApprovedPatchesShrink) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !dara.IsNil(request.ApprovedPatchesEnableNonSecurity) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RejectedPatchesShrink) {
		query["RejectedPatches"] = request.RejectedPatchesShrink
	}

	if !dara.IsNil(request.RejectedPatchesAction) {
		query["RejectedPatchesAction"] = request.RejectedPatchesAction
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourcesShrink) {
		query["Sources"] = request.SourcesShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePatchBaseline"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a patch baseline.
//
// @param request - UpdatePatchBaselineRequest
//
// @return UpdatePatchBaselineResponse
func (client *Client) UpdatePatchBaseline(request *UpdatePatchBaselineRequest) (_result *UpdatePatchBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.UpdatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an encryption parameter.
//
// @param tmpReq - UpdateSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretParameterResponse
func (client *Client) UpdateSecretParameterWithOptions(tmpReq *UpdateSecretParameterRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecretParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecretParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecretParameter"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an encryption parameter.
//
// @param request - UpdateSecretParameterRequest
//
// @return UpdateSecretParameterResponse
func (client *Client) UpdateSecretParameter(request *UpdateSecretParameterRequest) (_result *UpdateSecretParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.UpdateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a desired-state configuration.
//
// @param tmpReq - UpdateStateConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStateConfigurationResponse
func (client *Client) UpdateStateConfigurationWithOptions(tmpReq *UpdateStateConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateStateConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigureMode) {
		query["ConfigureMode"] = request.ConfigureMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleExpression) {
		query["ScheduleExpression"] = request.ScheduleExpression
	}

	if !dara.IsNil(request.ScheduleType) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.StateConfigurationId) {
		query["StateConfigurationId"] = request.StateConfigurationId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStateConfiguration"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a desired-state configuration.
//
// @param request - UpdateStateConfigurationRequest
//
// @return UpdateStateConfigurationResponse
func (client *Client) UpdateStateConfiguration(request *UpdateStateConfigurationRequest) (_result *UpdateStateConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.UpdateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an existing template.
//
// @param tmpReq - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(tmpReq *UpdateTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an existing template.
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check whether a template is valid.
//
// @param request - ValidateTemplateContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTemplateContentResponse
func (client *Client) ValidateTemplateContentWithOptions(request *ValidateTemplateContentRequest, runtime *dara.RuntimeOptions) (_result *ValidateTemplateContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateTemplateContent"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check whether a template is valid.
//
// @param request - ValidateTemplateContentRequest
//
// @return ValidateTemplateContentResponse
func (client *Client) ValidateTemplateContent(request *ValidateTemplateContentRequest) (_result *ValidateTemplateContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.ValidateTemplateContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
