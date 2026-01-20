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
	client.Endpoint, _err = client.GetEndpoint(dara.String("resourcesharing"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Accepts a resource sharing invitation.
//
// Description:
//
// ### [](#)
//
//   - A principal needs to accept or reject a resource sharing invitation only if the principal is not the management account or a member of a resource directory. If you share resources with an object in a resource directory, the system automatically accepts the resource sharing invitation for the object.
//
//   - A resource sharing invitation is valid for seven days. A principal must accept or reject the invitation within the validity period.
//
// This topic provides an example on how to call the API operation to accept the resource sharing invitation whose ID is `i-pMnItMX19fBJ****` in the `cn-hangzhou` region.
//
// @param request - AcceptResourceShareInvitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptResourceShareInvitationResponse
func (client *Client) AcceptResourceShareInvitationWithOptions(request *AcceptResourceShareInvitationRequest, runtime *dara.RuntimeOptions) (_result *AcceptResourceShareInvitationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceShareInvitationId) {
		query["ResourceShareInvitationId"] = request.ResourceShareInvitationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptResourceShareInvitation"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptResourceShareInvitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts a resource sharing invitation.
//
// Description:
//
// ### [](#)
//
//   - A principal needs to accept or reject a resource sharing invitation only if the principal is not the management account or a member of a resource directory. If you share resources with an object in a resource directory, the system automatically accepts the resource sharing invitation for the object.
//
//   - A resource sharing invitation is valid for seven days. A principal must accept or reject the invitation within the validity period.
//
// This topic provides an example on how to call the API operation to accept the resource sharing invitation whose ID is `i-pMnItMX19fBJ****` in the `cn-hangzhou` region.
//
// @param request - AcceptResourceShareInvitationRequest
//
// @return AcceptResourceShareInvitationResponse
func (client *Client) AcceptResourceShareInvitation(request *AcceptResourceShareInvitationRequest) (_result *AcceptResourceShareInvitationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptResourceShareInvitationResponse{}
	_body, _err := client.AcceptResourceShareInvitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates resources or principals with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to associate the vSwitch `vsw-bp183p93qs667muql****` and the member `172050525300****` with the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the association, the vSwitch is shared with the member.
//
// @param request - AssociateResourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResourceShareResponse
func (client *Client) AssociateResourceShareWithOptions(request *AssociateResourceShareRequest, runtime *dara.RuntimeOptions) (_result *AssociateResourceShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PermissionNames) {
		query["PermissionNames"] = request.PermissionNames
	}

	if !dara.IsNil(request.ResourceArns) {
		query["ResourceArns"] = request.ResourceArns
	}

	if !dara.IsNil(request.ResourceProperties) {
		query["ResourceProperties"] = request.ResourceProperties
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.TargetProperties) {
		query["TargetProperties"] = request.TargetProperties
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateResourceShare"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates resources or principals with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to associate the vSwitch `vsw-bp183p93qs667muql****` and the member `172050525300****` with the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the association, the vSwitch is shared with the member.
//
// @param request - AssociateResourceShareRequest
//
// @return AssociateResourceShareResponse
func (client *Client) AssociateResourceShare(request *AssociateResourceShareRequest) (_result *AssociateResourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateResourceShareResponse{}
	_body, _err := client.AssociateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates permissions with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to associate the `AliyunRSDefaultPermissionVSwitch` permission with the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
//
// @param request - AssociateResourceSharePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResourceSharePermissionResponse
func (client *Client) AssociateResourceSharePermissionWithOptions(request *AssociateResourceSharePermissionRequest, runtime *dara.RuntimeOptions) (_result *AssociateResourceSharePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PermissionName) {
		query["PermissionName"] = request.PermissionName
	}

	if !dara.IsNil(request.Replace) {
		query["Replace"] = request.Replace
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateResourceSharePermission"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateResourceSharePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates permissions with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to associate the `AliyunRSDefaultPermissionVSwitch` permission with the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
//
// @param request - AssociateResourceSharePermissionRequest
//
// @return AssociateResourceSharePermissionResponse
func (client *Client) AssociateResourceSharePermission(request *AssociateResourceSharePermissionRequest) (_result *AssociateResourceSharePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateResourceSharePermissionResponse{}
	_body, _err := client.AssociateResourceSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfers a resource share from one resource group to another.
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
		Version:     dara.String("2020-01-10"),
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
// Transfers a resource share from one resource group to another.
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
// Checks the status of resource sharing within a resource directory.
//
// @param request - CheckSharingWithResourceDirectoryStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSharingWithResourceDirectoryStatusResponse
func (client *Client) CheckSharingWithResourceDirectoryStatusWithOptions(runtime *dara.RuntimeOptions) (_result *CheckSharingWithResourceDirectoryStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSharingWithResourceDirectoryStatus"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSharingWithResourceDirectoryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the status of resource sharing within a resource directory.
//
// @return CheckSharingWithResourceDirectoryStatusResponse
func (client *Client) CheckSharingWithResourceDirectoryStatus() (_result *CheckSharingWithResourceDirectoryStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSharingWithResourceDirectoryStatusResponse{}
	_body, _err := client.CheckSharingWithResourceDirectoryStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource share.
//
// Description:
//
// Resource Sharing allows you to share your resources with one or more accounts and access the resources shared by other accounts. For more information, see [Resource Sharing overview](https://help.aliyun.com/document_detail/160622.html).
//
// This topic provides an example on how to call the API operation to create a resource share named `test` in the `cn-hangzhou` region to share the vSwitch `vsw-bp183p93qs667muql****` with the member `172050525300****` in a resource directory. In this example, the management account of the resource directory is used to call this API operation.
//
// @param request - CreateResourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceShareResponse
func (client *Client) CreateResourceShareWithOptions(request *CreateResourceShareRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowExternalTargets) {
		query["AllowExternalTargets"] = request.AllowExternalTargets
	}

	if !dara.IsNil(request.PermissionNames) {
		query["PermissionNames"] = request.PermissionNames
	}

	if !dara.IsNil(request.ResourceArns) {
		query["ResourceArns"] = request.ResourceArns
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceProperties) {
		query["ResourceProperties"] = request.ResourceProperties
	}

	if !dara.IsNil(request.ResourceShareName) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetProperties) {
		query["TargetProperties"] = request.TargetProperties
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceShare"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource share.
//
// Description:
//
// Resource Sharing allows you to share your resources with one or more accounts and access the resources shared by other accounts. For more information, see [Resource Sharing overview](https://help.aliyun.com/document_detail/160622.html).
//
// This topic provides an example on how to call the API operation to create a resource share named `test` in the `cn-hangzhou` region to share the vSwitch `vsw-bp183p93qs667muql****` with the member `172050525300****` in a resource directory. In this example, the management account of the resource directory is used to call this API operation.
//
// @param request - CreateResourceShareRequest
//
// @return CreateResourceShareResponse
func (client *Client) CreateResourceShare(request *CreateResourceShareRequest) (_result *CreateResourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceShareResponse{}
	_body, _err := client.CreateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource share.
//
// Description:
//
// After a resource share is deleted, all principals in the resource share can no longer access the resources in the resource share. However, the resources are not deleted with the resource share.
//
// A resource share that is deleted is in the `Deleted` state. The system deletes the record of the resource share within 48 hours to 96 hours.
//
// This topic provides an example on how to call the API operation to delete the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region.
//
// @param request - DeleteResourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceShareResponse
func (client *Client) DeleteResourceShareWithOptions(request *DeleteResourceShareRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceShare"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource share.
//
// Description:
//
// After a resource share is deleted, all principals in the resource share can no longer access the resources in the resource share. However, the resources are not deleted with the resource share.
//
// A resource share that is deleted is in the `Deleted` state. The system deletes the record of the resource share within 48 hours to 96 hours.
//
// This topic provides an example on how to call the API operation to delete the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region.
//
// @param request - DeleteResourceShareRequest
//
// @return DeleteResourceShareResponse
func (client *Client) DeleteResourceShare(request *DeleteResourceShareRequest) (_result *DeleteResourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceShareResponse{}
	_body, _err := client.DeleteResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions where the Resource Sharing service is available.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-01-10"),
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
// Queries the regions where the Resource Sharing service is available.
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
// Disassociates resources or principals from a resource share.
//
// Description:
//
//	  A resource owner can call this API operation to disassociate shared resources or principals from a resource share.
//
//		- If a principal does not belong to a resource directory, the principal can call this API operation to exit the resource share. For more information, see [Exit a resource share](https://help.aliyun.com/document_detail/440614.html).
//
// This topic provides an example on how to use the management account of a resource directory to call the API operation to disassociate the member `172050525300****` from the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the member is disassociated from the resource share, the member cannot share the resources in the resource share.
//
// @param request - DisassociateResourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateResourceShareResponse
func (client *Client) DisassociateResourceShareWithOptions(request *DisassociateResourceShareRequest, runtime *dara.RuntimeOptions) (_result *DisassociateResourceShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceArns) {
		query["ResourceArns"] = request.ResourceArns
	}

	if !dara.IsNil(request.ResourceOwner) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateResourceShare"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resources or principals from a resource share.
//
// Description:
//
//	  A resource owner can call this API operation to disassociate shared resources or principals from a resource share.
//
//		- If a principal does not belong to a resource directory, the principal can call this API operation to exit the resource share. For more information, see [Exit a resource share](https://help.aliyun.com/document_detail/440614.html).
//
// This topic provides an example on how to use the management account of a resource directory to call the API operation to disassociate the member `172050525300****` from the resource share `rs-6GRmdD3X****` in the `cn-hangzhou` region. After the member is disassociated from the resource share, the member cannot share the resources in the resource share.
//
// @param request - DisassociateResourceShareRequest
//
// @return DisassociateResourceShareResponse
func (client *Client) DisassociateResourceShare(request *DisassociateResourceShareRequest) (_result *DisassociateResourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateResourceShareResponse{}
	_body, _err := client.DisassociateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a permission from a resource share. You can disassociate a permission from a resource share only if the resource share does not contain resources of the type indicated by the permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to disassociate the `AliyunRSDefaultPermissionVSwitch` permission from the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
//
// @param request - DisassociateResourceSharePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateResourceSharePermissionResponse
func (client *Client) DisassociateResourceSharePermissionWithOptions(request *DisassociateResourceSharePermissionRequest, runtime *dara.RuntimeOptions) (_result *DisassociateResourceSharePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PermissionName) {
		query["PermissionName"] = request.PermissionName
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateResourceSharePermission"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateResourceSharePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a permission from a resource share. You can disassociate a permission from a resource share only if the resource share does not contain resources of the type indicated by the permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to disassociate the `AliyunRSDefaultPermissionVSwitch` permission from the `rs-6GRmdD3X****` resource share in the `cn-hangzhou` region.
//
// @param request - DisassociateResourceSharePermissionRequest
//
// @return DisassociateResourceSharePermissionResponse
func (client *Client) DisassociateResourceSharePermission(request *DisassociateResourceSharePermissionRequest) (_result *DisassociateResourceSharePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateResourceSharePermissionResponse{}
	_body, _err := client.DisassociateResourceSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables resource sharing for a resource directory.
//
// Description:
//
// You can share your resources with all members in your resource directory, all members in a specific folder in your resource directory, or a specific member in your resource directory as a resource owner only after you enable resource sharing for your resource directory.
//
// You can call this API operation only by using the management account of your resource directory or a RAM user or RAM role to which the required permissions are granted within the management account.
//
// @param request - EnableSharingWithResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSharingWithResourceDirectoryResponse
func (client *Client) EnableSharingWithResourceDirectoryWithOptions(runtime *dara.RuntimeOptions) (_result *EnableSharingWithResourceDirectoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSharingWithResourceDirectory"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSharingWithResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables resource sharing for a resource directory.
//
// Description:
//
// You can share your resources with all members in your resource directory, all members in a specific folder in your resource directory, or a specific member in your resource directory as a resource owner only after you enable resource sharing for your resource directory.
//
// You can call this API operation only by using the management account of your resource directory or a RAM user or RAM role to which the required permissions are granted within the management account.
//
// @return EnableSharingWithResourceDirectoryResponse
func (client *Client) EnableSharingWithResourceDirectory() (_result *EnableSharingWithResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSharingWithResourceDirectoryResponse{}
	_body, _err := client.EnableSharingWithResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the information about the `AliyunRSDefaultPermissionVSwitch` permission whose version is `v1` in the `cn-hangzhou` region.
//
// @param request - GetPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionResponse
func (client *Client) GetPermissionWithOptions(request *GetPermissionRequest, runtime *dara.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PermissionName) {
		query["PermissionName"] = request.PermissionName
	}

	if !dara.IsNil(request.PermissionVersion) {
		query["PermissionVersion"] = request.PermissionVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPermission"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the information about a permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the information about the `AliyunRSDefaultPermissionVSwitch` permission whose version is `v1` in the `cn-hangzhou` region.
//
// @param request - GetPermissionRequest
//
// @return GetPermissionResponse
func (client *Client) GetPermission(request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the versions of the `AliyunRSDefaultPermissionVSwitch` permission in the `cn-hangzhou` region.
//
// @param request - ListPermissionVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionVersionsResponse
func (client *Client) ListPermissionVersionsWithOptions(request *ListPermissionVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListPermissionVersionsResponse, _err error) {
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

	if !dara.IsNil(request.PermissionName) {
		query["PermissionName"] = request.PermissionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionVersions"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the versions of the `AliyunRSDefaultPermissionVSwitch` permission in the `cn-hangzhou` region.
//
// @param request - ListPermissionVersionsRequest
//
// @return ListPermissionVersionsResponse
func (client *Client) ListPermissionVersions(request *ListPermissionVersionsRequest) (_result *ListPermissionVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPermissionVersionsResponse{}
	_body, _err := client.ListPermissionVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the default permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the information about the default permission for the `VSwitch` resource type in the `cn-hangzhou` region.
//
// @param request - ListPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithOptions(request *ListPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissions"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the information about the default permission.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the information about the default permission for the `VSwitch` resource type in the `cn-hangzhou` region.
//
// @param request - ListPermissionsRequest
//
// @return ListPermissionsResponse
func (client *Client) ListPermissions(request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association records of resource shares.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the association records of the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows the following records:
//
//   - The resource `vsw-bp1upw03qyz8n7us9****` of the `VSwitch` type has been associated with the resource share `rs-6GRmdD3X****`. The resource is in the `Associated` state. This indicates that the resource is being shared.
//
//   - The resource `vsw-bp183p93qs667muql****` of the `VSwitch` type has been disassociated from the resource share `rs-6GRmdD3X****`. The resource is in the `Disassociated` state. This indicates that the sharing of the resource is stopped.
//
// @param request - ListResourceShareAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceShareAssociationsResponse
func (client *Client) ListResourceShareAssociationsWithOptions(request *ListResourceShareAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceShareAssociationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationStatus) {
		query["AssociationStatus"] = request.AssociationStatus
	}

	if !dara.IsNil(request.AssociationType) {
		query["AssociationType"] = request.AssociationType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceShareIds) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceShareAssociations"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceShareAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association records of resource shares.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the association records of the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows the following records:
//
//   - The resource `vsw-bp1upw03qyz8n7us9****` of the `VSwitch` type has been associated with the resource share `rs-6GRmdD3X****`. The resource is in the `Associated` state. This indicates that the resource is being shared.
//
//   - The resource `vsw-bp183p93qs667muql****` of the `VSwitch` type has been disassociated from the resource share `rs-6GRmdD3X****`. The resource is in the `Disassociated` state. This indicates that the sharing of the resource is stopped.
//
// @param request - ListResourceShareAssociationsRequest
//
// @return ListResourceShareAssociationsResponse
func (client *Client) ListResourceShareAssociations(request *ListResourceShareAssociationsRequest) (_result *ListResourceShareAssociationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceShareAssociationsResponse{}
	_body, _err := client.ListResourceShareAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource sharing invitations that are received.
//
// Description:
//
// ### [](#)
//
// This topic provides an example on how to call the API operation to query the resource sharing invitations that are received by the current account in the `cn-hangzhou` region. The response shows that one invitation is received by the current account and is waiting for confirmation.
//
// @param request - ListResourceShareInvitationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceShareInvitationsResponse
func (client *Client) ListResourceShareInvitationsWithOptions(request *ListResourceShareInvitationsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceShareInvitationsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceShareIds) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !dara.IsNil(request.ResourceShareInvitationIds) {
		query["ResourceShareInvitationIds"] = request.ResourceShareInvitationIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceShareInvitations"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceShareInvitationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource sharing invitations that are received.
//
// Description:
//
// ### [](#)
//
// This topic provides an example on how to call the API operation to query the resource sharing invitations that are received by the current account in the `cn-hangzhou` region. The response shows that one invitation is received by the current account and is waiting for confirmation.
//
// @param request - ListResourceShareInvitationsRequest
//
// @return ListResourceShareInvitationsResponse
func (client *Client) ListResourceShareInvitations(request *ListResourceShareInvitationsRequest) (_result *ListResourceShareInvitationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceShareInvitationsResponse{}
	_body, _err := client.ListResourceShareInvitationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions that are associated with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the permissions that are associated with the resource share created by using the current Alibaba Cloud account in the `cn-hangzhou` region.
//
// @param request - ListResourceSharePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceSharePermissionsResponse
func (client *Client) ListResourceSharePermissionsWithOptions(request *ListResourceSharePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceSharePermissionsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwner) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceSharePermissions"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceSharePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions that are associated with a resource share.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the permissions that are associated with the resource share created by using the current Alibaba Cloud account in the `cn-hangzhou` region.
//
// @param request - ListResourceSharePermissionsRequest
//
// @return ListResourceSharePermissionsResponse
func (client *Client) ListResourceSharePermissions(request *ListResourceSharePermissionsRequest) (_result *ListResourceSharePermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceSharePermissionsResponse{}
	_body, _err := client.ListResourceSharePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource shares.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows that the following resource shares are created within the account `151266687691****`:
//
//   - `rs-hX9wC5jO****`, which is in the `Deleted` state
//
//   - `rs-PqysnzIj****`, which is in the `Active` state
//
// @param request - ListResourceSharesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceSharesResponse
func (client *Client) ListResourceSharesWithOptions(request *ListResourceSharesRequest, runtime *dara.RuntimeOptions) (_result *ListResourceSharesResponse, _err error) {
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

	if !dara.IsNil(request.PermissionName) {
		query["PermissionName"] = request.PermissionName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwner) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !dara.IsNil(request.ResourceShareIds) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !dara.IsNil(request.ResourceShareName) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	if !dara.IsNil(request.ResourceShareStatus) {
		query["ResourceShareStatus"] = request.ResourceShareStatus
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceShares"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceSharesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource shares.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the resource shares that are created by using the current Alibaba Cloud account in the `cn-hangzhou` region. The response shows that the following resource shares are created within the account `151266687691****`:
//
//   - `rs-hX9wC5jO****`, which is in the `Deleted` state
//
//   - `rs-PqysnzIj****`, which is in the `Active` state
//
// @param request - ListResourceSharesRequest
//
// @return ListResourceSharesResponse
func (client *Client) ListResourceShares(request *ListResourceSharesRequest) (_result *ListResourceSharesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceSharesResponse{}
	_body, _err := client.ListResourceSharesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources you share with other accounts or the resources other accounts share with you.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the resources that you share with other accounts in the `cn-hangzhou` region. The response shows that in the resource share `rs-6GRmdD3X****`, you share the `vsw-bp1upw03qyz8n7us9****` resource of the `VSwitch` type with other accounts.
//
// @param request - ListSharedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedResourcesResponse
func (client *Client) ListSharedResourcesWithOptions(request *ListSharedResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListSharedResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceArns) {
		query["ResourceArns"] = request.ResourceArns
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceOwner) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !dara.IsNil(request.ResourceShareIds) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSharedResources"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSharedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources you share with other accounts or the resources other accounts share with you.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the resources that you share with other accounts in the `cn-hangzhou` region. The response shows that in the resource share `rs-6GRmdD3X****`, you share the `vsw-bp1upw03qyz8n7us9****` resource of the `VSwitch` type with other accounts.
//
// @param request - ListSharedResourcesRequest
//
// @return ListSharedResourcesResponse
func (client *Client) ListSharedResources(request *ListSharedResourcesRequest) (_result *ListSharedResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSharedResourcesResponse{}
	_body, _err := client.ListSharedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries principals.
//
// Description:
//
// If you are a resource owner, you can query the principals with which you share your resources. If you are a principal, you can query the resources that are shared with you.
//
// This topic provides an example on how to call the API operation to query the principals with which you share your resources in the `cn-hangzhou` region. The response shows that you share your resources with the principals `114240524784****` and `172050525300****`.
//
// @param request - ListSharedTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedTargetsResponse
func (client *Client) ListSharedTargetsWithOptions(request *ListSharedTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListSharedTargetsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwner) {
		query["ResourceOwner"] = request.ResourceOwner
	}

	if !dara.IsNil(request.ResourceShareIds) {
		query["ResourceShareIds"] = request.ResourceShareIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSharedTargets"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSharedTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries principals.
//
// Description:
//
// If you are a resource owner, you can query the principals with which you share your resources. If you are a principal, you can query the resources that are shared with you.
//
// This topic provides an example on how to call the API operation to query the principals with which you share your resources in the `cn-hangzhou` region. The response shows that you share your resources with the principals `114240524784****` and `172050525300****`.
//
// @param request - ListSharedTargetsRequest
//
// @return ListSharedTargetsResponse
func (client *Client) ListSharedTargets(request *ListSharedTargetsRequest) (_result *ListSharedTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSharedTargetsResponse{}
	_body, _err := client.ListSharedTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resource shares.
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
		Version:     dara.String("2020-01-10"),
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
// Queries the tags that are added to resource shares.
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
// 拒绝组织外共享邀请
//
// Description:
//
// This topic provides an example on how to call the API operation to reject the resource sharing invitation `i-yyTWbkjHArYh****` in the `cn-hangzhou` region.
//
// @param request - RejectResourceShareInvitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectResourceShareInvitationResponse
func (client *Client) RejectResourceShareInvitationWithOptions(request *RejectResourceShareInvitationRequest, runtime *dara.RuntimeOptions) (_result *RejectResourceShareInvitationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceShareInvitationId) {
		query["ResourceShareInvitationId"] = request.ResourceShareInvitationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectResourceShareInvitation"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectResourceShareInvitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拒绝组织外共享邀请
//
// Description:
//
// This topic provides an example on how to call the API operation to reject the resource sharing invitation `i-yyTWbkjHArYh****` in the `cn-hangzhou` region.
//
// @param request - RejectResourceShareInvitationRequest
//
// @return RejectResourceShareInvitationResponse
func (client *Client) RejectResourceShareInvitation(request *RejectResourceShareInvitationRequest) (_result *RejectResourceShareInvitationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectResourceShareInvitationResponse{}
	_body, _err := client.RejectResourceShareInvitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to a resource share.
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
		Version:     dara.String("2020-01-10"),
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
// Adds tags to a resource share.
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
// Removes tags from resource shares.
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
		Version:     dara.String("2020-01-10"),
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
// Removes tags from resource shares.
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
// 调用UpdateResourceShare修改共享单元基本信息。
//
// Description:
//
// You can call this API operation to change the name or resource sharing scope of a resource share.
//
// This topic provides an example on how to call the API operation to change the name of the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region from `test` to `new`.
//
// @param request - UpdateResourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceShareResponse
func (client *Client) UpdateResourceShareWithOptions(request *UpdateResourceShareRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowExternalTargets) {
		query["AllowExternalTargets"] = request.AllowExternalTargets
	}

	if !dara.IsNil(request.ResourceShareId) {
		query["ResourceShareId"] = request.ResourceShareId
	}

	if !dara.IsNil(request.ResourceShareName) {
		query["ResourceShareName"] = request.ResourceShareName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceShare"),
		Version:     dara.String("2020-01-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用UpdateResourceShare修改共享单元基本信息。
//
// Description:
//
// You can call this API operation to change the name or resource sharing scope of a resource share.
//
// This topic provides an example on how to call the API operation to change the name of the resource share `rs-qSkW1HBY****` in the `cn-hangzhou` region from `test` to `new`.
//
// @param request - UpdateResourceShareRequest
//
// @return UpdateResourceShareResponse
func (client *Client) UpdateResourceShare(request *UpdateResourceShareRequest) (_result *UpdateResourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceShareResponse{}
	_body, _err := client.UpdateResourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
