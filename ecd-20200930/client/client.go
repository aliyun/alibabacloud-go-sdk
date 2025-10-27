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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
func (client *Client) ActivateOfficeSiteWithOptions(request *ActivateOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *ActivateOfficeSiteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// @return ActivateOfficeSiteResponse
func (client *Client) ActivateOfficeSite(request *ActivateOfficeSiteRequest) (_result *ActivateOfficeSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateOfficeSiteResponse{}
	_body, _err := client.ActivateOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddDesktopOversoldUserGroupWithOptions(request *AddDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDesktopOversoldUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddDesktopOversoldUserGroupResponse
func (client *Client) AddDesktopOversoldUserGroup(request *AddDesktopOversoldUserGroupRequest) (_result *AddDesktopOversoldUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDesktopOversoldUserGroupResponse{}
	_body, _err := client.AddDesktopOversoldUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddDevicesWithOptions(request *AddDevicesRequest, runtime *dara.RuntimeOptions) (_result *AddDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddDevicesResponse
func (client *Client) AddDevices(request *AddDevicesRequest) (_result *AddDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDevicesResponse{}
	_body, _err := client.AddDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Shares a folder of a cloud disk with other users.
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
func (client *Client) AddFilePermissionWithOptions(tmpReq *AddFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *AddFilePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares a folder of a cloud disk with other users.
//
// Description:
//
// You can call this operation to share a specific folder with other users. You can also configure the folder permissions.
//
// @param request - AddFilePermissionRequest
//
// @return AddFilePermissionResponse
func (client *Client) AddFilePermission(request *AddFilePermissionRequest) (_result *AddFilePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFilePermissionResponse{}
	_body, _err := client.AddFilePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddUserToDesktopGroupWithOptions(request *AddUserToDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddUserToDesktopGroupResponse
func (client *Client) AddUserToDesktopGroup(request *AddUserToDesktopGroupRequest) (_result *AddUserToDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.AddUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddUserToDesktopOversoldUserGroupWithOptions(request *AddUserToDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToDesktopOversoldUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddUserToDesktopOversoldUserGroupResponse
func (client *Client) AddUserToDesktopOversoldUserGroup(request *AddUserToDesktopOversoldUserGroupRequest) (_result *AddUserToDesktopOversoldUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToDesktopOversoldUserGroupResponse{}
	_body, _err := client.AddUserToDesktopOversoldUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AllocateIpAddressWithOptions(request *AllocateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AllocateIpAddressResponse
func (client *Client) AllocateIpAddress(request *AllocateIpAddressRequest) (_result *AllocateIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateIpAddressResponse{}
	_body, _err := client.AllocateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ApplyAutoSnapshotPolicyWithOptions(request *ApplyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.ApplyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ApplyCoordinatePrivilegeWithOptions(request *ApplyCoordinatePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ApplyCoordinatePrivilegeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ApplyCoordinatePrivilegeResponse
func (client *Client) ApplyCoordinatePrivilege(request *ApplyCoordinatePrivilegeRequest) (_result *ApplyCoordinatePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyCoordinatePrivilegeResponse{}
	_body, _err := client.ApplyCoordinatePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ApplyCoordinationForMonitoringWithOptions(request *ApplyCoordinationForMonitoringRequest, runtime *dara.RuntimeOptions) (_result *ApplyCoordinationForMonitoringResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ApplyCoordinationForMonitoringResponse
func (client *Client) ApplyCoordinationForMonitoring(request *ApplyCoordinationForMonitoringRequest) (_result *ApplyCoordinationForMonitoringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyCoordinationForMonitoringResponse{}
	_body, _err := client.ApplyCoordinationForMonitoringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ApproveFotaUpdateWithOptions(request *ApproveFotaUpdateRequest, runtime *dara.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ApproveFotaUpdateResponse
func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (_result *ApproveFotaUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.ApproveFotaUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssociateIpAddressWithOptions(request *AssociateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AssociateIpAddressResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AssociateIpAddressResponse
func (client *Client) AssociateIpAddress(request *AssociateIpAddressRequest) (_result *AssociateIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateIpAddressResponse{}
	_body, _err := client.AssociateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssociateNetworkPackageWithOptions(request *AssociateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateNetworkPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AssociateNetworkPackageResponse
func (client *Client) AssociateNetworkPackage(request *AssociateNetworkPackageRequest) (_result *AssociateNetworkPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateNetworkPackageResponse{}
	_body, _err := client.AssociateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssociateRouteTableWithOptions(request *AssociateRouteTableRequest, runtime *dara.RuntimeOptions) (_result *AssociateRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AssociateRouteTableResponse
func (client *Client) AssociateRouteTable(request *AssociateRouteTableRequest) (_result *AssociateRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateRouteTableResponse{}
	_body, _err := client.AssociateRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AttachCenWithOptions(request *AttachCenRequest, runtime *dara.RuntimeOptions) (_result *AttachCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AttachCenResponse
func (client *Client) AttachCen(request *AttachCenRequest) (_result *AttachCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachCenResponse{}
	_body, _err := client.AttachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AttachEndUserWithOptions(request *AttachEndUserRequest, runtime *dara.RuntimeOptions) (_result *AttachEndUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AttachEndUserResponse
func (client *Client) AttachEndUser(request *AttachEndUserRequest) (_result *AttachEndUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachEndUserResponse{}
	_body, _err := client.AttachEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BindConfigGroupWithOptions(request *BindConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *BindConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BindConfigGroupResponse
func (client *Client) BindConfigGroup(request *BindConfigGroupRequest) (_result *BindConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindConfigGroupResponse{}
	_body, _err := client.BindConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelAutoSnapshotPolicyWithOptions(request *CancelAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CancelAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelCdsFileShareLinkWithOptions(request *CancelCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *CancelCdsFileShareLinkResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelCdsFileShareLinkResponse
func (client *Client) CancelCdsFileShareLink(request *CancelCdsFileShareLinkRequest) (_result *CancelCdsFileShareLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCdsFileShareLinkResponse{}
	_body, _err := client.CancelCdsFileShareLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelCoordinationForMonitoringWithOptions(request *CancelCoordinationForMonitoringRequest, runtime *dara.RuntimeOptions) (_result *CancelCoordinationForMonitoringResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelCoordinationForMonitoringResponse
func (client *Client) CancelCoordinationForMonitoring(request *CancelCoordinationForMonitoringRequest) (_result *CancelCoordinationForMonitoringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCoordinationForMonitoringResponse{}
	_body, _err := client.CancelCoordinationForMonitoringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelCopyImageWithOptions(request *CancelCopyImageRequest, runtime *dara.RuntimeOptions) (_result *CancelCopyImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelCopyImageResponse
func (client *Client) CancelCopyImage(request *CancelCopyImageRequest) (_result *CancelCopyImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCopyImageResponse{}
	_body, _err := client.CancelCopyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CloneCenterPolicyWithOptions(request *CloneCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *CloneCenterPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CloneCenterPolicyResponse
func (client *Client) CloneCenterPolicy(request *CloneCenterPolicyRequest) (_result *CloneCenterPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloneCenterPolicyResponse{}
	_body, _err := client.CloneCenterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ClonePolicyGroupWithOptions(request *ClonePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ClonePolicyGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ClonePolicyGroupResponse
func (client *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (_result *ClonePolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.ClonePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Completes a file uploading task.
//
// @param request - CompleteCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteCdsFileResponse
func (client *Client) CompleteCdsFileWithOptions(request *CompleteCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CompleteCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Completes a file uploading task.
//
// @param request - CompleteCdsFileRequest
//
// @return CompleteCdsFileResponse
func (client *Client) CompleteCdsFile(request *CompleteCdsFileRequest) (_result *CompleteCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompleteCdsFileResponse{}
	_body, _err := client.CompleteCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ConfigADConnectorTrustWithOptions(request *ConfigADConnectorTrustRequest, runtime *dara.RuntimeOptions) (_result *ConfigADConnectorTrustResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ConfigADConnectorTrustResponse
func (client *Client) ConfigADConnectorTrust(request *ConfigADConnectorTrustRequest) (_result *ConfigADConnectorTrustResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigADConnectorTrustResponse{}
	_body, _err := client.ConfigADConnectorTrustWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConfigADConnectorUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigADConnectorUserResponse
func (client *Client) ConfigADConnectorUserWithOptions(request *ConfigADConnectorUserRequest, runtime *dara.RuntimeOptions) (_result *ConfigADConnectorUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfigADConnectorUserRequest
//
// @return ConfigADConnectorUserResponse
func (client *Client) ConfigADConnectorUser(request *ConfigADConnectorUserRequest) (_result *ConfigADConnectorUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigADConnectorUserResponse{}
	_body, _err := client.ConfigADConnectorUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CopyCdsFileWithOptions(request *CopyCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CopyCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CopyCdsFileResponse
func (client *Client) CopyCdsFile(request *CopyCdsFileRequest) (_result *CopyCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyCdsFileResponse{}
	_body, _err := client.CopyCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CopyImageWithOptions(request *CopyImageRequest, runtime *dara.RuntimeOptions) (_result *CopyImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CopyImageResponse
func (client *Client) CopyImage(request *CopyImageRequest) (_result *CopyImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyImageResponse{}
	_body, _err := client.CopyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateADConnectorDirectoryWithOptions(request *CreateADConnectorDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateADConnectorDirectoryResponse
func (client *Client) CreateADConnectorDirectory(request *CreateADConnectorDirectoryRequest) (_result *CreateADConnectorDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CreateADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateADConnectorOfficeSiteWithOptions(request *CreateADConnectorOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateADConnectorOfficeSiteResponse
func (client *Client) CreateADConnectorOfficeSite(request *CreateADConnectorOfficeSiteRequest) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CreateADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAndBindNasFileSystemWithOptions(request *CreateAndBindNasFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateAndBindNasFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAndBindNasFileSystemResponse
func (client *Client) CreateAndBindNasFileSystem(request *CreateAndBindNasFileSystemRequest) (_result *CreateAndBindNasFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAndBindNasFileSystemResponse{}
	_body, _err := client.CreateAndBindNasFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAutoSnapshotPolicyWithOptions(request *CreateAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CreateAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBandwidthResourcePackagesWithOptions(request *CreateBandwidthResourcePackagesRequest, runtime *dara.RuntimeOptions) (_result *CreateBandwidthResourcePackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBandwidthResourcePackagesResponse
func (client *Client) CreateBandwidthResourcePackages(request *CreateBandwidthResourcePackagesRequest) (_result *CreateBandwidthResourcePackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBandwidthResourcePackagesResponse{}
	_body, _err := client.CreateBandwidthResourcePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateBundleWithOptions(request *CreateBundleRequest, runtime *dara.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateBundleResponse
func (client *Client) CreateBundle(request *CreateBundleRequest) (_result *CreateBundleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBundleResponse{}
	_body, _err := client.CreateBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCdsFileWithOptions(request *CreateCdsFileRequest, runtime *dara.RuntimeOptions) (_result *CreateCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCdsFileResponse
func (client *Client) CreateCdsFile(request *CreateCdsFileRequest) (_result *CreateCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCdsFileResponse{}
	_body, _err := client.CreateCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCdsFileShareLinkWithOptions(request *CreateCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *CreateCdsFileShareLinkResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCdsFileShareLinkResponse
func (client *Client) CreateCdsFileShareLink(request *CreateCdsFileShareLinkRequest) (_result *CreateCdsFileShareLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCdsFileShareLinkResponse{}
	_body, _err := client.CreateCdsFileShareLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCenterPolicyWithOptions(request *CreateCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCenterPolicyResponse, _err error) {
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

	if !dara.IsNil(request.AutoReconnect) {
		query["AutoReconnect"] = request.AutoReconnect
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCenterPolicyResponse
func (client *Client) CreateCenterPolicy(request *CreateCenterPolicyRequest) (_result *CreateCenterPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenterPolicyResponse{}
	_body, _err := client.CreateCenterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCloudDriveGroupWithOptions(request *CreateCloudDriveGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCloudDriveGroupResponse
func (client *Client) CreateCloudDriveGroup(request *CreateCloudDriveGroupRequest) (_result *CreateCloudDriveGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudDriveGroupResponse{}
	_body, _err := client.CreateCloudDriveGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCloudDriveServiceWithOptions(request *CreateCloudDriveServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCloudDriveServiceResponse
func (client *Client) CreateCloudDriveService(request *CreateCloudDriveServiceRequest) (_result *CreateCloudDriveServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudDriveServiceResponse{}
	_body, _err := client.CreateCloudDriveServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCloudDriveUsersWithOptions(request *CreateCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudDriveUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCloudDriveUsersResponse
func (client *Client) CreateCloudDriveUsers(request *CreateCloudDriveUsersRequest) (_result *CreateCloudDriveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudDriveUsersResponse{}
	_body, _err := client.CreateCloudDriveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateConfigGroupWithOptions(request *CreateConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateConfigGroupResponse
func (client *Client) CreateConfigGroup(request *CreateConfigGroupRequest) (_result *CreateConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConfigGroupResponse{}
	_body, _err := client.CreateConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDesktopGroupWithOptions(request *CreateDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDesktopGroupResponse
func (client *Client) CreateDesktopGroup(request *CreateDesktopGroupRequest) (_result *CreateDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDesktopOversoldGroupWithOptions(request *CreateDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDesktopOversoldGroupResponse
func (client *Client) CreateDesktopOversoldGroup(request *CreateDesktopOversoldGroupRequest) (_result *CreateDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDesktopOversoldGroupResponse{}
	_body, _err := client.CreateDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDesktopsWithOptions(tmpReq *CreateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateDesktopsRequest
//
// @return CreateDesktopsResponse
func (client *Client) CreateDesktops(request *CreateDesktopsRequest) (_result *CreateDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CreateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDiskEncryptionServiceWithOptions(request *CreateDiskEncryptionServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskEncryptionServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDiskEncryptionServiceResponse
func (client *Client) CreateDiskEncryptionService(request *CreateDiskEncryptionServiceRequest) (_result *CreateDiskEncryptionServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiskEncryptionServiceResponse{}
	_body, _err := client.CreateDiskEncryptionServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网盘
//
// @param request - CreateDriveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDriveResponse
func (client *Client) CreateDriveWithOptions(request *CreateDriveRequest, runtime *dara.RuntimeOptions) (_result *CreateDriveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网盘
//
// @param request - CreateDriveRequest
//
// @return CreateDriveResponse
func (client *Client) CreateDrive(request *CreateDriveRequest) (_result *CreateDriveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDriveResponse{}
	_body, _err := client.CreateDriveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建无影数据报表导出任务
//
// @param request - CreateEcdReportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEcdReportTaskResponse
func (client *Client) CreateEcdReportTaskWithOptions(request *CreateEcdReportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateEcdReportTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建无影数据报表导出任务
//
// @param request - CreateEcdReportTaskRequest
//
// @return CreateEcdReportTaskResponse
func (client *Client) CreateEcdReportTask(request *CreateEcdReportTaskRequest) (_result *CreateEcdReportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEcdReportTaskResponse{}
	_body, _err := client.CreateEcdReportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateForwardEntryWithOptions(request *CreateForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateForwardEntryResponse
func (client *Client) CreateForwardEntry(request *CreateForwardEntryRequest) (_result *CreateForwardEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateForwardEntryResponse{}
	_body, _err := client.CreateForwardEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *dara.RuntimeOptions) (_result *CreateImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateImageResponse
func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateNASFileSystemWithOptions(request *CreateNASFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateNASFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateNASFileSystemResponse
func (client *Client) CreateNASFileSystem(request *CreateNASFileSystemRequest) (_result *CreateNASFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CreateNASFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateNatGatewayWithOptions(request *CreateNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateNatGatewayResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateNatGatewayResponse
func (client *Client) CreateNatGateway(request *CreateNatGatewayRequest) (_result *CreateNatGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNatGatewayResponse{}
	_body, _err := client.CreateNatGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateNetworkPackageWithOptions(request *CreateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateNetworkPackageResponse
func (client *Client) CreateNetworkPackage(request *CreateNetworkPackageRequest) (_result *CreateNetworkPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CreateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePolicyGroupWithOptions(request *CreatePolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePolicyGroupResponse
func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRAMDirectoryWithOptions(request *CreateRAMDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRAMDirectoryResponse
func (client *Client) CreateRAMDirectory(request *CreateRAMDirectoryRequest) (_result *CreateRAMDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CreateRAMDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRouteEntryWithOptions(request *CreateRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRouteEntryResponse
func (client *Client) CreateRouteEntry(request *CreateRouteEntryRequest) (_result *CreateRouteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRouteEntryResponse{}
	_body, _err := client.CreateRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRouteTableWithOptions(request *CreateRouteTableRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRouteTableResponse
func (client *Client) CreateRouteTable(request *CreateRouteTableRequest) (_result *CreateRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRouteTableResponse{}
	_body, _err := client.CreateRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSimpleOfficeSiteWithOptions(request *CreateSimpleOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSimpleOfficeSiteResponse
func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CreateSimpleOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSnatEntryWithOptions(request *CreateSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateSnatEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSnatEntryResponse
func (client *Client) CreateSnatEntry(request *CreateSnatEntryRequest) (_result *CreateSnatEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSnatEntryResponse{}
	_body, _err := client.CreateSnatEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSubnetWithOptions(request *CreateSubnetRequest, runtime *dara.RuntimeOptions) (_result *CreateSubnetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSubnetResponse
func (client *Client) CreateSubnet(request *CreateSubnetRequest) (_result *CreateSubnetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSubnetResponse{}
	_body, _err := client.CreateSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模板
//
// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模板
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
// Deletes an automatic snapshot policy.
//
// @param request - DeleteAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicyWithOptions(request *DeleteAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DeleteAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteBundlesWithOptions(request *DeleteBundlesRequest, runtime *dara.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteBundlesResponse
func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (_result *DeleteBundlesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DeleteBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file from a cloud disk in Cloud Drive Service.
//
// @param request - DeleteCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCdsFileResponse
func (client *Client) DeleteCdsFileWithOptions(request *DeleteCdsFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file from a cloud disk in Cloud Drive Service.
//
// @param request - DeleteCdsFileRequest
//
// @return DeleteCdsFileResponse
func (client *Client) DeleteCdsFile(request *DeleteCdsFileRequest) (_result *DeleteCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCdsFileResponse{}
	_body, _err := client.DeleteCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCenterPolicyWithOptions(request *DeleteCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenterPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCenterPolicyResponse
func (client *Client) DeleteCenterPolicy(request *DeleteCenterPolicyRequest) (_result *DeleteCenterPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenterPolicyResponse{}
	_body, _err := client.DeleteCenterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCloudDriveGroupsWithOptions(request *DeleteCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudDriveGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCloudDriveGroupsResponse
func (client *Client) DeleteCloudDriveGroups(request *DeleteCloudDriveGroupsRequest) (_result *DeleteCloudDriveGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudDriveGroupsResponse{}
	_body, _err := client.DeleteCloudDriveGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteCloudDriveUsersWithOptions(request *DeleteCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudDriveUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteCloudDriveUsersResponse
func (client *Client) DeleteCloudDriveUsers(request *DeleteCloudDriveUsersRequest) (_result *DeleteCloudDriveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudDriveUsersResponse{}
	_body, _err := client.DeleteCloudDriveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteConfigGroupWithOptions(request *DeleteConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteConfigGroupResponse
func (client *Client) DeleteConfigGroup(request *DeleteConfigGroupRequest) (_result *DeleteConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConfigGroupResponse{}
	_body, _err := client.DeleteConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDesktopGroupWithOptions(request *DeleteDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDesktopGroupResponse
func (client *Client) DeleteDesktopGroup(request *DeleteDesktopGroupRequest) (_result *DeleteDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DeleteDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDesktopsWithOptions(request *DeleteDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDesktopsResponse
func (client *Client) DeleteDesktops(request *DeleteDesktopsRequest) (_result *DeleteDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DeleteDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDevicesWithOptions(request *DeleteDevicesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDevicesResponse
func (client *Client) DeleteDevices(request *DeleteDevicesRequest) (_result *DeleteDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.DeleteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDirectoriesWithOptions(request *DeleteDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDirectoriesResponse
func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (_result *DeleteDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DeleteDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网盘
//
// @param request - DeleteDriveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDriveResponse
func (client *Client) DeleteDriveWithOptions(request *DeleteDriveRequest, runtime *dara.RuntimeOptions) (_result *DeleteDriveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网盘
//
// @param request - DeleteDriveRequest
//
// @return DeleteDriveResponse
func (client *Client) DeleteDrive(request *DeleteDriveRequest) (_result *DeleteDriveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDriveResponse{}
	_body, _err := client.DeleteDriveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteEduRoomRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEduRoomResponse
func (client *Client) DeleteEduRoomWithOptions(request *DeleteEduRoomRequest, runtime *dara.RuntimeOptions) (_result *DeleteEduRoomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEduRoomRequest
//
// @return DeleteEduRoomResponse
func (client *Client) DeleteEduRoom(request *DeleteEduRoomRequest) (_result *DeleteEduRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEduRoomResponse{}
	_body, _err := client.DeleteEduRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteForwardEntryWithOptions(request *DeleteForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteForwardEntryResponse
func (client *Client) DeleteForwardEntry(request *DeleteForwardEntryRequest) (_result *DeleteForwardEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteForwardEntryResponse{}
	_body, _err := client.DeleteForwardEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *dara.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteImagesResponse
func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteNASFileSystemsWithOptions(request *DeleteNASFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DeleteNASFileSystemsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteNASFileSystemsResponse
func (client *Client) DeleteNASFileSystems(request *DeleteNASFileSystemsRequest) (_result *DeleteNASFileSystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DeleteNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteNatGatewayWithOptions(request *DeleteNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatGatewayResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteNatGatewayResponse
func (client *Client) DeleteNatGateway(request *DeleteNatGatewayRequest) (_result *DeleteNatGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNatGatewayResponse{}
	_body, _err := client.DeleteNatGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteNetworkPackagesWithOptions(request *DeleteNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkPackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteNetworkPackagesResponse
func (client *Client) DeleteNetworkPackages(request *DeleteNetworkPackagesRequest) (_result *DeleteNetworkPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DeleteNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteOfficeSitesWithOptions(request *DeleteOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DeleteOfficeSitesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteOfficeSitesResponse
func (client *Client) DeleteOfficeSites(request *DeleteOfficeSitesRequest) (_result *DeleteOfficeSitesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DeleteOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePolicyGroupsWithOptions(request *DeletePolicyGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePolicyGroupsResponse
func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (_result *DeletePolicyGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DeletePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRouteEntryWithOptions(request *DeleteRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRouteEntryResponse
func (client *Client) DeleteRouteEntry(request *DeleteRouteEntryRequest) (_result *DeleteRouteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteEntryResponse{}
	_body, _err := client.DeleteRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRouteTableWithOptions(request *DeleteRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRouteTableResponse
func (client *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (_result *DeleteRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteTableResponse{}
	_body, _err := client.DeleteRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSnatEntryWithOptions(request *DeleteSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnatEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSnatEntryResponse
func (client *Client) DeleteSnatEntry(request *DeleteSnatEntryRequest) (_result *DeleteSnatEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSnatEntryResponse{}
	_body, _err := client.DeleteSnatEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSubnetWithOptions(request *DeleteSubnetRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubnetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSubnetResponse
func (client *Client) DeleteSubnet(request *DeleteSubnetRequest) (_result *DeleteSubnetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSubnetResponse{}
	_body, _err := client.DeleteSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTemplatesWithOptions(request *DeleteTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAclEntriesWithOptions(request *DescribeAclEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAclEntriesResponse
func (client *Client) DescribeAclEntries(request *DescribeAclEntriesRequest) (_result *DescribeAclEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclEntriesResponse{}
	_body, _err := client.DescribeAclEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAutoSnapshotPolicyWithOptions(request *DescribeAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAutoSnapshotPolicyResponse
func (client *Client) DescribeAutoSnapshotPolicy(request *DescribeAutoSnapshotPolicyRequest) (_result *DescribeAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoSnapshotPolicyResponse{}
	_body, _err := client.DescribeAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeBundlesWithOptions(request *DescribeBundlesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeBundlesResponse
func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (_result *DescribeBundlesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DescribeBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCdsFileShareLinksWithOptions(request *DescribeCdsFileShareLinksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdsFileShareLinksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCdsFileShareLinksResponse
func (client *Client) DescribeCdsFileShareLinks(request *DescribeCdsFileShareLinksRequest) (_result *DescribeCdsFileShareLinksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdsFileShareLinksResponse{}
	_body, _err := client.DescribeCdsFileShareLinksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCensResponse
func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCenterPolicyListWithOptions(request *DescribeCenterPolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenterPolicyListResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCenterPolicyListResponse
func (client *Client) DescribeCenterPolicyList(request *DescribeCenterPolicyListRequest) (_result *DescribeCenterPolicyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenterPolicyListResponse{}
	_body, _err := client.DescribeCenterPolicyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeClientEventsWithOptions(request *DescribeClientEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeClientEventsResponse
func (client *Client) DescribeClientEvents(request *DescribeClientEventsRequest) (_result *DescribeClientEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DescribeClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCloudDriveGroupsWithOptions(request *DescribeCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDriveGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCloudDriveGroupsResponse
func (client *Client) DescribeCloudDriveGroups(request *DescribeCloudDriveGroupsRequest) (_result *DescribeCloudDriveGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudDriveGroupsResponse{}
	_body, _err := client.DescribeCloudDriveGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCloudDrivePermissionsWithOptions(request *DescribeCloudDrivePermissionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDrivePermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCloudDrivePermissionsResponse
func (client *Client) DescribeCloudDrivePermissions(request *DescribeCloudDrivePermissionsRequest) (_result *DescribeCloudDrivePermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudDrivePermissionsResponse{}
	_body, _err := client.DescribeCloudDrivePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCloudDriveUsersWithOptions(request *DescribeCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDriveUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCloudDriveUsersResponse
func (client *Client) DescribeCloudDriveUsers(request *DescribeCloudDriveUsersRequest) (_result *DescribeCloudDriveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudDriveUsersResponse{}
	_body, _err := client.DescribeCloudDriveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeConfigGroupWithOptions(request *DescribeConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeConfigGroupResponse
func (client *Client) DescribeConfigGroup(request *DescribeConfigGroupRequest) (_result *DescribeConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConfigGroupResponse{}
	_body, _err := client.DescribeConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCustomizedListHeadersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomizedListHeadersResponse
func (client *Client) DescribeCustomizedListHeadersWithOptions(request *DescribeCustomizedListHeadersRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomizedListHeadersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCustomizedListHeadersRequest
//
// @return DescribeCustomizedListHeadersResponse
func (client *Client) DescribeCustomizedListHeaders(request *DescribeCustomizedListHeadersRequest) (_result *DescribeCustomizedListHeadersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomizedListHeadersResponse{}
	_body, _err := client.DescribeCustomizedListHeadersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sessions in a desktop group.
//
// @param request - DescribeDesktopGroupSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopGroupSessionsResponse
func (client *Client) DescribeDesktopGroupSessionsWithOptions(request *DescribeDesktopGroupSessionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopGroupSessionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sessions in a desktop group.
//
// @param request - DescribeDesktopGroupSessionsRequest
//
// @return DescribeDesktopGroupSessionsResponse
func (client *Client) DescribeDesktopGroupSessions(request *DescribeDesktopGroupSessionsRequest) (_result *DescribeDesktopGroupSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopGroupSessionsResponse{}
	_body, _err := client.DescribeDesktopGroupSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopGroupsWithOptions(request *DescribeDesktopGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopGroupsResponse
func (client *Client) DescribeDesktopGroups(request *DescribeDesktopGroupsRequest) (_result *DescribeDesktopGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopGroupsResponse{}
	_body, _err := client.DescribeDesktopGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopInfoWithOptions(request *DescribeDesktopInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopInfoResponse
func (client *Client) DescribeDesktopInfo(request *DescribeDesktopInfoRequest) (_result *DescribeDesktopInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopInfoResponse{}
	_body, _err := client.DescribeDesktopInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopOversoldGroupWithOptions(request *DescribeDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopOversoldGroupResponse
func (client *Client) DescribeDesktopOversoldGroup(request *DescribeDesktopOversoldGroupRequest) (_result *DescribeDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopOversoldGroupResponse{}
	_body, _err := client.DescribeDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopOversoldUserWithOptions(request *DescribeDesktopOversoldUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopOversoldUserResponse
func (client *Client) DescribeDesktopOversoldUser(request *DescribeDesktopOversoldUserRequest) (_result *DescribeDesktopOversoldUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopOversoldUserResponse{}
	_body, _err := client.DescribeDesktopOversoldUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopOversoldUserGroupWithOptions(request *DescribeDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopOversoldUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopOversoldUserGroupResponse
func (client *Client) DescribeDesktopOversoldUserGroup(request *DescribeDesktopOversoldUserGroupRequest) (_result *DescribeDesktopOversoldUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopOversoldUserGroupResponse{}
	_body, _err := client.DescribeDesktopOversoldUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopSessionsWithOptions(request *DescribeDesktopSessionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopSessionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopSessionsResponse
func (client *Client) DescribeDesktopSessions(request *DescribeDesktopSessionsRequest) (_result *DescribeDesktopSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopSessionsResponse{}
	_body, _err := client.DescribeDesktopSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopTypesWithOptions(request *DescribeDesktopTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopTypesResponse
func (client *Client) DescribeDesktopTypes(request *DescribeDesktopTypesRequest) (_result *DescribeDesktopTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DescribeDesktopTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopsResponse
func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDesktopsInGroupWithOptions(request *DescribeDesktopsInGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopsInGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDesktopsInGroupResponse
func (client *Client) DescribeDesktopsInGroup(request *DescribeDesktopsInGroupRequest) (_result *DescribeDesktopsInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DescribeDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDevicesResponse
func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (_result *DescribeDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DescribeDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网盘列表
//
// @param request - DescribeDrivesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrivesResponse
func (client *Client) DescribeDrivesWithOptions(request *DescribeDrivesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrivesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网盘列表
//
// @param request - DescribeDrivesRequest
//
// @return DescribeDrivesResponse
func (client *Client) DescribeDrives(request *DescribeDrivesRequest) (_result *DescribeDrivesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDrivesResponse{}
	_body, _err := client.DescribeDrivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据报表导出任务列表
//
// @param request - DescribeEcdReportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEcdReportTasksResponse
func (client *Client) DescribeEcdReportTasksWithOptions(request *DescribeEcdReportTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeEcdReportTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据报表导出任务列表
//
// @param request - DescribeEcdReportTasksRequest
//
// @return DescribeEcdReportTasksResponse
func (client *Client) DescribeEcdReportTasks(request *DescribeEcdReportTasksRequest) (_result *DescribeEcdReportTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEcdReportTasksResponse{}
	_body, _err := client.DescribeEcdReportTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeFlowMetricWithOptions(request *DescribeFlowMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowMetricResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeFlowMetricResponse
func (client *Client) DescribeFlowMetric(request *DescribeFlowMetricRequest) (_result *DescribeFlowMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.DescribeFlowMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeFlowStatisticWithOptions(request *DescribeFlowStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowStatisticResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeFlowStatisticResponse
func (client *Client) DescribeFlowStatistic(request *DescribeFlowStatisticRequest) (_result *DescribeFlowStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowStatisticResponse{}
	_body, _err := client.DescribeFlowStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeForwardTableEntriesWithOptions(request *DescribeForwardTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeForwardTableEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeForwardTableEntriesResponse
func (client *Client) DescribeForwardTableEntries(request *DescribeForwardTableEntriesRequest) (_result *DescribeForwardTableEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeForwardTableEntriesResponse{}
	_body, _err := client.DescribeForwardTableEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeFotaPendingDesktopsWithOptions(request *DescribeFotaPendingDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFotaPendingDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeFotaPendingDesktopsResponse
func (client *Client) DescribeFotaPendingDesktops(request *DescribeFotaPendingDesktopsRequest) (_result *DescribeFotaPendingDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFotaPendingDesktopsResponse{}
	_body, _err := client.DescribeFotaPendingDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeFotaTasksWithOptions(request *DescribeFotaTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeFotaTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeFotaTasksResponse
func (client *Client) DescribeFotaTasks(request *DescribeFotaTasksRequest) (_result *DescribeFotaTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFotaTasksResponse{}
	_body, _err := client.DescribeFotaTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询全局桌面记录
//
// @param request - DescribeGlobalDesktopRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDesktopRecordsResponse
func (client *Client) DescribeGlobalDesktopRecordsWithOptions(request *DescribeGlobalDesktopRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDesktopRecordsResponse, _err error) {
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

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全局桌面记录
//
// @param request - DescribeGlobalDesktopRecordsRequest
//
// @return DescribeGlobalDesktopRecordsResponse
func (client *Client) DescribeGlobalDesktopRecords(request *DescribeGlobalDesktopRecordsRequest) (_result *DescribeGlobalDesktopRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDesktopRecordsResponse{}
	_body, _err := client.DescribeGlobalDesktopRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeGuestApplicationsWithOptions(request *DescribeGuestApplicationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGuestApplicationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeGuestApplicationsResponse
func (client *Client) DescribeGuestApplications(request *DescribeGuestApplicationsRequest) (_result *DescribeGuestApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGuestApplicationsResponse{}
	_body, _err := client.DescribeGuestApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeImageModifiedRecordsWithOptions(request *DescribeImageModifiedRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageModifiedRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeImageModifiedRecordsResponse
func (client *Client) DescribeImageModifiedRecords(request *DescribeImageModifiedRecordsRequest) (_result *DescribeImageModifiedRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageModifiedRecordsResponse{}
	_body, _err := client.DescribeImageModifiedRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeImagePermissionWithOptions(request *DescribeImagePermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeImagePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeImagePermissionResponse
func (client *Client) DescribeImagePermission(request *DescribeImagePermissionRequest) (_result *DescribeImagePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImagePermissionResponse{}
	_body, _err := client.DescribeImagePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeImagesResponse
func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeIpAddressesWithOptions(request *DescribeIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpAddressesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeIpAddressesResponse
func (client *Client) DescribeIpAddresses(request *DescribeIpAddressesRequest) (_result *DescribeIpAddressesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpAddressesResponse{}
	_body, _err := client.DescribeIpAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeKmsKeysWithOptions(request *DescribeKmsKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeKmsKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeKmsKeysResponse
func (client *Client) DescribeKmsKeys(request *DescribeKmsKeysRequest) (_result *DescribeKmsKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKmsKeysResponse{}
	_body, _err := client.DescribeKmsKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeModificationPriceWithOptions(request *DescribeModificationPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeModificationPriceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeModificationPriceResponse
func (client *Client) DescribeModificationPrice(request *DescribeModificationPriceRequest) (_result *DescribeModificationPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DescribeModificationPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeNASFileSystemsWithOptions(request *DescribeNASFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNASFileSystemsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeNASFileSystemsResponse
func (client *Client) DescribeNASFileSystems(request *DescribeNASFileSystemsRequest) (_result *DescribeNASFileSystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DescribeNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeNatGatewaysWithOptions(request *DescribeNatGatewaysRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatGatewaysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeNatGatewaysResponse
func (client *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (_result *DescribeNatGatewaysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatGatewaysResponse{}
	_body, _err := client.DescribeNatGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeNetworkPackagesWithOptions(request *DescribeNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkPackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeNetworkPackagesResponse
func (client *Client) DescribeNetworkPackages(request *DescribeNetworkPackagesRequest) (_result *DescribeNetworkPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DescribeNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeOfficeSitesResponse
func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePolicyGroupsWithOptions(request *DescribePolicyGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePolicyGroupsResponse
func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (_result *DescribePolicyGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DescribePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePriceResponse
func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePriceForCreateDesktopOversoldGroupWithOptions(request *DescribePriceForCreateDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForCreateDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePriceForCreateDesktopOversoldGroupResponse
func (client *Client) DescribePriceForCreateDesktopOversoldGroup(request *DescribePriceForCreateDesktopOversoldGroupRequest) (_result *DescribePriceForCreateDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceForCreateDesktopOversoldGroupResponse{}
	_body, _err := client.DescribePriceForCreateDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePriceForModifyDesktopOversoldGroupSaleWithOptions(request *DescribePriceForModifyDesktopOversoldGroupSaleRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForModifyDesktopOversoldGroupSaleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePriceForModifyDesktopOversoldGroupSaleResponse
func (client *Client) DescribePriceForModifyDesktopOversoldGroupSale(request *DescribePriceForModifyDesktopOversoldGroupSaleRequest) (_result *DescribePriceForModifyDesktopOversoldGroupSaleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceForModifyDesktopOversoldGroupSaleResponse{}
	_body, _err := client.DescribePriceForModifyDesktopOversoldGroupSaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePriceForRenewDesktopOversoldGroupWithOptions(request *DescribePriceForRenewDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceForRenewDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePriceForRenewDesktopOversoldGroupResponse
func (client *Client) DescribePriceForRenewDesktopOversoldGroup(request *DescribePriceForRenewDesktopOversoldGroupRequest) (_result *DescribePriceForRenewDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceForRenewDesktopOversoldGroupResponse{}
	_body, _err := client.DescribePriceForRenewDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRecordingsWithOptions(request *DescribeRecordingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordingsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRecordingsResponse
func (client *Client) DescribeRecordings(request *DescribeRecordingsRequest) (_result *DescribeRecordingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordingsResponse{}
	_body, _err := client.DescribeRecordingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRefundPriceWithOptions(request *DescribeRefundPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRefundPriceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRefundPriceResponse
func (client *Client) DescribeRefundPrice(request *DescribeRefundPriceRequest) (_result *DescribeRefundPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRefundPriceResponse{}
	_body, _err := client.DescribeRefundPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Queries the renewal price of an Alibaba Cloud Workspace service.
//
// @param request - DescribeRenewalPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenewalPriceResponse
func (client *Client) DescribeRenewalPriceWithOptions(request *DescribeRenewalPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenewalPriceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRenewalPriceResponse
func (client *Client) DescribeRenewalPrice(request *DescribeRenewalPriceRequest) (_result *DescribeRenewalPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRenewalPriceResponse{}
	_body, _err := client.DescribeRenewalPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeResourceByCenterPolicyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceByCenterPolicyIdResponse
func (client *Client) DescribeResourceByCenterPolicyIdWithOptions(request *DescribeResourceByCenterPolicyIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceByCenterPolicyIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeResourceByCenterPolicyIdRequest
//
// @return DescribeResourceByCenterPolicyIdResponse
func (client *Client) DescribeResourceByCenterPolicyId(request *DescribeResourceByCenterPolicyIdRequest) (_result *DescribeResourceByCenterPolicyIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceByCenterPolicyIdResponse{}
	_body, _err := client.DescribeResourceByCenterPolicyIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRouteEntryListWithOptions(request *DescribeRouteEntryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteEntryListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRouteEntryListResponse
func (client *Client) DescribeRouteEntryList(request *DescribeRouteEntryListRequest) (_result *DescribeRouteEntryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRouteEntryListResponse{}
	_body, _err := client.DescribeRouteEntryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRouteTableListWithOptions(request *DescribeRouteTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteTableListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRouteTableListResponse
func (client *Client) DescribeRouteTableList(request *DescribeRouteTableListRequest) (_result *DescribeRouteTableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRouteTableListResponse{}
	_body, _err := client.DescribeRouteTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeSessionStatisticWithOptions(request *DescribeSessionStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSessionStatisticResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeSessionStatisticResponse
func (client *Client) DescribeSessionStatistic(request *DescribeSessionStatisticRequest) (_result *DescribeSessionStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSessionStatisticResponse{}
	_body, _err := client.DescribeSessionStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeSnatTableEntriesWithOptions(request *DescribeSnatTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnatTableEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeSnatTableEntriesResponse
func (client *Client) DescribeSnatTableEntries(request *DescribeSnatTableEntriesRequest) (_result *DescribeSnatTableEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSnatTableEntriesResponse{}
	_body, _err := client.DescribeSnatTableEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeSubnetsWithOptions(request *DescribeSubnetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubnetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeSubnetsResponse
func (client *Client) DescribeSubnets(request *DescribeSubnetsRequest) (_result *DescribeSubnetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubnetsResponse{}
	_body, _err := client.DescribeSubnetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模板列表
//
// @param request - DescribeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模板列表
//
// @param request - DescribeTemplatesRequest
//
// @return DescribeTemplatesResponse
func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTimerGroupWithOptions(request *DescribeTimerGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeTimerGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTimerGroupResponse
func (client *Client) DescribeTimerGroup(request *DescribeTimerGroupRequest) (_result *DescribeTimerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTimerGroupResponse{}
	_body, _err := client.DescribeTimerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUserConnectTimeWithOptions(request *DescribeUserConnectTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserConnectTimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUserConnectTimeResponse
func (client *Client) DescribeUserConnectTime(request *DescribeUserConnectTimeRequest) (_result *DescribeUserConnectTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserConnectTimeResponse{}
	_body, _err := client.DescribeUserConnectTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUserConnectionRecordsWithOptions(request *DescribeUserConnectionRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserConnectionRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUserConnectionRecordsResponse
func (client *Client) DescribeUserConnectionRecords(request *DescribeUserConnectionRecordsRequest) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DescribeUserConnectionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUserProfilePathRulesWithOptions(request *DescribeUserProfilePathRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserProfilePathRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUserProfilePathRulesResponse
func (client *Client) DescribeUserProfilePathRules(request *DescribeUserProfilePathRulesRequest) (_result *DescribeUserProfilePathRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserProfilePathRulesResponse{}
	_body, _err := client.DescribeUserProfilePathRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUsersInGroupWithOptions(request *DescribeUsersInGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersInGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUsersInGroupResponse
func (client *Client) DescribeUsersInGroup(request *DescribeUsersInGroupRequest) (_result *DescribeUsersInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DescribeUsersInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUsersPasswordWithOptions(request *DescribeUsersPasswordRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersPasswordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUsersPasswordResponse
func (client *Client) DescribeUsersPassword(request *DescribeUsersPasswordRequest) (_result *DescribeUsersPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsersPasswordResponse{}
	_body, _err := client.DescribeUsersPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeVirtualMFADevicesWithOptions(request *DescribeVirtualMFADevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeVirtualMFADevicesResponse
func (client *Client) DescribeVirtualMFADevices(request *DescribeVirtualMFADevicesRequest) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DescribeVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DetachCenWithOptions(request *DetachCenRequest, runtime *dara.RuntimeOptions) (_result *DetachCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DetachCenResponse
func (client *Client) DetachCen(request *DetachCenRequest) (_result *DetachCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachCenResponse{}
	_body, _err := client.DetachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DetachEndUserWithOptions(request *DetachEndUserRequest, runtime *dara.RuntimeOptions) (_result *DetachEndUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DetachEndUserResponse
func (client *Client) DetachEndUser(request *DetachEndUserRequest) (_result *DetachEndUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachEndUserResponse{}
	_body, _err := client.DetachEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisableDesktopsInGroupWithOptions(request *DisableDesktopsInGroupRequest, runtime *dara.RuntimeOptions) (_result *DisableDesktopsInGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisableDesktopsInGroupResponse
func (client *Client) DisableDesktopsInGroup(request *DisableDesktopsInGroupRequest) (_result *DisableDesktopsInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDesktopsInGroupResponse{}
	_body, _err := client.DisableDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisconnectDesktopSessionsWithOptions(request *DisconnectDesktopSessionsRequest, runtime *dara.RuntimeOptions) (_result *DisconnectDesktopSessionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisconnectDesktopSessionsResponse
func (client *Client) DisconnectDesktopSessions(request *DisconnectDesktopSessionsRequest) (_result *DisconnectDesktopSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisconnectDesktopSessionsResponse{}
	_body, _err := client.DisconnectDesktopSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DissociateIpAddressWithOptions(request *DissociateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *DissociateIpAddressResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DissociateIpAddressResponse
func (client *Client) DissociateIpAddress(request *DissociateIpAddressRequest) (_result *DissociateIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateIpAddressResponse{}
	_body, _err := client.DissociateIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DissociateNetworkPackageWithOptions(request *DissociateNetworkPackageRequest, runtime *dara.RuntimeOptions) (_result *DissociateNetworkPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DissociateNetworkPackageResponse
func (client *Client) DissociateNetworkPackage(request *DissociateNetworkPackageRequest) (_result *DissociateNetworkPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateNetworkPackageResponse{}
	_body, _err := client.DissociateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DownloadCdsFileWithOptions(request *DownloadCdsFileRequest, runtime *dara.RuntimeOptions) (_result *DownloadCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DownloadCdsFileResponse
func (client *Client) DownloadCdsFile(request *DownloadCdsFileRequest) (_result *DownloadCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadCdsFileResponse{}
	_body, _err := client.DownloadCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportClientEventsWithOptions(request *ExportClientEventsRequest, runtime *dara.RuntimeOptions) (_result *ExportClientEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportClientEventsResponse
func (client *Client) ExportClientEvents(request *ExportClientEventsRequest) (_result *ExportClientEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportClientEventsResponse{}
	_body, _err := client.ExportClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportDesktopGroupInfoWithOptions(request *ExportDesktopGroupInfoRequest, runtime *dara.RuntimeOptions) (_result *ExportDesktopGroupInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportDesktopGroupInfoResponse
func (client *Client) ExportDesktopGroupInfo(request *ExportDesktopGroupInfoRequest) (_result *ExportDesktopGroupInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportDesktopGroupInfoResponse{}
	_body, _err := client.ExportDesktopGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExportDesktopListInfoWithOptions(request *ExportDesktopListInfoRequest, runtime *dara.RuntimeOptions) (_result *ExportDesktopListInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExportDesktopListInfoResponse
func (client *Client) ExportDesktopListInfo(request *ExportDesktopListInfoRequest) (_result *ExportDesktopListInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportDesktopListInfoResponse{}
	_body, _err := client.ExportDesktopListInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAsyncTaskWithOptions(request *GetAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTask(request *GetAsyncTaskRequest) (_result *GetAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.GetAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *dara.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCoordinateTicketWithOptions(request *GetCoordinateTicketRequest, runtime *dara.RuntimeOptions) (_result *GetCoordinateTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCoordinateTicketResponse
func (client *Client) GetCoordinateTicket(request *GetCoordinateTicketRequest) (_result *GetCoordinateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCoordinateTicketResponse{}
	_body, _err := client.GetCoordinateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDesktopGroupDetailWithOptions(request *GetDesktopGroupDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDesktopGroupDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDesktopGroupDetailResponse
func (client *Client) GetDesktopGroupDetail(request *GetDesktopGroupDetailRequest) (_result *GetDesktopGroupDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.GetDesktopGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOfficeSiteSsoStatusWithOptions(request *GetOfficeSiteSsoStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOfficeSiteSsoStatusResponse
func (client *Client) GetOfficeSiteSsoStatus(request *GetOfficeSiteSsoStatusRequest) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.GetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSpMetadataWithOptions(request *GetSpMetadataRequest, runtime *dara.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSpMetadataResponse
func (client *Client) GetSpMetadata(request *GetSpMetadataRequest) (_result *GetSpMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.GetSpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) HibernateDesktopsWithOptions(request *HibernateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *HibernateDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return HibernateDesktopsResponse
func (client *Client) HibernateDesktops(request *HibernateDesktopsRequest) (_result *HibernateDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HibernateDesktopsResponse{}
	_body, _err := client.HibernateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the files in a cloud disk.
//
// @param tmpReq - ListCdsFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCdsFilesResponse
func (client *Client) ListCdsFilesWithOptions(tmpReq *ListCdsFilesRequest, runtime *dara.RuntimeOptions) (_result *ListCdsFilesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the files in a cloud disk.
//
// @param request - ListCdsFilesRequest
//
// @return ListCdsFilesResponse
func (client *Client) ListCdsFiles(request *ListCdsFilesRequest) (_result *ListCdsFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCdsFilesResponse{}
	_body, _err := client.ListCdsFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDirectoryUsersWithOptions(request *ListDirectoryUsersRequest, runtime *dara.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDirectoryUsersResponse
func (client *Client) ListDirectoryUsers(request *ListDirectoryUsersRequest) (_result *ListDirectoryUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.ListDirectoryUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about shared files of cloud disks.
//
// @param request - ListFilePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilePermissionResponse
func (client *Client) ListFilePermissionWithOptions(request *ListFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *ListFilePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about shared files of cloud disks.
//
// @param request - ListFilePermissionRequest
//
// @return ListFilePermissionResponse
func (client *Client) ListFilePermission(request *ListFilePermissionRequest) (_result *ListFilePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilePermissionResponse{}
	_body, _err := client.ListFilePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListOfficeSiteOverviewWithOptions(request *ListOfficeSiteOverviewRequest, runtime *dara.RuntimeOptions) (_result *ListOfficeSiteOverviewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListOfficeSiteOverviewResponse
func (client *Client) ListOfficeSiteOverview(request *ListOfficeSiteOverviewRequest) (_result *ListOfficeSiteOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.ListOfficeSiteOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListOfficeSiteUsersWithOptions(request *ListOfficeSiteUsersRequest, runtime *dara.RuntimeOptions) (_result *ListOfficeSiteUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListOfficeSiteUsersResponse
func (client *Client) ListOfficeSiteUsers(request *ListOfficeSiteUsersRequest) (_result *ListOfficeSiteUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.ListOfficeSiteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 获取文件下载地址
//
// @param request - ListTransferFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransferFilesResponse
func (client *Client) ListTransferFilesWithOptions(request *ListTransferFilesRequest, runtime *dara.RuntimeOptions) (_result *ListTransferFilesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件下载地址
//
// @param request - ListTransferFilesRequest
//
// @return ListTransferFilesResponse
func (client *Client) ListTransferFiles(request *ListTransferFilesRequest) (_result *ListTransferFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransferFilesResponse{}
	_body, _err := client.ListTransferFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListUserAdOrganizationUnitsWithOptions(request *ListUserAdOrganizationUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListUserAdOrganizationUnitsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListUserAdOrganizationUnitsResponse
func (client *Client) ListUserAdOrganizationUnits(request *ListUserAdOrganizationUnitsRequest) (_result *ListUserAdOrganizationUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserAdOrganizationUnitsResponse{}
	_body, _err := client.ListUserAdOrganizationUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) LockVirtualMFADeviceWithOptions(request *LockVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return LockVirtualMFADeviceResponse
func (client *Client) LockVirtualMFADevice(request *LockVirtualMFADeviceRequest) (_result *LockVirtualMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.LockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MigrateDesktopsWithOptions(request *MigrateDesktopsRequest, runtime *dara.RuntimeOptions) (_result *MigrateDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MigrateDesktopsResponse
func (client *Client) MigrateDesktops(request *MigrateDesktopsRequest) (_result *MigrateDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MigrateDesktopsResponse{}
	_body, _err := client.MigrateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MigrateImageProtocolWithOptions(request *MigrateImageProtocolRequest, runtime *dara.RuntimeOptions) (_result *MigrateImageProtocolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MigrateImageProtocolResponse
func (client *Client) MigrateImageProtocol(request *MigrateImageProtocolRequest) (_result *MigrateImageProtocolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MigrateImageProtocolResponse{}
	_body, _err := client.MigrateImageProtocolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyADConnectorDirectoryWithOptions(request *ModifyADConnectorDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyADConnectorDirectoryResponse
func (client *Client) ModifyADConnectorDirectory(request *ModifyADConnectorDirectoryRequest) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.ModifyADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyADConnectorOfficeSiteWithOptions(request *ModifyADConnectorOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyADConnectorOfficeSiteResponse
func (client *Client) ModifyADConnectorOfficeSite(request *ModifyADConnectorOfficeSiteRequest) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.ModifyADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyAclEntriesWithOptions(request *ModifyAclEntriesRequest, runtime *dara.RuntimeOptions) (_result *ModifyAclEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyAclEntriesResponse
func (client *Client) ModifyAclEntries(request *ModifyAclEntriesRequest) (_result *ModifyAclEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAclEntriesResponse{}
	_body, _err := client.ModifyAclEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyAutoSnapshotPolicyWithOptions(request *ModifyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.ModifyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyBundleWithOptions(request *ModifyBundleRequest, runtime *dara.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyBundleResponse
func (client *Client) ModifyBundle(request *ModifyBundleRequest) (_result *ModifyBundleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBundleResponse{}
	_body, _err := client.ModifyBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the files in a cloud disk.
//
// @param request - ModifyCdsFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdsFileResponse
func (client *Client) ModifyCdsFileWithOptions(request *ModifyCdsFileRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the files in a cloud disk.
//
// @param request - ModifyCdsFileRequest
//
// @return ModifyCdsFileResponse
func (client *Client) ModifyCdsFile(request *ModifyCdsFileRequest) (_result *ModifyCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdsFileResponse{}
	_body, _err := client.ModifyCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCdsFileShareLinkWithOptions(request *ModifyCdsFileShareLinkRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdsFileShareLinkResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCdsFileShareLinkResponse
func (client *Client) ModifyCdsFileShareLink(request *ModifyCdsFileShareLinkRequest) (_result *ModifyCdsFileShareLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdsFileShareLinkResponse{}
	_body, _err := client.ModifyCdsFileShareLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCenterPolicyWithOptions(request *ModifyCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenterPolicyResponse, _err error) {
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

	if !dara.IsNil(request.AutoReconnect) {
		query["AutoReconnect"] = request.AutoReconnect
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCenterPolicyResponse
func (client *Client) ModifyCenterPolicy(request *ModifyCenterPolicyRequest) (_result *ModifyCenterPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCenterPolicyResponse{}
	_body, _err := client.ModifyCenterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCloudDriveGroupsWithOptions(request *ModifyCloudDriveGroupsRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDriveGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCloudDriveGroupsResponse
func (client *Client) ModifyCloudDriveGroups(request *ModifyCloudDriveGroupsRequest) (_result *ModifyCloudDriveGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudDriveGroupsResponse{}
	_body, _err := client.ModifyCloudDriveGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCloudDrivePermissionWithOptions(request *ModifyCloudDrivePermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDrivePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCloudDrivePermissionResponse
func (client *Client) ModifyCloudDrivePermission(request *ModifyCloudDrivePermissionRequest) (_result *ModifyCloudDrivePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudDrivePermissionResponse{}
	_body, _err := client.ModifyCloudDrivePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCloudDriveUsersWithOptions(request *ModifyCloudDriveUsersRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudDriveUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCloudDriveUsersResponse
func (client *Client) ModifyCloudDriveUsers(request *ModifyCloudDriveUsersRequest) (_result *ModifyCloudDriveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudDriveUsersResponse{}
	_body, _err := client.ModifyCloudDriveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyConfigGroupWithOptions(request *ModifyConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyConfigGroupResponse
func (client *Client) ModifyConfigGroup(request *ModifyConfigGroupRequest) (_result *ModifyConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyConfigGroupResponse{}
	_body, _err := client.ModifyConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCustomizedListHeadersWithOptions(request *ModifyCustomizedListHeadersRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomizedListHeadersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCustomizedListHeadersResponse
func (client *Client) ModifyCustomizedListHeaders(request *ModifyCustomizedListHeadersRequest) (_result *ModifyCustomizedListHeadersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomizedListHeadersResponse{}
	_body, _err := client.ModifyCustomizedListHeadersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopChargeTypeWithOptions(request *ModifyDesktopChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopChargeTypeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopChargeTypeResponse
func (client *Client) ModifyDesktopChargeType(request *ModifyDesktopChargeTypeRequest) (_result *ModifyDesktopChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopChargeTypeResponse{}
	_body, _err := client.ModifyDesktopChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopGroupWithOptions(request *ModifyDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopGroupResponse
func (client *Client) ModifyDesktopGroup(request *ModifyDesktopGroupRequest) (_result *ModifyDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.ModifyDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopHostNameWithOptions(request *ModifyDesktopHostNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopHostNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopHostNameResponse
func (client *Client) ModifyDesktopHostName(request *ModifyDesktopHostNameRequest) (_result *ModifyDesktopHostNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopHostNameResponse{}
	_body, _err := client.ModifyDesktopHostNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopNameWithOptions(request *ModifyDesktopNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopNameResponse
func (client *Client) ModifyDesktopName(request *ModifyDesktopNameRequest) (_result *ModifyDesktopNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.ModifyDesktopNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopOversoldGroupWithOptions(request *ModifyDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopOversoldGroupResponse
func (client *Client) ModifyDesktopOversoldGroup(request *ModifyDesktopOversoldGroupRequest) (_result *ModifyDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopOversoldGroupResponse{}
	_body, _err := client.ModifyDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopOversoldGroupSaleWithOptions(request *ModifyDesktopOversoldGroupSaleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldGroupSaleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopOversoldGroupSaleResponse
func (client *Client) ModifyDesktopOversoldGroupSale(request *ModifyDesktopOversoldGroupSaleRequest) (_result *ModifyDesktopOversoldGroupSaleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopOversoldGroupSaleResponse{}
	_body, _err := client.ModifyDesktopOversoldGroupSaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopOversoldUserGroupWithOptions(request *ModifyDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopOversoldUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopOversoldUserGroupResponse
func (client *Client) ModifyDesktopOversoldUserGroup(request *ModifyDesktopOversoldUserGroupRequest) (_result *ModifyDesktopOversoldUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopOversoldUserGroupResponse{}
	_body, _err := client.ModifyDesktopOversoldUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopSpecWithOptions(request *ModifyDesktopSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopSpecResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopSpecResponse
func (client *Client) ModifyDesktopSpec(request *ModifyDesktopSpecRequest) (_result *ModifyDesktopSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.ModifyDesktopSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopTimerWithOptions(request *ModifyDesktopTimerRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopTimerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopTimerResponse
func (client *Client) ModifyDesktopTimer(request *ModifyDesktopTimerRequest) (_result *ModifyDesktopTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopTimerResponse{}
	_body, _err := client.ModifyDesktopTimerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDesktopsPolicyGroupWithOptions(request *ModifyDesktopsPolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDesktopsPolicyGroupResponse
func (client *Client) ModifyDesktopsPolicyGroup(request *ModifyDesktopsPolicyGroupRequest) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.ModifyDesktopsPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyDiskSpecWithOptions(request *ModifyDiskSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskSpecResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyDiskSpecResponse
func (client *Client) ModifyDiskSpec(request *ModifyDiskSpecRequest) (_result *ModifyDiskSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDiskSpecResponse{}
	_body, _err := client.ModifyDiskSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyEntitlementWithOptions(request *ModifyEntitlementRequest, runtime *dara.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyEntitlementResponse
func (client *Client) ModifyEntitlement(request *ModifyEntitlementRequest) (_result *ModifyEntitlementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.ModifyEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyImageAttributeResponse
func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyImagePermissionWithOptions(request *ModifyImagePermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyImagePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyImagePermissionResponse
func (client *Client) ModifyImagePermission(request *ModifyImagePermissionRequest) (_result *ModifyImagePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyImagePermissionResponse{}
	_body, _err := client.ModifyImagePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyNASDefaultMountTargetWithOptions(request *ModifyNASDefaultMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyNASDefaultMountTargetResponse
func (client *Client) ModifyNASDefaultMountTarget(request *ModifyNASDefaultMountTargetRequest) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.ModifyNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyNetworkPackageBandwidthWithOptions(request *ModifyNetworkPackageBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkPackageBandwidthResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyNetworkPackageBandwidthResponse
func (client *Client) ModifyNetworkPackageBandwidth(request *ModifyNetworkPackageBandwidthRequest) (_result *ModifyNetworkPackageBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNetworkPackageBandwidthResponse{}
	_body, _err := client.ModifyNetworkPackageBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyNetworkPackageEnabledWithOptions(request *ModifyNetworkPackageEnabledRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyNetworkPackageEnabledResponse
func (client *Client) ModifyNetworkPackageEnabled(request *ModifyNetworkPackageEnabledRequest) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.ModifyNetworkPackageEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyOfficeSiteAttributeWithOptions(request *ModifyOfficeSiteAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyOfficeSiteAttributeResponse
func (client *Client) ModifyOfficeSiteAttribute(request *ModifyOfficeSiteAttributeRequest) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.ModifyOfficeSiteAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyOfficeSiteCrossDesktopAccessWithOptions(request *ModifyOfficeSiteCrossDesktopAccessRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyOfficeSiteCrossDesktopAccessResponse
func (client *Client) ModifyOfficeSiteCrossDesktopAccess(request *ModifyOfficeSiteCrossDesktopAccessRequest) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.ModifyOfficeSiteCrossDesktopAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改工作区DNS信息
//
// @param request - ModifyOfficeSiteDnsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfficeSiteDnsInfoResponse
func (client *Client) ModifyOfficeSiteDnsInfoWithOptions(request *ModifyOfficeSiteDnsInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteDnsInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改工作区DNS信息
//
// @param request - ModifyOfficeSiteDnsInfoRequest
//
// @return ModifyOfficeSiteDnsInfoResponse
func (client *Client) ModifyOfficeSiteDnsInfo(request *ModifyOfficeSiteDnsInfoRequest) (_result *ModifyOfficeSiteDnsInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOfficeSiteDnsInfoResponse{}
	_body, _err := client.ModifyOfficeSiteDnsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyOfficeSiteMfaEnabledWithOptions(request *ModifyOfficeSiteMfaEnabledRequest, runtime *dara.RuntimeOptions) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyOfficeSiteMfaEnabledResponse
func (client *Client) ModifyOfficeSiteMfaEnabled(request *ModifyOfficeSiteMfaEnabledRequest) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.ModifyOfficeSiteMfaEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyPolicyGroupWithOptions(request *ModifyPolicyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyPolicyGroupResponse
func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyResourceCenterPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceCenterPolicyResponse
func (client *Client) ModifyResourceCenterPolicyWithOptions(request *ModifyResourceCenterPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceCenterPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyResourceCenterPolicyRequest
//
// @return ModifyResourceCenterPolicyResponse
func (client *Client) ModifyResourceCenterPolicy(request *ModifyResourceCenterPolicyRequest) (_result *ModifyResourceCenterPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyResourceCenterPolicyResponse{}
	_body, _err := client.ModifyResourceCenterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifySecurityGroupAttributeWithOptions(request *ModifySecurityGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifySecurityGroupAttributeResponse
func (client *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (_result *ModifySecurityGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityGroupAttributeResponse{}
	_body, _err := client.ModifySecurityGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模板全量更新
//
// @param request - ModifyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplateWithOptions(request *ModifyTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模板全量更新
//
// @param request - ModifyTemplateRequest
//
// @return ModifyTemplateResponse
func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (_result *ModifyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTemplateResponse{}
	_body, _err := client.ModifyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a custom cloud computer template, including the template name and template description.
//
// Description:
//
// You can use this operation to modify only the name and description of a custom cloud computer template. To change other parameters of the template, use the [ModifyTemplate](https://help.aliyun.com/document_detail/2925841.html) operation.
//
// @param request - ModifyTemplateBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateBaseInfoResponse
func (client *Client) ModifyTemplateBaseInfoWithOptions(request *ModifyTemplateBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateBaseInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can use this operation to modify only the name and description of a custom cloud computer template. To change other parameters of the template, use the [ModifyTemplate](https://help.aliyun.com/document_detail/2925841.html) operation.
//
// @param request - ModifyTemplateBaseInfoRequest
//
// @return ModifyTemplateBaseInfoResponse
func (client *Client) ModifyTemplateBaseInfo(request *ModifyTemplateBaseInfoRequest) (_result *ModifyTemplateBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTemplateBaseInfoResponse{}
	_body, _err := client.ModifyTemplateBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyTimerGroupWithOptions(request *ModifyTimerGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyTimerGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyTimerGroupResponse
func (client *Client) ModifyTimerGroup(request *ModifyTimerGroupRequest) (_result *ModifyTimerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTimerGroupResponse{}
	_body, _err := client.ModifyTimerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyUserEntitlementWithOptions(request *ModifyUserEntitlementRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserEntitlementResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyUserEntitlementResponse
func (client *Client) ModifyUserEntitlement(request *ModifyUserEntitlementRequest) (_result *ModifyUserEntitlementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserEntitlementResponse{}
	_body, _err := client.ModifyUserEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyUserToDesktopGroupWithOptions(request *ModifyUserToDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserToDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyUserToDesktopGroupResponse
func (client *Client) ModifyUserToDesktopGroup(request *ModifyUserToDesktopGroupRequest) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.ModifyUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MoveCdsFileWithOptions(request *MoveCdsFileRequest, runtime *dara.RuntimeOptions) (_result *MoveCdsFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MoveCdsFileResponse
func (client *Client) MoveCdsFile(request *MoveCdsFileRequest) (_result *MoveCdsFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveCdsFileResponse{}
	_body, _err := client.MoveCdsFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RebootDesktopsResponse
func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RebuildDesktopsWithOptions(request *RebuildDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RebuildDesktopsResponse
func (client *Client) RebuildDesktops(request *RebuildDesktopsRequest) (_result *RebuildDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.RebuildDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReleaseIpAddressWithOptions(request *ReleaseIpAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleaseIpAddressResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReleaseIpAddressResponse
func (client *Client) ReleaseIpAddress(request *ReleaseIpAddressRequest) (_result *ReleaseIpAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseIpAddressResponse{}
	_body, _err := client.ReleaseIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the file sharing feature of a folder in a cloud disk.
//
// @param tmpReq - RemoveFilePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFilePermissionResponse
func (client *Client) RemoveFilePermissionWithOptions(tmpReq *RemoveFilePermissionRequest, runtime *dara.RuntimeOptions) (_result *RemoveFilePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the file sharing feature of a folder in a cloud disk.
//
// @param request - RemoveFilePermissionRequest
//
// @return RemoveFilePermissionResponse
func (client *Client) RemoveFilePermission(request *RemoveFilePermissionRequest) (_result *RemoveFilePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveFilePermissionResponse{}
	_body, _err := client.RemoveFilePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveUserFromDesktopGroupWithOptions(request *RemoveUserFromDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveUserFromDesktopGroupResponse
func (client *Client) RemoveUserFromDesktopGroup(request *RemoveUserFromDesktopGroupRequest) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.RemoveUserFromDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveUserFromDesktopOversoldUserGroupWithOptions(request *RemoveUserFromDesktopOversoldUserGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromDesktopOversoldUserGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveUserFromDesktopOversoldUserGroupResponse
func (client *Client) RemoveUserFromDesktopOversoldUserGroup(request *RemoveUserFromDesktopOversoldUserGroupRequest) (_result *RemoveUserFromDesktopOversoldUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromDesktopOversoldUserGroupResponse{}
	_body, _err := client.RemoveUserFromDesktopOversoldUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RenewDesktopGroupWithOptions(request *RenewDesktopGroupRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RenewDesktopGroupResponse
func (client *Client) RenewDesktopGroup(request *RenewDesktopGroupRequest) (_result *RenewDesktopGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewDesktopGroupResponse{}
	_body, _err := client.RenewDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RenewDesktopOversoldGroupWithOptions(request *RenewDesktopOversoldGroupRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopOversoldGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RenewDesktopOversoldGroupResponse
func (client *Client) RenewDesktopOversoldGroup(request *RenewDesktopOversoldGroupRequest) (_result *RenewDesktopOversoldGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewDesktopOversoldGroupResponse{}
	_body, _err := client.RenewDesktopOversoldGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RenewDesktopsWithOptions(request *RenewDesktopsRequest, runtime *dara.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RenewDesktopsResponse
func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (_result *RenewDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.RenewDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RenewNetworkPackagesWithOptions(request *RenewNetworkPackagesRequest, runtime *dara.RuntimeOptions) (_result *RenewNetworkPackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RenewNetworkPackagesResponse
func (client *Client) RenewNetworkPackages(request *RenewNetworkPackagesRequest) (_result *RenewNetworkPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewNetworkPackagesResponse{}
	_body, _err := client.RenewNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetDesktopsWithOptions(request *ResetDesktopsRequest, runtime *dara.RuntimeOptions) (_result *ResetDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ResetDesktopsResponse
func (client *Client) ResetDesktops(request *ResetDesktopsRequest) (_result *ResetDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetDesktopsResponse{}
	_body, _err := client.ResetDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetNASDefaultMountTargetWithOptions(request *ResetNASDefaultMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ResetNASDefaultMountTargetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ResetNASDefaultMountTargetResponse
func (client *Client) ResetNASDefaultMountTarget(request *ResetNASDefaultMountTargetRequest) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.ResetNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *dara.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ResetSnapshotResponse
func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RevokeCoordinatePrivilegeWithOptions(request *RevokeCoordinatePrivilegeRequest, runtime *dara.RuntimeOptions) (_result *RevokeCoordinatePrivilegeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RevokeCoordinatePrivilegeResponse
func (client *Client) RevokeCoordinatePrivilege(request *RevokeCoordinatePrivilegeRequest) (_result *RevokeCoordinatePrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeCoordinatePrivilegeResponse{}
	_body, _err := client.RevokeCoordinatePrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *dara.RuntimeOptions) (_result *RunCommandResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendVerifyCodeWithOptions(request *SendVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *SendVerifyCodeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SendVerifyCodeResponse
func (client *Client) SendVerifyCode(request *SendVerifyCodeRequest) (_result *SendVerifyCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendVerifyCodeResponse{}
	_body, _err := client.SendVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDesktopGroupScaleTimerWithOptions(request *SetDesktopGroupScaleTimerRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupScaleTimerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDesktopGroupScaleTimerResponse
func (client *Client) SetDesktopGroupScaleTimer(request *SetDesktopGroupScaleTimerRequest) (_result *SetDesktopGroupScaleTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDesktopGroupScaleTimerResponse{}
	_body, _err := client.SetDesktopGroupScaleTimerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDesktopGroupTimerWithOptions(request *SetDesktopGroupTimerRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupTimerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDesktopGroupTimerResponse
func (client *Client) SetDesktopGroupTimer(request *SetDesktopGroupTimerRequest) (_result *SetDesktopGroupTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDesktopGroupTimerResponse{}
	_body, _err := client.SetDesktopGroupTimerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDesktopGroupTimerStatusWithOptions(request *SetDesktopGroupTimerStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopGroupTimerStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDesktopGroupTimerStatusResponse
func (client *Client) SetDesktopGroupTimerStatus(request *SetDesktopGroupTimerStatusRequest) (_result *SetDesktopGroupTimerStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDesktopGroupTimerStatusResponse{}
	_body, _err := client.SetDesktopGroupTimerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置桌面维护模式
//
// @param request - SetDesktopMaintenanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDesktopMaintenanceResponse
func (client *Client) SetDesktopMaintenanceWithOptions(request *SetDesktopMaintenanceRequest, runtime *dara.RuntimeOptions) (_result *SetDesktopMaintenanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置桌面维护模式
//
// @param request - SetDesktopMaintenanceRequest
//
// @return SetDesktopMaintenanceResponse
func (client *Client) SetDesktopMaintenance(request *SetDesktopMaintenanceRequest) (_result *SetDesktopMaintenanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDesktopMaintenanceResponse{}
	_body, _err := client.SetDesktopMaintenanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetDirectorySsoStatusWithOptions(request *SetDirectorySsoStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetDirectorySsoStatusResponse
func (client *Client) SetDirectorySsoStatus(request *SetDirectorySsoStatusRequest) (_result *SetDirectorySsoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.SetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetIdpMetadataWithOptions(request *SetIdpMetadataRequest, runtime *dara.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetIdpMetadataResponse
func (client *Client) SetIdpMetadata(request *SetIdpMetadataRequest) (_result *SetIdpMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.SetIdpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetOfficeSiteSsoStatusWithOptions(request *SetOfficeSiteSsoStatusRequest, runtime *dara.RuntimeOptions) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetOfficeSiteSsoStatusResponse
func (client *Client) SetOfficeSiteSsoStatus(request *SetOfficeSiteSsoStatusRequest) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.SetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetUserProfilePathRulesWithOptions(tmpReq *SetUserProfilePathRulesRequest, runtime *dara.RuntimeOptions) (_result *SetUserProfilePathRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SetUserProfilePathRulesRequest
//
// @return SetUserProfilePathRulesResponse
func (client *Client) SetUserProfilePathRules(request *SetUserProfilePathRulesRequest) (_result *SetUserProfilePathRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetUserProfilePathRulesResponse{}
	_body, _err := client.SetUserProfilePathRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StartDesktopsResponse
func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *dara.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopDesktopsResponse
func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopInvocationWithOptions(request *StopInvocationRequest, runtime *dara.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopInvocationResponse
func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 文件传输审批回调
//
// @param request - TransferTaskApprovalCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferTaskApprovalCallbackResponse
func (client *Client) TransferTaskApprovalCallbackWithOptions(request *TransferTaskApprovalCallbackRequest, runtime *dara.RuntimeOptions) (_result *TransferTaskApprovalCallbackResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件传输审批回调
//
// @param request - TransferTaskApprovalCallbackRequest
//
// @return TransferTaskApprovalCallbackResponse
func (client *Client) TransferTaskApprovalCallback(request *TransferTaskApprovalCallbackRequest) (_result *TransferTaskApprovalCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferTaskApprovalCallbackResponse{}
	_body, _err := client.TransferTaskApprovalCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnbindConfigGroupWithOptions(request *UnbindConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *UnbindConfigGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnbindConfigGroupResponse
func (client *Client) UnbindConfigGroup(request *UnbindConfigGroupRequest) (_result *UnbindConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindConfigGroupResponse{}
	_body, _err := client.UnbindConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnbindUserDesktopWithOptions(request *UnbindUserDesktopRequest, runtime *dara.RuntimeOptions) (_result *UnbindUserDesktopResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnbindUserDesktopResponse
func (client *Client) UnbindUserDesktop(request *UnbindUserDesktopRequest) (_result *UnbindUserDesktopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.UnbindUserDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnlockVirtualMFADeviceWithOptions(request *UnlockVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnlockVirtualMFADeviceResponse
func (client *Client) UnlockVirtualMFADevice(request *UnlockVirtualMFADeviceRequest) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.UnlockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	_body, _err := client.CallApi(params, req, runtime)
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
func (client *Client) UpdateFotaTaskWithOptions(request *UpdateFotaTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateFotaTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFotaTaskResponse
func (client *Client) UpdateFotaTask(request *UpdateFotaTaskRequest) (_result *UpdateFotaTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFotaTaskResponse{}
	_body, _err := client.UpdateFotaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UploadImageWithOptions(request *UploadImageRequest, runtime *dara.RuntimeOptions) (_result *UploadImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UploadImageResponse
func (client *Client) UploadImage(request *UploadImageRequest) (_result *UploadImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadImageResponse{}
	_body, _err := client.UploadImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VerifyCenWithOptions(request *VerifyCenRequest, runtime *dara.RuntimeOptions) (_result *VerifyCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VerifyCenResponse
func (client *Client) VerifyCen(request *VerifyCenRequest) (_result *VerifyCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyCenResponse{}
	_body, _err := client.VerifyCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) WakeupDesktopsWithOptions(request *WakeupDesktopsRequest, runtime *dara.RuntimeOptions) (_result *WakeupDesktopsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return WakeupDesktopsResponse
func (client *Client) WakeupDesktops(request *WakeupDesktopsRequest) (_result *WakeupDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WakeupDesktopsResponse{}
	_body, _err := client.WakeupDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
