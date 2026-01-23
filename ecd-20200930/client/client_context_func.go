// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Unlocks a convenience office network that is automatically locked due to a long idle period of time.
//
// Description:
//
// If you do not create any cloud computer in a convenience office network within 15 days, the office network is automatically locked and virtual private cloud (VPC) resources are released. If you want to resume the office network, you can call this operation to unlock the office network.
//
// @param request - ActivateOfficeSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateOfficeSiteResponse
func (client *Client) ActivateOfficeSiteWithContext(ctx context.Context, request *ActivateOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *ActivateOfficeSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateOfficeSite"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateOfficeSiteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加桌面超卖用户组
//
// @param request - AddDesktopOversoldUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDesktopOversoldUserGroupResponse
func (client *Client) AddDesktopOversoldUserGroupWithContext(ctx context.Context, request *AddDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDesktopOversoldUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDesktopOversoldUserGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDesktopOversoldUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds trusted devices.
//
// Description:
//
// Each device can be registered in only one Alibaba Cloud account. If you register a device that has been registered in another Alibaba Cloud account, an error is reported.
//
// @param request - AddDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDevicesResponse
func (client *Client) AddDevicesWithContext(ctx context.Context, request *AddDevicesRequest, runtime *dara.RuntimeOptions) (_result *AddDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceIds) {
		query["DeviceIds"] = request.DeviceIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDevices"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a shared folder to the network disk.
//
// Description:
//
// You can call this operation to share a specific folder with other users. You can also configure the folder permissions.
//
// @param tmpReq - AddFilePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFilePermissionResponse
func (client *Client) AddFilePermissionWithContext(ctx context.Context, tmpReq *AddFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *AddFilePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddFilePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MemberList) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, dara.String("MemberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MemberListShrink) {
		query["MemberList"] = request.MemberListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFilePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFilePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds authorized users for a cloud computer share. The system automatically assigns cloud computers from a share to authorized users based on administrator-configured rules.
//
// @param request - AddUserToDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToDesktopGroupResponse
func (client *Client) AddUserToDesktopGroupWithContext(ctx context.Context, request *AddUserToDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToDesktopGroupResponse, _err error) {
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

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupIds) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SimpleUserGroupId) {
		query["SimpleUserGroupId"] = request.SimpleUserGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	if !dara.IsNil(request.UserOuPath) {
		query["UserOuPath"] = request.UserOuPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加用户到超卖用户组
//
// @param request - AddUserToDesktopOversoldUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToDesktopOversoldUserGroupResponse
func (client *Client) AddUserToDesktopOversoldUserGroupWithContext(ctx context.Context, request *AddUserToDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToDesktopOversoldUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddUserAmount) {
		query["AddUserAmount"] = request.AddUserAmount
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToDesktopOversoldUserGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToDesktopOversoldUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例开通公网IP
//
// @param request - AllocateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateIpAddressResponse
func (client *Client) AllocateIpAddressWithContext(ctx context.Context, request *AllocateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateIpAddress"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Apply an automatic snapshot policy to cloud computers. After the automatic snapshot policy is applied to the cloud computers, Elastic Desktop Service automatically creates snapshots for the cloud computers based on the time specified in the automatic snapshot policy.
//
// Description:
//
// You can also associate an automatic snapshot policy with a cloud desktop in the Elastic Desktop Service (EDS) console. To do so, perform the following steps: 1. Log on to the EDS console. 2. Choose Desktops and Groups > Desktops in the left-side navigation pane. 3. Find the cloud desktop that you want to manage on the Cloud Desktops page and choose More > Change Automatic Snapshot Policy in the Actions column. 4. Configure a policy for the cloud desktop as prompted in the Change Automatic Snapshot Policy panel.
//
// After you associate an automatic snapshot policy with the cloud desktop, the system creates snapshots for the cloud desktop based on the policy.
//
// @param request - ApplyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicyWithContext(ctx context.Context, request *ApplyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for the coordinate permissions.
//
// @param request - ApplyCoordinatePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyCoordinatePrivilegeResponse
func (client *Client) ApplyCoordinatePrivilegeWithContext(ctx context.Context, request *ApplyCoordinatePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ApplyCoordinatePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoId) {
		query["CoId"] = request.CoId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyCoordinatePrivilege"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyCoordinatePrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for coordination monitoring. This operation is mainly used in administrator assistance scenarios and education scenarios.
//
// @param request - ApplyCoordinationForMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyCoordinationForMonitoringResponse
func (client *Client) ApplyCoordinationForMonitoringWithContext(ctx context.Context, request *ApplyCoordinationForMonitoringRequest, runtime *dara.RuntimeOptions) (_result *ApplyCoordinationForMonitoringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoordinatePolicyType) {
		query["CoordinatePolicyType"] = request.CoordinatePolicyType
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InitiatorType) {
		query["InitiatorType"] = request.InitiatorType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceCandidates) {
		query["ResourceCandidates"] = request.ResourceCandidates
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyCoordinationForMonitoring"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyCoordinationForMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allows you to upgrade images.
//
// Description:
//
// The cloud computers for which you want to allow image updates must be in the Running state.
//
// @param request - ApproveFotaUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdateWithContext(ctx context.Context, request *ApproveFotaUpdateRequest, runtime *dara.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveFotaUpdate"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例绑定公网IP
//
// @param request - AssociateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateIpAddressResponse
func (client *Client) AssociateIpAddressWithContext(ctx context.Context, request *AssociateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AssociateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipId) {
		query["EipId"] = request.EipId
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateIpAddress"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a premium bandwidth plan to an office network. A premium bandwidth plan is used together with only one office network.
//
// @param request - AssociateNetworkPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateNetworkPackageResponse
func (client *Client) AssociateNetworkPackageWithContext(ctx context.Context, request *AssociateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateNetworkPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateNetworkPackage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateNetworkPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将创建的自定义路由表和同一VPC内的交换机绑定
//
// @param request - AssociateRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateRouteTableResponse
func (client *Client) AssociateRouteTableWithContext(ctx context.Context, request *AssociateRouteTableRequest, runtime *dara.RuntimeOptions) (_result *AssociateRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateRouteTable"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds an advanced office network to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Prerequisites
//
//   - A CEN instance is created.
//
//   - The office network is an advanced office network, and the account system type is convenient account.
//
// >  The office network is added to the CEN instance when you create the instance. An office network can be added to only one CEN instance.
//
// @param request - AttachCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachCenResponse
func (client *Client) AttachCenWithContext(ctx context.Context, request *AttachCenRequest, runtime *dara.RuntimeOptions) (_result *AttachCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachCen"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a hardware client to a user.
//
// @param request - AttachEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEndUserResponse
func (client *Client) AttachEndUserWithContext(ctx context.Context, request *AttachEndUserRequest, runtime *dara.RuntimeOptions) (_result *AttachEndUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachEndUser"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachEndUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns multiple cloud computers to users in a batch.
//
// Description:
//
//	  The cloud computers for which you want to change their policies must be in the Running state.
//
//		- After you call this operation, the assignment result is immediately returned. You can call the [DescribeDesktops](https://help.aliyun.com/document_detail/436815.html) operation to query the assignment of the cloud computer. The value of the `ManagementFlags` response parameter indicates the assignment of the cloud computer. A value of `ASSIGNING` indicates that the cloud computer is being assigned, and other values indicate that the cloud computer is assigned.
//
//		- We recommend that you check the assignment every 2 to 5 seconds and perform the checks within 50 seconds. Typically, 1 to 5 seconds are required to complete the assignment.
//
// @param request - BatchModifyEntitlementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchModifyEntitlementResponse
func (client *Client) BatchModifyEntitlementWithContext(ctx context.Context, request *BatchModifyEntitlementRequest, runtime *dara.RuntimeOptions) (_result *BatchModifyEntitlementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MaxDesktopPerUser) {
		query["MaxDesktopPerUser"] = request.MaxDesktopPerUser
	}

	if !dara.IsNil(request.MaxUserPerDesktop) {
		query["MaxUserPerDesktop"] = request.MaxUserPerDesktop
	}

	if !dara.IsNil(request.Preview) {
		query["Preview"] = request.Preview
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchModifyEntitlement"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchModifyEntitlementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a configuration group to resources.
//
// @param request - BindConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindConfigGroupResponse
func (client *Client) BindConfigGroupWithContext(ctx context.Context, request *BindConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *BindConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInfos) {
		query["ResourceInfos"] = request.ResourceInfos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an automatic snapshot policy for cloud computers.
//
// @param request - CancelAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicyWithContext(ctx context.Context, request *CancelAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a file sharing task.
//
// @param request - CancelCdsFileShareLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCdsFileShareLinkResponse
func (client *Client) CancelCdsFileShareLinkWithContext(ctx context.Context, request *CancelCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *CancelCdsFileShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.ShareId) {
		query["ShareId"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCdsFileShareLink"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCdsFileShareLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels monitoring on stream collaboration.
//
// @param request - CancelCoordinationForMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCoordinationForMonitoringResponse
func (client *Client) CancelCoordinationForMonitoringWithContext(ctx context.Context, request *CancelCoordinationForMonitoringRequest, runtime *dara.RuntimeOptions) (_result *CancelCoordinationForMonitoringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoIds) {
		query["CoIds"] = request.CoIds
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCoordinationForMonitoring"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCoordinationForMonitoringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the operation of copying an image to another region.
//
// @param request - CancelCopyImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCopyImageResponse
func (client *Client) CancelCopyImageWithContext(ctx context.Context, request *CancelCopyImageRequest, runtime *dara.RuntimeOptions) (_result *CancelCopyImageResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCopyImage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCopyImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a policy based on an existing global policy.
//
// @param request - CloneCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneCenterPolicyResponse
func (client *Client) CloneCenterPolicyWithContext(ctx context.Context, request *CloneCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *CloneCenterPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
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
		Action:      dara.String("CloneCenterPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneCenterPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones an existing policy to quickly create a policy.
//
// @param request - ClonePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClonePolicyGroupResponse
func (client *Client) ClonePolicyGroupWithContext(ctx context.Context, request *ClonePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ClonePolicyGroupResponse, _err error) {
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

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClonePolicyGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you create an object upload task, call this operation to upload the object.
//
// @param request - CompleteCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteCdsFileResponse
func (client *Client) CompleteCdsFileWithContext(ctx context.Context, request *CompleteCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CompleteCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UploadId) {
		query["UploadId"] = request.UploadId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a conditional forwarder and trust relationship for a high-definition experience (HDX)-based office network (formerly workspace). You can call the operation to configure a trust relationship for an enterprise Active Directory (AD) office network.
//
// @param request - ConfigADConnectorTrustRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigADConnectorTrustResponse
func (client *Client) ConfigADConnectorTrustWithContext(ctx context.Context, request *ConfigADConnectorTrustRequest, runtime *dara.RuntimeOptions) (_result *ConfigADConnectorTrustResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RdsLicenseDomain) {
		query["RdsLicenseDomain"] = request.RdsLicenseDomain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TrustKey) {
		query["TrustKey"] = request.TrustKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigADConnectorTrust"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigADConnectorTrustResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfigADConnectorUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigADConnectorUserResponse
func (client *Client) ConfigADConnectorUserWithContext(ctx context.Context, request *ConfigADConnectorUserRequest, runtime *dara.RuntimeOptions) (_result *ConfigADConnectorUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainPassword) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !dara.IsNil(request.DomainUserName) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !dara.IsNil(request.OUName) {
		query["OUName"] = request.OUName
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigADConnectorUser"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigADConnectorUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a file or a directory.
//
// @param request - CopyCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyCdsFileResponse
func (client *Client) CopyCdsFileWithContext(ctx context.Context, request *CopyCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CopyCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRename) {
		query["AutoRename"] = request.AutoRename
	}

	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileReceiverId) {
		query["FileReceiverId"] = request.FileReceiverId
	}

	if !dara.IsNil(request.FileReceiverType) {
		query["FileReceiverType"] = request.FileReceiverType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copy an image to another region. If you want to share an image across regions, you can call this operation to copy the image to the destination region and then share the image.
//
// @param request - CopyImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyImageResponse
func (client *Client) CopyImageWithContext(ctx context.Context, request *CopyImageRequest, runtime *dara.RuntimeOptions) (_result *CopyImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationDescription) {
		query["DestinationDescription"] = request.DestinationDescription
	}

	if !dara.IsNil(request.DestinationImageName) {
		query["DestinationImageName"] = request.DestinationImageName
	}

	if !dara.IsNil(request.DestinationRegionId) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyImage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory of the Active Directory (AD) type.
//
// Description:
//
// An AD directory is used to connect to an enterprise\\"s existing Active Directory and is suitable for large-scale cloud computer deployment. You are charged directory fees when you connect your AD to cloud computers. For more information, see [Billing overview](https://help.aliyun.com/document_detail/188395.html).
//
// @param request - CreateADConnectorDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateADConnectorDirectoryResponse
func (client *Client) CreateADConnectorDirectoryWithContext(ctx context.Context, request *CreateADConnectorDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DirectoryName) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !dara.IsNil(request.DnsAddress) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainPassword) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !dara.IsNil(request.DomainUserName) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !dara.IsNil(request.EnableAdminAccess) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !dara.IsNil(request.MfaEnabled) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.SubDomainDnsAddress) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !dara.IsNil(request.SubDomainName) {
		query["SubDomainName"] = request.SubDomainName
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateADConnectorDirectory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise Active Directory (AD) office network (formerly workspace). Elastic Desktop Service supports the following types of accounts: convenience accounts and enterprise AD accounts.
//
// Description:
//
// When you create an enterprise AD office network, the system automatically creates an AD connector to connect to an enterprise AD. You are charged for the AD connector. For more information, see [Billing overview](https://help.aliyun.com/document_detail/188395.html).
//
// After you call this operation to create an AD office network, you must perform the following steps to complete AD domain setting:
//
// 1.  Configure a conditional forwarder in a Domain Name System (DNS) server.
//
// 2.  Configure a trust relationship in an AD domain controller and call the [ConfigADConnectorTrust](https://help.aliyun.com/document_detail/311258.html) operation to configure the trust relationship with the AD office network.
//
// 3.  Call the [ListUserAdOrganizationUnits](https://help.aliyun.com/document_detail/311259.html) operation to query a list of organizational units (OUs) of the AD domain, and call the [ConfigADConnectorUser](https://help.aliyun.com/document_detail/311262.html) operation to specify an OU and administrator for the AD office network.
//
//	>  When you create the AD office network, take note of the DomainUserName and DomainPassword parameters. If you specify the parameters, you need to only configure a conditional forwarder. If you do not specify the parameters, you must configure a conditional forwarder, trust relationship, and OU as prompted.
//
// For more information, see [Create and manage enterprise AD office networks](https://help.aliyun.com/document_detail/214469.html).
//
// @param request - CreateADConnectorOfficeSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateADConnectorOfficeSiteResponse
func (client *Client) CreateADConnectorOfficeSiteWithContext(ctx context.Context, request *CreateADConnectorOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdHostname) {
		query["AdHostname"] = request.AdHostname
	}

	if !dara.IsNil(request.BackupDCHostname) {
		query["BackupDCHostname"] = request.BackupDCHostname
	}

	if !dara.IsNil(request.BackupDns) {
		query["BackupDns"] = request.BackupDns
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DnsAddress) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainPassword) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !dara.IsNil(request.DomainUserName) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !dara.IsNil(request.EnableAdminAccess) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !dara.IsNil(request.EnableInternetAccess) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !dara.IsNil(request.MfaEnabled) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.SubDomainDnsAddress) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !dara.IsNil(request.SubDomainName) {
		query["SubDomainName"] = request.SubDomainName
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateADConnectorOfficeSite"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a File Storage NAS (NAS) file system and mount the file system to the workspace in which a desktop group resides.
//
// @param request - CreateAndBindNasFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndBindNasFileSystemResponse
func (client *Client) CreateAndBindNasFileSystemWithContext(ctx context.Context, request *CreateAndBindNasFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateAndBindNasFileSystemResponse, _err error) {
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

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.FileSystemName) {
		query["FileSystemName"] = request.FileSystemName
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndBindNasFileSystem"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndBindNasFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic snapshot policy. WUYING WorkSpace automatically creates snapshots based on the time specified by the cron expression in the automatic snapshot policy.
//
// Description:
//
// You can call the operation to create an automatic snapshot policy based on a CRON expression. Then, the system automatically creates snapshots of a cloud desktop based on the policy.
//
// @param request - CreateAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates data transfer plans.
//
// @param request - CreateBandwidthResourcePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBandwidthResourcePackagesResponse
func (client *Client) CreateBandwidthResourcePackagesWithContext(ctx context.Context, request *CreateBandwidthResourcePackagesRequest, runtime *dara.RuntimeOptions) (_result *CreateBandwidthResourcePackagesResponse, _err error) {
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

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.PackageSize) {
		query["PackageSize"] = request.PackageSize
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBandwidthResourcePackages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBandwidthResourcePackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom cloud computer template.
//
// Description:
//
// Cloud computer templates include system templates and custom templates. A system template is the default template provided by Alibaba Cloud. You can call this operation to create a custom template.
//
// @param request - CreateBundleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBundleResponse
func (client *Client) CreateBundleWithContext(ctx context.Context, request *CreateBundleRequest, runtime *dara.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleName) {
		query["BundleName"] = request.BundleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RootDiskPerformanceLevel) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !dara.IsNil(request.RootDiskSizeGib) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !dara.IsNil(request.UserDiskPerformanceLevel) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !dara.IsNil(request.UserDiskSizeGib) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBundle"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBundleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a file to a cloud disk.
//
// Description:
//
// After the RAM permissions are authenticated, you can call the CreateCdsFile operation to obtain the upload URL of a file and upload the file to a cloud disk.
//
// @param request - CreateCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCdsFileResponse
func (client *Client) CreateCdsFileWithContext(ctx context.Context, request *CreateCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CreateCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.ConflictPolicy) {
		query["ConflictPolicy"] = request.ConflictPolicy
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileHash) {
		query["FileHash"] = request.FileHash
	}

	if !dara.IsNil(request.FileLength) {
		query["FileLength"] = request.FileLength
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ParentFileId) {
		query["ParentFileId"] = request.ParentFileId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file sharing task.
//
// @param request - CreateCdsFileShareLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCdsFileShareLinkResponse
func (client *Client) CreateCdsFileShareLinkWithContext(ctx context.Context, request *CreateCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *CreateCdsFileShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableDownload) {
		query["DisableDownload"] = request.DisableDownload
	}

	if !dara.IsNil(request.DisablePreview) {
		query["DisablePreview"] = request.DisablePreview
	}

	if !dara.IsNil(request.DisableSave) {
		query["DisableSave"] = request.DisableSave
	}

	if !dara.IsNil(request.DownloadLimit) {
		query["DownloadLimit"] = request.DownloadLimit
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.FileIds) {
		query["FileIds"] = request.FileIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PreviewLimit) {
		query["PreviewLimit"] = request.PreviewLimit
	}

	if !dara.IsNil(request.SaveLimit) {
		query["SaveLimit"] = request.SaveLimit
	}

	if !dara.IsNil(request.ShareName) {
		query["ShareName"] = request.ShareName
	}

	if !dara.IsNil(request.SharePwd) {
		query["SharePwd"] = request.SharePwd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCdsFileShareLink"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCdsFileShareLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a center policy.
//
// @param request - CreateCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenterPolicyResponse
func (client *Client) CreateCenterPolicyWithContext(ctx context.Context, request *CreateCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCenterPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcademicProxy) {
		query["AcademicProxy"] = request.AcademicProxy
	}

	if !dara.IsNil(request.AdminAccess) {
		query["AdminAccess"] = request.AdminAccess
	}

	if !dara.IsNil(request.AppContentProtection) {
		query["AppContentProtection"] = request.AppContentProtection
	}

	if !dara.IsNil(request.AuthorizeAccessPolicyRule) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !dara.IsNil(request.AuthorizeSecurityPolicyRule) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !dara.IsNil(request.AutoReconnect) {
		query["AutoReconnect"] = request.AutoReconnect
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.CameraRedirect) {
		query["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.ClientControlMenu) {
		query["ClientControlMenu"] = request.ClientControlMenu
	}

	if !dara.IsNil(request.ClientCreateSnapshot) {
		query["ClientCreateSnapshot"] = request.ClientCreateSnapshot
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Clipboard) {
		query["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.ClipboardGraineds) {
		query["ClipboardGraineds"] = request.ClipboardGraineds
	}

	if !dara.IsNil(request.ClipboardScope) {
		query["ClipboardScope"] = request.ClipboardScope
	}

	if !dara.IsNil(request.ColorEnhancement) {
		query["ColorEnhancement"] = request.ColorEnhancement
	}

	if !dara.IsNil(request.CpdDriveClipboard) {
		query["CpdDriveClipboard"] = request.CpdDriveClipboard
	}

	if !dara.IsNil(request.CpuDownGradeDuration) {
		query["CpuDownGradeDuration"] = request.CpuDownGradeDuration
	}

	if !dara.IsNil(request.CpuOverload) {
		query["CpuOverload"] = request.CpuOverload
	}

	if !dara.IsNil(request.CpuProcessors) {
		query["CpuProcessors"] = request.CpuProcessors
	}

	if !dara.IsNil(request.CpuProtectedMode) {
		query["CpuProtectedMode"] = request.CpuProtectedMode
	}

	if !dara.IsNil(request.CpuRateLimit) {
		query["CpuRateLimit"] = request.CpuRateLimit
	}

	if !dara.IsNil(request.CpuSampleDuration) {
		query["CpuSampleDuration"] = request.CpuSampleDuration
	}

	if !dara.IsNil(request.CpuSingleRateLimit) {
		query["CpuSingleRateLimit"] = request.CpuSingleRateLimit
	}

	if !dara.IsNil(request.DeviceConnectHint) {
		query["DeviceConnectHint"] = request.DeviceConnectHint
	}

	if !dara.IsNil(request.DeviceRedirects) {
		query["DeviceRedirects"] = request.DeviceRedirects
	}

	if !dara.IsNil(request.DeviceRules) {
		query["DeviceRules"] = request.DeviceRules
	}

	if !dara.IsNil(request.DisconnectKeepSession) {
		query["DisconnectKeepSession"] = request.DisconnectKeepSession
	}

	if !dara.IsNil(request.DisconnectKeepSessionTime) {
		query["DisconnectKeepSessionTime"] = request.DisconnectKeepSessionTime
	}

	if !dara.IsNil(request.DiskOverload) {
		query["DiskOverload"] = request.DiskOverload
	}

	if !dara.IsNil(request.DisplayMode) {
		query["DisplayMode"] = request.DisplayMode
	}

	if !dara.IsNil(request.DomainResolveRule) {
		query["DomainResolveRule"] = request.DomainResolveRule
	}

	if !dara.IsNil(request.DomainResolveRuleType) {
		query["DomainResolveRuleType"] = request.DomainResolveRuleType
	}

	if !dara.IsNil(request.EnableSessionRateLimiting) {
		query["EnableSessionRateLimiting"] = request.EnableSessionRateLimiting
	}

	if !dara.IsNil(request.EndUserApplyAdminCoordinate) {
		query["EndUserApplyAdminCoordinate"] = request.EndUserApplyAdminCoordinate
	}

	if !dara.IsNil(request.EndUserGroupCoordinate) {
		query["EndUserGroupCoordinate"] = request.EndUserGroupCoordinate
	}

	if !dara.IsNil(request.ExternalDrive) {
		query["ExternalDrive"] = request.ExternalDrive
	}

	if !dara.IsNil(request.FileMigrate) {
		query["FileMigrate"] = request.FileMigrate
	}

	if !dara.IsNil(request.FileTransferAddress) {
		query["FileTransferAddress"] = request.FileTransferAddress
	}

	if !dara.IsNil(request.FileTransferSpeed) {
		query["FileTransferSpeed"] = request.FileTransferSpeed
	}

	if !dara.IsNil(request.FileTransferSpeedLocation) {
		query["FileTransferSpeedLocation"] = request.FileTransferSpeedLocation
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.HoverConfigMsg) {
		query["HoverConfigMsg"] = request.HoverConfigMsg
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.InternetCommunicationProtocol) {
		query["InternetCommunicationProtocol"] = request.InternetCommunicationProtocol
	}

	if !dara.IsNil(request.InternetPrinter) {
		query["InternetPrinter"] = request.InternetPrinter
	}

	if !dara.IsNil(request.LocalDrive) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.MaxReconnectTime) {
		query["MaxReconnectTime"] = request.MaxReconnectTime
	}

	if !dara.IsNil(request.MemoryDownGradeDuration) {
		query["MemoryDownGradeDuration"] = request.MemoryDownGradeDuration
	}

	if !dara.IsNil(request.MemoryOverload) {
		query["MemoryOverload"] = request.MemoryOverload
	}

	if !dara.IsNil(request.MemoryProcessors) {
		query["MemoryProcessors"] = request.MemoryProcessors
	}

	if !dara.IsNil(request.MemoryProtectedMode) {
		query["MemoryProtectedMode"] = request.MemoryProtectedMode
	}

	if !dara.IsNil(request.MemoryRateLimit) {
		query["MemoryRateLimit"] = request.MemoryRateLimit
	}

	if !dara.IsNil(request.MemorySampleDuration) {
		query["MemorySampleDuration"] = request.MemorySampleDuration
	}

	if !dara.IsNil(request.MemorySingleRateLimit) {
		query["MemorySingleRateLimit"] = request.MemorySingleRateLimit
	}

	if !dara.IsNil(request.MobileRestart) {
		query["MobileRestart"] = request.MobileRestart
	}

	if !dara.IsNil(request.MobileSafeMenu) {
		query["MobileSafeMenu"] = request.MobileSafeMenu
	}

	if !dara.IsNil(request.MobileShutdown) {
		query["MobileShutdown"] = request.MobileShutdown
	}

	if !dara.IsNil(request.MobileWuyingKeeper) {
		query["MobileWuyingKeeper"] = request.MobileWuyingKeeper
	}

	if !dara.IsNil(request.MobileWyAssistant) {
		query["MobileWyAssistant"] = request.MobileWyAssistant
	}

	if !dara.IsNil(request.ModelLibrary) {
		query["ModelLibrary"] = request.ModelLibrary
	}

	if !dara.IsNil(request.MultiScreen) {
		query["MultiScreen"] = request.MultiScreen
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetRedirect) {
		query["NetRedirect"] = request.NetRedirect
	}

	if !dara.IsNil(request.NetRedirectRule) {
		query["NetRedirectRule"] = request.NetRedirectRule
	}

	if !dara.IsNil(request.NoOperationDisconnect) {
		query["NoOperationDisconnect"] = request.NoOperationDisconnect
	}

	if !dara.IsNil(request.NoOperationDisconnectTime) {
		query["NoOperationDisconnectTime"] = request.NoOperationDisconnectTime
	}

	if !dara.IsNil(request.PortProxy) {
		query["PortProxy"] = request.PortProxy
	}

	if !dara.IsNil(request.PrinterRedirect) {
		query["PrinterRedirect"] = request.PrinterRedirect
	}

	if !dara.IsNil(request.QualityEnhancement) {
		query["QualityEnhancement"] = request.QualityEnhancement
	}

	if !dara.IsNil(request.RecordEventDuration) {
		query["RecordEventDuration"] = request.RecordEventDuration
	}

	if !dara.IsNil(request.RecordEventFileExts) {
		query["RecordEventFileExts"] = request.RecordEventFileExts
	}

	if !dara.IsNil(request.RecordEventFilePaths) {
		query["RecordEventFilePaths"] = request.RecordEventFilePaths
	}

	if !dara.IsNil(request.RecordEventLevels) {
		query["RecordEventLevels"] = request.RecordEventLevels
	}

	if !dara.IsNil(request.RecordEventRegisters) {
		query["RecordEventRegisters"] = request.RecordEventRegisters
	}

	if !dara.IsNil(request.RecordEvents) {
		query["RecordEvents"] = request.RecordEvents
	}

	if !dara.IsNil(request.Recording) {
		query["Recording"] = request.Recording
	}

	if !dara.IsNil(request.RecordingAudio) {
		query["RecordingAudio"] = request.RecordingAudio
	}

	if !dara.IsNil(request.RecordingDuration) {
		query["RecordingDuration"] = request.RecordingDuration
	}

	if !dara.IsNil(request.RecordingEndTime) {
		query["RecordingEndTime"] = request.RecordingEndTime
	}

	if !dara.IsNil(request.RecordingExpires) {
		query["RecordingExpires"] = request.RecordingExpires
	}

	if !dara.IsNil(request.RecordingFps) {
		query["RecordingFps"] = request.RecordingFps
	}

	if !dara.IsNil(request.RecordingStartTime) {
		query["RecordingStartTime"] = request.RecordingStartTime
	}

	if !dara.IsNil(request.RecordingUserNotify) {
		query["RecordingUserNotify"] = request.RecordingUserNotify
	}

	if !dara.IsNil(request.RecordingUserNotifyMessage) {
		query["RecordingUserNotifyMessage"] = request.RecordingUserNotifyMessage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteCoordinate) {
		query["RemoteCoordinate"] = request.RemoteCoordinate
	}

	if !dara.IsNil(request.ResetDesktop) {
		query["ResetDesktop"] = request.ResetDesktop
	}

	if !dara.IsNil(request.ResolutionDpi) {
		query["ResolutionDpi"] = request.ResolutionDpi
	}

	if !dara.IsNil(request.ResolutionHeight) {
		query["ResolutionHeight"] = request.ResolutionHeight
	}

	if !dara.IsNil(request.ResolutionModel) {
		query["ResolutionModel"] = request.ResolutionModel
	}

	if !dara.IsNil(request.ResolutionWidth) {
		query["ResolutionWidth"] = request.ResolutionWidth
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SafeMenu) {
		query["SafeMenu"] = request.SafeMenu
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	if !dara.IsNil(request.ScreenDisplayMode) {
		query["ScreenDisplayMode"] = request.ScreenDisplayMode
	}

	if !dara.IsNil(request.SessionMaxRateKbps) {
		query["SessionMaxRateKbps"] = request.SessionMaxRateKbps
	}

	if !dara.IsNil(request.SmoothEnhancement) {
		query["SmoothEnhancement"] = request.SmoothEnhancement
	}

	if !dara.IsNil(request.StatusMonitor) {
		query["StatusMonitor"] = request.StatusMonitor
	}

	if !dara.IsNil(request.StreamingMode) {
		query["StreamingMode"] = request.StreamingMode
	}

	if !dara.IsNil(request.TargetFps) {
		query["TargetFps"] = request.TargetFps
	}

	if !dara.IsNil(request.Taskbar) {
		query["Taskbar"] = request.Taskbar
	}

	if !dara.IsNil(request.UsbRedirect) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !dara.IsNil(request.UsbSupplyRedirectRule) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !dara.IsNil(request.UseTime) {
		query["UseTime"] = request.UseTime
	}

	if !dara.IsNil(request.VideoEncAvgKbps) {
		query["VideoEncAvgKbps"] = request.VideoEncAvgKbps
	}

	if !dara.IsNil(request.VideoEncMaxQP) {
		query["VideoEncMaxQP"] = request.VideoEncMaxQP
	}

	if !dara.IsNil(request.VideoEncMinQP) {
		query["VideoEncMinQP"] = request.VideoEncMinQP
	}

	if !dara.IsNil(request.VideoEncPeakKbps) {
		query["VideoEncPeakKbps"] = request.VideoEncPeakKbps
	}

	if !dara.IsNil(request.VideoEncPolicy) {
		query["VideoEncPolicy"] = request.VideoEncPolicy
	}

	if !dara.IsNil(request.VideoRedirect) {
		query["VideoRedirect"] = request.VideoRedirect
	}

	if !dara.IsNil(request.VisualQuality) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	if !dara.IsNil(request.WatermarkAntiCam) {
		query["WatermarkAntiCam"] = request.WatermarkAntiCam
	}

	if !dara.IsNil(request.WatermarkColor) {
		query["WatermarkColor"] = request.WatermarkColor
	}

	if !dara.IsNil(request.WatermarkColumnAmount) {
		query["WatermarkColumnAmount"] = request.WatermarkColumnAmount
	}

	if !dara.IsNil(request.WatermarkCustomText) {
		query["WatermarkCustomText"] = request.WatermarkCustomText
	}

	if !dara.IsNil(request.WatermarkDegree) {
		query["WatermarkDegree"] = request.WatermarkDegree
	}

	if !dara.IsNil(request.WatermarkFontSize) {
		query["WatermarkFontSize"] = request.WatermarkFontSize
	}

	if !dara.IsNil(request.WatermarkFontStyle) {
		query["WatermarkFontStyle"] = request.WatermarkFontStyle
	}

	if !dara.IsNil(request.WatermarkPower) {
		query["WatermarkPower"] = request.WatermarkPower
	}

	if !dara.IsNil(request.WatermarkRowAmount) {
		query["WatermarkRowAmount"] = request.WatermarkRowAmount
	}

	if !dara.IsNil(request.WatermarkSecurity) {
		query["WatermarkSecurity"] = request.WatermarkSecurity
	}

	if !dara.IsNil(request.WatermarkShadow) {
		query["WatermarkShadow"] = request.WatermarkShadow
	}

	if !dara.IsNil(request.WatermarkTransparencyValue) {
		query["WatermarkTransparencyValue"] = request.WatermarkTransparencyValue
	}

	if !dara.IsNil(request.WatermarkType) {
		query["WatermarkType"] = request.WatermarkType
	}

	if !dara.IsNil(request.WuyingKeeper) {
		query["WuyingKeeper"] = request.WuyingKeeper
	}

	if !dara.IsNil(request.WyAssistant) {
		query["WyAssistant"] = request.WyAssistant
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenterPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenterPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to use a team space.
//
// Description:
//
// The list of teams of a cloud disk in Cloud Drive Service is synchronized from the Organization tab in the Elastic Desktop Service (EDS) console. You can choose Users > Manager User > User > Organization in the console. If you want to authorize a user to use a team space, you must move the user to the corresponding organization. After you move the user, the user can view the menu bar of the team space on a Cloud Drive Service client.
//
// @param request - CreateCloudDriveGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudDriveGroupResponse
func (client *Client) CreateCloudDriveGroupWithContext(ctx context.Context, request *CreateCloudDriveGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminUserIds) {
		query["AdminUserIds"] = request.AdminUserIds
	}

	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TotalSize) {
		query["TotalSize"] = request.TotalSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudDriveGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudDriveGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise drive.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Enterprise Drive Service (formerly Cloud Drive Service). For more information, see [Overview](https://help.aliyun.com/document_detail/386301.html).
//
// @param request - CreateCloudDriveServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudDriveServiceResponse
func (client *Client) CreateCloudDriveServiceWithContext(ctx context.Context, request *CreateCloudDriveServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CdsChargeType) {
		query["CdsChargeType"] = request.CdsChargeType
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MaxSize) {
		query["MaxSize"] = request.MaxSize
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteType) {
		query["OfficeSiteType"] = request.OfficeSiteType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	if !dara.IsNil(request.UserCount) {
		query["UserCount"] = request.UserCount
	}

	if !dara.IsNil(request.UserMaxSize) {
		query["UserMaxSize"] = request.UserMaxSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudDriveService"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudDriveServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates the users of a cloud disk.
//
// @param request - CreateCloudDriveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudDriveUsersResponse
func (client *Client) CreateCloudDriveUsersWithContext(ctx context.Context, request *CreateCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserMaxSize) {
		query["UserMaxSize"] = request.UserMaxSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudDriveUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudDriveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration group. A configuration group stores the setup details for scheduled tasks on cloud computers.
//
// @param request - CreateConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigGroupResponse
func (client *Client) CreateConfigGroupWithContext(ctx context.Context, request *CreateConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigTimers) {
		query["ConfigTimers"] = request.ConfigTimers
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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
		Action:      dara.String("CreateConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a shared group.
//
// Description:
//
//	  To learn about the features, application scenarios, usage limits, scaling policies, and other details of shared groups, refer to [Overview](https://help.aliyun.com/document_detail/290959.html).
//
//		- Before you call this operation, make sure that the required resources, such as the office network, cloud computer template, and policies, are created.
//
// @param request - CreateDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDesktopGroupResponse
func (client *Client) CreateDesktopGroupWithContext(ctx context.Context, request *CreateDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllClassifyUsers) {
		query["AllClassifyUsers"] = request.AllClassifyUsers
	}

	if !dara.IsNil(request.AllowAutoSetup) {
		query["AllowAutoSetup"] = request.AllowAutoSetup
	}

	if !dara.IsNil(request.AllowBufferCount) {
		query["AllowBufferCount"] = request.AllowBufferCount
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BindAmount) {
		query["BindAmount"] = request.BindAmount
	}

	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.BuyDesktopsCount) {
		query["BuyDesktopsCount"] = request.BuyDesktopsCount
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Classify) {
		query["Classify"] = request.Classify
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.ConnectDuration) {
		query["ConnectDuration"] = request.ConnectDuration
	}

	if !dara.IsNil(request.DataDiskCategory) {
		query["DataDiskCategory"] = request.DataDiskCategory
	}

	if !dara.IsNil(request.DataDiskPerLevel) {
		query["DataDiskPerLevel"] = request.DataDiskPerLevel
	}

	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.DefaultInitDesktopCount) {
		query["DefaultInitDesktopCount"] = request.DefaultInitDesktopCount
	}

	if !dara.IsNil(request.DefaultLanguage) {
		query["DefaultLanguage"] = request.DefaultLanguage
	}

	if !dara.IsNil(request.DeleteDuration) {
		query["DeleteDuration"] = request.DeleteDuration
	}

	if !dara.IsNil(request.DesktopGroupName) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.ExclusiveType) {
		query["ExclusiveType"] = request.ExclusiveType
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.GroupAmount) {
		query["GroupAmount"] = request.GroupAmount
	}

	if !dara.IsNil(request.GroupVersion) {
		query["GroupVersion"] = request.GroupVersion
	}

	if !dara.IsNil(request.Hostname) {
		query["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.IdleDisconnectDuration) {
		query["IdleDisconnectDuration"] = request.IdleDisconnectDuration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.KeepDuration) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !dara.IsNil(request.LoadPolicy) {
		query["LoadPolicy"] = request.LoadPolicy
	}

	if !dara.IsNil(request.MaxDesktopsCount) {
		query["MaxDesktopsCount"] = request.MaxDesktopsCount
	}

	if !dara.IsNil(request.MinDesktopsCount) {
		query["MinDesktopsCount"] = request.MinDesktopsCount
	}

	if !dara.IsNil(request.MultiResource) {
		query["MultiResource"] = request.MultiResource
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OwnType) {
		query["OwnType"] = request.OwnType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.ProfileFollowSwitch) {
		query["ProfileFollowSwitch"] = request.ProfileFollowSwitch
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RatioThreshold) {
		query["RatioThreshold"] = request.RatioThreshold
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResetType) {
		query["ResetType"] = request.ResetType
	}

	if !dara.IsNil(request.ScaleStrategyId) {
		query["ScaleStrategyId"] = request.ScaleStrategyId
	}

	if !dara.IsNil(request.SessionType) {
		query["SessionType"] = request.SessionType
	}

	if !dara.IsNil(request.SimpleUserGroupId) {
		query["SimpleUserGroupId"] = request.SimpleUserGroupId
	}

	if !dara.IsNil(request.SnapshotPolicyId) {
		query["SnapshotPolicyId"] = request.SnapshotPolicyId
	}

	if !dara.IsNil(request.StopDuration) {
		query["StopDuration"] = request.StopDuration
	}

	if !dara.IsNil(request.SystemDiskCategory) {
		query["SystemDiskCategory"] = request.SystemDiskCategory
	}

	if !dara.IsNil(request.SystemDiskPerLevel) {
		query["SystemDiskPerLevel"] = request.SystemDiskPerLevel
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TimerGroupId) {
		query["TimerGroupId"] = request.TimerGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	if !dara.IsNil(request.UserOuPath) {
		query["UserOuPath"] = request.UserOuPath
	}

	if !dara.IsNil(request.VolumeEncryptionEnabled) {
		query["VolumeEncryptionEnabled"] = request.VolumeEncryptionEnabled
	}

	if !dara.IsNil(request.VolumeEncryptionKey) {
		query["VolumeEncryptionKey"] = request.VolumeEncryptionKey
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建桌面超卖组
//
// @param request - CreateDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDesktopOversoldGroupResponse
func (client *Client) CreateDesktopOversoldGroupWithContext(ctx context.Context, request *CreateDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopOversoldGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrenceCount) {
		query["ConcurrenceCount"] = request.ConcurrenceCount
	}

	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.IdleDisconnectDuration) {
		query["IdleDisconnectDuration"] = request.IdleDisconnectDuration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.KeepDuration) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OversoldUserCount) {
		query["OversoldUserCount"] = request.OversoldUserCount
	}

	if !dara.IsNil(request.OversoldWarn) {
		query["OversoldWarn"] = request.OversoldWarn
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.StopDuration) {
		query["StopDuration"] = request.StopDuration
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates cloud computers. If you specify end users when you create cloud computers, the cloud computers are assigned to the end users after the cloud computers are created.
//
// Description:
//
// Before you create cloud computers, complete the following preparations:
//
//   - An office network (formerly called workspace) and users are created. For more information, see:
//
//   - Convenience office network: [CreateSimpleOfficeSite](https://help.aliyun.com/document_detail/215416.html) and [CreateUsers](https://help.aliyun.com/document_detail/437832.html).
//
//   - Active Directory (AD) office network: [CreateADConnectorOfficeSite](https://help.aliyun.com/document_detail/215417.html) and [Create an AD user](https://help.aliyun.com/document_detail/188619.html).
//
//   - Make sure a cloud computer template exists. If no cloud computer template exists, call the [CreateBundle](https://help.aliyun.com/document_detail/188883.html) operation to create a template.
//
//   - Make sure a policy exists. If no policy exists, call the [CreatePolicyGroup](https://help.aliyun.com/document_detail/188889.html) operation to create a policy.
//
// If you want the cloud computers to automatically execute a custom command script, you can use the `UserCommands` field to configure a custom command.
//
// @param tmpReq - CreateDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDesktopsResponse
func (client *Client) CreateDesktopsWithContext(ctx context.Context, tmpReq *CreateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDesktopsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DesktopAttachment) {
		request.DesktopAttachmentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DesktopAttachment, dara.String("DesktopAttachment"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AppRuleId) {
		query["AppRuleId"] = request.AppRuleId
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.BundleModels) {
		query["BundleModels"] = request.BundleModels
	}

	if !dara.IsNil(request.ChannelCookie) {
		query["ChannelCookie"] = request.ChannelCookie
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DesktopAttachmentShrink) {
		query["DesktopAttachment"] = request.DesktopAttachmentShrink
	}

	if !dara.IsNil(request.DesktopMemberIp) {
		query["DesktopMemberIp"] = request.DesktopMemberIp
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DesktopNameSuffix) {
		query["DesktopNameSuffix"] = request.DesktopNameSuffix
	}

	if !dara.IsNil(request.DesktopTimers) {
		query["DesktopTimers"] = request.DesktopTimers
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Hostname) {
		query["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.MonthDesktopSetting) {
		query["MonthDesktopSetting"] = request.MonthDesktopSetting
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.QosRuleId) {
		query["QosRuleId"] = request.QosRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SavingPlanId) {
		query["SavingPlanId"] = request.SavingPlanId
	}

	if !dara.IsNil(request.SnapshotPolicyId) {
		query["SnapshotPolicyId"] = request.SnapshotPolicyId
	}

	if !dara.IsNil(request.SubnetId) {
		query["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TimerGroupId) {
		query["TimerGroupId"] = request.TimerGroupId
	}

	if !dara.IsNil(request.UserAssignMode) {
		query["UserAssignMode"] = request.UserAssignMode
	}

	if !dara.IsNil(request.UserCommands) {
		query["UserCommands"] = request.UserCommands
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VolumeEncryptionEnabled) {
		query["VolumeEncryptionEnabled"] = request.VolumeEncryptionEnabled
	}

	if !dara.IsNil(request.VolumeEncryptionKey) {
		query["VolumeEncryptionKey"] = request.VolumeEncryptionKey
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the disk encryption feature and adds the service-linked role that is encrypted by Cloud Drive Service to a Resource Access Management (RAM) user.
//
// @param request - CreateDiskEncryptionServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskEncryptionServiceResponse
func (client *Client) CreateDiskEncryptionServiceWithContext(ctx context.Context, request *CreateDiskEncryptionServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskEncryptionServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiskEncryptionService"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiskEncryptionServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user-level storage resource.
//
// @param request - CreateDriveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDriveResponse
func (client *Client) CreateDriveWithContext(ctx context.Context, request *CreateDriveRequest, runtime *dara.RuntimeOptions) (_result *CreateDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DriveName) {
		query["DriveName"] = request.DriveName
	}

	if !dara.IsNil(request.ExternalDomainId) {
		query["ExternalDomainId"] = request.ExternalDomainId
	}

	if !dara.IsNil(request.ProfileRoaming) {
		query["ProfileRoaming"] = request.ProfileRoaming
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDrive"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDriveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data report export task.
//
// @param request - CreateEcdReportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEcdReportTaskResponse
func (client *Client) CreateEcdReportTaskWithContext(ctx context.Context, request *CreateEcdReportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateEcdReportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterList) {
		query["FilterList"] = request.FilterList
	}

	if !dara.IsNil(request.LangType) {
		query["LangType"] = request.LangType
	}

	if !dara.IsNil(request.ReportFileName) {
		query["ReportFileName"] = request.ReportFileName
	}

	if !dara.IsNil(request.SubType) {
		query["SubType"] = request.SubType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEcdReportTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEcdReportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加DNAT条目
//
// @param request - CreateForwardEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForwardEntryResponse
func (client *Client) CreateForwardEntryWithContext(ctx context.Context, request *CreateForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIp) {
		query["ExternalIp"] = request.ExternalIp
	}

	if !dara.IsNil(request.ExternalPort) {
		query["ExternalPort"] = request.ExternalPort
	}

	if !dara.IsNil(request.ForwardEntryName) {
		query["ForwardEntryName"] = request.ForwardEntryName
	}

	if !dara.IsNil(request.ForwardTableId) {
		query["ForwardTableId"] = request.ForwardTableId
	}

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.InternalPort) {
		query["InternalPort"] = request.InternalPort
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateForwardEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateForwardEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image based on a deployed cloud computer. Then, you can use the custom image to create cloud computers that have the same configurations. This prevents the repeated settings when you create cloud computers.
//
// @param request - CreateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageResponse
func (client *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest, runtime *dara.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCleanUserdata) {
		query["AutoCleanUserdata"] = request.AutoCleanUserdata
	}

	if !dara.IsNil(request.DataSnapshotIds) {
		query["DataSnapshotIds"] = request.DataSnapshotIds
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ImageResourceType) {
		query["ImageResourceType"] = request.ImageResourceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SnapshotIds) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a NAS file system.
//
// Description:
//
// <props="china">
//
// - Each standard workspace can create one NAS file system to meet the need for sharing files between cloud desktops in the workspace.
//
// - The system will automatically create a general-purpose NAS file system (with storage specifications of Capacity and Performance, with capacities of 10 PiB and 1 PiB respectively) and generate a default mount point.
//
// - The NAS file system uses pay-as-you-go by default. You need to pay for the actual storage usage. You can also purchase resource packages to offset the storage usage.
//
// For more information, see [Creating Shared Storage NAS](https://help.aliyun.com/document_detail/214481.html).
//
// <props="intl">
//
// - Each standard workspace can create one NAS file system to meet the need for sharing files between cloud desktops in the workspace.
//
// - The system will automatically create a general-purpose NAS file system (with storage specifications of Capacity and Performance, with capacities of 10 PiB and 1 PiB respectively) and generate a default mount point.
//
// - The NAS file system uses pay-as-you-go by default. You need to pay for the actual storage usage. You can also purchase storage packages to offset the storage usage.
//
// For more information, see [Creating Shared Storage NAS](https://help.aliyun.com/document_detail/214481.html).
//
// @param request - CreateNASFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNASFileSystemResponse
func (client *Client) CreateNASFileSystemWithContext(ctx context.Context, request *CreateNASFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateNASFileSystemResponse, _err error) {
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

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNASFileSystem"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建NAT网关
//
// @param request - CreateNatGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatGatewayResponse
func (client *Client) CreateNatGatewayWithContext(ctx context.Context, request *CreateNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateNatGatewayResponse, _err error) {
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

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNatGateway"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNatGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a premium bandwidth plan for an office network.
//
// @param request - CreateNetworkPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkPackageResponse
func (client *Client) CreateNetworkPackageWithContext(ctx context.Context, request *CreateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ChannelCookie) {
		query["ChannelCookie"] = request.ChannelCookie
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkPackage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud computer policy.
//
// Description:
//
// A cloud computer policy is a collection of rules to manage cloud computers in performance and security. For example, you can create a basic policy that involves the disk mapping, USB redirection, watermarking features and rules such as DNS rules. For more information, see [Policy overview](https://help.aliyun.com/document_detail/189345.html).
//
// @param request - CreatePolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroupWithContext(ctx context.Context, request *CreatePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminAccess) {
		query["AdminAccess"] = request.AdminAccess
	}

	if !dara.IsNil(request.AppContentProtection) {
		query["AppContentProtection"] = request.AppContentProtection
	}

	if !dara.IsNil(request.AuthorizeAccessPolicyRule) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !dara.IsNil(request.AuthorizeSecurityPolicyRule) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !dara.IsNil(request.CameraRedirect) {
		query["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Clipboard) {
		query["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.DeviceRedirects) {
		query["DeviceRedirects"] = request.DeviceRedirects
	}

	if !dara.IsNil(request.DeviceRules) {
		query["DeviceRules"] = request.DeviceRules
	}

	if !dara.IsNil(request.DomainList) {
		query["DomainList"] = request.DomainList
	}

	if !dara.IsNil(request.DomainResolveRule) {
		query["DomainResolveRule"] = request.DomainResolveRule
	}

	if !dara.IsNil(request.DomainResolveRuleType) {
		query["DomainResolveRuleType"] = request.DomainResolveRuleType
	}

	if !dara.IsNil(request.EndUserApplyAdminCoordinate) {
		query["EndUserApplyAdminCoordinate"] = request.EndUserApplyAdminCoordinate
	}

	if !dara.IsNil(request.EndUserGroupCoordinate) {
		query["EndUserGroupCoordinate"] = request.EndUserGroupCoordinate
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.Html5Access) {
		query["Html5Access"] = request.Html5Access
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.InternetCommunicationProtocol) {
		query["InternetCommunicationProtocol"] = request.InternetCommunicationProtocol
	}

	if !dara.IsNil(request.LocalDrive) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.MaxReconnectTime) {
		query["MaxReconnectTime"] = request.MaxReconnectTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetRedirect) {
		query["NetRedirect"] = request.NetRedirect
	}

	if !dara.IsNil(request.PreemptLogin) {
		query["PreemptLogin"] = request.PreemptLogin
	}

	if !dara.IsNil(request.PreemptLoginUser) {
		query["PreemptLoginUser"] = request.PreemptLoginUser
	}

	if !dara.IsNil(request.PrinterRedirection) {
		query["PrinterRedirection"] = request.PrinterRedirection
	}

	if !dara.IsNil(request.RecordContent) {
		query["RecordContent"] = request.RecordContent
	}

	if !dara.IsNil(request.RecordContentExpires) {
		query["RecordContentExpires"] = request.RecordContentExpires
	}

	if !dara.IsNil(request.Recording) {
		query["Recording"] = request.Recording
	}

	if !dara.IsNil(request.RecordingAudio) {
		query["RecordingAudio"] = request.RecordingAudio
	}

	if !dara.IsNil(request.RecordingDuration) {
		query["RecordingDuration"] = request.RecordingDuration
	}

	if !dara.IsNil(request.RecordingEndTime) {
		query["RecordingEndTime"] = request.RecordingEndTime
	}

	if !dara.IsNil(request.RecordingExpires) {
		query["RecordingExpires"] = request.RecordingExpires
	}

	if !dara.IsNil(request.RecordingFps) {
		query["RecordingFps"] = request.RecordingFps
	}

	if !dara.IsNil(request.RecordingStartTime) {
		query["RecordingStartTime"] = request.RecordingStartTime
	}

	if !dara.IsNil(request.RecordingUserNotify) {
		query["RecordingUserNotify"] = request.RecordingUserNotify
	}

	if !dara.IsNil(request.RecordingUserNotifyMessage) {
		query["RecordingUserNotifyMessage"] = request.RecordingUserNotifyMessage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteCoordinate) {
		query["RemoteCoordinate"] = request.RemoteCoordinate
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	if !dara.IsNil(request.UsbRedirect) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !dara.IsNil(request.UsbSupplyRedirectRule) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !dara.IsNil(request.VideoRedirect) {
		query["VideoRedirect"] = request.VideoRedirect
	}

	if !dara.IsNil(request.VisualQuality) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	if !dara.IsNil(request.WatermarkAntiCam) {
		query["WatermarkAntiCam"] = request.WatermarkAntiCam
	}

	if !dara.IsNil(request.WatermarkColor) {
		query["WatermarkColor"] = request.WatermarkColor
	}

	if !dara.IsNil(request.WatermarkDegree) {
		query["WatermarkDegree"] = request.WatermarkDegree
	}

	if !dara.IsNil(request.WatermarkFontSize) {
		query["WatermarkFontSize"] = request.WatermarkFontSize
	}

	if !dara.IsNil(request.WatermarkFontStyle) {
		query["WatermarkFontStyle"] = request.WatermarkFontStyle
	}

	if !dara.IsNil(request.WatermarkPower) {
		query["WatermarkPower"] = request.WatermarkPower
	}

	if !dara.IsNil(request.WatermarkRowAmount) {
		query["WatermarkRowAmount"] = request.WatermarkRowAmount
	}

	if !dara.IsNil(request.WatermarkSecurity) {
		query["WatermarkSecurity"] = request.WatermarkSecurity
	}

	if !dara.IsNil(request.WatermarkTransparency) {
		query["WatermarkTransparency"] = request.WatermarkTransparency
	}

	if !dara.IsNil(request.WatermarkTransparencyValue) {
		query["WatermarkTransparencyValue"] = request.WatermarkTransparencyValue
	}

	if !dara.IsNil(request.WatermarkType) {
		query["WatermarkType"] = request.WatermarkType
	}

	if !dara.IsNil(request.WyAssistant) {
		query["WyAssistant"] = request.WyAssistant
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) directory.
//
// Description:
//
// Before you create a RAM directory, complete the following preparations:
//
//   - Call the `CreateVpc` operation to create a virtual private cloud (VPC) in a region supported by Elastic Desktop Service.
//
//   - Call the `CreateVSwitch` operation to create a vSwitch in the VPC. The vSwitch is in a zone that is supported by Elastic Desktop Service. You can call the [DescribeZones](https://help.aliyun.com/document_detail/196648.html) operation to obtain the most recent zone list for a region supported by Elastic Desktop Service
//
// @param request - CreateRAMDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRAMDirectoryResponse
func (client *Client) CreateRAMDirectoryWithContext(ctx context.Context, request *CreateRAMDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DirectoryName) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !dara.IsNil(request.EnableAdminAccess) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !dara.IsNil(request.EnableInternetAccess) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRAMDirectory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建路由条目
//
// @param request - CreateRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteEntryResponse
func (client *Client) CreateRouteEntryWithContext(ctx context.Context, request *CreateRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteEntryResponse, _err error) {
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

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.NextHopId) {
		query["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteEntryName) {
		query["RouteEntryName"] = request.RouteEntryName
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建自定义路由表
//
// @param request - CreateRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteTableResponse
func (client *Client) CreateRouteTableWithContext(ctx context.Context, request *CreateRouteTableRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteTableName) {
		query["RouteTableName"] = request.RouteTableName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteTable"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an office network of the convenience account type. Elastic Desktop Service supports the following types of accounts: convenience accounts and enterprise AD accounts.
//
// @param request - CreateSimpleOfficeSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimpleOfficeSiteResponse
func (client *Client) CreateSimpleOfficeSiteWithContext(ctx context.Context, request *CreateSimpleOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.AuthorityHost) {
		query["AuthorityHost"] = request.AuthorityHost
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["ClientSecret"] = request.ClientSecret
	}

	if !dara.IsNil(request.CloudBoxOfficeSite) {
		query["CloudBoxOfficeSite"] = request.CloudBoxOfficeSite
	}

	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Eid) {
		query["Eid"] = request.Eid
	}

	if !dara.IsNil(request.EnableAdminAccess) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !dara.IsNil(request.EnableInternetAccess) {
		query["EnableInternetAccess"] = request.EnableInternetAccess
	}

	if !dara.IsNil(request.NeedVerifyZeroDevice) {
		query["NeedVerifyZeroDevice"] = request.NeedVerifyZeroDevice
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	if !dara.IsNil(request.VpcType) {
		query["VpcType"] = request.VpcType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimpleOfficeSite"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a snapshot for a disk of a cloud computer to back up or restore the data on the disk.
//
// Description:
//
// The cloud computer must be in the **Running*	- or **Stopped*	- state.
//
// @param request - CreateSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !dara.IsNil(request.SourceDiskType) {
		query["SourceDiskType"] = request.SourceDiskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加SNAT条目
//
// @param request - CreateSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnatEntryResponse
func (client *Client) CreateSnatEntryWithContext(ctx context.Context, request *CreateSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipAffinity) {
		query["EipAffinity"] = request.EipAffinity
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnatEntryName) {
		query["SnatEntryName"] = request.SnatEntryName
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	if !dara.IsNil(request.SnatTableId) {
		query["SnatTableId"] = request.SnatTableId
	}

	if !dara.IsNil(request.SourceCIDR) {
		query["SourceCIDR"] = request.SourceCIDR
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnatEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建子网
//
// @param request - CreateSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubnetResponse
func (client *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest, runtime *dara.RuntimeOptions) (_result *CreateSubnetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubnet"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom cloud computer template. A cloud computer template (or simply "template") simplifies the process of creating cloud computers by providing a predefined set of configurations. This eliminates the need to manually configure each setting, saving significant time and effort.
//
// Description:
//
// When you call this operation, take note of the following item:
//
//   - Most parameters in templates are optional. When you create a template, Elastic Desktop Service (EDS) does not validate the existence or correctness of the parameter values you specify. The parameter values in the template are only verified when you use the template to create cloud computers.
//
//   - For parameters that include the region attribute in the template, it\\"s important to note that if the specified region doesn’t match the region where the template is used to create a cloud computer, those parameters will not take effect.
//
// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithContext(ctx context.Context, request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskList) {
		bodyFlat["DataDiskList"] = request.DataDiskList
	}

	if !dara.IsNil(request.DefaultLanguage) {
		body["DefaultLanguage"] = request.DefaultLanguage
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		body["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PostPaidAfterUsedUp) {
		body["PostPaidAfterUsedUp"] = request.PostPaidAfterUsedUp
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionConfigList) {
		bodyFlat["RegionConfigList"] = request.RegionConfigList
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceTagList) {
		bodyFlat["ResourceTagList"] = request.ResourceTagList
	}

	if !dara.IsNil(request.SiteConfigList) {
		bodyFlat["SiteConfigList"] = request.SiteConfigList
	}

	if !dara.IsNil(request.SystemDiskPerformanceLevel) {
		body["SystemDiskPerformanceLevel"] = request.SystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TimerGroupId) {
		body["TimerGroupId"] = request.TimerGroupId
	}

	if !dara.IsNil(request.UserDuration) {
		body["UserDuration"] = request.UserDuration
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic snapshot policy.
//
// @param request - DeleteAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicyWithContext(ctx context.Context, request *DeleteAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom cloud computer templates.
//
// @param request - DeleteBundlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBundlesResponse
func (client *Client) DeleteBundlesWithContext(ctx context.Context, request *DeleteBundlesRequest, runtime *dara.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBundles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete files or folders from the network disk.
//
// @param request - DeleteCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCdsFileResponse
func (client *Client) DeleteCdsFileWithContext(ctx context.Context, request *DeleteCdsFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Deletes a center policy
//
// @param request - DeleteCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenterPolicyResponse
func (client *Client) DeleteCenterPolicyWithContext(ctx context.Context, request *DeleteCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenterPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.PolicyGroupIds) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
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
		Action:      dara.String("DeleteCenterPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenterPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes team spaces.
//
// @param request - DeleteCloudDriveGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudDriveGroupsResponse
func (client *Client) DeleteCloudDriveGroupsWithContext(ctx context.Context, request *DeleteCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudDriveGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudDriveGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudDriveGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除无影网盘中的终端用户
//
// @param request - DeleteCloudDriveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudDriveUsersResponse
func (client *Client) DeleteCloudDriveUsersWithContext(ctx context.Context, request *DeleteCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudDriveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudDriveUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudDriveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a configuration group.
//
// @param request - DeleteConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigGroupResponse
func (client *Client) DeleteConfigGroupWithContext(ctx context.Context, request *DeleteConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a cloud computer share.
//
// Description:
//
//	  Before releasing a cloud computer share, ensure that no cloud computers within it are in the Connected state and that no end users have access permissions to it.
//
//		- You cannot delete a cloud computer share with an active subscription if it contains cloud computers that have not yet expired.
//
//		- Deleting a pay-as-you-go cloud computer share will release all pay-as-you-go cloud computers within it.
//
// @param request - DeleteDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDesktopGroupResponse
func (client *Client) DeleteDesktopGroupWithContext(ctx context.Context, request *DeleteDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases pay-as-you-go cloud computers or expired subscription cloud computers.
//
// @param request - DeleteDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDesktopsResponse
func (client *Client) DeleteDesktopsWithContext(ctx context.Context, request *DeleteDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes trusted devices.
//
// Description:
//
// You can call the operation to manage client devices.
//
// @param request - DeleteDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDevicesResponse
func (client *Client) DeleteDevicesWithContext(ctx context.Context, request *DeleteDevicesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceIds) {
		query["DeviceIds"] = request.DeviceIds
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
		Action:      dara.String("DeleteDevices"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more directories.
//
// Description:
//
// You cannot delete a directory that has a cloud computer or is used by a cloud computer.
//
// @param request - DeleteDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoriesResponse
func (client *Client) DeleteDirectoriesWithContext(ctx context.Context, request *DeleteDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDirectories"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a drive.
//
// @param request - DeleteDriveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDriveResponse
func (client *Client) DeleteDriveWithContext(ctx context.Context, request *DeleteDriveRequest, runtime *dara.RuntimeOptions) (_result *DeleteDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		query["DriveId"] = request.DriveId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDrive"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDriveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEduRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEduRoomResponse
func (client *Client) DeleteEduRoomWithContext(ctx context.Context, request *DeleteEduRoomRequest, runtime *dara.RuntimeOptions) (_result *DeleteEduRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EduRoomId) {
		query["EduRoomId"] = request.EduRoomId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEduRoom"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEduRoomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除DNAT条目
//
// @param request - DeleteForwardEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteForwardEntryResponse
func (client *Client) DeleteForwardEntryWithContext(ctx context.Context, request *DeleteForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForwardEntryId) {
		query["ForwardEntryId"] = request.ForwardEntryId
	}

	if !dara.IsNil(request.ForwardTableId) {
		query["ForwardTableId"] = request.ForwardTableId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteForwardEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteForwardEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more custom images.
//
// Description:
//
//	  Images include system images and custom images. System images cannot be deleted.
//
//		- If an image that you want to delete is referenced by a cloud computer template, call the [DeleteBundles](https://help.aliyun.com/document_detail/436972.html) operation to delete the cloud computer template before you delete the image.
//
// @param request - DeleteImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImagesResponse
func (client *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest, runtime *dara.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCascadedBundle) {
		query["DeleteCascadedBundle"] = request.DeleteCascadedBundle
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes NAS file systems.
//
// Description:
//
// Before you delete a File Storage NAS (NAS) file system, make sure that the data you want to retain is backed up.
//
//	Warning: If a NAS file system is deleted, data stored in the NAS file system cannot be restored. Proceed with caution when you delete NAS file systems.
//
// @param request - DeleteNASFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNASFileSystemsResponse
func (client *Client) DeleteNASFileSystemsWithContext(ctx context.Context, request *DeleteNASFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DeleteNASFileSystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNASFileSystems"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除NAT网关
//
// @param request - DeleteNatGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatGatewayResponse
func (client *Client) DeleteNatGatewayWithContext(ctx context.Context, request *DeleteNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNatGateway"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNatGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more premium bandwidth plans.
//
// @param request - DeleteNetworkPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkPackagesResponse
func (client *Client) DeleteNetworkPackagesWithContext(ctx context.Context, request *DeleteNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkPackages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes office networks (formerly workspaces).
//
// Description:
//
// Before you delete an office network, make sure that the following operations are complete:
//
//   - All cloud computers in the office network are released.
//
//   - The data that you want to retain is backed up.
//
// >  Resources and data on cloud computers in an office network cannot be restored after you delete it. Proceed with caution.
//
// @param request - DeleteOfficeSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOfficeSitesResponse
func (client *Client) DeleteOfficeSitesWithContext(ctx context.Context, request *DeleteOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DeleteOfficeSitesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOfficeSites"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more custom cloud computer policies.
//
// Description:
//
//	  You cannot delete the cloud computer policy created by the Elastic Desktop Service (EDS) system.
//
//		- You cannot delete the cloud computer policies that are associated with cloud computers.
//
// @param request - DeletePolicyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyGroupsResponse
func (client *Client) DeletePolicyGroupsWithContext(ctx context.Context, request *DeletePolicyGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除路由条目
//
// @param request - DeleteRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteEntryResponse
func (client *Client) DeleteRouteEntryWithContext(ctx context.Context, request *DeleteRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.NextHopId) {
		query["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteEntryId) {
		query["RouteEntryId"] = request.RouteEntryId
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义路由表
//
// @param request - DeleteRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteTableResponse
func (client *Client) DeleteRouteTableWithContext(ctx context.Context, request *DeleteRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteTable"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more snapshots.
//
// Description:
//
// If the IDs of the snapshots that you specify do not exist, requests are ignored.
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除SNAT条目
//
// @param request - DeleteSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnatEntryResponse
func (client *Client) DeleteSnatEntryWithContext(ctx context.Context, request *DeleteSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnatEntryResponse, _err error) {
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

	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatTableId) {
		query["SnatTableId"] = request.SnatTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnatEntry"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除子网
//
// @param request - DeleteSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubnetResponse
func (client *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubnetResponse, _err error) {
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

	if !dara.IsNil(request.SubnetId) {
		query["SubnetId"] = request.SubnetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubnet"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom cloud computer templates.
//
// Description:
//
// Deleting a template does not affect cloud computers created from it or the associated resources.
//
// @param request - DeleteTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplatesResponse
func (client *Client) DeleteTemplatesWithContext(ctx context.Context, request *DeleteTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.TemplateIds) {
		body["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplates"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete virtual multi-factor authentication (MFA) devices.
//
// Description:
//
// If an MFA device is deleted, the device is unbound, reset, and disabled. When an Active Directory (AD) user wants to connect to the cloud desktop that is bound to the MFA device, the AD user must bind a new MFA device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADeviceWithContext(ctx context.Context, request *DeleteVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirtualMFADevice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access control list (ACL) of an office network or a cloud computer.
//
// @param request - DescribeAclEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclEntriesResponse
func (client *Client) DescribeAclEntriesWithContext(ctx context.Context, request *DescribeAclEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclEntriesResponse, _err error) {
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

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAclEntries"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic snapshot policy.
//
// Description:
//
// You can view an automatic snapshot policy that is associated with a cloud desktop in the Elastic Desktop Service (EDS) console. To view the automatic snapshot policy, you can go to the EDS console, choose Deployment > Snapshots in the left-side navigation pane, and then view an automatic snapshot policy on the Snapshots page.
//
// @param request - DescribeAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotPolicyResponse
func (client *Client) DescribeAutoSnapshotPolicyWithContext(ctx context.Context, request *DescribeAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotPolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cloud computer templates.
//
// @param request - DescribeBundlesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBundlesResponse
func (client *Client) DescribeBundlesWithContext(ctx context.Context, request *DescribeBundlesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.BundleType) {
		query["BundleType"] = request.BundleType
	}

	if !dara.IsNil(request.CheckStock) {
		query["CheckStock"] = request.CheckStock
	}

	if !dara.IsNil(request.CpuCount) {
		query["CpuCount"] = request.CpuCount
	}

	if !dara.IsNil(request.DesktopTypeFamily) {
		query["DesktopTypeFamily"] = request.DesktopTypeFamily
	}

	if !dara.IsNil(request.FotaChannel) {
		query["FotaChannel"] = request.FotaChannel
	}

	if !dara.IsNil(request.FromDesktopGroup) {
		query["FromDesktopGroup"] = request.FromDesktopGroup
	}

	if !dara.IsNil(request.GpuCount) {
		query["GpuCount"] = request.GpuCount
	}

	if !dara.IsNil(request.GpuDriverType) {
		query["GpuDriverType"] = request.GpuDriverType
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemorySize) {
		query["MemorySize"] = request.MemorySize
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SelectedBundle) {
		query["SelectedBundle"] = request.SelectedBundle
	}

	if !dara.IsNil(request.SessionType) {
		query["SessionType"] = request.SessionType
	}

	if !dara.IsNil(request.SupportMultiSession) {
		query["SupportMultiSession"] = request.SupportMultiSession
	}

	if !dara.IsNil(request.VolumeEncryptionEnabled) {
		query["VolumeEncryptionEnabled"] = request.VolumeEncryptionEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBundles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries file sharing links of a cloud disk in Cloud Drive Service.
//
// @param request - DescribeCdsFileShareLinksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdsFileShareLinksResponse
func (client *Client) DescribeCdsFileShareLinksWithContext(ctx context.Context, request *DescribeCdsFileShareLinksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdsFileShareLinksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.Creators) {
		query["Creators"] = request.Creators
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ShareId) {
		query["ShareId"] = request.ShareId
	}

	if !dara.IsNil(request.ShareName) {
		query["ShareName"] = request.ShareName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdsFileShareLinks"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdsFileShareLinksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all Cloud Enterprise Network (CEN) instances within an Alibaba Cloud account.
//
// @param request - DescribeCensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCensResponse
func (client *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCens"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries center policies.
//
// @param request - DescribeCenterPolicyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenterPolicyListResponse
func (client *Client) DescribeCenterPolicyListWithContext(ctx context.Context, request *DescribeCenterPolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenterPolicyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcademicProxy) {
		query["AcademicProxy"] = request.AcademicProxy
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.ModelLibrary) {
		query["ModelLibrary"] = request.ModelLibrary
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PortProxy) {
		query["PortProxy"] = request.PortProxy
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenterPolicyList"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenterPolicyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of end users. For example, the logs record the events that end users start and stop cloud desktops, and disconnect desktop sessions.
//
// Description:
//
// You can audit the operation logs of regular users to improve security. The operation logs record events such as desktop startup, shutdown, and session disconnection.
//
// @param request - DescribeClientEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientEventsResponse
func (client *Client) DescribeClientEventsWithContext(ctx context.Context, request *DescribeClientEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopIp) {
		query["DesktopIp"] = request.DesktopIp
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.EventTypes) {
		query["EventTypes"] = request.EventTypes
	}

	if !dara.IsNil(request.FillHardwareInfo) {
		query["FillHardwareInfo"] = request.FillHardwareInfo
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClientEvents"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云盘团队空间列表
//
// @param request - DescribeCloudDiskGroupDrivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDiskGroupDrivesResponse
func (client *Client) DescribeCloudDiskGroupDrivesWithContext(ctx context.Context, request *DescribeCloudDiskGroupDrivesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDiskGroupDrivesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
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
		Action:      dara.String("DescribeCloudDiskGroupDrives"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDiskGroupDrivesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云盘团队列表
//
// @param request - DescribeCloudDiskGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDiskGroupsResponse
func (client *Client) DescribeCloudDiskGroupsWithContext(ctx context.Context, request *DescribeCloudDiskGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDiskGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ParentOrgId) {
		query["ParentOrgId"] = request.ParentOrgId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudDiskGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDiskGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of authorized team spaces.
//
// @param request - DescribeCloudDriveGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDriveGroupsResponse
func (client *Client) DescribeCloudDriveGroupsWithContext(ctx context.Context, request *DescribeCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDriveGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DirectoryName) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !dara.IsNil(request.DriveStatus) {
		query["DriveStatus"] = request.DriveStatus
	}

	if !dara.IsNil(request.DriveType) {
		query["DriveType"] = request.DriveType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentGroupId) {
		query["ParentGroupId"] = request.ParentGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudDriveGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDriveGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pds用户权限
//
// @param request - DescribeCloudDrivePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDrivePermissionsResponse
func (client *Client) DescribeCloudDrivePermissionsWithContext(ctx context.Context, request *DescribeCloudDrivePermissionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDrivePermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudDrivePermissions"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDrivePermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有无影网盘终端用户的信息
//
// @param request - DescribeCloudDriveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDriveUsersResponse
func (client *Client) DescribeCloudDriveUsersWithContext(ctx context.Context, request *DescribeCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDriveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
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
		Action:      dara.String("DescribeCloudDriveUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDriveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration groups.
//
// @param request - DescribeConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigGroupResponse
func (client *Client) DescribeConfigGroupWithContext(ctx context.Context, request *DescribeConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
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

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCustomizedListHeadersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizedListHeadersResponse
func (client *Client) DescribeCustomizedListHeadersWithContext(ctx context.Context, request *DescribeCustomizedListHeadersRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomizedListHeadersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LangType) {
		query["LangType"] = request.LangType
	}

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomizedListHeaders"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomizedListHeadersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud computer shares.
//
// @param request - DescribeDesktopGroupSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopGroupSessionsResponse
func (client *Client) DescribeDesktopGroupSessionsWithContext(ctx context.Context, request *DescribeDesktopGroupSessionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopGroupSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupIds) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !dara.IsNil(request.DesktopGroupName) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FillTerminalInfo) {
		query["FillTerminalInfo"] = request.FillTerminalInfo
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnType) {
		query["OwnType"] = request.OwnType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionStatus) {
		query["SessionStatus"] = request.SessionStatus
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopGroupSessions"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopGroupSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud computer shares.
//
// @param request - DescribeDesktopGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopGroupsResponse
func (client *Client) DescribeDesktopGroupsWithContext(ctx context.Context, request *DescribeDesktopGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupIds) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !dara.IsNil(request.DesktopGroupName) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.ExcludedEndUserIds) {
		query["ExcludedEndUserIds"] = request.ExcludedEndUserIds
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MultiResource) {
		query["MultiResource"] = request.MultiResource
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OwnType) {
		query["OwnType"] = request.OwnType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about cloud computers.
//
// @param request - DescribeDesktopInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopInfoResponse
func (client *Client) DescribeDesktopInfoWithContext(ctx context.Context, request *DescribeDesktopInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.NeedExtraInfo) {
		query["NeedExtraInfo"] = request.NeedExtraInfo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopInfo"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云电脑基础元数据
//
// @param request - DescribeDesktopMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopMetadataResponse
func (client *Client) DescribeDesktopMetadataWithContext(ctx context.Context, request *DescribeDesktopMetadataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreationTimeStart) {
		query["CreationTimeStart"] = request.CreationTimeStart
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.IncludeDesktopGroup) {
		query["IncludeDesktopGroup"] = request.IncludeDesktopGroup
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OperationTimeStart) {
		query["OperationTimeStart"] = request.OperationTimeStart
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopMetadata"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询超卖组
//
// @param request - DescribeDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopOversoldGroupResponse
func (client *Client) DescribeDesktopOversoldGroupWithContext(ctx context.Context, request *DescribeDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldGroupResponse, _err error) {
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

	if !dara.IsNil(request.OversoldGroupIds) {
		query["OversoldGroupIds"] = request.OversoldGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询超卖组用户
//
// @param request - DescribeDesktopOversoldUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopOversoldUserResponse
func (client *Client) DescribeDesktopOversoldUserWithContext(ctx context.Context, request *DescribeDesktopOversoldUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldUserResponse, _err error) {
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

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.UserDesktopIds) {
		query["UserDesktopIds"] = request.UserDesktopIds
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopOversoldUser"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopOversoldUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询超卖用户组
//
// @param request - DescribeDesktopOversoldUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopOversoldUserGroupResponse
func (client *Client) DescribeDesktopOversoldUserGroupWithContext(ctx context.Context, request *DescribeDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopOversoldUserGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopOversoldUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed session information of a cloud computer.
//
// Description:
//
// You can only query data within the last 30 days.
//
// @param request - DescribeDesktopSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopSessionsResponse
func (client *Client) DescribeDesktopSessionsWithContext(ctx context.Context, request *DescribeDesktopSessionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckOsSession) {
		query["CheckOsSession"] = request.CheckOsSession
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EndUserIdFilter) {
		query["EndUserIdFilter"] = request.EndUserIdFilter
	}

	if !dara.IsNil(request.FillHardwareInfo) {
		query["FillHardwareInfo"] = request.FillHardwareInfo
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SessionStatus) {
		query["SessionStatus"] = request.SessionStatus
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubPayType) {
		query["SubPayType"] = request.SubPayType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopSessions"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance types of cloud computers.
//
// Description:
//
// When no values are specified for the `InstanceTypeFamily` and `DesktopTypeId` parameters for a cloud desktop, all types of cloud desktops are queried.
//
// @param request - DescribeDesktopTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopTypesResponse
func (client *Client) DescribeDesktopTypesWithContext(ctx context.Context, request *DescribeDesktopTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppliedScope) {
		query["AppliedScope"] = request.AppliedScope
	}

	if !dara.IsNil(request.CpuCount) {
		query["CpuCount"] = request.CpuCount
	}

	if !dara.IsNil(request.DesktopGroupIdForModify) {
		query["DesktopGroupIdForModify"] = request.DesktopGroupIdForModify
	}

	if !dara.IsNil(request.DesktopIdForModify) {
		query["DesktopIdForModify"] = request.DesktopIdForModify
	}

	if !dara.IsNil(request.DesktopTypeId) {
		query["DesktopTypeId"] = request.DesktopTypeId
	}

	if !dara.IsNil(request.DesktopTypeIdList) {
		query["DesktopTypeIdList"] = request.DesktopTypeIdList
	}

	if !dara.IsNil(request.GpuCount) {
		query["GpuCount"] = request.GpuCount
	}

	if !dara.IsNil(request.GpuDriverType) {
		query["GpuDriverType"] = request.GpuDriverType
	}

	if !dara.IsNil(request.GpuMemory) {
		query["GpuMemory"] = request.GpuMemory
	}

	if !dara.IsNil(request.InstanceTypeFamily) {
		query["InstanceTypeFamily"] = request.InstanceTypeFamily
	}

	if !dara.IsNil(request.MemorySize) {
		query["MemorySize"] = request.MemorySize
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeSet) {
		query["ScopeSet"] = request.ScopeSet
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.SupportMinSessionCount) {
		query["SupportMinSessionCount"] = request.SupportMinSessionCount
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopTypes"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cloud computers.
//
// @param request - DescribeDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopsResponse
func (client *Client) DescribeDesktopsWithContext(ctx context.Context, request *DescribeDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DesktopStatus) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !dara.IsNil(request.DesktopStatusList) {
		query["DesktopStatusList"] = request.DesktopStatusList
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExcludedEndUserId) {
		query["ExcludedEndUserId"] = request.ExcludedEndUserId
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.FillResourceGroup) {
		query["FillResourceGroup"] = request.FillResourceGroup
	}

	if !dara.IsNil(request.FilterDesktopGroup) {
		query["FilterDesktopGroup"] = request.FilterDesktopGroup
	}

	if !dara.IsNil(request.GpuInstanceGroupId) {
		query["GpuInstanceGroupId"] = request.GpuInstanceGroupId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.IncludeAutoSnapshotPolicy) {
		query["IncludeAutoSnapshotPolicy"] = request.IncludeAutoSnapshotPolicy
	}

	if !dara.IsNil(request.ManagementFlag) {
		query["ManagementFlag"] = request.ManagementFlag
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MultiResource) {
		query["MultiResource"] = request.MultiResource
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.OnlyDesktopGroup) {
		query["OnlyDesktopGroup"] = request.OnlyDesktopGroup
	}

	if !dara.IsNil(request.OsTypes) {
		query["OsTypes"] = request.OsTypes
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.QosRuleId) {
		query["QosRuleId"] = request.QosRuleId
	}

	if !dara.IsNil(request.QueryFotaUpdate) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SnapshotPolicyId) {
		query["SnapshotPolicyId"] = request.SnapshotPolicyId
	}

	if !dara.IsNil(request.SubPayType) {
		query["SubPayType"] = request.SubPayType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud computers in a share by billing method.
//
// @param request - DescribeDesktopsInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopsInGroupResponse
func (client *Client) DescribeDesktopsInGroupWithContext(ctx context.Context, request *DescribeDesktopsInGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopsInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomEndTimePeriod) {
		query["CustomEndTimePeriod"] = request.CustomEndTimePeriod
	}

	if !dara.IsNil(request.CustomStartTimePeriod) {
		query["CustomStartTimePeriod"] = request.CustomStartTimePeriod
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.IgnoreDeleted) {
		query["IgnoreDeleted"] = request.IgnoreDeleted
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktopsInGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of trusted devices.
//
// @param request - DescribeDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDevicesResponse
func (client *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDevices"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of directories.
//
// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithContext(ctx context.Context, request *DescribeDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DirectoryStatus) {
		query["DirectoryStatus"] = request.DirectoryStatus
	}

	if !dara.IsNil(request.DirectoryType) {
		query["DirectoryType"] = request.DirectoryType
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirectories"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user-level storage resources.
//
// @param request - DescribeDrivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrivesResponse
func (client *Client) DescribeDrivesWithContext(ctx context.Context, request *DescribeDrivesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrivesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainIds) {
		query["DomainIds"] = request.DomainIds
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

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrives"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrivesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data report export tasks.
//
// @param request - DescribeEcdReportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEcdReportTasksResponse
func (client *Client) DescribeEcdReportTasksWithContext(ctx context.Context, request *DescribeEcdReportTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeEcdReportTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SubType) {
		query["SubType"] = request.SubType
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEcdReportTasks"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEcdReportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询EIP监控
//
// @param request - DescribeFlowMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowMetricResponse
func (client *Client) DescribeFlowMetricWithContext(ctx context.Context, request *DescribeFlowMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowMetric"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud computer-level traffic statistics of a single office network.
//
// Description:
//
// > You can query only the traffic data in the last 90 days.
//
// @param request - DescribeFlowStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowStatisticResponse
func (client *Client) DescribeFlowStatisticWithContext(ctx context.Context, request *DescribeFlowStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
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

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowStatistic"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询DNAT条目
//
// @param request - DescribeForwardTableEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeForwardTableEntriesResponse
func (client *Client) DescribeForwardTableEntriesWithContext(ctx context.Context, request *DescribeForwardTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeForwardTableEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForwardEntryId) {
		query["ForwardEntryId"] = request.ForwardEntryId
	}

	if !dara.IsNil(request.ForwardTableId) {
		query["ForwardTableId"] = request.ForwardTableId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
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
		Action:      dara.String("DescribeForwardTableEntries"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeForwardTableEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the cloud computers whose images can be and are pending to be updated to the specified version.
//
// @param request - DescribeFotaPendingDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFotaPendingDesktopsResponse
func (client *Client) DescribeFotaPendingDesktopsWithContext(ctx context.Context, request *DescribeFotaPendingDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFotaPendingDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskUid) {
		query["TaskUid"] = request.TaskUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFotaPendingDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFotaPendingDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of update tasks.
//
// @param request - DescribeFotaTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFotaTasksResponse
func (client *Client) DescribeFotaTasksWithContext(ctx context.Context, request *DescribeFotaTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeFotaTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FotaStatus) {
		query["FotaStatus"] = request.FotaStatus
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.TaskUid) {
		query["TaskUid"] = request.TaskUid
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFotaTasks"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFotaTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information of all cloud computers and the corresponding usage duration records.
//
// Description:
//
//	  Domestic site users query site selection Shanghai, international site users choose Singapore.
//
//		- By default, you can query all cloud computers that are deleted or not deleted.
//
//		- Deleted cloud computers can be queried only if the deletion time is less than three months.
//
//		- Sort criteria cannot be shared with other criteria.
//
// @param request - DescribeGlobalDesktopRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDesktopRecordsResponse
func (client *Client) DescribeGlobalDesktopRecordsWithContext(ctx context.Context, request *DescribeGlobalDesktopRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDesktopRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DesktopStatusList) {
		query["DesktopStatusList"] = request.DesktopStatusList
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExcludeDesktopStatusList) {
		query["ExcludeDesktopStatusList"] = request.ExcludeDesktopStatusList
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubPayType) {
		query["SubPayType"] = request.SubPayType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDesktopRecords"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDesktopRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全局定时任务Batch记录
//
// @param request - DescribeGlobalTimerBatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalTimerBatchesResponse
func (client *Client) DescribeGlobalTimerBatchesWithContext(ctx context.Context, request *DescribeGlobalTimerBatchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalTimerBatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !dara.IsNil(request.TimerType) {
		query["TimerType"] = request.TimerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalTimerBatches"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalTimerBatchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution records of scheduled tasks on cloud computers.
//
// @param request - DescribeGlobalTimerRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalTimerRecordsResponse
func (client *Client) DescribeGlobalTimerRecordsWithContext(ctx context.Context, request *DescribeGlobalTimerRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalTimerRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.DisplayResultName) {
		query["DisplayResultName"] = request.DisplayResultName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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

	if !dara.IsNil(request.ResultCategory) {
		query["ResultCategory"] = request.ResultCategory
	}

	if !dara.IsNil(request.Retryable) {
		query["Retryable"] = request.Retryable
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !dara.IsNil(request.TimerResult) {
		query["TimerResult"] = request.TimerResult
	}

	if !dara.IsNil(request.TimerTypes) {
		query["TimerTypes"] = request.TimerTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalTimerRecords"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalTimerRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications and their processes of an end user.
//
// @param request - DescribeGuestApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGuestApplicationsResponse
func (client *Client) DescribeGuestApplicationsWithContext(ctx context.Context, request *DescribeGuestApplicationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGuestApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGuestApplications"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGuestApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the image modification records of cloud computers.
//
// @param request - DescribeImageModifiedRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageModifiedRecordsResponse
func (client *Client) DescribeImageModifiedRecordsWithContext(ctx context.Context, request *DescribeImageModifiedRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageModifiedRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
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
		Action:      dara.String("DescribeImageModifiedRecords"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageModifiedRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recipient Alibaba Cloud accounts with which an image is shared.
//
// Description:
//
// You can call the [ModifyImagePermission](https://help.aliyun.com/document_detail/436982.html) operation to share an image with another cloud computer user or unshare an image. You can call the DescribeImagePermission operation to obtain the Alibaba Cloud accounts with which the current image is shared.
//
// @param request - DescribeImagePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImagePermissionResponse
func (client *Client) DescribeImagePermissionWithContext(ctx context.Context, request *DescribeImagePermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeImagePermissionResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImagePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImagePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about images.
//
// @param request - DescribeImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImagesResponse
func (client *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopInstanceType) {
		query["DesktopInstanceType"] = request.DesktopInstanceType
	}

	if !dara.IsNil(request.FotaVersion) {
		query["FotaVersion"] = request.FotaVersion
	}

	if !dara.IsNil(request.GpuCategory) {
		query["GpuCategory"] = request.GpuCategory
	}

	if !dara.IsNil(request.GpuDriverVersion) {
		query["GpuDriverVersion"] = request.GpuDriverVersion
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ImageStatus) {
		query["ImageStatus"] = request.ImageStatus
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.LanguageType) {
		query["LanguageType"] = request.LanguageType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionType) {
		query["SessionType"] = request.SessionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
//	  After you run a command, it may not succeed. You can call this operation to query the execution result.
//
//		- You can query the information about execution in the last two weeks. A maximum of 100,000 lines of execution information can be retained.
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithContext(ctx context.Context, request *DescribeInvocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandType) {
		query["CommandType"] = request.CommandType
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.IncludeInvokeDesktops) {
		query["IncludeInvokeDesktops"] = request.IncludeInvokeDesktops
	}

	if !dara.IsNil(request.IncludeOutput) {
		query["IncludeOutput"] = request.IncludeOutput
	}

	if !dara.IsNil(request.InvokeId) {
		query["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.InvokeStatus) {
		query["InvokeStatus"] = request.InvokeStatus
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
		Action:      dara.String("DescribeInvocations"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公网IP
//
// @param request - DescribeIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpAddressesResponse
func (client *Client) DescribeIpAddressesWithContext(ctx context.Context, request *DescribeIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipId) {
		query["EipId"] = request.EipId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpAddresses"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Key Management Service (KMS) keys of users. The first time you call this operation, you can try to create a service key for Elastic Desktop Service (EDS) and call the operation to return results.
//
// @param request - DescribeKmsKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKmsKeysResponse
func (client *Client) DescribeKmsKeysWithContext(ctx context.Context, request *DescribeKmsKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeKmsKeysResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKmsKeys"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKmsKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for changing the specifications of a monthly subscription cloud computer with unlimited hours or a premium bandwidth plan.
//
// @param request - DescribeModificationPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModificationPriceResponse
func (client *Client) DescribeModificationPriceWithContext(ctx context.Context, request *DescribeModificationPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeModificationPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceSpecs) {
		query["ResourceSpecs"] = request.ResourceSpecs
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RootDiskPerformanceLevel) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !dara.IsNil(request.RootDiskSizeGib) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !dara.IsNil(request.UserDiskPerformanceLevel) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !dara.IsNil(request.UserDiskSizeGib) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModificationPrice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about File Storage NAS (NAS) file systems.
//
// @param request - DescribeNASFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNASFileSystemsResponse
func (client *Client) DescribeNASFileSystemsWithContext(ctx context.Context, request *DescribeNASFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNASFileSystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MatchCompatibleProfile) {
		query["MatchCompatibleProfile"] = request.MatchCompatibleProfile
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNASFileSystems"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询NAT详细列表
//
// @param request - DescribeNatGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatGatewaysResponse
func (client *Client) DescribeNatGatewaysWithContext(ctx context.Context, request *DescribeNatGatewaysRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatGatewaysResponse, _err error) {
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

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNatGateways"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more premium bandwidth plans.
//
// @param request - DescribeNetworkPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkPackagesResponse
func (client *Client) DescribeNetworkPackagesWithContext(ctx context.Context, request *DescribeNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
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
		Action:      dara.String("DescribeNetworkPackages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries office network properties, including office network ID, name, status, and creation time.
//
// @param request - DescribeOfficeSitesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSitesWithContext(ctx context.Context, request *DescribeOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteType) {
		query["OfficeSiteType"] = request.OfficeSiteType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityProtection) {
		query["SecurityProtection"] = request.SecurityProtection
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOfficeSites"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cloud computer policy.
//
// @param request - DescribePolicyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyGroupsResponse
func (client *Client) DescribePolicyGroupsWithContext(ctx context.Context, request *DescribePolicyGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.ExternalPolicyGroupIds) {
		query["ExternalPolicyGroupIds"] = request.ExternalPolicyGroupIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prices of Elastic Desktop Service (EDS) resources.
//
// Description:
//
// ## Usage notes
//
// The request parameters vary based on the type of desktop resources whose price you want to query. Take note of the following items:
//
//   - If you set ResourceType to OfficeSite, you must specify InstanceType.
//
//   - If you set ResourceType to Bandwidth, the pay-by-data-transfer metering method is used for network billing.
//
//   - If you set ResourceType to Desktop, you must specify InstanceType, RootDiskSizeGib, and UserDiskSizeGib. You can specify OsType, PeriodUnit, Period, and Amount based on your business requirements.
//
// > Before you call this operation to query the prices of cloud desktops by setting ResourceType to Desktop, you must know the desktop types and disk sizes that EDS provides. The disk sizes vary based on the desktop types. For more information, see [Cloud desktop types](https://help.aliyun.com/document_detail/188609.html).
//
// @param request - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
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

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.GroupDesktopCount) {
		query["GroupDesktopCount"] = request.GroupDesktopCount
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RootDiskCategory) {
		query["RootDiskCategory"] = request.RootDiskCategory
	}

	if !dara.IsNil(request.RootDiskPerformanceLevel) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !dara.IsNil(request.RootDiskSizeGib) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !dara.IsNil(request.UserDiskCategory) {
		query["UserDiskCategory"] = request.UserDiskCategory
	}

	if !dara.IsNil(request.UserDiskPerformanceLevel) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !dara.IsNil(request.UserDiskSizeGib) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询价格用于创建超卖组
//
// @param request - DescribePriceForCreateDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceForCreateDesktopOversoldGroupResponse
func (client *Client) DescribePriceForCreateDesktopOversoldGroupWithContext(ctx context.Context, request *DescribePriceForCreateDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForCreateDesktopOversoldGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrenceCount) {
		query["ConcurrenceCount"] = request.ConcurrenceCount
	}

	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.OversoldUserCount) {
		query["OversoldUserCount"] = request.OversoldUserCount
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePriceForCreateDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceForCreateDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询价格用于变配超卖组
//
// @param request - DescribePriceForModifyDesktopOversoldGroupSaleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceForModifyDesktopOversoldGroupSaleResponse
func (client *Client) DescribePriceForModifyDesktopOversoldGroupSaleWithContext(ctx context.Context, request *DescribePriceForModifyDesktopOversoldGroupSaleRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForModifyDesktopOversoldGroupSaleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrenceCount) {
		query["ConcurrenceCount"] = request.ConcurrenceCount
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.OversoldUserCount) {
		query["OversoldUserCount"] = request.OversoldUserCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePriceForModifyDesktopOversoldGroupSale"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceForModifyDesktopOversoldGroupSaleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询价格用于续费超卖组
//
// @param request - DescribePriceForRenewDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceForRenewDesktopOversoldGroupResponse
func (client *Client) DescribePriceForRenewDesktopOversoldGroupWithContext(ctx context.Context, request *DescribePriceForRenewDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForRenewDesktopOversoldGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePriceForRenewDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceForRenewDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询录屏文件列表
//
// @param request - DescribeRecordFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordFileResponse
func (client *Client) DescribeRecordFileWithContext(ctx context.Context, request *DescribeRecordFileRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderSort) {
		query["OrderSort"] = request.OrderSort
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of screen recording files.
//
// @param request - DescribeRecordingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordingsResponse
func (client *Client) DescribeRecordingsWithContext(ctx context.Context, request *DescribeRecordingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedSignedUrl) {
		query["NeedSignedUrl"] = request.NeedSignedUrl
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SignedUrlExpireMinutes) {
		query["SignedUrlExpireMinutes"] = request.SignedUrlExpireMinutes
	}

	if !dara.IsNil(request.StandardEndTime) {
		query["StandardEndTime"] = request.StandardEndTime
	}

	if !dara.IsNil(request.StandardStartTime) {
		query["StandardStartTime"] = request.StandardStartTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordings"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the refund amount for unsubscribing from a cloud computer.
//
// @param request - DescribeRefundPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRefundPriceResponse
func (client *Client) DescribeRefundPriceWithContext(ctx context.Context, request *DescribeRefundPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRefundPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RefundType) {
		query["RefundType"] = request.RefundType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRefundPrice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRefundPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are available for Elastic Desktop Service (EDS).
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the renewal price of an Alibaba Cloud Workspace service.
//
// @param request - DescribeRenewalPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenewalPriceResponse
func (client *Client) DescribeRenewalPriceWithContext(ctx context.Context, request *DescribeRenewalPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenewalPriceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenewalPrice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenewalPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeResourceByCenterPolicyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceByCenterPolicyIdResponse
func (client *Client) DescribeResourceByCenterPolicyIdWithContext(ctx context.Context, request *DescribeResourceByCenterPolicyIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceByCenterPolicyIdResponse, _err error) {
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

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceByCenterPolicyId"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceByCenterPolicyIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询路由条目列表
//
// @param request - DescribeRouteEntryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteEntryListResponse
func (client *Client) DescribeRouteEntryListWithContext(ctx context.Context, request *DescribeRouteEntryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteEntryListResponse, _err error) {
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

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRouteEntryList"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRouteEntryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询路由表列表
//
// @param request - DescribeRouteTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteTableListResponse
func (client *Client) DescribeRouteTableListWithContext(ctx context.Context, request *DescribeRouteTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteTableListResponse, _err error) {
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

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.RouteTableName) {
		query["RouteTableName"] = request.RouteTableName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRouteTableList"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRouteTableListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询办公网络维度安全组策略
//
// @param request - DescribeSecurityGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupAttributeResponse
func (client *Client) DescribeSecurityGroupAttributeWithContext(ctx context.Context, request *DescribeSecurityGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityGroupAttribute"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the session statistics of a region.
//
// Description:
//
//	  This is a central operation and can be called only by using services in the China (Shanghai) region.
//
//		- You can query session statistics for the past hour.
//
// @param request - DescribeSessionStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSessionStatisticResponse
func (client *Client) DescribeSessionStatisticWithContext(ctx context.Context, request *DescribeSessionStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSessionStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSessionStatistic"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSessionStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the snapshots that are created based on a cloud computer and the details of the snapshots.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !dara.IsNil(request.SnapshotType) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !dara.IsNil(request.SourceDiskType) {
		query["SourceDiskType"] = request.SourceDiskType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshots"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询SNAT条目
//
// @param request - DescribeSnatTableEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnatTableEntriesResponse
func (client *Client) DescribeSnatTableEntriesWithContext(ctx context.Context, request *DescribeSnatTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnatTableEntriesResponse, _err error) {
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

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatTableId) {
		query["SnatTableId"] = request.SnatTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnatTableEntries"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnatTableEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询子网
//
// @param request - DescribeSubnetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubnetsResponse
func (client *Client) DescribeSubnetsWithContext(ctx context.Context, request *DescribeSubnetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubnetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Eid) {
		query["Eid"] = request.Eid
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

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		query["SubnetId"] = request.SubnetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubnets"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubnetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cloud computer templates.
//
// @param request - DescribeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
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

	if !dara.IsNil(request.TemplateIds) {
		body["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplates"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a scheduled task configuration group.
//
// @param request - DescribeTimerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTimerGroupResponse
func (client *Client) DescribeTimerGroupWithContext(ctx context.Context, request *DescribeTimerGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeTimerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTimerGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTimerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询超卖组中用户连接数据
//
// @param request - DescribeUserConnectTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserConnectTimeResponse
func (client *Client) DescribeUserConnectTimeWithContext(ctx context.Context, request *DescribeUserConnectTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserConnectTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserDesktopId) {
		query["UserDesktopId"] = request.UserDesktopId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserConnectTime"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserConnectTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connection records of an authorized user to cloud computers in a cloud computer pool.
//
// @param request - DescribeUserConnectionRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserConnectionRecordsResponse
func (client *Client) DescribeUserConnectionRecordsWithContext(ctx context.Context, request *DescribeUserConnectionRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectDurationFrom) {
		query["ConnectDurationFrom"] = request.ConnectDurationFrom
	}

	if !dara.IsNil(request.ConnectDurationTo) {
		query["ConnectDurationTo"] = request.ConnectDurationTo
	}

	if !dara.IsNil(request.ConnectEndTimeFrom) {
		query["ConnectEndTimeFrom"] = request.ConnectEndTimeFrom
	}

	if !dara.IsNil(request.ConnectEndTimeTo) {
		query["ConnectEndTimeTo"] = request.ConnectEndTimeTo
	}

	if !dara.IsNil(request.ConnectStartTimeFrom) {
		query["ConnectStartTimeFrom"] = request.ConnectStartTimeFrom
	}

	if !dara.IsNil(request.ConnectStartTimeTo) {
		query["ConnectStartTimeTo"] = request.ConnectStartTimeTo
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EndUserType) {
		query["EndUserType"] = request.EndUserType
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
		Action:      dara.String("DescribeUserConnectionRecords"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the user profile management (UPM) directory blacklist and whitelist.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserProfilePathRules"),
		Version:     dara.String("2020-09-30"),
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
// Queries the information about authorized users of a cloud computer share, including the usernames, email addresses, mobile numbers, and cloud computer IDs.
//
// @param request - DescribeUsersInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersInGroupResponse
func (client *Client) DescribeUsersInGroupWithContext(ctx context.Context, request *DescribeUsersInGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectState) {
		query["ConnectState"] = request.ConnectState
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
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

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.QueryUserDetail) {
		query["QueryUserDetail"] = request.QueryUserDetail
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsersInGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the passwords of authorized users of a cloud computer.
//
// @param request - DescribeUsersPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersPasswordResponse
func (client *Client) DescribeUsersPasswordWithContext(ctx context.Context, request *DescribeUsersPasswordRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsersPassword"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsersPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices that are bound to an Active Directory (AD) account.
//
// @param request - DescribeVirtualMFADevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVirtualMFADevicesResponse
func (client *Client) DescribeVirtualMFADevicesWithContext(ctx context.Context, request *DescribeVirtualMFADevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
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

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVirtualMFADevices"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones in a region in which Elastic Desktop Service is supported.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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

	if !dara.IsNil(request.ZoneType) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an advanced office network from a CEN instance.
//
// @param request - DetachCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachCenResponse
func (client *Client) DetachCenWithContext(ctx context.Context, request *DetachCenRequest, runtime *dara.RuntimeOptions) (_result *DetachCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachCen"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a hardware client from a user.
//
// @param request - DetachEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEndUserResponse
func (client *Client) DetachEndUserWithContext(ctx context.Context, request *DetachEndUserRequest, runtime *dara.RuntimeOptions) (_result *DetachEndUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachEndUser"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachEndUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables specific cloud computers in a cloud computer share. After you call this operation to disable specific cloud computers, they enter the unavailable state.
//
// @param request - DisableDesktopsInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDesktopsInGroupResponse
func (client *Client) DisableDesktopsInGroupWithContext(ctx context.Context, request *DisableDesktopsInGroupRequest, runtime *dara.RuntimeOptions) (_result *DisableDesktopsInGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDesktopsInGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDesktopsInGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disconnects from desktop sessions.
//
// @param request - DisconnectDesktopSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisconnectDesktopSessionsResponse
func (client *Client) DisconnectDesktopSessionsWithContext(ctx context.Context, request *DisconnectDesktopSessionsRequest, runtime *dara.RuntimeOptions) (_result *DisconnectDesktopSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PreCheck) {
		query["PreCheck"] = request.PreCheck
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Sessions) {
		query["Sessions"] = request.Sessions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisconnectDesktopSessions"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisconnectDesktopSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例解绑/删除公网IP
//
// @param request - DissociateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateIpAddressResponse
func (client *Client) DissociateIpAddressWithContext(ctx context.Context, request *DissociateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *DissociateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipId) {
		query["EipId"] = request.EipId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateIpAddress"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a premium bandwidth plan from an office network.
//
// @param request - DissociateNetworkPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateNetworkPackageResponse
func (client *Client) DissociateNetworkPackageWithContext(ctx context.Context, request *DissociateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *DissociateNetworkPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateNetworkPackage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateNetworkPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download link of the target file.
//
// @param request - DownloadCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadCdsFileResponse
func (client *Client) DownloadCdsFileWithContext(ctx context.Context, request *DownloadCdsFileRequest, runtime *dara.RuntimeOptions) (_result *DownloadCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports events that occur on a cloud desktop from an Alibaba Cloud Workspace client.
//
// @param request - ExportClientEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportClientEventsResponse
func (client *Client) ExportClientEventsWithContext(ctx context.Context, request *ExportClientEventsRequest, runtime *dara.RuntimeOptions) (_result *ExportClientEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.EventTypes) {
		query["EventTypes"] = request.EventTypes
	}

	if !dara.IsNil(request.LangType) {
		query["LangType"] = request.LangType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportClientEvents"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportClientEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports cloud computer shares and saves the list as an XLSX file. Each entry includes the ID and name of the cloud computer share, the ID and name of the office network, the cloud computer share template, and the name of the security policy.
//
// @param request - ExportDesktopGroupInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDesktopGroupInfoResponse
func (client *Client) ExportDesktopGroupInfoWithContext(ctx context.Context, request *ExportDesktopGroupInfoRequest, runtime *dara.RuntimeOptions) (_result *ExportDesktopGroupInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupName) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.LangType) {
		query["LangType"] = request.LangType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportDesktopGroupInfo"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportDesktopGroupInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a cloud computer list as a CSV file.
//
// Description:
//
// The cloud computer list exported by calling this operation is saved as a CSV file. Each entry of data of a cloud computer includes the following fields:
//
//   - Cloud computer ID and name
//
//   - Office network ID and name
//
//   - The instance type, OS and protocol of the cloud computer
//
//   - System disk and data disk of the cloud computer
//
//   - The status
//
//   - Purchase method
//
//   - The time when the cloud computer expires
//
//   - Remaining duration and total duration
//
//   - Number of assigned users and number of current users
//
//   - Office network type
//
//   - The time when the cloud computer was created
//
//   - Tags
//
//   - Encryption status
//
//   - IP
//
//   - The hostname
//
// @param request - ExportDesktopListInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDesktopListInfoResponse
func (client *Client) ExportDesktopListInfoWithContext(ctx context.Context, request *ExportDesktopListInfoRequest, runtime *dara.RuntimeOptions) (_result *ExportDesktopListInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopName) {
		query["DesktopName"] = request.DesktopName
	}

	if !dara.IsNil(request.DesktopStatus) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LangType) {
		query["LangType"] = request.LangType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportDesktopListInfo"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportDesktopListInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about an asynchronous task based on the value of the AsyncTaskId parameter that you obtain by calling the CopyCdsFile operation.
//
// @param request - GetAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithContext(ctx context.Context, request *GetAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncTaskId) {
		query["AsyncTaskId"] = request.AsyncTaskId
	}

	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the credential that is used to connect to a cloud desktop.
//
// Description:
//
// The cloud computer must be in the Running state. The ticket obtained by calling this operation will expire in 10 minutes.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandContent) {
		query["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnectionTicket"),
		Version:     dara.String("2020-09-30"),
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
// # Obtains the credentials of the stream collaboration
//
// @param request - GetCoordinateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCoordinateTicketResponse
func (client *Client) GetCoordinateTicketWithContext(ctx context.Context, request *GetCoordinateTicketRequest, runtime *dara.RuntimeOptions) (_result *GetCoordinateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoId) {
		query["CoId"] = request.CoId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCoordinateTicket"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCoordinateTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a cloud computer share.
//
// @param request - GetDesktopGroupDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDesktopGroupDetailResponse
func (client *Client) GetDesktopGroupDetailWithContext(ctx context.Context, request *GetDesktopGroupDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDesktopGroupDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDesktopGroupDetail"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether single sign-on (SSO) is enabled for a workspace.
//
// @param request - GetOfficeSiteSsoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfficeSiteSsoStatusResponse
func (client *Client) GetOfficeSiteSsoStatusWithContext(ctx context.Context, request *GetOfficeSiteSsoStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOfficeSiteSsoStatus"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the metadata of a Security Assertion Markup Language (SAML) 2.0-based service provider (SP).
//
// Description:
//
// You can call this operation only for workspaces of the Active Directory (AD) and convenience account types.
//
// @param request - GetSpMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpMetadataResponse
func (client *Client) GetSpMetadataWithContext(ctx context.Context, request *GetSpMetadataRequest, runtime *dara.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpMetadata"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Hibernates cloud desktops.
//
// Description:
//
// Hibernating a cloud desktop is in private preview. If you want to try this feature, submit a ticket.
//
// @param request - HibernateDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HibernateDesktopsResponse
func (client *Client) HibernateDesktopsWithContext(ctx context.Context, request *HibernateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *HibernateDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HibernateDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HibernateDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of files in the network disk and obtain the download link of the file.
//
// @param tmpReq - ListCdsFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCdsFilesResponse
func (client *Client) ListCdsFilesWithContext(ctx context.Context, tmpReq *ListCdsFilesRequest, runtime *dara.RuntimeOptions) (_result *ListCdsFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCdsFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.ParentFileId) {
		query["ParentFileId"] = request.ParentFileId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCdsFiles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCdsFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the user information in the AD system if you use an AD directory to connect to an AD system.
//
// Description:
//
// If you use an AD directory to connect to an AD system, you can call this operation to obtain the user information in the AD system.
//
// @param request - ListDirectoryUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDirectoryUsersResponse
func (client *Client) ListDirectoryUsersWithContext(ctx context.Context, request *ListDirectoryUsersRequest, runtime *dara.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssignedInfo) {
		query["AssignedInfo"] = request.AssignedInfo
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IncludeAssignedUser) {
		query["IncludeAssignedUser"] = request.IncludeAssignedUser
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OUPath) {
		query["OUPath"] = request.OUPath
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDirectoryUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions on a shared file on a drive.
//
// @param request - ListFilePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilePermissionResponse
func (client *Client) ListFilePermissionWithContext(ctx context.Context, request *ListFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *ListFilePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFilePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries applications installed on a cloud computer.
//
// @param request - ListInstalledAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstalledAppsResponse
func (client *Client) ListInstalledAppsWithContext(ctx context.Context, request *ListInstalledAppsRequest, runtime *dara.RuntimeOptions) (_result *ListInstalledAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
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
		Action:      dara.String("ListInstalledApps"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstalledAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an office network, including its status, cloud computer quantity, virtual private cloud (VPC) type, and more.
//
// @param request - ListOfficeSiteOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOfficeSiteOverviewResponse
func (client *Client) ListOfficeSiteOverviewWithContext(ctx context.Context, request *ListOfficeSiteOverviewRequest, runtime *dara.RuntimeOptions) (_result *ListOfficeSiteOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceRefresh) {
		query["ForceRefresh"] = request.ForceRefresh
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.QueryRange) {
		query["QueryRange"] = request.QueryRange
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOfficeSiteOverview"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Active Directory (AD) accounts after an enterprise AD office network (formerly workspace) interconnects to an AD domain.
//
// @param request - ListOfficeSiteUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOfficeSiteUsersResponse
func (client *Client) ListOfficeSiteUsersWithContext(ctx context.Context, request *ListOfficeSiteUsersRequest, runtime *dara.RuntimeOptions) (_result *ListOfficeSiteUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssignedInfo) {
		query["AssignedInfo"] = request.AssignedInfo
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IncludeAssignedUser) {
		query["IncludeAssignedUser"] = request.IncludeAssignedUser
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OUPath) {
		query["OUPath"] = request.OUPath
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOfficeSiteUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of cloud computers.
//
// Description:
//
// You must use at least one of the following parameters in the request to determine the object that you want to query: `ResourceId.N`, `Tag.N.Key`, and `Tag.N.Value`.
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
		Version:     dara.String("2020-09-30"),
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
// 获取文件列表
//
// @param request - ListTransferFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransferFileDownloadUrlResponse
func (client *Client) ListTransferFileDownloadUrlWithContext(ctx context.Context, request *ListTransferFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *ListTransferFileDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileIds) {
		query["FileIds"] = request.FileIds
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransferFileDownloadUrl"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransferFileDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the file information of a file transmission task.
//
// @param request - ListTransferFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransferFilesResponse
func (client *Client) ListTransferFilesWithContext(ctx context.Context, request *ListTransferFilesRequest, runtime *dara.RuntimeOptions) (_result *ListTransferFilesResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransferFiles"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransferFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the organizational units (OUs) of an Active Directory (AD) domain that is connected to an enterprise AD office network (formerly workspace).
//
// @param request - ListUserAdOrganizationUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAdOrganizationUnitsResponse
func (client *Client) ListUserAdOrganizationUnitsWithContext(ctx context.Context, request *ListUserAdOrganizationUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListUserAdOrganizationUnitsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserAdOrganizationUnits"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAdOrganizationUnitsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks a multi-factor authentication (MFA) device that is in the NORMAL state.
//
// Description:
//
// After a virtual MFA device is locked, its status changes to LOCKED. The Active Directory (AD) user who uses the virtual MFA device is unable to pass MFA and is therefore unable to log on to the client. You can call the [UnlockVirtualMFADevice](https://help.aliyun.com/document_detail/206212.html) operation to unlock the device.
//
// @param request - LockVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockVirtualMFADeviceResponse
func (client *Client) LockVirtualMFADeviceWithContext(ctx context.Context, request *LockVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockVirtualMFADevice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates cloud computers from the current office network (formerly called workspace) to the new office network.
//
// @param request - MigrateDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateDesktopsResponse
func (client *Client) MigrateDesktopsWithContext(ctx context.Context, request *MigrateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *MigrateDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetOfficeSiteId) {
		query["TargetOfficeSiteId"] = request.TargetOfficeSiteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the protocols of images to Adaptive Streaming Protocol (ASP).
//
// @param request - MigrateImageProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateImageProtocolResponse
func (client *Client) MigrateImageProtocolWithContext(ctx context.Context, request *MigrateImageProtocolRequest, runtime *dara.RuntimeOptions) (_result *MigrateImageProtocolResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetProtocolType) {
		query["TargetProtocolType"] = request.TargetProtocolType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateImageProtocol"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateImageProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Active Directory (AD) directory.
//
// Description:
//
// You can modify the following domain name- and Domain Name System (DNS)-related parameters only for Active Directory (AD) directories that are in the ERROR or REGISTERING state: `DomainName`, `SubDomainName`, `DnsAddress.N`, and `SubDomainDnsAddress`.
//
// @param request - ModifyADConnectorDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyADConnectorDirectoryResponse
func (client *Client) ModifyADConnectorDirectoryWithContext(ctx context.Context, request *ModifyADConnectorDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdHostname) {
		query["AdHostname"] = request.AdHostname
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DirectoryName) {
		query["DirectoryName"] = request.DirectoryName
	}

	if !dara.IsNil(request.DnsAddress) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainPassword) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !dara.IsNil(request.DomainUserName) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !dara.IsNil(request.MfaEnabled) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !dara.IsNil(request.OUName) {
		query["OUName"] = request.OUName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubDomainDnsAddress) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !dara.IsNil(request.SubDomainName) {
		query["SubDomainName"] = request.SubDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyADConnectorDirectory"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic properties of an enterprise Active Directory (AD) office network, such as the office network name and domain names of the enterprise AD subdomains.
//
// Description:
//
// You can modify parameters of domain names and Domain Name System (DNS) for enterprise AD office networks that are in the `ERROR` or `REGISTERED` state. The parameters include `DomainName`, `SubDomainName`, `DnsAddress.N`, and `SubDomainDnsAddress.N`.
//
// @param request - ModifyADConnectorOfficeSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyADConnectorOfficeSiteResponse
func (client *Client) ModifyADConnectorOfficeSiteWithContext(ctx context.Context, request *ModifyADConnectorOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdHostname) {
		query["AdHostname"] = request.AdHostname
	}

	if !dara.IsNil(request.BackupDCHostname) {
		query["BackupDCHostname"] = request.BackupDCHostname
	}

	if !dara.IsNil(request.BackupDns) {
		query["BackupDns"] = request.BackupDns
	}

	if !dara.IsNil(request.DnsAddress) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainPassword) {
		query["DomainPassword"] = request.DomainPassword
	}

	if !dara.IsNil(request.DomainUserName) {
		query["DomainUserName"] = request.DomainUserName
	}

	if !dara.IsNil(request.MfaEnabled) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !dara.IsNil(request.OUName) {
		query["OUName"] = request.OUName
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubDomainDnsAddress) {
		query["SubDomainDnsAddress"] = request.SubDomainDnsAddress
	}

	if !dara.IsNil(request.SubDomainName) {
		query["SubDomainName"] = request.SubDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyADConnectorOfficeSite"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the Internet access control policy on the office network or cloud computer granularity.
//
// Description:
//
// You can set different Internet access control policies at different granularities to achieve the effect of composite policies. For example, you can disable the Internet access on the office network granularity and enable the Internet access on specific cloud computer granularity. The effect is that all cloud computers in the office network except the specified cloud computers are not allowed to access the Internet.
//
// @param request - ModifyAclEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAclEntriesResponse
func (client *Client) ModifyAclEntriesWithContext(ctx context.Context, request *ModifyAclEntriesRequest, runtime *dara.RuntimeOptions) (_result *ModifyAclEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAclEntries"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAclEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of an automatic snapshot policy, such as the policy name and snapshot retention period.
//
// @param request - ModifyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicyWithContext(ctx context.Context, request *ModifyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoSnapshotPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom cloud computer template.
//
// Description:
//
// Only custom desktop templates can be modified.
//
// @param request - ModifyBundleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBundleResponse
func (client *Client) ModifyBundleWithContext(ctx context.Context, request *ModifyBundleRequest, runtime *dara.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.BundleName) {
		query["BundleName"] = request.BundleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBundle"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBundleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a disk file or folder, such as the file name.
//
// @param request - ModifyCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdsFileResponse
func (client *Client) ModifyCdsFileWithContext(ctx context.Context, request *ModifyCdsFileRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.ConflictPolicy) {
		query["ConflictPolicy"] = request.ConflictPolicy
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the link for file sharing.
//
// @param request - ModifyCdsFileShareLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdsFileShareLinkResponse
func (client *Client) ModifyCdsFileShareLinkWithContext(ctx context.Context, request *ModifyCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdsFileShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableDownload) {
		query["DisableDownload"] = request.DisableDownload
	}

	if !dara.IsNil(request.DisablePreview) {
		query["DisablePreview"] = request.DisablePreview
	}

	if !dara.IsNil(request.DisableSave) {
		query["DisableSave"] = request.DisableSave
	}

	if !dara.IsNil(request.DownloadCount) {
		query["DownloadCount"] = request.DownloadCount
	}

	if !dara.IsNil(request.DownloadLimit) {
		query["DownloadLimit"] = request.DownloadLimit
	}

	if !dara.IsNil(request.Expiration) {
		query["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.PreviewCount) {
		query["PreviewCount"] = request.PreviewCount
	}

	if !dara.IsNil(request.PreviewLimit) {
		query["PreviewLimit"] = request.PreviewLimit
	}

	if !dara.IsNil(request.ReportCount) {
		query["ReportCount"] = request.ReportCount
	}

	if !dara.IsNil(request.SaveCount) {
		query["SaveCount"] = request.SaveCount
	}

	if !dara.IsNil(request.SaveLimit) {
		query["SaveLimit"] = request.SaveLimit
	}

	if !dara.IsNil(request.ShareId) {
		query["ShareId"] = request.ShareId
	}

	if !dara.IsNil(request.ShareName) {
		query["ShareName"] = request.ShareName
	}

	if !dara.IsNil(request.SharePwd) {
		query["SharePwd"] = request.SharePwd
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VideoPreviewCount) {
		query["VideoPreviewCount"] = request.VideoPreviewCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdsFileShareLink"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdsFileShareLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a center policy.
//
// @param request - ModifyCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenterPolicyResponse
func (client *Client) ModifyCenterPolicyWithContext(ctx context.Context, request *ModifyCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenterPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcademicProxy) {
		query["AcademicProxy"] = request.AcademicProxy
	}

	if !dara.IsNil(request.AdminAccess) {
		query["AdminAccess"] = request.AdminAccess
	}

	if !dara.IsNil(request.AppContentProtection) {
		query["AppContentProtection"] = request.AppContentProtection
	}

	if !dara.IsNil(request.AuthorizeAccessPolicyRule) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !dara.IsNil(request.AuthorizeSecurityPolicyRule) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !dara.IsNil(request.AutoReconnect) {
		query["AutoReconnect"] = request.AutoReconnect
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.CameraRedirect) {
		query["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.ClientControlMenu) {
		query["ClientControlMenu"] = request.ClientControlMenu
	}

	if !dara.IsNil(request.ClientCreateSnapshot) {
		query["ClientCreateSnapshot"] = request.ClientCreateSnapshot
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Clipboard) {
		query["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.ClipboardGraineds) {
		query["ClipboardGraineds"] = request.ClipboardGraineds
	}

	if !dara.IsNil(request.ClipboardScope) {
		query["ClipboardScope"] = request.ClipboardScope
	}

	if !dara.IsNil(request.ColorEnhancement) {
		query["ColorEnhancement"] = request.ColorEnhancement
	}

	if !dara.IsNil(request.CpdDriveClipboard) {
		query["CpdDriveClipboard"] = request.CpdDriveClipboard
	}

	if !dara.IsNil(request.CpuDownGradeDuration) {
		query["CpuDownGradeDuration"] = request.CpuDownGradeDuration
	}

	if !dara.IsNil(request.CpuOverload) {
		query["CpuOverload"] = request.CpuOverload
	}

	if !dara.IsNil(request.CpuProcessors) {
		query["CpuProcessors"] = request.CpuProcessors
	}

	if !dara.IsNil(request.CpuProtectedMode) {
		query["CpuProtectedMode"] = request.CpuProtectedMode
	}

	if !dara.IsNil(request.CpuRateLimit) {
		query["CpuRateLimit"] = request.CpuRateLimit
	}

	if !dara.IsNil(request.CpuSampleDuration) {
		query["CpuSampleDuration"] = request.CpuSampleDuration
	}

	if !dara.IsNil(request.CpuSingleRateLimit) {
		query["CpuSingleRateLimit"] = request.CpuSingleRateLimit
	}

	if !dara.IsNil(request.DeviceConnectHint) {
		query["DeviceConnectHint"] = request.DeviceConnectHint
	}

	if !dara.IsNil(request.DeviceRedirects) {
		query["DeviceRedirects"] = request.DeviceRedirects
	}

	if !dara.IsNil(request.DeviceRules) {
		query["DeviceRules"] = request.DeviceRules
	}

	if !dara.IsNil(request.DisconnectKeepSession) {
		query["DisconnectKeepSession"] = request.DisconnectKeepSession
	}

	if !dara.IsNil(request.DisconnectKeepSessionTime) {
		query["DisconnectKeepSessionTime"] = request.DisconnectKeepSessionTime
	}

	if !dara.IsNil(request.DiskOverload) {
		query["DiskOverload"] = request.DiskOverload
	}

	if !dara.IsNil(request.DisplayMode) {
		query["DisplayMode"] = request.DisplayMode
	}

	if !dara.IsNil(request.DomainResolveRule) {
		query["DomainResolveRule"] = request.DomainResolveRule
	}

	if !dara.IsNil(request.DomainResolveRuleType) {
		query["DomainResolveRuleType"] = request.DomainResolveRuleType
	}

	if !dara.IsNil(request.EnableSessionRateLimiting) {
		query["EnableSessionRateLimiting"] = request.EnableSessionRateLimiting
	}

	if !dara.IsNil(request.EndUserApplyAdminCoordinate) {
		query["EndUserApplyAdminCoordinate"] = request.EndUserApplyAdminCoordinate
	}

	if !dara.IsNil(request.EndUserGroupCoordinate) {
		query["EndUserGroupCoordinate"] = request.EndUserGroupCoordinate
	}

	if !dara.IsNil(request.ExternalDrive) {
		query["ExternalDrive"] = request.ExternalDrive
	}

	if !dara.IsNil(request.FileMigrate) {
		query["FileMigrate"] = request.FileMigrate
	}

	if !dara.IsNil(request.FileTransferAddress) {
		query["FileTransferAddress"] = request.FileTransferAddress
	}

	if !dara.IsNil(request.FileTransferSpeed) {
		query["FileTransferSpeed"] = request.FileTransferSpeed
	}

	if !dara.IsNil(request.FileTransferSpeedLocation) {
		query["FileTransferSpeedLocation"] = request.FileTransferSpeedLocation
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.HoverConfigMsg) {
		query["HoverConfigMsg"] = request.HoverConfigMsg
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.InternetCommunicationProtocol) {
		query["InternetCommunicationProtocol"] = request.InternetCommunicationProtocol
	}

	if !dara.IsNil(request.InternetPrinter) {
		query["InternetPrinter"] = request.InternetPrinter
	}

	if !dara.IsNil(request.LocalDrive) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.MaxReconnectTime) {
		query["MaxReconnectTime"] = request.MaxReconnectTime
	}

	if !dara.IsNil(request.MemoryDownGradeDuration) {
		query["MemoryDownGradeDuration"] = request.MemoryDownGradeDuration
	}

	if !dara.IsNil(request.MemoryOverload) {
		query["MemoryOverload"] = request.MemoryOverload
	}

	if !dara.IsNil(request.MemoryProcessors) {
		query["MemoryProcessors"] = request.MemoryProcessors
	}

	if !dara.IsNil(request.MemoryProtectedMode) {
		query["MemoryProtectedMode"] = request.MemoryProtectedMode
	}

	if !dara.IsNil(request.MemoryRateLimit) {
		query["MemoryRateLimit"] = request.MemoryRateLimit
	}

	if !dara.IsNil(request.MemorySampleDuration) {
		query["MemorySampleDuration"] = request.MemorySampleDuration
	}

	if !dara.IsNil(request.MemorySingleRateLimit) {
		query["MemorySingleRateLimit"] = request.MemorySingleRateLimit
	}

	if !dara.IsNil(request.MobileRestart) {
		query["MobileRestart"] = request.MobileRestart
	}

	if !dara.IsNil(request.MobileSafeMenu) {
		query["MobileSafeMenu"] = request.MobileSafeMenu
	}

	if !dara.IsNil(request.MobileShutdown) {
		query["MobileShutdown"] = request.MobileShutdown
	}

	if !dara.IsNil(request.MobileWuyingKeeper) {
		query["MobileWuyingKeeper"] = request.MobileWuyingKeeper
	}

	if !dara.IsNil(request.MobileWyAssistant) {
		query["MobileWyAssistant"] = request.MobileWyAssistant
	}

	if !dara.IsNil(request.ModelLibrary) {
		query["ModelLibrary"] = request.ModelLibrary
	}

	if !dara.IsNil(request.MultiScreen) {
		query["MultiScreen"] = request.MultiScreen
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetRedirect) {
		query["NetRedirect"] = request.NetRedirect
	}

	if !dara.IsNil(request.NetRedirectRule) {
		query["NetRedirectRule"] = request.NetRedirectRule
	}

	if !dara.IsNil(request.NoOperationDisconnect) {
		query["NoOperationDisconnect"] = request.NoOperationDisconnect
	}

	if !dara.IsNil(request.NoOperationDisconnectTime) {
		query["NoOperationDisconnectTime"] = request.NoOperationDisconnectTime
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PortProxy) {
		query["PortProxy"] = request.PortProxy
	}

	if !dara.IsNil(request.PrinterRedirect) {
		query["PrinterRedirect"] = request.PrinterRedirect
	}

	if !dara.IsNil(request.QualityEnhancement) {
		query["QualityEnhancement"] = request.QualityEnhancement
	}

	if !dara.IsNil(request.RecordEventDuration) {
		query["RecordEventDuration"] = request.RecordEventDuration
	}

	if !dara.IsNil(request.RecordEventFileExts) {
		query["RecordEventFileExts"] = request.RecordEventFileExts
	}

	if !dara.IsNil(request.RecordEventFilePaths) {
		query["RecordEventFilePaths"] = request.RecordEventFilePaths
	}

	if !dara.IsNil(request.RecordEventLevels) {
		query["RecordEventLevels"] = request.RecordEventLevels
	}

	if !dara.IsNil(request.RecordEventRegisters) {
		query["RecordEventRegisters"] = request.RecordEventRegisters
	}

	if !dara.IsNil(request.RecordEvents) {
		query["RecordEvents"] = request.RecordEvents
	}

	if !dara.IsNil(request.Recording) {
		query["Recording"] = request.Recording
	}

	if !dara.IsNil(request.RecordingAudio) {
		query["RecordingAudio"] = request.RecordingAudio
	}

	if !dara.IsNil(request.RecordingDuration) {
		query["RecordingDuration"] = request.RecordingDuration
	}

	if !dara.IsNil(request.RecordingEndTime) {
		query["RecordingEndTime"] = request.RecordingEndTime
	}

	if !dara.IsNil(request.RecordingExpires) {
		query["RecordingExpires"] = request.RecordingExpires
	}

	if !dara.IsNil(request.RecordingFps) {
		query["RecordingFps"] = request.RecordingFps
	}

	if !dara.IsNil(request.RecordingStartTime) {
		query["RecordingStartTime"] = request.RecordingStartTime
	}

	if !dara.IsNil(request.RecordingUserNotify) {
		query["RecordingUserNotify"] = request.RecordingUserNotify
	}

	if !dara.IsNil(request.RecordingUserNotifyMessage) {
		query["RecordingUserNotifyMessage"] = request.RecordingUserNotifyMessage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteCoordinate) {
		query["RemoteCoordinate"] = request.RemoteCoordinate
	}

	if !dara.IsNil(request.ResetDesktop) {
		query["ResetDesktop"] = request.ResetDesktop
	}

	if !dara.IsNil(request.ResolutionDpi) {
		query["ResolutionDpi"] = request.ResolutionDpi
	}

	if !dara.IsNil(request.ResolutionHeight) {
		query["ResolutionHeight"] = request.ResolutionHeight
	}

	if !dara.IsNil(request.ResolutionModel) {
		query["ResolutionModel"] = request.ResolutionModel
	}

	if !dara.IsNil(request.ResolutionWidth) {
		query["ResolutionWidth"] = request.ResolutionWidth
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RevokeAccessPolicyRule) {
		query["RevokeAccessPolicyRule"] = request.RevokeAccessPolicyRule
	}

	if !dara.IsNil(request.RevokeSecurityPolicyRule) {
		query["RevokeSecurityPolicyRule"] = request.RevokeSecurityPolicyRule
	}

	if !dara.IsNil(request.SafeMenu) {
		query["SafeMenu"] = request.SafeMenu
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	if !dara.IsNil(request.ScreenDisplayMode) {
		query["ScreenDisplayMode"] = request.ScreenDisplayMode
	}

	if !dara.IsNil(request.SessionMaxRateKbps) {
		query["SessionMaxRateKbps"] = request.SessionMaxRateKbps
	}

	if !dara.IsNil(request.SmoothEnhancement) {
		query["SmoothEnhancement"] = request.SmoothEnhancement
	}

	if !dara.IsNil(request.StatusMonitor) {
		query["StatusMonitor"] = request.StatusMonitor
	}

	if !dara.IsNil(request.StreamingMode) {
		query["StreamingMode"] = request.StreamingMode
	}

	if !dara.IsNil(request.TargetFps) {
		query["TargetFps"] = request.TargetFps
	}

	if !dara.IsNil(request.Taskbar) {
		query["Taskbar"] = request.Taskbar
	}

	if !dara.IsNil(request.UsbRedirect) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !dara.IsNil(request.UsbSupplyRedirectRule) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !dara.IsNil(request.UseTime) {
		query["UseTime"] = request.UseTime
	}

	if !dara.IsNil(request.VideoEncAvgKbps) {
		query["VideoEncAvgKbps"] = request.VideoEncAvgKbps
	}

	if !dara.IsNil(request.VideoEncMaxQP) {
		query["VideoEncMaxQP"] = request.VideoEncMaxQP
	}

	if !dara.IsNil(request.VideoEncMinQP) {
		query["VideoEncMinQP"] = request.VideoEncMinQP
	}

	if !dara.IsNil(request.VideoEncPeakKbps) {
		query["VideoEncPeakKbps"] = request.VideoEncPeakKbps
	}

	if !dara.IsNil(request.VideoEncPolicy) {
		query["VideoEncPolicy"] = request.VideoEncPolicy
	}

	if !dara.IsNil(request.VideoRedirect) {
		query["VideoRedirect"] = request.VideoRedirect
	}

	if !dara.IsNil(request.VisualQuality) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	if !dara.IsNil(request.WatermarkAntiCam) {
		query["WatermarkAntiCam"] = request.WatermarkAntiCam
	}

	if !dara.IsNil(request.WatermarkColor) {
		query["WatermarkColor"] = request.WatermarkColor
	}

	if !dara.IsNil(request.WatermarkColumnAmount) {
		query["WatermarkColumnAmount"] = request.WatermarkColumnAmount
	}

	if !dara.IsNil(request.WatermarkCustomText) {
		query["WatermarkCustomText"] = request.WatermarkCustomText
	}

	if !dara.IsNil(request.WatermarkDegree) {
		query["WatermarkDegree"] = request.WatermarkDegree
	}

	if !dara.IsNil(request.WatermarkFontSize) {
		query["WatermarkFontSize"] = request.WatermarkFontSize
	}

	if !dara.IsNil(request.WatermarkFontStyle) {
		query["WatermarkFontStyle"] = request.WatermarkFontStyle
	}

	if !dara.IsNil(request.WatermarkPower) {
		query["WatermarkPower"] = request.WatermarkPower
	}

	if !dara.IsNil(request.WatermarkRowAmount) {
		query["WatermarkRowAmount"] = request.WatermarkRowAmount
	}

	if !dara.IsNil(request.WatermarkSecurity) {
		query["WatermarkSecurity"] = request.WatermarkSecurity
	}

	if !dara.IsNil(request.WatermarkShadow) {
		query["WatermarkShadow"] = request.WatermarkShadow
	}

	if !dara.IsNil(request.WatermarkTransparencyValue) {
		query["WatermarkTransparencyValue"] = request.WatermarkTransparencyValue
	}

	if !dara.IsNil(request.WatermarkType) {
		query["WatermarkType"] = request.WatermarkType
	}

	if !dara.IsNil(request.WuyingKeeper) {
		query["WuyingKeeper"] = request.WuyingKeeper
	}

	if !dara.IsNil(request.WyAssistant) {
		query["WyAssistant"] = request.WyAssistant
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCenterPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCenterPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies team spaces.
//
// @param request - ModifyCloudDriveGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudDriveGroupsResponse
func (client *Client) ModifyCloudDriveGroupsWithContext(ctx context.Context, request *ModifyCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDriveGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TotalSize) {
		query["TotalSize"] = request.TotalSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudDriveGroups"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudDriveGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the user permissions on Cloud Drive Service, and configures users who have the download permissions and upload and download permissions. By default, the users that are not configured the preceding permissions only have the upload permissions.
//
// @param request - ModifyCloudDrivePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudDrivePermissionResponse
func (client *Client) ModifyCloudDrivePermissionWithContext(ctx context.Context, request *ModifyCloudDrivePermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDrivePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.DownloadEndUserIds) {
		query["DownloadEndUserIds"] = request.DownloadEndUserIds
	}

	if !dara.IsNil(request.DownloadUploadEndUserIds) {
		query["DownloadUploadEndUserIds"] = request.DownloadUploadEndUserIds
	}

	if !dara.IsNil(request.NoDownloadNoUploadEndUserIds) {
		query["NoDownloadNoUploadEndUserIds"] = request.NoDownloadNoUploadEndUserIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudDrivePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudDrivePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改无影网盘终端用户的属性
//
// @param request - ModifyCloudDriveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudDriveUsersResponse
func (client *Client) ModifyCloudDriveUsersWithContext(ctx context.Context, request *ModifyCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDriveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserMaxSize) {
		query["UserMaxSize"] = request.UserMaxSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudDriveUsers"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudDriveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a configuration group.
//
// @param request - ModifyConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyConfigGroupResponse
func (client *Client) ModifyConfigGroupWithContext(ctx context.Context, request *ModifyConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyConfigGroupResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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
		Action:      dara.String("ModifyConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the layouts of cloud computer list headers, such as the required fields and the display and hide settings.
//
// @param request - ModifyCustomizedListHeadersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomizedListHeadersResponse
func (client *Client) ModifyCustomizedListHeadersWithContext(ctx context.Context, request *ModifyCustomizedListHeadersRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomizedListHeadersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Headers) {
		query["Headers"] = request.Headers
	}

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomizedListHeaders"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomizedListHeadersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of cloud computers to subscription or pay-as-you-go.
//
// Description:
//
//	  Before you call this operation, make sure that you fully understand the billing methods of cloud computers. For more information, see [Billing overview](https://help.aliyun.com/document_detail/188395.html).
//
//		- Before you call this operation, make sure that the cloud computers whose billing method you want to change are in the Running or Stopped state and you have no overdue payments in your Alibaba Cloud account.
//
//		- After the order payment is completed, the system starts to change the billing method of the cloud computers. During the change, you cannot perform operations, such as starting or stopping the cloud computers, and changing configurations of the cloud computers.
//
// @param request - ModifyDesktopChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopChargeTypeResponse
func (client *Client) ModifyDesktopChargeTypeWithContext(ctx context.Context, request *ModifyDesktopChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.UseDuration) {
		query["UseDuration"] = request.UseDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopChargeType"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a cloud computer share.
//
// Description:
//
// Once a cloud computer share is created, the system automatically provisions cloud computers according to the auto-scaling policy and user connections, all based on the same template and security policy. You can adjust the cloud computer share\\"s configurations, including the share name, template, and policy, for different business scenarios.
//
// @param request - ModifyDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopGroupResponse
func (client *Client) ModifyDesktopGroupWithContext(ctx context.Context, request *ModifyDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowAutoSetup) {
		query["AllowAutoSetup"] = request.AllowAutoSetup
	}

	if !dara.IsNil(request.AllowBufferCount) {
		query["AllowBufferCount"] = request.AllowBufferCount
	}

	if !dara.IsNil(request.BindAmount) {
		query["BindAmount"] = request.BindAmount
	}

	if !dara.IsNil(request.BuyDesktopsCount) {
		query["BuyDesktopsCount"] = request.BuyDesktopsCount
	}

	if !dara.IsNil(request.Classify) {
		query["Classify"] = request.Classify
	}

	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.ConnectDuration) {
		query["ConnectDuration"] = request.ConnectDuration
	}

	if !dara.IsNil(request.DeleteDuration) {
		query["DeleteDuration"] = request.DeleteDuration
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupName) {
		query["DesktopGroupName"] = request.DesktopGroupName
	}

	if !dara.IsNil(request.DisableSessionConfig) {
		query["DisableSessionConfig"] = request.DisableSessionConfig
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.IdleDisconnectDuration) {
		query["IdleDisconnectDuration"] = request.IdleDisconnectDuration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.KeepDuration) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !dara.IsNil(request.LoadPolicy) {
		query["LoadPolicy"] = request.LoadPolicy
	}

	if !dara.IsNil(request.MaxDesktopsCount) {
		query["MaxDesktopsCount"] = request.MaxDesktopsCount
	}

	if !dara.IsNil(request.MinDesktopsCount) {
		query["MinDesktopsCount"] = request.MinDesktopsCount
	}

	if !dara.IsNil(request.OwnBundleId) {
		query["OwnBundleId"] = request.OwnBundleId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PolicyGroupIds) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
	}

	if !dara.IsNil(request.ProfileFollowSwitch) {
		query["ProfileFollowSwitch"] = request.ProfileFollowSwitch
	}

	if !dara.IsNil(request.RatioThreshold) {
		query["RatioThreshold"] = request.RatioThreshold
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetType) {
		query["ResetType"] = request.ResetType
	}

	if !dara.IsNil(request.ScaleStrategyId) {
		query["ScaleStrategyId"] = request.ScaleStrategyId
	}

	if !dara.IsNil(request.StopDuration) {
		query["StopDuration"] = request.StopDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the hostname of a Windows cloud computer in the Active Directory (AD) office network.
//
// Description:
//
// The Windows cloud computer whose hostname you want to modify must be in an AD office network. After the hostname is modified, the cloud computer is re-created.
//
// @param request - ModifyDesktopHostNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopHostNameResponse
func (client *Client) ModifyDesktopHostNameWithContext(ctx context.Context, request *ModifyDesktopHostNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopHostNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.NewHostName) {
		query["NewHostName"] = request.NewHostName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopHostName"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopHostNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a cloud computer to a new name.
//
// @param request - ModifyDesktopNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopNameResponse
func (client *Client) ModifyDesktopNameWithContext(ctx context.Context, request *ModifyDesktopNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.NewDesktopName) {
		query["NewDesktopName"] = request.NewDesktopName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopName"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改桌面超卖组
//
// @param request - ModifyDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopOversoldGroupResponse
func (client *Client) ModifyDesktopOversoldGroupWithContext(ctx context.Context, request *ModifyDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrenceCount) {
		query["ConcurrenceCount"] = request.ConcurrenceCount
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IdleDisconnectDuration) {
		query["IdleDisconnectDuration"] = request.IdleDisconnectDuration
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.KeepDuration) {
		query["KeepDuration"] = request.KeepDuration
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.OversoldUserCount) {
		query["OversoldUserCount"] = request.OversoldUserCount
	}

	if !dara.IsNil(request.OversoldWarn) {
		query["OversoldWarn"] = request.OversoldWarn
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.StopDuration) {
		query["StopDuration"] = request.StopDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改桌面超卖组售卖数据
//
// @param request - ModifyDesktopOversoldGroupSaleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopOversoldGroupSaleResponse
func (client *Client) ModifyDesktopOversoldGroupSaleWithContext(ctx context.Context, request *ModifyDesktopOversoldGroupSaleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldGroupSaleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrenceCount) {
		query["ConcurrenceCount"] = request.ConcurrenceCount
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.OversoldUserCount) {
		query["OversoldUserCount"] = request.OversoldUserCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopOversoldGroupSale"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopOversoldGroupSaleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改桌面超卖用户组
//
// @param request - ModifyDesktopOversoldUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopOversoldUserGroupResponse
func (client *Client) ModifyDesktopOversoldUserGroupWithContext(ctx context.Context, request *ModifyDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopOversoldUserGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopOversoldUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the instance type of a cloud computer and scales up the disks of the cloud computer.
//
// Description:
//
// Changing the configurations of a cloud computer includes changing the instance type of the cloud computer and scaling up the disks of the cloud computer.
//
//   - Before you change the configurations of a cloud computer, you must understand the instance types and disk sizes supported by cloud computers. For more information, see [Cloud computer types](https://help.aliyun.com/document_detail/188609.html). You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the instance types supported by cloud computers.
//
//   - You must change at least one of the following configurations: instance type, system disk size, and data disk size of the cloud computer. You must specify at least one of the following parameters: `DesktopType`, `RootDiskSizeGib`, and `UserDiskSizeGib`. Take note of the following items:
//
//   - The instance type of a cloud computer includes the configurations of vCPUs, memory, and GPUs. You can only change an instance type to another. You cannot change only one of the configurations.
//
//   - You cannot change a cloud computer between the General Office type and the non-General Office type. You cannot yet change a cloud computer between the Graphics type and the non-Graphics type.
//
//   - The system disk and data disks of a cloud computer can only be scaled up and cannot be scaled down.
//
//   - If the billing method of the cloud computer is subscription, the system calculates the price difference based on the configuration difference between the original cloud computer and the new cloud computer. You must make up for the price difference or receive a refund for the price difference.
//
//   - We recommend that you do not change the configurations of a cloud computer twice within 5 minutes.
//
//   - When you change the configurations of a cloud computer, the cloud computer must be in the Stopped state.
//
//   - After you change the configurations of a cloud computer, the personal data on the cloud computer is not affected.
//
// @param request - ModifyDesktopSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopSpecResponse
func (client *Client) ModifyDesktopSpecWithContext(ctx context.Context, request *ModifyDesktopSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceSpecs) {
		query["ResourceSpecs"] = request.ResourceSpecs
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RootDiskSizeGib) {
		query["RootDiskSizeGib"] = request.RootDiskSizeGib
	}

	if !dara.IsNil(request.UserDiskPerformanceLevel) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	if !dara.IsNil(request.UserDiskSizeGib) {
		query["UserDiskSizeGib"] = request.UserDiskSizeGib
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopSpec"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies scheduled tasks on cloud computers, such as starting, stopping, restarting, and resetting cloud computers on schedule.
//
// @param request - ModifyDesktopTimerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopTimerResponse
func (client *Client) ModifyDesktopTimerWithContext(ctx context.Context, request *ModifyDesktopTimerRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopTimerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopTimers) {
		query["DesktopTimers"] = request.DesktopTimers
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UseDesktopTimers) {
		query["UseDesktopTimers"] = request.UseDesktopTimers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopTimer"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopTimerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes an existing cloud computer policy for cloud computers.
//
// Description:
//
// The cloud computers for which you want to change their policies must be in the Running state.
//
// @param request - ModifyDesktopsPolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesktopsPolicyGroupResponse
func (client *Client) ModifyDesktopsPolicyGroupWithContext(ctx context.Context, request *ModifyDesktopsPolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PolicyGroupIds) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesktopsPolicyGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the performance level (PL) of a system disk or data disk.
//
// Description:
//
// When creating a cloud computer in Elastic Desktop Service (EDS) Enterprise, you can use a template to define specifications that align with your business needs. By default, Enterprise Graphics or High Frequency cloud computers utilize Enterprise SSDs (ESSDs). You can customize the disk capacity and performance level (PL) of these ESSDs, and adjust the PL for both system and data disks as needed.
//
// >  Only Enterprise Graphics or High Frequency cloud computers support disk PL adjustments.
//
// @param request - ModifyDiskSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskSpecResponse
func (client *Client) ModifyDiskSpecWithContext(ctx context.Context, request *ModifyDiskSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.RootDiskPerformanceLevel) {
		query["RootDiskPerformanceLevel"] = request.RootDiskPerformanceLevel
	}

	if !dara.IsNil(request.UserDiskPerformanceLevel) {
		query["UserDiskPerformanceLevel"] = request.UserDiskPerformanceLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskSpec"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a cloud computer to end users and removes all original end users of the cloud computer.
//
// Description:
//
//	  The cloud computer must be in the Running state.
//
//		- After you call this operation, the assignment result is immediately returned. You can call the [DescribeDesktops](https://help.aliyun.com/document_detail/436815.html) operation to query the assignment of the cloud computer. The value of the `ManagementFlags` response parameter indicates the assignment of the cloud computer. A value of `ASSIGNING` indicates that the cloud computer is being assigned, and other values indicate that the cloud computer is assigned.
//
//		- We recommend that you check the assignment every 2 to 5 seconds and perform the checks within 50 seconds. Typically, 1 to 5 seconds are required to complete the assignment.
//
// @param request - ModifyEntitlementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEntitlementResponse
func (client *Client) ModifyEntitlementWithContext(ctx context.Context, request *ModifyEntitlementRequest, runtime *dara.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEntitlement"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an image, including the name and description of the image.
//
// Description:
//
// You can call this operation to modify the attributes of only custom images that are in the Available state.
//
// @param request - ModifyImageAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImageAttributeResponse
func (client *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
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
		Action:      dara.String("ModifyImageAttribute"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares an image with other Alibaba Cloud accounts, or unshares an image from the recipient Alibaba Cloud accounts.
//
// Description:
//
// ### [](#)Security of shared images
//
// Elastic Desktop Service cannot guarantee the integrity and security of shared images. When you use a shared image, you must make sure that the image comes from a trusted sharer or account, and you are legally responsible for using the shared image.
//
// ### [](#)Quota and billing
//
//   - A shared image does not count against the image quotas of principals to which the image is shared.
//
//   - After a principal uses a shared image to create a cloud computer, the sharer is not charged for the shared image.
//
//   - You are not charged for shared images.
//
// ### [](#)Supported sharing behaviors
//
//   - You can share custom images with other Alibaba Cloud accounts.
//
//   - You can share custom images between accounts in the China site (aliyun.com) and the international site (alibabacloud.com).
//
// ### [](#)Unsupported sharing behaviors
//
//   - You cannot share images that are shared by other Alibaba Cloud accounts.
//
//   - You cannot share encrypted images.
//
//   - You cannot share images across regions. If you want to share an image across regions, you must copy the image to the destination region and then share the image. For more information, see [CopyImage](https://help.aliyun.com/document_detail/436978.html).
//
// @param request - ModifyImagePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImagePermissionResponse
func (client *Client) ModifyImagePermissionWithContext(ctx context.Context, request *ModifyImagePermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyImagePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddAccount) {
		query["AddAccount"] = request.AddAccount
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveAccount) {
		query["RemoveAccount"] = request.RemoveAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyImagePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyImagePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the mount target of a File Storage NAS (NAS) file system.
//
// Description:
//
// When you create a NAS file system, a mount target is automatically generated. By default, the mount target does not need to be changed. If the mount target is deleted by misoperation, you must specify a new mount target for the NAS file system in the workspace. You can call the [CreateMountTarget](https://help.aliyun.com/document_detail/62621.html) operation to create a mount target.
//
// @param request - ModifyNASDefaultMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNASDefaultMountTargetResponse
func (client *Client) ModifyNASDefaultMountTargetWithContext(ctx context.Context, request *ModifyNASDefaultMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNASDefaultMountTarget"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the bandwidth of a premium bandwidth plan.
//
// @param request - ModifyNetworkPackageBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkPackageBandwidthResponse
func (client *Client) ModifyNetworkPackageBandwidthWithContext(ctx context.Context, request *ModifyNetworkPackageBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkPackageBandwidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkPackageBandwidth"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkPackageBandwidthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores or disables a premium bandwidth plan.
//
// Description:
//
// If you want to temporarily disable the Internet access of your cloud computer after the Internet access is enabled for your cloud computer, you can disable the premium bandwidth plan and restore it as needed.
//
// @param request - ModifyNetworkPackageEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkPackageEnabledResponse
func (client *Client) ModifyNetworkPackageEnabledWithContext(ctx context.Context, request *ModifyNetworkPackageEnabledRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkPackageEnabled"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic properties of an office network, including the name and local administrator permission settings.
//
// @param request - ModifyOfficeSiteAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfficeSiteAttributeResponse
func (client *Client) ModifyOfficeSiteAttributeWithContext(ctx context.Context, request *ModifyOfficeSiteAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorityHost) {
		query["AuthorityHost"] = request.AuthorityHost
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["ClientSecret"] = request.ClientSecret
	}

	if !dara.IsNil(request.DesktopAccessType) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EnableAdminAccess) {
		query["EnableAdminAccess"] = request.EnableAdminAccess
	}

	if !dara.IsNil(request.NeedVerifyLoginRisk) {
		query["NeedVerifyLoginRisk"] = request.NeedVerifyLoginRisk
	}

	if !dara.IsNil(request.NeedVerifyZeroDevice) {
		query["NeedVerifyZeroDevice"] = request.NeedVerifyZeroDevice
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteName) {
		query["OfficeSiteName"] = request.OfficeSiteName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfficeSiteAttribute"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the communication between cloud computers in an office network (formerly workspace). If you enable the communication between cloud computers in an office network, the cloud computers can access each other.
//
// @param request - ModifyOfficeSiteCrossDesktopAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfficeSiteCrossDesktopAccessResponse
func (client *Client) ModifyOfficeSiteCrossDesktopAccessWithContext(ctx context.Context, request *ModifyOfficeSiteCrossDesktopAccessRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableCrossDesktopAccess) {
		query["EnableCrossDesktopAccess"] = request.EnableCrossDesktopAccess
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfficeSiteCrossDesktopAccess"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DNS information of an office network.
//
// @param request - ModifyOfficeSiteDnsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfficeSiteDnsInfoResponse
func (client *Client) ModifyOfficeSiteDnsInfoWithContext(ctx context.Context, request *ModifyOfficeSiteDnsInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteDnsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsAddress) {
		query["DnsAddress"] = request.DnsAddress
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfficeSiteDnsInfo"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfficeSiteDnsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables multi-factor authentication (MFA) for an enterprise Active Directory (AD) office network (formerly workspace).
//
// @param request - ModifyOfficeSiteMfaEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfficeSiteMfaEnabledResponse
func (client *Client) ModifyOfficeSiteMfaEnabledWithContext(ctx context.Context, request *ModifyOfficeSiteMfaEnabledRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MfaEnabled) {
		query["MfaEnabled"] = request.MfaEnabled
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfficeSiteMfaEnabled"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the cloud computer policy.
//
// @param request - ModifyPolicyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroupWithContext(ctx context.Context, request *ModifyPolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminAccess) {
		query["AdminAccess"] = request.AdminAccess
	}

	if !dara.IsNil(request.AppContentProtection) {
		query["AppContentProtection"] = request.AppContentProtection
	}

	if !dara.IsNil(request.AuthorizeAccessPolicyRule) {
		query["AuthorizeAccessPolicyRule"] = request.AuthorizeAccessPolicyRule
	}

	if !dara.IsNil(request.AuthorizeSecurityPolicyRule) {
		query["AuthorizeSecurityPolicyRule"] = request.AuthorizeSecurityPolicyRule
	}

	if !dara.IsNil(request.CameraRedirect) {
		query["CameraRedirect"] = request.CameraRedirect
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Clipboard) {
		query["Clipboard"] = request.Clipboard
	}

	if !dara.IsNil(request.DeviceRedirects) {
		query["DeviceRedirects"] = request.DeviceRedirects
	}

	if !dara.IsNil(request.DeviceRules) {
		query["DeviceRules"] = request.DeviceRules
	}

	if !dara.IsNil(request.DomainList) {
		query["DomainList"] = request.DomainList
	}

	if !dara.IsNil(request.DomainResolveRule) {
		query["DomainResolveRule"] = request.DomainResolveRule
	}

	if !dara.IsNil(request.DomainResolveRuleType) {
		query["DomainResolveRuleType"] = request.DomainResolveRuleType
	}

	if !dara.IsNil(request.EndUserApplyAdminCoordinate) {
		query["EndUserApplyAdminCoordinate"] = request.EndUserApplyAdminCoordinate
	}

	if !dara.IsNil(request.EndUserGroupCoordinate) {
		query["EndUserGroupCoordinate"] = request.EndUserGroupCoordinate
	}

	if !dara.IsNil(request.GpuAcceleration) {
		query["GpuAcceleration"] = request.GpuAcceleration
	}

	if !dara.IsNil(request.Html5Access) {
		query["Html5Access"] = request.Html5Access
	}

	if !dara.IsNil(request.Html5FileTransfer) {
		query["Html5FileTransfer"] = request.Html5FileTransfer
	}

	if !dara.IsNil(request.InternetCommunicationProtocol) {
		query["InternetCommunicationProtocol"] = request.InternetCommunicationProtocol
	}

	if !dara.IsNil(request.LocalDrive) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !dara.IsNil(request.MaxReconnectTime) {
		query["MaxReconnectTime"] = request.MaxReconnectTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetRedirect) {
		query["NetRedirect"] = request.NetRedirect
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PreemptLogin) {
		query["PreemptLogin"] = request.PreemptLogin
	}

	if !dara.IsNil(request.PreemptLoginUser) {
		query["PreemptLoginUser"] = request.PreemptLoginUser
	}

	if !dara.IsNil(request.PrinterRedirection) {
		query["PrinterRedirection"] = request.PrinterRedirection
	}

	if !dara.IsNil(request.RecordContent) {
		query["RecordContent"] = request.RecordContent
	}

	if !dara.IsNil(request.RecordContentExpires) {
		query["RecordContentExpires"] = request.RecordContentExpires
	}

	if !dara.IsNil(request.Recording) {
		query["Recording"] = request.Recording
	}

	if !dara.IsNil(request.RecordingAudio) {
		query["RecordingAudio"] = request.RecordingAudio
	}

	if !dara.IsNil(request.RecordingDuration) {
		query["RecordingDuration"] = request.RecordingDuration
	}

	if !dara.IsNil(request.RecordingEndTime) {
		query["RecordingEndTime"] = request.RecordingEndTime
	}

	if !dara.IsNil(request.RecordingExpires) {
		query["RecordingExpires"] = request.RecordingExpires
	}

	if !dara.IsNil(request.RecordingFps) {
		query["RecordingFps"] = request.RecordingFps
	}

	if !dara.IsNil(request.RecordingStartTime) {
		query["RecordingStartTime"] = request.RecordingStartTime
	}

	if !dara.IsNil(request.RecordingUserNotify) {
		query["RecordingUserNotify"] = request.RecordingUserNotify
	}

	if !dara.IsNil(request.RecordingUserNotifyMessage) {
		query["RecordingUserNotifyMessage"] = request.RecordingUserNotifyMessage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteCoordinate) {
		query["RemoteCoordinate"] = request.RemoteCoordinate
	}

	if !dara.IsNil(request.RevokeAccessPolicyRule) {
		query["RevokeAccessPolicyRule"] = request.RevokeAccessPolicyRule
	}

	if !dara.IsNil(request.RevokeSecurityPolicyRule) {
		query["RevokeSecurityPolicyRule"] = request.RevokeSecurityPolicyRule
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	if !dara.IsNil(request.UsbRedirect) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !dara.IsNil(request.UsbSupplyRedirectRule) {
		query["UsbSupplyRedirectRule"] = request.UsbSupplyRedirectRule
	}

	if !dara.IsNil(request.VideoRedirect) {
		query["VideoRedirect"] = request.VideoRedirect
	}

	if !dara.IsNil(request.VisualQuality) {
		query["VisualQuality"] = request.VisualQuality
	}

	if !dara.IsNil(request.Watermark) {
		query["Watermark"] = request.Watermark
	}

	if !dara.IsNil(request.WatermarkAntiCam) {
		query["WatermarkAntiCam"] = request.WatermarkAntiCam
	}

	if !dara.IsNil(request.WatermarkColor) {
		query["WatermarkColor"] = request.WatermarkColor
	}

	if !dara.IsNil(request.WatermarkDegree) {
		query["WatermarkDegree"] = request.WatermarkDegree
	}

	if !dara.IsNil(request.WatermarkFontSize) {
		query["WatermarkFontSize"] = request.WatermarkFontSize
	}

	if !dara.IsNil(request.WatermarkFontStyle) {
		query["WatermarkFontStyle"] = request.WatermarkFontStyle
	}

	if !dara.IsNil(request.WatermarkPower) {
		query["WatermarkPower"] = request.WatermarkPower
	}

	if !dara.IsNil(request.WatermarkRowAmount) {
		query["WatermarkRowAmount"] = request.WatermarkRowAmount
	}

	if !dara.IsNil(request.WatermarkSecurity) {
		query["WatermarkSecurity"] = request.WatermarkSecurity
	}

	if !dara.IsNil(request.WatermarkTransparency) {
		query["WatermarkTransparency"] = request.WatermarkTransparency
	}

	if !dara.IsNil(request.WatermarkTransparencyValue) {
		query["WatermarkTransparencyValue"] = request.WatermarkTransparencyValue
	}

	if !dara.IsNil(request.WatermarkType) {
		query["WatermarkType"] = request.WatermarkType
	}

	if !dara.IsNil(request.WyAssistant) {
		query["WyAssistant"] = request.WyAssistant
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicyGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyResourceCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceCenterPolicyResponse
func (client *Client) ModifyResourceCenterPolicyWithContext(ctx context.Context, request *ModifyResourceCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceCenterPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyGroupIds) {
		query["PolicyGroupIds"] = request.PolicyGroupIds
	}

	if !dara.IsNil(request.PolicyGroupType) {
		query["PolicyGroupType"] = request.PolicyGroupType
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyResourceCenterPolicy"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourceCenterPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改办公网络维度安全组策略
//
// @param request - ModifySecurityGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupAttributeResponse
func (client *Client) ModifySecurityGroupAttributeWithContext(ctx context.Context, request *ModifySecurityGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizeEgress) {
		query["AuthorizeEgress"] = request.AuthorizeEgress
	}

	if !dara.IsNil(request.AuthorizeIngress) {
		query["AuthorizeIngress"] = request.AuthorizeIngress
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RevokeEgress) {
		query["RevokeEgress"] = request.RevokeEgress
	}

	if !dara.IsNil(request.RevokeIngress) {
		query["RevokeIngress"] = request.RevokeIngress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityGroupAttribute"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom cloud computer template.
//
// Description:
//
// *
//
// **Warning*	- This operation employs the full parameter update logic to maintain compatibility between the no-configuration logic and the default update logic. In other words, any unspecified parameters are treated as empty.
//
// @param request - ModifyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplateWithContext(ctx context.Context, request *ModifyTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskList) {
		bodyFlat["DataDiskList"] = request.DataDiskList
	}

	if !dara.IsNil(request.DefaultLanguage) {
		body["DefaultLanguage"] = request.DefaultLanguage
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PolicyGroupId) {
		body["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.PostPaidAfterUsedUp) {
		body["PostPaidAfterUsedUp"] = request.PostPaidAfterUsedUp
	}

	if !dara.IsNil(request.RegionConfigList) {
		bodyFlat["RegionConfigList"] = request.RegionConfigList
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceTagList) {
		bodyFlat["ResourceTagList"] = request.ResourceTagList
	}

	if !dara.IsNil(request.SiteConfigList) {
		bodyFlat["SiteConfigList"] = request.SiteConfigList
	}

	if !dara.IsNil(request.SystemDiskPerformanceLevel) {
		body["SystemDiskPerformanceLevel"] = request.SystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.SystemDiskSize) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TimerGroupId) {
		body["TimerGroupId"] = request.TimerGroupId
	}

	if !dara.IsNil(request.UserDuration) {
		body["UserDuration"] = request.UserDuration
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTemplate"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a custom cloud computer template, including the template name and template description.
//
// Description:
//
// This operation allows you to modify only the name and description of a custom cloud computer template. To change other parameters of the template, call the [ModifyTemplate](https://help.aliyun.com/document_detail/2925841.html) operation.
//
// @param request - ModifyTemplateBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateBaseInfoResponse
func (client *Client) ModifyTemplateBaseInfoWithContext(ctx context.Context, request *ModifyTemplateBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateBaseInfoResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTemplateBaseInfo"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTemplateBaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scheduled task configuration group.
//
// @param request - ModifyTimerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTimerGroupResponse
func (client *Client) ModifyTimerGroupWithContext(ctx context.Context, request *ModifyTimerGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyTimerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigTimers) {
		query["ConfigTimers"] = request.ConfigTimers
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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
		Action:      dara.String("ModifyTimerGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTimerGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on cloud desktops to end users, or revokes the permissions from the end users.
//
// Description:
//
// You can modify end users only for cloud computers that are in the Running state.
//
// @param request - ModifyUserEntitlementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserEntitlementResponse
func (client *Client) ModifyUserEntitlementWithContext(ctx context.Context, request *ModifyUserEntitlementRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserEntitlementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizeDesktopId) {
		query["AuthorizeDesktopId"] = request.AuthorizeDesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RevokeDesktopId) {
		query["RevokeDesktopId"] = request.RevokeDesktopId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserEntitlement"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserEntitlementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Replaces the existing authorized users of a cloud computer share with different users
//
// @param request - ModifyUserToDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserToDesktopGroupResponse
func (client *Client) ModifyUserToDesktopGroupWithContext(ctx context.Context, request *ModifyUserToDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.NewEndUserIds) {
		query["NewEndUserIds"] = request.NewEndUserIds
	}

	if !dara.IsNil(request.OldEndUserIds) {
		query["OldEndUserIds"] = request.OldEndUserIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserToDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Move files or folders.
//
// @param request - MoveCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveCdsFileResponse
func (client *Client) MoveCdsFileWithContext(ctx context.Context, request *MoveCdsFileRequest, runtime *dara.RuntimeOptions) (_result *MoveCdsFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.ConflictPolicy) {
		query["ConflictPolicy"] = request.ConflictPolicy
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveCdsFile"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveCdsFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart cloud computers.
//
// Description:
//
// The cloud computers that you want to restart must be in the Running state.
//
// @param request - RebootDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktopsWithContext(ctx context.Context, request *RebootDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.OsUpdate) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds images for one or more cloud computers.
//
// Description:
//
// Before you proceed, take note of the following limits:
//
//   - You cannot convert a cloud computer\\"s operating system image from one type to another (e.g., Windows to Linux or vice versa) in China (Hong Kong) or overseas regions.
//
//   - GPU and non-GPU images are not interchangeable, as graphic-based cloud computers can only use GPU-accelerated images, while other cloud computers are limited to non-GPU-accelerated images.
//
// When a cloud computer’s image is updated, the system initializes its system disk by using the new image, resulting in the following effects:
//
//   - All data on the original system disk is erased. Snapshots created from the original system disk become unavailable and are automatically deleted.
//
//   - If the OS changes, data on the original data disk is cleared, and snapshots created from the original data disk become unavailable and are automatically deleted. If the OS remains the same, data on the original data disk is retained, and snapshots from the original data disk remain available.
//
// @param request - RebuildDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildDesktopsResponse
func (client *Client) RebuildDesktopsWithContext(ctx context.Context, request *RebuildDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AfterStatus) {
		query["AfterStatus"] = request.AfterStatus
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除公网IP
//
// @param request - ReleaseIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseIpAddressResponse
func (client *Client) ReleaseIpAddressWithContext(ctx context.Context, request *ReleaseIpAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleaseIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipId) {
		query["EipId"] = request.EipId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseIpAddress"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unshare a folder on the network disk.
//
// @param tmpReq - RemoveFilePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFilePermissionResponse
func (client *Client) RemoveFilePermissionWithContext(ctx context.Context, tmpReq *RemoveFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *RemoveFilePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveFilePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MemberList) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, dara.String("MemberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CdsId) {
		query["CdsId"] = request.CdsId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MemberListShrink) {
		query["MemberList"] = request.MemberListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFilePermission"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFilePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes user access permissions for a cloud computer share. Once access permissions for a cloud computer share are revoked from a user, the user can no longer access any cloud computers within that share.
//
// @param request - RemoveUserFromDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromDesktopGroupResponse
func (client *Client) RemoveUserFromDesktopGroupWithContext(ctx context.Context, request *RemoveUserFromDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupIds) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SimpleUserGroupId) {
		query["SimpleUserGroupId"] = request.SimpleUserGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	if !dara.IsNil(request.UserOuPath) {
		query["UserOuPath"] = request.UserOuPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除超卖用户组用户
//
// @param request - RemoveUserFromDesktopOversoldUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromDesktopOversoldUserGroupResponse
func (client *Client) RemoveUserFromDesktopOversoldUserGroupWithContext(ctx context.Context, request *RemoveUserFromDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromDesktopOversoldUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.UserDesktopId) {
		query["UserDesktopId"] = request.UserDesktopId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromDesktopOversoldUserGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromDesktopOversoldUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a shared cloud computer.
//
// @param request - RenewDesktopGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewDesktopGroupResponse
func (client *Client) RenewDesktopGroupWithContext(ctx context.Context, request *RenewDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewDesktopGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewDesktopGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 续费桌面超卖组
//
// @param request - RenewDesktopOversoldGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewDesktopOversoldGroupResponse
func (client *Client) RenewDesktopOversoldGroupWithContext(ctx context.Context, request *RenewDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopOversoldGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OversoldGroupId) {
		query["OversoldGroupId"] = request.OversoldGroupId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewDesktopOversoldGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewDesktopOversoldGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews monthly subscription cloud computers.
//
// @param request - RenewDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewDesktopsResponse
func (client *Client) RenewDesktopsWithContext(ctx context.Context, request *RenewDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews premium bandwidth plans.
//
// @param request - RenewNetworkPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewNetworkPackagesResponse
func (client *Client) RenewNetworkPackagesWithContext(ctx context.Context, request *RenewNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *RenewNetworkPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.NetworkPackageId) {
		query["NetworkPackageId"] = request.NetworkPackageId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewNetworkPackages"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewNetworkPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets cloud computers of a cloud computer share.
//
// Description:
//
// >  You can call this operation to reset only cloud computers from a cloud computer share.
//
// @param request - ResetDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDesktopsResponse
func (client *Client) ResetDesktopsWithContext(ctx context.Context, request *ResetDesktopsRequest, runtime *dara.RuntimeOptions) (_result *ResetDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopGroupIds) {
		query["DesktopGroupIds"] = request.DesktopGroupIds
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.LastRetryTime) {
		query["LastRetryTime"] = request.LastRetryTime
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetScope) {
		query["ResetScope"] = request.ResetScope
	}

	if !dara.IsNil(request.ResetType) {
		query["ResetType"] = request.ResetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the mount target of a File Storage NAS (NAS) file system.
//
// Description:
//
// When you create a NAS file system, a mount target is automatically generated. By default, you do not need to modify the mount target of the NAS file system. If the mount target is disabled, you need to reset the mount target of the NAS file system.
//
// @param request - ResetNASDefaultMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetNASDefaultMountTargetResponse
func (client *Client) ResetNASDefaultMountTargetWithContext(ctx context.Context, request *ResetNASDefaultMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetNASDefaultMountTarget"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores the data of a disk from a snapshot.
//
// Description:
//
// Before you call this operation, make sure that the following operations are performed:
//
//   - The data that you want to retain is backed up.
//
//     > The disk restoration operation is irreversible. After you call this operation, the disk is restored to the status at the point in time when the snapshot was created. Data that is generated between the snapshot creation time and the current time is lost. Before you restore the disk based on the snapshot, make sure that you back up data.
//
//   - The cloud computer to which the disk belongs is stopped.
//
// @param request - ResetSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshotWithContext(ctx context.Context, request *ResetSnapshotRequest, runtime *dara.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.StopDesktop) {
		query["StopDesktop"] = request.StopDesktop
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetSnapshot"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the coordinate permissions.
//
// @param request - RevokeCoordinatePrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeCoordinatePrivilegeResponse
func (client *Client) RevokeCoordinatePrivilegeWithContext(ctx context.Context, request *RevokeCoordinatePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *RevokeCoordinatePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoId) {
		query["CoId"] = request.CoId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeCoordinatePrivilege"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeCoordinatePrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a PowerShell or batch (.bat) script on Windows cloud desktops.
//
// Description:
//
// You can use the RunCommand operation to run scripts only on Windows cloud desktops.
//
// @param request - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithContext(ctx context.Context, request *RunCommandRequest, runtime *dara.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandContent) {
		query["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.CommandRole) {
		query["CommandRole"] = request.CommandRole
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommand"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification code that is required when you bind an advanced office network to a Cloud Enterprise Network (CEN) instance that belongs to another Alibaba Cloud account.
//
// Description:
//
// You must call this operation to obtain the verification code that is required when you bind an advanced office network to a CEN instance that belongs to another Alibaba Cloud account. After you call this operation, the system sends a verification code to the email address associated with the Alibaba Cloud account to which the CEN instance belongs.
//
// @param request - SendVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerifyCodeResponse
func (client *Client) SendVerifyCodeWithContext(ctx context.Context, request *SendVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *SendVerifyCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VerifyCodeAction) {
		query["VerifyCodeAction"] = request.VerifyCodeAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendVerifyCode"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendVerifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an auto scaling policy for a multi-session cloud computer. Elastic Desktop Service allows multiple end users to share a cloud computer in a multi-session cloud computer pool. This helps save costs.
//
// @param request - SetDesktopGroupScaleTimerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDesktopGroupScaleTimerResponse
func (client *Client) SetDesktopGroupScaleTimerWithContext(ctx context.Context, request *SetDesktopGroupScaleTimerRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupScaleTimerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScaleTimerInfos) {
		query["ScaleTimerInfos"] = request.ScaleTimerInfos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDesktopGroupScaleTimer"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDesktopGroupScaleTimerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a scheduled start, stop, restart, or reset task for a cloud computer share.
//
// @param request - SetDesktopGroupTimerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDesktopGroupTimerResponse
func (client *Client) SetDesktopGroupTimerWithContext(ctx context.Context, request *SetDesktopGroupTimerRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupTimerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetType) {
		query["ResetType"] = request.ResetType
	}

	if !dara.IsNil(request.TimerType) {
		query["TimerType"] = request.TimerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDesktopGroupTimer"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDesktopGroupTimerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of a scheduled task for a cloud computer share, such as enabling or disabling it.
//
// @param request - SetDesktopGroupTimerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDesktopGroupTimerStatusResponse
func (client *Client) SetDesktopGroupTimerStatusWithContext(ctx context.Context, request *SetDesktopGroupTimerStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupTimerStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimerType) {
		query["TimerType"] = request.TimerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDesktopGroupTimerStatus"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDesktopGroupTimerStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set the cloud computer maintenance mode.
//
// Description:
//
// If you need to perform some maintenance operations on the cloud computer and want to prohibit end user from connecting and using the cloud computer during this period, you can switch it to maintenance mode.
//
// @param request - SetDesktopMaintenanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDesktopMaintenanceResponse
func (client *Client) SetDesktopMaintenanceWithContext(ctx context.Context, request *SetDesktopMaintenanceRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopMaintenanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDesktopMaintenance"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDesktopMaintenanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the single sign-on (SSO) status of an Active Directory (AD) directory.
//
// Description:
//
// This operation is supported only for AD directories, not for RAM directories.
//
// @param request - SetDirectorySsoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDirectorySsoStatusResponse
func (client *Client) SetDirectorySsoStatusWithContext(ctx context.Context, request *SetDirectorySsoStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EnableSso) {
		query["EnableSso"] = request.EnableSso
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDirectorySsoStatus"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads the metadata of a Security Assertion Markup Language (SAML) 2.0-based identity provider (IdP).
//
// Description:
//
// You can call this operation only for workspaces of the Active Directory (AD) and convenience account types.
//
// @param request - SetIdpMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdpMetadataResponse
func (client *Client) SetIdpMetadataWithContext(ctx context.Context, request *SetIdpMetadataRequest, runtime *dara.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.IdpMetadata) {
		query["IdpMetadata"] = request.IdpMetadata
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdpMetadata"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables single sign-on (SSO) for a workspace.
//
// @param request - SetOfficeSiteSsoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetOfficeSiteSsoStatusResponse
func (client *Client) SetOfficeSiteSsoStatusWithContext(ctx context.Context, request *SetOfficeSiteSsoStatusRequest, runtime *dara.RuntimeOptions) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableSso) {
		query["EnableSso"] = request.EnableSso
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetOfficeSiteSsoStatus"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures directories in the blacklist and whitelist based on the user profile management (UPM) feature.
//
// @param tmpReq - SetUserProfilePathRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserProfilePathRulesResponse
func (client *Client) SetUserProfilePathRulesWithContext(ctx context.Context, tmpReq *SetUserProfilePathRulesRequest, runtime *dara.RuntimeOptions) (_result *SetUserProfilePathRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetUserProfilePathRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserProfilePathRule) {
		request.UserProfilePathRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserProfilePathRule, dara.String("UserProfilePathRule"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserProfilePathRuleShrink) {
		query["UserProfilePathRule"] = request.UserProfilePathRuleShrink
	}

	if !dara.IsNil(request.UserProfileRuleType) {
		query["UserProfileRuleType"] = request.UserProfileRuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserProfilePathRules"),
		Version:     dara.String("2020-09-30"),
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
// Starts stopped cloud computers. After the API operation is successfully called, the cloud computers enter the Running state.
//
// Description:
//
// The cloud computers that you want to start must be in the Stopped state.
//
// @param request - StartDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDesktopsResponse
func (client *Client) StartDesktopsWithContext(ctx context.Context, request *StartDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop cloud computers that are in the Running state. After the operation is successfully called, the cloud computers enter the Stopped state.
//
// Description:
//
// The cloud computers that you want to stop must be in the Running state.
//
// @param request - StopDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDesktopsResponse
func (client *Client) StopDesktopsWithContext(ctx context.Context, request *StopDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.OsUpdate) {
		query["OsUpdate"] = request.OsUpdate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StoppedMode) {
		query["StoppedMode"] = request.StoppedMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a Cloud Assistant command that is running on one or more cloud desktops.
//
// Description:
//
// When you stop a one-time execution of a command, the command continues to run on the cloud desktops where it has started to run, and will not run on the cloud desktops where it has not started to run.
//
// @param request - StopInvocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInvocationResponse
func (client *Client) StopInvocationWithContext(ctx context.Context, request *StopInvocationRequest, runtime *dara.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.InvokeId) {
		query["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInvocation"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to cloud computers. This allows you to filter and manage cloud computers by tag.
//
// Description:
//
// If TagKey is specified, the new TagValue value overrides the original TagValue value.
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
		Version:     dara.String("2020-09-30"),
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
// Queries the transmission and approval result for a submitted file.
//
// @param request - TransferTaskApprovalCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferTaskApprovalCallbackResponse
func (client *Client) TransferTaskApprovalCallbackWithContext(ctx context.Context, request *TransferTaskApprovalCallbackRequest, runtime *dara.RuntimeOptions) (_result *TransferTaskApprovalCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssBucketRegionId) {
		query["OssBucketRegionId"] = request.OssBucketRegionId
	}

	if !dara.IsNil(request.Result) {
		query["Result"] = request.Result
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferTaskApprovalCallback"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferTaskApprovalCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a configuration group from resources.
//
// @param request - UnbindConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindConfigGroupResponse
func (client *Client) UnbindConfigGroupWithContext(ctx context.Context, request *UnbindConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *UnbindConfigGroupResponse, _err error) {
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

	if !dara.IsNil(request.ResourceInfos) {
		query["ResourceInfos"] = request.ResourceInfos
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindConfigGroup"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindConfigGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑用户桌面
//
// @param request - UnbindUserDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktopWithContext(ctx context.Context, request *UnbindUserDesktopRequest, runtime *dara.RuntimeOptions) (_result *UnbindUserDesktopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopAgentIds) {
		query["DesktopAgentIds"] = request.DesktopAgentIds
	}

	if !dara.IsNil(request.DesktopGroupId) {
		query["DesktopGroupId"] = request.DesktopGroupId
	}

	if !dara.IsNil(request.DesktopIds) {
		query["DesktopIds"] = request.DesktopIds
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.UserDesktopIds) {
		query["UserDesktopIds"] = request.UserDesktopIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindUserDesktop"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a virtual multi-factor authentication (MFA) device that is in the LOCKED state.
//
// @param request - UnlockVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockVirtualMFADeviceResponse
func (client *Client) UnlockVirtualMFADeviceWithContext(ctx context.Context, request *UnlockVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockVirtualMFADevice"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from cloud computers. After you remove a tag, if the tag is not added to a cloud computer, the tag is automatically deleted.
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
		Version:     dara.String("2020-09-30"),
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
// Enables or disables the auto-push feature for an image update task.
//
// Description:
//
// You can call this operation to manage each image update task. This operation is valid only when the auto-update switch in the image update module for global image updates is turned off. If the auto-update switch is turned on, the switches for each image update task are always turned on. If you want to turn on or off the auto-update switch, go to the Elastic Desktop Service console and choose **Operations > Image Updates*	- in the left-side navigation pane.
//
// @param request - UpdateFotaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFotaTaskResponse
func (client *Client) UpdateFotaTaskWithContext(ctx context.Context, request *UpdateFotaTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateFotaTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskUid) {
		query["TaskUid"] = request.TaskUid
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFotaTask"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFotaTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads your custom Windows image.
//
// Description:
//
// >  You can upload only Windows images.
//
// @param request - UploadImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadImageResponse
func (client *Client) UploadImageWithContext(ctx context.Context, request *UploadImageRequest, runtime *dara.RuntimeOptions) (_result *UploadImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableSecurityCheck) {
		query["EnableSecurityCheck"] = request.EnableSecurityCheck
	}

	if !dara.IsNil(request.GpuCategory) {
		query["GpuCategory"] = request.GpuCategory
	}

	if !dara.IsNil(request.GpuDriverType) {
		query["GpuDriverType"] = request.GpuDriverType
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.LicenseType) {
		query["LicenseType"] = request.LicenseType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OssObjectPath) {
		query["OssObjectPath"] = request.OssObjectPath
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadImage"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ID of a Cloud Enterprise Network (CEN) instance and the ID of the Alibaba Cloud account to which the instance belongs and checks whether a CIDR block conflict exists between the routes of the instance and the IPv4 CIDR blocks of the associated office network.
//
// @param request - VerifyCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyCenResponse
func (client *Client) VerifyCenWithContext(ctx context.Context, request *VerifyCenRequest, runtime *dara.RuntimeOptions) (_result *VerifyCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyCen"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Wakes up cloud computers.
//
// Description:
//
// Only cloud computers that are in the Hibernated state can be waked up.
//
// @param request - WakeupDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WakeupDesktopsResponse
func (client *Client) WakeupDesktopsWithContext(ctx context.Context, request *WakeupDesktopsRequest, runtime *dara.RuntimeOptions) (_result *WakeupDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WakeupDesktops"),
		Version:     dara.String("2020-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WakeupDesktopsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
