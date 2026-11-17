// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sets the execution time for an OTA upgrade.
//
// @param request - ApproveOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOtaTaskResponse
func (client *Client) ApproveOtaTaskWithContext(ctx context.Context, request *ApproveOtaTaskRequest, runtime *dara.RuntimeOptions) (_result *ApproveOtaTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.OtaType) {
		body["OtaType"] = request.OtaType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveOtaTask"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns secondary private IP addresses to a development host.
//
// @param request - AssignWuyingServerPrivateAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignWuyingServerPrivateAddressesResponse
func (client *Client) AssignWuyingServerPrivateAddressesWithContext(ctx context.Context, request *AssignWuyingServerPrivateAddressesRequest, runtime *dara.RuntimeOptions) (_result *AssignWuyingServerPrivateAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SecondaryPrivateIpAddressCount) {
		body["SecondaryPrivateIpAddressCount"] = request.SecondaryPrivateIpAddressCount
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignWuyingServerPrivateAddresses"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignWuyingServerPrivateAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or removes assigned users for a delivery group. Only users added as assigned users can access cloud applications.
//
// Description:
//
// > After you change assigned users, the selected users receive notification emails. Changes typically take about 2 minutes to take effect on the client.
//
// @param tmpReq - AuthorizeInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeInstanceGroupResponse
func (client *Client) AuthorizeInstanceGroupWithContext(ctx context.Context, tmpReq *AuthorizeInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthorizeInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserMeta) {
		request.UserMetaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserMeta, dara.String("UserMeta"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizeUserGroupIds) {
		bodyFlat["AuthorizeUserGroupIds"] = request.AuthorizeUserGroupIds
	}

	if !dara.IsNil(request.AuthorizeUserIds) {
		bodyFlat["AuthorizeUserIds"] = request.AuthorizeUserIds
	}

	if !dara.IsNil(request.AvatarId) {
		body["AvatarId"] = request.AvatarId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.UnAuthorizeUserGroupIds) {
		bodyFlat["UnAuthorizeUserGroupIds"] = request.UnAuthorizeUserGroupIds
	}

	if !dara.IsNil(request.UnAuthorizeUserIds) {
		bodyFlat["UnAuthorizeUserIds"] = request.UnAuthorizeUserIds
	}

	if !dara.IsNil(request.UserMetaShrink) {
		body["UserMeta"] = request.UserMetaShrink
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or removes authorized users for a specified application deployed in an app instance group. Only authorized users can access the application.
//
// Description:
//
// ## Operation description
//
// This operation manages user authorization for an app instance group at the application level. The authorization result applies only to the application specified by AppId and does not affect the authorization relationships of other applications in the app instance group. To authorize users for the entire app instance group, call the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation.
//
// ## Before you begin
//
// - The app instance group is created, and **the application specified by AppId is deployed in the image used by the app instance group**. Otherwise, the error code `InvalidAppId.NotFound` is returned.
//
// - The app instance group **is not added to an active app instance group set**. App instance groups added to a set cannot be authorized individually and must be authorized through the set. Otherwise, the error code `InvalidAppInstanceGroup.AuthorizeBlockedBySet` is returned.
//
// - If the workspace to which the app instance group belongs is an AD workspace, **UserMeta must be specified**, the value of `UserMeta.Type` must be `ad`, and the value of `UserMeta.AdDomain` must be the same as the AD domain bound to the workspace.
//
// - If the app instance group has been authorized through a user group and mixed authorization of users and user groups is not supported, you cannot authorize individual users. Otherwise, the error code `AuthAppInstanceGroup.MixNotSupported` is returned.
//
// ## Parameter description
//
// - **Specify at least one of AuthorizeUserIds and UnAuthorizeUserIds**, or specify both. If both are empty, no authorization relationships are changed when you invoke this operation.
//
// - When adding authorization, the sum of the current number of authorized users for the application and the number of newly added users cannot exceed the authorized user quota for the application. If the quota is exceeded, the error code `ExceedAppAuthUserQuota` is returned. Removing authorization is not subject to the quota limit.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) or [GetAppInstanceGroup](https://help.aliyun.com/document_detail/600836.html) operation to obtain the app instance group ID (AppInstanceGroupId) and the application IDs of the applications deployed in the app instance group (AppId in the Apps list).
//
// 2. Call the [DescribeUsers](https://help.aliyun.com/document_detail/436936.html) operation to obtain the usernames of the users to be authorized or whose authorization is to be removed.
//
// 3. Call this operation to complete the authorization change.
//
// > After the authorization is changed, the selected users will receive a notification email. Generally, it takes about 2 minutes for the changes to take effect on the terminal.
//
// @param tmpReq - AuthorizeUsersForAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeUsersForAppResponse
func (client *Client) AuthorizeUsersForAppWithContext(ctx context.Context, tmpReq *AuthorizeUsersForAppRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeUsersForAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthorizeUsersForAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserMeta) {
		request.UserMetaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserMeta, dara.String("UserMeta"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AuthorizeUserIds) {
		body["AuthorizeUserIds"] = request.AuthorizeUserIds
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.UnAuthorizeUserIds) {
		body["UnAuthorizeUserIds"] = request.UnAuthorizeUserIds
	}

	if !dara.IsNil(request.UserMetaShrink) {
		body["UserMeta"] = request.UserMetaShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeUsersForApp"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeUsersForAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures third-party channels for multiple agent runtimes in batches, supports delivering different channel configurations to different runtimes, and returns the configuration result for each configuration item.
//
// @param request - BatchConfigRuntimeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchConfigRuntimeChannelResponse
func (client *Client) BatchConfigRuntimeChannelWithContext(ctx context.Context, request *BatchConfigRuntimeChannelRequest, runtime *dara.RuntimeOptions) (_result *BatchConfigRuntimeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.ChannelConfigs) {
		body["ChannelConfigs"] = request.ChannelConfigs
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchConfigRuntimeChannel"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchConfigRuntimeChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple LLM templates.
//
// Description:
//
// You can create multiple model templates under a model provider template in the Wuying Agent Management Center. You can add multiple models at a time and specify one of them as the default model. Existing models are automatically skipped to prevent duplicate model creation.
//
// Make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center before you call this operation.
//
// @param request - BatchCreateLlmTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateLlmTemplatesResponse
func (client *Client) BatchCreateLlmTemplatesWithContext(ctx context.Context, request *BatchCreateLlmTemplatesRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateLlmTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LlmTemplateItems) {
		body["LlmTemplateItems"] = request.LlmTemplateItems
	}

	if !dara.IsNil(request.ModelTemplateId) {
		body["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.ProviderTemplateId) {
		body["ProviderTemplateId"] = request.ProviderTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateLlmTemplates"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateLlmTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image based on an existing image and starts an image creation task.
//
// Description:
//
// Before you call this operation, make sure that the parent image is available and an available image distribution record exists in the creation region.
//
// 1. Call ListImage to obtain the parent image ID.
//
// 2. Specify the creation mode, cloud disk capacity, and resource type, and then call this operation.
//
// 3. Save the returned ImageId, VersionId, and TaskId.
//
// 4. Call GetTask to query the creation progress. **A successful request does not indicate that the image creation is complete.**
//
// @param tmpReq - BuilderImageByParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuilderImageByParentResponse
func (client *Client) BuilderImageByParentWithContext(ctx context.Context, tmpReq *BuilderImageByParentRequest, runtime *dara.RuntimeOptions) (_result *BuilderImageByParentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BuilderImageByParentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PropertyMap) {
		request.PropertyMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PropertyMap, dara.String("PropertyMap"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DataDiskSize) {
		body["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Distro) {
		body["Distro"] = request.Distro
	}

	if !dara.IsNil(request.DriverIdList) {
		body["DriverIdList"] = request.DriverIdList
	}

	if !dara.IsNil(request.ImageCreateMode) {
		body["ImageCreateMode"] = request.ImageCreateMode
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.NodeInstanceType) {
		body["NodeInstanceType"] = request.NodeInstanceType
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.ParentImageId) {
		body["ParentImageId"] = request.ParentImageId
	}

	if !dara.IsNil(request.PerformanceType) {
		body["PerformanceType"] = request.PerformanceType
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.PropertyMapShrink) {
		body["PropertyMap"] = request.PropertyMapShrink
	}

	if !dara.IsNil(request.ResourceList) {
		body["ResourceList"] = request.ResourceList
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.SourceRegion) {
		body["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.VSwitchIdList) {
		body["VSwitchIdList"] = request.VSwitchIdList
	}

	if !dara.IsNil(request.VersionName) {
		body["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuilderImageByParent"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuilderImageByParentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the memory of a specified agent instance.
//
// Description:
//
// Call `ListAgentInstanceMemory` to query the target instance ID and the associated user. **Verify the instance and user before clearing the memory.**
//
// @param request - ClearAgentInstanceMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearAgentInstanceMemoryResponse
func (client *Client) ClearAgentInstanceMemoryWithContext(ctx context.Context, request *ClearAgentInstanceMemoryRequest, runtime *dara.RuntimeOptions) (_result *ClearAgentInstanceMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentInstanceId) {
		body["AgentInstanceId"] = request.AgentInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearAgentInstanceMemory"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearAgentInstanceMemoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the model group for a resource group.
//
// Description:
//
// You can assign model groups to the resource groups that contain agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the WUYING Agent Management Center. These model groups serve as inference engines for agents within the resource groups to execute tasks.
//
// If an agent runtime has its own model group configured and its resource group also has a model group configured, the model group bound to the resource group takes effect. That is, the resource group settings take precedence over the agent runtime settings.
//
// **Before you begin**
//
// - The ResourceGroupId parameter in this operation specifically refers to the WUYING Enterprise Edition resource group ID, which starts with rg-. This resource group is managed by the WUYING Enterprise Edition control plane and must be created before you deploy agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent. You do not need to and cannot create it by using the API operations of this service. Although the account-level resource group IDs in Alibaba Cloud Resource Manager also start with rg-, they do not belong to WUYING resource groups and cannot be passed as this parameter.
//
// - Before you call this operation, obtain an existing WUYING resource group ID by using one of the following methods:
//
//   - Call the DescribeResourceGroups operation of WUYING Cloud Desktop eds-user (API version 2021-03-08). The ResourceGroup[].ResourceGroupId field in the response is the resource group ID required for this parameter. You can set the ResourceClassification request parameter to DesktopAgent and filter agent resource groups by AgentType (JVSClaw or HermesAgent). The AliyunResourceGroupId field in the response is the resource group ID of Alibaba Cloud Resource Manager and cannot be used for this operation.
//
//   - Call the GetRuntimeModelConfig operation of this service. The Data.ResourceGroupId field in the response is the WUYING resource group ID to which the specified agent runtime currently belongs. If the runtime does not belong to any resource group, this field is null.
//
//   - When you call the ListDesktopAgentRuntime operation of this service, you can use the ResourceGroupId request parameter to filter agent runtimes by resource group to confirm which runtimes are included in a specific resource group.
//
// - The ListModelTemplateResourceGroup operation returns only the resource groups that have been authorized for a specified model group. It cannot be used to discover resource groups that have not been authorized for any model group.
//
// Make sure that you are fully familiar with the operations and usage of the WUYING Agent Management Center before you call this operation.
//
// @param request - ConfigResourceGroupModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigResourceGroupModelTemplateResponse
func (client *Client) ConfigResourceGroupModelTemplateWithContext(ctx context.Context, request *ConfigResourceGroupModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *ConfigResourceGroupModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		body["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigResourceGroupModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigResourceGroupModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures third-party channels for agent runtimes.
//
// Description:
//
// You can configure third-party channels for agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. These channels serve as extended agent communication methods in addition to the AgentIM channel.
//
// Before you call this operation, make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - ConfigRuntimeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigRuntimeChannelResponse
func (client *Client) ConfigRuntimeChannelWithContext(ctx context.Context, request *ConfigRuntimeChannelRequest, runtime *dara.RuntimeOptions) (_result *ConfigRuntimeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.ConfigMode) {
		body["ConfigMode"] = request.ConfigMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RuntimeIds) {
		body["RuntimeIds"] = request.RuntimeIds
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigRuntimeChannel"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigRuntimeChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures model groups for agent runtime resources.
//
// Description:
//
// You can authorize model groups for agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center to serve as inference engines for agent task execution.
//
// Before you call this operation, make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - ConfigRuntimeModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigRuntimeModelTemplateResponse
func (client *Client) ConfigRuntimeModelTemplateWithContext(ctx context.Context, request *ConfigRuntimeModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *ConfigRuntimeModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		body["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.RuntimeIds) {
		body["RuntimeIds"] = request.RuntimeIds
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigRuntimeModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigRuntimeModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application and binds the application to a specified image version.
//
// Description:
//
// Only custom images and community images manageable by the current account are supported. If VersionId is omitted, the current active version of the image is used. If specified, it must belong to the image corresponding to ImageId.
//
// Call ListImage first to select an image. After creation, save the returned AppId, AppVersionId, ImageId, and VersionId. To dissociate the application, call UnbindApp.
//
// @param request - CreateAndBindAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndBindAppResponse
func (client *Client) CreateAndBindAppWithContext(ctx context.Context, request *CreateAndBindAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAndBindAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppIconUrl) {
		body["AppIconUrl"] = request.AppIconUrl
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.OriginAppType) {
		body["OriginAppType"] = request.OriginAppType
	}

	if !dara.IsNil(request.Port) {
		body["Port"] = request.Port
	}

	if !dara.IsNil(request.UrlSuffix) {
		body["UrlSuffix"] = request.UrlSuffix
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndBindApp"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndBindAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery group.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing and pricing](https://help.aliyun.com/document_detail/426039.html) of WUYING CloudApp.
//
// A delivery group is a logical grouping for delivering cloud applications to end users. It includes the underlying cloud application resources, images that contain cloud applications, resource management policies, and user allocation settings. For details, see [Publish a delivery group](https://help.aliyun.com/document_detail/426046.html).
//
// @param tmpReq - CreateAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInstanceGroupResponse
func (client *Client) CreateAppInstanceGroupWithContext(ctx context.Context, tmpReq *CreateAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Network) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, dara.String("Network"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodePool) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, dara.String("NodePool"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimePolicy) {
		request.RuntimePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimePolicy, dara.String("RuntimePolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityPolicy) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, dara.String("SecurityPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoragePolicy) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, dara.String("StoragePolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserDefinePolicy) {
		request.UserDefinePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserDefinePolicy, dara.String("UserDefinePolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoPolicy) {
		request.VideoPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoPolicy, dara.String("VideoPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserDefinePolicyShrink) {
		query["UserDefinePolicy"] = request.UserDefinePolicyShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCenterImageId) {
		body["AppCenterImageId"] = request.AppCenterImageId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppPackageType) {
		body["AppPackageType"] = request.AppPackageType
	}

	if !dara.IsNil(request.AppPolicyId) {
		body["AppPolicyId"] = request.AppPolicyId
	}

	if !dara.IsNil(request.AuthMode) {
		body["AuthMode"] = request.AuthMode
	}

	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeResourceMode) {
		body["ChargeResourceMode"] = request.ChargeResourceMode
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NetworkShrink) {
		body["Network"] = request.NetworkShrink
	}

	if !dara.IsNil(request.NodePoolShrink) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PreOpenAppId) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RuntimePolicyShrink) {
		body["RuntimePolicy"] = request.RuntimePolicyShrink
	}

	if !dara.IsNil(request.SecurityPolicyShrink) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !dara.IsNil(request.SessionTimeout) {
		body["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.StoragePolicyShrink) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	if !dara.IsNil(request.SubPayType) {
		body["SubPayType"] = request.SubPayType
	}

	if !dara.IsNil(request.UserGroupIds) {
		body["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	if !dara.IsNil(request.VideoPolicyShrink) {
		body["VideoPolicy"] = request.VideoPolicyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud browser group billed by monthly active users (MAU).
//
// Description:
//
// ## Prerequisites
//
// - Prepare an available office network, image, and specifications in the target business region, and ensure that the account has the corresponding browser configurations and resource quotas.
//
// - Specify `CloudBrowserName` and `BizRegionId`. You must set `OsType` to `Windows`.
//
// - Authorized users must be created in advance and match the account type. Authorized user groups must belong to the current account and match the account type of the office network.
//
// - **`Users` and `UserGroupIds` cannot be both non-empty.**
//
// ## MAU billing parameters
//
// - Set `ChargeType` to `PostPaid`.
//
// - **You must explicitly set `SubPayType` to `mau`. Omitting this field does not enable MAU billing.**
//
// - Set `ChargeResourceMode` to `AppInstance`.
//
// - You do not need to specify `Period`, `PeriodUnit`, `AppPackageType`, `AutoPay`, `AutoRenew`, or `NodePool`.
//
// ## Post-call processing
//
// **A successful API response does not indicate that the browser resources are ready.*	- After creation, query the status of the browser group and use it only after confirming that it is connectable.
//
// This operation creates a new cloud browser group and does not require a delivery group to be created in advance.
//
// ## Example descriptions
//
// Field examples are used to demonstrate how to specify values. Replace resource identifiers with the actual values under the current account. Capacity examples do not represent default values or upper limits.
//
// An example value of `-` indicates that the field is not required. When calling the API, omit the corresponding parameter and do not pass the `-` character.
//
// @param tmpReq - CreateBrowserInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBrowserInstanceGroupResponse
func (client *Client) CreateBrowserInstanceGroupWithContext(ctx context.Context, tmpReq *CreateBrowserInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateBrowserInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBrowserInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BrowserConfig) {
		request.BrowserConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BrowserConfig, dara.String("BrowserConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Network) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, dara.String("Network"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodePool) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, dara.String("NodePool"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityPolicy) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, dara.String("SecurityPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoragePolicy) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, dara.String("StoragePolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Timers) {
		request.TimersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Timers, dara.String("Timers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Users) {
		request.UsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Users, dara.String("Users"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppPackageType) {
		body["AppPackageType"] = request.AppPackageType
	}

	if !dara.IsNil(request.AuthNotificationEnabled) {
		body["AuthNotificationEnabled"] = request.AuthNotificationEnabled
	}

	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.BrowserConfigShrink) {
		body["BrowserConfig"] = request.BrowserConfigShrink
	}

	if !dara.IsNil(request.ChargeResourceMode) {
		body["ChargeResourceMode"] = request.ChargeResourceMode
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CloudBrowserName) {
		body["CloudBrowserName"] = request.CloudBrowserName
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MaxAmount) {
		body["MaxAmount"] = request.MaxAmount
	}

	if !dara.IsNil(request.NetworkShrink) {
		body["Network"] = request.NetworkShrink
	}

	if !dara.IsNil(request.NodePoolShrink) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.SecurityPolicyShrink) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !dara.IsNil(request.StoragePolicyShrink) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	if !dara.IsNil(request.SubPayType) {
		body["SubPayType"] = request.SubPayType
	}

	if !dara.IsNil(request.TagShrink) {
		body["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TimersShrink) {
		body["Timers"] = request.TimersShrink
	}

	if !dara.IsNil(request.UserGroupIds) {
		body["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	if !dara.IsNil(request.UsersShrink) {
		body["Users"] = request.UsersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBrowserInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBrowserInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an asynchronous export task to export the WUYING server list, user usage report, cloud browser bookmarks, or URL whitelist.
//
// Description:
//
// This operation creates an export task and returns the task ID and initial status. The exported file is not returned synchronously.
//
// After the task is created, save TaskId and call DescribeExportTasks to query the task status and progress. The recommended polling interval is 5 to 10 seconds. When the status is FINISHED, use the DownloadUrl in the latest query result to download the XLSX file. When the status is FAILED, check ErrorMsg and stop polling.
//
// The task is valid for 24 hours after creation. The download URL has an independent signature validity period of up to 1 hour, which does not exceed the remaining validity period of the task. Query the latest URL before downloading. You cannot obtain a valid download URL after the task expires.
//
// The WUYING server list export only accepts the filter fields listed in QueryParamJson. Pagination is controlled by the server. A single task can export up to 100,000 records. The cloud browser bookmarks and URL whitelist exports export the complete list of the target delivery group. This is not affected by the list search terms or the current page. The ownership and product type of the delivery group are verified before execution. Data is read in pages during execution, and no consistent snapshot is provided. A successful creation does not guarantee that subsequent verification, query, or file generation will succeed.
//
// @param request - CreateExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportTaskResponse
func (client *Client) CreateExportTaskWithContext(ctx context.Context, request *CreateExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExportFields) {
		body["ExportFields"] = request.ExportFields
	}

	if !dara.IsNil(request.QueryParamJson) {
		body["QueryParamJson"] = request.QueryParamJson
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExportTask"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image from a deployed WUYING instance to quickly create more WUYING instances with the same configurations and avoid repeated environment configurations.
//
// @param request - CreateImageByInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageByInstanceResponse
func (client *Client) CreateImageByInstanceWithContext(ctx context.Context, request *CreateImageByInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateImageByInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagList) {
		query["TagList"] = request.TagList
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoCleanUserdata) {
		body["AutoCleanUserdata"] = request.AutoCleanUserdata
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CopyProfile) {
		body["CopyProfile"] = request.CopyProfile
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DiskType) {
		body["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SubInstanceId) {
		body["SubInstanceId"] = request.SubInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageByInstance"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageByInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new image by debugging a delivery group.
//
// @param request - CreateImageFromAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageFromAppInstanceGroupResponse
func (client *Client) CreateImageFromAppInstanceGroupWithContext(ctx context.Context, request *CreateImageFromAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateImageFromAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCenterImageName) {
		body["AppCenterImageName"] = request.AppCenterImageName
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageFromAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageFromAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a provider template for model creation.
//
// Description:
//
// You can create a provider template for model creation under the model templates in the Wuying Agent Management Center to configure the connection information and keys for model services (such as Alibaba Cloud Model Studio, Token Plan, and Moonshot AI) that can be invoked by agents. The created model provider template is automatically associated with the specified model template.
//
// Make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center before you call this operation.
//
// @param request - CreateModelProviderTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelProviderTemplateResponse
func (client *Client) CreateModelProviderTemplateWithContext(ctx context.Context, request *CreateModelProviderTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateModelProviderTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableWuyingProxy) {
		query["EnableWuyingProxy"] = request.EnableWuyingProxy
	}

	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProviderName) {
		query["ProviderName"] = request.ProviderName
	}

	if !dara.IsNil(request.ProviderType) {
		query["ProviderType"] = request.ProviderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelProviderTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelProviderTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a template for model creation.
//
// Description:
//
// You can create a model creation group in the Wuying Agent Management Center to manage the model service providers and model scopes that can be invoked by agents. After creation, you can attach the model creation group to a cloud computer as the inference engine configuration for agents to execute nodes. Ensure that you are fully familiar with the operations and usage of the Wuying Agent Management Center before you call this API operation.
//
// @param request - CreateModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelTemplateResponse
func (client *Client) CreateModelTemplateWithContext(ctx context.Context, request *CreateModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RefScope) {
		query["RefScope"] = request.RefScope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates one or more workstations.
//
// Description:
//
// 1. A project is equivalent to the resource configuration module in the CloudFlow console.
//
// 2. If the ContentId parameter has multiple versions, this operation <notice>uses the default version</notice> for binding.
//
// 3. This operation succeeds only when the default version of the content is in the available state.
//
// @param request - CreateWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWuyingServerResponse
func (client *Client) CreateWuyingServerWithContext(ctx context.Context, request *CreateWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *CreateWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DataDisk) {
		bodyFlat["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.ErdmaEnabled) {
		body["ErdmaEnabled"] = request.ErdmaEnabled
	}

	if !dara.IsNil(request.GpuDriverVersion) {
		body["GpuDriverVersion"] = request.GpuDriverVersion
	}

	if !dara.IsNil(request.HostName) {
		body["HostName"] = request.HostName
	}

	if !dara.IsNil(request.IdempotenceToken) {
		body["IdempotenceToken"] = request.IdempotenceToken
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.MaxPrice) {
		body["MaxPrice"] = request.MaxPrice
	}

	if !dara.IsNil(request.NetworkStrategyType) {
		body["NetworkStrategyType"] = request.NetworkStrategyType
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.SavingPlanId) {
		body["SavingPlanId"] = request.SavingPlanId
	}

	if !dara.IsNil(request.ServerInstanceType) {
		body["ServerInstanceType"] = request.ServerInstanceType
	}

	if !dara.IsNil(request.ServerPortRange) {
		body["ServerPortRange"] = request.ServerPortRange
	}

	if !dara.IsNil(request.SubPayType) {
		body["SubPayType"] = request.SubPayType
	}

	if !dara.IsNil(request.SystemDiskCategory) {
		body["SystemDiskCategory"] = request.SystemDiskCategory
	}

	if !dara.IsNil(request.SystemDiskPerformanceLevel) {
		body["SystemDiskPerformanceLevel"] = request.SystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.VSwitchIds) {
		body["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VirtualNodePoolId) {
		body["VirtualNodePoolId"] = request.VirtualNodePoolId
	}

	if !dara.IsNil(request.WuyingServerName) {
		body["WuyingServerName"] = request.WuyingServerName
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified application adaptation or dissociates the adaptation from a specified image version.
//
// Description:
//
// *Confirm the deletion scope first.**
//
// - If both ImageId and VersionId are specified: only dissociates the adaptation from the specified image version.
//
// - If ImageId and VersionId are not both specified: initiates the adaptation deletion process. The adaptation cannot be deleted if it is still referenced by multiple images. If the application has other adaptations, it cannot be deleted together with the application.
//
// **AdaptorId, AppVersionId, and AppId must be provided at the same time and correspond to the same application adaptation.*	- Verify these identifiers through the application and image query results first.
//
// @param request - DeleteAdaptorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdaptorResponse
func (client *Client) DeleteAdaptorWithContext(ctx context.Context, request *DeleteAdaptorRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdaptorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdaptorId) {
		body["AdaptorId"] = request.AdaptorId
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersionId) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdaptor"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdaptorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified application version. If the version ID is omitted, deletes the entire application and its adapters.
//
// Description:
//
// *You must disassociate all images from the target application or version before deletion.**
//
// 1. Call ListApp to obtain the AppId and confirm whether the application can be deleted.
//
// 2. Use UnbindApp or DeleteAdaptor to disassociate the related images.
//
// 3. Specify AppVersionId to delete a single application version, or omit this parameter to delete the entire application.
//
// After an application version is deleted, if the application has no other valid adapters, the application record is also deleted.
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersionId) {
		body["AppVersionId"] = request.AppVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource-based pay-as-you-go delivery group.
//
// Description:
//
// > This operation does not support deleting delivery groups that use subscription resources.
//
// @param request - DeleteAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstanceGroupResponse
func (client *Client) DeleteAppInstanceGroupWithContext(ctx context.Context, request *DeleteAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specified application instances.
//
// Description:
//
// Only instances in the initializing or idle state can be deleted. This operation is available only to specific customers.
//
// @param request - DeleteAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstancesResponse
func (client *Client) DeleteAppInstancesWithContext(ctx context.Context, request *DeleteAppInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceIds) {
		body["AppInstanceIds"] = request.AppInstanceIds
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppInstances"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom WUYING image.
//
// Description:
//
// - You can only delete custom images that you own.
//
// - For images used in the WUYING Cloud Computer Pool, WUYING Cloud App, and WUYING Workstation product lines, ensure that no WUYING instances are using the image before you delete it.
//
// - If a WUYING cloud desktop template references the image, the template is also deleted when you delete the image.
//
// - If the image is available in multiple regions, deleting the image removes it from all regions.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an LLM template.
//
// Description:
//
// You can delete a model template created under a model service provider template in the Wuying Agent Management Center. Before deletion, make sure that the model is not the default model of an associated model group. Otherwise, the deletion fails. After the model template is deleted, the model configurations of the associated cloud desktops are automatically refreshed.
//
// Before you call this operation, make sure that you fully understand the operations and usage of the Wuying Agent Management Center.
//
// @param request - DeleteLlmTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLlmTemplateResponse
func (client *Client) DeleteLlmTemplateWithContext(ctx context.Context, request *DeleteLlmTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteLlmTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LlmTemplateId) {
		query["LlmTemplateId"] = request.LlmTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLlmTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLlmTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model provider template.
//
// Description:
//
// You can delete a model provider template created under model templates in the Wuying Agent Management Center. Before deletion, ensure that the model provider is not the provider of the default model and is not a system-preset provider (such as the Wuying credit package). After deletion, the associated models and key configurations are also removed.
//
// Before you call this operation, make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - DeleteModelProviderTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelProviderTemplateResponse
func (client *Client) DeleteModelProviderTemplateWithContext(ctx context.Context, request *DeleteModelProviderTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelProviderTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProviderTemplateId) {
		query["ProviderTemplateId"] = request.ProviderTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelProviderTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelProviderTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model template.
//
// Description:
//
// You can delete a model group created in the Wuying Agent Management Center. Before deletion, ensure that the template is not authorized for any resources. Otherwise, the deletion fails. After the model group is deleted, the model providers and models in the model group are also removed.
//
// Before you use this API operation, ensure that you fully understand the operations and usage of the Wuying Agent Management Center.
//
// @param request - DeleteModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelTemplateResponse
func (client *Client) DeleteModelTemplateWithContext(ctx context.Context, request *DeleteModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workstation.
//
// Description:
//
// Deletes a workstation.
//
// @param request - DeleteWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWuyingServerResponse
func (client *Client) DeleteWuyingServerWithContext(ctx context.Context, request *DeleteWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *DeleteWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures operation log delivery for specified products and writes the logs to a Simple Log Service (SLS) project and Logstore of the current account.
//
// Description:
//
// ## Before you begin
//
// - Specify the region where SLS is deployed and at least one product delivery scope.
//
// - Specify either `ProjectName` or `ExistedProjectName`. If you use an existing project, make sure that it belongs to the current account and is located in the specified region.
//
// - Only one enabled delivery configuration can exist for the same account and product type. Before you call this operation, you can call `DescribeSlsDeliveryConfig` to query existing configurations. If an enabled configuration already exists, calling this operation again returns `SlsDeliveryAlreadyExists`.
//
// ## Configuration management
//
// - This operation is used to configure delivery, not to query logs.
//
// - To pause delivery, call `StopDeliveryConfig`. Stopping delivery does not delete the delivery configuration.
//
// - To reconfigure the delivery destination or rebuild the configuration, query and delete the corresponding old configuration by calling `DeleteDeliveryConfig` before you call this operation. Deleting a delivery configuration does not delete the SLS project or Logstore.
//
// - Do not treat repeated calls for an existing configuration as update or resume operations. After the call, you can call `DescribeSlsDeliveryConfig` to confirm the configuration status.
//
// @param request - DeliverToUserSlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeliverToUserSlsResponse
func (client *Client) DeliverToUserSlsWithContext(ctx context.Context, request *DeliverToUserSlsRequest, runtime *dara.RuntimeOptions) (_result *DeliverToUserSlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryScopes) {
		bodyFlat["DeliveryScopes"] = request.DeliveryScopes
	}

	if !dara.IsNil(request.ExistedProjectName) {
		body["ExistedProjectName"] = request.ExistedProjectName
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SlsRegionId) {
		body["SlsRegionId"] = request.SlsRegionId
	}

	if !dara.IsNil(request.Ttl) {
		body["Ttl"] = request.Ttl
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeliverToUserSls"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeliverToUserSlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of the resource and session events of the current account within the last 30 days, supporting paging.
//
// Description:
//
// Queries the resource and session events of the current account. `StartTime` and `EndTime` must be specified at the same time, and the time range includes the start and end times. The service only retrieves logs from 30 days before the call time to the current time, so events outside this window are not returned. If multiple filter parameters are specified, the intersection is used. Results are returned in descending order of the event log time. To use page number-based pagination, specify `PageSize` and `PageNumber`. The current version does not use `MaxResults` and `NextToken`.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventType) {
		body["EventType"] = request.EventType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberEniIp) {
		body["MemberEniIp"] = request.MemberEniIp
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status, progress, and download URLs of export tasks by page.
//
// Description:
//
// Calls this operation to query the export tasks under the current Alibaba Cloud account by page to obtain the task status, execution progress, export file names, and download URLs. Export tasks are created by CreateExportTask. This operation only reads task information and does not create, modify, or retry tasks.
//
// The query results are returned in descending order of task creation time, with the most recently created tasks listed first. Only tasks in the INIT, RUNNING, FINISHED, and FAILED states are returned. Tasks marked as EXPIRED are not included in the results.
//
// When polling a single task by TaskId, set PageNumber to 1 and PageSize to 1, set the polling interval to 5 to 10 seconds, and stop polling after the task enters the FINISHED or FAILED state.
//
// A task is valid for 24 hours after creation and is marked as expired by a background task after 25 hours. Therefore, within the 24- to 25-hour window, a FINISHED task may still be queried, but its DownloadUrl is empty. After a task is marked as expired, the query result is an empty array.
//
// If the query result is an empty array, it may be because there are no visible tasks under the current account, the task type does not match, the page number is out of bounds, or the tasks have expired. This operation always returns an array structure and does not return an error when the query result is empty.
//
// The DownloadUrl is an OSS pre-signed URL. The default maximum validity period is 1 hour and does not exceed the remaining validity period of the task. The server automatically refreshes the URL when it is about to expire. Use the URL returned by the most recent query for downloading and do not cache the initially returned URL for a long time.
//
// @param request - DescribeExportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportTasksResponse
func (client *Client) DescribeExportTasksWithContext(ctx context.Context, request *DescribeExportTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeExportTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExportTasks"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Alibaba Cloud accounts with which the current image is shared.
//
// @param request - DescribeShareImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeShareImageResponse
func (client *Client) DescribeShareImageWithContext(ctx context.Context, request *DescribeShareImageRequest, runtime *dara.RuntimeOptions) (_result *DescribeShareImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeShareImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeShareImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Simple Log Service (SLS) delivery configurations of the current account.
//
// Description:
//
// ## Query description
//
// - Queries existing delivery configurations under the current account. This operation does not create or modify configurations.
//
// - You can filter configurations by SLS region, project name, Logstore name, and product type. If multiple conditions are specified, they take effect together.
//
// - If no filter conditions are specified, all delivery configurations under the current account are queried. If no configurations match the conditions, an empty list is returned.
//
// - The returned results are displayed hierarchically by project, Logstore, and delivery scope.
//
// ## Example description
//
// The examples of names and retention periods are used to demonstrate the format. Refer to the actual configurations. The retention period examples do not represent default values or upper limits.
//
// @param request - DescribeSlsDeliveryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsDeliveryConfigResponse
func (client *Client) DescribeSlsDeliveryConfigWithContext(ctx context.Context, request *DescribeSlsDeliveryConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsDeliveryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryScopes) {
		bodyFlat["DeliveryScopes"] = request.DeliveryScopes
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsDeliveryConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsDeliveryConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the system default path rules, or the user profile roaming rules for a specified app instance group or app instance group set.
//
// Description:
//
// When querying a specified target, specify either the app instance group ID or the app instance group set ID. When querying a set, the rules of the primary app instance group are returned. You can use the query results as a reference for modifying rules, and then call SetUserProfilePathRules to submit the complete rule list.
//
// @param request - DescribeUserProfilePathRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserProfilePathRulesResponse
func (client *Client) DescribeUserProfilePathRulesWithContext(ctx context.Context, request *DescribeUserProfilePathRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserProfilePathRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserProfilePathRules"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserProfilePathRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a development host.
//
// @param request - DescribeWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWuyingServerResponse
func (client *Client) DescribeWuyingServerWithContext(ctx context.Context, request *DescribeWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *DescribeWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the elastic IP address (EIP) information of a Wuying workstation.
//
// @param request - DescribeWuyingServerEipInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWuyingServerEipInfoResponse
func (client *Client) DescribeWuyingServerEipInfoWithContext(ctx context.Context, request *DescribeWuyingServerEipInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeWuyingServerEipInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Isp) {
		body["Isp"] = request.Isp
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWuyingServerEipInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWuyingServerEipInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms the successful verification of an image or application build.
//
// @param request - FinishValidateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishValidateResponse
func (client *Client) FinishValidateWithContext(ctx context.Context, request *FinishValidateRequest, runtime *dara.RuntimeOptions) (_result *FinishValidateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersionId) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishValidate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of runtimes of a specified type under the current Alibaba Cloud account, grouped by agent platform and agent provider. The statistics include the total number of runtimes, the number of runtimes without a configured model, the number of runtimes without a configured channel, the number of risky runtimes, and the number of expired runtimes.
//
// @param request - GetAgentRuntimeStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeStatisticResponse
func (client *Client) GetAgentRuntimeStatisticWithContext(ctx context.Context, request *GetAgentRuntimeStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.RuntimeType) {
		query["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentRuntimeStatistic"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentRuntimeStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified delivery group.
//
// @param request - GetAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceGroupResponse
func (client *Client) GetAppInstanceGroupWithContext(ctx context.Context, request *GetAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the connection authorization code and required resource information for an existing image or application creation task.
//
// Description:
//
// Before calling this operation, you must create a task, and the creation or validation resources corresponding to the task must be available.
//
// 1. Obtain the TaskId and actual TaskType from the response of the task creation operation or GetTask.
//
// 2. Specify ResourceType to obtain the connection information for the creation or validation environment.
//
// 3. Use the returned authorization code together with the user, resource group, and application identifiers in the same response for connection.
//
// The authorization code is a sensitive connection credential. Do not record it in public documents or logs. When specifying AppId, you must also specify the matching AppVersionId.
//
// @param request - GetAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthCodeResponse
func (client *Client) GetAuthCodeWithContext(ctx context.Context, request *GetAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *GetAuthCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersionId) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthCode"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration, status, and authorization statistics of a specified cloud browser group.
//
// Description:
//
// This topic describes the query usage for the monthly active user (MAU) billing scenario.
//
// ## Before you begin
//
// Obtain the ID of the cloud browser group within the current account. You can call the ListBrowserInstanceGroup operation to obtain the ID.
//
// ## Response content
//
// Returns the current configuration, status, and authorization statistics of the browser group. The bookmark and website access lists in the details return up to 20 entries each. To query the complete lists, call the ListBrowserBookmarks and ListBrowserRestrictedURLs operations, respectively.
//
// ## What to do next
//
// This operation only queries configurations and does not modify resources. After reading the returned status, perform connection or management operations that match the status.
//
// ## Example description
//
// In the examples, a hyphen (-) indicates that the field is not applicable to the current scenario or is not returned. It is not an actual string returned by the operation. The examples of common resource IDs are masked. Use the actual query results.
//
// @param request - GetBrowserInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBrowserInstanceGroupResponse
func (client *Client) GetBrowserInstanceGroupWithContext(ctx context.Context, request *GetBrowserInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetBrowserInstanceGroupResponse, _err error) {
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
		Action:      dara.String("GetBrowserInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBrowserInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the connection credential for a cloud application.
//
// Description:
//
// Call protocol description: operation_type: polling, required_steps: 1.
//
// This operation may require multiple calls (at least one) to obtain the connection credential.
//
// On the first call, an application instance is allocated to the specified convenience account and the application is started. If a Ticket is returned, the result is obtained synchronously. If a startup task ID (`TaskId`) is returned, subsequent calls are required.
//
// On subsequent calls, include the `TaskId` request parameter to invoke the operation and query whether the node is complete. When the returned node status (`TaskStatus`) is completed (`Finished`), the connection credential (`Ticket`) is also returned.
//
// > Prerequisites
//
// > - Before calling this operation, make sure that you have created a delivery group and authorized users for the delivery group:
//
// > - 1. The API for creating a delivery group is CreateAppInstanceGroup. For more information about the parameters, see the corresponding API documentation.
//
// > - 2. You can call the ListAppInstanceGroup operation to query the list of delivery groups. If the corresponding delivery group is not found, verify that the delivery group has been created and that the authentication credentials belong to the correct tenant.
//
// > - 3. The API for authorizing users for a delivery group is AuthorizeInstanceGroup. For more information about the parameters, see the corresponding API documentation.
//
// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithContext(ctx context.Context, request *GetConnectionTicketRequest, runtime *dara.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		body["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupIdList) {
		body["AppInstanceGroupIdList"] = request.AppInstanceGroupIdList
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !dara.IsNil(request.AppPolicyId) {
		body["AppPolicyId"] = request.AppPolicyId
	}

	if !dara.IsNil(request.AppStartParam) {
		body["AppStartParam"] = request.AppStartParam
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EnvironmentConfig) {
		body["EnvironmentConfig"] = request.EnvironmentConfig
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnectionTicket"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information about a debug application instance.
//
// @param request - GetDebugAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDebugAppInstanceResponse
func (client *Client) GetDebugAppInstanceWithContext(ctx context.Context, request *GetDebugAppInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetDebugAppInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDebugAppInstance"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDebugAppInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the effective versions of an image in various regions, along with its applications, distribution configurations, and tags.
//
// Description:
//
// Queries the effective versions of an image in various regions by ImageId. If a image version is configured in the same region, the image version is returned first. You can specify BizRegionId to limit the query to a specific region. The returned results are also constrained by the product type and the regions visible to the account. To query the details of a specific version, call GetImageById.
//
// @param request - GetEffectiveImageByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEffectiveImageByIdResponse
func (client *Client) GetEffectiveImageByIdWithContext(ctx context.Context, request *GetEffectiveImageByIdRequest, runtime *dara.RuntimeOptions) (_result *GetEffectiveImageByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ContentLocale) {
		body["ContentLocale"] = request.ContentLocale
	}

	if !dara.IsNil(request.IgnoreFilter) {
		body["IgnoreFilter"] = request.IgnoreFilter
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		body["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ShowHidden) {
		body["ShowHidden"] = request.ShowHidden
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEffectiveImageById"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEffectiveImageByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves file upload information.
//
// Description:
//
// Visitor information is filled in by visitors on the lead capture page when they execute a cloud flow. Therefore, the usage of cloud applications does not generate visitor information.
//
// @param request - GetFileUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileUploadInfoResponse
func (client *Client) GetFileUploadInfoWithContext(ctx context.Context, request *GetFileUploadInfoRequest, runtime *dara.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileUploadInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the image details, adapted applications, and region distribution information based on the image ID or version ID.
//
// Description:
//
// You must specify at least one of ImageId and VersionId. If only ImageId is specified, the online version is queried. If VersionId is specified, the specified version is queried. If both parameters are specified, VersionId must belong to the specified image. To query the effective version in each region, call GetEffectiveImageById.
//
// @param request - GetImageByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageByIdResponse
func (client *Client) GetImageByIdWithContext(ctx context.Context, request *GetImageByIdRequest, runtime *dara.RuntimeOptions) (_result *GetImageByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.IgnoreFilter) {
		body["IgnoreFilter"] = request.IgnoreFilter
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ShowHidden) {
		body["ShowHidden"] = request.ShowHidden
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageById"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution regions, distribution status, and progress of an image version.
//
// Description:
//
// Queries the region distribution status of an image by ImageId. If VersionId is not specified, the online version is queried. If VersionId is specified, the specified version is queried. Returns the non-deleted region distribution records and the list of corresponding region IDs.
//
// @param request - GetImageDistributeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDistributeInfoResponse
func (client *Client) GetImageDistributeInfoWithContext(ctx context.Context, request *GetImageDistributeInfoRequest, runtime *dara.RuntimeOptions) (_result *GetImageDistributeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDistributeInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDistributeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a model provider template.
//
// Description:
//
// You can query the details of a specified model provider template in the Wuying Agent Management Center, including the provider name, description, and connection configuration list.
//
// Before you call this operation, make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - GetModelProviderTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelProviderTemplateResponse
func (client *Client) GetModelProviderTemplateWithContext(ctx context.Context, request *GetModelProviderTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetModelProviderTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProviderTemplateId) {
		query["ProviderTemplateId"] = request.ProviderTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelProviderTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelProviderTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the authentication credentials, access endpoints, and topic list required for a client to establish an MQTT connection.
//
// Description:
//
// Call this operation to retrieve the credentials and connection information required for a client to establish an MQTT connection.
//
// 1. Use `MqttUsername` and `MqttPassword` as the username and password when connecting the MQTT client.
//
// 2. Use `MqttEndpoint` as the connection endpoint. If the client meets the required VPC network access conditions, you can use `MqttVPCEndpoint`. `MqttInstanceId` is used for token authentication, and `MqttGroupId` is used to concatenate the MQTT ClientId.
//
// 3. Distinguish between publish and subscribe topics based on the `topic`, `type`, and `permission` fields of each item in `TopicList`. `W` indicates publish permission, and `R` indicates subscribe permission.
//
// 4. `ExpireTimeStamp` is the UNIX timestamp in milliseconds when the password expires. Credentials may be reused. Retrieve new credentials based on the expiration time returned in the response.
//
// The device identity and request IDs in the documentation are demo data. The domain name uses the `example.com` sample domain. Credentials and instance identity are sample or desensitized values. These examples cannot be used directly for actual connections.
//
// @param tmpReq - GetMqttTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMqttTicketResponse
func (client *Client) GetMqttTicketWithContext(ctx context.Context, tmpReq *GetMqttTicketRequest, runtime *dara.RuntimeOptions) (_result *GetMqttTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMqttTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.PropertiesShrink) {
		body["Properties"] = request.PropertiesShrink
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMqttTicket"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMqttTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an over-the-air update task, including the available version and release notes.
//
// @param request - GetOtaTaskByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOtaTaskByTaskIdResponse
func (client *Client) GetOtaTaskByTaskIdWithContext(ctx context.Context, request *GetOtaTaskByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOtaTaskByTaskId"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the model settings of resource groups.
//
// @param request - GetResourceGroupAgentModelSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupAgentModelSettingsResponse
func (client *Client) GetResourceGroupAgentModelSettingsWithContext(ctx context.Context, request *GetResourceGroupAgentModelSettingsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupAgentModelSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupAgentModelSettings"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupAgentModelSettingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price and discount information required to scale out a subscription delivery group to a specified total number of nodes.
//
// Description:
//
// This operation queries the resource scale-out price for an existing delivery group. A successful call returns only a quote and does not create an order or add nodes.
//
// **Before you call this operation, verify the following conditions:**
//
// - The delivery group belongs to the current account and is located in a supported region in the Chinese mainland.
//
// - The product type is WUYING Cloud App or Cloud Browser, and the billing method is subscription and purchased by node. This operation does not apply to browser groups that use the pay-as-you-go billing method or are billed by monthly active users (MAU).
//
// - The order status of the delivery group is normal, a resource node pool is associated, and no paid but ineffective renewal or scale-in orders exist.
//
// - The target total number of nodes is greater than the current number of purchased nodes and does not exceed the node quota for the account in the corresponding region, product, and billing method.
//
// **Parameter relationships:**
//
// - ProductType, AppInstanceGroupId, and Amount are all required parameters and must be specified at the same time.
//
// - ProductType must match the product corresponding to AppInstanceGroupId.
//
// - **Amount is the total number of nodes after the scale-out, not the number of nodes added in this operation.*	- The region, node specifications, and billing method are determined by the existing delivery group and do not need to be passed separately.
//
// **Call sequence:**
//
// 1. Confirm the target delivery group and the current number of purchased nodes. For cloud app Common scenarios, you can first invoke [ListAppInstanceGroup](https://api.aliyun.com/document/appstream-center/2021-09-01/ListAppInstanceGroup) to obtain the identity, and then invoke [GetAppInstanceGroup](https://api.aliyun.com/document/appstream-center/2021-09-01/GetAppInstanceGroup) to query the configuration and NodePool.NodeAmount.
//
// 2. Invoke this operation with the same product type, delivery group identity, and target total number of nodes to read the price, coupons, and price calculation rules.
//
// 3. After confirming the quote, if you need to execute the scale-out, you can invoke [ModifyNodePoolAmount](https://api.aliyun.com/document/appstream-center/2021-09-01/ModifyNodePoolAmount) and fill in the target total number in NodePool.NodeAmount as required by the operation. **A successful price inquiry does not mean that an order has been created or the scale-out has been completed.**
//
// **Example scenario:**
//
// If a cloud app delivery group currently has 1 purchased subscription node and the target is to scale out to 2 nodes, set ProductType to CloudApp, AppInstanceGroupId to the ID of the delivery group, and Amount to 2. Replace the masked resource ID with the actual value under the current account. The example quantity does not represent a default value or quota limit. The price example only illustrates the relationship between the original price, discount amount, and discounted price, and does not represent the fixed selling price of the specification.
//
// @param request - GetResourceModifyPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceModifyPriceResponse
func (client *Client) GetResourceModifyPriceWithContext(ctx context.Context, request *GetResourceModifyPriceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceModifyPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceModifyPrice"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceModifyPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price information of a resource.
//
// @param request - GetResourcePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcePriceResponse
func (client *Client) GetResourcePriceWithContext(ctx context.Context, request *GetResourcePriceRequest, runtime *dara.RuntimeOptions) (_result *GetResourcePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AppInstanceType) {
		query["AppInstanceType"] = request.AppInstanceType
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.NodeInstanceType) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourcePrice"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum quantity quota applicable to the current account for a specified product, resource type, business region, and billing method.
//
// Description:
//
// This operation queries the maximum quantity quota. It does not query the used quota, remaining quota, or real-time resource inventory.
//
// **Parameter relationships:**
//
// - `ResourceType` and `ProductType` are required.
//
// - `BizRegionId` and `ChargeType` are optional query conditions and can be specified at the same time. They must be consistent with subsequent purchase or authorization operations.
//
// - If the region or billing method is omitted, the general quota for the corresponding dimension is queried. The quotas of different regions or billing methods are not added together.
//
// - If multiple quotas apply at the same time, the larger limit is returned.
//
// **Call sequence:**
//
// 1. Select the resource type, product, business region, and billing method of the payer based on subsequent business needs.
//
// 2. Invoke this operation to read `Data`. Query `AppInstanceGroup` before creating a delivery group, query `AppAuthUser` before authorizing an authorized user, and query `PoolNodeAmount` before purchasing or scaling out nodes.
//
// 3. Query the current usage in combination with resources, and then verify the current operation through subsequent creation, authorization, or scale-out operations. **This query does not reserve quotas or guarantee the success of subsequent purchases.**
//
// `Data=2147483647` indicates that no smaller quota limit is obtained in this query. It does not mean that the physical inventory is unlimited.
//
// @param request - GetResourceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceQuotaResponse
func (client *Client) GetResourceQuotaWithContext(ctx context.Context, request *GetResourceQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetResourceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceQuota"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the renewal price of WUYING Cloud Application resources.
//
// @param request - GetResourceRenewPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceRenewPriceResponse
func (client *Client) GetResourceRenewPriceWithContext(ctx context.Context, request *GetResourceRenewPriceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceRenewPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceRenewPrice"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceRenewPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the third-party channel configurations for agent runtimes.
//
// Description:
//
// You can query the third-party channel configuration status for agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center.
//
// Before calling this operation, make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - GetRuntimeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuntimeChannelResponse
func (client *Client) GetRuntimeChannelWithContext(ctx context.Context, request *GetRuntimeChannelRequest, runtime *dara.RuntimeOptions) (_result *GetRuntimeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.IncludeRiskInfo) {
		query["IncludeRiskInfo"] = request.IncludeRiskInfo
	}

	if !dara.IsNil(request.RuntimeId) {
		query["RuntimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.RuntimeType) {
		query["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuntimeChannel"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuntimeChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the QR code configuration process for a specific channel on a specified runtime, including the process status, QR code content, and the URL of the QR code configuration page.
//
// @param request - GetRuntimeChannelQrCodeDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuntimeChannelQrCodeDetailResponse
func (client *Client) GetRuntimeChannelQrCodeDetailWithContext(ctx context.Context, request *GetRuntimeChannelQrCodeDetailRequest, runtime *dara.RuntimeOptions) (_result *GetRuntimeChannelQrCodeDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.RuntimeId) {
		query["RuntimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.RuntimeType) {
		query["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuntimeChannelQrCodeDetail"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuntimeChannelQrCodeDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the model configuration details of a cloud desktop.
//
// Description:
//
// You can query the details of the model configuration currently bound to a specified cloud desktop in the Wuying Agent Management Center, including the model group, model provider list, and associated model information. When the risk information mode is enabled, you can also identify discrepancies between the end user\\"s actual configuration and the configuration delivered by the administrator.
//
// @param request - GetRuntimeModelConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuntimeModelConfigResponse
func (client *Client) GetRuntimeModelConfigWithContext(ctx context.Context, request *GetRuntimeModelConfigRequest, runtime *dara.RuntimeOptions) (_result *GetRuntimeModelConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.IncludeRiskInfo) {
		query["IncludeRiskInfo"] = request.IncludeRiskInfo
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.RuntimeId) {
		query["RuntimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.RuntimeType) {
		query["RuntimeType"] = request.RuntimeType
	}

	if !dara.IsNil(request.SmartModel) {
		query["SmartModel"] = request.SmartModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuntimeModelConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuntimeModelConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a task by task ID, or by resource information and task category.
//
// Description:
//
// Supports two query methods:
//
// - **By task ID**: Pass TaskId. This method takes precedence over resource information queries.
//
// - **By resource**: Omit TaskId, and pass EntityId, EntityType, and TaskCategory. You can use EntityVersionId to specify the resource version.
//
// The API returns the current task status and recorded step information. Determine the overall task result based on Status. The success of a single step does not mean the entire task is successful. You can call ListTaskStep to query the names and categories of image creation and application adaptation steps.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.EntityVersionId) {
		body["EntityVersionId"] = request.EntityVersionId
	}

	if !dara.IsNil(request.TaskCategory) {
		body["TaskCategory"] = request.TaskCategory
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the tenant-level memory configurations of the current account by agent provider, business platform, and business type.
//
// Description:
//
// Specify the agent provider, business platform, and business type corresponding to the configuration. After modification, you can use the same combination to query whether the configuration has been updated.
//
// @param request - GetTenantMemoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTenantMemoryConfigResponse
func (client *Client) GetTenantMemoryConfigWithContext(ctx context.Context, request *GetTenantMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *GetTenantMemoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentInstanceId) {
		body["AgentInstanceId"] = request.AgentInstanceId
	}

	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTenantMemoryConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTenantMemoryConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the OSS endpoint, directory, and STS temporary credentials required to upload image task files.
//
// Description:
//
// FileType currently supports only TASK. Use the returned OSS endpoint, bucket, file directory, and temporary credentials to upload files. After the file is uploaded, you can pass the complete OSS object path to ImportImage. Expiration is not supported.
//
// @param request - GetUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadInfoResponse
func (client *Client) GetUploadInfoWithContext(ctx context.Context, request *GetUploadInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUploadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports an image file from Object Storage Service (OSS) as a custom Cloud Desktop image and returns the image ID and version ID.
//
// Description:
//
// Currently, only ProductType=CloudDesktop is supported. Upload the image file first, and then pass its OSS object path to initiate the import. You can call GetUploadInfo to obtain upload credentials. After the API returns the image and version IDs, you can call GetImageById or GetImageDistributeInfo to query the creation and distribution status. A successful response indicates that the import request has been accepted, not that the image is already available.
//
// @param request - ImportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportImageResponse
func (client *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest, runtime *dara.RuntimeOptions) (_result *ImportImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BootMode) {
		body["BootMode"] = request.BootMode
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableSecurityCheck) {
		body["EnableSecurityCheck"] = request.EnableSecurityCheck
	}

	if !dara.IsNil(request.GpuCategory) {
		body["GpuCategory"] = request.GpuCategory
	}

	if !dara.IsNil(request.GpuDriverType) {
		body["GpuDriverType"] = request.GpuDriverType
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.LicenseType) {
		body["LicenseType"] = request.LicenseType
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OssObjectPath) {
		body["OssObjectPath"] = request.OssObjectPath
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ProtocolType) {
		body["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the memory configurations, memory statistics, and backup summaries of agent instances for the current account in a paged manner.
//
// Description:
//
// Filter by Agent provider, business platform, business type, end user, and instance memory switch as needed. Retain the filter conditions and increment the page number when navigating to the next page.
//
// @param request - ListAgentInstanceMemoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentInstanceMemoryResponse
func (client *Client) ListAgentInstanceMemoryWithContext(ctx context.Context, request *ListAgentInstanceMemoryRequest, runtime *dara.RuntimeOptions) (_result *ListAgentInstanceMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.ExcludedEndUserIds) {
		query["ExcludedEndUserIds"] = request.ExcludedEndUserIds
	}

	if !dara.IsNil(request.MemoryStatuses) {
		query["MemoryStatuses"] = request.MemoryStatuses
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyword) {
		query["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SearchType) {
		query["SearchType"] = request.SearchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentInstanceMemory"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentInstanceMemoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a deduplicated list of version numbers of undeleted agent instances under a specified tenant, aggregated by agent provider and optionally filtered by agent platform. The results are returned in ascending order of version numbers.
//
// @param request - ListAgentInstanceVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentInstanceVersionsResponse
func (client *Client) ListAgentInstanceVersionsWithContext(ctx context.Context, request *ListAgentInstanceVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentInstanceVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentInstanceVersions"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentInstanceVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries valid applications of the current account by using paging based on conditions such as the application ID, name, type, and associated product.
//
// Description:
//
// When multiple filter conditions are specified, they take effect together. AppName supports contains matching. If a single-value filter and the corresponding list filter are specified at the same time, the intersection is used. To avoid empty results, specify only one of them.
//
// If PageSize is omitted, the default value is 10. We recommend that you start from page 1 and query page by page until the number of retrieved records reaches the value of Count. Before you delete an application, check the Deletable field.
//
// @param tmpReq - ListAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppResponse
func (client *Client) ListAppWithContext(ctx context.Context, tmpReq *ListAppRequest, runtime *dara.RuntimeOptions) (_result *ListAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizTypeList) {
		request.BizTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizTypeList, dara.String("BizTypeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProductTypeList) {
		request.ProductTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductTypeList, dara.String("ProductTypeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BizTypeListShrink) {
		body["BizTypeList"] = request.BizTypeListShrink
	}

	if !dara.IsNil(request.OriginAppType) {
		body["OriginAppType"] = request.OriginAppType
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ProductTypeListShrink) {
		body["ProductTypeList"] = request.ProductTypeListShrink
	}

	if !dara.IsNil(request.SubAppType) {
		body["SubAppType"] = request.SubAppType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApp"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of multiple app instance groups that meet the specified conditions.
//
// @param request - ListAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstanceGroupResponse
func (client *Client) ListAppInstanceGroupWithContext(ctx context.Context, request *ListAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCenterImageId) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.NodeInstanceType) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludedUserGroupIds) {
		body["ExcludedUserGroupIds"] = request.ExcludedUserGroupIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.UserGroupIds) {
		body["UserGroupIds"] = request.UserGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query on the available session specifications for the current account based on the product, region, operating system, and specification conditions, returning results through paging.
//
// Description:
//
// This operation is used to select session specifications and does not create resources. The returned scope is affected by the available specifications and business region of the account. An empty result does not indicate a request failure.
//
// **Filtering and pagination:**
//
// - `ProductType` is required. Explicitly specify `BizRegionId`, `PageNumber`, and `PageSize`.
//
// - If `BizRegionId` is not provided, the region of the current request is used. You can call [ListRegions](https://api.aliyun.com/document/appstream-center/2021-09-01/ListRegions) to obtain the business region.
//
// - You can specify both `AppInstanceType` and `AppInstanceTypeList`. In this case, only the specifications that meet both conditions are returned. No priority override exists. The list can contain up to 100 items.
//
// - The specification family, CPU, memory, GPU memory, operating system, and category conditions are filtered by intersection. The union is not automatically applied.
//
// - Pass `PageNumber` starting from 1. Use `TotalCount`, `PageSize`, and `PageNumber` to continue pagination until all matching results are read.
//
// **Inventory and subsequent calls:**
//
// - Inventory is queried only when `QueryInventory=true`. If inventory information is not obtained, you cannot determine that the inventory is 0.
//
// - The inventory is the current query result and does not represent a successful reservation.
//
// - After obtaining `AppInstanceType`, you can query the price based on the session specification parameter requirements of [GetResourcePrice](https://api.aliyun.com/document/appstream-center/2021-09-01/GetResourcePrice).
//
// - The current query does not use `AppCenterImageId` to filter specifications. Do not use this parameter to determine the compatibility between images and specifications.
//
// Field examples are used to demonstrate how to fill in the fields. The existence of example specifications in the selected region or account is not guaranteed. The GPU count field does not currently provide values. Use the GPU memory field to understand the specifications.
//
// @param request - ListAppInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstanceTypeResponse
func (client *Client) ListAppInstanceTypeWithContext(ctx context.Context, request *ListAppInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCenterImageId) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !dara.IsNil(request.AppInstanceType) {
		query["AppInstanceType"] = request.AppInstanceType
	}

	if !dara.IsNil(request.AppInstanceTypeFamily) {
		query["AppInstanceTypeFamily"] = request.AppInstanceTypeFamily
	}

	if !dara.IsNil(request.AppInstanceTypeList) {
		query["AppInstanceTypeList"] = request.AppInstanceTypeList
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.GpuMemory) {
		query["GpuMemory"] = request.GpuMemory
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.QueryInventory) {
		query["QueryInventory"] = request.QueryInventory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstanceType"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstanceTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of session instances in a delivery group, including instance IDs, instance statuses, creation time, update time, session statuses, and public IP addresses of primary network interface controllers (NICs).
//
// @param request - ListAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstancesResponse
func (client *Client) ListAppInstancesWithContext(ctx context.Context, request *ListAppInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceId) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.IncludeDeleted) {
		query["IncludeDeleted"] = request.IncludeDeleted
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceIdList) {
		body["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstances"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployed applications in the image used by a specified delivery group with paging and returns the number of users with per-application authorization for each application.
//
// Description:
//
// ## Operation description
//
// This operation returns the list of deployed applications in the application image used by a specified delivery group, including the application ID, name, version, icon, and the number of users currently **authorized by application*	- for each application (AuthorizedUserCount).
//
// The returned AppId is the input for per-application authorization: when you call the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation to add or remove authorized users for a specified application in a delivery group, pass in the AppId returned by this operation.
//
// ## Before you begin
//
// - The delivery group is created, and **ProductType matches the product type of the delivery group**. If the delivery group does not exist or the product type does not match, the error code `InvalidAppInstanceGroup.NotFound` is returned.
//
// ## Parameter description
//
// - **AppInstanceGroupId is required**. This parameter is marked as optional in the parameter table, but the error code `InvalidParameter.AppInstanceGroupId` is returned if it is not specified.
//
// - PageNumber starts from 1. Valid values of PageSize: 1 to 100. If the values are invalid, the error codes `InvalidParameter.PageNumber` and `InvalidParameter.PageSize` are returned respectively.
//
// - If no applications are deployed in the delivery group image, the operation returns normally: Apps is an empty list and TotalCount is 0.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the delivery group ID (AppInstanceGroupId).
//
// 2. Call this operation to obtain the list of deployed applications in the delivery group and the AppId of each application.
//
// 3. To authorize by application, call the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation with the returned AppId.
//
// @param request - ListAppsByAppInstanceGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsByAppInstanceGroupIdResponse
func (client *Client) ListAppsByAppInstanceGroupIdWithContext(ctx context.Context, request *ListAppsByAppInstanceGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListAppsByAppInstanceGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppsByAppInstanceGroupId"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppsByAppInstanceGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query on the app instance groups for which a specified user has obtained access permissions through app instance group-level authorization, and supports fuzzy filtering by app instance group ID, app instance group name, application ID, and application name.
//
// Description:
//
// ## Operation description
//
// This operation queries the list of app instance groups to which a specified user (EndUserId) has been granted **app instance group-level authorization**. It returns the basic information of each app instance group (such as ID, name, status, region, creation time, and expiration time) and the list of applications deployed in the app instance group.
//
// Scope of returned results:
//
// - Only app instance groups that are authorized to the user as a whole through the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation are returned. Records authorized for individual applications through the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation are not included.
//
// - Only app instance groups whose product type matches ProductType and that are not deleted are returned. If no applications are deployed in the app instance group image, the app instance group is not returned.
//
// - The results are sorted in descending order based on the update time of the authorization records. The most recently authorized or modified app instance groups are ranked first.
//
// ## Before you begin
//
// - The [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation is called to authorize the app instance group to the user.
//
// ## Parameter description
//
// - **ProductType and EndUserId are required**. If ProductType is not specified, the error code `InvalidParameter.ProductType` is returned. If EndUserId is not specified, the error code `InvalidParameter.UserId` is returned.
//
// - EndUserId uses **exact match*	- based on the username. AppInstanceGroupId, AppInstanceGroupName, AppId, and AppName use **fuzzy match*	- (a hit occurs if the string is contained). If multiple filter conditions are specified, all conditions must be met.
//
// - PageNumber starts from 1. Valid values for PageSize: 1 to 100.
//
// - If the user does not have any app instance groups with authorization that meet the conditional requirements, the operation returns a normal response: AppInstanceGroupModels is an empty list and TotalCount is 0.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the app instance group ID, and then call the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation to authorize the app instance group to the user.
//
// 2. Call this operation to query the app instance groups authorized for the user and the applications deployed in the app instance groups.
//
// 3. To obtain the application connection ticket for the user, call the [GetConnectionTicket](~~GetConnectionTicket~~) operation based on the returned AppInstanceGroupId and the AppId in Apps.
//
// @param request - ListAuthorizedAppInstanceGroupByUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedAppInstanceGroupByUserResponse
func (client *Client) ListAuthorizedAppInstanceGroupByUserWithContext(ctx context.Context, request *ListAuthorizedAppInstanceGroupByUserRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedAppInstanceGroupByUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedAppInstanceGroupByUser"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedAppInstanceGroupByUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve the applications for which a specified user has obtained access permissions through application authorization, and supports fuzzy filtering by app instance group ID, app instance group name, application ID, or application name.
//
// Description:
//
// ## Operation description
//
// This operation queries the applications authorized for a specified user at the application level. **It returns only the records of applications authorized for the user at the application level*	- (for example, authorization completed by calling the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation). Authorization granted for an entire app instance group by calling the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation is not included in the response.
//
// The response is deduplicated by the combination of app instance group and application. Each record corresponds to an application within an app instance group.
//
// ## Before you begin
//
// - The user exists in the current account. If the user does not exist, the error code User.NotFound is returned. You can call the [DescribeUsers](https://help.aliyun.com/document_detail/436936.html) operation to obtain the username.
//
// - The app instance group to which the application belongs has been created, and the application has been authorized for the user at the application level. If no application-level authorization has been performed, an empty application list is returned.
//
// ## Metric description
//
// - **EndUserId and ProductType are required.*	- Set ProductType to CloudApp, which indicates WUYING Cloud App.
//
// - AppInstanceGroupId, AppId, AppInstanceGroupName, and AppName are optional filter conditions. Fuzzy match is used for all these parameters, and they can be used in any combination. If all of them are omitted, all application-level authorization records for the user are returned.
//
// - Use PageNumber and PageSize for paging. PageNumber starts from 1, and the valid values for PageSize range from 1 to 100. You can determine whether to continue the query based on the returned TotalCount.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) or [GetAppInstanceGroup](https://help.aliyun.com/document_detail/600836.html) operation to obtain the app instance group ID (AppInstanceGroupId) and the application IDs (AppId in the Apps list) of the applications deployed in the app instance group.
//
// 2. Call the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation to authorize the applications for the target user.
//
// 3. Call this operation to query the applications authorized for the user.
//
// @param request - ListAuthorizedAppsByUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedAppsByUserResponse
func (client *Client) ListAuthorizedAppsByUserWithContext(ctx context.Context, request *ListAuthorizedAppsByUserRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedAppsByUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedAppsByUser"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedAppsByUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user groups authorized for an app instance group.
//
// @param request - ListAuthorizedUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedUserGroupsResponse
func (client *Client) ListAuthorizedUserGroupsWithContext(ctx context.Context, request *ListAuthorizedUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedUserGroups"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedUserGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query on the authorized user list of a specified app instance group or app instance group set, and supports exact or fuzzy filtering by username.
//
// Description:
//
// ## Operation description
//
// This operation queries the currently authorized users of a specified app instance group (AppInstanceGroupId) or app instance group set (AppInstanceGroupSetId). It returns the username, account type, email address, phone number, and the authorization mode of the app instance group to which each user belongs.
//
// - This operation returns **authorization relationships**, which does not indicate that the users are currently online or have established connections.
//
// - In the Cloud Browser product, an app instance group is a cloud browser group, and the app instance group ID is the browser group ID.
//
// The return scope depends on the authorization mode of the app instance group (the AuthMode response parameter):
//
// - When the authorization mode is `App` (application authorization) or `AppInstanceGroup` (app instance group authorization): returns the users authorized by app instance group through the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation and the users authorized by application through the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation. If AppId is specified, only the users **specifically authorized for this application*	- are returned.
//
// - When the authorization mode is `Session` (session authorization): returns the users who are granted persistent sessions, and lists all persistent session IDs granted to each user in AppInstancePersistentIds. If AppInstancePersistentId is specified, only the users granted this session are returned.
//
// - When querying by app instance group set: returns the users authorized to the set. The AppInstanceGroupId response parameter is the primary app instance group ID of the set, and AppInstanceGroupSetId is the ID of the queried set.
//
// When querying by app instance group, the results are sorted in descending order of authorization time, with the most recently authorized users listed first.
//
// ## Before you begin
//
// - The target app instance group or app instance group set has been created, belongs to the current account, and its product type matches ProductType. Otherwise, a resource not found error code is returned.
//
// - Users have been authorized through the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) or [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation. If no users have been authorized, the operation returns a normal response with Users as an empty list and TotalCount as 0.
//
// ## Parameter description
//
// - **ProductType, PageNumber, and PageSize are required**. If the value of ProductType is invalid, the `ProductTypeInvalid` error code is returned.
//
// - **You must specify either AppInstanceGroupId or AppInstanceGroupSetId, but not both**. If both are specified or neither is specified, the `InvalidParameter.AppInstanceGroupId/AppInstanceGroupSetId` error code is returned.
//
// - **AppId and AppInstancePersistentId are not supported when querying by app instance group set**. If either is specified, the `InvalidParameter.AppId` or `InvalidParameter.AppInstancePersistentId` error code is returned.
//
// - EndUserId performs **exact match*	- by username, and UserIdFuzzy performs **fuzzy match*	- by username (a hit occurs if the username contains the keyword). Both can be specified at the same time, and both conditions must be met.
//
// - PageNumber starts from 1. The value range of PageSize is 1 to 100.
//
// - When the authorization mode is `App` or `AppInstanceGroup`, TotalCount is the **number of authorization records*	- that meet the conditions. If the same user has multiple authorization records (for example, authorized for multiple applications), the user is merged into one in Users. Therefore, the number of users returned on the current page may be less than PageSize. Use TotalCount to determine whether to continue paging. When the authorization mode is `Session`, TotalCount is the number of deduplicated users.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the app instance group ID. For cloud browser groups, call the [ListBrowserInstanceGroup](~~ListBrowserInstanceGroup~~) operation.
//
// 2. Call the [AuthorizeInstanceGroup](~~AuthorizeInstanceGroup~~) operation to authorize by app instance group, or call the [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~) operation to authorize by application.
//
// 3. Call this operation to query the authorized users. To remove authorization, pass the returned EndUserId to the UnAuthorizeUserIds parameter of the preceding authorization operations.
//
// ## Error codes
//
// - `ProductTypeInvalid`: The value of ProductType is invalid.
//
// - `InvalidParameter.AppInstanceGroupId/AppInstanceGroupSetId`: Both AppInstanceGroupId and AppInstanceGroupSetId are specified, or neither is specified.
//
// - `InvalidParameter.AppId`: AppId is specified when querying by app instance group set.
//
// - `InvalidParameter.AppInstancePersistentId`: AppInstancePersistentId is specified when querying by app instance group set.
//
// - `InvalidAppInstanceGroupSpecItem.NotFound`: The app instance group does not exist, does not belong to the current account, or the product type does not match.
//
// - `InvalidBrowserInstanceGroup.NotFound`: When ProductType is `CloudBrowser`, the cloud browser group does not exist, does not belong to the current account, or the product type does not match.
//
// - `InvalidAppInstanceGroupSet.NotFound`: The app instance group set does not exist, does not belong to the current account, the product type does not match, or there is no available primary app instance group in the set.
//
// - `InvalidAppInstanceGroupSet.ActivationFailed`: The app instance group set is not in an available state.
//
// @param request - ListAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedUsersResponse
func (client *Client) ListAuthorizedUsersWithContext(ctx context.Context, request *ListAuthorizedUsersRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.UserIdFuzzy) {
		query["UserIdFuzzy"] = request.UserIdFuzzy
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedUsers"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the binding information between users and resources.
//
// @param request - ListBindInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindInfoResponse
func (client *Client) ListBindInfoWithContext(ctx context.Context, request *ListBindInfoRequest, runtime *dara.RuntimeOptions) (_result *ListBindInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		body["AppIdList"] = request.AppIdList
	}

	if !dara.IsNil(request.AppInstanceGroupIdList) {
		body["AppInstanceGroupIdList"] = request.AppInstanceGroupIdList
	}

	if !dara.IsNil(request.AppInstanceIdList) {
		body["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserIdList) {
		body["UserIdList"] = request.UserIdList
	}

	if !dara.IsNil(request.WyIdList) {
		body["WyIdList"] = request.WyIdList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of bookmarks under a specified cloud browser group or delivery group set by using paging.
//
// @param request - ListBrowserBookmarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrowserBookmarksResponse
func (client *Client) ListBrowserBookmarksWithContext(ctx context.Context, request *ListBrowserBookmarksRequest, runtime *dara.RuntimeOptions) (_result *ListBrowserBookmarksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrowserInstanceGroupId) {
		query["BrowserInstanceGroupId"] = request.BrowserInstanceGroupId
	}

	if !dara.IsNil(request.BrowserInstanceGroupSetId) {
		query["BrowserInstanceGroupSetId"] = request.BrowserInstanceGroupSetId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

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
		Action:      dara.String("ListBrowserBookmarks"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBrowserBookmarksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud browser groups and their current status by using paging.
//
// Description:
//
// This topic describes how to use queries in the monthly active user (MAU) billing scenario.
//
// ## Query conditions
//
// You can filter by browser group identity, name, business region, office network, set, authorized user group, and status. You can specify only one value for status filtering at a time.
//
// ## Paging
//
// Use `PageNumber` and `PageSize` to perform a paged query. Use the returned `TotalCount` to determine whether to query the next page.
//
// ## What to do next
//
// To view the detailed configuration of a single browser group, invoke `GetBrowserInstanceGroup` by using the returned identity.
//
// ## Example description
//
// The `-` in the examples only indicates that the current scenario does not apply or the field is not returned. It is not a character string actually returned by the API operation. Common resource identity examples are masked. Use the actual query results when you use them.
//
// @param request - ListBrowserInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrowserInstanceGroupResponse
func (client *Client) ListBrowserInstanceGroupWithContext(ctx context.Context, request *ListBrowserInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListBrowserInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupSetId) {
		query["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.BrowserInstanceGroupId) {
		query["BrowserInstanceGroupId"] = request.BrowserInstanceGroupId
	}

	if !dara.IsNil(request.BrowserInstanceGroupName) {
		query["BrowserInstanceGroupName"] = request.BrowserInstanceGroupName
	}

	if !dara.IsNil(request.CloudBrowserName) {
		query["CloudBrowserName"] = request.CloudBrowserName
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Tier) {
		query["Tier"] = request.Tier
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludedUserGroupIds) {
		body["ExcludedUserGroupIds"] = request.ExcludedUserGroupIds
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		bodyFlat["Status"] = request.Status
	}

	if !dara.IsNil(request.UserGroupIds) {
		body["UserGroupIds"] = request.UserGroupIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBrowserInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBrowserInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query on the list of access restriction URLs configured for a specified browser delivery group or delivery group set.
//
// Description:
//
// Invoke this operation to perform a paging query on the list of access restriction URLs (Restricted URLs) configured for a cloud browser delivery group or delivery group set to view the details of the URLs configured in the current access control policy.
//
// You must specify either BrowserInstanceGroupId or BrowserInstanceGroupSetId, but not both. If you specify BrowserInstanceGroupId, the restricted URLs for the delivery group are queried. If you specify BrowserInstanceGroupSetId, the restricted URLs for the delivery group set are queried. Specifying both or neither returns a parameter error.
//
// You can use Keyword to filter restricted URLs by using a case-insensitive include match. If this parameter is not specified, all restricted URLs are returned. Pagination is controlled by PageNumber and PageSize. Valid values for PageSize: 1 to 200. Default value: 200. Default value for PageNumber: 0. If the value is less than 1, the first page is returned. In the response, TotalCount indicates the total number of records that meet the filter conditions, and Data contains the records on the current page. If no data is found, an empty array is returned.
//
// @param request - ListBrowserRestrictedURLsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrowserRestrictedURLsResponse
func (client *Client) ListBrowserRestrictedURLsWithContext(ctx context.Context, request *ListBrowserRestrictedURLsRequest, runtime *dara.RuntimeOptions) (_result *ListBrowserRestrictedURLsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrowserInstanceGroupId) {
		query["BrowserInstanceGroupId"] = request.BrowserInstanceGroupId
	}

	if !dara.IsNil(request.BrowserInstanceGroupSetId) {
		query["BrowserInstanceGroupSetId"] = request.BrowserInstanceGroupSetId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

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
		Action:      dara.String("ListBrowserRestrictedURLs"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBrowserRestrictedURLsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of commands delivered by calling RunAppInstanceCommand in persistent sessions, including the execution status, start and end times, and command output.
//
// Description:
//
// Call this operation to query the execution results of a command delivery (InvocationId) in each persistent session to track the execution status and obtain the command output. InvocationId is the command execution ID returned when you call RunAppInstanceCommand to deliver a command.
//
// Currently, only WuyingServer is supported for ProductType. You must specify at least one of AppInstanceGroupIds and AppInstanceGroupId. If both are specified, AppInstanceGroupIds takes precedence. When AppInstanceGroupIds is specified, you can pass up to 200 delivery group IDs, and these delivery groups must be in the same region. If any delivery group does not exist, a delivery group not found error is returned. If the delivery groups are in different regions, a region mismatch error is returned. When only AppInstanceGroupId is specified, the command execution results under this delivery group are queried.
//
// You can use AppInstancePersistentIds to query only the execution results of specified persistent sessions. When AppInstanceGroupIds is specified, only persistent sessions belonging to these delivery groups are retained. If no matching persistent sessions are found, an empty list is returned directly, and TotalCount is 0. For pagination, use PageNumber and PageSize. MaxResults and NextToken do not take effect currently.
//
// @param request - ListCommandInvocationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommandInvocationResultsResponse
func (client *Client) ListCommandInvocationResultsWithContext(ctx context.Context, request *ListCommandInvocationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListCommandInvocationResultsResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupIds) {
		body["AppInstanceGroupIds"] = request.AppInstanceGroupIds
	}

	if !dara.IsNil(request.AppInstancePersistentIds) {
		body["AppInstancePersistentIds"] = request.AppInstancePersistentIds
	}

	if !dara.IsNil(request.InvocationId) {
		body["InvocationId"] = request.InvocationId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommandInvocationResults"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommandInvocationResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of desktop Agent runtimes.
//
// @param request - ListDesktopAgentRuntimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDesktopAgentRuntimeResponse
func (client *Client) ListDesktopAgentRuntimeWithContext(ctx context.Context, request *ListDesktopAgentRuntimeRequest, runtime *dara.RuntimeOptions) (_result *ListDesktopAgentRuntimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentImOnlineStatus) {
		query["AgentImOnlineStatus"] = request.AgentImOnlineStatus
	}

	if !dara.IsNil(request.AgentInstanceStatuses) {
		query["AgentInstanceStatuses"] = request.AgentInstanceStatuses
	}

	if !dara.IsNil(request.AgentInstanceVersions) {
		query["AgentInstanceVersions"] = request.AgentInstanceVersions
	}

	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentPlatformList) {
		query["AgentPlatformList"] = request.AgentPlatformList
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.AgentProviderList) {
		query["AgentProviderList"] = request.AgentProviderList
	}

	if !dara.IsNil(request.AgentTemplateIds) {
		query["AgentTemplateIds"] = request.AgentTemplateIds
	}

	if !dara.IsNil(request.AuthUsers) {
		query["AuthUsers"] = request.AuthUsers
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ChannelConfigure) {
		query["ChannelConfigure"] = request.ChannelConfigure
	}

	if !dara.IsNil(request.DeploymentSource) {
		query["DeploymentSource"] = request.DeploymentSource
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.DesktopNames) {
		query["DesktopNames"] = request.DesktopNames
	}

	if !dara.IsNil(request.DesktopStatuses) {
		query["DesktopStatuses"] = request.DesktopStatuses
	}

	if !dara.IsNil(request.HasAuthUser) {
		query["HasAuthUser"] = request.HasAuthUser
	}

	if !dara.IsNil(request.HasRisk) {
		query["HasRisk"] = request.HasRisk
	}

	if !dara.IsNil(request.IncludeRiskInfo) {
		query["IncludeRiskInfo"] = request.IncludeRiskInfo
	}

	if !dara.IsNil(request.ManagementStatus) {
		query["ManagementStatus"] = request.ManagementStatus
	}

	if !dara.IsNil(request.ModelConfigure) {
		query["ModelConfigure"] = request.ModelConfigure
	}

	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDesktopAgentRuntime"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDesktopAgentRuntimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image information.
//
// @param request - ListImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageResponse
func (client *Client) ListImageWithContext(ctx context.Context, request *ListImageRequest, runtime *dara.RuntimeOptions) (_result *ListImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Distro) {
		query["Distro"] = request.Distro
	}

	if !dara.IsNil(request.TagList) {
		query["TagList"] = request.TagList
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionIdList) {
		body["BizRegionIdList"] = request.BizRegionIdList
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BizTypeList) {
		body["BizTypeList"] = request.BizTypeList
	}

	if !dara.IsNil(request.FeatureList) {
		body["FeatureList"] = request.FeatureList
	}

	if !dara.IsNil(request.FotaVersion) {
		body["FotaVersion"] = request.FotaVersion
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ImageType) {
		body["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.LanguageType) {
		body["LanguageType"] = request.LanguageType
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlatformName) {
		body["PlatformName"] = request.PlatformName
	}

	if !dara.IsNil(request.PlatformNameList) {
		body["PlatformNameList"] = request.PlatformNameList
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ProductTypeList) {
		body["ProductTypeList"] = request.ProductTypeList
	}

	if !dara.IsNil(request.ProtocolType) {
		body["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.ResourceInstanceType) {
		body["ResourceInstanceType"] = request.ResourceInstanceType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve the basic information of images, including the image name, type, status, and version identity, with paging support.
//
// Description:
//
// Supports filtering by image ID, name, status, type, package type, operating system, business line, product type, and protocol type. Multiple filter conditions take effect simultaneously, and the results are returned in descending order of image creation time. Returns the basic information and version identifiers of images. To obtain specific version and region distribution information, call ListImage.
//
// @param request - ListImageBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageBaseInfoResponse
func (client *Client) ListImageBaseInfoWithContext(ctx context.Context, request *ListImageBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListImageBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ImageType) {
		body["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.OsType) {
		body["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ProtocolType) {
		body["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageBaseInfo"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageBaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the built-in tags supported by images.
//
// @param request - ListImageTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageTagsResponse
func (client *Client) ListImageTagsWithContext(ctx context.Context, request *ListImageTagsRequest, runtime *dara.RuntimeOptions) (_result *ListImageTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

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
		Action:      dara.String("ListImageTags"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of LLM templates.
//
// Description:
//
// You can use paging to query the list of model templates under a model provider template in the Wuying Agent Management Center. You can filter the results by model group ID, model provider template ID, model template ID, and model encoding. When you query by model group, the default model is automatically pinned to the top.
//
// Before you call this operation, make sure that you fully understand the operations and usage of the Wuying Agent Management Center.
//
// @param tmpReq - ListLlmTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLlmTemplatesResponse
func (client *Client) ListLlmTemplatesWithContext(ctx context.Context, tmpReq *ListLlmTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListLlmTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListLlmTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LlmTemplateIds) {
		request.LlmTemplateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmTemplateIds, dara.String("LlmTemplateIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LlmCode) {
		query["LlmCode"] = request.LlmCode
	}

	if !dara.IsNil(request.LlmTemplateIdsShrink) {
		query["LlmTemplateIds"] = request.LlmTemplateIdsShrink
	}

	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProviderTemplateId) {
		query["ProviderTemplateId"] = request.ProviderTemplateId
	}

	if !dara.IsNil(request.SmartModel) {
		query["SmartModel"] = request.SmartModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLlmTemplates"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLlmTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of model provider endpoints.
//
// @param request - ListModelProviderEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProviderEndpointsResponse
func (client *Client) ListModelProviderEndpointsWithContext(ctx context.Context, request *ListModelProviderEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListModelProviderEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ProviderName) {
		query["ProviderName"] = request.ProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelProviderEndpoints"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProviderEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available LLM model codes for a specified model provider. Supports filtering by LLM code, model brand, recommended tier, and recommendation status, and returns model information such as the model name, description, release time, features, inference modality, pricing, credit multiplier, and risk type.
//
// @param request - ListModelProviderLlmCodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProviderLlmCodesResponse
func (client *Client) ListModelProviderLlmCodesWithContext(ctx context.Context, request *ListModelProviderLlmCodesRequest, runtime *dara.RuntimeOptions) (_result *ListModelProviderLlmCodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LlmCode) {
		query["LlmCode"] = request.LlmCode
	}

	if !dara.IsNil(request.ModelBrands) {
		query["ModelBrands"] = request.ModelBrands
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProviderName) {
		query["ProviderName"] = request.ProviderName
	}

	if !dara.IsNil(request.ProviderTemplateId) {
		query["ProviderTemplateId"] = request.ProviderTemplateId
	}

	if !dara.IsNil(request.Recommended) {
		query["Recommended"] = request.Recommended
	}

	if !dara.IsNil(request.RecommendedLevels) {
		query["RecommendedLevels"] = request.RecommendedLevels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelProviderLlmCodes"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProviderLlmCodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of model provider templates.
//
// Description:
//
// You can use paging to query the list of model provider templates under a specified model group in the Wuying Agent Management Center. You can filter the templates by provider name, model group ID, and provider template ID. Make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center before you call this operation.
//
// @param tmpReq - ListModelProviderTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProviderTemplatesResponse
func (client *Client) ListModelProviderTemplatesWithContext(ctx context.Context, tmpReq *ListModelProviderTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListModelProviderTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModelProviderTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProviderTemplateIds) {
		request.ProviderTemplateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProviderTemplateIds, dara.String("ProviderTemplateIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProviderName) {
		query["ProviderName"] = request.ProviderName
	}

	if !dara.IsNil(request.ProviderTemplateIdsShrink) {
		query["ProviderTemplateIds"] = request.ProviderTemplateIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelProviderTemplates"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProviderTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of resource groups associated with a model group.
//
// Description:
//
// You can call this operation to query the list of resource groups authorized by a model group in the WUYING Agent Management Center.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
//
// @param request - ListModelTemplateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelTemplateResourceGroupResponse
func (client *Client) ListModelTemplateResourceGroupWithContext(ctx context.Context, request *ListModelTemplateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListModelTemplateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelTemplateResourceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelTemplateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of model templates.
//
// Description:
//
// You can use paged query to retrieve the list of model groups created in the Wuying Agent Management Center. You can filter results by Agent provider, Agent platform, template group ID, and whether models have been configured. Paging is supported.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param tmpReq - ListModelTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelTemplatesResponse
func (client *Client) ListModelTemplatesWithContext(ctx context.Context, tmpReq *ListModelTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListModelTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModelTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModelTemplateIdList) {
		request.ModelTemplateIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelTemplateIdList, dara.String("ModelTemplateIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentPlatformList) {
		query["AgentPlatformList"] = request.AgentPlatformList
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.AgentProviderList) {
		query["AgentProviderList"] = request.AgentProviderList
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.HasModel) {
		query["HasModel"] = request.HasModel
	}

	if !dara.IsNil(request.ModelTemplateIdListShrink) {
		query["ModelTemplateIdList"] = request.ModelTemplateIdListShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RefScope) {
		query["RefScope"] = request.RefScope
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelTemplates"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource specifications available for selection when creating a delivery group.
//
// @param request - ListNodeInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeInstanceTypeResponse
func (client *Client) ListNodeInstanceTypeWithContext(ctx context.Context, request *ListNodeInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *ListNodeInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Gpu) {
		query["Gpu"] = request.Gpu
	}

	if !dara.IsNil(request.GpuMemory) {
		query["GpuMemory"] = request.GpuMemory
	}

	if !dara.IsNil(request.InstanceTypeForModify) {
		query["InstanceTypeForModify"] = request.InstanceTypeForModify
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NodeInstanceType) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !dara.IsNil(request.NodeInstanceTypeFamily) {
		query["NodeInstanceTypeFamily"] = request.NodeInstanceTypeFamily
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeInstanceType"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of resource nodes.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, request *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history of over-the-air updates.
//
// @param request - ListOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOtaTaskResponse
func (client *Client) ListOtaTaskWithContext(ctx context.Context, request *ListOtaTaskRequest, runtime *dara.RuntimeOptions) (_result *ListOtaTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.OtaType) {
		body["OtaType"] = request.OtaType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOtaTask"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of persistent session application instances in a delivery group.
//
// @param request - ListPersistentAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPersistentAppInstancesResponse
func (client *Client) ListPersistentAppInstancesWithContext(ctx context.Context, request *ListPersistentAppInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListPersistentAppInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstancePersistentIds) {
		query["AppInstancePersistentIds"] = request.AppInstancePersistentIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPersistentAppInstances"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPersistentAppInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries published delivery groups and their application information for a specified product type under the current Alibaba Cloud account by page, and supports filtering by delivery group, application, and user authorization status.
//
// Description:
//
// ## Prerequisites
//
// When you call this operation as a RAM user or by using STS credentials, you must have the appstreaming:ListPublishedAppInstanceGroup permission.
//
// Only delivery groups that are in the published state and have deployed applications in their images are included in the query results. This operation only queries information. It does not create delivery groups, assign users, or grant application access permissions.
//
// ## Query and pagination
//
// - **You must explicitly specify ProductType, PageNumber, and PageSize.*	- The page number starts from 1, and the number of entries per page ranges from 1 to 100.
//
// - AppInstanceGroupId, AppInstanceGroupName, AppId, and AppName all support fuzzy match. You can specify them individually or in combination. If multiple conditions are specified, all conditions must be met. If both AppId and AppName are specified, they must be met by the same application.
//
// - If an optional filter parameter is not specified or an empty string is specified, the condition is not used for filtering. The query results are sorted in descending order of the delivery group creation time. A delivery group is not returned multiple times even if it contains multiple matched applications.
//
// - ExcludeUserId is used to exclude delivery groups for which all applications have been directly authorized to a specified user. It cannot be used to determine whether the user has no access permissions at all.
//
// - AppId and AppName only filter delivery groups and **do not limit the returned Apps list to the matched applications**.
//
// ## Call sequence
//
// 1. Initiate a query with PageNumber=1 and the required PageSize. For Wuying Cloud App scenarios, set ProductType=CloudApp.
//
// 2. Read AppInstanceGroupModels. To retrieve the next page, keep the product type and filter conditions unchanged and increment PageNumber. If no delivery groups match the conditions, the total count is 0 and the list is empty. If the page number exceeds the range of the current results, the list may also be empty, but the total count still indicates the total number of all matched delivery groups.
//
// 3. To obtain the details of a single delivery group, pass the complete AppInstanceGroupId returned and the same ProductType to [GetAppInstanceGroup](~~GetAppInstanceGroup~~).
//
// The desensitized identity in the examples is used to demonstrate the format. Replace it with your actual identity when you call the operation.
//
// @param request - ListPublishedAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAppInstanceGroupResponse
func (client *Client) ListPublishedAppInstanceGroupWithContext(ctx context.Context, request *ListPublishedAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListPublishedAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ExcludeUserId) {
		query["ExcludeUserId"] = request.ExcludeUserId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query on applications deployed in published delivery groups of a specified product type under the current Alibaba Cloud account. Returns each application along with its delivery group and the number of users granted authorization for the application. Supports filtering by delivery group, application, and user authorization status.
//
// Description:
//
// ## Prerequisites
//
// To call this operation as a RAM user or by using STS credentials, you must have the `appstreaming:ListPublishedApps` permission.
//
// Only applications deployed in the images of published delivery groups are returned in the query results. This operation only queries information. It does not create delivery groups, nor does it grant or revoke application access permissions.
//
// ## Return granularity
//
// Results are returned item by item based on the "delivery group + application" combination. If the same application is deployed in multiple published delivery groups, multiple records are returned. `TotalCount` also counts the number of combinations. The list is sorted in descending order of the creation time of the associated delivery groups.
//
// ## Query and pagination
//
// - **You must explicitly specify `ProductType`, `PageNumber`, and `PageSize`.*	- The page number starts from `1`, and the number of entries per page ranges from `1` to `100`. If the values are invalid, the error codes `InvalidParameter.PageNumber` and `InvalidParameter.PageSize` are returned, respectively.
//
// - `AppInstanceGroupId`, `AppInstanceGroupName`, `AppId`, and `AppName` all support substring matching. You can specify them individually or in combination. If multiple conditions are specified, all conditions must be met.
//
// - If an optional filter parameter is not specified or an empty string is passed, filtering is not performed based on this condition.
//
// - `ExcludeUserId` is used to exclude applications that have been authorized for a specified user on a per-application basis, making it easier to filter out applications that can still be authorized for the user. **Access permissions obtained through overall delivery group authorization or user groups are not evaluated by this condition.**
//
// ## Call sequence
//
// 1. Initiate a query with `PageNumber=1` and the desired `PageSize`. For WUYING Cloud App scenarios, use `ProductType=CloudApp`.
//
// 2. Read `Apps`. To retrieve the next page, keep the product type and filter conditions unchanged and increment `PageNumber`. If no matching results are found, `TotalCount` is `0` and `Apps` is an empty list. If the page number exceeds the result range, `Apps` may also be empty, but `TotalCount` still indicates the total number of matching records.
//
// 3. When you need to grant authorization to a user for an application on a per-application basis, pass the returned `AppInstanceGroupId`, `AppId`, and the same `ProductType` to [AuthorizeUsersForApp](~~AuthorizeUsersForApp~~). To view delivery group details, pass `AppInstanceGroupId` and `ProductType` to [GetAppInstanceGroup](~~GetAppInstanceGroup~~).
//
// The masked identity in the examples is used to demonstrate the format. Replace it with your actual identity when you invoke the operation.
//
// @param request - ListPublishedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAppsResponse
func (client *Client) ListPublishedAppsWithContext(ctx context.Context, request *ListPublishedAppsRequest, runtime *dara.RuntimeOptions) (_result *ListPublishedAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ExcludeUserId) {
		query["ExcludeUserId"] = request.ExcludeUserId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedApps"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of region IDs available for image services under a specified product for the current account.
//
// Description:
//
// The returned results are affected by the available region configurations of the account and the product type. If ProductType is omitted, the query is performed based on CloudApp.
//
// This operation returns the regions available for image services. You cannot use this information to determine whether a specific image has been distributed to these regions. Before you create or distribute an image based on a parent image, you must also check the region information of the destination image.
//
// @param request - ListRegionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionIdResponse
func (client *Client) ListRegionIdWithContext(ctx context.Context, request *ListRegionIdRequest, runtime *dara.RuntimeOptions) (_result *ListRegionIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegionId"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions supported by WUYING Cloud Application.
//
// Description:
//
// > The regions returned by this operation are not all available regions. For information about available regions, see [Supported regions](https://help.aliyun.com/document_detail/426036.html).
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithContext(ctx context.Context, request *ListRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizSource) {
		query["BizSource"] = request.BizSource
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of associated templates, cloud desktops, and cloud desktop pools.
//
// @param request - ListRelatedResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRelatedResourceResponse
func (client *Client) ListRelatedResourceWithContext(ctx context.Context, request *ListRelatedResourceRequest, runtime *dara.RuntimeOptions) (_result *ListRelatedResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRelatedResource"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRelatedResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of one or more specified cloud resources.
//
// @param request - ListTagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagCloudResourcesResponse
func (client *Client) ListTagCloudResourcesWithContext(ctx context.Context, request *ListTagCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagCloudResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagCloudResources"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagCloudResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags bound to one or more delivery groups, or filters delivery groups that are bound to specified tags by tag keys and values.
//
// Description:
//
// This operation follows the Alibaba Cloud tagging specification and is used to query tags bound to Wuying Cloud App delivery groups. You can view the tags bound to resources by resource ID, or filter resources bound to specified tags by tag keys and values.
//
// ## Before you begin
//
// - A delivery group is created. You can obtain the delivery group ID by calling the ListAppInstanceGroup operation.
//
// - Tags are bound to the delivery group. You can bind tags by calling the TagResources operation, or bind them in the console or Tag service.
//
// ## Parameter description
//
// - **ResourceType is required**. Only `APPINSTANCEGROUP` (delivery group) is supported. If you specify other values, the error code `InvalidResourceType.Invalid` is returned.
//
// - **You must specify at least one of ResourceId.N and Tag.N**. If neither is specified, the error code `MissingParameter.ResourceIdsOrTags` is returned.
//
//   - If only ResourceId.N is specified: all tags bound to these resources are returned.
//
//   - If only Tag.N is specified: all resources bound to the specified tags and their matched tags are returned.
//
//   - If both are specified: only records of the specified resources bound to the specified tags are returned.
//
// - Tag.N.Key is required. If it is empty, the error code `InvalidTagPolicy.KeyInvalid` is returned. Tag.N.Value is optional. If it is not specified, the value corresponding to the tag key is not limited.
//
// - Multiple Tag.N parameters have an AND relationship. A resource is returned only if it is bound to all specified tags.
//
// ## Response description
//
// - The response returns one record for each resource and tag key-value pair. If a resource is bound to multiple tags, multiple records are returned.
//
// - TotalCount indicates the number of records returned in the current request.
//
// - This operation returns all matched results at once. An empty NextToken in the response indicates that no more data is available.
//
// ## Procedure
//
// 1. Call ListAppInstanceGroup to obtain the delivery group ID.
//
// 2. Call TagResources to bind tags to the delivery group.
//
// 3. Call this operation to query the binding relationships by resource ID or tag condition.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2021-09-01"),
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
// Queries the step identifiers, step categories, and interaction methods for application adaptation or image creation tasks.
//
// Description:
//
// Returns application adaptation steps when Type is set to ADAPTOR, and image creation steps when Type is set to IMAGE.
//
// This operation returns step definitions, not the actual progress of a specific task. Read the step definitions first, and then display the task progress based on the current step and step status returned by GetTask. The step sets vary across different task types. Do not use the step index as the completion percentage.
//
// @param request - ListTaskStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskStepResponse
func (client *Client) ListTaskStepWithContext(ctx context.Context, request *ListTaskStepRequest, runtime *dara.RuntimeOptions) (_result *ListTaskStepResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskStep"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskStepResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of workstations.
//
// Description:
//
// Retrieves the list of WUYING workstations.
//
// @param request - ListWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWuyingServerResponse
func (client *Client) ListWuyingServerWithContext(ctx context.Context, request *ListWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *ListWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddVirtualNodePoolStatusList) {
		bodyFlat["AddVirtualNodePoolStatusList"] = request.AddVirtualNodePoolStatusList
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.ExpiredTimeEnd) {
		body["ExpiredTimeEnd"] = request.ExpiredTimeEnd
	}

	if !dara.IsNil(request.ExpiredTimeStart) {
		body["ExpiredTimeStart"] = request.ExpiredTimeStart
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.NetworkInterfaceIp) {
		body["NetworkInterfaceIp"] = request.NetworkInterfaceIp
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ServerInstanceType) {
		body["ServerInstanceType"] = request.ServerInstanceType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Users) {
		bodyFlat["Users"] = request.Users
	}

	if !dara.IsNil(request.VirtualNodePoolId) {
		body["VirtualNodePoolId"] = request.VirtualNodePoolId
	}

	if !dara.IsNil(request.WuyingServerIdList) {
		bodyFlat["WuyingServerIdList"] = request.WuyingServerIdList
	}

	if !dara.IsNil(request.WuyingServerNameOrId) {
		body["WuyingServerNameOrId"] = request.WuyingServerNameOrId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available Wuying workstation workload templates for a specified virtual node pool.
//
// Description:
//
// Queries the available workload templates for a specified virtual node pool. Before calling this operation, call [ListWuyingServer](~~ListWuyingServer~~) to obtain the `VirtualNodePoolId` under the current account. This operation verifies the ownership of the node pool and returns up to the first 100 records sorted by template configuration order. The template content is Base64-encoded after being encoded in UTF-8. You must decode the Base64 string before use. The current version does not use `MaxResults` and `NextToken` for pagination.
//
// @param request - ListWuyingServerWorkloadTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWuyingServerWorkloadTemplateResponse
func (client *Client) ListWuyingServerWorkloadTemplateWithContext(ctx context.Context, request *ListWuyingServerWorkloadTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListWuyingServerWorkloadTemplateResponse, _err error) {
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

	if !dara.IsNil(request.VirtualNodePoolId) {
		query["VirtualNodePoolId"] = request.VirtualNodePoolId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWuyingServerWorkloadTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWuyingServerWorkloadTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available zone IDs for a specified product type and operating system type in a specified region.
//
// Description:
//
// ## Usage notes
//
// This operation returns a list of available zone IDs that can be selected by the current Alibaba Cloud account based on the region, product type, and operating system type. A typical use case is to confirm which zones are available before creating a resource that requires a specified vSwitch, and then select a vSwitch located in one of these zones.
//
// This operation only queries information. It does not create resources or incur fees.
//
// ## Before you begin
//
// - When you call this operation as a RAM user or by using STS credentials, you must have the appstreaming:ListAppInstanceType permission. If you do not have the required permissions, the error code Forbidden.NoPermission is returned.
//
// - **ProductType, BizRegionId, and OsType are all required.*	- If any of these parameters is missing, empty, or has an unrecognized value, the error code InvalidParameter.ValueInvalid is returned.
//
// - BizRegionId must be a region ID supported by App Streaming. You can call [ListRegions](~~ListRegions~~) to obtain the supported region IDs.
//
// ## Call sequence
//
// 1. Call [ListRegions](~~ListRegions~~) to obtain the supported region IDs.
//
// 2. Call this operation by using the region ID, target product type, and operating system type, and then read the returned ListZonesModel.Zones.
//
// 3. When you create a resource and need to specify a vSwitch, select a vSwitch located in one of the returned zones. For example, pass the corresponding vSwitch ID when you call [CreateAppInstanceGroup](~~CreateAppInstanceGroup~~) or [CreateWuyingServer](~~CreateWuyingServer~~).
//
// The returned list of available zones is determined by the available resources in the current region and may change over time. Query the list in real time before creating resources instead of caching the list for a long time.
//
// @param request - ListZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZonesResponse
func (client *Client) ListZonesWithContext(ctx context.Context, request *ListZonesRequest, runtime *dara.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListZones"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unregisters all sessions in a pay-as-you-go delivery group that has a scheduled auto scaling policy enabled.
//
// Description:
//
// > This operation is applicable only to pay-as-you-go resource delivery groups that have a scheduled auto scaling policy enabled. It can be called successfully only outside the scaling time windows specified in the scheduled auto scaling policy.
//
// @param request - LogOffAllSessionsInAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogOffAllSessionsInAppInstanceGroupResponse
func (client *Client) LogOffAllSessionsInAppInstanceGroupWithContext(ctx context.Context, request *LogOffAllSessionsInAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *LogOffAllSessionsInAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LogOffAllSessionsInAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the memory feature for the specified agent instances.
//
// Description:
//
// Call `ListAgentInstanceMemory` to obtain the target instance IDs. The server validates all instances in the request before updating the memory settings. After the update, you can query the memory settings of the instances.
//
// @param request - ModifyAgentInstanceMemoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAgentInstanceMemoryConfigResponse
func (client *Client) ModifyAgentInstanceMemoryConfigWithContext(ctx context.Context, request *ModifyAgentInstanceMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAgentInstanceMemoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentInstanceIds) {
		body["AgentInstanceIds"] = request.AgentInstanceIds
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAgentInstanceMemoryConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAgentInstanceMemoryConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the General Policy of a delivery group, including the number of concurrent sessions and the session retention duration after disconnection.
//
// @param tmpReq - ModifyAppInstanceGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppInstanceGroupAttributeResponse
func (client *Client) ModifyAppInstanceGroupAttributeWithContext(ctx context.Context, tmpReq *ModifyAppInstanceGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppInstanceGroupAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Network) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, dara.String("Network"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodePool) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, dara.String("NodePool"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityPolicy) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, dara.String("SecurityPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoragePolicy) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, dara.String("StoragePolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.NodePoolShrink) {
		query["NodePool"] = request.NodePoolShrink
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NetworkShrink) {
		body["Network"] = request.NetworkShrink
	}

	if !dara.IsNil(request.PerSessionPerApp) {
		body["PerSessionPerApp"] = request.PerSessionPerApp
	}

	if !dara.IsNil(request.PreOpenAppId) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !dara.IsNil(request.PreOpenMode) {
		body["PreOpenMode"] = request.PreOpenMode
	}

	if !dara.IsNil(request.SecurityPolicyShrink) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !dara.IsNil(request.StoragePolicyShrink) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppInstanceGroupAttribute"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of a specified app instance group to the target billing method and returns the ID of the billing conversion task to be created.
//
// Description:
//
// This operation converts the billing method, which may create an order and trigger subsequent billing changes. **A successful response or a returned task ID does not indicate that the new billing method has taken effect.**
//
// **Prerequisites:**
//
// - The resources belong to the current account, and the product and business region match the request. Only one app instance group can be specified per request.
//
// - The target billing method must be different from the current one, and the target resource specifications must support the corresponding sales and Billable methods.
//
// - App instance groups that have been added to a valid collection cannot have their Billable methods converted individually.
//
// - Multiple app instance groups created in the same batch of orders cannot have their Billable methods converted separately.
//
// **Parameter relationships and call sequence:**
//
// 1. Query the resource details first to confirm the product, business region, and current `ChargeType`. You can use [GetAppInstanceGroup](https://api.aliyun.com/document/appstream-center/2021-09-01/GetAppInstanceGroup) to query cloud app instance groups.
//
// 2. Explicitly specify the target `ChargeType`. If you need to purchase a subscription, specify the purchase duration by combining `Period` and `PeriodUnit`. The period combination must comply with the sales rules currently supported by the resource.
//
// 3. Set `AutoPay` based on whether automatic payment is allowed. If auto-renewal is involved, explicitly set `AutoRenew`.
//
// 4. After the call, save the returned `AppInstanceGroupModel.TaskId`. If there are unpaid orders, complete the payment first, and then query the resource details again to confirm that the billing method has been changed.
//
// This operation does not accept an expiration time parameter. Do not pass the expiration time from other operations into this request. The task examples only illustrate the purpose of the ID and do not define a fixed prefix or length for actual task IDs.
//
// @param tmpReq - ModifyAppInstanceGroupChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppInstanceGroupChargeTypeResponse
func (client *Client) ModifyAppInstanceGroupChargeTypeWithContext(ctx context.Context, tmpReq *ModifyAppInstanceGroupChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppInstanceGroupChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppInstanceGroupChargeTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppInstanceGroupIds) {
		request.AppInstanceGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInstanceGroupIds, dara.String("AppInstanceGroupIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupIdsShrink) {
		body["AppInstanceGroupIds"] = request.AppInstanceGroupIdsShrink
	}

	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppInstanceGroupChargeType"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppInstanceGroupChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the display policy of a delivery group, including settings such as frame rate, resolution, and protocol type.
//
// @param tmpReq - ModifyAppPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppPolicyResponse
func (client *Client) ModifyAppPolicyWithContext(ctx context.Context, tmpReq *ModifyAppPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VideoPolicy) {
		request.VideoPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoPolicy, dara.String("VideoPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppPolicyId) {
		query["AppPolicyId"] = request.AppPolicyId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.VideoPolicyShrink) {
		query["VideoPolicy"] = request.VideoPolicyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppPolicy"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of a Wuying cloud browser.
//
// Description:
//
// Modifies the properties of a Wuying cloud browser.
//
// @param tmpReq - ModifyBrowserInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBrowserInstanceGroupResponse
func (client *Client) ModifyBrowserInstanceGroupWithContext(ctx context.Context, tmpReq *ModifyBrowserInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyBrowserInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBrowserInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BrowserConfig) {
		request.BrowserConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BrowserConfig, dara.String("BrowserConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Network) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, dara.String("Network"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StoragePolicy) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, dara.String("StoragePolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Timers) {
		request.TimersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Timers, dara.String("Timers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BrowserConfigShrink) {
		query["BrowserConfig"] = request.BrowserConfigShrink
	}

	if !dara.IsNil(request.BrowserInstanceGroupId) {
		query["BrowserInstanceGroupId"] = request.BrowserInstanceGroupId
	}

	if !dara.IsNil(request.PolicyShrink) {
		query["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.TimersShrink) {
		query["Timers"] = request.TimersShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthNotificationEnabled) {
		body["AuthNotificationEnabled"] = request.AuthNotificationEnabled
	}

	if !dara.IsNil(request.CloudBrowserName) {
		body["CloudBrowserName"] = request.CloudBrowserName
	}

	if !dara.IsNil(request.MaxAmount) {
		body["MaxAmount"] = request.MaxAmount
	}

	if !dara.IsNil(request.NetworkShrink) {
		body["Network"] = request.NetworkShrink
	}

	if !dara.IsNil(request.StoragePolicyShrink) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBrowserInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBrowserInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the number of nodes in a subscription delivery group.
//
// @param tmpReq - ModifyNodePoolAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolAmountResponse
func (client *Client) ModifyNodePoolAmountWithContext(ctx context.Context, tmpReq *ModifyNodePoolAmountRequest, runtime *dara.RuntimeOptions) (_result *ModifyNodePoolAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyNodePoolAmountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodePool) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, dara.String("NodePool"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.NodePoolShrink) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodePoolAmount"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodePoolAmountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the elastic mode of a delivery group, including fixed quantity (no elastic scaling), scheduled scaling, and automatic scaling.
//
// Description:
//
// You can configure the elastic pattern for WUYING Cloud Application resources in Settings:
//
// - Fixed quantity: Elastic scaling is not used.
//
// - Automatic scaling: Automatically scales resources based on the number of connected sessions and the duration without session connections.
//
// - Scheduled scaling: Executes resource scaling during specified time periods on specified dates.
//
// Before using this operation, make sure that you fully understand the [billing method and pricing](https://help.aliyun.com/document_detail/426039.html) of WUYING Cloud Application.
//
// @param tmpReq - ModifyNodePoolAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolAttributeResponse
func (client *Client) ModifyNodePoolAttributeWithContext(ctx context.Context, tmpReq *ModifyNodePoolAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyNodePoolAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyNodePoolAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodePoolStrategy) {
		request.NodePoolStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePoolStrategy, dara.String("NodePoolStrategy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.NodeCapacity) {
		body["NodeCapacity"] = request.NodeCapacity
	}

	if !dara.IsNil(request.NodePoolStrategyShrink) {
		body["NodePoolStrategy"] = request.NodePoolStrategyShrink
	}

	if !dara.IsNil(request.PoolId) {
		body["PoolId"] = request.PoolId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodePoolAttribute"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares an image with other WUYING users or cancels image sharing.
//
// @param request - ModifyShareImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyShareImageResponse
func (client *Client) ModifyShareImageWithContext(ctx context.Context, request *ModifyShareImageRequest, runtime *dara.RuntimeOptions) (_result *ModifyShareImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddAccount) {
		body["AddAccount"] = request.AddAccount
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RemoveAccount) {
		body["RemoveAccount"] = request.RemoveAccount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyShareImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyShareImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the administrator account configuration, such as whether to enable resource expiration reminders.
//
// @param request - ModifyTenantConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTenantConfigResponse
func (client *Client) ModifyTenantConfigWithContext(ctx context.Context, request *ModifyTenantConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyTenantConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupExpireRemind) {
		body["AppInstanceGroupExpireRemind"] = request.AppInstanceGroupExpireRemind
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTenantConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTenantConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the tenant-level memory configuration for the current account based on the agent provider, business platform, and business type.
//
// Description:
//
// *Modifying the switch status also updates existing instances that have the corresponding memory feature configured under the same business combination.*	- After the modification, call `GetTenantMemoryConfig` and pass the same provider, platform, and business type to verify the configuration.
//
// @param request - ModifyTenantMemoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTenantMemoryConfigResponse
func (client *Client) ModifyTenantMemoryConfigWithContext(ctx context.Context, request *ModifyTenantMemoryConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyTenantMemoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTenantMemoryConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTenantMemoryConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of a workstation.
//
// @param request - ModifyWuyingServerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWuyingServerAttributeResponse
func (client *Client) ModifyWuyingServerAttributeWithContext(ctx context.Context, request *ModifyWuyingServerAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyWuyingServerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErdmaEnabled) {
		body["ErdmaEnabled"] = request.ErdmaEnabled
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	if !dara.IsNil(request.WuyingServerName) {
		body["WuyingServerName"] = request.WuyingServerName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWuyingServerAttribute"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWuyingServerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the specifications of an existing WUYING workstation.
//
// Description:
//
// ## Before you begin
//
// - The workstation is created and belongs to the current account.
//
// - When you change the specifications, the resource must be in the Deployed, Published Successfully, or Maintenance Failed state. The target specifications must be available in the region where the resource is located and compatible with the operating system and existing resource configurations.
//
// - Conversion between CPU specifications and GPU specifications is not supported.
//
// ## Post-processing
//
// A successful response from the API does not mean that the specification change is complete. If the generated order is unpaid, complete the payment first, and then query the workstation status and specifications to confirm the change result.
//
// @param request - ModifyWuyingServerSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWuyingServerSpecResponse
func (client *Client) ModifyWuyingServerSpecWithContext(ctx context.Context, request *ModifyWuyingServerSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyWuyingServerSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AfterStatus) {
		body["AfterStatus"] = request.AfterStatus
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.ServerInstanceType) {
		body["ServerInstanceType"] = request.ServerInstanceType
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWuyingServerSpec"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWuyingServerSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the allocated users added to a delivery group by paging.
//
// @param request - PageListAppInstanceGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageListAppInstanceGroupUserResponse
func (client *Client) PageListAppInstanceGroupUserWithContext(ctx context.Context, request *PageListAppInstanceGroupUserRequest, runtime *dara.RuntimeOptions) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageListAppInstanceGroupUser"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries delivery group reports by region, delivery group, status, user authorization relationship, or statistical date.
//
// Description:
//
// Queries the delivery group ID, name, and status in delivery group reports. You can call [ListRegions](~~ListRegions~~) to obtain the business region, and then call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the delivery group. AppInstanceGroupId matches the exact ID, and AppInstanceGroupName filters by inclusion. If multiple filter conditions are specified at the same time, the intersection is used. When filtering by date, StatDate specifies a single day, and StartDate and EndDate must be specified in pairs to represent a closed interval. If both a single day and an interval are specified, the results must meet both conditions. Before using EndUserId, make sure that the user has a valid authorization relationship.
//
// @param request - QueryAppInstanceGroupReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppInstanceGroupReportListResponse
func (client *Client) QueryAppInstanceGroupReportListWithContext(ctx context.Context, request *QueryAppInstanceGroupReportListRequest, runtime *dara.RuntimeOptions) (_result *QueryAppInstanceGroupReportListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupName) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDate) {
		body["StatDate"] = request.StatDate
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAppInstanceGroupReportList"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAppInstanceGroupReportListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of successful connections, the number of failed connections, the success rate, and the average connection duration within a specified time range.
//
// Description:
//
// Queries connection statistical reports. Before you invoke this operation, you can invoke [ListRegions](~~ListRegions~~) and [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the `BizRegionId` and `AppInstanceGroupId` parameters. The `ProductType`, `StartDate`, and `EndDate` parameters are required for server-side validation. The statistical interval is `[StartDate, EndDate)`, which includes the start time but excludes the end time. Results are returned by page, with a maximum of 1,500 entries per page.
//
// @param request - QueryConnectionReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConnectionReportResponse
func (client *Client) QueryConnectionReportWithContext(ctx context.Context, request *QueryConnectionReportRequest, runtime *dara.RuntimeOptions) (_result *QueryConnectionReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConnectionReport"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConnectionReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of daily active users and streaming daily active users by date.
//
// Description:
//
// Queries the daily active user (DAU) report. StartDate and EndDate are required. The date range is inclusive and cannot exceed 180 days. If ProductType is not specified or cannot be recognized, the query defaults to `CloudApp`. You can use BizRegionId and AppInstanceGroupId to narrow down the scope. If neither is specified, the query returns aggregated data for all business regions. Results are returned in pages, with a maximum of 1,500 records per page.
//
// @param request - QueryDauReportDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDauReportDataResponse
func (client *Client) QueryDauReportDataWithContext(ctx context.Context, request *QueryDauReportDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDauReportDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDauReportData"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDauReportDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of monthly active users for each month within a specified time range.
//
// Description:
//
// Queries the monthly active user (MAU) report. ProductType, StartDate, and EndDate are required for server-side validation, and the time span cannot exceed 60 months. The API returns statistical values by month. To adapt to offline data generated on the next day, the last month in the query range uses the day before EndDate as the statistical cutoff date. Results are returned by page, with a maximum of 1,500 entries per page.
//
// @param request - QueryMauReportDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMauReportDataResponse
func (client *Client) QueryMauReportDataWithContext(ctx context.Context, request *QueryMauReportDataRequest, runtime *dara.RuntimeOptions) (_result *QueryMauReportDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMauReportData"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMauReportDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network RTT and low-RTT user statistics within a specified time range.
//
// Description:
//
// Queries network round-trip time (RTT) reports. ProductType, StartDate, and EndDate are required for server-side validation. The statistical interval is [StartDate, EndDate). The minimum, maximum, and average RTT values are measured in milliseconds. The low-RTT threshold is determined by the report statistical caliber, and the current API operation does not return this threshold. Results are returned in pages, with a maximum of 1,500 entries per page.
//
// @param request - QueryNetworkRttReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNetworkRttReportResponse
func (client *Client) QueryNetworkRttReportWithContext(ctx context.Context, request *QueryNetworkRttReportRequest, runtime *dara.RuntimeOptions) (_result *QueryNetworkRttReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryNetworkRttReport"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryNetworkRttReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of nodes, the number of subscription nodes, and the number of pay-as-you-go nodes by statistical granularity.
//
// Description:
//
// Queries node quantity reports. StartDate, EndDate, and StatDateType are required. The query span limit is determined by the statistical granularity: up to 1 day for `1m` and `30m`, up to 15 days for `6h`, and up to 180 days for `1d`. Both the start and end boundaries are included in the query range. Results are returned in pages, with a maximum of 1,500 entries per page.
//
// @param request - QueryNodeReportDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNodeReportDataResponse
func (client *Client) QueryNodeReportDataWithContext(ctx context.Context, request *QueryNodeReportDataRequest, runtime *dara.RuntimeOptions) (_result *QueryNodeReportDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDateType) {
		body["StatDateType"] = request.StatDateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryNodeReportData"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryNodeReportDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the session count trend at various time points on a specified date, returned in 10-minute intervals.
//
// Description:
//
// Call this operation to query the session count trend data for the current Alibaba Cloud account on a specified date. The data consists of offline statistical results organized by day and returns the session counts at various time points on that date.
//
// The returned results are aggregated in 10-minute intervals. Each data point is identified by an hour (Hh, values 00-23) and a minute (Mi, values 00, 10, 20, 30, 40, 50) representing a time point, and SessionCntMi is the aggregated session count within the 10-minute interval for that time point. The results are sorted in ascending order by hour and minute.
//
// You can specify a business region by using BizRegionId and filter by delivery group by using AppInstanceGroupId. If neither parameter is provided, the aggregated session data for all regions and all delivery groups under the current account is returned. The actual region used to filter session data is determined by BizRegionId.
//
// StatDate is the statistical date in the yyyyMMdd format. If this parameter is not provided, the data for yesterday is queried by default. Only dates within the last 90 days are supported. If there is no session data on the specified date, an empty list is returned and the return code is NO_DATA.
//
// @param request - QuerySessionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionDataResponse
func (client *Client) QuerySessionDataWithContext(ctx context.Context, request *QuerySessionDataRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StatDate) {
		body["StatDate"] = request.StatDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionData"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of sessions, connected sessions, and streaming sessions by statistical granularity.
//
// Description:
//
// Queries session quantity reports. `StartDate`, `EndDate`, and `StatDateType` are required. The query span limit is determined by the statistical granularity: up to 1 day for `1m` and `30m`, up to 15 days for `6h`, and up to 180 days for `1d`. Both the start and end boundaries are included in the query range. Results are returned in pages, with a maximum of 1,500 records per page.
//
// @param request - QuerySessionReportDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionReportDataResponse
func (client *Client) QuerySessionReportDataWithContext(ctx context.Context, request *QuerySessionReportDataRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionReportDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDateType) {
		body["StatDateType"] = request.StatDateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionReportData"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionReportDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cumulative duration of application usage by a specified user on an hourly basis and returns the top five applications for each hour.
//
// Description:
//
// Queries the application usage of a specified user on an hourly basis. StartDate and EndDate use the yyyyMMdd format. If either parameter is not provided, it defaults to yesterday. Both dates must be within the last 90 days, and the start date cannot be later than the end date. The query range includes the start and end dates. UserId must be set to the ID of the user to be queried. A maximum of the top five applications by usage duration are returned for each hour.
//
// @param request - QueryUserAppUsageAggregateByHourRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserAppUsageAggregateByHourResponse
func (client *Client) QueryUserAppUsageAggregateByHourWithContext(ctx context.Context, request *QueryUserAppUsageAggregateByHourRequest, runtime *dara.RuntimeOptions) (_result *QueryUserAppUsageAggregateByHourResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserAppUsageAggregateByHour"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserAppUsageAggregateByHourResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ranking of cumulative application usage duration for a specified user within a date range.
//
// Description:
//
// Queries the ranking of application usage duration for a specified user within a statistical range. `StartDate` and `EndDate` use the `yyyyMMdd` format. If either parameter is not provided, it defaults to yesterday. Both dates must be within the last 90 days, and the start date cannot be later than the end date. The query range includes the start and end dates. `UserId` must be set to the ID of the user to be queried. The results are sorted in descending order of cumulative application usage duration.
//
// @param request - QueryUserAppUsageTimeRankRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserAppUsageTimeRankResponse
func (client *Client) QueryUserAppUsageTimeRankWithContext(ctx context.Context, request *QueryUserAppUsageTimeRankRequest, runtime *dara.RuntimeOptions) (_result *QueryUserAppUsageTimeRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserAppUsageTimeRank"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserAppUsageTimeRankResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connection status, last connection time, cumulative usage duration, and session persistence duration of users.
//
// Description:
//
// Queries user usage reports. `ProductType`, `StartDate`, and `EndDate` are required for server-side validation, and the time span cannot exceed 60 months. During the query, the start and end times are converted to the `yyyyMMdd` format and aggregated in a closed interval of natural days. The results are paginated in descending order of cumulative usage duration, with a maximum of 1,500 records per page. You can combine `ConnectionStatus`, `UserId`, `BizRegionId`, and `AppInstanceGroupId` for filtering.
//
// @param request - QueryUserUsageReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserUsageReportResponse
func (client *Client) QueryUserUsageReportWithContext(ctx context.Context, request *QueryUserUsageReportRequest, runtime *dara.RuntimeOptions) (_result *QueryUserUsageReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ConnectionStatus) {
		body["ConnectionStatus"] = request.ConnectionStatus
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserUsageReport"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserUsageReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a model group from resource groups.
//
// Description:
//
// You can authorize model groups for the resources to which agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent belong in the Wuying Agent Management Center. These model groups serve as inference engines for agents in the resource groups to execute tasks.
//
// If an agent runtime is configured with a model group and its resource group is also configured with a model group, the model group bound to the resource group takes effect. That is, the resource group settings take precedence over the agent runtime settings.
//
// When you remove the model group from the resource group to which the agent runtime belongs, the model group configured for the agent runtime automatically takes effect.
//
// Make sure that you fully understand the operations and usage of the Wuying Agent Management Center before you call this operation.
//
// @param request - RemoveResourceGroupModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveResourceGroupModelTemplateResponse
func (client *Client) RemoveResourceGroupModelTemplateWithContext(ctx context.Context, request *RemoveResourceGroupModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *RemoveResourceGroupModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		body["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		body["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveResourceGroupModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveResourceGroupModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the third-party channel configuration for an agent runtime.
//
// Description:
//
// You can call this operation to remove the specific third-party channel configuration for agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the WUYING Agent Management Center. After the configuration is removed, the agent cannot use the third-party channel for conversations.
//
// Make sure that you are fully familiar with the operations and usage of the WUYING Agent Management Center before you call this operation.
//
// Prerequisites:
//
// - The RuntimeIds passed to this operation must be agent runtimes managed by the WUYING Agent Management Center. The hosted cloud desktops are created by API operations such as CreateDesktop of the Elastic Desktop Service (EDS) product. Example of a cross-product call chain: CreateDesktop → ListDesktops → Use DesktopId as RuntimeIds.
//
// - OpenClaw scenario: RuntimeIds are cloud desktop IDs in the format of ecd-xxxx, which are the DesktopId values returned by the ListDesktops operation of the EDS product.
//
// - JVS Computer and Hermes Agent scenarios: RuntimeIds are the corresponding JVS IDs in the format of jvs-xxxx. The hosted cloud desktops are also created by API operations such as CreateDesktop of the EDS product.
//
// - You can call the ListDesktopAgentRuntime operation of this product to query the agent runtimes managed by the WUYING Agent Management Center. If the query result is empty, the corresponding cloud desktop is not managed. You must first call the CreateDesktop operation of the EDS product to create the desktop. The WUYING Agent Management Center automatically manages the desktop after receiving the desktop lifecycle event. Then, you can call this operation.
//
// @param request - RemoveRuntimeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRuntimeChannelResponse
func (client *Client) RemoveRuntimeChannelWithContext(ctx context.Context, request *RemoveRuntimeChannelRequest, runtime *dara.RuntimeOptions) (_result *RemoveRuntimeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.RuntimeIds) {
		body["RuntimeIds"] = request.RuntimeIds
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRuntimeChannel"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRuntimeChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the model group from an agent runtime resource.
//
// Description:
//
// You can remove the model group from agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. When an agent runtime resource needs to change its model group, you must first call this operation to remove the authorization relationship between the agent runtime resource and the original model group.
//
// Make sure that you are fully familiar with the operations and usage of the Wuying Agent Management Center before you use this operation.
//
// @param request - RemoveRuntimeModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRuntimeModelTemplateResponse
func (client *Client) RemoveRuntimeModelTemplateWithContext(ctx context.Context, request *RemoveRuntimeModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *RemoveRuntimeModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ModelTemplateId) {
		body["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.RuntimeIds) {
		body["RuntimeIds"] = request.RuntimeIds
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRuntimeModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRuntimeModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a delivery group.
//
// Description:
//
// Make sure you understand the [billing methods and pricing](https://help.aliyun.com/document_detail/426039.html) of WUYING Workspace before calling this operation.
//
// > Prerequisites:
//
// > - The delivery group must be in the PUBLISHED state and ChargeType must be PrePaid.
//
// @param tmpReq - RenewAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppInstanceGroupResponse
func (client *Client) RenewAppInstanceGroupWithContext(ctx context.Context, tmpReq *RenewAppInstanceGroupRequest, runtime *dara.RuntimeOptions) (_result *RenewAppInstanceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RenewAppInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RenewNodes) {
		request.RenewNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RenewNodes, dara.String("RenewNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RenewAmount) {
		query["RenewAmount"] = request.RenewAmount
	}

	if !dara.IsNil(request.RenewMode) {
		query["RenewMode"] = request.RenewMode
	}

	if !dara.IsNil(request.RenewNodesShrink) {
		query["RenewNodes"] = request.RenewNodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAppInstanceGroup"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAppInstanceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a workstation.
//
// @param request - RenewWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewWuyingServerResponse
func (client *Client) RenewWuyingServerWithContext(ctx context.Context, request *RenewWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *RenewWuyingServerResponse, _err error) {
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

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		body["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resizes the system cloud disk or user data cloud disk of a WUYING Workspace Enterprise Edition instance and returns the related order information.
//
// Description:
//
// This operation resizes the disks of a WUYING Workspace.
//
// @param tmpReq - ResizeWuyingServerDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeWuyingServerDiskResponse
func (client *Client) ResizeWuyingServerDiskWithContext(ctx context.Context, tmpReq *ResizeWuyingServerDiskRequest, runtime *dara.RuntimeOptions) (_result *ResizeWuyingServerDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResizeWuyingServerDiskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataDiskList) {
		request.DataDiskListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataDiskList, dara.String("DataDiskList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskListShrink) {
		body["DataDiskList"] = request.DataDiskListShrink
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SystemDiskNewSize) {
		body["SystemDiskNewSize"] = request.SystemDiskNewSize
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeWuyingServerDisk"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeWuyingServerDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a workstation.
//
// @param request - RestartWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartWuyingServerResponse
func (client *Client) RestartWuyingServerWithContext(ctx context.Context, request *RestartWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *RestartWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableRestartRetry) {
		query["EnableRestartRetry"] = request.EnableRestartRetry
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.WuyingServerIdList) {
		bodyFlat["WuyingServerIdList"] = request.WuyingServerIdList
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues Cloud Assistant commands to persistent sessions in WUYING Workspace delivery groups. This operation supports issuing commands to a single delivery group or multiple delivery groups in the same region, and returns the command execution ID.
//
// Description:
//
// Issues Cloud Assistant commands to persistent sessions in WUYING Workspace delivery groups. Currently, the ProductType parameter only supports WuyingServer, and the CommandType parameter only supports CloudAssistantCmd. Specifying other values results in a parameter error.
//
// You can specify the execution scope in two ways:
//
// - Single delivery group: Specify the delivery group by using AppInstanceGroupId. You can also specify persistent sessions within the delivery group by using AppInstancePersistentId or AppInstancePersistentIds.
//
// - Multiple delivery groups in the same region: Specify a list of delivery groups by using AppInstanceGroupIds. The command is executed on all persistent sessions in each delivery group in the list. All delivery groups must belong to the same region. Otherwise, the InvalidRegion.NotMatch error is returned. If any delivery group does not exist or none of the delivery groups have persistent sessions, the InvalidAppInstanceGroup.NotFound error is returned.
//
// You must specify at least one of AppInstanceGroupIds and AppInstanceGroupId. If both are specified, AppInstanceGroupIds takes precedence.
//
// This operation only issues the command. The command is executed asynchronously in the sessions. Regardless of whether the command is issued to one or multiple delivery groups, only one command execution ID (InvocationId) is returned. You can use this ID to call the ListCommandInvocationResults operation to query the execution status and output on each session.
//
// @param request - RunAppInstanceCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAppInstanceCommandResponse
func (client *Client) RunAppInstanceCommandWithContext(ctx context.Context, request *RunAppInstanceCommandRequest, runtime *dara.RuntimeOptions) (_result *RunAppInstanceCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupIds) {
		body["AppInstanceGroupIds"] = request.AppInstanceGroupIds
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !dara.IsNil(request.AppInstancePersistentIds) {
		body["AppInstancePersistentIds"] = request.AppInstancePersistentIds
	}

	if !dara.IsNil(request.CommandContent) {
		body["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.CommandParam) {
		body["CommandParam"] = request.CommandParam
	}

	if !dara.IsNil(request.CommandType) {
		body["CommandType"] = request.CommandType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RunScriptType) {
		body["RunScriptType"] = request.RunScriptType
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAppInstanceCommand"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunAppInstanceCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates the third-party channel QR code configuration process in batches for the specified cloud desktop runtimes, generates the URLs of the QR code configuration pages, and notifies the authorized users of the runtimes by email and SMS.
//
// Description:
//
// Initiates the third-party channel QR code configuration process in batches for managed cloud desktop runtimes in the Wuying Agent Management Center. The system generates a QR code configuration page URL for each runtime and notifies the authorized user of the runtime by email and SMS. The user can complete the channel configuration by opening the page and scanning the QR code.
//
// Usage notes:
//
// - Currently, only CloudDesktop is supported for RuntimeType. You can specify up to 200 runtime IDs in a single request.
//
// - Each runtime must have exactly one authorized user. Otherwise, the sending result for the runtime returns the Runtime.UserNotSupported error code.
//
// - The taskId in the QR code configuration page URL is valid for 1 hour. After it expires, you must call this operation again to initiate the process.
//
// - If the sending fails for some runtimes, the operation still returns a success response overall. Check the Success, Code, and Message fields for each runtime in Data. If the sending fails for all runtimes, the operation returns a failure response. If the error codes for all runtimes are the same, that error code is returned. Otherwise, SendQrCodeConfigMessage.SendFailed is returned.
//
// - Email and SMS notifications are sent on a best-effort basis. Notification failures do not affect the API response. After initiation, you can call GetRuntimeChannelQrCodeDetail to query the details of the QR code configuration process.
//
// @param request - SendRuntimeQrCodeConfigChannelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRuntimeQrCodeConfigChannelMessageResponse
func (client *Client) SendRuntimeQrCodeConfigChannelMessageWithContext(ctx context.Context, request *SendRuntimeQrCodeConfigChannelMessageRequest, runtime *dara.RuntimeOptions) (_result *SendRuntimeQrCodeConfigChannelMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentPlatform) {
		body["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		body["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.RuntimeIds) {
		body["RuntimeIds"] = request.RuntimeIds
	}

	if !dara.IsNil(request.RuntimeType) {
		body["RuntimeType"] = request.RuntimeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendRuntimeQrCodeConfigChannelMessage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendRuntimeQrCodeConfigChannelMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets user data roaming path rules for an app instance group or an app instance group set.
//
// Description:
//
// First, call `DescribeUserProfilePathRules` to query the target rules, and then submit the complete list of rules that you want to retain. **The submission overwrites the existing rules of the target. Passing an empty array or omitting the rule list sets the target rules to an empty list.*	- When a set ID is used, the rules apply to the members of the set. After the rules are configured, you can call the query operation to verify the settings.
//
// @param request - SetUserProfilePathRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserProfilePathRulesResponse
func (client *Client) SetUserProfilePathRulesWithContext(ctx context.Context, request *SetUserProfilePathRulesRequest, runtime *dara.RuntimeOptions) (_result *SetUserProfilePathRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceGroupSetId) {
		body["AppInstanceGroupSetId"] = request.AppInstanceGroupSetId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserProfilePathRules) {
		body["UserProfilePathRules"] = request.UserProfilePathRules
	}

	if !dara.IsNil(request.UserProfileRuleType) {
		body["UserProfileRuleType"] = request.UserProfileRuleType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserProfilePathRules"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserProfilePathRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an image creation task for an existing image version, or resumes an existing creation task.
//
// @param request - StartTaskForCreateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskForCreateImageResponse
func (client *Client) StartTaskForCreateImageWithContext(ctx context.Context, request *StartTaskForCreateImageRequest, runtime *dara.RuntimeOptions) (_result *StartTaskForCreateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageCreateMode) {
		body["ImageCreateMode"] = request.ImageCreateMode
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JumpStep) {
		body["JumpStep"] = request.JumpStep
	}

	if !dara.IsNil(request.PerformanceType) {
		body["PerformanceType"] = request.PerformanceType
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceList) {
		body["ResourceList"] = request.ResourceList
	}

	if !dara.IsNil(request.RetryType) {
		body["RetryType"] = request.RetryType
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.SourceRegion) {
		body["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SubInstanceId) {
		body["SubInstanceId"] = request.SubInstanceId
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTaskForCreateImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTaskForCreateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task to copy an image to other regions.
//
// @param request - StartTaskForDistributeImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTaskForDistributeImageResponse
func (client *Client) StartTaskForDistributeImageWithContext(ctx context.Context, request *StartTaskForDistributeImageRequest, runtime *dara.RuntimeOptions) (_result *StartTaskForDistributeImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationRegionList) {
		body["DestinationRegionList"] = request.DestinationRegionList
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RetryType) {
		body["RetryType"] = request.RetryType
	}

	if !dara.IsNil(request.SourceRegion) {
		body["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTaskForDistributeImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTaskForDistributeImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts one or more workstations.
//
// @param request - StartWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartWuyingServerResponse
func (client *Client) StartWuyingServerWithContext(ctx context.Context, request *StartWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *StartWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.WuyingServerIdList) {
		bodyFlat["WuyingServerIdList"] = request.WuyingServerIdList
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a workstation.
//
// @param request - StopWuyingServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopWuyingServerResponse
func (client *Client) StopWuyingServerWithContext(ctx context.Context, request *StopWuyingServerRequest, runtime *dara.RuntimeOptions) (_result *StopWuyingServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.WuyingServerIdList) {
		bodyFlat["WuyingServerIdList"] = request.WuyingServerIdList
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopWuyingServer"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopWuyingServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms the completion of image creation.
//
// @param request - SubmitImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageResponse
func (client *Client) SubmitImageWithContext(ctx context.Context, request *SubmitImageRequest, runtime *dara.RuntimeOptions) (_result *SubmitImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationRegionList) {
		body["DestinationRegionList"] = request.DestinationRegionList
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and binds tags to cloud resources, and updates the values of existing tags.
//
// @param request - TagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagCloudResourcesResponse
func (client *Client) TagCloudResourcesWithContext(ctx context.Context, request *TagCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagCloudResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagCloudResources"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagCloudResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and binds tags to specified app instance groups according to Alibaba Cloud tag specifications. If a tag with the same key already exists on a resource, the tag value is updated to the value specified in the current request. Currently, only the app instance group (`APPINSTANCEGROUP`) resource type is supported.
//
// Description:
//
// ## Prerequisites
//
// When you call this operation as a RAM user or by using STS credentials, you must have the `appstreaming:TagResources` permission.
//
// The resources to which you want to bind tags must belong to the current Alibaba Cloud account. **This operation currently supports only app instance group resources**. You must set ResourceType to `APPINSTANCEGROUP`. You can obtain the app instance group ID by calling [ListAppInstanceGroup](~~ListAppInstanceGroup~~).
//
// ## Tag rules
//
// - A single request can bind up to **20 tags*	- to up to **50 resources**. Duplicate resource IDs are automatically deduplicated.
//
// - A tag key must be 1 to 128 characters in length, and a tag value must be 0 to 256 characters in length. Both are case-sensitive.
//
// - A tag key cannot start with `aliyun` or `acs:` (case-insensitive). Neither tag keys nor tag values can contain `http://` or `https://`.
//
// - Tag keys in the same request cannot be duplicate. Otherwise, the error code `InvalidTag.Duplicated` is returned.
//
// - A tag key on a resource can correspond to only one tag value. If the tag key already exists on the resource, the tag value is updated to the new value.
//
// - A single resource can be bound to up to 20 custom tags. If this limit is exceeded, the error code `ResourceTag.CustomTagCountExceed` is returned.
//
// ## Execution results
//
// - **If any specified app instance group does not exist or does not belong to the current account, the entire request fails**, and the error code `InvalidAppInstanceGroup.NotFound` is returned. No tags are bound to any resources.
//
// - If the binding is successful, the response contains only RequestId and does not return tag details.
//
// - If multiple resources are specified and the binding fails for only some of the resources, the operation still returns a success response, and the response does not contain failure details. Call ListTagResources to verify the binding results. If only one resource is specified and the binding fails, the operation directly returns the corresponding error code.
//
// ## Call sequence
//
// 1. Call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the IDs of the app instance groups to which you want to bind tags.
//
// 2. Call this operation and specify ResourceType=`APPINSTANCEGROUP`, ResourceId.N, and Tag.N.Key/Tag.N.Value.
//
// 3. To view the tags bound to a resource, call ListTagResources. To unbind tags, call UntagResources.
//
// The masked identifiers in the examples are used to demonstrate the format. Replace them with your actual identifiers when you call the operation.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2021-09-01"),
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
// Disassociates secondary private IP addresses from a development host.
//
// @param request - UnassignWuyingServerPrivateAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassignWuyingServerPrivateAddressesResponse
func (client *Client) UnassignWuyingServerPrivateAddressesWithContext(ctx context.Context, request *UnassignWuyingServerPrivateAddressesRequest, runtime *dara.RuntimeOptions) (_result *UnassignWuyingServerPrivateAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddresses) {
		body["PrivateIpAddresses"] = request.PrivateIpAddresses
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassignWuyingServerPrivateAddresses"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassignWuyingServerPrivateAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a user from a session.
//
// @param request - UnbindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithContext(ctx context.Context, request *UnbindRequest, runtime *dara.RuntimeOptions) (_result *UnbindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Unbind"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates multiple applications from a specified image version.
//
// Description:
//
// Only custom images and community images manageable by the current account are supported. If VersionId is omitted, the current active version is used. If specified, the version must belong to the image specified by ImageId.
//
// Verify the association between the applications and the version through image query results before submitting AppIdList. **Disassociating an application does not delete the application record.*	- To delete an application, call DeleteApp.
//
// @param request - UnbindAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAppResponse
func (client *Client) UnbindAppWithContext(ctx context.Context, request *UnbindAppRequest, runtime *dara.RuntimeOptions) (_result *UnbindAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AppIdList) {
		bodyFlat["AppIdList"] = request.AppIdList
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindApp"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates tags from cloud resources.
//
// @param request - UntagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagCloudResourcesResponse
func (client *Client) UntagCloudResourcesWithContext(ctx context.Context, request *UntagCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagCloudResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		body["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagCloudResources"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagCloudResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from the specified Wuying app instance groups according to Alibaba Cloud tag specifications. You can remove specific tags by tag key, or remove all custom tags from the resources at once by setting All to true without specifying tag keys. Currently, only the app instance group (APPINSTANCEGROUP) resource type is supported.
//
// Description:
//
// ## Prerequisites
//
// When you call this operation as a RAM user or by using STS credentials, you must have the appstreaming:UntagResources permission.
//
// The resources from which you want to remove tags must belong to the current Alibaba Cloud account. **This operation currently supports only app instance group resources**. You must set ResourceType to APPINSTANCEGROUP. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the app instance group ID, and call the [ListTagResources](~~ListTagResources~~) operation to query the tags bound to the resources.
//
// ## metric description
//
// - A single request can remove labels from up to **50 resources**. Duplicate resource IDs are automatically deduplicated.
//
// - **You must specify at least one of TagKey.N and All**: If neither is specified, or if TagKey.N is not specified and All is set to false, the error code InvalidParameter.TagKeyListOrAll is returned.
//
//   - If TagKey.N is specified: Only the labels corresponding to these label keys are removed. A maximum of 20 label keys can be specified in a single request. In this case, the All parameter is ignored.
//
//   - If TagKey.N is not specified and All is set to true: All custom labels on the resources are removed, including Wuying system labels that start with `System/` and are attached by calling the [TagResources](~~TagResources~~) operation.
//
// - If a specified label key does not exist on a resource, the label key is skipped and no fault is returned.
//
// - Label keys that start with `System/` are Wuying system labels. Only System/Scheduler/GRAYSCALE and System/Scheduler/STOP_NEW_USER_CONNECTION are supported. If you specify other label keys that start with `System/`, the error code InvalidTagPolicy.KeyInvalid or InvalidTag.SystemKeyNotAllow is returned.
//
// ## Execution results
//
// - **If any specified app instance group does not exist or does not belong to the current account, the entire request fails**, and the error code InvalidAppInstanceGroup.NotFound is returned. No tags are removed from any resources.
//
// - If tags are successfully removed, the response contains only RequestId and does not return tag details.
//
// - If multiple resources are specified and tags are removed from only some of the resources, the operation still returns a success response, and the response does not contain failure details. You can call the [ListTagResources](~~ListTagResources~~) operation to verify the tag removal results. If only one resource is specified and tag removal fails, the operation directly returns the corresponding error code.
//
// ## Call sequence
//
// 1. Call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the app instance group ID. To confirm the tag keys bound to the resources, call the [ListTagResources](~~ListTagResources~~) operation.
//
// 2. Call this operation, set ResourceType to APPINSTANCEGROUP, specify ResourceId.N, and specify TagKey.N or set All to true to indicate the tags to be removed.
//
// 3. To bind tags again, call the [TagResources](~~TagResources~~) operation.
//
// The desensitized identifiers in the examples are used to demonstrate the format. Replace them with your actual identifiers when you call the operation.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		body["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		body["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2021-09-01"),
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
// Modifies the runtime parameters, optimization configurations, and user configurations of a specified application adapter.
//
// Description:
//
// The AdaptorId must correspond to a valid application adapter manageable by the current account. Pass only the configurations that need to be modified in this request.
//
// SilenceParam is used for application startup parameters. UserConfig is used for application runtime configurations. **This operation does not support modifying application scripts or executable file paths through AppScript or ExecuteFileAddress. Do not pass these two parameters.**
//
// @param tmpReq - UpdateAdaptorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdaptorResponse
func (client *Client) UpdateAdaptorWithContext(ctx context.Context, tmpReq *UpdateAdaptorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdaptorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAdaptorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserConfig) {
		request.UserConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserConfig, dara.String("UserConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AdaptorId) {
		body["AdaptorId"] = request.AdaptorId
	}

	if !dara.IsNil(request.AppScript) {
		body["AppScript"] = request.AppScript
	}

	if !dara.IsNil(request.ExecuteFileAddress) {
		body["ExecuteFileAddress"] = request.ExecuteFileAddress
	}

	if !dara.IsNil(request.OptimizeParam) {
		body["OptimizeParam"] = request.OptimizeParam
	}

	if !dara.IsNil(request.SilenceParam) {
		body["SilenceParam"] = request.SilenceParam
	}

	if !dara.IsNil(request.UserConfigShrink) {
		body["UserConfig"] = request.UserConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdaptor"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdaptorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the image of a delivery group.
//
// Description:
//
//	Warning: After the image update starts, sessions of end users who are accessing cloud applications will be disconnected. Proceed with caution to avoid data loss for end users.
//
// > Before calling this API, the delivery group must be in the PUBLISHED, DEPLOYED, or MAINTAIN_FAILED state. You can call GetAppInstanceGroup to query the current state of the delivery group.
//
// > After the update is published, you typically need to wait about 2 minutes for the changes to take effect on the client.
//
// @param request - UpdateAppInstanceGroupImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppInstanceGroupImageResponse
func (client *Client) UpdateAppInstanceGroupImageWithContext(ctx context.Context, request *UpdateAppInstanceGroupImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCenterImageId) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppInstanceGroupImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, version information, and associated application adapters of an image.
//
// Description:
//
// You can only modify images that are owned by the current account and allow modifications. Shared images cannot be modified.
//
// You must provide at least one of ImageId and VersionId. If both are provided, they must point to the same image. If VersionId is omitted, the current effective version of the image is used. Passing a non-empty AdaptorIdList updates the application adapter associations of the selected version based on the complete list. Make sure to include the adapters that you want to retain.
//
// @param request - UpdateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageResponse
func (client *Client) UpdateImageWithContext(ctx context.Context, request *UpdateImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdaptorIdList) {
		body["AdaptorIdList"] = request.AdaptorIdList
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Distro) {
		body["Distro"] = request.Distro
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		body["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	if !dara.IsNil(request.VersionName) {
		body["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a status update request for an existing image creation task to the corresponding creation resource.
//
// Description:
//
// The task must exist before you call this operation, and the corresponding creation resource and resource node must be available. The TaskType must match the actual type of the task. **This operation does not support modifying the task step, image version, or destination region by using TaskStep, VersionId, or BizRegionId. Do not pass these three parameters.*	- A successful request only indicates that the status update call is completed. To query the final status of the task, call GetTask.
//
// @param request - UpdateImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageTaskResponse
func (client *Client) UpdateImageTaskWithContext(ctx context.Context, request *UpdateImageTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskStep) {
		body["TaskStep"] = request.TaskStep
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageTask"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model provider template.
//
// Description:
//
// You can update a model provider template that has been created in the WUYING Agent Management Center, including the template name, description, model service connection configuration, and WUYING security proxy switch. Partial field updates are supported. You only need to pass in the fields that you want to modify.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
//
// @param tmpReq - UpdateModelProviderTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelProviderTemplateResponse
func (client *Client) UpdateModelProviderTemplateWithContext(ctx context.Context, tmpReq *UpdateModelProviderTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelProviderTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateModelProviderTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		query["Config"] = request.ConfigShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableWuyingProxy) {
		body["EnableWuyingProxy"] = request.EnableWuyingProxy
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProviderTemplateId) {
		body["ProviderTemplateId"] = request.ProviderTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelProviderTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelProviderTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model template.
//
// Description:
//
// You can update a model group that has been created in the Wuying Agent Management Center, including the group name, description, and model configuration information. You can modify the default model of a model group by updating the Config field. The updated configuration automatically takes effect on associated cloud desktops.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
//
// @param request - UpdateModelTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelTemplateResponse
func (client *Client) UpdateModelTemplateWithContext(ctx context.Context, request *UpdateModelTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModelTemplateId) {
		query["ModelTemplateId"] = request.ModelTemplateId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RefScope) {
		query["RefScope"] = request.RefScope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelTemplate"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workstation image.
//
// @param request - UpdateWuyingServerImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWuyingServerImageResponse
func (client *Client) UpdateWuyingServerImageWithContext(ctx context.Context, request *UpdateWuyingServerImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateWuyingServerImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWuyingServerImage"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWuyingServerImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the LLM template configuration.
//
// @param request - ValidateLlmTemplateConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateLlmTemplateConfigResponse
func (client *Client) ValidateLlmTemplateConfigWithContext(ctx context.Context, request *ValidateLlmTemplateConfigRequest, runtime *dara.RuntimeOptions) (_result *ValidateLlmTemplateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.ProviderTemplateId) {
		query["ProviderTemplateId"] = request.ProviderTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateLlmTemplateConfig"),
		Version:     dara.String("2021-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateLlmTemplateConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
