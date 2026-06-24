// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sets the execution time for an over-the-air update.
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
// 为研发主机分配辅助私有IP
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
// Add or remove assigned users for a delivery group. Only users added as assigned users can access cloud applications.
//
// Description:
//
// > After changing the assigned users, the selected users will receive corresponding notification emails. Generally, it takes about 2 minutes for the changes to take effect on the client.
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
// Creates LLM templates in batches.
//
// Description:
//
// You can create model templates in batches under a model provider template in the Wuying Agent Management Center. You can add multiple models at a time and specify one of them as the default model. Existing models are automatically skipped and are not created again.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// Configures the model group for a resource group.
//
// Description:
//
// You can assign a model group to resources associated with agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the WUYING Agent Management Center. The model group serves as the inference engine for tasks executed by agents within the resource group.
//
// When both an agent runtime and its resource group have model groups configured, the model group bound to the resource group takes effect. The resource group setting takes priority over the agent runtime setting.
//
// Before calling this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// Configures a third-party channel for Agent runtime.
//
// Description:
//
// You can configure third-party channels for Agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. These channels serve as extended Agent communication methods beyond the AgentIM channel.
//
// Before using this operation, make sure you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// Configures model groups for Agent runtime resources.
//
// Description:
//
// You can authorize model groups for Agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. The model groups serve as inference engines for Agent task execution.
//
// Make sure that you are familiar with the operations and usage of the Wuying Agent Management Center before calling this operation.
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
// Creates a delivery group.
//
// Description:
//
// Make sure that you are familiar with the [billing and pricing](https://help.aliyun.com/document_detail/426039.html) of WUYING Cloud Application before you call this operation.
//
// A delivery group is a logical grouping for delivering cloud applications to end users. It includes the underlying cloud application resources, images that contain cloud applications, resource management policies, and user assignment settings. For details, see [Publish a delivery group](https://help.aliyun.com/document_detail/426046.html).
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
// Creates a custom image from a deployed WUYING instance. You can use the custom image to quickly create more WUYING instances with the same configurations, without having to repeatedly configure the instance environment each time.
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
// Creates a new image from a debug delivery group.
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
// # Create Model Provider Template
//
// Description:
//
// You can create a model provider template under a model template in the Wuying Agent Management Center. This template is used to configure the connection information and keys for model services (such as Alibaba Cloud Bailian, Token Plan, and Moonshot) that Agents can call. After creation, the model provider template is automatically associated with the specified model template.
//
// Make sure you are fully familiar with the operations and usage of the Wuying Agent Management Center before calling this API.
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
// Creates a model creation template.
//
// Description:
//
// You can create a model group in the WUYING Agent Management Center to manage the model providers and model scope that an Agent can invoke. After creation, you can attach the model group to a cloud computer as the inference engine configuration for Agent task execution.
//
// Make sure that you are familiar with the operations and usage of the WUYING Agent Management Center before calling this operation.
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
// 1. A project corresponds to the resource configuration module in the CloudFlow console.
//
// 2. If the ContentId specified in the request parameters has multiple versions, this API operation <notice>uses the default version</notice> for binding.
//
// 3. This operation succeeds only when the default version of the content is in an available state.
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
// Deletes a pay-as-you-go resource-based delivery group.
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
// Deletes a specified application instance.
//
// Description:
//
// Only instances in the init or idle state can be deleted. This operation is available only to specific customers.
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
// Deletes a custom AppStream image.
//
// Description:
//
// - You can delete only custom images that belong to you.
//
// - For images associated with the AppStream Cloud Computer Pool, AppStream Cloud Application, or AppStream Workstation product lines, you must ensure that no AppStream instances are using the image before you can delete it.
//
// - If an AppStream Cloud Desktop template references an image, the template is also deleted when the image is deleted.
//
// - If an image is available in multiple regions, deleting the image removes it from all regions.
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
// You can delete a model template that has been created under a model provider template in the Wuying Agent Management Center. Before deletion, ensure that the model is not the default model of an associated model group. Otherwise, the deletion fails. After deletion, the model configurations of associated cloud computers are automatically refreshed.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// You can delete a model provider template that has been created under model templates in the WUYING Agent Management Center. Before deletion, make sure that the model provider is not the provider of the default model and is not a system preset type provider (such as WUYING credits package). After deletion, the associated models and key configurations are also removed.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// You can delete a model group that has been created in the WUYING Agent Management Center. Before deletion, ensure that the template has not been authorized to any resource. Otherwise, the deletion fails. After deletion, the model providers and models under the model group are also removed.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// Deletes a cloud graphics workstation.
//
// Description:
//
// Deletes a cloud graphics workstation.
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
// # Configure LogShipper for Simple Log Service
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
// 查询研发主机详情
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
// Queries the Elastic IP Address (EIP) information of a Wuying workspace.
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
// Retrieves connection credentials for a cloud application.
//
// Description:
//
// This operation requires multiple invokes (at least two) to obtain the connection credentials.
//
// On the first invoke, an application instance is allocated to the specified convenience account and the application is started. A startup task ID (`TaskID`) is returned.
//
// On subsequent invokes, pass the `TaskID` request parameter to query whether the task is complete. When the returned task status (`TaskStatus`) is completed (`Finished`), the connection credentials (`Ticket`) are also returned.
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
// Queries the details of a model provider template.
//
// Description:
//
// You can query the details of a specified model provider template in the WUYING Agent Management Center, including the provider name, description, and connection configuration list.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// Queries the details of an over-the-air update task, including the available version and version description.
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
// Queries the third-party channel configurations of an Agent runtime.
//
// Description:
//
// You can query the third-party channel configuration status of Agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// Queries the model configuration details of a cloud computer.
//
// Description:
//
// You can query the model configuration details currently bound to a specified cloud computer in the Wuying Agent Management Center, including model groups, model provider lists, and associated model information. After you enable the risk information mode, you can also identify differences between the end user\\"s actual configuration and the configuration delivered by the administrator.
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
// Queries the details of multiple delivery groups. This operation does not specify a particular delivery group but queries the details of all delivery groups that meet the specified conditions.
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
// Queries the details of session instances in a delivery group, including instance IDs, instance statuses, creation time, update time, session statuses, and public IP addresses of primary network interfaces.
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
// Queries the list of user groups authorized by a specified delivery group.
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
// Queries the list of desktop agent runtimes.
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
	if !dara.IsNil(request.AgentInstanceStatuses) {
		query["AgentInstanceStatuses"] = request.AgentInstanceStatuses
	}

	if !dara.IsNil(request.AgentInstanceVersions) {
		query["AgentInstanceVersions"] = request.AgentInstanceVersions
	}

	if !dara.IsNil(request.AgentPlatform) {
		query["AgentPlatform"] = request.AgentPlatform
	}

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
	}

	if !dara.IsNil(request.AuthUsers) {
		query["AuthUsers"] = request.AuthUsers
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
// Queries a list of LLM templates.
//
// Description:
//
// You can use paging to retrieve the list of model templates under a model provider template in the Wuying Agent Management Center. You can filter results by model group ID, model provider template ID, model template ID, and model encoding. When you query by model group dimension, the default model is automatically pinned to the top.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// 查询模型提供商 Endpoint 列表
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
// Queries the list of model provider templates.
//
// Description:
//
// You can perform a paged query to retrieve the list of model provider templates under a specified model group in the WUYING Agent Management Center. You can filter results by provider name, model group ID, and provider template ID. Paging is supported.
//
// Before calling this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// You can call this operation to query the list of resource groups authorized by a model group in the Wuying Agent Management Center.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// You can use paged query to retrieve model groups that have been created in the Wuying Agent Management Center, with paging support. You can filter results by Agent provider, Agent platform, template group ID, and whether models have been configured.
//
// Before calling this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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

	if !dara.IsNil(request.AgentProvider) {
		query["AgentProvider"] = request.AgentProvider
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
// Queries the regions supported by WUYING Cloud Application.
//
// Description:
//
// > The regions returned by this operation are not necessarily all available regions. For information about available regions, see [Supported regions](https://help.aliyun.com/document_detail/426036.html).
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
// Queries the tag list of one or more specified cloud resources.
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
// Queries the list of workstations.
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

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
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
// Logs off all sessions in a pay-as-you-go delivery group that has scheduled auto scaling policies enabled.
//
// Description:
//
// > This operation is applicable only to pay-as-you-go resource delivery groups that have scheduled auto scaling policies enabled, and can be called successfully only outside the scaling time periods configured in the scheduled auto scaling policies.
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
// Modifies the attributes of a Wuying Cloud Browser.
//
// Description:
//
// Modifies the attributes of a Wuying Cloud Browser.
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
// Modifies the scaling mode of a delivery group, including fixed quantity (no elastic scaling), scheduled scaling, and automatic scaling.
//
// Description:
//
// You can configure the scaling pattern for WUYING Cloud Application resources in Settings:
//
// - Fixed quantity: Elastic scaling is not used.
//
// - Automatic scaling: Automatically scales resources based on the number of connected sessions and the idle duration without session connections.
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
// Modifies the configuration of an administrator account, such as whether to enable resource expiration reminders.
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
// Modifies the properties of a cloud graphics workstation.
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
// Performs a paged query on allocated users added to a delivery group.
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
// Removes model groups from a resource group.
//
// Description:
//
// You can authorize model groups for resources that belong to Agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the WUYING Agent Management Center. The model groups serve as inference engines for Agents to execute tasks within the resource group.
//
// When an Agent runtime has its own model group configured and the resource group it belongs to also has a model group configured, the model group bound to the resource group takes effect. The resource group setting takes priority over the Agent runtime setting.
//
// When you remove the model group from the resource group to which an Agent runtime belongs, the model group configured on the Agent runtime itself automatically takes effect.
//
// Before calling this operation, make sure that you are familiar with the operations and usage of the WUYING Agent Management Center.
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
// Removes a third-party channel configuration from an agent runtime.
//
// Description:
//
// You can call this operation to remove a specific third-party channel configuration from agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. After the configuration is removed, the agent can no longer use the third-party channel for conversations.
//
// Before calling this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// Removes a model group from an Agent runtime resource.
//
// Description:
//
// You can remove model groups from Agent runtime resources such as JVS Computer, OpenClaw, and Hermes Agent in the Wuying Agent Management Center. When an Agent runtime resource needs to switch to a different model group, call this operation first to remove the authorization relationship between the Agent runtime resource and the existing model group.
//
// Make sure that you are familiar with the operations and usage of the Wuying Agent Management Center before calling this operation.
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
// Before you call this operation, make sure that you are familiar with the [Billable methods and pricing](https://help.aliyun.com/document_detail/426039.html) of WUYING Workspace.
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
// Starts a workstation.
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
// Creates and attaches tags to cloud resources. If a tag already exists on a resource, the tag value is updated.
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
// 解绑研发主机的辅助私有IP
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
// Unbinds tags from cloud resources in a unified manner.
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
// Updates the image of a delivery group.
//
// Description:
//
//	Warning: After the image update starts, sessions of end users who are accessing cloud applications will be disconnected. Proceed with caution to avoid data loss for end users.
//
// > After the update is published, changes typically take about 2 minutes to take effect on the end user side.
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
// Updates a model provider template.
//
// Description:
//
// You can update a model provider template that has been created in the Wuying Agent Management Center, including the template name, description, model service connection configuration, and Wuying security proxy switch. Partial field updates are supported. You only need to pass in the fields that you want to modify.
//
// Before you call this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
// You can update a model group that has been created in the Wuying Agent Management Center, including the group name, description, and model configuration information. The updated configuration automatically takes effect on associated cloud desktops.
//
// Before using this operation, make sure that you are familiar with the operations and usage of the Wuying Agent Management Center.
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
