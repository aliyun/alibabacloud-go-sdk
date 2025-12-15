// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sets the execution time of an over-the-air (OTA) update task.
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
// 授权用户
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
// Creates a delivery group.
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
// Create a custom image from a deployed instance. This allows you to quickly create more instances with the same configurations and avoid repeatedly configuring the instance environment each time you create the instance.
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
		Body: openapiutil.ParseToMap(body),
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
// Creates a new image by debugging the delivery group.
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
// Create one or more workstations.
//
// Description:
//
// 1.  Project is equivalent to the Resource Configuration module of the Cloud Flow console
//
// 2.  If there are multiple versions behind the input parameter ContentId:
//
//	**
//
//	**Note*	- The default version is used.
//
//	Bind simultaneously
//
// 3.  You can call the current interface only if the default version of Content is available.
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
// Deletes a delivery group that uses the By Resource - Pay-as-you-go billing method.
//
// Description:
//
// >  You cannot call this operation to delete a subscription delivery group.
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
// Deletes an application instance.
//
// Description:
//
// Only application instances that are in the Initializing or Idle state can be deleted. The operation can be called only by specific customers.
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
// # Delete a custom RDS image
//
// Description:
//
//	  You can only delete custom images to which a user belongs.
//
//		- If the product line is an image of the RDS cloud computer pool, RDS cloud application, and RDS workstation, make sure that no RDS instances use the image before you delete it.
//
//		- The RDS CloudDesktop template references an image. When you delete an image, the template is also deleted.
//
//		- If the image contains multiple regions, the images in all regions are deleted when the image is deleted.
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
// 删除工作站
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
// Queries the Elastic IP Addresses (EIPs) of workstations.
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
// Queries the details of a delivery group.
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
// Queries the credential that is used to connect to App Streaming.
//
// Description:
//
// You must call this operation at least twice to obtain a connection credential.
//
// The first time you call this operation, the system assigns an application instance to the specified convenience account and then starts the application. In this case, the ID of the started task, which is indicated by `TaskID`, is returned.
//
// In subsequent calls, you must configure `TaskID` to query whether the task is completed. If the value of `TaskStatus` in the response is `Finished`, the connection credential, which is indicated by `Ticket`, is returned.
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
// Queries information that is used to debug an application instance.
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
// Queries the details of an over-the-air (OTA) update task, including the available versions and version description.
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
// Queries resource prices.
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
// Queries the renewal prices of App Streaming resources.
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
// Queries the details of multiple delivery groups that meet the query conditions.
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
// Queries the details of application instances in a delivery group, including the IDs, status, creation time, update time, session status, and public IP addresses associated with the primary NICs of the instances.
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
// Queries the user groups authorized by a delivery group.
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
// Queries the bindings between users and resources.
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
// Queries the image information about an ECS instance.
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
// Queries the resource types that are available for purchase when you create a delivery group.
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
// Queries resource nodes.
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
// Queries the information about over-the-air (OTA) update tasks.
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
// Queries app instances of the persistent session type in a delivery group.
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
// Queries the regions that are supported by App Streaming.
//
// Description:
//
// >  All supported regions instead of available regions are returned by this operation. For more information, see [Supported regions](https://help.aliyun.com/document_detail/426036.html).
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
// Queries the tags added to one or more cloud resources.
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

	if !dara.IsNil(request.ServerInstanceType) {
		body["ServerInstanceType"] = request.ServerInstanceType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
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
// Closes all sessions in a pay-as-you-go delivery group for which a scheduled scaling policy is used.
//
// Description:
//
// >  This operation can be called only if you use a pay-as-you-go delivery group for which a scheduled scaling policy is used and if you call the operation at a time other than the scheduled time.
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
// Modifies the general policies of a delivery group, including the number of concurrent sessions and the retention period of disconnected sessions.
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
// Modify the delivery group display policy, including settings such as frame rate, resolution, and protocol type.
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
// Modifies the properties of the cloud browser.
//
// Description:
//
// Modifies the properties of the cloud browser.
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

	if !dara.IsNil(request.NetworkShrink) {
		body["Network"] = request.NetworkShrink
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
// Changes the number of nodes in a subscription delivery group.
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
// Changes the scaling policy of a delivery group. The following scaling policies are supported: fixed resource number, scheduled scaling, and auto scaling.
//
// Description:
//
// You can select one of the following scaling policies for cloud app resources:
//
//   - No scaling: Resources are not scaled.
//
//   - Auto scaling: Resources are automatically scaled based on the number of connected sessions and the duration during which no session is connected.
//
//   - Scheduled scaling: Resources are scaled during specific periods of time on specific dates.
//
// Before you call this operation, make sure that you fully understand the [billing methods and prices](https://help.aliyun.com/document_detail/426039.html) of App Streaming.
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
// Modifies the configurations of the administrator account, such as whether to enable the resource expiration reminder feature.
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
// Modify workstation properties.
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
// Queries the assigned users that are added to a delivery group by page.
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
// Renews a delivery group.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing methods and prices](https://help.aliyun.com/document_detail/426039.html) of App Streaming.
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
// Renew one workstation.
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
// Restarts the workstation.
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
// Initiates a task to replicate an image to another region.
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
// Start the workstation.
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
// Stops the workstation.
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
// 为云资源创建并绑定标签
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
// Unbinds a user and a session.
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
// Removes tags from cloud resources.
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
// *
//
// **Warning*	- After the image is updated, the end user session accessing the cloud application will be disconnected. Exercise caution to avoid end user data loss.
//
// >  After the image of the delivery group is updated, the change takes effect on the terminal in approximately 2 minutes.
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
// 更新研发主机镜像
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
