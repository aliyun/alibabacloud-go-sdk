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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aiworkspace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Receives and processes system event messages sent by DataWorks.
//
// Description:
//
// This operation can be called only by the internal system and cannot be called by external users.
//
// @param request - AcceptDataworksEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptDataworksEventResponse
func (client *Client) AcceptDataworksEventWithOptions(request *AcceptDataworksEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AcceptDataworksEventResponse, _err error) {
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

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptDataworksEvent"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/action/acceptdataworksevent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptDataworksEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Receives and processes system event messages sent by DataWorks.
//
// Description:
//
// This operation can be called only by the internal system and cannot be called by external users.
//
// @param request - AcceptDataworksEventRequest
//
// @return AcceptDataworksEventResponse
func (client *Client) AcceptDataworksEvent(request *AcceptDataworksEventRequest) (_result *AcceptDataworksEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AcceptDataworksEventResponse{}
	_body, _err := client.AcceptDataworksEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a custom image to a workspace.
//
// @param request - AddImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithOptions(request *AddImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageUri) {
		body["ImageUri"] = request.ImageUri
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImage"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom image to a workspace.
//
// @param request - AddImageRequest
//
// @return AddImageResponse
func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to an image.
//
// @param request - AddImageLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageLabelsResponse
func (client *Client) AddImageLabelsWithOptions(ImageId *string, request *AddImageLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddImageLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImageLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images/" + dara.PercentEncode(dara.StringValue(ImageId)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to an image.
//
// @param request - AddImageLabelsRequest
//
// @return AddImageLabelsResponse
func (client *Client) AddImageLabels(ImageId *string, request *AddImageLabelsRequest) (_result *AddImageLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageLabelsResponse{}
	_body, _err := client.AddImageLabelsWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a role to a member in a workspace. After you add a role to a member, the member is granted the permissions of the role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberRoleResponse
func (client *Client) AddMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMemberRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMemberRole"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/members/" + dara.PercentEncode(dara.StringValue(MemberId)) + "/roles/" + dara.PercentEncode(dara.StringValue(RoleName))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMemberRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a role to a member in a workspace. After you add a role to a member, the member is granted the permissions of the role.
//
// @return AddMemberRoleResponse
func (client *Client) AddMemberRole(WorkspaceId *string, MemberId *string, RoleName *string) (_result *AddMemberRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMemberRoleResponse{}
	_body, _err := client.AddMemberRoleWithOptions(WorkspaceId, MemberId, RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a resource belongs based on the ID.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourcegroups/action/changeresourcegroup"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Changes the resource group to which a resource belongs based on the ID.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a code build in Platform for AI (PAI). You can configure Git branches and commit IDs. After the code build is created, you can reference the code build in a Deep Learning Containers (DLC) job.
//
// @param request - CreateCodeSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCodeSourceResponse
func (client *Client) CreateCodeSourceWithOptions(request *CreateCodeSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCodeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.CodeBranch) {
		body["CodeBranch"] = request.CodeBranch
	}

	if !dara.IsNil(request.CodeCommit) {
		body["CodeCommit"] = request.CodeCommit
	}

	if !dara.IsNil(request.CodeRepo) {
		body["CodeRepo"] = request.CodeRepo
	}

	if !dara.IsNil(request.CodeRepoAccessToken) {
		body["CodeRepoAccessToken"] = request.CodeRepoAccessToken
	}

	if !dara.IsNil(request.CodeRepoUserName) {
		body["CodeRepoUserName"] = request.CodeRepoUserName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.MountPath) {
		body["MountPath"] = request.MountPath
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCodeSource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a code build in Platform for AI (PAI). You can configure Git branches and commit IDs. After the code build is created, you can reference the code build in a Deep Learning Containers (DLC) job.
//
// @param request - CreateCodeSourceRequest
//
// @return CreateCodeSourceResponse
func (client *Client) CreateCodeSource(request *CreateCodeSourceRequest) (_result *CreateCodeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CreateCodeSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a connection. This API is used to connect Platform for AI (PAI) to customer models and databases in LangStudio and multimodal dataset search scenarios.
//
// @param request - CreateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnectionResponse
func (client *Client) CreateConnectionWithOptions(request *CreateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.ConnectionType) {
		body["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Models) {
		body["Models"] = request.Models
	}

	if !dara.IsNil(request.ResourceMeta) {
		body["ResourceMeta"] = request.ResourceMeta
	}

	if !dara.IsNil(request.Secrets) {
		body["Secrets"] = request.Secrets
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConnection"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/connections"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a connection. This API is used to connect Platform for AI (PAI) to customer models and databases in LangStudio and multimodal dataset search scenarios.
//
// @param request - CreateConnectionRequest
//
// @return CreateConnectionResponse
func (client *Client) CreateConnection(request *CreateConnectionRequest) (_result *CreateConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConnectionResponse{}
	_body, _err := client.CreateConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dataset.
//
// @param request - CreateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.DataCount) {
		body["DataCount"] = request.DataCount
	}

	if !dara.IsNil(request.DataSize) {
		body["DataSize"] = request.DataSize
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Edition) {
		body["Edition"] = request.Edition
	}

	if !dara.IsNil(request.ImportInfo) {
		body["ImportInfo"] = request.ImportInfo
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MountAccessReadWriteRoleIdList) {
		body["MountAccessReadWriteRoleIdList"] = request.MountAccessReadWriteRoleIdList
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.Property) {
		body["Property"] = request.Property
	}

	if !dara.IsNil(request.Provider) {
		body["Provider"] = request.Provider
	}

	if !dara.IsNil(request.ProviderType) {
		body["ProviderType"] = request.ProviderType
	}

	if !dara.IsNil(request.SourceDatasetId) {
		body["SourceDatasetId"] = request.SourceDatasetId
	}

	if !dara.IsNil(request.SourceDatasetVersion) {
		body["SourceDatasetVersion"] = request.SourceDatasetVersion
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.VersionDescription) {
		body["VersionDescription"] = request.VersionDescription
	}

	if !dara.IsNil(request.VersionLabels) {
		body["VersionLabels"] = request.VersionLabels
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset.
//
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates the metadata records of multiple files in a dataset at a time.
//
// @param request - CreateDatasetFileMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetFileMetasResponse
func (client *Client) CreateDatasetFileMetasWithOptions(DatasetId *string, request *CreateDatasetFileMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetFileMetasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetFileMetas) {
		body["DatasetFileMetas"] = request.DatasetFileMetas
	}

	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetFileMetas"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetFileMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates the metadata records of multiple files in a dataset at a time.
//
// @param request - CreateDatasetFileMetasRequest
//
// @return CreateDatasetFileMetasResponse
func (client *Client) CreateDatasetFileMetas(DatasetId *string, request *CreateDatasetFileMetasRequest) (_result *CreateDatasetFileMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetFileMetasResponse{}
	_body, _err := client.CreateDatasetFileMetasWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dataset job.
//
// @param request - CreateDatasetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetJobResponse
func (client *Client) CreateDatasetJobWithOptions(DatasetId *string, request *CreateDatasetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.JobAction) {
		body["JobAction"] = request.JobAction
	}

	if !dara.IsNil(request.JobMode) {
		body["JobMode"] = request.JobMode
	}

	if !dara.IsNil(request.JobSpec) {
		body["JobSpec"] = request.JobSpec
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetJob"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset job.
//
// @param request - CreateDatasetJobRequest
//
// @return CreateDatasetJobResponse
func (client *Client) CreateDatasetJob(DatasetId *string, request *CreateDatasetJobRequest) (_result *CreateDatasetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetJobResponse{}
	_body, _err := client.CreateDatasetJobWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a job configuration for a dataset.
//
// @param request - CreateDatasetJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetJobConfigResponse
func (client *Client) CreateDatasetJobConfigWithOptions(DatasetId *string, request *CreateDatasetJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.ConfigType) {
		body["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetJobConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobconfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job configuration for a dataset.
//
// @param request - CreateDatasetJobConfigRequest
//
// @return CreateDatasetJobConfigResponse
func (client *Client) CreateDatasetJobConfig(DatasetId *string, request *CreateDatasetJobConfigRequest) (_result *CreateDatasetJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetJobConfigResponse{}
	_body, _err := client.CreateDatasetJobConfigWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags for a dataset.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - The tag key and value are not empty strings and cannot exceed 128 characters in length.
//
//   - The tag key cannot start with any of the following strings: "aliyun", "acs", "http://", and "https://".
//
// @param request - CreateDatasetLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetLabelsResponse
func (client *Client) CreateDatasetLabelsWithOptions(DatasetId *string, request *CreateDatasetLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates tags for a dataset.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - The tag key and value are not empty strings and cannot exceed 128 characters in length.
//
//   - The tag key cannot start with any of the following strings: "aliyun", "acs", "http://", and "https://".
//
// @param request - CreateDatasetLabelsRequest
//
// @return CreateDatasetLabelsResponse
func (client *Client) CreateDatasetLabels(DatasetId *string, request *CreateDatasetLabelsRequest) (_result *CreateDatasetLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetLabelsResponse{}
	_body, _err := client.CreateDatasetLabelsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dataset version.
//
// @param request - CreateDatasetVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetVersionResponse
func (client *Client) CreateDatasetVersionWithOptions(DatasetId *string, request *CreateDatasetVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataCount) {
		body["DataCount"] = request.DataCount
	}

	if !dara.IsNil(request.DataSize) {
		body["DataSize"] = request.DataSize
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ImportInfo) {
		body["ImportInfo"] = request.ImportInfo
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.Property) {
		body["Property"] = request.Property
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset version.
//
// @param request - CreateDatasetVersionRequest
//
// @return CreateDatasetVersionResponse
func (client *Client) CreateDatasetVersion(DatasetId *string, request *CreateDatasetVersionRequest) (_result *CreateDatasetVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetVersionResponse{}
	_body, _err := client.CreateDatasetVersionWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags for a dataset version.
//
// @param request - CreateDatasetVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetVersionLabelsResponse
func (client *Client) CreateDatasetVersionLabelsWithOptions(DatasetId *string, VersionName *string, request *CreateDatasetVersionLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetVersionLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetVersionLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName)) + "/labels"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates tags for a dataset version.
//
// @param request - CreateDatasetVersionLabelsRequest
//
// @return CreateDatasetVersionLabelsResponse
func (client *Client) CreateDatasetVersionLabels(DatasetId *string, VersionName *string, request *CreateDatasetVersionLabelsRequest) (_result *CreateDatasetVersionLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetVersionLabelsResponse{}
	_body, _err := client.CreateDatasetVersionLabelsWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an experiment.
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.ArtifactUri) {
		body["ArtifactUri"] = request.ArtifactUri
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperiment"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateExperimentRequest
//
// @return CreateExperimentResponse
func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a workspace as a member. You can add multiple users as members.
//
// @param request - CreateMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberResponse
func (client *Client) CreateMemberWithOptions(WorkspaceId *string, request *CreateMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Members) {
		body["Members"] = request.Members
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMember"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a workspace as a member. You can add multiple users as members.
//
// @param request - CreateMemberRequest
//
// @return CreateMemberResponse
func (client *Client) CreateMember(WorkspaceId *string, request *CreateMemberRequest) (_result *CreateMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemberResponse{}
	_body, _err := client.CreateMemberWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model. A model is a collection of model versions. When you create a model, you must specify the model name and description.
//
// @param request - CreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithOptions(request *CreateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.ModelDescription) {
		body["ModelDescription"] = request.ModelDescription
	}

	if !dara.IsNil(request.ModelDoc) {
		body["ModelDoc"] = request.ModelDoc
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelType) {
		body["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.OrderNumber) {
		body["OrderNumber"] = request.OrderNumber
	}

	if !dara.IsNil(request.Origin) {
		body["Origin"] = request.Origin
	}

	if !dara.IsNil(request.ParameterSize) {
		body["ParameterSize"] = request.ParameterSize
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Task) {
		body["Task"] = request.Task
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model. A model is a collection of model versions. When you create a model, you must specify the model name and description.
//
// @param request - CreateModelRequest
//
// @return CreateModelResponse
func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tag for a model.
//
// @param request - CreateModelLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelLabelsResponse
func (client *Client) CreateModelLabelsWithOptions(ModelId *string, request *CreateModelLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag for a model.
//
// @param request - CreateModelLabelsRequest
//
// @return CreateModelLabelsResponse
func (client *Client) CreateModelLabels(ModelId *string, request *CreateModelLabelsRequest) (_result *CreateModelLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelLabelsResponse{}
	_body, _err := client.CreateModelLabelsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new version for the specified model.
//
// @param request - CreateModelVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelVersionResponse
func (client *Client) CreateModelVersionWithOptions(ModelId *string, request *CreateModelVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalStatus) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !dara.IsNil(request.CompressionSpec) {
		body["CompressionSpec"] = request.CompressionSpec
	}

	if !dara.IsNil(request.DistillationSpec) {
		body["DistillationSpec"] = request.DistillationSpec
	}

	if !dara.IsNil(request.EvaluationSpec) {
		body["EvaluationSpec"] = request.EvaluationSpec
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.FormatType) {
		body["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.FrameworkType) {
		body["FrameworkType"] = request.FrameworkType
	}

	if !dara.IsNil(request.InferenceSpec) {
		body["InferenceSpec"] = request.InferenceSpec
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Metrics) {
		body["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.TrainingSpec) {
		body["TrainingSpec"] = request.TrainingSpec
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.VersionDescription) {
		body["VersionDescription"] = request.VersionDescription
	}

	if !dara.IsNil(request.VersionName) {
		body["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new version for the specified model.
//
// @param request - CreateModelVersionRequest
//
// @return CreateModelVersionResponse
func (client *Client) CreateModelVersion(ModelId *string, request *CreateModelVersionRequest) (_result *CreateModelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelVersionResponse{}
	_body, _err := client.CreateModelVersionWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tag for a model version.
//
// @param request - CreateModelVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelVersionLabelsResponse
func (client *Client) CreateModelVersionLabelsWithOptions(ModelId *string, VersionName *string, request *CreateModelVersionLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelVersionLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelVersionLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag for a model version.
//
// @param request - CreateModelVersionLabelsRequest
//
// @return CreateModelVersionLabelsResponse
func (client *Client) CreateModelVersionLabels(ModelId *string, VersionName *string, request *CreateModelVersionLabelsRequest) (_result *CreateModelVersionLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelVersionLabelsResponse{}
	_body, _err := client.CreateModelVersionLabelsWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go order for DataWorks, OSS, PAI, or MaxCompute.
//
// @param request - CreateProductOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductOrdersResponse
func (client *Client) CreateProductOrdersWithOptions(request *CreateProductOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProductOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Products) {
		body["Products"] = request.Products
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProductOrders"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/productorders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProductOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go order for DataWorks, OSS, PAI, or MaxCompute.
//
// @param request - CreateProductOrdersRequest
//
// @return CreateProductOrdersResponse
func (client *Client) CreateProductOrders(request *CreateProductOrdersRequest) (_result *CreateProductOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductOrdersResponse{}
	_body, _err := client.CreateProductOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a run. A run is an experiment that can be associated with a specific workload or simply a code execution.
//
// @param request - CreateRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRunResponse
func (client *Client) CreateRunWithOptions(request *CreateRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExperimentId) {
		body["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRun"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a run. A run is an experiment that can be associated with a specific workload or simply a code execution.
//
// @param request - CreateRunRequest
//
// @return CreateRunResponse
func (client *Client) CreateRun(request *CreateRunRequest) (_result *CreateRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRunResponse{}
	_body, _err := client.CreateRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EnvTypes) {
		body["EnvTypes"] = request.EnvTypes
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.WorkspaceName) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates resources with a workspace.
//
// @param request - CreateWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResourceResponse
func (client *Client) CreateWorkspaceResourceWithOptions(WorkspaceId *string, request *CreateWorkspaceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Option) {
		body["Option"] = request.Option
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspaceResource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates resources with a workspace.
//
// @param request - CreateWorkspaceResourceRequest
//
// @return CreateWorkspaceResourceResponse
func (client *Client) CreateWorkspaceResource(WorkspaceId *string, request *CreateWorkspaceResourceRequest) (_result *CreateWorkspaceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResourceResponse{}
	_body, _err := client.CreateWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a code source based on the provided ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCodeSourceResponse
func (client *Client) DeleteCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCodeSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCodeSource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources/" + dara.PercentEncode(dara.StringValue(CodeSourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a code source based on the provided ID.
//
// @return DeleteCodeSourceResponse
func (client *Client) DeleteCodeSource(CodeSourceId *string) (_result *DeleteCodeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.DeleteCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes workspace configurations.
//
// @param request - DeleteConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfigWithOptions(WorkspaceId *string, ConfigKey *string, request *DeleteConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/configs/" + dara.PercentEncode(dara.StringValue(ConfigKey))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes workspace configurations.
//
// @param request - DeleteConfigRequest
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfig(WorkspaceId *string, ConfigKey *string, request *DeleteConfigRequest) (_result *DeleteConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigResponse{}
	_body, _err := client.DeleteConfigWithOptions(WorkspaceId, ConfigKey, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a connection.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnectionWithOptions(ConnectionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnection"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/connections/" + dara.PercentEncode(dara.StringValue(ConnectionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection.
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnection(ConnectionId *string) (_result *DeleteConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConnectionResponse{}
	_body, _err := client.DeleteConnectionWithOptions(ConnectionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset.
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(DatasetId *string) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the metadata records of multiple files in a dataset at a time.
//
// @param request - DeleteDatasetFileMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetFileMetasResponse
func (client *Client) DeleteDatasetFileMetasWithOptions(DatasetId *string, request *DeleteDatasetFileMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetFileMetasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetFileMetaIds) {
		query["DatasetFileMetaIds"] = request.DatasetFileMetaIds
	}

	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetFileMetas"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetFileMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the metadata records of multiple files in a dataset at a time.
//
// @param request - DeleteDatasetFileMetasRequest
//
// @return DeleteDatasetFileMetasResponse
func (client *Client) DeleteDatasetFileMetas(DatasetId *string, request *DeleteDatasetFileMetasRequest) (_result *DeleteDatasetFileMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetFileMetasResponse{}
	_body, _err := client.DeleteDatasetFileMetasWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataset job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetJobResponse
func (client *Client) DeleteDatasetJobWithOptions(DatasetId *string, DatasetJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetJob"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs/" + dara.PercentEncode(dara.StringValue(DatasetJobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset job.
//
// @return DeleteDatasetJobResponse
func (client *Client) DeleteDatasetJob(DatasetId *string, DatasetJobId *string) (_result *DeleteDatasetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetJobResponse{}
	_body, _err := client.DeleteDatasetJobWithOptions(DatasetId, DatasetJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a job configuration for a dataset.
//
// @param request - DeleteDatasetJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetJobConfigResponse
func (client *Client) DeleteDatasetJobConfigWithOptions(DatasetId *string, DatasetJobConfigId *string, request *DeleteDatasetJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetJobConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobconfigs/" + dara.PercentEncode(dara.StringValue(DatasetJobConfigId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a job configuration for a dataset.
//
// @param request - DeleteDatasetJobConfigRequest
//
// @return DeleteDatasetJobConfigResponse
func (client *Client) DeleteDatasetJobConfig(DatasetId *string, DatasetJobConfigId *string, request *DeleteDatasetJobConfigRequest) (_result *DeleteDatasetJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetJobConfigResponse{}
	_body, _err := client.DeleteDatasetJobConfigWithOptions(DatasetId, DatasetJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dataset tag.
//
// @param request - DeleteDatasetLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetLabelsResponse
func (client *Client) DeleteDatasetLabelsWithOptions(DatasetId *string, request *DeleteDatasetLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset tag.
//
// @param request - DeleteDatasetLabelsRequest
//
// @return DeleteDatasetLabelsResponse
func (client *Client) DeleteDatasetLabels(DatasetId *string, request *DeleteDatasetLabelsRequest) (_result *DeleteDatasetLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetLabelsResponse{}
	_body, _err := client.DeleteDatasetLabelsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the information about a specified version of a dataset. Version v1 cannot be deleted by using this operation. When you call the DeleteDataset operation to delete a dataset, it can be deleted at the same time.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetVersionResponse
func (client *Client) DeleteDatasetVersionWithOptions(DatasetId *string, VersionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information about a specified version of a dataset. Version v1 cannot be deleted by using this operation. When you call the DeleteDataset operation to delete a dataset, it can be deleted at the same time.
//
// @return DeleteDatasetVersionResponse
func (client *Client) DeleteDatasetVersion(DatasetId *string, VersionName *string) (_result *DeleteDatasetVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetVersionResponse{}
	_body, _err := client.DeleteDatasetVersionWithOptions(DatasetId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes tags for a dataset version.
//
// @param request - DeleteDatasetVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetVersionLabelsResponse
func (client *Client) DeleteDatasetVersionLabelsWithOptions(DatasetId *string, VersionName *string, request *DeleteDatasetVersionLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetVersionLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keys) {
		query["Keys"] = request.Keys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetVersionLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tags for a dataset version.
//
// @param request - DeleteDatasetVersionLabelsRequest
//
// @return DeleteDatasetVersionLabelsResponse
func (client *Client) DeleteDatasetVersionLabels(DatasetId *string, VersionName *string, request *DeleteDatasetVersionLabelsRequest) (_result *DeleteDatasetVersionLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetVersionLabelsResponse{}
	_body, _err := client.DeleteDatasetVersionLabelsWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an experiment.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperiment"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an experiment.
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(ExperimentId *string) (_result *DeleteExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an experiment tag.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentLabelResponse
func (client *Client) DeleteExperimentLabelWithOptions(ExperimentId *string, Key *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperimentLabelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperimentLabel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/labels/" + dara.PercentEncode(dara.StringValue(Key))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperimentLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an experiment tag.
//
// @return DeleteExperimentLabelResponse
func (client *Client) DeleteExperimentLabel(ExperimentId *string, Key *string) (_result *DeleteExperimentLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentLabelResponse{}
	_body, _err := client.DeleteExperimentLabelWithOptions(ExperimentId, Key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a member from a workspace.
//
// @param request - DeleteMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMembersResponse
func (client *Client) DeleteMembersWithOptions(WorkspaceId *string, request *DeleteMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMembers"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/members"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a member from a workspace.
//
// @param request - DeleteMembersRequest
//
// @return DeleteMembersResponse
func (client *Client) DeleteMembers(WorkspaceId *string, request *DeleteMembersRequest) (_result *DeleteMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMembersResponse{}
	_body, _err := client.DeleteMembersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithOptions(ModelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @return DeleteModelResponse
func (client *Client) DeleteModel(ModelId *string) (_result *DeleteModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(ModelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the labels of a model.
//
// @param request - DeleteModelLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelLabelsResponse
func (client *Client) DeleteModelLabelsWithOptions(ModelId *string, request *DeleteModelLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the labels of a model.
//
// @param request - DeleteModelLabelsRequest
//
// @return DeleteModelLabelsResponse
func (client *Client) DeleteModelLabels(ModelId *string, request *DeleteModelLabelsRequest) (_result *DeleteModelLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelLabelsResponse{}
	_body, _err := client.DeleteModelLabelsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelVersionResponse
func (client *Client) DeleteModelVersionWithOptions(ModelId *string, VersionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model version.
//
// @return DeleteModelVersionResponse
func (client *Client) DeleteModelVersion(ModelId *string, VersionName *string) (_result *DeleteModelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelVersionResponse{}
	_body, _err := client.DeleteModelVersionWithOptions(ModelId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a model version tag.
//
// @param request - DeleteModelVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelVersionLabelsResponse
func (client *Client) DeleteModelVersionLabelsWithOptions(ModelId *string, VersionName *string, request *DeleteModelVersionLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelVersionLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelVersionLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a model version tag.
//
// @param request - DeleteModelVersionLabelsRequest
//
// @return DeleteModelVersionLabelsResponse
func (client *Client) DeleteModelVersionLabels(ModelId *string, VersionName *string, request *DeleteModelVersionLabelsRequest) (_result *DeleteModelVersionLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelVersionLabelsResponse{}
	_body, _err := client.DeleteModelVersionLabelsWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a run.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRunResponse
func (client *Client) DeleteRunWithOptions(RunId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRun"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a run.
//
// @return DeleteRunResponse
func (client *Client) DeleteRun(RunId *string) (_result *DeleteRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRunResponse{}
	_body, _err := client.DeleteRunWithOptions(RunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a tag that is added to a run.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRunLabelResponse
func (client *Client) DeleteRunLabelWithOptions(RunId *string, Key *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRunLabelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRunLabel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId)) + "/labels/" + dara.PercentEncode(dara.StringValue(Key))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRunLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag that is added to a run.
//
// @return DeleteRunLabelResponse
func (client *Client) DeleteRunLabel(RunId *string, Key *string) (_result *DeleteRunLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRunLabelResponse{}
	_body, _err := client.DeleteRunLabelWithOptions(RunId, Key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes user configurations.
//
// @param request - DeleteUserConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserConfigResponse
func (client *Client) DeleteUserConfigWithOptions(CategoryName *string, request *DeleteUserConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/userconfigs/" + dara.PercentEncode(dara.StringValue(CategoryName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes user configurations.
//
// @param request - DeleteUserConfigRequest
//
// @return DeleteUserConfigResponse
func (client *Client) DeleteUserConfig(CategoryName *string, request *DeleteUserConfigRequest) (_result *DeleteUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserConfigResponse{}
	_body, _err := client.DeleteUserConfigWithOptions(CategoryName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace. After you delete a workspace, the associated resources are not automatically released. You must manually release the resources.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(WorkspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace. After you delete a workspace, the associated resources are not automatically released. You must manually release the resources.
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(WorkspaceId *string) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource from a workspace. The resource is not deleted at the underlying layer.
//
// @param request - DeleteWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResourceResponse
func (client *Client) DeleteWorkspaceResourceWithOptions(WorkspaceId *string, request *DeleteWorkspaceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspaceResource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/resources"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource from a workspace. The resource is not deleted at the underlying layer.
//
// @param request - DeleteWorkspaceResourceRequest
//
// @return DeleteWorkspaceResourceResponse
func (client *Client) DeleteWorkspaceResource(WorkspaceId *string, request *DeleteWorkspaceResourceRequest) (_result *DeleteWorkspaceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResourceResponse{}
	_body, _err := client.DeleteWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a code source.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeSourceResponse
func (client *Client) GetCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCodeSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCodeSource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources/" + dara.PercentEncode(dara.StringValue(CodeSourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a code source.
//
// @return GetCodeSourceResponse
func (client *Client) GetCodeSource(CodeSourceId *string) (_result *GetCodeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeSourceResponse{}
	_body, _err := client.GetCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a workspace configuration item.
//
// @param request - GetConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigResponse
func (client *Client) GetConfigWithOptions(WorkspaceId *string, request *GetConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a workspace configuration item.
//
// @param request - GetConfigRequest
//
// @return GetConfigResponse
func (client *Client) GetConfig(WorkspaceId *string, request *GetConfigRequest) (_result *GetConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConfigResponse{}
	_body, _err := client.GetConfigWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the connection details.
//
// @param request - GetConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionResponse
func (client *Client) GetConnectionWithOptions(ConnectionId *string, request *GetConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptOption) {
		query["EncryptOption"] = request.EncryptOption
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnection"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/connections/" + dara.PercentEncode(dara.StringValue(ConnectionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the connection details.
//
// @param request - GetConnectionRequest
//
// @return GetConnectionResponse
func (client *Client) GetConnection(ConnectionId *string, request *GetConnectionRequest) (_result *GetConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConnectionResponse{}
	_body, _err := client.GetConnectionWithOptions(ConnectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a dataset.
//
// @return GetDatasetResponse
func (client *Client) GetDataset(DatasetId *string) (_result *GetDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata records of specific files in a dataset.
//
// @param request - GetDatasetFileMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetFileMetaResponse
func (client *Client) GetDatasetFileMetaWithOptions(DatasetId *string, DatasetFileMetaId *string, request *GetDatasetFileMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetFileMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetFileMeta"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas/" + dara.PercentEncode(dara.StringValue(DatasetFileMetaId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetFileMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata records of specific files in a dataset.
//
// @param request - GetDatasetFileMetaRequest
//
// @return GetDatasetFileMetaResponse
func (client *Client) GetDatasetFileMeta(DatasetId *string, DatasetFileMetaId *string, request *GetDatasetFileMetaRequest) (_result *GetDatasetFileMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetFileMetaResponse{}
	_body, _err := client.GetDatasetFileMetaWithOptions(DatasetId, DatasetFileMetaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains metadata statistics of a dataset.
//
// @param request - GetDatasetFileMetasStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetFileMetasStatisticsResponse
func (client *Client) GetDatasetFileMetasStatisticsWithOptions(DatasetId *string, request *GetDatasetFileMetasStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetFileMetasStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregateBy) {
		query["AggregateBy"] = request.AggregateBy
	}

	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetFileMetasStatistics"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/statistics/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetFileMetasStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains metadata statistics of a dataset.
//
// @param request - GetDatasetFileMetasStatisticsRequest
//
// @return GetDatasetFileMetasStatisticsResponse
func (client *Client) GetDatasetFileMetasStatistics(DatasetId *string, request *GetDatasetFileMetasStatisticsRequest) (_result *GetDatasetFileMetasStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetFileMetasStatisticsResponse{}
	_body, _err := client.GetDatasetFileMetasStatisticsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a dataset job.
//
// @param request - GetDatasetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetJobResponse
func (client *Client) GetDatasetJobWithOptions(DatasetId *string, DatasetJobId *string, request *GetDatasetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetJob"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs/" + dara.PercentEncode(dara.StringValue(DatasetJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a dataset job.
//
// @param request - GetDatasetJobRequest
//
// @return GetDatasetJobResponse
func (client *Client) GetDatasetJob(DatasetId *string, DatasetJobId *string, request *GetDatasetJobRequest) (_result *GetDatasetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetJobResponse{}
	_body, _err := client.GetDatasetJobWithOptions(DatasetId, DatasetJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a job configuration for a dataset.
//
// @param request - GetDatasetJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetJobConfigResponse
func (client *Client) GetDatasetJobConfigWithOptions(DatasetId *string, DatasetJobConfigId *string, request *GetDatasetJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetJobConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobconfigs/" + dara.PercentEncode(dara.StringValue(DatasetJobConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a job configuration for a dataset.
//
// @param request - GetDatasetJobConfigRequest
//
// @return GetDatasetJobConfigResponse
func (client *Client) GetDatasetJobConfig(DatasetId *string, DatasetJobConfigId *string, request *GetDatasetJobConfigRequest) (_result *GetDatasetJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetJobConfigResponse{}
	_body, _err := client.GetDatasetJobConfigWithOptions(DatasetId, DatasetJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a specified version of a dataset.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetVersionResponse
func (client *Client) GetDatasetVersionWithOptions(DatasetId *string, VersionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a specified version of a dataset.
//
// @return GetDatasetVersionResponse
func (client *Client) GetDatasetVersion(DatasetId *string, VersionName *string) (_result *GetDatasetVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetVersionResponse{}
	_body, _err := client.GetDatasetVersionWithOptions(DatasetId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the default workspace.
//
// @param request - GetDefaultWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultWorkspaceResponse
func (client *Client) GetDefaultWorkspaceWithOptions(request *GetDefaultWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDefaultWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/defaultWorkspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the default workspace.
//
// @param request - GetDefaultWorkspaceRequest
//
// @return GetDefaultWorkspaceResponse
func (client *Client) GetDefaultWorkspace(request *GetDefaultWorkspaceRequest) (_result *GetDefaultWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDefaultWorkspaceResponse{}
	_body, _err := client.GetDefaultWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains an experiment.
//
// @param request - GetExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(ExperimentId *string, request *GetExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperiment"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an experiment.
//
// @param request - GetExperimentRequest
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(ExperimentId *string, request *GetExperimentRequest) (_result *GetExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about an image.
//
// @param request - GetImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(ImageId *string, request *GetImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImage"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images/" + dara.PercentEncode(dara.StringValue(ImageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about an image.
//
// @param request - GetImageRequest
//
// @return GetImageResponse
func (client *Client) GetImage(ImageId *string, request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a member in a workspace.
//
// @param request - GetMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemberResponse
func (client *Client) GetMemberWithOptions(WorkspaceId *string, request *GetMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMember"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/member"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a member in a workspace.
//
// @param request - GetMemberRequest
//
// @return GetMemberResponse
func (client *Client) GetMember(WorkspaceId *string, request *GetMemberRequest) (_result *GetMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemberResponse{}
	_body, _err := client.GetMemberWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified model.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelResponse
func (client *Client) GetModelWithOptions(ModelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified model.
//
// @return GetModelResponse
func (client *Client) GetModel(ModelId *string) (_result *GetModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelResponse{}
	_body, _err := client.GetModelWithOptions(ModelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a model version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelVersionResponse
func (client *Client) GetModelVersionWithOptions(ModelId *string, VersionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a model version.
//
// @return GetModelVersionResponse
func (client *Client) GetModelVersion(ModelId *string, VersionName *string) (_result *GetModelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelVersionResponse{}
	_body, _err := client.GetModelVersionWithOptions(ModelId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains permissions on a workspace.
//
// @param tmpReq - GetPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionResponse
func (client *Client) GetPermissionWithOptions(WorkspaceId *string, PermissionCode *string, tmpReq *GetPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPermission"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/permissions/" + dara.PercentEncode(dara.StringValue(PermissionCode))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains permissions on a workspace.
//
// @param request - GetPermissionRequest
//
// @return GetPermissionResponse
func (client *Client) GetPermission(WorkspaceId *string, PermissionCode *string, request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(WorkspaceId, PermissionCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the run information.
//
// @param request - GetRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunResponse
func (client *Client) GetRunWithOptions(RunId *string, request *GetRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRun"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run information.
//
// @param request - GetRunRequest
//
// @return GetRunResponse
func (client *Client) GetRun(RunId *string, request *GetRunRequest) (_result *GetRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRunResponse{}
	_body, _err := client.GetRunWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a workspace.
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(WorkspaceId *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a workspace.
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(WorkspaceId *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists code sources. Pagination, sorting, and filtering by condition are supported.
//
// @param request - ListCodeSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCodeSourcesResponse
func (client *Client) ListCodeSourcesWithOptions(request *ListCodeSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCodeSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
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
		Action:      dara.String("ListCodeSources"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists code sources. Pagination, sorting, and filtering by condition are supported.
//
// @param request - ListCodeSourcesRequest
//
// @return ListCodeSourcesResponse
func (client *Client) ListCodeSources(request *ListCodeSourcesRequest) (_result *ListCodeSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.ListCodeSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of workspace configurations.
//
// @param request - ListConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigsResponse
func (client *Client) ListConfigsWithOptions(WorkspaceId *string, request *ListConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.ConfigKeys) {
		query["ConfigKeys"] = request.ConfigKeys
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of workspace configurations.
//
// @param request - ListConfigsRequest
//
// @return ListConfigsResponse
func (client *Client) ListConfigs(WorkspaceId *string, request *ListConfigsRequest) (_result *ListConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigsResponse{}
	_body, _err := client.ListConfigsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists connections.
//
// @param tmpReq - ListConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectionsResponse
func (client *Client) ListConnectionsWithOptions(tmpReq *ListConnectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListConnectionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConnectionIds) {
		request.ConnectionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConnectionIds, dara.String("ConnectionIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ConnectionTypes) {
		request.ConnectionTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConnectionTypes, dara.String("ConnectionTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ModelTypes) {
		request.ModelTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelTypes, dara.String("ModelTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionIdsShrink) {
		query["ConnectionIds"] = request.ConnectionIdsShrink
	}

	if !dara.IsNil(request.ConnectionName) {
		query["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.ConnectionTypesShrink) {
		query["ConnectionTypes"] = request.ConnectionTypesShrink
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.EncryptOption) {
		query["EncryptOption"] = request.EncryptOption
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.ModelTypesShrink) {
		query["ModelTypes"] = request.ModelTypesShrink
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.ToolCall) {
		query["ToolCall"] = request.ToolCall
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnections"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/connections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists connections.
//
// @param request - ListConnectionsRequest
//
// @return ListConnectionsResponse
func (client *Client) ListConnections(request *ListConnectionsRequest) (_result *ListConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConnectionsResponse{}
	_body, _err := client.ListConnectionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of dataset files.
//
// @param tmpReq - ListDatasetFileMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetFileMetasResponse
func (client *Client) ListDatasetFileMetasWithOptions(DatasetId *string, tmpReq *ListDatasetFileMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetFileMetasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDatasetFileMetasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryContentTypeIncludeAny) {
		request.QueryContentTypeIncludeAnyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryContentTypeIncludeAny, dara.String("QueryContentTypeIncludeAny"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QueryFileTypeIncludeAny) {
		request.QueryFileTypeIncludeAnyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryFileTypeIncludeAny, dara.String("QueryFileTypeIncludeAny"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QueryTagsExclude) {
		request.QueryTagsExcludeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryTagsExclude, dara.String("QueryTagsExclude"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QueryTagsIncludeAll) {
		request.QueryTagsIncludeAllShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryTagsIncludeAll, dara.String("QueryTagsIncludeAll"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.QueryTagsIncludeAny) {
		request.QueryTagsIncludeAnyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryTagsIncludeAny, dara.String("QueryTagsIncludeAny"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.EndFileUpdateTime) {
		query["EndFileUpdateTime"] = request.EndFileUpdateTime
	}

	if !dara.IsNil(request.EndTagUpdateTime) {
		query["EndTagUpdateTime"] = request.EndTagUpdateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryContentTypeIncludeAnyShrink) {
		query["QueryContentTypeIncludeAny"] = request.QueryContentTypeIncludeAnyShrink
	}

	if !dara.IsNil(request.QueryExpression) {
		query["QueryExpression"] = request.QueryExpression
	}

	if !dara.IsNil(request.QueryFileDir) {
		query["QueryFileDir"] = request.QueryFileDir
	}

	if !dara.IsNil(request.QueryFileName) {
		query["QueryFileName"] = request.QueryFileName
	}

	if !dara.IsNil(request.QueryFileTypeIncludeAnyShrink) {
		query["QueryFileTypeIncludeAny"] = request.QueryFileTypeIncludeAnyShrink
	}

	if !dara.IsNil(request.QueryImage) {
		query["QueryImage"] = request.QueryImage
	}

	if !dara.IsNil(request.QueryTagsExcludeShrink) {
		query["QueryTagsExclude"] = request.QueryTagsExcludeShrink
	}

	if !dara.IsNil(request.QueryTagsIncludeAllShrink) {
		query["QueryTagsIncludeAll"] = request.QueryTagsIncludeAllShrink
	}

	if !dara.IsNil(request.QueryTagsIncludeAnyShrink) {
		query["QueryTagsIncludeAny"] = request.QueryTagsIncludeAnyShrink
	}

	if !dara.IsNil(request.QueryText) {
		query["QueryText"] = request.QueryText
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ScoreThreshold) {
		query["ScoreThreshold"] = request.ScoreThreshold
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartFileUpdateTime) {
		query["StartFileUpdateTime"] = request.StartFileUpdateTime
	}

	if !dara.IsNil(request.StartTagUpdateTime) {
		query["StartTagUpdateTime"] = request.StartTagUpdateTime
	}

	if !dara.IsNil(request.ThumbnailMode) {
		query["ThumbnailMode"] = request.ThumbnailMode
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetFileMetas"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetFileMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of dataset files.
//
// @param request - ListDatasetFileMetasRequest
//
// @return ListDatasetFileMetasResponse
func (client *Client) ListDatasetFileMetas(DatasetId *string, request *ListDatasetFileMetasRequest) (_result *ListDatasetFileMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetFileMetasResponse{}
	_body, _err := client.ListDatasetFileMetasWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dataset job configurations at a time.
//
// @param request - ListDatasetJobConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetJobConfigsResponse
func (client *Client) ListDatasetJobConfigsWithOptions(DatasetId *string, request *ListDatasetJobConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetJobConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetJobConfigs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobconfigs/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetJobConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dataset job configurations at a time.
//
// @param request - ListDatasetJobConfigsRequest
//
// @return ListDatasetJobConfigsResponse
func (client *Client) ListDatasetJobConfigs(DatasetId *string, request *ListDatasetJobConfigsRequest) (_result *ListDatasetJobConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetJobConfigsResponse{}
	_body, _err := client.ListDatasetJobConfigsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists jobs in a dataset.
//
// @param request - ListDatasetJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetJobsResponse
func (client *Client) ListDatasetJobsWithOptions(DatasetId *string, request *ListDatasetJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		query["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.JobAction) {
		query["JobAction"] = request.JobAction
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetJobs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists jobs in a dataset.
//
// @param request - ListDatasetJobsRequest
//
// @return ListDatasetJobsResponse
func (client *Client) ListDatasetJobs(DatasetId *string, request *ListDatasetJobsRequest) (_result *ListDatasetJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetJobsResponse{}
	_body, _err := client.ListDatasetJobsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists dataset versions.
//
// @param request - ListDatasetVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetVersionsResponse
func (client *Client) ListDatasetVersionsWithOptions(DatasetId *string, request *ListDatasetVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	if !dara.IsNil(request.LabelValues) {
		query["LabelValues"] = request.LabelValues
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

	if !dara.IsNil(request.Properties) {
		query["Properties"] = request.Properties
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceTypes) {
		query["SourceTypes"] = request.SourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetVersions"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists dataset versions.
//
// @param request - ListDatasetVersionsRequest
//
// @return ListDatasetVersionsResponse
func (client *Client) ListDatasetVersions(DatasetId *string, request *ListDatasetVersionsRequest) (_result *ListDatasetVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetVersionsResponse{}
	_body, _err := client.ListDatasetVersionsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the datasets in a workspace.
//
// @param request - ListDatasetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.DataSourceTypes) {
		query["DataSourceTypes"] = request.DataSourceTypes
	}

	if !dara.IsNil(request.DataTypes) {
		query["DataTypes"] = request.DataTypes
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
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

	if !dara.IsNil(request.Properties) {
		query["Properties"] = request.Properties
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.ShareScope) {
		query["ShareScope"] = request.ShareScope
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceDatasetId) {
		query["SourceDatasetId"] = request.SourceDatasetId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceTypes) {
		query["SourceTypes"] = request.SourceTypes
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the datasets in a workspace.
//
// @param request - ListDatasetsRequest
//
// @return ListDatasetsResponse
func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists experiments.
//
// @param tmpReq - ListExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentResponse
func (client *Client) ListExperimentWithOptions(tmpReq *ListExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListExperimentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OptionsShrink) {
		query["Options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PageToken) {
		query["PageToken"] = request.PageToken
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperiment"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists experiments.
//
// @param request - ListExperimentRequest
//
// @return ListExperimentResponse
func (client *Client) ListExperiment(request *ListExperimentRequest) (_result *ListExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentResponse{}
	_body, _err := client.ListExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举特性
//
// @param request - ListFeaturesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFeaturesResponse
func (client *Client) ListFeaturesWithOptions(request *ListFeaturesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFeaturesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFeatures"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/features"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举特性
//
// @param request - ListFeaturesRequest
//
// @return ListFeaturesResponse
func (client *Client) ListFeatures(request *ListFeaturesRequest) (_result *ListFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeaturesResponse{}
	_body, _err := client.ListFeaturesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all tags of an image.
//
// @param request - ListImageLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageLabelsResponse
func (client *Client) ListImageLabelsWithOptions(request *ListImageLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListImageLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.LabelFilter) {
		query["LabelFilter"] = request.LabelFilter
	}

	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/image/labels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all tags of an image.
//
// @param request - ListImageLabelsRequest
//
// @return ListImageLabelsResponse
func (client *Client) ListImageLabels(request *ListImageLabelsRequest) (_result *ListImageLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImageLabelsResponse{}
	_body, _err := client.ListImageLabelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of images.
//
// @param request - ListImagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.ImageUri) {
		query["ImageUri"] = request.ImageUri
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
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

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of images.
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the members in a workspace.
//
// @param request - ListMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
func (client *Client) ListMembersWithOptions(WorkspaceId *string, request *ListMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberName) {
		query["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Roles) {
		query["Roles"] = request.Roles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMembers"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the members in a workspace.
//
// @param request - ListMembersRequest
//
// @return ListMembersResponse
func (client *Client) ListMembers(WorkspaceId *string, request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of model versions.
//
// @param request - ListModelVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelVersionsResponse
func (client *Client) ListModelVersionsWithOptions(ModelId *string, request *ListModelVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalStatus) {
		query["ApprovalStatus"] = request.ApprovalStatus
	}

	if !dara.IsNil(request.FormatType) {
		query["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.FrameworkType) {
		query["FrameworkType"] = request.FrameworkType
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
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

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelVersions"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of model versions.
//
// @param request - ListModelVersionsRequest
//
// @return ListModelVersionsResponse
func (client *Client) ListModelVersions(ModelId *string, request *ListModelVersionsRequest) (_result *ListModelVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelVersionsResponse{}
	_body, _err := client.ListModelVersionsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of models.
//
// @param tmpReq - ListModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
func (client *Client) ListModelsWithOptions(tmpReq *ListModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Conditions) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, dara.String("Conditions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collections) {
		query["Collections"] = request.Collections
	}

	if !dara.IsNil(request.ConditionsShrink) {
		query["Conditions"] = request.ConditionsShrink
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelType) {
		query["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Task) {
		query["Task"] = request.Task
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of models.
//
// @param request - ListModelsRequest
//
// @return ListModelsResponse
func (client *Client) ListModels(request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the permissions that a user has in a workspace.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithOptions(WorkspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissions"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/permissions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the permissions that a user has in a workspace.
//
// @return ListPermissionsResponse
func (client *Client) ListPermissions(WorkspaceId *string) (_result *ListPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举产品
//
// @param request - ListProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCodes) {
		query["ProductCodes"] = request.ProductCodes
	}

	if !dara.IsNil(request.ServiceCodes) {
		query["ServiceCodes"] = request.ServiceCodes
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举产品
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of quotas.
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotas"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of quotas.
//
// @param request - ListQuotasRequest
//
// @return ListQuotasResponse
func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated with a workspace.
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductTypes) {
		query["ProductTypes"] = request.ProductTypes
	}

	if !dara.IsNil(request.QuotaIds) {
		query["QuotaIds"] = request.QuotaIds
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.VerboseFields) {
		query["VerboseFields"] = request.VerboseFields
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated with a workspace.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the metrics for a run.
//
// @param request - ListRunMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunMetricsResponse
func (client *Client) ListRunMetricsWithOptions(RunId *string, request *ListRunMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRunMetricsResponse, _err error) {
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

	if !dara.IsNil(request.PageToken) {
		query["PageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRunMetrics"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId)) + "/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRunMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the metrics for a run.
//
// @param request - ListRunMetricsRequest
//
// @return ListRunMetricsResponse
func (client *Client) ListRunMetrics(RunId *string, request *ListRunMetricsRequest) (_result *ListRunMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRunMetricsResponse{}
	_body, _err := client.ListRunMetricsWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of runs.
//
// @param request - ListRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunsResponse
func (client *Client) ListRunsWithOptions(request *ListRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExperimentId) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !dara.IsNil(request.GmtCreateTime) {
		query["GmtCreateTime"] = request.GmtCreateTime
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PageToken) {
		query["PageToken"] = request.PageToken
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRuns"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of runs.
//
// @param request - ListRunsRequest
//
// @return ListRunsResponse
func (client *Client) ListRuns(request *ListRunsRequest) (_result *ListRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRunsResponse{}
	_body, _err := client.ListRunsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user information.
//
// @param request - ListUserConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserConfigsResponse
func (client *Client) ListUserConfigsWithOptions(request *ListUserConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryNames) {
		query["CategoryNames"] = request.CategoryNames
	}

	if !dara.IsNil(request.ConfigKeys) {
		query["ConfigKeys"] = request.ConfigKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserConfigs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/userconfigs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user information.
//
// @param request - ListUserConfigsRequest
//
// @return ListUserConfigsResponse
func (client *Client) ListUserConfigs(request *ListUserConfigsRequest) (_result *ListUserConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserConfigsResponse{}
	_body, _err := client.ListUserConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the users who do not belong to a workspace. These users can be added to the workspace as members.
//
// @param request - ListWorkspaceUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceUsersResponse
func (client *Client) ListWorkspaceUsersWithOptions(WorkspaceId *string, request *ListWorkspaceUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspaceUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceUsers"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the users who do not belong to a workspace. These users can be added to the workspace as members.
//
// @param request - ListWorkspaceUsersRequest
//
// @return ListWorkspaceUsersResponse
func (client *Client) ListWorkspaceUsers(WorkspaceId *string, request *ListWorkspaceUsersRequest) (_result *ListWorkspaceUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspaceUsersResponse{}
	_body, _err := client.ListWorkspaceUsersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all workspaces in a region.
//
// Description:
//
// You can use the option parameter to specify query options, so as to obtain different information about the workspaces.
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.ModuleList) {
		query["ModuleList"] = request.ModuleList
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all workspaces in a region.
//
// Description:
//
// You can use the option parameter to specify query options, so as to obtain different information about the workspaces.
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Logs multiple metrics for a run at a time.
//
// @param request - LogRunMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogRunMetricsResponse
func (client *Client) LogRunMetricsWithOptions(RunId *string, request *LogRunMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *LogRunMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Metrics) {
		body["Metrics"] = request.Metrics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LogRunMetrics"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId)) + "/metrics/action/log"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LogRunMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Logs multiple metrics for a run at a time.
//
// @param request - LogRunMetricsRequest
//
// @return LogRunMetricsResponse
func (client *Client) LogRunMetrics(RunId *string, request *LogRunMetricsRequest) (_result *LogRunMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LogRunMetricsResponse{}
	_body, _err := client.LogRunMetricsWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a private code source to a workspace to make the code source publicly accessible.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishCodeSourceResponse
func (client *Client) PublishCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishCodeSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishCodeSource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources/" + dara.PercentEncode(dara.StringValue(CodeSourceId)) + "/publish"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a private code source to a workspace to make the code source publicly accessible.
//
// @return PublishCodeSourceResponse
func (client *Client) PublishCodeSource(CodeSourceId *string) (_result *PublishCodeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishCodeSourceResponse{}
	_body, _err := client.PublishCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a private dataset in a workspace.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDatasetResponse
func (client *Client) PublishDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishDatasetResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDataset"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/publish"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a private dataset in a workspace.
//
// @return PublishDatasetResponse
func (client *Client) PublishDataset(DatasetId *string) (_result *PublishDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishDatasetResponse{}
	_body, _err := client.PublishDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes an image. After the image is published, the visibility of the image is changed from PRIVATE to PUBLIC.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishImageResponse
func (client *Client) PublishImageWithOptions(ImageId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishImageResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishImage"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images/" + dara.PercentEncode(dara.StringValue(ImageId)) + "/publish"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an image. After the image is published, the visibility of the image is changed from PRIVATE to PUBLIC.
//
// @return PublishImageResponse
func (client *Client) PublishImage(ImageId *string) (_result *PublishImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishImageResponse{}
	_body, _err := client.PublishImageWithOptions(ImageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an image.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageResponse
func (client *Client) RemoveImageWithOptions(ImageId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveImage"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images/" + dara.PercentEncode(dara.StringValue(ImageId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an image.
//
// @return RemoveImageResponse
func (client *Client) RemoveImage(ImageId *string) (_result *RemoveImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageResponse{}
	_body, _err := client.RemoveImageWithOptions(ImageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an image tag.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageLabelsResponse
func (client *Client) RemoveImageLabelsWithOptions(ImageId *string, LabelKey *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveImageLabelsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveImageLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/images/" + dara.PercentEncode(dara.StringValue(ImageId)) + "/labels/" + dara.PercentEncode(dara.StringValue(LabelKey))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an image tag.
//
// @return RemoveImageLabelsResponse
func (client *Client) RemoveImageLabels(ImageId *string, LabelKey *string) (_result *RemoveImageLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.RemoveImageLabelsWithOptions(ImageId, LabelKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a member role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMemberRoleResponse
func (client *Client) RemoveMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveMemberRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveMemberRole"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/members/" + dara.PercentEncode(dara.StringValue(MemberId)) + "/roles/" + dara.PercentEncode(dara.StringValue(RoleName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveMemberRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member role.
//
// @return RemoveMemberRoleResponse
func (client *Client) RemoveMemberRole(WorkspaceId *string, MemberId *string, RoleName *string) (_result *RemoveMemberRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveMemberRoleResponse{}
	_body, _err := client.RemoveMemberRoleWithOptions(WorkspaceId, MemberId, RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a experiment tag.
//
// @param request - SetExperimentLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetExperimentLabelsResponse
func (client *Client) SetExperimentLabelsWithOptions(ExperimentId *string, request *SetExperimentLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SetExperimentLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetExperimentLabels"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SetExperimentLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a experiment tag.
//
// @param request - SetExperimentLabelsRequest
//
// @return SetExperimentLabelsResponse
func (client *Client) SetExperimentLabels(ExperimentId *string, request *SetExperimentLabelsRequest) (_result *SetExperimentLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetExperimentLabelsResponse{}
	_body, _err := client.SetExperimentLabelsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the user configurations.
//
// @param request - SetUserConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserConfigsResponse
func (client *Client) SetUserConfigsWithOptions(request *SetUserConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SetUserConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserConfigs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/userconfigs"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the user configurations.
//
// @param request - SetUserConfigsRequest
//
// @return SetUserConfigsResponse
func (client *Client) SetUserConfigs(request *SetUserConfigsRequest) (_result *SetUserConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetUserConfigsResponse{}
	_body, _err := client.SetUserConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a dataset job.
//
// @param request - StopDatasetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDatasetJobResponse
func (client *Client) StopDatasetJobWithOptions(DatasetId *string, DatasetJobId *string, request *StopDatasetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopDatasetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDatasetJob"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs/" + dara.PercentEncode(dara.StringValue(DatasetJobId)) + "/action/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDatasetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a dataset job.
//
// @param request - StopDatasetJobRequest
//
// @return StopDatasetJobResponse
func (client *Client) StopDatasetJob(DatasetId *string, DatasetJobId *string, request *StopDatasetJobRequest) (_result *StopDatasetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopDatasetJobResponse{}
	_body, _err := client.StopDatasetJobWithOptions(DatasetId, DatasetJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a code build.
//
// @param request - UpdateCodeSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCodeSourceResponse
func (client *Client) UpdateCodeSourceWithOptions(CodeSourceId *string, request *UpdateCodeSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCodeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeBranch) {
		body["CodeBranch"] = request.CodeBranch
	}

	if !dara.IsNil(request.CodeCommit) {
		body["CodeCommit"] = request.CodeCommit
	}

	if !dara.IsNil(request.CodeRepo) {
		body["CodeRepo"] = request.CodeRepo
	}

	if !dara.IsNil(request.CodeRepoAccessToken) {
		body["CodeRepoAccessToken"] = request.CodeRepoAccessToken
	}

	if !dara.IsNil(request.CodeRepoUserName) {
		body["CodeRepoUserName"] = request.CodeRepoUserName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.MountPath) {
		body["MountPath"] = request.MountPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCodeSource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/codesources/" + dara.PercentEncode(dara.StringValue(CodeSourceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a code build.
//
// @param request - UpdateCodeSourceRequest
//
// @return UpdateCodeSourceResponse
func (client *Client) UpdateCodeSource(CodeSourceId *string, request *UpdateCodeSourceRequest) (_result *UpdateCodeSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCodeSourceResponse{}
	_body, _err := client.UpdateCodeSourceWithOptions(CodeSourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates or adds a workspace configuration item.
//
// @param request - UpdateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfigWithOptions(WorkspaceId *string, request *UpdateConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.ConfigKey) {
		body["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.ConfigValue) {
		body["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates or adds a workspace configuration item.
//
// @param request - UpdateConfigRequest
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfig(WorkspaceId *string, request *UpdateConfigRequest) (_result *UpdateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigResponse{}
	_body, _err := client.UpdateConfigWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates or adds workspace configurations in batches.
//
// @param request - UpdateConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigsResponse
func (client *Client) UpdateConfigsWithOptions(WorkspaceId *string, request *UpdateConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigs"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/configs"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates or adds workspace configurations in batches.
//
// @param request - UpdateConfigsRequest
//
// @return UpdateConfigsResponse
func (client *Client) UpdateConfigs(WorkspaceId *string, request *UpdateConfigsRequest) (_result *UpdateConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigsResponse{}
	_body, _err := client.UpdateConfigsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a connection.
//
// @param request - UpdateConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectionResponse
func (client *Client) UpdateConnectionWithOptions(ConnectionId *string, request *UpdateConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Models) {
		body["Models"] = request.Models
	}

	if !dara.IsNil(request.Secrets) {
		body["Secrets"] = request.Secrets
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnection"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/connections/" + dara.PercentEncode(dara.StringValue(ConnectionId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a connection.
//
// @param request - UpdateConnectionRequest
//
// @return UpdateConnectionResponse
func (client *Client) UpdateConnection(ConnectionId *string, request *UpdateConnectionRequest) (_result *UpdateConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConnectionResponse{}
	_body, _err := client.UpdateConnectionWithOptions(ConnectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name, description, and other information about a dataset.
//
// @param request - UpdateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithOptions(DatasetId *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Edition) {
		body["Edition"] = request.Edition
	}

	if !dara.IsNil(request.MountAccessReadWriteRoleIdList) {
		body["MountAccessReadWriteRoleIdList"] = request.MountAccessReadWriteRoleIdList
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.SharingConfig) {
		body["SharingConfig"] = request.SharingConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name, description, and other information about a dataset.
//
// @param request - UpdateDatasetRequest
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDataset(DatasetId *string, request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the metadata records of multiple files in a dataset at a time.
//
// @param request - UpdateDatasetFileMetasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetFileMetasResponse
func (client *Client) UpdateDatasetFileMetasWithOptions(DatasetId *string, request *UpdateDatasetFileMetasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetFileMetasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetFileMetas) {
		body["DatasetFileMetas"] = request.DatasetFileMetas
	}

	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.TagJobId) {
		body["TagJobId"] = request.TagJobId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetFileMetas"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetfilemetas"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetFileMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the metadata records of multiple files in a dataset at a time.
//
// @param request - UpdateDatasetFileMetasRequest
//
// @return UpdateDatasetFileMetasResponse
func (client *Client) UpdateDatasetFileMetas(DatasetId *string, request *UpdateDatasetFileMetasRequest) (_result *UpdateDatasetFileMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetFileMetasResponse{}
	_body, _err := client.UpdateDatasetFileMetasWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a dataset job.
//
// @param request - UpdateDatasetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetJobResponse
func (client *Client) UpdateDatasetJobWithOptions(DatasetId *string, DatasetJobId *string, request *UpdateDatasetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetVersion) {
		body["DatasetVersion"] = request.DatasetVersion
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetJob"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobs/" + dara.PercentEncode(dara.StringValue(DatasetJobId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a dataset job.
//
// @param request - UpdateDatasetJobRequest
//
// @return UpdateDatasetJobResponse
func (client *Client) UpdateDatasetJob(DatasetId *string, DatasetJobId *string, request *UpdateDatasetJobRequest) (_result *UpdateDatasetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetJobResponse{}
	_body, _err := client.UpdateDatasetJobWithOptions(DatasetId, DatasetJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a job configuration for a dataset.
//
// @param request - UpdateDatasetJobConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetJobConfigResponse
func (client *Client) UpdateDatasetJobConfigWithOptions(DatasetId *string, DatasetJobConfigId *string, request *UpdateDatasetJobConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetJobConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.ConfigType) {
		body["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetJobConfig"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/datasetjobconfigs/" + dara.PercentEncode(dara.StringValue(DatasetJobConfigId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetJobConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a job configuration for a dataset.
//
// @param request - UpdateDatasetJobConfigRequest
//
// @return UpdateDatasetJobConfigResponse
func (client *Client) UpdateDatasetJobConfig(DatasetId *string, DatasetJobConfigId *string, request *UpdateDatasetJobConfigRequest) (_result *UpdateDatasetJobConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetJobConfigResponse{}
	_body, _err := client.UpdateDatasetJobConfigWithOptions(DatasetId, DatasetJobConfigId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a specified version of a dataset.
//
// @param request - UpdateDatasetVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetVersionResponse
func (client *Client) UpdateDatasetVersionWithOptions(DatasetId *string, VersionName *string, request *UpdateDatasetVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataCount) {
		body["DataCount"] = request.DataCount
	}

	if !dara.IsNil(request.DataSize) {
		body["DataSize"] = request.DataSize
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(DatasetId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a specified version of a dataset.
//
// @param request - UpdateDatasetVersionRequest
//
// @return UpdateDatasetVersionResponse
func (client *Client) UpdateDatasetVersion(DatasetId *string, VersionName *string, request *UpdateDatasetVersionRequest) (_result *UpdateDatasetVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetVersionResponse{}
	_body, _err := client.UpdateDatasetVersionWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a workspace as the default workspace.
//
// @param request - UpdateDefaultWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDefaultWorkspaceResponse
func (client *Client) UpdateDefaultWorkspaceWithOptions(request *UpdateDefaultWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDefaultWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDefaultWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/defaultWorkspaces"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDefaultWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a workspace as the default workspace.
//
// @param request - UpdateDefaultWorkspaceRequest
//
// @return UpdateDefaultWorkspaceResponse
func (client *Client) UpdateDefaultWorkspace(request *UpdateDefaultWorkspaceRequest) (_result *UpdateDefaultWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDefaultWorkspaceResponse{}
	_body, _err := client.UpdateDefaultWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an experiment.
//
// @param request - UpdateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperimentWithOptions(ExperimentId *string, request *UpdateExperimentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExperiment"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/experiments/" + dara.PercentEncode(dara.StringValue(ExperimentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an experiment.
//
// @param request - UpdateExperimentRequest
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperiment(ExperimentId *string, request *UpdateExperimentRequest) (_result *UpdateExperimentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentResponse{}
	_body, _err := client.UpdateExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic configuration information about a model.
//
// @param request - UpdateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelResponse
func (client *Client) UpdateModelWithOptions(ModelId *string, request *UpdateModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.ModelDescription) {
		body["ModelDescription"] = request.ModelDescription
	}

	if !dara.IsNil(request.ModelDoc) {
		body["ModelDoc"] = request.ModelDoc
	}

	if !dara.IsNil(request.ModelName) {
		body["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelType) {
		body["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.OrderNumber) {
		body["OrderNumber"] = request.OrderNumber
	}

	if !dara.IsNil(request.Origin) {
		body["Origin"] = request.Origin
	}

	if !dara.IsNil(request.ParameterSize) {
		body["ParameterSize"] = request.ParameterSize
	}

	if !dara.IsNil(request.Task) {
		body["Task"] = request.Task
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModel"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic configuration information about a model.
//
// @param request - UpdateModelRequest
//
// @return UpdateModelResponse
func (client *Client) UpdateModel(ModelId *string, request *UpdateModelRequest) (_result *UpdateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelResponse{}
	_body, _err := client.UpdateModelWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a model version.
//
// @param request - UpdateModelVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelVersionResponse
func (client *Client) UpdateModelVersionWithOptions(ModelId *string, VersionName *string, request *UpdateModelVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalStatus) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !dara.IsNil(request.CompressionSpec) {
		body["CompressionSpec"] = request.CompressionSpec
	}

	if !dara.IsNil(request.DistillationSpec) {
		body["DistillationSpec"] = request.DistillationSpec
	}

	if !dara.IsNil(request.EvaluationSpec) {
		body["EvaluationSpec"] = request.EvaluationSpec
	}

	if !dara.IsNil(request.ExtraInfo) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.InferenceSpec) {
		body["InferenceSpec"] = request.InferenceSpec
	}

	if !dara.IsNil(request.Metrics) {
		body["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.TrainingSpec) {
		body["TrainingSpec"] = request.TrainingSpec
	}

	if !dara.IsNil(request.VersionDescription) {
		body["VersionDescription"] = request.VersionDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelVersion"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/versions/" + dara.PercentEncode(dara.StringValue(VersionName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model version.
//
// @param request - UpdateModelVersionRequest
//
// @return UpdateModelVersionResponse
func (client *Client) UpdateModelVersion(ModelId *string, VersionName *string, request *UpdateModelVersionRequest) (_result *UpdateModelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelVersionResponse{}
	_body, _err := client.UpdateModelVersionWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the run information.
//
// @param request - UpdateRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRunResponse
func (client *Client) UpdateRunWithOptions(RunId *string, request *UpdateRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRun"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/runs/" + dara.PercentEncode(dara.StringValue(RunId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the run information.
//
// @param request - UpdateRunRequest
//
// @return UpdateRunResponse
func (client *Client) UpdateRun(RunId *string, request *UpdateRunRequest) (_result *UpdateRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRunResponse{}
	_body, _err := client.UpdateRunWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name and description of a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(WorkspaceId *string, request *UpdateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspace"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and description of a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(WorkspaceId *string, request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the resources of a workspace.
//
// @param request - UpdateWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResourceResponse
func (client *Client) UpdateWorkspaceResourceWithOptions(WorkspaceId *string, request *UpdateWorkspaceResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IsDefault) {
		body["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspaceResource"),
		Version:     dara.String("2021-02-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/workspaces/" + dara.PercentEncode(dara.StringValue(WorkspaceId)) + "/resources"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resources of a workspace.
//
// @param request - UpdateWorkspaceResourceRequest
//
// @return UpdateWorkspaceResourceResponse
func (client *Client) UpdateWorkspaceResource(WorkspaceId *string, request *UpdateWorkspaceResourceRequest) (_result *UpdateWorkspaceResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResourceResponse{}
	_body, _err := client.UpdateWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
