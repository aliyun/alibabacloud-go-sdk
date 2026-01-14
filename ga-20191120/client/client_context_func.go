// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds IP addresses or CIDR blocks to an access control list (ACL). You can add IP addresses or CIDR blocks to an ACL and configure a whitelist or blacklist to allow or deny requests from clients.
//
// Description:
//
//	  **AddEntriesToAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) operation to query the status of the ACL to which you want to add IP entries.
//
//	    	- If the ACL is in the **configuring*	- state, it indicates that IP entries are added to the ACL. In this case, you can perform only query operations.
//
//	    	- If the ACL is in the **active*	- state, it indicates that IP entries are added to the ACL.
//
//		- The **AddEntriesToAcl*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - AddEntriesToAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEntriesToAclResponse
func (client *Client) AddEntriesToAclWithContext(ctx context.Context, request *AddEntriesToAclRequest, runtime *dara.RuntimeOptions) (_result *AddEntriesToAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntries) {
		query["AclEntries"] = request.AclEntries
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEntriesToAcl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates access control lists (ACLs) with a listener.
//
// Description:
//
// ## Description
//
//   - **AssociateAclsWithListener*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the state of the listener with which you attempt to associate an ACL.
//
//   - If the listener is in the **updating*	- state, it indicates that the ACL is being associated. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, it indicates that the ACL is associated.
//
//   - The **AssociateAclsWithListener*	- operation cannot be called repeatedly for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - AssociateAclsWithListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateAclsWithListenerResponse
func (client *Client) AssociateAclsWithListenerWithContext(ctx context.Context, request *AssociateAclsWithListenerRequest, runtime *dara.RuntimeOptions) (_result *AssociateAclsWithListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclIds) {
		query["AclIds"] = request.AclIds
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateAclsWithListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates additional certificates with an HTTPS listener. You can associate multiple certificates with an HTTPS listener and configure virtual endpoint groups and forwarding rules to accelerate access to multiple HTTPS-capable domain names.
//
// Description:
//
//	  Only HTTPS listeners can be associated with additional certificates.
//
//		- **AssociateAdditionalCertificatesWithListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of the listener with which you want to associate an additional certificate.
//
//	    	- If the listener is in the **updating*	- state, it indicates that the additional certificate is being associated. In this case, you can perform only query operations.
//
//	    	- If the listener is in the **active*	- state, it indicates that the additional certificate is associated.
//
//		- The **AssociateAdditionalCertificatesWithListener*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - AssociateAdditionalCertificatesWithListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateAdditionalCertificatesWithListenerResponse
func (client *Client) AssociateAdditionalCertificatesWithListenerWithContext(ctx context.Context, request *AssociateAdditionalCertificatesWithListenerRequest, runtime *dara.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateAdditionalCertificatesWithListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GA集成云产品
//
// @param request - AssociateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateResourcesResponse
func (client *Client) AssociateResourcesWithContext(ctx context.Context, request *AssociateResourcesRequest, runtime *dara.RuntimeOptions) (_result *AssociateResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AssociatedMode) {
		query["AssociatedMode"] = request.AssociatedMode
	}

	if !dara.IsNil(request.AssociatedResourceId) {
		query["AssociatedResourceId"] = request.AssociatedResourceId
	}

	if !dara.IsNil(request.AssociatedResourceRegionId) {
		query["AssociatedResourceRegionId"] = request.AssociatedResourceRegionId
	}

	if !dara.IsNil(request.AssociatedResourceType) {
		query["AssociatedResourceType"] = request.AssociatedResourceType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateResources"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an Anti-DDoS Pro or Anti-DDoS Premium instance with a Global Accelerator (GA) instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - **AttachDdosToAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) operation to query the status of the GA instance.
//
//   - If the GA instance is in the **configuring*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is being associated with the GA instance. In this case, you can perform only query operations.
//
//   - If the GA instance is in the **active*	- state, the Anti-DDoS Pro or Anti-DDoS Premium instance is associated with the GA instance.
//
//   - You cannot repeatedly call the **AttachDdosToAccelerator*	- operation for the same GA instance within a specific period of time.
//
// @param request - AttachDdosToAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDdosToAcceleratorResponse
func (client *Client) AttachDdosToAcceleratorWithContext(ctx context.Context, request *AttachDdosToAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *AttachDdosToAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.DdosConfigList) {
		query["DdosConfigList"] = request.DdosConfigList
	}

	if !dara.IsNil(request.DdosId) {
		query["DdosId"] = request.DdosId
	}

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDdosToAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDdosToAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Log Service Logstore with an endpoint group.
//
// Description:
//
//	  **AttachLogStoreToEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns a request ID, but this operation is still being performed in the system background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the state of an endpoint group.
//
//	    	- If the endpoint group is in the **updating*	- state, it indicates that a Logstore is being associated with the group. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, it indicates that a Logstore is associated with the group.
//
//		- The **AttachLogStoreToEndpointGroup*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - AttachLogStoreToEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLogStoreToEndpointGroupResponse
func (client *Client) AttachLogStoreToEndpointGroupWithContext(ctx context.Context, request *AttachLogStoreToEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachLogStoreToEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AccessLogRecordCustomizedHeaderList) {
		query["AccessLogRecordCustomizedHeaderList"] = request.AccessLogRecordCustomizedHeaderList
	}

	if !dara.IsNil(request.AccessLogRecordCustomizedHeadersEnabled) {
		query["AccessLogRecordCustomizedHeadersEnabled"] = request.AccessLogRecordCustomizedHeadersEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointGroupIds) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlsLogStoreName) {
		query["SlsLogStoreName"] = request.SlsLogStoreName
	}

	if !dara.IsNil(request.SlsProjectName) {
		query["SlsProjectName"] = request.SlsProjectName
	}

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachLogStoreToEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachLogStoreToEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a bandwidth plan with a Global Accelerator (GA) instance.
//
// Description:
//
//	  **BandwidthPackageAddAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the bandwidth plan that you want to associate.
//
//	    	- If the bandwidth plan is in the **binding*	- state, it indicates that the bandwidth plan is being associated. In this case, you can perform only query operations.
//
//	    	- If the bandwidth plan is in the **active*	- state, it indicates that the bandwidth plan is associated.
//
//		- The **BandwidthPackageAddAccelerator*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - BandwidthPackageAddAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BandwidthPackageAddAcceleratorResponse
func (client *Client) BandwidthPackageAddAcceleratorWithContext(ctx context.Context, request *BandwidthPackageAddAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *BandwidthPackageAddAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BandwidthPackageAddAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BandwidthPackageAddAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a bandwidth plan from a Global Accelerator (GA) instance.
//
// Description:
//
//	  **BandwidthPackageRemoveAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the bandwidth plan that you attempt to disassociate.
//
//	    	- If the bandwidth plan is in the **unbinding*	- state, it indicates that the bandwidth plan is being disassociated. In this case, you can perform only query operations.
//
//	    	- If the bandwidth plan is in the **active*	- state, it indicates that the bandwidth plan is disassociated.
//
//		- The **BandwidthPackageRemoveAccelerator*	- cannot be called repeatedly for the same GA instance.
//
// @param request - BandwidthPackageRemoveAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BandwidthPackageRemoveAcceleratorResponse
func (client *Client) BandwidthPackageRemoveAcceleratorWithContext(ctx context.Context, request *BandwidthPackageRemoveAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *BandwidthPackageRemoveAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BandwidthPackageRemoveAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BandwidthPackageRemoveAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a Global Accelerator (GA) resource belongs.
//
// Description:
//
// The **ChangeResourceGroup*	- operation cannot be repeatedly called for the same GA instance within a specific period of time.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

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
		Version:     dara.String("2019-11-20"),
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
// Configures latency monitoring for an endpoint.
//
// Description:
//
//	  **ConfigEndpointProbe*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the status of the endpoint group to which an endpoint belongs and determine whether latency monitoring is configured for the endpoint.
//
//	    	- If the endpoint group is in the **updating*	- state, it indicates that latency monitoring is being configured for the endpoint. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, it indicates that latency monitoring is configured for the endpoint.
//
//		- The **ConfigEndpointProbe*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - ConfigEndpointProbeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigEndpointProbeResponse
func (client *Client) ConfigEndpointProbeWithContext(ctx context.Context, request *ConfigEndpointProbeRequest, runtime *dara.RuntimeOptions) (_result *ConfigEndpointProbeResponse, _err error) {
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

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.ProbePort) {
		query["ProbePort"] = request.ProbePort
	}

	if !dara.IsNil(request.ProbeProtocol) {
		query["ProbeProtocol"] = request.ProbeProtocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigEndpointProbe"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigEndpointProbeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Global Accelerator (GA) instance. GA is a high-availability and high-performance network acceleration service for global users. By leveraging the high-quality BGP bandwidth and global network of Alibaba Cloud, GA allows service providers to deploy applications across regions and users to connect to the nearest access points for content delivery acceleration. This reduces network issues, such as network latency, network jitters, and packet loss.
//
// Description:
//
// ## Description
//
// **CreateAccelerator*	- is an asynchronous operation. After you send a request, the system returns the ID of a GA instance, but the operation is still being performed in the system background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the state of a GA instance.
//
//   - If the GA instance is in the **init*	- state, it indicates that the GA instance is being created. In this case, you can perform only query operations.
//
//   - If the GA instance is in the **active*	- state, it indicates that the GA instance is created.
//
// @param request - CreateAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAcceleratorResponse
func (client *Client) CreateAcceleratorWithContext(ctx context.Context, request *CreateAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *CreateAcceleratorResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthBillingType) {
		query["BandwidthBillingType"] = request.BandwidthBillingType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.IpSetConfig) {
		query["IpSetConfig"] = request.IpSetConfig
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL).
//
// Description:
//
// *CreateAcl*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) operation to query the state of an ACL.
//
//   - If the ACL is in the **init*	- state, the ACL is being created. In this case, you can only perform only query operations.
//
//   - If the ACL is in the **active*	- state, the ACL is created.
//
// @param request - CreateAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclResponse
func (client *Client) CreateAclWithContext(ctx context.Context, request *CreateAclRequest, runtime *dara.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntries) {
		query["AclEntries"] = request.AclEntries
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAcl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an origin probing task.
//
// Description:
//
// You can call the **CreateApplicationMonitor*	- operation to create an origin probing task. An origin probing task monitors the network quality between a client and an origin server and checks the availability of the origin server.
//
// Before you call this operation, take note of the following items:
//
//   - You can create origin detection tasks only for subscription Standard Global Accelerator (GA) instances whose specification is Medium Ⅰ.
//
//   - You cannot create an origin probe task for a UDP listener.
//
//   - The service port of the URL or IP address that is probed must be within the listening port range.
//
//   - **CreateApplicationMonitor*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) operation to query the status of the origin probing task.
//
//   - If the origin probing task is in the **init*	- state, it indicates that the task is being created. You can perform only query operations.
//
//   - If the origin probing task is in the **active*	- state, it indicates that the task is created.
//
//   - The **CreateApplicationMonitor*	- operation cannot be called repeatedly for the same GA instance within a specific period of time.
//
// @param request - CreateApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationMonitorResponse
func (client *Client) CreateApplicationMonitorWithContext(ctx context.Context, request *CreateApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DetectEnable) {
		query["DetectEnable"] = request.DetectEnable
	}

	if !dara.IsNil(request.DetectThreshold) {
		query["DetectThreshold"] = request.DetectThreshold
	}

	if !dara.IsNil(request.DetectTimes) {
		query["DetectTimes"] = request.DetectTimes
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a bandwidth plan.
//
// Description:
//
// To use Global Accelerator (GA) for acceleration, you must purchase a basic bandwidth plan. A basic bandwidth plan supports the following bandwidth types:
//
//   - **Basic**: Both the default acceleration region and the default service region are in the Chinese mainland. The accelerated service is deployed on Alibaba Cloud.
//
//   - **Enhanced**: Both the default acceleration region and the default service region are in the Chinese mainland. The accelerated service can be deployed on and off Alibaba Cloud.
//
//   - **Premium**: Both the default acceleration region and the default service region are outside the Chinese mainland. The accelerated service can be deployed on and off Alibaba Cloud. If you want to accelerate data transfer for clients in the Chinese mainland, you must select China (Hong Kong) as the acceleration region.
//
// When you call this operation, take note of the following items:
//
//   - **CreateBandwidthPackage*	- is an asynchronous operation. After you send a request, the system returns the ID of a bandwidth plan, but the bandwidth plan is still being created in the system background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the bandwidth plan.
//
//   - If the bandwidth plan is in the **init*	- state, it indicates that the bandwidth plan is being created. In this case, you can perform only query operations.
//
//   - If the bandwidth plan is in the **active*	- state, it indicates that the bandwidth plan is created.
//
//   - The **CreateBandwidthPackage*	- operation cannot be repeatedly called for the same GA instance within a specific period of time.
//
// @param request - CreateBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBandwidthPackageResponse
func (client *Client) CreateBandwidthPackageWithContext(ctx context.Context, request *CreateBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateBandwidthPackageResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthType) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !dara.IsNil(request.BillingType) {
		query["BillingType"] = request.BillingType
	}

	if !dara.IsNil(request.CbnGeographicRegionIdA) {
		query["CbnGeographicRegionIdA"] = request.CbnGeographicRegionIdA
	}

	if !dara.IsNil(request.CbnGeographicRegionIdB) {
		query["CbnGeographicRegionIdB"] = request.CbnGeographicRegionIdB
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Ratio) {
		query["Ratio"] = request.Ratio
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBandwidthPackage"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an accelerated IP address for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **CreateBasicAccelerateIp*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) operation to query the status of an accelerated IP address:
//
//	    	- If no status information is returned, the accelerated IP address is being created. In this case, you can perform only query operations.
//
//	    	- If the accelerated IP address is in the **active*	- state, the accelerated IP address is created.
//
//		- The **CreateBasicAccelerateIp*	- operation cannot be repeatedly called for the same GA instance within a specific period of time.
//
// @param request - CreateBasicAccelerateIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAccelerateIpResponse
func (client *Client) CreateBasicAccelerateIpWithContext(ctx context.Context, request *CreateBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicAccelerateIp"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicAccelerateIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mapping between an accelerated IP address and an endpoint for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **CreateBasicAccelerateIpEndpointRelation*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) API operation to query the status of an accelerated IP address or an endpoint to determine the association status between the accelerated IP address and endpoint.
//
//	    	- If the status of the accelerated IP address and endpoint is **binding**, the accelerated IP address is being associated with the endpoint. In this case, you can query the accelerated IP address and endpoint but cannot perform other operations.
//
//	    	- If the status of the accelerated IP address and endpoint is **bound*	- and the status returned by the [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) API operation is **active**, the accelerated IP address is associated with the endpoint.
//
//		- The **CreateBasicAccelerateIpEndpointRelation*	- API operation cannot be repeatedly called for the same basic GA instance within a period of time.
//
// @param request - CreateBasicAccelerateIpEndpointRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAccelerateIpEndpointRelationResponse
func (client *Client) CreateBasicAccelerateIpEndpointRelationWithContext(ctx context.Context, request *CreateBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicAccelerateIpEndpointRelation"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates mappings between accelerated IP addresses and endpoints for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  The **CreateBasicAccelerateIpEndpointRelations*	- is asynchronous. After you send a request, the system returns a request ID and runs the task in the system background. You can call the [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) or [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) API operation to query the status of an accelerated IP address or an endpoint to determine the association status.
//
//	    	- If an accelerated IP address and the endpoint are in the **binding*	- state, the accelerated IP address is being associated with the endpoint. In this case, you can only query the accelerated IP address and endpoint, but cannot perform other operations.
//
//	    	- If all the accelerated IP addresses and the endpoint are in the **bound*	- state, and the association status returned by the [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) API operation is **active**, the accelerated IP addresses are associated with the endpoints.
//
//		- The **CreateBasicAccelerateIpEndpointRelations*	- API operation cannot be repeatedly called for the same basic GA instance within a period of time.
//
// @param request - CreateBasicAccelerateIpEndpointRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAccelerateIpEndpointRelationsResponse
func (client *Client) CreateBasicAccelerateIpEndpointRelationsWithContext(ctx context.Context, request *CreateBasicAccelerateIpEndpointRelationsRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpEndpointRelations) {
		query["AccelerateIpEndpointRelations"] = request.AccelerateIpEndpointRelations
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicAccelerateIpEndpointRelations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Basic Global Accelerator (GA) instances leverage the immense bandwidth of the high-quality global network of Alibaba Cloud to provide end-to-end acceleration services. You can use basic GA instances to accelerate content delivery at Layer 3 (IP). You can call the CreateBasicAccelerator operation to create a basic GA instance.
//
// Description:
//
// *CreateBasicAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) or [ListBasicAccelerators](https://help.aliyun.com/document_detail/353189.html) operation to query the status of the task.
//
//   - If the basic GA instance is in the **init*	- state, it indicates that the basic GA instance is being created. In this case, you can perform only query operations.
//
//   - If the basic GA instance is in the **active*	- state, it indicates that the basic GA instance is created.
//
// @param request - CreateBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicAcceleratorResponse
func (client *Client) CreateBasicAcceleratorWithContext(ctx context.Context, request *CreateBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicAcceleratorResponse, _err error) {
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

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BandwidthBillingType) {
		query["BandwidthBillingType"] = request.BandwidthBillingType
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **CreateBasicEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) API operation to query the status of an endpoint.
//
//	    	- If the endpoint is in the **init*	- state, the endpoint is being created. In this case, you can perform only query operations.
//
//	    	- If the endpoint is in the **active*	- state, the endpoint is created.
//
//		- The **CreateBasicEndpoint*	- API operation cannot be repeatedly called for the same basic GA instance within a specific period of time.
//
// @param request - CreateBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicEndpointResponse
func (client *Client) CreateBasicEndpointWithContext(ctx context.Context, request *CreateBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointAddress) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointSubAddress) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !dara.IsNil(request.EndpointSubAddressType) {
		query["EndpointSubAddressType"] = request.EndpointSubAddressType
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.EndpointZoneId) {
		query["EndpointZoneId"] = request.EndpointZoneId
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
		Action:      dara.String("CreateBasicEndpoint"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint group for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **CreateBasicEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns an endpoint group ID and runs the task in the background. You can call the [GetBasicEndpointGroup](https://help.aliyun.com/document_detail/362984.html) operation to query the status of the task.
//
//	    	- If the endpoint group is in the **init*	- state, the endpoint is being created. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, the endpoint group is created.
//
//		- You cannot call the **CreateBasicEndpointGroup*	- operation again on the same GA instance before the previous request is completed.
//
// @param request - CreateBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicEndpointGroupResponse
func (client *Client) CreateBasicEndpointGroupWithContext(ctx context.Context, request *CreateBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndpointAddress) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !dara.IsNil(request.EndpointGroupRegion) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !dara.IsNil(request.EndpointSubAddress) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
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
		Action:      dara.String("CreateBasicEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple endpoints for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **CreateBasicEndpoints*	- is an asynchronous operation. After you call this operation, the system returns a request ID and runs the task in the background. You can call the [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) operation to query the status of endpoints. - If one or more endpoints are in the **init*	- state, it indicates that the endpoints are being created. In this case, you can continue to perform query operations on the endpoints. If all endpoints are in the **active*	- state, it indicates that the endpoints are created.
//
//		- You cannot call the **CreateBasicEndpoints*	- operation again on the same GA instance before the previous operation is complete.
//
// @param request - CreateBasicEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicEndpointsResponse
func (client *Client) CreateBasicEndpointsWithContext(ctx context.Context, request *CreateBasicEndpointsRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.Endpoints) {
		query["Endpoints"] = request.Endpoints
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an acceleration region for a basic Global Accelerator (GA) instance.
//
// Description:
//
// Take note of the following limits:
//
//   - You can specify only one acceleration region for each basic GA instance, and only IPv4 clients can connect to basic GA instances.
//
//   - **CreateBasicIpSet*	- is an asynchronous operation. After you send a request, the system returns an acceleration region ID and runs the task in the background. You can call the [GetBasicIpSet](https://help.aliyun.com/document_detail/362987.html) operation to query the status of the task.
//
//   - If the acceleration region is in the **init*	- state, the acceleration region is being created. In this case, you can perform only query operations.
//
//   - If the acceleration region is in the **active*	- state, the acceleration region is created.
//
//   - You cannot call the **CreateBasicIpSet*	- operation again on the same GA instance before the previous task is completed.
//
// @param request - CreateBasicIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBasicIpSetResponse
func (client *Client) CreateBasicIpSetWithContext(ctx context.Context, request *CreateBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *CreateBasicIpSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateRegionId) {
		query["AccelerateRegionId"] = request.AccelerateRegionId
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IspType) {
		query["IspType"] = request.IspType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBasicIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBasicIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you configure a custom routing listener for a Global Accelerator (GA) instance, the GA instance generates a port mapping table based on the listener port range, mapping information (protocols and port ranges) of the associated endpoint groups, and IP addresses of endpoints (vSwitches), and forwards client requests to the specified IP addresses and ports in the vSwitches.
//
// You can call this operation to create mappings for an endpoint group of a custom routing listener. Take note of the following items:
//
//   - **CreateCustomRoutingEndpointGroupDestinations*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the task.
//
//   - If the endpoint group is in the **updating*	- state, it indicates that the mappings are being created for the endpoint group. In this case, you can perform only query operations.
//
//   - If the endpoint group is in the **active*	- state, it indicates that the mappings are created for the endpoint group.
//
//   - You cannot call the **CreateCustomRoutingEndpointGroupDestinations*	- operation again on the same GA instance before the previous task is completed.
//
// ### Prerequisites
//
// Make sure that the following prerequisites are met before you call this operation:
//
//   - A standard GA instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
//   - A bandwidth plan is associated with the standard GA instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
//   - An application is deployed as an endpoint to receive requests that are forwarded from GA. You can specify only vSwitches as endpoints for custom routing listeners.
//
//   - The permissions to use custom routing listeners are acquired and a custom routing listener is created for the GA instance. The custom routing listener feature is in invitational preview. To use the feature, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.11182188.console-base-top.dworkorder.18ae4882n3v6ZW#/ticket/createIndex). For information about how to create a custom routing listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
//   - An endpoint group is created for the custom routing listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// Description:
//
// readAndWrite
//
// @param request - CreateCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) CreateCustomRoutingEndpointGroupDestinationsWithContext(ctx context.Context, request *CreateCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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

	if !dara.IsNil(request.DestinationConfigurations) {
		query["DestinationConfigurations"] = request.DestinationConfigurations
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomRoutingEndpointGroupDestinations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates endpoint groups for a custom routing listener.
//
// Description:
//
// Global Accelerator (GA) forwards client requests to endpoints in an endpoint group based on the routing type of the listener that is associated with the endpoint group.
//
//   - After you configure an intelligent routing listener for a GA instance, the GA instance selects a nearby and healthy endpoint group and forwards client requests to a healthy endpoint in the endpoint group.
//
//   - After you configure a custom routing listener for a GA instance, the instance generates a port mapping table based on the listener port range, protocols and port ranges of the associated endpoint groups, and IP addresses of endpoints (vSwitches), and forwards client requests to specified IP addresses and ports in the vSwitches.
//
// You can call this operation to create endpoint groups for custom routing listeners. For information about how to create endpoint groups for intelligent routing listeners, see [CreateEndpointGroup](https://help.aliyun.com/document_detail/153259.html).
//
// When you call this operation, take note of the following items:
//
//   - **CreateCustomRoutingEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) or [ListCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449374.html) operation to query the status of the endpoint groups that are associated with custom routing listeners.
//
//   - If one or more endpoint groups are in the **init*	- state, it indicates that the endpoint groups are being created. In this case, you can perform only query operations.
//
//   - If all endpoint groups are in the **active*	- state, it indicates that the endpoint groups are created.
//
//   - The **CreateCustomRoutingEndpointGroups*	- operation cannot be called repeatedly for the same GA instance within a specific period of time.
//
// ### Prerequisites
//
// Make sure that the following requirements are met before you call this operation:
//
//   - A standard GA instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
//   - A bandwidth plan is associated with the standard GA instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
//   - An application is deployed to receive requests that are forwarded from GA. You can specify only vSwitches as endpoints for custom routing listeners.
//
//   - The permissions to use custom routing listeners are acquired and a custom routing listener is created for the GA instance. Custom routing listeners are in invitational preview. To use custom routing listeners, contact your account manager. For more information about how to create a custom routing listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
// @param request - CreateCustomRoutingEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointGroupsResponse
func (client *Client) CreateCustomRoutingEndpointGroupsWithContext(ctx context.Context, request *CreateCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupConfigurations) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomRoutingEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates traffic destinations for an endpoint that is associated with a custom routing listener.
//
// Description:
//
// This operation takes effect only when the traffic access policy of an endpoint allows traffic to specified destinations. You can call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation to query the traffic access policy of an endpoint. This operation takes effect only if the value of **TrafficToEndpointPolicy*	- is set to **AllowCustom**, which allows traffic to specific destinations.
//
// When you call this operation, take note of the following items:
//
//   - **CreateCustomRoutingEndpointTrafficPolicies*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the task.
//
//   - If the endpoint group is in the **updating*	- state, the traffic destinations are being created. In this state, you can only query the traffic destinations.
//
//   - If the endpoint group is in the **active*	- state, the traffic destinations are created.
//
//   - You cannot call the **CreateCustomRoutingEndpointTrafficPolicies*	- operation repeatedly for the same GA instance in a specific period of time.
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - A standard GA instance is created. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
//   - If the bandwidth metering method of the standard GA instance is **pay-by-bandwidth**, a bandwidth plan must be associated with the standard GA instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
//   - An application that serves as the endpoint of the standard GA instance is deployed to receive requests that are forwarded from GA. You can specify only vSwitches as endpoints for custom routing listeners.
//
//   - The permissions to use custom routing listeners are acquired, and a custom routing listener is created. Custom routing listeners are in invitational preview. To use custom routing listeners, contact your account manager. For more information about how to create a custom routing listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
//   - Endpoint groups are created for the custom routing listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
//   - Endpoints are created for the custom routing listener. For more information, see [CreateCustomRoutingEndpoints](https://help.aliyun.com/document_detail/449382.html).
//
// @param request - CreateCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) CreateCustomRoutingEndpointTrafficPoliciesWithContext(ctx context.Context, request *CreateCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.PolicyConfigurations) {
		query["PolicyConfigurations"] = request.PolicyConfigurations
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomRoutingEndpointTrafficPolicies"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates endpoints for a custom routing listener.
//
// Description:
//
// After you configure a custom routing listener for a Global Accelerator (GA) instance, the instance generates a port mapping table based on the listener port range, the protocols and port ranges of the associated endpoint groups, and the IP addresses of endpoints (vSwitches), and forwards client requests to specified IP addresses and ports in the vSwitches.
//
// This operation is used to create endpoints for custom routing listeners. When you call this operation, take note of the following items:
//
//   - **CreateCustomRoutingEndpoints*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of an endpoint group and determine whether endpoints are created in the endpoint group.
//
//   - If the endpoint group is in the **updating*	- state, it indicates that endpoints are being created. In this case, you can perform only query operations.
//
//   - If the endpoint group is in the **active*	- state, it indicates that endpoints are created.
//
//   - The **CreateCustomRoutingEndpoints*	- operation cannot be called repeatedly for the same GA instance within a specific period of time.
//
// ### Prerequisites
//
// The following operations are complete before you call this operation:
//
//   - Create a standard GA instance. For more information, see [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html).
//
//   - Associate a bandwidth plan with the standard GA instance. For more information, see [BandwidthPackageAddAccelerator](https://help.aliyun.com/document_detail/153239.html).
//
//   - Deploy an application that serves as the endpoint of the GA instance. The application is used to receive requests that are forwarded from GA. You can specify only vSwitches as endpoints for custom routing listeners.
//
//   - Apply for permissions to use custom routing listeners and create a custom routing listener for the standard GA instance. Custom routing listeners are in invitational preview. To use custom routing listeners, contact your account manager. For more information about how to create a custom routing listener, see [CreateListener](https://help.aliyun.com/document_detail/153253.html).
//
//   - Create an endpoint group for the custom routing listener. For more information, see [CreateCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/449363.html).
//
// @param request - CreateCustomRoutingEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomRoutingEndpointsResponse
func (client *Client) CreateCustomRoutingEndpointsWithContext(ctx context.Context, request *CreateCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomRoutingEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.EndpointConfigurations) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomRoutingEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name and associates the domain name with Global Accelerator (GA) instances.
//
// Description:
//
// After you associate an accelerated domain name that has obtained an ICP number with a Global Accelerator (GA) instance, you do not need to complete filing for the accelerated domain name or its subdomains on Alibaba Cloud.
//
// You can call this operation to add an accelerated domain name and associate the accelerated domain name with GA instances. When you call this operation, take note of the following items:
//
//   - If your accelerated domain name is hosted in the Chinese mainland, you must obtain an ICP number for the domain name.
//
//   - The same accelerated domain name cannot be repeatedly associated with the same GA instance.
//
//   - You cannot repeatedly call the **CreateDomain*	- operation by using the same Alibaba Cloud account within a specific period of time.
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorIds) {
		query["AcceleratorIds"] = request.AcceleratorIds
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint group.
//
// Description:
//
//	  When you call this operation to create a virtual endpoint group for a Layer 4 listener, make sure that a default endpoint group is created.
//
//		- **CreateEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns the ID of an endpoint group, but the endpoint group is still being created in the system background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the state of the endpoint group.
//
//	    	- If the endpoint group is in the **init*	- state, it indicates that the endpoint group is being created. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, it indicates that the endpoint group is created.
//
//		- The **CreateEndpointGroup*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - CreateEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointGroupResponse
func (client *Client) CreateEndpointGroupWithContext(ctx context.Context, request *CreateEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointConfigurations) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !dara.IsNil(request.EndpointGroupRegion) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !dara.IsNil(request.EndpointGroupType) {
		query["EndpointGroupType"] = request.EndpointGroupType
	}

	if !dara.IsNil(request.EndpointIpVersion) {
		query["EndpointIpVersion"] = request.EndpointIpVersion
	}

	if !dara.IsNil(request.EndpointProtocolVersion) {
		query["EndpointProtocolVersion"] = request.EndpointProtocolVersion
	}

	if !dara.IsNil(request.EndpointRequestProtocol) {
		query["EndpointRequestProtocol"] = request.EndpointRequestProtocol
	}

	if !dara.IsNil(request.HealthCheckEnabled) {
		query["HealthCheckEnabled"] = request.HealthCheckEnabled
	}

	if !dara.IsNil(request.HealthCheckHost) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !dara.IsNil(request.HealthCheckIntervalSeconds) {
		query["HealthCheckIntervalSeconds"] = request.HealthCheckIntervalSeconds
	}

	if !dara.IsNil(request.HealthCheckPath) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !dara.IsNil(request.HealthCheckPort) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !dara.IsNil(request.HealthCheckProtocol) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortOverrides) {
		query["PortOverrides"] = request.PortOverrides
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ThresholdCount) {
		query["ThresholdCount"] = request.ThresholdCount
	}

	if !dara.IsNil(request.TrafficPercentage) {
		query["TrafficPercentage"] = request.TrafficPercentage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple endpoint groups at a time.
//
// Description:
//
//	  You can call this operation to create multiple endpoint groups at a time. However, you cannot create a default endpoint group and a virtual endpoint group at the same time.
//
//		- You cannot create a virtual endpoint group for a Layer 4 listener. To create a virtual endpoint group for a Layer 4 listener, call the [CreateEndpointGroup](https://help.aliyun.com/document_detail/2302394.html) operation.
//
//		- **CreateEndpointGroups*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) or [ListEndpointGroups](https://help.aliyun.com/document_detail/153261.html) operation to query the status of endpoint groups.
//
//	    	- If the endpoint groups are in the **init*	- state, the endpoint groups are being created. In this case, you can perform only query operations.
//
//	    	- If all endpoint groups are in the **active*	- state, the endpoint groups are created.
//
//		- The **CreateEndpointGroups*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - CreateEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointGroupsResponse
func (client *Client) CreateEndpointGroupsWithContext(ctx context.Context, request *CreateEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *CreateEndpointGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EndpointGroupConfigurations) {
		bodyFlat["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you want to distribute and process traffic based on request attributes, such as domain names and paths, or information in requests, such as HTTP headers and cookies, you can create custom forwarding rules for a listener. The listener forwards requests based on the forwarding rules. You can call the CreateForwardingRules operation to create forwarding rules.
//
// Description:
//
// Before you call this operation to create forwarding rules, we recommend that you learn how forwarding rules work and how requests are matched against forwarding rules. For more information, see [Configure forwarding rules](https://help.aliyun.com/document_detail/204224.html).
//
// When you call this operation, take note of the following items:
//
//   - **CreateForwardingRules*	- is an asynchronous operation. After you send a request, the system returns a forwarding rule ID and runs the task in the background. You can call the [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) operation to query the status of a forwarding rule.
//
//   - If the forwarding rule is in the **configuring*	- state, the rule is being created. In this case, you can only perform query operations.
//
//   - If the forwarding rule is in the **active*	- state, the rule is created.
//
//   - The **CreateForwardingRules*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - CreateForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForwardingRulesResponse
func (client *Client) CreateForwardingRulesWithContext(ctx context.Context, request *CreateForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ForwardingRules) {
		bodyFlat["ForwardingRules"] = request.ForwardingRules
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateForwardingRules"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateForwardingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates acceleration regions.
//
// Description:
//
//	  **CreateIpSets*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeIpSet](https://help.aliyun.com/document_detail/153246.html) operation to query the status of the task.
//
//	    	- If acceleration regions are in the **init*	- state, it indicates that the acceleration regions are being created. In this case, you can perform only query operations.
//
//	    	- If acceleration regions are in the **active*	- state, it indicates that the acceleration regions are created.
//
//		- You cannot call the **CreateIpSets*	- operation again on the same GA instance before the previous operation is completed.
//
// @param request - CreateIpSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpSetsResponse
func (client *Client) CreateIpSetsWithContext(ctx context.Context, request *CreateIpSetsRequest, runtime *dara.RuntimeOptions) (_result *CreateIpSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateRegion) {
		query["AccelerateRegion"] = request.AccelerateRegion
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpSets"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A listener checks connection requests and distributes the requests to endpoints based on forwarding rules that are defined by the scheduling algorithm. You can call the CreateListener operation to create a listener for a GA instance.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - **CreateListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of the task.
//
//   - If the listener is in the **init*	- state, the listener is being created. In this state, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, the listener is created.
//
//   - You cannot repeatedly call the **CreateListener*	- operation for the same GA instance within the specified period of time.
//
// @param request - CreateListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateListenerResponse
func (client *Client) CreateListenerWithContext(ctx context.Context, request *CreateListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
	}

	if !dara.IsNil(request.ClientAffinity) {
		query["ClientAffinity"] = request.ClientAffinity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomRoutingEndpointGroupConfigurations) {
		query["CustomRoutingEndpointGroupConfigurations"] = request.CustomRoutingEndpointGroupConfigurations
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndpointGroupConfigurations) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !dara.IsNil(request.HttpVersion) {
		query["HttpVersion"] = request.HttpVersion
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortRanges) {
		query["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.XForwardedForConfig) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates secondary IP addresses for a CNAME that is assigned to a Global Accelerator (GA) instance. If an acceleration area of the GA instance becomes unavailable, access traffic is redirected to the secondary IP addresses.
//
// Description:
//
//	  **CreateSpareIps*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of a GA instance.
//
//	    	- If the GA instance is in the **configuring*	- state, it indicates that secondary IP addresses are being created for the CNAME that is assigned to the GA instance. In this case, you can only perform query operations.
//
//	    	- If the GA instance is in the **active*	- state, it indicates that secondary IP addresses are created for the CNAME that is assigned to the GA instance.
//
//		- The **CreateSpareIps*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - CreateSpareIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSpareIpsResponse
func (client *Client) CreateSpareIpsWithContext(ctx context.Context, request *CreateSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *CreateSpareIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpareIps) {
		query["SpareIps"] = request.SpareIps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSpareIps"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSpareIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Global Accelerator (GA) instance.
//
// Description:
//
//	  You cannot delete subscription GA instances.
//
//		- **DeleteAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the task.
//
//	    	- If the GA instance is in the **deleting*	- state, the GA instance is being deleted. In this case, you can perform only query operations.
//
//	    	- If the GA instance cannot be queried, the GA instance is deleted.
//
// @param request - DeleteAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAcceleratorResponse
func (client *Client) DeleteAcceleratorWithContext(ctx context.Context, request *DeleteAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a access control list (ACL) of a Global Accelerator (GA) instance.
//
// Description:
//
// *DeleteAcl*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [GetAcl](https://help.aliyun.com/document_detail/258292.html) operation to query the status of an ACL.
//
//   - If the ACL is in the **deleting*	- state, it indicates that the ACL is being deleted. In this case, you can perform only query operations.
//
//   - If the ACL cannot be queried, it indicates that the ACL is deleted.
//
// @param request - DeleteAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclResponse
func (client *Client) DeleteAclWithContext(ctx context.Context, request *DeleteAclRequest, runtime *dara.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAcl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an origin probing task.
//
// Description:
//
//	  **DeleteApplicationMonitor*	- is an asynchronous operation. After you call this operation, the system returns a request ID, but the operation is still being performed in the system background. You can call the [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) operation to query the state of an origin probing task.
//
//	    	- If the origin probing task is in the **deleting*	- state, it indicates that the task is being deleted. In this case, you can perform only query operations.
//
//	    	- If the origin probing task cannot be queried, it indicates that the task is deleted.
//
//		- The **DeleteApplicationMonitor*	- operation cannot be called repeatedly for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationMonitorResponse
func (client *Client) DeleteApplicationMonitorWithContext(ctx context.Context, request *DeleteApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bandwidth plan.
//
// Description:
//
//	  By default, subscription bandwidth plans cannot be deleted. If you want to unsubscribe from subscription bandwidth plans, go to the [Unsubscribe](https://usercenter2-intl.aliyun.com/refund/refund) page. Before you can unsubscribe from a subscription bandwidth plan that is associated with a Global Accelerator (GA) instance, you must disassociate the bandwidth plan from the GA instance. For information about how to disassociate a bandwidth plan from a GA instance, see [BandwidthPackageRemoveAccelerator](https://help.aliyun.com/document_detail/153240.html).
//
//		- Bandwidth plans that are associated with GA instances cannot be deleted. Before you can delete a bandwidth plan that is associated with a GA instance, you must disassociate the bandwidth plan from the GA instance. For information about how to disassociate a bandwidth plan from a GA instance, see [BandwidthPackageRemoveAccelerator](https://help.aliyun.com/document_detail/153240.html).
//
//		- **DeleteBandwidthPackage*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the task.
//
//	    	- If the bandwidth plan is in the **deleting*	- state, the bandwidth plan is being deleted. In this case, you can perform only query operations.
//
//	    	- If the bandwidth plan cannot be found, the bandwidth plan is deleted.
//
//		- The **DeleteBandwidthPackage*	- operation cannot be repeatedly called for the same bandwidth plan within a specific period of time.
//
// @param request - DeleteBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBandwidthPackageResponse
func (client *Client) DeleteBandwidthPackageWithContext(ctx context.Context, request *DeleteBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DeleteBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBandwidthPackage"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an accelerated IP address of a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **DeleteBasicAccelerateIp*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) operation to query the status of an accelerated IP address.
//
//	    	- If an accelerated IP address is in the **deleting*	- state, the accelerated IP address is being deleted. In this case, you can perform only query operations.
//
//	    	- If the system fails to return information about an accelerated IP address, the accelerated IP address is deleted.
//
//		- You cannot repeatedly call the **DeleteBasicAccelerateIp*	- operation for the same basic GA instance within a specific period of time.
//
// @param request - DeleteBasicAccelerateIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicAccelerateIpResponse
func (client *Client) DeleteBasicAccelerateIpWithContext(ctx context.Context, request *DeleteBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAccelerateIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicAccelerateIp"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicAccelerateIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a mapping between an accelerated IP address and an endpoint for a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **DeleteBasicAccelerateIpEndpointRelation*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the following operations to check whether an accelerated IP address is disassociated from an endpoint:
//
//	    	- You can call the [GetBasicAccelerateIp](https://help.aliyun.com/document_detail/466794.html) and [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) operations to query the status of an accelerated IP address and an endpoint. If the accelerated IP address and the endpoint are in the **unbinding*	- state, the accelerated IP address is being disassociated from the endpoint. In this case, you can query the IP address and endpoint but cannot perform other operations.
//
//	    	- If the association status between the accelerated IP address and the endpoint cannot be queried by calling the [ListBasicAccelerateIpEndpointRelations](https://help.aliyun.com/document_detail/466803.html) operation, the accelerated IP address is disassociated from the endpoint.
//
//		- The **DeleteBasicAccelerateIpEndpointRelation*	- API operation cannot be repeatedly called for the same basic GA instance within a period of time.
//
// @param request - DeleteBasicAccelerateIpEndpointRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicAccelerateIpEndpointRelationResponse
func (client *Client) DeleteBasicAccelerateIpEndpointRelationWithContext(ctx context.Context, request *DeleteBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAccelerateIpEndpointRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicAccelerateIpEndpointRelation"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  You cannot delete subscription basic GA instances. You can unsubscribe from a basic GA instance on the [Unsubscribe](https://usercenter2-intl.aliyun.com/refund/refund) page. Before you unsubscribe from a basic GA instance, make sure that the acceleration areas and endpoint groups of the GA instance are deleted and no bandwidth plans are associated with the GA instance.
//
//	    	- For information about how to delete an acceleration area, see [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//	    	- For information about how to delete an endpoint group, see [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
//	    	- For information about how to disassociate a bandwidth plan from a basic GA instance, see [BandwidthPackageRemoveAccelerator](https://help.aliyun.com/document_detail/153240.html).
//
//		- Before you call this operation to delete a pay-as-you-go basic GA instance, make sure that all data is migrated and the acceleration areas and endpoint groups of the instance are deleted.
//
//	    	- For information about how to delete an acceleration area, see [DeleteBasicIpSet](https://help.aliyun.com/document_detail/2253388.html).
//
//	    	- For information about how to delete an endpoint group, see [DeleteBasicEndpointGroup](https://help.aliyun.com/document_detail/2253399.html).
//
//		- **DeleteBasicAccelerator*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) operation to query the status of the task.
//
//	    	- If the basic GA instance is in the **deleting*	- state, it indicates that the instance is being deleted. In this case, you can perform only query operations.
//
//	    	- If the information about the basic GA instance is not displayed in the response, it indicates that the instance is deleted.
//
// @param request - DeleteBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicAcceleratorResponse
func (client *Client) DeleteBasicAcceleratorWithContext(ctx context.Context, request *DeleteBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint that is associated with a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **DeleteBasicEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListBasicEndpoints](https://help.aliyun.com/document_detail/466831.html) operation to query the status of endpoints.
//
//	    	- If the endpoint is in the **deleting*	- state, it indicates that the endpoint is being deleted. In this case, you can perform only query operations.
//
//	    	- If the endpoint cannot be found, it indicates that the endpoint is deleted.
//
//		- The **DeleteBasicEndpoint*	- API operation cannot be repeatedly called for the same basic GA instance within a period of time.
//
// @param request - DeleteBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicEndpointResponse
func (client *Client) DeleteBasicEndpointWithContext(ctx context.Context, request *DeleteBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicEndpoint"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint group that is associated with a basic Global Accelerator (GA) instance.
//
// Description:
//
// Before you delete an endpoint group, take note of the following items:
//
//   - If an endpoint in the endpoint group is associated with an accelerated IP address, you cannot delete the endpoint group. You can call the [DeleteBasicAccelerateIpEndpointRelation](https://help.aliyun.com/document_detail/2253413.html) operation to disassociate the endpoint from the accelerated IP address.
//
//   - If no endpoint in the endpoint group is associated with an accelerated IP address, you can delete the endpoint group. When you delete an endpoint group, all endpoints in the endpoint group are deleted.
//
//   - **DeleteBasicEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicEndpointGroup](https://help.aliyun.com/document_detail/362984.html) operation to query the status of the task.
//
//   - If the endpoint group is in the **deleting*	- state, the endpoint group is being deleted. In this case, you can perform only query operations.
//
//   - If the endpoint group cannot be queried, the endpoint group is deleted.
//
//   - The **DeleteBasicEndpointGroup*	- operation cannot be repeatedly called for the same GA instance within a specific period of time.
//
// @param request - DeleteBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicEndpointGroupResponse
func (client *Client) DeleteBasicEndpointGroupWithContext(ctx context.Context, request *DeleteBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicEndpointGroupResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the acceleration region of a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  If an accelerated IP address is associated with an endpoint, you cannot delete the acceleration region. You can call the [DeleteBasicAccelerateIpEndpointRelation](https://help.aliyun.com/document_detail/2253413.html) operation to disassociate the accelerated IP address from the endpoint.
//
//		- \\*\\*DeleteBasicIpSet\\*\\	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicIpSet](https://help.aliyun.com/document_detail/362987.html) operation to query the status of the task.
//
//	    	- If the acceleration region is in the **deleting*	- state, it indicates that the acceleration region is being deleted. In this case, you can perform only query operations.
//
//	    	- If the acceleration region cannot be queried, it indicates that the acceleration region is deleted.
//
//		- The \\*\\*DeleteBasicIpSet\\*\\	- operation cannot be repeatedly called for the same basic GA instance within a specific period of time.
//
// @param request - DeleteBasicIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBasicIpSetResponse
func (client *Client) DeleteBasicIpSetWithContext(ctx context.Context, request *DeleteBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteBasicIpSetResponse, _err error) {
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

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBasicIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBasicIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes mappings from an endpoint group that is associated with a custom routing listener.
//
// Description:
//
//	  **DeleteCustomRoutingEndpointGroupDestinations*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) to query the status of the task.
//
//	    	- If the endpoint group is in the **updating*	- state, it indicates that mappings are being deleted from the endpoint group. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state and no information about the mappings that you want to delete is found in the response when you call the [DescribeCustomRoutingEndpointGroupDestinations](https://help.aliyun.com/document_detail/449378.html) operation, it indicates the mappings are deleted from the endpoint group.
//
//		- You cannot call the **DeleteCustomRoutingEndpointGroupDestinations*	- operation again on the same Global Accelerator (GA) instance before the previous request is completed.
//
// @param request - DeleteCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) DeleteCustomRoutingEndpointGroupDestinationsWithContext(ctx context.Context, request *DeleteCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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

	if !dara.IsNil(request.DestinationIds) {
		query["DestinationIds"] = request.DestinationIds
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomRoutingEndpointGroupDestinations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple endpoint groups that are associated with a custom routing listener.
//
// Description:
//
//	  **DeleteCustomRoutingEndpointGroups*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the state of the endpoint groups associated with a custom routing listener that you attempt to delete.
//
//	    	- If the endpoint groups are in the **deleting*	- state, the endpoint groups are being deleted. In this case, you can perform only query operations.
//
//	    	- If the endpoint groups cannot be queried, the endpoint groups are deleted.
//
//		- You cannot use the **DeleteCustomRoutingEndpointGroups*	- operation on the same Global Accelerator (GA) instance before the previous operation is complete.
//
// @param request - DeleteCustomRoutingEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRoutingEndpointGroupsResponse
func (client *Client) DeleteCustomRoutingEndpointGroupsWithContext(ctx context.Context, request *DeleteCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupIds) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomRoutingEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes traffic destinations from an endpoint.
//
// Description:
//
//	  **DeleteCustomRoutingEndpointTrafficPolicies*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of an endpoint group to check whether the traffic destinations are deleted.
//
//	    	- If the endpoint group is in the **updating*	- state, the traffic destinations are being deleted. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state and the traffic destinations that you want to delete cannot be queried by calling the [DescribeCustomRoutingEndPointTrafficPolicy](https://help.aliyun.com/document_detail/449392.html) operation, the traffic destinations are deleted.
//
//		- The **DeleteCustomRoutingEndpointTrafficPolicies*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) DeleteCustomRoutingEndpointTrafficPoliciesWithContext(ctx context.Context, request *DeleteCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.PolicyIds) {
		query["PolicyIds"] = request.PolicyIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomRoutingEndpointTrafficPolicies"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes endpoints from a custom routing listener.
//
// Description:
//
//	  **DeleteCustomRoutingEndpoints*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) to query the status of the task.
//
//	    	- If an endpoint group is in the **updating*	- state, the endpoint is being deleted. In this case, you can perform only query operations.
//
//	    	- If an endpoint group is in the **active*	- state and the endpoint cannot be found after you call the [DescribeCustomRoutingEndpoint](https://help.aliyun.com/document_detail/449386.html) operation, the endpoint is deleted.
//
//		- You cannot call the **DeleteCustomRoutingEndpoints*	- operation again on the same Global Accelerator (GA) instance before the previous task is completed.
//
// @param request - DeleteCustomRoutingEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomRoutingEndpointsResponse
func (client *Client) DeleteCustomRoutingEndpointsWithContext(ctx context.Context, request *DeleteCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomRoutingEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointIds) {
		query["EndpointIds"] = request.EndpointIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomRoutingEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a domain name from Global Accelerator (GA) instances.
//
// Description:
//
// You cannot call the **DeleteDomainAcceleratorRelation*	- operation again by using the same Alibaba Cloud account before the previous operation is complete.
//
// @param request - DeleteDomainAcceleratorRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainAcceleratorRelationResponse
func (client *Client) DeleteDomainAcceleratorRelationWithContext(ctx context.Context, request *DeleteDomainAcceleratorRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainAcceleratorRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorIds) {
		query["AcceleratorIds"] = request.AcceleratorIds
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainAcceleratorRelation"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainAcceleratorRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint group.
//
// Description:
//
//	  **DeleteEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the status of the endpoint group.
//
//	    	- If the endpoint group is in the **deleting*	- state, it indicates that the endpoint group is being deleted. In this case, you can perform only query operations.
//
//	    	- If the endpoint group cannot be queried, it indicates that the endpoint group is deleted.
//
//		- The **DeleteEndpointGroup*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - DeleteEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointGroupResponse
func (client *Client) DeleteEndpointGroupWithContext(ctx context.Context, request *DeleteEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes endpoint groups.
//
// Description:
//
//	  **DeleteEndpointGroups*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the status of the task.
//
//	    	- If an endpoint group is in the **deleting*	- state, the endpoint group is being deleted. In this case, you can perform only query operations.
//
//	    	- If an endpoint group cannot be queried, the endpoint group is deleted.
//
//		- The **DeleteEndpointGroups*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointGroupsResponse
func (client *Client) DeleteEndpointGroupsWithContext(ctx context.Context, request *DeleteEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEndpointGroupsResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupIds) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes forwarding rules.
//
// Description:
//
//	  **DeleteForwardingRules*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) operation to query the status of the task.
//
//	    	- If a forwarding rule is in the **deleting*	- state, the forwarding rule is being deleted. In this case, you can perform only query operations.
//
//	    	- If a forwarding rule cannot be queried, the forwarding rule is deleted.
//
//		- The **DeleteForwardingRules*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteForwardingRulesResponse
func (client *Client) DeleteForwardingRulesWithContext(ctx context.Context, request *DeleteForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForwardingRuleIds) {
		query["ForwardingRuleIds"] = request.ForwardingRuleIds
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteForwardingRules"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteForwardingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an acceleration region.
//
// Description:
//
//	  **DeleteIpSet*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeIpSet](https://help.aliyun.com/document_detail/153246.html) operation to query the status of an acceleration region.
//
//	    	- If the acceleration region is in the **deleting*	- state, it indicates that the acceleration region is being deleted. In this case, you can perform only query operations.
//
//	    	- If the acceleration region cannot be queried, it indicates that the acceleration region is deleted.
//
//		- The **DeleteIpSet*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - DeleteIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpSetResponse
func (client *Client) DeleteIpSetWithContext(ctx context.Context, request *DeleteIpSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes acceleration regions.
//
// Description:
//
//	  **DeleteIpSets*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeIpSet](https://help.aliyun.com/document_detail/153246.html) operation to query the status of the task.
//
//	    	- If the acceleration region is in the **deleting*	- state, the acceleration region is being deleted. In this case, you can perform only query operations.
//
//	    	- If you cannot query the acceleration region, the acceleration region is deleted.
//
//		- You cannot repeatedly call the **DeleteIpSets*	- operation for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteIpSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpSetsResponse
func (client *Client) DeleteIpSetsWithContext(ctx context.Context, request *DeleteIpSetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpSetIds) {
		query["IpSetIds"] = request.IpSetIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpSets"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a listener.
//
// Description:
//
//	  Before you call the **DeleteListener*	- operation, make sure that no endpoint groups are associated with the listener that you want to delete. For information about how to delete an endpoint group, see the following topics:
//
//	    	- [DeleteEndpointGroup](https://help.aliyun.com/document_detail/2253305.html): deletes an endpoint group that is associated with an intelligent routing listener.
//
//	    	- [DeleteEndpointGroups](https://help.aliyun.com/document_detail/2253311.html): deletes multiple endpoint groups that are associated with intelligent routing listeners at the same time.
//
//	    	- [DeleteCustomRoutingEndpointGroups](https://help.aliyun.com/document_detail/2303183.html): deletes multiple endpoint groups that are associated with custom routing listeners at the same time.
//
//		- **DeleteListener*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of the listener.
//
//	    	- If the listener is in the **deleting*	- state, the listener is being deleted. In this case, you can perform only query operations.
//
//	    	- If the listener cannot be queried, the listener is deleted.
//
//		- You cannot repeatedly call the **DeleteListener*	- operation to delete the listeners of the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DeleteListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteListenerResponse
func (client *Client) DeleteListenerWithContext(ctx context.Context, request *DeleteListenerRequest, runtime *dara.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the secondary IP addresses that are associated with a CNAME.
//
// Description:
//
//	  **DeleteSpareIps*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of a GA instance.
//
//	    	- If the GA instance is in the **configuring*	- state, it indicates that the secondary IP addresses for the CNAME are being deleted. In this case, you can perform only query operations.
//
//	    	- If the GA instance is in the **active*	- state and the secondary IP addresses for the CNAME cannot be queried by calling the [ListSpareIps](https://help.aliyun.com/document_detail/262121.html) operation, it indicates that the IP addresses are deleted.
//
//		- The **DeleteSpareIps*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - DeleteSpareIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSpareIpsResponse
func (client *Client) DeleteSpareIpsWithContext(ctx context.Context, request *DeleteSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSpareIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpareIps) {
		query["SpareIps"] = request.SpareIps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSpareIps"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSpareIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Global Accelerator (GA) instance.
//
// @param request - DescribeAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAcceleratorResponse
func (client *Client) DescribeAcceleratorWithContext(ctx context.Context, request *DescribeAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto-renewal status of a Global Accelerator (GA) instance.
//
// @param request - DescribeAcceleratorAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAcceleratorAutoRenewAttributeResponse
func (client *Client) DescribeAcceleratorAutoRenewAttributeWithContext(ctx context.Context, request *DescribeAcceleratorAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAcceleratorAutoRenewAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a pay-as-you-go Global Accelerator (GA) instance.
//
// @param request - DescribeAcceleratorServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAcceleratorServiceStatusResponse
func (client *Client) DescribeAcceleratorServiceStatusWithContext(ctx context.Context, request *DescribeAcceleratorServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAcceleratorServiceStatusResponse, _err error) {
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
		Action:      dara.String("DescribeAcceleratorServiceStatus"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAcceleratorServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an origin probing task.
//
// @param request - DescribeApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationMonitorResponse
func (client *Client) DescribeApplicationMonitorWithContext(ctx context.Context, request *DescribeApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a bandwidth plan.
//
// @param request - DescribeBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBandwidthPackageResponse
func (client *Client) DescribeBandwidthPackageWithContext(ctx context.Context, request *DescribeBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBandwidthPackage"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto-renewal status of a bandwidth plan.
//
// @param request - DescribeBandwidthPackageAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBandwidthPackageAutoRenewAttributeResponse
func (client *Client) DescribeBandwidthPackageAutoRenewAttributeWithContext(ctx context.Context, request *DescribeBandwidthPackageAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeBandwidthPackageAutoRenewAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBandwidthPackageAutoRenewAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBandwidthPackageAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about commodities.
//
// @param request - DescribeCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommodityResponse
func (client *Client) DescribeCommodityWithContext(ctx context.Context, request *DescribeCommodityRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommodityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCommodity"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCommodityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prices of commodities.
//
// Description:
//
// You can call the [DescribeCommodity](https://help.aliyun.com/document_detail/2253233.html) operation to query information about the commodity modules.
//
// @param request - DescribeCommodityPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommodityPriceResponse
func (client *Client) DescribeCommodityPriceWithContext(ctx context.Context, request *DescribeCommodityPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommodityPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Orders) {
		query["Orders"] = request.Orders
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCommodityPrice"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCommodityPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a traffic destination of an endpoint.
//
// @param request - DescribeCustomRoutingEndPointTrafficPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomRoutingEndPointTrafficPolicyResponse
func (client *Client) DescribeCustomRoutingEndPointTrafficPolicyWithContext(ctx context.Context, request *DescribeCustomRoutingEndPointTrafficPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndPointTrafficPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
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
		Action:      dara.String("DescribeCustomRoutingEndPointTrafficPolicy"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomRoutingEndPointTrafficPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified endpoint that is associated with a custom routing listener.
//
// @param request - DescribeCustomRoutingEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomRoutingEndpointResponse
func (client *Client) DescribeCustomRoutingEndpointWithContext(ctx context.Context, request *DescribeCustomRoutingEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointGroup) {
		query["EndpointGroup"] = request.EndpointGroup
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomRoutingEndpoint"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specific endpoint group that is associated with a custom routing listener.
//
// @param request - DescribeCustomRoutingEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomRoutingEndpointGroupResponse
func (client *Client) DescribeCustomRoutingEndpointGroupWithContext(ctx context.Context, request *DescribeCustomRoutingEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomRoutingEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mapping configuration of a specified endpoint group that is associated with a custom routing listener.
//
// @param request - DescribeCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) DescribeCustomRoutingEndpointGroupDestinationsWithContext(ctx context.Context, request *DescribeCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationId) {
		query["DestinationId"] = request.DestinationId
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomRoutingEndpointGroupDestinations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an endpoint group.
//
// @param request - DescribeEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointGroupResponse
func (client *Client) DescribeEndpointGroupWithContext(ctx context.Context, request *DescribeEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an acceleration region.
//
// @param request - DescribeIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpSetResponse
func (client *Client) DescribeIpSetWithContext(ctx context.Context, request *DescribeIpSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration information about a listener of a Global Accelerator (GA) instance.
//
// Description:
//
// This operation is used to query configuration information about a listener of a GA instance. The information includes the routing type of the listener, the status of the listener, the timestamp that indicates when the listener was created, and the listener ports.
//
// @param request - DescribeListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListenerResponse
func (client *Client) DescribeListenerWithContext(ctx context.Context, request *DescribeListenerRequest, runtime *dara.RuntimeOptions) (_result *DescribeListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Simple Log Service project and Logstore associated with an endpoint group.
//
// @param request - DescribeLogStoreOfEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogStoreOfEndpointGroupResponse
func (client *Client) DescribeLogStoreOfEndpointGroupWithContext(ctx context.Context, request *DescribeLogStoreOfEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogStoreOfEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogStoreOfEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogStoreOfEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where Global Accelerator (GA) instances are deployed.
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-11-20"),
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
// Disassociates a Global Accelerator (GA) instance from an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
//	  The **DetachDdosFromAccelerator*	- operation is asynchronous. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) operation to query the status of the GA instance.
//
//	    	- If the GA instance is in the **configuring*	- state, the Anti-DDoS Pro/Premium instance is being disassociated from the GA instance. In this case, you can perform only query operations.
//
//	    	- If the GA instance is in the **active*	- state, the Anti-DDoS Pro/Premium instance is disassociated from the GA instance.
//
//		- **DetachDdosFromAccelerator*	- cannot be repeatedly called for the same GA instance within a specific period of time.
//
// @param request - DetachDdosFromAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDdosFromAcceleratorResponse
func (client *Client) DetachDdosFromAcceleratorWithContext(ctx context.Context, request *DetachDdosFromAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *DetachDdosFromAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.DdosConfigList) {
		query["DdosConfigList"] = request.DdosConfigList
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDdosFromAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDdosFromAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Log Service Logstore from an endpoint group.
//
// Description:
//
// ## Description
//
//   - **DetachLogStoreFromEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the state of an endpoint group.
//
//   - If the endpoint group is in the **updating*	- state, the Log Service Logstore is being disassociated from the endpoint group. In this case, you can perform only query operations.
//
//     <!---->
//
//   - If the endpoint group is in the **active*	- state, the Log Service Logstore is disassociated from the endpoint group.
//
//   - The **DetachLogStoreFromEndpointGroup*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DetachLogStoreFromEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLogStoreFromEndpointGroupResponse
func (client *Client) DetachLogStoreFromEndpointGroupWithContext(ctx context.Context, request *DetachLogStoreFromEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachLogStoreFromEndpointGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointGroupIds) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachLogStoreFromEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachLogStoreFromEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the diagnostic feature for an origin probing task.
//
// @param request - DetectApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectApplicationMonitorResponse
func (client *Client) DetectApplicationMonitorWithContext(ctx context.Context, request *DetectApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DetectApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an origin probing task.
//
// @param request - DisableApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationMonitorResponse
func (client *Client) DisableApplicationMonitorWithContext(ctx context.Context, request *DisableApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑GA集成云产品
//
// @param request - DisassociateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateResourcesResponse
func (client *Client) DisassociateResourcesWithContext(ctx context.Context, request *DisassociateResourcesRequest, runtime *dara.RuntimeOptions) (_result *DisassociateResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AssociatedResourceId) {
		query["AssociatedResourceId"] = request.AssociatedResourceId
	}

	if !dara.IsNil(request.AssociatedResourceRegionId) {
		query["AssociatedResourceRegionId"] = request.AssociatedResourceRegionId
	}

	if !dara.IsNil(request.AssociatedResourceType) {
		query["AssociatedResourceType"] = request.AssociatedResourceType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateResources"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates access control lists (ACLs) from a listener.
//
// Description:
//
// ## Description
//
//   - **DissociateAclsFromListener*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the state of a listener:
//
//   - If the listener is in the **updating*	- state, ACLs are being disassociated from the listener. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, ACLs are disassociated from the listener.
//
//   - The **DissociateAclsFromListener*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - DissociateAclsFromListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateAclsFromListenerResponse
func (client *Client) DissociateAclsFromListenerWithContext(ctx context.Context, request *DissociateAclsFromListenerRequest, runtime *dara.RuntimeOptions) (_result *DissociateAclsFromListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclIds) {
		query["AclIds"] = request.AclIds
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateAclsFromListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an additional certificate from an HTTPS listener.
//
// Description:
//
// ## Description
//
//   - **DissociateAdditionalCertificatesFromListener*	- is an asynchronous operation. After you send a request, the system returns a request ID, but this operation is still being performed in the system background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the state of an HTTPS listener.
//
//   - If the listener is in the **updating*	- state, it indicates that the additional certificate is being dissociated from the listener. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, it indicates that the additional certificate is dissociated from the listener.
//
//   - The **DissociateAdditionalCertificatesFromListener*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance with a specific period of time.
//
// @param request - DissociateAdditionalCertificatesFromListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateAdditionalCertificatesFromListenerResponse
func (client *Client) DissociateAdditionalCertificatesFromListenerWithContext(ctx context.Context, request *DissociateAdditionalCertificatesFromListenerRequest, runtime *dara.RuntimeOptions) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateAdditionalCertificatesFromListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an origin probing task.
//
// @param request - EnableApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationMonitorResponse
func (client *Client) EnableApplicationMonitorWithContext(ctx context.Context, request *EnableApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an access control list (ACL).
//
// @param request - GetAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAclResponse
func (client *Client) GetAclWithContext(ctx context.Context, request *GetAclRequest, runtime *dara.RuntimeOptions) (_result *GetAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAcl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an accelerated IP address of a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAccelerateIpResponse
func (client *Client) GetBasicAccelerateIpWithContext(ctx context.Context, request *GetBasicAccelerateIpRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicAccelerateIp"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicAccelerateIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the accelerated IP address of a basic Global Accelerator (GA) instance is associated with an endpoint.
//
// @param request - GetBasicAccelerateIpEndpointRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAccelerateIpEndpointRelationResponse
func (client *Client) GetBasicAccelerateIpEndpointRelationWithContext(ctx context.Context, request *GetBasicAccelerateIpEndpointRelationRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpEndpointRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicAccelerateIpEndpointRelation"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of idle accelerated IP addresses of a Global Accelerator (GA) instance.
//
// @param request - GetBasicAccelerateIpIdleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAccelerateIpIdleCountResponse
func (client *Client) GetBasicAccelerateIpIdleCountWithContext(ctx context.Context, request *GetBasicAccelerateIpIdleCountRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAccelerateIpIdleCountResponse, _err error) {
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

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicAccelerateIpIdleCount"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicAccelerateIpIdleCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicAcceleratorResponse
func (client *Client) GetBasicAcceleratorWithContext(ctx context.Context, request *GetBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *GetBasicAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about an endpoint that is associated with a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicEndpointResponse
func (client *Client) GetBasicEndpointWithContext(ctx context.Context, request *GetBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetBasicEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicEndpoint"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the endpoint group of a basic GA instance.
//
// @param request - GetBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicEndpointGroupResponse
func (client *Client) GetBasicEndpointGroupWithContext(ctx context.Context, request *GetBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *GetBasicEndpointGroupResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the acceleration region of a basic Global Accelerator (GA) instance.
//
// @param request - GetBasicIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBasicIpSetResponse
func (client *Client) GetBasicIpSetWithContext(ctx context.Context, request *GetBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *GetBasicIpSetResponse, _err error) {
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

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBasicIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBasicIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取GA实例关联的云产品
//
// @param request - GetGlobalAcceleratorResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGlobalAcceleratorResourcesResponse
func (client *Client) GetGlobalAcceleratorResourcesWithContext(ctx context.Context, request *GetGlobalAcceleratorResourcesRequest, runtime *dara.RuntimeOptions) (_result *GetGlobalAcceleratorResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AssociatedResourceId) {
		query["AssociatedResourceId"] = request.AssociatedResourceId
	}

	if !dara.IsNil(request.AssociatedResourceRegionId) {
		query["AssociatedResourceRegionId"] = request.AssociatedResourceRegionId
	}

	if !dara.IsNil(request.AssociatedResourceType) {
		query["AssociatedResourceType"] = request.AssociatedResourceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGlobalAcceleratorResources"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGlobalAcceleratorResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of the endpoints and endpoint groups of a listener.
//
// @param request - GetHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthStatusResponse
func (client *Client) GetHealthStatusWithContext(ctx context.Context, request *GetHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *GetHealthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHealthStatus"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHealthStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of invalid domain names.
//
// @param request - GetInvalidDomainCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInvalidDomainCountResponse
func (client *Client) GetInvalidDomainCountWithContext(ctx context.Context, request *GetInvalidDomainCountRequest, runtime *dara.RuntimeOptions) (_result *GetInvalidDomainCountResponse, _err error) {
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
		Action:      dara.String("GetInvalidDomainCount"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInvalidDomainCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum bandwidth of an acceleration area.
//
// @param request - GetIpsetsBandwidthLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpsetsBandwidthLimitResponse
func (client *Client) GetIpsetsBandwidthLimitWithContext(ctx context.Context, request *GetIpsetsBandwidthLimitRequest, runtime *dara.RuntimeOptions) (_result *GetIpsetsBandwidthLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIpsetsBandwidthLimit"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIpsetsBandwidthLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a secondary IP address that is associated with a CNAME.
//
// @param request - GetSpareIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpareIpResponse
func (client *Client) GetSpareIpWithContext(ctx context.Context, request *GetSpareIpRequest, runtime *dara.RuntimeOptions) (_result *GetSpareIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpareIp) {
		query["SpareIp"] = request.SpareIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpareIp"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpareIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available acceleration areas and regions.
//
// @param request - ListAccelerateAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccelerateAreasResponse
func (client *Client) ListAccelerateAreasWithContext(ctx context.Context, request *ListAccelerateAreasRequest, runtime *dara.RuntimeOptions) (_result *ListAccelerateAreasResponse, _err error) {
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
		Action:      dara.String("ListAccelerateAreas"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccelerateAreasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Global Accelerator (GA) instances.
//
// @param request - ListAcceleratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcceleratorsResponse
func (client *Client) ListAcceleratorsWithContext(ctx context.Context, request *ListAcceleratorsRequest, runtime *dara.RuntimeOptions) (_result *ListAcceleratorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
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

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccelerators"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAcceleratorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of access control lists (ACLs).
//
// @param request - ListAclsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAclsResponse
func (client *Client) ListAclsWithContext(ctx context.Context, request *ListAclsRequest, runtime *dara.RuntimeOptions) (_result *ListAclsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclIds) {
		query["AclIds"] = request.AclIds
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAcls"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAclsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin probing tasks.
//
// @param request - ListApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationMonitorResponse
func (client *Client) ListApplicationMonitorWithContext(ctx context.Context, request *ListApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationMonitorResponse, _err error) {
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

	if !dara.IsNil(request.SearchValue) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostic results of origin probing tasks.
//
// @param request - ListApplicationMonitorDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationMonitorDetectResultResponse
func (client *Client) ListApplicationMonitorDetectResultWithContext(ctx context.Context, request *ListApplicationMonitorDetectResultRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationMonitorDetectResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationMonitorDetectResult"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationMonitorDetectResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available acceleration regions.
//
// @param request - ListAvailableAccelerateAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableAccelerateAreasResponse
func (client *Client) ListAvailableAccelerateAreasWithContext(ctx context.Context, request *ListAvailableAccelerateAreasRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableAccelerateAreasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableAccelerateAreas"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableAccelerateAreasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available acceleration regions of a Global Accelerator (GA) instance.
//
// @param request - ListAvailableBusiRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableBusiRegionsResponse
func (client *Client) ListAvailableBusiRegionsWithContext(ctx context.Context, request *ListAvailableBusiRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableBusiRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableBusiRegions"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableBusiRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth plans.
//
// @param request - ListBandwidthPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBandwidthPackagesResponse
func (client *Client) ListBandwidthPackagesWithContext(ctx context.Context, request *ListBandwidthPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListBandwidthPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
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

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBandwidthPackages"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBandwidthPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth plans.
//
// Description:
//
// To query the detailed information about a bandwidth plan, call the **ListBandwidthPackages*	- operation. For more information, see [ListBandwidthPackages](https://help.aliyun.com/document_detail/2253239.html).
//
// @param request - ListBandwidthackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBandwidthackagesResponse
func (client *Client) ListBandwidthackagesWithContext(ctx context.Context, request *ListBandwidthackagesRequest, runtime *dara.RuntimeOptions) (_result *ListBandwidthackagesResponse, _err error) {
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
		Action:      dara.String("ListBandwidthackages"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBandwidthackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the accelerated IP address of a basic Global Accelerator (GA) instance is associated with an endpoint.
//
// @param request - ListBasicAccelerateIpEndpointRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBasicAccelerateIpEndpointRelationsResponse
func (client *Client) ListBasicAccelerateIpEndpointRelationsWithContext(ctx context.Context, request *ListBasicAccelerateIpEndpointRelationsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAccelerateIpEndpointRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
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
		Action:      dara.String("ListBasicAccelerateIpEndpointRelations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated IP addresses in the acceleration region of a basic Global Accelerator (GA) instance.
//
// @param request - ListBasicAccelerateIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBasicAccelerateIpsResponse
func (client *Client) ListBasicAccelerateIpsWithContext(ctx context.Context, request *ListBasicAccelerateIpsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAccelerateIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateIpAddress) {
		query["AccelerateIpAddress"] = request.AccelerateIpAddress
	}

	if !dara.IsNil(request.AccelerateIpId) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
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
		Action:      dara.String("ListBasicAccelerateIps"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBasicAccelerateIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries basic Global Accelerator (GA) instances.
//
// @param request - ListBasicAcceleratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBasicAcceleratorsResponse
func (client *Client) ListBasicAcceleratorsWithContext(ctx context.Context, request *ListBasicAcceleratorsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicAcceleratorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
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

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBasicAccelerators"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBasicAcceleratorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints that are associated with a basic Global Accelerator (GA) instance.
//
// @param request - ListBasicEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBasicEndpointsResponse
func (client *Client) ListBasicEndpointsWithContext(ctx context.Context, request *ListBasicEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListBasicEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBasicEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBasicEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the acceleration regions that are supported by Global Accelerator (GA).
//
// @param request - ListBusiRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBusiRegionsResponse
func (client *Client) ListBusiRegionsWithContext(ctx context.Context, request *ListBusiRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListBusiRegionsResponse, _err error) {
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
		Action:      dara.String("ListBusiRegions"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBusiRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available acceleration areas and regions.
//
// Description:
//
// You can call this operation to query the acceleration areas and regions that you can specify on the wizard page of Global Accelerator (GA) and for free-trial GA instances. You can filter acceleration areas and regions based on specified conditions.
//
// @param request - ListCommonAreasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonAreasResponse
func (client *Client) ListCommonAreasWithContext(ctx context.Context, request *ListCommonAreasRequest, runtime *dara.RuntimeOptions) (_result *ListCommonAreasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IsEpg) {
		query["IsEpg"] = request.IsEpg
	}

	if !dara.IsNil(request.IsIpSet) {
		query["IsIpSet"] = request.IsIpSet
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommonAreas"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommonAreasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries endpoint group mapping configurations of a custom routing listener of a Global Accelerator (GA) instance.
//
// @param request - ListCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) ListCustomRoutingEndpointGroupDestinationsWithContext(ctx context.Context, request *ListCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.FromPort) {
		query["FromPort"] = request.FromPort
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocols) {
		query["Protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ToPort) {
		query["ToPort"] = request.ToPort
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingEndpointGroupDestinations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoint groups that are associated with a custom routing listener.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code for different SDKs.](https://api.aliyun.com/#product=Ga\\&api=ListCustomRoutingEndpointGroups\\&type=RPC\\&version=2019-11-20)
//
// @param request - ListCustomRoutingEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointGroupsResponse
func (client *Client) ListCustomRoutingEndpointGroupsWithContext(ctx context.Context, request *ListCustomRoutingEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AccessLogSwitch) {
		query["AccessLogSwitch"] = request.AccessLogSwitch
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic policies of an endpoint that belongs to a custom routing listener.
//
// @param request - ListCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) ListCustomRoutingEndpointTrafficPoliciesWithContext(ctx context.Context, request *ListCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingEndpointTrafficPolicies"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an endpoint.
//
// @param request - ListCustomRoutingEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingEndpointsResponse
func (client *Client) ListCustomRoutingEndpointsWithContext(ctx context.Context, request *ListCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the port mapping table of a custom routing listener.
//
// Description:
//
// After you configure a custom routing listener for a Global Accelerator (GA) instance, the instance generates a port mapping table based on the listener port range, backend service protocols and port ranges of the associated endpoint groups, and IP addresses of endpoints (vSwitches). The custom routing listener forwards client requests to specified IP addresses and ports in the vSwitches based on the port mapping table. This operation is used to query the generated port mapping table.
//
// @param request - ListCustomRoutingPortMappingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingPortMappingsResponse
func (client *Client) ListCustomRoutingPortMappingsWithContext(ctx context.Context, request *ListCustomRoutingPortMappingsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingPortMappingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingPortMappings"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingPortMappingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the port mapping table of a specified backend instance that is associated with a custom routing listener.
//
// @param request - ListCustomRoutingPortMappingsByDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomRoutingPortMappingsByDestinationResponse
func (client *Client) ListCustomRoutingPortMappingsByDestinationWithContext(ctx context.Context, request *ListCustomRoutingPortMappingsByDestinationRequest, runtime *dara.RuntimeOptions) (_result *ListCustomRoutingPortMappingsByDestinationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationAddress) {
		query["DestinationAddress"] = request.DestinationAddress
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomRoutingPortMappingsByDestination"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomRoutingPortMappingsByDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names.
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithContext(ctx context.Context, request *ListDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
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

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the public CIDR blocks to which the endpoint group IP addresses belong. The CIDR blocks can be used to configure ACLs in terminals.
//
// @param request - ListEndpointGroupIpAddressCidrBlocksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEndpointGroupIpAddressCidrBlocksResponse
func (client *Client) ListEndpointGroupIpAddressCidrBlocksWithContext(ctx context.Context, request *ListEndpointGroupIpAddressCidrBlocksRequest, runtime *dara.RuntimeOptions) (_result *ListEndpointGroupIpAddressCidrBlocksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.EndpointGroupRegion) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEndpointGroupIpAddressCidrBlocks"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEndpointGroupIpAddressCidrBlocksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of endpoint groups.
//
// @param request - ListEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEndpointGroupsResponse
func (client *Client) ListEndpointGroupsWithContext(ctx context.Context, request *ListEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListEndpointGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AccessLogSwitch) {
		query["AccessLogSwitch"] = request.AccessLogSwitch
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointGroupType) {
		query["EndpointGroupType"] = request.EndpointGroupType
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries forwarding rules.
//
// Description:
//
// >  This operation is used to query only custom forwarding rules, not the default forwarding rule.
//
// @param request - ListForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListForwardingRulesResponse
func (client *Client) ListForwardingRulesWithContext(ctx context.Context, request *ListForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *ListForwardingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForwardingRuleId) {
		query["ForwardingRuleId"] = request.ForwardingRuleId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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
		Action:      dara.String("ListForwardingRules"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListForwardingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries acceleration regions.
//
// @param request - ListIpSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpSetsResponse
func (client *Client) ListIpSetsWithContext(ctx context.Context, request *ListIpSetsRequest, runtime *dara.RuntimeOptions) (_result *ListIpSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpSets"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the line types of elastic IP addresses (EIPs) that are supported in an acceleration region.
//
// @param request - ListIspTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIspTypesResponse
func (client *Client) ListIspTypesWithContext(ctx context.Context, request *ListIspTypesRequest, runtime *dara.RuntimeOptions) (_result *ListIspTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AcceleratorType) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !dara.IsNil(request.BusinessRegionId) {
		query["BusinessRegionId"] = request.BusinessRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIspTypes"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIspTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates associated with a listener.
//
// @param request - ListListenerCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenerCertificatesResponse
func (client *Client) ListListenerCertificatesWithContext(ctx context.Context, request *ListListenerCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
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

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListenerCertificates"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the listeners of a Global Accelerator (GA) instance.
//
// Description:
//
// This operation is used to query information about the listeners of a GA instance, including the status of each listener, the timestamp that indicates when each listener was created, and the listener ports.
//
// @param request - ListListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenersResponse
func (client *Client) ListListenersWithContext(ctx context.Context, request *ListListenersRequest, runtime *dara.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListeners"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the secondary IP addresses that are associated with a CNAME.
//
// @param request - ListSpareIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpareIpsResponse
func (client *Client) ListSpareIpsWithContext(ctx context.Context, request *ListSpareIpsRequest, runtime *dara.RuntimeOptions) (_result *ListSpareIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSpareIps"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSpareIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the TLS security policies that are supported by HTTPS listeners.
//
// Description:
//
// You can select a TLS security policy when you create an HTTPS listener. This API operation is used to query the TLS security policies that are supported by HTTPS listeners.
//
// @param request - ListSystemSecurityPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemSecurityPoliciesResponse
func (client *Client) ListSystemSecurityPoliciesWithContext(ctx context.Context, request *ListSystemSecurityPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListSystemSecurityPoliciesResponse, _err error) {
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
		Action:      dara.String("ListSystemSecurityPolicies"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to Global Accelerator (GA) resources.
//
// Description:
//
//	  You must specify **ResourceId*	- or **Tag*	- in the request to specify the object that you want to query.********
//
//		- **Tag*	- is a resource tag that consists of a key-value pair (Key and Value). If you specify only **Key**, all tag values that are associated with the specified tag key are returned. If you specify only **Value**, an error message is returned.
//
//		- If you specify **Tag*	- and **ResourceId*	- to filter tags, **ResourceId*	- must match all specified key-value pairs.
//
//		- If you specify multiple key-value pairs, resources that contain the key-value pairs are returned.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Version:     dara.String("2019-11-20"),
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
// Activates the pay-as-you-go Global Accelerator (GA) service. If you want to use pay-as-you-go GA instances, you must activate the pay-as-you-go GA service first.
//
// @param request - OpenAcceleratorServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAcceleratorServiceResponse
func (client *Client) OpenAcceleratorServiceWithContext(ctx context.Context, request *OpenAcceleratorServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenAcceleratorServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAcceleratorService"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAcceleratorServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Inquire about the approval status of cross-border permissions for an Alibaba Cloud account (main account).
//
// @param request - QueryCrossBorderApprovalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCrossBorderApprovalStatusResponse
func (client *Client) QueryCrossBorderApprovalStatusWithContext(ctx context.Context, request *QueryCrossBorderApprovalStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryCrossBorderApprovalStatusResponse, _err error) {
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
		Action:      dara.String("QueryCrossBorderApprovalStatus"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCrossBorderApprovalStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes IP entries from an access control list (ACL).
//
// Description:
//
//	  **RemoveEntriesFromAcl*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetAcl](https://help.aliyun.com/document_detail/258292.html) or [ListAcls](https://help.aliyun.com/document_detail/258291.html) operation to query the status of the ACL from which you want to delete IP entries.
//
//	    	- If the ACL is in the **configuring*	- state, it indicates that the IP entries are being deleted. In this case, you can perform only query operations.
//
//	    	- If the ACL is in the **active*	- state, it indicates that the IP entries are deleted.
//
//		- The **RemoveEntriesFromAcl*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - RemoveEntriesFromAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEntriesFromAclResponse
func (client *Client) RemoveEntriesFromAclWithContext(ctx context.Context, request *RemoveEntriesFromAclRequest, runtime *dara.RuntimeOptions) (_result *RemoveEntriesFromAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntries) {
		query["AclEntries"] = request.AclEntries
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEntriesFromAcl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the bandwidth plans of Global Accelerator (GA) instances.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - The GA instance continues to forward network traffic.
//
//   - **ReplaceBandwidthPackage*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) or [ListAccelerators](https://help.aliyun.com/document_detail/153236.html) operation to query the status of the GA instance.
//
//   - If the GA instance is in the **configuring*	- state, it indicates that the associated bandwidth plan is being replaced. In this case, you can perform only query operations.
//
//   - If the GA instance is in the **active*	- state, it indicates that the associated bandwidth plan is replaced.
//
//   - The **ReplaceBandwidthPackage*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - ReplaceBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceBandwidthPackageResponse
func (client *Client) ReplaceBandwidthPackageWithContext(ctx context.Context, request *ReplaceBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *ReplaceBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetBandwidthPackageId) {
		query["TargetBandwidthPackageId"] = request.TargetBandwidthPackageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceBandwidthPackage"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to Global Accelerator (GA) resources.
//
// Description:
//
// You can add up to 20 tags to each GA resource. When you call this operation, Alibaba Cloud first checks the number of existing tags attached to the resource. If the quota is reached, an error message is returned.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-11-20"),
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
// Removes tags from Global Accelerator (GA) resources.
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Version:     dara.String("2019-11-20"),
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
// Modifies a Global Accelerator (GA) instance.
//
// Description:
//
//	  **UpdateAccelerator*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of a GA instance.
//
//	    	- If the GA instance is in the **configuring*	- state, the GA instance is being modified. In this case, you can perform only query operations.
//
//	    	- If the GA instance is in the **active*	- state, the GA instance is modified.
//
//		- The **UpdateAccelerator*	- operation cannot be repeatedly called for the same GA instance within a specific period of time.
//
// @param request - UpdateAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorResponse
func (client *Client) UpdateAcceleratorWithContext(ctx context.Context, request *UpdateAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
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

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the auto-renewal settings of a Global Accelerator (GA) instance.
//
// Description:
//
// You cannot repeatedly call the **UpdateAcceleratorAutoRenewAttribute*	- operation for the same GA instance within a specific period of time.
//
// @param request - UpdateAcceleratorAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorAutoRenewAttributeResponse
func (client *Client) UpdateAcceleratorAutoRenewAttributeWithContext(ctx context.Context, request *UpdateAcceleratorAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAcceleratorAutoRenewAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms the updated specification of a Global Accelerator (GA) instance.
//
// Description:
//
// After you modify the specification of a GA instance, you must confirm the modification. The **UpdateAcceleratorConfirm*	- operation is used to confirm the specification of a GA instance that is modified. When you call this operation to confirm the specification of a GA instance, take note of the following items:
//
//   - **UpdateAcceleratorConfirm*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of a GA instance.
//
//   - If the GA instance is in the **configuring*	- state, it indicates that the specification of the instance is being modified. In this case, you can perform only query operations.
//
//   - If the GA instance is in the **active*	- state, it indicates that the specification of the instance is modified.
//
//   - The **UpdateAcceleratorConfirm*	- operation cannot be called repeatedly for the same GA instance within a specific period of time.
//
// @param request - UpdateAcceleratorConfirmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorConfirmResponse
func (client *Client) UpdateAcceleratorConfirmWithContext(ctx context.Context, request *UpdateAcceleratorConfirmRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorConfirmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAcceleratorConfirm"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAcceleratorConfirmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the type of transmission network for a Global Accelerator (GA) instance.
//
// Description:
//
// You can call this operation to change the type of transmission network for a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**. Before you call this operation, make sure that the following requirements are met:
//
//   - Cloud Data Transfer (CDT) is activated. When you call the [CreateAccelerator](https://help.aliyun.com/document_detail/206786.html) operation and set **BandwidthBillingType*	- to **CDT*	- to create a **standard*	- GA instance whose bandwidth metering method is **pay-by-data-transfer**, CDT is automatically activated. The data transfer fees are managed by CDT.
//
//   - If you want to set **CrossBorderMode*	- to **private**, which specifies cross-border Express Connect circuit as the type of transmission network, make sure that real-name verification is complete for your enterprise account. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/52595.html).
//
// @param request - UpdateAcceleratorCrossBorderModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorCrossBorderModeResponse
func (client *Client) UpdateAcceleratorCrossBorderModeWithContext(ctx context.Context, request *UpdateAcceleratorCrossBorderModeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorCrossBorderModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossBorderMode) {
		query["CrossBorderMode"] = request.CrossBorderMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAcceleratorCrossBorderMode"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAcceleratorCrossBorderModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables cross-border data transmission for a Global Accelerator (GA) instance.
//
// Description:
//
// You can call this operation to enable or disable cross-border data transmission for basic or standard GA instances that use Cloud Data Transfer (CDT) to bill data transfers.
//
// @param request - UpdateAcceleratorCrossBorderStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAcceleratorCrossBorderStatusResponse
func (client *Client) UpdateAcceleratorCrossBorderStatusWithContext(ctx context.Context, request *UpdateAcceleratorCrossBorderStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAcceleratorCrossBorderStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossBorderStatus) {
		query["CrossBorderStatus"] = request.CrossBorderStatus
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAcceleratorCrossBorderStatus"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAcceleratorCrossBorderStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an access control list (ACL) of a Global Accelerator (GA) instance.
//
// @param request - UpdateAclAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAclAttributeResponse
func (client *Client) UpdateAclAttributeWithContext(ctx context.Context, request *UpdateAclAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAclAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAclAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces an expired additional certificate for an HTTPS listener.
//
// Description:
//
// The UpdateAdditionalCertificateWithListener operation is used to replace an additional certificate. This operation is suitable for scenarios in which the original certificate expires and needs to be replaced with a new certificate and the associated domain name remains unchanged.
//
//   - **UpdateAdditionalCertificateWithListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListListenerCertificates](https://help.aliyun.com/document_detail/307743.html) to query the status of the task:
//
//   - If the certificate to be replaced is in the **updating*	- state, the certificate is being replaced. In this case, you can only query the certificate.
//
//   - If the new certificate is in the **active*	- state, the new certificate is in effect.
//
//   - You cannot repeatedly call the **UpdateAdditionalCertificateWithListener*	- operation for the same GA instance within a specific period of time.
//
// @param request - UpdateAdditionalCertificateWithListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdditionalCertificateWithListenerResponse
func (client *Client) UpdateAdditionalCertificateWithListenerWithContext(ctx context.Context, request *UpdateAdditionalCertificateWithListenerRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdditionalCertificateWithListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdditionalCertificateWithListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an origin probing task.
//
// Description:
//
// *UpdateApplicationMonitor*	- is an asynchronous operation. After you send a request, the system returns a request ID, but this operation is still being performed in the system background. You can call the [DescribeApplicationMonitor](https://help.aliyun.com/document_detail/408463.html) or [ListApplicationMonitor](https://help.aliyun.com/document_detail/408462.html) operation to check whether the configurations of an origin probing task are modified.
//
//   - If the values of modified parameters remain unchanged, it indicates that the origin probing task is being modified. In this case, you can perform only query operations.
//
//   - If the values of modified parameters change, it indicates that the origin probing task is modified.
//
// @param request - UpdateApplicationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationMonitorResponse
func (client *Client) UpdateApplicationMonitorWithContext(ctx context.Context, request *UpdateApplicationMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DetectEnable) {
		query["DetectEnable"] = request.DetectEnable
	}

	if !dara.IsNil(request.DetectThreshold) {
		query["DetectThreshold"] = request.DetectThreshold
	}

	if !dara.IsNil(request.DetectTimes) {
		query["DetectTimes"] = request.DetectTimes
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationMonitor"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the auto-renewal status of a bandwidth plan.
//
// Description:
//
// You cannot repeatedly call the **UpdateBandwidthPackagaAutoRenewAttribute*	- operation to modify the auto-renewal settings of a bandwidth plan.
//
// @param request - UpdateBandwidthPackagaAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBandwidthPackagaAutoRenewAttributeResponse
func (client *Client) UpdateBandwidthPackagaAutoRenewAttributeWithContext(ctx context.Context, request *UpdateBandwidthPackagaAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateBandwidthPackagaAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBandwidthPackagaAutoRenewAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBandwidthPackagaAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a bandwidth plan.
//
// Description:
//
// Take note of the following items:
//
//   - **UpdateBandwidthPackage*	- is a synchronous operation when you call the operation to modify the configuration excluding the bandwidth value of a bandwidth plan. The new configuration immediately takes effect after the operation is performed.
//
//   - **UpdateBandwidthPackage*	- is an asynchronous operation when you call the operation to modify the configuration including the bandwidth value of a bandwidth plan that is not associated with a Global Accelerator (GA) instance. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeBandwidthPackage](https://help.aliyun.com/document_detail/153241.html) operation to query the status of the task.
//
//   - If the parameter values of the bandwidth plan remain unchanged, the bandwidth plan is being modified. In this case, you can perform only query operations.
//
//   - If the parameter values of the bandwidth plan are changed, the bandwidth plan is modified.
//
//   - **UpdateBandwidthPackage*	- is an asynchronous operation when you call the operation to modify the configuration including the bandwidth value of a bandwidth plan that is associated with a GA instance. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeAccelerator](https://help.aliyun.com/document_detail/153235.html) operation to query the status of the task.
//
//   - If the GA instance is in the **configuring*	- state, the bandwidth plan is being modified. In this case, you can perform only query operations.
//
//   - If the GA instance is in the **active*	- state, the bandwidth plan is modified.
//
//   - You cannot repeatedly call the **UpdateBandwidthPackage*	- operation for the same bandwidth plan within a specific period of time.
//
// @param request - UpdateBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBandwidthPackageResponse
func (client *Client) UpdateBandwidthPackageWithContext(ctx context.Context, request *UpdateBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *UpdateBandwidthPackageResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthPackageId) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !dara.IsNil(request.BandwidthType) {
		query["BandwidthType"] = request.BandwidthType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBandwidthPackage"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a basic Global Accelerator (GA) instance.
//
// @param request - UpdateBasicAcceleratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicAcceleratorResponse
func (client *Client) UpdateBasicAcceleratorWithContext(ctx context.Context, request *UpdateBasicAcceleratorRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicAcceleratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBasicAccelerator"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBasicAcceleratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of an endpoint that is associated with a basic Global Accelerator (GA) instance.
//
// @param request - UpdateBasicEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicEndpointResponse
func (client *Client) UpdateBasicEndpointWithContext(ctx context.Context, request *UpdateBasicEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
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
		Action:      dara.String("UpdateBasicEndpoint"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBasicEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an endpoint group that is associated with a basic Global Accelerator (GA) instance.
//
// Description:
//
//	  **UpdateBasicEndpointGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. The system modifies the configurations of an endpoint group that is associated with a basic GA instance by deleting the endpoint group and creating a new endpoint group. You can call the [GetBasicAccelerator](https://help.aliyun.com/document_detail/353188.html) operation to query the status of the task.
//
//	    	- If the basic GA instance is in the **configuring*	- state, the configurations of the endpoint group are being modified. In this case, you can perform only query operations.
//
//	    	- If the basic GA instance is in the **active*	- state, the configurations of the endpoint group are modified.
//
//		- The **UpdateBasicEndpointGroup*	- operation cannot be repeatedly called for the same basic GA instance within a specific period of time.
//
// @param request - UpdateBasicEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicEndpointGroupResponse
func (client *Client) UpdateBasicEndpointGroupWithContext(ctx context.Context, request *UpdateBasicEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicEndpointGroupResponse, _err error) {
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

	if !dara.IsNil(request.EndpointAddress) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointSubAddress) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
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
		Action:      dara.String("UpdateBasicEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBasicEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the bandwidth of an acceleration region of a basic Global Accelerator (GA) instance.
//
// Description:
//
// Before you call this operation, take note of the following limits:
//
//   - You can modify the bandwidth of an acceleration region of a basic GA instance only if the bandwidth metering method of the basic GA instance is **pay-by-data-transfer**.
//
//   - **UpdateBasicIpSet*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetBasicIpSet](https://help.aliyun.com/document_detail/362987.html) operation to query the status of the task.
//
//   - If the acceleration region is in the **updating*	- state, it indicates that the bandwidth of the acceleration region is being modified. In this case, you can perform only query operations.
//
//   - If the acceleration region is in the **active*	- state, it indicates that the bandwidth of the acceleration region is modified.
//
//   - You cannot repeatedly call the **UpdateBasicIpSet*	- operation for the same basic GA instance within a specific period of time.
//
// @param request - UpdateBasicIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBasicIpSetResponse
func (client *Client) UpdateBasicIpSetWithContext(ctx context.Context, request *UpdateBasicIpSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateBasicIpSetResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBasicIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBasicIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of an endpoint group that is associated with a custom routing listener.
//
// @param request - UpdateCustomRoutingEndpointGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoutingEndpointGroupAttributeResponse
func (client *Client) UpdateCustomRoutingEndpointGroupAttributeWithContext(ctx context.Context, request *UpdateCustomRoutingEndpointGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupAttributeResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
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
		Action:      dara.String("UpdateCustomRoutingEndpointGroupAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the mapping configurations of an endpoint group that is associated with a custom routing listener.
//
// Description:
//
//	  **UpdateCustomRoutingEndpointGroupDestinations*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of an endpoint group associated with a custom routing listener to check whether the mapping configurations of the endpoint group are modified.
//
//	    	- If the endpoint group is in the **updating*	- state, the mapping configurations of the endpoint group are being modified. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, the mapping configurations of the endpoint group are modified.
//
//		- The **UpdateCustomRoutingEndpointGroupDestinations*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - UpdateCustomRoutingEndpointGroupDestinationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoutingEndpointGroupDestinationsResponse
func (client *Client) UpdateCustomRoutingEndpointGroupDestinationsWithContext(ctx context.Context, request *UpdateCustomRoutingEndpointGroupDestinationsRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
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

	if !dara.IsNil(request.DestinationConfigurations) {
		query["DestinationConfigurations"] = request.DestinationConfigurations
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomRoutingEndpointGroupDestinations"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the traffic policies for an endpoint that is associated with a custom routing listener.
//
// Description:
//
//	  **UpdateCustomRoutingEndpointTrafficPolicies*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the status of the task.
//
//	    	- If the endpoint group is in the **updating*	- state, traffic policies are being modified for endpoints in the endpoint group. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, traffic policies are modified for endpoints in the endpoint group.
//
//		- The **UpdateCustomRoutingEndpointTrafficPolicies*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - UpdateCustomRoutingEndpointTrafficPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoutingEndpointTrafficPoliciesResponse
func (client *Client) UpdateCustomRoutingEndpointTrafficPoliciesWithContext(ctx context.Context, request *UpdateCustomRoutingEndpointTrafficPoliciesRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.PolicyConfigurations) {
		query["PolicyConfigurations"] = request.PolicyConfigurations
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomRoutingEndpointTrafficPolicies"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the endpoints of a custom routing listener.
//
// Description:
//
// ## Description
//
//   - **UpdateCustomRoutingEndpoints*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeCustomRoutingEndpointGroup](https://help.aliyun.com/document_detail/449373.html) operation to query the state of the endpoint groups associated with a custom routing listener to check whether the endpoints in the endpoint groups are modified.
//
//   - If an endpoint group is in the **updating*	- state, the endpoints in the endpoint group are being modified. In this case, you can perform only query operations.
//
//   - If an endpoint group is in the **active*	- state, the endpoints in the endpoint group are modified.
//
//   - The **UpdateCustomRoutingEndpoints*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - UpdateCustomRoutingEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomRoutingEndpointsResponse
func (client *Client) UpdateCustomRoutingEndpointsWithContext(ctx context.Context, request *UpdateCustomRoutingEndpointsRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomRoutingEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.EndpointConfigurations) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomRoutingEndpoints"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an accelerated domain name.
//
// Description:
//
// You can call this operation to modify an accelerated domain name. If the new accelerated domain name is hosted in the Chinese mainland, you must obtain an Internet content provider (ICP) number for the domain name.
//
// You cannot call the **UpdateDomain*	- operation again by using the same Alibaba Cloud account before the previous request is completed.
//
// @param request - UpdateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithContext(ctx context.Context, request *UpdateDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetDomain) {
		query["TargetDomain"] = request.TargetDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomain"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the ICP filing status of an accelerated domain name.
//
// Description:
//
// You can call this operation to query and update the ICP filing status of an accelerated domain name.
//
// The **UpdateDomainState*	- operation holds an exclusive lock on the GA instance. While the operation is in progress, you cannot call the same operation with the same Alibaba Cloud account.
//
// @param request - UpdateDomainStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainStateResponse
func (client *Client) UpdateDomainStateWithContext(ctx context.Context, request *UpdateDomainStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainState"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an endpoint group.
//
// Description:
//
//	  **UpdateEndpointGroup*	- is an asynchronous operation. After you send a request, the system returns a request ID, but the operation is still being performed in the system background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) operation to query the state of an endpoint group.
//
//	    	- If the endpoint group is in the **updating*	- state, it indicates that the configurations of the endpoint group are being modified. In this case, you can perform only query operations.
//
//	    	- If the endpoint group is in the **active*	- state, it indicates that the configurations of the endpoint group are modified.
//
//		- The **UpdateEndpointGroup*	- operation cannot be repeatedly called for the same Global Accelerator (GA) instance within a specific period of time.
//
// @param request - UpdateEndpointGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointGroupResponse
func (client *Client) UpdateEndpointGroupWithContext(ctx context.Context, request *UpdateEndpointGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupResponse, _err error) {
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

	if !dara.IsNil(request.EndpointConfigurations) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.EndpointGroupRegion) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !dara.IsNil(request.EndpointIpVersion) {
		query["EndpointIpVersion"] = request.EndpointIpVersion
	}

	if !dara.IsNil(request.EndpointProtocolVersion) {
		query["EndpointProtocolVersion"] = request.EndpointProtocolVersion
	}

	if !dara.IsNil(request.EndpointRequestProtocol) {
		query["EndpointRequestProtocol"] = request.EndpointRequestProtocol
	}

	if !dara.IsNil(request.HealthCheckEnabled) {
		query["HealthCheckEnabled"] = request.HealthCheckEnabled
	}

	if !dara.IsNil(request.HealthCheckHost) {
		query["HealthCheckHost"] = request.HealthCheckHost
	}

	if !dara.IsNil(request.HealthCheckIntervalSeconds) {
		query["HealthCheckIntervalSeconds"] = request.HealthCheckIntervalSeconds
	}

	if !dara.IsNil(request.HealthCheckPath) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !dara.IsNil(request.HealthCheckPort) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !dara.IsNil(request.HealthCheckProtocol) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortOverrides) {
		query["PortOverrides"] = request.PortOverrides
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ThresholdCount) {
		query["ThresholdCount"] = request.ThresholdCount
	}

	if !dara.IsNil(request.TrafficPercentage) {
		query["TrafficPercentage"] = request.TrafficPercentage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEndpointGroup"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEndpointGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of an endpoint group.
//
// @param request - UpdateEndpointGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointGroupAttributeResponse
func (client *Client) UpdateEndpointGroupAttributeWithContext(ctx context.Context, request *UpdateEndpointGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupAttributeResponse, _err error) {
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

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
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
		Action:      dara.String("UpdateEndpointGroupAttribute"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEndpointGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the endpoint groups that are associated with a listener.
//
// Description:
//
// ### Description
//
//   - **UpdateEndpointGroups*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeEndpointGroup](https://help.aliyun.com/document_detail/153260.html) or [ListEndpointGroups](https://help.aliyun.com/document_detail/153261.html) operation to query the status of an endpoint group.
//
//   - If the endpoint group is in the **updating*	- state, it indicates that the configuration of the endpoint group is being modified. In this case, you can perform only query operations.
//
//   - If the endpoint group is in the **active*	- state, it indicates that the configuration of the endpoint group is modified.
//
//   - The **UpdateEndpointGroups*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - UpdateEndpointGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointGroupsResponse
func (client *Client) UpdateEndpointGroupsWithContext(ctx context.Context, request *UpdateEndpointGroupsRequest, runtime *dara.RuntimeOptions) (_result *UpdateEndpointGroupsResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointGroupConfigurations) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEndpointGroups"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEndpointGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a forwarding rule.
//
// Description:
//
//	  **UpdateForwardingRules*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListForwardingRules](https://help.aliyun.com/document_detail/205817.html) operation to query the status of a forwarding rule.
//
//	    	- If the forwarding rule is in the **configuring*	- state, it indicates that the forwarding rule is being modified. In this case, you can perform only query operations.
//
//	    	- If the forwarding rule is in the **active*	- state, it indicates that the forwarding rule is modified.
//
//		- The **UpdateForwardingRules*	- operation holds an exclusive lock on the Global Accelerator (GA) instance. While the operation is in progress, you cannot call the same operation in the same Alibaba Cloud account.
//
// @param request - UpdateForwardingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateForwardingRulesResponse
func (client *Client) UpdateForwardingRulesWithContext(ctx context.Context, request *UpdateForwardingRulesRequest, runtime *dara.RuntimeOptions) (_result *UpdateForwardingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ForwardingRules) {
		query["ForwardingRules"] = request.ForwardingRules
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateForwardingRules"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateForwardingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an acceleration region in an acceleration area for a Global Accelerator (GA) instance.
//
// Description:
//
//	  **UpdateIpSet*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeIpSet](https://help.aliyun.com/document_detail/153246.html) operation to query the status of an acceleration region.
//
//	    	- If the acceleration region is in the **updating*	- state, it indicates that the acceleration region is being modified. In this case, you can continue to perform query operations on the acceleration regions.
//
//	    	- If the acceleration region is in the **active*	- state, it indicates that the acceleration region is modified.
//
//		- You cannot call the **UpdateIpSet*	- operation again on the same GA instance before the previous operation is complete.
//
// @param request - UpdateIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpSetResponse
func (client *Client) UpdateIpSetWithContext(ctx context.Context, request *UpdateIpSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpSetResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpSetId) {
		query["IpSetId"] = request.IpSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIpSet"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of acceleration regions in an acceleration area for a Global Accelerator (GA) instance.
//
// Description:
//
//	  **UpdateIpSets*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [DescribeIpSet](https://help.aliyun.com/document_detail/153246.html) operation to query the status of the task.
//
//	    	- If an acceleration region is in the **updating*	- state, the acceleration region is being modified. In this case, you can perform only query operations.
//
//	    	- If an acceleration region is in the **active*	- state, the acceleration region is modified.
//
//		- You cannot repeatedly call the **UpdateIpSets*	- operation for the same GA instance within a specific period of time.
//
// @param request - UpdateIpSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpSetsResponse
func (client *Client) UpdateIpSetsWithContext(ctx context.Context, request *UpdateIpSetsRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpSets) {
		query["IpSets"] = request.IpSets
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIpSets"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a listener for a Global Accelerator (GA) instance.
//
// Description:
//
// This operation can be called to modify the configurations such as the protocol and ports of a listener to meet your business requirements.
//
// When you call this operation, take note of the following items:
//
//   - **UpdateListener*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [DescribeListener](https://help.aliyun.com/document_detail/153254.html) operation to query the status of a listener.
//
//   - If the listener is in the **updating*	- state, it indicates that its configurations are being modified. In this case, you can perform only query operations.
//
//   - If the listener is in the **active*	- state, it indicates that its configurations are modified.
//
//   - The **UpdateListener*	- operation cannot be repeatedly called to modify listener configurations for the same GA instance within a specific period of time.
//
// @param request - UpdateListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListenerResponse
func (client *Client) UpdateListenerWithContext(ctx context.Context, request *UpdateListenerRequest, runtime *dara.RuntimeOptions) (_result *UpdateListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendPorts) {
		query["BackendPorts"] = request.BackendPorts
	}

	if !dara.IsNil(request.Certificates) {
		query["Certificates"] = request.Certificates
	}

	if !dara.IsNil(request.ClientAffinity) {
		query["ClientAffinity"] = request.ClientAffinity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HttpVersion) {
		query["HttpVersion"] = request.HttpVersion
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortRanges) {
		query["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ProxyProtocol) {
		query["ProxyProtocol"] = request.ProxyProtocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.XForwardedForConfig) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateListener"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改SLS日志配置
//
// @param request - UpdateLogStoreConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLogStoreConfigResponse
func (client *Client) UpdateLogStoreConfigWithContext(ctx context.Context, request *UpdateLogStoreConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateLogStoreConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorId) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !dara.IsNil(request.AccessLogRecordCustomizedHeaderList) {
		query["AccessLogRecordCustomizedHeaderList"] = request.AccessLogRecordCustomizedHeaderList
	}

	if !dara.IsNil(request.AccessLogRecordCustomizedHeadersEnabled) {
		query["AccessLogRecordCustomizedHeadersEnabled"] = request.AccessLogRecordCustomizedHeadersEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointGroupId) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlsLogStoreName) {
		query["SlsLogStoreName"] = request.SlsLogStoreName
	}

	if !dara.IsNil(request.SlsProjectName) {
		query["SlsProjectName"] = request.SlsProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLogStoreConfig"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLogStoreConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the control mode of a resource from managed mode to unmanaged mode.
//
// Description:
//
//	  This operation is applicable only to **managed*	- Global Accelerator (GA) instances.
//
//		- After you change the control mode of a GA instance from managed mode to unmanaged mode, you cannot change the mode of the instance to managed mode.
//
//		- After you change the control mode of a GA instance from managed mode to unmanaged mode, you can obtain all operation permissions on the instance.
//
//	  <warning>If you change or delete a configuration of a GA instance whose control mode is changed from managed mode to unmanaged mode, the cloud services that depend on the instance may not work as expected. Proceed with caution.
//
// @param request - UpdateServiceManagedControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceManagedControlResponse
func (client *Client) UpdateServiceManagedControlWithContext(ctx context.Context, request *UpdateServiceManagedControlRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceManagedControlResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceManaged) {
		query["ServiceManaged"] = request.ServiceManaged
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceManagedControl"),
		Version:     dara.String("2019-11-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceManagedControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
